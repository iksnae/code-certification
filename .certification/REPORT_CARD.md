# 🔵 Code Certification — Full Report

**Repository:** `//github.com/iksnae/code-certification`  
**Commit:** `8842f82`  
**Generated:** 2026-03-09T19:29:15  

---

## Summary

| Metric | Value |
|--------|-------|
| **Overall Grade** | 🔵 **B+** |
| **Overall Score** | 87.4% |
| **Total Units** | 446 |
| **Passing** | 446 |
| **Failing** | 0 |
| **Pass Rate** | 100.0% |
| **Observations** | 0 |
| **Expired** | 0 |

## Grade Distribution

| Grade | Count | % | Bar |
|:-----:|------:|----:|-----|
| B+ | 318 | 71.3% | ███████████████████████████████████ |
| B | 128 | 28.7% | ██████████████ |

## Dimension Averages

| Dimension | Score | Bar |
|-----------|------:|-----|
| architectural_fitness | 80.0% | ████████████████░░░░ |
| change_risk | 90.0% | █████████████████░░░ |
| correctness | 95.0% | ██████████████████░░ |
| maintainability | 93.2% | ██████████████████░░ |
| operational_quality | 85.0% | █████████████████░░░ |
| performance_appropriateness | 80.0% | ████████████████░░░░ |
| readability | 93.2% | ██████████████████░░ |
| security | 80.0% | ████████████████░░░░ |
| testability | 90.0% | █████████████████░░░ |

## By Language

### go — 🔵 B+ (87.4%)

- **Units:** 435
- **Score range:** 83.9% – 87.8%
- **Grades:** 314×B+, 121×B

### file — 🔵 B (86.5%)

- **Units:** 11
- **Score range:** 83.9% – 87.8%
- **Grades:** 4×B+, 7×B

## All Units

### `./` (5 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `CLAUDE.md` | file | B+ | 87.8% | certified | 2026-04-23 |
| `FEATURES.md` | file | B | 86.1% | certified | 2026-04-23 |
| `README.md` | file | B | 86.1% | certified | 2026-04-23 |
| `go.mod` | file | B+ | 87.8% | certified | 2026-04-23 |
| `go.sum` | file | B+ | 87.8% | certified | 2026-04-23 |

### `cmd/certify/` (28 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `certifyContext` | class | B+ | 87.8% | certified | 2026-04-23 |
| `certifyUnit` | method | B | 86.7% | certified | 2026-04-23 |
| `collectRepoEvidence` | method | B+ | 87.8% | certified | 2026-04-23 |
| `defaultConfigObj` | function | B+ | 87.8% | certified | 2026-04-23 |
| `filterUnits` | function | B+ | 87.8% | certified | 2026-04-23 |
| `init` | function | B+ | 87.8% | certified | 2026-04-23 |
| `loadCertifyContext` | function | B | 86.7% | certified | 2026-04-23 |
| `loadQueue` | function | B+ | 87.8% | certified | 2026-04-23 |
| `printQueueStatus` | method | B+ | 87.8% | certified | 2026-04-23 |
| `printSummary` | method | B+ | 87.8% | certified | 2026-04-23 |
| `processQueue` | method | B | 86.1% | certified | 2026-04-23 |
| `runCertify` | function | B+ | 87.8% | certified | 2026-04-23 |
| `setupAgent` | function | B+ | 87.8% | certified | 2026-04-23 |
| `cli_test.go` | file | B+ | 87.8% | certified | 2026-04-23 |
| `init` | function | B+ | 87.8% | certified | 2026-04-23 |
| `generateConfig` | function | B+ | 87.8% | certified | 2026-04-23 |
| `init` | function | B+ | 87.8% | certified | 2026-04-23 |
| `languagePolicy` | function | B | 85.6% | certified | 2026-04-23 |
| `main` | function | B+ | 87.8% | certified | 2026-04-23 |
| `detectCommit` | function | B+ | 87.8% | certified | 2026-04-23 |
| `detectRepoName` | function | B+ | 87.8% | certified | 2026-04-23 |
| `init` | function | B+ | 87.8% | certified | 2026-04-23 |
| `saveBadge` | function | B+ | 87.8% | certified | 2026-04-23 |
| `saveReportCard` | function | B+ | 87.8% | certified | 2026-04-23 |
| `init` | function | B+ | 87.8% | certified | 2026-04-23 |
| `init` | function | B+ | 87.8% | certified | 2026-04-23 |
| `init` | function | B+ | 87.8% | certified | 2026-04-23 |
| `version.go` | file | B+ | 87.8% | certified | 2026-04-23 |

<details>
<summary>languagePolicy — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 75.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 9 exceeds threshold 0

</details>

### `docs/` (4 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `README.md` | file | B | 86.7% | certified | 2026-04-23 |
| `architecture.md` | file | B | 86.7% | certified | 2026-04-23 |
| `policy-authoring.md` | file | B | 86.7% | certified | 2026-04-23 |
| `troubleshooting.md` | file | B+ | 87.8% | certified | 2026-04-23 |

### `docs/internal/` (2 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `PRD.md` | file | B | 83.9% | certified | 2026-04-23 |
| `STORIES.md` | file | B | 83.9% | certified | 2026-04-23 |

### `internal/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `integration_test.go` | file | B | 86.1% | certified | 2026-04-23 |

### `internal/agent/` (97 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `attribution_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Chat` | method | B+ | 87.8% | certified | 2026-04-23 |
| `CircuitBreaker` | class | B+ | 87.8% | certified | 2026-04-23 |
| `IsOpen` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Name` | method | B+ | 87.8% | certified | 2026-04-23 |
| `NewCircuitBreaker` | function | B+ | 87.8% | certified | 2026-04-23 |
| `AdaptiveMessages` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Chat` | method | B | 86.7% | certified | 2026-04-23 |
| `FallbackProvider` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ModelChain` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Name` | method | B+ | 87.8% | certified | 2026-04-23 |
| `NewFallbackProvider` | function | B+ | 87.8% | certified | 2026-04-23 |
| `NewModelChain` | function | B+ | 87.8% | certified | 2026-04-23 |
| `modelPinnedProvider` | class | B+ | 87.8% | certified | 2026-04-23 |
| `fallback_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `APIError` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Chat` | method | B | 86.1% | certified | 2026-04-23 |
| `Error` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Name` | method | B+ | 87.8% | certified | 2026-04-23 |
| `NewOpenRouterProvider` | function | B+ | 87.8% | certified | 2026-04-23 |
| `OpenRouterProvider` | class | B+ | 87.8% | certified | 2026-04-23 |
| `doRequest` | method | B | 86.7% | certified | 2026-04-23 |
| `isAuthError` | function | B+ | 87.8% | certified | 2026-04-23 |
| `isBudgetError` | function | B+ | 87.8% | certified | 2026-04-23 |
| `isRetryable` | function | B+ | 87.8% | certified | 2026-04-23 |
| `openrouter_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Coordinator` | class | B+ | 87.8% | certified | 2026-04-23 |
| `CoordinatorConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `NewCoordinator` | function | B+ | 87.8% | certified | 2026-04-23 |
| `NewPipeline` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Pipeline` | class | B+ | 87.8% | certified | 2026-04-23 |
| `PipelineConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ReviewUnit` | method | B | 86.7% | certified | 2026-04-23 |
| `Run` | method | B | 86.7% | certified | 2026-04-23 |
| `Stats` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Strategy` | class | B+ | 87.8% | certified | 2026-04-23 |
| `toResult` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Get` | method | B | 86.7% | certified | 2026-04-23 |
| `LoadPrompt` | function | B+ | 87.8% | certified | 2026-04-23 |
| `NewPromptRegistry` | function | B+ | 87.8% | certified | 2026-04-23 |
| `PromptRegistry` | class | B+ | 87.8% | certified | 2026-04-23 |
| `PromptTemplate` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Render` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Version` | method | B+ | 87.8% | certified | 2026-04-23 |
| `prompts_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Provider` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Allow` | method | B+ | 87.8% | certified | 2026-04-23 |
| `NewRateLimiter` | function | B+ | 87.8% | certified | 2026-04-23 |
| `RateLimiter` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Wait` | method | B+ | 87.8% | certified | 2026-04-23 |
| `refill` | method | B+ | 87.8% | certified | 2026-04-23 |
| `ratelimit_test.go` | file | B+ | 87.8% | certified | 2026-04-23 |
| `NewReviewer` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Review` | method | B | 85.0% | certified | 2026-04-23 |
| `ReviewInput` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ReviewResult` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Reviewer` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ToEvidence` | method | B+ | 87.8% | certified | 2026-04-23 |
| `joinModels` | function | B+ | 87.8% | certified | 2026-04-23 |
| `reviewer_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `ModelFor` | method | B+ | 87.8% | certified | 2026-04-23 |
| `NewRouter` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Router` | class | B+ | 87.8% | certified | 2026-04-23 |
| `directAssignment` | method | B | 86.7% | certified | 2026-04-23 |
| `router_test.go` | file | B+ | 87.8% | certified | 2026-04-23 |
| `DecisionResponse` | class | B+ | 87.8% | certified | 2026-04-23 |
| `PrescreenResponse` | class | B+ | 87.8% | certified | 2026-04-23 |
| `RemediationResponse` | class | B+ | 87.8% | certified | 2026-04-23 |
| `RemediationStep` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ScoringResponse` | class | B+ | 87.8% | certified | 2026-04-23 |
| `schemas_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Execute` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Name` | method | B+ | 87.8% | certified | 2026-04-23 |
| `NewPrescreenStage` | function | B+ | 87.8% | certified | 2026-04-23 |
| `NewReviewStage` | function | B+ | 87.8% | certified | 2026-04-23 |
| `NewScoringStage` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Stage` | class | B+ | 87.8% | certified | 2026-04-23 |
| `StageInput` | class | B+ | 87.8% | certified | 2026-04-23 |
| `StageResult` | class | B+ | 87.8% | certified | 2026-04-23 |
| `defaultScores` | function | B+ | 87.8% | certified | 2026-04-23 |
| `extractJSON` | function | B | 86.7% | certified | 2026-04-23 |
| `looseParseNeedsReview` | function | B+ | 87.8% | certified | 2026-04-23 |
| `prescreenStage` | class | B+ | 87.8% | certified | 2026-04-23 |
| `reviewStage` | class | B+ | 87.8% | certified | 2026-04-23 |
| `scoringStage` | class | B+ | 87.8% | certified | 2026-04-23 |
| `stage_test.go` | file | B | 83.9% | certified | 2026-04-23 |
| `ChatRequest` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ChatResponse` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Choice` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Content` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Message` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ModelConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ResponseFormat` | class | B+ | 87.8% | certified | 2026-04-23 |
| `String` | method | B+ | 87.8% | certified | 2026-04-23 |
| `TaskType` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Usage` | class | B+ | 87.8% | certified | 2026-04-23 |
| `types_test.go` | file | B | 86.7% | certified | 2026-04-23 |

### `internal/config/` (21 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Load` | function | B | 86.1% | certified | 2026-04-23 |
| `LoadFile` | function | B+ | 87.8% | certified | 2026-04-23 |
| `LoadFromDir` | function | B+ | 87.8% | certified | 2026-04-23 |
| `rawConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `validate` | function | B+ | 87.8% | certified | 2026-04-23 |
| `loader_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `FilterPolicyPacks` | function | B | 86.7% | certified | 2026-04-23 |
| `NewPolicyMatcher` | function | B+ | 87.8% | certified | 2026-04-23 |
| `LoadPolicyPack` | function | B+ | 87.8% | certified | 2026-04-23 |
| `LoadPolicyPacks` | function | B | 86.7% | certified | 2026-04-23 |
| `parseDimension` | function | B+ | 87.8% | certified | 2026-04-23 |
| `parsePolicyPack` | function | B | 86.7% | certified | 2026-04-23 |
| `parseSeverity` | function | B+ | 87.8% | certified | 2026-04-23 |
| `rawPolicyPack` | class | B+ | 87.8% | certified | 2026-04-23 |
| `rawPolicyRule` | class | B+ | 87.8% | certified | 2026-04-23 |
| `policy_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Error` | method | B+ | 87.8% | certified | 2026-04-23 |
| `ValidateConfig` | function | B | 86.7% | certified | 2026-04-23 |
| `ValidatePolicyPack` | function | B | 86.7% | certified | 2026-04-23 |
| `ValidationError` | class | B+ | 87.8% | certified | 2026-04-23 |
| `validator_test.go` | file | B | 86.7% | certified | 2026-04-23 |

### `internal/discovery/` (39 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `DetectLanguages` | function | B | 86.7% | certified | 2026-04-23 |
| `DetectedAdapters` | function | B+ | 87.8% | certified | 2026-04-23 |
| `LanguageInfo` | class | B+ | 87.8% | certified | 2026-04-23 |
| `buildLanguageList` | function | B+ | 87.8% | certified | 2026-04-23 |
| `detect_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `ChangedFiles` | function | B+ | 87.8% | certified | 2026-04-23 |
| `DetectMoves` | function | B+ | 87.8% | certified | 2026-04-23 |
| `FilterByPaths` | function | B | 86.7% | certified | 2026-04-23 |
| `FilterChanged` | function | B+ | 87.8% | certified | 2026-04-23 |
| `MovedFile` | class | B+ | 87.8% | certified | 2026-04-23 |
| `diff_test.go` | file | B+ | 87.8% | certified | 2026-04-23 |
| `GenericScanner` | class | B+ | 87.8% | certified | 2026-04-23 |
| `NewGenericScanner` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Scan` | method | B | 86.1% | certified | 2026-04-23 |
| `matchAny` | function | B | 86.7% | certified | 2026-04-23 |
| `GoAdapter` | class | B+ | 87.8% | certified | 2026-04-23 |
| `NewGoAdapter` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Scan` | method | B | 86.1% | certified | 2026-04-23 |
| `parseFile` | method | B | 86.7% | certified | 2026-04-23 |
| `go_adapter_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Diff` | function | B | 86.7% | certified | 2026-04-23 |
| `DiffResult` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Index` | class | B+ | 87.8% | certified | 2026-04-23 |
| `LoadIndex` | function | B+ | 87.8% | certified | 2026-04-23 |
| `NewIndex` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Save` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Units` | method | B+ | 87.8% | certified | 2026-04-23 |
| `indexEntry` | class | B+ | 87.8% | certified | 2026-04-23 |
| `index_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `DeduplicateFileLevel` | function | B | 86.7% | certified | 2026-04-23 |
| `Merge` | function | B | 86.7% | certified | 2026-04-23 |
| `Scanner` | class | B+ | 87.8% | certified | 2026-04-23 |
| `UnitList` | class | B+ | 87.8% | certified | 2026-04-23 |
| `scanner_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `NewTSAdapter` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Scan` | method | B | 86.1% | certified | 2026-04-23 |
| `TSAdapter` | class | B+ | 87.8% | certified | 2026-04-23 |
| `parseFile` | method | B | 86.7% | certified | 2026-04-23 |
| `ts_adapter_test.go` | file | B | 86.7% | certified | 2026-04-23 |

### `internal/domain/` (69 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `AgentConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `AnalyzerConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `CertificationMode` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Config` | class | B+ | 87.8% | certified | 2026-04-23 |
| `DefaultConfig` | function | B+ | 87.8% | certified | 2026-04-23 |
| `EnforcingConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ExpiryConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `IssueConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ModelAssignments` | class | B+ | 87.8% | certified | 2026-04-23 |
| `PolicyConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ProviderConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `RateLimitConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ScheduleConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ScopeConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `SignoffConfig` | class | B+ | 87.8% | certified | 2026-04-23 |
| `String` | method | B+ | 87.8% | certified | 2026-04-23 |
| `config_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `AllDimensions` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Dimension` | class | B+ | 87.8% | certified | 2026-04-23 |
| `DimensionScores` | class | B+ | 87.8% | certified | 2026-04-23 |
| `DimensionWeights` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Grade` | class | B+ | 87.8% | certified | 2026-04-23 |
| `GradeFromScore` | function | B | 86.7% | certified | 2026-04-23 |
| `String` | method | B+ | 87.8% | certified | 2026-04-23 |
| `WeightedAverage` | method | B | 86.7% | certified | 2026-04-23 |
| `dimension_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Evidence` | class | B+ | 87.8% | certified | 2026-04-23 |
| `EvidenceKind` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ParseSeverity` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Severity` | class | B+ | 87.8% | certified | 2026-04-23 |
| `String` | method | B+ | 87.8% | certified | 2026-04-23 |
| `init` | function | B+ | 87.8% | certified | 2026-04-23 |
| `evidence_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Duration` | method | B+ | 87.8% | certified | 2026-04-23 |
| `ExpiryFactors` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ExpiryWindow` | class | B+ | 87.8% | certified | 2026-04-23 |
| `IsExpired` | method | B+ | 87.8% | certified | 2026-04-23 |
| `RemainingAt` | method | B+ | 87.8% | certified | 2026-04-23 |
| `expiry_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Override` | class | B+ | 87.8% | certified | 2026-04-23 |
| `OverrideAction` | class | B+ | 87.8% | certified | 2026-04-23 |
| `String` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Validate` | method | B+ | 87.8% | certified | 2026-04-23 |
| `override_test.go` | file | B+ | 87.8% | certified | 2026-04-23 |
| `IsGlobal` | method | B+ | 87.8% | certified | 2026-04-23 |
| `PolicyPack` | class | B+ | 87.8% | certified | 2026-04-23 |
| `PolicyRule` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Violation` | class | B+ | 87.8% | certified | 2026-04-23 |
| `policy_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `CertificationRecord` | class | B+ | 87.8% | certified | 2026-04-23 |
| `IsPassing` | method | B+ | 87.8% | certified | 2026-04-23 |
| `ParseStatus` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Status` | class | B+ | 87.8% | certified | 2026-04-23 |
| `String` | method | B+ | 87.8% | certified | 2026-04-23 |
| `init` | function | B+ | 87.8% | certified | 2026-04-23 |
| `record_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Language` | method | B+ | 87.8% | certified | 2026-04-23 |
| `NewUnit` | function | B+ | 87.8% | certified | 2026-04-23 |
| `NewUnitID` | function | B+ | 87.8% | certified | 2026-04-23 |
| `ParseUnitID` | function | B+ | 87.8% | certified | 2026-04-23 |
| `ParseUnitType` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Path` | method | B+ | 87.8% | certified | 2026-04-23 |
| `String` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Symbol` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Unit` | class | B+ | 87.8% | certified | 2026-04-23 |
| `UnitID` | class | B+ | 87.8% | certified | 2026-04-23 |
| `UnitType` | class | B+ | 87.8% | certified | 2026-04-23 |
| `init` | function | B+ | 87.8% | certified | 2026-04-23 |
| `unit_test.go` | file | B | 86.1% | certified | 2026-04-23 |

### `internal/engine/` (10 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `CertifyUnit` | function | B+ | 87.8% | certified | 2026-04-23 |
| `pipeline_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Score` | function | B | 86.1% | certified | 2026-04-23 |
| `StatusFromScore` | function | B+ | 87.8% | certified | 2026-04-23 |
| `extractSummaryFloat` | function | B | 86.7% | certified | 2026-04-23 |
| `extractSummaryInt` | function | B | 86.1% | certified | 2026-04-23 |
| `scoreFromGitHistory` | function | B+ | 87.8% | certified | 2026-04-23 |
| `scoreFromMetrics` | function | B | 85.6% | certified | 2026-04-23 |
| `severityPenalty` | function | B+ | 87.8% | certified | 2026-04-23 |
| `scorer_test.go` | file | B | 85.6% | certified | 2026-04-23 |

<details>
<summary>scoreFromMetrics — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 1 exceeds threshold 0

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 75.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 1 exceeds threshold 0

</details>

### `internal/evidence/` (39 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Collector` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ComputeGoComplexity` | function | B | 86.1% | certified | 2026-04-23 |
| `ComputeSymbolMetrics` | function | B | 85.0% | certified | 2026-04-23 |
| `funcName` | function | B+ | 87.8% | certified | 2026-04-23 |
| `complexity_test.go` | file | B | 85.6% | certified | 2026-04-23 |
| `CollectAll` | method | B+ | 87.8% | certified | 2026-04-23 |
| `HasGoMod` | method | B+ | 87.8% | certified | 2026-04-23 |
| `HasPackageJSON` | method | B+ | 87.8% | certified | 2026-04-23 |
| `NewToolExecutor` | function | B+ | 87.8% | certified | 2026-04-23 |
| `ToolExecutor` | class | B+ | 87.8% | certified | 2026-04-23 |
| `runGitStats` | method | B+ | 87.8% | certified | 2026-04-23 |
| `runGoTest` | method | B | 86.7% | certified | 2026-04-23 |
| `runGoVet` | method | B+ | 87.8% | certified | 2026-04-23 |
| `runGolangciLint` | method | B+ | 87.8% | certified | 2026-04-23 |
| `ChurnRate` | method | B+ | 87.8% | certified | 2026-04-23 |
| `GitStats` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ParseGitLog` | function | B+ | 87.8% | certified | 2026-04-23 |
| `ToEvidence` | method | B+ | 87.8% | certified | 2026-04-23 |
| `git_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `LintFinding` | class | B+ | 87.8% | certified | 2026-04-23 |
| `LintResult` | class | B+ | 87.8% | certified | 2026-04-23 |
| `TestResult` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ToEvidence` | method | B+ | 87.8% | certified | 2026-04-23 |
| `lint_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `CodeMetrics` | class | B | 86.7% | certified | 2026-04-23 |
| `ComputeMetrics` | function | B | 85.0% | certified | 2026-04-23 |
| `ToEvidence` | method | B | 86.7% | certified | 2026-04-23 |
| `containsTodo` | function | B | 86.7% | certified | 2026-04-23 |
| `metrics_test.go` | file | B | 85.6% | certified | 2026-04-23 |
| `ParseCoverProfile` | function | B | 86.7% | certified | 2026-04-23 |
| `ParseGitLogWithAge` | function | B+ | 87.8% | certified | 2026-04-23 |
| `ParseGoTestJSON` | function | B | 86.7% | certified | 2026-04-23 |
| `ParseGoVet` | function | B | 86.7% | certified | 2026-04-23 |
| `ParseGolangciLintJSON` | function | B | 86.7% | certified | 2026-04-23 |
| `goTestEvent` | class | B+ | 87.8% | certified | 2026-04-23 |
| `golangciLintIssue` | class | B+ | 87.8% | certified | 2026-04-23 |
| `golangciLintOutput` | class | B+ | 87.8% | certified | 2026-04-23 |
| `simpleAtoi` | function | B+ | 87.8% | certified | 2026-04-23 |
| `runner_test.go` | file | B | 86.7% | certified | 2026-04-23 |

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 75.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 5 exceeds threshold 0

</details>

<details>
<summary>CodeMetrics — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 1 exceeds threshold 0

</details>

<details>
<summary>ComputeMetrics — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 80.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 4 exceeds threshold 0

</details>

<details>
<summary>ToEvidence — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 1 exceeds threshold 0

</details>

<details>
<summary>containsTodo — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 2 exceeds threshold 0

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 75.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 7 exceeds threshold 0

</details>

### `internal/expiry/` (2 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Calculate` | function | B | 86.7% | certified | 2026-04-23 |
| `calculator_test.go` | file | B | 86.7% | certified | 2026-04-23 |

### `internal/github/` (17 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `BuildIssueCloseCommand` | function | B+ | 87.8% | certified | 2026-04-23 |
| `BuildIssueCreateCommand` | function | B+ | 87.8% | certified | 2026-04-23 |
| `BuildIssueSearchCommand` | function | B+ | 87.8% | certified | 2026-04-23 |
| `BuildIssueUpdateCommand` | function | B+ | 87.8% | certified | 2026-04-23 |
| `FormatGroupedIssueBody` | function | B+ | 87.8% | certified | 2026-04-23 |
| `FormatIssueBody` | function | B+ | 87.8% | certified | 2026-04-23 |
| `FormatIssueTitle` | function | B+ | 87.8% | certified | 2026-04-23 |
| `issues_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `BuildPRCommentCommand` | function | B+ | 87.8% | certified | 2026-04-23 |
| `ComputeTrustDelta` | function | B | 86.1% | certified | 2026-04-23 |
| `FormatPRComment` | function | B | 86.1% | certified | 2026-04-23 |
| `TrustDelta` | class | B+ | 87.8% | certified | 2026-04-23 |
| `pr_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `GenerateNightlyWorkflow` | function | B+ | 87.8% | certified | 2026-04-23 |
| `GeneratePRWorkflow` | function | B+ | 87.8% | certified | 2026-04-23 |
| `GenerateWeeklyWorkflow` | function | B+ | 87.8% | certified | 2026-04-23 |
| `workflows_test.go` | file | B | 86.7% | certified | 2026-04-23 |

### `internal/override/` (9 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Apply` | function | B | 86.7% | certified | 2026-04-23 |
| `ApplyAll` | function | B+ | 87.8% | certified | 2026-04-23 |
| `applier_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `LoadDir` | function | B | 86.7% | certified | 2026-04-23 |
| `LoadFile` | function | B | 86.7% | certified | 2026-04-23 |
| `parseAction` | function | B+ | 87.8% | certified | 2026-04-23 |
| `rawOverride` | class | B+ | 87.8% | certified | 2026-04-23 |
| `rawOverrideFile` | class | B+ | 87.8% | certified | 2026-04-23 |
| `loader_test.go` | file | B | 86.7% | certified | 2026-04-23 |

### `internal/policy/` (14 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Evaluate` | function | B+ | 87.8% | certified | 2026-04-23 |
| `EvaluationResult` | class | B+ | 87.8% | certified | 2026-04-23 |
| `evaluateRule` | function | B+ | 87.8% | certified | 2026-04-23 |
| `extractComplexity` | function | B | 86.7% | certified | 2026-04-23 |
| `extractCoverage` | function | B+ | 87.8% | certified | 2026-04-23 |
| `extractMetric` | function | B | 85.0% | certified | 2026-04-23 |
| `extractTodoCount` | function | B | 85.0% | certified | 2026-04-23 |
| `evaluator_test.go` | file | B | 85.6% | certified | 2026-04-23 |
| `Match` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Matcher` | class | B+ | 87.8% | certified | 2026-04-23 |
| `NewMatcher` | function | B+ | 87.8% | certified | 2026-04-23 |
| `matchPath` | function | B | 86.7% | certified | 2026-04-23 |
| `matchesPack` | method | B | 86.7% | certified | 2026-04-23 |
| `matcher_test.go` | file | B | 86.7% | certified | 2026-04-23 |

<details>
<summary>extractMetric — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 80.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 2 exceeds threshold 0

</details>

<details>
<summary>extractTodoCount — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 80.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 7 exceeds threshold 0

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 75.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 2 exceeds threshold 0

</details>

### `internal/queue/` (17 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `BatchNext` | method | B | 86.7% | certified | 2026-04-23 |
| `Complete` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Enqueue` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Fail` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Item` | class | B+ | 87.8% | certified | 2026-04-23 |
| `ItemStatus` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Len` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Load` | function | B+ | 87.8% | certified | 2026-04-23 |
| `New` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Next` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Queue` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Reset` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Save` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Skip` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Stats` | class | B | 86.7% | certified | 2026-04-23 |
| `persistedQueue` | class | B+ | 87.8% | certified | 2026-04-23 |
| `queue_test.go` | file | B | 86.1% | certified | 2026-04-23 |

### `internal/record/` (17 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `AppendHistory` | method | B+ | 87.8% | certified | 2026-04-23 |
| `ListAll` | method | B | 86.7% | certified | 2026-04-23 |
| `Load` | method | B+ | 87.8% | certified | 2026-04-23 |
| `LoadHistory` | method | B | 86.7% | certified | 2026-04-23 |
| `NewStore` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Save` | method | B+ | 87.8% | certified | 2026-04-23 |
| `Store` | class | B+ | 87.8% | certified | 2026-04-23 |
| `dimensionsToMap` | function | B+ | 87.8% | certified | 2026-04-23 |
| `fromJSON` | function | B+ | 87.8% | certified | 2026-04-23 |
| `historyEntry` | class | B+ | 87.8% | certified | 2026-04-23 |
| `historyPathFor` | method | B+ | 87.8% | certified | 2026-04-23 |
| `mapToDimensions` | function | B+ | 87.8% | certified | 2026-04-23 |
| `parseGrade` | function | B+ | 87.8% | certified | 2026-04-23 |
| `pathFor` | method | B+ | 87.8% | certified | 2026-04-23 |
| `recordJSON` | class | B+ | 87.8% | certified | 2026-04-23 |
| `toJSON` | function | B+ | 87.8% | certified | 2026-04-23 |
| `store_test.go` | file | B | 86.7% | certified | 2026-04-23 |

### `internal/report/` (55 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Badge` | class | B+ | 87.8% | certified | 2026-04-23 |
| `BadgeMarkdown` | function | B+ | 87.8% | certified | 2026-04-23 |
| `FormatBadgeJSON` | function | B+ | 87.8% | certified | 2026-04-23 |
| `GenerateBadge` | function | B+ | 87.8% | certified | 2026-04-23 |
| `badgeColor` | function | B | 86.7% | certified | 2026-04-23 |
| `badgeMessage` | function | B+ | 87.8% | certified | 2026-04-23 |
| `badge_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `Card` | class | B+ | 87.8% | certified | 2026-04-23 |
| `FormatCardMarkdown` | function | B | 85.0% | certified | 2026-04-23 |
| `FormatCardText` | function | B | 85.0% | certified | 2026-04-23 |
| `GenerateCard` | function | B | 86.7% | certified | 2026-04-23 |
| `IssueCard` | class | B+ | 87.8% | certified | 2026-04-23 |
| `LanguageCard` | class | B+ | 87.8% | certified | 2026-04-23 |
| `buildLanguageCards` | function | B+ | 87.8% | certified | 2026-04-23 |
| `buildTopIssues` | function | B | 86.7% | certified | 2026-04-23 |
| `gradeEmoji` | function | B | 86.7% | certified | 2026-04-23 |
| `card_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `AreaSummary` | class | B+ | 87.8% | certified | 2026-04-23 |
| `Detailed` | function | B+ | 87.8% | certified | 2026-04-23 |
| `DetailedReport` | class | B+ | 87.8% | certified | 2026-04-23 |
| `FormatDetailedText` | function | B | 85.0% | certified | 2026-04-23 |
| `LanguageBreakdown` | class | B+ | 87.8% | certified | 2026-04-23 |
| `UnitSummary` | class | B+ | 87.8% | certified | 2026-04-23 |
| `computeDimensionAverages` | function | B+ | 87.8% | certified | 2026-04-23 |
| `computeLanguageBreakdowns` | function | B+ | 87.8% | certified | 2026-04-23 |
| `explainStatus` | function | B+ | 87.8% | certified | 2026-04-23 |
| `findExpiringSoon` | function | B | 86.7% | certified | 2026-04-23 |
| `findFailing` | function | B+ | 87.8% | certified | 2026-04-23 |
| `findHighestRisk` | function | B+ | 87.8% | certified | 2026-04-23 |
| `findRecurrentlyFailing` | function | B+ | 87.8% | certified | 2026-04-23 |
| `unitSummaryFrom` | function | B+ | 87.8% | certified | 2026-04-23 |
| `detailed_test.go` | file | B | 86.1% | certified | 2026-04-23 |
| `FormatFullMarkdown` | function | B+ | 87.8% | certified | 2026-04-23 |
| `FullReport` | class | B+ | 87.8% | certified | 2026-04-23 |
| `GenerateFullReport` | function | B+ | 87.8% | certified | 2026-04-23 |
| `LanguageDetail` | class | B+ | 87.8% | certified | 2026-04-23 |
| `UnitReport` | class | B+ | 87.8% | certified | 2026-04-23 |
| `buildLanguageDetail` | function | B | 86.7% | certified | 2026-04-23 |
| `dirOf` | function | B+ | 87.8% | certified | 2026-04-23 |
| `shortFile` | function | B+ | 87.8% | certified | 2026-04-23 |
| `sortedKeys` | function | B+ | 87.8% | certified | 2026-04-23 |
| `unitReportFrom` | function | B+ | 87.8% | certified | 2026-04-23 |
| `writeAllUnits` | function | B | 86.7% | certified | 2026-04-23 |
| `writeDimensionAverages` | function | B+ | 87.8% | certified | 2026-04-23 |
| `writeGradeDistribution` | function | B+ | 87.8% | certified | 2026-04-23 |
| `writeHeader` | function | B+ | 87.8% | certified | 2026-04-23 |
| `writeLanguageDetail` | function | B | 86.7% | certified | 2026-04-23 |
| `writeSummary` | function | B+ | 87.8% | certified | 2026-04-23 |
| `writeUnitDetails` | function | B | 86.7% | certified | 2026-04-23 |
| `full_test.go` | file | B | 86.7% | certified | 2026-04-23 |
| `FormatJSON` | function | B+ | 87.8% | certified | 2026-04-23 |
| `FormatText` | function | B+ | 87.8% | certified | 2026-04-23 |
| `Health` | function | B | 86.7% | certified | 2026-04-23 |
| `HealthReport` | class | B+ | 87.8% | certified | 2026-04-23 |
| `health_test.go` | file | B | 86.7% | certified | 2026-04-23 |

---

*446 units certified. Generated by [certify](https://github.com/iksnae/code-certification).*
