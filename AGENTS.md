# AGENTS.md

Agent-focused documentation for the 52vibes repository.

---

## Project Overview

52vibes is a year-long (52 weeks) AI agent collaboration experiment. Each week produces a shipped artifact while documenting human-AI collaboration patterns.

**Key Resources:**
- `PROJECT.md` — Philosophy, constraints, themes
- `ROADMAP.md` — 52-week project list with dependencies

---

## Repository Structure

```
52vibes/
├── PROJECT.md              # Project philosophy and constraints
├── ROADMAP.md              # 52-week project roadmap
├── wk1-blog/               # Week 1: Static Blog Platform
│   ├── README.md           # Setup and development guide
│   ├── SETUP_GUIDE.md      # Detailed setup instructions
│   ├── hugo.toml           # Hugo site configuration
│   ├── content/            # Site content (blog/, about/, weeks/, community/)
│   ├── themes/52vibes/     # Custom Hugo theme
│   │   ├── layouts/        # Templates and partials
│   │   └── assets/css/     # Stylesheets
│   ├── static/             # Static files (fonts, _headers)
│   ├── data/               # Data files (weeks.yaml)
│   └── design/             # Formal design documents
│       ├── 1.1_NEEDS.md       # Customer needs (RFC 2119)
│       ├── 1.2_ACCEPTANCE.md  # Acceptance criteria
│       ├── 2.1_REQUIREMENTS.md   # Product requirements
│       ├── 2.2_QA.md             # QA test specifications
│       ├── 3.1_TECH_REQUIREMENTS.md  # Technical requirements
│       └── 3.2_SYSTEM_TESTS.md       # System test specifications
├── crush-export/           # Tool to export Crush sessions to markdown
│   ├── cmd/crush-export/   # CLI entrypoint
│   ├── internal/           # Go packages (db, export, model)
│   ├── go.mod
│   └── crush-export        # Compiled binary
├── sessions/               # Exported session logs (markdown)
├── notes/                  # Zettelkasten design notes
├── .github/workflows/      # GitHub Actions
│   └── wiki-sync.yml       # Sync sessions to wiki
└── .crush/                 # Crush AI agent data (SQLite database)
```

---

## Tools

### crush-export

Exports Crush AI agent sessions from SQLite database to markdown files.

**Location:** `crush-export/`

**Build:**
```bash
cd crush-export
go build -o crush-export ./cmd/crush-export
```

**Usage:**
```bash
# List all sessions
./crush-export -l

# Export all sessions to sessions/ directory
./crush-export -db ../.crush/crush.db -out ../sessions/

# Export specific session
./crush-export -db ../.crush/crush.db -session <session-id> -out ../sessions/
```

**Flags:**
- `-db` — Path to crush.db (default: `.crush/crush.db`)
- `-out` — Output directory for markdown files (default: `sessions`)
- `-session` — Export specific session ID (optional)
- `-l` — List all sessions

---

## Design Documents

Week projects follow a formal design process with numbered documents:

| File | Purpose |
|------|---------|
| `1.1_NEEDS.md` | Customer needs (RFC 2119 keywords) |
| `1.2_ACCEPTANCE.md` | Acceptance criteria mapped to needs |
| `2.1_REQUIREMENTS.md` | Product requirements |
| `2.2_QA.md` | QA test specifications |
| `3.1_TECH_REQUIREMENTS.md` | Technical requirements and architecture |
| `3.2_SYSTEM_TESTS.md` | System test specifications |

**Conventions:**
- Needs use IDs like `N001`, `N002`, etc.
- Acceptance criteria use `A001`, `A002`, etc. (maps 1:1 to needs)
- Requirements use product prefixes: `BLOG001`, `CX001`, `SL001`, `COM001`
- QA specs use `BLOG_QA_001`, `CX_QA_001`, etc.
- Tech requirements use `BLOG_TECH_001`, `CX_TECH_001`, etc.

**Zettelkasten Notes:**
Design documents are also stored in `notes/` as linked zettelkasten entries for cross-referencing.

---

## Project Themes

| Quarter | Weeks | Theme |
|---------|-------|-------|
| Q1 | 1-13 | Agentic Infrastructure |
| Q2 | 14-26 | Production Tools |
| Q3 | 27-39 | Complex Workflows |
| Q4 | 40-52 | Synthesis & Edge Cases |

---

## Week 1: Static Blog Platform

**Goal:** Static blog at `https://52vibes.dev/` with tmux-inspired design.

**Key Decisions:**
- Hugo static site generator (v0.140.0+)
- Cloudflare Pages deployment (native build)
- No JavaScript except: theme toggle (~500 bytes), community page API (~500 bytes)
- Gruvbox color scheme (dark/light)
- IBM Plex Sans + JetBrains Mono fonts (self-hosted WOFF2)
- CLI browser support (lynx, w3m)

**Products:**
1. **Blog Platform (BLOG)** — Hugo site with custom `52vibes` theme
2. **crush-export (CX)** — Session export CLI tool (Go)
3. **Session Log Sync (SL)** — GitHub Actions for wiki sync
4. **Community Features (COM)** — Fork showcase page

**Local Development:**
```bash
cd wk1-blog
hugo server -D        # Development server with drafts
hugo --minify         # Production build
```

**Cloudflare Pages Settings:**
- Build command: `hugo --minify`
- Output directory: `public`
- Environment: `HUGO_VERSION=0.140.0`

---

## Conventions

### File Naming
- Week directories: `wk1-blog/`, `wk2-security/`, etc.
- Session exports: `YYYY-MM-DD_HH-MM-SS_<title>.md`

### Markdown
- RFC 2119 keywords (MUST, SHOULD, MAY) used in needs/requirements
- Bulleted lists with `*` for specifications
- Tables for structured data

### Code
- Go for tooling (crush-export)
- Hugo templates for blog

---

## Session Logs

Session data is stored in `.crush/crush.db` (SQLite). Use `crush-export` to generate markdown for wiki publication.

**Workflow:**
1. `crush-export` generates markdown in `sessions/`
2. Human reviews and curates (adds Lessons Learned)
3. Push to main triggers GitHub Actions
4. Actions syncs `sessions/*.md` to GitHub Wiki

---

## Gotchas

1. **Crush database location** — Sessions are in `.crush/crush.db`, not `.crush/sessions/`
2. **No external JavaScript** — All JS must be embedded inline, under 1KB total
3. **CLI browser testing** — Test with both lynx (strict) and w3m (tables)
4. **Week numbering** — Zero-padded: `week-01`, `week-02`, etc.
5. **Font files required** — Download IBM Plex Sans and JetBrains Mono WOFF2 to `static/fonts/`
6. **Wiki sync PAT** — Requires `WIKI_TOKEN` secret with repo scope for wiki write access
7. **Hugo version** — Use Extended build v0.120.0+ for CSS processing
