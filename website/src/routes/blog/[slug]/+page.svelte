<script lang="ts">
	import { page } from '$app/stores';
	import type { PageData } from './$types';
	import { marked } from 'marked';
	
	let { data }: { data: PageData } = $props();
	let post = data.post;
	
	const baseUrl = 'https://docfetch.dev';
	const canonicalUrl = `${baseUrl}/blog/${post.slug}`;
	
	// For now, using placeholder content - will load from markdown file in production
	const placeholderContent = `
## Quick Start

Install DocFetch and convert your first documentation site:

\`\`\`bash
npm install -g doc-fetch
doc-fetch --url https://golang.org/doc/ --output go-docs.md --llm-txt
\`\`\`

**Full article coming soon!** This is a placeholder until we properly integrate markdown loading.

The complete guide covers:
- Why LLMs need complete documentation context
- Three methods compared (manual, extensions, automated)
- Step-by-step DocFetch tutorial
- LLM.txt indexing explained
- Token optimization tips
- Real-world examples

Stay tuned!
`;
	
	const htmlContent = marked(placeholderContent);
	
	// Helper functions for related posts
	function getRelatedPostTitle(slug: string): string {
		const titles: Record<string, string> = {
			'llm-txt-index-guide': 'LLM.txt Explained: The Secret to Better AI Context Navigation',
			'ai-agent-documentation-problem': 'Why AI Agents Can\'t Read Documentation (And How to Fix It)',
			'best-practices-rag-context-preparation': 'RAG Context Preparation: 7 Best Practices from Production',
			'compare-documentation-fetchers': 'Documentation Fetchers Compared: DocFetch vs Alternatives',
			'token-efficiency-llm-context': 'Token Efficiency: How to Fit More Context in Less Tokens'
		};
		return titles[slug] || 'Related Article';
	}
	
	function getRelatedPostExcerpt(slug: string): string {
		const excerpts: Record<string, string> = {
			'llm-txt-index-guide': 'What is llm.txt? How does it supercharge your AI agents? Complete guide to semantic documentation indexing.',
			'ai-agent-documentation-problem': 'The fundamental problem with AI agents and web navigation. Real-world solutions for production deployments.',
			'best-practices-rag-context-preparation': 'Learn from real deployments: chunking strategies, metadata enrichment, and indexing approaches that work.',
			'compare-documentation-fetchers': 'Honest comparison of tools for converting docs to markdown. Features, limitations, and when to use each.',
			'token-efficiency-llm-context': 'Reduce LLM costs by 60%+ with smart context preparation. Cleaning strategies and compression techniques.'
		};
		return excerpts[slug] || 'Read more about this topic.';
	}
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
	
	<!-- Structured Data: BlogPosting Schema -->
	<script type="application/ld+json">
		{
			"@context": "https://schema.org",
			"@type": "BlogPosting",
			"headline": {post.title},
			"description": "{post.excerpt}",
			"url": "{canonicalUrl}",
			"author": {
				"@type": "Person",
				"name": {post.author},
				"url": "https://github.com/AlphaTechini"
			},
			"datePublished": "{post.date}",
			"dateModified": "{post.modifiedDate || post.date}",
			"publisher": {
				"@type": "Organization",
				"name": "DocFetch",
				"logo": {
					"@type": "ImageObject",
					"url": "{baseUrl}/favicon.svg"
				}
			},
			"mainEntityOfPage": {
				"@type": "WebPage",
				"@id": canonicalUrl
			},
			"articleBody": "Complete guide to converting documentation websites into AI-ready markdown with LLM.txt indexing",
			"wordCount": 3500,
			"inLanguage": "en-US",
			"keywords": {post.tags.join(', ')},
			"articleSection": {post.category || 'Tutorial'}
		}
	</script>
	
	<!-- Structured Data: FAQPage Schema (for People Also Ask) -->
	<script type="application/ld+json">
		{
			"@context": "https://schema.org",
			"@type": "FAQPage",
			"mainEntity": {post.faqs.map(faq => ({
				"@type": "Question",
				"name": faq.question,
				"acceptedAnswer": {
					"@type": "Answer",
					"text": faq.answer
				}
			}))}
		}
	</script>
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
			{@html htmlContent}
		</div>

		<footer class="post-footer">
			<!-- Related Posts Section (Internal Linking for Topical Authority) -->
			<section class="related-posts">
				<h3>Related Articles</h3>
				<div class="related-grid">
					{#each post.relatedPosts || [] as relatedSlug}
						<article class="related-card">
							<h4>
								<a href={`/blog/${relatedSlug}`}>
									{getRelatedPostTitle(relatedSlug)}
								</a>
							</h4>
							<p>{getRelatedPostExcerpt(relatedSlug)}</p>
						</article>
					{/each}
				</div>
			</section>
			
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
					<a href="https://twitter.com/intent/tweet?text={encodeURIComponent(post.title)}&url={encodeURIComponent(canonicalUrl)}" target="_blank" rel="noopener noreferrer" class="share-btn twitter">Twitter</a>
					<a href="https://www.linkedin.com/sharing/share-offsite/?url={encodeURIComponent(canonicalUrl)}" target="_blank" rel="noopener noreferrer" class="share-btn linkedin">LinkedIn</a>
					<a href="https://news.ycombinator.com/submitlink?u={encodeURIComponent(canonicalUrl)}&t={encodeURIComponent(post.title)}" target="_blank" rel="noopener noreferrer" class="share-btn hn">Hacker News</a>
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
		color: #d4d4d4;
	}

	.content pre code {
		color: #d4d4d4;
	}

	.content :global(code:not(pre code)) {
		background: var(--bg-secondary);
		padding: 0.2rem 0.5rem;
		border-radius: 3px;
		font-size: 0.9em;
		color: var(--text-primary);
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

	.related-posts {
		margin-bottom: 3rem;
		padding-bottom: 2rem;
		border-bottom: 1px solid var(--border);
	}

	.related-posts h3 {
		font-size: 1.4rem;
		margin-bottom: 1.5rem;
	}

	.related-grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
		gap: 1.5rem;
	}

	.related-card {
		background: var(--bg-secondary);
		padding: 1.5rem;
		border-radius: 6px;
		transition: transform 0.2s;
	}

	.related-card:hover {
		transform: translateY(-2px);
	}

	.related-card h4 {
		font-size: 1.1rem;
		margin: 0 0 0.75rem;
	}

	.related-card h4 a {
		color: var(--text-primary);
		text-decoration: none;
	}

	.related-card h4 a:hover {
		color: var(--accent);
	}

	.related-card p {
		font-size: 0.9rem;
		color: var(--text-secondary);
		margin: 0;
		line-height: 1.6;
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
