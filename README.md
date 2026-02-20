# DocFetch - Dynamic Documentation Fetcher 📚

**Transform entire documentation sites into AI-ready, single-file markdown with intelligent LLM.txt indexing**

Most AIs can't navigate documentation like humans do. They can't scroll through sections, click sidebar links, or explore related pages. **DocFetch solves this fundamental problem** by converting entire documentation sites into comprehensive, clean markdown files that contain every section and piece of information in a format that LLMs love.

## 🚀 Why DocFetch is Essential for AI Development

### 🤖 **AI/LLM Optimization**
- **Single-file consumption**: No more fragmented context across multiple pages
- **Clean, structured markdown**: Perfect token efficiency for LLM context windows  
- **Intelligent LLM.txt generation**: AI-friendly index with semantic categorization
- **Noise removal**: Automatically strips navigation, headers, footers, ads, and buttons

### ⚡ **Developer Productivity**
- **One command automation**: Replace hours of manual copy-pasting with a single CLI command
- **Complete documentation access**: Give your AI agents full access to official documentation
- **Consistent formatting**: Uniform structure across different documentation sites
- **Version control friendly**: Markdown files work perfectly with Git

### 🎯 **Smart Content Intelligence**
- **Automatic page classification**: Identifies APIs, guides, references, and examples
- **Semantic descriptions**: Generates concise, relevant descriptions for each section
- **URL preservation**: Maintains original source links for verification
- **Adaptive content extraction**: Works with diverse documentation site structures

### 🔧 **Production Ready**
- **Concurrent fetching**: Fast downloads with configurable concurrency
- **Respectful crawling**: Honors robots.txt and includes rate limiting
- **Cross-platform**: Works on Windows, macOS, and Linux
- **Multiple installation options**: NPM, Go install, or direct binary download

## 📦 Installation

### NPM (Recommended for JavaScript/Node.js developers)
```bash
npm install -g doc-fetch
```

### Go (For Go developers)
```bash
go install github.com/AlphaTechini/doc-fetch/cmd/docfetch@latest
```

### Direct Binary Download
Visit [Releases](https://github.com/AlphaTechini/doc-fetch/releases) and download your platform's binary.

## 🎯 Usage

### Basic Usage
```bash
# Fetch entire documentation site to single markdown file
doc-fetch --url https://golang.org/doc/ --output ./docs/golang-full.md

# With LLM.txt generation for AI optimization
doc-fetch --url https://react.dev/learn --output docs.md --llm-txt
```

### Advanced Usage
```bash
# Comprehensive documentation fetch with all features
doc-fetch \
  --url https://docs.example.com \
  --output ./internal/docs.md \
  --depth 4 \
  --concurrent 10 \
  --llm-txt \
  --user-agent "MyBot/1.0"
```

### Command Options
| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--url` | `-u` | Base URL to fetch documentation from | **Required** |
| `--output` | `-o` | Output file path | `docs.md` |
| `--depth` | `-d` | Maximum crawl depth | `2` |
| `--concurrent` | `-c` | Number of concurrent fetchers | `3` |
| `--llm-txt` | | Generate AI-friendly llm.txt index | `false` |
| `--user-agent` | | Custom user agent string | `DocFetch/1.0` |

## 📁 Output Files

When using `--llm-txt`, DocFetch generates two files:

### `docs.md` - Complete Documentation
```markdown
# Documentation

This file contains documentation fetched by DocFetch.

---

## Getting Started

This guide covers installation, setup, and first program...

---

## Language Specification

Complete Go language specification and syntax...
```

### `docs.llm.txt` - AI-Friendly Index
```txt
# llm.txt - AI-friendly documentation index

[GUIDE] Getting Started
https://golang.org/doc/install
Covers installation, setup, and first program.

[REFERENCE] Language Specification  
https://golang.org/ref/spec
Complete Go language specification and syntax.

[API] net/http
https://pkg.go.dev/net/http
HTTP client/server implementation.
```

## 🌟 Real-World Examples

### Fetch Go Documentation
```bash
doc-fetch --url https://golang.org/doc/ --output ./docs/go-documentation.md --depth 4 --llm-txt
```

### Fetch React Documentation  
```bash
doc-fetch --url https://react.dev/learn --output ./docs/react-learn.md --concurrent 10 --llm-txt
```

### Fetch Your Own Project Docs
```bash
doc-fetch --url https://your-project.com/docs/ --output ./internal/docs.md --llm-txt
```

## 🤖 How LLM.txt Supercharges Your AI

The generated `llm.txt` file acts as a **semantic roadmap** for your AI agents:

1. **Precise Navigation**: Agents can query specific sections without scanning entire documents
2. **Context Awareness**: Know whether they're looking at an API reference vs. a tutorial
3. **Efficient Retrieval**: Jump directly to relevant content based on query intent
4. **Source Verification**: Always maintain links back to original documentation

**Example AI Prompt Enhancement:**
```
Instead of: "What does the net/http package do?"
Your AI can now: "Check the [API] net/http section in llm.txt for HTTP client/server implementation details"
```

## 🏗️ How It Works

1. **Link Discovery**: Parses the base URL to find all internal documentation links
2. **Content Fetching**: Downloads all pages concurrently with respect for robots.txt  
3. **HTML Cleaning**: Removes non-content elements (navigation, headers, footers, etc.)
4. **Markdown Conversion**: Converts cleaned HTML to structured markdown
5. **Intelligent Classification**: Categorizes pages as API, GUIDE, REFERENCE, or EXAMPLE
6. **Description Generation**: Creates concise, relevant descriptions for each section
7. **Single File Output**: Combines all documentation into one comprehensive file
8. **LLM.txt Generation**: Creates AI-friendly index with semantic categorization

## 🚀 Future Features

- **Incremental updates**: Only fetch changed pages on subsequent runs
- **Custom selectors**: Allow users to specify content areas for different sites  
- **Multiple formats**: Support PDF, JSON, and other output formats
- **Token counting**: Estimate token usage for LLM context planning
- **Advanced classification**: Machine learning-based page type detection

## 💡 Why This Exists

Traditional documentation sites are designed for **human navigation**, not **AI consumption**. When working with LLMs, you often need to manually copy-paste multiple sections or provide incomplete context. DocFetch automates this process, giving your AI agents complete access to documentation without the manual overhead.

**Stop wasting time copying documentation. Start building AI agents with complete knowledge.**

## 🤝 Contributing

Contributions are welcome! Please open an issue or pull request on GitHub.

## 📄 License

MIT License

---

**Built with ❤️ for AI developers who deserve better documentation access**