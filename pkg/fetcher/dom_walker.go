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

// walkNode recursively walks DOM nodes preserving text flow.
// Looks ahead past each link to capture trailing text in the same block.
func (dw *DOMWalker) walkNode(s *goquery.Selection, links *[]LinkInfo, currentText string) {
	nodes := s.Nodes
	if len(nodes) == 0 {
		return
	}

	for i, node := range nodes {
		if node.Type == html.TextNode {
			text := strings.TrimSpace(node.Data)
			if text != "" {
				if currentText != "" {
					currentText += " "
				}
				currentText += text
			}
		} else if node.Type == html.ElementNode {
			if node.Data == "a" {
				singleNode := goquery.NewDocumentFromNode(node)
				href, exists := singleNode.Attr("href")
				if exists && href != "" {
					resolvedURL := dw.resolveURL(href)
					linkText := strings.TrimSpace(singleNode.Text())

					// Look ahead at remaining siblings for trailing text
					trailingText := dw.collectTrailingText(nodes, i+1)

					context := currentText
					if context != "" {
						context += " "
					}
					context += linkText
					if trailingText != "" {
						context += " " + trailingText
					}
					context = strings.TrimSpace(context)

					*links = append(*links, LinkInfo{
						URL:     resolvedURL,
						Text:    linkText,
						Context: context,
					})

					currentText = ""
				}
			} else if node.Data == "p" || node.Data == "li" || node.Data == "div" {
				childSel := goquery.NewDocumentFromNode(node).Selection.Contents()
				dw.walkNode(childSel, links, currentText)
				currentText = ""
			} else {
				childSel := goquery.NewDocumentFromNode(node).Selection.Contents()
				dw.walkNode(childSel, links, currentText)
			}
		}
	}
}

// collectTrailingText walks sibling nodes after a link until the next link or block element
func (dw *DOMWalker) collectTrailingText(nodes []*html.Node, startIdx int) string {
	var parts []string
	for j := startIdx; j < len(nodes); j++ {
		node := nodes[j]
		if node.Type == html.TextNode {
			text := strings.TrimSpace(node.Data)
			if text != "" {
				parts = append(parts, text)
			}
		} else if node.Type == html.ElementNode {
			switch node.Data {
			case "a", "p", "li", "div", "ul", "ol", "table", "blockquote",
				"h1", "h2", "h3", "h4", "h5", "h6", "hr", "pre", "br", "img":
				return strings.Join(parts, " ")
			default:
				if text := dw.collectInlineText(node); text != "" {
					parts = append(parts, text)
				}
			}
		}
	}
	return strings.Join(parts, " ")
}

// collectInlineText extracts text from inline elements for trailing context
func (dw *DOMWalker) collectInlineText(node *html.Node) string {
	s := goquery.NewDocumentFromNode(node)
	return strings.TrimSpace(s.Text())
}

// ExtractNavLinks extracts links from navigation, sidebar, TOC, and menu areas
func (dw *DOMWalker) ExtractNavLinks(doc *goquery.Document) []LinkInfo {
	var links []LinkInfo
	seen := make(map[string]bool)

	navSelectors := []string{
		"nav", "aside",
		"[class*='sidebar']", "[class*='nav']", "[class*='menu']", "[class*='toc']",
		"[id*='sidebar']", "[id*='nav']", "[id*='menu']", "[id*='toc']",
		"[role='navigation']",
	}

	for _, selector := range navSelectors {
		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			// Skip if this element is inside the main content area (already crawled)
			if s.Closest("main, article, .content, .docs-content, #content, .document, .documentation").Length() > 0 {
				return
			}

			s.Find("a[href]").Each(func(j int, a *goquery.Selection) {
				href, exists := a.Attr("href")
				if !exists || href == "" {
					return
				}

				linkText := strings.TrimSpace(a.Text())
				if linkText == "" {
					return
				}

				resolvedURL := dw.resolveURL(href)
				if seen[resolvedURL] {
					return
				}
				seen[resolvedURL] = true

				links = append(links, LinkInfo{
					URL:     resolvedURL,
					Text:    linkText,
					Context: linkText,
				})
			})
		})
	}

	return links
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

// ExtractNavLinksFromPage extracts sidebar/nav links from a page
func ExtractNavLinksFromPage(doc *goquery.Document, pageURL string) []LinkInfo {
	walker, err := NewDOMWalker(pageURL)
	if err != nil {
		return nil
	}

	return walker.ExtractNavLinks(doc)
}
