# Chore: High-Value Observation Reduction

## Chore Description

The self-certification report card shows 773 units with 29 C-grade units, 85.8% overall score, and 175 active observations across the codebase. Analysis reveals four high-value wins that will significantly improve the score and promote C-grade units to B with targeted, low-risk changes:

1. **Stale record pruning** — 25 records exist for units no longer in the index (deleted `init()` functions, renamed types). These phantom records pollute the report card, architect review, and observation counts. Adding a `Store.Prune()` method and calling it during `certify` and `report` eliminates ghost data.

2. **`todo_count` false positives** — The `containsTodo` function in `metrics.go` uses naive substring matching (`strings.Contains(upper, "TODO")`), causing false positives on: test data struct literals (`TodoCount: 0`), policy YAML strings (`"No TODO/FIXME comments"`), and the TODO-detection code itself (`containsTodo`, `extractTodoCount`). Fixing the detection to only match actual TODO/FIXME comment markers eliminates ~10 of 13 `todo_count` observations.

3. **Split long functions** — 8 functions exceed the 100-line threshold, causing `func_lines` observations that push units to C grade. The top offenders: `FormatForLLM` (212 lines), `Certify` (144 lines), `runArchitect` (125 lines), `generateUnitPages` (123 lines), `formatUnitMarkdownWithNav` (122 lines), `reviewer.Review` (121 lines), `scoreFromStructural` (105 lines), `architect_review.Review` (105 lines).

4. **Fix remaining `errors_ignored`** — 28 active observations from blank identifiers in multi-value assignments. Many are legitimate probe patterns (`http.Get` for URL checking) but can be restructured to handle errors properly.

**Expected impact:** 22 C-grade units → <10, observations 175 → ~80, score 85.8% → ~87%.

## Relevant Files

### Stale Record Pruning
- `internal/record/store.go` — Record store; needs new `Prune(validIDs set)` method to delete records not in current index
- `internal/record/store_test.go` — Tests for store; needs `TestPrune` 
- `cmd/certify/certify_cmd.go` — Certification command; should call `Prune()` after scan to clean stale records before certification
- `cmd/certify/report_cmd.go` — Report command; should call `Prune()` before loading records for report generation

### TODO False Positive Fix
- `internal/evidence/metrics.go` — Contains `containsTodo()` with naive substring matching; needs to exclude string literals and only match comment-context TODOs
- `internal/evidence/metrics_test.go` — Tests for metrics; needs test cases for false positive exclusion
- `internal/evidence/metrics.go#ComputeMetrics` — The `ComputeMetrics` function walks lines and calls `containsTodo` — already distinguishes comment vs code lines but the issue is that `containsTodo` is called on code lines that contain struct field names like `TodoCount`

### Split Long Functions
- `internal/agent/architect.go` — `FormatForLLM` at 212 lines; extract helpers for each section (metrics, packages, hotspots, coupling, dependency graph)
- `internal/engine/certifier.go` — `Certify` at 144 lines; extract evidence collection, policy evaluation, and record creation into helpers
- `cmd/certify/architect_cmd.go` — `runArchitect` at 125 lines; extract provider setup, context building, and output formatting
- `internal/report/report_tree.go` — `formatUnitMarkdownWithNav` at 122 lines; extract dimension table, observation list, evidence sections into helpers
- `internal/report/site.go` — `generateUnitPages` at 123 lines; extract template rendering sections
- `internal/agent/reviewer.go` — `Review` at 121 lines; extract prompt building and response parsing
- `internal/engine/scorer.go` — `scoreFromStructural` at 105 lines; extract metric extraction into helper
- `internal/agent/architect_review.go` — `Review` at 105 lines; extract phase execution loop

### Fix errors_ignored
- `internal/evidence/executor.go` — 5 occurrences: `runGolangciLint` (2), `HasPackageJSON` (1), `HasGoMod` (1), `runGoVet` (1), `runGoTest` (1), `runGitStats` (1)
- `internal/record/store.go` — 4 occurrences: `fromJSON` (5 — already addressed?), `AppendRun` (1), `SaveSnapshot` (1), `AppendHistory` (2)
- `internal/policy/evaluator.go` — 3 occurrences: `extractCoverage` (1), `extractComplexity` (2), `extractTodoCount` (1)
- `internal/policy/matcher.go` — 3 occurrences: `matchPath` (3)
- `internal/discovery/index.go` — 3 occurrences: `LoadIndex` (1), `Diff` (2)
- `internal/discovery/generic.go` — 1 occurrence: `matchAny` (1)
- `internal/engine/scorer.go` — 4 occurrences: `extractSummaryInt` (3), `extractSummaryFloat` (1)
- `internal/engine/certifier.go` — 1 occurrence: `SaveReportArtifacts` (1)
- `internal/report/full.go` — 1 occurrence: `writeAllUnits` (1)
- `internal/report/report_tree.go` — 1 occurrence: `GenerateReportTree` (1)
- `internal/workspace/workspace.go` — 3 occurrences: `LoadSubmoduleCard` (2), `CheckHasConfig` (1)
- `internal/agent/models.go` — 2 occurrences: `listOllamaModels` (1), `listOpenAIModels` (1)
- `cmd/certify/certify_cmd.go` — 1 occurrence: `detectAPIKeyOnly` (1)
- `internal/queue/queue.go` — 1 occurrence: `Enqueue` (1)

### New Files
- None — all changes are to existing files.

## Step by Step Tasks

### Step 1: Add `Store.Prune()` method (TDD)

- Write `TestPrune` in `internal/record/store_test.go`:
  - Create a store with 3 records (A, B, C)
  - Call `Prune(map[string]bool{"A": true, "C": true})`
  - Assert record B is deleted, A and C remain
  - Assert returned count is 1 (pruned)
- Run test — expect fail
- Implement `func (s *Store) Prune(validIDs map[string]bool) (int, error)` in `internal/record/store.go`:
  - Read all `.json` files in `s.dir`
  - For each, unmarshal to get `unit_id`
  - If `unit_id` not in `validIDs`, delete the file
  - Return count of deleted files
- Run test — expect pass

### Step 2: Wire pruning into certify and report commands

- In `cmd/certify/certify_cmd.go`, after loading the index (`loadIndex`), call `store.Prune(indexIDs)` and log the count if verbose
- In `cmd/certify/report_cmd.go`, after creating the store and before `store.ListAll()`, load the index and call `store.Prune(indexIDs)` to clean stale records
- Build and run `just build` to verify compilation

### Step 3: Fix `containsTodo` false positives (TDD)

- In `internal/evidence/metrics.go`, modify `ComputeMetrics` to only count TODOs in comment lines:
  - The function already distinguishes comment lines (starting with `//` or within `/* */` blocks) from code lines
  - Currently `containsTodo` is called on BOTH comment lines AND code lines (line 43 is inside the `inBlockComment || strings.HasPrefix(trimmed, "//")` branch, but line 49 calls it on code lines too)
  - Fix: remove the `containsTodo` call from the code-line branch (line 49) — only count TODOs in actual comments
- Update `internal/evidence/metrics_test.go`:
  - Add test case: code line with `TodoCount: 0` should NOT count as a TODO
  - Add test case: code line with `"No TODO/FIXME comments"` in a string literal should NOT count
  - Add test case: `// TODO: fix this` should still count
- Run tests — expect pass

### Step 4: Split `FormatForLLM` (212 → <100 per function)

- In `internal/agent/architect.go`, extract from `FormatForLLM`:
  - `formatSnapshotMetrics(b *strings.Builder, snap *ArchitectSnapshot)` — aggregate metrics section
  - `formatPackageTable(b *strings.Builder, pkgs []PackageInfo)` — package table
  - `formatHotspots(b *strings.Builder, hotspots []Hotspot)` — hotspots section
  - `formatCoupling(b *strings.Builder, edges []CouplingEdge)` — coupling analysis
  - `formatDependencyGraph(b *strings.Builder, deps map[string][]string)` — dependency graph
  - `formatGradeDistribution(b *strings.Builder, dist map[string]int)` — grade distribution
- `FormatForLLM` becomes an orchestrator calling these helpers
- Run `go build ./...` to verify

### Step 5: Split `Certify` (144 → <100 per function)

- In `internal/engine/certifier.go`, extract from `Certify`:
  - `collectAndMergeEvidence(unit domain.Unit, repoEvidence []domain.Evidence) []domain.Evidence` — evidence collection + merge logic
  - `evaluateAndScore(unit domain.Unit, evidence []domain.Evidence, policies []domain.PolicyRule) (domain.Status, domain.Grade, float64, domain.DimensionScores, []string, []domain.PolicyResult)` — policy eval + scoring
- `Certify` becomes an orchestrator calling these helpers
- Run `go test ./internal/engine/...` to verify

### Step 6: Split `runArchitect` (125 → <100)

- In `cmd/certify/architect_cmd.go`, extract:
  - `buildArchitectContext(certDir string, cfg domain.Config) (*agent.ProjectContext, error)` — snapshot + context building
  - `writeArchitectOutput(review *agent.ArchitectReview, certDir string, verbose bool) error` — output formatting + file writing
- Run `go build ./cmd/certify/` to verify

### Step 7: Split `formatUnitMarkdownWithNav` (122 → <100)

- In `internal/report/report_tree.go`, extract:
  - `formatDimensionsTable(b *strings.Builder, dims domain.DimensionScores)` — dimension score table
  - `formatEvidenceSection(b *strings.Builder, evidence []domain.Evidence)` — evidence rendering
- Run `go test ./internal/report/...` to verify

### Step 8: Split `generateUnitPages` (123 → <100)

- In `internal/report/site.go`, extract:
  - `renderUnitPageContent(rec domain.CertificationRecord) string` — content generation for a single unit page
- Run `go build ./...` to verify

### Step 9: Split `reviewer.Review` (121 → <100)

- In `internal/agent/reviewer.go`, extract:
  - `buildReviewPrompt(unit domain.Unit, evidence []domain.Evidence) string` — prompt construction
  - `parseReviewResponse(raw string) (ReviewResult, error)` — response parsing + JSON extraction
- Run `go test ./internal/agent/...` to verify

### Step 10: Split `scoreFromStructural` (105 → <100)

- In `internal/engine/scorer.go`, extract:
  - `applyStructuralPenalty(dim string, metric string, value float64, threshold float64, scores map[string]float64)` — common penalty application logic
- Run `go test ./internal/engine/...` to verify

### Step 11: Split `architect_review.Review` (105 → <100)

- In `internal/agent/architect_review.go`, extract:
  - `executePhase(ctx context.Context, provider agent.Provider, phase PhaseConfig, priorOutputs []string) (string, error)` — single phase execution with think-tag stripping
- Run `go test ./internal/agent/...` to verify

### Step 12: Fix `errors_ignored` — executor.go

- In `internal/evidence/executor.go`:
  - `HasPackageJSON`: change `_, err := os.Stat(...)` with `_ =` to proper `if err != nil` check
  - `HasGoMod`: same pattern
  - `runGolangciLint`: handle `cmd.CombinedOutput()` error (already done? check for `_ =` patterns)
  - `runGoVet`, `runGoTest`, `runGitStats`: same — ensure no `_ =` on multi-return calls
- Run `go test ./internal/evidence/...`

### Step 13: Fix `errors_ignored` — remaining files

- `internal/policy/evaluator.go`: `extractCoverage`, `extractComplexity`, `extractTodoCount` — replace `fmt.Sscanf` `_ =` with proper error handling
- `internal/policy/matcher.go`: `matchPath` — replace `filepath.Match` `_ =` with error check
- `internal/discovery/index.go`: `LoadIndex`, `Diff` — handle JSON decode errors
- `internal/discovery/generic.go`: `matchAny` — handle `filepath.Match` error
- `internal/engine/scorer.go`: `extractSummaryInt`, `extractSummaryFloat` — handle `fmt.Sscanf` errors
- `internal/engine/certifier.go`: `SaveReportArtifacts` — handle errors
- `internal/report/full.go`: `writeAllUnits` — handle errors
- `internal/report/report_tree.go`: `GenerateReportTree` — handle errors
- `internal/record/store.go`: `AppendRun`, `SaveSnapshot`, `AppendHistory` — handle marshal/write errors
- `internal/workspace/workspace.go`: `LoadSubmoduleCard`, `CheckHasConfig` — handle errors
- `internal/agent/models.go`: `listOllamaModels`, `listOpenAIModels` — handle HTTP/JSON errors
- `internal/queue/queue.go`: `Enqueue` — handle errors
- `cmd/certify/certify_cmd.go`: `detectAPIKeyOnly` — handle errors
- Run `go test ./...`

### Step 14: Validate — full build, test, certify

- `gofmt -l ./...` — must produce no output
- `go vet ./...` — must pass
- `go test ./... -count=1` — all packages must pass
- `go build -o build/bin/certify ./cmd/certify/` — must succeed
- `./build/bin/certify scan --root .` — rescan to update index
- `./build/bin/certify certify --root . --skip-agent --batch --reset-queue` — full re-certification
- `./build/bin/certify report` — generate report card
- Verify: C-grade units < 10, observations significantly reduced, score ≥ 87%

## Validation Commands

```bash
# Format check
gofmt -l .

# Vet check  
go vet ./...

# All tests pass
go test ./... -count=1

# Build succeeds
go build -o build/bin/certify ./cmd/certify/

# Scan
./build/bin/certify scan --root .

# Full re-certification (with pruning now built-in)
./build/bin/certify certify --root . --skip-agent --batch --reset-queue

# Generate report
./build/bin/certify report

# Verify targets met
python3 -c "
import json, glob
with open('.certification/index.json') as f:
    index = json.load(f)
index_ids = set(u['id'] for u in index)
c_count = 0
obs_count = 0
total_score = 0
total = 0
for f in glob.glob('.certification/records/*.json'):
    with open(f) as fh:
        r = json.load(fh)
    if r['unit_id'] not in index_ids:
        print(f'ERROR: stale record found: {r[\"unit_id\"]}')
        continue
    total += 1
    total_score += r.get('score', 0)
    if r.get('grade') == 'C':
        c_count += 1
        print(f'C: {r[\"unit_id\"]} ({r.get(\"score\",0):.3f})')
    obs_count += len(r.get('observations', []))
avg = total_score / total * 100 if total else 0
print(f'\nTotal: {total}, C-grade: {c_count}, Observations: {obs_count}, Score: {avg:.1f}%')
assert c_count < 10, f'C-grade units {c_count} >= 10'
assert avg >= 87.0, f'Score {avg:.1f}% < 87%'
print('ALL TARGETS MET')
"
```

## Notes

- **`global_mutable_count` is deferred**: Go doesn't support `const` maps. The 24 observations from `var` map declarations in `domain`, `discovery`, and `agent` would require wrapping every map access in a function call — high churn, low value. Accept these as structural Go limitations.
- **AI-generated observations (99 total) are not addressed here**: These come from LLM review during certification and are subjective. They account for most of the 175 total but don't directly cause C grades. Addressing them would require code changes based on AI suggestions (e.g., `isLocalURL` security), which is a separate effort.
- **`os_exit_calls` in `main.go` is standard Go**: The `os.Exit(1)` in `main()` after `rootCmd.Execute()` fails is idiomatic Go CLI. Removing it would change program behavior. Accept this observation.
- **Test file TODOs**: Test files like `metrics_test.go` contain string literals with "TODO" for testing the TODO detection feature. The fix in Step 3 (only counting comment-line TODOs) handles this correctly since these appear in code lines, not comment lines.
- **Stale records come from refactoring**: When functions are renamed, deleted, or consolidated (like our init() elimination), old records persist. The prune mechanism ensures reports always reflect current code.
