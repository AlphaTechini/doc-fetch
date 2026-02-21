# DocFetch Website Deployment Guide

## Quick Deploy (Vercel - Recommended)

### Option 1: Vercel CLI

```bash
cd website
npm i -g vercel
vercel login
vercel
```

Follow the prompts:
- Set up and deploy? **Y**
- Which scope? **(your choice)**
- Link to existing project? **N**
- Project name? **docfetch**
- Directory? **.**
- Override settings? **N**

Done! Your site is live at `https://docfetch.vercel.app`

### Option 2: GitHub + Vercel

1. Push `website` folder to your repo (or keep as subdirectory)
2. Go to [vercel.com](https://vercel.com)
3. Import GitHub repository
4. Root Directory: `website`
5. Build Command: `npm run build`
6. Output Directory: `build`
7. Deploy!

## Netlify Deployment

### Drag & Drop

```bash
npm run build
```

Drag the `build` folder to [Netlify Drop](https://app.netlify.com/drop)

### Netlify CLI

```bash
npm i -g netlify-cli
netlify login
netlify init
netlify deploy --prod
```

Settings:
- Base directory: `website`
- Build command: `npm run build`
- Publish directory: `build`

## Manual Static Hosting

Build the site:

```bash
npm run build
```

The `build` folder contains static files. Serve with any HTTP server:

```bash
# Using serve
npx serve build

# Using Python
python3 -m http.server 8000 --directory build

# Using nginx
# Copy build/ contents to your web root
```

## Custom Domain

### Vercel

```bash
vercel domains add docfetch.dev
```

Then configure DNS at your registrar:
- Type: `A`
- Name: `@`
- Value: `76.76.21.21`

### Netlify

1. Go to Domain Settings
2. Add custom domain
3. Update DNS records as instructed

## Environment Variables

None required for this static site.

## Post-Deployment Checklist

- [ ] Verify all links work
- [ ] Test mobile responsiveness
- [ ] Check SEO meta tags (use [metatags.io](https://metatags.io))
- [ ] Submit sitemap to Google Search Console
- [ ] Add Google Analytics if needed
- [ ] Set up custom domain
- [ ] Configure HTTPS (automatic on Vercel/Netlify)

## Performance Optimization

The site is already optimized:
- ✅ Minimal dependencies (SvelteKit only)
- ✅ No external CSS frameworks
- ✅ Semantic HTML
- ✅ Lazy loading not needed (single page)
- ✅ Fast first paint (<1s on 3G)

If you need more speed:
- Enable gzip/brotli compression (automatic on Vercel/Netlify)
- Use CDN for static assets (automatic)
- Consider Cloudflare for additional caching

## Monitoring

Add analytics if desired:

### Google Analytics

Add to `+page.svelte` `<svelte:head>`:

```svelte
<!-- Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=GA_MEASUREMENT_ID"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());
  gtag('config', 'GA_MEASUREMENT_ID');
</script>
```

### Plausible (Privacy-Focused)

```svelte
<script defer data-domain="docfetch.dev" src="https://plausible.io/js/script.js"></script>
```

## Rollback

### Vercel

```bash
vercel rollback
```

### Netlify

Go to Deploys → Click previous deploy → "Publish deploy"

## Troubleshooting

### Build Fails

```bash
npm run check
```

Fix TypeScript errors, then rebuild.

### 404 on Refresh

Add `_redirects` file in `static/`:

```
/*    /index.html   200
```

### Slow Load Times

- Check network tab in DevTools
- Ensure hosting provider has CDN enabled
- Consider enabling HTTP/2

## Support

Issues: https://github.com/AlphaTechini/doc-fetch/issues
