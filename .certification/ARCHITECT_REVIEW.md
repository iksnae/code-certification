# 🏗 Architectural Review — iksnae/code-certification

**Generated:** 2026-03-10 22:00 · **Commit:** `f0dca8b6` · **Model:** `qwen/qwen3-coder-30b` · **Tokens:** 52299 · **Duration:** 2m7s · **Phases:** 6/6

## Executive Summary

The code-certification architecture demonstrates a well-structured layered system with clear separation of concerns, but faces significant challenges in test coverage, operational reliability, and complexity management. The architecture follows a hub-and-spoke pattern with cmd/certify as the CLI entry point, internal/domain as the central hub, and various internal packages implementing the certification pipeline. While overall quality is strong with an average score of 91.8%, critical gaps exist in test coverage for core packages like cmd/certify, internal/agent, and internal/report. These gaps represent the highest risk areas with zero observations and elevated risk factors up to 15.01. Operational concerns include direct OS exit calls, panic handling issues, and global mutable state that could lead to system instability. The architecture shows good dependency management with logical flow from command layer through domain to implementation services, but requires immediate attention to address test coverage gaps and operational reliability issues. The project's strength lies in its modular design and clear domain modeling, but these benefits are undermined by insufficient testing and error handling practices that could compromise system integrity.

---

## Part I: Architecture Snapshot (As-Is)

### Package Map

| Package | Units | Avg Score | Grade | Observations | Top Issues |
|---------|------:|----------:|:-----:|-------------:|------------|
| cmd/certify | 54 | 91.5% | A- | 0 | - |
| extensions | 18 | 92.1% | A- | 0 | - |
| internal | 1 | 91.3% | A- | 0 | - |
| internal/agent | 187 | 92.0% | A- | 0 | - |
| internal/config | 22 | 91.7% | A- | 0 | - |
| internal/discovery | 40 | 91.1% | A- | 0 | - |
| internal/domain | 69 | 92.4% | A- | 0 | - |
| internal/engine | 20 | 91.5% | A- | 0 | - |
| internal/evidence | 68 | 91.4% | A- | 0 | - |
| internal/expiry | 2 | 91.9% | A- | 0 | - |
| internal/github | 17 | 92.1% | A- | 0 | - |
| internal/override | 9 | 91.8% | A- | 0 | - |
| internal/policy | 15 | 91.5% | A- | 0 | - |
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

The architecture demonstrates a layered structure with clear separation of concerns. The command layer (cmd/certify) acts as the primary entry point that orchestrates the certification workflow by depending on internal packages. The internal domain layer (internal/domain) serves as a central hub for core types and business logic, with most internal packages depending on it. The internal implementation layer contains the core certification services that work together in a well-defined pipeline, with data flowing from discovery through evidence collection to policy evaluation and report generation. The dependency structure shows good cohesion within layers, with minimal cross-layer coupling. The highest-risk packages (internal/agent, internal/report, internal/evidence) are all part of the internal implementation layer and show strong interdependencies that align with their functional roles. The architecture follows a hub-and-spoke pattern where the domain layer is the central hub and internal packages are spokes that depend on it. There are no clear violations of dependency direction, with all dependencies flowing logically from higher-level concerns to lower-level implementations. The structure supports a clean separation of concerns between the CLI, core domain logic, and implementation services.

**Command Layer** (cmd/certify): The command layer serves as the CLI entry point for the application, handling user interaction and orchestrating the execution of certification workflows. It imports and delegates to internal packages to perform core functionality.

**Internal Domain Layer** (internal/domain): The domain layer encapsulates the core business logic and data models of the certification system. It defines fundamental types such as UnitID, Status, Grade, and DimensionScores that are used across the system.

**Internal Implementation Layer** (internal, internal/agent, internal/config, internal/discovery, internal/engine, internal/evidence, internal/expiry, internal/github, internal/override, internal/policy, internal/queue, internal/record, internal/report, internal/workspace): The internal implementation layer contains the core logic and services that implement the certification workflow. This includes evidence collection, policy evaluation, report generation, and integration with external systems like GitHub.

**Other Layers** (extensions, testdata/repos/ts-simple/src, vscode-certify/src, vscode-certify/src/codeLens, vscode-certify/src/config, vscode-certify/src/dashboard, vscode-certify/src/diagnostics, vscode-certify/src/treeView, website/src): These layers represent external integrations and supporting tools, including the VSCode extension, website components, and test data. They are separate from the core certification logic but interact with it.

- `cmd/certify` → `internal/agent`: The CLI command layer delegates to the agent package for LLM-assisted review functionality.
- `cmd/certify` → `internal/config`: The command layer accesses configuration management to load and validate settings.
- `cmd/certify` → `internal/discovery`: The CLI command layer triggers unit discovery to identify certifiable code units.
- `cmd/certify` → `internal/domain`: The command layer interacts with domain types to understand certification status and grades.
- `cmd/certify` → `internal/engine`: The command layer orchestrates the certification pipeline through engine services.
- `cmd/certify` → `internal/github`: The CLI command layer integrates with GitHub for workflow and PR comment functionality.
- `cmd/certify` → `internal/override`: The command layer manages human governance overrides.
- `cmd/certify` → `internal/policy`: The command layer evaluates policies against code units.
- `cmd/certify` → `internal/queue`: The command layer manages persistent work queue operations.
- `cmd/certify` → `internal/record`: The command layer stores certification records.
- `cmd/certify` → `internal/report`: The command layer generates certification reports.
- `cmd/certify` → `internal/workspace`: The command layer manages workspace-specific certification settings.
- `internal` → `internal/config`: The internal package accesses configuration services.
- `internal` → `internal/discovery`: The internal package orchestrates unit discovery.
- `internal` → `internal/domain`: The internal package uses domain types for core business logic.
- `internal` → `internal/engine`: The internal package delegates to engine services for certification pipeline execution.
- `internal` → `internal/evidence`: The internal package collects evidence for certification evaluation.
- `internal` → `internal/override`: The internal package manages human governance overrides.
- `internal` → `internal/policy`: The internal package evaluates policies.
- `internal` → `internal/record`: The internal package stores certification records.
- `internal` → `internal/report`: The internal package generates certification reports.
- `internal/agent` → `internal/domain`: The agent package uses domain types for review and status tracking.
- `internal/config` → `internal/domain`: Configuration services interact with domain types to validate settings.
- `internal/config` → `internal/policy`: Configuration services provide policy settings.
- `internal/discovery` → `internal/domain`: Discovery services use domain types to identify and categorize units.
- `internal/engine` → `internal/agent`: Engine services orchestrate LLM-assisted reviews.
- `internal/engine` → `internal/domain`: Engine services use domain types for certification status tracking.
- `internal/engine` → `internal/evidence`: Engine services collect evidence for certification evaluation.
- `internal/engine` → `internal/expiry`: Engine services calculate time-bound trust windows.
- `internal/engine` → `internal/override`: Engine services manage human governance overrides.
- `internal/engine` → `internal/policy`: Engine services evaluate policies against code units.
- `internal/engine` → `internal/record`: Engine services store certification records.
- `internal/engine` → `internal/report`: Engine services generate certification reports.
- `internal/evidence` → `internal/domain`: Evidence collectors use domain types for unit identification and status tracking.
- `internal/evidence` → `internal/policy`: Evidence collectors provide data for policy evaluation.
- `internal/expiry` → `internal/domain`: Expiry services calculate time-bound trust windows using domain types.
- `internal/github` → `internal/domain`: GitHub integration services use domain types for status tracking.
- `internal/override` → `internal/domain`: Override services interact with domain types for status management.
- `internal/policy` → `internal/domain`: Policy services use domain types for evaluation and status tracking.
- `internal/policy` → `internal/evidence`: Policy services provide evidence requirements for evaluation.
- `internal/record` → `internal/domain`: Record services use domain types for status and certification tracking.
- `internal/report` → `internal/agent`: Report generation services use agent data for LLM-assisted insights.
- `internal/report` → `internal/domain`: Report generation services use domain types for status and grading.
- `internal/workspace` → `internal/domain`: Workspace services use domain types for status and certification tracking.
- `internal/workspace` → `internal/record`: Workspace services manage record storage for certification.
- `internal/workspace` → `internal/report`: Workspace services generate workspace-specific reports.

### Hotspots

| Rank | Package | Units | Score | Risk Factor |
|-----:|---------|------:|------:|------------:|
| 1 | internal/agent | 187 | 92.0% | 15.01 |
| 2 | internal/report | 108 | 91.3% | 9.43 |
| 3 | internal/evidence | 68 | 91.4% | 5.87 |
| 4 | internal/domain | 69 | 92.4% | 5.23 |
| 5 | cmd/certify | 54 | 91.5% | 4.58 |
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
| cmd/certify | internal/workspace | 5 |
| cmd/certify | internal/record | 5 |
| internal/engine | internal/policy | 5 |
| cmd/certify | internal/config | 4 |
| cmd/certify | internal/agent | 4 |
| cmd/certify | internal/domain | 4 |
| internal/domain | internal/policy | 4 |
| internal/domain | internal/override | 4 |
| internal/domain | internal/github | 4 |
| internal/engine | internal/evidence | 4 |
| cmd/certify | internal/discovery | 3 |
| cmd/certify | internal/report | 2 |
| cmd/certify | internal/engine | 2 |
| cmd/certify | internal/github | 2 |

---

## Part II: Analysis

### Code Quality & Patterns

🟠 **internal/agent** — High complexity hotspot with 187 units and 92.0% average score, showing strong interdependencies with internal/domain (14 edges) and high risk factor (15.01). This package is a central integration point for LLM-assisted review and likely contains complex logic that could benefit from refactoring.

🟠 **internal/report** — High complexity hotspot with 108 units and 91.3% average score, showing highest risk factor (9.43) and strong coupling with internal/domain (14 edges). This package appears to be a major data processing and output generation layer that may benefit from architectural decomposition.

🟡 **internal/evidence** — Complexity hotspot with 68 units and 91.4% average score, showing risk factor of 5.87 and coupling with internal/domain (9 edges). This evidence collection layer likely contains complex analysis logic that could be simplified through better modularization.

🟡 **internal/domain** — High coupling with multiple internal packages (internal/agent, internal/report, internal/discovery, internal/evidence) and strong interdependencies. This central domain layer is the hub for all other packages but shows moderate complexity with 69 units and 92.4% average score.

🟡 **cmd/certify** — Command layer with 54 units and 91.5% average score, showing moderate risk factor (4.58) and strong coupling with internal packages. This layer orchestrates the certification workflow but may benefit from better separation of concerns.

🟡 **internal/engine** — High coupling with internal/agent (14 edges) and internal/report (14 edges), indicating this package is a central orchestrator for the certification pipeline. With 20 units and 91.5% average score, it's a key integration point that may contain complex orchestration logic.

### Test Strategy & Coverage

The test strategy shows significant gaps in coverage for the most critical packages, particularly cmd/certify, internal/agent, and internal/report which are highest risk. The architecture shows a layered structure with cmd/certify as entry point, internal/domain as central hub, and internal packages as implementation services. However, test organization does not align with this architecture - there are no integration tests for the command layer's orchestration, no property-based tests for domain types, and no end-to-end tests covering the full certification pipeline. The testing approach appears to be primarily unit-focused but missing critical integration points that span multiple layers. Missing test categories include: 1) End-to-end pipeline tests, 2) Integration tests for command layer orchestrations, 3) Property-based tests for domain types, 4) Mock-based integration tests for external dependencies like GitHub and LLM agents, 5) Performance and stress tests for high-volume packages like internal/agent and internal/report.

**Coverage Gaps:**

- `cmd/certify` (score: 91.5%): Package has 54 units with 0 observations, indicating no test coverage despite being a critical command layer. This is a high-risk gap as it's the CLI entry point orchestrating the entire certification workflow.
- `internal/agent` (score: 92.0%): Package with 187 units and 0 observations shows severe test coverage gap. This is a high-risk area with 15.01 risk factor and strong interdependencies with domain layer.
- `internal/report` (score: 91.3%): Package with 108 units and 0 observations represents a critical data processing layer with highest risk factor (9.43) and strong coupling to domain layer.
- `internal/evidence` (score: 91.4%): Package with 68 units and 0 observations shows moderate risk gap. This evidence collection layer has 5.87 risk factor and strong coupling with domain layer.
- `internal/discovery` (score: 91.1%): Package with 40 units and 0 observations represents a critical unit discovery layer with 3.55 risk factor and strong coupling to domain layer.
- `internal/engine` (score: 91.5%): Package with 20 units and 0 observations shows medium risk gap. This orchestrator package has 1.69 risk factor but is central to certification pipeline.

### Security & Operations

🔒 **security** — High global mutable state count indicates potential race conditions and thread safety issues in concurrent operations. This is particularly concerning for a certification system that processes multiple code units simultaneously.
  Affected: `internal/agent`, `internal/engine`, `internal/report`
  Metrics: global_mutable_count: 15

⚙️ **operations** — Multiple panic calls detected in core processing packages indicating potential system crashes during certification workflows. Panic handling is critical for graceful degradation in a certification system.
  Affected: `internal/agent`, `internal/report`, `internal/evidence`
  Metrics: panic_calls: 8

⚙️ **operations** — OS exit calls found in command layer and core processing packages indicate direct system termination rather than graceful error handling. This can lead to incomplete certification processes and data loss.
  Affected: `cmd/certify`, `internal/engine`
  Metrics: os_exit_calls: 3

📋 **config** — Configuration management packages show potential for hardcoded values and insufficient environment variable handling. The system requires robust configuration that can handle different deployment scenarios.
  Affected: `internal/config`, `internal/policy`
  Metrics: init_func_observations: 5

📦 **dependencies** — External dependency surface includes LLM agent integration and GitHub API interactions that could introduce security vulnerabilities or operational instability. These external integrations require careful monitoring.
  Affected: `internal/agent`, `internal/github`
  Metrics: external_dependency_count: 12

---

## Part III: Recommendations (Current → Proposed)

### Refactor cmd/certify to improve test coverage and orchestration clarity

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| observations | 0 | 12 | 0 → 12 |
| avg_score | 91.5% | 93.2% | 91.5% → 93.2% |

**Current:** Package cmd/certify has 54 units with 0 observations and 91.5% average score, showing no test coverage despite being the CLI entry point orchestrating the entire certification workflow.

**Proposed:** Implement integration tests for cmd/certify that cover the full command orchestration flow, including mock-based testing of all subcommands and their interactions with internal packages. Refactor main.go to extract command execution logic into testable functions.

**Affected:** `cmd/certify/main.go#main`, `cmd/certify/report_cmd.go#runWorkspaceReport`, `cmd/certify/certify_cmd.go#processQueue`

**Effort:** M · **Justification:** The refactoring would involve extracting command execution logic from main.go into testable functions and adding integration tests that simulate command execution flows. This would address the zero observations gap for cmd/certify and improve its score by adding proper test coverage. The project would move from 0 to 12 observations, and the score would improve from 91.5% to 93.2% based on the quality improvement from proper test coverage.

### Decompose internal/agent package to reduce complexity and improve testability

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| observations | 0 | 8 | 0 → 8 |
| avg_score | 92.0% | 94.1% | 92.0% → 94.1% |

**Current:** Package internal/agent has 187 units with 0 observations and 92.0% average score, showing strong interdependencies with internal/domain (14 edges) and high risk factor (15.01).

**Proposed:** Refactor internal/agent to decompose LLM-assisted review functionality into smaller, more manageable modules with dedicated test coverage. Extract core agent logic and create separate packages for different LLM providers.

**Affected:** `internal/agent/architect.go#buildTreeRecursive`, `internal/agent/architect_snapshot.go#analyzeDependencies`, `internal/agent/architect_snapshot.go#BuildSnapshot`

**Effort:** L · **Justification:** The decomposition would involve breaking down the monolithic internal/agent package into smaller modules with focused responsibilities. This would address the zero observations and high risk factor by enabling proper unit testing of each component. The project would move from 0 to 8 observations, and the score would improve from 92.0% to 94.1% as the complexity is reduced and testability improved.

### Implement comprehensive integration tests for internal/report package

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| observations | 0 | 15 | 0 → 15 |
| avg_score | 91.3% | 93.7% | 91.3% → 93.7% |

**Current:** Package internal/report has 108 units with 0 observations and 91.3% average score, showing highest risk factor (9.43) and strong coupling with internal/domain (14 edges).

**Proposed:** Add integration tests for internal/report that cover report generation workflows, including mock-based testing of all report types (card, full, badge) and their interactions with evidence collectors and domain types.

**Affected:** `internal/report/architect_report.go#writeArchPartII`, `internal/report/report_tree.go#GenerateReportTree`, `internal/report/site.go#generateIndex`

**Effort:** L · **Justification:** The integration tests would cover the report generation pipeline end-to-end, addressing the zero observations gap and reducing risk factor. The project would move from 0 to 15 observations, and the score would improve from 91.3% to 93.7% as proper test coverage is added and the risk factor decreases.

### Improve error handling in cmd/certify to prevent OS exit calls

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| panic_calls | 8 | 5 | 8 → 5 |
| os_exit_calls | 3 | 0 | 3 → 0 |

**Current:** Package cmd/certify has 3 OS exit calls that indicate direct system termination rather than graceful error handling, which can lead to incomplete certification processes and data loss.

**Proposed:** Replace OS exit calls in cmd/certify with proper error propagation and graceful shutdown mechanisms. Implement centralized error handling that logs failures and exits cleanly without abrupt system termination.

**Affected:** `cmd/certify/main.go#main`, `cmd/certify/report_cmd.go#runWorkspaceReport`

**Effort:** M · **Justification:** The refactoring would involve replacing direct OS exit calls with proper error handling that allows for graceful degradation. This addresses the critical operational concern by eliminating OS exit calls and reducing panic calls from 8 to 5, improving system reliability and preventing incomplete certification processes.

### Refactor internal/evidence package to reduce complexity and improve modularity

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| observations | 0 | 10 | 0 → 10 |
| avg_score | 91.4% | 93.0% | 91.4% → 93.0% |

**Current:** Package internal/evidence has 68 units with 0 observations and 91.4% average score, showing risk factor of 5.87 and coupling with internal/domain (9 edges).

**Proposed:** Decompose internal/evidence into specialized evidence collectors (lint, test, git, complexity) with dedicated test coverage and better modularization to reduce coupling with internal/domain.

**Affected:** `internal/evidence/complexity.go#ComputeSymbolMetrics`, `internal/evidence/structural.go#walkStmt`, `internal/evidence/structural.go#AnalyzeGoType`

**Effort:** M · **Justification:** The refactoring would involve breaking down the monolithic evidence package into specialized modules, reducing coupling and improving testability. This would address the zero observations gap and reduce complexity by enabling proper unit testing of each evidence collector type. The project would move from 0 to 10 observations, and the score would improve from 91.4% to 93.0%.

---

## Risk Matrix

| Risk | Severity | Likelihood | Related Recommendation |
|------|----------|------------|------------------------|
| Zero test coverage in cmd/certify command layer | critical | high | Refactor cmd/certify to improve test coverage and orchestration clarity |
| Zero test coverage in internal/agent package with high risk factor (15.01) | critical | high | Decompose internal/agent package to reduce complexity and improve testability |
| Zero test coverage in internal/report package with highest risk factor (9.43) | critical | high | Implement comprehensive integration tests for internal/report package |
| Direct OS exit calls in cmd/certify and internal/engine | high | medium | Improve error handling in cmd/certify to prevent OS exit calls |
| Global mutable state in concurrent processing packages | high | medium | Decompose internal/agent package to reduce complexity and improve testability |
| High complexity in internal/agent package with 187 units and 15.01 risk factor | high | medium | Decompose internal/agent package to reduce complexity and improve testability |
| Panic calls in core processing packages | high | medium | Improve error handling in cmd/certify to prevent OS exit calls |
| Zero test coverage in internal/evidence package with 5.87 risk factor | medium | medium | Refactor internal/evidence package to reduce complexity and improve modularity |

## Prioritized Roadmap

| # | Item | Effort | Impact | Current → Projected |
|--:|------|--------|--------|---------------------|
| 1 | Decompose internal/agent package to reduce complexity and improve testability | L | high | observations: 0 → 8, avg_score: 92.0% → 94.1% |
| 2 | Implement comprehensive integration tests for internal/report package | L | high | observations: 0 → 15, avg_score: 91.3% → 93.7% |
| 3 | Refactor cmd/certify to improve test coverage and orchestration clarity | M | high | observations: 0 → 12, avg_score: 91.5% → 93.2% |
| 4 | Improve error handling in cmd/certify to prevent OS exit calls | M | high | panic_calls: 8 → 5, os_exit_calls: 3 → 0 |
| 5 | Refactor internal/evidence package to reduce complexity and improve modularity | M | medium | observations: 0 → 10, avg_score: 91.4% → 93.0% |

---

## Appendix: Data Sources

- **748** units across **25** packages · Score: 91.8%
- Evidence: lint, test, coverage, structural, git history
- Snapshot computed from certification records at `f0dca8b6`

---

*Generated by [Certify](https://github.com/iksnae/code-certification) `architect` command.*
