<script lang="ts">
	import { onMount } from 'svelte';
	
	let isDark = false;
	let initialized = false;
	
	onMount(() => {
		// Check localStorage first
		const stored = localStorage.getItem('theme');
		if (stored) {
			isDark = stored === 'dark';
		} else {
			// Fall back to system preference
			isDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
		}
		
		applyTheme(isDark);
		initialized = true;
		
		// Listen for system theme changes
		const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
		const handleChange = (e: MediaQueryListEvent) => {
			if (!localStorage.getItem('theme')) {
				isDark = e.matches;
				applyTheme(isDark);
			}
		};
		
		mediaQuery.addEventListener('change', handleChange);
		
		return () => {
			mediaQuery.removeEventListener('change', handleChange);
		};
	});
	
	function toggle() {
		isDark = !isDark;
		applyTheme(isDark);
		localStorage.setItem('theme', isDark ? 'dark' : 'light');
	}
	
	function applyTheme(dark: boolean) {
		if (dark) {
			document.documentElement.classList.add('dark');
			document.documentElement.style.colorScheme = 'dark';
		} else {
			document.documentElement.classList.remove('dark');
			document.documentElement.style.colorScheme = 'light';
		}
	}
</script>

<button class="theme-toggle" on:click={toggle} aria-label="Toggle dark mode" title={isDark ? 'Switch to light mode' : 'Switch to dark mode'}>
	{#if isDark}
		<!-- Sun icon -->
		<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<circle cx="12" cy="12" r="5"></circle>
			<line x1="12" y1="1" x2="12" y2="3"></line>
			<line x1="12" y1="21" x2="12" y2="23"></line>
			<line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
			<line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
			<line x1="1" y1="12" x2="3" y2="12"></line>
			<line x1="21" y1="12" x2="23" y2="12"></line>
			<line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
			<line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
		</svg>
	{:else}
		<!-- Moon icon -->
		<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
		</svg>
	{/if}
</button>

<style>
	.theme-toggle {
		background: transparent;
		border: 1px solid var(--border);
		border-radius: 6px;
		padding: 0.5rem;
		cursor: pointer;
		color: var(--text-secondary);
		transition: all 0.2s;
		display: flex;
		align-items: center;
		justify-content: center;
	}
	
	.theme-toggle:hover {
		background: var(--bg-secondary);
		color: var(--accent);
		border-color: var(--accent);
	}
	
	.theme-toggle:active {
		transform: scale(0.95);
	}
	
	:global(.dark) {
		--bg-primary: #0d1117;
		--bg-secondary: #161b22;
		--text-primary: #f0f6fc;
		--text-secondary: #8b949e;
		--text-muted: #6e7681;
		--border: #30363d;
		--accent: #58a6ff;
		--accent-hover: #79c0ff;
	}
	
	:global(.dark) .copy-code-button {
		background: rgba(255, 255, 255, 0.05);
	}
	
	:global(.dark) .copy-code-button:hover {
		background: rgba(255, 255, 255, 0.1);
	}
	
	:global(.dark) pre {
		background: #161b22;
	}
	
	:global(.dark) .toc {
		background: var(--bg-secondary);
		border-color: var(--border);
	}
	
	:global(.dark) .install-card {
		background: var(--bg-secondary);
	}
	
	:global(.dark) .feature-card,
	:global(.dark) .example-card,
	:global(.dark) .related-card {
		background: var(--bg-secondary);
	}
	
	:global(.dark) .cta-box {
		background: var(--bg-secondary);
	}
</style>
