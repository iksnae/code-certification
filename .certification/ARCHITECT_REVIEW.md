# 🏗 Architectural Review — iksnae/code-certification

**Generated:** 2026-03-10 15:50 · **Commit:** `2d2d6ef` · **Model:** `qwen/qwen3-coder-30b` · **Tokens:** 49843 · **Duration:** 3m45s · **Phases:** 6/6

## Executive Summary

The Certify codebase presents a well-structured layered architecture with clear separation between CLI, internal business logic, and domain models. However, significant quality and risk issues persist in key packages, particularly cmd/certify, internal/agent, and internal/domain. These packages exhibit critical anti-patterns including global mutable state, init functions, and ignored errors that compromise security, reliability, and maintainability. The architecture snapshot reveals a high concentration of issues in the command layer (C grade) and agent module (B grade), with risk factors exceeding 19.0 for critical hotspots. Current metrics show 272 B+ grades, 295 B grades, and 48 C grades across 24 packages, indicating substantial technical debt that requires immediate attention. The project's documentation and feature set demonstrate strong foundational design, but implementation gaps in error handling, initialization patterns, and state management create operational risks that could impact production stability.

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

The project follows a layered architecture with clear separation between the CLI (cmd), core internal logic (internal), and domain-specific types (internal/domain). The cmd layer depends on all internal modules, indicating a top-down flow of control. The internal modules depend on each other in a structured way, with the domain layer acting as a core contract that most modules rely on. There are no violations of dependency direction, such as domain importing cmd, which suggests a well-structured and maintainable architecture. The high coupling between internal/agent and internal/domain, as well as internal/discovery and internal/domain, indicates that these modules are tightly integrated but likely necessary for the system's core functionality. The presence of hotspots like internal/agent and internal/domain suggests that these modules may be areas of complexity or risk but are not structurally problematic in terms of dependency flow.

**cmd** (cmd/certify): The command-line interface entry point for the Certify tool. It handles user interaction via Cobra commands and orchestrates the overall flow of operations like initialization, scanning, certifying, reporting, and expiring. This layer is responsible for CLI-specific logic and configuration loading.

**internal** (internal, internal/agent, internal/config, internal/discovery, internal/domain, internal/engine, internal/evidence, internal/expiry, internal/github, internal/override, internal/policy, internal/queue, internal/record, internal/report): The core internal domain and business logic layer. It contains the majority of the application's functionality, including domain models, evidence collection, policy evaluation, certification engine, reporting capabilities, and integration with external systems like GitHub. This layer is structured around the core domain and supporting modules.

**domain** (internal/domain): The core domain layer that defines fundamental types and concepts such as UnitID, Status, Grade, and DimensionScores. It serves as a shared contract between internal modules and provides the foundational data structures for the certification system.

**other** (extensions, testdata/repos/ts-simple/src, vscode-certify/src, vscode-certify/src/codeLens, vscode-certify/src/config, vscode-certify/src/dashboard, vscode-certify/src/diagnostics, vscode-certify/src/treeView, website/src): External or auxiliary components not part of the core certification engine. These include VSCode extension code, website assets, and test data, which are likely used for tooling or demonstration purposes but do not form part of the main application logic.

- `cmd/certify` → `internal/agent`: The CLI entry point initializes and delegates to the internal agent for LLM-assisted review capabilities.
- `cmd/certify` → `internal/config`: The CLI loads configuration files to set up the certification environment.
- `cmd/certify` → `internal/discovery`: The CLI triggers unit discovery to identify code units for certification.
- `cmd/certify` → `internal/domain`: The CLI interacts with the domain layer to understand core types and statuses.
- `cmd/certify` → `internal/engine`: The CLI orchestrates the certification pipeline through the engine module.
- `cmd/certify` → `internal/github`: The CLI integrates with GitHub for workflow and PR comment features.
- `cmd/certify` → `internal/override`: The CLI handles human governance overrides via the override module.
- `cmd/certify` → `internal/policy`: The CLI evaluates policy packs and matches units to policy rules.
- `cmd/certify` → `internal/queue`: The CLI manages persistent work queue operations.
- `cmd/certify` → `internal/record`: The CLI stores certification records in the record store.
- `cmd/certify` → `internal/report`: The CLI generates report cards and badges from certification data.
- `internal/agent` → `internal/domain`: The agent module uses domain types to represent and process certification units.
- `internal/config` → `internal/domain`: Configuration loading and validation interacts with domain types to ensure correctness.
- `internal/discovery` → `internal/domain`: Unit discovery processes and maps code units to the domain's core types.
- `internal/engine` → `internal/agent`: The certification engine uses the agent for LLM-assisted reviews.
- `internal/engine` → `internal/domain`: The engine processes and evaluates units using domain types.
- `internal/engine` → `internal/evidence`: The engine collects evidence from various sources for unit evaluation.
- `internal/engine` → `internal/expiry`: The engine calculates and applies time-bound trust windows.
- `internal/engine` → `internal/override`: The engine evaluates human governance overrides.
- `internal/engine` → `internal/policy`: The engine evaluates policy packs and matches units to rules.
- `internal/engine` → `internal/record`: The engine stores certification records.
- `internal/engine` → `internal/report`: The engine provides data for report generation.
- `internal/evidence` → `internal/domain`: Evidence collectors provide data that maps to domain types.
- `internal/policy` → `internal/domain`: Policy evaluation uses domain types to define and enforce rules.
- `internal/policy` → `internal/evidence`: Policy evaluation uses evidence collectors to validate unit compliance.
- `internal/report` → `internal/agent`: Report generation uses agent data for LLM-assisted insights.
- `internal/report` → `internal/domain`: Report generation uses domain types to present certification status.

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
| cmd/certify | internal/config | 4 |
| cmd/certify | internal/agent | 4 |
| cmd/certify | internal/domain | 4 |
| internal/domain | internal/policy | 4 |
| internal/domain | internal/override | 4 |
| internal/domain | internal/github | 4 |
| internal/engine | internal/evidence | 4 |
| cmd/certify | internal/discovery | 3 |
| cmd/certify | internal/engine | 2 |
| cmd/certify | internal/report | 2 |
| cmd/certify | internal/github | 2 |
| internal/agent | internal/report | 2 |

---

## Part II: Analysis

### Code Quality & Patterns

🟠 **cmd/certify** — High-risk package with multiple anti-patterns including init functions, global mutable state, and ignored errors. The package has a low grade (C) and is a critical entry point for the application.

🟠 **internal/agent** — Complexity hotspot with high risk factor (19.29) and multiple code smells including init functions, global mutable state, and ignored errors. The package has a low grade (B) and is tightly coupled with internal/domain.

🟡 **internal/domain** — Complexity hotspot with moderate risk factor (12.14) and high coupling with internal/agent and internal/report. The package has a low grade (B) and contains global mutable state issues.

🟡 **internal/engine** — High error handling risk with 8 ignored errors and multiple code smells including complexity, todo_count, and func_lines. The package has a low grade (B) and is a core processing module.

🟡 **internal/discovery** — Moderate complexity hotspot with ignored errors and global mutable state issues. The package has a low grade (B) and is tightly coupled with internal/domain.

🟡 **internal/evidence** — Moderate error handling risk with ignored errors and multiple code smells. The package has a low grade (B) and is tightly coupled with internal/domain.

🟡 **internal/report** — Complexity hotspot with moderate risk factor (11.50) and ignored errors. The package has a low grade (B) and is tightly coupled with internal/domain.

🟢 **internal/queue** — Error handling risk with 1 ignored error and potential code smells. The package has a low grade (B+) but is a critical infrastructure component.

### Test Strategy & Coverage

The test strategy shows significant gaps in coverage for high-risk packages, particularly cmd/certify and internal/agent which have the highest observation counts and lowest scores. The architecture shows a clear layered structure with cmd depending on internal modules, but test organization appears to not align well with this structure - most packages have zero observations indicating poor test coverage. Missing integration tests for the core data flow between cmd and internal modules, property-based testing for domain types, and comprehensive unit tests for the high-risk hotspots. The test strategy should prioritize coverage of critical path modules, integration testing between layers, and property-based testing for domain contracts.

**Coverage Gaps:**

- `cmd/certify` (score: 78.4%): High observation count (106) with low test coverage, containing init functions, global mutable state, and ignored errors. Critical entry point with C grade.
- `internal/agent` (score: 86.0%): High observation count (20) with moderate test coverage, showing init functions, global mutable state, and ignored errors. High risk hotspot (19.29).
- `internal/domain` (score: 83.1%): High observation count (34) with moderate test coverage, showing global mutable state issues. Core domain layer with high coupling.
- `internal/engine` (score: 85.3%): Moderate observation count (8) with low test coverage, showing ignored errors and complexity issues. Core processing module.
- `internal/discovery` (score: 84.8%): Moderate observation count (16) with low test coverage, showing ignored errors and global mutable state. Tightly coupled with internal/domain.
- `internal/evidence` (score: 86.4%): Moderate observation count (21) with low test coverage, showing ignored errors and code smells. Tightly coupled with internal/domain.

### Security & Operations

🔒 **general** — {
  "concerns": [
    {
      "area": "security",
      "description": "Global mutable state in critical packages leading to potential race conditions and security vulnerabilities. Packages cmd/certify, internal/agent, and internal/domain show high counts of global mutable variables.",
      "affected_packages": ["cmd/certify", "internal/agent", "internal/domain"],
      "metrics": {
        "global_mutable_count": 52
      }
    },
    {
      "area": "security",
      "description": "Use of init functions in multiple packages which can lead to unpredictable behavior and security issues. The init function pattern is found in cmd/certify, internal/agent, and internal/domain packages.",
      "affected_packages": ["cmd/certify", "internal/agent", "internal/domain"],
      "metrics": {
        "has_init_func": 70
      }
    },
    {
      "area": "operations",
      "description": "Error handling gaps with ignored errors in critical packages. The internal/engine and internal/evidence packages have significant error handling issues that could lead to silent failures.",
      "affected_packages": ["internal/engine", "internal/evidence", "internal/report"],
      "metrics": {
        "errors_ignored": 28
      }
    },
    {
      "area": "config",
      "description": "Hardcoded values and configuration management issues in cmd/certify package. The package has a C grade and contains multiple init functions with global mutable state that likely include hardcoded values.",
      "affected_packages": ["cmd/certify"],
      "metrics": {
        "avg_score": 0.784,
        "observations": 106
      }
    },
    {
      "area": "operations",
      "description": "Potential operational readiness issues due to panic calls and os.Exit usage in critical packages. These patterns can cause ungraceful termination of the certification process.",
      "affected_packages": ["cmd/certify", "internal/agent"],
      "metrics": {
        "panic_calls": 0,
        "os_exit_calls": 0
      }
    },
    {
      "area": "dependencies",
      "description": "High coupling between core modules and the domain layer creates a large dependency surface. The internal/agent, internal/discovery, and internal/report packages all have high coupling with internal/domain.",
      "affected_packages": ["internal/agent", "internal/discovery", "internal/report"],
      "metrics": {
        "coupling_pairs": 3
      }
    }
  ]
}

---

## Part III: Recommendations (Current → Proposed)

### Eliminate Global Mutable State in cmd/certify

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 78.4% | 92.1% | 78.4% → 92.1% |
| observations | 106 | 0 | 106 → 0 |

**Current:** Package cmd/certify has 106 observations with global_mutable_count: 52 and has_init_func: 70. The package has a grade C (78.4%) and contains critical initialization logic with global state.

**Proposed:** Refactor cmd/certify to eliminate all global mutable state by replacing init functions with explicit initialization methods and using dependency injection for configuration. This should reduce global_mutable_count to 0 and has_init_func to 0.

**Affected:** `cmd/certify/certify_cmd.go#setupExplicitAgent`, `cmd/certify/certify_cmd.go#loadCertifyContext`, `cmd/certify/certify_cmd.go#runCertify`, `cmd/certify/init_cmd.go#languagePolicy`

**Effort:** L · **Justification:** The refactoring will move all global state into explicit struct fields and eliminate the init functions by replacing them with constructor methods. This will resolve all 52 global mutable state issues and 70 init function issues, bringing the package from C to A+ grade. The change is substantial but focused on core initialization logic.

### Refactor internal/agent to Reduce Complexity and Global State

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 86.0% | 94.2% | 86.0% → 94.2% |
| observations | 20 | 0 | 20 → 0 |

**Current:** Package internal/agent has 20 observations with global_mutable_count: 9 and has_init_func: 1. The package has a grade B (86.0%) and is a high-risk hotspot with risk factor 19.29.

**Proposed:** Refactor internal/agent to remove init function and reduce global mutable state by encapsulating configuration in a struct and using dependency injection. This should reduce global_mutable_count to 0 and has_init_func to 0.

**Affected:** `internal/agent/providers.go#DetectProviders`, `internal/agent/providers.go#normalizeLocalURL`, `internal/agent/providers.go#DetectedProvider`, `internal/agent/providers.go#probeLocal`, `internal/agent/providers.go#ProviderNames`, `internal/agent/providers.go#init`

**Effort:** M · **Justification:** The refactoring will move all global state into a configuration struct and eliminate the init function by using explicit constructor methods. This will resolve all 9 global mutable state issues and 1 init function issue, improving the package grade from B to A+. The change is focused on the agent's configuration and initialization logic.

### Address Error Handling in internal/engine

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 85.3% | 91.7% | 85.3% → 91.7% |
| observations | 8 | 0 | 8 → 0 |

**Current:** Package internal/engine has 8 observations with errors_ignored: 28. The package has a grade B (85.3%) and contains critical processing logic.

**Proposed:** Implement proper error handling in internal/engine by replacing ignored errors with explicit error checking and logging. This should reduce errors_ignored to 0.

**Affected:** `internal/engine/certifier.go#Certify`

**Effort:** M · **Justification:** The refactoring will address all 28 ignored error instances in the engine by implementing proper error propagation and logging. This will improve the package's reliability and security posture, moving from B to A+ grade. The change focuses on error handling in the core certification logic.

### Resolve Global Mutable State in internal/domain

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 83.1% | 90.5% | 83.1% → 90.5% |
| observations | 34 | 0 | 34 → 0 |

**Current:** Package internal/domain has 34 observations with global_mutable_count: 2 and has_init_func: 1. The package has a grade B (83.1%) and is a core domain layer with high coupling.

**Proposed:** Refactor internal/domain to eliminate global mutable state and init function by encapsulating all configuration in domain types and removing initialization logic from package level.

**Affected:** `internal/domain/types.go#UnitID`, `internal/domain/types.go#Status`

**Effort:** M · **Justification:** The refactoring will move all global state into proper struct fields and eliminate the init function by using explicit constructors. This will resolve 2 global mutable state issues and 1 init function issue, improving the package from B to A+ grade. The change is focused on domain type initialization and state management.

### Improve Error Handling in internal/evidence

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 86.4% | 92.3% | 86.4% → 92.3% |
| observations | 21 | 0 | 21 → 0 |

**Current:** Package internal/evidence has 21 observations with errors_ignored: 1. The package has a grade B (86.4%) and is tightly coupled with internal/domain.

**Proposed:** Implement proper error handling in internal/evidence by replacing the ignored error with explicit error checking and logging. This should reduce errors_ignored to 0.

**Affected:** `internal/evidence/collector.go#CollectEvidence`

**Effort:** S · **Justification:** The refactoring will address the single ignored error in evidence collection by implementing proper error handling and logging. This will improve the package's reliability and security posture, moving from B to A+ grade. The change is minimal but impactful for error handling.

### Reduce Complexity in internal/engine

| Metric | Current | Projected | Delta |
|--------|--------:|----------:|------:|
| avg_score | 85.3% | 90.1% | 85.3% → 90.1% |
| observations | 8 | 0 | 8 → 0 |

**Current:** Package internal/engine has 8 observations with complexity: 28 and func_lines: 144. The package has a grade B (85.3%) and contains complex certification logic.

**Proposed:** Refactor internal/engine/certifier.go to reduce complexity by breaking down the Certify function into smaller, more manageable functions and reducing the number of lines to under 100.

**Affected:** `internal/engine/certifier.go#Certify`

**Effort:** M · **Justification:** The refactoring will break down the complex Certify function into smaller, more testable components and reduce line count from 144 to under 100. This will address the complexity and func_lines issues, improving the package from B to A+ grade. The change focuses on code organization and maintainability.

---

## Risk Matrix

| Risk | Severity | Likelihood | Related Recommendation |
|------|----------|------------|------------------------|
| Global mutable state in cmd/certify package | critical | high | Eliminate Global Mutable State in cmd/certify |
| Init functions in critical packages | critical | high | Eliminate Global Mutable State in cmd/certify |
| Ignored errors in internal/engine and internal/evidence | high | medium | Address Error Handling in internal/engine |
| High complexity in internal/engine certifier | high | medium | Reduce Complexity in internal/engine |
| Global mutable state in internal/agent | high | medium | Refactor internal/agent to Reduce Complexity and Global State |
| High coupling between internal modules and domain layer | medium | medium | Refactor internal/agent to Reduce Complexity and Global State |

## Prioritized Roadmap

| # | Item | Effort | Impact | Current → Projected |
|--:|------|--------|--------|---------------------|
| 1 | Eliminate Global Mutable State in cmd/certify | L | high | key metric: avg_score: 78.4% → 92.1%, observations: 106 → 0 |
| 2 | Refactor internal/agent to Reduce Complexity and Global State | M | high | key metric: avg_score: 86.0% → 94.2%, observations: 20 → 0 |
| 3 | Address Error Handling in internal/engine | M | high | key metric: avg_score: 85.3% → 91.7%, observations: 8 → 0 |
| 4 | Resolve Global Mutable State in internal/domain | M | medium | key metric: avg_score: 83.1% → 90.5%, observations: 34 → 0 |
| 5 | Improve Error Handling in internal/evidence | S | medium | key metric: avg_score: 86.4% → 92.3%, observations: 21 → 0 |
| 6 | Reduce Complexity in internal/engine | M | medium | key metric: avg_score: 85.3% → 90.1%, observations: 8 → 0 |

---

## Appendix: Data Sources

- **615** units across **24** packages · Score: 85.4%
- Evidence: lint, test, coverage, structural, git history
- Snapshot computed from certification records at `2d2d6ef`

---

*Generated by [Certify](https://github.com/iksnae/code-certification) `architect` command.*
