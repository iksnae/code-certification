---
title: Doctor & Onboard
description: Diagnose setup issues and follow guided onboarding.
---

Certify includes two support commands that help you get set up correctly and troubleshoot problems.

## Onboard — Guided Setup

New to Certify? `certify onboard` shows exactly where you are and what to do next:

```bash
certify onboard
```

```
🚀 Certify Onboarding — /path/to/your-repo

  👉 1. Initialize Certify
     → certify init
       Run this in your repository root
  ⬜ 2. Discover Code Units
     Complete step 1 first
  ⬜ 3. Run Certification
     Complete step 2 first
  ⬜ 4. Generate Report Card
     Complete step 3 first
  ⬜ 5. Architect Review (optional)
     Complete step 3 first
  ⬜ 6. Add Badge to README

  Next step: certify init
```

The checklist detects your progress automatically by checking the filesystem. Steps transition through three states:

| Icon | State | Meaning |
|------|-------|---------|
| ⬜ | Pending | Dependencies not yet met |
| 👉 | Ready | Dependencies met — run this next |
| ✅ | Done | Already completed |

Run `certify onboard` after each step to see your progress update. When everything is done:

```
🚀 Certify Onboarding — /path/to/your-repo

  ✅ 1. Initialize Certify
  ✅ 2. Discover Code Units
  ✅ 3. Run Certification
  ✅ 4. Generate Report Card
  ✅ 5. Architect Review (optional)
  ✅ 6. Add Badge to README

  🎉 All done! Your project is fully onboarded.
     Run 'certify doctor' to check health at any time.
```

### The 6 Steps

| Step | Command | What it does |
|------|---------|-------------|
| 1 | `certify init` | Creates `.certification/` with config, policies, CI workflows |
| 2 | `certify scan` | Discovers all functions, methods, types, files |
| 3 | `certify certify` | Evaluates units against policies, assigns grades |
| 4 | `certify report` | Generates report card, badge, report tree |
| 5 | `certify architect` | AI-powered 6-phase architectural review (optional) |
| 6 | `certify report --badge` | Generates badge.json for your README |

---

## Doctor — Health Checks

`certify doctor` runs comprehensive diagnostics on your setup:

```bash
certify doctor
```

```
🩺 Certify Doctor — /path/to/your-repo

  ── Environment ──
    ✅ Go compiler: go version go1.22.0 darwin/arm64
    ✅ Git: git version 2.43.0
  ── Project Setup ──
    ✅ Configuration file: .certification/config.yml exists
    ✅ Policy packs: 2 policy pack(s): go-standard.yml, go-library.yml
    ✅ Unit index: .certification/index.json (12345 bytes)
    ✅ Certification records: 195 certification record(s)
    ✅ Report card: REPORT_CARD.md exists
    ✅ Badge endpoint: badge.json exists
  ── Configuration ──
    ✅ Configuration valid: mode=advisory, expiry=90d
    ✅ Scope includes: scope.include is empty (includes everything)
  ── Policy Packs ──
    ✅ Policy: go-standard@1.1.0: 13 rule(s)
    ✅ Policy: go-library@1.0.0: 2 rule(s)
  ── Optional Tools ──
    ✅ golangci-lint: golangci-lint has version 2.7.2
    ✅ GitHub CLI (gh): gh version 2.86.0
  ── AI Providers ──
    ✅ Provider: openrouter: https://openrouter.ai/api/v1 (key: 73 chars)
    ✅ Provider: ollama: http://localhost:11434/v1 (local)

  Summary: 18 passed, 0 warnings, 0 failed, 0 skipped

  ✅ Everything looks good!
```

### What Doctor Checks

| Group | Checks |
|-------|--------|
| **Environment** | Go compiler available, Git available |
| **Project Setup** | config.yml, policies dir, index.json, records, report card, badge |
| **Configuration** | YAML valid, expiry settings correct, scope patterns, agent config |
| **Policy Packs** | Each policy YAML validated (name, version, rules, dimensions) |
| **Agent Config** | Base URL set, API key env var available, model configured |
| **Optional Tools** | `golangci-lint` installed, `gh` CLI installed |
| **AI Providers** | Auto-detects OpenRouter, OpenAI, Groq, Ollama, LM Studio |

### Check Statuses

| Icon | Status | Meaning |
|------|--------|---------|
| ✅ | Pass | Check passed |
| ⚠️ | Warning | Non-critical issue — things work but could be better |
| ❌ | Fail | Something is broken — fix before proceeding |
| ⏭️ | Skip | Optional feature not configured (not a problem) |

Failed checks include a `→` fix suggestion telling you exactly what to do.

### Exit Codes

`certify doctor` exits with code **1** if any check fails, making it usable in CI:

```yaml
# In a GitHub Actions workflow
- name: Check Certify setup
  run: certify doctor
```

### When to Use Doctor

- **Before filing a bug** — run doctor first to check your environment
- **After upgrading** — verify everything still works
- **In CI** — catch setup issues before they cause confusing failures
- **Debugging "0 units found"** — doctor checks scope config and index
- **AI not working** — doctor checks provider config, API keys, and connectivity
