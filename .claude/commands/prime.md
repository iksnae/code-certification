---
description: Quick-start agent understanding of the Code Certification System
---

# Purpose

Orient an agent to the Code Certification System — a Go CLI that evaluates code against versioned policies and assigns time-bound certification status to meaningful code units.

## Project Overview

The Code Certification System is a **repository governance platform** built as a standalone **Go CLI** (`certify`). It discovers certifiable code units, evaluates them against versioned policy packs, collects evidence from linters/type-checkers/test-runners, optionally performs agent-assisted review via OpenRouter, and assigns time-bound certification status with expiration.

**Current State**: Pre-implementation. Documentation and tooling only. No Go source code yet.

**Development Approach**: Strict TDD — every feature is implemented test-first (test → implement → refactor).

## Product Documentation

| Document | Purpose |
|----------|---------|
| `PRD.md` | Full product requirements (27 sections) |
| `FEATURES.md` | Feature acceptance checklist (200+ measurable criteria) |
| `STORIES.md` | User stories organized by 15 epics |
| `specs/` | Implementation plans and architecture specs |

## Key Concepts

- **Certification Unit**: Smallest meaningful code element (function, method, class, module, file)
- **Certification Record**: Structured trust status with policy version, evidence, scores, dates
- **Policy Pack**: Versioned rule sets defining expectations (coding standards, security, architecture)
- **Evidence**: Data from tools (lint, type-check, test, static analysis, git history, agent review)
- **Certification Status**: `certified`, `certified_with_observations`, `probationary`, `expired`, `decertified`, `exempt`
- **Certification Expiry**: Trust windows that decay over time based on risk, churn, and stability
- **Certification Dimensions**: correctness, maintainability, readability, testability, security, architectural_fitness, operational_quality, performance_appropriateness, change_risk

## Target Architecture

```
certify CLI (Go binary)
├── cmd/certify/              CLI entry point
├── internal/engine/          Certification pipeline orchestrator
├── internal/policy/          Policy loading and evaluation
├── internal/evidence/        Evidence collection and normalization
├── internal/agent/           OpenRouter agent-assisted review (optional)
├── internal/discovery/       Code unit discovery + language adapters
├── internal/record/          Certification record management
└── internal/report/          Report generation

.certification/               Repository-local state (in target repos)
├── config.yml                Configuration
├── policies/                 Policy packs
├── prompts/                  Agent review prompt templates
├── units/index.json          Unit inventory
├── records/                  Certification records
└── reports/                  Generated reports
```

## Build & Test

```bash
just doctor       # Verify environment
just build        # Build CLI → build/bin/certify
just test         # Run all tests
just lint         # golangci-lint
just check        # All quality gates (fmt + vet + lint + test)
just run -- help  # Run CLI with args
```

## Agent-Assisted Review

Optional LLM-powered review via **OpenRouter** (free-tier models). See `specs/project-bootstrap-justfile-commands-initial-commit.md` for complete architecture:

- 7 free models researched with capabilities (tool calling, structured output, reasoning)
- Task-to-model routing: prescreen → review → scoring → decision → remediation
- Provider abstraction interface for swappable backends
- `OPENROUTER_API_KEY` in GitHub repo secrets; graceful degradation when absent

## GitHub Issue Tracking

Implementation is tracked via **8 GitHub epic issues** (#1–#8), one per phase. Each epic includes a step checklist and FEATURES section mapping. See `CLAUDE.md` for the full table.

**Development workflow for every change**:
1. Reference the relevant epic issue (#1–#8)
2. TDD cycle: failing test → implement → refactor
3. Update `FEATURES.md`: check off completed criteria (`- [ ]` → `- [x]`)
4. Update the epic: check off completed steps in the GitHub issue
5. Commit with `Refs #N` to link to the epic
6. Push to keep GitHub and code in sync

## Workflow

1. Read `CLAUDE.md` for project context and tracking guidance
2. Read `PRD.md` for full product requirements
3. Read `FEATURES.md` for acceptance checklist (living document — update as you implement)
4. Read `STORIES.md` for user stories
5. Check `specs/` for implementation plans
6. Check GitHub issues (`gh issue list`) for current epic status
7. Run `just doctor` to verify tooling
8. Explore Go source (when it exists)

## Report

Provide a summary covering:
- **Current state**: Pre-implementation, documentation only
- **Architecture**: Go CLI, policy-as-code, OpenRouter agent review
- **Development approach**: TDD — test-first for every feature
- **Recommended next steps**: Based on user's task
