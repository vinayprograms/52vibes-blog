---
title: "Why I'm letting an AI agent drive for 52 weeks"
date: 2026-01-07T09:00:00-05:00
week: 1
quarter: 1
theme: "Agentic Infrastructure"
description: "The launch of 52vibes: a year-long experiment to find where AI agents break"
repository: "https://github.com/vinayprograms/52vibes"
tags: ["hugo", "infrastructure", "week-1", "experiment"]
---

Social media is flooded with posts talking about AI agents revolutionizing coding. Threads, videos, shorts/reels, blog posts — all claiming great things. Every week someone says they shipped something in 48 hours using Claude, Cursor or some other coding agent. The demos are impressive. The hype is real.

I have been running small experiments in this space since 2023. In my experiments I've noticed gaps - confident hallucinations, agents spiraling into increasingly wrong solutions, subtle bugs that only surface in production, fundamental mistakes in security design, etc. The list is long enough to wonder, "Where should AI stop and human take over (or the other way around)?"

People do post about these failures. But the "for" and "against" arguments are typically polarized, and many skip the details. There's no clear way to articulate collaboration boundaries.

So I'm running an year long experiment: 52 weeks, 52 projects, all built using an agent-first process. This is not to prove AI agents are amazing (we have enough of that content) or that they are useless where it matters. Instead, to find exactly where they break — and document it without skipping any details.

## The Setup

One rule: **the agent handles details, I provide direction and creative input.**

I'm not using the agent as a fancy auto-complete. I'm giving it problems to solve, letting it make implementation decisions, write tests, debug failures. My job is to set the vision, provide constraints and steer when needed.

The agent is [Crush](https://github.com/charmbracelet/crush) and uses Claude Opus 4.5 for LLM. Every session is logged. Every failure documented. No cherry-picking the wins.

## What actually happened this week

The numbers:

| Metric | Value |
|--------|-------|
| Sessions | 30 |
| Total messages | 2,034 |
| My messages | 122 |
| Agent messages | 627 |
| Message ratio | 5.1x |

For every prompt I gave, the agent generated five responses on average — reading files, making decisions, writing code, iterating. Some sessions were more autonomous. In the implementation session, I sent **5 messages** while the agent worked through **83 responses** — a 16.6x ratio — converting technical requirements into working code.

Here's what I actually observed.

### The collaboration model

I designed a 3-stage formal requirements elaboration process:

`NEEDS`+`ACCEPTANCE` → `REQUIREMENTS`+`QA` → `TECH REQUIREMENTS`+`SYSTEM TESTS`

For each stage, I wrote detailed prompts explaining exactly how to transform the previous stage's output. The first stage — capturing customer needs — was collaborative. I described what I wanted (tmux-inspired theme, minimal JavaScript, Gruvbox colors) and we iterated together to formalize those into proper need statements with acceptance criteria.

Then I stepped back.

The agent took those needs and autonomously generated:
- Product requirements with traceable IDs (BLOG001, CX001, etc.) plus QA test specifications for each requirement.
- Technical specifications mapping requirements to implementation details and system test criteria.
- Implementation TODOs linked to technical specifications.

303 messages of requirements work from essentially a handful of prompts. Detailed traceability across stages emerged with minimal prompting. Decision on which RFC 2119 keywords (MUST, SHOULD, MAY) to use happened with almost no guidance, except for me asking the agent to use it.

### When it worked

The agent excelled at transformations:
- Vague requirements → formal specifications
- Specifications → implementation TODOs
- TODOs → working code

For `crush-export` — a Go CLI tool for exporting Crush sessions to markdown — the agent:
- Consistently applied TDD all along with me only asking for it somewhere in the initial prompt
- Handled filename collisions with numeric suffixes
- Caught unicode edge cases in test data
- Designed a clean layered architecture (cmd → internal/db → internal/export → internal/model)

I didn't review individual lines of code during implementation. I described what I wanted, the agent built it, and the tests passed.

### Where the boundary might be

Cloudflare Pages deployment had to be done manually — not because the agent couldn't, but because it was not easy to automate without giving access to Cloudflare API/MCP. The agent clearly "understood" (I'm anthropomorphizing here) that these are manual steps I need to perform on the Cloudflare dashboard.

And I made mistakes:
* I connected the wrong GitHub repo initially.
* I had to manually trigger rebuilds and ended up pasting build logs back to the agent for interpretation.

The agent couldn't click buttons for me. It couldn't see the Cloudflare dashboard. But it could interpret error logs, suggest fixes, and guide me through recovery. When the build failed due to Hugo version mismatches, it diagnosed the problem from the logs I provided.

This also revealed the human side of this challenge - when should I curb the urge to copy-paste everything into an agent and do what it says, without even a cursory review.

This week, the agent seemed to excel at text transformations, I handled integrations. But I'm not ready to call that a rule yet — it's one data point.

### When the human overdoes it

Here's something that doesn't get discussed: humans can overengineer too.

That 3-stage process? That was me. The agent didn't ask for it. I imposed it.

Was it necessary for a week-1 blog? Almost certainly not. A simpler "describe what you want, let's build it" approach would have shipped faster. Instead, I spent significant time writing detailed prompts for each stage, reviewing intermediate outputs, and randomly verifying traceability to confirm that agent is following my instructions.

When an agent produces excessive code, we cry "hallucination". When a human imposes excessive process, we call it... being thorough? Best practices?

In a corporate setting, this matters. All that extra hours of work adds to project cost. The agent would have happily built the blog in a fraction of the time if I hadn't insisted on the ceremony. The agent doesn't push back on process any more than it pushes back on scope. It implements what you describe. **The discipline has to come from the human — including discipline about when NOT to be "over-disciplined".**

## The Artifacts

What shipped:

1. **This blog** — Hugo static site at [52vibes.blog](https://52vibes.blog), custom theme, dark/light toggle, CLI browser support
2. **crush-export** — Go CLI for exporting Crush sessions to markdown
3. **Design documents** — Full requirements traceability from needs to system tests

What's pending:
- Session log publication (curation in progress)
- Community page content

## Why 2026

The honest answer: because we're at an inflection point.

In 2024, AI assistants were impressive demos. In 2025, they started showing some practical results. In 2026, I'm think we'll see agents that can sustain multi-hour autonomous work sessions and give good quality results.

But "can sustain" and "should sustain" are different questions. When an agent generates 83 responses from 5 prompts, how do you verify the work? How do you catch subtle errors compounding over dozens of decisions? And when a human over-specifies the process, how much is lost to unnecessary ceremony?

This is the question 52vibes is trying to answer. Not "can AI agents code?" but "how should humans collaborate with AI agents?" The boundaries aren't just about agent limitations — they're about human tendencies too.

52 weeks. 52 chances to find out.

---

[View all session logs for this week →](/logs/week-01-blog-platform/)
