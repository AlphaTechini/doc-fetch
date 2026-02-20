# DocFetch - Dynamic Documentation Fetcher

**Fetch entire documentation sites into single, AI-friendly markdown files**

Most AIs can't navigate documentation like humans do. They can't scroll through sections, click sidebar links, or explore related pages. DocFetch solves this by converting entire documentation sites into comprehensive markdown files that contain every section and piece of information in a format that LLMs love.

## Features

- 🌐 **Full site crawling**: Automatically discovers and fetches all documentation pages
- 📝 **Clean markdown conversion**: Converts HTML to clean, structured markdown
- 🧹 **Content cleaning**: Removes navigation, headers, footers, ads, and other noise
- ⚡ **Concurrent fetching**: Fast downloads with configurable concurrency
- 🤖 **Agent-optimized**: Output format designed specifically for LLM consumption
- 🔮 **Future LLM.txt support**: Planned AI-friendly index generation

## Installation

```bash
go install github.com/AlphaTechini/doc-fetch/cmd/docfetch@latest
```

## Usage

### Basic Usage
```bash
# Fetch entire documentation site to single markdown file
docfetch https://golang.org/doc/ --output ./docs/golang-full.md

# With custom settings
docfetch https://docs.example.com --output docs.md --depth 3 --concurrent 5
```

### Command Options
```
--url, -u          Base URL to fetch documentation from (required)
--output, -o       Output file path (default: docs.md)
--depth, -d        Maximum crawl depth (default: 2)
--concurrent, -c   Number of concurrent fetchers (default: 3)
--user-agent       Custom user agent string
--timeout          Request timeout in seconds (default: 30)
```

## Examples

### Fetch Go Documentation
```bash
docfetch https://golang.org/doc/ --output ./docs/go-documentation.md --depth 4
```

### Fetch React Documentation  
```bash
docfetch https://react.dev/learn --output ./docs/react-learn.md --concurrent 10
```

### Fetch Your Own Project Docs
```bash
docfetch https://your-project.com/docs/ --output ./internal/docs.md
```

## How It Works

1. **Link Discovery**: Parses the base URL to find all internal documentation links
2. **Content Fetching**: Downloads all pages concurrently with respect for robots.txt
3. **HTML Cleaning**: Removes non-content elements (navigation, headers, footers, etc.)
4. **Markdown Conversion**: Converts cleaned HTML to structured markdown
5. **Single File Output**: Combines all documentation into one comprehensive file

## Future Features

- [ ] **LLM.txt generation**: Create AI-friendly index files with section descriptions
- [ ] **Incremental updates**: Only fetch changed pages on subsequent runs
- [ ] **Custom selectors**: Allow users to specify content areas for different sites
- [ ] **Multiple formats**: Support PDF, JSON, and other output formats
- [ ] **Token counting**: Estimate token usage for LLM context planning

## Why This Exists

Traditional documentation sites are designed for human navigation, not AI consumption. When working with LLMs, you often need to manually copy-paste multiple sections or provide incomplete context. DocFetch automates this process, giving your AI agents complete access to documentation without the manual overhead.

## Contributing

Contributions are welcome! Please open an issue or pull request on GitHub.

## License

MIT License