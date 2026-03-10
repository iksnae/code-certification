# Release v0.1.4

**Date:** 2026-03-09

## Highlights

Per-unit report cards are now a first-class deliverable — every certified unit gets its own detailed markdown report, and all links from the report card resolve correctly on GitHub.

## What's Changed

### Bug Fixes

- **fix: broken unit report links in report card** — Unit name links in `REPORT_CARD.md` previously pointed to files in `.certification/reports/` which were gitignored, causing 404s on GitHub. Reports are now tracked in git and all links resolve correctly. ([#7c0397b](../../commit/7c0397b))

- **fix: missing inline anchors for certified units** — The report card's `<details>` blocks were only generated for units with observations, leaving certified units without anchor targets. All units now get inline anchors for consistent back-navigation from individual reports. ([#7c0397b](../../commit/7c0397b))

### Features

- **feat: per-unit report cards** — Each unit gets a standalone markdown report with identity, certification status, dimension scores, AI assessment, suggestions, and observations. Reports link back to the main report card. ([#6a451cb](../../commit/6a451cb))

- **feat: deep review for local models** — Local LLM providers now support deep review stages alongside remote providers. ([#6a451cb](../../commit/6a451cb))

- **feat: unit table links to per-unit detail sections** — Report card unit tables link directly to expandable detail sections for quick navigation. ([#df2cdd4](../../commit/df2cdd4))

- **feat: AI guidance in report card + prescreen suggestions** — Report cards now include an AI Insights section with top suggestions aggregated across all units, powered by the configured LLM. ([#84930b9](../../commit/84930b9))

## Upgrade Notes

- `.certification/reports/` is no longer gitignored. After upgrading, run `certify report` to generate unit reports, then commit the new `reports/` directory.
- No breaking changes to CLI flags or configuration.

## Full Changelog

```
7c0397b fix: track unit reports and ensure all units have inline anchors
6a451cb feat: deep review for local models + per-unit report cards
df2cdd4 feat: unit table links to per-unit detail sections in report card
84930b9 feat: AI guidance in report card + prescreen suggestions
```
