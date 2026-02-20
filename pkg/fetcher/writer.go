package fetcher

import (
	"bufio"
	"os"
	"strings"
)

// writeResults writes the fetched documentation to the output file
func writeResults(outputPath string, resultsChan <-chan string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	
	// Write header
	header := "# Documentation\n\nThis file contains documentation fetched by DocFetch.\n\n---\n\n"
	writer.WriteString(header)
	
	// Write all results
	for result := range resultsChan {
		if strings.TrimSpace(result) != "" {
			writer.WriteString(result)
		}
	}
	
	return nil
}