package fetcher

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/time/rate"
)

// OptimizedFetcher uses worker pool pattern for controlled concurrency
type OptimizedFetcher struct {
	config      Config
	httpClient  *http.Client
	jobs        chan string    // Buffered job queue (worker pool pattern)
	visited     sync.Map       // Concurrent-safe URL deduplication
	activeJobs  sync.WaitGroup // Track active jobs to shut down gracefully
	resultsChan chan string    // Results writer channel
	llmEntries  []LLMTxtEntry
	llmMutex    sync.Mutex
	pageCount   int32
	errorCount  int32
	ctx         context.Context
	cancel      context.CancelFunc
	rateLimiter *rate.Limiter // Rate limiting to prevent DDoS
}

// RunOptimized executes documentation fetching with maximum concurrency
func RunOptimized(config Config) error {
	if err := ValidateConfig(&config); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	log.Printf("🚀 Starting HIGH-PERFORMANCE documentation fetch from: %s", config.BaseURL)
	log.Printf("   Workers: %d | Max Depth: %d | Rate Limit: %d req/sec", config.Workers, config.MaxDepth, config.Workers*2)

	// Create rate limiter (prevent DDoS, be civilized)
	// Allow burst of workers*2, then sustain at workers per second
	rateLimiter := rate.NewLimiter(rate.Limit(config.Workers), config.Workers*2)

	// Buffered channels for worker pool pattern
	jobs := make(chan string, config.Workers*10)         // Job queue
	resultsChan := make(chan string, config.Workers*100) // Results buffer

	fetcher := &OptimizedFetcher{
		config:      config,
		jobs:        jobs,
		resultsChan: resultsChan,
		httpClient:  createOptimizedHTTPClient(config.Workers),
		rateLimiter: rateLimiter,
		ctx:         context.Background(),
	}

	fetcher.ctx, fetcher.cancel = context.WithTimeout(context.Background(), 10*time.Minute)
	defer fetcher.cancel()

	startTime := time.Now()

	// Start result writer in background
	var writeWg sync.WaitGroup
	if !config.GenerateLLMTxt {
		writeWg.Add(1)
		go func() {
			defer writeWg.Done()
			if err := writeResultsOptimized(config.OutputPath, resultsChan); err != nil {
				log.Printf("❌ Error writing results: %v", err)
			}
		}()
	}

	// Start worker pool (controlled concurrency)
	var workerWg sync.WaitGroup
	for i := 0; i < config.Workers; i++ {
		workerWg.Add(1)
		go fetcher.worker(i, &workerWg, rateLimiter)
	}

	// Submit initial URL to job queue
	fetcher.submitPage(config.BaseURL, 0)

	// Graceful shutdown of job queue when all processing finishes
	go func() {
		fetcher.activeJobs.Wait()
		close(jobs)
	}()

	// Wait for all workers to complete
	workerWg.Wait()

	// Close results channel (no more results will be written)
	close(resultsChan)

	// Wait for results writer to finish
	writeWg.Wait()

	elapsed := time.Since(startTime)
	pagesFetched := atomic.LoadInt32(&fetcher.pageCount)
	errors := atomic.LoadInt32(&fetcher.errorCount)

	log.Printf("✅ Fetch completed!")
	log.Printf("   📊 Pages fetched: %d", pagesFetched)
	log.Printf("   ⏱️  Time elapsed: %v (%.2f pages/sec)", elapsed, float64(pagesFetched)/elapsed.Seconds())
	log.Printf("   ❌ Errors: %d", errors)
	log.Printf("💡 Tip: Use --concurrent <N> to adjust parallelism (current: %d)", config.Workers)

	// Generate LLM.txt if requested
	if config.GenerateLLMTxt {
		if len(fetcher.llmEntries) > 0 {
			if err := GenerateLLMTxt(fetcher.llmEntries, config.OutputPath); err != nil {
				log.Printf("⚠️  Warning: Failed to generate llm.txt: %v", err)
			} else {
				log.Printf("📝 LLM.txt generated: %s (%d entries)", config.OutputPath, len(fetcher.llmEntries))
			}
		} else {
			log.Printf("📝 LLM.txt generated: %s (0 entries)", config.OutputPath)
		}
	}

	return nil
}

// createOptimizedHTTPClient creates a high-performance HTTP client with connection pooling
func createOptimizedHTTPClient(workers int) *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        workers * 2,
			MaxIdleConnsPerHost: workers,
			IdleConnTimeout:     90 * time.Second,
			DisableCompression:  false,
			DisableKeepAlives:   false,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}

// worker processes URLs from job queue with rate limiting
func (f *OptimizedFetcher) worker(id int, wg *sync.WaitGroup, limiter *rate.Limiter) {
	defer wg.Done()

	for pageURL := range f.jobs {
		select {
		case <-f.ctx.Done():
			f.activeJobs.Done()
			return
		default:
			// Wait for rate limiter (be civilized, don't DDoS)
			if err := limiter.Wait(f.ctx); err != nil {
				log.Printf("⚠️  Rate limit error: %v", err)
				f.activeJobs.Done()
				continue
			}

			f.processURL(pageURL, 0)
			f.activeJobs.Done()
		}
	}
}

// submitPage adds a URL to job queue (with depth tracking and deduplication)
func (f *OptimizedFetcher) submitPage(pageURL string, depth int) {
	if depth > f.config.MaxDepth {
		return
	}

	// Deduplicate URLs (concurrent-safe)
	if _, loaded := f.visited.LoadOrStore(pageURL, true); loaded {
		return
	}

	// Submit to job queue via goroutine with context guard to avoid deadlocks
	f.activeJobs.Add(1)
	go func() {
		select {
		case f.jobs <- pageURL:
		case <-f.ctx.Done():
			f.activeJobs.Done()
		}
	}()
}

// processURL fetches and processes a single URL
func (f *OptimizedFetcher) processURL(pageURL string, depth int) {
	atomic.AddInt32(&f.pageCount, 1)

	startTime := time.Now()

	// Validate URL
	if err := isValidURL(pageURL); err != nil {
		atomic.AddInt32(&f.errorCount, 1)
		log.Printf("❌ Invalid URL %s: %v", pageURL, err)
		return
	}

	// Fetch the page
	resp, err := f.httpClient.Get(pageURL)
	if err != nil {
		atomic.AddInt32(&f.errorCount, 1)
		log.Printf("❌ Error fetching %s: %v", pageURL, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		atomic.AddInt32(&f.errorCount, 1)
		log.Printf("❌ Non-200 status %d for %s", resp.StatusCode, pageURL)
		return
	}

	// Parse HTML concurrently
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		atomic.AddInt32(&f.errorCount, 1)
		log.Printf("❌ Error parsing HTML for %s: %v", pageURL, err)
		return
	}

	// Extract links FIRST (always, even for pages with no content)
	links := ExtractAllLinksFromPage(doc, pageURL)
	navLinks := ExtractNavLinksFromPage(doc, pageURL)

	// Submit extracted links for crawling
	for _, link := range links {
		if !isCrawlableURL(link.URL) {
			continue
		}
		if depth < f.config.MaxDepth {
			f.submitPage(link.URL, depth+1)
		}
	}

	// Submit non-anchor nav links for crawling
	for _, link := range navLinks {
		if !isCrawlableURL(link.URL) || strings.HasPrefix(link.URL, "#") {
			continue
		}
		if depth < f.config.MaxDepth {
			f.submitPage(link.URL, depth+1)
		}
	}

	// Build hierarchical structure using SectionBuilder
	builder := NewSectionBuilder(pageURL)
	builder.BuildFromDocument(doc)

	// Render to markdown
	content := builder.GenerateMarkdown()

	// Send to results if content exists and not in llm-txt mode
	if strings.TrimSpace(content) != "" && !f.config.GenerateLLMTxt {
		title := builder.GetRoot().Title
		if title == "Documentation" || title == "" {
			title = doc.Find("title").Text()
			if title == "" {
				title = pageURL
			}
		}
		f.resultsChan <- fmt.Sprintf("%s\t%s\t%s", pageURL, title, content)
	} else if strings.TrimSpace(content) == "" {
		log.Printf("⚠️  No content found for %s", pageURL)
	}

	// Add to llm.txt entries if requested
	if f.config.GenerateLLMTxt {
		f.llmMutex.Lock()
		for _, link := range links {
			if !strings.Contains(link.URL, f.config.BaseURL) {
				continue
			}

			f.llmEntries = append(f.llmEntries, LLMTxtEntry{
				Type:        ClassifyPage(link.URL, link.Text),
				Title:       link.Text,
				URL:         link.URL,
				Description: link.Context,
			})
		}

		// Add nav/sidebar links with anchor-target content for better context
		for _, link := range navLinks {
			if !strings.Contains(link.URL, f.config.BaseURL) {
				continue
			}

			context := link.Context
			hashIdx := strings.Index(link.URL, "#")
			if hashIdx >= 0 {
				fragment := link.URL[hashIdx+1:]
				if fragment != "" {
					if anchorContent := extractAnchorContent(doc, fragment); anchorContent != "" {
						context = anchorContent
					}
				}
			}

			f.llmEntries = append(f.llmEntries, LLMTxtEntry{
				Type:        ClassifyPage(link.URL, link.Text),
				Title:       link.Text,
				URL:         link.URL,
				Description: context,
			})
		}
		f.llmMutex.Unlock()
	}

	elapsed := time.Since(startTime)
	log.Printf("✅ Fetched %s (%.2fs)", pageURL, elapsed.Seconds())
}

// findLinkTextForURL finds the anchor text for a given URL in the document
func findLinkTextForURL(doc *goquery.Document, targetURL string) string {
	var linkText string

	doc.Find("a[href]").Each(func(i int, a *goquery.Selection) {
		href, exists := a.Attr("href")
		if exists && strings.Contains(href, targetURL) || strings.Contains(targetURL, href) {
			text := strings.TrimSpace(a.Text())
			if text != "" && linkText == "" {
				linkText = text
			}
		}
	})

	return linkText
}

// writeResultsOptimized writes results using hierarchical section building
func writeResultsOptimized(outputPath string, resultsChan <-chan string) error {
	// Collect all results first
	type pageResult struct {
		url     string
		title   string
		content string
	}

	var pages []pageResult

	for result := range resultsChan {
		if strings.TrimSpace(result) == "" {
			continue
		}

		// Format: "URL\tTitle\tContent"
		parts := strings.SplitN(result, "\t", 3)
		if len(parts) < 3 {
			continue
		}

		pages = append(pages, pageResult{
			url:     strings.TrimSpace(parts[0]),
			title:   strings.TrimSpace(parts[1]),
			content: strings.TrimSpace(parts[2]),
		})
	}

	// Build hierarchical structure from collected pages
	// For now, use simple approach - each page is a top-level section
	// TODO: Implement full DOM-based hierarchy across pages

	var output strings.Builder
	output.WriteString("# Documentation\n\n")
	output.WriteString("Generated by DocFetch - Intelligent documentation fetcher\n\n")
	output.WriteString("---\n\n")

	// Write table of contents
	output.WriteString("## Table of Contents\n\n")
	for _, page := range pages {
		slug := slugify(page.title)
		output.WriteString(fmt.Sprintf("- [%s](#%s)\n", page.title, slug))
	}
	output.WriteString("\n---\n\n")

	// Write each page
	for _, page := range pages {
		output.WriteString(fmt.Sprintf("## %s\n\n", page.title))
		output.WriteString(fmt.Sprintf("*Source: [%s](%s)*\n\n", page.url, page.url))
		output.WriteString(page.content)
		output.WriteString("\n\n---\n\n")
	}

	// Write to file
	return os.WriteFile(outputPath, []byte(output.String()), 0644)
}

// slugify converts text to URL-friendly slug
func slugify(text string) string {
	text = strings.ToLower(text)
	text = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(text, "-")
	text = strings.Trim(text, "-")
	return text
}

// isCrawlableURL returns true if the URL has an HTTP scheme
func isCrawlableURL(rawURL string) bool {
	return strings.HasPrefix(rawURL, "http://") || strings.HasPrefix(rawURL, "https://")
}

// extractAnchorContent finds the element targeted by a fragment ID and returns
// the nearest heading plus the first paragraph of content as a description.
func extractAnchorContent(doc *goquery.Document, fragment string) string {
	target := doc.Find("#" + fragment).First()
	if target.Length() == 0 {
		return ""
	}

	// Find the nearest preceding heading sibling
	var headingText string
	prev := target.PrevAll()
	if prev.Length() > 0 {
		prev.Find("h1, h2, h3, h4, h5, h6").Last().Each(func(i int, s *goquery.Selection) {
			if headingText == "" {
				headingText = strings.TrimSpace(s.Text())
			}
		})
	}

	// If no preceding sibling heading, look at parent's previous siblings
	if headingText == "" {
		target.Parent().PrevAll().Find("h1, h2, h3, h4, h5, h6").Last().Each(func(i int, s *goquery.Selection) {
			if headingText == "" {
				headingText = strings.TrimSpace(s.Text())
			}
		})
	}

	// If still no heading, use the target's own text as heading
	if headingText == "" {
		headingText = strings.TrimSpace(target.Text())
	}

	// Extract first meaningful text after the heading/target
	var firstPara string
	target.NextUntil("h1, h2, h3, h4, h5, h6").Each(func(i int, s *goquery.Selection) {
		if firstPara != "" {
			return
		}
		if s.Is("p") || s.Is("div") {
			text := strings.TrimSpace(s.Text())
			if text != "" {
				firstPara = text
			}
		}
	})

	if headingText != "" {
		return headingText
	}
	if firstPara != "" {
		return firstPara
	}
	return ""
}
