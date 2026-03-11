---
title: Workspace Mode
description: Multi-repo certification across git submodules.
---

Certify's workspace mode operates across multiple repositories organized as git submodules. Each submodule is certified independently, then results are aggregated into a workspace-level report.

## When to Use Workspace Mode

Use `--workspace` when your project is composed of multiple git repositories linked as submodules — a common pattern for:

- **Monorepo-style workspaces** with separate repos per service
- **Shared library architectures** where libraries and consumers are separate repos
- **Platform teams** managing multiple related components

## Setup

### 1. Initialize

```bash
cd your-workspace-root
certify init --workspace
```

This discovers all git submodules and runs `certify init` in each unconfigured one. It also creates a workspace-level `.certification/` directory.

### 2. Scan

```bash
certify scan --workspace
```

Scans all configured submodules, discovering code units in each.

### 3. Certify

```bash
certify certify --workspace
```

Runs certification in each submodule. Uses the submodule's own `.certification/config.yml` for policies and agent configuration.

### 4. Report

```bash
certify report --workspace
```

Generates a workspace-level report card that aggregates results across all submodules:

```
🟢 Workspace: A- (91.2%)
  Units: 1,247 · Passing: 1,230 · Failing: 17

  api-service                    🟢 A-  91.8%  612 units
  shared-lib                     🟢 A   94.2%  203 units
  web-frontend                   🟢 B+  88.1%  432 units
```

### 5. Architect Review

```bash
certify architect --workspace
```

Runs a holistic architectural review that treats all submodules as components of a single system. See [Architect Review → Workspace Mode](/code-certification/advanced/architect-review/#workspace-mode).

## How It Works

### Submodule Discovery

Workspace mode uses `git submodule status` to find all submodules. Each submodule is checked for a `.certification/config.yml` — only configured submodules are processed.

### Independent Certification

Each submodule is certified independently with its own:
- Configuration (`.certification/config.yml`)
- Policy packs (`.certification/policies/`)
- Records and state

The workspace doesn't change how individual certification works.

### Aggregation

The workspace report card computes:

| Metric | How computed |
|--------|-------------|
| **Overall Score** | Weighted average by unit count |
| **Overall Grade** | Grade from weighted average score |
| **Total Units** | Sum across all submodules |
| **Pass Rate** | Total passing / total units |

Submodules without certify configured are listed but excluded from score aggregation.

### Workspace Reports

Workspace mode generates:

```
.certification/
├── REPORT_CARD.md           # workspace-level aggregate card
├── reports/
│   ├── index.md             # workspace report tree root
│   ├── api-service.md       # per-submodule summary
│   └── shared-lib.md        # per-submodule summary
```

Each submodule summary links to the submodule's own full report tree.

## Architecture Review in Workspace Mode

The `certify architect --workspace` command goes beyond per-submodule analysis. It builds a `WorkspaceArchSnapshot` that includes:

- **Per-submodule architecture snapshots** — package maps, dependencies, metrics
- **Submodule role classification** — automatically detects service, library, or tool
- **Cross-submodule dependencies** — detected from `go.mod` replace directives
- **Infrastructure files** — workspace-level CI, Docker, build orchestration
- **Aggregate metrics** — total units, weighted score, best/worst submodule

The 6-phase LLM review uses workspace-specific prompts that focus on system-level concerns: integration boundaries, shared library quality, deployment coupling, cross-submodule security surface.

## Workspace-Level Infrastructure

Certify treats workspace-level files as **infrastructure**, not application code:

| File Type | Examples |
|-----------|---------|
| Build orchestration | `Justfile`, `Makefile` |
| CI/CD | `.github/workflows/*.yml` |
| Container | `Dockerfile`, `docker-compose.yml` |
| Deployment | `fly.toml`, `netlify.toml`, `vercel.json` |
| Dev environment | `Tiltfile`, `Vagrantfile` |

These are listed in the architect review snapshot but not individually certified.

## Commands with Workspace Support

| Command | `--workspace` behavior |
|---------|----------------------|
| `certify init` | Initialize all submodules |
| `certify scan` | Scan all configured submodules |
| `certify certify` | Certify all configured submodules |
| `certify report` | Aggregate workspace report card |
| `certify architect` | Holistic system-level review |
| `certify expire` | Expire across all submodules |
