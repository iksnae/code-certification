---
title: Introduction
description: What Certify is, why it exists, and how it works.
---

## What is Certify?

Certify is a **continuous code certification** tool. It evaluates every code unit in your repository — functions, methods, types, files — against versioned policies, scores them across 9 quality dimensions, and assigns **time-bound certification** that expires intentionally.

## Why?

CI pipelines validate whether code passes at a specific moment. But they don't answer a harder question:

> **Should this code still be trusted?**

Standards evolve. Dependencies change. Systems grow more complex. Code that once met expectations slowly drifts away from them.

Certify replaces the assumption of permanent trust with **measurable, time-bound certification**. When certification expires, code must be re-evaluated against current standards.

## How It Works

```
Discover → Evaluate → Certify → Report
```

1. **Discover** — Language-aware adapters find every certifiable unit (functions, methods, types, files)
2. **Evaluate** — Deterministic evidence collected from linters, test runners, git history, AST analysis, code metrics
3. **Certify** — Units scored across 9 dimensions, assigned status with expiration date
4. **Report** — Complete report card generated with per-unit breakdown

## What You Get

- **Report Card** — Every unit in your repo with grade, score, and dimension breakdown
- **Certification Badge** — Live shields.io badge for your README
- **Architect Review** — AI-powered 6-phase architectural analysis with grounded metrics
- **Static Site** — Browsable, searchable HTML report for large repos
- **CI Integration** — GitHub Actions for PR review, nightly sweeps, weekly reports
- **Policy-as-Code** — Versioned YAML policies with path scoping
- **Workspace Mode** — Multi-repo certification across git submodules
- **Doctor & Onboard** — Built-in diagnostics and guided setup
- **VSCode Extension** — Inline grades, interactive dashboard, AI provider configuration

## Language Support

| Language | Adapter | Discovery |
|----------|---------|-----------|
| **Go** | Full AST | Functions, methods, types, structural analysis |
| **TypeScript** | Regex | Classes, functions, exports |
| **Everything else** | File-level | One unit per file |

## Next Steps

## Next Steps

- New to Certify? [Install →](/code-certification/guides/installation/) then run `certify onboard` for guided setup
- Already installed? Run `certify doctor` to check your environment
