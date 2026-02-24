package fetcher

import (
	"net/url"
)

// Config holds the fetcher configuration
type Config struct {
	BaseURL        string
	OutputPath     string
	MaxDepth       int
	Workers        int
	UserAgent      string
	GenerateLLMTxt bool
}

// LLMTxtEntry represents a single entry in llm.txt
type LLMTxtEntry struct {
	Type        string // API, GUIDE, REFERENCE, etc.
	Title       string // Page/link title
	URL         string // Absolute URL
	Description string // Context/description
}

// isValidURL checks if a string is a valid URL
func isValidURL(s string) error {
	_, err := url.Parse(s)
	return err
}
