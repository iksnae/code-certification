# Release v0.8.0

**Date:** 2026-03-10

## Highlights

**Rule-level policy scoping** and **code quality remediation** eliminate all observations. The project now self-certifies at **A- (91.8%) with 0 observations** across 748 units.

## What's Changed

### New Features

- **feat: rule-level path scoping for policy rules** — `PolicyRule` now supports `path_patterns` (include) and `exclude_patterns` (exclude) fields. Rules can be scoped to specific paths or excluded from test files. `Evaluate()` filters rules by unit path before evaluation.

- **feat: split policy packs** — `go-standard.yml` split into two packs:
  - `go-standard.yml` (v1.1.0): 13 universal rules. `no-todos` excludes `*_test.go`.
  - `go-library.yml` (v1.0.0): `no-panic` and `no-os-exit` scoped to `internal/**` only.
  
  This means `os.Exit` in CLI entry points and TODO strings in test fixtures no longer generate false observations.

### Bug Fixes

- **fix: TODO false positives** — `containsTodo` now enforces word boundaries (non-letter before/after) and skips TODO/FIXME inside quoted strings in comments. Eliminates false positives from identifiers like `extractTodoCount` and comments like `// Parse "N TODOs"`.

- **fix: convert mutable global vars to functions** — 8 `var` declarations in `providers.go` and `autodetect.go` converted to functions returning fresh slices/maps. Eliminates `global_mutable_count` for the agent package entirely.

- **fix: BasicLit const-like detection** — `var Version = "dev"` (ldflags pattern) now recognized as const-like. Added `*ast.BasicLit` case to `isConstLikeExpr`.

- **fix: extract countGlobalMutables helper** — Reduces `AnalyzeGoFile` nesting depth from 5 to 3.

- **fix: wire PathPatterns/ExcludePatterns through config loader** — `rawPolicyRule` was missing the new fields, causing YAML values to be silently dropped during parsing.

## Results

| Metric | v0.7.0 | v0.8.0 | Change |
|--------|--------|--------|--------|
| Score | 91.7% | **91.8%** | +0.1% |
| Observations | 12 | **0** | -12 |
| A- units | 669 | **676** | +7 |
| Policy packs | 2 | **3** | +1 |
| go-standard rules | 15 | **13** | -2 (moved to go-library) |

## Tests Added

- `TestEvaluate_RuleExcludePatterns` — rule with exclude_patterns skips test files
- `TestEvaluate_RulePathPatterns` — rule with path_patterns only fires for matching paths
- `TestEvaluate_NoPatterns_AppliesToAll` — existing behavior unchanged
- `TestLoadPolicyPack_GoLibrary` — validates go-library.yml structure
- `TestLoadPolicyPack_GoStandard_ExcludePatterns` — validates YAML exclude_patterns loading
- `TestCodeMetrics_TodoCount_QuotedNotFlagged` — TODO inside quoted strings
- `TestCodeMetrics_TodoCount_IdentifierNotFlagged` — TODO inside identifiers
- 2 new const-like var tests (string literal, int literal)
- All 16 packages pass with zero regressions

## Full Changelog

```
f0dca8b6 docs: mark policy scoping plan complete
c3c6411d chore: re-certify — 0 observations with scoped policy rules
fb391766 fix: wire PathPatterns/ExcludePatterns through config loader
016e0e01 feat: split policy packs — go-library.yml for internal-only rules
1da4e293 feat: rule-level path scoping for policy rules
efb88d27 chore: re-certify after architect remediation — 3 observations, 91.8%
a75f0acb fix: TODO detection with word boundary + quoted string checks
6d22c97f fix: extract countGlobalMutables to reduce AnalyzeGoFile nesting depth
cc9b4fe1 fix: detect BasicLit vars as const-like (ldflags pattern)
01cac4b7 fix: convert mutable global vars to functions in agent package
30a6235a fix: TODO false positives — skip TODO/FIXME inside quoted strings in comments
```
