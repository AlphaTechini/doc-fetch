package fetcher

import (
	"fmt"
	"log"
	"net/http"
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
	}
	
	for page := range pagesChan {
		mutex.Lock()
		if visited[page.URL] {
			mutex.Unlock()
			continue
		}
		visited[page.URL] = true
		mutex.Unlock
		
		log.Printf("Fetching: %s", page.URL)
		
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

// cleanContent extracts and cleans the main documentation content
func cleanContent(doc *goquery.Document) string {
	// Common selectors for documentation content
	selectors := []string{
		"main",
		"article",
		".content",
		".docs-content",
		"#main-content",
		".documentation",
		".post-content",
		".markdown-body",
		".content-wrapper",
		".doc-content",
	}
	
	// Try each selector
	for _, selector := range selectors {
		if el := doc.Find(selector); el.Length() > 0 {
			// Remove unwanted elements
			el.Find("nav, header, footer, .sidebar, .toc, .navigation, script, style, .ad, .advertisement").Remove()
			
			// Convert to HTML and then clean
			htmlContent, err := el.Html()
			if err != nil {
				continue
			}
			
			// Basic HTML cleaning
			cleaned := cleanHTML(htmlContent)
			if cleaned != "" {
				return cleaned
			}
		}
	}
	
	// Fallback: try to get body content
	body := doc.Find("body")
	if body.Length() > 0 {
		body.Find("nav, header, footer, .sidebar, .toc, .navigation, script, style, .ad, .advertisement").Remove()
		htmlContent, _ := body.Html()
		return cleanHTML(htmlContent)
	}
	
	return ""
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