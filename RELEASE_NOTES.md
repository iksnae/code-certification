# Release v0.12.0

**Date:** 2026-03-11

## Highlights

**Deep analysis for Go and multi-language LSP support.** Certify now performs type-aware cross-file analysis — call graphs, fan-in/fan-out, dead code detection, dependency graphs, interface compliance, and coupling metrics. Go gets built-in deep analysis via `go/packages` + SSA + VTA; TypeScript, Python, and Rust gain Tier 2 analysis via optional LSP servers. The new `certify doctor` output shows analysis tiers per language.

## What's New

### Deep Go Analysis (Sprint 7)
- **Call graph via SSA/VTA** — fan-in (callers) and fan-out (callees) per function
- **Dead export detection** — exported symbols with zero external references
- New metrics: `fan_in`, `fan_out`, `is_dead_code`
- Built on `golang.org/x/tools` (go/packages, SSA, VTA) — no CGo required

### Interface Compliance & Dependency Graph (Sprint 8)
- **Dependency depth** — transitive local import chain depth
- **Instability metric** — Robert C. Martin's Ce/(Ca+Ce)
- **Parameter abstraction** — detect concrete external struct params (hard to mock)
- **Coupling score** — fan_in × fan_out normalized
- New metrics: `dep_depth`, `instability`, `concrete_deps`, `coupling_score`

### Type-Aware Refinement (Sprint 9)
- **Unused parameters** — function params never referenced in body
- **Interface size** — ISP violation detection (large interfaces)
- **Type-aware error wrapping** — verify errors returned without `fmt.Errorf("%w")`
- New metrics: `unused_params`, `interface_size`, `type_aware_unwrapped`

### LSP Client Infrastructure (Sprint 10)
- JSON-RPC 2.0 client with Content-Length framing
- LSP types: initialize, textDocument/documentSymbol, textDocument/references, callHierarchy, textDocument/publishDiagnostics

### Multi-Language LSP Analysis (Sprint 11)
- **LSP Analyzer** — fan-in/fan-out/dead-code via callHierarchy for any language
- Auto-detect: `typescript-language-server`, `pyright`, `rust-analyzer`
- Graceful degradation: returns `nil, nil` when server not installed

### Architect Snapshot v3 (Sprint 12)
- **14 deep analysis aggregates** in architect snapshot: avg/max fan-in, avg/max fan-out, dead export count, concrete deps, cognitive complexity, error wrapping, unsafe imports, secrets, dep depth, instability
- Schema bumped to v3 — LLM prompts now reference deep analysis data

### Doctor Analysis Tiers (Sprint 13)
- `certify doctor` reports per-language analysis tier (0/1/2)
- Tier upgrade hints with install commands for LSP servers
- New "Analysis Tiers" section in doctor output

### Documentation Alignment
- All website pages updated with new metrics, analysis tiers, deep analysis
- Available Metrics table expanded from 12 to 35+ entries
- Architecture docs include `internal/analysis/` and `internal/analysis/lsp/`
- VSCode extension guide updated for v0.5.0 features

### VSCode Extension v0.5.0
- CodeLens for Python, Rust, JavaScript (in addition to Go, TypeScript)
- Deep analysis dashboard section with fan-in hotspots table
- Dead code + high fan-in diagnostics in Problems panel
- `Certify: Run Doctor` command
- 12+ language detection patterns

### Scoring & Policy Updates
- `go-standard` policy bumped to v1.4.0 with 4 new rules: `max-fan-out` (15), `max-fan-in` (20), `no-dead-exports`, `max-dep-depth` (8)
- Deep analysis metrics scored into: maintainability (fan-out, dead code), testability (concrete deps), arch fitness (dep depth, instability), change risk (fan-in), operational quality (error wrapping)
- Multi-language policy packs: `ts-standard`, `python-standard`, `rust-standard`
- Lint/test tool integration for ESLint, ruff, pytest, cargo clippy/test

## New Dependencies

- `golang.org/x/tools v0.42.0` (go/packages, SSA, VTA)
- `golang.org/x/mod v0.33.0`, `golang.org/x/sync v0.19.0` (transitive)

## Upgrade Notes

- **Backward compatible** — all changes are additive. Existing records, policies, and configs continue to work.
- **Re-run recommended** — run `certify certify` after upgrading to collect deep analysis evidence.
- **LSP servers optional** — TS/Py/Rs Tier 2 analysis only activates when the language server is installed. Run `certify doctor` to check.
- **VSCode extension** — update to v0.5.0 for deep analysis dashboard. New features only appear when data is present.

## Files Changed

| Area | Key Files |
|------|-----------|
| Deep Go analysis | `internal/analysis/go_deep.go`, `go_deps.go`, `go_refine.go` + tests |
| LSP infrastructure | `internal/analysis/lsp/client.go`, `types.go` + tests |
| LSP analyzer | `internal/analysis/lsp_analyzer.go`, `lsp_config.go` + tests |
| Architect snapshot | `internal/agent/architect_snapshot.go` + tests |
| Doctor tiers | `internal/doctor/doctor.go` + tests |
| Scoring | `internal/engine/scorer.go`, `certifier.go` + tests |
| Policy | `.certification/policies/go-standard.yml` (v1.4.0) |
| VSCode extension | `vscode-certify/src/` (8 files) |
| Documentation | 11 doc files across repo root, `docs/`, `website/` |

## Test Results

All 19 Go packages pass. 0 vet/fmt issues. Website builds clean (21 pages).
