# Chore: Architect Review Remediation — Highest Impact Wins

## Chore Description

The architect review (`.certification/ARCHITECT_REVIEW.md`) identified 7 recommendations across the codebase. The project currently scores **B (85.4%)** with **48 observations** across 615 units. This chore remediates the highest-impact issues first, targeting the root causes that generate the most observations.

**Key insight from analysis:** The majority of observations are **file-level metrics bleeding into every function**. The `has_init_func` and `global_mutable_count` metrics are computed per-file and then applied to every function within that file. A single `init()` function in `certify_cmd.go` generates 18 `has_init_func` observations (one per function), and 7 global vars generate 18 `global_mutable_count` observations. Fixing file-level issues has multiplicative impact.

**Observation breakdown (129 records with observations, 48 unique units):**
- `has_init_func: 1 exceeds threshold 0` — **70 occurrences** (10 init() in cmd/certify, 4 in domain, 2 in agent — affects every function in those files)
- `errors_ignored: N exceeds threshold 0` — **25 occurrences** (actual code quality issue: `_ =` assignments)
- `global_mutable_count: N exceeds threshold 2` — **52 occurrences** (certify_cmd.go has 7, report_cmd.go has 6, providers.go has 6+)
- `todo_count: N exceeds threshold 0` — **~10 occurrences** (real TODOs scattered in source)
- `param_count: N exceeds threshold 5` — **~3 occurrences** (buildCertificationRun has 8 params)
- AI-generated observations — **~8 occurrences** (isLocalURL security concerns)

**Remediation strategy (ordered by observation reduction):**

| # | Fix | Observations Eliminated | Score Impact |
|---|-----|------------------------|-------------|
| 1 | Consolidate `cmd/certify` init() into single registration | ~36 (39 units × init) | cmd/certify 78.4% → ~82% |
| 2 | Reduce `cmd/certify` global vars (flags → local scope) | ~36 (39 units × globals) | cmd/certify → ~85% |
| 3 | Replace `internal/domain` init() with compile-time maps | ~28 (72 units × 4 inits) | domain 83.1% → ~86% |
| 4 | Replace `internal/agent` init() with map literal merges | ~20 (138 units, but only ~10 affected) | agent 86.0% → ~87% |
| 5 | Fix `errors_ignored` across packages | ~25 direct | scattered +1-2% |
| 6 | Fix remaining TODOs and param_count | ~13 | scattered +1% |

## Relevant Files

Use these files to resolve the chore:

**cmd/certify (init consolidation + global var reduction):**
- `cmd/certify/root.go` — defines `rootCmd`, `workspaceMode` global var, `init()` registers subcommands + persistent flags
- `cmd/certify/certify_cmd.go` — 6 flag globals (`certifyPath`, `certifySkipAgent`, `certifyBatch`, `certifyResetQueue`, `certifyTarget`, `certifyDiffBase`), `certifyCmd` global, `init()` binds flags
- `cmd/certify/scan.go` — `scanPath` global, `scanCmd` global, `init()` binds flags
- `cmd/certify/report_cmd.go` — 6 flag globals, `reportCmd` global, `init()` binds flags
- `cmd/certify/architect_cmd.go` — 5 flag globals, `architectCmd` global, `init()` binds flags
- `cmd/certify/init_cmd.go` — 3 flag globals, `initCmd` global, `init()` binds flags
- `cmd/certify/expire.go` — `expirePath` global, `expireCmd` global, `init()` binds flags
- `cmd/certify/models_cmd.go` — 2 flag globals, `modelsCmd` global, `init()` binds flags
- `cmd/certify/review.go` — `reviewPath` global, `reviewCmd` global, `init()` binds flags
- `cmd/certify/version.go` — 3 version globals (ldflags), `versionCmd` global, `init()` registers
- `cmd/certify/main.go` — calls `rootCmd.Execute()`, entry point
- `cmd/certify/workspace_dispatch.go` — helper function, no globals, no init

**internal/domain (init elimination):**
- `internal/domain/record.go` — `stringToStatus` var built in `init()` by inverting `statusStrings`
- `internal/domain/unit.go` — `stringToUnitType` var built in `init()` by inverting `unitTypeStrings`
- `internal/domain/evidence.go` — `stringToEvidenceKind` and `stringToSeverity` vars built in 2 `init()` functions

**internal/agent (init elimination):**
- `internal/agent/autodetect.go` — `init()` sets `DefaultModels["openrouter"] = ConservativeModels`
- `internal/agent/providers.go` — `init()` populates 4 entries in `DefaultModels` map

**errors_ignored fixes:**
- `cmd/certify/certify_cmd.go` — `root, _ = os.Getwd()` (lines 135, 378), `apiKey, _ = agent.DetectAPIKey()` (line 447)
- `cmd/certify/architect_cmd.go` — `root, _ = os.Getwd()` (line 60), `apiKey, _ = agent.DetectAPIKey()` (line 194)
- `cmd/certify/expire.go` — `_ =` pattern
- `cmd/certify/init_cmd.go` — `_ =` pattern
- `cmd/certify/report_cmd.go` — `_ =` pattern
- `cmd/certify/review.go` — `_ =` pattern
- `cmd/certify/scan.go` — `_ =` pattern
- `internal/evidence/structural.go` — 2 `_ =` patterns
- `internal/policy/evaluator.go` — `_, _ = fmt.Sscanf(...)` (line 152)
- `internal/record/store.go` — `_ = json.Unmarshal(...)` (line 456)
- `internal/report/architect_report.go` — `_ = grade` (line 372)

**TODO fixes:**
- `internal/engine/scorer.go` — line 111: `// TODO count already handled by violation penalties`

**param_count fix:**
- `cmd/certify/certify_cmd.go` — `buildCertificationRun` has 8 parameters, threshold is 5

**Test files (verify no regressions):**
- `cmd/certify/cli_test.go`
- `internal/domain/evidence_test.go`
- `internal/agent/architect_test.go`
- `internal/evidence/structural_test.go`
- `internal/engine/scorer_test.go`

### New Files
- None. All changes are to existing files.

## Step by Step Tasks

### Step 1: Eliminate init() in internal/domain — replace with compile-time reverse maps

The 4 `init()` functions in `internal/domain` build reverse lookup maps (`stringToStatus`, `stringToUnitType`, `stringToEvidenceKind`, `stringToSeverity`) by iterating the forward maps. Replace with explicit compile-time `map[string]T` literals. This eliminates `has_init_func` observations for all 72 units in this package.

- In `internal/domain/record.go`:
  - Remove `var stringToStatus map[string]Status`
  - Remove the `init()` function
  - Replace with `var stringToStatus = map[string]Status{ "certified": StatusCertified, ... }` using all values from `statusStrings`
  - Verify `StatusFromString()` still works

- In `internal/domain/unit.go`:
  - Remove `var stringToUnitType map[string]UnitType`
  - Remove the `init()` function
  - Replace with `var stringToUnitType = map[string]UnitType{ "function": UnitTypeFunction, ... }`

- In `internal/domain/evidence.go`:
  - Remove `var stringToEvidenceKind map[string]EvidenceKind` and its `init()`
  - Replace with `var stringToEvidenceKind = map[string]EvidenceKind{ "lint": EvidenceKindLint, ... }`
  - Remove `var stringToSeverity map[string]Severity` and its `init()`
  - Replace with `var stringToSeverity = map[string]Severity{ "info": SeverityInfo, ... }`

- Run `go test ./internal/domain/... -count=1` to verify.

### Step 2: Eliminate init() in internal/agent — merge into map literals

The 2 `init()` functions in `internal/agent` add entries to `DefaultModels`. Move these into the map literal declaration.

- In `internal/agent/providers.go`:
  - Move the 4 entries from `init()` into the `DefaultModels` map literal at declaration
  - Remove the `init()` function

- In `internal/agent/autodetect.go`:
  - Move the `"openrouter": ConservativeModels` entry into the `DefaultModels` map literal (already declared in `providers.go`)
  - Since `ConservativeModels` is declared in `autodetect.go`, the map literal in `providers.go` can't reference it at init time. Instead, merge: move `ConservativeModels` to `providers.go` alongside the other `Default*Models` vars, then add `"openrouter": ConservativeModels` to the `DefaultModels` literal, and remove both `init()` functions
  - Remove the `init()` function

- Run `go test ./internal/agent/... -count=1` to verify.

### Step 3: Consolidate cmd/certify init() into a single registerCommands() function

Cobra's standard pattern uses `init()` for flag binding, but having 10 `init()` functions means every unit in the file gets the `has_init_func` observation. Consolidate all subcommand registration and flag binding into a single `registerCommands()` function called from `main()`.

- In `cmd/certify/main.go`:
  - Add a call to `registerCommands()` before `rootCmd.Execute()`

- In `cmd/certify/root.go`:
  - Replace the `init()` function with `func registerCommands()` that:
    - Adds the persistent `--workspace` flag
    - Registers all subcommands: `versionCmd`, `initCmd`, `scanCmd`, `certifyCmd`, `reportCmd`, `modelsCmd`, `architectCmd`, `reviewCmd`, `expireCmd`
    - Calls individual flag-binding functions: `bindCertifyFlags()`, `bindScanFlags()`, `bindReportFlags()`, etc.

- In every other `cmd/certify/*.go` file that has `init()`:
  - Rename `init()` to `bind<Command>Flags()` (e.g., `bindCertifyFlags()`, `bindScanFlags()`, etc.)
  - The function body stays the same (just flag bindings)

- Specific renames:
  - `certify_cmd.go`: `init()` → `bindCertifyFlags()`
  - `scan.go`: `init()` → `bindScanFlags()`
  - `report_cmd.go`: `init()` → `bindReportFlags()`
  - `architect_cmd.go`: `init()` → `bindArchitectFlags()`
  - `init_cmd.go`: `init()` → `bindInitFlags()`
  - `expire.go`: `init()` → `bindExpireFlags()`
  - `models_cmd.go`: `init()` → `bindModelsFlags()`
  - `review.go`: `init()` → `bindReviewFlags()`
  - `version.go`: remove the `init()` entirely (it only calls `rootCmd.AddCommand(versionCmd)` which is now in `registerCommands()`)

- Run `go build ./cmd/certify/` and `go test ./cmd/certify/... -count=1` to verify.

### Step 4: Reduce cmd/certify global variables — move flag vars into cobra flag lookups

Each file's `var` block for flags (e.g., `certifyPath`, `certifySkipAgent`, etc.) inflates `global_mutable_count` for all functions in that file. Replace with cobra's `cmd.Flags().GetString()` / `cmd.Flags().GetBool()` lookups at use time, or consolidate into a per-command struct.

**Strategy:** For each command file, replace the standalone `var` flag variables with direct flag value lookups in the `RunE` function. Keep command variables (`var certifyCmd = ...`) since cobra needs them.

- In `cmd/certify/certify_cmd.go`:
  - Remove: `var certifyPath, certifySkipAgent, certifyBatch, certifyResetQueue, certifyTarget, certifyDiffBase`
  - In `bindCertifyFlags()`: use `certifyCmd.Flags().StringVar(&v, ...)` where `v` is a local or use `certifyCmd.Flags().String("path", "", "...")` without a var binding
  - In `runCertify()` and helpers: read flags via `cmd.Flags().GetString("path")`, `cmd.Flags().GetBool("skip-agent")`, etc.
  - The `certifyContext` struct can hold the resolved flag values
  - Update `loadCertifyContext()` to accept the flag values as params or read from a passed `*cobra.Command`
  - Update `runWorkspaceCertify()` to accept flag values

- In `cmd/certify/scan.go`:
  - Remove `var scanPath`
  - Read via `cmd.Flags().GetString("path")` in RunE

- In `cmd/certify/report_cmd.go`:
  - Remove `var reportPath, reportFormat, reportOutput, reportSite, siteOutput`
  - Read via `cmd.Flags().Get*()` in RunE

- In `cmd/certify/architect_cmd.go`:
  - Remove `var architectPath, architectModel, architectPhase, architectOutput, architectVerbose`
  - Read via `cmd.Flags().Get*()` in RunE

- In `cmd/certify/init_cmd.go`:
  - Remove `var initPath, initForce`
  - Read via `cmd.Flags().Get*()` in RunE

- In `cmd/certify/expire.go`:
  - Remove `var expirePath`
  - Read via `cmd.Flags().GetString("path")` in RunE

- In `cmd/certify/models_cmd.go`:
  - Remove `var modelsBaseURL, modelsAPIKey`
  - Read via `cmd.Flags().Get*()` in RunE

- In `cmd/certify/review.go`:
  - Remove `var reviewPath`
  - Read via `cmd.Flags().GetString("path")` in RunE

- In `cmd/certify/root.go`:
  - Remove `var workspaceMode`
  - Read via `cmd.Flags().GetBool("workspace")` — but since it's a persistent flag used in RunE functions, pass `cmd` to helpers

- In `cmd/certify/version.go`:
  - Keep `var Version, Commit, Date` as these are ldflags-injected, not flag vars

- Run `go build ./cmd/certify/` and `go test ./cmd/certify/... -count=1` to verify.

### Step 5: Fix errors_ignored — handle all discarded error returns

Replace `_ = ` and `val, _ = ` patterns with proper error handling.

- `cmd/certify/certify_cmd.go` line 135, 378: `root, _ = os.Getwd()` → `root, err := os.Getwd(); if err != nil { return fmt.Errorf("getting working directory: %w", err) }`
- `cmd/certify/certify_cmd.go` line 447: `apiKey, _ = agent.DetectAPIKey()` → this is intentional (key is empty string if not found, which is the fallback). Add an explicit comment: `apiKey, _ = agent.DetectAPIKey() // intentional: empty key triggers fallback to local`
  - Actually, the structural analyzer detects `_ =` as error ignored. The second return is `envVar string`, not error. So this shouldn't trigger errors_ignored. Check if it does. If it does, the pattern is fine but might need a named discard like `apiKey, _envVar = ...` — however, that won't help. Better to just accept this one or refactor to `apiKey = agent.DetectAPIKeyValue()`.
- `cmd/certify/architect_cmd.go` line 60: same `root, _ = os.Getwd()` fix
- `cmd/certify/architect_cmd.go` line 194: same `apiKey, _ = agent.DetectAPIKey()` — same as above
- All other `cmd/certify/*.go` files: find and fix `_, _ =` patterns
- `internal/policy/evaluator.go` line 152: `_, _ = fmt.Sscanf(...)` — Sscanf returns (n int, err error). The error is intentionally discarded (pattern matching, not error). Replace with explicit handling or a helper.
- `internal/record/store.go` line 456: `_ = json.Unmarshal(...)` — if unmarshal fails, `details` stays zero-value which is the intended fallback. Add a comment or handle.
- `internal/report/architect_report.go` line 372: `_ = grade` — this is a dummy usage to avoid unused variable. Remove the variable if unused.
- `internal/evidence/structural.go`: 2 patterns — check which they are and fix.

- Run `go test ./... -count=1` after each fix.

### Step 6: Fix remaining TODO comment and param_count

- `internal/engine/scorer.go` line 111: Change `// TODO count already handled by violation penalties` to `// NOTE: todo_count observation already handled by policy violation penalties`. The word "TODO" triggers the metric. Just rewording avoids the false positive.

- `cmd/certify/certify_cmd.go` — `buildCertificationRun` has 8 params. Refactor to accept a struct:
  ```go
  type runResult struct {
      runID          string
      startedAt      time.Time
      commit         string
      policyVersions []string
      certified      int
      failed         int
      processed      int
  }
  func buildCertificationRun(r runResult, store *record.Store) domain.CertificationRun
  ```

- Run `go test ./... -count=1` after each fix.

### Step 7: Run full validation and re-certify

- Run `go fmt ./...`
- Run `go vet ./...`
- Run `go test ./... -count=1`
- Run `go build -o build/bin/certify ./cmd/certify/`
- Run `./build/bin/certify certify --skip-agent --reset-queue` (re-certify with new code)
- Run `./build/bin/certify report` (regenerate report card)
- Verify observation count dropped significantly (target: 48 → <15)
- Verify average score improved (target: 85.4% → >88%)
- Verify cmd/certify grade improved (target: C 78.4% → B or better)

## Validation Commands

Execute every command to validate the chore is complete with zero regressions.

```bash
# Format check
gofmt -l ./cmd/certify/ ./internal/domain/ ./internal/agent/ ./internal/engine/ ./internal/evidence/ ./internal/policy/ ./internal/record/ ./internal/report/

# Vet check
go vet ./...

# Full test suite — all 16 packages must pass
go test ./... -count=1

# Build succeeds
go build -o build/bin/certify ./cmd/certify/

# Verify no init() in domain or agent
grep -rn "^func init()" internal/domain/ internal/agent/ --include="*.go" | grep -v _test.go

# Verify no init() in cmd/certify (should return empty)
grep -rn "^func init()" cmd/certify/ --include="*.go" | grep -v _test.go

# Verify reduced global var count in cmd/certify
grep -c "^var " cmd/certify/*.go | grep -v "_test.go"

# Self-certify and check results
./build/bin/certify certify --skip-agent --reset-queue
./build/bin/certify report
```

## Report

**Date:** 2026-03-10

### What Was Implemented

All 7 steps from the plan were executed in order:

1. **internal/domain init() elimination** — 4 init() functions replaced with compile-time map literals for `stringToStatus`, `stringToUnitType`, `stringToEvidenceKind`, `stringToSeverity`
2. **internal/agent init() elimination** — 2 init() functions merged into `DefaultModels` map literal declaration; `ConservativeModels` referenced directly via Go same-package init order guarantee
3. **cmd/certify init() consolidation** — 10 init() functions converted to named `bind*Flags()` functions; single `registerCommands()` called from `main()`; `TestMain` added for test setup
4. **cmd/certify global var elimination** — All flag `var` blocks removed; replaced with `flagString/flagBool/flagInt/flagStringSlice` helpers wrapping cobra `GetString/GetBool/GetInt`; `certifyFlags` struct for certify command; `workspaceMode` global removed
5. **errors_ignored fixes** — `json.Unmarshal`, `ParseEvidenceKind`, `time.Parse`, `fmt.Sscanf` error handling fixed; dead `grade` variable removed; `detectAPIKeyOnly()` wrapper added; `os.ReadDir` error handled
6. **TODO + param_count fixes** — Reworded "TODO count" comment to "Note:"; `buildCertificationRun` 8 params → `runParams` struct (2 params)
7. **Full validation + re-certification** — gofmt, go vet, go test (16/16 pass), build, scan (748 units), certify (773 units), report

### Results

| Metric | Before | After | Delta |
|--------|--------|-------|-------|
| Overall Score | 85.4% | 85.8% | +0.4% |
| cmd/certify Grade | C (78.4%) | B (83.3%) | +4.9% |
| Observations | 48 | 29 | -19 (40% reduction) |
| init() functions | 16 | 0 | -16 |
| has_init_func observations (fresh run) | 70 | 0 | -70 |
| Global var count (cmd/certify) | 3-7 per file | 1 per file | -80% |
| C-grade units | 48 | 29 | -19 |
| Total units | 615 | 773 | +158 (new code) |

### Issues Encountered

1. **Cobra flag pattern creates errors_ignored**: Replacing `var` flag bindings with `cmd.Flags().GetString()` introduced `val, _ := ...` patterns that the structural analyzer counts as `errors_ignored`. Fixed by creating `flagString/flagBool/flagInt/flagStringSlice` helper functions that properly handle the error return.

2. **Stale records inflate observation count**: Old certification records from removed units (like deleted `init()` functions) persist in `records/` and are counted in reports. The fresh run had zero `has_init_func` observations but the total shows 14 from stale records.

3. **DetectAPIKey returns (string, string) not (string, error)**: The structural analyzer counts blank identifiers in ALL multi-value assignments, not just error returns. Created `detectAPIKeyOnly()` wrapper to isolate the pattern.

4. **gofmt needed after compile-time map edits**: The map literal alignment in `record.go` required reformatting.

### Refactoring Done

- **Flag helpers in root.go**: `flagString`, `flagBool`, `flagInt`, `flagStringSlice` — clean abstraction over cobra's error-returning flag getters
- **certifyFlags struct**: Groups all certify command flag values, passed to helpers instead of global state
- **runParams struct**: Replaces 8-parameter `buildCertificationRun` with structured input
- **registerCommands()**: Single entry point for all subcommand registration, called from `main()`
- **TestMain in cli_test.go**: Explicit test setup replacing implicit init() registration

### Commits

```
248ca74 chore: eliminate init() from internal/domain and internal/agent
ad8723f chore: consolidate cmd/certify init() into registerCommands()
844581c chore: eliminate global flag vars from cmd/certify — use cobra flag lookups
79b9ff3 fix: eliminate errors_ignored across codebase
a89ec97 fix: resolve TODO comment and param_count violations
8cdfcc8 chore: gofmt
562ec83 chore: self-certify after architect remediation — 773 units, B (85.8%)
```

### FEATURES.md

All criteria were already checked off — this chore improved code quality metrics without changing feature behavior.

## Notes

- **Cobra flag pattern change:** Moving from `var` + `init()` + `StringVar()` to `cmd.Flags().String()` + `cmd.Flags().GetString()` in RunE is a well-established alternative pattern. The key tradeoff is that flag values are only available inside RunE (not in package-level helpers) — helpers must receive values as params or via a context struct.

- **Domain init() is safe to remove:** The reverse maps are trivially constructed from the forward maps. Since Go doesn't support `const` maps, the `var` declarations remain, but eliminating `init()` removes the `has_init_func` observation from all 72 units.

- **Go maps are not const:** Even after removing `init()`, `var mapName = map[...]...{}` still counts as `global_mutable_count`. The discovery package's 8 map vars are effectively constant but counted as mutable. To eliminate those observations, we'd need to wrap them in functions (`func certifiableExts() map[...]... { return map[...]...{} }`) — which is a larger refactor deferred for a future chore. The threshold of 2 means packages with ≤2 map vars are fine.

- **`isLocalURL` AI observations:** The LLM flagged `isLocalURL` as a security issue. The function is correct for its use case (detecting localhost URLs for timeout/strategy selection) but the observation is valid — `net/url.Parse` would be more robust. This is a low-priority fix compared to the structural wins above.

- **Version vars must stay global:** `Version`, `Commit`, `Date` in `version.go` are set via `-ldflags` at build time. They must remain package-level `var` declarations.

- **Expected final state:**
  - `cmd/certify`: 0 init(), ~2 global vars (rootCmd, versionCmd-level), grade C→B+
  - `internal/domain`: 0 init(), ~same global vars but no init penalty, grade B→B+
  - `internal/agent`: 0 init(), ~same global vars, grade B→B+
  - Overall: 85.4% → ~89%, observations 48→<15
