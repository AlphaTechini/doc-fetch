package fetcher

import (
	"strings"
)

// ConvertHTMLToMarkdown converts HTML content to clean markdown
func ConvertHTMLToMarkdown(htmlContent string) string {
	if htmlContent == "" {
		return ""
	}
	
	// Basic HTML to markdown conversion
	markdownContent := basicHTMLToMarkdown(htmlContent)
	
	return strings.TrimSpace(markdownContent)
}

// basicHTMLToMarkdown provides basic HTML to markdown conversion
func basicHTMLToMarkdown(html string) string {
	// Replace common HTML tags with markdown equivalents
	replacements := map[string]string{
		"<h1>": "# ",
		"</h1>": "\n\n",
		"<h2>": "## ",
		"</h2>": "\n\n", 
		"<h3>": "### ",
		"</h3>": "\n\n",
		"<h4>": "#### ",
		"</h4>": "\n\n",
		"<h5>": "##### ",
		"</h5>": "\n\n",
		"<h6>": "###### ",
		"</h6>": "\n\n",
		"<p>": "",
		"</p>": "\n\n",
		"<br>": "\n",
		"<br/>": "\n",
		"<strong>": "**",
		"</strong>": "**",
		"<b>": "**",
		"</b>": "**",
		"<em>": "*",
		"</em>": "*",
		"<i>": "*",
		"</i>": "*",
		"<code>": "`",
		"</code>": "`",
		"<pre>": "```",
		"</pre>": "```",
		"<ul>": "",
		"</ul>": "\n",
		"<ol>": "",
		"</ol>": "\n",
		"<li>": "- ",
		"</li>": "\n",
		"<blockquote>": "> ",
		"</blockquote>": "\n\n",
	}
	
	result := html
	for htmlTag, mdReplacement := range replacements {
		result = strings.ReplaceAll(result, htmlTag, mdReplacement)
	}
	
	// Clean up extra whitespace
	result = strings.ReplaceAll(result, "\n\n\n", "\n\n")
	result = strings.TrimSpace(result)
	
	return result
}