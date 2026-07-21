package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AlphaTechini/doc-fetch/pkg/fetcher"
)

var version = "dev"

func main() {
	versionFlag := flag.Bool("version", false, "Show version")
	url := flag.String("url", "", "Base URL to fetch documentation from")
	output := flag.String("output", "docs.md", "Output file path")
	depth := flag.Int("depth", 2, "Maximum crawl depth")
	concurrent := flag.Int("concurrent", 5, "Concurrent fetchers (default: 5)")
	userAgent := flag.String("user-agent", "DocFetch/1.0", "Custom user agent")
	llmTxt := flag.Bool("llm-txt", false, "Generate llm.txt index file")
	doctorFlag := flag.Bool("doctor", false, "Check for duplicate global installations")

	flag.Parse()

	if *doctorFlag {
		runDoctor()
		return
	}

	warnIfOutdated(version)

	if *versionFlag {
		fmt.Printf("doc-fetch version %s\n", version)
		os.Exit(0)
	}

	if *url == "" {
		log.Fatal("Error: URL is required\nUsage: doc-fetch --url <base-url> --output <file-path>")
	}

	effectiveDepth := *depth

	// Default to depth 4 for --llm-txt if depth wasn't explicitly set
	if *llmTxt {
		depthSet := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name == "depth" {
				depthSet = true
			}
		})
		if !depthSet {
			effectiveDepth = 4
		}
	}

	// Validate configuration for security
	config := fetcher.Config{
		BaseURL:        *url,
		OutputPath:     *output,
		MaxDepth:       effectiveDepth,
		Workers:        *concurrent,
		UserAgent:      *userAgent,
		GenerateLLMTxt: *llmTxt,
	}

	if err := fetcher.ValidateConfig(&config); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Use optimized high-performance fetcher
	err := fetcher.RunOptimized(config)
	if err != nil {
		log.Fatalf("Failed to fetch documentation: %v", err)
	}

	if *llmTxt {
		log.Printf("LLM.txt index generated: %s", *output)
	} else {
		log.Printf("Documentation successfully saved to %s", *output)
	}
}
