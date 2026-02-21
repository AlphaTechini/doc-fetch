<script lang="ts">
	import { onMount } from 'svelte';
	
	export let codeElement: HTMLElement | null = null;
	let copied = false;
	let showTooltip = false;
	
	onMount(() => {
		if (codeElement) {
			const pre = codeElement.parentElement;
			if (pre) {
				pre.style.position = 'relative';
			}
		}
	});
	
	async function copyCode() {
		if (!codeElement) return;
		
		try {
			await navigator.clipboard.writeText(codeElement.textContent || '');
			copied = true;
			showTooltip = true;
			
			setTimeout(() => {
				copied = false;
				showTooltip = false;
			}, 2000);
		} catch (err) {
			console.error('Failed to copy:', err);
		}
	}
</script>

<button class="copy-btn" class:visible={copied || showTooltip} on:click={copyCode} aria-label="Copy code">
	{#if copied}
		<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<polyline points="20 6 9 17 4 12"></polyline>
		</svg>
		<span class="tooltip">Copied!</span>
	{:else}
		<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
			<path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
		</svg>
		<span class="tooltip">Copy</span>
	{/if}
</button>

<style>
	.copy-btn {
		position: absolute;
		top: 0.75rem;
		right: 0.75rem;
		background: rgba(255, 255, 255, 0.1);
		border: none;
		border-radius: 4px;
		padding: 0.5rem;
		cursor: pointer;
		color: #d4d4d4;
		transition: all 0.2s;
		opacity: 0;
		z-index: 10;
		display: flex;
		align-items: center;
		gap: 0.5rem;
		font-size: 0.875rem;
	}

	.copy-btn:hover,
	.copy-btn.visible {
		opacity: 1;
		background: rgba(255, 255, 255, 0.15);
	}

	.copy-btn:active {
		transform: scale(0.95);
	}

	.copy-btn svg {
		display: block;
	}

	.tooltip {
		background: rgba(0, 0, 0, 0.8);
		color: white;
		padding: 0.25rem 0.5rem;
		border-radius: 4px;
		font-size: 0.75rem;
		white-space: nowrap;
		pointer-events: none;
	}

	pre:hover .copy-btn {
		opacity: 1;
	}
</style>
