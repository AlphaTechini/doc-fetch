package fetcher

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GenerateLLMTxt creates an llm.txt file with AI-friendly documentation index
// Format: "Description Text: URL" (extracted from link context)
func GenerateLLMTxt(entries []LLMTxtEntry, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create llm.txt file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Write minimal header
	writer.WriteString("# llm.txt\n")
	writer.WriteString("# Link index with descriptions\n\n")

	// Write entries: use description (link context text) + URL
	for _, entry := range entries {
		// Skip entries without URL
		if entry.URL == "" {
			continue
		}
		
		// Use description if available, otherwise use title
		text := entry.Description
		if text == "" || text == "Documentation page content." {
			text = entry.Title
		}
		
		// Clean text (remove newlines, extra spaces)
		text = strings.TrimSpace(text)
		text = strings.ReplaceAll(text, "\n", " ")
		text = strings.ReplaceAll(text, "  ", " ")
		
		// Skip if no meaningful text
		if text == "" {
			continue
		}
		
		// Write in format: "Description: URL"
		writer.WriteString(fmt.Sprintf("%s: %s\n", text, entry.URL))
	}

	return nil
}