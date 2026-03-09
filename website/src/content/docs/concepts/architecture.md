---
title: Architecture
description: How Certify is built — packages, data flow, and design decisions.
---

## System Design

```
┌─────────────────────────────────────────────────────────────┐
│                     CLI (cmd/certify/)                       │
│  init │ scan │ certify │ report │ expire │ review │ version │
└────┬──────┬──────┬────────┬────────┬────────┬──────────────┘
     │      │      │        │        │        │
┌────▼──────▼──────▼────────▼────────▼────────▼──────────────┐
│                    Internal Packages                        │
│  discovery │ evidence │ engine │ agent │ report │ record    │
│  config    │ policy   │ expiry │ queue │ github │ override  │
└─────────────────────────────────────────────────────────────┘
```

## Package Map

| Package | Responsibility |
|---------|---------------|
| `cmd/certify/` | CLI entry point — Cobra commands |
| `internal/discovery/` | Language-aware unit discovery (Go AST, TS regex, generic) |
| `internal/evidence/` | Evidence collection (lint, test, git, complexity) |
| `internal/engine/` | Certification pipeline — scoring and status assignment |
| `internal/agent/` | LLM-assisted review via OpenRouter |
| `internal/report/` | Report generation (card, full, badge, health) |
| `internal/record/` | Record persistence (JSON files + history) |
| `internal/config/` | Configuration and policy loading |
| `internal/policy/` | Policy evaluation and matching |
| `internal/expiry/` | Time-bound certification window calculation |
| `internal/queue/` | Persistent work queue for incremental processing |
| `internal/github/` | GitHub integration (PR comments, issues, workflows) |
| `internal/override/` | Human governance overrides |
| `internal/domain/` | Core types — UnitID, Status, Grade, Dimension |

## Data Flow

### Discovery
```
Repository → Scanner → Language Adapters → index.json
```

The scanner walks the filesystem, detects languages, and dispatches to adapters. Go adapter uses `go/ast` for precise symbol extraction. TypeScript uses regex patterns. Everything else gets file-level units.

### Certification
```
index.json → Queue → Evidence Collection → Policy Evaluation → Scoring → Records
```

The work queue enables incremental processing — interrupted runs resume where they left off. Evidence is collected per-unit from linters, test runners, git history, and AST analysis. The engine evaluates evidence against policy packs, scores across 9 dimensions, and assigns status.

### Reporting
```
Records → Report Generator → REPORT_CARD.md + badge.json
```

The report aggregates all records into summary statistics, dimension averages, language breakdowns, and per-unit tables.

## Key Types

```go
// Every certifiable unit has a unique ID
type UnitID  // go://internal/engine/scorer.go#Score

// Certification status
type Status  // certified | probationary | decertified | expired | exempt

// Quality grade
type Grade   // A | A- | B+ | B | C | D | F

// Scores per dimension
type DimensionScores  // map[Dimension]float64
```

## Storage

All state lives in `.certification/`:

```
.certification/
├── config.yml        # configuration
├── policies/         # policy pack YAML files
├── records/          # per-unit JSON records
│   ├── abc123.json
│   └── abc123.history.jsonl
├── overrides/        # human governance overrides
├── REPORT_CARD.md    # generated report card
└── badge.json        # shields.io badge endpoint
```

Records use SHA256 hash of UnitID for filenames — flat directory, no nesting. History is append-only JSON-lines for trend tracking.
