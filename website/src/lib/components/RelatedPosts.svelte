<script lang="ts">
	export let posts: Array<{
		slug: string;
		title: string;
		excerpt: string;
		readTime?: string;
	}> = [];
	
	// Generate placeholder OG images (in production, these would be real images)
	function getThumbnailUrl(slug: string): string {
		// Using gradient placeholders with text overlay
		const colors = [
			'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
			'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)',
			'linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)',
			'linear-gradient(135deg, #43e97b 0%, #38f9d7 100%)',
			'linear-gradient(135deg, #fa709a 0%, #fee140 100%)'
		];
		
		const index = slug.length % colors.length;
		return colors[index];
	}
	
	function getInitials(title: string): string {
		return title
			.split(' ')
			.filter(word => word.length > 0)
			.slice(0, 2)
			.map(word => word[0].toUpperCase())
			.join('');
	}
</script>

{#if posts.length > 0}
	<section class="related-posts-component">
		<h3>Related Articles</h3>
		<div class="related-grid-with-thumbs">
			{#each posts as post}
				<article class="related-card-with-thumb">
					<a href={`/blog/${post.slug}`} class="related-thumbnail" style="background: {getThumbnailUrl(post.slug)}">
						<span class="thumbnail-initials">{getInitials(post.title)}</span>
					</a>
					<div class="related-content">
						<h4>
							<a href={`/blog/${post.slug}`}>{post.title}</a>
						</h4>
						<p class="related-excerpt">{post.excerpt}</p>
						{#if post.readTime}
							<span class="related-read-time">{post.readTime}</span>
						{/if}
					</div>
				</article>
			{/each}
		</div>
	</section>
{/if}

<style>
	.related-posts-component {
		margin-bottom: 3rem;
		padding-bottom: 2rem;
		border-bottom: 1px solid var(--border);
	}
	
	.related-posts-component h3 {
		font-size: 1.4rem;
		margin-bottom: 1.5rem;
	}
	
	.related-grid-with-thumbs {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
		gap: 2rem;
	}
	
	.related-card-with-thumb {
		background: var(--bg-secondary);
		border-radius: 8px;
		overflow: hidden;
		transition: transform 0.2s, box-shadow 0.2s;
		display: flex;
		flex-direction: column;
	}
	
	.related-card-with-thumb:hover {
		transform: translateY(-4px);
		box-shadow: 0 8px 16px rgba(0, 0, 0, 0.08);
	}
	
	.related-thumbnail {
		display: flex;
		align-items: center;
		justify-content: center;
		height: 180px;
		text-decoration: none;
		position: relative;
		overflow: hidden;
	}
	
	.thumbnail-initials {
		font-size: 3rem;
		font-weight: 800;
		color: rgba(255, 255, 255, 0.9);
		text-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
	}
	
	.related-content {
		padding: 1.5rem;
		flex: 1;
		display: flex;
		flex-direction: column;
	}
	
	.related-content h4 {
		margin: 0 0 0.75rem;
		font-size: 1.1rem;
		line-height: 1.4;
	}
	
	.related-content h4 a {
		color: var(--text-primary);
		text-decoration: none;
	}
	
	.related-content h4 a:hover {
		color: var(--accent);
	}
	
	.related-excerpt {
		color: var(--text-secondary);
		font-size: 0.9rem;
		line-height: 1.6;
		margin: 0 0 1rem;
		flex: 1;
	}
	
	.related-read-time {
		font-size: 0.8rem;
		color: var(--text-muted);
		font-weight: 500;
		display: inline-flex;
		align-items: center;
		gap: 0.25rem;
	}
	
	@media (max-width: 768px) {
		.related-grid-with-thumbs {
			grid-template-columns: 1fr;
		}
	}
</style>
