---
title: "When Features Ship But Performance Doesn't"
date: 2026-01-20T09:00:00-05:00
week: 2-3
quarter: 1
theme: "Agentic Infrastructure"
description: "Vibe coding ships features fast. Performance, platform quirks, and verification? Not so much."
repository: "https://github.com/vinayprograms/52vibes"
tags: ["memory", "infrastructure", "week-2", "week-3", "golang", "c", "vibe-coding", "containers"]
---

I was listening to the Pragmatic Engineer podcast with Martin Fowler at the beginning of the 2nd week. He made an observation that framed my next two weeks: with vibe coding, you're removing the learning loop - you can optimize the LLM's feedback loop to get things right at the cost of dropping your understanding of the code it wrote. In control system theory terms, **you become a runaway process**.

I ran three experiments to dig into this further.

## Experiment 1: Security-hardened containers

I built [claudebox](https://github.com/vinayprograms/claudebox) — a security-hardened container to run Claude Code with defense-in-depth protections as well as apply the "Deny by default, allow by exception" principle. The idea was to limit the damage an arbitrary / malicious NPM package used by Claude Code can do to my host system. It can only access specific locations on the internet access the internet, cannot read files outside your project, or access privileged operations from host kernel.

LLM's proposed architecture worked - multi-container setup with Squid proxy, domain allowlisting, rate limiting, credential isolation, built as ~1500 lines of shell, Dockerfiles, and configs.

**Where it struggled:** Platform-level details. When Claude Code, running on macOS, spawns parallel subagents reading many files simultaneously, the VirtioFS file sharing layer in Docker/Podman VMs gets overwhelmed. You get `Unknown system error -35` — transient failures that look like bugs but are platform limitations. The agent kept trying to "fix" code that wasn't broken. The real fix was architectural: staging directories, smaller workspaces, accepting that some things can't be solved in code.

**Insights:**
1. Vibe coding infrastructure features requires a more complex agentic system. We already know that infrastructure code requires testing against real environments — you can't mock a container runtime or other system dependencies, as compared to application code that can be sandboxed and tested in isolation. Current agentic coding environments do not provide such testing environments nor do they provide tools to make it easy to build such environments. This effectively negates the velocity advantage of vibe coding for infrastructure work.
2. Coding agents cannot readily figure out what testing environments and tools are required to test infrastructure code as opposed to application code. I only need to ask the agent "use TDD to test your code" and it will know what to do. In contrast, infrastructure requires lot more reasoning about the environment, dependencies and tools required to test the code and agents today cannot do it out of the box.

## Experiment 2: Memory service in C

I built a [memory service](https://github.com/vinayprograms/agent-memory/tree/main/memory-c) for AI agents — semantic search over conversation history so agents can build context through selective querying shoving everything into its context window leading to "context rot". The target: <10ms p99 latency.

Why C? My argument was that Go's GC pauses could hit 1-5ms under load — half the 10ms latency budget gone to garbage collection. C has zero GC, direct ONNX API access, predictable performance under concurrency. The trade-off? C means slower development. But I reasoned: "Development velocity concern is mitigated since Claude implements the code."

That was the bet. Let AI handle C's velocity disadvantage. See if performance follows.

The result was ~12,300 lines of C that implemented HNSW index, LMDB storage, ONNX embeddings and JSON-RPC API. Features shipped fast. The velocity bet paid off.

But, it missed <10ms by an order of magnitude. Smaller embedding models took a few hundred milliseconds. Larger models: 2-3 seconds. I asked for Apple Silicon and NVIDIA acceleration to be added. The agent claimed that it had implemented it, but it always fell back to CPU during execution (when running on the host OS).

**Insight:** Agents are not mature enough to meet aggressive non-functional requirements (like performance). This is a classic problem in software engineering. Designs that have agressive non-functional requirements are fundamentally different from those where there is enough room for delays and latencies. AI delivered on development velocity — the reason I thought I could use C despite its overhead. But it didn't deliver on performance — the reason I chose C in the first place. The language choice was correct. The execution didn't follow.

## Experiment 3: C to Go migration

I asked the agent to identify Go alternatives for C libraries (ONNX, HNSW, WAL) and [migrate](https://github.com/vinayprograms/agent-memory/tree/main/memory-go) the architecture. It generated ~7,200 lines of Go. The translation preserved design decisions while adapting to Go idioms.

This worked well. Language migration with architectural preservation is achievable.

## The testing trap

Across all three experiments, I asked the agent to generate tests. It did. They passed.

Here's the problem: the agent's tests verify its implementation is internally consistent. They don't verify the implementation matches what I "wanted". When you vibe code tests alongside implementation, you get a closed loop — the agent checking its own work against its own understanding.

**To verify the agent's output, you need tests the agent didn't write.**

## Brains need tools

The podcast pointed out two things relevant to this discussion -
1. Everything we experiment with in LLM space is greenfield. Brownfield success at scale remains unproven.
2. Things like renaming a class is a solved problem with IDEs. Yet LLMs cannot reliably do it across a codebase. Understanding code and merely modifying code based on plaintext specifications are different capabilities.

This also points to something important: **brains without powerful tools aren't effective.** Tools that handle complex relationships and modify targets based on those relationships give orders of magnitude more leverage. Humans became the dominant species not because of intelligence alone. The human body containint powerfule sensory and motor systems combined with brains gave us true leverage.


## What shipped

1. **Claudebox** — Security-hardened container for Claude Code. Currently usable in development setups. Still has a few limitations in ways it can be used.
2. **Memory Service (C)** — 12,300 lines. Works functionally. Misses latency target.
3. **Memory Service (Go)** — 7,200 lines. Cleaner migration of the same architecture.

This was part 1 of N. The memory service work continues.

## The takeaways

**One:** Vibe coding excels at functional requirements and fails at non-functional ones. Features ship; performance doesn't follow automatically.

**Two:** Coding agents don't come with the ability to automatically build elaborate, sandboxed infrastructure testing rigs. This limitation makes it difficult to use agents for vibe coding infrastructure stuff.

**Three:** Agent-generated tests create a verification loop for the agent, not the user. To verify that work truly meets your needs, you must spend time building tests as you understand them instead of letting the agent run fully automated TDD loops.

**Four:** You can optimize the LLM's feedback loop while breaking your own. If you're not learning as you build, you're accumulating debt you don't understand.

## The boundary

Vibe coding works when:
- Functional correctness is verifiable by inspection
- Performance isn't critical
- You can catch failures before they compound

It struggles when:
- Non-functional requirements matter
- Platform/infra details are involved
- You need the code to work in ways you can't directly observe

---
**Footnotes**
1. Anthropic blocked third-party agents from Claude Code Max. I had spend initial days of week-2 migrating all my customizations from Crush to Claude Code. Hence the reason for merging week-2 and week-3 into a single experiment.
2. I may be prompting wrong. I hope readers can point out my mistakes, bad/wrong choices after looking at the logs.

[View all session logs for this week →](/logs/week-02-03/)
