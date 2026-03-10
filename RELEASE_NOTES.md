# Release v0.6.0

**Date:** 2026-03-10

## Highlights

New `certify architect` command — an AI-powered architectural review that produces a comprehensive report with a deterministic architecture snapshot, 6-phase LLM analysis, and comparative recommendations with current→projected metric deltas.

## What's Changed

### New Features

- **feat: `certify architect` command** — AI-driven architectural review of an entire project. Builds a deterministic architecture snapshot from certification data (package graph, dependency analysis, hotspots, coupling), then runs 6 LLM phases producing a bespoke report with comparative recommendations. Output: `.certification/ARCHITECT_REVIEW.md`.

- **feat: architecture snapshot** — `BuildSnapshot()` computes a complete structural model from certification records: package metrics, Go import dependency graph, layer classification, hotspot ranking (`units × (1 - score)`), coupling pair analysis. Fully deterministic — same records always produce identical output.

- **feat: 6-phase review pipeline** — Sequential LLM phases with feed-forward context:
  1. Architecture Narration (as-is description, no recommendations)
  2. Code Quality & Patterns (findings citing snapshot metrics)
  3. Test Strategy & Coverage (coverage gaps per package)
  4. Security & Operations (global state, error handling, init funcs)
  5. Comparative Recommendations (current→projected deltas)
  6. Synthesis & Roadmap (executive summary, risk matrix, prioritized roadmap)

- **feat: comparative recommendation format** — Every recommendation includes a delta table (`| Metric | Current | Projected | Delta |`), affected units, effort estimate, and justification grounded in snapshot data. Phase 5 validates that all recommendations have deltas.

- **feat: chain-of-thought capture** — `<think>` tags from reasoning models (qwen3, etc.) are captured per-phase in `result.Thinking[]` and rendered in the report under collapsible "🧠 Agent Reasoning" sections.

- **feat: 3-part report structure** —
  - Part I: Architecture Snapshot (deterministic tables from data, always present even if LLM fails)
  - Part II: Analysis (LLM narrative grounded in snapshot numbers)
  - Part III: Recommendations (comparative before/after with delta tables)

### Bug Fixes

- **fix: Phase 4 JSON parsing** — `ArchConcern.Metrics` changed from `map[string]string` to `map[string]any` to handle numeric metric values from LLM responses.
- **fix: think tag interference** — Qwen3 `<think>` blocks containing braces confused `extractJSON`. Added `stripThinkTags()` before JSON extraction.
- **fix: timeout for local models** — Added `SetTimeout()` on `OpenRouterProvider`. Architect command sets 10-minute timeout for local models that need several minutes per phase.

### CLI

```
certify architect                    # full 6-phase review
certify architect --model gpt-4o     # use specific model
certify architect --phase 1          # run only architecture narration
certify architect --verbose          # print full LLM responses
certify architect --output FILE      # custom output path
```

## New Files

| File | Purpose |
|------|---------|
| `cmd/certify/architect_cmd.go` | CLI command + provider setup |
| `internal/agent/architect_snapshot.go` | ArchSnapshot, BuildSnapshot, import analysis |
| `internal/agent/architect.go` | ProjectContext, GatherContext, FormatForLLM |
| `internal/agent/architect_review.go` | 6-phase orchestrator, response types |
| `internal/agent/architect_prompts.go` | Phase system prompts |
| `internal/report/architect_report.go` | Report formatter (Part I/II/III) |

## Stats

- **9 new files**, 2,837 lines added
- **23 new tests** across 3 test files, all passing
- **16 packages** pass with zero regressions
- Dogfood: 615 units, 24 packages, 6/6 phases, 7 recommendations, 50K tokens, 3m46s

## Full Changelog

```
3f696eb chore: architect dogfood — 6/6 phases, 7 recommendations, 615 units reviewed
6f62d32 fix: Phase 4 JSON parsing — ArchConcern.Metrics map[string]any
2d2d6ef fix: architect — 10min timeout, 8K-12K token limits, think tag capture
1d6d4ea feat: architect E2E integration test
ec3f3ff feat: certify architect CLI command
7c28fa9 feat: architect report formatter — Part I/II/III
cf63155 feat: architect review pipeline — 6-phase orchestrator
108abc3 feat: architect snapshot + context gathering
```
