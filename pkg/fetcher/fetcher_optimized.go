package fetcher

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
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
	if err := validateConfig(&config); err != nil {
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

	// Extract content
	content := cleanContent(doc)
	if content == "" {
		atomic.AddInt32(&f.errorCount, 1)
		log.Printf("⚠️  No content found for %s", pageURL)
		return
	}

	// Extract title
	title := doc.Find("title").Text()
	if title == "" {
		title = pageURL
	}

	// Send result with URL for proper grouping
	f.resultsChan <- fmt.Sprintf("%s\t%s\t%s", pageURL, title, content)

	// Generate LLM.txt entry if requested
	if f.config.GenerateLLMTxt {
		cleanTitle := CleanTitle(title)
		
		// Try to find this URL in parent page's link context
		linkText := findLinkTextForURL(doc, pageURL)
		
		entry := LLMTxtEntry{
			Type:        ClassifyPage(pageURL, cleanTitle),
			Title:       cleanTitle,
			URL:         pageURL,
			Description: linkText, // Use link context as description
		}

		f.llmMutex.Lock()
		f.llmEntries = append(f.llmEntries, entry)
		f.llmMutex.Unlock()
	}

	// Extract links for crawling (if depth allows)
	if depth < f.config.MaxDepth {
		f.extractAndSubmitLinks(doc, pageURL, depth+1)
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

// extractAndSubmitLinks finds and queues all internal links (including data attributes)
func (f *OptimizedFetcher) extractAndSubmitLinks(doc *goquery.Document, baseURL string, depth int) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return
	}

	// Extract links from multiple sources (proper crawling strategy)
	// Check: href, data-href, data-url (frontend engineers love hiding things)
	linkSelectors := []string{"a[href]", "a[data-href]", "a[data-url]"}
	
	for _, selector := range linkSelectors {
		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			var href string
			var exists bool
			
			// Try href first, then data attributes
			href, exists = s.Attr("href")
			if !exists {
				href, exists = s.Attr("data-href")
			}
			if !exists {
				href, exists = s.Attr("data-url")
			}
			
			if !exists || href == "" {
				return
			}

			// Resolve relative URLs
			resolvedURL, err := base.Parse(href)
			if err != nil {
				return
			}

			// Only follow same-domain links
			if resolvedURL.Host != "" && resolvedURL.Host != base.Host {
				return
			}

			// Skip non-HTML resources
			if isNonHTMLResource(resolvedURL.Path) {
				return
			}

			f.submitPage(resolvedURL.String(), depth)
		})
	}
}

// isNonHTMLResource checks if URL points to non-HTML resources
func isNonHTMLResource(path string) bool {
	extensions := []string{".pdf", ".zip", ".tar", ".gz", ".exe", ".dmg", ".pkg", ".deb", ".rpm"}
	pathLower := strings.ToLower(path)
	
	for _, ext := range extensions {
		if strings.HasSuffix(pathLower, ext) {
			return true
		}
	}
	return false
}

// writeResultsOptimized writes results to file with intelligent grouping
func writeResultsOptimized(outputPath string, resultsChan <-chan string) error {
	// Use ContentGrouper for beautiful organization
	grouper := NewContentGrouper()
	
	// Parse incoming results and add to grouper
	for result := range resultsChan {
		if strings.TrimSpace(result) == "" {
			continue
		}
		
		// Format: "URL\tTitle\tContent" (tab-separated)
		parts := strings.SplitN(result, "\t", 3)
		if len(parts) < 3 {
			log.Printf("⚠️  Skipping malformed result")
			continue
		}
		
		url := strings.TrimSpace(parts[0])
		title := strings.TrimSpace(parts[1])
		content := strings.TrimSpace(parts[2])
		
		// Add to appropriate group
		grouper.AddPage(url, title, content)
	}
	
	// Generate beautifully organized markdown
	output := grouper.GenerateMarkdown()
	
	// Write to file
	return os.WriteFile(outputPath, []byte(output), 0644)
}
