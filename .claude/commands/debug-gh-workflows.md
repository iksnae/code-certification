# GH Workflow Debugging

Resolve failing GitHub Actions workflows using the `gh` CLI.  
Run `debug-gh-workflows $ARGUMENTS` first, then iteratively investigate, fix, push, and watch runs until passing.

## Instructions

- Start with `debug-gh-workflows $ARGUMENTS` to collect failing workflows, run IDs, and context.
- Always work from the latest failing run on the relevant branch/PR.
- Prefer `gh` CLI over manual inspection. Use `gh api` only if data isn’t exposed.
- Capture exact failing job, step, and error text before changing anything.
- Keep fixes minimal and root-cause focused.
- After changes: commit → push → trigger workflow → watch run.
- Repeat until all workflows pass.

## Relevant Commands

Primary commands for investigation:

- `gh workflow list`
- `gh workflow view <workflow>`
- `gh run list --workflow <workflow> --limit 20`
- `gh run view <run-id>`
- `gh run view <run-id> --log-failed`
- `gh run view <run-id> --json jobs`
- `gh run download <run-id> --dir .artifacts`
- `gh run watch <run-id> --exit-status --compact`
- `gh workflow run <workflow> --ref <branch>` (if manual trigger needed)

Use API fallback only when necessary:

- `gh api repos/{owner}/{repo}/actions/runs/<run_id>`
- `gh api repos/{owner}/{repo}/actions/workflows/<workflow_id>/runs`

## Debug Loop

### 1. Discover Failures
- Execute `debug-gh-workflows $ARGUMENTS`.
- Identify failing workflow(s), branch/PR, latest run ID.

### 2. Inspect Workflow + Runs
- Confirm workflow definition with `gh workflow view`.
- List recent runs and select latest failure.

### 3. Examine Logs
- Start with `gh run view <run-id> --log-failed`.
- If needed, inspect jobs via JSON output.
- Download artifacts if logs are insufficient.

### 4. Root Cause Analysis
Classify failure type:
- Tests/build errors
- Dependency/cache issues
- Secrets/permissions
- Environment/config mismatch
- Workflow YAML logic
- Flaky infrastructure

Identify minimal corrective change.

### 5. Implement Fix
- Modify code/workflow/scripts accordingly.
- Run local validation where applicable.
- Ensure no unrelated changes.

### 6. Push + Trigger
- Commit clearly.
- Push branch.
- Trigger workflow if required.

### 7. Watch Execution
- Monitor using:
  - `gh run watch <run-id> --exit-status --compact`
- Confirm successful completion.

### 8. Repeat Until Passing
- If failure persists, compare logs/artifacts.
- Refine fix and rerun loop until green.

## Validation Commands

Run these before concluding:

- `gh run list --limit 5`
- `gh run view <latest-run-id>`
- `gh run watch <latest-run-id> --exit-status`

All must complete without failures.

## Notes

- Prefer `--log-failed` first for efficiency.
- Use artifacts for deep debugging.
- Stabilize flaky workflows (pin versions, improve caching, deterministic setup).
- Avoid overengineering fixes.