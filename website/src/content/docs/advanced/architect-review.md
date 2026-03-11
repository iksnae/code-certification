---
title: Architect Review
description: AI-powered architectural analysis with grounded data and comparative recommendations.
---

The `certify architect` command runs a comprehensive, 6-phase AI-driven architectural review of your codebase. It builds a deterministic architecture snapshot from certification data, then uses an LLM to analyze structure, quality patterns, test strategy, security posture, and generate prioritized recommendations with projected impact.

## Quick Start

```bash
# First, certify your codebase (needed for data)
certify certify

# Then run the architect review
certify architect
```

Output: `.certification/ARCHITECT_REVIEW.md`

## What It Produces

The architect review generates a detailed markdown report with:

- **Executive Summary** — 2–3 paragraph overview of codebase health
- **Part I: Architecture Snapshot** — Package map, dependency graph, layer classification, hotspots, coupling pairs (all deterministic, no LLM)
- **Part II: Analysis** — Code quality findings, test strategy assessment, security/operations concerns
- **Part III: Recommendations** — Each with current→projected deltas, affected units, effort estimates
- **Risk Matrix** — Prioritized risks with severity and likelihood
- **Roadmap** — Ordered improvement plan referencing specific recommendations

## The 6-Phase Pipeline

Each phase runs as a separate LLM call, with prior phase outputs fed forward as context:

| Phase | Name | What it does |
|-------|------|-------------|
| 1 | Architecture Narration | Describes the as-is architecture — layers, data flow, dependency health |
| 2 | Code Quality & Patterns | Identifies anti-patterns, complexity hotspots, error handling issues |
| 3 | Test Strategy & Coverage | Evaluates test approach, coverage gaps, missing test categories |
| 4 | Security & Operations | Assesses structural risk metrics, operational readiness, config management |
| 5 | Comparative Recommendations | Generates specific improvements with current→projected metric deltas |
| 6 | Synthesis & Roadmap | Executive summary, risk matrix, prioritized improvement roadmap |

### Run a Single Phase

```bash
certify architect --phase 4    # only security & operations
```

## Data Grounding

Every metric the LLM sees is **computed from real AST analysis and certification data**, not estimated. The architecture snapshot includes:

### Structural Metrics (from AST)
16 metrics aggregated across all units:
- `panic_calls`, `os_exit_calls`, `global_mutable_count`, `defer_in_loop`
- `init_func_count`, `context_not_first`, `errors_ignored`, `naked_returns`
- `recursive_calls`, `max_nesting_depth`, `nested_loop_pairs`, `quadratic_patterns`
- `total_func_lines`, `total_params`, `total_returns`, `total_methods`

### Deep Analysis Metrics (schema v3)
14 type-aware metrics from call graph and dependency analysis:
- `avg_fan_in`, `max_fan_in` — call site counts (change risk hotspots)
- `avg_fan_out`, `max_fan_out` — outgoing call counts (dependency coupling)
- `dead_export_count` — exported symbols with zero external references
- `concrete_deps_count` — function params accepting concrete external types
- `avg_cognitive_complexity`, `max_cognitive_complexity` — Sonar-style readability
- `errors_not_wrapped` — error returns without wrapping context
- `unsafe_import_count`, `hardcoded_secrets` — security signals
- `max_dep_depth` — deepest transitive import chain
- `avg_instability` — average package instability (0=stable, 1=unstable)

### Coverage Metrics
- Units with/without coverage data
- Average, min, max coverage percentages

### Code Metrics
- Total code lines, comment lines, TODOs
- Total/max/average cyclomatic complexity

### Package-Level Data
- Per-package unit count, average score, grade, observation count
- Dependency edges between internal packages
- Layer classification (cmd, internal, domain, pkg)
- Hotspot ranking by risk factor (units × (1 − avg_score))
- Coupling pairs (most cross-referenced package pairs)

### Anti-Hallucination

All 6 phase prompts include explicit grounding language:
- "Reference only data that appears in the snapshot"
- "Do not fabricate specific numeric values"
- "If a metric is not present, do not reference it"

The snapshot carries a `SchemaVersion` (currently v2) so reports can be distinguished from pre-grounding versions.

## Provider Setup

The architect command needs generous timeouts — local 30B+ models can take several minutes per phase. The timeout is automatically set to 10 minutes per phase.

### Use Config

```yaml
# .certification/config.yml
agent:
  enabled: true
  provider:
    base_url: https://openrouter.ai/api/v1
    api_key_env: OPENROUTER_API_KEY
  models:
    review: qwen/qwen3-coder-30b
```

### Override Model

```bash
certify architect --model gpt-4o
certify architect --model qwen/qwen3-coder-30b
```

### Auto-Detection

If `OPENROUTER_API_KEY`, `OPENAI_API_KEY`, or a local provider (Ollama, LM Studio) is available, the architect command auto-detects and uses it.

## Workspace Mode

```bash
certify architect --workspace
```

In workspace mode, the architect review analyzes all submodules as components of a single system:

- **Per-submodule snapshots** built from each submodule's certification records
- **Cross-submodule dependencies** detected from `go.mod` replace directives
- **Submodule role classification** — service, library, tool (based on filesystem heuristics)
- **Infrastructure files** — workspace-level Justfile, Makefile, CI workflows, Docker configs
- **System-level prompts** — LLM reasons about integration boundaries, shared library quality, deployment coupling

The workspace review produces the same 6-phase report but with system-level perspective:

| Phase | Workspace Focus |
|-------|----------------|
| 1 | System composition — what role each submodule plays |
| 2 | Cross-submodule quality patterns, shared library impact |
| 3 | Integration testing gaps, contract testing between services |
| 4 | Cross-submodule security surface, deployment coupling |
| 5 | System-level improvements, boundary changes |
| 6 | Workspace-level executive summary |

## Verbose Output

```bash
certify architect --verbose
```

Prints the full raw LLM response for each phase, including chain-of-thought reasoning from models that support it (e.g., Qwen3's `<think>` tags).

## Example Output

```markdown
# 🏗 Architectural Review — iksnae/code-certification

**Generated:** 2026-03-11 12:00 · **Model:** qwen/qwen3-coder-30b · **Tokens:** 24,391 · **Phases:** 6/6

## Executive Summary

The project demonstrates a well-structured layered architecture with clear
separation between CLI, engine, and domain layers...

## Part I — Architecture (Deterministic)

### Package Map
| Package | Units | Avg Score | Grade | Observations |
|---------|------:|----------:|:-----:|-------------:|
| internal/agent | 142 | 91.2% | A- | 4 |
| internal/engine | 89 | 93.1% | A | 0 |
...

## Part III — Recommendations

### 1. Extract Error Handling Middleware
| Metric | Current | Projected |
|--------|---------|-----------|
| errors_ignored | 5 | 0 |
| avg_score (engine/) | 93.1% | 95.2% |
**Effort:** S · **Affected:** 3 units
```
