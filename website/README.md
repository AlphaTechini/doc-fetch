# DocFetch Website

Landing page for [DocFetch](https://github.com/AlphaTechini/doc-fetch) - Transform documentation sites into AI-ready markdown.

## Tech Stack

- **Framework**: SvelteKit
- **Styling**: Vanilla CSS (no frameworks - classic, fast, no bloat)
- **Deployment**: Vercel/Netlify ready (adapter-auto)

## Development

```bash
cd website
npm install
npm run dev
```

Open http://localhost:5173

## Build

```bash
npm run build
npm run preview
```

## SEO Features

- ✅ Semantic HTML structure
- ✅ Meta tags (title, description, keywords)
- ✅ Open Graph tags for social sharing
- ✅ Twitter Card support
- ✅ Canonical URL
- ✅ Structured data (JSON-LD)
- ✅ Mobile responsive
- ✅ Fast load time (minimal dependencies)
- ✅ Accessible (proper heading hierarchy, alt text)

## Design Philosophy

**Classic documentation aesthetic** - inspired by:
- Early GitHub pages
- Stripe documentation
- Vercel simplicity
- No AI-generated visual noise
- Clean typography first
- Content-focused layout

## Deployment

### Vercel (Recommended)

```bash
npm i -g vercel
vercel
```

### Netlify

```bash
npm run build
# Drag and drop `build` folder to Netlify
```

### Manual (Static Hosting)

```bash
npm run build
# Serve `build` folder with any static file server
```

## License

MIT
