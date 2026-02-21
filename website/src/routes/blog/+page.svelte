<script lang="ts">
	import { onMount } from 'svelte';
	
	let scrolled = false;
	
	const posts = [
		{
			slug: 'convert-docs-to-markdown-for-llm',
			title: 'How to Convert Documentation to Markdown for LLMs (Complete Guide)',
			excerpt: 'Step-by-step guide: Transform entire documentation websites into clean, AI-ready markdown. Includes tools, techniques, and best practices for optimal LLM context.',
			date: 'February 21, 2026',
			readTime: '8 min read',
			tags: ['Tutorial', 'LLM', 'Markdown']
		},
		{
			slug: 'llm-txt-index-guide',
			title: 'LLM.txt Explained: The Secret to Better AI Context Navigation',
			excerpt: 'What is llm.txt? How does it supercharge your AI agents? Complete guide to implementing semantic documentation indexing for better RAG performance.',
			date: 'February 20, 2026',
			readTime: '6 min read',
			tags: ['AI', 'RAG', 'Best Practices']
		},
		{
			slug: 'ai-agent-documentation-problem',
			title: 'Why AI Agents Can\'t Read Documentation (And How to Fix It)',
			excerpt: 'The fundamental problem with AI agents and web navigation. Why single-page context matters. Real-world solutions for production deployments.',
			date: 'February 19, 2026',
			readTime: '7 min read',
			tags: ['AI Agents', 'Problem/Solution']
		},
		{
			slug: 'best-practices-rag-context-preparation',
			title: 'RAG Context Preparation: 7 Best Practices from Production',
			excerpt: 'Learn from real deployments: chunking strategies, metadata enrichment, cleaning techniques, and indexing approaches that actually work.',
			date: 'February 18, 2026',
			readTime: '10 min read',
			tags: ['RAG', 'Production', 'Best Practices']
		},
		{
			slug: 'compare-documentation-fetchers',
			title: 'Documentation Fetchers Compared: DocFetch vs Alternatives',
			excerpt: 'Honest comparison of tools for converting docs to markdown. Features, limitations, pricing, and when to use each tool.',
			date: 'February 17, 2026',
			readTime: '9 min read',
			tags: ['Comparison', 'Tools']
		},
		{
			slug: 'token-efficiency-llm-context',
			title: 'Token Efficiency: How to Fit More Context in Less Tokens',
			excerpt: 'Reduce LLM costs by 60%+ with smart context preparation. Cleaning strategies, compression techniques, and token optimization tips.',
			date: 'February 16, 2026',
			readTime: '8 min read',
			tags: ['Cost Optimization', 'LLM']
		}
	];
	
	onMount(() => {
		const handleScroll = () => {
			scrolled = window.scrollY > 50;
		};
		window.addEventListener('scroll', handleScroll);
		return () => window.removeEventListener('scroll', handleScroll);
	});
</script>

<svelte:head>
	<title>DocFetch Blog - AI Documentation & LLM Context Guides</title>
	<meta name="description" content="Practical guides on converting documentation to markdown, optimizing LLM context, RAG best practices, and AI agent development." />
	<meta name="keywords" content="LLM context, RAG, documentation conversion, AI agents, markdown, token efficiency, AI tutorials" />
	
	<!-- Open Graph -->
	<meta property="og:type" content="website" />
	<meta property="og:title" content="DocFetch Blog - AI Documentation Guides" />
	<meta property="og:description" content="Practical guides on LLM context preparation, RAG, and AI agent development" />
	
	<!-- Twitter Card -->
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content="DocFetch Blog" />
	<meta name="twitter:description" content="Guides on AI documentation and LLM context optimization" />
</svelte:head>

<div class="container">
	<header class={scrolled ? 'scrolled' : ''}>
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

	<main>
		<section class="hero">
			<h1>DocFetch Blog</h1>
			<p class="subtitle">
				Practical guides on documentation conversion, LLM context optimization, 
				RAG best practices, and AI agent development.
			</p>
		</section>

		<section class="posts">
			<div class="post-grid">
				{#each posts as post}
					<article class="post-card">
						<div class="post-meta">
							<time datetime={post.date}>{post.date}</time>
							<span class="read-time">{post.readTime}</span>
						</div>
						<h2>
							<a href={`/blog/${post.slug}`}>{post.title}</a>
						</h2>
						<p class="excerpt">{post.excerpt}</p>
						<div class="tags">
							{#each post.tags as tag}
								<span class="tag">{tag}</span>
							{/each}
						</div>
						<a href={`/blog/${post.slug}`} class="read-more">Read more →</a>
					</article>
				{/each}
			</div>
		</section>

		<section class="newsletter">
			<h2>Stay Updated</h2>
			<p>Get notified when we publish new guides. No spam, unsubscribe anytime.</p>
			<form class="newsletter-form" onsubmit|preventDefault>
				<input type="email" placeholder="your@email.com" required />
				<button type="submit">Subscribe</button>
			</form>
			<p class="note">We respect your inbox. Max 1-2 emails per month.</p>
		</section>
	</main>

	<footer>
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
		--max-width: 1200px;
	}

	:global(body) {
		margin: 0;
		padding: 0;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
		background: var(--bg-primary);
		color: var(--text-primary);
		line-height: 1.6;
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
		border-bottom: 1px solid transparent;
		z-index: 1000;
		transition: all 0.3s ease;
	}

	header.scrolled {
		border-bottom-color: var(--border);
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
	}

	nav {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 1rem 2rem;
		max-width: var(--max-width);
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

	.hero {
		padding: 8rem 0 3rem;
		text-align: center;
	}

	h1 {
		font-size: 3.5rem;
		font-weight: 800;
		margin: 0 0 1rem;
		letter-spacing: -0.02em;
	}

	.subtitle {
		font-size: 1.25rem;
		color: var(--text-secondary);
		max-width: 700px;
		margin: 0 auto;
		line-height: 1.7;
	}

	.posts {
		padding: 3rem 0 5rem;
		border-top: 1px solid var(--border);
	}

	.post-grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
		gap: 2.5rem;
	}

	.post-card {
		display: flex;
		flex-direction: column;
		padding: 2rem;
		background: var(--bg-secondary);
		border-radius: 8px;
		transition: transform 0.2s, box-shadow 0.2s;
	}

	.post-card:hover {
		transform: translateY(-4px);
		box-shadow: 0 8px 16px rgba(0, 0, 0, 0.08);
	}

	.post-meta {
		display: flex;
		gap: 1rem;
		font-size: 0.875rem;
		color: var(--text-muted);
		margin-bottom: 1rem;
	}

	.post-card h2 {
		font-size: 1.5rem;
		margin: 0 0 1rem;
		line-height: 1.3;
	}

	.post-card h2 a {
		color: var(--text-primary);
		text-decoration: none;
	}

	.post-card h2 a:hover {
		color: var(--accent);
	}

	.excerpt {
		color: var(--text-secondary);
		line-height: 1.7;
		margin: 0 0 1.5rem;
		flex-grow: 1;
	}

	.tags {
		display: flex;
		gap: 0.5rem;
		flex-wrap: wrap;
		margin-bottom: 1.5rem;
	}

	.tag {
		background: rgba(0, 102, 204, 0.1);
		color: var(--accent);
		padding: 0.25rem 0.75rem;
		border-radius: 12px;
		font-size: 0.8rem;
		font-weight: 500;
	}

	.read-more {
		color: var(--accent);
		text-decoration: none;
		font-weight: 600;
		transition: color 0.2s;
	}

	.read-more:hover {
		color: var(--accent-hover);
	}

	.newsletter {
		background: var(--bg-secondary);
		border-radius: 8px;
		padding: 4rem 3rem;
		text-align: center;
		margin: 3rem 0;
	}

	.newsletter h2 {
		font-size: 2rem;
		margin: 0 0 1rem;
	}

	.newsletter p {
		color: var(--text-secondary);
		margin: 0 0 2rem;
		font-size: 1.1rem;
	}

	.newsletter-form {
		display: flex;
		gap: 1rem;
		max-width: 500px;
		margin: 0 auto 1rem;
	}

	.newsletter-form input {
		flex: 1;
		padding: 0.875rem 1rem;
		border: 1px solid var(--border);
		border-radius: 6px;
		font-size: 1rem;
		font-family: inherit;
	}

	.newsletter-form input:focus {
		outline: none;
		border-color: var(--accent);
		box-shadow: 0 0 0 3px rgba(0, 102, 204, 0.1);
	}

	.newsletter-form button {
		background: var(--accent);
		color: white;
		border: none;
		padding: 0.875rem 2rem;
		border-radius: 6px;
		font-weight: 600;
		font-size: 1rem;
		cursor: pointer;
		transition: background 0.2s;
	}

	.newsletter-form button:hover {
		background: var(--accent-hover);
	}

	.note {
		font-size: 0.875rem;
		color: var(--text-muted);
		margin: 0;
	}

	footer {
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
			font-size: 2.5rem;
		}

		.post-grid {
			grid-template-columns: 1fr;
		}

		.newsletter-form {
			flex-direction: column;
		}

		.footer-content {
			flex-direction: column;
			gap: 2rem;
			text-align: center;
		}

		.container {
			padding: 0 1rem;
		}
	}
</style>
