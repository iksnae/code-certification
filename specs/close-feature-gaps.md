# Chore: Close FEATURES.md Gaps — 48/275 → 275/275

## Chore Description
FEATURES.md has 275 acceptance criteria. Only 48 are checked off (17%). However, the codebase already implements the majority of these — they were never checked off during implementation. This chore has two parts:

**Part A — Audit & Check Off (est. ~120 criteria):** Walk every unchecked criterion, verify it is already satisfied by existing code, and check it off. No code changes needed.

**Part B — Implement Missing Features (est. ~107 criteria):** For criteria that genuinely require new code, implement them in priority waves:

1. **Wave 1 — Low-hanging fruit** (checkmarks via minor code fixes): Sections 5, 7, 8, 10, 11, 16, 17, 18, 19 — features that exist but need small wiring, a flag, or a format tweak.
2. **Wave 2 — Discovery & evidence hardening**: Section 5 (diff detection, rename handling), Section 6 (language-agnostic model verification), Section 7 (evidence normalization, missing evidence handling).
3. **Wave 3 — GitHub integration**: Sections 13, 14, 15 — PR workflow behavior, scheduled workflows, issue sync. Code exists in `internal/github/` but needs CI workflow wiring.
4. **Wave 4 — Documentation & operational quality**: Sections 1, 21, 22, 23, 24, 25 — docs, security model, schema validation, rollout guidance.
5. **Wave 5 — Graduation criteria**: Sections 26, 27 — cross-cutting readiness validation.

## Relevant Files
Use these files to resolve the chore:

**Acceptance criteria source:**
- `FEATURES.md` — The 275-criterion checklist. Every `- [ ]` → `- [x]` change must be justified by existing code or new implementation.

**Domain model (verify against §1, §8, §9, §10):**
- `internal/domain/unit.go` — UnitID with language/path/symbol, UnitType enum, Unit struct
- `internal/domain/record.go` — CertificationRecord with all fields (status, grade, score, dimensions, evidence, observations, actions, timestamps, source)
- `internal/domain/evidence.go` — Evidence struct, EvidenceKind enum, Severity enum
- `internal/domain/dimension.go` — 9 dimensions, DimensionScores, Grade, WeightedAverage
- `internal/domain/config.go` — Config, CertificationMode (advisory/enforcing), ExpiryConfig, AgentConfig
- `internal/domain/policy.go` — PolicyPack, PolicyRule with language/path targeting
- `internal/domain/expiry.go` — ExpiryWindow, ExpiryFactors
- `internal/domain/override.go` — Override with action types and rationale validation

**Configuration & policy (verify against §3, §4):**
- `internal/config/loader.go` — LoadFile, Load, LoadFromDir for YAML config
- `internal/config/policy.go` — LoadPolicyPack, LoadPolicyPacks for policy YAML
- `internal/config/matcher.go` — NewPolicyMatcher matching policies to units by language/path

**Discovery (verify against §5, §6, §20):**
- `internal/discovery/scanner.go` — GenericScanner with include/exclude globs, Merge, DeduplicateFileLevel
- `internal/discovery/go_adapter.go` — GoAdapter using go/ast for functions, methods, types
- `internal/discovery/ts_adapter.go` — TSAdapter using regex for exports, classes, consts
- `internal/discovery/detect.go` — DetectLanguages, DetectedAdapters for auto-detection
- `internal/discovery/index.go` — Index with JSON save/load, Diff (added/removed/unchanged)
- `internal/discovery/generic.go` — Generic file-level scanner

**Evidence collection (verify against §7):**
- `internal/evidence/executor.go` — ToolExecutor running go vet, go test, golangci-lint, git
- `internal/evidence/runner.go` — ParseGoVet, ParseGoTestJSON, ParseCoverProfile, ParseGolangciLintJSON, ParseGitLogWithAge
- `internal/evidence/metrics.go` — ComputeMetrics (lines, comments, TODOs), ToEvidence
- `internal/evidence/complexity.go` — ComputeGoComplexity (AST cyclomatic), ComputeSymbolMetrics
- `internal/evidence/git.go` — ParseGitLog, GitStats, ChurnRate
- `internal/evidence/lint.go` — LintResult, LintFinding, TestResult with ToEvidence
- `internal/evidence/collector.go` — Evidence collector interface

**Evaluation engine (verify against §8):**
- `internal/policy/evaluator.go` — Evaluate rules against evidence, extractMetric, extractTodoCount, extractComplexity
- `internal/policy/matcher.go` — Match policies to units
- `internal/engine/scorer.go` — Score dimensions from evidence, complexity/size/git scoring, StatusFromScore
- `internal/engine/pipeline.go` — CertifyUnit full pipeline (evaluate → score → status → expiry → record)

**Expiry (verify against §11):**
- `internal/expiry/calculator.go` — Calculate expiry windows with risk-based adjustment

**Override (verify against §17):**
- `internal/override/loader.go` — LoadDir for override YAML
- `internal/override/applier.go` — ApplyAll applying overrides to records

**Agent review (verify against §16):**
- `internal/agent/stage.go` — PrescreenStage, ReviewStage, ScoringStage with loose JSON parsing
- `internal/agent/pipeline.go` — Pipeline (strategy-based stage composition), Coordinator (dedup, budget)
- `internal/agent/circuit.go` — CircuitBreaker wrapping Provider
- `internal/agent/fallback.go` — FallbackProvider, ModelChain, AdaptiveMessages
- `internal/agent/reviewer.go` — ReviewResult with ModelsUsed attribution, ToEvidence
- `internal/agent/openrouter.go` — OpenRouterProvider with retry/backoff
- `internal/agent/prompts.go` — PromptTemplate, PromptRegistry

**Record store (verify against §10, §22):**
- `internal/record/store.go` — Store with SHA256-hashed filenames, Save/Load/LoadAll

**Reporting (verify against §18):**
- `internal/report/health.go` — HealthReport with FormatText, FormatJSON

**GitHub integration (verify against §2, §13, §14, §15):**
- `internal/github/workflows.go` — GeneratePRWorkflow, GenerateNightlyWorkflow, GenerateWeeklyWorkflow
- `internal/github/pr.go` — FormatPRComment (advisory + enforcing), BuildPRCommentCommand
- `internal/github/issues.go` — FormatIssueTitle/Body, BuildIssueCreateCommand/CloseCommand

**Queue (verify against §14 incremental):**
- `internal/queue/queue.go` — Persistent work queue with save-after-each-item

**CLI (verify against §19):**
- `cmd/certify/init_cmd.go` — init with language detection, policy/workflow generation
- `cmd/certify/scan.go` — scan with smart adapter selection
- `cmd/certify/certify_cmd.go` — certify with queue, agent coordinator, --batch/--skip-agent/--path
- `cmd/certify/report_cmd.go` — report with --format text/json
- `cmd/certify/expire.go` — expire overdue records
- `cmd/certify/review.go` — PR review annotation
- `cmd/certify/version.go` — version info

**CI:**
- `.github/workflows/ci.yml` — test + certify jobs on push/PR

### New Files

- `docs/README.md` — Installation, quickstart, CLI reference (§24)
- `docs/policy-authoring.md` — Policy pack authoring guide (§24)
- `docs/architecture.md` — Architecture overview with diagrams (§1, §24)
- `docs/troubleshooting.md` — Common issues and fixes (§24)
- `internal/config/validator.go` — Config and policy schema validation (§22)
- `internal/config/validator_test.go` — Tests for validation
- `internal/discovery/diff.go` — Richer diff: moved/renamed detection (§5)
- `internal/discovery/diff_test.go` — Tests for diff
- `internal/report/dimension.go` — Dimension-level breakdown reports (§9, §18)
- `internal/report/dimension_test.go` — Tests for dimension reports
- `.github/workflows/nightly.yml` — Nightly sweep workflow (§14)
- `.github/workflows/weekly.yml` — Weekly report workflow (§14)

## Step by Step Tasks
IMPORTANT: Execute every step in order, top to bottom.

### Step 1 — Audit Part A: Check Off Already-Implemented Criteria

Systematically walk FEATURES.md sections 1–27. For each unchecked `- [ ]`, verify the criterion is already satisfied by existing code. If yes, change to `- [x]`. Expected to close ~120 criteria with zero code changes.

- §5 Certifiable Unit Discovery: UnitID has language/path/symbol (check ✓), generic fallback exists (check ✓), include/exclude respected in GenericScanner (check ✓), index is persistent JSON (check ✓), Diff computes added/removed/unchanged (check ✓)
- §6 Language-Agnostic: system certifies Go + TS + generic files (check ✓), adapters are pluggable (check ✓), core engine has no language hardcoding (check ✓)
- §7 Evidence Collection: executor runs lint/test/git tools (check ✓), results normalized to Evidence struct (check ✓), evidence attached to records (check ✓)
- §8 Evaluation Engine: evaluates units against policy packs (check ✓), records violations (check ✓), computes dimension scores (check ✓), weighted scoring (check ✓), assigns status/grade/score/confidence (check ✓)
- §10 Trust Ledger: records have unit_id, type, path, policy_version, status, grade, score, evidence, observations, actions, certified_at, expires_at (check all ✓)
- §11 Expiry: every certified unit gets expires_at (check ✓), windows computed from config (check ✓), min/max configurable (check ✓)
- §16 Agent Review: optional (checked), can enable/disable via config (check ✓), output distinguished in Source field (check ✓), cannot override deterministic failures (check ✓ — agent evidence is additive), operates without agent (check ✓)
- §17 Overrides: supports exemptions (check ✓), manual overrides (check ✓), stored in YAML (check ✓), rationale required (check ✓), can extend/shorten/force (check ✓)
- §18 Reporting: machine-readable JSON (check ✓), human-readable text (check ✓), counts by status (check ✓), stored in predictable locations (check ✓)
- §19 CLI: provides local CLI (check ✓), init/scan/certify/expire/report/review commands (check all ✓), respects config (check ✓)
- §20 Multi-Language: adapter boundary exists (check ✓), Go + TS adapters (check ✓), adding adapters doesn't change core (check ✓)
- §22 Storage: structured schemas exist (checked), data files reviewable (check ✓), tolerates missing state (check ✓)
- §25 Architecture: core domain free of host concerns (check ✓), lint-compliant via go vet (check ✓), testable and modular (check ✓ — 256 tests, 15 packages)

### Step 2 — Wave 1: Low-Hanging Code Fixes

Small code additions that close criteria in bulk:

- **§7 Missing evidence marking**: In `internal/policy/evaluator.go`, when `extractMetric` returns -1, the violation already says "missing evidence for metric X". Verify this covers §7's "missing evidence explicitly marked" criterion. Add a test.
- **§8 Human-readable explanation**: `CertificationRecord.Observations` already carries violation descriptions. Add a `func (r CertificationRecord) Explanation() string` method that renders a readable summary. Test it.
- **§9 Dimension-level breakdown in reports**: Create `internal/report/dimension.go` — extend HealthReport to include per-dimension averages. Update FormatText and FormatJSON to include dimension breakdowns. Test.
- **§10 Policy version in records**: `CertifyUnit` in `internal/engine/pipeline.go` sets `PolicyVersion: ""`. Wire the matched policy pack versions into this field. Test.
- **§11 Expired units marked in reports**: Update `internal/report/health.go` FormatText to list expired units. Already counts them; add names.
- **§17 Overrides visible in reports**: When record has override-applied status, include in report output.
- **§18 Expiring-soon / highest-risk / failing areas in reports**: Extend HealthReport to compute and include these lists. Test.

### Step 3 — Wave 2: Discovery & Evidence Hardening

- **§5 Moved/renamed file detection**: Create `internal/discovery/diff.go` with `DetectMoves(old, new []Unit) []MoveResult` using path similarity + symbol matching. Test with fixtures.
- **§5 Changed unit detection between revisions**: Extend `Diff` to accept git commit SHAs and compute changed files via `git diff --name-only`. Test.
- **§6 Polyglot verification**: Add integration test that certifies a repo containing Go + TS + Python files. Verify all three get discovered and certified (Go/TS with adapters, Python with generic).
- **§7 Evidence collection failure reporting**: In `ToolExecutor.CollectAll`, return `[]EvidenceResult` that includes errors alongside successes. Surface in report. Test.
- **§7 Partial evidence not represented as complete**: Add `Completeness` field to Evidence (0.0–1.0). When go test runs but coverage is missing, set Completeness < 1.0. Test.

### Step 4 — Wave 3: GitHub Integration Wiring

- **§2 Init via workflow dispatch**: Add `workflow_dispatch` trigger to `.github/workflows/ci.yml` with `init` job that runs `certify init` and creates a PR via `gh pr create`.
- **§2 Init PR with summary**: Update `cmd/certify/init_cmd.go` to output a markdown summary of what was detected and generated. Capture and use in PR body.
- **§13 PR workflow**: Create `.github/workflows/pr-certify.yml` that runs scan → certify on changed files only (use `git diff --name-only origin/main...HEAD`). Post comment via `certify review`.
- **§13 Trust delta**: Add `--diff-base` flag to `certify certify` that computes and reports trust delta vs a base commit.
- **§14 Scheduled workflows**: Create `.github/workflows/nightly.yml` and `.github/workflows/weekly.yml` from the generators, customized for this repo.
- **§14 Incremental scheduled runs**: Use the queue with `--batch` in nightly workflow.
- **§15 Issue sync**: Add `certify issues` command that reads records, finds failing/expired units, and calls `gh issue create`/`gh issue close` via the existing `internal/github/issues.go` builders. Dedup via label search.

### Step 5 — Wave 4: Config Validation & Documentation

- **§22 Config/policy validation**: Create `internal/config/validator.go` with `ValidateConfig(cfg Config) []ValidationError` and `ValidatePolicyPack(p PolicyPack) []ValidationError`. Check required fields, valid severity values, valid metric names. Test. Wire into `certify certify` as a pre-flight check.
- **§1 Documented architecture**: Create `docs/architecture.md` with component diagram, data flow, package dependency description.
- **§21 Security model**: Create `docs/security.md` documenting workflow permissions, fork PR limitations, token scoping. Audit `.github/workflows/*.yml` for minimal permissions.
- **§23 Rollout guidance**: Create `docs/rollout.md` with incremental adoption playbook (advisory → scoped → enforcing).
- **§24 Operational quality docs**: Create `docs/README.md` (install + quickstart), `docs/policy-authoring.md`, `docs/troubleshooting.md`, `docs/cli-reference.md`.

### Step 6 — Wave 5: Graduation Criteria (§26, §27)

- Walk every criterion in §26 (Minimum v1 Readiness) and §27 (Program-Level Success Indicators).
- For each, either verify it passes by running the tool against the test repo, or implement the missing piece.
- Create an integration test `TestV1ReadinessCriteria` that exercises the full lifecycle: init → scan → certify → expire → report → review → issues.
- Check off all 19 criteria in §26 and §27.

### Step 7 — Final FEATURES.md Audit & Validation

- Run `grep -c '\- \[ \]' FEATURES.md` — should be 0.
- Run `grep -c '\- \[x\]' FEATURES.md` — should be 275.
- Run all validation commands below.
- Commit with message `chore: close all 275 FEATURES.md criteria — v1 complete. Refs #5`

## Validation Commands
Execute every command to validate the chore is complete with zero regressions.

```bash
# All tests pass (expect 270+)
go test ./... -count=1

# Zero vet issues
go vet ./...

# Format compliance
gofmt -l $(find . -name '*.go' -not -path './testdata/*') | wc -l  # expect 0

# Build succeeds
go build -o build/bin/certify ./cmd/certify/

# CLI commands functional
./build/bin/certify version
./build/bin/certify init --path /tmp/certify-validate-test
./build/bin/certify scan --path /tmp/certify-validate-test
./build/bin/certify certify --skip-agent --path /tmp/certify-validate-test
./build/bin/certify report --path /tmp/certify-validate-test
./build/bin/certify report --format json --path /tmp/certify-validate-test
./build/bin/certify expire --path /tmp/certify-validate-test

# Self-certification passes
./build/bin/certify scan
./build/bin/certify certify --skip-agent --reset-queue
./build/bin/certify report

# FEATURES.md fully checked
test $(grep -c '\- \[ \]' FEATURES.md) -eq 0
test $(grep -c '\- \[x\]' FEATURES.md) -eq 275

# Documentation exists
test -f docs/README.md
test -f docs/architecture.md
test -f docs/policy-authoring.md
test -f docs/troubleshooting.md

# Workflows exist
test -f .github/workflows/ci.yml
test -f .github/workflows/nightly.yml
test -f .github/workflows/weekly.yml

# Cleanup
rm -rf /tmp/certify-validate-test
```

## Notes
- **Part A is the biggest win**: Roughly 120 criteria are already implemented but never checked off. The audit alone should bring us from 17% → ~60% with zero code changes.
- **TDD applies to all new code**: Every new function/method in Waves 1–5 gets a test written first.
- **Commit after each wave**: Each wave should be a separate commit with `Refs #5` and updated FEATURES.md counts.
- **Agent review criteria (§16)**: Several depend on API availability. Test with mocks; verify live when free-tier limits allow.
- **§12 Invalidation**: Partial implementation exists (Diff detects changes). Full invalidation logic (auto-recertify on change) is best wired through the PR workflow (§13).
- **§27 Program-Level Success**: These are high-level "can a team do X" criteria. Verified by the integration test + docs, not individual unit tests.
- **Estimated effort by wave**: Part A (2h audit), Wave 1 (3h), Wave 2 (4h), Wave 3 (6h), Wave 4 (4h), Wave 5 (2h) = ~21h total.
