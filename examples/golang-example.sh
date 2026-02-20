#!/bin/bash

# Example usage for Go documentation
echo "Fetching Go documentation..."

# Basic usage
doc-fetch --url https://golang.org/doc/ --output ./docs/golang-full.md

# With custom settings
doc-fetch --url https://pkg.go.dev/std --output ./docs/go-stdlib.md --depth 3 --concurrent 5

echo "Documentation saved to docs/ directory"