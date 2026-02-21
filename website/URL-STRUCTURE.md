# DocFetch URL Structure & Information Architecture

## 2026 SEO Best Practices for URL Hierarchy

Google uses URL structure to understand:
1. **Topical hierarchy** - How content relates
2. **Site architecture** - Content organization
3. **Context signals** - What the page is about

## Current Structure (❌ Bad)

```
/blog/convert-docs-to-markdown-for-llm
/blog/llm-txt-index-guide
/blog/ai-agent-documentation-problem
```

**Problems:**
- Flat structure (no hierarchy)
- No topical clustering
- Missed semantic signals

## Target Structure (✅ Good)

### Tier 1: Category Pages (Broad Topics)

```
/web3/                    - Web3 + AI integration
/ai-infra/                - AI infrastructure
/rag/                     - RAG systems
/llm-tools/               - LLM developer tools
```

### Tier 2: Subcategory Pages (Specific Topics)

```
/web3/vector-databases/
/web3/agent-economics/
/ai-infra/cost-optimization/
/ai-infra/security/
/rag/context-preparation/
/rag/retrieval-strategies/
/llm-tools/documentation/
/llm-tools/token-efficiency/
```

### Tier 3: Individual Articles (Long-tail)

```
/web3/vector-databases/hollowdb-review
/web3/agent-economics/instant-rag-analysis
/ai-infra/cost-optimization/multi-provider-routing
/rag/context-preparation/convert-docs-to-markdown
/rag/context-preparation/llm-txt-guide
/llm-tools/documentation/docfetch-tutorial
/llm-tools/token-efficiency/compression-techniques
```

## Implementation Plan

### Phase 1: Restructure Existing Content

**Current:** `/blog/convert-docs-to-markdown-for-llm`  
**New:** `/rag/context-preparation/convert-docs-to-markdown`

**Current:** `/blog/llm-txt-index-guide`  
**New:** `/rag/context-preparation/llm-txt-guide`

**Current:** `/blog/ai-agent-documentation-problem`  
**New:** `/llm-tools/documentation/ai-agents-cant-read-docs`

### Phase 2: Create Category Landing Pages

Each category gets a landing page that:
1. Introduces the topic
2. Lists all articles in that category
3. Links to subcategories
4. Targets broad keywords

Example: `/rag/context-preparation/`
```markdown
# RAG Context Preparation

Complete guide to preparing, cleaning, and structuring context for Retrieval Augmented Generation systems.

## Topics Covered

### Documentation Conversion
- [Convert Docs to Markdown for LLMs](/rag/context-preparation/convert-docs-to-markdown)
- [LLM.txt Explained](/rag/context-preparation/llm-txt-guide)

### Chunking Strategies
- [Optimal Chunk Sizes for RAG](/rag/context-preparation/chunk-sizes)
- [Semantic vs Fixed-Size Chunking](/rag/context-preparation/semantic-chunking)

### Cleaning Techniques
- [Remove HTML Noise from Documents](/rag/context-preparation/remove-html-noise)
- [Code Block Preservation](/rag/context-preparation/preserve-code-blocks)
```

### Phase 3: Implement in SvelteKit

#### File Structure

```
src/routes/
├── rag/
│   ├── +page.svelte              # /rag/ category page
│   └── context-preparation/
│       ├── +page.svelte          # /rag/context-preparation/ subcategory
│       └── [slug]/
│           └── +page.svelte      # /rag/context-preparation/[slug]
├── llm-tools/
│   ├── +page.svelte
│   └── documentation/
│       ├── +page.svelte
│       └── [slug]/
│           └── +page.svelte
└── web3/
    ├── +page.svelte
    └── [subcategory]/
        └── [slug]/
            └── +page.svelte
```

#### Load Function Updates

```typescript
// src/routes/rag/context-preparation/[slug]/+page.ts
export const load: PageLoad = async ({ params }) => {
	const post = getPost(params.slug);
	
	return {
		post,
		category: 'RAG',
		subcategory: 'Context Preparation',
		breadcrumb: [
			{ label: 'Home', href: '/' },
			{ label: 'RAG', href: '/rag' },
			{ label: 'Context Preparation', href: '/rag/context-preparation' },
			{ label: post.title, href: `/rag/context-preparation/${post.slug}` }
		]
	};
};
```

#### Breadcrumb Navigation

Add structured breadcrumbs for UX + SEO:

```svelte
<nav class="breadcrumb" aria-label="Breadcrumb">
	<ol>
		{#each data.breadcrumb as crumb}
			<li>
				<a href={crumb.href}>{crumb.label}</a>
			</li>
		{/each}
	</ol>
</nav>
```

With Schema.org markup:

```json
{
  "@context": "https://schema.org",
  "@type": "BreadcrumbList",
  "itemListElement": [
    {
      "@type": "ListItem",
      "position": 1,
      "name": "Home",
      "item": "https://docfetch.dev"
    },
    {
      "@type": "ListItem",
      "position": 2,
      "name": "RAG",
      "item": "https://docfetch.dev/rag"
    },
    {
      "@type": "ListItem",
      "position": 3,
      "name": "Context Preparation",
      "item": "https://docfetch.dev/rag/context-preparation"
    },
    {
      "@type": "ListItem",
      "position": 4,
      "name": "How to Convert Documentation",
      "item": "https://docfetch.dev/rag/context-preparation/convert-docs-to-markdown"
    }
  ]
}
```

## Internal Linking Strategy

### Automated Contextual Linking

Create a utility that auto-links keywords:

```typescript
// lib/auto-link.ts
const keywordLinks = {
	'llm.txt': '/rag/context-preparation/llm-txt-guide',
	'RAG': '/rag',
	'documentation conversion': '/rag/context-preparation/convert-docs-to-markdown',
	'token efficiency': '/llm-tools/token-efficiency'
};

export function autoLinkContent(content: string): string {
	let linked = content;
	
	for (const [keyword, url] of Object.entries(keywordLinks)) {
		const regex = new RegExp(`\\b(${keyword})\\b`, 'gi');
		// Don't link inside existing <a> tags
		linked = linked.replace(regex, `<a href="${url}">$1</a>`);
	}
	
	return linked;
}
```

### Related Posts Algorithm

Instead of hardcoded related posts, use tag-based matching:

```typescript
function findRelatedPosts(currentPost: Post, allPosts: Post[]): Post[] {
	return allPosts
		.filter(post => post.slug !== currentPost.slug)
		.map(post => ({
			...post,
			relevanceScore: calculateRelevance(currentPost, post)
		}))
		.sort((a, b) => b.relevanceScore - a.relevanceScore)
		.slice(0, 3);
}

function calculateRelevance(post1: Post, post2: Post): number {
	let score = 0;
	
	// Same category: +3 points
	if (post1.category === post2.category) score += 3;
	
	// Same subcategory: +5 points
	if (post1.subcategory === post2.subcategory) score += 5;
	
	// Shared tags: +1 per tag
	const sharedTags = post1.tags.filter(tag => post2.tags.includes(tag));
	score += sharedTags.length;
	
	return score;
}
```

## Redirect Strategy (Preserve SEO)

When restructuring URLs, set up 301 redirects:

```javascript
// vercel.json or netlify.toml
{
  "redirects": [
    {
      "source": "/blog/convert-docs-to-markdown-for-llm",
      "destination": "/rag/context-preparation/convert-docs-to-markdown",
      "permanent": true
    },
    {
      "source": "/blog/llm-txt-index-guide",
      "destination": "/rag/context-preparation/llm-txt-guide",
      "permanent": true
    }
  ]
}
```

## URL Naming Conventions

### ✅ DO

- Use hyphens: `/context-preparation`
- Lowercase only: `/llm-tools` not `/LLM-Tools`
- Descriptive slugs: `/convert-docs-to-markdown`
- Hierarchical: `/category/subcategory/article`
- Remove stop words: `/rag/context-preparation` not `/rag/how-to-do-context-preparation`

### ❌ DON'T

- Underscores: `/context_preparation` (hard to read)
- Uppercase: `/RAG/Context-Preparation` (case sensitivity issues)
- Vague slugs: `/guide-1`, `/article-2`
- Dates in URL: `/2026/02/21/...` (dates age, content doesn't)
- Overly long: `/rag/context-preparation/how-to-convert-documentation-to-markdown-for-llms-complete-guide`

## Target Keyword Mapping

Map URLs to target keywords:

| URL | Target Keyword | Search Volume |
|-----|----------------|---------------|
| `/rag/context-preparation/convert-docs-to-markdown` | convert documentation to markdown | 2,400/mo |
| `/rag/context-preparation/llm-txt-guide` | llm.txt generator | 90/mo |
| `/llm-tools/token-efficiency` | reduce LLM token costs | 720/mo |
| `/web3/vector-databases/hollowdb-review` | hollowdb vector database | 320/mo |
| `/ai-infra/cost-optimization` | AI infrastructure cost optimization | 1,200/mo |

## Migration Checklist

- [ ] Create new directory structure in `src/routes/`
- [ ] Move existing blog posts to new locations
- [ ] Update all internal links
- [ ] Set up 301 redirects in `vercel.json`
- [ ] Update sitemap.xml
- [ ] Submit new sitemap to Google Search Console
- [ ] Monitor 404 errors in Search Console
- [ ] Update social sharing URLs
- [ ] Test breadcrumb navigation
- [ ] Verify canonical URLs point to new structure

---

**Why This Matters for SEO:**

1. **Topical Authority** - Google sees deep coverage of each topic
2. **Semantic Signals** - URL structure reinforces content themes
3. **Internal Linking** - Natural link flow from category → article
4. **User Experience** - Clear navigation, users understand where they are
5. **Crawl Efficiency** - Googlebot understands site structure

This is how you build a **content empire**, not just a blog.
