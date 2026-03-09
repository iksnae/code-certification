---
name: repo-root-cleanup
description: Clean up clutter in the repository root by identifying and safely removing or relocating generated, temporary, and tool-specific files.
---

# Repository Root Cleanup Instructions

When the user asks to clean up the repository root, tidy the project, or remove clutter from the top-level directory, follow these steps **carefully and conservatively**.

## Goals

- Improve repository hygiene and readability
- Remove or relocate generated, temporary, or accidental files
- Never delete source code or critical config without confirmation

## Step-by-Step Process

### 1. Identify Repository Root

- Confirm the current workspace root
- List **only top-level files and directories**
- Do not recurse unless explicitly asked

### 2. Categorize Items

Classify each root item into one of these groups:

#### ✅ Keep (do not modify)
- `.git/`
- `.roo/`
- `src/`, `app/`, `packages/`, `lib/`
- `README*`, `LICENSE*`, `CONTRIBUTING*`
- `package.json`, `pyproject.toml`, `go.mod`, `Cargo.toml`
- `.gitignore`, `.editorconfig`, `.env.example`

#### 🧹 Cleanup Candidates (safe but confirm)
- Build artifacts: `dist/`, `build/`, `out/`, `coverage/`
- Tool output: `.turbo/`, `.next/`, `.vite/`, `.cache/`
- Logs: `*.log`
- OS junk: `.DS_Store`, `Thumbs.db`
- Temporary files: `tmp/`, `temp/`, `*.tmp`, `*.bak`

#### ❓ Review Before Action
- Unfamiliar root files
- One-off scripts
- Large binaries or archives
- Stray config files

### 3. Propose a Cleanup Plan

Before making changes:

- Summarize findings in a table or list
- Clearly state:
  - What will be **deleted**
  - What will be **moved** (and where)
  - What will be **kept**
- Ask for confirmation if any deletion is non-trivial

### 4. Execute Carefully

Once approved:

- Prefer deletion only for clearly disposable files
- Suggest moving reusable artifacts into:
  - `scripts/`
  - `tools/`
  - `docs/`
- Never rewrite history or touch `.git/`

### 5. Optional Enhancements

If appropriate, suggest:
- Updating `.gitignore`
- Adding a `scripts/clean.*` helper
- Documenting build artifacts in `README.md`

## Safety Rules (Non-Negotiable)

- ❌ Never delete source code by default
- ❌ Never delete config files without explanation
- ❌ Never assume generated vs manual — verify
- ✅ Ask when unsure
- ✅ Explain every action

## Output Format

Always conclude with:
- A brief summary of actions taken
- Any follow-up recommendations
