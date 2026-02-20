package fetcher

import (
	"strings"
	"regexp"
)

// ExtractDescription creates a concise description from page content
func ExtractDescription(content string) string {
	// Clean up the content
	content = strings.TrimSpace(content)
	
	// Remove extra whitespace and newlines
	content = regexp.MustCompile(`\s+`).ReplaceAllString(content, " ")
	
	// Split into sentences
	sentences := strings.Split(content, ". ")
	
	// Take first 1-2 sentences, but keep it under 200 characters
	if len(sentences) >= 2 {
		desc := sentences[0] + ". " + sentences[1] + "."
		if len(desc) > 200 {
			desc = sentences[0] + "."
		}
		return desc
	} else if len(sentences) == 1 {
		desc := sentences[0] + "."
		if len(desc) > 200 {
			// Truncate to 200 chars and add ellipsis
			desc = desc[:197] + "..."
		}
		return desc
	}
	
	// Fallback description
	return "Documentation page content."
}

// CleanTitle removes common suffixes and prefixes
func CleanTitle(title string) string {
	title = strings.TrimSpace(title)
	
	// Common patterns to remove
	patterns := []string{
		" - Documentation",
		" | Documentation", 
		" - Go",
		" | Go",
		" - React",
		" | React",
		" Documentation",
		" Docs",
		" API Reference",
	}
	
	for _, pattern := range patterns {
		title = strings.ReplaceAll(title, pattern, "")
	}
	
	return strings.TrimSpace(title)
}