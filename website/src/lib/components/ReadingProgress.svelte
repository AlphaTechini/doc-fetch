<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	
	let progress = 0;
	let animationFrame: number;
	
	onMount(() => {
		const updateProgress = () => {
			const scrollTop = window.scrollY;
			const docHeight = document.documentElement.scrollHeight - window.innerHeight;
			progress = docHeight > 0 ? (scrollTop / docHeight) * 100 : 0;
			animationFrame = requestAnimationFrame(updateProgress);
		};
		
		updateProgress();
		
		return () => {
			cancelAnimationFrame(animationFrame);
		};
	});
</script>

<div class="progress-bar" style="--progress: {progress}%"></div>

<style>
	.progress-bar {
		position: fixed;
		top: 0;
		left: 0;
		width: var(--progress);
		height: 3px;
		background: linear-gradient(90deg, var(--accent), var(--accent-hover));
		z-index: 9999;
		transition: width 0.1s ease-out;
	}
</style>
