# Plan: Rule-Level Path Scoping + Policy Pack Split

## Problem

Our policies are too blunt:
- `no-os-exit` says "library code" but applies to `cmd/certify/main.go` (CLI entry point)
- `no-todos` applies to test files that test TODO detection
- The architect review treats every observation as a real concern because policies can't express scope

## Solution: A + B Combined

### A: Rule-Level Path Scoping

Add `path_patterns` and `exclude_patterns` to `PolicyRule`. When present, the evaluator skips rules that don't match the unit's path.

```yaml
- id: no-todos
  dimension: readability
  description: "No TODO/FIXME comments in certified code"
  severity: warning
  metric: todo_count
  threshold: 0
  exclude_patterns: ["*_test.go"]   # ← NEW
```

### B: Split Policy Packs

Split `go-standard.yml` into:
- **`go-standard.yml`** — universal rules (lint, tests, complexity, params, nesting, func-lines, defer-in-loop, context-first, init-func, global-mutable, max-methods)
- **`go-library.yml`** — library-only rules (`internal/**`): no-os-exit, no-panic
- Keep `no-todos` in go-standard with `exclude_patterns: ["*_test.go"]`

## Implementation

### Step 1: Add fields to PolicyRule (domain)

```go
type PolicyRule struct {
    ID              string    `json:"id" yaml:"id"`
    Dimension       Dimension `json:"dimension" yaml:"dimension"`
    Description     string    `json:"description" yaml:"description"`
    Severity        Severity  `json:"severity" yaml:"severity"`
    Threshold       float64   `json:"threshold,omitempty" yaml:"threshold,omitempty"`
    Metric          string    `json:"metric,omitempty" yaml:"metric,omitempty"`
    PathPatterns    []string  `json:"path_patterns,omitempty" yaml:"path_patterns,omitempty"`       // NEW
    ExcludePatterns []string  `json:"exclude_patterns,omitempty" yaml:"exclude_patterns,omitempty"` // NEW
}
```

### Step 2: Add unit path to Evaluate signature

Change `Evaluate(rules, ev)` → `Evaluate(rules, ev, unitPath)`.

Filter rules by path before evaluating:
- If `rule.PathPatterns` is set, unit path must match at least one
- If `rule.ExcludePatterns` is set, unit path must NOT match any

### Step 3: Split policy packs

`go-standard.yml`:
- lint-clean, test-pass, no-todos (exclude *_test.go), max-complexity, max-params, max-nesting, no-ignored-errors, max-func-lines, no-defer-in-loop, context-first, no-init-func, limit-global-mutable, max-methods

`go-library.yml` (path_patterns: ["internal/**"]):
- no-panic, no-os-exit

### Step 4: Update testdata policies

Mirror the split in `testdata/policies/`.

### Step 5: Update callers

`CertifyUnit` in pipeline.go passes `unit.ID.Path()` to `Evaluate`.

## Testing

1. `TestEvaluate_RulePathPatterns` — rule with path_patterns only fires for matching paths
2. `TestEvaluate_RuleExcludePatterns` — rule with exclude_patterns skips matching paths  
3. `TestEvaluate_NoPatterns_AppliesToAll` — existing behavior unchanged
4. `TestPolicyPackSplit` — go-library.yml rules only match internal/** units
5. Existing evaluator tests still pass (add unitPath param)

## Acceptance Criteria

- [x] `PolicyRule` has `PathPatterns` and `ExcludePatterns` fields
- [x] `Evaluate` accepts unit path and filters rules accordingly
- [x] `go-standard.yml` excludes `*_test.go` from `no-todos`
- [x] `go-library.yml` contains `no-os-exit` and `no-panic` scoped to `internal/**`
- [x] `main.go#main` has 0 observations after re-certification
- [x] Test files with TODO fixtures have 0 observations
- [x] All existing tests pass
- [x] Re-certification shows 0 observations

## Completion Report

**Date:** March 10, 2026

### What was implemented

Rule-level path scoping for policy rules (Option A) plus policy pack split (Option B):

1. **`PolicyRule` struct** — Added `PathPatterns` and `ExcludePatterns` fields with JSON/YAML tags.

2. **`Evaluate()` signature** — Changed from `Evaluate(rules, ev)` to `Evaluate(rules, ev, unitPath)`. New `ruleAppliesToPath()` function filters rules by matching unit path against basename and full path using `filepath.Match`.

3. **Config loader** — Added `PathPatterns` and `ExcludePatterns` to `rawPolicyRule` and wired them through `parsePolicyPack`. (This was the bug that initially prevented YAML fields from loading.)

4. **Policy pack split**:
   - `go-standard.yml` (v1.1.0): 13 universal rules. `no-todos` has `exclude_patterns: ["*_test.go"]`. Removed `no-panic` and `no-os-exit`.
   - `go-library.yml` (v1.0.0): 2 library-only rules (`no-panic`, `no-os-exit`) scoped to `internal/**` via pack-level `path_patterns`.

5. **testdata/policies/** mirrored the same split.

### Tests added

- `TestEvaluate_RuleExcludePatterns` — rule with exclude_patterns skips test files
- `TestEvaluate_RulePathPatterns` — rule with path_patterns only fires for matching paths
- `TestEvaluate_NoPatterns_AppliesToAll` — existing behavior unchanged
- `TestLoadPolicyPack_GoLibrary` — validates go-library.yml structure
- `TestLoadPolicyPack_GoStandard_ExcludePatterns` — validates YAML exclude_patterns loading

### Issues encountered

- **Config loader silently dropped new fields**: `rawPolicyRule` didn't have `PathPatterns`/`ExcludePatterns`, so YAML parsing worked but the values were lost during conversion to `domain.PolicyRule`. Caught by the ExcludePatterns test.
- **All existing `Evaluate()` callers needed updating**: 12 call sites in evaluator_test.go + 1 in pipeline.go needed the new `unitPath` parameter.

### Results

| Metric | Before | After |
|--------|--------|-------|
| Observations | 3 | **0** |
| Policy packs | 2 (go-standard, global) | **3** (go-standard, go-library, global) |
| go-standard rules | 15 | **13** |
| go-library rules | 0 | **2** (scoped to internal/**) |
