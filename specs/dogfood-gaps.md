# Dogfood Gap Analysis — What Breaks on an Unknown Repo

After self-certifying `code-certification`, here's everything that would fail or produce garbage results on a fresh unknown repository.

## Critical: Evidence Collection Is Fake

**The #1 problem.** The `certify certify` command currently:
- Hardcodes `Passed: true` for lint and test evidence with a comment "verified at build time"
- This is a lie — we never actually ran any lint or test tools
- On an unknown repo, every unit gets a free pass regardless of actual code quality

### What's needed

#### Tool Runners (`internal/evidence/runner.go`)
Actual tool execution with output parsing:

| Tool | Language | Output Format | Parser Status |
|------|----------|--------------|---------------|
| `go vet` | Go | stderr text | ❌ Not built |
| `golangci-lint run --out-format json` | Go | JSON | ❌ Not built |
| `go test -json ./...` | Go | JSON stream | ❌ Not built |
| `go test -coverprofile` | Go | coverage file | ❌ Not built |
| `eslint --format json` | TS/JS | JSON | ❌ Not built |
| `jest --json` | TS/JS | JSON | ❌ Not built |
| `tsc --noEmit` | TS | stderr text | ❌ Not built |

We have `LintResult`, `TestResult`, `CodeMetrics`, `GitStats` structs and `ToEvidence()` methods — but nothing that actually **runs the tools and parses their output**.

#### Git Evidence Runner
We have `ParseGitLog()` but no function that actually runs `git log --format='%H\t%an\t%ad' -- <file>` and feeds it in. Need:
- `git log` runner per file
- Age calculation (first commit date → now)
- Churn rate integration into expiry factors

#### Evidence-to-Rule Metric Mapping Is Fragile
`extractMetric()` in `policy/evaluator.go` uses string matching (`"lint_errors"`, `"test_failures"`) and tries to pull numbers from evidence summaries via `fmt.Sscanf`. This breaks easily:
- Summary format changes → metric extraction fails → "missing evidence" violation
- Should use typed evidence details instead of string parsing

## Critical: Init Doesn't Detect Languages

`certify init` creates a static config and one global policy. On an unknown repo it should:

1. **Walk the repo and detect languages** — count `.go`, `.ts`, `.py`, `.rs`, etc. files
2. **Generate language-specific starter policies** — Go repo gets go-standard.yml, TS repo gets ts-standard.yml, Python repo gets python-standard.yml
3. **Detect available tooling** — check if `golangci-lint`, `eslint`, `jest`, `pytest` are installed or in package.json/go.mod
4. **Set scope patterns automatically** — detect `vendor/`, `node_modules/`, `dist/`, common generated paths
5. **Generate GitHub workflow files** — we have the generators but init doesn't call them
6. **Offer to create a bootstrap PR** — currently just writes to disk, should optionally `gh pr create`

## High: Scan Blindly Runs All Adapters

`certify scan` runs Generic + Go + TS adapters on every repo regardless of content. On a Python repo:
- Go adapter finds zero `.go` files → fine but wasted work
- TS adapter finds zero `.ts` files → same
- Generic scanner finds everything → produces file-level units only

### What's needed
- Detect languages first, only run relevant adapters
- Registry pattern: `map[string]Scanner` keyed by detected language
- Adapter auto-detection: if repo has `go.mod` → use Go adapter, if `package.json` → TS adapter

## High: TODO Count Applies Per-File Not Per-Symbol

Every symbol in `evaluator.go` gets the same "todo_count: 11 exceeds threshold" because metrics are computed at file level but rules are evaluated per symbol-level unit. This means:
- A file with 1 TODO in function A penalizes functions B, C, D equally
- Need per-symbol metric scoping, or at minimum per-file dedup

## High: Agent Review Not Wired Into Pipeline

Config has `agent.enabled: true` with model assignments, but `certify certify` completely ignores it. The `Reviewer` exists in `internal/agent/` but is never called from the CLI command. Need:
1. Check `cfg.Agent.Enabled` in certify command
2. Read `OPENROUTER_API_KEY` from env
3. Create provider + router + reviewer
4. Call `reviewer.Review()` for units that need it
5. Merge agent evidence with deterministic evidence
6. Respect the "agent cannot override deterministic failures" rule

## Medium: No `--path` Flag

CLI commands use `os.Getwd()` always. No way to certify a different directory:
```bash
certify certify --path /some/other/repo  # doesn't work
```

## Medium: Missing `certify expire` Logic for Recertification

`certify expire` marks records as expired but doesn't trigger recertification. On a real repo you'd want:
- `certify expire --recertify` to immediately re-evaluate expired units
- Integration with nightly workflow to auto-recertify

## Medium: No Complexity Measurement

The `max-complexity` rule references metric `"complexity"` but `extractMetric()` returns `0` with a `// TODO: implement complexity extraction` comment. Need:
- Go: use `go/ast` to count branches (if/for/switch/select/case) per function
- TS: regex-based branch counting
- Feed actual complexity values into evidence

## Medium: Index Dedup Creates Noise

`Merge()` in `discovery/scanner.go` combines generic file-level units with Go symbol-level units. A file like `main.go` appears as:
- `go://main.go` (file-level from generic)
- `go://main.go#main` (function from Go adapter)
- `go://main.go#helper` (function from Go adapter)

The file-level entry is redundant when symbol-level entries exist. Should filter out file-level entries when a language adapter provided symbol-level ones for the same file.

## Low: Workflow Files Not Written to Disk

`internal/github/workflows.go` generates YAML strings but `certify init` doesn't write them to `.github/workflows/`. Need:
```go
os.MkdirAll(".github/workflows", 0o755)
os.WriteFile(".github/workflows/certification-pr.yml", []byte(GeneratePRWorkflow()), 0o644)
```

## Low: No Config Merge/Override

Can't override config via CLI flags or environment variables. Useful for CI:
```bash
certify certify --mode enforcing  # override config.yml mode
CERTIFY_MODE=enforcing certify certify  # env override
```

## Low: Override Loader Not Called from CLI

`internal/override/` exists but `certify certify` doesn't load or apply overrides from `.certification/overrides/`.

## Summary: Priority Order

1. **Tool runners** — without real evidence, certification is meaningless
2. **Language detection in init** — unknown repos get wrong setup
3. **Agent wiring** — config says enabled but nothing happens
4. **Smart scanning** — adapter selection based on detected languages
5. **Per-symbol metric scoping** — file-level metrics unfairly penalize all symbols
6. **Complexity measurement** — promised but not delivered
7. **Workflow file output** — init should produce `.github/workflows/`
8. **Override loading** — exists but not wired
9. **CLI --path flag** — convenience for CI
10. **Config overrides** — env/flag merging
