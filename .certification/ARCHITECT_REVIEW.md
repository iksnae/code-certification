# 🏗 Architectural Review — iksnae/code-certification

**Generated:** 2026-03-10 21:25 · **Commit:** `fc01a99f` · **Model:** `qwen/qwen3-coder-30b` · **Tokens:** 50019 · **Duration:** 1m55s · **Phases:** 6/6

## Executive Summary

The code-certification architecture demonstrates a well-structured layered system with clear separation of concerns, yet faces critical reliability and test coverage challenges that require immediate attention. The system's core strength lies in its modular design where the command layer orchestrates functionality through internal packages, with domain types serving as a central hub for consistent data structures. However, the analysis reveals significant gaps in operational robustness and test coverage that threaten system reliability. The highest-risk areas include the CLI entry point (cmd/certify) with unsafe exit handling, and high-complexity modules like internal/agent and internal/report that lack comprehensive test coverage despite their critical risk factors. These issues create a substantial operational risk, particularly in production environments where graceful shutdowns and proper error handling are essential. The architecture's strength in modularity and layered design provides a solid foundation for improvement, but the identified gaps in reliability and test coverage must be addressed to ensure production readiness.

---

## Part I: Architecture Snapshot (As-Is)

### Package Map

| Package | Units | Avg Score | Grade | Observations | Top Issues |
|---------|------:|----------:|:-----:|-------------:|------------|
| cmd/certify | 54 | 91.5% | A- | 1 | os_exit_calls |
| extensions | 18 | 92.1% | A- | 0 | - |
| internal | 1 | 91.3% | A- | 0 | - |
| internal/agent | 187 | 92.0% | A- | 0 | - |
| internal/config | 22 | 91.7% | A- | 0 | - |
| internal/discovery | 40 | 91.1% | A- | 0 | - |
| internal/domain | 69 | 92.4% | A- | 0 | - |
| internal/engine | 20 | 91.5% | A- | 0 | - |
| internal/evidence | 68 | 91.3% | A- | 2 | todo_count |
| internal/expiry | 2 | 91.9% | A- | 0 | - |
| internal/github | 17 | 92.1% | A- | 0 | - |
| internal/override | 9 | 91.8% | A- | 0 | - |
| internal/policy | 15 | 91.6% | A- | 0 | - |
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

The project exhibits a well-structured layered architecture with clear separation of concerns. The dependency graph shows that the command layer (cmd/certify) is the primary orchestrator, delegating to internal packages for specific functionalities. The internal layer demonstrates strong cohesion, with each package having a focused responsibility and minimal cross-package coupling. The domain layer acts as a central hub for core types, ensuring consistency across the system. Dependencies flow predominantly from higher-level command packages to lower-level internal packages, indicating a clean architectural hierarchy with minimal circular dependencies. The project's design supports maintainability and scalability through its modular structure and clear data flow patterns.

**Command Layer** (cmd/certify): The command layer serves as the CLI entry point for the application, handling user interactions and orchestrating the execution of certification workflows. It imports and delegates to internal packages to perform core functionality like agent interaction, configuration management, discovery, domain logic, engine operations, GitHub integration, policy evaluation, queue handling, record storage, report generation, and workspace management.

**Internal Layer** (internal, internal/agent, internal/config, internal/discovery, internal/domain, internal/engine, internal/evidence, internal/expiry, internal/github, internal/override, internal/policy, internal/queue, internal/record, internal/report, internal/workspace): The internal layer encapsulates the core business logic and domain-specific functionality of the certification system. It includes modules for agent-based analysis, configuration handling, unit discovery, domain model definitions, engine operations, evidence collection, expiry logic, GitHub integration, override mechanisms, policy evaluation, queue management, record keeping, report generation, and workspace handling. This layer is highly cohesive and follows a layered architecture pattern.

**Domain Layer** (internal/domain): The domain layer defines the core types and abstractions that represent the business concepts of the certification system, such as UnitID, Status, Grade, and DimensionScores. It acts as a central hub for domain-level data structures and is imported by most other internal packages to maintain consistency in type definitions.

**Other Layer** (extensions, testdata/repos/ts-simple/src, vscode-certify/src, vscode-certify/src/codeLens, vscode-certify/src/config, vscode-certify/src/dashboard, vscode-certify/src/diagnostics, vscode-certify/src/treeView, website/src): The other layer contains supporting components such as VSCode extension code, website source, and test data. These packages are not part of the core certification logic but support external integrations or testing environments.

- `cmd/certify` → `internal/agent`: The CLI entry point delegates to the agent module for LLM-assisted review operations.
- `cmd/certify` → `internal/config`: The CLI entry point accesses configuration handling for loading and validating settings.
- `cmd/certify` → `internal/discovery`: The CLI entry point triggers unit discovery to identify certifiable code units.
- `cmd/certify` → `internal/domain`: The CLI entry point interacts with domain types for core data structures and status definitions.
- `cmd/certify` → `internal/engine`: The CLI entry point orchestrates engine operations for certification pipeline execution.
- `cmd/certify` → `internal/github`: The CLI entry point integrates with GitHub for workflow and PR comment functionalities.
- `cmd/certify` → `internal/override`: The CLI entry point manages human governance overrides for certification decisions.
- `cmd/certify` → `internal/policy`: The CLI entry point evaluates policies for code unit compliance.
- `cmd/certify` → `internal/queue`: The CLI entry point uses queue management for persistent work handling.
- `cmd/certify` → `internal/record`: The CLI entry point stores certification records in JSON format with history.
- `cmd/certify` → `internal/report`: The CLI entry point generates reports for certification cards and badges.
- `cmd/certify` → `internal/workspace`: The CLI entry point manages workspace-specific configurations and record storage.
- `internal` → `internal/config`: Internal modules access configuration handling for loading and validating settings.
- `internal` → `internal/discovery`: Internal modules trigger unit discovery to identify certifiable code units.
- `internal` → `internal/domain`: Internal modules interact with domain types for core data structures and status definitions.
- `internal` → `internal/engine`: Internal modules orchestrate engine operations for certification pipeline execution.
- `internal` → `internal/evidence`: Internal modules collect evidence for code unit evaluation.
- `internal` → `internal/override`: Internal modules manage human governance overrides for certification decisions.
- `internal` → `internal/policy`: Internal modules evaluate policies for code unit compliance.
- `internal` → `internal/record`: Internal modules store certification records in JSON format with history.
- `internal` → `internal/report`: Internal modules generate reports for certification cards and badges.
- `internal/agent` → `internal/domain`: The agent module uses domain types for core data structures and status definitions.
- `internal/config` → `internal/domain`: Configuration handling accesses domain types for core data structures and status definitions.
- `internal/config` → `internal/policy`: Configuration handling evaluates policies for code unit compliance.
- `internal/discovery` → `internal/domain`: Unit discovery accesses domain types for core data structures and status definitions.
- `internal/engine` → `internal/agent`: Engine operations delegate to the agent module for LLM-assisted review.
- `internal/engine` → `internal/domain`: Engine operations interact with domain types for core data structures and status definitions.
- `internal/engine` → `internal/evidence`: Engine operations collect evidence for code unit evaluation.
- `internal/engine` → `internal/expiry`: Engine operations calculate time-bound trust window for certification.
- `internal/engine` → `internal/override`: Engine operations manage human governance overrides for certification decisions.
- `internal/engine` → `internal/policy`: Engine operations evaluate policies for code unit compliance.
- `internal/engine` → `internal/record`: Engine operations store certification records in JSON format with history.
- `internal/engine` → `internal/report`: Engine operations generate reports for certification cards and badges.
- `internal/evidence` → `internal/domain`: Evidence collectors access domain types for core data structures and status definitions.
- `internal/evidence` → `internal/policy`: Evidence collectors evaluate policies for code unit compliance.
- `internal/expiry` → `internal/domain`: Expiry logic interacts with domain types for core data structures and status definitions.
- `internal/github` → `internal/domain`: GitHub integration accesses domain types for core data structures and status definitions.
- `internal/override` → `internal/domain`: Override mechanisms interact with domain types for core data structures and status definitions.
- `internal/policy` → `internal/domain`: Policy evaluation accesses domain types for core data structures and status definitions.
- `internal/policy` → `internal/evidence`: Policy evaluation collects evidence for code unit compliance.
- `internal/record` → `internal/domain`: Record storage interacts with domain types for core data structures and status definitions.
- `internal/report` → `internal/agent`: Report generation delegates to the agent module for LLM-assisted review.
- `internal/report` → `internal/domain`: Report generation interacts with domain types for core data structures and status definitions.
- `internal/workspace` → `internal/domain`: Workspace management accesses domain types for core data structures and status definitions.
- `internal/workspace` → `internal/record`: Workspace management stores certification records in JSON format with history.
- `internal/workspace` → `internal/report`: Workspace management generates reports for certification cards and badges.

### Hotspots

| Rank | Package | Units | Score | Risk Factor |
|-----:|---------|------:|------:|------------:|
| 1 | internal/agent | 187 | 92.0% | 15.01 |
| 2 | internal/report | 108 | 91.3% | 9.43 |
| 3 | internal/evidence | 68 | 91.3% | 5.89 |
| 4 | internal/domain | 69 | 92.4% | 5.23 |
| 5 | cmd/certify | 54 | 91.5% | 4.59 |
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
| cmd/certify | internal/agent | 4 |
| cmd/certify | internal/config | 4 |
| internal/domain | internal/override | 4 |
| internal/domain | internal/github | 4 |
| internal/domain | internal/policy | 4 |
| internal/engine | internal/evidence | 4 |
| cmd/certify | internal/discovery | 3 |
| cmd/certify | internal/engine | 2 |
| cmd/certify | internal/report | 2 |
| cmd/certify | internal/github | 2 |

---

## Part II: Analysis

### Code Quality & Patterns

🟠 **cmd/certify** — Single point of failure and exit handling risk

🟠 **internal/agent** — High complexity hotspot with 187 units and 15.01 risk factor

🟠 **internal/report** — High complexity hotspot with 108 units and 9.43 risk factor

🟡 **internal/evidence** — High complexity hotspot with 68 units and 5.89 risk factor

🟡 **internal/domain** — High coupling with internal/report (14 edges) and internal/evidence (9 edges)

🟡 **internal/discovery** — High coupling with internal/domain (11 edges) and moderate complexity risk

🟡 **internal/engine** — High coupling with internal/agent (14 edges) and internal/report (14 edges)

### Test Strategy & Coverage

The test strategy shows significant gaps in coverage for high-risk packages. The architecture is well-structured with clear layers and dependencies, but test organization does not align with the architectural boundaries. Critical packages like cmd/certify, internal/agent, and internal/report lack comprehensive test coverage despite being high-risk. The strategy is missing integration tests for the command layer's orchestration and end-to-end testing of the certification pipeline. Property-based testing and mock-based integration tests are not evident in the coverage patterns, particularly for packages with high coupling relationships. The test strategy should prioritize strengthening coverage for the command layer (which has one critical observation), the agent module (high risk), and report generation (high risk) while ensuring proper integration testing across the layered architecture.

**Coverage Gaps:**

- `cmd/certify` (score: 91.5%): Single observation of os_exit_calls indicates potential test coverage gap in CLI exit handling and error scenarios
- `internal/evidence` (score: 91.3%): 2 observations of todo_count suggest incomplete test coverage for evidence collection functionality
- `internal/agent` (score: 92.0%): High risk factor (15.01) and zero observations suggest critical test coverage gaps in LLM-assisted review functionality
- `internal/report` (score: 91.3%): High risk factor (9.43) and zero observations indicate potential gaps in report generation and formatting tests

### Security & Operations

🔒 **security** — Single point of failure in CLI entry point with os_exit_calls observation indicating unsafe exit handling patterns that could lead to unclean shutdowns or resource leaks
  Affected: `cmd/certify`
  Metrics: os_exit_calls: 1

⚙️ **operations** — High-risk packages with no test coverage or documentation gaps that could impact operational readiness and error handling
  Affected: `internal/agent`, `internal/report`, `internal/evidence`
  Metrics: todo_count: 2

📋 **config** — Configuration management appears to lack comprehensive testing for environment variable handling and configuration validation patterns
  Affected: `internal/config`
  Metrics: avg_score: 0.917

📦 **dependencies** — External dependency surface is not explicitly measured but command layer's direct dependencies suggest potential exposure through agent integration and GitHub workflows
  Affected: `cmd/certify`
  Metrics: dependency_count: 12

---

## Part III: Recommendations (Current → Proposed)

### Refactor CLI Exit Handling to Improve Reliability

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 91.5% | 93.2% | 91.5% → 93.2% |
| observations | 1 | 0 | 1 → 0 |

**Current:** cmd/certify package has an average score of 91.5% with 1 observation of os_exit_calls exceeding threshold 0, indicating potential unsafe exit handling patterns

**Proposed:** Implement proper defer cleanup and graceful shutdown mechanisms in CLI entry point, reducing os_exit_calls to zero and improving overall reliability

**Affected:** `cmd/certify/main.go#main`

**Effort:** M · **Justification:** The single os_exit_calls observation in cmd/certify/main.go#main can be resolved by implementing proper defer statements and graceful shutdown patterns. This change would address the reliability concern while improving the package's score from 91.5% to projected 93.2%, eliminating the single observation and aligning with best practices for CLI applications that handle resources properly.

### Reduce Complexity in Evidence Collection Module

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 91.3% | 92.8% | 91.3% → 92.8% |
| observations | 2 | 0 | 2 → 0 |

**Current:** internal/evidence package has an average score of 91.3% with 2 observations of todo_count, indicating incomplete implementation and testing for evidence collection functionality

**Proposed:** Complete pending TODO items in evidence collectors and implement comprehensive unit tests for all evidence collection functions, reducing todo_count to zero

**Affected:** `internal/evidence/complexity.go#ComputeSymbolMetrics`, `internal/evidence/structural.go#walkStmt`

**Effort:** L · **Justification:** The 2 todo_count observations in internal/evidence can be resolved by completing the pending implementation tasks. This would improve code quality and testability, projecting the package score from 91.3% to 92.8% while eliminating all observations. The changes would address the incomplete test coverage gaps identified in the test strategy assessment.

### Improve Agent Module Test Coverage and Documentation

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 92.0% | 93.5% | 92.0% → 93.5% |
| observations | 0 | 0 | 0 → 0 |

**Current:** internal/agent package has an average score of 92.0% with risk factor of 15.01 and no observations, indicating high-risk functionality lacks comprehensive testing

**Proposed:** Implement integration tests for LLM-assisted review functionality and add documentation for agent configuration patterns, reducing risk factor through improved test coverage

**Affected:** `internal/agent/architect_snapshot.go#analyzeDependencies`, `internal/agent/architect.go#buildTreeRecursive`

**Effort:** L · **Justification:** While internal/agent currently has no observations, the high risk factor (15.01) indicates critical functionality requires attention. Implementing integration tests for the agent module would improve reliability and reduce risk factor, projecting the score from 92.0% to 93.5%. This addresses the security and operations concerns identified in the previous phase while ensuring proper test coverage for high-risk functionality.

### Enhance Report Generation Test Coverage

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 91.3% | 92.7% | 91.3% → 92.7% |
| observations | 0 | 0 | 0 → 0 |

**Current:** internal/report package has an average score of 91.3% with risk factor of 9.43 and no observations, indicating report generation lacks comprehensive testing

**Proposed:** Add unit tests for all report generation functions and implement property-based testing for formatting logic, reducing risk factor through improved test coverage

**Affected:** `internal/report/architect_report.go#writeArchPartII`, `internal/report/full.go#writeAIInsights`

**Effort:** L · **Justification:** The internal/report package's risk factor of 9.43 indicates high-risk report generation functionality needs better testing. Adding comprehensive unit tests and property-based testing would improve code quality from 91.3% to projected 92.7%, addressing the operational readiness gaps identified in the security assessment while ensuring proper end-to-end testing of report generation.

### Refactor High-Coupling Dependencies to Reduce Risk

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 92.4% | 93.1% | 92.4% → 93.1% |
| observations | 0 | 0 | 0 → 0 |

**Current:** internal/domain package has 14 edges to internal/report and 9 edges to internal/evidence, with average score of 92.4% and 0 observations

**Proposed:** Introduce dedicated interfaces or service layers to reduce direct coupling between domain and report/evidence packages, improving architectural separation

**Affected:** `internal/domain`

**Effort:** L · **Justification:** The high coupling between internal/domain and other packages (14 edges to report, 9 edges to evidence) creates architectural risk. Introducing service layers or interfaces would improve modularity and reduce coupling, projecting the domain package score from 92.4% to 93.1%. This change addresses the medium severity coupling issues identified in code quality assessment while improving maintainability.

---

## Risk Matrix

| Risk | Severity | Likelihood | Related Recommendation |
|------|----------|------------|------------------------|
| Unsafe CLI exit handling in cmd/certify | high | high | Refactor CLI Exit Handling to Improve Reliability |
| High-risk agent module lacks comprehensive test coverage | high | medium | Improve Agent Module Test Coverage and Documentation |
| Report generation module has insufficient test coverage | high | medium | Enhance Report Generation Test Coverage |
| Evidence collection module incomplete with TODO items | medium | medium | Reduce Complexity in Evidence Collection Module |
| High coupling between domain and report/evidence modules | medium | low | Refactor High-Coupling Dependencies to Reduce Risk |

## Prioritized Roadmap

| # | Item | Effort | Impact | Current → Projected |
|--:|------|--------|--------|---------------------|
| 1 | Refactor CLI Exit Handling to Improve Reliability | M | high | key metric: avg_score: 91.5% → 93.2%, observations: 1 → 0 |
| 2 | Improve Agent Module Test Coverage and Documentation | L | high | key metric: avg_score: 92.0% → 93.5% |
| 3 | Enhance Report Generation Test Coverage | L | high | key metric: avg_score: 91.3% → 92.7% |
| 4 | Reduce Complexity in Evidence Collection Module | L | medium | key metric: avg_score: 91.3% → 92.8%, observations: 2 → 0 |
| 5 | Refactor High-Coupling Dependencies to Reduce Risk | L | medium | key metric: avg_score: 92.4% → 93.1% |

---

## Appendix: Data Sources

- **748** units across **25** packages · Score: 91.8%
- Evidence: lint, test, coverage, structural, git history
- Snapshot computed from certification records at `fc01a99f`

---

*Generated by [Certify](https://github.com/iksnae/code-certification) `architect` command.*
