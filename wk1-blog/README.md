# 52vibes Blog

Hugo static site for the 52vibes experiment.

## Requirements

- Hugo Extended v0.120.0+ ([install guide](https://gohugo.io/installation/))

## Local Development

```bash
# Start development server
hugo server -D

# Build for production
hugo --minify
```

## Deployment

Deployed automatically via Cloudflare Pages on push to `main`.

### Cloudflare Pages Settings

- **Build command**: `hugo --minify`
- **Build output directory**: `public`
- **Environment variable**: `HUGO_VERSION=0.140.0`

## Structure

```
wk1-blog/
├── content/
│   ├── blog/           # Weekly posts
│   ├── about/          # About page
│   ├── community/      # Community forks
│   └── weeks/          # Week index
├── themes/52vibes/     # Custom theme
│   ├── layouts/        # Templates
│   ├── assets/css/     # Styles
│   └── static/         # Static files
├── static/
│   ├── fonts/          # Self-hosted fonts
│   └── _headers        # Cloudflare headers
└── hugo.toml           # Site config
```

## Theme

Custom tmux-inspired theme with:
- Gruvbox color scheme (dark/light)
- Responsive layout (mobile-first)
- Minimal JavaScript (~600 bytes total)
- WCAG-compliant accessibility
- CLI browser support (lynx, w3m)

## Fonts

Download and place in `static/fonts/`:
- [IBM Plex Sans](https://github.com/IBM/plex) (WOFF2)
- [JetBrains Mono](https://github.com/JetBrains/JetBrainsMono) (WOFF2)
