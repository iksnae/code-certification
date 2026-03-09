# Code Certification System

[![CI](https://github.com/iksnae/code-certification/actions/workflows/ci.yml/badge.svg)](https://github.com/iksnae/code-certification/actions/workflows/ci.yml)
[![Certification](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/iksnae/code-certification/main/.certification/badge.json)](https://github.com/iksnae/code-certification/blob/main/.certification/REPORT_CARD.md)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A repository governance platform that continuously evaluates software code against defined engineering policies and assigns **time-bound certification status** to meaningful code units.

Unlike CI pipelines that only validate whether code passes at a specific moment, Code Certification establishes **explicit trust with expiration**, ensuring that code quality, security posture, maintainability, and architectural compliance are periodically re-evaluated as standards evolve.

## Status

**v1 complete** — 275/275 feature criteria implemented. 282 tests across 15 packages.

| Metric | Value |
|--------|-------|
| Code units | 416 |
| Pass rate | 100% |
| Average score | 0.861 |
| Test count | 282 |
| Packages | 15 |

## Installation

### From Source

```bash
git clone https://github.com/iksnae/code-certification.git
cd code-certification
go build -o certify ./cmd/certify/
```

### Prerequisites

- [Go](https://go.dev/) 1.22+
- [git](https://git-scm.com/)
- [golangci-lint](https://golangci-lint.run/) (optional, for enhanced lint evidence)
- [just](https://github.com/casey/just) command runner (optional, for dev workflows)
- [gh](https://cli.github.com/) GitHub CLI (optional, for PR/issue integration)

## Quick Start

```bash
# 1. Initialize certification in your repository
cd /path/to/your-repo
certify init

# 2. Discover code units
certify scan

# 3. Run certification
certify certify --skip-agent

# 4. View results
certify report
certify report --detailed
certify report --format json
```

## CLI Reference

| Command | Description |
|---------|-------------|
| `certify init` | Bootstrap `.certification/` directory with config, policies, workflows |
| `certify scan` | Discover certifiable code units, save to index |
| `certify certify` | Evaluate units against policies, collect evidence, assign status |
| `certify report` | Generate health report (text, JSON, or detailed) |
| `certify expire` | Mark overdue certifications as expired |
| `certify review` | Generate PR review annotation |
| `certify version` | Show version information |

### Key Flags

```bash
certify certify --skip-agent          # deterministic only, no LLM review
certify certify --batch 20            # process 20 units then stop (incremental)
certify certify --target internal/    # only certify units under internal/
certify certify --diff-base origin/main  # only certify changed files (for PRs)
certify certify --reset-queue         # rebuild queue, start over
certify init --pr                     # create initialization as a pull request
certify report --detailed             # include dimensions, risk, expiring units
```

## How It Works

```
┌─────────────┐     ┌──────────┐     ┌──────────┐     ┌──────────┐
│  Discovery  │────▶│ Evidence │────▶│  Engine  │────▶│  Record  │
│  (scan)     │     │(collect) │     │(certify) │     │ (store)  │
└─────────────┘     └──────────┘     └──────────┘     └──────────┘
       │                  │                │                │
  Go/TS/generic    go vet, test,     Policy eval,     JSON files in
  AST adapters     lint, git stats   9 dimensions,    .certification/
                                     weighted score    records/
```

1. **Discover** — Finds certifiable units (functions, methods, types, files) using language-aware adapters
2. **Collect** — Gathers evidence from linters, test runners, git history, and code metrics
3. **Evaluate** — Checks evidence against versioned policy packs, scores across 9 quality dimensions
4. **Certify** — Assigns time-bound status (certified, probationary, expired, decertified) with grade and score
5. **Report** — Generates health reports with dimension breakdowns, risk analysis, and expiring-soon detection

## Features

- **Language-agnostic** — Go and TypeScript adapters built-in, generic file-level for everything else
- **Policy-as-code** — Versioned YAML policies with language/path targeting and severity levels
- **9 quality dimensions** — Correctness, maintainability, readability, testability, security, architecture, operational quality, performance, change risk
- **Time-bound trust** — Every certification expires; risk factors shorten/lengthen windows
- **Incremental processing** — Persistent work queue processes across multiple runs
- **Agent-assisted review** — Optional LLM review via OpenRouter (open-weight models, Apache 2.0)
- **GitHub integration** — PR certification, nightly sweeps, weekly reports, remediation issues
- **Human governance** — Overrides with required rationale, exemptions, forced review
- **Self-certifying** — The tool certifies its own codebase (416 units, 100% pass)

## Agent-Assisted Review

Optional LLM-powered code review using open-weight models via [OpenRouter](https://openrouter.ai/):

- **Primary**: `qwen/qwen3-coder:free` — Apache 2.0, code-specialized, 262k context
- **Fallback**: `mistralai/mistral-nemo` — Apache 2.0, 12B, clean JSON output

Agent review supplements deterministic evidence — it never overrides tool results. The system works fully without it.

```bash
export OPENROUTER_API_KEY=sk-or-v1-your-key
certify certify --batch 20    # process with agent review
```

## Documentation

- **[docs/README.md](docs/README.md)** — Installation & quickstart
- **[docs/architecture.md](docs/architecture.md)** — System design & package structure
- **[docs/policy-authoring.md](docs/policy-authoring.md)** — Writing custom policies
- **[docs/troubleshooting.md](docs/troubleshooting.md)** — Common issues & solutions
- **[FEATURES.md](FEATURES.md)** — Feature acceptance checklist (275/275 ✅)
- **[PRD.md](PRD.md)** — Full product requirements
- **[STORIES.md](STORIES.md)** — User stories by epic

## Development

```bash
just doctor     # verify development environment
just build      # build CLI → build/bin/certify
just test       # run all tests
just lint       # run golangci-lint
just check      # all quality gates (fmt + vet + lint + test)
just cover      # generate test coverage report
```

## License

[MIT](LICENSE)
