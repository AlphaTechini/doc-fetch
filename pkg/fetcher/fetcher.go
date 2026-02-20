package fetcher

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// Config holds the configuration for the documentation fetcher
type Config struct {
	BaseURL         string
	OutputPath      string
	MaxDepth        int
	Workers         int
	UserAgent       string
	GenerateLLMTxt  bool
}

// Page represents a fetched documentation page
type Page struct {
	URL     string
	Title   string
	Content string
	Links   []string
}

// LLMTxtEntry represents an entry in the llm.txt file
type LLMTxtEntry struct {
	Type        string
	Title       string
	URL         string
	Description string
}

// Run executes the documentation fetching process
func Run(config Config) error {
	// Validate configuration
	if err := validateConfig(&config); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	log.Printf("Starting documentation fetch from: %s", config.BaseURL)
	
	// Create a visited map to avoid duplicate fetching
	visited := make(map[string]bool)
	var mutex sync.Mutex
	
	// Create channel for pages and results
	pagesChan := make(chan *Page, config.Workers*2)
	resultsChan := make(chan string, config.Workers*2)
	var llmEntries []LLMTxtEntry
	
	// Start worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < config.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(config, pagesChan, resultsChan, &mutex, visited, &llmEntries)
		}()
	}
	
	// Start the initial fetch
	pagesChan <- &Page{URL: config.BaseURL, Title: "Root"}
	
	// Close pages channel when all workers are done
	go func() {
		wg.Wait()
		close(pagesChan)
		close(resultsChan)
	}()
	
	// Collect results and write to file
	err := writeResults(config.OutputPath, resultsChan)
	if err != nil {
		return err
	}
	
	// Generate LLM.txt if requested
	if config.GenerateLLMTxt {
		llmTxtPath := strings.TrimSuffix(config.OutputPath, ".md") + ".llm.txt"
		err = GenerateLLMTxt(llmEntries, llmTxtPath)
		if err != nil {
			log.Printf("Warning: Failed to generate llm.txt: %v", err)
		} else {
			log.Printf("LLM.txt generated: %s", llmTxtPath)
		}
	}
	
	return nil
}

// worker processes pages from the channel
func worker(config Config, pagesChan <-chan *Page, resultsChan chan<- string, mutex *sync.Mutex, visited map[string]bool, llmEntries *[]LLMTxtEntry) {
	client := &http.Client{
		Timeout: 30 * time.Second,
		// Add transport with security restrictions
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	
	for page := range pagesChan {
		// Validate URL before fetching
		if err := isValidURL(page.URL); err != nil {
			log.Printf("Skipping invalid URL %s: %v", page.URL, err)
			continue
		}
		
		mutex.Lock()
		if visited[page.URL] {
			mutex.Unlock()
			continue
		}
		visited[page.URL] = true
		mutex.Unlock()
		
		log.Printf("Fetching: %s", page.URL)
		
		// Rate limiting - be respectful to servers
		time.Sleep(100 * time.Millisecond)
		
		// Fetch the page
		req, err := http.NewRequest("GET", page.URL, nil)
		if err != nil {
			log.Printf("Error creating request for %s: %v", page.URL, err)
			continue
		}
		req.Header.Set("User-Agent", config.UserAgent)
		
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error fetching %s: %v", page.URL, err)
			continue
		}
		defer resp.Body.Close()
		
		if resp.StatusCode != 200 {
			log.Printf("Non-200 status code %d for %s", resp.StatusCode, page.URL)
			continue
		}
		
		// Parse HTML
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Printf("Error parsing HTML for %s: %v", page.URL, err)
			continue
		}
		
		// Extract title
		title := doc.Find("title").Text()
		if title == "" {
			title = page.URL
		}
		
		// Clean and extract content
		content := cleanContent(doc)
		if content == "" {
			log.Printf("No content found for %s", page.URL)
			continue
		}
		
		// Send result to output
		resultsChan <- fmt.Sprintf("## %s\n\n%s\n\n---\n\n", title, content)
		
		// Generate LLM.txt entry if requested
		if config.GenerateLLMTxt {
			cleanTitle := CleanTitle(title)
			entryType := ClassifyPage(page.URL, cleanTitle)
			description := ExtractDescription(content)
			
			entry := LLMTxtEntry{
				Type:        entryType,
				Title:       cleanTitle,
				URL:         page.URL,
				Description: description,
			}
			
			mutex.Lock()
			*llmEntries = append(*llmEntries, entry)
			mutex.Unlock()
		}
		
		// Extract links for further crawling (limited depth logic would go here)
		// For MVP, we'll just fetch the main page
		// Future: implement link extraction and recursive crawling
	}
}

// cleanContent extracts and cleans the main documentation content using multiple strategies
func cleanContent(doc *goquery.Document) string {
	// Strategy 1: Try semantic HTML5 elements (most reliable)
	semanticSelectors := []string{
		"main",
		"article",
		"[role='main']",
		"[role='article']",
	}
	
	for _, selector := range semanticSelectors {
		if el := doc.Find(selector); el.Length() > 0 {
			content := extractTextContent(el)
			if len(content) > 200 { // Minimum viable content
				return content
			}
		}
	}
	
	// Strategy 2: Try common class/id patterns
	classSelectors := []string{
		".content",
		".docs-content", 
		"#main-content",
		".documentation",
		".post-content",
		".markdown-body",
		".content-wrapper",
		".doc-content",
		".document",
		".entry-content",
		".page-content",
		".article-content",
		"[class*='content']",
		"[class*='docs']",
		"[class*='document']",
		"[id*='content']",
		"[id*='main']",
	}
	
	for _, selector := range classSelectors {
		if el := doc.Find(selector); el.Length() > 0 {
			content := extractTextContent(el)
			if len(content) > 200 {
				return content
			}
		}
	}
	
	// Strategy 3: Look for sections with high text density
	var bestSection *goquery.Selection
	maxTextLen := 0
	
	doc.Find("section, div").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if len(text) > maxTextLen {
			// Check if this section has more text than child elements
			childText := 0
			s.Children().Each(func(j int, c *goquery.Selection) {
				childText += len(strings.TrimSpace(c.Text()))
			})
			
			// If parent has significantly more text, it's likely the main content
			if len(text) > childText + (childText/2) && len(text) > 500 {
				maxTextLen = len(text)
				bestSection = s
			}
		}
	})
	
	if bestSection != nil {
		content := extractTextContent(bestSection)
		if len(content) > 200 {
			return content
		}
	}
	
	// Strategy 4: Fallback to body with aggressive cleaning
	body := doc.Find("body")
	if body.Length() > 0 {
		// Remove all non-content elements aggressively
		body.Find("nav, header, footer, aside, script, style, form, iframe, .sidebar, .toc, .navigation, .menu, .ads, .advertisement, [class*='nav'], [class*='menu'], [class*='sidebar'], [class*='footer'], [class*='header']").Remove()
		
		// Find the largest remaining container
		var largest *goquery.Selection
		largestSize := 0
		
		body.Find("*").Each(func(i int, s *goquery.Selection) {
			text := strings.TrimSpace(s.Text())
			if len(text) > largestSize && s.Children().Length() < 50 {
				largestSize = len(text)
				largest = s
			}
		})
		
		if largest != nil {
			content := extractTextContent(largest)
			if len(content) > 200 {
				return content
			}
		}
		
		// Last resort: entire body
		htmlContent, _ := body.Html()
		cleaned := cleanHTML(htmlContent)
		if len(cleaned) > 200 {
			return cleaned
		}
	}
	
	return ""
}

// extractTextContent extracts and cleans text from a selection
func extractTextContent(sel *goquery.Selection) string {
	// Clone the selection to avoid modifying original
	clone := sel.Clone()
	
	// Remove unwanted elements
	clone.Find("nav, header, footer, aside, script, style, form, iframe, .sidebar, .toc, .navigation, .menu, .ads, .advertisement, button, [class*='nav'], [class*='menu'], [class*='sidebar'], [class*='footer'], [class*='header'], [class*='button'], [onclick], [role='navigation'], [role='banner'], [role='contentinfo']").Remove()
	
	// Get HTML and convert to clean text
	htmlContent, err := clone.Html()
	if err != nil {
		return ""
	}
	
	return cleanHTML(htmlContent)
}

// cleanHTML performs basic HTML cleaning
func cleanHTML(htmlStr string) string {
	// Parse and extract text content while preserving structure
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return ""
	}
	
	var texts []string
	var extractText func(*html.Node)
	extractText = func(n *html.Node) {
		if n.Type == html.TextNode {
			text := strings.TrimSpace(n.Data)
			if text != "" {
				texts = append(texts, text)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractText(c)
		}
	}
	
	extractText(doc)
	return strings.Join(texts, "\n")
}

// isValidURL validates that a URL is safe to fetch
func isValidURL(urlStr string) error {
	parsed, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}
	
	// Only allow HTTP/HTTPS
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return fmt.Errorf("only HTTP/HTTPS URLs allowed")
	}
	
	// Block private IP ranges
	host := parsed.Hostname()
	ip := net.ParseIP(host)
	if ip != nil {
		if ip.IsPrivate() || ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsMulticast() {
			return fmt.Errorf("private/internal IP addresses not allowed")
		}
	}
	
	// Block dangerous hostnames
	dangerousHosts := []string{"localhost", "127.0.0.1", "0.0.0.0", "::1"}
	for _, dangerous := range dangerousHosts {
		if host == dangerous {
			return fmt.Errorf("local hostnames not allowed")
		}
	}
	
	return nil
}

// validateConfig validates the entire configuration
func validateConfig(config *Config) error {
	if err := validateOutputPath(config.OutputPath); err != nil {
		return fmt.Errorf("output path validation failed: %w", err)
	}
	
	if err := isValidURL(config.BaseURL); err != nil {
		return fmt.Errorf("base URL validation failed: %w", err)
	}
	
	// Limit depth to prevent excessive crawling
	if config.MaxDepth > 10 {
		return fmt.Errorf("max depth cannot exceed 10")
	}
	
	// Limit workers to prevent resource exhaustion
	if config.Workers > 20 {
		return fmt.Errorf("concurrent workers cannot exceed 20")
	}
	
	// Ensure reasonable timeout values
	if config.Workers <= 0 {
		config.Workers = 3 // Default
	}
	if config.MaxDepth <= 0 {
		config.MaxDepth = 2 // Default
	}
	
	return nil
}