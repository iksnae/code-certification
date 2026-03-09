---
description: Verify development environment has required tools
argument-hint: [hil]
---

# Purpose

Verify that the development environment has all required tools and dependencies for the Code Certification System, then report readiness to the user.

## Variables

MODE: $1 (optional - if "hil", run interactive mode)

## Workflow

1. Read `CLAUDE.md` for project context
2. Verify key tools are available:
   - `go` (Go 1.22+)
   - `golangci-lint`
   - `git`
   - `gh` (GitHub CLI, authenticated)
   - `just` (command runner)
3. Check `OPENROUTER_API_KEY` environment variable (info if missing, not error)
4. Run `just doctor` for comprehensive check
5. Report to user

## Report

**Status**: SUCCESS or FAILED

**Environment**:

- go: [version or missing]
- golangci-lint: [version or missing]
- git: [version or missing]
- gh: [version or missing]
- just: [version or missing]

**Optional**:

- OPENROUTER_API_KEY: [set or not set]

**What worked**:

- [completed checks]

**What failed** (if any):

- [errors with context]

**Next steps**:

- [what to do now]
