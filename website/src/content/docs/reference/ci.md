---
title: CI Integration
description: GitHub Actions workflows for continuous certification.
---

`certify init` generates three GitHub Actions workflows in `.github/workflows/`.

## PR Certification

**`certification-pr.yml`** — Runs on every pull request.

- Certifies only files changed in the PR (`--diff-base origin/main`)
- Posts a certification summary as a PR comment
- Fast — only processes changed units

## Nightly Sweep

**`certification-nightly.yml`** — Runs on a cron schedule (default: 2 AM UTC).

- Runs `certify expire` to mark overdue certifications
- Commits updated records back to the repository
- Keeps certification state current without manual intervention

## Weekly Report

**`certification-weekly.yml`** — Runs weekly (default: Monday 6 AM UTC).

- Full certification run across the entire repository
- Generates updated report card and badge
- Commits `.certification/REPORT_CARD.md` and `badge.json`
- Uploads report as a workflow artifact

## Required Secrets

| Secret | Required | Purpose |
|--------|----------|---------|
| `GITHUB_TOKEN` | Auto | Provided by GitHub Actions — used for commits and PR comments |
| `OPENROUTER_API_KEY` | Optional | Enables agent-assisted review |

## Customizing Workflows

The generated workflows are standard GitHub Actions YAML. You can customize:

- **Schedule** — Change cron expressions
- **Triggers** — Add branch filters or path filters
- **Steps** — Add notifications, Slack alerts, or other integrations

## Self-Hosted CI

If you're not using GitHub Actions, the same commands work in any CI:

```bash
# Install
go install github.com/iksnae/code-certification/cmd/certify@latest

# Run
certify scan
certify certify --skip-agent
certify report --format full
```
