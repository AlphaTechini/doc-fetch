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
	jobs        chan string        // Buffered job queue (worker pool pattern)
	visited     sync.Map           // Concurrent-safe URL deduplication
	resultsChan chan string        // Results writer channel
	llmEntries  []LLMTxtEntry
	llmMutex    sync.Mutex
	pageCount   int32
	errorCount  int32
	ctx         context.Context
	cancel      context.CancelFunc
	rateLimiter *rate.Limiter      // Rate limiting to prevent DDoS
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
	jobs := make(chan string, config.Workers*10)  // Job queue
	resultsChan := make(chan string, config.Workers*100)  // Results buffer

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
	writeWg.Add(1)
	go func() {
		defer writeWg.Done()
		if err := writeResultsOptimized(config.OutputPath, resultsChan); err != nil {
			log.Printf("❌ Error writing results: %v", err)
		}
	}()

	// Start worker pool (controlled concurrency)
	var workerWg sync.WaitGroup
	for i := 0; i < config.Workers; i++ {
		workerWg.Add(1)
		go fetcher.worker(i, &workerWg, rateLimiter)
	}

	// Submit initial URL to job queue
	fetcher.submitPage(config.BaseURL, 0)

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
	if config.GenerateLLMTxt && len(fetcher.llmEntries) > 0 {
		llmTxtPath := strings.TrimSuffix(config.OutputPath, ".md") + ".llm.txt"
		if err := GenerateLLMTxt(fetcher.llmEntries, llmTxtPath); err != nil {
			log.Printf("⚠️  Warning: Failed to generate llm.txt: %v", err)
		} else {
			log.Printf("📝 LLM.txt generated: %s (%d entries)", llmTxtPath, len(fetcher.llmEntries))
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
			return
		default:
			// Wait for rate limiter (be civilized, don't DDoS)
			if err := limiter.Wait(f.ctx); err != nil {
				log.Printf("⚠️  Rate limit error: %v", err)
				continue
			}
			
			f.processURL(pageURL, 0)
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

	// Submit to job queue (buffered, so rarely blocks)
	f.jobs <- pageURL
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

	// Build hierarchical structure using SectionBuilder (NEW method)
	builder := NewSectionBuilder(pageURL)
	builder.BuildFromDocument(doc)
	
	// Render to markdown
	content := builder.GenerateMarkdown()
	if strings.TrimSpace(content) == "" {
		atomic.AddInt32(&f.errorCount, 1)
		log.Printf("⚠️  No content found for %s", pageURL)
		return
	}

	// Extract title from first heading or use page title
	title := builder.GetRoot().Title
	if title == "Documentation" || title == "" {
		title = doc.Find("title").Text()
		if title == "" {
			title = pageURL
		}
	}

	// Send structured markdown result
	f.resultsChan <- fmt.Sprintf("%s\t%s\t%s", pageURL, title, content)

	// Extract links with proper DOM traversal (always, for both crawling and llm.txt)
	links := ExtractAllLinksFromPage(doc, pageURL)
	
	// Submit extracted links for crawling
	for _, link := range links {
		if depth < f.config.MaxDepth {
			f.submitPage(link.URL, depth+1)
		}
	}
	
	// Add to llm.txt entries if requested
	if f.config.GenerateLLMTxt {
		f.llmMutex.Lock()
		for _, link := range links {
			// Skip anchor links and external links
			if strings.HasPrefix(link.URL, "#") || !strings.Contains(link.URL, f.config.BaseURL) {
				continue
			}
			
			f.llmEntries = append(f.llmEntries, LLMTxtEntry{
				Type:        ClassifyPage(link.URL, link.Text),
				Title:       link.Text,
				URL:         link.URL,
				Description: link.Context, // Full context including surrounding text
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
