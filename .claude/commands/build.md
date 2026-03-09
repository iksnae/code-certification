---
description: Implement the plan
argument-hint: [path-to-plan]
---

# Build

Follow the `Workflow` to implement the `PATH_TO_PLAN` then `Report` the completed work.

## Variables

PATH_TO_PLAN: $ARGUMENTS

## Workflow

- If no `PATH_TO_PLAN` is provided, STOP immediately and ask the user to provide it (AskUserQuestion).
- Read and execute the plan at `PATH_TO_PLAN`. Think hard about the plan and implement it into the codebase.
- After each completed step, update `FEATURES.md` — check off any criteria that are now satisfied (`- [ ]` → `- [x]`).
- Include `Refs #N` in commit messages to link to the relevant GitHub epic issue (see `CLAUDE.md` for epic table).

## Report

- Present the `## Report` section of the plan.
- List which `FEATURES.md` criteria were checked off.
- List which GitHub epic steps were completed.
