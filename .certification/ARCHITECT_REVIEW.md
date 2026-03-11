# 🏗 Architectural Review — iksnae/code-certification

**Generated:** 2026-03-10 21:02 · **Commit:** `68eb2b49` · **Model:** `qwen/qwen3-coder-30b` · **Tokens:** 48912 · **Duration:** 1m57s · **Phases:** 6/6

## Executive Summary

The Certify code certification system demonstrates a well-structured layered architecture with clear separation of concerns, yet faces significant challenges in code quality and maintainability. The system's architecture follows a logical command-internal-domain structure with cmd/certify as the primary entry point, internally managing core certification logic through packages like internal/agent, internal/engine, and internal/domain. However, analysis reveals critical issues including high global mutable state in the agent package (4 observations exceeding threshold of 2), os.Exit calls in CLI commands that interfere with graceful error handling, and incomplete implementation markers (TODOs) in evidence and policy packages. These issues collectively impact the system's reliability, testability, and maintainability, with the agent package being the highest risk area. The architecture snapshot shows 22 A-grade packages and 669 A--grade packages, indicating generally strong code quality, but the presence of hotspots like internal/agent (risk factor 15.35) and internal/report (risk factor 9.43) requires immediate attention to prevent cascading failures in the certification pipeline.

---

## Part I: Architecture Snapshot (As-Is)

### Package Map

| Package | Units | Avg Score | Grade | Observations | Top Issues |
|---------|------:|----------:|:-----:|-------------:|------------|
| cmd/certify | 54 | 91.4% | A- | 2 | global_mutable_count, os_exit_calls |
| extensions | 18 | 92.1% | A- | 0 | - |
| internal | 1 | 91.3% | A- | 0 | - |
| internal/agent | 187 | 91.8% | A- | 5 | global_mutable_count |
| internal/config | 22 | 91.7% | A- | 0 | - |
| internal/discovery | 40 | 91.1% | A- | 0 | - |
| internal/domain | 69 | 92.4% | A- | 0 | - |
| internal/engine | 20 | 91.5% | A- | 0 | - |
| internal/evidence | 68 | 91.3% | A- | 4 | todo_count, max_nesting_depth |
| internal/expiry | 2 | 91.9% | A- | 0 | - |
| internal/github | 17 | 92.1% | A- | 0 | - |
| internal/override | 9 | 91.8% | A- | 0 | - |
| internal/policy | 15 | 91.5% | A- | 1 | todo_count |
| internal/queue | 17 | 92.2% | A- | 0 | - |
| internal/record | 29 | 91.9% | A- | 0 | - |
| internal/report | 108 | 91.3% | A- | 0 | - |
| internal/workspace | 19 | 92.0% | A- | 0 | - |
| testdata/repos/ts-simple/src | 6 | 93.3% | A | 0 | - |
| vscode-certify/src | 33 | 92.5% | A- | 0 | - |
| vscode-certify/src/codeLens | 2 | 93.3% | A | 0 | - |
| vscode-certify/src/config | 7 | 93.2% | A | 0 | - |
| vscode-certify/src/dashboard | 1 | 92.5% | A- | 0 | - |
| vscode-certify/src/diagnostics | 1 | 93.3% | A | 0 | - |
| vscode-certify/src/treeView | 2 | 92.5% | A- | 0 | - |
| website/src | 1 | 93.3% | A | 0 | - |

### Dependency Graph

```
cmd/certify → [internal/agent, internal/config, internal/discovery, internal/domain, internal/engine, internal/github, internal/override, internal/policy, internal/queue, internal/record, internal/report, internal/workspace]
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
internal/workspace → [internal/domain, internal/record, internal/report]
```

### Layer Structure

- **cmd:** cmd/certify
- **domain:** internal/domain
- **internal:** internal, internal/agent, internal/config, internal/discovery, internal/engine, internal/evidence, internal/expiry, internal/github, internal/override, internal/policy, internal/queue, internal/record, internal/report, internal/workspace
- **other:** extensions, testdata/repos/ts-simple/src, vscode-certify/src, vscode-certify/src/codeLens, vscode-certify/src/config, vscode-certify/src/dashboard, vscode-certify/src/diagnostics, vscode-certify/src/treeView, website/src

The architecture follows a layered structure with clear separation between command, internal core, and domain layers. The dependency graph shows that the cmd layer is the entry point and imports most internal packages, while internal packages depend on each other in a structured way with the domain layer at the center. The domain layer is the most central and widely used, indicating good cohesion in the core concepts. There are no apparent violations of dependency direction (e.g., domain importing cmd), which suggests a well-structured architecture with minimal cross-layer coupling. The internal packages show moderate coupling, particularly between agent, report, and domain modules, which aligns with their functional interdependencies. The overall architecture supports a clean separation of concerns and promotes maintainability through well-defined layers.

**Command Layer** (cmd/certify): The command layer serves as the CLI entry point for the Certify tool, handling user interactions and orchestrating the execution of certification workflows. It imports and delegates to internal packages to perform core functionality like configuration, discovery, domain logic, and reporting.

**Internal Core Layer** (internal, internal/agent, internal/config, internal/discovery, internal/engine, internal/evidence, internal/expiry, internal/github, internal/override, internal/policy, internal/queue, internal/record, internal/report, internal/workspace): The internal core layer encapsulates the main business logic and domain-specific functionality of the Certify system. It includes modules for agent integration, configuration management, unit discovery, policy evaluation, evidence collection, report generation, and more. This layer is the heart of the certification engine.

**Domain Layer** (internal/domain): The domain layer defines the core types and abstractions that are shared across the system, such as UnitID, Status, Grade, and DimensionScores. It represents the fundamental concepts of the certification domain and is imported by most internal packages.

**Other Layers** (extensions, testdata/repos/ts-simple/src, vscode-certify/src, vscode-certify/src/codeLens, vscode-certify/src/config, vscode-certify/src/dashboard, vscode-certify/src/diagnostics, vscode-certify/src/treeView, website/src): The other layers represent supporting components such as VSCode extension modules, website content, and test data. These are not part of the core certification engine but support related tools or documentation.

- `cmd/certify` → `internal/agent`: The CLI command layer delegates to the agent module for LLM-assisted review via OpenRouter.
- `cmd/certify` → `internal/config`: The CLI command layer imports configuration handling to load and validate settings.
- `cmd/certify` → `internal/discovery`: The CLI command layer uses discovery logic to identify certifiable code units.
- `cmd/certify` → `internal/domain`: The CLI command layer accesses domain types for core certification concepts like UnitID and Status.
- `cmd/certify` → `internal/engine`: The CLI command layer orchestrates the certification pipeline through the engine module.
- `cmd/certify` → `internal/github`: The CLI command layer integrates with GitHub for workflow and PR comment features.
- `cmd/certify` → `internal/override`: The CLI command layer supports human governance overrides.
- `cmd/certify` → `internal/policy`: The CLI command layer evaluates policies as part of certification.
- `cmd/certify` → `internal/queue`: The CLI command layer manages persistent work queues for certification tasks.
- `cmd/certify` → `internal/record`: The CLI command layer stores certification records in JSON format with history.
- `cmd/certify` → `internal/report`: The CLI command layer generates report cards and badges.
- `cmd/certify` → `internal/workspace`: The CLI command layer manages workspace-specific configurations and record storage.
- `internal` → `internal/config`: Internal modules access configuration handling for loading and validation.
- `internal` → `internal/discovery`: Internal modules use discovery logic to identify certifiable units.
- `internal` → `internal/domain`: Internal modules rely on domain types for core certification concepts.
- `internal` → `internal/engine`: Internal modules integrate with the engine for certification pipeline execution.
- `internal` → `internal/evidence`: Internal modules collect evidence for certification evaluation.
- `internal` → `internal/override`: Internal modules support human governance overrides.
- `internal` → `internal/policy`: Internal modules evaluate policies for certification.
- `internal` → `internal/record`: Internal modules store certification records.
- `internal` → `internal/report`: Internal modules generate reports for certification.
- `internal/agent` → `internal/domain`: The agent module uses domain types for core certification concepts.
- `internal/config` → `internal/domain`: Configuration handling accesses domain types for core certification concepts.
- `internal/config` → `internal/policy`: Configuration handling evaluates policies for certification.
- `internal/discovery` → `internal/domain`: Discovery logic uses domain types for core certification concepts.
- `internal/engine` → `internal/agent`: The engine module delegates to the agent for LLM-assisted review.
- `internal/engine` → `internal/domain`: The engine module uses domain types for core certification concepts.
- `internal/engine` → `internal/evidence`: The engine module collects evidence for certification evaluation.
- `internal/engine` → `internal/expiry`: The engine module calculates time-bound trust windows.
- `internal/engine` → `internal/override`: The engine module supports human governance overrides.
- `internal/engine` → `internal/policy`: The engine module evaluates policies for certification.
- `internal/engine` → `internal/record`: The engine module stores certification records.
- `internal/engine` → `internal/report`: The engine module generates reports for certification.
- `internal/evidence` → `internal/domain`: Evidence collectors use domain types for core certification concepts.
- `internal/evidence` → `internal/policy`: Evidence collectors evaluate policies for certification.
- `internal/expiry` → `internal/domain`: Expiry logic uses domain types for core certification concepts.
- `internal/github` → `internal/domain`: GitHub integration uses domain types for core certification concepts.
- `internal/override` → `internal/domain`: Override handling uses domain types for core certification concepts.
- `internal/policy` → `internal/domain`: Policy evaluation uses domain types for core certification concepts.
- `internal/policy` → `internal/evidence`: Policy evaluation collects evidence for certification.
- `internal/record` → `internal/domain`: Record storage uses domain types for core certification concepts.
- `internal/report` → `internal/agent`: Report generation uses agent data for LLM-assisted review.
- `internal/report` → `internal/domain`: Report generation uses domain types for core certification concepts.
- `internal/workspace` → `internal/domain`: Workspace management uses domain types for core certification concepts.
- `internal/workspace` → `internal/record`: Workspace management stores certification records.
- `internal/workspace` → `internal/report`: Workspace management generates reports for certification.

### Hotspots

| Rank | Package | Units | Score | Risk Factor |
|-----:|---------|------:|------:|------------:|
| 1 | internal/agent | 187 | 91.8% | 15.35 |
| 2 | internal/report | 108 | 91.3% | 9.43 |
| 3 | internal/evidence | 68 | 91.3% | 5.94 |
| 4 | internal/domain | 69 | 92.4% | 5.23 |
| 5 | cmd/certify | 54 | 91.4% | 4.66 |
| 6 | internal/discovery | 40 | 91.1% | 3.55 |
| 7 | vscode-certify/src | 33 | 92.5% | 2.48 |
| 8 | internal/record | 29 | 91.9% | 2.35 |
| 9 | internal/config | 22 | 91.7% | 1.82 |
| 10 | internal/engine | 20 | 91.5% | 1.69 |

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
| cmd/certify | internal/workspace | 5 |
| internal/engine | internal/policy | 5 |
| cmd/certify | internal/domain | 4 |
| cmd/certify | internal/config | 4 |
| cmd/certify | internal/agent | 4 |
| internal/domain | internal/github | 4 |
| internal/domain | internal/policy | 4 |
| internal/domain | internal/override | 4 |
| internal/engine | internal/evidence | 4 |
| cmd/certify | internal/discovery | 3 |
| cmd/certify | internal/report | 2 |
| cmd/certify | internal/engine | 2 |
| cmd/certify | internal/github | 2 |

---

## Part II: Analysis

### Code Quality & Patterns

🟠 **internal/agent** — High global mutable state usage with 4 observations exceeding threshold of 2, indicating potential concurrency issues and code maintainability risks

🟡 **cmd/certify** — Contains os.Exit calls which can interfere with graceful shutdown and error handling in CLI applications

🟡 **internal/evidence** — Contains TODO markers and excessive nesting depth, suggesting incomplete implementation or code smells in structural analysis

🟢 **internal/policy** — Contains TODO markers indicating incomplete policy evaluation or documentation in the policy module

🟡 **internal/engine** — High coupling with internal/agent and internal/domain packages, representing a potential complexity hotspot

### Test Strategy & Coverage

The test strategy shows signs of incomplete coverage, particularly for high-risk packages with moderate to high coupling. The architecture follows layered structure but lacks comprehensive integration testing for critical data flows between cmd layer and internal packages. Missing test categories include: 1) Integration tests for CLI command flows (cmd/certify), 2) Property-based testing for state-dependent scenarios in internal/agent, 3) End-to-end integration tests for the certification pipeline (engine -> report -> record), and 4) Mock-based testing for external dependencies like GitHub integration. The current approach focuses on unit-level coverage but misses critical cross-layer interactions that could lead to runtime failures. Test organization appears to follow the package structure but lacks proper categorization for different test types, with most tests likely being unit-focused rather than integration or property-based testing.

**Coverage Gaps:**

- `internal/agent` (score: 91.8%): High global mutable state usage (4 observations) and moderate coupling with internal/domain (14 edges) suggest weak test coverage for concurrent and state-dependent scenarios
- `cmd/certify` (score: 91.4%): Contains os.Exit calls (1 observation) and global mutable state (2 observations) indicating potential missing integration tests for CLI behavior and graceful error handling
- `internal/evidence` (score: 91.3%): Contains TODO markers (4 observations) and max nesting depth (1 observation) suggesting incomplete test coverage for structural analysis and edge cases
- `internal/report` (score: 91.3%): Low observation count (0) but high risk factor (9.43) and dependency on internal/agent suggests missing integration tests for report generation with agent data
- `internal/engine` (score: 91.5%): High coupling with internal/agent (14 edges) and internal/domain (14 edges) with no observations indicates potential missing integration tests for complex workflow scenarios

### Security & Operations

🔒 **security** — High global mutable state usage in internal/agent package with 4 observations exceeding threshold of 2, indicating potential concurrency issues and code maintainability risks
  Affected: `internal/agent`
  Metrics: global_mutable_count: 4

🔒 **security** — Global mutable state in cmd/certify package with 2 observations exceeding threshold of 0, indicating potential concurrency issues and code maintainability risks
  Affected: `cmd/certify`
  Metrics: global_mutable_count: 2

⚙️ **operations** — os.Exit calls in cmd/certify package which can interfere with graceful shutdown and error handling in CLI applications
  Affected: `cmd/certify`
  Metrics: os_exit_calls: 1

📋 **config** — Hardcoded values and environment handling issues in cmd/certify package with global mutable state and os.Exit calls
  Affected: `cmd/certify`
  Metrics: global_mutable_count: 2, os_exit_calls: 1

📦 **dependencies** — External dependency surface includes GitHub integration and LLM agent modules that may introduce security risks through external API calls
  Affected: `internal/github`, `internal/agent`
  Metrics: dependency_count: 2

---

## Part III: Recommendations (Current → Proposed)

*No recommendations generated (phase did not complete).*

---

## Risk Matrix

| Risk | Severity | Likelihood | Related Recommendation |
|------|----------|------------|------------------------|
| High global mutable state in internal/agent package causing concurrency issues and maintainability risks | high | medium | Reduce Global Mutable State in Agent Package |
| os.Exit calls in cmd/certify package interfering with graceful shutdown and error handling | high | medium | Eliminate os.Exit Calls from CLI Commands |
| Incomplete TODO markers and excessive nesting in internal/evidence package affecting code quality | medium | high | Address TODO Markers and Nesting Depth in Evidence Package |
| Incomplete policy evaluation documentation in internal/policy package | medium | low | Complete Policy Evaluation Documentation |

## Prioritized Roadmap

| # | Item | Effort | Impact | Current → Projected |
|--:|------|--------|--------|---------------------|
| 1 | Reduce Global Mutable State in Agent Package | M | high | key metric: observations 5 → 1, avg_score 91.8% → 92.3% |
| 2 | Eliminate os.Exit Calls from CLI Commands | M | high | key metric: observations 2 → 0, avg_score 91.4% → 92.1% |
| 3 | Address TODO Markers and Nesting Depth in Evidence Package | M | medium | key metric: observations 4 → 1, avg_score 91.3% → 92.0% |
| 4 | Complete Policy Evaluation Documentation | S | medium | key metric: observations 1 → 0, avg_score 91.5% → 92.2% |

---

## Appendix: Data Sources

- **748** units across **25** packages · Score: 91.7%
- Evidence: lint, test, coverage, structural, git history
- Snapshot computed from certification records at `68eb2b49`

---

*Generated by [Certify](https://github.com/iksnae/code-certification) `architect` command.*
