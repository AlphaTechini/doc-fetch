# DocFetch Usage Guide

## Installation

```bash
# Clone the repository
git clone https://github.com/AlphaTechini/doc-fetch.git
cd doc-fetch

# Build the binary
go build -o doc-fetch ./cmd/docfetch

# Or install directly
go install github.com/AlphaTechini/doc-fetch/cmd/docfetch@latest
```

## Basic Usage

```bash
# Fetch documentation from a URL to a markdown file
doc-fetch --url https://golang.org/doc/ --output docs.md

# Specify custom output path
doc-fetch --url https://docs.example.com --output ./documentation/full-docs.md
```

## Advanced Options

```bash
# Control crawl depth (default: 2)
doc-fetch --url https://docs.example.com --output docs.md --depth 3

# Set concurrent workers (default: 3)
doc-fetch --url https://docs.example.com --output docs.md --concurrent 5

# Custom user agent
doc-fetch --url https://docs.example.com --output docs.md --user-agent "MyBot/1.0"
```

## Supported Documentation Sites

DocFetch works best with sites that have:
- Clear content structure
- Standard HTML markup
- Proper semantic HTML elements

Common selectors used for content extraction:
- `<main>`
- `<article>` 
- `.content`, `.docs-content`
- `#main-content`
- `.documentation`

## Output Format

The output is clean markdown that includes:
- Page titles as H2 headings
- Cleaned content with formatting preserved
- Separation between different pages with `---`

## Future Features

- [ ] Recursive link crawling
- [ ] LLM.txt generation
- [ ] PDF and other format support
- [ ] Incremental updates
- [ ] Custom CSS selectors per site