---
name: incremental-commit-all
description: Examine git history and working tree changes, then stage and make incremental, meaningful commits (with clear messages) until the repo is clean; ask the user to confirm whenever a change seems ambiguous or potentially should not be committed.
---

# Incremental Commit All Changes (Safe + Meaningful)

## Goal
Turn the current working tree into a clean state by creating a sequence of small, logically grouped commits with meaningful messages, while **avoiding accidental commits** (secrets, generated artifacts, local-only configs, WIP junk). When uncertain, **ask the user** before proceeding.

## Pre-flight Safety Checks
1. Confirm this is a git repo:
   - `git rev-parse --show-toplevel`
2. Identify current branch + cleanliness baseline:
   - `git status --porcelain=v1`
   - `git branch --show-current`
3. Collect context:
   - Recent history + message style:
     - `git log -n 20 --oneline --decorate`
     - If repo seems to use Conventional Commits (e.g., `feat:`, `fix:`), follow it.
   - Current changes overview:
     - `git diff --stat`
     - `git diff`
     - `git diff --cached --stat` (should be empty at start typically)
     - `git status`

## Change Classification (Decide what’s commit-worthy)
For each item from `git status --porcelain` (tracked modifications, untracked files, deletions, renames):
1. Determine whether it is likely:
   - **Source/code change** (usually commit)
   - **Test change** (usually commit)
   - **Docs change** (usually commit)
   - **Config change** (often commit, but confirm if local-only)
   - **Generated/build artifacts** (usually do NOT commit)
   - **Secrets/credentials** (NEVER commit)
2. Quick checks:
   - Is it ignored? `git check-ignore -v <path>` (if ignored, generally don’t add)
   - Does it look like a secret or key material?
     - Files like: `.env`, `*.pem`, `*.key`, `id_rsa`, `credentials.json`, `*secret*`, `*token*`
     - If suspected, STOP and ask user what to do; suggest adding to `.gitignore` if appropriate.

## Grouping Strategy (How to split into incremental commits)
Prefer multiple small commits over one big commit.

Create a commit plan by grouping changes into coherent units, in this order:
1. **Pure refactors / formatting** (no behavior change)
2. **Mechanical changes** (renames, moves)
3. **Feature work** (one feature slice at a time)
4. **Bug fixes**
5. **Tests**
6. **Docs**
7. **Chores** (tooling, CI, dependency bumps)
8. **Lockfiles** (only when intentionally updated; otherwise confirm)

Guidelines:
- Don’t mix refactor + feature in one commit if you can avoid it.
- Keep commits “reviewable”: ideally one clear purpose.
- If a file includes multiple unrelated hunks, use patch staging to split.

## Staging Workflow (Incremental)
Repeat until `git status --porcelain` is empty:

### A) Propose the next commit
1. Pick the smallest logical group not yet committed.
2. Show evidence for what will be included:
   - Files: `git diff --name-status`
   - Detailed diff for selected files: `git diff -- <paths...>`
3. If hunks are mixed:
   - Use patch staging: `git add -p -- <paths...>` (or equivalent interactive staging tool available in the environment)

### B) Decide whether to ask the user (Mandatory when unclear)
Before staging/committing, ask the user to confirm if ANY of these are true:
- Untracked files that could be local-only (example: `.env`, `.vscode/`, `*.log`, local cache dirs)
- Large generated output (dist/, build/, coverage/, `.next/`, `.turbo/`, etc.)
- Minified bundles or vendor blobs not normally tracked
- Lockfile changes (`package-lock.json`, `pnpm-lock.yaml`, `yarn.lock`, `poetry.lock`, etc.) where it’s not obvious they’re intended
- Changes that look like debug prints, temporary scaffolding, “WIP”
- Suspected secrets/keys/tokens
- The correct commit grouping is ambiguous (could go in more than one commit)

When asking:
- Present a short summary of the questionable items.
- Provide recommended action (commit vs ignore vs split vs revert).
- Ask a direct question with options.

### C) Create a meaningful commit message
Derive the message from:
- The repo’s recent message style (`git log -n 20`)
- The scope of the change

Preferred message structure:
- If Conventional Commits style is present:
  - `feat(scope): <summary>`
  - `fix(scope): <summary>`
  - `refactor(scope): <summary>`
  - `test(scope): <summary>`
  - `docs(scope): <summary>`
  - `chore(scope): <summary>`
- Otherwise:
  - Imperative mood: “Add …”, “Fix …”, “Refactor …”
  - Keep first line ≤ 72 chars when possible

Always ensure the message answers: “What changed, and why?”

### D) Commit
1. Verify staged contents:
   - `git diff --cached --stat`
   - `git diff --cached`
2. Commit:
   - `git commit -m "<message>"`
   - If needed, include a short body (multi-line) explaining why.

### E) Post-commit validation
After each commit:
- `git status --porcelain`
- If remaining changes exist, repeat with the next group.

## Edge Cases + Rules
- **Never commit secrets.** If detected, stop and ask user; recommend remediation (rotate keys, add ignore, remove from history if already committed).
- **Don’t commit broken state** unless the repo history suggests WIP commits are acceptable; otherwise ensure tests/build still make sense.
- If the repo uses pre-commit hooks/formatters and they modify files, include those changes in an appropriate commit (or rerun staging plan).
- If there are submodules:
  - Treat updates carefully; ask user to confirm submodule pointer changes.

## Final Output (When finished)
Provide a concise summary:
- Number of commits created
- Commit SHAs + messages (`git log -n <N> --oneline`)
- Any files intentionally left uncommitted (should be none unless user approved)
- Any `.gitignore` updates suggested/performed
