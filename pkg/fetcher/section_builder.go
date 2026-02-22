package fetcher

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// Section represents a hierarchical document section
type Section struct {
	Title    string
	Level    int        // 1-6 for h1-h6
	Content  strings.Builder
	Children []*Section
	URL      string // Source URL
}

// SectionBuilder builds hierarchical document structure from DOM
type SectionBuilder struct {
	stack    []*Section
	root     *Section
	baseURL  string
}

// NewSectionBuilder creates a new section builder
func NewSectionBuilder(baseURL string) *SectionBuilder {
	root := &Section{
		Title:    "Documentation",
		Level:    0, // Root level
		Children: make([]*Section, 0),
		URL:      baseURL,
	}

	return &SectionBuilder{
		stack:   []*Section{root},
		root:    root,
		baseURL: baseURL,
	}
}

// BuildFromDocument walks DOM and builds section hierarchy
func (sb *SectionBuilder) BuildFromDocument(doc *goquery.Document) {
	// Find main content container
	contentContainer := doc.Find("main, article, .content, .docs-content").First()
	if contentContainer.Length() == 0 {
		contentContainer = doc.Find("body")
	}

	// Walk through all nodes in order
	contentContainer.Contents().Each(func(i int, s *goquery.Selection) {
		sb.processNode(s)
	})
}

// processNode handles a single node
func (sb *SectionBuilder) processNode(s *goquery.Selection) {
	nodes := s.Nodes
	if len(nodes) == 0 {
		return
	}

	for _, node := range nodes {
		if node.Type == html.ElementNode {
			// Check if heading
			level := sb.headingLevel(node.Data)
			if level > 0 {
				// Create new section
				title := strings.TrimSpace(goquery.NewDocumentFromNode(node).Text())
				sb.createSection(title, level)
			} else {
				// Append content to current section
				content := sb.renderNode(node)
				if strings.TrimSpace(content) != "" {
					sb.currentSection().Content.WriteString(content)
					sb.currentSection().Content.WriteString("\n\n")
				}
			}
		}
	}
}

// headingLevel returns 1-6 for h1-h6, 0 otherwise
func (sb *SectionBuilder) headingLevel(tagName string) int {
	if len(tagName) != 2 || tagName[0] != 'h' {
		return 0
	}

	level := int(tagName[1] - '0')
	if level >= 1 && level <= 6 {
		return level
	}
	return 0
}

// createSection adds a new section at appropriate hierarchy level
func (sb *SectionBuilder) createSection(title string, level int) {
	newSection := &Section{
		Title:    title,
		Level:    level,
		Children: make([]*Section, 0),
	}

	// Pop stack until we find parent with lower level
	for len(sb.stack) > 1 && sb.stack[len(sb.stack)-1].Level >= level {
		sb.stack = sb.stack[:len(sb.stack)-1]
	}

	// Add as child of current top of stack
	parent := sb.stack[len(sb.stack)-1]
	parent.Children = append(parent.Children, newSection)

	// Push new section onto stack
	sb.stack = append(sb.stack, newSection)
}

// currentSection returns the section at top of stack
func (sb *SectionBuilder) currentSection() *Section {
	if len(sb.stack) == 0 {
		return sb.root
	}
	return sb.stack[len(sb.stack)-1]
}

// renderNode converts HTML node to markdown text
func (sb *SectionBuilder) renderNode(node *html.Node) string {
	switch node.Data {
	case "p":
		return sb.renderChildren(node)
	case "br":
		return "\n"
	case "strong", "b":
		return "**" + sb.renderChildren(node) + "**"
	case "em", "i":
		return "*" + sb.renderChildren(node) + "*"
	case "code":
		return "`" + sb.renderChildren(node) + "`"
	case "pre":
		return "```\n" + sb.renderChildren(node) + "\n```"
	case "ul":
		return sb.renderList(node, false)
	case "ol":
		return sb.renderList(node, true)
	case "a":
		return sb.renderLink(node)
	default:
		return sb.renderChildren(node)
	}
}

// renderChildren renders all child nodes
func (sb *SectionBuilder) renderChildren(node *html.Node) string {
	var result strings.Builder

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.TextNode {
			result.WriteString(child.Data)
		} else if child.Type == html.ElementNode {
			result.WriteString(sb.renderNode(child))
		}
	}

	return result.String()
}

// renderList renders ul/ol as markdown list
func (sb *SectionBuilder) renderList(node *html.Node, ordered bool) string {
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
			result.WriteString(sb.renderChildren(child))
			itemNum++
		}
	}

	return result.String()
}

// renderLink renders <a> as markdown link
func (sb *SectionBuilder) renderLink(node *html.Node) string {
	var href, text string

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			href = attr.Val
			break
		}
	}

	text = sb.renderChildren(node)

	if href != "" {
		return "[" + text + "](" + href + ")"
	}
	return text
}

// GetRoot returns the root section
func (sb *SectionBuilder) GetRoot() *Section {
	return sb.root
}

// GenerateMarkdown produces structured markdown output
func (sb *SectionBuilder) GenerateMarkdown() string {
	var result strings.Builder

	sb.writeSection(&result, sb.root, 0)

	return result.String()
}

// writeSection recursively writes sections to markdown
func (sb *SectionBuilder) writeSection(result *strings.Builder, section *Section, depth int) {
	// Write section title (except root)
	if section.Level > 0 {
		prefix := strings.Repeat("#", section.Level+1)
		result.WriteString(prefix + " " + section.Title + "\n\n")
		
		// Add source URL if available
		if section.URL != "" && section.Level == 1 {
			result.WriteString("*Source: " + section.URL + "*\n\n")
		}
	}

	// Write content
	if section.Content.Len() > 0 {
		result.WriteString(section.Content.String())
	}

	// Write children
	for _, child := range section.Children {
		sb.writeSection(result, child, depth+1)
	}
}
