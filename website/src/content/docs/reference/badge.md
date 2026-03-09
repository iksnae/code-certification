---
title: Badge Setup
description: Add a live certification badge to your README.
---

Certify generates a [shields.io](https://shields.io) endpoint badge that reflects your current certification status. It updates automatically when CI commits new results.

## Get Your Badge

```bash
certify report --badge
```

This prints the markdown snippet for your README:

```markdown
[![Certification](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/OWNER/REPO/main/.certification/badge.json)](https://github.com/OWNER/REPO/blob/main/.certification/REPORT_CARD.md)
```

## How It Works

1. Every `certify report` run generates `.certification/badge.json`
2. The JSON follows the [shields.io endpoint schema](https://shields.io/badges/endpoint-badge)
3. shields.io reads the JSON from your repo's raw content URL
4. The badge renders with your current grade, pass rate, and unit count

## Badge Colors

Colors map to certification status per the [brand guide](/code-certification/docs/brand/):

| Grade | Color | Status |
|-------|-------|--------|
| A | Green (#2E8B57) | 🟢 Certified |
| A- | Green (#3DA06A) | 🟢 Certified |
| B+, B | Steel Blue (#4A6B82) | 🟢 Certified |
| C | Amber (#E0A100) | 🟡 Observations |
| D | Warning (#F59E0B) | 🟠 Probationary |
| F | Red (#DC2626) | 🔴 Decertified |

## Example Badge JSON

```json
{
  "schemaVersion": 1,
  "label": "certification",
  "message": "B+ · 100% · 476 units",
  "color": "4A6B82",
  "namedLogo": "checkmarx",
  "logoColor": "white"
}
```

## Keeping It Updated

The badge updates when `.certification/badge.json` changes in your repository. With the CI workflows from `certify init`, this happens automatically on every push to main.

Click the badge in any README → opens your full report card.
