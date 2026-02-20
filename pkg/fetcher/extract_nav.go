package fetcher

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ExtractNavigationStructure extracts nav elements with h2/h3, ul/li, and hrefs
func ExtractNavigationStructure(doc *goquery.Document) string {
	var result strings.Builder
	
	result.WriteString("# Navigation Structure\n\n")
	
	// Find all nav elements
	doc.Find("nav").Each(func(i int, nav *goquery.Selection) {
		result.WriteString(fmt.Sprintf("## Navigation Block %d\n\n", i+1))
		
		// Look for headings in nav
		nav.Find("h1, h2, h3, h4, h5, h6").Each(func(j int, h *goquery.Selection) {
			tagName := h.Get(0).Data
			text := strings.TrimSpace(h.Text())
			if text != "" {
				result.WriteString(fmt.Sprintf("### %s: %s\n\n", tagName, text))
			}
			
			// Find ul under this heading
			h.NextFiltered("ul").Each(func(k int, ul *goquery.Selection) {
				result.WriteString(extractListWithLinks(ul, 1))
			})
		})
		
		// Also find ul directly in nav
		nav.ChildrenFiltered("ul").Each(func(k int, ul *goquery.Selection) {
			result.WriteString(extractListWithLinks(ul, 1))
		})
		
		result.WriteString("---\n\n")
	})
	
	// Also look for elements with navigation-related classes/ids
	navSelectors := []string{
		"[class*='nav']",
		"[id*='nav']",
		"[class*='menu']",
		"[id*='menu']",
		"[role='navigation']",
		".toc",
		"#toc",
		"[class*='toc']",
		"[id*='toc']",
	}
	
	for _, selector := range navSelectors {
		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			// Skip if already processed as nav element
			if s.Parent().Is("nav") {
				return
			}
			
			result.WriteString(fmt.Sprintf("## Navigation Element (matched: %s)\n\n", selector))
			
			// Extract headings
			s.Find("h1, h2, h3, h4, h5, h6").Each(func(j int, h *goquery.Selection) {
				tagName := h.Get(0).Data
				text := strings.TrimSpace(h.Text())
				if text != "" {
					result.WriteString(fmt.Sprintf("### %s: %s\n\n", tagName, text))
				}
			})
			
			// Extract lists with links
			s.Find("ul, ol").Each(func(k int, list *goquery.Selection) {
				result.WriteString(extractListWithLinks(list, 1))
			})
			
			result.WriteString("---\n\n")
		})
	}
	
	return result.String()
}

// extractListWithLinks extracts list items with their href attributes
func extractListWithLinks(list *goquery.Selection, indentLevel int) string {
	var result strings.Builder
	
	indent := strings.Repeat("  ", indentLevel)
	
	list.Find("> li").Each(func(i int, li *goquery.Selection) {
		// Get the text
		text := strings.TrimSpace(li.Text())
		
		// Find any links in this li
		li.Find("a[href]").Each(func(j int, a *goquery.Selection) {
			href, exists := a.Attr("href")
			linkText := strings.TrimSpace(a.Text())
			if exists && href != "" {
				result.WriteString(fmt.Sprintf("%s- [%s](%s)\n", indent, linkText, href))
			}
		})
		
		// If no links found, just add the text
		if li.Find("a[href]").Length() == 0 && text != "" {
			result.WriteString(fmt.Sprintf("%s- %s\n", indent, text))
		}
		
		// Recursively process nested lists
		li.ChildrenFiltered("ul, ol").Each(func(k int, nested *goquery.Selection) {
			result.WriteString(extractListWithLinks(nested, indentLevel+1))
		})
	})
	
	return result.String()
}

// ExtractAllLinks extracts all links from the page with context
func ExtractAllLinks(doc *goquery.Document, baseURL string) string {
	var result strings.Builder
	
	result.WriteString("# All Links Found\n\n")
	
	linksFound := 0
	
	// Group links by section
	doc.Find("section, article, div[class*='content'], div[id*='content']").Each(func(i int, section *goquery.Selection) {
		sectionLinks := 0
		var sectionResult strings.Builder
		
		// Get section title
		title := ""
		section.Find("h1, h2, h3").First().Each(func(j int, h *goquery.Selection) {
			title = strings.TrimSpace(h.Text())
		})
		
		if title == "" {
			title = fmt.Sprintf("Section %d", i+1)
		}
		
		sectionResult.WriteString(fmt.Sprintf("## %s\n\n", title))
		
		// Find all links in this section
		section.Find("a[href]").Each(func(j int, a *goquery.Selection) {
			href, exists := a.Attr("href")
			text := strings.TrimSpace(a.Text())
			if exists && href != "" && text != "" {
				sectionResult.WriteString(fmt.Sprintf("- [%s](%s)\n", text, href))
				sectionLinks++
				linksFound++
			}
		})
		
		if sectionLinks > 0 {
			result.WriteString(sectionResult.String())
			result.WriteString("\n")
		}
	})
	
	result.WriteString(fmt.Sprintf("\n**Total links found: %d**\n", linksFound))
	
	return result.String()
}
