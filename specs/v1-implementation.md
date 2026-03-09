# Plan: Code Certification System v1 Implementation

## Task Description

Implement the complete Code Certification System v1 as defined by PRD.md (§21 MVP Scope), FEATURES.md (§1–§27), and STORIES.md (20 epics, 50+ stories). This is the full product: a Go CLI (`certify`) that discovers certifiable code units in a repository, evaluates them against versioned policies, collects deterministic evidence, optionally performs agent-assisted review via OpenRouter, assigns time-bound certification status, manages expiration and recertification, integrates with GitHub workflows, creates remediation issues, and generates reports.

Every line of code is implemented via TDD: failing test → minimal implementation → refactor.

## Objective

When this plan is complete, the `certify` CLI can:
1. Bootstrap certification in any repository (`certify init`)
2. Discover and index code units with stable identifiers (`certify scan`)
3. Evaluate units against YAML policy packs (`certify certify`)
4. Collect evidence from linters, tests, static analysis, git history, and metrics
5. Score across 9 quality dimensions with configurable weights
6. Assign certification status (certified / certified_with_observations / probationary / expired / decertified / exempt)
7. Compute risk-based expiration windows with trust decay
8. Optionally run agent-assisted review via OpenRouter free models
9. Store structured certification records in `.certification/records/`
10. Generate machine-readable and human-readable reports (`certify report`)
11. Expire outdated certifications (`certify expire`)
12. Operate in advisory or enforcing mode
13. Target specific scopes (file, directory, changed set)
14. Integrate with GitHub via workflow YAML files and the `gh` CLI for PR annotations and issue sync

## Problem Statement

The Code Certification System exists only as product documentation (PRD, FEATURES, STORIES) and project tooling. Zero Go source code exists. The entire product must be implemented from scratch following strict TDD, producing a working v1 CLI that satisfies the MVP scope (PRD §21) and the minimum v1 readiness criteria (FEATURES §26).

## Solution Approach

Build the system bottom-up in 8 phases, each independently testable and committable. Each phase produces working, tested code that can be built and run. Later phases compose earlier ones. The architecture follows clean separation: domain types → storage → policy engine → evidence → discovery → evaluation engine → agent review → CLI → GitHub integration.

**Module path**: `github.com/code-certification/certify`

**Key architectural boundaries**:
- `internal/domain/` — Pure domain types with zero external dependencies
- `internal/config/` — YAML configuration loading and validation
- `internal/policy/` — Policy loading, matching, and evaluation
- `internal/discovery/` — Unit discovery with language adapter interface
- `internal/evidence/` — Evidence collection with analyzer adapter interface
- `internal/engine/` — Certification pipeline orchestrator
- `internal/agent/` — OpenRouter provider, model routing, agent review pipeline
- `internal/record/` — Certification record persistence (JSON files)
- `internal/report/` — Report generation (JSON + text)
- `internal/expiry/` — Expiration computation and trust decay
- `internal/github/` — GitHub integration (PR annotations, issue sync, workflow generation)
- `cmd/certify/` — CLI entry point using `cobra`

## Relevant Files

### Existing Files
- `PRD.md` — Full product requirements, certification model, architecture, MVP scope
- `FEATURES.md` — 200+ measurable acceptance criteria across 27 sections
- `STORIES.md` — 50+ user stories across 20 epics
- `CLAUDE.md` — Project context and architecture decisions
- `specs/project-bootstrap-justfile-commands-initial-commit.md` — Agent architecture, OpenRouter model catalog, structured output schemas, provider interface design
- `justfile` — Build, test, lint recipes (already handles missing source gracefully)

### New Files

Every file below is created via TDD (test file first, then implementation).

#### Go Module
- `go.mod` — Module definition
- `go.sum` — Dependency checksums

#### Domain Types (`internal/domain/`)
- `internal/domain/unit.go` — CertificationUnit, UnitID, UnitType
- `internal/domain/unit_test.go`
- `internal/domain/record.go` — CertificationRecord, Status, Grade
- `internal/domain/record_test.go`
- `internal/domain/policy.go` — PolicyPack, PolicyRule, Severity, Scope
- `internal/domain/policy_test.go`
- `internal/domain/evidence.go` — Evidence, EvidenceSource, EvidenceKind
- `internal/domain/evidence_test.go`
- `internal/domain/dimension.go` — Dimension, DimensionScores, dimension constants
- `internal/domain/dimension_test.go`
- `internal/domain/expiry.go` — ExpiryWindow, ExpiryFactors
- `internal/domain/expiry_test.go`
- `internal/domain/override.go` — Override, OverrideAction
- `internal/domain/override_test.go`
- `internal/domain/config.go` — Config, AgentConfig, ScopeConfig
- `internal/domain/config_test.go`

#### Configuration (`internal/config/`)
- `internal/config/loader.go` — YAML config loading, validation, defaults
- `internal/config/loader_test.go`
- `internal/config/policy_loader.go` — Policy pack loading from `.certification/policies/`
- `internal/config/policy_loader_test.go`

#### Discovery (`internal/discovery/`)
- `internal/discovery/scanner.go` — Scanner interface and orchestrator
- `internal/discovery/scanner_test.go`
- `internal/discovery/generic.go` — Generic file-level adapter (all languages)
- `internal/discovery/generic_test.go`
- `internal/discovery/go_adapter.go` — Go-specific symbol discovery (go/ast)
- `internal/discovery/go_adapter_test.go`
- `internal/discovery/ts_adapter.go` — TypeScript symbol discovery (regex-based for v1)
- `internal/discovery/ts_adapter_test.go`
- `internal/discovery/index.go` — Unit index persistence (JSON)
- `internal/discovery/index_test.go`
- `internal/discovery/diff.go` — Change detection between revisions
- `internal/discovery/diff_test.go`

#### Evidence (`internal/evidence/`)
- `internal/evidence/collector.go` — Collector interface and orchestrator
- `internal/evidence/collector_test.go`
- `internal/evidence/lint.go` — Lint result ingestion (golangci-lint, ESLint)
- `internal/evidence/lint_test.go`
- `internal/evidence/test_results.go` — Test result ingestion (go test, jest)
- `internal/evidence/test_results_test.go`
- `internal/evidence/metrics.go` — Code metrics (complexity, size, churn)
- `internal/evidence/metrics_test.go`
- `internal/evidence/git.go` — Git history analysis (churn, age, authors)
- `internal/evidence/git_test.go`

#### Policy Evaluation (`internal/policy/`)
- `internal/policy/evaluator.go` — Policy evaluation engine
- `internal/policy/evaluator_test.go`
- `internal/policy/matcher.go` — Policy-to-unit matching (language, path, pattern)
- `internal/policy/matcher_test.go`
- `internal/policy/rules.go` — Built-in rule implementations
- `internal/policy/rules_test.go`

#### Scoring & Expiry (`internal/engine/`)
- `internal/engine/scorer.go` — Dimension scoring with configurable weights
- `internal/engine/scorer_test.go`
- `internal/engine/grader.go` — Grade computation (A through F)
- `internal/engine/grader_test.go`
- `internal/engine/status.go` — Status determination logic
- `internal/engine/status_test.go`
- `internal/engine/pipeline.go` — Certification pipeline orchestrator
- `internal/engine/pipeline_test.go`

#### Expiry (`internal/expiry/`)
- `internal/expiry/calculator.go` — Expiry window computation with risk factors
- `internal/expiry/calculator_test.go`
- `internal/expiry/decay.go` — Trust decay and window adjustment
- `internal/expiry/decay_test.go`

#### Records (`internal/record/`)
- `internal/record/store.go` — Record persistence (JSON files in `.certification/records/`)
- `internal/record/store_test.go`
- `internal/record/history.go` — Historical record queries
- `internal/record/history_test.go`

#### Reporting (`internal/report/`)
- `internal/report/health.go` — Repository health report
- `internal/report/health_test.go`
- `internal/report/risk.go` — Risk-focused report
- `internal/report/risk_test.go`
- `internal/report/trend.go` — Trend report (over time)
- `internal/report/trend_test.go`
- `internal/report/formatter.go` — JSON and text output formatters
- `internal/report/formatter_test.go`

#### Agent Review (`internal/agent/`)
- `internal/agent/types.go` — ChatRequest, ChatResponse, Message, Tool types
- `internal/agent/types_test.go`
- `internal/agent/provider.go` — Provider interface
- `internal/agent/openrouter.go` — OpenRouterProvider implementation
- `internal/agent/openrouter_test.go`
- `internal/agent/router.go` — Task-to-model routing with fallback chain
- `internal/agent/router_test.go`
- `internal/agent/ratelimit.go` — Token-bucket rate limiter
- `internal/agent/ratelimit_test.go`
- `internal/agent/schemas.go` — Structured output JSON schemas
- `internal/agent/schemas_test.go`
- `internal/agent/prompts.go` — Prompt template loading and versioning
- `internal/agent/prompts_test.go`
- `internal/agent/reviewer.go` — Agent review pipeline (prescreen → review → score → decide → remediate)
- `internal/agent/reviewer_test.go`

#### GitHub Integration (`internal/github/`)
- `internal/github/pr.go` — PR annotation via `gh` CLI
- `internal/github/pr_test.go`
- `internal/github/issues.go` — Issue creation and sync via `gh` CLI
- `internal/github/issues_test.go`
- `internal/github/workflows.go` — Workflow YAML generation
- `internal/github/workflows_test.go`

#### Override System (`internal/override/`)
- `internal/override/loader.go` — Override loading from `.certification/overrides/`
- `internal/override/loader_test.go`
- `internal/override/applier.go` — Override application to certification results
- `internal/override/applier_test.go`

#### CLI (`cmd/certify/`)
- `cmd/certify/main.go` — Entry point
- `cmd/certify/root.go` — Root command setup (cobra)
- `cmd/certify/init.go` — `certify init` command (bootstrap)
- `cmd/certify/scan.go` — `certify scan` command (discovery)
- `cmd/certify/certify_cmd.go` — `certify certify` command (evaluate)
- `cmd/certify/expire.go` — `certify expire` command
- `cmd/certify/report.go` — `certify report` command
- `cmd/certify/review.go` — `certify review` command (PR-focused)
- `cmd/certify/version.go` — `certify version` command

#### Test Fixtures
- `testdata/repos/go-simple/` — Simple Go repo fixture
- `testdata/repos/ts-simple/` — Simple TypeScript repo fixture
- `testdata/repos/polyglot/` — Multi-language repo fixture
- `testdata/policies/` — Test policy packs
- `testdata/evidence/` — Sample evidence outputs
- `testdata/config/` — Test configuration files

#### Certification Templates (shipped with CLI)
- `templates/config.yml` — Default configuration template
- `templates/policies/global.yml` — Global starter policy
- `templates/policies/go.yml` — Go starter policy
- `templates/policies/typescript.yml` — TypeScript starter policy
- `templates/policies/security.yml` — Security starter policy
- `templates/workflows/certification-pr.yml` — PR workflow template
- `templates/workflows/certification-nightly.yml` — Nightly workflow template
- `templates/workflows/certification-weekly.yml` — Weekly workflow template
- `templates/prompts/prescreen.v1.md` — Agent prescreen prompt
- `templates/prompts/review.v1.md` — Agent review prompt
- `templates/prompts/scoring.v1.md` — Agent scoring prompt
- `templates/prompts/decision.v1.md` — Agent decision prompt
- `templates/prompts/remediation.v1.md` — Agent remediation prompt

## Implementation Phases

### Phase 1: Domain Foundation
Pure domain types with zero external dependencies. All types are immutable value objects or well-defined structs with validation. This phase produces the vocabulary that all other code speaks.

**Covers**: FEATURES §1 (core domain model), §8 (status model), §9 (dimensions), §10 (record schema), §22 (structured schemas)

### Phase 2: Configuration & Policy
YAML loading for configuration and policy packs. Policy matching logic (language, path, pattern targeting). Policy evaluation against evidence.

**Covers**: FEATURES §3 (repository-local config), §4 (policy-as-code), §22 (schema validation)

### Phase 3: Discovery & Indexing
Code unit discovery with language adapter interface. Generic file-level adapter. Go-specific adapter using `go/ast`. TypeScript regex adapter. Unit index persistence. Change detection.

**Covers**: FEATURES §5 (unit discovery), §6 (language-agnostic model), §20 (multi-language adapters)

### Phase 4: Evidence Collection
Evidence collector interface with analyzer adapters. Lint result ingestion. Test result ingestion. Code metrics (complexity, size). Git history analysis (churn, age, stability).

**Covers**: FEATURES §7 (evidence collection)

### Phase 5: Certification Engine
Scoring across 9 dimensions with configurable weights. Grade computation. Status determination. Expiry window calculation with risk factors. Trust decay. Override application. Full certification pipeline orchestrator that composes discovery + evidence + policy + scoring + expiry.

**Covers**: FEATURES §8 (evaluation engine), §9 (dimensions), §10 (records), §11 (expiry), §12 (invalidation), §17 (overrides)

### Phase 6: Agent-Assisted Review
OpenRouter provider implementation. Task-to-model routing with fallback chain. Rate limiting. Structured output schemas. Prompt template loading. 5-step agent review pipeline (prescreen → review → score → decide → remediate). Full graceful degradation.

**Covers**: FEATURES §16 (agent-assisted review)

### Phase 7: Records, Reports & CLI
Record persistence (JSON files). Report generation (health, risk, trend). JSON and text formatters. Cobra CLI with commands: init, scan, certify, expire, report, review, version. Advisory and enforcing modes. Scoped execution (file, directory, changed set).

**Covers**: FEATURES §10 (trust ledger), §18 (reporting), §19 (CLI), §23 (rollout/advisory mode)

### Phase 8: GitHub Integration
Workflow YAML generation (PR, nightly, weekly). PR annotation via `gh` CLI. Issue creation and sync. Bootstrap PR creation. Merge blocking support.

**Covers**: FEATURES §2 (bootstrap), §12 (invalidation via PR), §13 (PR workflow), §14 (scheduled workflows), §15 (issue sync), §21 (security/permissions)

## Step by Step Tasks
IMPORTANT: Execute every step in order, top to bottom.

### 1. Initialize Go Module
- Run `go mod init github.com/code-certification/certify`
- Add initial dependencies: `gopkg.in/yaml.v3`, `github.com/spf13/cobra`
- Run `go mod tidy`
- Verify: `cat go.mod`

### 2. Domain Types — Unit & UnitID (TDD)
- **Test first**: `internal/domain/unit_test.go`
  - Test UnitID construction from language, path, symbol
  - Test UnitID string representation (`go://internal/service/sync.go#Apply`)
  - Test UnitID parsing from string
  - Test file-level fallback ID (`file://scripts/release.sh`)
  - Test UnitType enum (File, Function, Method, Class, Module, Package)
  - Test CertificationUnit struct construction and validation
- **Implement**: `internal/domain/unit.go`
- **Refactor**: Ensure immutability, add doc comments
- Verify: `go test ./internal/domain/ -run TestUnit -v`

### 3. Domain Types — Certification Status & Dimensions (TDD)
- **Test first**: `internal/domain/record_test.go`, `internal/domain/dimension_test.go`
  - Test Status enum (Certified, CertifiedWithObservations, Probationary, Expired, Decertified, Exempt)
  - Test Status.String() and ParseStatus()
  - Test Grade enum (A, A-, B+, B, C, D, F) with numeric thresholds
  - Test all 9 Dimension constants
  - Test DimensionScores construction with all 9 dimensions
  - Test DimensionScores.WeightedAverage() with configurable weights
  - Test DimensionScores.Grade() computation
- **Implement**: `internal/domain/record.go`, `internal/domain/dimension.go`
- Verify: `go test ./internal/domain/ -v`

### 4. Domain Types — Evidence & Policy (TDD)
- **Test first**: `internal/domain/evidence_test.go`, `internal/domain/policy_test.go`
  - Test Evidence struct (Source, Kind, RawData, Normalized, Timestamp, Missing flag)
  - Test EvidenceKind enum (Lint, TypeCheck, Test, StaticAnalysis, Metrics, GitHistory, AgentReview)
  - Test PolicyPack struct (Version, Name, Language targets, Path patterns, Rules)
  - Test PolicyRule struct (Dimension, Threshold, Severity, Description)
  - Test Severity enum (Info, Warning, Error, Critical)
- **Implement**: `internal/domain/evidence.go`, `internal/domain/policy.go`
- Verify: `go test ./internal/domain/ -v`

### 5. Domain Types — Expiry, Override, Config (TDD)
- **Test first**: `internal/domain/expiry_test.go`, `internal/domain/override_test.go`, `internal/domain/config_test.go`
  - Test ExpiryWindow struct with CertifiedAt, ExpiresAt
  - Test ExpiryFactors (base window, churn rate, test coverage, complexity, prior history)
  - Test Override struct (UnitID, Action, Rationale, Actor, Timestamp)
  - Test OverrideAction enum (Exempt, ExtendWindow, ShortenWindow, ForceReview)
  - Test CertificationRecord struct (full record with all fields from FEATURES §10)
  - Test Config struct with AgentConfig, ScopeConfig, ScheduleConfig
- **Implement**: the remaining domain files
- Verify: `go test ./internal/domain/ -v`
- **Commit**: `git add -A && git commit -m "feat: domain types — units, records, policies, evidence, dimensions, expiry, overrides"`

### 6. Configuration Loading (TDD)
- **Create test fixtures**: `testdata/config/valid.yml`, `testdata/config/minimal.yml`, `testdata/config/invalid.yml`
- **Test first**: `internal/config/loader_test.go`
  - Test loading valid config from YAML
  - Test defaults applied for missing optional fields
  - Test validation errors for invalid config
  - Test minimal config (only required fields)
  - Test include/exclude pattern parsing
  - Test agent config parsing (provider, models, rate limits)
  - Test advisory vs enforcing mode
- **Implement**: `internal/config/loader.go`
- Verify: `go test ./internal/config/ -v`

### 7. Policy Pack Loading (TDD)
- **Create test fixtures**: `testdata/policies/global.yml`, `testdata/policies/go.yml`, `testdata/policies/typescript.yml`
- **Test first**: `internal/config/policy_loader_test.go`
  - Test loading policy pack from YAML
  - Test language-targeted policies
  - Test path-pattern-targeted policies
  - Test policy version tracking
  - Test multiple policy packs loaded together
  - Test policy validation (missing required fields)
- **Implement**: `internal/config/policy_loader.go`
- Verify: `go test ./internal/config/ -v`
- **Commit**: `git add -A && git commit -m "feat: configuration and policy loading with YAML validation"`

### 8. Unit Discovery — Generic Adapter (TDD)
- **Create test fixtures**: `testdata/repos/go-simple/` (a few .go files), `testdata/repos/ts-simple/` (a few .ts files)
- **Test first**: `internal/discovery/scanner_test.go`, `internal/discovery/generic_test.go`
  - Test Scanner interface contract
  - Test generic adapter discovers all files in a directory
  - Test include/exclude pattern filtering
  - Test generated/vendor path exclusion
  - Test stable file-level UnitID assignment
  - Test discovery across nested directories
- **Implement**: `internal/discovery/scanner.go`, `internal/discovery/generic.go`
- Verify: `go test ./internal/discovery/ -run TestGeneric -v`

### 9. Unit Discovery — Go Adapter (TDD)
- **Test first**: `internal/discovery/go_adapter_test.go`
  - Test function-level discovery using go/ast
  - Test method-level discovery (receiver methods)
  - Test type/struct discovery
  - Test package-level discovery
  - Test stable UnitID with symbol (`go://path/file.go#FuncName`)
  - Test graceful fallback to file-level on parse error
- **Implement**: `internal/discovery/go_adapter.go`
- Verify: `go test ./internal/discovery/ -run TestGo -v`

### 10. Unit Discovery — TypeScript Adapter (TDD)
- **Test first**: `internal/discovery/ts_adapter_test.go`
  - Test function/const/class/export discovery via regex
  - Test .ts and .tsx file support
  - Test stable UnitID with symbol (`ts://path/file.ts#functionName`)
  - Test graceful fallback on unparseable files
- **Implement**: `internal/discovery/ts_adapter.go`
- Verify: `go test ./internal/discovery/ -run TestTS -v`

### 11. Unit Index & Change Detection (TDD)
- **Test first**: `internal/discovery/index_test.go`, `internal/discovery/diff_test.go`
  - Test index serialization/deserialization (JSON)
  - Test index update: new units, removed units, changed units
  - Test change detection between two index snapshots
  - Test git-based change detection (mock git diff output)
  - Test moved/renamed file handling
- **Implement**: `internal/discovery/index.go`, `internal/discovery/diff.go`
- Verify: `go test ./internal/discovery/ -v`
- **Commit**: `git add -A && git commit -m "feat: unit discovery with Go, TypeScript, and generic adapters + index management"`

### 12. Evidence Collection — Git History (TDD)
- **Test first**: `internal/evidence/git_test.go`
  - Test commit count extraction for a file
  - Test churn rate calculation (changes per time period)
  - Test file age detection
  - Test author count
  - Test evidence normalization to domain Evidence struct
  - Mock git command output for deterministic tests
- **Implement**: `internal/evidence/git.go`
- Verify: `go test ./internal/evidence/ -run TestGit -v`

### 13. Evidence Collection — Lint & Test Results (TDD)
- **Create test fixtures**: `testdata/evidence/golangci-lint.json`, `testdata/evidence/eslint.json`, `testdata/evidence/go-test.json`
- **Test first**: `internal/evidence/lint_test.go`, `internal/evidence/test_results_test.go`
  - Test golangci-lint JSON output parsing
  - Test ESLint JSON output parsing
  - Test go test JSON output parsing
  - Test evidence normalization for each tool
  - Test partial/missing evidence handling
  - Test collection failure reporting
- **Implement**: `internal/evidence/lint.go`, `internal/evidence/test_results.go`
- Verify: `go test ./internal/evidence/ -v`

### 14. Evidence Collection — Metrics (TDD)
- **Test first**: `internal/evidence/metrics_test.go`
  - Test file size measurement
  - Test line count
  - Test cyclomatic complexity approximation (Go)
  - Test test file presence detection
  - Test evidence normalization
- **Implement**: `internal/evidence/metrics.go`
- **Implement orchestrator**: `internal/evidence/collector.go`
- Verify: `go test ./internal/evidence/ -v`
- **Commit**: `git add -A && git commit -m "feat: evidence collection — git history, lint, test results, code metrics"`

### 15. Policy Evaluation (TDD)
- **Test first**: `internal/policy/evaluator_test.go`, `internal/policy/matcher_test.go`
  - Test policy matching: global policy matches all units
  - Test policy matching: language-specific policy matches only that language
  - Test policy matching: path-pattern matching
  - Test rule evaluation: threshold pass/fail
  - Test rule evaluation: severity assignment
  - Test multi-policy evaluation (global + language-specific combined)
  - Test violation recording
  - Test evaluation produces pass/observation/probation/failure outcomes
- **Implement**: `internal/policy/evaluator.go`, `internal/policy/matcher.go`, `internal/policy/rules.go`
- Verify: `go test ./internal/policy/ -v`
- **Commit**: `git add -A && git commit -m "feat: policy evaluation engine with matching and rule evaluation"`

### 16. Dimension Scoring & Grading (TDD)
- **Test first**: `internal/engine/scorer_test.go`, `internal/engine/grader_test.go`
  - Test score computation from evidence + policy results across 9 dimensions
  - Test configurable dimension weights
  - Test default equal weights
  - Test grade computation from weighted average (A=0.93+, A-=0.90+, B+=0.87+, B=0.80+, C=0.70+, D=0.60+, F=<0.60)
  - Test confidence computation
  - Test edge cases: missing evidence for a dimension, partial data
- **Implement**: `internal/engine/scorer.go`, `internal/engine/grader.go`
- Verify: `go test ./internal/engine/ -run "TestScor|TestGrad" -v`

### 17. Status Determination (TDD)
- **Test first**: `internal/engine/status_test.go`
  - Test certified: all dimensions above threshold, no critical violations
  - Test certified_with_observations: above threshold but minor violations exist
  - Test probationary: below threshold on some dimensions, no critical failures
  - Test decertified: critical violations or below minimum thresholds
  - Test exempt: override applied
  - Test deterministic failures always produce decertified regardless of agent scores
- **Implement**: `internal/engine/status.go`
- Verify: `go test ./internal/engine/ -run TestStatus -v`

### 18. Expiry Computation (TDD)
- **Test first**: `internal/expiry/calculator_test.go`, `internal/expiry/decay_test.go`
  - Test base window selection (new=30d, standard=90d, high-confidence=180d, critical=30d)
  - Test churn rate shortens window
  - Test low test coverage shortens window
  - Test high complexity shortens window
  - Test repeated passes lengthen window (up to 365d max)
  - Test poor history shortens window
  - Test configurable min/max bounds
  - Test trust decay: window shrinks on each failure, grows on each pass
- **Implement**: `internal/expiry/calculator.go`, `internal/expiry/decay.go`
- Verify: `go test ./internal/expiry/ -v`
- **Commit**: `git add -A && git commit -m "feat: certification engine — scoring, grading, status determination, expiry computation"`

### 19. Certification Pipeline (TDD)
- **Test first**: `internal/engine/pipeline_test.go`
  - Test full pipeline: discovery → evidence → policy eval → scoring → status → expiry → record
  - Test pipeline with advisory mode (no blocking)
  - Test pipeline with enforcing mode
  - Test pipeline with scoped execution (single file, directory, changed set)
  - Test pipeline produces valid CertificationRecord for each unit
  - Test pipeline handles missing evidence gracefully
  - Test pipeline respects overrides
  - Test pipeline skips excluded paths
- **Implement**: `internal/engine/pipeline.go`
- Verify: `go test ./internal/engine/ -run TestPipeline -v`

### 20. Override System (TDD)
- **Create test fixtures**: `testdata/config/overrides.yml`
- **Test first**: `internal/override/loader_test.go`, `internal/override/applier_test.go`
  - Test override loading from YAML
  - Test exempt override applied to unit
  - Test extend-window override applied
  - Test shorten-window override applied
  - Test force-review override applied
  - Test override requires rationale (validation)
  - Test overrides visible in certification record
- **Implement**: `internal/override/loader.go`, `internal/override/applier.go`
- Verify: `go test ./internal/override/ -v`
- **Commit**: `git add -A && git commit -m "feat: certification pipeline orchestrator with override support"`

### 21. Record Persistence (TDD)
- **Test first**: `internal/record/store_test.go`, `internal/record/history_test.go`
  - Test writing certification record to JSON file
  - Test reading certification record from JSON file
  - Test record file path derivation from UnitID
  - Test updating existing record (preserves history)
  - Test listing all records
  - Test querying records by status
  - Test historical record queries (last N certifications for a unit)
  - Test handling corrupted/missing record files gracefully
- **Implement**: `internal/record/store.go`, `internal/record/history.go`
- Verify: `go test ./internal/record/ -v`
- **Commit**: `git add -A && git commit -m "feat: certification record persistence with history tracking"`

### 22. Agent Types & Provider Interface (TDD)
- **Test first**: `internal/agent/types_test.go`
  - Test ChatRequest serialization to JSON (OpenAI-compatible)
  - Test ChatResponse deserialization from JSON
  - Test Message, Tool, ToolCall types
  - Test ModelConfig struct
  - Test TaskType enum
- **Implement**: `internal/agent/types.go`, `internal/agent/provider.go`
- Verify: `go test ./internal/agent/ -run TestTypes -v`

### 23. OpenRouter Provider (TDD)
- **Test first**: `internal/agent/openrouter_test.go`
  - Test successful chat completion with httptest mock server
  - Test request headers (Authorization, HTTP-Referer, X-Title, Content-Type)
  - Test 401 error handling (disable agent review)
  - Test 429 error handling (retry with backoff)
  - Test 502/503 error handling (fallback)
  - Test timeout handling
  - Test missing API key (graceful skip)
  - Test structured output request format
  - Test tool calling request format
- **Implement**: `internal/agent/openrouter.go`
- Verify: `go test ./internal/agent/ -run TestOpenRouter -v`

### 24. Rate Limiter & Model Router (TDD)
- **Test first**: `internal/agent/ratelimit_test.go`, `internal/agent/router_test.go`
  - Test token-bucket rate limiter (20 req/min)
  - Test rate limiter blocks when exhausted
  - Test rate limiter refills over time
  - Test model router selects correct model for each TaskType
  - Test model router falls back on primary model failure
  - Test model router falls back to emergency fallback
  - Test model router respects capability requirements (structured_outputs, tools, reasoning)
- **Implement**: `internal/agent/ratelimit.go`, `internal/agent/router.go`
- Verify: `go test ./internal/agent/ -run "TestRate|TestRouter" -v`

### 25. Agent Schemas & Prompts (TDD)
- **Test first**: `internal/agent/schemas_test.go`, `internal/agent/prompts_test.go`
  - Test dimension scoring schema validation
  - Test prescreen response schema validation
  - Test prompt template loading from files
  - Test prompt version extraction
  - Test prompt template variable substitution
- **Create prompt templates**: `templates/prompts/prescreen.v1.md`, `review.v1.md`, `scoring.v1.md`, `decision.v1.md`, `remediation.v1.md`
- **Implement**: `internal/agent/schemas.go`, `internal/agent/prompts.go`
- Verify: `go test ./internal/agent/ -run "TestSchema|TestPrompt" -v`

### 26. Agent Review Pipeline (TDD)
- **Test first**: `internal/agent/reviewer_test.go`
  - Test full pipeline: prescreen → review → score → decide → remediate
  - Test prescreen returns false → skip remaining steps
  - Test code review returns observations
  - Test dimension scoring returns valid structured JSON
  - Test status determination with reasoning
  - Test remediation generation for failing units
  - Test any step failure → fallback → skip, continue pipeline
  - Test all steps optional — valid certification without agent input
  - Test agent results marked with `source: "agent"` in evidence
  - Test agent cannot override deterministic failures
- **Implement**: `internal/agent/reviewer.go`
- Verify: `go test ./internal/agent/ -v`
- **Commit**: `git add -A && git commit -m "feat: agent-assisted review — OpenRouter provider, model routing, 5-step pipeline"`

### 27. Report Generation (TDD)
- **Test first**: `internal/report/health_test.go`, `internal/report/risk_test.go`, `internal/report/trend_test.go`, `internal/report/formatter_test.go`
  - Test health report: unit counts by status, average grade, expiring-soon count
  - Test risk report: highest-risk units, recurring failures, policy drift areas
  - Test trend report: certification coverage over time, debt trend
  - Test JSON formatter produces valid JSON
  - Test text formatter produces readable output
  - Test report storage to `.certification/reports/`
- **Implement**: `internal/report/health.go`, `internal/report/risk.go`, `internal/report/trend.go`, `internal/report/formatter.go`
- Verify: `go test ./internal/report/ -v`
- **Commit**: `git add -A && git commit -m "feat: reporting — health, risk, trend reports with JSON and text formatters"`

### 28. CLI — Root & Version Commands (TDD)
- **Test first**: Test cobra command setup, version output
- **Implement**: `cmd/certify/main.go`, `cmd/certify/root.go`, `cmd/certify/version.go`
- Verify: `go build -o build/bin/certify ./cmd/certify && ./build/bin/certify version`

### 29. CLI — Init Command (TDD)
- **Test first**: Test init creates `.certification/` directory structure, generates config, policies, workflow files
- **Implement**: `cmd/certify/init.go`
- **Create templates**: `templates/config.yml`, `templates/policies/*.yml`, `templates/workflows/*.yml`
- Verify: `./build/bin/certify init --path /tmp/test-repo`

### 30. CLI — Scan, Certify, Expire, Report, Review Commands (TDD)
- **Test first**: Test each command's argument parsing, option handling, and output
- **Implement**: `cmd/certify/scan.go`, `cmd/certify/certify_cmd.go`, `cmd/certify/expire.go`, `cmd/certify/report.go`, `cmd/certify/review.go`
- Verify: `./build/bin/certify --help` (all commands visible)
- **Commit**: `git add -A && git commit -m "feat: CLI — init, scan, certify, expire, report, review commands"`

### 31. GitHub Integration — Workflow Generation (TDD)
- **Test first**: `internal/github/workflows_test.go`
  - Test PR workflow YAML generation with correct permissions
  - Test nightly workflow YAML with schedule cron
  - Test weekly workflow YAML
  - Test OPENROUTER_API_KEY secret wiring in generated workflows
  - Test generated YAML is valid
- **Implement**: `internal/github/workflows.go`
- Verify: `go test ./internal/github/ -run TestWorkflow -v`

### 32. GitHub Integration — PR Annotations (TDD)
- **Test first**: `internal/github/pr_test.go`
  - Test PR annotation formatting (certified count, decertified count, observations)
  - Test `gh` CLI command construction for PR comments
  - Test advisory mode annotation (no blocking)
  - Test enforcing mode annotation (with exit code for blocking)
  - Test delta view (what improved, what worsened)
- **Implement**: `internal/github/pr.go`
- Verify: `go test ./internal/github/ -run TestPR -v`

### 33. GitHub Integration — Issue Sync (TDD)
- **Test first**: `internal/github/issues_test.go`
  - Test issue creation for decertified unit
  - Test issue body formatting (unit ID, violations, evidence, remediation)
  - Test duplicate issue detection (no spam)
  - Test issue update when status changes
  - Test issue close when unit re-certifies
  - Test grouped issue mode
  - Test label application
- **Implement**: `internal/github/issues.go`
- Verify: `go test ./internal/github/ -v`
- **Commit**: `git add -A && git commit -m "feat: GitHub integration — workflow generation, PR annotations, issue sync"`

### 34. Integration Tests
- **Test**: End-to-end test using `testdata/repos/go-simple/`:
  - `certify init` → creates `.certification/`
  - `certify scan` → discovers units, creates index
  - `certify certify` → evaluates all units, creates records
  - `certify report` → generates health report
  - `certify expire` → expires old certifications
- **Test**: End-to-end test with polyglot repo
- **Test**: End-to-end test with advisory vs enforcing mode
- **Test**: End-to-end test with agent review disabled (no API key)
- Verify: `go test ./... -v -count=1`

### 35. Final Quality Gates
- Run: `just fmt` — all files formatted
- Run: `just vet` — no vet issues
- Run: `just lint` — no lint issues
- Run: `just test` — all tests pass
- Run: `just build` — binary builds successfully
- Run: `just check` — all quality gates pass
- Run: `./build/bin/certify version` — prints version
- Run: `./build/bin/certify --help` — all commands listed
- **Commit**: `git add -A && git commit -m "feat: Code Certification System v1 — complete implementation"`

### 36. Validate Against Acceptance Criteria
- Run every command in `Validation Commands` section below
- Cross-reference results against FEATURES §26 (Minimum v1 Readiness Criteria)
- Verify all 12 v1 criteria are met

## Testing Strategy

### TDD Workflow
Every step follows: **Test → Implement → Refactor → Commit**. No code exists without a test that demands it.

### Unit Test Coverage
Every `internal/` package has a corresponding `_test.go` file. Target: >80% coverage on core domain, engine, and policy packages. Agent package tests use `httptest` mock servers.

### Test Fixtures
`testdata/` contains deterministic fixtures: sample repos, policy files, evidence outputs, configuration files. Tests never depend on external services or network.

### Integration Tests
End-to-end tests in Step 34 exercise the full pipeline using fixture repos. These verify that all components compose correctly.

### Manual Verification
After Step 35, manually run:
1. `certify init --path /tmp/manual-test` on a real Go repository
2. `certify scan --path /tmp/manual-test`
3. `certify certify --path /tmp/manual-test`
4. `certify report --path /tmp/manual-test`
5. Inspect `.certification/` directory contents

### Edge Cases
- Empty repository (no code files)
- Repository with only unsupported languages
- Corrupted/missing `.certification/` state
- Policy with impossible thresholds (all fail)
- Agent API key missing (graceful skip)
- Agent API rate limited (backoff and fallback)
- Very large repository (10,000+ files)
- Symlinked files and directories
- Binary files in repository
- Unicode filenames

## Acceptance Criteria

These map directly to FEATURES §26 (Minimum v1 Readiness Criteria):

1. A repository can be onboarded through `certify init` producing a reviewable `.certification/` structure
2. Repository-local configuration is created with sensible defaults
3. Starter policy packs are generated for detected languages
4. An initial unit index is created with stable identifiers
5. `certify certify` evaluates units and produces certification records
6. `certify expire` marks overdue certifications as expired
7. Certification records are written as structured JSON with all required fields
8. `certify report` generates machine-readable (JSON) and human-readable (text) reports
9. GitHub issue creation works via `gh` CLI integration
10. Go receives language-aware unit discovery (function/method level)
11. TypeScript receives at least basic symbol discovery
12. Unsupported languages receive generic file-level certification
13. Agent-assisted review works with OpenRouter when `OPENROUTER_API_KEY` is set
14. Agent-assisted review degrades gracefully when API key is absent
15. All tests pass: `just test`
16. All lint checks pass: `just lint`
17. Binary builds successfully: `just build`
18. CLI help is complete: `certify --help` lists all commands

## Validation Commands

```bash
# Full quality gate
just check

# Build
just build
./build/bin/certify version
./build/bin/certify --help

# Test coverage
just cover

# Test count
go test ./... -v -count=1 2>&1 | grep -c "^--- PASS"

# Integration: init a test repo
mkdir -p /tmp/certify-test && cd /tmp/certify-test && git init
certify init --path /tmp/certify-test
ls /tmp/certify-test/.certification/
cat /tmp/certify-test/.certification/config.yml

# Integration: scan and certify
certify scan --path /tmp/certify-test
certify certify --path /tmp/certify-test
certify report --path /tmp/certify-test --format json
certify report --path /tmp/certify-test --format text

# Cleanup
rm -rf /tmp/certify-test
```

## Notes

- **Go version**: 1.22+ (matching justfile and CLAUDE.md)
- **Dependencies**: Kept minimal. `cobra` for CLI, `yaml.v3` for YAML, `go/ast` for Go parsing (stdlib). No ORMs, no heavy frameworks.
- **No `go:embed` for templates in v1**: Templates loaded from filesystem. Embedding is a v1.1 enhancement.
- **Agent prompts are first drafts**: The prompt templates in `templates/prompts/` will need iteration based on real-world testing with OpenRouter models. Version them from day one (`v1`) so they can evolve.
- **TypeScript adapter is regex-based**: Full AST parsing (via node/swc) is a v1.1 enhancement. Regex covers 80% of cases (exported functions, classes, const).
- **GitHub integration uses `gh` CLI**: Direct GitHub API usage is a future enhancement. `gh` CLI provides auth handling, pagination, and error reporting for free.
- **File size limit**: Keep individual files under 500 lines. Split when approaching.
- **Advisory mode is default**: Bootstrap always generates advisory-mode config. Users explicitly opt into enforcing.
- **This plan produces ~70 test files and ~60 implementation files**: Budget approximately 35-40 TDD cycles (some steps combine related tests).
