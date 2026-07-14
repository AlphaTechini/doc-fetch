package fetcher

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// HTMLToMarkdown converts HTML node to markdown text
// This is a standalone converter used by SectionBuilder
func HTMLToMarkdown(node *html.Node) string {
	if node == nil {
		return ""
	}

	switch node.Data {
	case "h1":
		return "# " + strings.TrimSpace(renderChildren(node)) + "\n\n"
	case "h2":
		return "## " + strings.TrimSpace(renderChildren(node)) + "\n\n"
	case "h3":
		return "### " + strings.TrimSpace(renderChildren(node)) + "\n\n"
	case "h4":
		return "#### " + strings.TrimSpace(renderChildren(node)) + "\n\n"
	case "h5":
		return "##### " + strings.TrimSpace(renderChildren(node)) + "\n\n"
	case "h6":
		return "###### " + strings.TrimSpace(renderChildren(node)) + "\n\n"
	case "p":
		return renderChildren(node) + "\n\n"
	case "br":
		return "\n"
	case "hr":
		return "---\n\n"
	case "strong", "b":
		return "**" + renderChildren(node) + "**"
	case "em", "i":
		return "*" + renderChildren(node) + "*"
	case "code":
		// Check if parent is pre (block code) or inline
		if node.Parent != nil && node.Parent.Data == "pre" {
			return renderChildren(node)
		}
		return "`" + renderChildren(node) + "`"
	case "pre":
		return "```\n" + renderChildren(node) + "\n```\n\n"
	case "ul":
		return renderList(node, false) + "\n"
	case "ol":
		return renderList(node, true) + "\n"
	case "li":
		return renderChildren(node)
	case "a":
		return renderLink(node)
	case "table":
		return renderTable(node) + "\n\n"
	case "blockquote":
		return "> " + renderChildren(node) + "\n\n"
	default:
		return renderChildren(node)
	}
}

// renderChildren renders all child nodes recursively
func renderChildren(node *html.Node) string {
	var result strings.Builder

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.TextNode {
			result.WriteString(child.Data)
		} else if child.Type == html.ElementNode {
			result.WriteString(HTMLToMarkdown(child))
		}
	}

	return result.String()
}

// renderList renders ul/ol as markdown list
func renderList(node *html.Node, ordered bool) string {
	var result strings.Builder

	prefix := "- "
	if ordered {
		prefix = "1. "
	}

	itemNum := 0
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Data == "li" {
			if itemNum > 0 {
				result.WriteString("\n")
			}
			result.WriteString(prefix)
			result.WriteString(renderChildren(child))
			itemNum++
		}
	}

	return result.String()
}

// renderLink renders <a> as markdown link
func renderLink(node *html.Node) string {
	var href, text string

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			href = attr.Val
			break
		}
	}

	text = renderChildren(node)

	if href != "" {
		return "[" + text + "](" + href + ")"
	}
	return text
}

// renderTable converts HTML table to markdown table
func renderTable(node *html.Node) string {
	var result strings.Builder

	rowNum := 0
	for row := node.FirstChild; row != nil; row = row.NextSibling {
		if row.Data != "tr" {
			continue
		}

		result.WriteString("| ")
		colNum := 0

		for cell := row.FirstChild; cell != nil; cell = cell.NextSibling {
			if cell.Data != "td" && cell.Data != "th" {
				continue
			}

			if colNum > 0 {
				result.WriteString(" | ")
			}
			result.WriteString(strings.TrimSpace(renderChildren(cell)))
			colNum++
		}
		result.WriteString(" |\n")

		// Add separator after header row
		if rowNum == 0 {
			result.WriteString("|")
			for i := 0; i < colNum; i++ {
				result.WriteString(" --- |")
			}
			result.WriteString("\n")
		}

		rowNum++
	}

	return result.String()
}

// ExtractContentAsMarkdown extracts main content from document and converts to markdown
// Uses SectionBuilder for hierarchical structure
func ExtractContentAsMarkdown(doc *goquery.Document, baseURL string) string {
	builder := NewSectionBuilder(baseURL)
	builder.BuildFromDocument(doc)
	return builder.GenerateMarkdown()
}
