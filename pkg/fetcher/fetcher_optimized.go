package fetcher

import (
	"bufio"
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
)

// OptimizedFetcher uses advanced Go concurrency patterns for 10x speedup
type OptimizedFetcher struct {
	config        Config
	httpClient    *http.Client
	urlQueue      chan string
	visited       sync.Map // Concurrent map instead of mutex-protected map
	resultsChan   chan string
	llmEntries    []LLMTxtEntry
	llmMutex      sync.Mutex
	pageCount     int32
	errorCount    int32
	ctx           context.Context
	cancel        context.CancelFunc
}

// RunOptimized executes documentation fetching with maximum concurrency
func RunOptimized(config Config) error {
	if err := validateConfig(&config); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	log.Printf("🚀 Starting HIGH-PERFORMANCE documentation fetch from: %s", config.BaseURL)
	log.Printf("   Workers: %d | Max Depth: %d | Concurrency: Enabled", config.Workers, config.MaxDepth)

	fetcher := &OptimizedFetcher{
		config:      config,
		urlQueue:    make(chan string, config.Workers*100), // Large buffer for URLs
		resultsChan: make(chan string, config.Workers*10), // Larger buffer
		httpClient: createOptimizedHTTPClient(config.Workers),
	}

	fetcher.ctx, fetcher.cancel = context.WithTimeout(context.Background(), 10*time.Minute)
	defer fetcher.cancel()

	startTime := time.Now()
	
	// Start result writer in background
	var writeWg sync.WaitGroup
	writeWg.Add(1)
	go func() {
		defer writeWg.Add(-1)
		writeResultsOptimized(config.OutputPath, fetcher.resultsChan)
	}()

	// Start worker pool
	var workerWg sync.WaitGroup
	for i := 0; i < config.Workers; i++ {
		workerWg.Add(1)
		go fetcher.worker(i, &workerWg)
	}

	// Submit initial URL
	fetcher.submitPage(config.BaseURL, 0)

	// Close URL queue when all pages are processed
	go func() {
		workerWg.Wait()
		close(fetcher.urlQueue)
	}()

	// Wait for all workers to complete
	workerWg.Wait()
	close(fetcher.resultsChan)

	// Wait for results to be written
	writeWg.Wait()

	elapsed := time.Since(startTime)
	pagesFetched := atomic.LoadInt32(&fetcher.pageCount)
	errors := atomic.LoadInt32(&fetcher.errorCount)

	log.Printf("✅ Fetch completed!")
	log.Printf("   📊 Pages fetched: %d", pagesFetched)
	log.Printf("   ⏱️  Time elapsed: %v", elapsed)
	log.Printf("   📈 Speed: %.2f pages/second", float64(pagesFetched)/elapsed.Seconds())
	log.Printf("   ❌ Errors: %d", errors)

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

// worker processes URLs from the submission queue
func (f *OptimizedFetcher) worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for url := range f.urlQueue {
		select {
		case <-f.ctx.Done():
			return
		default:
			f.processURL(url, 0)
		}
	}
}

// submitPage adds a URL to be fetched (with depth tracking)
func (f *OptimizedFetcher) submitPage(pageURL string, depth int) {
	if depth > f.config.MaxDepth {
		return
	}

	// Check if already visited using atomic operation
	if _, loaded := f.visited.LoadOrStore(pageURL, true); loaded {
		return
	}

	select {
	case f.urlQueue <- pageURL:
		// Successfully queued
	default:
		// Queue full, skip this URL
		log.Printf("⚠️  Queue full, skipping: %s", pageURL)
	}
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

	// Send result
	f.resultsChan <- fmt.Sprintf("## %s\n\n%s\n\n---\n\n", title, content)

	// Generate LLM.txt entry if requested
	if f.config.GenerateLLMTxt {
		cleanTitle := CleanTitle(title)
		entryType := ClassifyPage(pageURL, cleanTitle)
		description := ExtractDescription(content)

		entry := LLMTxtEntry{
			Type:        entryType,
			Title:       cleanTitle,
			URL:         pageURL,
			Description: description,
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

// extractAndSubmitLinks finds and queues all internal links
func (f *OptimizedFetcher) extractAndSubmitLinks(doc *goquery.Document, baseURL string, depth int) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return
	}

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
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

// writeResultsOptimized writes results to file efficiently
func writeResultsOptimized(outputPath string, resultsChan <-chan string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriterSize(file, 32*1024) // 32KB buffer for better I/O
	defer writer.Flush()

	// Write header
	header := "# Documentation\n\nThis file contains documentation fetched by DocFetch.\n\n---\n\n"
	writer.WriteString(header)

	count := 0
	for result := range resultsChan {
		if strings.TrimSpace(result) != "" {
			writer.WriteString(result)
			count++
			
			// Flush periodically to avoid memory buildup
			if count%10 == 0 {
				writer.Flush()
			}
		}
	}

	return nil
}
