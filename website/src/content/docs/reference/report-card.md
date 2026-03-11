---
title: Report Card
description: Understanding the Certify report card and report formats.
---

The report card is Certify's primary output — a complete per-unit certification of your entire codebase.

## Generate

```bash
certify report                 # terminal card (default)
certify report --format full   # complete markdown report
certify report --format site   # browsable HTML site
```

## Report Formats

| Format | Command | Output | Use case |
|--------|---------|--------|----------|
| **Card** | `--format card` | Terminal | Quick visual summary |
| **Full** | `--format full` | `.certification/REPORT_CARD.md` | Complete markdown report |
| **Site** | `--format site` or `--site` | `.certification/site/` | Browsable HTML for large repos |
| **JSON** | `--format json` | stdout | Machine-readable for tooling |
| **Text** | `--format text` | stdout | Brief health summary |

Every run also saves `.certification/badge.json` for your README badge.

## Report Card Sections

### Summary

Overall certification status at a glance:

```
| Overall Grade | 🟢 A- |
| Overall Score | 91.8% |
| Total Units   | 816   |
| Pass Rate     | 100%  |
```

### Grade Distribution

How units are distributed across grades:

```
| Grade | Count | %     | Bar                              |
| A     |   412 | 50.5% | ██████████████████████████████████ |
| A-    |   287 | 35.2% | ████████████████████████           |
| B+    |    98 | 12.0% | ████████                           |
| B     |    19 |  2.3% | ██                                 |
```

### Dimension Averages

Average scores across all 9 quality dimensions:

```
| correctness         | 95.0% | ██████████████████░░ |
| maintainability     | 93.3% | ██████████████████░░ |
| readability         | 92.4% | ██████████████████░░ |
```

### By Language

Per-language breakdown with score ranges and grade distribution.

### All Units

Every unit organized by directory, showing:
- Unit name and type (function, method, class, file)
- Grade and score
- Certification status
- Expiration date

Units with observations get expandable detail sections with per-dimension scores.

## Report Tree

In addition to the report card, `certify report` generates a **report tree** — individual markdown files for every unit:

```
.certification/reports/
├── index.md
├── internal/
│   ├── engine/
│   │   └── scorer.go/
│   │       ├── Score.md
│   │       └── CertifyUnit.md
│   └── agent/
│       └── architect.go/
│           └── GatherContext.md
```

Each unit file contains the full certification record with dimensions, evidence, observations, and navigation links.

## Static Site

For large repos, the static site format provides a browsable, searchable experience:

```bash
certify report --site
```

Generates `.certification/site/` with:
- Index page with overall metrics and package list
- Per-package pages with unit tables
- Per-unit pages with full dimension scores
- Client-side search across all units
- Works offline — no external dependencies

Open `site/index.html` in any browser, or serve with:

```bash
cd .certification/site && python3 -m http.server 8080
```

## Grading Scale

| Grade | Score Range | Status |
|-------|------------|--------|
| **A** | ≥ 93% | 🟢 Certified |
| **A-** | ≥ 90% | 🟢 Certified |
| **B+** | ≥ 87% | 🟢 Certified |
| **B** | ≥ 80% | 🟢 Certified |
| **C** | ≥ 70% | 🟡 Observations |
| **D** | ≥ 60% | 🟠 Probationary |
| **F** | < 60% | 🔴 Decertified |
