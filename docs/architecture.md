# Architecture

## Overview

Certify is a Go CLI that continuously evaluates code units against versioned policies and assigns time-bound certification status.

```
┌─────────────────────────────────────────────────────────────┐
│                          CLI (cmd/certify/)                          │
│  init │ scan │ certify │ report │ expire │ architect │ doctor │ …    │
└────┬──────┬──────┬────────┬────────┬────────┬──────────────┘
     │      │      │        │        │        │
┌────▼──────▼──────▼────────▼────────▼────────▼──────────────┐
│                    Internal Packages                        │
│                                                             │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌───────────┐  │
│  │ discovery │  │ evidence │  │  engine  │  │   agent   │  │
│  │ (scan)    │  │ (collect)│  │ (certify)│  │ (review)  │  │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  └─────┬─────┘  │
│       │              │             │              │         │
│  ┌────▼─────┐  ┌─────▼────┐  ┌────▼─────┐  ┌────▼──────┐  │
│  │  config  │  │  policy  │  │  record  │  │  report   │  │
│  │ (load)   │  │ (eval)   │  │  (store) │  │ (format)  │  │
│  └──────────┘  └──────────┘  └──────────┘  └───────────┘  │
│                                                             │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌───────────┐  │
│  │  domain  │  │  expiry  │  │ override │  │  github   │  │
│  │ (types)  │  │ (calc)   │  │ (apply)  │  │ (actions) │  │
│  └──────────┘  └──────────┘  └──────────┘  └───────────┘  │
│                                                             │
│  ┌──────────┐                                               │
│  │  queue   │                                               │
│  │ (persist)│                                               │
│  └──────────┘                                               │
└─────────────────────────────────────────────────────────────┘
```

## Package Responsibilities

### `internal/domain/`
Core types with zero external dependencies. Defines: Unit, UnitID, Evidence, PolicyPack, CertificationRecord, Status, Grade, Dimension, Override, Config.

### `internal/analysis/`
Unified code analysis interface. Go analyzer wraps `go/ast` for 27 structural metrics. TypeScript, Python, and Rust analyzers use tree-sitter for 22+ metrics. `DeepGoAnalyzer` uses `go/packages` + SSA + VTA for type-aware call graph analysis: fan-in/fan-out, dead code detection, dependency depth, instability, interface compliance, unused params.

### `internal/analysis/lsp/`
Generic JSON-RPC 2.0 LSP client for type-aware analysis of non-Go languages. Communicates with `typescript-language-server`, `pyright`, `rust-analyzer` via stdin/stdout. Provides call hierarchy (fan-in/fan-out), references (dead code), and diagnostics.

### `internal/discovery/`
Code unit discovery. Language adapters (Go via `go/ast`, TS/Py/Rs via tree-sitter, generic file-level). Nested module root detection. Index management with JSON persistence. Diff computation for added/removed/changed units.

### `internal/evidence/`
Evidence collection from external tools. Runs `go vet`/`golangci-lint`, `go test`, ESLint, ruff, pytest, cargo clippy/test, `git log`. Parses results into normalized `Evidence` structs. Module root discovery for multi-language repos.

### `internal/config/`
Loads YAML config and policy packs. Policy matching (language, path targeting). Config validation.

### `internal/policy/`
Evaluates policy rules against evidence. Extracts metrics (lint errors, test failures, complexity, TODOs). Records violations with severity and dimension.

### `internal/engine/`
Certification pipeline: evaluate → score → status → grade → expiry → record. Multi-dimension scoring across 9 quality dimensions. Weighted averages.

### `internal/agent/`
Optional LLM-assisted review via OpenRouter. Pipeline architecture: Prescreen → Review → Scoring stages. Circuit breaker, model chain with fallback, adaptive messages. Model attribution tracking.

### `internal/expiry/`
Computes certification expiry windows based on risk factors: churn, complexity, coverage, security sensitivity, prior pass/fail history.

### `internal/record/`
Flat JSON file store for certification records. SHA256-hashed filenames for deterministic storage.

### `internal/report/`
Report generation in multiple formats. Health reports (summary + detailed). Full reports with dimension breakdowns, by-language analysis, expiring-soon detection, highest-risk identification. Card reports for terminal display. Static HTML site generation with dashboard, per-package/unit pages, and client-side search. Badge generation for shields.io. Uses `LanguageDetail` as the unified language summary type across all formats.

### `internal/override/`
Manual governance: exempt, extend/shorten windows, force review. YAML-based override definitions with required rationale.

### `internal/github/`
GitHub integration: PR comment formatting, issue creation/close commands, workflow YAML generation.

### `internal/queue/`
Persistent JSON-backed work queue for incremental processing across runs. Crash-safe (saves after each item).

## Data Model

### Configuration
```
.certification/
├── config.yml          # Mode, scope, agent, expiry, issues
├── policies/           # Versioned YAML policy packs
│   ├── global.yml
│   └── go-standard.yml
├── overrides/          # Manual exemptions/adjustments
├── records/            # Certification record JSON files (gitignored)
├── state.json          # Full state snapshot (tracked in git)
├── runs.jsonl          # Certification run history (tracked in git)
├── index.json          # Discovered unit index
├── queue.json          # Processing queue state (gitignored)
├── REPORT_CARD.md      # Markdown report card (tracked)
├── badge.json          # Shields.io badge endpoint (tracked)
├── reports/            # Per-unit markdown reports (gitignored)
└── site/               # Interactive HTML report site (gitignored)
```

### Certification Record
Each record contains: unit identity, policy version, run ID, status, grade, score, confidence, dimension scores, evidence details (as JSON), observations, actions, timestamps (certified_at, expires_at), source attribution.

### Certification State
- `state.json` — Snapshot of all records + runs, tracked in git for post-clone completeness
- `runs.jsonl` — Append-only JSONL log of certification runs with overall grade/score
- `records/` — Individual JSON files per unit (gitignored, derived from state)
- `reports/` — Per-unit markdown reports (gitignored, regenerated on demand)
- `site/` — Interactive HTML report (gitignored, regenerated on demand)

## Design Principles

1. **Language-agnostic core**: Domain types have no language-specific assumptions
2. **Adapter pattern**: Language support added via Scanner interface implementations
3. **Deterministic first**: Agent review is optional and additive — never overrides deterministic evidence
4. **Repository-native**: All state lives in `.certification/` within the target repo
5. **Incremental**: Queue-based processing across multiple runs
6. **Auditable**: Records are JSON, versioned in git, human-reviewable
