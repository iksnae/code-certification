---
title: Policy Packs
description: Define and customize certification policies.
---

Policy packs are versioned YAML files that define the rules your code is evaluated against. They live in `.certification/policies/`.

## Structure

```yaml
name: go-standard
version: "1.0.0"
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
```

## Fields

### Policy-Level

| Field | Required | Description |
|-------|----------|-------------|
| `name` | Yes | Unique policy pack name |
| `version` | Yes | Semantic version string |
| `language` | No | Target language (go, ts, py). Omit for universal. |

### Rule-Level

| Field | Required | Description |
|-------|----------|-------------|
| `id` | Yes | Unique rule identifier |
| `dimension` | Yes | Quality dimension this rule evaluates |
| `description` | Yes | Human-readable explanation |
| `severity` | Yes | `error` (must pass) or `warning` (observation) |
| `metric` | Yes | Evidence metric to check |
| `threshold` | Yes | Maximum acceptable value |

## Available Metrics

| Metric | Source | Description |
|--------|--------|-------------|
| `lint_errors` | Linter output | Number of lint errors |
| `test_failures` | Test runner | Number of failing tests |
| `todo_count` | Code scan | TODO/FIXME comment count |
| `cyclomatic_complexity` | AST analysis | Cyclomatic complexity score |
| `line_count` | Code metrics | Total lines in unit |
| `function_count` | Code metrics | Number of functions |

## Dimensions

Rules can target any of the 9 quality dimensions:

`correctness`, `maintainability`, `readability`, `testability`, `security`, `architectural_fitness`, `operational_quality`, `performance_appropriateness`, `change_risk`

## Severity

| Severity | Effect |
|----------|--------|
| `error` | Failing this rule blocks certification |
| `warning` | Creates an observation but unit can still be certified |

## Language Targeting

Policies with a `language` field only apply to units of that language:

```yaml
language: go    # Only applies to go:// units
```

Policies without a `language` field apply to all units.

## Default Policies

`certify init` generates:
- **`global.yml`** — Universal rules (lint clean, tests pass)
- **Language-specific** — Auto-detected (e.g., `go.yml`, `ts.yml`)

## Custom Policies

Add any `.yml` file to `.certification/policies/` — it's automatically loaded on the next certification run.
