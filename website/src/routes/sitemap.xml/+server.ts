// Dynamic sitemap generation for SEO
import type { RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async () => {
	// In production, fetch from CMS/markdown files
	const baseUrl = 'https://docfetch.dev';
	
	const staticPages = [
		{ path: '/', priority: '1.0', changefreq: 'weekly' },
		{ path: '/#features', priority: '0.8', changefreq: 'monthly' },
		{ path: '/#installation', priority: '0.8', changefreq: 'monthly' },
		{ path: '/#usage', priority: '0.7', changefreq: 'monthly' },
		{ path: '/blog', priority: '0.9', changefreq: 'daily' }
	];
	
	const blogPosts = [
		{
			slug: 'convert-docs-to-markdown-for-llm',
			path: '/rag/context-preparation/convert-docs-to-markdown',
			lastmod: '2026-02-21',
			priority: '0.8',
			changefreq: 'weekly'
		},
		{
			slug: 'llm-txt-index-guide',
			path: '/rag/context-preparation/llm-txt-guide',
			lastmod: '2026-02-21',
			priority: '0.7',
			changefreq: 'weekly'
		},
		{
			slug: 'ai-agent-documentation-problem',
			path: '/llm-tools/documentation/ai-agents-cant-read-docs',
			lastmod: '2026-02-21',
			priority: '0.7',
			changefreq: 'weekly'
		}
	];
	
	const allPages = [...staticPages, ...blogPosts];
	
	const sitemap = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
        xmlns:news="http://www.google.com/schemas/sitemap-news/0.9"
        xmlns:xhtml="http://www.w3.org/1999/xhtml"
        xmlns:image="http://www.google.com/schemas/sitemap-image/1.1"
        xmlns:video="http://www.google.com/schemas/sitemap-video/1.1">
${allPages.map(page => `  <url>
    <loc>${baseUrl}${page.path}</loc>
    <lastmod>${page.lastmod || new Date().toISOString().split('T')[0]}</lastmod>
    <changefreq>${page.changefreq}</changefreq>
    <priority>${page.priority}</priority>
  </url>`).join('\n')}
</urlset>`;
	
	return new Response(sitemap, {
		headers: {
			'Content-Type': 'application/xml',
			'Cache-Control': 'max-age=0, must-revalidate'
		}
	});
};
