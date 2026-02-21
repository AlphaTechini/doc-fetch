<script lang="ts">
	import { onMount, tick } from 'svelte';
	
	export let contentSelector = '.content';
	export let headingSelector = 'h2, h3';
	export let title = 'On this page';
	
	let headings: Array<{
		id: string;
		text: string;
		level: number;
	}> = [];
	
	let activeId = '';
	let observer: IntersectionObserver | null = null;
	
	onMount(() => {
		tick().then(() => {
			extractHeadings();
			setupObserver();
		});
		
		return () => {
			if (observer) observer.disconnect();
		};
	});
	
	function extractHeadings() {
		const content = document.querySelector(contentSelector);
		if (!content) return;
		
		const elements = content.querySelectorAll(headingSelector);
		headings = Array.from(elements).map((el) => {
			// Generate ID if not present
			if (!el.id) {
				el.id = el.textContent
					?.toLowerCase()
					.replace(/[^a-z0-9]+/g, '-')
					.replace(/(^-|-$)/g, '') || '';
			}
			
			return {
				id: el.id,
				text: el.textContent || '',
				level: el.tagName === 'H2' ? 2 : 3
			};
		});
	}
	
	function setupObserver() {
		const options = {
			root: null,
			rootMargin: '-20% 0px -80% 0px',
			threshold: 0
		};
		
		observer = new IntersectionObserver((entries) => {
			entries.forEach((entry) => {
				if (entry.isIntersecting) {
					activeId = entry.target.id;
				}
			});
		}, options);
		
		headings.forEach((heading) => {
			const el = document.getElementById(heading.id);
			if (el) observer?.observe(el);
		});
	}
	
	function scrollToHeading(id: string, event: MouseEvent) {
		event.preventDefault();
		const el = document.getElementById(id);
		if (el) {
			el.scrollIntoView({ behavior: 'smooth', block: 'start' });
			history.pushState(null, '', `#${id}`);
			activeId = id;
		}
	}
</script>

{#if headings.length > 0}
	<nav class="toc" aria-label={title}>
		<h3 class="toc-title">{title}</h3>
		<ul class="toc-list">
			{#each headings as heading}
				<li class="toc-item level-{heading.level}">
					<a
						href={`#${heading.id}`}
						class:active={activeId === heading.id}
						on:click={(e) => scrollToHeading(heading.id, e)}
					>
						{heading.text}
					</a>
				</li>
			{/each}
		</ul>
	</nav>
{/if}

<style>
	.toc {
		position: sticky;
		top: 6rem;
		background: var(--bg-secondary);
		padding: 1.5rem;
		border-radius: 8px;
		border: 1px solid var(--border);
		max-height: calc(100vh - 8rem);
		overflow-y: auto;
	}
	
	.toc-title {
		font-size: 0.875rem;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		color: var(--text-muted);
		margin: 0 0 1rem;
	}
	
	.toc-list {
		list-style: none;
		padding: 0;
		margin: 0;
	}
	
	.toc-item {
		margin-bottom: 0.5rem;
	}
	
	.toc-item a {
		display: block;
		padding: 0.5rem 0.75rem;
		color: var(--text-secondary);
		text-decoration: none;
		font-size: 0.9rem;
		line-height: 1.5;
		border-radius: 4px;
		transition: all 0.2s;
		border-left: 2px solid transparent;
	}
	
	.toc-item a:hover {
		background: rgba(0, 102, 204, 0.05);
		color: var(--accent);
	}
	
	.toc-item a.active {
		background: rgba(0, 102, 204, 0.1);
		color: var(--accent);
		border-left-color: var(--accent);
		font-weight: 600;
	}
	
	.toc-item.level-3 a {
		padding-left: 1.5rem;
		font-size: 0.85rem;
	}
	
	/* Scrollbar styling */
	.toc::-webkit-scrollbar {
		width: 6px;
	}
	
	.toc::-webkit-scrollbar-track {
		background: transparent;
	}
	
	.toc::-webkit-scrollbar-thumb {
		background: var(--border);
		border-radius: 3px;
	}
	
	.toc::-webkit-scrollbar-thumb:hover {
		background: var(--text-muted);
	}
	
	@media (max-width: 1200px) {
		.toc {
			display: none;
		}
	}
</style>
