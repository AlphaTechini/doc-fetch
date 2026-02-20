package fetcher

import (
	"strings"
)

// ClassifyPage determines the type of documentation page
func ClassifyPage(url, title string) string {
	urlLower := strings.ToLower(url)
	titleLower := strings.ToLower(title)
	
	// API detection
	if strings.Contains(urlLower, "/api/") || 
	   strings.Contains(urlLower, "/pkg/") ||
	   strings.Contains(urlLower, "/reference/pkg/") ||
	   strings.Contains(titleLower, "api") ||
	   strings.Contains(titleLower, "package") {
		return "API"
	}
	
	// Guide/Tutorial detection  
	if strings.Contains(urlLower, "/guide/") ||
	   strings.Contains(urlLower, "/tutorial/") ||
	   strings.Contains(urlLower, "/learn/") ||
	   strings.Contains(urlLower, "/docs/guides/") ||
	   strings.Contains(titleLower, "guide") ||
	   strings.Contains(titleLower, "tutorial") ||
	   strings.Contains(titleLower, "getting started") {
		return "GUIDE"
	}
	
	// Reference documentation
	if strings.Contains(urlLower, "/ref/") ||
	   strings.Contains(urlLower, "/reference/") ||
	   strings.Contains(urlLower, "/spec/") ||
	   strings.Contains(titleLower, "reference") ||
	   strings.Contains(titleLower, "specification") {
		return "REFERENCE"
	}
	
	// Examples
	if strings.Contains(urlLower, "/example/") ||
	   strings.Contains(urlLower, "/examples/") ||
	   strings.Contains(titleLower, "example") {
		return "EXAMPLE"
	}
	
	// Default to section
	return "SECTION"
}