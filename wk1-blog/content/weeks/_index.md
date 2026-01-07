---
title: "Weekly Projects"
description: "List of all weekly projects"
---

# Weekly Index

All 52 weeks of the experiment.

| Week | Quarter | Theme | Project |
|------|---------|-------|---------|
{{ range $i := seq 52 }}
{{ $weekNum := printf "%02d" $i }}
{{ $quarter := add (div (sub $i 1) 13) 1 }}
{{ $theme := index (slice "Agentic Infrastructure" "Production Tools" "Complex Workflows" "Synthesis") (sub $quarter 1) }}
| {{ $i }} | Q{{ $quarter }} | {{ $theme }} | {{ if eq $i 1 }}[Blog Platform](/blog/week-01-blog-platform/){{ else }}—{{ end }} |
{{ end }}
