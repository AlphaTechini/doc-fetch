<script lang="ts">
	import { injectSpeedInsights } from '@vercel/speed-insights/sveltekit';
	
	let { children } = $props();
	
	// Inject Vercel Speed Insights for Core Web Vitals tracking
	injectSpeedInsights();
</script>

<svelte:head>
	<!-- Preload Critical Assets for LCP Optimization -->
	<link rel="preload" href="/favicon.svg" as="image" type="image/svg+xml" />
	
	<!-- DNS Prefetch for External Resources -->
	<link rel="dns-prefetch" href="https://img.shields.io" />
	<link rel="dns-prefetch" href="https://github.com" />
	<link rel="dns-prefetch" href="https://vitals.vercel-insights.com" />
	
	<!-- Preconnect to CDN -->
	<link rel="preconnect" href="https://img.shields.io" crossorigin />
	<link rel="preconnect" href="https://vitals.vercel-insights.com" crossorigin />
	
	<!-- Font Optimization (if using custom fonts) -->
	<!-- <link rel="preload" href="/fonts/inter-var.woff2" as="font" type="font/woff2" crossorigin> -->
	
	<!-- Theme Color -->
	<meta name="theme-color" content="#0066cc" />
	
	<!-- Performance Hints -->
	<link rel="prefetch" href="/blog" as="document" />
</svelte:head>

{@render children()}

<style>
	/* Critical CSS inlined for faster LCP */
	:global(body) {
		margin: 0;
		padding: 0;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
		background: #ffffff;
		color: #1a1a1a;
		line-height: 1.6;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
	}
	
	/* Prevent CLS by setting image max-width */
	:global(img) {
		max-width: 100%;
		height: auto;
		display: block;
	}
	
	/* Reduce layout shifts with contain */
	:global(.feature-card, .post-card, .example-card) {
		contain: layout style;
	}
</style>
