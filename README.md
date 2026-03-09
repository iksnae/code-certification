# certify

[![CI](https://github.com/iksnae/code-certification/actions/workflows/ci.yml/badge.svg)](https://github.com/iksnae/code-certification/actions/workflows/ci.yml)
[![Certification](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/iksnae/code-certification/main/.certification/badge.json)](https://github.com/iksnae/code-certification/blob/main/.certification/REPORT_CARD.md)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Certify** evaluates every code unit in your repository against versioned policies, scores them across 9 quality dimensions, and generates a detailed **report card** with time-bound certification status.

[📋 See our own report card →](.certification/REPORT_CARD.md)

---

## Install

```bash
go install github.com/iksnae/code-certification/cmd/certify@latest
```

Or build from source:

```bash
git clone https://github.com/iksnae/code-certification.git
cd code-certification
go build -o certify ./cmd/certify/
```

**Requires:** Go 1.22+, Git

## Quick Start

```bash
cd your-repo

# 1. Bootstrap — creates config, policies, and CI workflows
certify init

# 2. Discover — finds every function, method, type, and file
certify scan

# 3. Certify — collects evidence, evaluates, scores, certifies
certify certify

# 4. Report — generates your report card
certify report --format full
```

That's it. Your report card is at `.certification/REPORT_CARD.md`.

## What You Get

### Report Card

A complete per-unit scoring of your entire codebase:

```
# 🔵 Code Certification — Full Report

## Summary
| Overall Grade | 🔵 B |
| Total Units   | 447  |
| Pass Rate     | 100% |

## Dimension Averages
| correctness              | 95.0% | ██████████████████░░ |
| maintainability          | 93.3% | ██████████████████░░ |
| readability              | 92.4% | ██████████████████░░ |
| testability              | 90.0% | █████████████████░░░ |
| security                 | 80.0% | ████████████████░░░░ |

## All Units (organized by directory)
| `Score`       | function | B | 86.7% | certified | 2026-04-23 |
| `CertifyUnit` | function | B | 85.6% | certified | 2026-04-23 |
... every unit in your repo
```

### Badge

Add this to your README — it updates automatically:

```bash
certify report --badge
```

Outputs:

```markdown
[![Certification](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/OWNER/REPO/main/.certification/badge.json)](https://github.com/OWNER/REPO/blob/main/.certification/REPORT_CARD.md)
```

Click the badge → full report card.

### CI Integration

`certify init` generates GitHub Actions workflows:

- **PR** — Certifies changed files, posts review summary
- **Nightly** — Sweeps for expired certifications
- **Weekly** — Full certification run + report card update

## 9 Quality Dimensions

Every code unit is scored across:

| Dimension | What it measures |
|-----------|-----------------|
| **Correctness** | Lint errors, vet issues, test failures |
| **Maintainability** | Cyclomatic complexity, function length |
| **Readability** | Line length, documentation, TODO count |
| **Testability** | Test coverage, test existence |
| **Security** | Security-sensitive patterns |
| **Architectural Fitness** | Package structure, dependency patterns |
| **Operational Quality** | Git churn, contributor count |
| **Performance** | Algorithmic complexity indicators |
| **Change Risk** | Recent changes, author concentration |

Dimensions are weighted and combined into a single score → grade (A through F).

## Certification Status

Certifications are **time-bound** — they expire:

| Status | Meaning |
|--------|---------|
| `certified` | Meets all policies, score above threshold |
| `certified_with_observations` | Passes but has warnings |
| `probationary` | Below threshold, grace period |
| `decertified` | Fails policy requirements |
| `expired` | Certification window elapsed, needs re-evaluation |
| `exempt` | Excluded by human override |

Default window: 90 days. Risk factors adjust the window (high churn → shorter).

## Commands

| Command | What it does |
|---------|-------------|
| `certify init` | Bootstrap `.certification/` with config + policies |
| `certify scan` | Discover all certifiable code units |
| `certify certify` | Evaluate, score, and certify units |
| `certify report` | Generate reports (text, card, full, json) |
| `certify expire` | Mark overdue certifications as expired |
| `certify version` | Show version |

### Useful Flags

```bash
certify certify --skip-agent         # no LLM review, deterministic only
certify certify --batch 20           # process 20 units at a time
certify certify --diff-base main     # only changed files (for PRs)
certify certify --target internal/   # scope to specific paths

certify report --format full         # complete report card (markdown)
certify report --format card         # terminal report card
certify report --format json         # machine-readable
certify report --badge               # print README badge snippet
certify report --output report.md    # write to file
```

## Configuration

`certify init` creates `.certification/config.yml`:

```yaml
mode: advisory        # advisory (report only) or enforcing (block on failure)

scope:
  include: []         # empty = everything
  exclude:
    - "vendor/**"
    - "node_modules/**"
    - "**/*_test.go"

expiry:
  default_window_days: 90
```

### Custom Policies

Add YAML policy files to `.certification/policies/`:

```yaml
name: my-team-standards
version: "1.0.0"
language: go

rules:
  - id: no-todos
    dimension: readability
    description: "No TODO comments in certified code"
    severity: warning
    metric: todo_count
    threshold: 0

  - id: low-complexity
    dimension: maintainability
    description: "Cyclomatic complexity under 15"
    severity: error
    metric: cyclomatic_complexity
    threshold: 15
```

### Agent-Assisted Review (Optional)

Add LLM-powered code review using open-weight models:

```yaml
agent:
  enabled: true
  provider:
    type: openrouter
    api_key_env: OPENROUTER_API_KEY
```

Uses Apache 2.0 licensed models (Qwen, Mistral). Agent review supplements deterministic evidence — the system works fully without it.

## Language Support

| Language | Adapter | Discovery |
|----------|---------|-----------|
| **Go** | Full | Functions, methods, types via `go/ast` |
| **TypeScript** | Basic | Classes, functions, exports via regex |
| **Everything else** | File-level | One unit per file |

## Files Generated

```
.certification/
├── config.yml          # your configuration
├── policies/           # policy packs (YAML)
├── records/            # per-unit certification records (JSON)
├── overrides/          # human governance overrides
├── index.json          # discovered unit index
├── REPORT_CARD.md      # ← the report card
└── badge.json          # shields.io badge endpoint
```

## Documentation

- [Architecture](docs/architecture.md) — system design + package diagram
- [Policy Authoring](docs/policy-authoring.md) — writing custom policies
- [Troubleshooting](docs/troubleshooting.md) — common issues
- [Contributor Guide](CLAUDE.md) — development setup

## License

[MIT](LICENSE)
