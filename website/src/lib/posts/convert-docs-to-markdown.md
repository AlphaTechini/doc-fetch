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

The `onRequest` hook fires before the request is processed...
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
- `fastify-docs.md` - Complete documentation (all pages)
- `fastify-docs.llm.txt` - Semantic index for AI navigation

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

**Solution:** Ensure your converter preserves `<pre><code>` blocks:
```javascript
turndown.addRule('codeBlock', {
  filter: ['pre'],
  replacement: (content, node) => {
    const code = node.querySelector('code');
    const language = code?.className || '';
    return `\`\`\`${language}\n${code?.textContent || content}\n\`\`\``;
  }
});
```

### Problem 3: Relative Links Broken

**Symptom:** Links like `./guide.md` don't work

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

**Solution:** Always check `https://example.com/robots.txt` first

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

- [ ] Install DocFetch: `npm install -g doc-fetch`
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
4. **Command:** `doc-fetch --url https://docs.site --llm-txt`
5. **Result:** AI agents with complete, navigable documentation knowledge

Stop copying documentation manually. Start building AI agents with full context.

**Try it now:**
```bash
npm install -g doc-fetch
doc-fetch --url https://golang.org/doc/ --output go-docs.md --llm-txt
```

Your AI agents will thank you. 🚀
