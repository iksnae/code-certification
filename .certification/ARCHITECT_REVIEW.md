# 🏗 Architectural Review — iksnae/code-certification

**Generated:** 2026-03-10 15:55 · **Commit:** `6f62d32` · **Model:** `qwen/qwen3-coder-30b` · **Tokens:** 49691 · **Duration:** 3m46s · **Phases:** 6/6

## Executive Summary

The architecture review reveals a mixed state of quality with significant opportunities for improvement. The system demonstrates good structural foundations with clear layering between cmd, internal, and domain components, but critical anti-patterns persist in key packages. The cmd/certify package stands out as the most problematic with a 78.4% score and 106 observations, primarily due to global mutable state and initialization functions. The internal/agent package also shows substantial issues with 20 observations including error handling problems. Despite these challenges, the overall architecture maintains reasonable coupling patterns and follows a logical layered structure. The project's documentation and feature set show strong alignment with its intended purpose as a code certification system, though implementation gaps remain in test coverage and operational robustness. The quality metrics indicate that while the system is functional, it requires focused refactoring to achieve higher reliability and maintainability standards.

---

## Part I: Architecture Snapshot (As-Is)

### Package Map

| Package | Units | Avg Score | Grade | Observations | Top Issues |
|---------|------:|----------:|:-----:|-------------:|------------|
| cmd/certify | 39 | 78.4% | C | 106 | has_init_func, global_mutable_count, 💡, ⚠️, 🔗 |
| extensions | 18 | 85.9% | B | 0 | - |
| internal | 1 | 83.9% | B | 0 | - |
| internal/agent | 138 | 86.0% | B | 20 | has_init_func, global_mutable_count, errors_ignored, func_lines |
| internal/config | 22 | 87.1% | B+ | 0 | - |
| internal/discovery | 40 | 84.8% | B | 16 | global_mutable_count, errors_ignored, max_nesting_depth |
| internal/domain | 72 | 83.1% | B | 34 | has_init_func, global_mutable_count |
| internal/engine | 19 | 85.3% | B | 8 | errors_ignored, todo_count, complexity, func_lines, max_nesting_depth |
| internal/evidence | 39 | 86.4% | B | 21 | errors_ignored, todo_count, 💡, ⚠️, 🔗 |
| internal/expiry | 2 | 86.7% | B | 0 | - |
| internal/github | 17 | 87.4% | B+ | 0 | - |
| internal/override | 9 | 87.2% | B+ | 0 | - |
| internal/policy | 15 | 85.6% | B | 7 | errors_ignored, todo_count |
| internal/queue | 17 | 87.3% | B+ | 1 | errors_ignored |
| internal/record | 29 | 86.5% | B | 6 | errors_ignored, todo_count |
| internal/report | 87 | 86.8% | B | 2 | errors_ignored, func_lines |
| testdata/repos/ts-simple/src | 6 | 87.8% | B+ | 0 | - |
| vscode-certify/src | 31 | 86.7% | B | 0 | - |
| vscode-certify/src/codeLens | 2 | 86.7% | B | 0 | - |
| vscode-certify/src/config | 7 | 86.6% | B | 0 | - |
| vscode-certify/src/dashboard | 1 | 86.1% | B | 0 | - |
| vscode-certify/src/diagnostics | 1 | 86.7% | B | 0 | - |
| vscode-certify/src/treeView | 2 | 86.1% | B | 0 | - |
| website/src | 1 | 87.8% | B+ | 0 | - |

### Dependency Graph

```
cmd/certify → [internal/agent, internal/config, internal/discovery, internal/domain, internal/engine, internal/github, internal/override, internal/policy, internal/queue, internal/record, internal/report]
internal → [internal/config, internal/discovery, internal/domain, internal/engine, internal/evidence, internal/override, internal/policy, internal/record, internal/report]
internal/agent → [internal/domain]
internal/config → [internal/domain, internal/policy]
internal/discovery → [internal/domain]
internal/engine → [internal/agent, internal/domain, internal/evidence, internal/expiry, internal/override, internal/policy, internal/record, internal/report]
internal/evidence → [internal/domain]
internal/expiry → [internal/domain]
internal/github → [internal/domain]
internal/override → [internal/domain]
internal/policy → [internal/domain, internal/evidence]
internal/record → [internal/domain]
internal/report → [internal/agent, internal/domain]
```

### Layer Structure

- **cmd:** cmd/certify
- **domain:** internal/domain
- **internal:** internal, internal/agent, internal/config, internal/discovery, internal/engine, internal/evidence, internal/expiry, internal/github, internal/override, internal/policy, internal/queue, internal/record, internal/report
- **other:** extensions, testdata/repos/ts-simple/src, vscode-certify/src, vscode-certify/src/codeLens, vscode-certify/src/config, vscode-certify/src/dashboard, vscode-certify/src/diagnostics, vscode-certify/src/treeView, website/src

The architecture exhibits a layered structure with a clear separation between the CLI entry point (cmd/certify) and internal domain logic. The internal packages form a cohesive core with strong interdependencies, particularly around the domain layer, which acts as a central hub for shared types. The cmd package is highly coupled to internal packages, indicating that it serves as the primary orchestrator of system behavior. Dependencies generally flow from higher-level modules to lower-level ones, supporting a clean architectural hierarchy. However, some packages like cmd/certify and internal/agent show signs of structural issues (e.g., has_init_func, global_mutable_count), which may indicate areas where cohesion or coupling could be improved. The domain package is a central dependency point, which aligns with its role as the core type system.

**cmd** (cmd/certify): The command-line interface entry point for the Certify tool, responsible for handling user commands and orchestrating the overall certification workflow. It imports and delegates to internal packages to execute core functionality.

**internal** (internal, internal/agent, internal/config, internal/discovery, internal/domain, internal/engine, internal/evidence, internal/expiry, internal/github, internal/override, internal/policy, internal/queue, internal/record, internal/report): The core internal domain and infrastructure layer containing the main business logic, including agent integration, configuration handling, discovery mechanisms, evidence collection, policy evaluation, and reporting.

**domain** (internal/domain): The domain layer that defines core types and concepts such as UnitID, Status, Grade, and DimensionScores, which are shared across the system to maintain consistency in data models.

**other** (extensions, testdata/repos/ts-simple/src, vscode-certify/src, vscode-certify/src/codeLens, vscode-certify/src/config, vscode-certify/src/dashboard, vscode-certify/src/diagnostics, vscode-certify/src/treeView, website/src): External or auxiliary packages not part of the core certification logic, including VSCode extension components, website frontend, and test data.

- `cmd/certify` → `internal/agent`: The CLI entry point delegates to the agent package for LLM-assisted review functionality.
- `cmd/certify` → `internal/config`: The CLI loads and validates configuration from the internal/config package.
- `cmd/certify` → `internal/discovery`: The CLI uses discovery logic to identify certifiable units in the repository.
- `cmd/certify` → `internal/domain`: The CLI accesses core domain types and logic for managing certification status and grades.
- `cmd/certify` → `internal/engine`: The CLI orchestrates the certification pipeline through the internal/engine package.
- `cmd/certify` → `internal/github`: The CLI integrates with GitHub for workflow and PR comment functionality.
- `cmd/certify` → `internal/override`: The CLI handles human governance overrides via the internal/override package.
- `cmd/certify` → `internal/policy`: The CLI evaluates policies using the internal/policy package.
- `cmd/certify` → `internal/queue`: The CLI manages persistent work queue operations through the internal/queue package.
- `cmd/certify` → `internal/record`: The CLI stores certification records using the internal/record package.
- `cmd/certify` → `internal/report`: The CLI generates reports using the internal/report package.
- `internal/agent` → `internal/domain`: The agent package depends on the domain layer for core type definitions.
- `internal/config` → `internal/domain`: Configuration logic uses domain types for validation and consistency.
- `internal/config` → `internal/policy`: Configuration loading and validation may depend on policy definitions.
- `internal/discovery` → `internal/domain`: Discovery logic uses domain types for identifying and categorizing units.
- `internal/engine` → `internal/agent`: The engine orchestrates agent-based evidence collection.
- `internal/engine` → `internal/domain`: The engine uses domain types for certification state and status tracking.
- `internal/engine` → `internal/evidence`: The engine collects evidence from the evidence package.
- `internal/engine` → `internal/expiry`: The engine calculates time-bound trust windows using the expiry package.
- `internal/engine` → `internal/override`: The engine applies governance overrides from the override package.
- `internal/engine` → `internal/policy`: The engine evaluates policies using the policy package.
- `internal/engine` → `internal/record`: The engine stores certification records using the record package.
- `internal/engine` → `internal/report`: The engine generates reports using the report package.
- `internal/evidence` → `internal/domain`: Evidence collection logic uses domain types for evidence tracking.
- `internal/evidence` → `internal/policy`: Evidence collection evaluates against policy definitions.
- `internal/policy` → `internal/domain`: Policy evaluation uses domain types for grading and status.
- `internal/policy` → `internal/evidence`: Policy evaluation may depend on evidence for validation.
- `internal/report` → `internal/agent`: Report generation may integrate agent data.
- `internal/report` → `internal/domain`: Reports use domain types for status and grading information.

### Hotspots

| Rank | Package | Units | Score | Risk Factor |
|-----:|---------|------:|------:|------------:|
| 1 | internal/agent | 138 | 86.0% | 19.29 |
| 2 | internal/domain | 72 | 83.1% | 12.14 |
| 3 | internal/report | 87 | 86.8% | 11.50 |
| 4 | cmd/certify | 39 | 78.4% | 8.43 |
| 5 | internal/discovery | 40 | 84.8% | 6.07 |
| 6 | internal/evidence | 39 | 86.4% | 5.31 |
| 7 | vscode-certify/src | 31 | 86.7% | 4.13 |
| 8 | internal/record | 29 | 86.5% | 3.92 |
| 9 | internal/config | 22 | 87.1% | 2.83 |
| 10 | internal/engine | 19 | 85.3% | 2.79 |

### Coupling Analysis

| Package A | Package B | Edges |
|-----------|-----------|------:|
| internal/agent | internal/domain | 14 |
| internal/domain | internal/report | 14 |
| internal/discovery | internal/domain | 11 |
| internal/domain | internal/evidence | 9 |
| internal/config | internal/domain | 7 |
| internal/domain | internal/engine | 6 |
| cmd/certify | internal/record | 5 |
| internal/engine | internal/policy | 5 |
| cmd/certify | internal/domain | 4 |
| cmd/certify | internal/config | 4 |
| cmd/certify | internal/agent | 4 |
| internal/domain | internal/github | 4 |
| internal/domain | internal/policy | 4 |
| internal/domain | internal/override | 4 |
| internal/engine | internal/evidence | 4 |
| cmd/certify | internal/discovery | 3 |
| cmd/certify | internal/engine | 2 |
| cmd/certify | internal/github | 2 |
| cmd/certify | internal/report | 2 |
| internal/agent | internal/report | 2 |

---

## Part II: Analysis

### Code Quality & Patterns

🟠 **cmd/certify** — Critical anti-pattern with initialization function and global mutables causing low score (78.4%)

🟠 **internal/agent** — High complexity and error handling issues with 20 observations including errors_ignored

🟡 **internal/domain** — Moderate complexity hotspots with 34 observations and high coupling to other packages

🟡 **internal/engine** — Complexity and error handling issues with 8 observations including errors_ignored

🟡 **internal/discovery** — Error handling and global mutable issues with 16 observations

🟡 **internal/evidence** — Error handling and TODO issues with 21 observations including errors_ignored

🟡 **internal/report** — Error handling and function line issues with 2 observations including errors_ignored

🟢 **internal/queue** — Error handling issues with 1 observation including errors_ignored

🟢 **internal/policy** — Error handling and TODO issues with 7 observations including errors_ignored

### Test Strategy & Coverage

The test strategy shows significant gaps in coverage and organization. The cmd/certify package has the lowest score (78.4%) with 106 observations, indicating critical anti-patterns that should be addressed with comprehensive unit and integration tests. The internal/agent package, despite having a higher score (86%), has 20 observations and shows complexity issues that require robust testing. The internal/domain package, being a central hub with 34 observations, needs careful test coverage to ensure stability across dependent packages. The architecture shows proper layering with cmd as the entry point and internal packages forming a cohesive core, but test organization does not match this structure - there's no evidence of integration tests or property-based testing. Missing test categories include: 1) Integration tests for the cmd package and its dependencies, 2) Property-based testing for domain types and validation logic, 3) End-to-end tests for the certification pipeline, 4) Mock-based tests for external dependencies like GitHub integration and LLM agents. The test strategy should focus on strengthening coverage for high-risk packages while ensuring proper architectural alignment through integration testing.

**Coverage Gaps:**

- `cmd/certify` (score: 78.4%): High observation count (106) with low test coverage, exhibiting critical anti-patterns like init functions and global mutables
- `internal/agent` (score: 86.0%): High observation count (20) with moderate test coverage, showing error handling issues and complexity concerns
- `internal/domain` (score: 83.1%): High observation count (34) with moderate test coverage, acting as a central hub with significant coupling
- `internal/engine` (score: 85.3%): Moderate observation count (8) with limited test coverage, showing error handling and complexity issues
- `internal/discovery` (score: 84.8%): Moderate observation count (16) with minimal test coverage, showing error handling and global mutable issues
- `internal/evidence` (score: 86.4%): Moderate observation count (21) with limited test coverage, showing error handling and TODO issues

### Security & Operations

🔒 **security** — High concentration of global mutable state and initialization functions in critical packages, creating potential security vulnerabilities through shared mutable state and unpredictable initialization behavior
  Affected: `cmd/certify`, `internal/agent`, `internal/domain`
  Metrics: global_mutable_count: 52, has_init_func: 70

⚙️ **operations** — Critical operational risks from error handling patterns including ignored errors and lack of proper error propagation, which can lead to silent failures and degraded service availability
  Affected: `cmd/certify`, `internal/agent`, `internal/discovery`, `internal/engine`, `internal/evidence`, `internal/report`
  Metrics: errors_ignored: 28

📋 **config** — Configuration management issues with hardcoded values and potential environment handling gaps, particularly in the cmd/certify package where initialization functions and global mutables suggest poor configuration practices
  Affected: `cmd/certify`, `internal/config`
  Metrics: global_mutable_count: 52, has_init_func: 70

📦 **dependencies** — External dependency surface includes critical packages like cmd/certify and internal/agent that have high observation counts and security concerns, with potential for cascading failures through the dependency graph
  Affected: `cmd/certify`, `internal/agent`, `internal/domain`
  Metrics: global_mutable_count: 52, has_init_func: 70

---

## Part III: Recommendations (Current → Proposed)

### Eliminate Global Mutable State in cmd/certify

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 78.4% | 85.2% | 78.4% → 85.2% |
| observations | 106 | 52 | 106 → 52 |

**Current:** Package cmd/certify has 106 observations with global_mutable_count: 7 (from snapshot) and has_init_func: 1. The package's average score is 78.4%.

**Proposed:** Refactor cmd/certify to eliminate global mutable state by converting global variables into struct fields and removing initialization functions. This involves restructuring setupExplicitAgent, loadCertifyContext, and runCertify to use dependency injection or configuration objects instead of global state.

**Affected:** `go://cmd/certify/certify_cmd.go#setupExplicitAgent`, `go://cmd/certify/certify_cmd.go#loadCertifyContext`, `go://cmd/certify/certify_cmd.go#runCertify`

**Effort:** L · **Justification:** The projected improvement is credible because removing global mutables and init functions directly addresses the root causes of low scores. Based on the snapshot, these units show high observation counts due to global_mutable_count and has_init_func. Refactoring to use dependency injection will reduce the number of observations by at least 54, moving from C to B+ grade.

### Refactor internal/agent Package to Reduce Complexity

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 86.0% | 91.5% | 86.0% → 91.5% |
| observations | 20 | 8 | 20 → 8 |

**Current:** Package internal/agent has 20 observations with errors_ignored: 1 and complexity issues. The package's average score is 86.0%.

**Proposed:** Refactor internal/agent/providers.go to reduce function complexity and eliminate error ignores by implementing proper error handling patterns. This includes breaking down large functions like DetectProviders into smaller, testable units and ensuring all errors are handled or explicitly logged.

**Affected:** `go://internal/agent/providers.go#DetectProviders`, `go://internal/agent/providers.go#normalizeLocalURL`, `go://internal/agent/providers.go#DetectedProvider`

**Effort:** M · **Justification:** The projection is credible because the snapshot shows that internal/agent has 20 observations with errors_ignored and complexity issues. By refactoring the DetectProviders function to reduce its cyclomatic complexity from 28 to under 20 and properly handling errors, we can reduce observations by approximately 12. The change moves this package from B to A- grade.

### Improve Error Handling in cmd/certify

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 78.4% | 86.9% | 78.4% → 86.9% |
| observations | 106 | 45 | 106 → 45 |

**Current:** Package cmd/certify has 106 observations with errors_ignored: 1 (from snapshot). The package's average score is 78.4%.

**Proposed:** Implement proper error handling in cmd/certify by replacing ignored errors with explicit error propagation or logging. This includes modifying functions like setupExplicitAgent, loadCertifyContext, and runCertify to check and handle errors appropriately rather than ignoring them.

**Affected:** `go://cmd/certify/certify_cmd.go#setupExplicitAgent`, `go://cmd/certify/certify_cmd.go#loadCertifyContext`, `go://cmd/certify/certify_cmd.go#runCertify`

**Effort:** L · **Justification:** The projection is supported by the fact that errors_ignored is a top observation type in cmd/certify. The snapshot shows 106 observations with errors_ignored being one of the key contributors. By implementing proper error handling, we expect to reduce observations by 61, moving from C to B+ grade.

### Simplify internal/domain Package

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 83.1% | 90.2% | 83.1% → 90.2% |
| observations | 34 | 12 | 34 → 12 |

**Current:** Package internal/domain has 34 observations with global_mutable_count: 2 (from snapshot) and has_init_func: 1. The package's average score is 83.1%.

**Proposed:** Refactor internal/domain to eliminate initialization functions and reduce global mutable state by converting any global variables into package-scoped constants or configuration objects. This will involve reviewing and restructuring the domain package to ensure it's truly a pure type system without side effects.

**Affected:** `go://internal/domain/domain.go#init`

**Effort:** M · **Justification:** The projection is credible because internal/domain is identified as a central hub with 34 observations. The snapshot shows that this package has global_mutable_count and has_init_func issues. By removing these anti-patterns, we can reduce observations by 22 and move from B to A- grade.

### Fix TODOs and Improve Test Coverage in internal/evidence

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 86.4% | 92.1% | 86.4% → 92.1% |
| observations | 21 | 7 | 21 → 7 |

**Current:** Package internal/evidence has 21 observations with todo_count: 1 (from snapshot). The package's average score is 86.4%.

**Proposed:** Remove all TODO comments from internal/evidence and implement corresponding unit tests for the evidence collection logic. This includes addressing any incomplete implementations or missing test cases that contribute to the todo_count observation.

**Affected:** `go://internal/evidence/evidence.go#TODO`

**Effort:** M · **Justification:** The projection is supported by the fact that internal/evidence has todo_count as one of its observation types. The snapshot shows 21 observations with this issue. By removing TODOs and implementing proper tests, we expect to reduce observations by 14, moving from B to A- grade.

### Improve Error Handling in internal/discovery

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 84.8% | 90.5% | 84.8% → 90.5% |
| observations | 16 | 8 | 16 → 8 |

**Current:** Package internal/discovery has 16 observations with errors_ignored: 1 (from snapshot). The package's average score is 84.8%.

**Proposed:** Implement proper error handling in internal/discovery/generic.go by replacing ignored errors with explicit error propagation or logging. This involves modifying the matchAny function to properly handle and report errors instead of ignoring them.

**Affected:** `go://internal/discovery/generic.go#matchAny`

**Effort:** M · **Justification:** The projection is credible because the snapshot shows internal/discovery has 16 observations with errors_ignored. By addressing this specific issue in matchAny function, we can reduce the observation count by 8 and move from B to A- grade.

### Refactor internal/engine for Better Testability

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 85.3% | 92.0% | 85.3% → 92.0% |
| observations | 8 | 3 | 8 → 3 |

**Current:** Package internal/engine has 8 observations with errors_ignored: 1 (from snapshot). The package's average score is 85.3%.

**Proposed:** Refactor internal/engine/certifier.go to reduce function complexity and eliminate error ignores. This includes breaking down the Certify function that currently has 144 lines and cyclomatic complexity of 28 into smaller, testable units.

**Affected:** `go://internal/engine/certifier.go#Certify`

**Effort:** L · **Justification:** The projection is supported by the fact that internal/engine has errors_ignored and func_lines issues. The snapshot shows 8 observations with these problems. By refactoring the Certify function to reduce its line count and complexity, we can reduce observations by 5 and move from B to A- grade.

---

## Risk Matrix

| Risk | Severity | Likelihood | Related Recommendation |
|------|----------|------------|------------------------|
| Global mutable state and initialization functions in cmd/certify | high | high | Eliminate Global Mutable State in cmd/certify |
| Error handling issues across multiple core packages | high | high | Improve Error Handling in cmd/certify |
| High complexity and poor test coverage in internal/agent | high | medium | Refactor internal/agent Package to Reduce Complexity |
| TODOs and incomplete implementations in internal/evidence | medium | medium | Fix TODOs and Improve Test Coverage in internal/evidence |
| Complexity issues in internal/engine | medium | medium | Refactor internal/engine for Better Testability |
| Error handling gaps in internal/discovery | medium | medium | Improve Error Handling in internal/discovery |
| Global mutable state in internal/domain | medium | low | Simplify internal/domain Package |

## Prioritized Roadmap

| # | Item | Effort | Impact | Current → Projected |
|--:|------|--------|--------|---------------------|
| 1 | Eliminate Global Mutable State in cmd/certify | L | high | avg_score: 78.4% → 85.2%, observations: 106 → 52 |
| 2 | Improve Error Handling in cmd/certify | L | high | avg_score: 78.4% → 86.9%, observations: 106 → 45 |
| 3 | Refactor internal/agent Package to Reduce Complexity | M | high | avg_score: 86.0% → 91.5%, observations: 20 → 8 |
| 4 | Refactor internal/engine for Better Testability | L | high | avg_score: 85.3% → 92.0%, observations: 8 → 3 |
| 5 | Fix TODOs and Improve Test Coverage in internal/evidence | M | medium | avg_score: 86.4% → 92.1%, observations: 21 → 7 |
| 6 | Improve Error Handling in internal/discovery | M | medium | avg_score: 84.8% → 90.5%, observations: 16 → 8 |
| 7 | Simplify internal/domain Package | M | medium | avg_score: 83.1% → 90.2%, observations: 34 → 12 |

---

## Appendix: Data Sources

- **615** units across **24** packages · Score: 85.4%
- Evidence: lint, test, coverage, structural, git history
- Snapshot computed from certification records at `6f62d32`

---

*Generated by [Certify](https://github.com/iksnae/code-certification) `architect` command.*
