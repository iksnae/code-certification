---
title: Quick Start
description: Certify your first repository in 4 commands.
---

## Step 1 — Bootstrap

Navigate to your repository and initialize Certify:

```bash
cd your-repo
certify init
```

This creates:

```
.certification/
├── config.yml          # configuration
├── policies/
│   ├── global.yml      # universal policy pack
│   └── go.yml          # language-specific (auto-detected)
```

It also generates GitHub Actions workflows in `.github/workflows/`.

## Step 2 — Discover

Scan your repository to find all certifiable code units:

```bash
certify scan
```

Output:

```
  Detecting languages...
  • go: 47 files (config found)
  Go adapter: 182 symbols
✓ Discovered 195 code units (saved to .certification/index.json)
```

Certify finds functions, methods, types, and files using language-aware adapters.

## Step 3 — Certify

Evaluate every unit against your policies:

```bash
certify certify
```

This collects evidence (lint results, test status, git history, complexity metrics), evaluates against policy packs, scores across 9 quality dimensions, and assigns certification status.

```
  Collected 4 repo-level evidence items
  Processing... 195/195
✓ Certified 195/195 units
```

:::tip
Use `--skip-agent` to run deterministic-only certification (no LLM review). This is the default when no OpenRouter API key is configured.
:::

## Step 4 — Report

Generate your report card:

```bash
certify report --format full
```

Your complete report card is saved to `.certification/REPORT_CARD.md`.

For a quick terminal summary:

```bash
certify report --format card
```

```
╔══════════════════════════════════════════════════════════════╗
║              CERTIFY — REPORT CARD                          ║
╠══════════════════════════════════════════════════════════════╣
║       Overall Grade:  🟢 B+       Score: 87.3%              ║
║  Total Units:     195       Pass Rate:   100.0%              ║
╚══════════════════════════════════════════════════════════════╝
```

## What's Next?

- [Add a badge to your README →](/code-certification/reference/badge/)
- [Customize policies →](/code-certification/reference/policies/)
- [Set up CI integration →](/code-certification/reference/ci/)
- [Read your first full report card →](/code-certification/guides/first-report/)
