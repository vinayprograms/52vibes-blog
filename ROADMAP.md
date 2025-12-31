# 52vibes: Draft Project List

> **Status:** Draft - subject to revision based on prior week results
> **Last Updated:** December 30, 2024

---

## Q1: Agentic Infrastructure (Weeks 1-13)

Building the foundation that all future weeks depend on.

| Week | Dates | Project | Artifact | Tests Capability |
|------|-------|---------|----------|------------------|
| 1 | Jan 2-8 | Static Blog Platform | Hugo + theme + CI/CD | File generation, config, deployment |
| 2 | Jan 9-15 | Session Logger | Go package for conversation capture | Structured data, file I/O |
| 3 | Jan 16-22 | CLI Framework | Go CLI scaffold with plugin architecture | Code generation patterns |
| 4 | Jan 23-29 | Error Taxonomy | Package for categorizing agent failures | Self-reflection, classification |
| 5 | Jan 30-Feb 5 | Prompt Library | Versioned prompt management system | Template handling, context |
| 6 | Feb 6-12 | MCP Server Template | Reusable MCP server scaffold | Protocol implementation |
| 7 | Feb 13-19 | Eval Harness | Testing framework for agent outputs | Test generation, assertions |
| 8 | Feb 20-26 | Context Manager | Long-context handling utilities | Memory, summarization |
| 9 | Feb 27-Mar 5 | Config Generator | YAML/TOML/JSON config scaffolding | Schema inference |
| 10 | Mar 6-12 | Git Automation | Git workflow automation library | Multi-step operations |
| 11 | Mar 13-19 | Documentation Generator | Auto-generate docs from code | Code understanding |
| 12 | Mar 20-26 | API Client Generator | Generate typed clients from OpenAPI | Spec parsing, code gen |
| 13 | Mar 27-Apr 2 | **Q1 Integration** | Combine Q1 tools into cohesive toolkit | System integration |

---

## Q2: Production Tools (Weeks 14-26)

Real-world tools that stress-test Q1 infrastructure.

| Week | Dates | Project | Artifact | Tests Capability |
|------|-------|---------|----------|------------------|
| 14 | Apr 3-9 | Code Review Assistant | Review automation library + CLI | Code analysis, feedback |
| 15 | Apr 10-16 | Test Generator | Generate tests from implementation | Test reasoning |
| 16 | Apr 17-23 | Migration Assistant | Database/API migration helper | Breaking change detection |
| 17 | Apr 24-30 | Changelog Generator | Semantic changelog from commits | Commit understanding |
| 18 | May 1-7 | Dependency Auditor | Security/update checker | Vulnerability analysis |
| 19 | May 8-14 | Performance Profiler | Automated perf analysis | Metric interpretation |
| 20 | May 15-21 | Log Analyzer | Structured log parsing + insights | Pattern recognition |
| 21 | May 22-28 | Schema Validator | Multi-format schema validation | Constraint reasoning |
| 22 | May 29-Jun 4 | Release Manager | Automated release workflow | Multi-step coordination |
| 23 | Jun 5-11 | Incident Reporter | Postmortem generation | Causal reasoning |
| 24 | Jun 12-18 | Runbook Generator | Operational runbooks from code | Operational knowledge |
| 25 | Jun 19-25 | Cost Estimator | Cloud cost estimation | Numerical reasoning |
| 26 | Jun 26-Jul 2 | **Q2 Integration** | Production toolkit consolidation | System integration |

---

## Q3: Complex Workflows (Weeks 27-39)

Multi-step, multi-file, human-in-the-loop challenges.

| Week | Dates | Project | Artifact | Tests Capability |
|------|-------|---------|----------|------------------|
| 27 | Jul 3-9 | Refactoring Engine | Large-scale refactoring library | Multi-file coordination |
| 28 | Jul 10-16 | Full-Stack Scaffold | Complete app generator | Architecture decisions |
| 29 | Jul 17-23 | Workflow Orchestrator | Multi-agent task coordination | Agent collaboration |
| 30 | Jul 24-30 | Approval System | Human-in-the-loop workflows | Handoff protocols |
| 31 | Jul 31-Aug 6 | Recovery Manager | Graceful failure recovery | Error handling |
| 32 | Aug 7-13 | State Machine Builder | Complex state management | State reasoning |
| 33 | Aug 14-20 | Contract Testing | Cross-service contract tests | Interface understanding |
| 34 | Aug 21-27 | Feature Flag Manager | Gradual rollout system | Risk assessment |
| 35 | Aug 28-Sep 3 | A/B Test Framework | Experiment infrastructure | Statistical reasoning |
| 36 | Sep 4-10 | Data Pipeline Builder | ETL pipeline generator | Data flow reasoning |
| 37 | Sep 11-17 | Compliance Checker | Policy/regulation validation | Rule interpretation |
| 38 | Sep 18-24 | Capacity Planner | Resource planning tool | Forecasting |
| 39 | Sep 25-Oct 1 | **Q3 Integration** | Complex workflow toolkit | System integration |

---

## Q4: Synthesis & Edge Cases (Weeks 40-52)

Pushing boundaries and consolidating learnings.

| Week | Dates | Project | Artifact | Tests Capability |
|------|-------|---------|----------|------------------|
| 40 | Oct 2-8 | Ambiguity Handler | Unclear requirement resolution | Clarification strategies |
| 41 | Oct 9-15 | Context Recovery | Resume interrupted sessions | Long-term memory |
| 42 | Oct 16-22 | Multi-Language Support | Cross-language tooling | Language switching |
| 43 | Oct 23-29 | Legacy Modernizer | Legacy code updater | Old pattern recognition |
| 44 | Oct 30-Nov 5 | Security Analyzer | Vulnerability detection | Security reasoning |
| 45 | Nov 6-12 | Accessibility Auditor | A11y compliance checker | Standards interpretation |
| 46 | Nov 13-19 | Internationalization | i18n automation | Cultural reasoning |
| 47 | Nov 20-26 | Edge Case Generator | Boundary condition finder | Edge case reasoning |
| 48 | Nov 27-Dec 3 | Benchmark Suite | Comprehensive agent benchmarks | Self-evaluation |
| 49 | Dec 4-10 | Recommendation Engine | Collaboration recommendations | Meta-analysis |
| 50 | Dec 11-17 | Report Generator | Automated report creation | Synthesis |
| 51 | Dec 18-24 | **Final Report Draft** | Year-end report v1 | Comprehensive synthesis |
| 52 | Dec 25-31 | **Final Report & Retrospective** | Published findings | Reflection |

---

## Dependencies Graph (Key Connections)

```
Week 1 (Blog) ──────────────────────────────────────► All weekly posts
Week 2 (Logger) ────► Week 7 (Eval) ────► Week 48 (Benchmark)
Week 3 (CLI) ────► Weeks 14-26 (all CLI tools)
Week 4 (Error Taxonomy) ────► Week 31 (Recovery) ────► Week 51 (Report)
Week 5 (Prompts) ────► Week 8 (Context) ────► Week 41 (Context Recovery)
Week 6 (MCP) ────► Week 29 (Orchestrator)
Week 10 (Git) ────► Week 17 (Changelog) ────► Week 22 (Release)
```

---

## Review Schedule

- **Weekly:** Adjust next week based on current week results
- **Monthly:** Review month's projects, update upcoming month
- **Quarterly:** Major revision of remaining quarters

---

## Notes

- Projects are intentionally ambitious; scope will be cut to fit time
- "Integration" weeks (13, 26, 39) are buffers for catching up or consolidating
- Q4 final weeks prioritize documentation over new features
- Holiday weeks (Thanksgiving, Christmas) have lighter technical scope
