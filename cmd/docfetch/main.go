package main

import (
	"flag"
	"log"
	"strings"

	"github.com/AlphaTechini/doc-fetch/pkg/fetcher"
)

func main() {
	url := flag.String("url", "", "Base URL to fetch documentation from")
	output := flag.String("output", "docs.md", "Output file path")
	depth := flag.Int("depth", 2, "Maximum crawl depth")
	concurrent := flag.Int("concurrent", 5, "Concurrent fetchers (default: 5)")
	userAgent := flag.String("user-agent", "DocFetch/1.0", "Custom user agent")
	llmTxt := flag.Bool("llm-txt", false, "Generate llm.txt index file")

	flag.Parse()

	if *url == "" {
		log.Fatal("Error: URL is required\nUsage: doc-fetch --url <base-url> --output <file-path>")
	}

	// Auto-increase depth when --llm-txt is enabled (need to crawl more pages)
	effectiveDepth := *depth
	if *llmTxt && *depth < 4 {
		effectiveDepth = 4 // Minimum depth for comprehensive llm.txt
		log.Printf("📝 llm.txt enabled: increased depth from %d to %d for comprehensive link extraction", *depth, effectiveDepth)
	}

	// Validate configuration for security
	config := fetcher.Config{
		BaseURL:         *url,
		OutputPath:      *output,
		MaxDepth:        effectiveDepth,
		Workers:         *concurrent,
		UserAgent:       *userAgent,
		GenerateLLMTxt:  *llmTxt,
		LLMTxtPath:      getLLMTxtPath(*output),
	}

	if err := fetcher.ValidateConfig(&config); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Use optimized high-performance fetcher
	err := fetcher.RunOptimized(config)
	if err != nil {
		log.Fatalf("Failed to fetch documentation: %v", err)
	}

	log.Printf("Documentation successfully saved to %s", *output)
	if *llmTxt {
		llmTxtPath := *output
		if strings.HasSuffix(*output, ".md") {
			llmTxtPath = strings.TrimSuffix(*output, ".md") + ".llm.txt"
		} else {
			llmTxtPath = *output + ".llm.txt"
		}
		log.Printf("LLM.txt index generated: %s", llmTxtPath)
	}
}

// getLLMTxtPath returns the llm.txt output path based on main output path
func getLLMTxtPath(outputPath string) string {
	if strings.HasSuffix(outputPath, ".md") {
		return strings.TrimSuffix(outputPath, ".md") + ".llm.txt"
	}
	return outputPath + ".llm.txt"
}
