<p align="center">
  <img src="assets/logo-256.png" alt="Certify" width="128" />
</p>

<h1 align="center">Certify</h1>

<p align="center">
  <a href="https://github.com/iksnae/code-certification/actions/workflows/ci.yml"><img src="https://github.com/iksnae/code-certification/actions/workflows/ci.yml/badge.svg" alt="CI" /></a>
  <a href="https://github.com/iksnae/code-certification/blob/main/.certification/REPORT_CARD.md"><img src="https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/iksnae/code-certification/main/.certification/badge.json" alt="Certification" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License: MIT" /></a>
</p>

**Code trust, with an expiration date.**

Certify continuously evaluates every code unit in your repository, scores it against versioned policies, and assigns time-bound certification you can actually trust.

CI tells you whether code passes right now. **Certify tells you whether code should still be trusted.**

[📋 View our report card →](.certification/REPORT_CARD.md) · [📖 Documentation →](https://iksnae.github.io/code-certification/)

---

## Why Certify

Code that once passed review doesn't stay trustworthy forever. Standards evolve, dependencies change, systems grow more complex.

Certify introduces **continuous code certification** — measurable quality scores with certification that expires intentionally. When certification lapses, code must be re-evaluated against current standards.

Instead of treating quality as a one-time event, Certify makes it a **continuous process of trust, verification, and renewal**.

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

---

## Quick Start

```bash
cd your-repo

# 1. Bootstrap — creates config, policies, and CI workflows
certify init

# 2. Discover — finds every function, method, type, and file
certify scan

# 3. Certify — collects evidence, evaluates, scores
certify certify

# 4. Report — generates your report card
certify report --format full
```

Your report card is at `.certification/REPORT_CARD.md`.

---

## What You Get

### Report Card

A complete per-unit certification of your entire codebase:

```
# 🟢 Certify — Full Report

## Summary
| Overall Grade | 🟢 B |
| Total Units   | 474  |
| Pass Rate     | 100% |

## Dimension Averages
| correctness              | 95.0% | ██████████████████░░ |
| maintainability          | 93.3% | ██████████████████░░ |
| readability              | 92.4% | ██████████████████░░ |
| testability              | 90.0% | █████████████████░░░ |
| security                 | 80.0% | ████████████████░░░░ |

## All Units (organized by directory)
| Score         | function | B | 86.7% | certified | 2026-06-07 |
| CertifyUnit   | function | B | 85.6% | certified | 2026-06-07 |
... every unit in your repo
```

### Certification Badge

Add a live badge to your README — it updates automatically:

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

---

## Quality Dimensions

Every code unit is scored across 9 quality dimensions:

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

---

## Certification Status

Certifications are **time-bound** — they expire by design.

| Status | Meaning |
|--------|---------|
| 🟢 **Certified** | Meets all required policies |
| 🟡 **Certified with Observations** | Acceptable but with minor issues |
| 🟠 **Probationary** | Requires improvement soon |
| 🔴 **Decertified** | Fails required policies |
| ⚪ **Expired** | Certification window has elapsed, needs recertification |
| **Exempt** | Explicitly excluded by human override |

Default certification window: **90 days**. Risk factors adjust the window — high churn shortens it, stable code extends it.

---

## Commands

| Command | Description |
|---------|-------------|
| `certify init` | Bootstrap `.certification/` with config and policies |
| `certify scan` | Discover all certifiable code units |
| `certify certify` | Evaluate, score, and certify units |
| `certify report` | Generate report card and badge |
| `certify expire` | Mark overdue certifications as expired |
| `certify version` | Show version |

### Flags

```bash
certify certify --skip-agent         # deterministic only, no LLM review
certify certify --batch 20           # process 20 units at a time
certify certify --diff-base main     # only changed files (for PRs)
certify certify --target internal/   # scope to specific paths

certify report --format full         # complete report card (markdown)
certify report --format card         # terminal report card
certify report --format json         # machine-readable
certify report --badge               # print README badge snippet
certify report --output report.md    # write to file
```

---

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

### Policy Packs

Add YAML policy packs to `.certification/policies/`:

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

Supplement deterministic evidence with LLM-powered code review:

```yaml
agent:
  enabled: true
  provider:
    type: openrouter
    api_key_env: OPENROUTER_API_KEY
```

Uses Apache 2.0 licensed models (Qwen, Mistral). Agent review supplements — it never overrides — deterministic evidence. Certify works fully without it.

---

## Language Support

| Language | Adapter | Discovery |
|----------|---------|-----------|
| **Go** | Full | Functions, methods, types via `go/ast` |
| **TypeScript** | Basic | Classes, functions, exports via regex |
| **Everything else** | File-level | One code unit per file |

---

## Repository Structure

```
.certification/
├── config.yml          # configuration
├── policies/           # policy packs
├── records/            # per-unit certification records
├── overrides/          # human governance overrides
├── REPORT_CARD.md      # ← the report card
└── badge.json          # shields.io badge endpoint
```

---

## Documentation

- [Brand Guide](docs/brand.md) — identity, terminology, visual direction
- [Architecture](docs/architecture.md) — system design and package structure
- [Policy Authoring](docs/policy-authoring.md) — writing custom policy packs
- [Troubleshooting](docs/troubleshooting.md) — common issues and solutions
- [Contributor Guide](CLAUDE.md) — development setup

---

## License

[MIT](LICENSE)
