# 🟢 Certify — Report Card

**Repository:** `iksnae/code-certification`  
**Commit:** `bdcd842`  
**Generated:** 2026-03-09T20:13:36  

---

## Summary

| Metric | Value |
|--------|-------|
| **Overall Grade** | 🟢 **B+** |
| **Overall Score** | 89.0% |
| **Total Units** | 542 |
| **Passing** | 542 |
| **Failing** | 0 |
| **Pass Rate** | 100.0% |
| **Observations** | 0 |
| **Expired** | 0 |

## Grade Distribution

| Grade | Count | % | Bar |
|:-----:|------:|----:|-----|
| B+ | 542 | 100.0% | ██████████████████████████████████████████████████ |

## Dimension Averages

| Dimension | Score | Bar |
|-----------|------:|-----|
| architectural_fitness | 85.0% | █████████████████░░░ |
| change_risk | 90.0% | █████████████████░░░ |
| correctness | 95.0% | ██████████████████░░ |
| maintainability | 93.6% | ██████████████████░░ |
| operational_quality | 85.0% | █████████████████░░░ |
| performance_appropriateness | 85.0% | █████████████████░░░ |
| readability | 92.6% | ██████████████████░░ |
| security | 85.0% | █████████████████░░░ |
| testability | 90.0% | █████████████████░░░ |

## 🤖 AI Insights

*Powered by `qwen/qwen3-coder-30b` — 542 units analyzed*

### Top Suggestions

- 💡 Consider adding more descriptive comments for complex logic paths ×12
- 💡 Verify that all edge cases are covered by existing tests ×7
- 💡 Address the TODOs in the code ×7
- 💡 Add more comprehensive test coverage for edge cases ×7
- 💡 Evaluate if the current complexity (5) could be reduced through refactoring ×7
- 💡 Consider refactoring to reduce complexity from 9 ×7
- 💡 Verify that all grade emojis are properly localized for internationalization ×5
- 💡 Add unit tests for edge cases like empty files or invalid formats ×5
- 💡 Consider adding JSDoc comments for better documentation ×5
- 💡 Evaluate if the emoji constants could be exported from a dedicated constants file ×5
- 💡 Check if error handling could be made more explicit ×5
- 💡 Evaluate if the function name clearly reflects its purpose ×5
- 💡 Consider adding unit tests for edge cases ×5
- 💡 Document the expected file format in comments ×5
- 💡 Consider adding more detailed error handling for file operations ×5
- 💡 Verify token length validation logic ×4
- 💡 Consider adding unit tests for edge cases in the ConnectionTestResult class ×4
- 💡 Verify that the 66% coverage threshold meets project requirements ×4
- 💡 Consider adding godoc comments for better documentation ×4
- 💡 Add documentation comments for public methods and fields ×4

*...and 35 more suggestions across individual units*

## By Language

### go — 🟢 B+ (89.1%)

- **Units:** 475
- **Score range:** 87.2% – 89.4%
- **Grades:** 475×B+

### ts — 🟢 B+ (88.5%)

- **Units:** 67
- **Score range:** 88.3% – 89.4%
- **Grades:** 67×B+

## All Units

### `cmd/certify/` (36 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `certifyContext` | class | B+ | 89.4% | certified | 2026-04-23 |
| `certifyUnit` | method | B+ | 87.2% | certified | 2026-04-23 |
| `collectRepoEvidence` | method | B+ | 89.4% | certified | 2026-04-23 |
| `defaultConfigObj` | function | B+ | 89.4% | certified | 2026-04-23 |
| `filterUnits` | function | B+ | 89.4% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `isLocalURL` | function | B+ | 89.4% | certified | 2026-04-23 |
| `loadCertifyContext` | function | B+ | 88.3% | certified | 2026-04-23 |
| `loadQueue` | function | B+ | 89.4% | certified | 2026-04-23 |
| `printQueueStatus` | method | B+ | 89.4% | certified | 2026-04-23 |
| `printSummary` | method | B+ | 89.4% | certified | 2026-04-23 |
| `processQueue` | method | B+ | 88.3% | certified | 2026-04-23 |
| `runCertify` | function | B+ | 89.4% | certified | 2026-04-23 |
| `saveReportArtifacts` | method | B+ | 89.4% | certified | 2026-04-23 |
| `setupAgent` | function | B+ | 89.4% | certified | 2026-04-23 |
| `setupConservativeAgent` | function | B+ | 89.4% | certified | 2026-04-23 |
| `setupExplicitAgent` | function | B+ | 88.3% | certified | 2026-04-23 |
| `cli_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `generateConfig` | function | B+ | 89.4% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `languagePolicy` | function | B+ | 87.2% | certified | 2026-04-23 |
| `main` | function | B+ | 89.4% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `listFromProvider` | function | B+ | 89.4% | certified | 2026-04-23 |
| `runModels` | function | B+ | 89.4% | certified | 2026-04-23 |
| `detectCommit` | function | B+ | 89.4% | certified | 2026-04-23 |
| `detectRepoName` | function | B+ | 89.4% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `saveBadge` | function | B+ | 89.4% | certified | 2026-04-23 |
| `saveReportCard` | function | B+ | 89.4% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `tryScanSuggestions` | function | B+ | 89.4% | certified | 2026-04-23 |
| `version.go` | file | B+ | 89.4% | certified | 2026-04-23 |

<details>
<summary>certifyContext — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>certifyUnit — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>collectRepoEvidence — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>defaultConfigObj — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>filterUnits — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>isLocalURL — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>loadCertifyContext — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>loadQueue — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>printQueueStatus — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>printSummary — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>processQueue — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>runCertify — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>saveReportArtifacts — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>setupAgent — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>setupConservativeAgent — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>setupExplicitAgent — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary>generateConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues, but has 9 TODOs indicating incomplete functionality
- 💡 Address the 9 TODO items to complete functionality
- 💡 Consider adding more comprehensive test coverage for edge cases
- 💡 Evaluate if the current complexity (4) can be reduced through refactoring

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues, but has 9 TODOs indicating incomplete functionality
- 💡 Address the 9 TODO items to complete functionality
- 💡 Consider adding more comprehensive test coverage for edge cases
- 💡 Evaluate if the current complexity (4) can be reduced through refactoring

</details>

<details>
<summary>languagePolicy — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 75.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 9 exceeds threshold 0
- 🤖 Code passes all tests and linting with no issues, but has 9 TODOs indicating incomplete functionality
- 💡 Address the 9 TODO items to complete functionality
- 💡 Consider adding more comprehensive test coverage for edge cases
- 💡 Evaluate if the current complexity (4) can be reduced through refactoring

</details>

<details>
<summary>main — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and low complexity

</details>

<details>
<summary>listFromProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and low complexity

</details>

<details>
<summary>runModels — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and low complexity

</details>

<details>
<summary>detectCommit — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the function name clearly reflects its purpose
- 💡 Check if error handling could be made more explicit

</details>

<details>
<summary>detectRepoName — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the function name clearly reflects its purpose
- 💡 Check if error handling could be made more explicit

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the function name clearly reflects its purpose
- 💡 Check if error handling could be made more explicit

</details>

<details>
<summary>saveBadge — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the function name clearly reflects its purpose
- 💡 Check if error handling could be made more explicit

</details>

<details>
<summary>saveReportCard — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the function name clearly reflects its purpose
- 💡 Check if error handling could be made more explicit

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues found

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity

</details>

<details>
<summary>tryScanSuggestions — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, all tests passing, and low complexity indicating good quality

</details>

### `extensions/` (18 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `agent-chain.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `agent-team.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `cross-agent.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `damage-control.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `minimal.ts` | file | B+ | 89.4% | certified | 2026-04-23 |
| `pi-pi.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `pure-focus.ts` | file | B+ | 89.4% | certified | 2026-04-23 |
| `purpose-gate.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `session-replay.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `subagent-widget.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `system-select.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `theme-cycler.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `THEME_MAP` | function | B+ | 88.3% | certified | 2026-04-23 |
| `applyExtensionDefaults` | function | B+ | 88.3% | certified | 2026-04-23 |
| `applyExtensionTheme` | function | B+ | 88.3% | certified | 2026-04-23 |
| `tilldone.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `tool-counter-widget.ts` | file | B+ | 88.3% | certified | 2026-04-23 |
| `tool-counter.ts` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider adding more comprehensive unit tests for edge cases
- 💡 Evaluate if the current complexity score of 0 is accurate and document any assumptions
- 💡 Verify that the 66% test coverage threshold meets project quality standards

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>THEME_MAP — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>applyExtensionDefaults — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>applyExtensionTheme — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider addressing the TODOs to ensure complete functionality
- 💡 Evaluate if the current complexity score of 0 is accurate for the code's actual complexity
- 💡 Verify that 66% test coverage is sufficient for production use case

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `internal/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `integration_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with full test coverage and no linting issues

</details>

### `internal/agent/` (127 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `attribution_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `DetectAPIKey` | function | B+ | 89.4% | certified | 2026-04-23 |
| `FormatProviderSummary` | function | B+ | 89.4% | certified | 2026-04-23 |
| `HasAnyProvider` | function | B+ | 89.4% | certified | 2026-04-23 |
| `NewConservativeCoordinator` | function | B+ | 88.3% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `autodetect_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Chat` | method | B+ | 89.4% | certified | 2026-04-23 |
| `CircuitBreaker` | class | B+ | 89.4% | certified | 2026-04-23 |
| `IsOpen` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Name` | method | B+ | 89.4% | certified | 2026-04-23 |
| `NewCircuitBreaker` | function | B+ | 89.4% | certified | 2026-04-23 |
| `AdaptiveMessages` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Chat` | method | B+ | 88.3% | certified | 2026-04-23 |
| `FallbackProvider` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ModelChain` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Name` | method | B+ | 89.4% | certified | 2026-04-23 |
| `NewFallbackProvider` | function | B+ | 89.4% | certified | 2026-04-23 |
| `NewModelChain` | function | B+ | 89.4% | certified | 2026-04-23 |
| `modelPinnedProvider` | class | B+ | 89.4% | certified | 2026-04-23 |
| `fallback_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `ListModels` | function | B+ | 89.4% | certified | 2026-04-23 |
| `ModelInfo` | class | B+ | 89.4% | certified | 2026-04-23 |
| `listOllamaModels` | function | B+ | 88.3% | certified | 2026-04-23 |
| `listOpenAIModels` | function | B+ | 88.3% | certified | 2026-04-23 |
| `ollamaModel` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ollamaTagsResponse` | class | B+ | 89.4% | certified | 2026-04-23 |
| `openAIModel` | class | B+ | 89.4% | certified | 2026-04-23 |
| `openAIModelsResponse` | class | B+ | 89.4% | certified | 2026-04-23 |
| `models_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `APIError` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Chat` | method | B+ | 88.3% | certified | 2026-04-23 |
| `Error` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Name` | method | B+ | 89.4% | certified | 2026-04-23 |
| `NewLocalProvider` | function | B+ | 89.4% | certified | 2026-04-23 |
| `NewOpenRouterProvider` | function | B+ | 89.4% | certified | 2026-04-23 |
| `OpenRouterProvider` | class | B+ | 89.4% | certified | 2026-04-23 |
| `doRequest` | method | B+ | 88.3% | certified | 2026-04-23 |
| `isAPIError` | function | B+ | 89.4% | certified | 2026-04-23 |
| `isAuthError` | function | B+ | 89.4% | certified | 2026-04-23 |
| `isBudgetError` | function | B+ | 89.4% | certified | 2026-04-23 |
| `isRetryable` | function | B+ | 89.4% | certified | 2026-04-23 |
| `openrouter_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Coordinator` | class | B+ | 89.4% | certified | 2026-04-23 |
| `CoordinatorConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `NewCoordinator` | function | B+ | 89.4% | certified | 2026-04-23 |
| `NewPipeline` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Pipeline` | class | B+ | 89.4% | certified | 2026-04-23 |
| `PipelineConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ReviewUnit` | method | B+ | 88.3% | certified | 2026-04-23 |
| `Run` | method | B+ | 88.3% | certified | 2026-04-23 |
| `Stats` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Strategy` | class | B+ | 89.4% | certified | 2026-04-23 |
| `toResult` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Get` | method | B+ | 88.3% | certified | 2026-04-23 |
| `LoadPrompt` | function | B+ | 89.4% | certified | 2026-04-23 |
| `NewPromptRegistry` | function | B+ | 89.4% | certified | 2026-04-23 |
| `PromptRegistry` | class | B+ | 89.4% | certified | 2026-04-23 |
| `PromptTemplate` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Render` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Version` | method | B+ | 89.4% | certified | 2026-04-23 |
| `prompts_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Provider` | class | B+ | 89.4% | certified | 2026-04-23 |
| `provider_multi_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `DetectProviders` | function | B+ | 87.2% | certified | 2026-04-23 |
| `DetectedProvider` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ProviderNames` | function | B+ | 89.4% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `normalizeLocalURL` | function | B+ | 89.4% | certified | 2026-04-23 |
| `probeLocal` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Allow` | method | B+ | 89.4% | certified | 2026-04-23 |
| `NewRateLimiter` | function | B+ | 89.4% | certified | 2026-04-23 |
| `RateLimiter` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Wait` | method | B+ | 89.4% | certified | 2026-04-23 |
| `refill` | method | B+ | 89.4% | certified | 2026-04-23 |
| `ratelimit_test.go` | file | B+ | 89.4% | certified | 2026-04-23 |
| `NewReviewer` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Review` | method | B+ | 87.2% | certified | 2026-04-23 |
| `ReviewInput` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ReviewResult` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Reviewer` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ToEvidence` | method | B+ | 89.4% | certified | 2026-04-23 |
| `ToPrescreenEvidence` | method | B+ | 89.4% | certified | 2026-04-23 |
| `joinModels` | function | B+ | 89.4% | certified | 2026-04-23 |
| `reviewer_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `ModelFor` | method | B+ | 89.4% | certified | 2026-04-23 |
| `NewRouter` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Router` | class | B+ | 89.4% | certified | 2026-04-23 |
| `directAssignment` | method | B+ | 88.3% | certified | 2026-04-23 |
| `router_test.go` | file | B+ | 89.4% | certified | 2026-04-23 |
| `DecisionResponse` | class | B+ | 89.4% | certified | 2026-04-23 |
| `PrescreenResponse` | class | B+ | 89.4% | certified | 2026-04-23 |
| `RemediationResponse` | class | B+ | 89.4% | certified | 2026-04-23 |
| `RemediationStep` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ScoringResponse` | class | B+ | 89.4% | certified | 2026-04-23 |
| `schemas_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Execute` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Name` | method | B+ | 89.4% | certified | 2026-04-23 |
| `NewPrescreenStage` | function | B+ | 89.4% | certified | 2026-04-23 |
| `NewReviewStage` | function | B+ | 89.4% | certified | 2026-04-23 |
| `NewScoringStage` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Stage` | class | B+ | 89.4% | certified | 2026-04-23 |
| `StageInput` | class | B+ | 89.4% | certified | 2026-04-23 |
| `StageResult` | class | B+ | 89.4% | certified | 2026-04-23 |
| `defaultScores` | function | B+ | 89.4% | certified | 2026-04-23 |
| `extractJSON` | function | B+ | 88.3% | certified | 2026-04-23 |
| `looseParseNeedsReview` | function | B+ | 89.4% | certified | 2026-04-23 |
| `prescreenStage` | class | B+ | 89.4% | certified | 2026-04-23 |
| `reviewStage` | class | B+ | 89.4% | certified | 2026-04-23 |
| `scoringStage` | class | B+ | 89.4% | certified | 2026-04-23 |
| `stage_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `RepoSummary` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ScanSuggestion` | class | B+ | 89.4% | certified | 2026-04-23 |
| `SuggestForRepo` | function | B+ | 89.4% | certified | 2026-04-23 |
| `buildSuggestPrompt` | function | B+ | 89.4% | certified | 2026-04-23 |
| `suggest_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `ChatRequest` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ChatResponse` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Choice` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Content` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Message` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ModelConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ResponseFormat` | class | B+ | 89.4% | certified | 2026-04-23 |
| `String` | method | B+ | 89.4% | certified | 2026-04-23 |
| `TaskType` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Usage` | class | B+ | 89.4% | certified | 2026-04-23 |
| `types_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>DetectAPIKey — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage and low complexity

</details>

<details>
<summary>FormatProviderSummary — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage and low complexity

</details>

<details>
<summary>HasAnyProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage and low complexity

</details>

<details>
<summary>NewConservativeCoordinator — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage and low complexity

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage and low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>Chat — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no identified issues

</details>

<details>
<summary>CircuitBreaker — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no identified issues

</details>

<details>
<summary>IsOpen — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no identified issues

</details>

<details>
<summary>Name — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no identified issues

</details>

<details>
<summary>NewCircuitBreaker — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no identified issues

</details>

<details>
<summary>AdaptiveMessages — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>Chat — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>FallbackProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>ModelChain — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>Name — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>NewFallbackProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>NewModelChain — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>modelPinnedProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>ListModels — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality.

</details>

<details>
<summary>ModelInfo — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality.

</details>

<details>
<summary>listOllamaModels — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality.

</details>

<details>
<summary>listOpenAIModels — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality.

</details>

<details>
<summary>ollamaModel — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality.

</details>

<details>
<summary>ollamaTagsResponse — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality.

</details>

<details>
<summary>openAIModel — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality.

</details>

<details>
<summary>openAIModelsResponse — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality.

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>APIError — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary>Chat — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary>Error — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary>Name — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary>NewLocalProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary>NewOpenRouterProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary>OpenRouterProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary>doRequest — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary>isAPIError — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary>isAuthError — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary>isBudgetError — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary>isRetryable — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting or testing issues and low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>Coordinator — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested with full coverage, and passes all static analysis checks

</details>

<details>
<summary>CoordinatorConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested with full coverage, and passes all static analysis checks

</details>

<details>
<summary>NewCoordinator — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested with full coverage, and passes all static analysis checks

</details>

<details>
<summary>NewPipeline — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested with full coverage, and passes all static analysis checks

</details>

<details>
<summary>Pipeline — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested with full coverage, and passes all static analysis checks

</details>

<details>
<summary>PipelineConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>ReviewUnit — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested with full coverage, and passes all static analysis checks

</details>

<details>
<summary>Run — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested with full coverage, and passes all static analysis checks

</details>

<details>
<summary>Stats — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested with full coverage, and passes all static analysis checks

</details>

<details>
<summary>Strategy — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested with full coverage, and passes all static analysis checks

</details>

<details>
<summary>toResult — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested with full coverage, and passes all static analysis checks

</details>

<details>
<summary>Get — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple method with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>LoadPrompt — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple method with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>NewPromptRegistry — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple method with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>PromptRegistry — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple method with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>PromptTemplate — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple method with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>Render — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple method with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>Version — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple method with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>Provider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>DetectProviders — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>DetectedProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>ProviderNames — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>normalizeLocalURL — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>probeLocal — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>Allow — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested, and within complexity limits with no linting issues

</details>

<details>
<summary>NewRateLimiter — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested, and within complexity limits with no linting issues

</details>

<details>
<summary>RateLimiter — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested, and within complexity limits with no linting issues

</details>

<details>
<summary>Wait — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested, and within complexity limits with no linting issues

</details>

<details>
<summary>refill — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested, and within complexity limits with no linting issues

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>NewReviewer — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>Review — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>ReviewInput — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>ReviewResult — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>Reviewer — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>ToEvidence — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>ToPrescreenEvidence — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>joinModels — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>ModelFor — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>NewRouter — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>Router — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>directAssignment — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>DecisionResponse — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>PrescreenResponse — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>RemediationResponse — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>RemediationStep — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>ScoringResponse — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>Execute — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>Name — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>NewPrescreenStage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>NewReviewStage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>NewScoringStage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>Stage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>StageInput — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>StageResult — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>defaultScores — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>extractJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>looseParseNeedsReview — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>prescreenStage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>reviewStage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary>scoringStage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates high quality

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with full test coverage and no linting issues

</details>

<details>
<summary>RepoSummary — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary>ScanSuggestion — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary>SuggestForRepo — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary>buildSuggestPrompt — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>ChatRequest — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality

</details>

<details>
<summary>ChatResponse — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality

</details>

<details>
<summary>Choice — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality

</details>

<details>
<summary>Content — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality

</details>

<details>
<summary>Message — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality

</details>

<details>
<summary>ModelConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality

</details>

<details>
<summary>ResponseFormat — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality

</details>

<details>
<summary>String — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality

</details>

<details>
<summary>TaskType — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality

</details>

<details>
<summary>Usage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `internal/config/` (22 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Load` | function | B+ | 87.2% | certified | 2026-04-23 |
| `LoadFile` | function | B+ | 89.4% | certified | 2026-04-23 |
| `LoadFromDir` | function | B+ | 89.4% | certified | 2026-04-23 |
| `rawAgent` | class | B+ | 89.4% | certified | 2026-04-23 |
| `rawConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `validate` | function | B+ | 89.4% | certified | 2026-04-23 |
| `loader_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `FilterPolicyPacks` | function | B+ | 88.3% | certified | 2026-04-23 |
| `NewPolicyMatcher` | function | B+ | 89.4% | certified | 2026-04-23 |
| `LoadPolicyPack` | function | B+ | 89.4% | certified | 2026-04-23 |
| `LoadPolicyPacks` | function | B+ | 88.3% | certified | 2026-04-23 |
| `parseDimension` | function | B+ | 89.4% | certified | 2026-04-23 |
| `parsePolicyPack` | function | B+ | 88.3% | certified | 2026-04-23 |
| `parseSeverity` | function | B+ | 89.4% | certified | 2026-04-23 |
| `rawPolicyPack` | class | B+ | 89.4% | certified | 2026-04-23 |
| `rawPolicyRule` | class | B+ | 89.4% | certified | 2026-04-23 |
| `policy_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Error` | method | B+ | 89.4% | certified | 2026-04-23 |
| `ValidateConfig` | function | B+ | 88.3% | certified | 2026-04-23 |
| `ValidatePolicyPack` | function | B+ | 88.3% | certified | 2026-04-23 |
| `ValidationError` | class | B+ | 89.4% | certified | 2026-04-23 |
| `validator_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>Load — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary>LoadFile — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary>LoadFromDir — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary>rawAgent — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary>rawConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary>validate — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>FilterPolicyPacks — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>NewPolicyMatcher — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>LoadPolicyPack — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates good quality

</details>

<details>
<summary>LoadPolicyPacks — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates good quality

</details>

<details>
<summary>parseDimension — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates good quality

</details>

<details>
<summary>parsePolicyPack — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates good quality

</details>

<details>
<summary>parseSeverity — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates good quality

</details>

<details>
<summary>rawPolicyPack — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates good quality

</details>

<details>
<summary>rawPolicyRule — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity indicates good quality

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>Error — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary>ValidateConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary>ValidatePolicyPack — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary>ValidationError — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `internal/discovery/` (39 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `DetectLanguages` | function | B+ | 88.3% | certified | 2026-04-23 |
| `DetectedAdapters` | function | B+ | 89.4% | certified | 2026-04-23 |
| `LanguageInfo` | class | B+ | 89.4% | certified | 2026-04-23 |
| `buildLanguageList` | function | B+ | 89.4% | certified | 2026-04-23 |
| `detect_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `ChangedFiles` | function | B+ | 89.4% | certified | 2026-04-23 |
| `DetectMoves` | function | B+ | 89.4% | certified | 2026-04-23 |
| `FilterByPaths` | function | B+ | 88.3% | certified | 2026-04-23 |
| `FilterChanged` | function | B+ | 89.4% | certified | 2026-04-23 |
| `MovedFile` | class | B+ | 89.4% | certified | 2026-04-23 |
| `diff_test.go` | file | B+ | 89.4% | certified | 2026-04-23 |
| `GenericScanner` | class | B+ | 89.4% | certified | 2026-04-23 |
| `NewGenericScanner` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Scan` | method | B+ | 87.2% | certified | 2026-04-23 |
| `matchAny` | function | B+ | 88.3% | certified | 2026-04-23 |
| `GoAdapter` | class | B+ | 89.4% | certified | 2026-04-23 |
| `NewGoAdapter` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Scan` | method | B+ | 88.3% | certified | 2026-04-23 |
| `parseFile` | method | B+ | 88.3% | certified | 2026-04-23 |
| `go_adapter_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Diff` | function | B+ | 88.3% | certified | 2026-04-23 |
| `DiffResult` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Index` | class | B+ | 89.4% | certified | 2026-04-23 |
| `LoadIndex` | function | B+ | 89.4% | certified | 2026-04-23 |
| `NewIndex` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Save` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Units` | method | B+ | 89.4% | certified | 2026-04-23 |
| `indexEntry` | class | B+ | 89.4% | certified | 2026-04-23 |
| `index_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `DeduplicateFileLevel` | function | B+ | 88.3% | certified | 2026-04-23 |
| `Merge` | function | B+ | 88.3% | certified | 2026-04-23 |
| `Scanner` | class | B+ | 89.4% | certified | 2026-04-23 |
| `UnitList` | class | B+ | 89.4% | certified | 2026-04-23 |
| `scanner_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `NewTSAdapter` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Scan` | method | B+ | 88.3% | certified | 2026-04-23 |
| `TSAdapter` | class | B+ | 89.4% | certified | 2026-04-23 |
| `parseFile` | method | B+ | 88.3% | certified | 2026-04-23 |
| `ts_adapter_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>DetectLanguages — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>DetectedAdapters — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested, and within complexity limits with no linting issues

</details>

<details>
<summary>LanguageInfo — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested, and within complexity limits with no linting issues

</details>

<details>
<summary>buildLanguageList — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested, and within complexity limits with no linting issues

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>ChangedFiles — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>DetectMoves — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>FilterByPaths — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>FilterChanged — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>MovedFile — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>GenericScanner — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no identified issues

</details>

<details>
<summary>NewGenericScanner — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no identified issues

</details>

<details>
<summary>Scan — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no identified issues

</details>

<details>
<summary>matchAny — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no identified issues

</details>

<details>
<summary>GoAdapter — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary>NewGoAdapter — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary>Scan — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary>parseFile — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>Diff — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>DiffResult — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>Index — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>LoadIndex — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>NewIndex — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>Save — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>Units — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>indexEntry — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>DeduplicateFileLevel — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary>Merge — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary>Scanner — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary>UnitList — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>NewTSAdapter — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linters with good coverage and minimal complexity

</details>

<details>
<summary>Scan — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linters with good coverage and minimal complexity

</details>

<details>
<summary>TSAdapter — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linters with good coverage and minimal complexity

</details>

<details>
<summary>parseFile — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linters with good coverage and minimal complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `internal/domain/` (69 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `AgentConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `AnalyzerConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `CertificationMode` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Config` | class | B+ | 89.4% | certified | 2026-04-23 |
| `DefaultConfig` | function | B+ | 89.4% | certified | 2026-04-23 |
| `EnforcingConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ExpiryConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `IssueConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ModelAssignments` | class | B+ | 89.4% | certified | 2026-04-23 |
| `PolicyConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ProviderConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `RateLimitConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ScheduleConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ScopeConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `SignoffConfig` | class | B+ | 89.4% | certified | 2026-04-23 |
| `String` | method | B+ | 89.4% | certified | 2026-04-23 |
| `config_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `AllDimensions` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Dimension` | class | B+ | 89.4% | certified | 2026-04-23 |
| `DimensionScores` | class | B+ | 89.4% | certified | 2026-04-23 |
| `DimensionWeights` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Grade` | class | B+ | 89.4% | certified | 2026-04-23 |
| `GradeFromScore` | function | B+ | 88.3% | certified | 2026-04-23 |
| `String` | method | B+ | 89.4% | certified | 2026-04-23 |
| `WeightedAverage` | method | B+ | 88.3% | certified | 2026-04-23 |
| `dimension_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Evidence` | class | B+ | 89.4% | certified | 2026-04-23 |
| `EvidenceKind` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ParseSeverity` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Severity` | class | B+ | 89.4% | certified | 2026-04-23 |
| `String` | method | B+ | 89.4% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `evidence_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Duration` | method | B+ | 89.4% | certified | 2026-04-23 |
| `ExpiryFactors` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ExpiryWindow` | class | B+ | 89.4% | certified | 2026-04-23 |
| `IsExpired` | method | B+ | 89.4% | certified | 2026-04-23 |
| `RemainingAt` | method | B+ | 89.4% | certified | 2026-04-23 |
| `expiry_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Override` | class | B+ | 89.4% | certified | 2026-04-23 |
| `OverrideAction` | class | B+ | 89.4% | certified | 2026-04-23 |
| `String` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Validate` | method | B+ | 89.4% | certified | 2026-04-23 |
| `override_test.go` | file | B+ | 89.4% | certified | 2026-04-23 |
| `IsGlobal` | method | B+ | 89.4% | certified | 2026-04-23 |
| `PolicyPack` | class | B+ | 89.4% | certified | 2026-04-23 |
| `PolicyRule` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Violation` | class | B+ | 89.4% | certified | 2026-04-23 |
| `policy_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `CertificationRecord` | class | B+ | 89.4% | certified | 2026-04-23 |
| `IsPassing` | method | B+ | 89.4% | certified | 2026-04-23 |
| `ParseStatus` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Status` | class | B+ | 89.4% | certified | 2026-04-23 |
| `String` | method | B+ | 89.4% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `record_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Language` | method | B+ | 89.4% | certified | 2026-04-23 |
| `NewUnit` | function | B+ | 89.4% | certified | 2026-04-23 |
| `NewUnitID` | function | B+ | 89.4% | certified | 2026-04-23 |
| `ParseUnitID` | function | B+ | 89.4% | certified | 2026-04-23 |
| `ParseUnitType` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Path` | method | B+ | 89.4% | certified | 2026-04-23 |
| `String` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Symbol` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Unit` | class | B+ | 89.4% | certified | 2026-04-23 |
| `UnitID` | class | B+ | 89.4% | certified | 2026-04-23 |
| `UnitType` | class | B+ | 89.4% | certified | 2026-04-23 |
| `init` | function | B+ | 89.4% | certified | 2026-04-23 |
| `unit_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>AgentConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>AnalyzerConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>CertificationMode — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>Config — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>DefaultConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>EnforcingConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>ExpiryConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>IssueConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>ModelAssignments — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>PolicyConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>ProviderConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>RateLimitConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>ScheduleConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>ScopeConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>SignoffConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary>String — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detectable issues or complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>AllDimensions — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>Dimension — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>DimensionScores — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>DimensionWeights — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>Grade — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>GradeFromScore — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>String — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>WeightedAverage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>Evidence — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>EvidenceKind — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>ParseSeverity — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>Severity — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>String — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and reasonable code coverage
- 💡 Consider adding more comprehensive test cases for edge cases
- 💡 Evaluate if the current test coverage adequately addresses business logic
- 💡 Review code comments for clarity and completeness

</details>

<details>
<summary>Duration — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detected issues and full test coverage

</details>

<details>
<summary>ExpiryFactors — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detected issues and full test coverage

</details>

<details>
<summary>ExpiryWindow — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detected issues and full test coverage

</details>

<details>
<summary>IsExpired — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detected issues and full test coverage

</details>

<details>
<summary>RemainingAt — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detected issues and full test coverage

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>Override — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity suggests solid quality
- 💡 Consider adding godoc comments for better documentation
- 💡 Add example tests to demonstrate usage patterns
- 💡 Verify that the 66% coverage threshold meets project requirements

</details>

<details>
<summary>OverrideAction — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity suggests solid quality
- 💡 Consider adding godoc comments for better documentation
- 💡 Add example tests to demonstrate usage patterns
- 💡 Verify that the 66% coverage threshold meets project requirements

</details>

<details>
<summary>String — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity suggests solid quality
- 💡 Consider adding godoc comments for better documentation
- 💡 Add example tests to demonstrate usage patterns
- 💡 Verify that the 66% coverage threshold meets project requirements

</details>

<details>
<summary>Validate — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity suggests solid quality
- 💡 Consider adding godoc comments for better documentation
- 💡 Add example tests to demonstrate usage patterns
- 💡 Verify that the 66% coverage threshold meets project requirements

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>IsGlobal — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues found, passing all tests and linting checks

</details>

<details>
<summary>PolicyPack — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues found, passing all tests and linting checks

</details>

<details>
<summary>PolicyRule — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues found, passing all tests and linting checks

</details>

<details>
<summary>Violation — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues found, passing all tests and linting checks

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>CertificationRecord — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>IsPassing — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>ParseStatus — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>Status — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>String — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>Language — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>NewUnit — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>NewUnitID — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>ParseUnitID — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>ParseUnitType — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>Path — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>String — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>Symbol — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>Unit — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>UnitID — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>UnitType — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary>init — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests pass, and low complexity indicate high quality.

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `internal/engine/` (10 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `CertifyUnit` | function | B+ | 88.3% | certified | 2026-04-23 |
| `pipeline_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Score` | function | B+ | 88.3% | certified | 2026-04-23 |
| `StatusFromScore` | function | B+ | 89.4% | certified | 2026-04-23 |
| `extractSummaryFloat` | function | B+ | 88.3% | certified | 2026-04-23 |
| `extractSummaryInt` | function | B+ | 88.3% | certified | 2026-04-23 |
| `scoreFromGitHistory` | function | B+ | 89.4% | certified | 2026-04-23 |
| `scoreFromMetrics` | function | B+ | 87.2% | certified | 2026-04-23 |
| `severityPenalty` | function | B+ | 89.4% | certified | 2026-04-23 |
| `scorer_test.go` | file | B+ | 87.2% | certified | 2026-04-23 |

<details>
<summary>CertifyUnit — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with good test coverage and no linting issues

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>Score — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage, but has TODOs and moderate complexity
- 💡 Address the TODOs in the code
- 💡 Consider refactoring to reduce complexity from 9
- 💡 Add more comprehensive test coverage for edge cases

</details>

<details>
<summary>StatusFromScore — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage, but has TODOs and moderate complexity
- 💡 Address the TODOs in the code
- 💡 Consider refactoring to reduce complexity from 9
- 💡 Add more comprehensive test coverage for edge cases

</details>

<details>
<summary>extractSummaryFloat — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage, but has TODOs and moderate complexity
- 💡 Address the TODOs in the code
- 💡 Consider refactoring to reduce complexity from 9
- 💡 Add more comprehensive test coverage for edge cases

</details>

<details>
<summary>extractSummaryInt — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage, but has TODOs and moderate complexity
- 💡 Address the TODOs in the code
- 💡 Consider refactoring to reduce complexity from 9
- 💡 Add more comprehensive test coverage for edge cases

</details>

<details>
<summary>scoreFromGitHistory — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage, but has TODOs and moderate complexity
- 💡 Address the TODOs in the code
- 💡 Consider refactoring to reduce complexity from 9
- 💡 Add more comprehensive test coverage for edge cases

</details>

<details>
<summary>scoreFromMetrics — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 1 exceeds threshold 0
- 🤖 Code passes all tests and linting with good coverage, but has TODOs and moderate complexity
- 💡 Address the TODOs in the code
- 💡 Consider refactoring to reduce complexity from 9
- 💡 Add more comprehensive test coverage for edge cases

</details>

<details>
<summary>severityPenalty — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage, but has TODOs and moderate complexity
- 💡 Address the TODOs in the code
- 💡 Consider refactoring to reduce complexity from 9
- 💡 Add more comprehensive test coverage for edge cases

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 75.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 1 exceeds threshold 0
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and reasonable code coverage
- 💡 Consider removing the TODO comment if the task is complete
- 💡 Add more comprehensive test cases for edge scenarios
- 💡 Evaluate if the current test coverage (66%) is sufficient for production use

</details>

### `internal/evidence/` (39 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Collector` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ComputeGoComplexity` | function | B+ | 88.3% | certified | 2026-04-23 |
| `ComputeSymbolMetrics` | function | B+ | 87.2% | certified | 2026-04-23 |
| `funcName` | function | B+ | 89.4% | certified | 2026-04-23 |
| `complexity_test.go` | file | B+ | 87.2% | certified | 2026-04-23 |
| `CollectAll` | method | B+ | 89.4% | certified | 2026-04-23 |
| `HasGoMod` | method | B+ | 89.4% | certified | 2026-04-23 |
| `HasPackageJSON` | method | B+ | 89.4% | certified | 2026-04-23 |
| `NewToolExecutor` | function | B+ | 89.4% | certified | 2026-04-23 |
| `ToolExecutor` | class | B+ | 89.4% | certified | 2026-04-23 |
| `runGitStats` | method | B+ | 89.4% | certified | 2026-04-23 |
| `runGoTest` | method | B+ | 88.3% | certified | 2026-04-23 |
| `runGoVet` | method | B+ | 89.4% | certified | 2026-04-23 |
| `runGolangciLint` | method | B+ | 89.4% | certified | 2026-04-23 |
| `ChurnRate` | method | B+ | 89.4% | certified | 2026-04-23 |
| `GitStats` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ParseGitLog` | function | B+ | 89.4% | certified | 2026-04-23 |
| `ToEvidence` | method | B+ | 89.4% | certified | 2026-04-23 |
| `git_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `LintFinding` | class | B+ | 89.4% | certified | 2026-04-23 |
| `LintResult` | class | B+ | 89.4% | certified | 2026-04-23 |
| `TestResult` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ToEvidence` | method | B+ | 89.4% | certified | 2026-04-23 |
| `lint_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `CodeMetrics` | class | B+ | 88.3% | certified | 2026-04-23 |
| `ComputeMetrics` | function | B+ | 87.2% | certified | 2026-04-23 |
| `ToEvidence` | method | B+ | 88.3% | certified | 2026-04-23 |
| `containsTodo` | function | B+ | 88.3% | certified | 2026-04-23 |
| `metrics_test.go` | file | B+ | 87.2% | certified | 2026-04-23 |
| `ParseCoverProfile` | function | B+ | 88.3% | certified | 2026-04-23 |
| `ParseGitLogWithAge` | function | B+ | 89.4% | certified | 2026-04-23 |
| `ParseGoTestJSON` | function | B+ | 88.3% | certified | 2026-04-23 |
| `ParseGoVet` | function | B+ | 88.3% | certified | 2026-04-23 |
| `ParseGolangciLintJSON` | function | B+ | 88.3% | certified | 2026-04-23 |
| `goTestEvent` | class | B+ | 89.4% | certified | 2026-04-23 |
| `golangciLintIssue` | class | B+ | 89.4% | certified | 2026-04-23 |
| `golangciLintOutput` | class | B+ | 89.4% | certified | 2026-04-23 |
| `simpleAtoi` | function | B+ | 89.4% | certified | 2026-04-23 |
| `runner_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>Collector — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors, tests passing, and low complexity

</details>

<details>
<summary>ComputeGoComplexity — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>ComputeSymbolMetrics — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>funcName — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 75.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 5 exceeds threshold 0
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>CollectAll — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no issues found.

</details>

<details>
<summary>HasGoMod — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no issues found.

</details>

<details>
<summary>HasPackageJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no issues found.

</details>

<details>
<summary>NewToolExecutor — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no issues found.

</details>

<details>
<summary>ToolExecutor — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no issues found.

</details>

<details>
<summary>runGitStats — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no issues found.

</details>

<details>
<summary>runGoTest — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no issues found.

</details>

<details>
<summary>runGoVet — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>runGolangciLint — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with no issues found.

</details>

<details>
<summary>ChurnRate — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>GitStats — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>ParseGitLog — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>ToEvidence — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>LintFinding — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, tests passing, and low complexity

</details>

<details>
<summary>LintResult — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, tests passing, and low complexity

</details>

<details>
<summary>TestResult — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, tests passing, and low complexity

</details>

<details>
<summary>ToEvidence — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is clean, well-tested, and passes all static analysis checks with no issues found.

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>CodeMetrics — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 1 exceeds threshold 0
- 🤖 Minimal code size with no linting issues and full test coverage

</details>

<details>
<summary>ComputeMetrics — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 4 exceeds threshold 0
- 🤖 Code passes all tests and linting with good coverage, but contains TODOs that should be addressed
- 💡 Address the 4 TODOs in the codebase
- 💡 Consider breaking down function complexity if it exceeds 15
- 💡 Add unit tests for edge cases in ComputeMetrics

</details>

<details>
<summary>ToEvidence — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 1 exceeds threshold 0
- 🤖 Code passes all tests and linting with good coverage, but contains TODOs that should be addressed
- 💡 Address the 4 TODOs in the codebase
- 💡 Consider breaking down function complexity if it exceeds 15
- 💡 Add unit tests for edge cases in ComputeMetrics

</details>

<details>
<summary>containsTodo — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 2 exceeds threshold 0
- 🤖 Code passes all tests and linting with good coverage, but contains TODOs that should be addressed
- 💡 Address the 4 TODOs in the codebase
- 💡 Consider breaking down function complexity if it exceeds 15
- 💡 Add unit tests for edge cases in ComputeMetrics

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 75.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 7 exceeds threshold 0
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean structure

</details>

<details>
<summary>ParseCoverProfile — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>ParseGitLogWithAge — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>ParseGoTestJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>ParseGoVet — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>ParseGolangciLintJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>goTestEvent — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>golangciLintIssue — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>golangciLintOutput — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary>simpleAtoi — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

### `internal/expiry/` (2 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Calculate` | function | B+ | 88.3% | certified | 2026-04-23 |
| `calculator_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>Calculate — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity
- 💡 Consider adding more detailed comments for complex logic paths
- 💡 Evaluate if the current complexity (10) can be reduced through refactoring
- 💡 Verify that all edge cases are properly tested in the 66% coverage

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `internal/github/` (17 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `BuildIssueCloseCommand` | function | B+ | 89.4% | certified | 2026-04-23 |
| `BuildIssueCreateCommand` | function | B+ | 89.4% | certified | 2026-04-23 |
| `BuildIssueSearchCommand` | function | B+ | 89.4% | certified | 2026-04-23 |
| `BuildIssueUpdateCommand` | function | B+ | 89.4% | certified | 2026-04-23 |
| `FormatGroupedIssueBody` | function | B+ | 89.4% | certified | 2026-04-23 |
| `FormatIssueBody` | function | B+ | 89.4% | certified | 2026-04-23 |
| `FormatIssueTitle` | function | B+ | 89.4% | certified | 2026-04-23 |
| `issues_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `BuildPRCommentCommand` | function | B+ | 89.4% | certified | 2026-04-23 |
| `ComputeTrustDelta` | function | B+ | 88.3% | certified | 2026-04-23 |
| `FormatPRComment` | function | B+ | 88.3% | certified | 2026-04-23 |
| `TrustDelta` | class | B+ | 89.4% | certified | 2026-04-23 |
| `pr_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `GenerateNightlyWorkflow` | function | B+ | 89.4% | certified | 2026-04-23 |
| `GeneratePRWorkflow` | function | B+ | 89.4% | certified | 2026-04-23 |
| `GenerateWeeklyWorkflow` | function | B+ | 89.4% | certified | 2026-04-23 |
| `workflows_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>BuildIssueCloseCommand — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>BuildIssueCreateCommand — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>BuildIssueSearchCommand — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>BuildIssueUpdateCommand — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>FormatGroupedIssueBody — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>FormatIssueBody — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary>FormatIssueTitle — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and passes all static analysis checks with low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>BuildPRCommentCommand — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal function with low complexity and no linting issues

</details>

<details>
<summary>ComputeTrustDelta — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal function with low complexity and no linting issues

</details>

<details>
<summary>FormatPRComment — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal function with low complexity and no linting issues

</details>

<details>
<summary>TrustDelta — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues detected across multiple quality checks

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with full test coverage and no linting issues

</details>

<details>
<summary>GenerateNightlyWorkflow — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>GeneratePRWorkflow — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>GenerateWeeklyWorkflow — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `internal/override/` (9 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Apply` | function | B+ | 88.3% | certified | 2026-04-23 |
| `ApplyAll` | function | B+ | 89.4% | certified | 2026-04-23 |
| `applier_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `LoadDir` | function | B+ | 88.3% | certified | 2026-04-23 |
| `LoadFile` | function | B+ | 88.3% | certified | 2026-04-23 |
| `parseAction` | function | B+ | 89.4% | certified | 2026-04-23 |
| `rawOverride` | class | B+ | 89.4% | certified | 2026-04-23 |
| `rawOverrideFile` | class | B+ | 89.4% | certified | 2026-04-23 |
| `loader_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>Apply — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting issues, full test coverage, and low complexity

</details>

<details>
<summary>ApplyAll — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Clean, simple function with no linting issues, full test coverage, and low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>LoadDir — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity
- 💡 Consider adding more detailed error handling for file operations
- 💡 Add unit tests for edge cases like empty files or invalid formats
- 💡 Document the expected file format in comments

</details>

<details>
<summary>LoadFile — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity
- 💡 Consider adding more detailed error handling for file operations
- 💡 Add unit tests for edge cases like empty files or invalid formats
- 💡 Document the expected file format in comments

</details>

<details>
<summary>parseAction — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity
- 💡 Consider adding more detailed error handling for file operations
- 💡 Add unit tests for edge cases like empty files or invalid formats
- 💡 Document the expected file format in comments

</details>

<details>
<summary>rawOverride — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity
- 💡 Consider adding more detailed error handling for file operations
- 💡 Add unit tests for edge cases like empty files or invalid formats
- 💡 Document the expected file format in comments

</details>

<details>
<summary>rawOverrideFile — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and reasonable complexity
- 💡 Consider adding more detailed error handling for file operations
- 💡 Add unit tests for edge cases like empty files or invalid formats
- 💡 Document the expected file format in comments

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `internal/policy/` (14 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Evaluate` | function | B+ | 89.4% | certified | 2026-04-23 |
| `EvaluationResult` | class | B+ | 89.4% | certified | 2026-04-23 |
| `evaluateRule` | function | B+ | 89.4% | certified | 2026-04-23 |
| `extractComplexity` | function | B+ | 88.3% | certified | 2026-04-23 |
| `extractCoverage` | function | B+ | 89.4% | certified | 2026-04-23 |
| `extractMetric` | function | B+ | 87.2% | certified | 2026-04-23 |
| `extractTodoCount` | function | B+ | 87.2% | certified | 2026-04-23 |
| `evaluator_test.go` | file | B+ | 87.2% | certified | 2026-04-23 |
| `Match` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Matcher` | class | B+ | 89.4% | certified | 2026-04-23 |
| `NewMatcher` | function | B+ | 89.4% | certified | 2026-04-23 |
| `matchPath` | function | B+ | 88.3% | certified | 2026-04-23 |
| `matchesPack` | method | B+ | 88.3% | certified | 2026-04-23 |
| `matcher_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>Evaluate — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no static analysis issues, full test coverage, and minimal complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the current complexity (5) could be reduced through refactoring
- 💡 Verify that all edge cases are covered by existing tests

</details>

<details>
<summary>EvaluationResult — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no static analysis issues, full test coverage, and minimal complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the current complexity (5) could be reduced through refactoring
- 💡 Verify that all edge cases are covered by existing tests

</details>

<details>
<summary>evaluateRule — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no static analysis issues, full test coverage, and minimal complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the current complexity (5) could be reduced through refactoring
- 💡 Verify that all edge cases are covered by existing tests

</details>

<details>
<summary>extractComplexity — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no static analysis issues, full test coverage, and minimal complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the current complexity (5) could be reduced through refactoring
- 💡 Verify that all edge cases are covered by existing tests

</details>

<details>
<summary>extractCoverage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no static analysis issues, full test coverage, and minimal complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the current complexity (5) could be reduced through refactoring
- 💡 Verify that all edge cases are covered by existing tests

</details>

<details>
<summary>extractMetric — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 2 exceeds threshold 0
- 🤖 Code quality is acceptable with no static analysis issues, full test coverage, and minimal complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the current complexity (5) could be reduced through refactoring
- 💡 Verify that all edge cases are covered by existing tests

</details>

<details>
<summary>extractTodoCount — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 7 exceeds threshold 0
- 🤖 Code quality is acceptable with no static analysis issues, full test coverage, and minimal complexity
- 💡 Consider adding more descriptive comments for complex logic paths
- 💡 Evaluate if the current complexity (5) could be reduced through refactoring
- 💡 Verify that all edge cases are covered by existing tests

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 75.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 2 exceeds threshold 0
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and reasonable test coverage
- 💡 Consider reducing the number of TODOs to improve code completeness
- 💡 Evaluate if 66% test coverage is sufficient for this policy evaluation logic
- 💡 Check if the 81 lines of code could be refactored for better readability

</details>

<details>
<summary>Match — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage and minimal complexity

</details>

<details>
<summary>Matcher — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage and minimal complexity

</details>

<details>
<summary>NewMatcher — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage and minimal complexity

</details>

<details>
<summary>matchPath — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage and minimal complexity

</details>

<details>
<summary>matchesPack — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with good coverage and minimal complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

### `internal/queue/` (17 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `BatchNext` | method | B+ | 88.3% | certified | 2026-04-23 |
| `Complete` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Enqueue` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Fail` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Item` | class | B+ | 89.4% | certified | 2026-04-23 |
| `ItemStatus` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Len` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Load` | function | B+ | 89.4% | certified | 2026-04-23 |
| `New` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Next` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Queue` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Reset` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Save` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Skip` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Stats` | class | B+ | 88.3% | certified | 2026-04-23 |
| `persistedQueue` | class | B+ | 89.4% | certified | 2026-04-23 |
| `queue_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>BatchNext — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Complete — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Enqueue — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Fail — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Item — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>ItemStatus — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Len — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Load — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>New — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Next — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Queue — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Reset — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Save — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Skip — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>Stats — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary>persistedQueue — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting issues, all tests passing, and low complexity indicate high quality

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `internal/record/` (17 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `AppendHistory` | method | B+ | 89.4% | certified | 2026-04-23 |
| `ListAll` | method | B+ | 88.3% | certified | 2026-04-23 |
| `Load` | method | B+ | 89.4% | certified | 2026-04-23 |
| `LoadHistory` | method | B+ | 88.3% | certified | 2026-04-23 |
| `NewStore` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Save` | method | B+ | 89.4% | certified | 2026-04-23 |
| `Store` | class | B+ | 89.4% | certified | 2026-04-23 |
| `dimensionsToMap` | function | B+ | 89.4% | certified | 2026-04-23 |
| `fromJSON` | function | B+ | 89.4% | certified | 2026-04-23 |
| `historyEntry` | class | B+ | 89.4% | certified | 2026-04-23 |
| `historyPathFor` | method | B+ | 89.4% | certified | 2026-04-23 |
| `mapToDimensions` | function | B+ | 89.4% | certified | 2026-04-23 |
| `parseGrade` | function | B+ | 89.4% | certified | 2026-04-23 |
| `pathFor` | method | B+ | 89.4% | certified | 2026-04-23 |
| `recordJSON` | class | B+ | 89.4% | certified | 2026-04-23 |
| `toJSON` | function | B+ | 89.4% | certified | 2026-04-23 |
| `store_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>AppendHistory — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>ListAll — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>Load — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>LoadHistory — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>NewStore — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>Save — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>Store — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>dimensionsToMap — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>fromJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>historyEntry — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>historyPathFor — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>mapToDimensions — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>parseGrade — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>pathFor — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>recordJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary>toJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with low complexity and no linting or testing issues indicates high quality

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `internal/report/` (56 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `Badge` | class | B+ | 89.4% | certified | 2026-04-23 |
| `BadgeMarkdown` | function | B+ | 89.4% | certified | 2026-04-23 |
| `FormatBadgeJSON` | function | B+ | 89.4% | certified | 2026-04-23 |
| `GenerateBadge` | function | B+ | 89.4% | certified | 2026-04-23 |
| `badgeColor` | function | B+ | 88.3% | certified | 2026-04-23 |
| `badgeMessage` | function | B+ | 89.4% | certified | 2026-04-23 |
| `badge_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `Card` | class | B+ | 89.4% | certified | 2026-04-23 |
| `FormatCardMarkdown` | function | B+ | 87.2% | certified | 2026-04-23 |
| `FormatCardText` | function | B+ | 87.2% | certified | 2026-04-23 |
| `GenerateCard` | function | B+ | 88.3% | certified | 2026-04-23 |
| `IssueCard` | class | B+ | 89.4% | certified | 2026-04-23 |
| `LanguageCard` | class | B+ | 89.4% | certified | 2026-04-23 |
| `buildLanguageCards` | function | B+ | 89.4% | certified | 2026-04-23 |
| `buildTopIssues` | function | B+ | 88.3% | certified | 2026-04-23 |
| `gradeEmoji` | function | B+ | 89.4% | certified | 2026-04-23 |
| `card_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `AreaSummary` | class | B+ | 89.4% | certified | 2026-04-23 |
| `Detailed` | function | B+ | 89.4% | certified | 2026-04-23 |
| `DetailedReport` | class | B+ | 89.4% | certified | 2026-04-23 |
| `FormatDetailedText` | function | B+ | 87.2% | certified | 2026-04-23 |
| `LanguageBreakdown` | class | B+ | 89.4% | certified | 2026-04-23 |
| `UnitSummary` | class | B+ | 89.4% | certified | 2026-04-23 |
| `computeDimensionAverages` | function | B+ | 89.4% | certified | 2026-04-23 |
| `computeLanguageBreakdowns` | function | B+ | 89.4% | certified | 2026-04-23 |
| `explainStatus` | function | B+ | 89.4% | certified | 2026-04-23 |
| `findExpiringSoon` | function | B+ | 88.3% | certified | 2026-04-23 |
| `findFailing` | function | B+ | 89.4% | certified | 2026-04-23 |
| `findHighestRisk` | function | B+ | 89.4% | certified | 2026-04-23 |
| `findRecurrentlyFailing` | function | B+ | 89.4% | certified | 2026-04-23 |
| `unitSummaryFrom` | function | B+ | 89.4% | certified | 2026-04-23 |
| `detailed_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `FormatFullMarkdown` | function | B+ | 89.4% | certified | 2026-04-23 |
| `FullReport` | class | B+ | 89.4% | certified | 2026-04-23 |
| `GenerateFullReport` | function | B+ | 89.4% | certified | 2026-04-23 |
| `LanguageDetail` | class | B+ | 89.4% | certified | 2026-04-23 |
| `UnitReport` | class | B+ | 89.4% | certified | 2026-04-23 |
| `buildLanguageDetail` | function | B+ | 88.3% | certified | 2026-04-23 |
| `dirOf` | function | B+ | 89.4% | certified | 2026-04-23 |
| `shortFile` | function | B+ | 89.4% | certified | 2026-04-23 |
| `sortedKeys` | function | B+ | 89.4% | certified | 2026-04-23 |
| `unitReportFrom` | function | B+ | 89.4% | certified | 2026-04-23 |
| `writeAIInsights` | function | B+ | 87.2% | certified | 2026-04-23 |
| `writeAllUnits` | function | B+ | 88.3% | certified | 2026-04-23 |
| `writeDimensionAverages` | function | B+ | 89.4% | certified | 2026-04-23 |
| `writeGradeDistribution` | function | B+ | 89.4% | certified | 2026-04-23 |
| `writeHeader` | function | B+ | 89.4% | certified | 2026-04-23 |
| `writeLanguageDetail` | function | B+ | 88.3% | certified | 2026-04-23 |
| `writeSummary` | function | B+ | 89.4% | certified | 2026-04-23 |
| `writeUnitDetails` | function | B+ | 88.3% | certified | 2026-04-23 |
| `full_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |
| `FormatJSON` | function | B+ | 89.4% | certified | 2026-04-23 |
| `FormatText` | function | B+ | 89.4% | certified | 2026-04-23 |
| `Health` | function | B+ | 88.3% | certified | 2026-04-23 |
| `HealthReport` | class | B+ | 89.4% | certified | 2026-04-23 |
| `health_test.go` | file | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>Badge — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>BadgeMarkdown — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>FormatBadgeJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>GenerateBadge — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>badgeColor — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>badgeMessage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is excellent with full test coverage, no linting errors, and clean structure

</details>

<details>
<summary>Card — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues detected across multiple quality checks

</details>

<details>
<summary>FormatCardMarkdown — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues detected across multiple quality checks

</details>

<details>
<summary>FormatCardText — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues detected across multiple quality checks

</details>

<details>
<summary>GenerateCard — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues detected across multiple quality checks

</details>

<details>
<summary>IssueCard — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues detected across multiple quality checks

</details>

<details>
<summary>LanguageCard — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues detected across multiple quality checks

</details>

<details>
<summary>buildLanguageCards — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues detected across multiple quality checks

</details>

<details>
<summary>buildTopIssues — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues detected across multiple quality checks

</details>

<details>
<summary>gradeEmoji — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no issues detected across multiple quality checks

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, all tests passing, and clean commit history

</details>

<details>
<summary>AreaSummary — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>Detailed — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>DetailedReport — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>FormatDetailedText — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>LanguageBreakdown — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>UnitSummary — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>computeDimensionAverages — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>computeLanguageBreakdowns — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>explainStatus — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>findExpiringSoon — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>findFailing — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>findHighestRisk — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>findRecurrentlyFailing — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary>unitSummaryFrom — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>FormatFullMarkdown — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>FullReport — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>GenerateFullReport — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>LanguageDetail — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>UnitReport — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>buildLanguageDetail — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>dirOf — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>shortFile — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>sortedKeys — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>unitReportFrom — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>writeAIInsights — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>writeAllUnits — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>writeDimensionAverages — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>writeGradeDistribution — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>writeHeader — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>writeLanguageDetail — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>writeSummary — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary>writeUnitDetails — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code is minimal, well-tested, and free of linting issues with low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with full test coverage and no linting issues

</details>

<details>
<summary>FormatJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>FormatText — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>Health — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary>HealthReport — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and low complexity

</details>

<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `testdata/repos/ts-simple/src/` (6 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `DialogueParser` | class | B+ | 89.4% | certified | 2026-04-23 |
| `MAX_TOKENS` | function | B+ | 89.4% | certified | 2026-04-23 |
| `parseNode` | function | B+ | 89.4% | certified | 2026-04-23 |
| `tokenizeDialogue` | function | B+ | 89.4% | certified | 2026-04-23 |
| `formatDate` | function | B+ | 89.4% | certified | 2026-04-23 |
| `log` | function | B+ | 89.4% | certified | 2026-04-23 |

<details>
<summary>DialogueParser — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity
- 💡 Consider adding unit tests for edge cases
- 💡 Add documentation comments for public functions
- 💡 Verify token length validation logic

</details>

<details>
<summary>MAX_TOKENS — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity
- 💡 Consider adding unit tests for edge cases
- 💡 Add documentation comments for public functions
- 💡 Verify token length validation logic

</details>

<details>
<summary>parseNode — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity
- 💡 Consider adding unit tests for edge cases
- 💡 Add documentation comments for public functions
- 💡 Verify token length validation logic

</details>

<details>
<summary>tokenizeDialogue — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity
- 💡 Consider adding unit tests for edge cases
- 💡 Add documentation comments for public functions
- 💡 Verify token length validation logic

</details>

<details>
<summary>formatDate — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detected issues, passing all tests and linting checks

</details>

<details>
<summary>log — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no detected issues, passing all tests and linting checks

</details>

### `vscode-certify/src/` (31 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `RunResult` | class | B+ | 88.3% | certified | 2026-04-23 |
| `findCertifyBinary` | function | B+ | 88.3% | certified | 2026-04-23 |
| `listModels` | function | B+ | 88.3% | certified | 2026-04-23 |
| `promptInstall` | function | B+ | 88.3% | certified | 2026-04-23 |
| `runCertify` | function | B+ | 88.3% | certified | 2026-04-23 |
| `runCertifyJSON` | function | B+ | 88.3% | certified | 2026-04-23 |
| `runInTerminal` | function | B+ | 88.3% | certified | 2026-04-23 |
| `BRAND_COLORS` | function | B+ | 88.3% | certified | 2026-04-23 |
| `DIMENSION_NAMES` | function | B+ | 88.3% | certified | 2026-04-23 |
| `GRADE_COLORS` | function | B+ | 88.3% | certified | 2026-04-23 |
| `GRADE_EMOJI` | function | B+ | 88.3% | certified | 2026-04-23 |
| `PROVIDER_PRESETS` | function | B+ | 88.3% | certified | 2026-04-23 |
| `CertifyDataLoader` | class | B+ | 88.3% | certified | 2026-04-23 |
| `activate` | function | B+ | 88.3% | certified | 2026-04-23 |
| `deactivate` | function | B+ | 88.3% | certified | 2026-04-23 |
| `createStatusBarItem` | function | B+ | 89.4% | certified | 2026-04-23 |
| `AgentConfig` | class | B+ | 88.3% | certified | 2026-04-23 |
| `BadgeJSON` | class | B+ | 88.3% | certified | 2026-04-23 |
| `CertifyCard` | class | B+ | 88.3% | certified | 2026-04-23 |
| `CertifyConfig` | class | B+ | 88.3% | certified | 2026-04-23 |
| `FullReport` | class | B+ | 88.3% | certified | 2026-04-23 |
| `IndexEntry` | class | B+ | 88.3% | certified | 2026-04-23 |
| `IssueCard` | class | B+ | 88.3% | certified | 2026-04-23 |
| `LanguageCard` | class | B+ | 88.3% | certified | 2026-04-23 |
| `LanguageDetail` | class | B+ | 88.3% | certified | 2026-04-23 |
| `ModelAssignments` | class | B+ | 88.3% | certified | 2026-04-23 |
| `ModelInfo` | class | B+ | 88.3% | certified | 2026-04-23 |
| `ProviderConfig` | class | B+ | 88.3% | certified | 2026-04-23 |
| `ProviderPreset` | class | B+ | 88.3% | certified | 2026-04-23 |
| `RecordJSON` | class | B+ | 88.3% | certified | 2026-04-23 |
| `UnitReport` | class | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>RunResult — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>findCertifyBinary — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>listModels — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>promptInstall — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with full test coverage and no linting issues

</details>

<details>
<summary>runCertify — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with full test coverage and no linting issues

</details>

<details>
<summary>runCertifyJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>runInTerminal — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>BRAND_COLORS — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider adding JSDoc comments for better documentation
- 💡 Evaluate if the emoji constants could be exported from a dedicated constants file
- 💡 Verify that all grade emojis are properly localized for internationalization

</details>

<details>
<summary>DIMENSION_NAMES — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider adding JSDoc comments for better documentation
- 💡 Evaluate if the emoji constants could be exported from a dedicated constants file
- 💡 Verify that all grade emojis are properly localized for internationalization

</details>

<details>
<summary>GRADE_COLORS — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider adding JSDoc comments for better documentation
- 💡 Evaluate if the emoji constants could be exported from a dedicated constants file
- 💡 Verify that all grade emojis are properly localized for internationalization

</details>

<details>
<summary>GRADE_EMOJI — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider adding JSDoc comments for better documentation
- 💡 Evaluate if the emoji constants could be exported from a dedicated constants file
- 💡 Verify that all grade emojis are properly localized for internationalization

</details>

<details>
<summary>PROVIDER_PRESETS — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider adding JSDoc comments for better documentation
- 💡 Evaluate if the emoji constants could be exported from a dedicated constants file
- 💡 Verify that all grade emojis are properly localized for internationalization

</details>

<details>
<summary>CertifyDataLoader — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>activate — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all static analysis checks with no errors or warnings, and has comprehensive test coverage
- 💡 Consider adding more detailed logging for debugging purposes
- 💡 Evaluate if the current test coverage (66%) can be improved for edge cases
- 💡 Check if any of the 72 lines could benefit from refactoring to reduce complexity

</details>

<details>
<summary>deactivate — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all static analysis checks with no errors or warnings, and has comprehensive test coverage
- 💡 Consider adding more detailed logging for debugging purposes
- 💡 Evaluate if the current test coverage (66%) can be improved for edge cases
- 💡 Check if any of the 72 lines could benefit from refactoring to reduce complexity

</details>

<details>
<summary>createStatusBarItem — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and minimal complexity
- 💡 Consider adding unit tests for edge cases
- 💡 Add JSDoc comments for better documentation
- 💡 Implement error handling for status bar item creation

</details>

<details>
<summary>AgentConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>BadgeJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>CertifyCard — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>CertifyConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>FullReport — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>IndexEntry — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>IssueCard — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>LanguageCard — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>LanguageDetail — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>ModelAssignments — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>ModelInfo — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>ProviderConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>ProviderPreset — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>RecordJSON — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

<details>
<summary>UnitReport — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality appears acceptable with no linting errors, full test coverage, and clean commit history

</details>

### `vscode-certify/src/codeLens/` (2 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `CertifyCodeLensProvider` | class | B+ | 88.3% | certified | 2026-04-23 |
| `showDimensionScores` | function | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>CertifyCodeLensProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is good with no linting errors, full test coverage, and clean structure

</details>

<details>
<summary>showDimensionScores — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is good with no linting errors, full test coverage, and clean structure

</details>

### `vscode-certify/src/config/` (5 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `ConfigPanel` | class | B+ | 88.3% | certified | 2026-04-23 |
| `ConnectionTestResult` | class | B+ | 88.3% | certified | 2026-04-23 |
| `readConfig` | function | B+ | 88.3% | certified | 2026-04-23 |
| `testConnection` | function | B+ | 88.3% | certified | 2026-04-23 |
| `writeConfig` | function | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>ConfigPanel — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>ConnectionTestResult — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider adding unit tests for edge cases in the ConnectionTestResult class
- 💡 Add documentation comments for public methods and fields
- 💡 Implement validation logic for connection test results

</details>

<details>
<summary>readConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider adding unit tests for edge cases in the ConnectionTestResult class
- 💡 Add documentation comments for public methods and fields
- 💡 Implement validation logic for connection test results

</details>

<details>
<summary>testConnection — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider adding unit tests for edge cases in the ConnectionTestResult class
- 💡 Add documentation comments for public methods and fields
- 💡 Implement validation logic for connection test results

</details>

<details>
<summary>writeConfig — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider adding unit tests for edge cases in the ConnectionTestResult class
- 💡 Add documentation comments for public methods and fields
- 💡 Implement validation logic for connection test results

</details>

### `vscode-certify/src/dashboard/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `DashboardPanel` | class | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>DashboardPanel — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all automated checks with full test coverage and no linting issues

</details>

### `vscode-certify/src/diagnostics/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `CertifyDiagnostics` | class | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>CertifyDiagnostics — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code quality is acceptable with no linting errors, full test coverage, and clean commit history
- 💡 Consider adding unit tests for edge cases in diagnostics
- 💡 Add documentation comments for public methods
- 💡 Implement logging for diagnostic events

</details>

### `vscode-certify/src/treeView/` (2 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `CertificationTreeProvider` | class | B+ | 88.3% | certified | 2026-04-23 |
| `CertifyTreeItem` | class | B+ | 88.3% | certified | 2026-04-23 |

<details>
<summary>CertificationTreeProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

<details>
<summary>CertifyTreeItem — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Code passes all tests and linting with no issues found

</details>

### `website/src/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| `collections` | function | B+ | 89.4% | certified | 2026-04-23 |

<details>
<summary>collections — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 Minimal code size with no linting errors and full test coverage

</details>

---

*542 units certified. Generated by [Certify](https://github.com/iksnae/code-certification).*
