# Policy Authoring Guide

## Overview

Policies are versioned YAML files stored in `.certification/policies/`. They define rules that code units must satisfy to earn certification.

## Policy Pack Structure

```yaml
name: my-policy
version: "1.0.0"
language: go              # optional: target specific language (go, ts, py, etc.)
path_pattern: "internal/" # optional: target specific paths

rules:
  - id: lint-clean
    dimension: correctness
    description: "Code must pass linter with zero errors"
    severity: error
    metric: lint_errors
    threshold: 0

  - id: max-complexity
    dimension: maintainability
    description: "Functions must have cyclomatic complexity ≤ 20"
    severity: warning
    metric: complexity
    threshold: 20
```

## Fields

### Pack-Level

| Field | Required | Description |
|-------|----------|-------------|
| `name` | Yes | Unique pack name |
| `version` | Yes | Semantic version string |
| `language` | No | Target language (empty = all) |
| `path_pattern` | No | Target path prefix (empty = all) |

### Rule-Level

| Field | Required | Description |
|-------|----------|-------------|
| `id` | Yes | Unique rule identifier |
| `dimension` | No | Quality dimension this rule evaluates |
| `description` | No | Human-readable explanation |
| `severity` | Yes | `info`, `warning`, `error`, or `critical` |
| `metric` | Yes | Evidence metric to check |
| `threshold` | Yes | Maximum allowed value (metric must be ≤ threshold) |

## Available Metrics

| Metric | Source | Description |
|--------|--------|-------------|
| `lint_errors` | go vet, golangci-lint | Number of lint errors |
| `test_failures` | go test | Number of failed tests |
| `todo_count` | Code metrics | Number of TODO/FIXME comments |
| `complexity` | AST analysis | Cyclomatic complexity |
| `test_coverage` | go test -cover | Test coverage percentage (0.0–1.0) |

## Severity Levels

| Level | Effect |
|-------|--------|
| `info` | Recorded but doesn't affect status |
| `warning` | Creates observation, may lower score |
| `error` | Blocks certification (decertified) |
| `critical` | Blocks certification, shorter trust window |

## Quality Dimensions

- `correctness` — Does the code do what it should?
- `maintainability` — Is the code easy to modify?
- `readability` — Is the code easy to understand?
- `testability` — Is the code well-tested?
- `security` — Is the code secure?
- `architectural_fitness` — Does the code fit the architecture?
- `operational_quality` — Is the code operationally sound?
- `performance` — Is performance appropriate?
- `change_risk` — How risky is changing this code?

## Examples

### Strict Security Policy

```yaml
name: security-strict
version: "1.0.0"
path_pattern: "internal/auth/"

rules:
  - id: no-lint-errors
    dimension: security
    severity: critical
    metric: lint_errors
    threshold: 0

  - id: low-complexity
    dimension: security
    severity: error
    metric: complexity
    threshold: 10
```

### Lenient Documentation Policy

```yaml
name: docs-advisory
version: "1.0.0"

rules:
  - id: no-todos
    dimension: readability
    severity: warning
    metric: todo_count
    threshold: 3
```

## Validation

Validate policies before committing:

```bash
certify certify --path . 2>&1 | head -5
# Any policy syntax errors will be reported
```

## Enabling/Disabling Packs

In `.certification/config.yml`:

```yaml
policies:
  disabled:
    - docs-advisory  # Skip this pack
```

Or enable only specific packs:

```yaml
policies:
  enabled:
    - security-strict
    - go-standard
```
