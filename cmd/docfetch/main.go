package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AlphaTechini/doc-fetch/v2/pkg/fetcher"
)

var version = "dev"

func main() {
	log.SetFlags(0)
	versionFlag := flag.Bool("version", false, "Show version")
	url := flag.String("url", "", "Base URL to fetch documentation from")
	output := flag.String("output", "docs.md", "Output file path")
	depth := flag.Int("depth", 2, "Maximum crawl depth")
	concurrent := flag.Int("concurrent", 5, "Concurrent fetchers (default: 5)")
	userAgent := flag.String("user-agent", "DocFetch/1.0", "Custom user agent")
	llmTxt := flag.Bool("llm-txt", false, "Generate llm.txt index file")
	doctorFlag := flag.Bool("doctor", false, "Check for duplicate global installations")
	fixFlag := flag.Bool("fix", false, "Remove stale duplicate global installations")
	yesFlag := flag.Bool("yes", false, "Skip confirmation for --fix")
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Show detailed crawl output")
	flag.BoolVar(&verbose, "verbose", false, "Show detailed crawl output")

	flag.Parse()

	if *doctorFlag {
		runDoctor()
		return
	}
	if *fixFlag {
		runFix(*yesFlag)
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
		Verbose:        verbose,
	}

	if err := fetcher.ValidateConfig(&config); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	display := newProgressDisplay(config.BaseURL, !verbose)
	config.Progress = display.Update
	display.Start()
	result, err := fetcher.RunOptimized(config)
	display.Stop()
	if err != nil {
		log.Fatalf("Failed to fetch documentation: %v", err)
	}

	failureSummary := ""
	if result.Failed > 0 {
		failureSummary = fmt.Sprintf(" (%d failed)", result.Failed)
	}
	pageLabel := "pages"
	if result.Processed == 1 {
		pageLabel = "page"
	}
	fmt.Printf("Processed %d %s in %.1fs%s. Saved to %s.\n",
		result.Processed, pageLabel, result.Elapsed.Seconds(), failureSummary, *output)
}
