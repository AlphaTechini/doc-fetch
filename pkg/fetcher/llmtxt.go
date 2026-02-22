package fetcher

import (
	"bufio"
	"fmt"
	"os"
)

// GenerateLLMTxt creates clean structured llm.txt index
// Format follows hierarchical structure: Title → Section → Subsection
func GenerateLLMTxt(entries []LLMTxtEntry, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create llm.txt file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Group entries by page/title for hierarchical output
	pageSections := make(map[string][]LLMTxtEntry)
	for _, entry := range entries {
		pageSections[entry.Title] = append(pageSections[entry.Title], entry)
	}

	// Write header
	writer.WriteString("# Documentation Index\n\n")
	
	// Write each page with sections
	for title, sections := range pageSections {
		// Page title
		writer.WriteString(fmt.Sprintf("# %s\n\n", title))
		
		// Sections/subsections
		for _, section := range sections {
			if section.Description != "" && section.Description != "Documentation page content." {
				// Description from link context
				writer.WriteString(fmt.Sprintf("%s\n", section.Description))
			}
			
			// Source URL
			writer.WriteString(fmt.Sprintf("---\nSource: %s\n\n", section.URL))
		}
		
		writer.WriteString("---\n\n")
	}

	return nil
}