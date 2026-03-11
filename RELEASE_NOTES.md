# Release v0.9.0

**Date:** 2026-03-11

## Highlights

**Complete snapshot↔prompt data alignment** eliminates all known LLM hallucination vectors in the architect review pipeline. Every metric cited by any phase prompt is now present — with exact AST-computed values — in the snapshot sent to the model. The project self-certifies at **A- (91.8%) with 4 observations** across 816 units.

## What's Changed

### Bug Fixes

- **fix(architect): complete snapshot↔prompt data alignment** — Extends PR #20's structural metrics fix to close all remaining gaps where architect prompts referenced data absent from the snapshot.

  **Structural metrics expanded (7 → 16):** Added `naked_returns`, `recursive_calls`, `max_nesting_depth`, `nested_loop_pairs`, `quadratic_patterns`, `total_func_lines`, `total_params`, `total_returns`, `total_methods` to `StructuralAggregates`.

  **Coverage aggregates added:** New `CoverageAggregates` struct provides `units_with_coverage`, `units_without_coverage`, `avg_coverage`, `min_coverage`, `max_coverage` — computed from `EvidenceKindTest` evidence. Phase 3 previously fabricated coverage percentages; now references exact data.

  **Code metrics aggregates added:** New `CodeMetricsAggregates` struct provides `total_code_lines`, `total_comment_lines`, `total_complexity`, `max_complexity`, `avg_complexity`, `total_todos` — computed from `EvidenceKindMetrics` evidence. Phase 2 previously invented complexity numbers.

  **`context_not_first` counting fixed:** Changed from boolean per-file counting (`> 0 → count++`) to per-unit summing (`+= int(...)`) to match the prompt's "functions with context.Context not as first param" language.

  **Anti-hallucination grounding in all 6 phases:** Previously only Phases 4–5 had grounding language. Now all phases include explicit instructions to cite only data present in the snapshot and never fabricate values.

  **Pipeline version tagging:** `ArchSnapshot.SchemaVersion` field (v2) is set by `BuildSnapshot()`, rendered in `FormatForLLM()` header, and included in the architect report appendix. Consumers can distinguish pre- and post-fix reports.

### Evidence of Fix

Tested with Qwen3-Coder-30b against 816 units:
- Phase 3 correctly reports "128 units lack coverage data" and uses exact 74.4% average (previously fabricated)
- Phase 4 reports `panic_calls: 0`, `os_exit_calls: 1`, `errors_ignored: 5` — all exact (previously Qwen claimed 23 panic calls)
- Phase 5 recommendations cite exact snapshot values in all deltas
- Phase 6 synthesis traces claims to specific phases

### Tests Added

9 new tests covering all changes:
- `TestBuildSnapshot_StructuralMetrics_Extended` — 9 new metric aggregations
- `TestBuildSnapshot_ContextNotFirst_SumsNotBool` — per-unit summing
- `TestBuildSnapshot_CoverageAggregates` — coverage min/max/avg
- `TestBuildSnapshot_CodeMetricsAggregates` — complexity/lines/todos
- `TestBuildSnapshot_SchemaVersion` — version field
- `TestFormatForLLM_CoverageMetrics` — rendered coverage section
- `TestFormatForLLM_CodeMetrics` — rendered code metrics section
- `TestFormatForLLM_SchemaVersion` — schema version in header
- `TestArchitectPrompts_AllContainGrounding` — all 6 phases grounded

## Files Changed

| File | Change |
|------|--------|
| `internal/agent/architect_snapshot.go` | Expanded types, aggregation logic, schema version |
| `internal/agent/architect.go` | New format functions, schema version in header |
| `internal/agent/architect_prompts.go` | Grounding language in all 6 phases |
| `internal/agent/architect_snapshot_test.go` | 5 new aggregation tests |
| `internal/agent/architect_test.go` | 4 new format/prompt tests |
| `internal/report/architect_report.go` | Schema version in appendix |
| `specs/snapshot-data-alignment.md` | Full implementation plan |

## Upgrade Notes

- **Backward compatible:** All changes are additive. Old snapshot JSON with `schema_version: 0` (Go zero value) correctly indicates pre-alignment data.
- **VS Code extension:** No changes needed — extension reads `ARCHITECT_REVIEW.md` as text, not snapshot types.
- **Re-run recommended:** Run `certify architect` after upgrading to generate reports with complete data grounding.
