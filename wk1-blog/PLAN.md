# Week 1: Static Blog Infrastructure

**Goal:** A working static blog at `https://52vibes.dev/blog` with CI/CD pipeline that converts markdown → HTML/CSS and deploys to Cloudflare Pages.

**Time Budget:** 10-20 hours

---

## Deliverables

1. Hugo-based static site generator setup
2. Custom minimal theme (pure HTML/CSS, no JavaScript)
3. GitHub Actions workflow for automated builds
4. Cloudflare Pages deployment
5. Working blog at `https://52vibes.dev/blog`

---

## Technical Decisions

### Static Site Generator: Hugo

- Fastest builds, single binary
- Extensive theme ecosystem
- Outputs pure HTML/CSS
- Mature and well-documented

### Hosting Architecture

```
GitHub repo (markdown) 
    → GitHub Actions (build) 
    → Cloudflare Pages (host)
```

### Repository Structure

```
52vibes-blog/
├── content/
│   └── posts/
│       └── week-01-infrastructure.md
├── themes/
│   └── 52vibes/          # Custom minimal theme
├── static/
├── config.toml
└── .github/workflows/
    └── deploy.yml
```

### Design Principles

- Clean typography (system fonts or self-hosted)
- Responsive CSS Grid/Flexbox layout
- Dark/light mode via `prefers-color-scheme`
- Minimal, fast-loading
- Good mobile reading experience

---

## Open Questions (Resolve Day 1)

1. **Theme approach** - Build custom theme from scratch, or start with minimal Hugo theme and strip it down?
2. **Content structure** - Just blog posts, or also pages like "About", "Weekly Index"?
3. **Writing workflow** - Draft posts in repo directly, or separate drafts location?
4. **Domain setup** - Is `52vibes.dev` already pointing to Cloudflare, or need DNS config?
5. **First post** - Include writing/publishing an intro post, or just infrastructure?

---

## Success Criteria

- [ ] Blog accessible at `https://52vibes.dev/blog`
- [ ] Push to `main` triggers automatic build and deploy
- [ ] Pages render correctly on desktop and mobile
- [ ] No JavaScript in final output
- [ ] Lighthouse performance score > 95
