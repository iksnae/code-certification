---
title: CLI Reference
description: Complete reference for all Certify commands and flags.
---

All commands support `--workspace` for multi-repo operation across git submodules.

## Commands

### `certify init`

Bootstrap certification in a repository.

```bash
certify init [--path <dir>] [--pr] [--workspace]
```

| Flag | Description |
|------|-------------|
| `--path` | Repository root (default: current directory) |
| `--pr` | Create initialization as a pull request via `gh` CLI |
| `--workspace` | Discover submodules and initialize each one |

Creates:
- `.certification/config.yml`
- `.certification/policies/` with auto-detected language policies
- `.github/workflows/` with PR, nightly, and weekly workflows

### `certify scan`

Discover all certifiable code units.

```bash
certify scan [--path <dir>] [--workspace]
```

| Flag | Description |
|------|-------------|
| `--path` | Repository root (default: current directory) |
| `--workspace` | Scan all configured submodules |

Saves discovered units to `.certification/index.json`.

### `certify certify`

Evaluate, score, and certify code units.

```bash
certify certify [flags]
```

| Flag | Description |
|------|-------------|
| `--skip-agent` | Deterministic only, no LLM review |
| `--batch <n>` | Process n units per run (0 = all) |
| `--reset-queue` | Rebuild work queue from scratch |
| `--target <path>` | Only certify units under given path(s) |
| `--diff-base <ref>` | Only certify files changed since git ref |
| `--path <dir>` | Repository root (default: current directory) |
| `--workspace` | Certify all configured submodules |

The work queue is persistent — interrupted runs resume where they left off.

### `certify report`

Generate certification reports.

```bash
certify report [flags]
```

| Flag | Description |
|------|-------------|
| `--format <fmt>` | Output format: `card` (default), `full`, `json`, `text`, `site` |
| `--badge` | Print shields.io badge markdown for your README |
| `--site` | Generate a static HTML site (shorthand for `--format site`) |
| `--output <file>` | Write to file instead of stdout |
| `--detailed` | Include dimension breakdowns, risk analysis, expiring units |
| `--path <dir>` | Repository root |
| `--workspace` | Generate workspace-level report across submodules |

Every run saves `.certification/REPORT_CARD.md` and `.certification/badge.json`.

### `certify architect`

AI-powered architectural review of your entire codebase.

```bash
certify architect [flags]
```

| Flag | Description |
|------|-------------|
| `--model <id>` | Override model for all phases |
| `--phase <n>` | Run only phase n (1–6, default: all) |
| `--output <file>` | Output file path (default: `.certification/ARCHITECT_REVIEW.md`) |
| `--verbose` | Print full LLM responses |
| `--path <dir>` | Repository root |
| `--workspace` | Workspace-level review across submodules |

Requires an AI provider. See [Architect Review →](/code-certification/advanced/architect-review/)

### `certify expire`

Mark overdue certifications as expired.

```bash
certify expire [--path <dir>] [--workspace]
```

Checks all records against their expiration dates and updates status to `expired` where the certification window has elapsed.

### `certify models`

List available models from an AI provider.

```bash
certify models [flags]
```

| Flag | Description |
|------|-------------|
| `--provider-url <url>` | Provider API base URL |
| `--api-key-env <var>` | Environment variable containing the API key |

Auto-detects providers if no URL is specified. Works with any OpenAI-compatible endpoint.

### `certify review`

Generate PR review annotation for GitHub Actions.

```bash
certify review [--path <dir>]
```

Formats certification results as a PR comment. Used in CI workflows to post certification summaries on pull requests.

### `certify doctor`

Check setup and diagnose issues.

```bash
certify doctor [--path <dir>]
```

Runs health checks on:
- **Environment** — Go compiler, Git
- **Project setup** — `.certification/` directory, config, policies, index, records, report card, badge
- **Configuration** — validates `config.yml`, checks agent provider settings, scope patterns
- **Policy packs** — validates all policy YAML files
- **Optional tools** — `golangci-lint`, `gh` CLI
- **AI providers** — auto-detects cloud and local providers

Exits with code 1 if any check fails.

### `certify onboard`

Interactive onboarding guide.

```bash
certify onboard [--path <dir>]
```

Shows a step-by-step checklist for setting up Certify in your project. Checks which steps are already complete and tells you what to do next:

1. Initialize (certify init)
2. Discover code units (certify scan)
3. Run certification (certify certify)
4. Generate report card (certify report)
5. Architect review — optional (certify architect)
6. Add badge to README (certify report --badge)

Re-run at any time to see your progress.

### `certify version`

Print version information.

```bash
certify version
```

## Workspace Mode

The `--workspace` flag enables multi-repo operation across git submodules. When set, commands discover all submodules via `git submodule status` and operate on each configured one.

```bash
certify init --workspace       # initialize all submodules
certify scan --workspace       # scan all submodules
certify certify --workspace    # certify all submodules
certify report --workspace     # aggregate workspace report card
certify architect --workspace  # holistic architecture review
certify expire --workspace     # expire across all submodules
```

See [Workspace Mode →](/code-certification/advanced/workspace/)
