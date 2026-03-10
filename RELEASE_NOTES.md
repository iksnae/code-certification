# Release v0.2.0

**Date:** 2026-03-10

## Highlights

Major architecture improvements across the certification engine, plus a new interactive HTML site report for browsing certification results at scale. All 10 architecture issues identified in the evaluation are now resolved. Three redundant types eliminated through report type consolidation. Derived report files are now gitignored — certification state is tracked via `state.json` for post-clone completeness.

## What's Changed

### New Features

- **feat: interactive HTML site report** — `certify report --site` generates a self-contained static HTML site under `.certification/site/` with a dashboard, per-package roll-ups, per-unit detail pages, and client-side search. Works offline via `file://`, zero external dependencies, dark mode support. At 559 units, generates 584 pages in under 2 seconds. ([#b5805f9](../../commit/b5805f9))

### Architecture Improvements

- **feat: persist evidence in records** — Evidence details stored as `json.RawMessage` in certification records. Records now carry full evidence context for auditability. ([#a23f933](../../commit/a23f933), Refs #9)

- **feat: extract Certifier service** — `internal/engine/certifier.go` with `Certifier` struct owning the full certification pipeline. CLI constructs it, calls `Certify()` per unit. Clean separation of orchestration from CLI concerns. ([#b617e53](../../commit/b617e53), Refs #10)

- **feat: typed evidence metrics** — `Metrics map[string]float64` on `Evidence`, replacing string-based metric extraction. All 5 evidence producers updated, evaluator rewritten for map lookup, 11 new tests. ([#df9fb30](../../commit/df9fb30), Refs #11)

- **feat: track certification state in git** — `SaveSnapshot()`/`LoadSnapshot()` on `record.Store`, `state.json` tracked in git for post-clone completeness. 5 new tests. ([#136f372](../../commit/136f372), Refs #12)

- **feat: certification run tracking** — `CertificationRun` domain type with `GenerateRunID()`, JSONL persistence in `.certification/runs.jsonl`, `RunID`/`PolicyVersion` populated on all records. 7 new tests. ([#eb4d464](../../commit/eb4d464), Refs #13)

- **feat: consolidate report generation** — `SaveReportArtifacts()` accepts `FullReport`, eliminates redundant `GenerateFullReport()` calls (was 2-3×, now 1×). ([#784c165](../../commit/784c165), Refs #14)

- **feat: wire/remove unused interfaces** — Deleted dead `evidence.Collector` interface, added `discovery.Scanners()` registry for polymorphic dispatch. 2 new tests. ([#95aa6ea](../../commit/95aa6ea), Refs #15)

### Chores

- **chore: gitignore per-unit markdown reports** — `.certification/reports/` (559 files, 2.2MB) removed from git tracking. Reports regenerate on demand via `certify report`. State tracked via `state.json`. ([#df55ebd](../../commit/df55ebd), Refs #17)

- **chore: unify language summary types** — Deleted `LanguageCard`, `LanguageBreakdown`, and `langRow`. `LanguageDetail` is now the single language summary type with `Passing` count. `Card.Languages` and `DetailedReport.ByLanguage` both use `LanguageDetail`. JSON backward-compatible (additive fields only). 4 new tests. ([#7791751](../../commit/7791751), Refs #18)

### Documentation

- **docs: update README, architecture, CLI docs for site report** — Added site report to quick start, flags, repository structure, and feature descriptions. Updated architecture docs with certification state model and report package description.

## Breaking Changes

- `Card.Languages` type changed from `[]LanguageCard` to `[]LanguageDetail`. JSON output gains new fields (`passing`, `grade_distribution`, `top_score`, `bottom_score`) — existing fields unchanged.
- `DetailedReport.ByLanguage` type changed from `map[string]LanguageBreakdown` to `map[string]LanguageDetail`. The `.Total` field is now `.Units`.
- `.certification/reports/` is now gitignored. Run `certify report` to regenerate. No data loss — all state is in `state.json`.

## Stats

- **10 architecture issues resolved** (#9–#18, all closed)
- **3 types deleted** (LanguageCard, LanguageBreakdown, langRow)
- **1 dead interface deleted** (evidence.Collector)
- **~40 new tests** across all changes
- **15 test packages**, all passing with zero regressions
- **559 units certified**, 100% pass rate, B+ overall

## Full Changelog

```
b7a7b83 chore: unify language summary types into LanguageDetail (Refs #18)
df55ebd chore: gitignore per-unit markdown reports (Refs #17)
95aa6ea feat: wire/remove unused interfaces (Refs #15)
784c165 feat: consolidate report generation (Refs #14)
eb4d464 feat: certification run tracking with JSONL persistence (Refs #13)
136f372 feat: track certification state in git via state.json (Refs #12)
df9fb30 feat: typed metrics map on Evidence (Refs #11)
b617e53 feat: extract Certifier service (Refs #10)
a23f933 feat: persist evidence in records (Refs #9)
b5805f9 feat: add static HTML site report output
```
