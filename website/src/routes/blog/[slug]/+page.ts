// SSR Load function for proper SEO
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
	const posts = {
		'convert-docs-to-markdown-for-llm': {
			slug: 'convert-docs-to-markdown-for-llm',
			title: 'How to Convert Documentation to Markdown for LLMs (Complete Guide)',
			excerpt: 'Step-by-step guide: Transform entire documentation websites into clean, AI-ready markdown. Includes tools, techniques, and best practices for optimal LLM context.',
			date: '2026-02-21',
			author: 'AlphaTechini',
			readTime: '8 min read',
			tags: ['Tutorial', 'LLM', 'Markdown'],
			category: 'rag',
			subcategory: 'context-preparation',
			modifiedDate: '2026-02-21',
			relatedPosts: [
				'llm-txt-index-guide',
				'ai-agent-documentation-problem',
				'best-practices-rag-context-preparation'
			],
			faqs: [
				{
					question: 'What is the best way to convert documentation to markdown?',
					answer: 'Automated tools like DocFetch are most efficient. They crawl entire documentation sites, extract clean content, and generate structured markdown with semantic indexing (llm.txt).'
				},
				{
					question: 'Why convert documentation to markdown for LLMs?',
					answer: 'LLMs need complete context in a single prompt. Markdown provides clean, structured text without HTML bloat, navigation, or ads that waste tokens.'
				},
				{
					question: 'What is llm.txt and how does it help?',
					answer: 'llm.txt is a semantic index file that categorizes documentation sections (GUIDE, API, TUTORIAL) with descriptions. It helps AI agents navigate large documentation efficiently.'
				},
				{
					question: 'Can I automate documentation conversion?',
					answer: 'Yes. Tools like DocFetch can automatically fetch, clean, and convert entire documentation sites with one command. Set up cron jobs for regular updates.'
				}
			],
			content: '' // Will be loaded dynamically
		}
	};
	
	const post = posts[params.slug as keyof typeof posts];
	
	if (!post) {
		throw new Error('Post not found');
	}
	
	return {
		post
	};
};
