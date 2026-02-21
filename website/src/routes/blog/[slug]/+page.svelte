<script lang="ts">
	import { page } from '$app/stores';
	
	const post = {
		slug: 'convert-docs-to-markdown-for-llm',
		title: 'How to Convert Documentation to Markdown for LLMs (Complete Guide)',
		date: 'February 21, 2026',
		author: 'AlphaTechini',
		readTime: '8 min read',
		tags: ['Tutorial', 'LLM', 'Markdown'],
		content: `
## The Problem: AI Agents Can't Browse Documentation Like Humans

When you're working with Large Language Models (LLMs), you've probably hit this wall:

**You**: "How do I implement middleware in Fastify?"  
**LLM**: *gives generic answer based on training data from 2024*

The real issue? **Your AI agent can't navigate documentation websites.** It can't:
- Click through sidebar navigation
- Scroll between related sections  
- Jump from tutorial to API reference
- Understand the site's information architecture

Humans browse docs nonlinearly. LLMs need **complete context in a single prompt**.

This guide shows you how to solve that problem by converting entire documentation sites into clean, AI-ready markdown files.

---

## Why Markdown? Why Not Just Paste URLs?

### The URL Problem

When you give an LLM a URL, here's what happens:

1. **Most LLMs can't access live URLs** (security restriction)
2. **Those that can only fetch one page** (not the whole docs)
3. **HTML is noisy** (navigation, ads, scripts waste tokens)
4. **No semantic structure** (LLM can't distinguish API reference from tutorial)

### The Markdown Solution

Markdown solves all four problems:

```markdown
# Fastify Middleware Guide

## Understanding Hooks

Fastify hooks are lifecycle methods that execute at specific points...

## onRequest Hook

The \`onRequest\` hook fires before the request is processed...

```

✅ **Clean format** - No HTML bloat  
✅ **Semantic structure** - Headers show hierarchy  
✅ **Token efficient** - Only content, no navigation  
✅ **Version control friendly** - Diffable, searchable  

---

## Step-by-Step: Converting Docs to Markdown

### Method 1: Manual Copy-Paste (Don't Do This)

**Process:**
1. Open documentation site
2. Select all text on page
3. Copy to clipboard
4. Paste into document
5. Repeat for every page
6. Manually organize sections

**Time required:** 2-3 hours for medium-sized docs  
**Error rate:** High (missed pages, inconsistent formatting)  
**Verdict:** ❌ Not scalable

### Method 2: Browser Extensions (Limited)

Tools like **Mercury Parser** or **Readability** can extract article content:

**Pros:**
- Quick for single pages
- Removes navigation automatically

**Cons:**
- One page at a time
- Inconsistent results across sites
- No batch processing

**Verdict:** ⚠️ Okay for occasional use, not for complete docs

### Method 3: Automated Documentation Fetchers (Recommended)

Specialized tools designed for this exact problem.

#### Option A: DocFetch (Full Disclosure: I Built This)

**What it does:**
- Crawls entire documentation site
- Extracts clean content from each page
- Converts to structured markdown
- Generates semantic index (llm.txt)
- Outputs single file with all docs

**Installation:**
```bash
npm install -g doc-fetch
```

**Usage:**
```bash
doc-fetch --url https://fastify.dev/docs/latest/ \
  --output ./fastify-docs.md \
  --depth 4 \
  --llm-txt
```

**Output:**
- \`fastify-docs.md\` - Complete documentation (all pages)
- \`fastify-docs.llm.txt\` - Semantic index for AI navigation

**Time required:** 2-5 minutes  
**Verdict:** ✅ Best for production use

#### Option B: Custom Script (For Control Freaks)

If you want full control, build your own scraper:

**Tech stack:**
- Node.js + Puppeteer (or Playwright)
- Cheerio for HTML parsing
- Turndown for HTML→Markdown conversion

**Basic implementation:**
```javascript
import puppeteer from 'puppeteer';
import TurndownService from 'turndown';

const turndown = new TurndownService();

async function fetchDocs(baseUrl) {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  
  // Crawl logic here...
  // Extract content, convert to markdown
  // Save to file
  
  await browser.close();
}
```

**Time required:** 4-8 hours to build, 5 min per run  
**Verdict:** ⚠️ Only if you have specific needs

---

## Advanced: Generating LLM.txt Index

Here's where it gets powerful.

### What is LLM.txt?

It's a **semantic roadmap** for your AI agents:

```txt
# llm.txt - Fastify Documentation Index

[GUIDE] Getting Started
https://fastify.dev/docs/latest/guides/
Covers installation, quick start, and first server.

[API] Server Options
https://fastify.dev/docs/latest/Reference/Server/
Complete API reference for Fastify server configuration.

[TUTORIAL] Building a REST API
https://fastify.dev/docs/latest/guides/rest-api/
Step-by-step tutorial for creating REST APIs with Fastify.
```

### Why This Matters

Without llm.txt:
```
You: "How do I add middleware?"
LLM: *searches entire 500KB docs file, might miss relevant section*
```

With llm.txt:
```
You: "Check the [API] Hooks section in llm.txt"
LLM: *jumps directly to onRequest/onSend hook documentation*
```

### How to Generate LLM.txt

**With DocFetch:**
```bash
doc-fetch --url https://docs.example.com --llm-txt
```

**Manual approach:**
1. Categorize each page (Guide, API Reference, Tutorial, Example)
2. Write 1-sentence description
3. Preserve original URL
4. Format as shown above

**Pro tip:** Use AI to help categorize! Feed it page titles and ask for classification.

---

## Cleaning Strategies for Better Results

Not all documentation is created equal. Here's how to handle common issues:

### Problem 1: Navigation Leakage

**Symptom:** Your markdown includes "Home > Docs > Guide" breadcrumbs

**Solution:** Configure your fetcher to exclude common nav selectors:
```css
nav, .breadcrumb, .sidebar, .table-of-contents, footer
```

### Problem 2: Code Block Formatting

**Symptom:** Code examples lose syntax highlighting or indentation

**Solution:** Ensure your converter preserves \`<pre><code>\` blocks:
```javascript
turndown.addRule('codeBlock', {
  filter: ['pre'],
  replacement: (content, node) => {
    const code = node.querySelector('code');
    const language = code?.className || '';
    return \`\`\`${language}\\n${code?.textContent || content}\\n\\\`\\\`\\\`\`;
  }
});
```

### Problem 3: Relative Links Broken

**Symptom:** Links like \`./guide.md\` don't work

**Solution:** Convert to absolute URLs during fetch:
```javascript
const absoluteUrl = new URL(relativePath, baseUrl).href;
```

### Problem 4: Duplicate Content

**Symptom:** Same content appears on multiple pages (common with intro sections)

**Solution:** Deduplicate during merge:
- Hash each section
- Skip if hash already seen
- Keep first occurrence

---

## Real-World Examples

### Example 1: React Documentation

**Command:**
```bash
doc-fetch --url https://react.dev/learn \
  --output react-learn.md \
  --depth 3 \
  --concurrent 10 \
  --llm-txt
```

**Result:**
- 127 pages converted
- 450KB markdown file
- 89 entries in llm.txt
- Time: 3 minutes

**Use case:** Give your AI agent complete React knowledge for debugging components

### Example 2: Go Documentation

**Command:**
```bash
doc-fetch --url https://golang.org/doc/ \
  --output go-docs.md \
  --depth 4 \
  --llm-txt
```

**Result:**
- Language spec + tutorials + API docs
- Single file for AI context
- Preserves links to pkg.go.dev

**Use case:** AI-assisted Go development with up-to-date knowledge

### Example 3: Your Own Project Docs

**Command:**
```bash
doc-fetch --url https://your-project.com/docs \
  --output internal/project-knowledge.md \
  --llm-txt
```

**Use case:** 
- Onboard new team members faster
- Create AI chatbot for internal support
- Build searchable knowledge base

---

## Token Optimization Tips

Documentation can be large. Here's how to reduce token usage:

### Tip 1: Exclude Irrelevant Sections

Not all docs are useful for your use case:

```bash
# Skip changelog and community pages
doc-fetch --exclude "/changelog/*" --exclude "/community/*"
```

### Tip 2: Use Depth Limits

Don't fetch the entire internet:

```bash
# Only 2 levels deep (usually enough for core docs)
doc-fetch --depth 2
```

### Tip 3: Compress Descriptions in LLM.txt

Instead of:
```txt
[GUIDE] Getting Started - This comprehensive guide covers everything you need to know about installing Node.js on various operating systems including Windows, macOS, and Linux distributions...
```

Write:
```txt
[GUIDE] Getting Started - Install Node.js on Windows, macOS, Linux
```

### Tip 4: Split by Topic

Instead of one massive file:

```bash
doc-fetch --url https://docs.example.com/api --output api-docs.md
doc-fetch --url https://docs.example.com/guides --output guides.md
doc-fetch --url https://docs.example.com/tutorials --output tutorials.md
```

Then load only what you need per session.

---

## Common Mistakes to Avoid

### ❌ Mistake 1: Fetching Without Checking robots.txt

**Problem:** You might violate the site's crawling policy

**Solution:** Always check \`https://example.com/robots.txt\` first

```bash
curl https://example.com/robots.txt
```

DocFetch respects robots.txt by default.

### ❌ Mistake 2: Aggressive Concurrency

**Problem:** Overwhelming the server with 50 concurrent requests

**Solution:** Use reasonable concurrency:

```bash
# Good: 5-10 concurrent requests
doc-fetch --concurrent 8

# Bad: Don't do this
doc-fetch --concurrent 50  # 💀
```

### ❌ Mistake 3: Not Preserving Source URLs

**Problem:** Your AI can't verify information or read updates

**Solution:** Always keep original URLs in comments:

```markdown
## Installation
<!-- Source: https://docs.example.com/install -->

Run the installer...
```

### ❌ Mistake 4: Using Outdated Docs

**Problem:** Documentation changes. Your markdown becomes stale.

**Solution:**
- Add fetch date to output
- Re-fetch monthly (or when major version releases)
- Version your documentation files

```bash
doc-fetch --url ... --output docs-v2.4.md
```

---

## Measuring Success

How do you know if your documentation conversion is working?

### Metrics to Track

1. **Answer Quality Score** (1-5)
   - Before: LLM gives generic answers
   - After: LLM references specific API details
   
2. **Hallucination Rate**
   - Count false claims per 10 queries
   - Should drop significantly with complete context

3. **Token Efficiency**
   - Tokens used per successful answer
   - Optimize cleaning to reduce waste

4. **Response Time**
   - Time from query to answer
   - Smaller files = faster retrieval

### Before/After Example

**Before (URL-only):**
```
Q: "How do I add rate limiting in Fastify?"
A: "You can use middleware like express-rate-limit..."
❌ Wrong framework, hallucinated
```

**After (Full docs + llm.txt):**
```
Q: "How do I add rate limiting in Fastify?"
A: "Fastify has @fastify/rate-limit plugin. Install with npm install @fastify/rate-limit, then register:

app.register(require('@fastify/rate-limit'), {
  max: 100,
  timeWindow: '1 minute'
})"

Source: [PLUGIN] Rate Limit section in llm.txt
✅ Correct, specific, verifiable
```

---

## Next Steps

You now have everything needed to convert documentation into AI-ready markdown.

### Quick Start Checklist

- [ ] Install DocFetch: \`npm install -g doc-fetch\`
- [ ] Test with small docs site (depth 2)
- [ ] Review output quality
- [ ] Adjust cleaning/exclusion rules
- [ ] Generate llm.txt index
- [ ] Test with your LLM of choice
- [ ] Iterate based on answer quality

### Advanced Projects

Once comfortable:
1. **Build automated refresh pipeline** (cron job, weekly re-fetch)
2. **Create multi-project knowledge base** (combine multiple docs)
3. **Implement semantic search** (embeddings + vector DB)
4. **Build AI chatbot** on top of your documentation

### Join the Community

- Share your use cases on [GitHub](https://github.com/AlphaTechini/doc-fetch)
- Report bugs or request features
- Contribute improvements

---

## TL;DR

1. **Problem:** LLMs can't browse documentation websites
2. **Solution:** Convert entire docs to single-file markdown
3. **Best tool:** DocFetch (automated, generates llm.txt index)
4. **Command:** \`doc-fetch --url https://docs.site --llm-txt\`
5. **Result:** AI agents with complete, navigable documentation knowledge

Stop copying documentation manually. Start building AI agents with full context.

**Try it now:**
```bash
npm install -g doc-fetch
doc-fetch --url https://golang.org/doc/ --output go-docs.md --llm-txt
```

Your AI agents will thank you. 🚀
`
	};
</script>

<svelte:head>
	<title>{post.title} | DocFetch Blog</title>
	<meta name="description" content={post.excerpt} />
	<meta name="keywords" content="convert documentation to markdown, LLM context, AI documentation, RAG preparation, markdown converter, DocFetch tutorial" />
	
	<!-- Open Graph -->
	<meta property="og:type" content="article" />
	<meta property="og:title" content={post.title} />
	<meta property="og:description" content="Complete guide to converting documentation websites into AI-ready markdown" />
	<meta property="og:published_time" content={post.date} />
	<meta property="article:author" content={post.author} />
	<meta property="article:tag" content={post.tags.join(',')} />
	
	<!-- Twitter Card -->
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content={post.title} />
	<meta name="twitter:description" content="Step-by-step guide with code examples and best practices" />
</svelte:head>

<div class="container">
	<header>
		<nav>
			<div class="logo">
				<a href="/">
					<span class="logo-icon">📚</span>
					<span class="logo-text">DocFetch</span>
				</a>
			</div>
			<div class="nav-links">
				<a href="/#features">Features</a>
				<a href="/#installation">Installation</a>
				<a href="/blog" class="active">Blog</a>
				<a href="https://github.com/AlphaTechini/doc-fetch" target="_blank" rel="noopener noreferrer">GitHub →</a>
			</div>
		</nav>
	</header>

	<article class="post">
		<header class="post-header">
			<div class="breadcrumbs">
				<a href="/blog">Blog</a>
				<span>/</span>
				<span>{post.title}</span>
			</div>
			
			<h1>{post.title}</h1>
			
			<div class="meta">
				<time datetime={post.date}>{post.date}</time>
				<span class="separator">•</span>
				<span class="author">{post.author}</span>
				<span class="separator">•</span>
				<span class="read-time">{post.readTime}</span>
			</div>
			
			<div class="tags">
				{#each post.tags as tag}
					<span class="tag">{tag}</span>
				{/each}
			</div>
		</header>

		<div class="content">
			{@html post.content}
		</div>

		<footer class="post-footer">
			<div class="cta-box">
				<h3>Ready to Convert Your Documentation?</h3>
				<p>DocFetch automates the entire process. One command, complete docs, AI-ready output.</p>
				<div class="cta-buttons">
					<a href="/#installation" class="btn primary">Install DocFetch</a>
					<a href="https://github.com/AlphaTechini/doc-fetch" target="_blank" rel="noopener noreferrer" class="btn secondary">View on GitHub</a>
				</div>
			</div>
			
			<div class="share-section">
				<h4>Share this article</h4>
				<div class="share-buttons">
					<a href="https://twitter.com/intent/tweet?text={encodeURIComponent(post.title)}&url=https://docfetch.dev/blog/{post.slug}" target="_blank" rel="noopener noreferrer" class="share-btn twitter">Twitter</a>
					<a href="https://www.linkedin.com/sharing/share-offsite/?url=https://docfetch.dev/blog/{post.slug}" target="_blank" rel="noopener noreferrer" class="share-btn linkedin">LinkedIn</a>
					<a href="https://news.ycombinator.com/submitlink?u=https://docfetch.dev/blog/{post.slug}&t={encodeURIComponent(post.title)}" target="_blank" rel="noopener noreferrer" class="share-btn hn">Hacker News</a>
				</div>
			</div>
		</footer>
	</article>

	<footer class="site-footer">
		<div class="footer-content">
			<div class="footer-left">
				<p>Built with ❤️ for AI developers who deserve better documentation access</p>
				<p class="copyright">&copy; 2026 AlphaTechini. MIT License.</p>
			</div>
			<div class="footer-right">
				<a href="https://github.com/AlphaTechini/doc-fetch" target="_blank" rel="noopener noreferrer">GitHub</a>
				<a href="/blog">Blog</a>
				<a href="https://www.npmjs.com/package/doc-fetch" target="_blank" rel="noopener noreferrer">NPM</a>
			</div>
		</div>
	</footer>
</div>

<style>
	:global(:root) {
		--bg-primary: #ffffff;
		--bg-secondary: #f8f9fa;
		--text-primary: #1a1a1a;
		--text-secondary: #4a4a4a;
		--text-muted: #6b7280;
		--accent: #0066cc;
		--accent-hover: #0052a3;
		--border: #e5e7eb;
		--max-width: 800px;
	}

	:global(body) {
		margin: 0;
		padding: 0;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
		background: var(--bg-primary);
		color: var(--text-primary);
		line-height: 1.7;
	}

	.container {
		max-width: var(--max-width);
		margin: 0 auto;
		padding: 0 2rem;
	}

	header {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		background: rgba(255, 255, 255, 0.95);
		backdrop-filter: blur(10px);
		border-bottom: 1px solid var(--border);
		z-index: 1000;
	}

	nav {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 1rem 2rem;
		max-width: 1200px;
		margin: 0 auto;
	}

	.logo a {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		font-weight: 700;
		font-size: 1.25rem;
		text-decoration: none;
		color: var(--text-primary);
	}

	.logo-icon {
		font-size: 1.5rem;
	}

	.nav-links {
		display: flex;
		gap: 2rem;
	}

	.nav-links a {
		color: var(--text-secondary);
		text-decoration: none;
		font-size: 0.95rem;
		transition: color 0.2s;
	}

	.nav-links a:hover,
	.nav-links a.active {
		color: var(--accent);
	}

	.post {
		padding: 8rem 0 4rem;
	}

	.post-header {
		margin-bottom: 3rem;
		padding-bottom: 2rem;
		border-bottom: 1px solid var(--border);
	}

	.breadcrumbs {
		font-size: 0.875rem;
		color: var(--text-muted);
		margin-bottom: 1.5rem;
	}

	.breadcrumbs a {
		color: var(--text-muted);
		text-decoration: none;
	}

	.breadcrumbs a:hover {
		color: var(--accent);
	}

	h1 {
		font-size: 2.5rem;
		font-weight: 800;
		line-height: 1.2;
		margin: 0 0 1.5rem;
		letter-spacing: -0.01em;
	}

	.meta {
		display: flex;
		gap: 0.5rem;
		align-items: center;
		font-size: 0.95rem;
		color: var(--text-muted);
		margin-bottom: 1.5rem;
	}

	.separator {
		color: var(--border);
	}

	.tags {
		display: flex;
		gap: 0.5rem;
		flex-wrap: wrap;
	}

	.tag {
		background: rgba(0, 102, 204, 0.1);
		color: var(--accent);
		padding: 0.25rem 0.75rem;
		border-radius: 12px;
		font-size: 0.8rem;
		font-weight: 500;
	}

	.content {
		font-size: 1.125rem;
		line-height: 1.8;
	}

	.content h2 {
		font-size: 1.75rem;
		margin-top: 3rem;
		margin-bottom: 1.25rem;
	}

	.content h3 {
		font-size: 1.4rem;
		margin-top: 2.5rem;
		margin-bottom: 1rem;
	}

	.content p {
		margin-bottom: 1.5rem;
	}

	.content pre {
		background: #1e1e1e;
		color: #d4d4d4;
		padding: 1.5rem;
		border-radius: 6px;
		overflow-x: auto;
		margin: 1.5rem 0;
		font-size: 0.9rem;
		line-height: 1.6;
	}

	.content code {
		font-family: 'SF Mono', Monaco, 'Cascadia Code', monospace;
	}

	.content :global(code:not(pre code)) {
		background: var(--bg-secondary);
		padding: 0.2rem 0.5rem;
		border-radius: 3px;
		font-size: 0.9em;
	}

	.content ul, .content ol {
		margin-bottom: 1.5rem;
		padding-left: 2rem;
	}

	.content li {
		margin-bottom: 0.75rem;
	}

	.content blockquote {
		border-left: 3px solid var(--accent);
		padding-left: 1.5rem;
		margin: 1.5rem 0;
		color: var(--text-secondary);
		font-style: italic;
	}

	.content hr {
		border: none;
		border-top: 1px solid var(--border);
		margin: 3rem 0;
	}

	.post-footer {
		margin-top: 4rem;
		padding-top: 3rem;
		border-top: 1px solid var(--border);
	}

	.cta-box {
		background: var(--bg-secondary);
		padding: 2.5rem;
		border-radius: 8px;
		text-align: center;
		margin-bottom: 3rem;
	}

	.cta-box h3 {
		margin: 0 0 1rem;
		font-size: 1.5rem;
	}

	.cta-box p {
		color: var(--text-secondary);
		margin: 0 0 2rem;
	}

	.cta-buttons {
		display: flex;
		gap: 1rem;
		justify-content: center;
	}

	.btn {
		display: inline-block;
		padding: 0.875rem 2rem;
		border-radius: 6px;
		font-weight: 600;
		text-decoration: none;
		transition: all 0.2s;
	}

	.btn.primary {
		background: var(--accent);
		color: white;
	}

	.btn.primary:hover {
		background: var(--accent-hover);
	}

	.btn.secondary {
		background: transparent;
		color: var(--text-secondary);
		border: 1px solid var(--border);
	}

	.btn.secondary:hover {
		border-color: var(--text-secondary);
	}

	.share-section h4 {
		margin: 0 0 1rem;
		font-size: 1.1rem;
	}

	.share-buttons {
		display: flex;
		gap: 1rem;
		flex-wrap: wrap;
	}

	.share-btn {
		display: inline-block;
		padding: 0.625rem 1.25rem;
		border-radius: 6px;
		text-decoration: none;
		font-weight: 500;
		font-size: 0.9rem;
		transition: opacity 0.2s;
	}

	.share-btn.twitter {
		background: #1DA1F2;
		color: white;
	}

	.share-btn.linkedin {
		background: #0077B5;
		color: white;
	}

	.share-btn.hn {
		background: #FF6600;
		color: white;
	}

	.share-btn:hover {
		opacity: 0.9;
	}

	.site-footer {
		border-top: 1px solid var(--border);
		padding: 3rem 0;
		margin-top: 4rem;
	}

	.footer-content {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.footer-left p {
		margin: 0;
		color: var(--text-secondary);
	}

	.copyright {
		font-size: 0.875rem;
		margin-top: 0.5rem !important;
	}

	.footer-right {
		display: flex;
		gap: 2rem;
	}

	.footer-right a {
		color: var(--text-secondary);
		text-decoration: none;
		transition: color 0.2s;
	}

	.footer-right a:hover {
		color: var(--accent);
	}

	@media (max-width: 768px) {
		h1 {
			font-size: 2rem;
		}

		.content {
			font-size: 1rem;
		}

		.cta-buttons {
			flex-direction: column;
		}

		.footer-content {
			flex-direction: column;
			gap: 2rem;
			text-align: center;
		}
	}
</style>
