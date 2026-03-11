# Chore: Documentation & Marketing Accuracy Audit

## Chore Description

Review and update all public-facing documentation, the Astro website, the VSCode extension README, and the root README to accurately reflect the current state of the CLI (v0.7.0) and VSCode extension (v0.4.0). The audit identified the following categories of staleness:

### Missing commands/features from documentation

1. **`certify architect`** command — AI-powered architectural review. Not mentioned anywhere in docs, website CLI reference, or README command table.
2. **`certify models`** command — List available models from an AI provider. Not in website CLI reference or docs/README.md CLI table.
3. **`--workspace` flag** — Multi-repo operation across git submodules. Not documented anywhere except the root command long description.
4. **`--site` flag on `certify report`** — Missing from website CLI reference page (present in root README and docs/README.md but not the website reference).
5. **Algorithmic complexity metrics** — 5 new structural metrics (`algo_complexity`, `loop_nesting_depth`, `recursive_calls`, `nested_loop_pairs`, `quadratic_patterns`) not mentioned in website dimensions page or architecture docs.
6. **Const-like var detection** — Not documented; users may wonder why their lookup tables no longer count as mutable state.
7. **Graduated scoring** — Git history graduation (op_quality/change_risk), file-level readability thresholds, and complexity tiers are not documented.

### Stale version numbers and examples

8. **Website installation page** — Shows `certify v0.1.0` as expected version output. Should show `v0.7.0`.
9. **VSCode extension README** — References `certify-vscode-0.1.0.vsix` for VSIX install. Should be `0.4.0`.
10. **Website VSCode extension guide** — References `certify-vscode-0.1.0.vsix`. Should be `0.4.0`.
11. **Website quickstart** — Example output shows `B+` grade, `195 units`. Self-certification is now `A-`, `748 units`.
12. **Website index (landing page)** — Example report card shows `B` grade. Our self-cert is `A-`.
13. **Root README "What You Get" section** — Shows `B` grade example with `474 units`. Should reflect current reality.
14. **Website report-card page** — Shows `B+` grade, `476 units` in example summary.

### Stale architecture/evidence documentation

15. **Architecture page (docs/ and website)** — `evidence` package description doesn't mention structural analysis (AST-based metrics, algorithmic complexity, const-like detection). Only mentions "cyclomatic complexity."
16. **Architecture page** — CLI diagram missing `architect` and `models` commands.
17. **Website dimensions page** — Security dimension says "dependency vulnerabilities, hardcoded credentials" but actual evidence is `global_mutable_count`. Performance says "resource usage patterns" but actual evidence is loop nesting depth, recursive calls.
18. **docs/README.md** — Agent review section says "circuit breaker" with "5 consecutive failures" but the code uses 3 failures.
19. **Website agent-review page** — Says "circuit breaker" with "5 consecutive agent calls fail" — should be 3.

### Missing documentation pages

20. **No `certify architect` guide or reference page** on the website.
21. **No workspace mode documentation** on the website.

## Relevant Files

Use these files to resolve the chore:

- `README.md` — Root project README. Needs: architect/models/workspace commands in table, updated example grades/unit counts, mention algorithmic complexity in Performance dimension row.
- `docs/README.md` — Main docs page. Needs: architect/models commands in CLI table, `--site` and `--workspace` flags, fix circuit breaker count (5→3).
- `docs/architecture.md` — Architecture overview. Needs: architect/models in CLI diagram, updated evidence package description with structural analysis + algo complexity.
- `website/src/content/docs/index.mdx` — Landing page. Needs: updated example grades in report card snippet.
- `website/src/content/docs/guides/installation.md` — Needs: version output updated from `v0.1.0` to current.
- `website/src/content/docs/guides/quickstart.md` — Needs: updated example output grades/counts.
- `website/src/content/docs/guides/vscode-extension.md` — Needs: VSIX version `0.1.0` → `0.4.0`, mention performance_appropriateness dimension now always visible.
- `website/src/content/docs/reference/cli.md` — Needs: `certify architect`, `certify models` commands, `--site` flag on report, `--workspace` persistent flag.
- `website/src/content/docs/reference/report-card.md` — Needs: updated example summary (grade/units), mention site report format.
- `website/src/content/docs/concepts/dimensions.md` — Needs: accurate evidence descriptions for Security, Performance, Operational Quality, Change Risk reflecting actual metrics.
- `website/src/content/docs/advanced/agent-review.md` — Needs: circuit breaker count 5→3, mention architect review command.
- `vscode-certify/README.md` — Needs: VSIX version `0.1.0` → `0.4.0`.

### New Files

- `website/src/content/docs/reference/architect.md` — New reference page for the `certify architect` command.
- `website/src/content/docs/guides/workspace.md` — New guide for workspace (multi-repo) mode.

## Step by Step Tasks

IMPORTANT: Execute every step in order, top to bottom.

### Step 1: Update root README.md

- Add `certify architect` and `certify models` to the Commands table.
- Add `--workspace` to the Flags section or add a new "Workspace Mode" section.
- Update "What You Get → Report Card" example to show `A-` grade and `748` units.
- Update "Quality Dimensions" table → Performance row: change "Algorithmic complexity indicators" to "Algorithmic complexity (loop nesting, recursion, quadratic patterns)".
- Verify all example output reflects current CLI behavior.

### Step 2: Update docs/README.md

- Add `certify architect` and `certify models` to the CLI Reference table.
- Add `--site` flag to the `certify report` Flags table.
- Add `--workspace` flag documentation (persistent flag on all commands).
- Add `certify architect` Flags table (`--model`, `--output`, `--phase`, `--verbose`, `--path`).
- Fix circuit breaker section: "5 consecutive failures" → "3 consecutive failures".
- Update Agent-Assisted Review to mention `certify architect` command.
- Update evidence package description to mention structural analysis, algorithmic complexity.

### Step 3: Update docs/architecture.md

- Add `architect` and `models` to the CLI diagram box.
- Update `internal/evidence/` description: add "AST-based structural analysis (doc comments, parameters, nesting, function length, panic calls, os.Exit, defer-in-loop, context position, method count, global mutable state, algorithmic complexity)."
- Update `internal/agent/` description: add "Architect review: 6-phase architectural analysis with deterministic snapshot."
- Add `internal/workspace/` package to the diagram and description: "Multi-repo workspace orchestration across git submodules."

### Step 4: Update website landing page (index.mdx)

- Change example report card snippet from `B` grade to `A-`.
- Update "Four Commands" section to mention `certify architect` as a fifth command (or note it as advanced).
- Update Dimensions table Performance row: "Algorithmic complexity (loop nesting, recursion detection)".

### Step 5: Update website installation page

- Change `certify v0.1.0 (abc1234) built 2026-03-09T12:00:00Z` to `certify v0.7.0`.

### Step 6: Update website quickstart page

- Update example `certify scan` output from `195 code units` to a more realistic number.
- Update example `certify report` card output from `B+` / `87.3%` to `A-` / `91.7%`.

### Step 7: Update website CLI reference page

- Add `certify architect` section with all flags (`--model`, `--output`, `--phase`, `--verbose`, `--path`).
- Add `certify models` section with flags (`--provider-url`, `--api-key-env`).
- Add `--site` flag to `certify report` flags table.
- Add `--workspace` persistent flag note to the top of the page.

### Step 8: Update website dimensions page

- **Security** evidence: change "dependency vulnerabilities, hardcoded credentials" to "Global mutable state (package-level `var` declarations). Const-like vars (lookup tables, error sentinels, compiled regexes) are excluded."
- **Performance Appropriateness** evidence: change "resource usage patterns" to "Algorithmic complexity via AST analysis: loop nesting depth (O(n²), O(n³)), recursive calls (O(2^n)), and quadratic anti-patterns (string concatenation in loops)."
- **Operational Quality** evidence: change "contributor count" to "Graduated by commit count (>50→0.95, >20→0.90, >10→0.85) and contributor count."
- **Change Risk** evidence: add "Graduated by author count (≥3→0.95, ≥2→0.90, 1→0.70)."
- **Architectural Fitness** evidence: change "Package structure, dependency patterns" to "API design patterns: context.Context position, god objects (>15 methods). Penalty-only — appears in score only when violations exist."

### Step 9: Update website report-card page

- Update example summary from `B+` / `476` units to realistic numbers.
- Add `site` format row to the formats table: `--site` / `--format site` → Interactive HTML dashboard.

### Step 10: Update website agent-review page

- Fix circuit breaker: "5 consecutive agent calls fail" → "3 consecutive failures".
- Add section mentioning `certify architect` — AI-powered architectural review (distinct from per-unit agent review).

### Step 11: Create website architect reference page

- Create `website/src/content/docs/reference/architect.md` with:
  - Command description, synopsis, all flags
  - 6-phase pipeline description (Architecture Narration → Code Quality → Test Strategy → Security/Ops → Recommendations → Synthesis)
  - Output: `.certification/ARCHITECT_REVIEW.md`
  - Report structure: Part I (deterministic snapshot), Part II (LLM analysis), Part III (recommendations with deltas)
  - Example usage

### Step 12: Create website workspace guide

- Create `website/src/content/docs/guides/workspace.md` with:
  - What workspace mode is (multi-repo across git submodules)
  - `--workspace` flag usage on `certify init`, `certify certify`, `certify report`
  - Submodule discovery and `.certification/` bootstrapping per submodule
  - Aggregate report generation

### Step 13: Update VSCode extension README

- Change VSIX install reference from `certify-vscode-0.1.0.vsix` to `certify-vscode-0.4.0.vsix`.
- Add note about `performance_appropriateness` dimension now always visible in dashboard (v0.4.0+).

### Step 14: Update website VSCode extension guide

- Change VSIX reference from `certify-vscode-0.1.0.vsix` to `certify-vscode-0.4.0.vsix`.
- Add note about `performance_appropriateness` dimension visibility.

### Step 15: Build website and verify

- Run `cd website && npm run build` to verify Astro builds without errors.
- Spot-check generated HTML for broken links or missing content.

### Step 16: Run validation commands

- Run all `Validation Commands` below to confirm zero regressions.

### Step 17: Commit and push

- Commit with message `docs: accuracy audit — update all docs/website/marketing for v0.7.0`.
- Push to trigger the Pages deploy workflow.

## Validation Commands

Execute every command to validate the chore is complete with zero regressions.

```bash
# Go tests still pass (no code changes, but verify nothing broken)
go test ./... -count=1

# Go build still works
go build -o build/bin/certify ./cmd/certify/

# Website builds without errors
cd website && npx astro build 2>&1 | tail -5

# VSCode extension still compiles
cd vscode-certify && npm run lint && npm run build

# Verify no stale version references remain
grep -rn 'v0\.1\.0\|v0\.2\.0\|v0\.3\.0\|v0\.4\.0\|v0\.5\.0\|v0\.6\.' \
  README.md docs/ website/src/content/ vscode-certify/README.md \
  2>/dev/null | grep -v 'CHANGELOG\|node_modules\|package.json\|specs/' || echo "PASS: no stale versions"

# Verify architect command is documented
grep -l 'architect' README.md docs/README.md docs/architecture.md \
  website/src/content/docs/reference/cli.md \
  website/src/content/docs/reference/architect.md

# Verify models command is documented
grep -l 'models' README.md docs/README.md \
  website/src/content/docs/reference/cli.md

# Verify workspace is documented
grep -l 'workspace' README.md docs/README.md \
  website/src/content/docs/guides/workspace.md

# Verify dimensions page has accurate evidence descriptions
grep -c 'loop nesting\|recursive calls\|global_mutable\|mutable state' \
  website/src/content/docs/concepts/dimensions.md

# Verify circuit breaker count is 3 (not 5)
grep -rn 'consecutive.*fail' docs/README.md website/src/content/docs/advanced/agent-review.md | grep -v '3 consecutive'
# ^ should produce no output (all references should say 3)

# Verify VSIX version references are current
grep -rn '0\.1\.0\.vsix' website/src/content/ vscode-certify/README.md
# ^ should produce no output
```

## Notes

- The website uses **Astro Starlight** and deploys to GitHub Pages via `.github/workflows/pages.yml`. Changes to `website/` or `docs/` auto-trigger deployment on push to `main`.
- Example grades/counts in documentation should be realistic but don't need to match self-certification exactly — they serve as illustrative examples. Using `A-` / `91%` / `748 units` is fine since that's the current state, but a generic example repo with `B+` / `195 units` is also acceptable for the quickstart (since new repos won't have 748 units). The key is that the examples shouldn't show capabilities we don't have or grades that are impossible to achieve.
- The `docs/` directory is the authoritative source; the `website/src/content/docs/` mirrors it for the Astro site. Keep them consistent but the website version can have richer formatting (Astro components, tabs, etc.).
- The `certify architect` command is an advanced feature — it should get a reference page but doesn't need to be prominently featured in the quickstart flow.
- **Do NOT change any Go source code.** This is a docs-only chore.
