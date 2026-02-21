<script lang="ts">
	import { onMount } from 'svelte';
	
	let scrolled = false;
	
	onMount(() => {
		const handleScroll = () => {
			scrolled = window.scrollY > 50;
		};
		window.addEventListener('scroll', handleScroll);
		return () => window.removeEventListener('scroll', handleScroll);
	});
</script>

<svelte:head>
	<title>About DocFetch - Built by AlphaTechini</title>
	<meta name="description" content="DocFetch is built by AlphaTechini, a senior infrastructure engineer specializing in Web2+Web3 systems and AI orchestration." />
	<meta name="author" content="AlphaTechini" />
	
	<!-- Open Graph -->
	<meta property="og:type" content="profile" />
	<meta property="og:title" content="About DocFetch" />
	<meta property="og:description" content="Built by a senior infrastructure engineer with deep expertise in distributed systems" />
	
	<!-- Structured Data: Person Schema -->
	<script type="application/ld+json">
		{
			"@context": "https://schema.org",
			"@type": "Person",
			"name": "AlphaTechini",
			"url": "https://github.com/AlphaTechini",
			"sameAs": [
				"https://github.com/AlphaTechini",
				"https://www.npmjs.com/~alphatechini",
				"https://pypi.org/user/alphatechini/"
			],
			"jobTitle": "Senior Infrastructure Engineer",
			"worksFor": {
				"@type": "Organization",
				"name": "Open Source"
			},
			"knowsAbout": [
				"Backend Engineering",
				"Web3 Infrastructure",
				"AI Orchestration",
				"Distributed Systems",
				"Developer Tools"
			]
		}
	</script>
	
	<!-- Structured Data: SoftwareApplication -->
	<script type="application/ld+json">
		{
			"@context": "https://schema.org",
			"@type": "SoftwareApplication",
			"name": "DocFetch",
			"applicationCategory": "DeveloperApplication",
			"author": {
				"@type": "Person",
				"name": "AlphaTechini"
			},
			"description": "Transform documentation sites into AI-ready markdown",
			"url": "https://docfetch.dev",
			"dateCreated": "2026-02-20",
			"version": "2.0.4"
		}
	</script>
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
				<a href="/blog">Blog</a>
				<a href="/about" class="active">About</a>
				<a href="https://github.com/AlphaTechini/doc-fetch" target="_blank" rel="noopener noreferrer">GitHub →</a>
			</div>
		</nav>
	</header>

	<main>
		<section class="hero">
			<h1>About DocFetch</h1>
			<p class="subtitle">
				Built by a senior infrastructure engineer to solve a real problem in AI development.
			</p>
		</section>

		<section class="content">
			<article>
				<h2>The Problem</h2>
				<p>
					When building AI agents and RAG systems, I constantly hit the same wall: 
					<strong>LLMs can't navigate documentation websites like humans do.</strong>
				</p>
				<p>
					They can't click through sidebars, jump between related sections, or understand 
					the information architecture of documentation sites. This meant manually 
					copy-pasting dozens of pages to give my AI agents complete context.
				</p>
				<p>
					This is unsustainable and error-prone. There had to be a better way.
				</p>

				<h2>The Solution</h2>
				<p>
					DocFetch automates the entire process. One command converts an entire 
					documentation website into clean, structured markdown with semantic indexing 
					(llm.txt) that helps AI agents navigate efficiently.
				</p>
				<p>
					It's not just a scraper—it's purpose-built for AI/LLM workflows with features like:
				</p>
				<ul>
					<li>Intelligent page classification (API, GUIDE, TUTORIAL, EXAMPLE)</li>
					<li>Semantic descriptions for each section</li>
					<li>Token-efficient output (no HTML bloat)</li>
					<li>Respectful crawling (robots.txt, rate limiting)</li>
					<li>Cross-platform support (Python, Node.js, Go, binaries)</li>
				</ul>

				<h2>About the Author</h2>
				<div class="author-box">
					<div class="author-info">
						<h3>AlphaTechini</h3>
						<p class="role">Senior Infrastructure Engineer</p>
						<p>
							Specializing in Web2 + Web3 systems integration, AI orchestration, 
							and distributed automation. Deep expertise in building fault-tolerant 
							event systems and developer tools.
						</p>
						<p>
							<strong>Current focus:</strong> Building infrastructure-level AI 
							integrations that reduce human dependency and create strategic leverage.
						</p>
						
						<div class="links">
							<a href="https://github.com/AlphaTechini" target="_blank" rel="noopener noreferrer" class="link">
								<span class="icon">🐙</span> GitHub
							</a>
							<a href="https://www.npmjs.com/~alphatechini" target="_blank" rel="noopener noreferrer" class="link">
								<span class="icon">📦</span> NPM
							</a>
							<a href="https://pypi.org/user/alphatechini/" target="_blank" rel="noopener noreferrer" class="link">
								<span class="icon">🐍</span> PyPI
							</a>
						</div>
					</div>
				</div>

				<h2>Expertise &amp; Technologies</h2>
				<div class="tech-grid">
					<div class="tech-category">
						<h3>Backend</h3>
						<ul>
							<li>Node.js + Fastify</li>
							<li>Go (systems programming)</li>
							<li>Python</li>
							<li>MongoDB</li>
						</ul>
					</div>
					
					<div class="tech-category">
						<h3>Frontend</h3>
						<ul>
							<li>Svelte + SvelteKit</li>
							<li>Tailwind CSS</li>
							<li>Plotly.js</li>
						</ul>
					</div>
					
					<div class="tech-category">
						<h3>Web3</h3>
						<ul>
							<li>Solidity</li>
							<li>ERC-20, ERC-721</li>
							<li>Uniswap Integration</li>
							<li>On-chain Event Systems</li>
						</ul>
					</div>
					
					<div class="tech-category">
						<h3>AI/ML</h3>
						<ul>
							<li>LLM Orchestration</li>
							<li>RAG Systems</li>
							<li>Agent Architecture</li>
							<li>Prompt Engineering</li>
						</ul>
					</div>
				</div>

				<h2>Philosophy</h2>
				<blockquote>
					"Build control layers, not features. Own infrastructure where possible. 
					Prefer composability over abstraction. Operate as startup founder + 
					systems engineer + security auditor simultaneously."
				</blockquote>
				
				<p>
					This philosophy shapes every aspect of DocFetch:
				</p>
				<ul>
					<li><strong>No vendor lock-in:</strong> Works with any LLM provider</li>
					<li><strong>Own your infra:</strong> Self-hosted, open source</li>
					<li><strong>Composable:</strong> CLI tool that integrates with existing workflows</li>
					<li><strong>Security-aware:</strong> Respects robots.txt, rate limiting, SSRF protection</li>
				</ul>

				<h2>Projects</h2>
				<div class="projects">
					<div class="project">
						<h3>DocFetch</h3>
						<p>
							Transform documentation sites into AI-ready markdown with intelligent 
							LLM.txt indexing. Published on NPM, PyPI, and Go modules.
						</p>
						<a href="https://github.com/AlphaTechini/doc-fetch" target="_blank" rel="noopener noreferrer">View on GitHub →</a>
					</div>
					
					<div class="project">
						<h3>Reactive Network Trading Portfolio dApp</h3>
						<p>
							ERC-20 limited universe trading platform with stop-loss, auto-sell, 
							and RCS automation triggers. Built with Svelte + Fastify.
						</p>
						<a href="https://github.com/AlphaTechini/Reactive" target="_blank" rel="noopener noreferrer">View on GitHub →</a>
					</div>
					
					<div class="project">
						<h3>Fault-Tolerant On-Chain Event Webhook Service</h3>
						<p>
							At-least-once delivery system for blockchain events with replayability, 
							ABI-only decoding, and operational reliability focus.
						</p>
						<a href="https://github.com/AlphaTechini/Fault-Tolerant-On-Chain-Event-Webhook-Service" target="_blank" rel="noopener noreferrer">View on GitHub →</a>
					</div>
					
					<div class="project">
						<h3>xCreator (AI Social Post Engine)</h3>
						<p>
							Algorithm-aware social post generation with strict prompt template 
							injection and cost-controlled AI usage.
						</p>
						<a href="https://github.com/AlphaTechini/xCreator" target="_blank" rel="noopener noreferrer">View on GitHub →</a>
					</div>
				</div>

				<h2>Contact &amp; Contributions</h2>
				<p>
					DocFetch is open source and community-driven. Contributions, issues, and 
					feature requests are welcome!
				</p>
				<div class="cta-buttons">
					<a href="https://github.com/AlphaTechini/doc-fetch" target="_blank" rel="noopener noreferrer" class="btn primary">
						Contribute on GitHub
					</a>
					<a href="https://github.com/AlphaTechini/doc-fetch/issues" target="_blank" rel="noopener noreferrer" class="btn secondary">
						Report Issue
					</a>
				</div>
			</article>
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
				<a href="/about">About</a>
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

	.hero {
		padding: 8rem 0 3rem;
		text-align: center;
	}

	h1 {
		font-size: 3rem;
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

	.content {
		padding: 3rem 0 5rem;
		border-top: 1px solid var(--border);
	}

	h2 {
		font-size: 2rem;
		margin-top: 3rem;
		margin-bottom: 1.5rem;
	}

	h3 {
		font-size: 1.4rem;
		margin-top: 2rem;
		margin-bottom: 1rem;
	}

	p {
		margin-bottom: 1.5rem;
		color: var(--text-secondary);
	}

	ul {
		margin-bottom: 1.5rem;
		padding-left: 2rem;
	}

	li {
		margin-bottom: 0.75rem;
		color: var(--text-secondary);
	}

	blockquote {
		border-left: 3px solid var(--accent);
		padding-left: 1.5rem;
		margin: 2rem 0;
		font-style: italic;
		background: var(--bg-secondary);
		padding: 1.5rem;
		border-radius: 0 6px 6px 0;
	}

	.author-box {
		background: var(--bg-secondary);
		padding: 2rem;
		border-radius: 8px;
		margin: 2rem 0;
	}

	.author-info h3 {
		margin: 0 0 0.5rem;
	}

	.role {
		color: var(--accent);
		font-weight: 600;
		margin: 0 0 1rem;
	}

	.links {
		display: flex;
		gap: 1.5rem;
		margin-top: 1.5rem;
		flex-wrap: wrap;
	}

	.link {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		color: var(--accent);
		text-decoration: none;
		font-weight: 500;
	}

	.link:hover {
		text-decoration: underline;
	}

	.icon {
		font-size: 1.2rem;
	}

	.tech-grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
		gap: 1.5rem;
		margin: 2rem 0;
	}

	.tech-category {
		background: var(--bg-secondary);
		padding: 1.5rem;
		border-radius: 6px;
	}

	.tech-category h3 {
		margin: 0 0 1rem;
		font-size: 1.1rem;
	}

	.tech-category ul {
		margin: 0;
		padding-left: 1.25rem;
	}

	.projects {
		display: grid;
		gap: 1.5rem;
		margin: 2rem 0;
	}

	.project {
		background: var(--bg-secondary);
		padding: 1.5rem;
		border-radius: 6px;
	}

	.project h3 {
		margin: 0 0 0.75rem;
	}

	.project p {
		margin: 0 0 1rem;
	}

	.cta-buttons {
		display: flex;
		gap: 1rem;
		margin-top: 2rem;
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
