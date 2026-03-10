# 🏗 Architectural Review — iksnae/code-certification

**Generated:** 2026-03-10 15:42 · **Commit:** `1d6d4ea` · **Model:** `qwen/qwen3-coder-30b` · **Tokens:** 17385 · **Duration:** 2m28s · **Phases:** 3/6

## Executive Summary

*Executive summary not available (synthesis phase did not complete).*

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
| cmd/certify | internal/agent | 4 |
| cmd/certify | internal/config | 4 |
| internal/domain | internal/override | 4 |
| internal/domain | internal/github | 4 |
| internal/domain | internal/policy | 4 |
| internal/engine | internal/evidence | 4 |
| cmd/certify | internal/discovery | 3 |
| cmd/certify | internal/report | 2 |
| cmd/certify | internal/engine | 2 |
| cmd/certify | internal/github | 2 |
| internal/agent | internal/report | 2 |

---

## Part II: Analysis

### Code Quality & Patterns

🟠 **cmd/certify** — Contains multiple anti-patterns including init functions, global mutable state, and error handling issues

🟠 **internal/agent** — High risk due to complexity and error handling issues with 20 observations including errors_ignored

🟡 **internal/domain** — Moderate complexity hotspot with 34 observations including global mutable state and init functions

🟡 **internal/engine** — Contains code smells such as high complexity, excessive todo_count, and error handling issues

🟡 **internal/discovery** — Error handling and global mutable state issues present in 16 observations

🟡 **internal/evidence** — Error handling and todo_count issues with 21 observations

🟡 **internal/policy** — Error handling and todo_count issues with 7 observations

🟢 **internal/queue** — Error handling issues with 1 observation

🟡 **internal/record** — Error handling and todo_count issues with 6 observations

🟢 **internal/report** — Error handling and function line count issues with 2 observations

### Test Strategy & Coverage

The test strategy shows significant gaps in coverage for high-risk packages. The cmd/certify package has the lowest score (78.4%) with 106 observations, indicating critical anti-patterns that likely lack proper test coverage. Internal/agent and internal/domain packages also show high observation counts with moderate scores, suggesting insufficient unit testing. The architecture shows clear dependencies from cmd/certify to internal packages, but test organization appears to be missing integration and property-based testing categories. There's a mismatch between the layered architecture (cmd -> internal -> domain) and test coverage, with most packages having zero observations but also zero test coverage indicators. The strategy should focus on: 1) Adding integration tests for the main command flow, 2) Implementing property-based testing for domain logic, 3) Ensuring proper coverage for high-risk packages, and 4) Establishing test categories that align with the architecture layers.

**Coverage Gaps:**

- `cmd/certify` (score: 78.4%): High observation count (106) with low test coverage, contains init functions and global mutable state issues
- `internal/agent` (score: 86.0%): High observation count (20) with complexity and error handling issues, likely insufficient unit test coverage
- `internal/domain` (score: 83.1%): Moderate observation count (34) with init functions and global mutable state issues, potential test coverage gap
- `internal/engine` (score: 85.3%): Medium observation count (8) with complexity and todo issues, likely missing integration tests
- `internal/discovery` (score: 84.8%): Medium observation count (16) with error handling issues, may lack property-based or integration tests
- `internal/evidence` (score: 86.4%): Medium observation count (21) with error handling and todo issues, potential missing test categories

### Security & Operations

🔒 **general** — {
  "concerns": [
    {
      "area": "security",
      "description": "High global mutable state count in critical packages including cmd/certify and internal/agent, creating potential race conditions and security vulnerabilities",
      "affected_packages": ["cmd/certify", "internal/agent", "internal/domain"],
      "metrics": {
        "global_mutable_count": 52
      }
    },
    {
      "area": "operations",
      "description": "Multiple packages contain init functions that can cause unpredictable initialization order and make system behavior harder to reason about",
      "affected_packages": ["cmd/certify", "internal/agent", "internal/domain"],
      "metrics": {
        "has_init_func": 70
      }
    },
    {
      "area": "config",
      "description": "Command line interface in cmd/certify contains hardcoded configuration values and lacks proper environment variable handling",
      "affected_packages": ["cmd/certify"],
      "metrics": {
        "global_mutable_count": 7,
        "has_init_func": 1
      }
    },
    {
      "area": "dependencies",
      "description": "High dependency coupling between cmd/certify and internal packages creates operational fragility with potential for cascading failures",
      "affected_packages": ["cmd/certify", "internal/agent", "internal/domain"],
      "metrics": {
        "coupling_edges": 14
      }
    },
    {
      "area": "operations",
      "description": "Error handling issues in critical packages including cmd/certify and internal/agent where errors are ignored, potentially masking operational problems",
      "affected_packages": ["cmd/certify", "internal/agent", "internal/discovery", "internal/evidence", "internal/engine", "internal/policy", "internal/queue", "internal/record", "internal/report"],
      "metrics": {
        "errors_ignored": 28
      }
    }
  ]
}

---

## Part III: Recommendations (Current → Proposed)

*No recommendations generated (phase did not complete).*

---

## Risk Matrix

*No risk matrix (synthesis phase did not complete).*

## Prioritized Roadmap

*No roadmap (synthesis phase did not complete).*

## ⚠️ Incomplete Phases

- Phase 1 (Architecture Narration): openrouter: HTTP request: Post "http://localhost:1234/v1/chat/completions": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
- Phase 5 (Comparative Recommendations): openrouter: HTTP request: Post "http://localhost:1234/v1/chat/completions": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
- Phase 6 (Synthesis & Roadmap): openrouter: HTTP request: Post "http://localhost:1234/v1/chat/completions": context deadline exceeded (Client.Timeout exceeded while awaiting headers)

---

## Appendix: Data Sources

- **615** units across **24** packages · Score: 85.4%
- Evidence: lint, test, coverage, structural, git history
- Snapshot computed from certification records at `1d6d4ea`

---

*Generated by [Certify](https://github.com/iksnae/code-certification) `architect` command.*
