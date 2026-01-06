# 52vibes Blog Setup Guide

Step-by-step instructions for completing the remaining setup tasks.

---

## 1. Set Up Self-Hosted Fonts

The theme uses IBM Plex Sans and JetBrains Mono. Download WOFF2 files and place in `static/fonts/`.

### Steps

1. **Download IBM Plex Sans**
   ```bash
   # From https://github.com/IBM/plex/releases
   # Download: IBMPlexSans-Regular.woff2, IBMPlexSans-Bold.woff2, IBMPlexSans-Italic.woff2
   ```

2. **Download JetBrains Mono**
   ```bash
   # From https://github.com/JetBrains/JetBrainsMono/releases
   # Download: JetBrainsMono-Regular.woff2
   ```

3. **Place files**
   ```
   static/fonts/
   ├── IBMPlexSans-Regular.woff2
   ├── IBMPlexSans-Bold.woff2
   ├── IBMPlexSans-Italic.woff2
   └── JetBrainsMono-Regular.woff2
   ```

4. **Add @font-face rules** to `themes/52vibes/assets/css/main.css`:
   ```css
   @font-face {
     font-family: 'IBM Plex Sans';
     src: url('/fonts/IBMPlexSans-Regular.woff2') format('woff2');
     font-weight: 400;
     font-style: normal;
     font-display: swap;
   }

   @font-face {
     font-family: 'IBM Plex Sans';
     src: url('/fonts/IBMPlexSans-Bold.woff2') format('woff2');
     font-weight: 700;
     font-style: normal;
     font-display: swap;
   }

   @font-face {
     font-family: 'JetBrains Mono';
     src: url('/fonts/JetBrainsMono-Regular.woff2') format('woff2');
     font-weight: 400;
     font-style: normal;
     font-display: swap;
   }
   ```

---

## 2. Create Weeks Data File (Optional)

A data file can provide structured week metadata. Currently using content-based approach instead.

### If you want structured data

1. **Create** `data/weeks.yaml`:
   ```yaml
   - week: 1
     title: "Blog Platform"
     quarter: 1
     theme: "Agentic Infrastructure"
     status: "complete"
     post: "/blog/week-01-blog-platform/"
   
   - week: 2
     title: "Security Tooling"
     quarter: 1
     theme: "Agentic Infrastructure"
     status: "upcoming"
   ```

2. **Access in templates** via `{{ site.Data.weeks }}`

---

## 3. Create Weekly Index Template (Optional)

The current `content/weeks/_index.md` uses inline Hugo templating. For a dedicated template:

1. **Create** `themes/52vibes/layouts/weeks/list.html`:
   ```html
   {{ define "main" }}
   <section>
     <h1>{{ .Title }}</h1>
     
     <table>
       <thead>
         <tr>
           <th>Week</th>
           <th>Quarter</th>
           <th>Theme</th>
           <th>Project</th>
         </tr>
       </thead>
       <tbody>
         {{ range site.Data.weeks }}
         <tr>
           <td>{{ .week }}</td>
           <td>Q{{ .quarter }}</td>
           <td>{{ .theme }}</td>
           <td>
             {{ if eq .status "complete" }}
             <a href="{{ .post }}">{{ .title }}</a>
             {{ else }}
             {{ .title }}
             {{ end }}
           </td>
         </tr>
         {{ end }}
       </tbody>
     </table>
   </section>
   {{ end }}
   ```

---

## 4. Create External Link Partial (Optional)

Adds visual indicator and security attributes to external links.

1. **Create** `themes/52vibes/layouts/partials/external-link.html`:
   ```html
   <a href="{{ .url }}" target="_blank" rel="noopener noreferrer">
     {{ .text }} <span aria-hidden="true">↗</span>
   </a>
   ```

2. **Usage in templates**:
   ```html
   {{ partial "external-link.html" (dict "url" "https://example.com" "text" "Example") }}
   ```

---

## 5. Configure Cloudflare Pages Deployment

### Steps

1. **Go to** [Cloudflare Pages Dashboard](https://dash.cloudflare.com/?to=/:account/pages)

2. **Create a project**
   - Connect to GitHub repository
   - Select `52vibes` repository

3. **Configure build settings**
   | Setting | Value |
   |---------|-------|
   | Production branch | `main` |
   | Build command | `cd wk1-blog && hugo --minify` |
   | Build output directory | `wk1-blog/public` |
   | Root directory | `/` |

4. **Add environment variable**
   | Variable | Value |
   |----------|-------|
   | `HUGO_VERSION` | `0.140.0` |

5. **Save and deploy**

---

## 6. Configure Custom Domain

### Steps

1. **In Cloudflare Pages project** → Custom domains

2. **Add domain**: `52vibes.dev`

3. **DNS Configuration** (if using Cloudflare DNS):
   - Automatic CNAME added

4. **Add www redirect** (optional):
   - Add `www.52vibes.dev`
   - Configure redirect rule: `www.52vibes.dev` → `52vibes.dev`

5. **Verify HTTPS**:
   - Certificate auto-provisioned
   - Check "Always use HTTPS" is enabled

---

## 7. Verify Preview Deployments

### Steps

1. **Create a PR** to the repository

2. **Check Cloudflare Pages** creates a preview deployment

3. **Preview URL format**: `<commit-hash>.<project>.pages.dev`

4. **Verify**:
   - Site builds successfully
   - Theme renders correctly
   - Links work

---

## 8. Verify Local Development Parity

### Steps

1. **Install Hugo** (Extended edition):
   ```bash
   # macOS
   brew install hugo

   # Verify extended edition
   hugo version  # Should show "extended"
   ```

2. **Run local server**:
   ```bash
   cd wk1-blog
   hugo server -D
   ```

3. **Check**:
   - http://localhost:1313 loads
   - Theme matches production
   - Hot reload works on file changes
   - Draft posts visible with `-D` flag

---

## 9. Implement Image Processing Pipeline (Optional)

Hugo can process images at build time for responsive images.

### Steps

1. **Create image shortcode** `themes/52vibes/layouts/shortcodes/img.html`:
   ```html
   {{ $src := .Get "src" }}
   {{ $alt := .Get "alt" }}
   {{ $img := resources.Get $src }}
   {{ if $img }}
     {{ $small := $img.Resize "480x" }}
     {{ $medium := $img.Resize "800x" }}
     {{ $large := $img.Resize "1200x" }}
     <picture>
       <source srcset="{{ $large.RelPermalink }}" media="(min-width: 1024px)">
       <source srcset="{{ $medium.RelPermalink }}" media="(min-width: 768px)">
       <img src="{{ $small.RelPermalink }}" alt="{{ $alt }}" loading="lazy">
     </picture>
   {{ else }}
     <img src="{{ $src }}" alt="{{ $alt }}" loading="lazy">
   {{ end }}
   ```

2. **Place images** in `assets/images/` (not `static/`)

3. **Use in content**:
   ```markdown
   {{</* img src="images/screenshot.png" alt="Screenshot" */>}}
   ```

---

## 10. Set Up Social Card Generation (Optional)

Generate Open Graph images for social sharing.

### Option A: Static default image

1. **Create** `static/images/og-default.png` (1200x630px)
2. Already referenced in `meta.html` partial

### Option B: Per-post images

1. **Add to front matter**:
   ```yaml
   image: "/images/week-01-card.png"
   ```

2. **Create image** for each post

### Option C: Dynamic generation (advanced)

Requires external tool like [hugo-og-image](https://github.com/Ladicle/hugo-og-image) or custom build step.

---

## 11. Implement Performance Optimizations (Optional)

### Already implemented
- CSS minification via Hugo Pipes
- HTML minification in `hugo.toml`
- Cache headers in `_headers`

### Additional optimizations

1. **Critical CSS inlining** (if needed):
   ```html
   {{ $critical := resources.Get "css/critical.css" | minify }}
   <style>{{ $critical.Content | safeCSS }}</style>
   ```

2. **Preload fonts**:
   ```html
   <link rel="preload" href="/fonts/IBMPlexSans-Regular.woff2" as="font" type="font/woff2" crossorigin>
   ```

3. **Lazy load images** (already in image shortcode):
   ```html
   <img loading="lazy" ...>
   ```

---

## Verification Checklist

After completing setup:

- [ ] `hugo server` runs without errors
- [ ] Fonts load correctly (check Network tab)
- [ ] Dark/light theme toggle works
- [ ] Mobile layout renders correctly
- [ ] Cloudflare Pages deploys successfully
- [ ] Custom domain resolves
- [ ] HTTPS certificate valid
- [ ] Security headers present (check securityheaders.com)
- [ ] RSS feed accessible at `/blog/index.xml`
- [ ] Sitemap accessible at `/sitemap.xml`
