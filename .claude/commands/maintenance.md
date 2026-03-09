---
description: Run project maintenance checks and report health status
---

# Purpose

Execute project maintenance tasks: verify build, test, lint status, and report overall health.

## Workflow

1. Read `CLAUDE.md` for project context
2. Check git repository health: `git status` (clean working tree)
3. Check Go module: `go mod tidy` (verify no changes needed), if go.mod exists
4. Run lint: `just lint`
5. Run tests: `just test`
6. Check build: `just build`
7. Report to user

## Report

**Status**: SUCCESS or FAILED

**Health Checks**:

- Git repository: [clean/dirty]
- Go module: [tidy/needs tidy/not initialized]
- Lint: [pass/fail/no Go files]
- Tests: [pass/fail/no test files]
- Build: [pass/fail/no source]

**What worked**:

- [completed actions]

**What failed** (if any):

- [errors with context]

**Next steps**:

- [what to do now]
