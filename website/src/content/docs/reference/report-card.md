---
title: Report Card
description: Understanding the Certify report card.
---

The report card is Certify's primary output — a complete per-unit certification of your entire codebase.

## Generate

```bash
certify report --format full
```

Saved to `.certification/REPORT_CARD.md`.

## Sections

### Summary

Overall certification status at a glance:

```
| Overall Grade | 🟢 B+ |
| Overall Score | 87.3% |
| Total Units   | 476   |
| Pass Rate     | 100%  |
```

### Grade Distribution

How units are distributed across grades:

```
| Grade | Count | %     | Bar                              |
| B+    | 328   | 68.9% | ██████████████████████████████████ |
| B     | 148   | 31.1% | ███████████████                    |
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
