---
title: Architecture
description: How Certify is built — packages, data flow, and design decisions.
---

## System Design

```
┌────────────────────────────────────────────────────────────────────────┐
│                          CLI (cmd/certify/)                            │
│  init │ scan │ certify │ report │ expire │ architect │ doctor │ onboard │
└────┬──────┬──────┬────────┬────────┬────────┬────────┬─────────┬──────┘
     │      │      │        │        │        │        │         │
┌────▼──────▼──────▼────────▼────────▼────────▼────────▼─────────▼──────┐
│                         Internal Packages                             │
│  discovery │ evidence │ engine │ agent  │ report │ record │ workspace │
│  config   │ policy   │ expiry │ queue  │ github │ override │ doctor  │
└───────────────────────────────────────────────────────────────────────┘
```

## Package Map

| Package | Responsibility |
|---------|---------------|
| `cmd/certify/` | CLI entry point — Cobra commands, workspace dispatch |
| `internal/analysis/` | Unified analyzer interface — Go (`go/ast`), TS/Py/Rs (tree-sitter), deep Go analysis (`go/packages` + SSA + VTA call graph) |
| `internal/analysis/lsp/` | LSP JSON-RPC 2.0 client — fan-in/fan-out/dead code for TS/Py/Rs via language servers |
| `internal/discovery/` | Language-aware unit discovery (Go AST, TS/Py/Rs tree-sitter, generic) |
| `internal/evidence/` | Evidence collection — lint, test, git, complexity, structural AST analysis |
| `internal/engine/` | Certification pipeline — scoring across 9 dimensions, status assignment |
| `internal/agent/` | LLM-assisted review, architect review pipeline, workspace snapshots |
| `internal/report/` | Report generation — card, full, badge, health, site, report tree, architect report |
| `internal/record/` | Record persistence — JSON files, history tracking, state snapshots |
| `internal/config/` | Configuration loading, validation, policy pack parsing |
| `internal/policy/` | Policy evaluation — rule matching, path scoping, threshold checking |
| `internal/expiry/` | Time-bound certification window calculation |
| `internal/queue/` | Persistent work queue for incremental processing |
| `internal/github/` | GitHub integration — PR comments, issue creation, workflow generation |
| `internal/override/` | Human governance overrides (exempt, force-certify) |
| `internal/workspace/` | Multi-repo workspace support — submodule discovery, aggregation, reporting |
| `internal/doctor/` | Health checks (doctor) and onboarding plan |
| `internal/domain/` | Core types — UnitID, Status, Grade, Dimension, Evidence, Policy, Config |

## Data Flow

### Discovery
```
Repository → Scanner → Language Adapters → index.json
```

The scanner walks the filesystem, detects languages, and dispatches to adapters. Go adapter uses `go/ast` for precise symbol extraction. TypeScript, Python, and Rust use tree-sitter for symbol-level discovery. Everything else gets file-level units. Nested module roots (`go.mod`, `package.json`, `Cargo.toml`, `pyproject.toml` in subdirectories) are automatically detected.

### Evidence Collection
```
Unit → Tool Executor → Evidence Items
```

For each unit, multiple evidence collectors run:
- **Lint** — `go vet`, `golangci-lint`, ESLint, ruff, cargo clippy findings attributed to specific units
- **Test** — `go test`, Jest/Vitest, pytest, cargo test results with per-unit coverage
- **Git** — Change frequency, author count, file age from `git log`
- **Structural** — 27+ AST metrics: complexity, error handling, nesting, security patterns, documentation
- **Deep Analysis** — Type-aware cross-file: fan-in/fan-out (call graph), dead code, dep depth, instability, interface compliance, coupling (Go built-in; TS/Py/Rs via optional LSP)
- **Metrics** — Code lines, comment lines, cyclomatic complexity, TODO count

### Certification
```
index.json → Queue → Evidence → Policy Evaluation → Scoring → Records
```

The work queue enables incremental processing — interrupted runs resume where they left off. The engine evaluates evidence against policy packs (with path scoping), scores across 9 dimensions, and assigns certification status with expiration dates.

### Reporting
```
Records → Report Generator → REPORT_CARD.md + badge.json + reports/ + site/
```

Reports come in multiple formats:
- **Card** — Terminal summary with grade and key metrics
- **Full** — Complete per-unit markdown report
- **Report Tree** — Per-unit markdown files in `reports/` directory
- **Site** — Static searchable HTML site for large repos
- **JSON** — Machine-readable for tooling

### Architect Review
```
Records → Architecture Snapshot → 6-Phase LLM Pipeline → ARCHITECT_REVIEW.md
```

The architect review builds a deterministic snapshot (package graph, structural metrics, coverage data, dependencies) then runs a 6-phase LLM analysis pipeline. See [Architect Review →](/code-certification/advanced/architect-review/).

### Workspace
```
Submodules → Per-Submodule Certification → Workspace Aggregation → Workspace Report
```

In workspace mode, each submodule is certified independently. Results are aggregated into a workspace-level report card. The workspace architect review treats all submodules as components of a single system. See [Workspace Mode →](/code-certification/advanced/workspace/).

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

// Evidence collected per unit
type Evidence  // kind, source, metrics map, severity
```

## Storage

All state lives in `.certification/`:

```
.certification/
├── config.yml              # configuration
├── policies/               # policy pack YAML files
│   ├── go-standard.yml
│   └── go-library.yml
├── records/                # per-unit JSON records
│   ├── abc123.json
│   └── abc123.history.jsonl
├── overrides/              # human governance overrides
├── reports/                # per-unit markdown report tree
│   ├── index.md
│   └── internal/engine/scorer.go/
│       └── Score.md
├── state.json              # fast-load state snapshot
├── runs.jsonl              # certification run history
├── REPORT_CARD.md          # generated report card
├── ARCHITECT_REVIEW.md     # architect review (if run)
├── badge.json              # shields.io badge endpoint
└── site/                   # static HTML site (if generated)
```

Records use SHA256 hash of UnitID for filenames — flat directory, no nesting. History is append-only JSON-lines for trend tracking. The `state.json` snapshot enables fast loading without scanning the records directory.
