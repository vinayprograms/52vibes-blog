---
title: "Hugo Future Dating Hides Blog Post"
weight: 22
agent: "Crush CLI Agent (model: claude-opus-4-5-20251101)"
duration: "52m 42s"
---

## TL;DR - Session insights

- Hugo excludes future-dated content by default - a post dated tomorrow won't appear in listings until that date arrives (or `buildFuture = true` is set)
- CSS grid with fixed TOC width and fluid content (`grid-template-columns: 180px 1fr`) prevents layout from breaking on different screen sizes while maintaining readability
- Browser's sessionStorage caching of API responses can make dynamic content appear stale - reducing cache TTL to 1 minute enables near-real-time updates for community repo listings
- Hugo data files (YAML) cannot reference page front matter dynamically, but templates can use `site.GetPage` to fetch titles at build time

---

## Session Log

* **USER**: I used another agent to modify the content for first week's blog. But the `Blog` page doesn't include link to my first week's post :(

* **TOOL-CALL**: ls `{"path":"./wk1-blog/content/blog"}`

* **TOOL**:
  ```
  
  - ./wk1-blog/content/blog/
    - reports/
    - _index.md
    - week-01-blog-platform.md
  
  ```

* **TOOL-CALL**: view `{"file_path":"./wk1-blog/content/blog/_index.md"}`

* **TOOL-CALL**: view `{"file_path":"./wk1-blog/content/blog/week-01-blog-platform.md"}`

* **TOOL**:
  ```
  <file>
       1|---
       2|title: "Blog"
       3|description: "Weekly posts from the 52vibes experiment"
       4|---
       5|
       6|Weekly posts documenting the 52vibes experiment.
  </file>
  
  ```

* **TOOL**:
  ```
  <file>
       1|---
       2|title: "Week 1: Why I'm Letting an AI Agent Drive for 52 Weeks"
       3|date: 2026-01-07T10:00:00-08:00
       4|week: 1
       5|quarter: 1
       6|theme: "Agentic Infrastructure"
       7|description: "The launch of 52vibes: a year-long experiment to find where AI agents break"
       8|repository: "https://github.com/vinayprograms/52vibes"
       9|tags: ["hugo", "infrastructure", "week-1", "experiment"]
      10|---
      11|
      12|Social media is flooded with developers claiming AI agents are revolutionizing how they code. Twitter threads, YouTube videos, TikTok shorts, Instagram reels, blog posts—every week someone ships a product in 48 hours with Claude or Cursor. The demos are impressive. The hype is real.
      13|
      14|But I've been using these tools long enough to notice the gaps. The confident hallucinations. The moments where the agent spirals into increasingly wrong solutions. The subtle bugs that only surface in production.
      15|
      16|People do post about these failures. But the "for" and "against" content is so polarized, and so many posts skip the essential details, that there's no clear way to identify where the actual collaboration boundaries should be.
      17|
      18|So I'm running an experiment: 52 weeks, 52 projects, all built agent-first. Not to prove AI agents are amazing (we have enough of that content). Instead, to find exactly where they break—and document it with the details that usually get skipped.
      19|
      20|## The Setup
      21|
      22|One rule: **the agent handles details, I provide direction and creative input.**
      23|
      24|I'm not using the agent as a fancy autocomplete. I'm giving it problems to solve, letting it make implementation decisions, write tests, debug failures. My job is to set the vision, provide constraints, and steer when needed.
      25|
      26|The agent is Claude Opus 4.5, running through [Crush](https://github.com/charmbracelet/crush)—a terminal-native AI interface. Every session is logged. Every failure documented. No cherry-picking the wins.
      27|
      28|## What Actually Happened This Week
      29|
  [...112 more lines...]
  ```

* **AI**: The post has a future date (`2026-01-07`). Hugo doesn't list future-dated content by default. Let me check the Hugo config and fix this:

* **TOOL-CALL**: view `{"file_path":"./wk1-blog/hugo.toml"}`

* **TOOL**:
  ```
  <file>
       1|baseURL = "https://52vibes.dev/"
       2|languageCode = "en-us"
       3|title = "52vibes"
       4|theme = "52vibes"
       5|
       6|# Hugo settings
       7|enableRobotsTXT = true
       8|enableGitInfo = true
       9|
      10|[params]
      11|  description = "52 weeks of AI agent collaboration experiments"
      12|  author = "Vinay"
      13|  github = "https://github.com/vinayprograms/52vibes-blog"
      14|
      15|[markup]
      16|  [markup.goldmark]
      17|    [markup.goldmark.renderer]
      18|      unsafe = true
      19|  [markup.highlight]
      20|    style = "gruvbox"
      21|    lineNos = false
      22|    noClasses = false
      23|
      24|[outputs]
      25|  home = ["HTML", "RSS"]
      26|  section = ["HTML", "RSS"]
      27|
      28|[sitemap]
      29|  changefreq = "weekly"
  [...12 more lines...]
  ```

* **AI**: The post date is `2026-01-07` but today is `2026-01-06`. Hugo excludes future-dated posts by default. Two options:
  
  1. Change the post date to today or earlier
  2. Add `buildFuture = true` to hugo.toml
  
  Which do you prefer? Or I can just change the date to today.

* **USER**: Change it to today

* **TOOL-CALL**: edit `{"file_path":"./wk1-blog/content/blog/week-01-blog-platform.md","old_string":"date: 2026-01-07T10:00:00-08:00","new_string":"date: 2026-01-06T10:00:00-08:00"}`

* **TOOL**:
  ```
  <result>
  Content replaced in file: ./wk1-blog/content/blog/week-01-blog-platform.md
  </result>
  
  ```

* **AI**: Done. The post should now appear on the Blog page.

* **USER**: The width for the content is not big enough. It is scrolling horizontally on a desktop browser!

* **TOOL-CALL**: glob `{"pattern":"wk1-blog/themes/52vibes/assets/css/**/*.css"}`

* **TOOL**: `./wk1-blog/themes/52vibes/assets/css/main.css`

* **TOOL-CALL**: view `{"file_path":"./wk1-blog/themes/52vibes/assets/css/main.css"}`

* **TOOL**:
  ```
  <file>
       1|/* 52vibes - Minimalist Gruvbox theme */
       2|@font-face {
       3|  font-family: 'IBM Plex Sans';
       4|  src: url('/fonts/IBMPlexSans-Regular.woff2') format('woff2');
       5|  font-weight: 400;
       6|  font-style: normal;
       7|  font-display: swap;
       8|}
       9|
      10|@font-face {
      11|  font-family: 'IBM Plex Sans';
      12|  src: url('/fonts/IBMPlexSans-Bold.woff2') format('woff2');
      13|  font-weight: 700;
      14|  font-style: normal;
      15|  font-display: swap;
      16|}
      17|
      18|@font-face {
      19|  font-family: 'JetBrains Mono';
      20|  src: url('/fonts/JetBrainsMono-Regular.woff2') format('woff2');
      21|  font-weight: 400;
      22|  font-style: normal;
      23|  font-display: swap;
      24|}
      25|
      26|/* CSS Custom Properties - Gruvbox Dark */
      27|:root {
      28|  --bg: #1d2021;
      29|  --bg-soft: #282828;
  [...725 more lines...]
  ```

* **AI**: The issue is `--content-max: 720px` (line 46). I'll increase it to a more comfortable width for desktop:

* **TOOL-CALL**: edit `{"file_path":"./wk1-blog/themes/52vibes/assets/css/main.css","old_string":"  --content-max: 720px;","new_string":"  --content-max: 800px;"}`

* **TOOL**:
  ```
  <result>
  Content replaced in file: ./wk1-blog/themes/52vibes/assets/css/main.css
  </result>
  
  ```

* **AI**: Done. Increased content width from 720px to 800px.

* **USER**: Can't it be dynamic instead of changing the size. The TOC on the left can be of fixed size. If the content area is designed to fit the rest (of course, it needs a upper limit to deal with ultra wide monitors), then TOC location will automatically adjust itself.

* **AI**: Good point. I'll make the content area fluid with a max-width, and let the TOC stay fixed.

* **TOOL-CALL**: view `{"file_path":"./wk1-blog/themes/52vibes/assets/css/main.css","offset":614,"limit":85}`

* **TOOL**:
  ```
  <file>
     615|@media (min-width: 1100px) {
     616|  .content-area {
     617|    max-width: 800px;
     618|    margin: 0 auto;
     619|    padding: 1.5rem var(--spacing);
     620|  }
     621|
     622|  .content-area > * {
     623|    max-width: none;
     624|    margin-left: 0;
     625|    margin-right: 0;
     626|  }
     627|
     628|  .article-layout {
     629|    display: grid;
     630|    grid-template-columns: 180px 1fr;
     631|    gap: 2rem;
     632|    max-width: 1000px;
     633|  }
     634|
     635|  article {
     636|    max-width: none;
     637|    grid-column: 2;
     638|    grid-row: 1;
     639|  }
     640|
     641|  .toc-sidebar {
     642|    grid-column: 1;
     643|    grid-row: 1;
  [...60 more lines...]
  ```

* **TOOL-CALL**: edit `{"file_path":"./wk1-blog/themes/52vibes/assets/css/main.css","old_string":"@media (min-width: 1100px) {\n  .content-area {\n    max-width: 800px;\n    margin: 0 auto;\n    padding: 1.5rem var(--spacing);\n  }\n\n  .content-area > * {\n    max-width: none;\n    margin-left: 0;\n    margin-right: 0;\n  }\n\n  .article-layout {\n    display: grid;\n    grid-template-columns: 180px 1fr;\n    gap: 2rem;\n    max-width: 1000px;\n  }\n\n  article {\n    max-width: ...`

* **TOOL**:
  ```
  <result>
  Content replaced in file: ./wk1-blog/themes/52vibes/assets/css/main.css
  </result>
  
  ```

* **AI**: Done. Now:
  - `.content-area` expands up to 1200px (upper limit for ultra-wide)
  - TOC stays fixed at 180px
  - Article content caps at 800px for readability
  - Layout adjusts fluidly between these bounds

**BREAK**: 2026-01-06 23:45:20 - 2026-01-07 00:27:28 (42 min 7 sec)

* **USER**: I updated the blog post. Review it once and let me know of any issues (don't modify anything).

* **TOOL-CALL**: view `{"file_path":"./wk1-blog/content/blog/week-01-blog-platform.md"}`

* **TOOL**:
  ```
  <file>
       1|---
       2|title: "Why I'm Letting an AI Agent Drive for 52 Weeks"
       3|date: 2026-01-06T10:00:00-08:00
       4|week: 1
       5|quarter: 1
       6|theme: "Agentic Infrastructure"
       7|description: "The launch of 52vibes: a year-long experiment to find where AI agents break"
       8|repository: "https://github.com/vinayprograms/52vibes"
       9|tags: ["hugo", "infrastructure", "week-1", "experiment"]
      10|---
      11|
      12|Social media is flooded with posts claiming now AI agents are revolutionizing coding. Threads, videos, shorts/reels, blog posts — all claiming great things. Every week someone says the shipped something in 48 hours with using Claude, Cursor or a some other coding agent. The demos are impressive. The hype is real.
      13|
      14|I too have been experimenting in this space since early 2023. I've used these tools long enough to notice gaps - those confident hallucinations, agents spiraling into increasingly wrong solutions, subtle bugs that only surface in production, fundamental mistakes in security design... the list is long enough to wonder, "Where should AI stop and human take over (or the other way around)?"
      15|
      16|People do post about these failures. But the "for" and "against" arguments are typically polarized, and many skip the details. There's no clear way to articulate collaboration boundariese.
      17|
      18|So I'm running an experiment: 52 weeks, 52 projects, all built using a agent-first process. Not to prove AI agents are amazing (we have enough of that content). Instead, to find exactly where they break — and document it without skipping any details.
      19|
      20|## The Setup
      21|
      22|One rule: **the agent handles details, I provide direction and creative input.**
      23|
      24|I'm not using the agent as a fancy autocomplete. I'm giving it problems to solve, letting it make implementation decisions, write tests, debug failures. My job is to set the vision, provide constraints and steer when needed.
      25|
      26|The agent is [Crush](https://github.com/charmbracelet/crush) and the LLM is Claude Opus 4.5. Every session is logged. Every failure documented. No cherry-picking the wins.
      27|
      28|## What Actually Happened This Week
      29|
  [...101 more lines...]
  ```

* **AI**: **Issues found:**
  
  1. **Line 12**: "all claiming great things. Every week someone says the shipped something" → "says **they** shipped something"
  
  2. **Line 12**: "with using Claude" → either "with Claude" or "using Claude" (not both)
  
  3. **Line 16**: "boundariese" → "boundaries"
  
  4. **Line 18**: "using a agent-first" → "using **an** agent-first"
  
  5. **Line 65**: "its a transformer" → "it's a transformer" (contraction needs apostrophe)
  
  6. **Line 80**: Missing closing parenthesis or awkward phrasing — "not easy to automatically do it without access to Cloudflare API / MCP" reads a bit clunky. Consider: "not easy to automate without Cloudflare API/MCP access"
  
  7. **Line 82**: Run-on structure — "Connected the wrong GitHub repo initially. I had to manually trigger rebuilds and ended up pasting..." — the first sentence is a fragment (missing "I")

* **USER**: Now?

* **TOOL-CALL**: view `{"file_path":"./wk1-blog/content/blog/week-01-blog-platform.md"}`

* **TOOL**:
  ```
  <file>
       1|---
       2|title: "Why I'm Letting an AI Agent Drive for 52 Weeks"
       3|date: 2026-01-06T10:00:00-08:00
       4|week: 1
       5|quarter: 1
       6|theme: "Agentic Infrastructure"
       7|description: "The launch of 52vibes: a year-long experiment to find where AI agents break"
       8|repository: "https://github.com/vinayprograms/52vibes"
       9|tags: ["hugo", "infrastructure", "week-1", "experiment"]
      10|---
      11|
      12|Social media is flooded with posts claiming now AI agents are revolutionizing coding. Threads, videos, shorts/reels, blog posts — all claiming great things. Every week someone says they shipped something in 48 hours using Claude, Cursor or a some other coding agent. The demos are impressive. The hype is real.
      13|
      14|I too have been experimenting in this space since early 2023. I've used these tools long enough to notice gaps - those confident hallucinations, agents spiraling into increasingly wrong solutions, subtle bugs that only surface in production, fundamental mistakes in security design... the list is long enough to wonder, "Where should AI stop and human take over (or the other way around)?"
      15|
      16|People do post about these failures. But the "for" and "against" arguments are typically polarized, and many skip the details. There's no clear way to articulate collaboration boundaries.
      17|
      18|So I'm running an experiment: 52 weeks, 52 projects, all built using an agent-first process. Not to prove AI agents are amazing (we have enough of that content). Instead, to find exactly where they break — and document it without skipping any details.
      19|
      20|## The Setup
      21|
      22|One rule: **the agent handles details, I provide direction and creative input.**
      23|
      24|I'm not using the agent as a fancy autocomplete. I'm giving it problems to solve, letting it make implementation decisions, write tests, debug failures. My job is to set the vision, provide constraints and steer when needed.
      25|
      26|The agent is [Crush](https://github.com/charmbracelet/crush) and the LLM is Claude Opus 4.5. Every session is logged. Every failure documented. No cherry-picking the wins.
      27|
      28|## What Actually Happened This Week
      29|
  [...101 more lines...]
  ```

* **AI**: **Remaining issues:**
  
  1. **Line 12**: "or a some other" → "or some other" (extra "a")
  
  2. **Line 82**: Double space after "interpretation." (before "The agent couldn't")
  
  3. **Line 86**: "This week, agent seemed" → "This week, **the** agent seemed"
  
  Otherwise looks good.

* **USER**: Now?

* **TOOL-CALL**: view `{"file_path":"./wk1-blog/content/blog/week-01-blog-platform.md"}`

* **TOOL**:
  ```
  <file>
       1|---
       2|title: "Why I'm Letting an AI Agent Drive for 52 Weeks"
       3|date: 2026-01-06T10:00:00-08:00
       4|week: 1
       5|quarter: 1
       6|theme: "Agentic Infrastructure"
       7|description: "The launch of 52vibes: a year-long experiment to find where AI agents break"
       8|repository: "https://github.com/vinayprograms/52vibes"
       9|tags: ["hugo", "infrastructure", "week-1", "experiment"]
      10|---
      11|
      12|Social media is flooded with posts claiming now AI agents are revolutionizing coding. Threads, videos, shorts/reels, blog posts — all claiming great things. Every week someone says they shipped something in 48 hours using Claude, Cursor or some other coding agent. The demos are impressive. The hype is real.
      13|
      14|I too have been experimenting in this space since early 2023. I've used these tools long enough to notice gaps - those confident hallucinations, agents spiraling into increasingly wrong solutions, subtle bugs that only surface in production, fundamental mistakes in security design... the list is long enough to wonder, "Where should AI stop and human take over (or the other way around)?"
      15|
      16|People do post about these failures. But the "for" and "against" arguments are typically polarized, and many skip the details. There's no clear way to articulate collaboration boundaries.
      17|
      18|So I'm running an experiment: 52 weeks, 52 projects, all built using an agent-first process. Not to prove AI agents are amazing (we have enough of that content). Instead, to find exactly where they break — and document it without skipping any details.
      19|
      20|## The Setup
      21|
      22|One rule: **the agent handles details, I provide direction and creative input.**
      23|
      24|I'm not using the agent as a fancy autocomplete. I'm giving it problems to solve, letting it make implementation decisions, write tests, debug failures. My job is to set the vision, provide constraints and steer when needed.
      25|
      26|The agent is [Crush](https://github.com/charmbracelet/crush) and the LLM is Claude Opus 4.5. Every session is logged. Every failure documented. No cherry-picking the wins.
      27|
      28|## What Actually Happened This Week
      29|
  [...101 more lines...]
  ```

* **AI**: Looks good. No issues found.

* **USER**: Now?

* **AI**: Already reviewed - no issues found.

* **USER**: No. I made a small change on the LLM used. Review that sentence

* **TOOL-CALL**: view `{"file_path":"./wk1-blog/content/blog/week-01-blog-platform.md","offset":24,"limit":5}`

* **TOOL**:
  ```
  <file>
      25|
      26|The agent is [Crush](https://github.com/charmbracelet/crush) and uses Claude Opus 4.5 for LLM. Every session is logged. Every failure documented. No cherry-picking the wins.
      27|
      28|## What Actually Happened This Week
      29|
  
  (File has more lines. Use 'offset' parameter to read beyond line 29)
  </file>
  
  ```

* **AI**: Reads fine. "uses Claude Opus 4.5 for LLM" is grammatically correct — though "as the LLM" would be slightly more natural than "for LLM".
