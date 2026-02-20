package fetcher

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LLMTxtEntry represents a single entry in the llm.txt file
type LLMTxtEntry struct {
	Type        string // SECTION, API, GUIDE, REFERENCE, TUTORIAL
	Title       string
	URL         string
	Description string
}

// GenerateLLMTxt creates an llm.txt file with AI-friendly documentation index
func GenerateLLMTxt(entries []LLMTxtEntry, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create llm.txt file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.WriteString("# llm.txt - AI-friendly documentation index\n")
	writer.WriteString("# This file helps LLMs quickly find relevant documentation sections\n\n")

	for _, entry := range entries {
		// Write entry in the format: [TYPE] Title
		writer.WriteString(fmt.Sprintf("[%s] %s\n", 
			strings.ToUpper(entry.Type), entry.Title))
		// Write URL
		writer.WriteString(entry.URL + "\n")
		// Write description
		writer.WriteString(entry.Description + "\n\n")
	}

	return nil
}