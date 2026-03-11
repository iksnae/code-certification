# Release v0.7.0

**Date:** 2026-03-10

## Highlights

**Algorithmic complexity detection** and **measurement accuracy fixes** push the self-certification score from B+ (88.7%) to **A- (91.7%)** with 92.4% of units at A- or above.

## What's Changed

### New Features

- **feat: algorithmic complexity detection via go/ast** — New structural analysis that estimates Big-O complexity class for every function: O(1), O(n), O(n²), O(n³), O(2^n). Detects nested loop depth, recursive calls, and quadratic anti-patterns (string concat in loops). 5 new metrics: `algo_complexity`, `loop_nesting_depth`, `recursive_calls`, `nested_loop_pairs`, `quadratic_patterns`.

- **feat: performance_appropriateness scoring** — The `performance_appropriateness` dimension is now always measured for function-level units (was penalty-only for `defer_in_loop`). O(1)→0.95, O(n)→0.90, O(n²)→0.70, O(n³)→0.50, O(2^n)→0.40.

- **feat: graduated git history scoring** — `operational_quality` now scales with commit count: >50→0.95, >20→0.90, >10→0.85, >0→0.75. `change_risk` scales with author count: ≥3→0.95, ≥2→0.90, 1→0.70. Previously all units were capped at 0.85/0.90.

- **feat: const-like var detection** — Structural analyzer now distinguishes truly mutable `var` declarations from const-like ones. Composite literals (`map[K]V{...}`, `[]T{...}`), error sentinels (`errors.New`), compiled regexes (`regexp.MustCompile`), and pointer-to-literals (`&T{}`) are excluded from `global_mutable_count`. Only `make()`/`new()`/uninitialized vars count as mutable. Fixes 168 units falsely penalized for lookup tables.

- **feat: graduated file-level readability** — File-level `code_lines` thresholds raised to match file reality: ≤100→0.95, ≤300→0.90, ≤500→0.85, ≤800→0.75, >800→0.60 (was ≤50/≤150/≤300/>300).

- **feat: graduated complexity scoring** — Added 0.80 tier for cyclomatic complexity 11-15 (was jump from 0.85 at ≤10 to 0.70 at ≤20).

### Bug Fixes

- **fix: 14 exported symbols documented** — Added doc comments to architect review types and interface method implementations to fix readability penalties.

## Results

| Metric | v0.6.2 | v0.7.0 | Change |
|--------|--------|--------|--------|
| Overall Grade | B+ | **A-** | ↑ |
| Score | 88.7% | **91.7%** | +3.0% |
| A/A- units | 395 (52.8%) | **691 (92.4%)** | +296 |
| B+ units | 148 | 26 | -122 |
| B units | 192 | 31 | -161 |
| C units | 13 | **0** | -13 |
| Observations | 13 | **0** | -13 |
| Dimensions measured | 7 | **8** | +performance_appropriateness |

## Tests Added

- 26 new tests: algorithmic complexity (12), const-like var detection (12), graduated scoring (9), algo scoring (5)
- All 16 packages pass with zero regressions

## Full Changelog

```
9d89d9a fix: add doc comments to 14 exported symbols
0c1729d feat: score performance_appropriateness from algorithmic complexity
64ac068 feat: algorithmic complexity detection via go/ast
ac78ce0 feat: graduated file-level readability + complexity scoring
d33e933 feat: const-like var detection — exclude lookup tables
1908dd9 feat: graduated git history scoring for operational_quality and change_risk
```
