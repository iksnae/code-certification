---
title: CLI Reference
description: Complete reference for all Certify commands and flags.
---

## Commands

### `certify init`

Bootstrap certification in a repository.

```bash
certify init [--path <dir>] [--pr]
```

| Flag | Description |
|------|-------------|
| `--path` | Repository root (default: current directory) |
| `--pr` | Create initialization as a pull request via `gh` CLI |

Creates:
- `.certification/config.yml`
- `.certification/policies/` with auto-detected language policies
- `.github/workflows/` with PR, nightly, and weekly workflows

### `certify scan`

Discover all certifiable code units.

```bash
certify scan [--path <dir>]
```

| Flag | Description |
|------|-------------|
| `--path` | Repository root (default: current directory) |

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

The work queue is persistent — interrupted runs resume where they left off.

### `certify report`

Generate certification reports.

```bash
certify report [flags]
```

| Flag | Description |
|------|-------------|
| `--format <fmt>` | Output format: `text`, `card`, `full`, `json` |
| `--badge` | Print shields.io badge markdown for your README |
| `--output <file>` | Write to file instead of stdout |
| `--detailed` | Add dimension breakdowns to text format |
| `--path <dir>` | Repository root |

Every run saves `.certification/REPORT_CARD.md` and `.certification/badge.json`.

### `certify expire`

Mark overdue certifications as expired.

```bash
certify expire [--path <dir>]
```

Checks all records against their expiration dates and updates status to `expired` where the certification window has elapsed.

### `certify version`

Print version information.

```bash
certify version
```
