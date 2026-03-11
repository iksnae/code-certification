---
title: Quality Dimensions
description: The 9 quality dimensions Certify scores every code unit against.
---

Certify evaluates every code unit across 9 quality dimensions. Each dimension is scored independently from deterministic evidence, then combined using weighted averaging to produce a final score and grade.

## The Dimensions

### Correctness
**Does the code do what it claims?**

Evidence: lint errors (`go vet`, `golangci-lint`, ESLint, ruff, cargo clippy), test failures, `errors_ignored`, `panic_calls`, `os_exit_calls`, `defer_in_loop`, `empty_catch_blocks`.

### Maintainability
**How easy is it to modify safely?**

Evidence: cyclomatic complexity, function length, nesting depth, parameter count, `fan_out` (too many dependencies), `is_dead_code` (unused exports), `unused_params`, method count.

### Readability
**How clear and understandable is the code?**

Evidence: `cognitive_complexity` (Sonar-style), documentation presence, `max_nesting_depth`, `naked_returns`, `func_lines`, TODO/FIXME count.

### Testability
**How well is the code tested?**

Evidence: test coverage percentage, test file existence, `concrete_deps` (params using concrete types instead of interfaces — hard to mock), `os_exit_calls` (untestable), `init_func_count`.

### Security
**Are there security concerns?**

Evidence:
- `unsafe_import_count` — dangerous imports (os/exec, eval, subprocess, libc)
- `hardcoded_secrets` — string literals matching secret patterns (API keys, passwords)
- `global_mutable_count` — package-level mutable variables (race condition risk)

### Architectural Fitness
**Does the code fit the system's architecture?**

Evidence:
- `dep_depth` — transitive import chain depth
- `instability` — Robert C. Martin's instability metric (Ce / (Ca + Ce))
- `coupling_score` — fan-in × fan-out normalized
- `concrete_deps` — function params accepting concrete external types
- `interface_size` — methods in implemented interfaces (ISP violations)
- `context_not_first` — Go context.Context not as first parameter
- `method_count` — god object detection (>15 methods)

### Operational Quality
**How stable and observable is the code?**

Evidence: `errors_not_wrapped` / `type_aware_unwrapped` (errors returned without context), git churn (change frequency), contributor count, file age from `git log`.

### Performance Appropriateness
**Are there performance concerns?**

Evidence from AST analysis:
- `loop_nesting_depth` — deepest loop nesting
- `nested_loop_pairs` — nested loop pairs (O(n²) risk)
- `quadratic_patterns` — detected quadratic algorithm patterns
- `recursive_calls` — direct recursive function calls
- `defer_in_loop` — defer inside loops (resource leak + performance)

### Change Risk
**How risky is this code to modify?**

Evidence:
- `fan_in` — number of call sites invoking this function (high fan-in = many dependents affected by changes)
- Recent change frequency, author concentration from git history

## Evidence Sources

| Source | What it provides |
|--------|-----------------|
| **Lint** | `go vet`, `golangci-lint`, ESLint, ruff, cargo clippy — errors and warnings per unit |
| **Test** | `go test`, Jest/Vitest, pytest, cargo test — pass/fail status, per-unit coverage |
| **Git** | Churn rate, author count, last change date, file age |
| **Structural (AST)** | 27+ metrics from language-specific analysis — Go via `go/ast`, TS/Py/Rs via tree-sitter. Includes complexity, error handling, security patterns, documentation, nesting |
| **Deep Analysis** | Go: call graph via `go/packages` + SSA/VTA — `fan_in`, `fan_out`, `is_dead_code`, `dep_depth`, `instability`, `concrete_deps`, `unused_params`, `interface_size`, `type_aware_unwrapped`. TS/Py/Rs: via LSP servers (optional) |
| **Metrics** | Code lines, comment lines, cyclomatic complexity, TODO count |

### Analysis Tiers

| Tier | What | Go | TS/Py/Rs |
|------|------|-----|----------|
| **Tier 0** | Universal (git, line counts, TODOs) | ✅ Built-in | ✅ Built-in |
| **Tier 1** | Syntax AST (structural metrics, complexity, security) | ✅ `go/ast` | ✅ tree-sitter |
| **Tier 2** | Type-aware (call graph, dead code, dep graph) | ✅ `go/types` + `go/packages` | Optional LSP server |

Run `certify doctor` to see which tiers are available per language.

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
