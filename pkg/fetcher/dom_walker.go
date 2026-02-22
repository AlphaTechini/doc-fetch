package fetcher

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// DOMWalker traverses HTML DOM sequentially preserving structure
type DOMWalker struct {
	baseURL     *url.URL
	linkContext map[string]string // URL -> surrounding text
}

// NewDOMWalker creates a new DOM walker
func NewDOMWalker(pageURL string) (*DOMWalker, error) {
	base, err := url.Parse(pageURL)
	if err != nil {
		return nil, err
	}

	return &DOMWalker{
		baseURL:     base,
		linkContext: make(map[string]string),
	}, nil
}

// ExtractLinks walks DOM sequentially and extracts links with context
func (dw *DOMWalker) ExtractLinks(doc *goquery.Document) []LinkInfo {
	var links []LinkInfo

	// Find main content container (don't parse entire document)
	contentContainer := doc.Find("main, article, .content, .docs-content, #content, .document, .documentation").First()
	
	// If no container found, use body as fallback
	if contentContainer.Length() == 0 {
		contentContainer = doc.Find("body")
	}
	
	// If still nothing (no body tag), use entire document
	if contentContainer.Length() == 0 {
		contentContainer = doc.Selection
	}

	// Walk nodes sequentially
	contentContainer.Contents().Each(func(i int, s *goquery.Selection) {
		dw.walkNode(s, &links, "")
	})

	return links
}

// walkNode recursively walks DOM nodes preserving text flow
func (dw *DOMWalker) walkNode(s *goquery.Selection, links *[]LinkInfo, currentText string) {
	// Get the underlying HTML node
	nodes := s.Nodes
	if len(nodes) == 0 {
		return
	}

	for _, node := range nodes {
		if node.Type == html.TextNode {
			// Accumulate text
			text := strings.TrimSpace(node.Data)
			if text != "" {
				currentText += " " + text
			}
		} else if node.Type == html.ElementNode {
			// Check if this is a link
			if node.Data == "a" {
				// Extract href
				singleNode := goquery.NewDocumentFromNode(node)
				href, exists := singleNode.Attr("href")
				if exists && href != "" {
					// Resolve relative URL
					resolvedURL := dw.resolveURL(href)
					
					// Get link text
					linkText := strings.TrimSpace(singleNode.Text())
					
					// Combine with accumulated context
					context := strings.TrimSpace(currentText)
					if context != "" {
						context = context + " "
					}
					context += linkText

					// Store link with context
					*links = append(*links, LinkInfo{
						URL:     resolvedURL,
						Text:    linkText,
						Context: context,
					})
					
					// Reset text accumulator after link
					currentText = ""
				}
			} else if node.Data == "p" || node.Data == "li" || node.Data == "div" {
				// Block element - process its contents
				childSel := goquery.NewDocumentFromNode(node).Selection
				dw.walkNode(childSel, links, currentText)
				currentText = "" // Reset after block
			} else {
				// Inline element or other - recurse
				childSel := goquery.NewDocumentFromNode(node).Selection
				dw.walkNode(childSel, links, currentText)
			}
		}
	}
}

// resolveURL converts relative URLs to absolute
func (dw *DOMWalker) resolveURL(href string) string {
	rel, err := url.Parse(href)
	if err != nil {
		return href
	}

	absolute := dw.baseURL.ResolveReference(rel)
	return absolute.String()
}

// LinkInfo contains extracted link with context
type LinkInfo struct {
	URL     string // Absolute URL
	Text    string // Link text
	Context string // Surrounding text including link
}

// ExtractAllLinksFromPage is a helper for extracting all links from a page
func ExtractAllLinksFromPage(doc *goquery.Document, pageURL string) []LinkInfo {
	walker, err := NewDOMWalker(pageURL)
	if err != nil {
		return nil
	}

	return walker.ExtractLinks(doc)
}
