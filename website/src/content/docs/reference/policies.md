---
title: Policy Packs
description: Define and customize certification policies.
---

Policy packs are versioned YAML files that define the rules your code is evaluated against. They live in `.certification/policies/`.

## Structure

```yaml
name: go-standard
version: "1.1.0"
language: go         # optional — targets specific language

rules:
  - id: go-vet-clean
    dimension: correctness
    description: "go vet must report zero issues"
    severity: error
    metric: lint_errors
    threshold: 0

  - id: low-complexity
    dimension: maintainability
    description: "Cyclomatic complexity under 15"
    severity: warning
    metric: cyclomatic_complexity
    threshold: 15

  - id: no-todos
    dimension: readability
    description: "No TODO/FIXME markers in production code"
    severity: warning
    metric: todo_count
    threshold: 0
    exclude_patterns:
      - "*_test.go"
```

## Fields

### Policy-Level

| Field | Required | Description |
|-------|----------|-------------|
| `name` | Yes | Unique policy pack name |
| `version` | Yes | Semantic version string |
| `language` | No | Target language (go, ts, py). Omit for universal. |
| `path_patterns` | No | Glob patterns — only apply this pack to matching paths |

### Rule-Level

| Field | Required | Description |
|-------|----------|-------------|
| `id` | Yes | Unique rule identifier |
| `dimension` | Yes | Quality dimension this rule evaluates |
| `description` | Yes | Human-readable explanation |
| `severity` | Yes | `error` (must pass) or `warning` (observation) |
| `metric` | Yes | Evidence metric to check |
| `threshold` | Yes | Maximum acceptable value |
| `path_patterns` | No | Glob patterns — only apply this rule to matching paths |
| `exclude_patterns` | No | Glob patterns — skip this rule for matching paths |

## Path Scoping

Rules and entire packs can be scoped to specific paths using glob patterns.

### Rule-Level Scoping

Apply a rule only to specific paths:

```yaml
rules:
  - id: no-panic
    dimension: security
    description: "Library code should not call panic()"
    severity: error
    metric: panic_calls
    threshold: 0
    path_patterns:
      - "internal/**"
    exclude_patterns:
      - "*_test.go"
```

This rule applies only to files under `internal/` and skips test files.

### Pack-Level Scoping

Scope an entire policy pack to a path:

```yaml
name: go-library
version: "1.0.0"
language: go
path_patterns:
  - "internal/**"

rules:
  - id: no-panic
    # ...
  - id: no-os-exit
    # ...
```

### Common Patterns

```yaml
# Only apply to non-test production code
exclude_patterns:
  - "*_test.go"
  - "testdata/**"

# Only apply to internal library code
path_patterns:
  - "internal/**"
  - "pkg/**"

# Only apply to CLI entry points
path_patterns:
  - "cmd/**"
```

## Available Metrics

| Metric | Source | Description |
|--------|--------|-------------|
| `lint_errors` | Linter output | Number of lint errors |
| `test_failures` | Test runner | Number of failing tests |
| `todo_count` | Code scan | TODO/FIXME comment count |
| `cyclomatic_complexity` | AST analysis | Cyclomatic complexity score |
| `line_count` | Code metrics | Total lines in unit |
| `function_count` | Code metrics | Number of functions |
| `panic_calls` | AST analysis | Number of panic() calls |
| `os_exit_calls` | AST analysis | Number of os.Exit() calls |
| `errors_ignored` | AST analysis | Error returns assigned to `_` |
| `global_mutable_count` | AST analysis | Package-level mutable variables |
| `defer_in_loop` | AST analysis | Defer statements inside loops |
| `naked_returns` | AST analysis | Bare return in named-return functions |

## Dimensions

Rules can target any of the 9 quality dimensions:

`correctness`, `maintainability`, `readability`, `testability`, `security`, `architectural_fitness`, `operational_quality`, `performance_appropriateness`, `change_risk`

## Severity

| Severity | Effect |
|----------|--------|
| `error` | Failing this rule blocks certification (unit is decertified) |
| `warning` | Creates an observation but unit can still be certified |

## Language Targeting

Policies with a `language` field only apply to units of that language:

```yaml
language: go    # Only applies to go:// units
```

Policies without a `language` field apply to all units.

## Default Policies

`certify init` generates auto-detected policy packs:

- **`go-standard.yml`** — Universal Go rules (vet clean, low complexity, no TODOs)
- **`go-library.yml`** — Library-specific rules scoped to `internal/**` (no panic, no os.Exit)

The split ensures `os.Exit` in CLI entry points and TODO strings in test fixtures don't generate false observations.

## Custom Policies

Add any `.yml` file to `.certification/policies/` — it's automatically loaded on the next certification run. Policy packs are matched to units by language and path patterns.
