// Svelte action to automatically add copy buttons to all code blocks
export function addCopyButtons(node: HTMLElement) {
	const observer = new MutationObserver(() => {
		addButtonsToCodeBlocks(node);
	});
	
	// Initial setup
	addButtonsToCodeBlocks(node);
	
	// Watch for dynamically added content
	observer.observe(node, {
		childList: true,
		subtree: true
	});
	
	return {
		destroy() {
			observer.disconnect();
		}
	};
}

function addButtonsToCodeBlocks(container: HTMLElement) {
	const pres = container.querySelectorAll('pre');
	
	pres.forEach((pre) => {
		// Skip if already has a copy button
		if (pre.querySelector('.copy-btn')) return;
		
		const code = pre.querySelector('code');
		if (!code) return;
		
		// Create copy button
		const button = document.createElement('button');
		button.className = 'copy-code-button';
		button.setAttribute('aria-label', 'Copy code');
		button.innerHTML = `
			<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
				<rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
				<path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
			</svg>
			<span class="tooltip">Copy</span>
		`;
		
		button.addEventListener('click', async () => {
			try {
				await navigator.clipboard.writeText(code.textContent || '');
				button.classList.add('copied');
				button.innerHTML = `
					<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
						<polyline points="20 6 9 17 4 12"></polyline>
					</svg>
					<span class="tooltip">Copied!</span>
				`;
				
				setTimeout(() => {
					button.classList.remove('copied');
					button.innerHTML = `
						<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
							<rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
							<path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
						</svg>
						<span class="tooltip">Copy</span>
					`;
				}, 2000);
			} catch (err) {
				console.error('Failed to copy:', err);
			}
		});
		
		pre.appendChild(button);
	});
}
