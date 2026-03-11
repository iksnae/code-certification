# Plan: Snapshot ↔ Prompt Data Alignment

## Task Description

Close all remaining gaps where architect review prompts reference data that is absent from the `ArchSnapshot` sent to the LLM. PR #20 fixed the structural metrics gap for Phase 4; this plan extends that fix to every phase and every evidence type, then hardens the pipeline against future drift.

## Objective

Guarantee that every metric, count, or data point cited by any architect phase prompt is present — with exact values — in the `FormatForLLM()` output. Eliminate all known hallucination vectors. Tag reports with the pipeline version so consumers can distinguish pre- and post-fix artifacts.

## Problem Statement

PR #20 (commit `cde8b1fd`) aggregated 7 structural metrics into the snapshot and grounded Phase 4 + 5 prompts. However:

| # | Gap | Affected Phases | Root Cause |
|---|-----|-----------------|------------|
| 1 | 9 structural metrics collected but NOT aggregated (`naked_returns`, `recursive_calls`, `loop_nesting_depth`, `nested_loop_pairs`, `quadratic_patterns`, `func_lines`, `param_count`, `return_count`, `method_count`) | 2, 5 | `StructuralAggregates` only includes 7 of 16 metrics |
| 2 | Per-unit coverage data (`unit_test_coverage`) never aggregated into snapshot | 3 | `EvidenceKindTest` with `coverage:unit` source is not read by `BuildSnapshot` |
| 3 | Code metrics (`complexity`, `code_lines`, `todo_count`, etc.) never aggregated | 2, 5 | `EvidenceKindMetrics` not read by `BuildSnapshot` |
| 4 | Phases 1–3 have no anti-hallucination grounding | 1, 2, 3 | Grounding language only added to Phases 4–5 |
| 5 | Phase 6 passes through unvalidated hallucinated values from earlier phases | 6 | No post-hoc validation of cited numbers |
| 6 | `context_not_first` counted as boolean (per-file) but prompt says "functions with" (per-function) | 4 | `if ev.Metrics["context_not_first"] > 0 { count++ }` instead of `+= int(...)` |
| 7 | Reports generated before this fix contain hallucinated data with no version marker | all | No `pipeline_version` or `schema_version` in snapshot or report output |

### Evidence of the Problem

Phase 3 prompt: *"Reference current observation counts and coverage-related metrics from the snapshot."*  
Snapshot contents: `TotalUnits`, `AvgScore`, `GradeDistribution`, `TopObservations` — **no coverage fields**.  
Result: LLM fabricates coverage percentages.

Phase 2 prompt: *"Complexity hotspots (reference the Hotspots table)"*  
Hotspots table columns: `Package | Units | Score | Risk Factor` — **no complexity, nesting, or line count data**.  
Result: LLM invents complexity numbers when asked about hotspots.

## Solution Approach

Six work items, ordered by hallucination severity:

1. **Expand structural aggregates** — add the 9 missing metrics to `StructuralAggregates` and the rendered table
2. **Add coverage aggregates** — new `CoverageAggregates` struct, populated from `EvidenceKindTest` evidence
3. **Add code metrics aggregates** — new `CodeMetricsAggregates` struct, populated from `EvidenceKindMetrics` evidence
4. **Ground all 6 phase prompts** — standardize anti-hallucination preamble across every phase
5. **Fix `context_not_first` counting** — sum values instead of treating as boolean
6. **Add pipeline version tagging** — version field in snapshot, rendered in `FormatForLLM()`, included in report output

## Relevant Files

### Core changes
- `internal/agent/architect_snapshot.go` — `StructuralAggregates`, `SnapshotMetrics`, `BuildSnapshot()`
- `internal/agent/architect.go` — `formatStructuralMetrics()`, `FormatForLLM()`, new format functions
- `internal/agent/architect_prompts.go` — all 6 phase system prompts

### Test files
- `internal/agent/architect_snapshot_test.go` — aggregation tests
- `internal/agent/architect_test.go` — `FormatForLLM` output tests

### Report output (version tagging)
- `internal/report/architect_report.go` — `writeArchAppendix()`, add pipeline version
- `internal/report/full.go` — `writeHeader()`, add schema version

### Reference (read-only, for field names)
- `internal/evidence/structural.go` — `ToEvidence()` metric keys (16 fields)
- `internal/evidence/metrics.go` — `CodeMetrics.ToEvidence()` metric keys (6 fields)
- `internal/engine/certifier.go` — `buildCoverageEvidence()` metric key (`unit_test_coverage`)
- `internal/domain/evidence.go` — `EvidenceKind` enum values

## Implementation Phases

### Phase 1: Expand Structural Aggregates + Fix `context_not_first`
Lowest risk, highest impact. Extends the pattern already established by PR #20.

### Phase 2: Add Coverage & Code Metrics Aggregates
New aggregate structs and format functions. Fixes the Phase 3 data gap.

### Phase 3: Ground All Prompts
Prompt-only changes. No data model changes.

### Phase 4: Pipeline Version Tagging
Cross-cutting metadata addition.

## Step by Step Tasks

### 1. Expand `StructuralAggregates` (TDD)

**Test first** — `architect_snapshot_test.go`:

```go
func TestBuildSnapshot_StructuralMetrics_Extended(t *testing.T) {
    records := []domain.CertificationRecord{
        makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
            Kind:   domain.EvidenceKindStructural,
            Source: "structural",
            Passed: true,
            Metrics: map[string]float64{
                "naked_returns":      2,
                "recursive_calls":    1,
                "loop_nesting_depth": 3,
                "nested_loop_pairs":  1,
                "quadratic_patterns": 0,
                "func_lines":         45,
                "param_count":        3,
                "return_count":       2,
                "method_count":       0,
            },
        }),
        makeRecordWithEvidence("go://pkg/b.go#Bar", 0.80, domain.Evidence{
            Kind:   domain.EvidenceKindStructural,
            Source: "structural",
            Passed: true,
            Metrics: map[string]float64{
                "naked_returns":      0,
                "recursive_calls":    2,
                "loop_nesting_depth": 4,
                "nested_loop_pairs":  0,
                "quadratic_patterns": 1,
                "func_lines":         120,
                "param_count":        5,
                "return_count":       1,
                "method_count":       8,
            },
        }),
    }

    snap := BuildSnapshot(records, "")
    s := snap.Metrics.Structural

    // Sums
    assertEqual(t, "naked_returns", s.NakedReturns, 2)
    assertEqual(t, "recursive_calls", s.RecursiveCalls, 3)
    assertEqual(t, "nested_loop_pairs", s.NestedLoopPairs, 1)
    assertEqual(t, "quadratic_patterns", s.QuadraticPatterns, 1)

    // Totals (summed across units)
    assertEqual(t, "total_func_lines", s.TotalFuncLines, 165)
    assertEqual(t, "total_params", s.TotalParams, 8)
    assertEqual(t, "total_returns", s.TotalReturns, 3)
    assertEqual(t, "total_methods", s.TotalMethods, 8)

    // Max (worst case across units)
    assertEqual(t, "max_nesting_depth", s.MaxNestingDepth, 4)
}
```

**Implementation** — `architect_snapshot.go`:

Add to `StructuralAggregates`:
```go
NakedReturns       int `json:"naked_returns"`
RecursiveCalls     int `json:"recursive_calls"`
MaxNestingDepth    int `json:"max_nesting_depth"`    // max across all units
NestedLoopPairs    int `json:"nested_loop_pairs"`
QuadraticPatterns  int `json:"quadratic_patterns"`
TotalFuncLines     int `json:"total_func_lines"`
TotalParams        int `json:"total_params"`
TotalReturns       int `json:"total_returns"`
TotalMethods       int `json:"total_methods"`
```

In `BuildSnapshot()` structural evidence loop, add:
```go
snap.Metrics.Structural.NakedReturns += int(ev.Metrics["naked_returns"])
snap.Metrics.Structural.RecursiveCalls += int(ev.Metrics["recursive_calls"])
snap.Metrics.Structural.NestedLoopPairs += int(ev.Metrics["nested_loop_pairs"])
snap.Metrics.Structural.QuadraticPatterns += int(ev.Metrics["quadratic_patterns"])
snap.Metrics.Structural.TotalFuncLines += int(ev.Metrics["func_lines"])
snap.Metrics.Structural.TotalParams += int(ev.Metrics["param_count"])
snap.Metrics.Structural.TotalReturns += int(ev.Metrics["return_count"])
snap.Metrics.Structural.TotalMethods += int(ev.Metrics["method_count"])
if nd := int(ev.Metrics["loop_nesting_depth"]); nd > snap.Metrics.Structural.MaxNestingDepth {
    snap.Metrics.Structural.MaxNestingDepth = nd
}
```

**Format** — `architect.go`, extend `formatStructuralMetrics()`:
```go
fmt.Fprintf(b, "| naked_returns | %d | Bare return statements in named-return functions |\n", s.NakedReturns)
fmt.Fprintf(b, "| recursive_calls | %d | Direct recursive function calls |\n", s.RecursiveCalls)
fmt.Fprintf(b, "| max_nesting_depth | %d | Deepest loop nesting across all units |\n", s.MaxNestingDepth)
fmt.Fprintf(b, "| nested_loop_pairs | %d | Nested loop pairs (O(n²) risk) |\n", s.NestedLoopPairs)
fmt.Fprintf(b, "| quadratic_patterns | %d | Detected quadratic algorithm patterns |\n", s.QuadraticPatterns)
fmt.Fprintf(b, "| total_func_lines | %d | Sum of function body lines |\n", s.TotalFuncLines)
fmt.Fprintf(b, "| total_params | %d | Sum of function parameter counts |\n", s.TotalParams)
fmt.Fprintf(b, "| total_returns | %d | Sum of function return value counts |\n", s.TotalReturns)
fmt.Fprintf(b, "| total_methods | %d | Sum of type method counts |\n", s.TotalMethods)
```

**Validation**: `go test ./internal/agent/... -count=1`

### 2. Fix `context_not_first` counting (TDD)

**Test first** — `architect_snapshot_test.go`:

```go
func TestBuildSnapshot_ContextNotFirst_SumsNotBool(t *testing.T) {
    records := []domain.CertificationRecord{
        makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
            Kind:    domain.EvidenceKindStructural,
            Source:  "structural",
            Passed:  true,
            Metrics: map[string]float64{"context_not_first": 1},
        }),
        makeRecordWithEvidence("go://pkg/b.go#Bar", 0.80, domain.Evidence{
            Kind:    domain.EvidenceKindStructural,
            Source:  "structural",
            Passed:  true,
            Metrics: map[string]float64{"context_not_first": 1},
        }),
    }
    snap := BuildSnapshot(records, "")
    // Should be 2, not 1 (boolean would give 1 after first > 0 check)
    if snap.Metrics.Structural.ContextNotFirst != 2 {
        t.Errorf("expected 2 context_not_first, got %d", snap.Metrics.Structural.ContextNotFirst)
    }
}
```

**Implementation** — `architect_snapshot.go`:

Change:
```go
if ev.Metrics["context_not_first"] > 0 {
    snap.Metrics.Structural.ContextNotFirst++
}
```
To:
```go
snap.Metrics.Structural.ContextNotFirst += int(ev.Metrics["context_not_first"])
```

Note: `context_not_first` is emitted as `contextVal` (0.0 or 1.0) in `structural.go` `ToEvidence()` because the underlying `StructuralMetrics.ContextNotFirst` is a `bool`. The per-unit value is always 0 or 1, but multiple units can each contribute 1, so summing is correct — "2 functions with context not first" rather than "1 file had the issue". The prompt says "Functions with context.Context not as first param" which matches per-unit summing.

**Validation**: `go test ./internal/agent/... -count=1`

### 3. Add `CoverageAggregates` to snapshot (TDD)

**Test first** — `architect_snapshot_test.go`:

```go
func TestBuildSnapshot_CoverageAggregates(t *testing.T) {
    records := []domain.CertificationRecord{
        makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
            Kind:    domain.EvidenceKindTest,
            Source:  "coverage:unit",
            Passed:  true,
            Metrics: map[string]float64{"unit_test_coverage": 0.85},
        }),
        makeRecordWithEvidence("go://pkg/b.go#Bar", 0.80, domain.Evidence{
            Kind:    domain.EvidenceKindTest,
            Source:  "coverage:unit",
            Passed:  true,
            Metrics: map[string]float64{"unit_test_coverage": 0.60},
        }),
        // Unit with no coverage evidence
        makeRecord("go://pkg/c.go#Baz", 0.90, nil),
    }

    snap := BuildSnapshot(records, "")
    c := snap.Metrics.Coverage

    if c.UnitsWithCoverage != 2 {
        t.Errorf("expected 2 units with coverage, got %d", c.UnitsWithCoverage)
    }
    if c.UnitsWithoutCoverage != 1 {
        t.Errorf("expected 1 unit without coverage, got %d", c.UnitsWithoutCoverage)
    }
    // Average of 0.85 and 0.60 = 0.725
    if diff := c.AvgCoverage - 0.725; diff > 0.001 || diff < -0.001 {
        t.Errorf("expected avg coverage ~0.725, got %.3f", c.AvgCoverage)
    }
    if c.MinCoverage != 0.60 {
        t.Errorf("expected min coverage 0.60, got %.2f", c.MinCoverage)
    }
}
```

**Implementation** — `architect_snapshot.go`:

New struct:
```go
// CoverageAggregates holds aggregated test coverage data across all units.
type CoverageAggregates struct {
    UnitsWithCoverage    int     `json:"units_with_coverage"`
    UnitsWithoutCoverage int     `json:"units_without_coverage"`
    AvgCoverage          float64 `json:"avg_coverage"`
    MinCoverage          float64 `json:"min_coverage"`
    MaxCoverage          float64 `json:"max_coverage"`
}
```

Add to `SnapshotMetrics`:
```go
Coverage  CoverageAggregates  `json:"coverage"`
```

In `BuildSnapshot()`, track coverage per record:
```go
var coverageValues []float64
// ... in the per-record loop:
for _, ev := range r.Evidence {
    if ev.Source == "coverage:unit" {
        if cov, ok := ev.Metrics["unit_test_coverage"]; ok {
            coverageValues = append(coverageValues, cov)
        }
    }
}
// ... after the loop:
snap.Metrics.Coverage.UnitsWithCoverage = len(coverageValues)
snap.Metrics.Coverage.UnitsWithoutCoverage = len(records) - len(coverageValues)
if len(coverageValues) > 0 {
    min, max, sum := coverageValues[0], coverageValues[0], 0.0
    for _, v := range coverageValues {
        sum += v
        if v < min { min = v }
        if v > max { max = v }
    }
    snap.Metrics.Coverage.AvgCoverage = sum / float64(len(coverageValues))
    snap.Metrics.Coverage.MinCoverage = min
    snap.Metrics.Coverage.MaxCoverage = max
}
```

**Format** — `architect.go`, new function:
```go
func formatCoverageMetrics(b *strings.Builder, snap *ArchSnapshot) {
    c := snap.Metrics.Coverage
    if c.UnitsWithCoverage == 0 && c.UnitsWithoutCoverage == 0 {
        return
    }
    b.WriteString("## Coverage Metrics (aggregated from all units)\n")
    fmt.Fprintf(b, "- Units with coverage data: %d\n", c.UnitsWithCoverage)
    fmt.Fprintf(b, "- Units without coverage data: %d\n", c.UnitsWithoutCoverage)
    if c.UnitsWithCoverage > 0 {
        fmt.Fprintf(b, "- Average coverage: %.1f%%\n", c.AvgCoverage*100)
        fmt.Fprintf(b, "- Min coverage: %.1f%%\n", c.MinCoverage*100)
        fmt.Fprintf(b, "- Max coverage: %.1f%%\n", c.MaxCoverage*100)
    }
    b.WriteString("\n")
}
```

Add `formatCoverageMetrics(&b, snap)` call in `FormatForLLM()` after `formatStructuralMetrics`.

**Test FormatForLLM output** — `architect_test.go`:
```go
func TestFormatForLLM_CoverageMetrics(t *testing.T) { ... }
```

**Validation**: `go test ./internal/agent/... -count=1`

### 4. Add `CodeMetricsAggregates` to snapshot (TDD)

**Test first** — `architect_snapshot_test.go`:

```go
func TestBuildSnapshot_CodeMetricsAggregates(t *testing.T) {
    records := []domain.CertificationRecord{
        makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
            Kind:    domain.EvidenceKindMetrics,
            Source:  "metrics",
            Passed:  true,
            Metrics: map[string]float64{
                "code_lines":  80,
                "complexity":  5,
                "todo_count":  1,
            },
        }),
        makeRecordWithEvidence("go://pkg/b.go#Bar", 0.80, domain.Evidence{
            Kind:    domain.EvidenceKindMetrics,
            Source:  "metrics",
            Passed:  true,
            Metrics: map[string]float64{
                "code_lines":  200,
                "complexity":  12,
                "todo_count":  0,
            },
        }),
    }

    snap := BuildSnapshot(records, "")
    cm := snap.Metrics.CodeMetrics

    assertEqual(t, "total_code_lines", cm.TotalCodeLines, 280)
    assertEqual(t, "total_complexity", cm.TotalComplexity, 17)
    assertEqual(t, "max_complexity", cm.MaxComplexity, 12)
    assertEqual(t, "total_todos", cm.TotalTodos, 1)
}
```

**Implementation** — `architect_snapshot.go`:

```go
// CodeMetricsAggregates holds aggregated code metrics across all units.
type CodeMetricsAggregates struct {
    TotalCodeLines   int     `json:"total_code_lines"`
    TotalComplexity  int     `json:"total_complexity"`
    MaxComplexity    int     `json:"max_complexity"`
    AvgComplexity    float64 `json:"avg_complexity"`
    TotalTodos       int     `json:"total_todos"`
    TotalCommentLines int    `json:"total_comment_lines"`
}
```

Add to `SnapshotMetrics`:
```go
CodeMetrics  CodeMetricsAggregates  `json:"code_metrics"`
```

In `BuildSnapshot()`, aggregate from `EvidenceKindMetrics`:
```go
for _, ev := range r.Evidence {
    if ev.Kind == domain.EvidenceKindMetrics {
        snap.Metrics.CodeMetrics.TotalCodeLines += int(ev.Metrics["code_lines"])
        c := int(ev.Metrics["complexity"])
        snap.Metrics.CodeMetrics.TotalComplexity += c
        if c > snap.Metrics.CodeMetrics.MaxComplexity {
            snap.Metrics.CodeMetrics.MaxComplexity = c
        }
        snap.Metrics.CodeMetrics.TotalTodos += int(ev.Metrics["todo_count"])
        snap.Metrics.CodeMetrics.TotalCommentLines += int(ev.Metrics["comment_lines"])
        metricsUnitCount++
    }
}
// After loop:
if metricsUnitCount > 0 {
    snap.Metrics.CodeMetrics.AvgComplexity = float64(snap.Metrics.CodeMetrics.TotalComplexity) / float64(metricsUnitCount)
}
```

**Format** — `architect.go`, new function:
```go
func formatCodeMetrics(b *strings.Builder, snap *ArchSnapshot) {
    cm := snap.Metrics.CodeMetrics
    if cm.TotalCodeLines == 0 {
        return
    }
    b.WriteString("## Code Metrics (aggregated from all units)\n")
    b.WriteString("| Metric | Value | Description |\n")
    b.WriteString("|--------|------:|-------------|\n")
    fmt.Fprintf(b, "| total_code_lines | %d | Lines of code (excluding blanks/comments) |\n", cm.TotalCodeLines)
    fmt.Fprintf(b, "| total_comment_lines | %d | Lines of comments |\n", cm.TotalCommentLines)
    fmt.Fprintf(b, "| total_complexity | %d | Sum of cyclomatic complexity |\n", cm.TotalComplexity)
    fmt.Fprintf(b, "| max_complexity | %d | Highest single-unit complexity |\n", cm.MaxComplexity)
    fmt.Fprintf(b, "| avg_complexity | %.1f | Average complexity per unit |\n", cm.AvgComplexity)
    fmt.Fprintf(b, "| total_todos | %d | TODO/FIXME markers in code |\n", cm.TotalTodos)
    b.WriteString("\n")
}
```

**Validation**: `go test ./internal/agent/... -count=1`

### 5. Ground all 6 phase prompts (no data model changes)

Add a shared grounding preamble to every phase prompt. This is prompt-only — no Go types change.

**Shared preamble** (new const in `architect_prompts.go`):
```go
const architectGroundingPreamble = `
GROUNDING RULES (apply to all phases):
- Only cite metrics, counts, and values that appear in the Architecture Snapshot below.
- If a metric is not present in the snapshot, do not reference it.
- Never fabricate, estimate, or round specific numeric values.
- If you cannot determine a value from the data, say "unknown" or "not available".
- When citing a number from a table, use the EXACT value shown.
`
```

**Per-phase changes:**

**Phase 1** — append to system prompt:
```
Reference only package names, unit counts, scores, and dependency edges that appear in the snapshot.
Do not cite metrics or values that are not explicitly present.
```

**Phase 2** — replace the citation example line:
```
For each finding, cite EXACT metric values from the snapshot tables.
Example: "engine/ has avg score 78.2% with 12 errors_ignored observations" — but only if those exact numbers appear in the Package Map and Top Observations tables.
Do not cite complexity values, coverage percentages, or function line counts unless they appear in the Code Metrics or Structural Metrics tables.
```

**Phase 3** — replace coverage reference:
```
Reference the "Coverage Metrics" section in the snapshot for exact coverage data.
If no coverage data is present, state "coverage data not available" — do not estimate percentages.
Reference observation counts from the Package Map table to identify packages with likely quality gaps.
```

**Phase 4** — already grounded by PR #20. No changes needed.

**Phase 5** — add after existing grounding rule:
```
- Use exact values from the "Code Metrics" table for complexity and line count claims
- Use exact values from the "Coverage Metrics" section for coverage claims
```

**Phase 6** — add grounding:
```
Rules:
- Every claim in the executive summary must be traceable to a specific phase output or snapshot metric
- Risk severity and likelihood must be grounded in actual metric values, not inferred
- Do not introduce new metrics or numbers not present in prior phase outputs
- If a prior phase output appears to contain estimates, flag them as "unverified" rather than repeating as fact
```

**Test** — `architect_test.go`:
```go
func TestArchitectPrompts_AllContainGrounding(t *testing.T) {
    prompts := ArchitectPhasePrompts()
    for i, p := range prompts {
        if !strings.Contains(p, "do not") && !strings.Contains(p, "Do not") &&
           !strings.Contains(p, "DO NOT") && !strings.Contains(p, "Never") {
            t.Errorf("Phase %d prompt missing anti-hallucination grounding", i+1)
        }
    }
}
```

**Validation**: `go test ./internal/agent/... -count=1`

### 6. Add pipeline version tagging

**Implementation** — `architect_snapshot.go`:

Add version constant and field:
```go
// SnapshotSchemaVersion tracks the data contract between BuildSnapshot and FormatForLLM.
// Increment when adding/removing/renaming fields in SnapshotMetrics or its sub-structs.
const SnapshotSchemaVersion = 2 // v1: PR #20 (structural only), v2: this PR (full alignment)
```

Add to `ArchSnapshot`:
```go
type ArchSnapshot struct {
    SchemaVersion  int              `json:"schema_version"`
    // ... existing fields
}
```

Set in `BuildSnapshot()`:
```go
snap.SchemaVersion = SnapshotSchemaVersion
```

**Format** — `architect.go`, add to `formatHeader()`:
```go
fmt.Fprintf(b, "**Snapshot Schema:** v%d\n", snap.SchemaVersion)
```

**Report output** — `internal/report/architect_report.go`, update `writeArchAppendix()`:
```go
if snap != nil {
    fmt.Fprintf(b, "- **%d** units across **%d** packages", snap.Metrics.TotalUnits, snap.Metrics.TotalPackages)
    if snap.Metrics.AvgScore > 0 {
        fmt.Fprintf(b, " · Score: %.1f%%", snap.Metrics.AvgScore*100)
    }
    fmt.Fprintf(b, " · Schema: v%d", snap.SchemaVersion)
    b.WriteString("\n")
}
```

This lets consumers of ARCHITECT_REVIEW.md see whether the report was generated with complete data grounding.

**Test** — `architect_snapshot_test.go`:
```go
func TestBuildSnapshot_SchemaVersion(t *testing.T) {
    snap := BuildSnapshot(nil, "")
    if snap.SchemaVersion != SnapshotSchemaVersion {
        t.Errorf("expected schema version %d, got %d", SnapshotSchemaVersion, snap.SchemaVersion)
    }
}
```

**Validation**: `go test ./internal/agent/... -count=1` and `go test ./internal/report/... -count=1`

## Testing Strategy

TDD for every change — write failing test, implement, verify pass.

| Step | Test File | Test Name | What It Verifies |
|------|-----------|-----------|-----------------|
| 1 | `architect_snapshot_test.go` | `TestBuildSnapshot_StructuralMetrics_Extended` | 9 new metrics aggregate correctly |
| 2 | `architect_snapshot_test.go` | `TestBuildSnapshot_ContextNotFirst_SumsNotBool` | Per-function counting, not per-file |
| 3 | `architect_snapshot_test.go` | `TestBuildSnapshot_CoverageAggregates` | Coverage from `EvidenceKindTest` with min/max/avg |
| 3 | `architect_test.go` | `TestFormatForLLM_CoverageMetrics` | Coverage section appears in LLM output |
| 4 | `architect_snapshot_test.go` | `TestBuildSnapshot_CodeMetricsAggregates` | Code metrics from `EvidenceKindMetrics` |
| 4 | `architect_test.go` | `TestFormatForLLM_CodeMetrics` | Code metrics section appears in LLM output |
| 5 | `architect_test.go` | `TestArchitectPrompts_AllContainGrounding` | Every phase has anti-hallucination language |
| 6 | `architect_snapshot_test.go` | `TestBuildSnapshot_SchemaVersion` | Version field set |

Run full suite after each step:
```bash
go test ./internal/agent/... -count=1
```

Run full project suite after all steps:
```bash
go test ./... -count=1
```

## Acceptance Criteria

- [ ] All 16 structural metrics from `ToEvidence()` are aggregated in `StructuralAggregates` (up from 7)
- [ ] Coverage data aggregated into `CoverageAggregates` with `units_with_coverage`, `avg_coverage`, `min_coverage`, `max_coverage`
- [ ] Code metrics aggregated into `CodeMetricsAggregates` with `total_complexity`, `max_complexity`, `avg_complexity`, `total_code_lines`, `total_todos`
- [ ] `context_not_first` sums per-unit values (not boolean count)
- [ ] All 6 phase prompts contain anti-hallucination grounding language
- [ ] `ArchSnapshot.SchemaVersion` is set and rendered in `FormatForLLM()` output
- [ ] Architect report appendix includes schema version
- [ ] All new tests pass: `go test ./internal/agent/... -count=1`
- [ ] Full suite passes: `go test ./... -count=1`
- [ ] `go vet ./...` clean
- [ ] `gofmt -l .` clean (excluding testdata)

## Validation Commands

```bash
# All tests pass
go test ./... -count=1

# Build succeeds
go build -o build/bin/certify ./cmd/certify/

# Format + vet clean
go vet ./...
gofmt -l . | grep -v testdata | head -5  # expect no output

# Verify FormatForLLM contains all new sections
go test -run TestFormatForLLM -v ./internal/agent/...

# Verify all prompts are grounded
go test -run TestArchitectPrompts_AllContainGrounding -v ./internal/agent/...

# Verify schema version
go test -run TestBuildSnapshot_SchemaVersion -v ./internal/agent/...

# End-to-end: generate a fresh architect review and inspect
go install ./cmd/certify/
certify certify --skip-agent --batch 0
certify architect --phases 4 2>/dev/null | grep -c "Structural Metrics"  # should be >= 1
certify architect --phases 3 2>/dev/null | grep -c "Coverage Metrics"    # should be >= 1
```

## Risk Assessment

| Risk | Likelihood | Impact | Mitigation |
|------|-----------|--------|------------|
| Expanded structural table exceeds LLM context window | Low | Medium | `FormatForLLM` already has `maxTokensHint`; new sections add ~500 chars total |
| Coverage evidence missing for non-Go projects | Medium | Low | `formatCoverageMetrics` already handles zero case with early return |
| Anti-hallucination prompts reduce LLM creativity in recommendations | Low | Low | Only constrains fabrication of numbers, not analysis or recommendations |
| Schema version breaks JSON consumers | Low | Low | Additive field — existing consumers ignore unknown fields |
| `context_not_first` counting change breaks existing tests | Low | Low | PR #20 test uses value 0 and 1 — sum and boolean give same result for those inputs |

## Notes

- **Backward compatibility**: All changes are additive to `SnapshotMetrics`. The JSON `schema_version` field defaults to 0 for old snapshots (Go zero value), which correctly indicates "pre-alignment".
- **Token budget**: The three new format sections add approximately 400–600 characters to `FormatForLLM()` output. At ~4 chars/token, that's ~100–150 tokens — well within the 4000-token hint.
- **Why not per-package aggregates?** Phase 2 asks about per-package patterns, but the Package Map table already provides per-package `AvgScore`, `Units`, `Observations`, and `TopIssues`. Adding per-package complexity or coverage would significantly increase output size. If needed, that's a follow-up — the global aggregates close the immediate hallucination gap.
- **Nightly sweep re-run**: After merging, the next nightly certification sweep will generate reports with schema v2 data. No manual invalidation of old reports is needed — the version field lets consumers distinguish.

## Commit Plan

Single PR with atomic commit:
```
fix(architect): complete snapshot↔prompt data alignment

Close all remaining gaps where architect prompts reference data not
present in the snapshot. Extends PR #20's structural metrics fix to
coverage, code metrics, and complexity data.

Changes:
- Expand StructuralAggregates: 7 → 16 metrics (add naked_returns,
  recursive_calls, loop_nesting_depth, nested_loop_pairs,
  quadratic_patterns, func_lines, param_count, return_count,
  method_count)
- Fix context_not_first: sum per-unit values instead of boolean count
- Add CoverageAggregates: units_with/without_coverage, avg/min/max
- Add CodeMetricsAggregates: total_complexity, max_complexity,
  avg_complexity, total_code_lines, total_todos, total_comment_lines
- Add anti-hallucination grounding to all 6 phase prompts
- Add SnapshotSchemaVersion (v2) to snapshot and report output
- Add tests for all new aggregation, formatting, and prompt grounding

Fixes #21
```
