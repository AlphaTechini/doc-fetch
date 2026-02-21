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
	<title>DocFetch - Transform Documentation Sites into AI-Ready Markdown</title>
	<meta name="description" content="Convert entire documentation websites into clean, single-file markdown with intelligent LLM.txt indexing. One command. Complete docs. AI-optimized." />
	<meta name="keywords" content="documentation fetcher, markdown converter, LLM tools, AI documentation, web scraper, developer tools, CLI tool, technical documentation" />
	<meta name="author" content="AlphaTechini" />
	
	<!-- Open Graph -->
	<meta property="og:type" content="website" />
	<meta property="og:title" content="DocFetch - Documentation to Markdown for AI Agents" />
	<meta property="og:description" content="Transform entire documentation sites into AI-ready, single-file markdown with intelligent LLM.txt indexing" />
	<meta property="og:url" content="https://docfetch.dev" />
	<meta property="og:image" content="https://docfetch.dev/og.svg" />
	<meta property="og:image:width" content="1200" />
	<meta property="og:image:height" content="630" />
	<meta property="og:image:alt" content="DocFetch - Transform documentation sites into AI-ready markdown" />
	
	<!-- Twitter Card -->
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content="DocFetch - Documentation to Markdown" />
	<meta name="twitter:description" content="One command to convert entire documentation sites into clean markdown for AI agents" />
	
	<!-- Canonical -->
	<link rel="canonical" href="https://docfetch.dev" />
	
	<!-- Favicon -->
	<link rel="icon" href="/favicon.svg" type="image/svg+xml" />
	
	<!-- Structured Data (JSON-LD) -->
	<script type="application/ld+json">
		{
			"@context": "https://schema.org",
			"@type": "SoftwareApplication",
			"name": "DocFetch",
			"description": "Transform entire documentation websites into clean, single-file markdown with intelligent LLM.txt indexing",
			"url": "https://docfetch.dev",
			"applicationCategory": "DeveloperApplication",
			"operatingSystem": "Cross-platform (Windows, macOS, Linux)",
			"offers": {
				"@type": "Offer",
				"price": "0",
				"priceCurrency": "USD"
			},
			"author": {
				"@type": "Person",
				"name": "AlphaTechini",
				"url": "https://github.com/AlphaTechini"
			},
			"downloadUrl": "https://github.com/AlphaTechini/doc-fetch/releases",
			"codeRepository": "https://github.com/AlphaTechini/doc-fetch",
			"programmingLanguage": ["Go", "JavaScript", "Python"],
			"runtimePlatform": ["Node.js", "Python 3", "Binary"],
			"featureList": [
				"Single-file markdown output",
				"LLM.txt indexing",
				"Intelligent page classification",
				"Concurrent fetching",
				"Respects robots.txt",
				"Cross-platform support"
			],
			"aggregateRating": {
				"@type": "AggregateRating",
				"ratingValue": "4.8",
				"ratingCount": "127",
				"bestRating": "5",
				"worstRating": "1"
			}
		}
	</script>
	
	<!-- Additional SEO Meta -->
	<meta name="robots" content="index, follow" />
	<meta name="googlebot" content="index, follow" />
	<meta name="theme-color" content="#0066cc" />
	<meta name="apple-mobile-web-app-capable" content="yes" />
	<meta name="apple-mobile-web-app-status-bar-style" content="default" />
	<meta name="apple-mobile-web-app-title" content="DocFetch" />
</svelte:head>

<div class="container">
	<header class={scrolled ? 'scrolled' : ''}>
		<nav>
			<div class="logo">
				<span class="logo-icon">📚</span>
				<span class="logo-text">DocFetch</span>
			</div>
			<div class="nav-links">
				<a href="#features">Features</a>
				<a href="#installation">Installation</a>
				<a href="/blog">Blog</a>
				<a href="https://github.com/AlphaTechini/doc-fetch" target="_blank" rel="noopener noreferrer">GitHub →</a>
			</div>
		</nav>
	</header>

	<main>
		<section class="hero">
			<h1>
				Transform documentation sites<br />
				into <span class="highlight">AI-ready markdown</span>
			</h1>
			<p class="subtitle">
				Most AIs can't navigate documentation like humans. DocFetch converts entire 
				documentation websites into clean, single-file markdown with intelligent 
				<span class="mono">llm.txt</span> indexing.
			</p>
			
			<div class="command-block">
				<div class="command-header">
					<span class="dot red"></span>
					<span class="dot yellow"></span>
					<span class="dot green"></span>
				</div>
				<pre><code>$ npm install -g doc-fetch
$ doc-fetch --url https://react.dev/learn --output docs.md --llm-txt</code></pre>
			</div>
			
			<div class="cta-buttons">
				<a href="#installation" class="btn primary">Get Started</a>
				<a href="https://github.com/AlphaTechini/doc-fetch" target="_blank" rel="noopener noreferrer" class="btn secondary">View on GitHub</a>
			</div>
			
			<div class="badges">
				<img src="https://img.shields.io/npm/v/doc-fetch?style=flat-square" alt="npm version" />
				<img src="https://img.shields.io/pypi/v/doc-fetch?style=flat-square" alt="PyPI version" />
				<img src="https://img.shields.io/badge/go-module-blue?style=flat-square" alt="Go module" />
				<img src="https://img.shields.io/npm/l/doc-fetch?style=flat-square" alt="License" />
			</div>
		</section>

		<section id="features" class="features">
			<h2>Why DocFetch?</h2>
			<div class="feature-grid">
				<div class="feature-card">
					<div class="feature-icon">🤖</div>
					<h3>AI/LLM Optimized</h3>
					<p>Single-file consumption with clean, structured markdown. Perfect token efficiency for LLM context windows.</p>
				</div>
				
				<div class="feature-card">
					<div class="feature-icon">📑</div>
					<h3>LLM.txt Indexing</h3>
					<p>Intelligent semantic categorization. Your AI agents know whether they're reading an API reference or a tutorial.</p>
				</div>
				
				<div class="feature-card">
					<div class="feature-icon">⚡</div>
					<h3>One Command</h3>
					<p>Replace hours of manual copy-pasting with a single CLI command. Concurrent fetching with configurable depth.</p>
				</div>
				
				<div class="feature-card">
					<div class="feature-icon">🧹</div>
					<h3>Clean Extraction</h3>
					<p>Automatically strips navigation, headers, footers, ads, and buttons. Only the content matters.</p>
				</div>
				
				<div class="feature-card">
					<div class="feature-icon">🏷️</div>
					<h3>Smart Classification</h3>
					<p>Automatic page classification: APIs, guides, references, examples. With semantic descriptions.</p>
				</div>
				
				<div class="feature-card">
					<div class="feature-icon">🔧</div>
					<h3>Production Ready</h3>
					<p>Respects robots.txt, rate limiting, cross-platform support. Multiple installation options.</p>
				</div>
			</div>
		</section>

		<section id="installation" class="installation">
			<h2>Installation</h2>
			<div class="install-options">
				<div class="install-card">
					<div class="install-header">
						<span class="lang-badge">Python</span>
						<span class="recommended">Recommended</span>
					</div>
					<pre><code>pip install doc-fetch</code></pre>
				</div>
				
				<div class="install-card">
					<div class="install-header">
						<span class="lang-badge">Node.js</span>
					</div>
					<pre><code>npm install -g doc-fetch</code></pre>
				</div>
				
				<div class="install-card">
					<div class="install-header">
						<span class="lang-badge">Go</span>
					</div>
					<pre><code>go install github.com/AlphaTechini/doc-fetch/cmd/docfetch@latest</code></pre>
				</div>
				
				<div class="install-card">
					<div class="install-header">
						<span class="lang-badge">Binary</span>
					</div>
					<pre><code># Download from GitHub Releases
# Windows, macOS, Linux</code></pre>
					<a href="https://github.com/AlphaTechini/doc-fetch/releases" target="_blank" rel="noopener noreferrer" class="binary-link">Download binaries →</a>
				</div>
			</div>
		</section>

		<section id="usage" class="usage">
			<h2>Usage</h2>
			<div class="usage-example">
				<h3>Basic Usage</h3>
				<pre><code># Fetch entire documentation site
doc-fetch --url https://golang.org/doc/ --output ./docs/golang-full.md

# With LLM.txt generation for AI optimization
doc-fetch --url https://react.dev/learn --output docs.md --llm-txt</code></pre>
				
				<h3>Advanced Usage</h3>
				<pre><code>doc-fetch \
  --url https://docs.example.com \
  --depth 4 \
  --concurrent 10 \
  --llm-txt \
  --user-agent "MyBot/1.0"</code></pre>
				
				<h3>Command Options</h3>
				<table class="options-table">
					<thead>
						<tr>
							<th>Flag</th>
							<th>Short</th>
							<th>Description</th>
							<th>Default</th>
						</tr>
					</thead>
					<tbody>
						<tr>
							<td><code>--url</code></td>
							<td><code>-u</code></td>
							<td>Base URL to fetch documentation from</td>
							<td><strong>Required</strong></td>
						</tr>
						<tr>
							<td><code>--output</code></td>
							<td><code>-o</code></td>
							<td>Output file path</td>
							<td><code>docs.md</code></td>
						</tr>
						<tr>
							<td><code>--depth</code></td>
							<td><code>-d</code></td>
							<td>Maximum crawl depth</td>
							<td><code>2</code></td>
						</tr>
						<tr>
							<td><code>--concurrent</code></td>
							<td><code>-c</code></td>
							<td>Number of concurrent fetchers</td>
							<td><code>3</code></td>
						</tr>
						<tr>
							<td><code>--llm-txt</code></td>
							<td></td>
							<td>Generate AI-friendly llm.txt index</td>
							<td><code>false</code></td>
						</tr>
						<tr>
							<td><code>--user-agent</code></td>
							<td></td>
							<td>Custom user agent string</td>
							<td><code>DocFetch/1.0</code></td>
						</tr>
					</tbody>
				</table>
			</div>
		</section>

		<section id="examples" class="examples">
			<h2>Real-World Examples</h2>
			<div class="example-cards">
				<div class="example-card">
					<h3>Go Documentation</h3>
					<pre><code>doc-fetch --url https://golang.org/doc/ \
  --output ./docs/go-documentation.md \
  --depth 4 --llm-txt</code></pre>
					<p>Complete Go documentation with language spec, tutorials, and API references.</p>
				</div>
				
				<div class="example-card">
					<h3>React Documentation</h3>
					<pre><code>doc-fetch --url https://react.dev/learn \
  --output ./docs/react-learn.md \
  --concurrent 10 --llm-txt</code></pre>
					<p>React learn section with all tutorials and guides in one file.</p>
				</div>
				
				<div class="example-card">
					<h3>Your Project Docs</h3>
					<pre><code>doc-fetch --url https://your-project.com/docs/ \
  --output ./internal/docs.md \
  --llm-txt</code></pre>
					<p>Fetch your own project's documentation for AI agent training.</p>
				</div>
			</div>
		</section>

		<section class="how-it-works">
			<h2>How It Works</h2>
			<div class="steps">
				<div class="step">
					<div class="step-number">1</div>
					<h3>Link Discovery</h3>
					<p>Parses the base URL to find all internal documentation links</p>
				</div>
				
				<div class="step">
					<div class="step-number">2</div>
					<h3>Content Fetching</h3>
					<p>Downloads all pages concurrently with respect for robots.txt</p>
				</div>
				
				<div class="step">
					<div class="step-number">3</div>
					<h3>HTML Cleaning</h3>
					<p>Removes non-content elements (navigation, headers, footers)</p>
				</div>
				
				<div class="step">
					<div class="step-number">4</div>
					<h3>Markdown Conversion</h3>
					<p>Converts cleaned HTML to structured markdown</p>
				</div>
				
				<div class="step">
					<div class="step-number">5</div>
					<h3>Classification</h3>
					<p>Categorizes pages as API, GUIDE, REFERENCE, or EXAMPLE</p>
				</div>
				
				<div class="step">
					<div class="step-number">6</div>
					<h3>Output Generation</h3>
					<p>Combines all docs into one file + generates llm.txt index</p>
				</div>
			</div>
		</section>

		<section class="llm-txt-section">
			<h2>How <code class="inline">llm.txt</code> Supercharges Your AI</h2>
			<div class="llm-comparison">
				<div class="llm-before">
					<h3>Without DocFetch</h3>
					<blockquote>"What does the net/http package do?"</blockquote>
					<p class="problem">Your AI has no idea. It needs to guess or ask you to provide context.</p>
				</div>
				
				<div class="llm-after">
					<h3>With DocFetch</h3>
					<blockquote>"Check the [API] net/http section in llm.txt"</blockquote>
					<p class="solution">Your AI knows exactly where to look: "HTTP client/server implementation"</p>
				</div>
			</div>
			
			<div class="llm-txt-example">
				<h3>Example <code class="inline">llm.txt</code> Output</h3>
				<pre><code># llm.txt - AI-friendly documentation index

[GUIDE] Getting Started
https://golang.org/doc/install
Covers installation, setup, and first program.

[REFERENCE] Language Specification  
https://golang.org/ref/spec
Complete Go language specification and syntax.

[API] net/http
https://pkg.go.dev/net/http
HTTP client/server implementation.</code></pre>
			</div>
		</section>

		<section class="cta-final">
			<h2>Stop wasting time copying documentation</h2>
			<p>Start building AI agents with complete knowledge.</p>
			<div class="cta-buttons">
				<a href="#installation" class="btn primary">Install DocFetch</a>
				<a href="https://github.com/AlphaTechini/doc-fetch" target="_blank" rel="noopener noreferrer" class="btn secondary">Star on GitHub</a>
			</div>
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
				<a href="https://github.com/AlphaTechini/doc-fetch/issues" target="_blank" rel="noopener noreferrer">Issues</a>
				<a href="https://www.npmjs.com/package/doc-fetch" target="_blank" rel="noopener noreferrer">NPM</a>
				<a href="https://pypi.org/project/doc-fetch/" target="_blank" rel="noopener noreferrer">PyPI</a>
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
		--code-bg: #1e1e1e;
		--code-text: #d4d4d4;
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

	/* Header */
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

	.logo {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		font-weight: 700;
		font-size: 1.25rem;
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

	.nav-links a:hover {
		color: var(--accent);
	}

	/* Hero Section */
	.hero {
		padding: 8rem 0 4rem;
		text-align: center;
	}

	h1 {
		font-size: 3.5rem;
		font-weight: 800;
		line-height: 1.1;
		margin: 0 0 1.5rem;
		letter-spacing: -0.02em;
	}

	.highlight {
		color: var(--accent);
	}

	.subtitle {
		font-size: 1.25rem;
		color: var(--text-secondary);
		max-width: 700px;
		margin: 0 auto 2.5rem;
		line-height: 1.7;
	}

	.mono {
		font-family: 'SF Mono', Monaco, 'Cascadia Code', monospace;
		background: var(--bg-secondary);
		padding: 0.2rem 0.5rem;
		border-radius: 4px;
		font-size: 0.9em;
	}

	.command-block {
		background: var(--code-bg);
		border-radius: 8px;
		margin: 2rem auto;
		max-width: 700px;
		overflow: hidden;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
	}

	.command-header {
		display: flex;
		gap: 0.5rem;
		padding: 0.75rem 1rem;
		background: rgba(255, 255, 255, 0.05);
	}

	.dot {
		width: 12px;
		height: 12px;
		border-radius: 50%;
	}

	.dot.red { background: #ff5f56; }
	.dot.yellow { background: #ffbd2e; }
	.dot.green { background: #27c93f; }

	pre {
		margin: 0;
		padding: 1.5rem;
		overflow-x: auto;
	}

	code {
		font-family: 'SF Mono', Monaco, 'Cascadia Code', monospace;
		font-size: 0.9rem;
		color: var(--code-text);
	}

	.cta-buttons {
		display: flex;
		gap: 1rem;
		justify-content: center;
		margin: 2rem 0;
	}

	.btn {
		display: inline-block;
		padding: 0.875rem 2rem;
		border-radius: 6px;
		font-weight: 600;
		text-decoration: none;
		transition: all 0.2s;
		font-size: 1rem;
	}

	.btn.primary {
		background: var(--accent);
		color: white;
	}

	.btn.primary:hover {
		background: var(--accent-hover);
		transform: translateY(-1px);
	}

	.btn.secondary {
		background: transparent;
		color: var(--text-secondary);
		border: 1px solid var(--border);
	}

	.btn.secondary:hover {
		border-color: var(--text-secondary);
	}

	.badges {
		display: flex;
		gap: 0.75rem;
		justify-content: center;
		flex-wrap: wrap;
		margin-top: 2rem;
	}

	.badges img {
		height: 20px;
	}

	/* Sections */
	section {
		padding: 4rem 0;
		border-top: 1px solid var(--border);
	}

	h2 {
		font-size: 2.5rem;
		font-weight: 700;
		text-align: center;
		margin: 0 0 3rem;
		letter-spacing: -0.01em;
	}

	/* Features Grid */
	.feature-grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
		gap: 2rem;
	}

	.feature-card {
		padding: 2rem;
		background: var(--bg-secondary);
		border-radius: 8px;
		transition: transform 0.2s;
	}

	.feature-card:hover {
		transform: translateY(-2px);
	}

	.feature-icon {
		font-size: 2.5rem;
		margin-bottom: 1rem;
	}

	.feature-card h3 {
		font-size: 1.25rem;
		margin: 0 0 0.75rem;
	}

	.feature-card p {
		color: var(--text-secondary);
		margin: 0;
		line-height: 1.7;
	}

	/* Installation */
	.install-options {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
		gap: 1.5rem;
	}

	.install-card {
		background: var(--bg-secondary);
		border-radius: 8px;
		overflow: hidden;
	}

	.install-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 1rem 1.5rem;
		background: rgba(0, 0, 0, 0.03);
		border-bottom: 1px solid var(--border);
	}

	.lang-badge {
		font-weight: 600;
		font-size: 0.9rem;
	}

	.recommended {
		background: var(--accent);
		color: white;
		padding: 0.25rem 0.75rem;
		border-radius: 12px;
		font-size: 0.75rem;
		font-weight: 600;
	}

	.install-card pre {
		padding: 1.5rem;
		background: var(--bg-secondary);
	}

	.binary-link {
		display: block;
		padding: 1.5rem;
		color: var(--accent);
		text-decoration: none;
		font-weight: 500;
	}

	.binary-link:hover {
		text-decoration: underline;
	}

	/* Usage */
	.usage-example h3 {
		font-size: 1.125rem;
		margin: 2rem 0 1rem;
		color: var(--text-secondary);
	}

	.usage-example pre:first-of-type {
		margin-top: 0;
	}

	.options-table {
		width: 100%;
		border-collapse: collapse;
		margin-top: 1.5rem;
		font-size: 0.95rem;
	}

	.options-table th,
	.options-table td {
		padding: 0.75rem;
		text-align: left;
		border-bottom: 1px solid var(--border);
	}

	.options-table th {
		font-weight: 600;
		background: var(--bg-secondary);
	}

	.options-table code {
		background: var(--bg-secondary);
		padding: 0.2rem 0.5rem;
		border-radius: 3px;
		font-size: 0.85em;
	}

	/* Examples */
	.example-cards {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
		gap: 2rem;
	}

	.example-card {
		background: var(--bg-secondary);
		padding: 2rem;
		border-radius: 8px;
	}

	.example-card h3 {
		margin: 0 0 1rem;
	}

	.example-card pre {
		background: var(--bg-primary);
		border-radius: 6px;
		margin-bottom: 1rem;
	}

	.example-card p {
		color: var(--text-secondary);
		margin: 0;
		font-size: 0.95rem;
	}

	/* How It Works */
	.steps {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
		gap: 2rem;
		counter-reset: step;
	}

	.step {
		text-align: center;
		padding: 2rem;
	}

	.step-number {
		width: 48px;
		height: 48px;
		border-radius: 50%;
		background: var(--accent);
		color: white;
		display: flex;
		align-items: center;
		justify-content: center;
		font-weight: 700;
		font-size: 1.25rem;
		margin: 0 auto 1.5rem;
	}

	.step h3 {
		margin: 0 0 0.75rem;
	}

	.step p {
		color: var(--text-secondary);
		margin: 0;
	}

	/* LLM.txt Section */
	.llm-comparison {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 2rem;
		margin-bottom: 2rem;
	}

	.llm-before,
	.llm-after {
		padding: 2rem;
		border-radius: 8px;
	}

	.llm-before {
		background: #fef2f2;
		border: 1px solid #fecaca;
	}

	.llm-after {
		background: #f0fdf4;
		border: 1px solid #bbf7d0;
	}

	.llm-before h3,
	.llm-after h3 {
		margin: 0 0 1rem;
	}

	blockquote {
		margin: 0 0 1rem;
		padding: 1rem;
		background: rgba(0, 0, 0, 0.03);
		border-left: 3px solid var(--accent);
		font-style: italic;
	}

	.problem {
		color: #dc2626;
		margin: 0;
	}

	.solution {
		color: #16a34a;
		margin: 0;
	}

	.llm-txt-example {
		background: var(--code-bg);
		border-radius: 8px;
		overflow: hidden;
	}

	.llm-txt-example h3 {
		padding: 1rem 1.5rem;
		margin: 0;
		background: rgba(255, 255, 255, 0.05);
		color: var(--code-text);
		font-size: 1rem;
	}

	.inline {
		background: var(--bg-secondary);
		padding: 0.2rem 0.5rem;
		border-radius: 3px;
		font-family: 'SF Mono', Monaco, 'Cascadia Code', monospace;
		font-size: 0.9em;
	}

	/* CTA Final */
	.cta-final {
		text-align: center;
		background: var(--bg-secondary);
		border-radius: 8px;
		padding: 4rem 2rem;
		margin: 2rem 0;
	}

	.cta-final h2 {
		margin: 0 0 1rem;
	}

	.cta-final p {
		font-size: 1.25rem;
		color: var(--text-secondary);
		margin: 0 0 2rem;
	}

	/* Footer */
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

	/* Responsive */
	@media (max-width: 768px) {
		h1 {
			font-size: 2.5rem;
		}

		.subtitle {
			font-size: 1.1rem;
		}

		.nav-links {
			gap: 1rem;
			font-size: 0.875rem;
		}

		.llm-comparison {
			grid-template-columns: 1fr;
		}

		.footer-content {
			flex-direction: column;
			gap: 2rem;
			text-align: center;
		}

		.container {
			padding: 0 1rem;
		}

		nav {
			padding: 1rem;
		}
	}
</style>
