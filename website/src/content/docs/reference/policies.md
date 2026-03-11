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

### Universal (all languages)

| Metric | Source | Description |
|--------|--------|-------------|
| `lint_errors` | Linter output | Number of lint errors |
| `test_failures` | Test runner | Number of failing tests |
| `todo_count` | Code scan | TODO/FIXME comment count |
| `code_lines` | Code metrics | Lines of code |
| `complexity` | Code metrics | Cyclomatic complexity |

### Structural (AST analysis — all languages with analyzer)

| Metric | Source | Description |
|--------|--------|-------------|
| `param_count` | AST | Number of function parameters |
| `return_count` | AST | Number of return values |
| `func_lines` | AST | Lines in function body |
| `max_nesting_depth` | AST | Deepest nesting level |
| `has_doc_comment` | AST | Has documentation comment |
| `cognitive_complexity` | AST | Sonar-style cognitive complexity |
| `errors_ignored` | AST | Error returns assigned to `_` |
| `errors_not_wrapped` | AST | Errors returned without wrapping |
| `naked_returns` | AST | Bare return in named-return functions |
| `panic_calls` | AST | panic(), throw, unwrap() calls |
| `empty_catch_blocks` | AST | catch/except/recover with empty body |
| `os_exit_calls` | AST | os.Exit() / sys.exit() calls |
| `defer_in_loop` | AST | Defer/finally inside loops |
| `method_count` | AST | Methods on a type (type-level) |
| `context_not_first` | AST | context.Context not first param (Go) |
| `has_init_func` | AST | File contains init() |
| `global_mutable_count` | AST | Package-level mutable variables |
| `unsafe_import_count` | AST | Dangerous imports (os/exec, eval, subprocess) |
| `hardcoded_secrets` | AST | String literals matching secret patterns |
| `loop_nesting_depth` | AST | Max nested loop depth |
| `recursive_calls` | AST | Direct recursive calls |
| `nested_loop_pairs` | AST | Inner loops nested in outer loops |
| `quadratic_patterns` | AST | Known O(n²) anti-patterns |

### Deep Analysis (type-aware — Go built-in, TS/Py/Rs via LSP)

| Metric | Source | Description |
|--------|--------|-------------|
| `fan_in` | Call graph | Number of call sites invoking this function |
| `fan_out` | Call graph | Number of distinct functions called |
| `is_dead_code` | References | Exported symbol with zero external references |
| `dep_depth` | Import graph | Transitive local import depth |
| `instability` | Import graph | Ce/(Ca+Ce) — 0=stable, 1=unstable |
| `concrete_deps` | Type analysis | Params accepting concrete external struct types |
| `coupling_score` | Call graph | fan_in × fan_out normalized |
| `unused_params` | Type analysis | Parameters never referenced in function body |
| `interface_size` | Type analysis | Methods in interfaces this type implements |
| `type_aware_unwrapped` | Type analysis | Error returns without wrapping (type-verified) |

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

- **`go-standard.yml`** (v1.4.0) — 22 rules: lint clean, complexity, fan-in/fan-out limits, dead code, dep depth, cognitive complexity, unsafe imports, hardcoded secrets, error wrapping
- **`go-library.yml`** — Library-specific rules scoped to `internal/**` (no panic, no os.Exit)
- **`ts-standard.yml`** — TypeScript rules: complexity, nesting, empty catch, unsafe imports
- **`python-standard.yml`** — Python rules: complexity, bare except, subprocess usage
- **`rust-standard.yml`** — Rust rules: complexity, unwrap() usage, unsafe blocks

The split ensures `os.Exit` in CLI entry points and TODO strings in test fixtures don't generate false observations.

## Custom Policies

Add any `.yml` file to `.certification/policies/` — it's automatically loaded on the next certification run. Policy packs are matched to units by language and path patterns.
