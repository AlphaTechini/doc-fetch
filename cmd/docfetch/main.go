package main

import (
	"flag"
	"log"
	"os"

	"github.com/AlphaTechini/doc-fetch/pkg/fetcher"
)

func main() {
	url := flag.String("url", "", "Base URL to fetch documentation from")
	output := flag.String("output", "docs.md", "Output file path")
	depth := flag.Int("depth", 2, "Maximum crawl depth")
	concurrent := flag.Int("concurrent", 3, "Concurrent fetchers")
	userAgent := flag.String("user-agent", "DocFetch/1.0", "Custom user agent")
	llmTxt := flag.Bool("llm-txt", false, "Generate llm.txt index file")

	flag.Parse()

	if *url == "" {
		log.Fatal("Error: URL is required\nUsage: doc-fetch --url <base-url> --output <file-path>")
	}

	config := fetcher.Config{
		BaseURL:         *url,
		OutputPath:      *output,
		MaxDepth:        *depth,
		Workers:         *concurrent,
		UserAgent:       *userAgent,
		GenerateLLMTxt:  *llmTxt,
	}

	err := fetcher.Run(config)
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