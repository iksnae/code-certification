# Certify — Contributor Guide

## Overview

Go CLI that continuously evaluates every code unit in a repository, scores it against versioned policies, and assigns time-bound certification you can actually trust.

**Module:** `github.com/iksnae/code-certification`  
**Binary:** `certify` (built from `cmd/certify/`)  

## Build & Test

```bash
just build          # → build/bin/certify
just test           # run all tests
just check          # fmt + vet + lint + test
just certify        # self-certify this repo
just report-card    # generate .certification/REPORT_CARD.md
```

Or directly:

```bash
go build -o build/bin/certify ./cmd/certify/
go test ./... -count=1
```

## Package Structure

```
cmd/certify/          CLI entry point (cobra commands)
internal/
  agent/              LLM-assisted review via OpenRouter
  config/             Configuration loading + validation
  discovery/          Language-aware unit discovery (Go, TS, generic)
  domain/             Core types: UnitID, Status, Grade, Dimension, etc.
  engine/             Certification pipeline + scoring
  evidence/           Evidence collectors (lint, test, git, complexity)
  expiry/             Time-bound trust window calculation
  github/             GitHub integration (PR comments, issues, workflows)
  override/           Human governance overrides
  policy/             Policy-as-code evaluation + matching
  queue/              Persistent work queue
  record/             Record store (JSON + history)
  report/             Report generation (card, full, badge)
```

## Development Rules

1. **TDD** — Write test first, then implement, then refactor.
2. **`gofmt`** — All code must be formatted. CI checks this.
3. **Module imports** — Use `github.com/iksnae/code-certification/internal/...`.
4. **Commit messages** — `feat:`, `fix:`, `chore:`, `docs:` prefixes.
5. **No external test frameworks** — stdlib `testing` only.

## Key Types

- `domain.UnitID` — `lang://path#symbol` (e.g., `go://internal/engine/scorer.go#Score`)
- `domain.Status` — `certified | probationary | decertified | expired | exempt`
- `domain.Grade` — `A | A- | B+ | B | C | D | F` from weighted dimension scores
- `domain.DimensionScores` — 9 quality dimensions: correctness, maintainability, readability, testability, security, architectural_fitness, operational_quality, performance_appropriateness, change_risk

## CLI Flow

```
certify init     → bootstrap .certification/ in target repo
certify scan     → discover units → .certification/index.json
certify certify  → collect evidence → evaluate → store records
certify report   → generate report card + badge
certify expire   → mark overdue records as expired
```
