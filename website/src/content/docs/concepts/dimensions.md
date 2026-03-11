---
title: Quality Dimensions
description: The 9 quality dimensions Certify scores every code unit against.
---

Certify evaluates every code unit across 9 quality dimensions. Each dimension is scored independently from deterministic evidence, then combined using weighted averaging to produce a final score and grade.

## The Dimensions

### Correctness
**Does the code do what it claims?**

Evidence: lint errors (`go vet`, `golangci-lint`), test failures, type errors.

### Maintainability
**How easy is it to modify safely?**

Evidence: cyclomatic complexity, function length, nesting depth, parameter count.

### Readability
**How clear and understandable is the code?**

Evidence: line length, documentation presence, TODO/FIXME count, comment-to-code ratio, naming conventions.

### Testability
**How well is the code tested?**

Evidence: test coverage percentage, test file existence, test-to-code ratio.

### Security
**Are there security concerns?**

Evidence from AST analysis:
- `panic_calls` — panic() in production code
- `os_exit_calls` — os.Exit() calls
- `errors_ignored` — error returns assigned to `_`
- `global_mutable_count` — package-level mutable variables (race condition risk)
- `defer_in_loop` — defer inside for/range loops (resource leak risk)

### Architectural Fitness
**Does the code fit the system's architecture?**

Evidence: package structure, dependency direction, import patterns, `context_not_first` (functions with context.Context not as first parameter), `init_func_count` (hidden initialization).

### Operational Quality
**How stable is the code in practice?**

Evidence: git churn (change frequency), contributor count, file age from `git log`.

### Performance Appropriateness
**Are there performance concerns?**

Evidence from AST analysis:
- `max_nesting_depth` — deepest loop nesting
- `nested_loop_pairs` — nested loop pairs (O(n²) risk)
- `quadratic_patterns` — detected quadratic algorithm patterns
- `recursive_calls` — direct recursive function calls
- `naked_returns` — bare returns in named-return functions

### Change Risk
**How risky is this code to modify?**

Evidence: recent change frequency, author concentration, coupling metrics, number of dependents.

## Evidence Sources

| Source | What it provides |
|--------|-----------------|
| **Lint** | `go vet`, `golangci-lint` errors and warnings |
| **Test** | Pass/fail status, per-unit coverage percentage |
| **Git** | Churn rate, author count, last change date, file age |
| **Structural (AST)** | 16 metrics from Go AST analysis — panic calls, error handling, nesting depth, complexity patterns, function signatures |
| **Metrics** | Code lines, comment lines, cyclomatic complexity, TODO count |

## Scoring

Each dimension produces a score from 0.0 to 1.0. These are combined using configurable weights:

```
Final Score = Σ (dimension_score × dimension_weight) / Σ weights
```

Default weights give equal importance to all dimensions. Customize in your policy configuration.

## Grades

| Grade | Score | Meaning |
|-------|-------|---------|
| **A** | ≥ 93% | Excellent across all dimensions |
| **A-** | ≥ 90% | Very strong, minor areas for improvement |
| **B+** | ≥ 87% | Strong, above average quality |
| **B** | ≥ 80% | Good, meets expectations |
| **C** | ≥ 70% | Acceptable but has notable issues |
| **D** | ≥ 60% | Below expectations, needs improvement |
| **F** | < 60% | Fails to meet quality standards |
