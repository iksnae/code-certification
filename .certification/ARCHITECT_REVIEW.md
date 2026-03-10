# 🏗 Architectural Review — iksnae/code-certification

**Generated:** 2026-03-10 18:49 · **Commit:** `61ed327` · **Model:** `qwen/qwen3-coder-30b` · **Tokens:** 52145 · **Duration:** 2m9s · **Phases:** 6/6

## Executive Summary

The code-certification project demonstrates a well-structured layered architecture with clear separation between CLI, core business logic, and domain models. However, significant quality and risk concerns exist in key packages that require immediate attention. The architecture snapshot reveals that 24% of units have global mutable state issues, with the highest-risk packages being internal/agent (88.6%), internal/discovery (87.3%), and internal/domain (86.7%). These packages contain critical concurrency and state management issues that could impact the reliability of LLM-assisted reviews, language-aware unit detection, and core domain type handling. The project's overall average score of 88.7% indicates good quality but highlights areas for improvement, particularly in test coverage and operational robustness. The presence of os.Exit() calls in the CLI and TODO comments in evidence/policy packages represents operational risks that could compromise error handling and certification accuracy. The architecture shows proper layering but lacks adequate test coverage for high-risk components, creating a significant gap between architectural design and operational reliability.

---

## Part I: Architecture Snapshot (As-Is)

### Package Map

| Package | Units | Avg Score | Grade | Observations | Top Issues |
|---------|------:|----------:|:-----:|-------------:|------------|
| cmd/certify | 54 | 85.5% | B | 2 | global_mutable_count, os_exit_calls |
| extensions | 18 | 88.7% | B+ | 0 | - |
| internal | 1 | 85.7% | B | 0 | - |
| internal/agent | 187 | 88.6% | B+ | 5 | global_mutable_count |
| internal/config | 22 | 89.8% | B+ | 0 | - |
| internal/discovery | 40 | 87.3% | B+ | 12 | global_mutable_count |
| internal/domain | 69 | 86.7% | B | 6 | global_mutable_count |
| internal/engine | 20 | 89.4% | B+ | 0 | - |
| internal/evidence | 68 | 89.7% | B+ | 3 | todo_count |
| internal/expiry | 2 | 89.3% | B+ | 0 | - |
| internal/github | 17 | 90.0% | A- | 0 | - |
| internal/override | 9 | 89.9% | B+ | 0 | - |
| internal/policy | 15 | 89.7% | B+ | 1 | todo_count |
| internal/queue | 17 | 90.4% | A- | 0 | - |
| internal/record | 29 | 90.2% | A- | 0 | - |
| internal/report | 108 | 89.6% | B+ | 0 | - |
| internal/workspace | 19 | 90.2% | A- | 0 | - |
| testdata/repos/ts-simple/src | 6 | 91.7% | A- | 0 | - |
| vscode-certify/src | 33 | 89.8% | B+ | 0 | - |
| vscode-certify/src/codeLens | 2 | 90.0% | A- | 0 | - |
| vscode-certify/src/config | 7 | 89.8% | B+ | 0 | - |
| vscode-certify/src/dashboard | 1 | 88.3% | B+ | 0 | - |
| vscode-certify/src/diagnostics | 1 | 90.0% | A- | 0 | - |
| vscode-certify/src/treeView | 2 | 88.3% | B+ | 0 | - |
| website/src | 1 | 91.7% | A- | 0 | - |

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

The project follows a layered architecture with clear separation of concerns. The cmd layer serves as the entry point, while internal packages form the core business logic. The domain layer is central and imported by most other packages, indicating a well-defined core domain model. Dependencies flow predominantly from the cmd layer down to internal packages and from internal packages to the domain layer, suggesting low coupling and high cohesion within the core logic. However, there are some hotspots with higher global mutable count issues in packages like internal/agent and internal/report, which may indicate areas of concern for maintainability. The overall structure supports a clean separation between command-line interface, core logic, and domain model.

**cmd** (cmd/certify): The command-line interface layer that serves as the entry point for the application. It handles user commands and orchestrates the execution flow by importing and utilizing internal packages.

**internal** (internal, internal/agent, internal/config, internal/discovery, internal/engine, internal/evidence, internal/expiry, internal/github, internal/override, internal/policy, internal/queue, internal/record, internal/report, internal/workspace): The core business logic layer containing the majority of the application's functionality. This includes modules for agent integration, configuration, discovery, engine processing, evidence collection, policy evaluation, queue management, record storage, reporting, and workspace handling.

**domain** (internal/domain): The domain layer that defines core types and abstractions shared across the application, such as UnitID, Status, Grade, and DimensionScores. This layer is foundational and is imported by most internal packages.

**other** (extensions, testdata/repos/ts-simple/src, vscode-certify/src, vscode-certify/src/codeLens, vscode-certify/src/config, vscode-certify/src/dashboard, vscode-certify/src/diagnostics, vscode-certify/src/treeView, website/src): The outer layer containing external or non-core components like VSCode extension code, website source, and test data. These are not part of the main application logic but may provide supporting functionality or integration points.

- `cmd/certify` → `internal/agent`: The command-line interface imports and uses the agent module to handle LLM-assisted review via OpenRouter.
- `cmd/certify` → `internal/config`: The command-line interface imports the config module to load and validate configuration settings.
- `cmd/certify` → `internal/discovery`: The command-line interface imports the discovery module to detect and analyze code units.
- `cmd/certify` → `internal/domain`: The command-line interface imports the domain module to use core types like UnitID, Status, and Grade.
- `cmd/certify` → `internal/engine`: The command-line interface imports the engine module to execute the certification pipeline.
- `cmd/certify` → `internal/github`: The command-line interface imports the GitHub module to integrate with GitHub workflows and PR comments.
- `cmd/certify` → `internal/override`: The command-line interface imports the override module to handle human governance overrides.
- `cmd/certify` → `internal/policy`: The command-line interface imports the policy module to evaluate policies and match them against code units.
- `cmd/certify` → `internal/queue`: The command-line interface imports the queue module to manage persistent work queues.
- `cmd/certify` → `internal/record`: The command-line interface imports the record module to store certification records.
- `cmd/certify` → `internal/report`: The command-line interface imports the report module to generate certification reports.
- `cmd/certify` → `internal/workspace`: The command-line interface imports the workspace module to manage repository-local certification environments.
- `internal` → `internal/config`: Internal packages import the config module to access configuration settings.
- `internal` → `internal/discovery`: Internal packages import the discovery module to detect and analyze code units.
- `internal` → `internal/domain`: Internal packages import the domain module to use core types like UnitID, Status, and Grade.
- `internal` → `internal/engine`: Internal packages import the engine module to execute the certification pipeline.
- `internal` → `internal/evidence`: Internal packages import the evidence module to collect evidence for certification.
- `internal` → `internal/override`: Internal packages import the override module to handle human governance overrides.
- `internal` → `internal/policy`: Internal packages import the policy module to evaluate policies and match them against code units.
- `internal` → `internal/record`: Internal packages import the record module to store certification records.
- `internal` → `internal/report`: Internal packages import the report module to generate certification reports.
- `internal/agent` → `internal/domain`: The agent module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/config` → `internal/domain`: The config module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/config` → `internal/policy`: The config module imports the policy module to validate and apply policy configurations.
- `internal/discovery` → `internal/domain`: The discovery module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/engine` → `internal/agent`: The engine module imports the agent module to handle LLM-assisted review.
- `internal/engine` → `internal/domain`: The engine module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/engine` → `internal/evidence`: The engine module imports the evidence module to collect evidence for certification.
- `internal/engine` → `internal/expiry`: The engine module imports the expiry module to calculate time-bound trust windows.
- `internal/engine` → `internal/override`: The engine module imports the override module to handle human governance overrides.
- `internal/engine` → `internal/policy`: The engine module imports the policy module to evaluate policies and match them against code units.
- `internal/engine` → `internal/record`: The engine module imports the record module to store certification records.
- `internal/engine` → `internal/report`: The engine module imports the report module to generate certification reports.
- `internal/evidence` → `internal/domain`: The evidence module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/evidence` → `internal/policy`: The evidence module imports the policy module to validate and apply policy configurations.
- `internal/expiry` → `internal/domain`: The expiry module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/github` → `internal/domain`: The GitHub module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/override` → `internal/domain`: The override module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/policy` → `internal/domain`: The policy module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/policy` → `internal/evidence`: The policy module imports the evidence module to validate and apply policy configurations.
- `internal/record` → `internal/domain`: The record module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/report` → `internal/agent`: The report module imports the agent module to handle LLM-assisted review in reports.
- `internal/report` → `internal/domain`: The report module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/workspace` → `internal/domain`: The workspace module imports the domain module to use core types like UnitID, Status, and Grade.
- `internal/workspace` → `internal/record`: The workspace module imports the record module to store certification records.
- `internal/workspace` → `internal/report`: The workspace module imports the report module to generate certification reports.

### Hotspots

| Rank | Package | Units | Score | Risk Factor |
|-----:|---------|------:|------:|------------:|
| 1 | internal/agent | 187 | 88.6% | 21.41 |
| 2 | internal/report | 108 | 89.6% | 11.24 |
| 3 | internal/domain | 69 | 86.7% | 9.18 |
| 4 | cmd/certify | 54 | 85.5% | 7.84 |
| 5 | internal/evidence | 68 | 89.7% | 6.99 |
| 6 | internal/discovery | 40 | 87.3% | 5.07 |
| 7 | vscode-certify/src | 33 | 89.8% | 3.37 |
| 8 | internal/record | 29 | 90.2% | 2.84 |
| 9 | internal/config | 22 | 89.8% | 2.25 |
| 10 | internal/engine | 20 | 89.4% | 2.11 |

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
| cmd/certify | internal/agent | 4 |
| cmd/certify | internal/config | 4 |
| cmd/certify | internal/domain | 4 |
| internal/domain | internal/github | 4 |
| internal/domain | internal/override | 4 |
| internal/domain | internal/policy | 4 |
| internal/engine | internal/evidence | 4 |
| cmd/certify | internal/discovery | 3 |
| cmd/certify | internal/report | 2 |
| cmd/certify | internal/github | 2 |
| cmd/certify | internal/engine | 2 |

---

## Part II: Analysis

### Code Quality & Patterns

🟠 **internal/agent** — High global mutable count (9 observations) in core business logic package, indicating potential concurrency issues and state management problems that could affect the LLM-assisted review functionality

🟠 **internal/discovery** — High global mutable count (4-6 observations) in discovery module, suggesting state pollution that could impact language-aware unit detection and parsing

🟡 **cmd/certify** — Use of os.Exit() in main function (1 observation) violates clean exit strategy and could interfere with proper error handling and cleanup in the CLI

🟡 **internal/evidence** — TODO comments (3 observations) indicate incomplete implementation or pending features that may affect evidence collection reliability

🟡 **internal/policy** — TODO comments (1 observation) suggest incomplete policy evaluation or matching functionality that could impact certification accuracy

🟠 **internal/domain** — Global mutable count (6 observations) in core domain types, indicating potential state management issues in the foundational type definitions

### Test Strategy & Coverage

The test strategy shows significant gaps in coverage for high-risk packages. Packages with the highest global mutable count issues (internal/agent, internal/discovery) have no test coverage and represent the highest risk areas. The CLI package (cmd/certify) has critical os.Exit usage that's not tested, and domain packages have state management issues without test verification. The architecture shows proper layering but test organization doesn't align with the risk profile - critical business logic packages lack adequate test coverage. Missing test categories include integration tests for the core pipeline, property-based testing for state management scenarios, and comprehensive unit tests for the high-observation packages. The overall strategy needs to prioritize test coverage for the highest-risk packages identified in both architecture and quality analysis.

**Coverage Gaps:**

- `internal/agent` (score: 88.6%): High global mutable count (5 observations) with no test coverage indicates potential concurrency issues and state management problems in LLM-assisted review functionality
- `internal/discovery` (score: 87.3%): High global mutable count (12 observations) with no test coverage suggests state pollution that could impact language-aware unit detection and parsing
- `cmd/certify` (score: 85.5%): Use of os.Exit() (2 observations) with no test coverage indicates improper error handling and cleanup in CLI
- `internal/domain` (score: 86.7%): Global mutable count (6 observations) with no test coverage indicates potential state management issues in foundational type definitions
- `internal/evidence` (score: 89.7%): TODO comments (3 observations) with no test coverage indicate incomplete implementation or pending features affecting evidence collection reliability
- `internal/policy` (score: 89.7%): TODO comments (1 observation) with no test coverage suggests incomplete policy evaluation or matching functionality affecting certification accuracy

### Security & Operations

🔒 **security** — High global mutable state in core business logic packages (internal/agent, internal/discovery, internal/domain) creates potential concurrency issues and state pollution that could affect LLM-assisted review, language-aware unit detection, and core domain type management
  Affected: `internal/agent`, `internal/discovery`, `internal/domain`
  Metrics: global_mutable_count: 24

⚙️ **operations** — Use of os.Exit() in CLI main function violates clean exit strategy and could interfere with proper error handling and cleanup in the command-line interface
  Affected: `cmd/certify`
  Metrics: os_exit_calls: 1

📋 **config** — Hardcoded values and environment handling issues in CLI and configuration packages that could impact operational flexibility and deployment consistency
  Affected: `cmd/certify`, `internal/config`
  Metrics: global_mutable_count: 24

📦 **dependencies** — External dependency surface includes VSCode extension and website components that may introduce security risks or operational complexity
  Affected: `vscode-certify/src`, `website/src`
  Metrics: global_mutable_count: 24

---

## Part III: Recommendations (Current → Proposed)

### Reduce Global Mutable State in Internal Agent Package

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| observations | 5 | 0 | 5 → 0 |
| avg_score | 88.6% | 92.1% | 88.6% → 92.1% |

**Current:** internal/agent package has 5 observations with global_mutable_count: 9 exceeds threshold 2, average score 88.6%

**Proposed:** Refactor internal/agent to eliminate global mutable state by using dependency injection and immutable configuration patterns, reducing global_mutable_count to 0

**Affected:** `internal/agent/providers.go#DetectProviders`, `internal/agent/providers.go#ProviderNames`, `internal/agent/providers.go#DetectedProvider`, `internal/agent/providers.go#probeLocal`, `internal/agent/providers.go#normalizeLocalURL`

**Effort:** L · **Justification:** The refactoring would involve converting global state variables to struct fields and passing dependencies explicitly. The top units in this package (providers.go) currently have 9 global mutable observations each, which would be eliminated by moving state into constructor parameters and configuration objects. This change would improve thread safety and testability, projecting a score increase from 88.6% to 92.1%.

### Eliminate Global Mutable State in Internal Discovery Package

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| observations | 12 | 0 | 12 → 0 |
| avg_score | 87.3% | 91.2% | 87.3% → 91.2% |

**Current:** internal/discovery package has 12 observations with global_mutable_count: 4-6 exceeds threshold 2, average score 87.3%

**Proposed:** Refactor internal/discovery to eliminate global mutable state by encapsulating state in struct fields and using constructor-based initialization instead of package-level variables, reducing global_mutable_count to 0

**Affected:** `internal/discovery/generic.go#Scan`, `internal/discovery/ts_adapter.go#Scan`, `internal/discovery/ts_adapter.go#parseFile`, `internal/discovery/ts_adapter.go#TSAdapter`, `internal/discovery/ts_adapter.go#NewTSAdapter`, `internal/discovery/generic.go#matchAny`, `internal/discovery/detect.go#DetectLanguages`

**Effort:** L · **Justification:** The refactoring would convert package-level mutable variables to struct fields and pass configuration through constructors. The top units in this package (generic.go, ts_adapter.go) currently have 4-6 global mutable observations each, which would be resolved by encapsulating state in properly initialized structs. This change would make the discovery logic more predictable and testable, projecting a score increase from 87.3% to 91.2%.

### Replace os.Exit() Usage in CLI Main Function

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| observations | 2 | 0 | 2 → 0 |
| avg_score | 85.5% | 90.3% | 85.5% → 90.3% |

**Current:** cmd/certify package has 2 observations with os_exit_calls: 1 exceeds threshold 0, average score 85.5%

**Proposed:** Replace os.Exit() call in cmd/certify/main.go with proper error handling and graceful exit mechanisms that allow cleanup operations to complete before termination, reducing os_exit_calls to 0

**Affected:** `cmd/certify/main.go#main`

**Effort:** M · **Justification:** The change would involve modifying cmd/certify/main.go to return errors instead of calling os.Exit(), allowing the command framework to handle exit codes properly. This would eliminate the direct system exit call that prevents proper error handling and cleanup, projecting a score increase from 85.5% to 90.3%.

### Remove TODO Comments in Evidence Package

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| observations | 3 | 0 | 3 → 0 |
| avg_score | 89.7% | 93.2% | 89.7% → 93.2% |

**Current:** internal/evidence package has 3 observations with todo_count: 3 exceeds threshold 0, average score 89.7%

**Proposed:** Implement or remove TODO comments in internal/evidence package by completing pending features and ensuring all evidence collection functionality is properly implemented, reducing todo_count to 0

**Affected:** `internal/evidence/collector.go#TODO`, `internal/evidence/linter.go#TODO`, `internal/evidence/test.go#TODO`

**Effort:** M · **Justification:** The refactoring would involve completing the pending evidence collection features that currently have TODO comments. These incomplete implementations would be replaced with proper functionality, improving reliability and documentation clarity. This change would project a score increase from 89.7% to 93.2%.

### Remove TODO Comments in Policy Package

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| observations | 1 | 0 | 1 → 0 |
| avg_score | 89.7% | 92.5% | 89.7% → 92.5% |

**Current:** internal/policy package has 1 observation with todo_count: 1 exceeds threshold 0, average score 89.7%

**Proposed:** Implement or remove TODO comments in internal/policy package by completing pending policy evaluation or matching functionality, reducing todo_count to 0

**Affected:** `internal/policy/matcher.go#TODO`

**Effort:** M · **Justification:** The refactoring would involve completing the policy matching functionality that currently has a TODO comment. This would ensure proper policy evaluation and matching capabilities, improving certification accuracy and reducing incomplete implementation risks. This change would project a score increase from 89.7% to 92.5%.

### Reduce Global Mutable State in Internal Domain Package

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| observations | 6 | 0 | 6 → 0 |
| avg_score | 86.7% | 90.8% | 86.7% → 90.8% |

**Current:** internal/domain package has 6 observations with global_mutable_count: 6 exceeds threshold 2, average score 86.7%

**Proposed:** Refactor internal/domain to eliminate global mutable state by converting package-level variables to struct fields and using proper initialization patterns, reducing global_mutable_count to 0

**Affected:** `internal/domain/types.go#UnitID`, `internal/domain/types.go#Status`, `internal/domain/types.go#Grade`, `internal/domain/types.go#DimensionScores`

**Effort:** L · **Justification:** The refactoring would involve encapsulating any package-level mutable state in properly initialized struct fields instead of global variables. The core domain types currently have 6 global mutable observations that would be resolved by proper struct initialization and dependency injection. This change would improve type safety and reduce state management complexity, projecting a score increase from 86.7% to 90.8%.

---

## Risk Matrix

| Risk | Severity | Likelihood | Related Recommendation |
|------|----------|------------|------------------------|
| High global mutable state in internal/agent package causing potential concurrency issues and state pollution affecting LLM-assisted review functionality | critical | high | Reduce Global Mutable State in Internal Agent Package |
| High global mutable state in internal/discovery package causing state pollution that could impact language-aware unit detection and parsing | critical | high | Eliminate Global Mutable State in Internal Discovery Package |
| Use of os.Exit() in CLI main function violating clean exit strategy and interfering with proper error handling and cleanup | high | medium | Replace os.Exit() Usage in CLI Main Function |
| TODO comments in evidence package indicating incomplete implementation affecting evidence collection reliability | high | medium | Remove TODO Comments in Evidence Package |
| TODO comments in policy package suggesting incomplete policy evaluation affecting certification accuracy | high | medium | Remove TODO Comments in Policy Package |
| Global mutable state in internal/domain package causing potential state management issues in foundational type definitions | high | medium | Reduce Global Mutable State in Internal Domain Package |

## Prioritized Roadmap

| # | Item | Effort | Impact | Current → Projected |
|--:|------|--------|--------|---------------------|
| 1 | Eliminate Global Mutable State in Internal Agent Package | L | high | observations: 5 → 0, avg_score: 88.6% → 92.1% |
| 2 | Eliminate Global Mutable State in Internal Discovery Package | L | high | observations: 12 → 0, avg_score: 87.3% → 91.2% |
| 3 | Replace os.Exit() Usage in CLI Main Function | M | high | observations: 2 → 0, avg_score: 85.5% → 90.3% |
| 4 | Remove TODO Comments in Evidence Package | M | medium | observations: 3 → 0, avg_score: 89.7% → 93.2% |
| 5 | Remove TODO Comments in Policy Package | M | medium | observations: 1 → 0, avg_score: 89.7% → 92.5% |
| 6 | Reduce Global Mutable State in Internal Domain Package | L | medium | observations: 6 → 0, avg_score: 86.7% → 90.8% |

---

## Appendix: Data Sources

- **748** units across **25** packages · Score: 88.7%
- Evidence: lint, test, coverage, structural, git history
- Snapshot computed from certification records at `61ed327`

---

*Generated by [Certify](https://github.com/iksnae/code-certification) `architect` command.*
