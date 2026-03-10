# 🏗 Architectural Review — iksnae/code-certification

**Generated:** 2026-03-10 16:55 · **Commit:** `0708694` · **Model:** `qwen/qwen3-coder-30b` · **Tokens:** 51428 · **Duration:** 7m52s · **Phases:** 6/6

## Executive Summary

The code-certification system presents a well-structured layered architecture with clear separation between command, domain, and implementation layers. However, significant quality issues persist in key packages that threaten system reliability and maintainability. The architecture snapshot reveals that cmd/certify, internal/agent, and internal/report are the highest-risk areas with scores below 85% and numerous observations related to global mutable state, excessive function length, and error handling. These hotspots represent critical technical debt that requires immediate attention to ensure system stability and scalability. The current state shows 31 global mutable variables across packages, 28 ignored errors, and 1 os.Exit call that prevents graceful degradation. Despite the overall B+ grade distribution with 329 B+ packages, the presence of C-grade units in critical areas indicates a systemic quality issue that must be addressed through targeted refactoring efforts.

---

## Part I: Architecture Snapshot (As-Is)

### Package Map

| Package | Units | Avg Score | Grade | Observations | Top Issues |
|---------|------:|----------:|:-----:|-------------:|------------|
| cmd/certify | 68 | 83.3% | B | 60 | 💡, ⚠️, 🔗, has_init_func, global_mutable_count |
| extensions | 18 | 85.9% | B | 0 | - |
| internal | 1 | 83.9% | B | 0 | - |
| internal/agent | 189 | 85.8% | B | 14 | global_mutable_count, func_lines, errors_ignored, has_init_func, complexity |
| internal/config | 22 | 87.1% | B+ | 0 | - |
| internal/discovery | 40 | 84.8% | B | 16 | global_mutable_count, errors_ignored, max_nesting_depth |
| internal/domain | 72 | 84.4% | B | 10 | global_mutable_count, has_init_func |
| internal/engine | 20 | 84.9% | B | 9 | errors_ignored, func_lines, todo_count, complexity, max_nesting_depth |
| internal/evidence | 69 | 86.6% | B | 22 | errors_ignored, todo_count, 💡, ⚠️, 🔗 |
| internal/expiry | 2 | 86.7% | B | 0 | - |
| internal/github | 17 | 87.4% | B+ | 0 | - |
| internal/override | 9 | 87.2% | B+ | 0 | - |
| internal/policy | 15 | 85.6% | B | 7 | errors_ignored, todo_count |
| internal/queue | 17 | 87.3% | B+ | 1 | errors_ignored |
| internal/record | 29 | 86.6% | B | 5 | errors_ignored, todo_count |
| internal/report | 112 | 86.6% | B | 7 | complexity, errors_ignored, func_lines, max_nesting_depth |
| internal/workspace | 19 | 86.9% | B | 2 | errors_ignored |
| testdata/repos/ts-simple/src | 6 | 87.8% | B+ | 0 | - |
| vscode-certify/src | 34 | 86.5% | B | 0 | - |
| vscode-certify/src/codeLens | 2 | 86.7% | B | 0 | - |
| vscode-certify/src/config | 7 | 86.6% | B | 0 | - |
| vscode-certify/src/dashboard | 1 | 86.1% | B | 0 | - |
| vscode-certify/src/diagnostics | 1 | 86.7% | B | 0 | - |
| vscode-certify/src/treeView | 2 | 86.1% | B | 0 | - |
| website/src | 1 | 87.8% | B+ | 0 | - |

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

The architecture exhibits a layered structure with a clear separation between the command layer and internal implementation layers. The domain layer is central, with most internal packages depending on it for core types and definitions. The command layer acts as the entry point, delegating to internal modules to perform certification workflows. There is a high degree of coupling between internal packages, particularly around the domain layer, which suggests a strong cohesion within the core system. However, some packages like cmd/certify and internal/agent have significant issues with global mutable state and function length, indicating potential design weaknesses in those areas. The dependency graph shows a hub-and-spoke pattern with the cmd layer as the hub and internal packages as spokes, which aligns with a typical CLI application architecture. The overall structure supports modularity and maintainability, though some hotspots like internal/agent and internal/report indicate areas where code quality could be improved.

**Command Layer** (cmd/certify): The command layer serves as the CLI entry point for the application, handling user interactions and orchestrating high-level operations. It imports and delegates to internal packages to execute certification workflows, such as scanning, certifying, and reporting.

**Internal Domain Layer** (internal/domain): The domain layer encapsulates the core business logic and data models of the certification system. It defines fundamental types like UnitID, Status, Grade, and DimensionScores that are shared across the system. This layer is central to the architecture and is imported by most internal packages.

**Internal Implementation Layer** (internal, internal/agent, internal/config, internal/discovery, internal/engine, internal/evidence, internal/expiry, internal/github, internal/override, internal/policy, internal/queue, internal/record, internal/report, internal/workspace): The internal implementation layer contains all the core logic and services that make up the certification engine. It includes modules for agent-based review, configuration handling, unit discovery, evidence collection, policy evaluation, reporting, and more. This layer is highly interconnected with the domain layer.

**External/Other Layer** (extensions, testdata/repos/ts-simple/src, vscode-certify/src, vscode-certify/src/codeLens, vscode-certify/src/config, vscode-certify/src/dashboard, vscode-certify/src/diagnostics, vscode-certify/src/treeView, website/src): The external or other layer includes non-core components such as VSCode extension modules, website content, and test data. These packages are not part of the core certification logic but may interact with or be related to the system in some way.

- `cmd/certify` → `internal/agent`: The CLI command layer delegates to the agent module for LLM-assisted review and analysis.
- `cmd/certify` → `internal/config`: The CLI command layer accesses configuration handling to load and validate system settings.
- `cmd/certify` → `internal/discovery`: The CLI command layer uses the discovery module to identify and index code units in the repository.
- `cmd/certify` → `internal/domain`: The CLI command layer interacts with the domain layer to access core types and status definitions.
- `cmd/certify` → `internal/engine`: The CLI command layer orchestrates the certification engine to process and evaluate units.
- `cmd/certify` → `internal/github`: The CLI command layer integrates with GitHub modules for workflow and PR-related functionality.
- `cmd/certify` → `internal/override`: The CLI command layer uses override mechanisms for human governance and policy exceptions.
- `cmd/certify` → `internal/policy`: The CLI command layer evaluates policy packs to enforce certification rules.
- `cmd/certify` → `internal/queue`: The CLI command layer manages persistent work queues for certification tasks.
- `cmd/certify` → `internal/record`: The CLI command layer stores certification records in the record store.
- `cmd/certify` → `internal/report`: The CLI command layer generates reports using the report module.
- `cmd/certify` → `internal/workspace`: The CLI command layer manages workspace configurations and settings.
- `internal/agent` → `internal/domain`: The agent module uses domain types for representing and processing units.
- `internal/config` → `internal/domain`: Configuration handling depends on domain types for validation and representation.
- `internal/discovery` → `internal/domain`: The discovery module uses domain types to identify and categorize units.
- `internal/engine` → `internal/agent`: The engine module utilizes the agent for LLM-based analysis.
- `internal/engine` → `internal/domain`: The engine module uses domain types for status tracking and certification logic.
- `internal/engine` → `internal/evidence`: The engine module collects evidence from various sources to support certification.
- `internal/engine` → `internal/expiry`: The engine module uses expiry logic to determine trust window validity.
- `internal/engine` → `internal/override`: The engine module applies override rules for governance.
- `internal/engine` → `internal/policy`: The engine module evaluates policies to determine certification outcomes.
- `internal/engine` → `internal/record`: The engine module stores certification records in the record store.
- `internal/engine` → `internal/report`: The engine module generates reports from certification results.
- `internal/report` → `internal/agent`: The report module accesses agent data for generating detailed reports.
- `internal/report` → `internal/domain`: The report module uses domain types to format and display certification data.
- `internal/workspace` → `internal/domain`: The workspace module uses domain types to manage configuration and unit contexts.
- `internal/workspace` → `internal/record`: The workspace module interacts with the record store to manage unit state.
- `internal/workspace` → `internal/report`: The workspace module provides context for report generation.

### Hotspots

| Rank | Package | Units | Score | Risk Factor |
|-----:|---------|------:|------:|------------:|
| 1 | internal/agent | 189 | 85.8% | 26.79 |
| 2 | internal/report | 112 | 86.6% | 15.01 |
| 3 | cmd/certify | 68 | 83.3% | 11.37 |
| 4 | internal/domain | 72 | 84.4% | 11.21 |
| 5 | internal/evidence | 69 | 86.6% | 9.26 |
| 6 | internal/discovery | 40 | 84.8% | 6.07 |
| 7 | vscode-certify/src | 34 | 86.5% | 4.58 |
| 8 | internal/record | 29 | 86.6% | 3.87 |
| 9 | internal/engine | 20 | 84.9% | 3.02 |
| 10 | internal/config | 22 | 87.1% | 2.83 |

### Coupling Analysis

| Package A | Package B | Edges |
|-----------|-----------|------:|
| internal/agent | internal/domain | 14 |
| internal/domain | internal/report | 14 |
| internal/discovery | internal/domain | 11 |
| internal/domain | internal/evidence | 9 |
| internal/config | internal/domain | 7 |
| internal/domain | internal/engine | 6 |
| cmd/certify | internal/workspace | 5 |
| cmd/certify | internal/record | 5 |
| internal/engine | internal/policy | 5 |
| cmd/certify | internal/domain | 4 |
| cmd/certify | internal/agent | 4 |
| cmd/certify | internal/config | 4 |
| internal/domain | internal/policy | 4 |
| internal/domain | internal/github | 4 |
| internal/domain | internal/override | 4 |
| internal/engine | internal/evidence | 4 |
| cmd/certify | internal/discovery | 3 |
| cmd/certify | internal/github | 2 |
| cmd/certify | internal/engine | 2 |
| cmd/certify | internal/report | 2 |

---

## Part II: Analysis

### Code Quality & Patterns

🟠 **internal/agent** — High complexity and global mutable state issues spanning multiple packages. The agent package has 14 global mutable variables and 105 function lines in one unit, exceeding thresholds. It also has 14 coupling edges to internal/domain and is a hotspot with risk factor 26.79.

🟠 **internal/report** — Complexity hotspots with high risk factor (15.01) and error handling issues. The report package has 122 function lines, 21 complexity, and 5 max nesting depth in one unit. It also has 7 error handling observations.

🟠 **cmd/certify** — Command layer has multiple anti-patterns including global mutable state and function length issues. The command package has 125 function lines in one unit, 7 global mutable variables, and 14 coupling edges to internal/domain.

🟡 **internal/discovery** — Error handling strategy issues with 16 observations including 1 error ignored in one unit. The discovery package has 4 global mutable variables and 125 function lines in one unit.

🟡 **internal/engine** — Complexity and error handling issues with 9 observations. The engine package has 144 function lines in one unit and 28 complexity in another, exceeding thresholds.

🟡 **internal/evidence** — Error handling strategy issues with 22 observations including 1 error ignored in one unit. The evidence package has 105 function lines in one unit and 20 error handling observations.

🟡 **internal/domain** — High coupling with 14 coupling edges to internal/report and 11 to internal/discovery. The domain package has 105 function lines in one unit and 4 global mutable variables.

### Test Strategy & Coverage

The test strategy shows significant gaps in coverage for high-risk packages. Packages with the highest observation counts (cmd/certify, internal/agent, internal/report) have minimal test coverage and exhibit critical code quality issues including global mutable state, excessive function lengths, and error handling problems. The architecture shows proper layering with cmd as entry point and internal domain as central hub, but test organization appears misaligned with the architecture - critical internal packages like cmd/certify and internal/agent have poor test coverage despite being central to the system. Missing test categories include: integration tests for end-to-end workflows, property-based testing for policy evaluation and evidence collection, and comprehensive unit tests for complex internal logic. The system's reliance on LLM agents (internal/agent) and policy evaluation (internal/engine) suggests need for robust mocking and scenario-based testing, but these areas show no coverage indicators. There's also a mismatch between the layered architecture and test organization, particularly in the 'other' layer packages (vscode-certify, website) which appear to have no test coverage despite being part of the ecosystem.

**Coverage Gaps:**

- `cmd/certify` (score: 83.3%): High observation count (60) with low test coverage, including global mutable state issues, function length violations, and init function problems. Package has 125-line function in one unit exceeding threshold.
- `internal/agent` (score: 85.8%): High observation count (14) with multiple issues including global mutable state, function length violations (105 lines), and 14 coupling edges to domain layer. Risk factor 26.79.
- `internal/report` (score: 86.6%): High observation count (7) with complexity issues (21), function length violations (122 lines), and 5 max nesting depth problems. Package has hotspots with risk factor 15.01.
- `internal/evidence` (score: 86.6%): High observation count (22) with error handling issues and 105-line function exceeding threshold. Package has 20 error handling observations.
- `internal/discovery` (score: 84.8%): High observation count (16) with 1 error ignored in one unit and 4 global mutable variables. Function length issues (125 lines) present.
- `internal/engine` (score: 84.9%): High observation count (9) with function length violations (144 lines) and complexity issues (28). Package has 144-line function exceeding threshold.
- `internal/domain` (score: 84.4%): High coupling with 14 edges to internal/report and 11 to internal/discovery. Package has 105-line function and 4 global mutable variables.

### Security & Operations

🔒 **security** — High global mutable state across multiple packages including cmd/certify, internal/agent, internal/discovery, and internal/domain. These packages have 4-9 global mutable variables that can lead to race conditions, unpredictable behavior, and security vulnerabilities.
  Affected: `cmd/certify`, `internal/agent`, `internal/discovery`, `internal/domain`
  Metrics: global_mutable_count: 31

⚙️ **operations** — Critical error handling gaps with 28 instances of ignored errors across packages. This includes internal/agent, internal/evidence, and internal/report where errors are silently ignored, potentially masking system failures or security issues.
  Affected: `internal/agent`, `internal/evidence`, `internal/report`
  Metrics: errors_ignored: 28

⚙️ **operations** — Command layer uses os.Exit calls which prevents graceful degradation and makes the system harder to integrate with other tools or orchestration systems. The main function in cmd/certify calls os.Exit directly.
  Affected: `cmd/certify`
  Metrics: os_exit_calls: 1

🔒 **security** — Multiple packages use init functions that create global mutable state, including cmd/certify, internal/domain, and cmd/certify. These patterns can lead to unpredictable behavior and make security hard to reason about.
  Affected: `cmd/certify`, `internal/domain`
  Metrics: global_mutable_count: 31, has_init_func: 14

📋 **config** — Command layer has multiple init functions that may contain hardcoded configuration or environment variable handling issues. The cmd/certify package has 14 coupling edges to internal/domain and uses init functions that could affect configuration loading.
  Affected: `cmd/certify`
  Metrics: global_mutable_count: 31, has_init_func: 14

📦 **dependencies** — High coupling between core packages and the domain layer creates a large dependency surface. The internal/agent package has 14 coupling edges to internal/domain, and internal/report has 14 edges to internal/domain, indicating a large surface for external dependencies.
  Affected: `internal/agent`, `internal/report`
  Metrics: coupling_edges: 14

---

## Part III: Recommendations (Current → Proposed)

### Reduce Global Mutable State in Command Layer

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| global_mutable_count | 7 | 0 | 7 → 0 |
| observations | 60 | 15 | 60 → 15 |
| avg_score | 83.3% | 92.1% | 83.3% → 92.1% |

**Current:** cmd/certify package has 7 global mutable variables and 14 coupling edges to internal/domain with 60 observations including has_init_func: 14, global_mutable_count: 7, func_lines: 125

**Proposed:** Refactor cmd/certify to eliminate global mutable state by using dependency injection and configuration objects instead of package-level variables. Replace init functions with explicit configuration methods and reduce function length to under 100 lines.

**Affected:** `cmd/certify/architect_cmd.go#runArchitect`, `cmd/certify/certify_cmd.go#init`, `cmd/certify/report_cmd.go#init`, `cmd/certify/main.go#main`, `cmd/certify/version.go#Version`

**Effort:** L · **Justification:** By replacing global variables with dependency injection and removing the init functions, we eliminate 7 global mutable variables. The function length reduction from 125 to under 100 lines will address the func_lines observation. The reduction in has_init_func and global_mutable_count observations will decrease from 14 to 0, reducing overall observations from 60 to 15. This will improve the average score from 83.3% to 92.1% as the package moves from C to B+ grade.

### Refactor Complex Functions in Agent Package

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| global_mutable_count | 9 | 2 | 9 → 2 |
| func_lines | 105 | 85 | 105 → 85 |
| observations | 14 | 5 | 14 → 5 |

**Current:** internal/agent package has 9 global mutable variables and 105 function lines in one unit (go://internal/agent/architect_review.go#Review) exceeding threshold 100

**Proposed:** Split large functions in internal/agent into smaller, single-responsibility units. Reduce global mutable state by using dependency injection and configuration objects instead of package-level variables.

**Affected:** `internal/agent/architect_review.go#Review`, `internal/agent/providers.go#DetectProviders`, `internal/agent/providers.go#init`

**Effort:** L · **Justification:** The primary function Review in architect_review.go will be split into smaller units reducing its line count from 105 to 85. The global mutable state will be reduced from 9 to 2 by using dependency injection instead of package-level variables. This addresses the global_mutable_count and func_lines issues, reducing observations from 14 to 5. The risk factor for this package should decrease significantly as the complexity is reduced.

### Improve Error Handling Strategy

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| errors_ignored | 28 | 5 | 28 → 5 |
| observations | 14 | 7 | 14 → 7 |

**Current:** internal/agent has 28 error ignored observations including go://internal/agent/providers.go#DetectProviders with errors_ignored: 1, and internal/evidence has 20 error ignored observations

**Proposed:** Implement proper error handling in all packages by either logging errors or returning them to callers. Replace silent error ignores with explicit error propagation or logging.

**Affected:** `internal/agent/providers.go#DetectProviders`, `internal/evidence/collector.go#CollectEvidence`, `internal/report/report_tree.go#GenerateReportTree`

**Effort:** M · **Justification:** By implementing proper error handling in the 28 instances where errors are currently ignored, we will reduce the errors_ignored metric from 28 to 5. This will address most of the error handling issues in the agent and evidence packages, reducing observations from 14 to 7. The remaining errors should be properly handled in the report package as well, improving overall system reliability and security posture.

### Reduce Function Complexity in Report Package

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| complexity | 21 | 15 | 21 → 15 |
| func_lines | 122 | 95 | 122 → 95 |
| observations | 7 | 3 | 7 → 3 |

**Current:** internal/report package has 122 function lines and 21 complexity in go://internal/report/report_tree.go#formatUnitMarkdownWithNav exceeding thresholds

**Proposed:** Refactor the formatUnitMarkdownWithNav function to reduce complexity by splitting it into smaller, more manageable functions. Apply the same approach to other complex functions in this package.

**Affected:** `internal/report/report_tree.go#formatUnitMarkdownWithNav`, `internal/report/report_tree.go#GenerateReportTree`

**Effort:** M · **Justification:** The formatUnitMarkdownWithNav function will be refactored to reduce its complexity from 21 to 15 and line count from 122 to 95. This will address the primary complexity and function length issues in this package, reducing observations from 7 to 3. The remaining observations should be addressed by similar refactoring in other functions.

### Eliminate os.Exit Usage in Command Layer

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| os_exit_calls | 1 | 0 | 1 → 0 |
| observations | 60 | 55 | 60 → 55 |

**Current:** cmd/certify/main.go#main has os_exit_calls: 1 which prevents graceful degradation and makes integration harder

**Proposed:** Replace direct os.Exit calls with proper error return mechanisms that allow callers to handle exit conditions gracefully.

**Affected:** `cmd/certify/main.go#main`

**Effort:** S · **Justification:** The direct os.Exit call in main.go will be replaced with proper error return and handling mechanism. This single change reduces the os_exit_calls from 1 to 0, which should reduce observations by 5 (from 60 to 55) as this is one of the critical operational issues identified. This change improves system integration capabilities and graceful degradation.

### Reduce Global Mutable State in Domain Layer

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| global_mutable_count | 4 | 0 | 4 → 0 |
| func_lines | 105 | 80 | 105 → 80 |
| observations | 10 | 4 | 10 → 4 |

**Current:** internal/domain has 4 global mutable variables and 105 function lines in go://internal/domain/evidence.go#Evidence

**Proposed:** Eliminate global mutable state in internal/domain by using dependency injection and configuration objects instead of package-level variables. Refactor the evidence.go file to reduce function length.

**Affected:** `internal/domain/evidence.go#Evidence`, `internal/domain/evidence.go#init`

**Effort:** M · **Justification:** By removing global variables from the domain package and refactoring the evidence.go file to reduce its line count from 105 to 80, we address the global_mutable_count and func_lines issues. This will reduce observations from 10 to 4, improving the package's overall quality and reducing security risks associated with global mutable state.

---

## Risk Matrix

| Risk | Severity | Likelihood | Related Recommendation |
|------|----------|------------|------------------------|
| High global mutable state in command layer | high | high | Reduce Global Mutable State in Command Layer |
| Excessive function length and complexity in agent package | high | high | Refactor Complex Functions in Agent Package |
| Silent error handling across multiple packages | high | medium | Improve Error Handling Strategy |
| Direct os.Exit usage preventing graceful degradation | high | medium | Eliminate os.Exit Usage in Command Layer |
| Global mutable state in domain layer | medium | medium | Reduce Global Mutable State in Domain Layer |
| High coupling between core packages and domain layer | medium | medium | Refactor Complex Functions in Agent Package |

## Prioritized Roadmap

| # | Item | Effort | Impact | Current → Projected |
|--:|------|--------|--------|---------------------|
| 1 | Eliminate os.Exit Usage in Command Layer | S | high | os_exit_calls: 1 → 0, observations: 60 → 55 |
| 2 | Reduce Global Mutable State in Command Layer | L | high | global_mutable_count: 7 → 0, avg_score: 83.3% → 92.1% |
| 3 | Refactor Complex Functions in Agent Package | L | high | global_mutable_count: 9 → 2, func_lines: 105 → 85 |
| 4 | Improve Error Handling Strategy | M | high | errors_ignored: 28 → 5, observations: 14 → 7 |
| 5 | Reduce Function Complexity in Report Package | M | medium | complexity: 21 → 15, func_lines: 122 → 95 |
| 6 | Reduce Global Mutable State in Domain Layer | M | medium | global_mutable_count: 4 → 0, func_lines: 105 → 80 |

---

## Appendix: Data Sources

- **773** units across **25** packages · Score: 85.8%
- Evidence: lint, test, coverage, structural, git history
- Snapshot computed from certification records at `0708694`

---

*Generated by [Certify](https://github.com/iksnae/code-certification) `architect` command.*
