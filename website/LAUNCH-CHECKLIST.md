# DocFetch Website Launch Checklist

## Pre-Launch

### Content Review
- [ ] Verify all installation commands are correct
- [ ] Test all example commands locally
- [ ] Check GitHub links point to correct repo
- [ ] Verify NPM/PyPI badges show correct versions
- [ ] Proofread all copy for typos

### Technical Checks
- [ ] Run `npm run build` successfully
- [ ] Test production build with `npm run preview`
- [ ] Verify responsive design (mobile, tablet, desktop)
- [ ] Check all internal anchor links work
- [ ] Test external links open in new tabs
- [ ] Verify favicon displays correctly

### SEO Verification
- [ ] Meta title: "DocFetch - Transform Documentation Sites into AI-Ready Markdown"
- [ ] Meta description: Present and under 160 characters
- [ ] Open Graph tags render correctly (test with [metatags.io](https://metatags.io))
- [ ] Twitter Card preview works
- [ ] JSON-LD structured data validates ([Google Rich Results Test](https://search.google.com/test/rich-results))
- [ ] Canonical URL set correctly

### Performance
- [ ] Lighthouse score >90 (run in Chrome DevTools)
- [ ] First contentful paint <1.5s
- [ ] Time to interactive <3s
- [ ] No console errors or warnings
- [ ] Images optimized (SVG for logo/icons)

### Accessibility
- [ ] All images have alt text
- [ ] Proper heading hierarchy (h1 → h2 → h3)
- [ ] Color contrast meets WCAG AA
- [ ] Keyboard navigation works
- [ ] Focus states visible

## Deployment

### Vercel Deployment
```bash
cd website
vercel --prod
```

- [ ] Custom domain configured (docfetch.dev)
- [ ] HTTPS enabled (automatic on Vercel)
- [ ] DNS propagated
- [ ] Redirect www → non-www (or vice versa)

### Post-Deployment
- [ ] Live site loads without errors
- [ ] All links work on production
- [ ] Mobile view tested on real device
- [ ] Social sharing preview looks correct

## Post-Launch

### Analytics & Monitoring
- [ ] Google Analytics installed (optional)
- [ ] Google Search Console property added
- [ ] Sitemap submitted to Google
- [ ] Bing Webmaster Tools added (optional)

### Marketing
- [ ] Announce on Twitter/X
- [ ] Post to Hacker News "Show HN"
- [ ] Share in relevant subreddits (r/webdev, r/sveltejs)
- [ ] Add to Product Hunt (optional)
- [ ] Update main README.md with website link
- [ ] Add website link to NPM package
- [ ] Add website link to PyPI package
- [ ] Update GitHub repo description with website URL

### Documentation
- [ ] Add website link to doc-fetch main README
- [ ] Create screenshot gallery (optional)
- [ ] Write launch blog post (optional)

## Ongoing Maintenance

### Monthly
- [ ] Check for broken links
- [ ] Review analytics for issues
- [ ] Update dependencies (`npm update`)
- [ ] Check for SvelteKit updates

### Quarterly
- [ ] Refresh screenshots if UI changed
- [ ] Update version numbers in examples
- [ ] Review and update features section
- [ ] Check competitor landscape

## Rollback Plan

If something goes wrong:

1. **Immediate**: Revert to previous Vercel deployment
   ```bash
   vercel rollback
   ```

2. **Fix locally**: Address the issue in code

3. **Redeploy**: 
   ```bash
   vercel --prod
   ```

## Success Metrics

Track these after launch:

- **Traffic**: Page views, unique visitors
- **Engagement**: Time on page, bounce rate
- **Conversions**: GitHub stars, npm installs, pypi downloads
- **SEO**: Search rankings for "documentation fetcher", "markdown converter"
- **Social**: Shares, mentions, backlinks

## Contact

For issues: https://github.com/AlphaTechini/doc-fetch/issues

---

**Launch Date**: _________________

**Launched By**: _________________

**Notes**: _______________________________________________
