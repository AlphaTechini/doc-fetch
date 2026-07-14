# DocFetch - Dynamic Documentation Fetcher 📚

**Transform entire documentation sites into AI-ready, single-file markdown with intelligent LLM.txt indexing**

🌐 **Website**: [doc-fetch.vercel.app](https://doc-fetch.vercel.app/)

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

### PyPI (Recommended for Python developers) ✨ NEW
```bash
pip install doc-fetch
```

### NPM (Recommended for JavaScript/Node.js developers)
```bash
npm install -g doc-fetch-cli
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
doc-fetch --url https://go.dev/doc/ --output ./docs/golang-full.md

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
| Flag | Description | Default |
|------|-------------|---------|
| `--url` | Base URL to fetch documentation from | **Required** |
| `--output` | Output file path | `docs.md` |
| `--depth` | Maximum crawl depth | `2` |
| `--concurrent` | Number of concurrent workers | `5` |
| `--llm-txt` | Generate AI-friendly llm.txt index | `false` |
| `--user-agent` | Custom user agent string | `DocFetch/1.0` |

**Note**: Short flags (e.g., `-c`, `-d`) have been removed for clarity. Use full flag names only.

## ⚡ Advanced Tips for Large Documentation Sites

### Faster Scraping for Large Sites (1000+ pages)

```bash
# Increase concurrency for faster crawling
doc-fetch --url https://docs.example.com --output docs.md --concurrent 15

# Reduce depth if you only need top-level pages
doc-fetch --url https://docs.example.com --output docs.md --depth 2 --concurrent 20

# For massive sites, use multiple passes with different starting URLs
doc-fetch --url https://docs.example.com/guide --output guide.md --depth 3 --concurrent 10
doc-fetch --url https://docs.example.com/api --output api.md --depth 3 --concurrent 10
```

### Recommended Settings by Site Size

| Site Size | Pages | Concurrency | Depth | Time Estimate |
|-----------|-------|-------------|-------|---------------|
| Small | <100 | 5 | 3 | ~30 seconds |
| Medium | 100-500 | 10 | 3 | ~2 minutes |
| Large | 500-2000 | 15 | 4 | ~5-10 minutes |
| Very Large | 2000+ | 20 | 4 | ~15-30 minutes |

### Troubleshooting

**"Queue full" warnings**: Increase buffer size by using higher concurrency (`--concurrent 15`)

**Slow initial crawl**: Normal - speed increases as more workers find pages

**Missing pages**: Increase depth (`--depth 4`) or start from multiple entry points

**Rate limiting**: Add delay between requests or reduce concurrency

### Best Practices

1. **Start with conservative settings** (`--concurrent 5`, `--depth 2`)
2. **Monitor output** for missing sections
3. **Adjust based on site structure** (some sites have deeper nav trees)
4. **Use --llm-txt** for AI agent consumption (generates link index)
5. **Respect robots.txt** - DocFetch honors it automatically

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

### `docs.llm.txt` - Link Index (v2.5+ Format)
```txt
# llm.txt - Documentation URL Index
# Generated by DocFetch - List of discovered documentation URLs

[Installing Go](https://go.dev/doc/install): Official Go installation guide for all platforms
[Tutorial: Getting started](https://go.dev/doc/tutorial/getting-started.html): In this tutorial you'll get a brief introduction to Go programming
[Effective Go](https://go.dev/doc/effective_go.html): A document that gives tips for writing clear idiomatic Go code
[net/http Package](https://pkg.go.dev/net/http): Package http provides HTTP client and server implementations
```

**v2.5+**: Context descriptions are now included alongside each link for better AI understanding. Sidebar and anchor-linked sub-sections are also captured.

## 🌟 Real-World Examples

### Fetch Go Documentation
```bash
doc-fetch --url https://go.dev/doc/ --output ./docs/go-documentation.md --depth 4 --llm-txt
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