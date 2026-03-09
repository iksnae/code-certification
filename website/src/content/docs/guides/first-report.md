---
title: Your First Report Card
description: Understanding what Certify tells you about your codebase.
---

After running `certify certify`, you can generate your full report card:

```bash
certify report --format full
```

This creates `.certification/REPORT_CARD.md` — a complete per-unit certification of your entire codebase.

## Report Card Structure

### Summary

The top section shows your overall certification status:

| Metric | What it means |
|--------|--------------|
| **Overall Grade** | Weighted average across all units (A through F) |
| **Overall Score** | Percentage score (0–100%) |
| **Total Units** | Number of certifiable code units found |
| **Passing** | Units meeting all policy requirements |
| **Failing** | Units that fail required policies |
| **Pass Rate** | Percentage of units passing |
| **Observations** | Passing units with minor warnings |
| **Expired** | Units whose certification window has elapsed |

### Dimension Averages

Your codebase scored across 9 quality dimensions with visual bars:

```
| correctness         | 95.0% | ██████████████████░░ |
| maintainability     | 93.3% | ██████████████████░░ |
| readability         | 92.4% | ██████████████████░░ |
```

This tells you which quality areas are strongest and where to focus improvement.

### By Language

Breakdown by detected language — each showing unit count, average score, score range, and grade distribution.

### All Units

Every single unit organized by directory:

```
| Score         | function | B+ | 87.4% | certified | 2026-06-07 |
| CertifyUnit   | function | B  | 85.6% | certified | 2026-06-07 |
```

Units with observations or non-passing status get expandable detail sections showing per-dimension scores and specific observations.

## Report Formats

| Format | Command | Use case |
|--------|---------|----------|
| **Full** | `--format full` | Complete markdown report card |
| **Card** | `--format card` | Quick terminal summary |
| **JSON** | `--format json` | Machine-readable for tooling |
| **Text** | (default) | Brief health summary |

## Where Reports Are Saved

Every `certify report` run saves:

- `.certification/REPORT_CARD.md` — full per-unit report card
- `.certification/badge.json` — shields.io badge endpoint

These are committed to your repository so they're always visible on GitHub.
