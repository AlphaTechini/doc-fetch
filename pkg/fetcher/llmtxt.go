package fetcher

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GenerateLLMTxt creates an llm.txt file with AI-friendly documentation index
// Format: "Title: URL" (one line per page)
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
	writer.WriteString("# Documentation index for AI agents\n\n")

	// Write entries in simple format: "Title: URL"
	for _, entry := range entries {
		// Skip entries without title or URL
		if entry.Title == "" || entry.URL == "" {
			continue
		}
		
		// Clean title (remove special chars that might break parsing)
		cleanTitle := strings.TrimSpace(entry.Title)
		cleanTitle = strings.ReplaceAll(cleanTitle, "\n", " ")
		cleanTitle = strings.ReplaceAll(cleanTitle, "  ", " ")
		
		// Write in format: "Title: URL"
		writer.WriteString(fmt.Sprintf("%s: %s\n", cleanTitle, entry.URL))
	}

	return nil
}