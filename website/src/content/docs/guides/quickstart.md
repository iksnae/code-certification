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
│   ├── go-standard.yml # universal Go rules
│   └── go-library.yml  # library-specific (auto-detected)
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

This collects evidence (lint results, test status, git history, AST structural analysis, complexity metrics), evaluates against policy packs, scores across 9 quality dimensions, and assigns certification status.

```
  Collected 4 repo-level evidence items
  Processing... 195/195
✓ Certified 195/195 units
```

:::tip[AI-Assisted Mode]
If `OPENROUTER_API_KEY` is set in your environment or as a GitHub secret, Certify automatically enables conservative AI-assisted review — no config changes needed. Use `--skip-agent` to disable it for a single run.
:::

## Step 4 — Report

Generate your report card:

```bash
certify report
```

```
╔══════════════════════════════════════════════════════════════╗
║              CERTIFY — REPORT CARD                          ║
╠══════════════════════════════════════════════════════════════╣
║       Overall Grade:  🟢 B+       Score: 87.3%              ║
║  Total Units:     195       Pass Rate:   100.0%              ║
╚══════════════════════════════════════════════════════════════╝
```

For the complete markdown report card:

```bash
certify report --format full
```

For a browsable static HTML site:

```bash
certify report --site
```

## Step 5 — Architect Review (Optional)

If you have an AI provider configured, run a 6-phase architectural analysis:

```bash
certify architect
```

This analyzes your package structure, dependencies, quality patterns, test strategy, and security posture — producing a comprehensive `ARCHITECT_REVIEW.md` with prioritized recommendations and projected metric improvements.

## What's Next?

- [Add a badge to your README →](/code-certification/reference/badge/)
- [Customize policies →](/code-certification/reference/policies/)
- [Set up CI integration →](/code-certification/reference/ci/)
- [Read your first full report card →](/code-certification/guides/first-report/)
- [Run an architect review →](/code-certification/advanced/architect-review/)
- [Configure workspace mode →](/code-certification/advanced/workspace/)
