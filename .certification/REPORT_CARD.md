# 🟢 Certify — Report Card

**Repository:** `iksnae/code-certification`  
**Commit:** `84930b9`  
**Generated:** 2026-03-09T20:16:34  

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
- 💡 Address the TODOs in the code ×7
- 💡 Verify that all edge cases are covered by existing tests ×7
- 💡 Add more comprehensive test coverage for edge cases ×7
- 💡 Evaluate if the current complexity (5) could be reduced through refactoring ×7
- 💡 Consider refactoring to reduce complexity from 9 ×7
- 💡 Check if error handling could be made more explicit ×5
- 💡 Evaluate if the function name clearly reflects its purpose ×5
- 💡 Consider adding more detailed error handling for file operations ×5
- 💡 Consider adding unit tests for edge cases ×5
- 💡 Add unit tests for edge cases like empty files or invalid formats ×5
- 💡 Evaluate if the emoji constants could be exported from a dedicated constants file ×5
- 💡 Verify that all grade emojis are properly localized for internationalization ×5
- 💡 Consider adding JSDoc comments for better documentation ×5
- 💡 Document the expected file format in comments ×5
- 💡 Add example tests to demonstrate usage patterns ×4
- 💡 Verify that the 66% coverage threshold meets project requirements ×4
- 💡 Consider adding godoc comments for better documentation ×4
- 💡 Add documentation comments for public methods and fields ×4
- 💡 Implement validation logic for connection test results ×4

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
| [`certifyContext`](#cmd-certify-certify-cmd-go-certifycontext) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`certifyUnit`](#cmd-certify-certify-cmd-go-certifyunit) | method | B+ | 87.2% | certified | 2026-04-23 |
| [`collectRepoEvidence`](#cmd-certify-certify-cmd-go-collectrepoevidence) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`defaultConfigObj`](#cmd-certify-certify-cmd-go-defaultconfigobj) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`filterUnits`](#cmd-certify-certify-cmd-go-filterunits) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](#cmd-certify-certify-cmd-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`isLocalURL`](#cmd-certify-certify-cmd-go-islocalurl) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`loadCertifyContext`](#cmd-certify-certify-cmd-go-loadcertifycontext) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`loadQueue`](#cmd-certify-certify-cmd-go-loadqueue) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`printQueueStatus`](#cmd-certify-certify-cmd-go-printqueuestatus) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`printSummary`](#cmd-certify-certify-cmd-go-printsummary) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`processQueue`](#cmd-certify-certify-cmd-go-processqueue) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`runCertify`](#cmd-certify-certify-cmd-go-runcertify) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`saveReportArtifacts`](#cmd-certify-certify-cmd-go-savereportartifacts) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`setupAgent`](#cmd-certify-certify-cmd-go-setupagent) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`setupConservativeAgent`](#cmd-certify-certify-cmd-go-setupconservativeagent) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`setupExplicitAgent`](#cmd-certify-certify-cmd-go-setupexplicitagent) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`cli_test.go`](#cmd-certify-cli-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`init`](#cmd-certify-expire-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`generateConfig`](#cmd-certify-init-cmd-go-generateconfig) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](#cmd-certify-init-cmd-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`languagePolicy`](#cmd-certify-init-cmd-go-languagepolicy) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`main`](#cmd-certify-main-go-main) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](#cmd-certify-models-cmd-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`listFromProvider`](#cmd-certify-models-cmd-go-listfromprovider) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`runModels`](#cmd-certify-models-cmd-go-runmodels) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`detectCommit`](#cmd-certify-report-cmd-go-detectcommit) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`detectRepoName`](#cmd-certify-report-cmd-go-detectreponame) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](#cmd-certify-report-cmd-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`saveBadge`](#cmd-certify-report-cmd-go-savebadge) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`saveReportCard`](#cmd-certify-report-cmd-go-savereportcard) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](#cmd-certify-review-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](#cmd-certify-root-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](#cmd-certify-scan-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`tryScanSuggestions`](#cmd-certify-scan-go-tryscansuggestions) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`version.go`](#cmd-certify-version-go) | file | B+ | 89.4% | certified | 2026-04-23 |

<a id="cmd-certify-certify-cmd-go-certifycontext"></a>
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

<a id="cmd-certify-certify-cmd-go-certifyunit"></a>
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

<a id="cmd-certify-certify-cmd-go-collectrepoevidence"></a>
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

<a id="cmd-certify-certify-cmd-go-defaultconfigobj"></a>
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

<a id="cmd-certify-certify-cmd-go-filterunits"></a>
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

<a id="cmd-certify-certify-cmd-go-init"></a>
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

<a id="cmd-certify-certify-cmd-go-islocalurl"></a>
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

<a id="cmd-certify-certify-cmd-go-loadcertifycontext"></a>
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

<a id="cmd-certify-certify-cmd-go-loadqueue"></a>
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

<a id="cmd-certify-certify-cmd-go-printqueuestatus"></a>
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

<a id="cmd-certify-certify-cmd-go-printsummary"></a>
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

<a id="cmd-certify-certify-cmd-go-processqueue"></a>
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

<a id="cmd-certify-certify-cmd-go-runcertify"></a>
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

<a id="cmd-certify-certify-cmd-go-savereportartifacts"></a>
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

<a id="cmd-certify-certify-cmd-go-setupagent"></a>
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

<a id="cmd-certify-certify-cmd-go-setupconservativeagent"></a>
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

<a id="cmd-certify-certify-cmd-go-setupexplicitagent"></a>
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

<a id="cmd-certify-cli-test-go"></a>
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

<a id="cmd-certify-expire-go-init"></a>
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

<a id="cmd-certify-init-cmd-go-generateconfig"></a>
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

<a id="cmd-certify-init-cmd-go-init"></a>
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

<a id="cmd-certify-init-cmd-go-languagepolicy"></a>
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

<a id="cmd-certify-main-go-main"></a>
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

<a id="cmd-certify-models-cmd-go-init"></a>
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

<a id="cmd-certify-models-cmd-go-listfromprovider"></a>
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

<a id="cmd-certify-models-cmd-go-runmodels"></a>
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

<a id="cmd-certify-report-cmd-go-detectcommit"></a>
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

<a id="cmd-certify-report-cmd-go-detectreponame"></a>
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

<a id="cmd-certify-report-cmd-go-init"></a>
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

<a id="cmd-certify-report-cmd-go-savebadge"></a>
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

<a id="cmd-certify-report-cmd-go-savereportcard"></a>
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

<a id="cmd-certify-review-go-init"></a>
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

<a id="cmd-certify-root-go-init"></a>
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

<a id="cmd-certify-scan-go-init"></a>
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

<a id="cmd-certify-scan-go-tryscansuggestions"></a>
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

<a id="cmd-certify-version-go"></a>
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
| [`agent-chain.ts`](#extensions-agent-chain-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`agent-team.ts`](#extensions-agent-team-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`cross-agent.ts`](#extensions-cross-agent-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`damage-control.ts`](#extensions-damage-control-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`minimal.ts`](#extensions-minimal-ts) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`pi-pi.ts`](#extensions-pi-pi-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`pure-focus.ts`](#extensions-pure-focus-ts) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`purpose-gate.ts`](#extensions-purpose-gate-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`session-replay.ts`](#extensions-session-replay-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`subagent-widget.ts`](#extensions-subagent-widget-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`system-select.ts`](#extensions-system-select-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`theme-cycler.ts`](#extensions-theme-cycler-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`THEME_MAP`](#extensions-thememap-ts-theme-map) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`applyExtensionDefaults`](#extensions-thememap-ts-applyextensiondefaults) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`applyExtensionTheme`](#extensions-thememap-ts-applyextensiontheme) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`tilldone.ts`](#extensions-tilldone-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`tool-counter-widget.ts`](#extensions-tool-counter-widget-ts) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`tool-counter.ts`](#extensions-tool-counter-ts) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="extensions-agent-chain-ts"></a>
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

<a id="extensions-agent-team-ts"></a>
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

<a id="extensions-cross-agent-ts"></a>
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

<a id="extensions-damage-control-ts"></a>
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

<a id="extensions-minimal-ts"></a>
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

<a id="extensions-pi-pi-ts"></a>
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

<a id="extensions-pure-focus-ts"></a>
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

<a id="extensions-purpose-gate-ts"></a>
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

<a id="extensions-session-replay-ts"></a>
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

<a id="extensions-subagent-widget-ts"></a>
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

<a id="extensions-system-select-ts"></a>
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

<a id="extensions-theme-cycler-ts"></a>
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

<a id="extensions-thememap-ts-theme-map"></a>
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

<a id="extensions-thememap-ts-applyextensiondefaults"></a>
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

<a id="extensions-thememap-ts-applyextensiontheme"></a>
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

<a id="extensions-tilldone-ts"></a>
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

<a id="extensions-tool-counter-widget-ts"></a>
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

<a id="extensions-tool-counter-ts"></a>
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
| [`integration_test.go`](#internal-integration-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-integration-test-go"></a>
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
| [`attribution_test.go`](#internal-agent-attribution-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`DetectAPIKey`](#internal-agent-autodetect-go-detectapikey) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatProviderSummary`](#internal-agent-autodetect-go-formatprovidersummary) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`HasAnyProvider`](#internal-agent-autodetect-go-hasanyprovider) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewConservativeCoordinator`](#internal-agent-autodetect-go-newconservativecoordinator) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`init`](#internal-agent-autodetect-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`autodetect_test.go`](#internal-agent-autodetect-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Chat`](#internal-agent-circuit-go-chat) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`CircuitBreaker`](#internal-agent-circuit-go-circuitbreaker) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`IsOpen`](#internal-agent-circuit-go-isopen) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Name`](#internal-agent-circuit-go-name) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewCircuitBreaker`](#internal-agent-circuit-go-newcircuitbreaker) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`AdaptiveMessages`](#internal-agent-fallback-go-adaptivemessages) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Chat`](#internal-agent-fallback-go-chat) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`FallbackProvider`](#internal-agent-fallback-go-fallbackprovider) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ModelChain`](#internal-agent-fallback-go-modelchain) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Name`](#internal-agent-fallback-go-name) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewFallbackProvider`](#internal-agent-fallback-go-newfallbackprovider) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewModelChain`](#internal-agent-fallback-go-newmodelchain) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`modelPinnedProvider`](#internal-agent-fallback-go-modelpinnedprovider) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`fallback_test.go`](#internal-agent-fallback-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`ListModels`](#internal-agent-models-go-listmodels) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ModelInfo`](#internal-agent-models-go-modelinfo) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`listOllamaModels`](#internal-agent-models-go-listollamamodels) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`listOpenAIModels`](#internal-agent-models-go-listopenaimodels) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ollamaModel`](#internal-agent-models-go-ollamamodel) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ollamaTagsResponse`](#internal-agent-models-go-ollamatagsresponse) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`openAIModel`](#internal-agent-models-go-openaimodel) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`openAIModelsResponse`](#internal-agent-models-go-openaimodelsresponse) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`models_test.go`](#internal-agent-models-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`APIError`](#internal-agent-openrouter-go-apierror) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Chat`](#internal-agent-openrouter-go-chat) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`Error`](#internal-agent-openrouter-go-error) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Name`](#internal-agent-openrouter-go-name) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewLocalProvider`](#internal-agent-openrouter-go-newlocalprovider) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewOpenRouterProvider`](#internal-agent-openrouter-go-newopenrouterprovider) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`OpenRouterProvider`](#internal-agent-openrouter-go-openrouterprovider) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`doRequest`](#internal-agent-openrouter-go-dorequest) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`isAPIError`](#internal-agent-openrouter-go-isapierror) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`isAuthError`](#internal-agent-openrouter-go-isautherror) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`isBudgetError`](#internal-agent-openrouter-go-isbudgeterror) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`isRetryable`](#internal-agent-openrouter-go-isretryable) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`openrouter_test.go`](#internal-agent-openrouter-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Coordinator`](#internal-agent-pipeline-go-coordinator) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`CoordinatorConfig`](#internal-agent-pipeline-go-coordinatorconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`NewCoordinator`](#internal-agent-pipeline-go-newcoordinator) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewPipeline`](#internal-agent-pipeline-go-newpipeline) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Pipeline`](#internal-agent-pipeline-go-pipeline) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`PipelineConfig`](#internal-agent-pipeline-go-pipelineconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ReviewUnit`](#internal-agent-pipeline-go-reviewunit) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`Run`](#internal-agent-pipeline-go-run) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`Stats`](#internal-agent-pipeline-go-stats) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Strategy`](#internal-agent-pipeline-go-strategy) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`toResult`](#internal-agent-pipeline-go-toresult) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Get`](#internal-agent-prompts-go-get) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`LoadPrompt`](#internal-agent-prompts-go-loadprompt) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewPromptRegistry`](#internal-agent-prompts-go-newpromptregistry) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`PromptRegistry`](#internal-agent-prompts-go-promptregistry) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`PromptTemplate`](#internal-agent-prompts-go-prompttemplate) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Render`](#internal-agent-prompts-go-render) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Version`](#internal-agent-prompts-go-version) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`prompts_test.go`](#internal-agent-prompts-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Provider`](#internal-agent-provider-go-provider) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`provider_multi_test.go`](#internal-agent-provider-multi-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`DetectProviders`](#internal-agent-providers-go-detectproviders) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`DetectedProvider`](#internal-agent-providers-go-detectedprovider) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ProviderNames`](#internal-agent-providers-go-providernames) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](#internal-agent-providers-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`normalizeLocalURL`](#internal-agent-providers-go-normalizelocalurl) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`probeLocal`](#internal-agent-providers-go-probelocal) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Allow`](#internal-agent-ratelimit-go-allow) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewRateLimiter`](#internal-agent-ratelimit-go-newratelimiter) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`RateLimiter`](#internal-agent-ratelimit-go-ratelimiter) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Wait`](#internal-agent-ratelimit-go-wait) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`refill`](#internal-agent-ratelimit-go-refill) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ratelimit_test.go`](#internal-agent-ratelimit-test-go) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`NewReviewer`](#internal-agent-reviewer-go-newreviewer) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Review`](#internal-agent-reviewer-go-review) | method | B+ | 87.2% | certified | 2026-04-23 |
| [`ReviewInput`](#internal-agent-reviewer-go-reviewinput) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ReviewResult`](#internal-agent-reviewer-go-reviewresult) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Reviewer`](#internal-agent-reviewer-go-reviewer) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ToEvidence`](#internal-agent-reviewer-go-toevidence) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ToPrescreenEvidence`](#internal-agent-reviewer-go-toprescreenevidence) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`joinModels`](#internal-agent-reviewer-go-joinmodels) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`reviewer_test.go`](#internal-agent-reviewer-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`ModelFor`](#internal-agent-router-go-modelfor) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewRouter`](#internal-agent-router-go-newrouter) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Router`](#internal-agent-router-go-router) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`directAssignment`](#internal-agent-router-go-directassignment) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`router_test.go`](#internal-agent-router-test-go) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`DecisionResponse`](#internal-agent-schemas-go-decisionresponse) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`PrescreenResponse`](#internal-agent-schemas-go-prescreenresponse) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`RemediationResponse`](#internal-agent-schemas-go-remediationresponse) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`RemediationStep`](#internal-agent-schemas-go-remediationstep) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ScoringResponse`](#internal-agent-schemas-go-scoringresponse) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`schemas_test.go`](#internal-agent-schemas-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Execute`](#internal-agent-stage-go-execute) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Name`](#internal-agent-stage-go-name) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewPrescreenStage`](#internal-agent-stage-go-newprescreenstage) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewReviewStage`](#internal-agent-stage-go-newreviewstage) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewScoringStage`](#internal-agent-stage-go-newscoringstage) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Stage`](#internal-agent-stage-go-stage) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`StageInput`](#internal-agent-stage-go-stageinput) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`StageResult`](#internal-agent-stage-go-stageresult) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`defaultScores`](#internal-agent-stage-go-defaultscores) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`extractJSON`](#internal-agent-stage-go-extractjson) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`looseParseNeedsReview`](#internal-agent-stage-go-looseparseneedsreview) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`prescreenStage`](#internal-agent-stage-go-prescreenstage) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`reviewStage`](#internal-agent-stage-go-reviewstage) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`scoringStage`](#internal-agent-stage-go-scoringstage) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`stage_test.go`](#internal-agent-stage-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`RepoSummary`](#internal-agent-suggest-go-reposummary) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ScanSuggestion`](#internal-agent-suggest-go-scansuggestion) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`SuggestForRepo`](#internal-agent-suggest-go-suggestforrepo) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`buildSuggestPrompt`](#internal-agent-suggest-go-buildsuggestprompt) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`suggest_test.go`](#internal-agent-suggest-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`ChatRequest`](#internal-agent-types-go-chatrequest) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ChatResponse`](#internal-agent-types-go-chatresponse) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Choice`](#internal-agent-types-go-choice) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Content`](#internal-agent-types-go-content) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Message`](#internal-agent-types-go-message) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ModelConfig`](#internal-agent-types-go-modelconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ResponseFormat`](#internal-agent-types-go-responseformat) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`String`](#internal-agent-types-go-string) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`TaskType`](#internal-agent-types-go-tasktype) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Usage`](#internal-agent-types-go-usage) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`types_test.go`](#internal-agent-types-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-agent-attribution-test-go"></a>
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

<a id="internal-agent-autodetect-go-detectapikey"></a>
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

<a id="internal-agent-autodetect-go-formatprovidersummary"></a>
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

<a id="internal-agent-autodetect-go-hasanyprovider"></a>
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

<a id="internal-agent-autodetect-go-newconservativecoordinator"></a>
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

<a id="internal-agent-autodetect-go-init"></a>
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

<a id="internal-agent-autodetect-test-go"></a>
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

<a id="internal-agent-circuit-go-chat"></a>
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

<a id="internal-agent-circuit-go-circuitbreaker"></a>
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

<a id="internal-agent-circuit-go-isopen"></a>
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

<a id="internal-agent-circuit-go-name"></a>
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

<a id="internal-agent-circuit-go-newcircuitbreaker"></a>
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

<a id="internal-agent-fallback-go-adaptivemessages"></a>
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

<a id="internal-agent-fallback-go-chat"></a>
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

<a id="internal-agent-fallback-go-fallbackprovider"></a>
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

<a id="internal-agent-fallback-go-modelchain"></a>
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

<a id="internal-agent-fallback-go-name"></a>
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

<a id="internal-agent-fallback-go-newfallbackprovider"></a>
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

<a id="internal-agent-fallback-go-newmodelchain"></a>
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

<a id="internal-agent-fallback-go-modelpinnedprovider"></a>
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

<a id="internal-agent-fallback-test-go"></a>
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

<a id="internal-agent-models-go-listmodels"></a>
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

<a id="internal-agent-models-go-modelinfo"></a>
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

<a id="internal-agent-models-go-listollamamodels"></a>
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

<a id="internal-agent-models-go-listopenaimodels"></a>
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

<a id="internal-agent-models-go-ollamamodel"></a>
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

<a id="internal-agent-models-go-ollamatagsresponse"></a>
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

<a id="internal-agent-models-go-openaimodel"></a>
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

<a id="internal-agent-models-go-openaimodelsresponse"></a>
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

<a id="internal-agent-models-test-go"></a>
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

<a id="internal-agent-openrouter-go-apierror"></a>
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

<a id="internal-agent-openrouter-go-chat"></a>
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

<a id="internal-agent-openrouter-go-error"></a>
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

<a id="internal-agent-openrouter-go-name"></a>
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

<a id="internal-agent-openrouter-go-newlocalprovider"></a>
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

<a id="internal-agent-openrouter-go-newopenrouterprovider"></a>
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

<a id="internal-agent-openrouter-go-openrouterprovider"></a>
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

<a id="internal-agent-openrouter-go-dorequest"></a>
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

<a id="internal-agent-openrouter-go-isapierror"></a>
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

<a id="internal-agent-openrouter-go-isautherror"></a>
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

<a id="internal-agent-openrouter-go-isbudgeterror"></a>
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

<a id="internal-agent-openrouter-go-isretryable"></a>
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

<a id="internal-agent-openrouter-test-go"></a>
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

<a id="internal-agent-pipeline-go-coordinator"></a>
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

<a id="internal-agent-pipeline-go-coordinatorconfig"></a>
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

<a id="internal-agent-pipeline-go-newcoordinator"></a>
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

<a id="internal-agent-pipeline-go-newpipeline"></a>
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

<a id="internal-agent-pipeline-go-pipeline"></a>
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

<a id="internal-agent-pipeline-go-pipelineconfig"></a>
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

<a id="internal-agent-pipeline-go-reviewunit"></a>
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

<a id="internal-agent-pipeline-go-run"></a>
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

<a id="internal-agent-pipeline-go-stats"></a>
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

<a id="internal-agent-pipeline-go-strategy"></a>
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

<a id="internal-agent-pipeline-go-toresult"></a>
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

<a id="internal-agent-prompts-go-get"></a>
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

<a id="internal-agent-prompts-go-loadprompt"></a>
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

<a id="internal-agent-prompts-go-newpromptregistry"></a>
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

<a id="internal-agent-prompts-go-promptregistry"></a>
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

<a id="internal-agent-prompts-go-prompttemplate"></a>
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

<a id="internal-agent-prompts-go-render"></a>
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

<a id="internal-agent-prompts-go-version"></a>
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

<a id="internal-agent-prompts-test-go"></a>
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

<a id="internal-agent-provider-go-provider"></a>
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

<a id="internal-agent-provider-multi-test-go"></a>
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

<a id="internal-agent-providers-go-detectproviders"></a>
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

<a id="internal-agent-providers-go-detectedprovider"></a>
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

<a id="internal-agent-providers-go-providernames"></a>
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

<a id="internal-agent-providers-go-init"></a>
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

<a id="internal-agent-providers-go-normalizelocalurl"></a>
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

<a id="internal-agent-providers-go-probelocal"></a>
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

<a id="internal-agent-ratelimit-go-allow"></a>
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

<a id="internal-agent-ratelimit-go-newratelimiter"></a>
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

<a id="internal-agent-ratelimit-go-ratelimiter"></a>
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

<a id="internal-agent-ratelimit-go-wait"></a>
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

<a id="internal-agent-ratelimit-go-refill"></a>
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

<a id="internal-agent-ratelimit-test-go"></a>
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

<a id="internal-agent-reviewer-go-newreviewer"></a>
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

<a id="internal-agent-reviewer-go-review"></a>
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

<a id="internal-agent-reviewer-go-reviewinput"></a>
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

<a id="internal-agent-reviewer-go-reviewresult"></a>
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

<a id="internal-agent-reviewer-go-reviewer"></a>
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

<a id="internal-agent-reviewer-go-toevidence"></a>
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

<a id="internal-agent-reviewer-go-toprescreenevidence"></a>
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

<a id="internal-agent-reviewer-go-joinmodels"></a>
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

<a id="internal-agent-reviewer-test-go"></a>
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

<a id="internal-agent-router-go-modelfor"></a>
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

<a id="internal-agent-router-go-newrouter"></a>
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

<a id="internal-agent-router-go-router"></a>
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

<a id="internal-agent-router-go-directassignment"></a>
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

<a id="internal-agent-router-test-go"></a>
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

<a id="internal-agent-schemas-go-decisionresponse"></a>
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

<a id="internal-agent-schemas-go-prescreenresponse"></a>
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

<a id="internal-agent-schemas-go-remediationresponse"></a>
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

<a id="internal-agent-schemas-go-remediationstep"></a>
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

<a id="internal-agent-schemas-go-scoringresponse"></a>
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

<a id="internal-agent-schemas-test-go"></a>
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

<a id="internal-agent-stage-go-execute"></a>
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

<a id="internal-agent-stage-go-name"></a>
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

<a id="internal-agent-stage-go-newprescreenstage"></a>
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

<a id="internal-agent-stage-go-newreviewstage"></a>
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

<a id="internal-agent-stage-go-newscoringstage"></a>
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

<a id="internal-agent-stage-go-stage"></a>
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

<a id="internal-agent-stage-go-stageinput"></a>
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

<a id="internal-agent-stage-go-stageresult"></a>
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

<a id="internal-agent-stage-go-defaultscores"></a>
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

<a id="internal-agent-stage-go-extractjson"></a>
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

<a id="internal-agent-stage-go-looseparseneedsreview"></a>
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

<a id="internal-agent-stage-go-prescreenstage"></a>
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

<a id="internal-agent-stage-go-reviewstage"></a>
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

<a id="internal-agent-stage-go-scoringstage"></a>
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

<a id="internal-agent-stage-test-go"></a>
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

<a id="internal-agent-suggest-go-reposummary"></a>
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

<a id="internal-agent-suggest-go-scansuggestion"></a>
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

<a id="internal-agent-suggest-go-suggestforrepo"></a>
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

<a id="internal-agent-suggest-go-buildsuggestprompt"></a>
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

<a id="internal-agent-suggest-test-go"></a>
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

<a id="internal-agent-types-go-chatrequest"></a>
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

<a id="internal-agent-types-go-chatresponse"></a>
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

<a id="internal-agent-types-go-choice"></a>
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

<a id="internal-agent-types-go-content"></a>
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

<a id="internal-agent-types-go-message"></a>
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

<a id="internal-agent-types-go-modelconfig"></a>
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

<a id="internal-agent-types-go-responseformat"></a>
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

<a id="internal-agent-types-go-string"></a>
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

<a id="internal-agent-types-go-tasktype"></a>
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

<a id="internal-agent-types-go-usage"></a>
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

<a id="internal-agent-types-test-go"></a>
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
| [`Load`](#internal-config-loader-go-load) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`LoadFile`](#internal-config-loader-go-loadfile) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`LoadFromDir`](#internal-config-loader-go-loadfromdir) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`rawAgent`](#internal-config-loader-go-rawagent) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`rawConfig`](#internal-config-loader-go-rawconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`validate`](#internal-config-loader-go-validate) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`loader_test.go`](#internal-config-loader-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`FilterPolicyPacks`](#internal-config-matcher-go-filterpolicypacks) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`NewPolicyMatcher`](#internal-config-matcher-go-newpolicymatcher) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`LoadPolicyPack`](#internal-config-policy-go-loadpolicypack) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`LoadPolicyPacks`](#internal-config-policy-go-loadpolicypacks) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`parseDimension`](#internal-config-policy-go-parsedimension) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`parsePolicyPack`](#internal-config-policy-go-parsepolicypack) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`parseSeverity`](#internal-config-policy-go-parseseverity) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`rawPolicyPack`](#internal-config-policy-go-rawpolicypack) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`rawPolicyRule`](#internal-config-policy-go-rawpolicyrule) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`policy_test.go`](#internal-config-policy-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Error`](#internal-config-validator-go-error) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ValidateConfig`](#internal-config-validator-go-validateconfig) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ValidatePolicyPack`](#internal-config-validator-go-validatepolicypack) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ValidationError`](#internal-config-validator-go-validationerror) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`validator_test.go`](#internal-config-validator-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-config-loader-go-load"></a>
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

<a id="internal-config-loader-go-loadfile"></a>
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

<a id="internal-config-loader-go-loadfromdir"></a>
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

<a id="internal-config-loader-go-rawagent"></a>
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

<a id="internal-config-loader-go-rawconfig"></a>
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

<a id="internal-config-loader-go-validate"></a>
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

<a id="internal-config-loader-test-go"></a>
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

<a id="internal-config-matcher-go-filterpolicypacks"></a>
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

<a id="internal-config-matcher-go-newpolicymatcher"></a>
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

<a id="internal-config-policy-go-loadpolicypack"></a>
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

<a id="internal-config-policy-go-loadpolicypacks"></a>
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

<a id="internal-config-policy-go-parsedimension"></a>
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

<a id="internal-config-policy-go-parsepolicypack"></a>
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

<a id="internal-config-policy-go-parseseverity"></a>
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

<a id="internal-config-policy-go-rawpolicypack"></a>
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

<a id="internal-config-policy-go-rawpolicyrule"></a>
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

<a id="internal-config-policy-test-go"></a>
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

<a id="internal-config-validator-go-error"></a>
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

<a id="internal-config-validator-go-validateconfig"></a>
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

<a id="internal-config-validator-go-validatepolicypack"></a>
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

<a id="internal-config-validator-go-validationerror"></a>
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

<a id="internal-config-validator-test-go"></a>
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
| [`DetectLanguages`](#internal-discovery-detect-go-detectlanguages) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`DetectedAdapters`](#internal-discovery-detect-go-detectedadapters) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`LanguageInfo`](#internal-discovery-detect-go-languageinfo) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`buildLanguageList`](#internal-discovery-detect-go-buildlanguagelist) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`detect_test.go`](#internal-discovery-detect-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`ChangedFiles`](#internal-discovery-diff-go-changedfiles) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`DetectMoves`](#internal-discovery-diff-go-detectmoves) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FilterByPaths`](#internal-discovery-diff-go-filterbypaths) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`FilterChanged`](#internal-discovery-diff-go-filterchanged) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`MovedFile`](#internal-discovery-diff-go-movedfile) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`diff_test.go`](#internal-discovery-diff-test-go) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`GenericScanner`](#internal-discovery-generic-go-genericscanner) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`NewGenericScanner`](#internal-discovery-generic-go-newgenericscanner) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Scan`](#internal-discovery-generic-go-scan) | method | B+ | 87.2% | certified | 2026-04-23 |
| [`matchAny`](#internal-discovery-generic-go-matchany) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`GoAdapter`](#internal-discovery-go-adapter-go-goadapter) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`NewGoAdapter`](#internal-discovery-go-adapter-go-newgoadapter) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Scan`](#internal-discovery-go-adapter-go-scan) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`parseFile`](#internal-discovery-go-adapter-go-parsefile) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`go_adapter_test.go`](#internal-discovery-go-adapter-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Diff`](#internal-discovery-index-go-diff) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`DiffResult`](#internal-discovery-index-go-diffresult) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Index`](#internal-discovery-index-go-index) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`LoadIndex`](#internal-discovery-index-go-loadindex) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewIndex`](#internal-discovery-index-go-newindex) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Save`](#internal-discovery-index-go-save) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Units`](#internal-discovery-index-go-units) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`indexEntry`](#internal-discovery-index-go-indexentry) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`index_test.go`](#internal-discovery-index-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`DeduplicateFileLevel`](#internal-discovery-scanner-go-deduplicatefilelevel) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`Merge`](#internal-discovery-scanner-go-merge) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`Scanner`](#internal-discovery-scanner-go-scanner) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`UnitList`](#internal-discovery-scanner-go-unitlist) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`scanner_test.go`](#internal-discovery-scanner-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`NewTSAdapter`](#internal-discovery-ts-adapter-go-newtsadapter) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Scan`](#internal-discovery-ts-adapter-go-scan) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`TSAdapter`](#internal-discovery-ts-adapter-go-tsadapter) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`parseFile`](#internal-discovery-ts-adapter-go-parsefile) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`ts_adapter_test.go`](#internal-discovery-ts-adapter-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-discovery-detect-go-detectlanguages"></a>
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

<a id="internal-discovery-detect-go-detectedadapters"></a>
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

<a id="internal-discovery-detect-go-languageinfo"></a>
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

<a id="internal-discovery-detect-go-buildlanguagelist"></a>
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

<a id="internal-discovery-detect-test-go"></a>
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

<a id="internal-discovery-diff-go-changedfiles"></a>
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

<a id="internal-discovery-diff-go-detectmoves"></a>
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

<a id="internal-discovery-diff-go-filterbypaths"></a>
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

<a id="internal-discovery-diff-go-filterchanged"></a>
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

<a id="internal-discovery-diff-go-movedfile"></a>
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

<a id="internal-discovery-diff-test-go"></a>
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

<a id="internal-discovery-generic-go-genericscanner"></a>
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

<a id="internal-discovery-generic-go-newgenericscanner"></a>
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

<a id="internal-discovery-generic-go-scan"></a>
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

<a id="internal-discovery-generic-go-matchany"></a>
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

<a id="internal-discovery-go-adapter-go-goadapter"></a>
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

<a id="internal-discovery-go-adapter-go-newgoadapter"></a>
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

<a id="internal-discovery-go-adapter-go-scan"></a>
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

<a id="internal-discovery-go-adapter-go-parsefile"></a>
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

<a id="internal-discovery-go-adapter-test-go"></a>
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

<a id="internal-discovery-index-go-diff"></a>
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

<a id="internal-discovery-index-go-diffresult"></a>
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

<a id="internal-discovery-index-go-index"></a>
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

<a id="internal-discovery-index-go-loadindex"></a>
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

<a id="internal-discovery-index-go-newindex"></a>
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

<a id="internal-discovery-index-go-save"></a>
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

<a id="internal-discovery-index-go-units"></a>
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

<a id="internal-discovery-index-go-indexentry"></a>
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

<a id="internal-discovery-index-test-go"></a>
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

<a id="internal-discovery-scanner-go-deduplicatefilelevel"></a>
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

<a id="internal-discovery-scanner-go-merge"></a>
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

<a id="internal-discovery-scanner-go-scanner"></a>
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

<a id="internal-discovery-scanner-go-unitlist"></a>
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

<a id="internal-discovery-scanner-test-go"></a>
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

<a id="internal-discovery-ts-adapter-go-newtsadapter"></a>
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

<a id="internal-discovery-ts-adapter-go-scan"></a>
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

<a id="internal-discovery-ts-adapter-go-tsadapter"></a>
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

<a id="internal-discovery-ts-adapter-go-parsefile"></a>
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

<a id="internal-discovery-ts-adapter-test-go"></a>
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
| [`AgentConfig`](#internal-domain-config-go-agentconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`AnalyzerConfig`](#internal-domain-config-go-analyzerconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`CertificationMode`](#internal-domain-config-go-certificationmode) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Config`](#internal-domain-config-go-config) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`DefaultConfig`](#internal-domain-config-go-defaultconfig) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`EnforcingConfig`](#internal-domain-config-go-enforcingconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ExpiryConfig`](#internal-domain-config-go-expiryconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`IssueConfig`](#internal-domain-config-go-issueconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ModelAssignments`](#internal-domain-config-go-modelassignments) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`PolicyConfig`](#internal-domain-config-go-policyconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ProviderConfig`](#internal-domain-config-go-providerconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`RateLimitConfig`](#internal-domain-config-go-ratelimitconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ScheduleConfig`](#internal-domain-config-go-scheduleconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ScopeConfig`](#internal-domain-config-go-scopeconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`SignoffConfig`](#internal-domain-config-go-signoffconfig) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`String`](#internal-domain-config-go-string) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`config_test.go`](#internal-domain-config-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`AllDimensions`](#internal-domain-dimension-go-alldimensions) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Dimension`](#internal-domain-dimension-go-dimension) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`DimensionScores`](#internal-domain-dimension-go-dimensionscores) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`DimensionWeights`](#internal-domain-dimension-go-dimensionweights) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Grade`](#internal-domain-dimension-go-grade) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`GradeFromScore`](#internal-domain-dimension-go-gradefromscore) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`String`](#internal-domain-dimension-go-string) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`WeightedAverage`](#internal-domain-dimension-go-weightedaverage) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`dimension_test.go`](#internal-domain-dimension-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Evidence`](#internal-domain-evidence-go-evidence) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`EvidenceKind`](#internal-domain-evidence-go-evidencekind) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseSeverity`](#internal-domain-evidence-go-parseseverity) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Severity`](#internal-domain-evidence-go-severity) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`String`](#internal-domain-evidence-go-string) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](#internal-domain-evidence-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`evidence_test.go`](#internal-domain-evidence-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Duration`](#internal-domain-expiry-go-duration) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ExpiryFactors`](#internal-domain-expiry-go-expiryfactors) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ExpiryWindow`](#internal-domain-expiry-go-expirywindow) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`IsExpired`](#internal-domain-expiry-go-isexpired) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`RemainingAt`](#internal-domain-expiry-go-remainingat) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`expiry_test.go`](#internal-domain-expiry-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Override`](#internal-domain-override-go-override) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`OverrideAction`](#internal-domain-override-go-overrideaction) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`String`](#internal-domain-override-go-string) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Validate`](#internal-domain-override-go-validate) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`override_test.go`](#internal-domain-override-test-go) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`IsGlobal`](#internal-domain-policy-go-isglobal) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`PolicyPack`](#internal-domain-policy-go-policypack) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`PolicyRule`](#internal-domain-policy-go-policyrule) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Violation`](#internal-domain-policy-go-violation) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`policy_test.go`](#internal-domain-policy-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`CertificationRecord`](#internal-domain-record-go-certificationrecord) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`IsPassing`](#internal-domain-record-go-ispassing) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseStatus`](#internal-domain-record-go-parsestatus) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Status`](#internal-domain-record-go-status) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`String`](#internal-domain-record-go-string) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](#internal-domain-record-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`record_test.go`](#internal-domain-record-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Language`](#internal-domain-unit-go-language) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewUnit`](#internal-domain-unit-go-newunit) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewUnitID`](#internal-domain-unit-go-newunitid) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseUnitID`](#internal-domain-unit-go-parseunitid) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseUnitType`](#internal-domain-unit-go-parseunittype) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Path`](#internal-domain-unit-go-path) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`String`](#internal-domain-unit-go-string) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Symbol`](#internal-domain-unit-go-symbol) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Unit`](#internal-domain-unit-go-unit) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`UnitID`](#internal-domain-unit-go-unitid) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`UnitType`](#internal-domain-unit-go-unittype) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](#internal-domain-unit-go-init) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`unit_test.go`](#internal-domain-unit-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-domain-config-go-agentconfig"></a>
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

<a id="internal-domain-config-go-analyzerconfig"></a>
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

<a id="internal-domain-config-go-certificationmode"></a>
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

<a id="internal-domain-config-go-config"></a>
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

<a id="internal-domain-config-go-defaultconfig"></a>
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

<a id="internal-domain-config-go-enforcingconfig"></a>
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

<a id="internal-domain-config-go-expiryconfig"></a>
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

<a id="internal-domain-config-go-issueconfig"></a>
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

<a id="internal-domain-config-go-modelassignments"></a>
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

<a id="internal-domain-config-go-policyconfig"></a>
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

<a id="internal-domain-config-go-providerconfig"></a>
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

<a id="internal-domain-config-go-ratelimitconfig"></a>
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

<a id="internal-domain-config-go-scheduleconfig"></a>
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

<a id="internal-domain-config-go-scopeconfig"></a>
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

<a id="internal-domain-config-go-signoffconfig"></a>
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

<a id="internal-domain-config-go-string"></a>
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

<a id="internal-domain-config-test-go"></a>
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

<a id="internal-domain-dimension-go-alldimensions"></a>
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

<a id="internal-domain-dimension-go-dimension"></a>
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

<a id="internal-domain-dimension-go-dimensionscores"></a>
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

<a id="internal-domain-dimension-go-dimensionweights"></a>
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

<a id="internal-domain-dimension-go-grade"></a>
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

<a id="internal-domain-dimension-go-gradefromscore"></a>
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

<a id="internal-domain-dimension-go-string"></a>
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

<a id="internal-domain-dimension-go-weightedaverage"></a>
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

<a id="internal-domain-dimension-test-go"></a>
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

<a id="internal-domain-evidence-go-evidence"></a>
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

<a id="internal-domain-evidence-go-evidencekind"></a>
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

<a id="internal-domain-evidence-go-parseseverity"></a>
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

<a id="internal-domain-evidence-go-severity"></a>
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

<a id="internal-domain-evidence-go-string"></a>
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

<a id="internal-domain-evidence-go-init"></a>
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

<a id="internal-domain-evidence-test-go"></a>
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

<a id="internal-domain-expiry-go-duration"></a>
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

<a id="internal-domain-expiry-go-expiryfactors"></a>
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

<a id="internal-domain-expiry-go-expirywindow"></a>
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

<a id="internal-domain-expiry-go-isexpired"></a>
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

<a id="internal-domain-expiry-go-remainingat"></a>
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

<a id="internal-domain-expiry-test-go"></a>
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

<a id="internal-domain-override-go-override"></a>
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

<a id="internal-domain-override-go-overrideaction"></a>
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

<a id="internal-domain-override-go-string"></a>
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

<a id="internal-domain-override-go-validate"></a>
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

<a id="internal-domain-override-test-go"></a>
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

<a id="internal-domain-policy-go-isglobal"></a>
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

<a id="internal-domain-policy-go-policypack"></a>
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

<a id="internal-domain-policy-go-policyrule"></a>
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

<a id="internal-domain-policy-go-violation"></a>
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

<a id="internal-domain-policy-test-go"></a>
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

<a id="internal-domain-record-go-certificationrecord"></a>
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

<a id="internal-domain-record-go-ispassing"></a>
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

<a id="internal-domain-record-go-parsestatus"></a>
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

<a id="internal-domain-record-go-status"></a>
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

<a id="internal-domain-record-go-string"></a>
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

<a id="internal-domain-record-go-init"></a>
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

<a id="internal-domain-record-test-go"></a>
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

<a id="internal-domain-unit-go-language"></a>
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

<a id="internal-domain-unit-go-newunit"></a>
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

<a id="internal-domain-unit-go-newunitid"></a>
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

<a id="internal-domain-unit-go-parseunitid"></a>
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

<a id="internal-domain-unit-go-parseunittype"></a>
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

<a id="internal-domain-unit-go-path"></a>
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

<a id="internal-domain-unit-go-string"></a>
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

<a id="internal-domain-unit-go-symbol"></a>
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

<a id="internal-domain-unit-go-unit"></a>
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

<a id="internal-domain-unit-go-unitid"></a>
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

<a id="internal-domain-unit-go-unittype"></a>
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

<a id="internal-domain-unit-go-init"></a>
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

<a id="internal-domain-unit-test-go"></a>
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
| [`CertifyUnit`](#internal-engine-pipeline-go-certifyunit) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`pipeline_test.go`](#internal-engine-pipeline-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Score`](#internal-engine-scorer-go-score) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`StatusFromScore`](#internal-engine-scorer-go-statusfromscore) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`extractSummaryFloat`](#internal-engine-scorer-go-extractsummaryfloat) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`extractSummaryInt`](#internal-engine-scorer-go-extractsummaryint) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`scoreFromGitHistory`](#internal-engine-scorer-go-scorefromgithistory) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`scoreFromMetrics`](#internal-engine-scorer-go-scorefrommetrics) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`severityPenalty`](#internal-engine-scorer-go-severitypenalty) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`scorer_test.go`](#internal-engine-scorer-test-go) | file | B+ | 87.2% | certified | 2026-04-23 |

<a id="internal-engine-pipeline-go-certifyunit"></a>
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

<a id="internal-engine-pipeline-test-go"></a>
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

<a id="internal-engine-scorer-go-score"></a>
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

<a id="internal-engine-scorer-go-statusfromscore"></a>
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

<a id="internal-engine-scorer-go-extractsummaryfloat"></a>
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

<a id="internal-engine-scorer-go-extractsummaryint"></a>
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

<a id="internal-engine-scorer-go-scorefromgithistory"></a>
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

<a id="internal-engine-scorer-go-scorefrommetrics"></a>
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

<a id="internal-engine-scorer-go-severitypenalty"></a>
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

<a id="internal-engine-scorer-test-go"></a>
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
| [`Collector`](#internal-evidence-collector-go-collector) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ComputeGoComplexity`](#internal-evidence-complexity-go-computegocomplexity) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ComputeSymbolMetrics`](#internal-evidence-complexity-go-computesymbolmetrics) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`funcName`](#internal-evidence-complexity-go-funcname) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`complexity_test.go`](#internal-evidence-complexity-test-go) | file | B+ | 87.2% | certified | 2026-04-23 |
| [`CollectAll`](#internal-evidence-executor-go-collectall) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`HasGoMod`](#internal-evidence-executor-go-hasgomod) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`HasPackageJSON`](#internal-evidence-executor-go-haspackagejson) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewToolExecutor`](#internal-evidence-executor-go-newtoolexecutor) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ToolExecutor`](#internal-evidence-executor-go-toolexecutor) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`runGitStats`](#internal-evidence-executor-go-rungitstats) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`runGoTest`](#internal-evidence-executor-go-rungotest) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`runGoVet`](#internal-evidence-executor-go-rungovet) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`runGolangciLint`](#internal-evidence-executor-go-rungolangcilint) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ChurnRate`](#internal-evidence-git-go-churnrate) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`GitStats`](#internal-evidence-git-go-gitstats) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseGitLog`](#internal-evidence-git-go-parsegitlog) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ToEvidence`](#internal-evidence-git-go-toevidence) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`git_test.go`](#internal-evidence-git-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`LintFinding`](#internal-evidence-lint-go-lintfinding) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`LintResult`](#internal-evidence-lint-go-lintresult) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`TestResult`](#internal-evidence-lint-go-testresult) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ToEvidence`](#internal-evidence-lint-go-toevidence) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`lint_test.go`](#internal-evidence-lint-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`CodeMetrics`](#internal-evidence-metrics-go-codemetrics) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ComputeMetrics`](#internal-evidence-metrics-go-computemetrics) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`ToEvidence`](#internal-evidence-metrics-go-toevidence) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`containsTodo`](#internal-evidence-metrics-go-containstodo) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`metrics_test.go`](#internal-evidence-metrics-test-go) | file | B+ | 87.2% | certified | 2026-04-23 |
| [`ParseCoverProfile`](#internal-evidence-runner-go-parsecoverprofile) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ParseGitLogWithAge`](#internal-evidence-runner-go-parsegitlogwithage) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseGoTestJSON`](#internal-evidence-runner-go-parsegotestjson) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ParseGoVet`](#internal-evidence-runner-go-parsegovet) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ParseGolangciLintJSON`](#internal-evidence-runner-go-parsegolangcilintjson) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`goTestEvent`](#internal-evidence-runner-go-gotestevent) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`golangciLintIssue`](#internal-evidence-runner-go-golangcilintissue) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`golangciLintOutput`](#internal-evidence-runner-go-golangcilintoutput) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`simpleAtoi`](#internal-evidence-runner-go-simpleatoi) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`runner_test.go`](#internal-evidence-runner-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-evidence-collector-go-collector"></a>
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

<a id="internal-evidence-complexity-go-computegocomplexity"></a>
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

<a id="internal-evidence-complexity-go-computesymbolmetrics"></a>
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

<a id="internal-evidence-complexity-go-funcname"></a>
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

<a id="internal-evidence-complexity-test-go"></a>
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

<a id="internal-evidence-executor-go-collectall"></a>
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

<a id="internal-evidence-executor-go-hasgomod"></a>
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

<a id="internal-evidence-executor-go-haspackagejson"></a>
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

<a id="internal-evidence-executor-go-newtoolexecutor"></a>
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

<a id="internal-evidence-executor-go-toolexecutor"></a>
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

<a id="internal-evidence-executor-go-rungitstats"></a>
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

<a id="internal-evidence-executor-go-rungotest"></a>
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

<a id="internal-evidence-executor-go-rungovet"></a>
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

<a id="internal-evidence-executor-go-rungolangcilint"></a>
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

<a id="internal-evidence-git-go-churnrate"></a>
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

<a id="internal-evidence-git-go-gitstats"></a>
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

<a id="internal-evidence-git-go-parsegitlog"></a>
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

<a id="internal-evidence-git-go-toevidence"></a>
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

<a id="internal-evidence-git-test-go"></a>
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

<a id="internal-evidence-lint-go-lintfinding"></a>
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

<a id="internal-evidence-lint-go-lintresult"></a>
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

<a id="internal-evidence-lint-go-testresult"></a>
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

<a id="internal-evidence-lint-go-toevidence"></a>
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

<a id="internal-evidence-lint-test-go"></a>
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

<a id="internal-evidence-metrics-go-codemetrics"></a>
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

<a id="internal-evidence-metrics-go-computemetrics"></a>
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

<a id="internal-evidence-metrics-go-toevidence"></a>
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

<a id="internal-evidence-metrics-go-containstodo"></a>
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

<a id="internal-evidence-metrics-test-go"></a>
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

<a id="internal-evidence-runner-go-parsecoverprofile"></a>
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

<a id="internal-evidence-runner-go-parsegitlogwithage"></a>
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

<a id="internal-evidence-runner-go-parsegotestjson"></a>
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

<a id="internal-evidence-runner-go-parsegovet"></a>
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

<a id="internal-evidence-runner-go-parsegolangcilintjson"></a>
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

<a id="internal-evidence-runner-go-gotestevent"></a>
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

<a id="internal-evidence-runner-go-golangcilintissue"></a>
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

<a id="internal-evidence-runner-go-golangcilintoutput"></a>
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

<a id="internal-evidence-runner-go-simpleatoi"></a>
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

<a id="internal-evidence-runner-test-go"></a>
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
| [`Calculate`](#internal-expiry-calculator-go-calculate) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`calculator_test.go`](#internal-expiry-calculator-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-expiry-calculator-go-calculate"></a>
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

<a id="internal-expiry-calculator-test-go"></a>
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
| [`BuildIssueCloseCommand`](#internal-github-issues-go-buildissueclosecommand) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`BuildIssueCreateCommand`](#internal-github-issues-go-buildissuecreatecommand) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`BuildIssueSearchCommand`](#internal-github-issues-go-buildissuesearchcommand) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`BuildIssueUpdateCommand`](#internal-github-issues-go-buildissueupdatecommand) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatGroupedIssueBody`](#internal-github-issues-go-formatgroupedissuebody) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatIssueBody`](#internal-github-issues-go-formatissuebody) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatIssueTitle`](#internal-github-issues-go-formatissuetitle) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`issues_test.go`](#internal-github-issues-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`BuildPRCommentCommand`](#internal-github-pr-go-buildprcommentcommand) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ComputeTrustDelta`](#internal-github-pr-go-computetrustdelta) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`FormatPRComment`](#internal-github-pr-go-formatprcomment) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`TrustDelta`](#internal-github-pr-go-trustdelta) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`pr_test.go`](#internal-github-pr-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`GenerateNightlyWorkflow`](#internal-github-workflows-go-generatenightlyworkflow) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`GeneratePRWorkflow`](#internal-github-workflows-go-generateprworkflow) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`GenerateWeeklyWorkflow`](#internal-github-workflows-go-generateweeklyworkflow) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`workflows_test.go`](#internal-github-workflows-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-github-issues-go-buildissueclosecommand"></a>
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

<a id="internal-github-issues-go-buildissuecreatecommand"></a>
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

<a id="internal-github-issues-go-buildissuesearchcommand"></a>
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

<a id="internal-github-issues-go-buildissueupdatecommand"></a>
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

<a id="internal-github-issues-go-formatgroupedissuebody"></a>
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

<a id="internal-github-issues-go-formatissuebody"></a>
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

<a id="internal-github-issues-go-formatissuetitle"></a>
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

<a id="internal-github-issues-test-go"></a>
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

<a id="internal-github-pr-go-buildprcommentcommand"></a>
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

<a id="internal-github-pr-go-computetrustdelta"></a>
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

<a id="internal-github-pr-go-formatprcomment"></a>
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

<a id="internal-github-pr-go-trustdelta"></a>
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

<a id="internal-github-pr-test-go"></a>
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

<a id="internal-github-workflows-go-generatenightlyworkflow"></a>
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

<a id="internal-github-workflows-go-generateprworkflow"></a>
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

<a id="internal-github-workflows-go-generateweeklyworkflow"></a>
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

<a id="internal-github-workflows-test-go"></a>
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
| [`Apply`](#internal-override-applier-go-apply) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ApplyAll`](#internal-override-applier-go-applyall) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`applier_test.go`](#internal-override-applier-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`LoadDir`](#internal-override-loader-go-loaddir) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`LoadFile`](#internal-override-loader-go-loadfile) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`parseAction`](#internal-override-loader-go-parseaction) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`rawOverride`](#internal-override-loader-go-rawoverride) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`rawOverrideFile`](#internal-override-loader-go-rawoverridefile) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`loader_test.go`](#internal-override-loader-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-override-applier-go-apply"></a>
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

<a id="internal-override-applier-go-applyall"></a>
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

<a id="internal-override-applier-test-go"></a>
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

<a id="internal-override-loader-go-loaddir"></a>
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

<a id="internal-override-loader-go-loadfile"></a>
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

<a id="internal-override-loader-go-parseaction"></a>
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

<a id="internal-override-loader-go-rawoverride"></a>
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

<a id="internal-override-loader-go-rawoverridefile"></a>
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

<a id="internal-override-loader-test-go"></a>
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
| [`Evaluate`](#internal-policy-evaluator-go-evaluate) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`EvaluationResult`](#internal-policy-evaluator-go-evaluationresult) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`evaluateRule`](#internal-policy-evaluator-go-evaluaterule) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`extractComplexity`](#internal-policy-evaluator-go-extractcomplexity) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`extractCoverage`](#internal-policy-evaluator-go-extractcoverage) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`extractMetric`](#internal-policy-evaluator-go-extractmetric) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`extractTodoCount`](#internal-policy-evaluator-go-extracttodocount) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`evaluator_test.go`](#internal-policy-evaluator-test-go) | file | B+ | 87.2% | certified | 2026-04-23 |
| [`Match`](#internal-policy-matcher-go-match) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Matcher`](#internal-policy-matcher-go-matcher) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`NewMatcher`](#internal-policy-matcher-go-newmatcher) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`matchPath`](#internal-policy-matcher-go-matchpath) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`matchesPack`](#internal-policy-matcher-go-matchespack) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`matcher_test.go`](#internal-policy-matcher-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-policy-evaluator-go-evaluate"></a>
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

<a id="internal-policy-evaluator-go-evaluationresult"></a>
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

<a id="internal-policy-evaluator-go-evaluaterule"></a>
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

<a id="internal-policy-evaluator-go-extractcomplexity"></a>
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

<a id="internal-policy-evaluator-go-extractcoverage"></a>
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

<a id="internal-policy-evaluator-go-extractmetric"></a>
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

<a id="internal-policy-evaluator-go-extracttodocount"></a>
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

<a id="internal-policy-evaluator-test-go"></a>
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

<a id="internal-policy-matcher-go-match"></a>
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

<a id="internal-policy-matcher-go-matcher"></a>
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

<a id="internal-policy-matcher-go-newmatcher"></a>
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

<a id="internal-policy-matcher-go-matchpath"></a>
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

<a id="internal-policy-matcher-go-matchespack"></a>
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

<a id="internal-policy-matcher-test-go"></a>
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
| [`BatchNext`](#internal-queue-queue-go-batchnext) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`Complete`](#internal-queue-queue-go-complete) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Enqueue`](#internal-queue-queue-go-enqueue) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Fail`](#internal-queue-queue-go-fail) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Item`](#internal-queue-queue-go-item) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ItemStatus`](#internal-queue-queue-go-itemstatus) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Len`](#internal-queue-queue-go-len) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Load`](#internal-queue-queue-go-load) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`New`](#internal-queue-queue-go-new) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Next`](#internal-queue-queue-go-next) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Queue`](#internal-queue-queue-go-queue) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Reset`](#internal-queue-queue-go-reset) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Save`](#internal-queue-queue-go-save) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Skip`](#internal-queue-queue-go-skip) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Stats`](#internal-queue-queue-go-stats) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`persistedQueue`](#internal-queue-queue-go-persistedqueue) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`queue_test.go`](#internal-queue-queue-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-queue-queue-go-batchnext"></a>
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

<a id="internal-queue-queue-go-complete"></a>
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

<a id="internal-queue-queue-go-enqueue"></a>
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

<a id="internal-queue-queue-go-fail"></a>
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

<a id="internal-queue-queue-go-item"></a>
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

<a id="internal-queue-queue-go-itemstatus"></a>
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

<a id="internal-queue-queue-go-len"></a>
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

<a id="internal-queue-queue-go-load"></a>
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

<a id="internal-queue-queue-go-new"></a>
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

<a id="internal-queue-queue-go-next"></a>
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

<a id="internal-queue-queue-go-queue"></a>
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

<a id="internal-queue-queue-go-reset"></a>
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

<a id="internal-queue-queue-go-save"></a>
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

<a id="internal-queue-queue-go-skip"></a>
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

<a id="internal-queue-queue-go-stats"></a>
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

<a id="internal-queue-queue-go-persistedqueue"></a>
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

<a id="internal-queue-queue-test-go"></a>
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
| [`AppendHistory`](#internal-record-store-go-appendhistory) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ListAll`](#internal-record-store-go-listall) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`Load`](#internal-record-store-go-load) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`LoadHistory`](#internal-record-store-go-loadhistory) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`NewStore`](#internal-record-store-go-newstore) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Save`](#internal-record-store-go-save) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Store`](#internal-record-store-go-store) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`dimensionsToMap`](#internal-record-store-go-dimensionstomap) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`fromJSON`](#internal-record-store-go-fromjson) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`historyEntry`](#internal-record-store-go-historyentry) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`historyPathFor`](#internal-record-store-go-historypathfor) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`mapToDimensions`](#internal-record-store-go-maptodimensions) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`parseGrade`](#internal-record-store-go-parsegrade) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`pathFor`](#internal-record-store-go-pathfor) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`recordJSON`](#internal-record-store-go-recordjson) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`toJSON`](#internal-record-store-go-tojson) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`store_test.go`](#internal-record-store-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-record-store-go-appendhistory"></a>
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

<a id="internal-record-store-go-listall"></a>
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

<a id="internal-record-store-go-load"></a>
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

<a id="internal-record-store-go-loadhistory"></a>
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

<a id="internal-record-store-go-newstore"></a>
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

<a id="internal-record-store-go-save"></a>
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

<a id="internal-record-store-go-store"></a>
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

<a id="internal-record-store-go-dimensionstomap"></a>
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

<a id="internal-record-store-go-fromjson"></a>
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

<a id="internal-record-store-go-historyentry"></a>
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

<a id="internal-record-store-go-historypathfor"></a>
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

<a id="internal-record-store-go-maptodimensions"></a>
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

<a id="internal-record-store-go-parsegrade"></a>
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

<a id="internal-record-store-go-pathfor"></a>
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

<a id="internal-record-store-go-recordjson"></a>
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

<a id="internal-record-store-go-tojson"></a>
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

<a id="internal-record-store-test-go"></a>
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
| [`Badge`](#internal-report-badge-go-badge) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`BadgeMarkdown`](#internal-report-badge-go-badgemarkdown) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatBadgeJSON`](#internal-report-badge-go-formatbadgejson) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`GenerateBadge`](#internal-report-badge-go-generatebadge) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`badgeColor`](#internal-report-badge-go-badgecolor) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`badgeMessage`](#internal-report-badge-go-badgemessage) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`badge_test.go`](#internal-report-badge-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Card`](#internal-report-card-go-card) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatCardMarkdown`](#internal-report-card-go-formatcardmarkdown) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`FormatCardText`](#internal-report-card-go-formatcardtext) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`GenerateCard`](#internal-report-card-go-generatecard) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`IssueCard`](#internal-report-card-go-issuecard) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`LanguageCard`](#internal-report-card-go-languagecard) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`buildLanguageCards`](#internal-report-card-go-buildlanguagecards) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`buildTopIssues`](#internal-report-card-go-buildtopissues) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`gradeEmoji`](#internal-report-card-go-gradeemoji) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`card_test.go`](#internal-report-card-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`AreaSummary`](#internal-report-detailed-go-areasummary) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Detailed`](#internal-report-detailed-go-detailed) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`DetailedReport`](#internal-report-detailed-go-detailedreport) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatDetailedText`](#internal-report-detailed-go-formatdetailedtext) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`LanguageBreakdown`](#internal-report-detailed-go-languagebreakdown) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`UnitSummary`](#internal-report-detailed-go-unitsummary) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`computeDimensionAverages`](#internal-report-detailed-go-computedimensionaverages) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`computeLanguageBreakdowns`](#internal-report-detailed-go-computelanguagebreakdowns) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`explainStatus`](#internal-report-detailed-go-explainstatus) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`findExpiringSoon`](#internal-report-detailed-go-findexpiringsoon) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`findFailing`](#internal-report-detailed-go-findfailing) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`findHighestRisk`](#internal-report-detailed-go-findhighestrisk) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`findRecurrentlyFailing`](#internal-report-detailed-go-findrecurrentlyfailing) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`unitSummaryFrom`](#internal-report-detailed-go-unitsummaryfrom) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`detailed_test.go`](#internal-report-detailed-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`FormatFullMarkdown`](#internal-report-full-go-formatfullmarkdown) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FullReport`](#internal-report-full-go-fullreport) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`GenerateFullReport`](#internal-report-full-go-generatefullreport) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`LanguageDetail`](#internal-report-full-go-languagedetail) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`UnitReport`](#internal-report-full-go-unitreport) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`buildLanguageDetail`](#internal-report-full-go-buildlanguagedetail) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`dirOf`](#internal-report-full-go-dirof) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`shortFile`](#internal-report-full-go-shortfile) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`sortedKeys`](#internal-report-full-go-sortedkeys) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`unitReportFrom`](#internal-report-full-go-unitreportfrom) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`writeAIInsights`](#internal-report-full-go-writeaiinsights) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`writeAllUnits`](#internal-report-full-go-writeallunits) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`writeDimensionAverages`](#internal-report-full-go-writedimensionaverages) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`writeGradeDistribution`](#internal-report-full-go-writegradedistribution) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`writeHeader`](#internal-report-full-go-writeheader) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`writeLanguageDetail`](#internal-report-full-go-writelanguagedetail) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`writeSummary`](#internal-report-full-go-writesummary) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`writeUnitDetails`](#internal-report-full-go-writeunitdetails) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`full_test.go`](#internal-report-full-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`FormatJSON`](#internal-report-health-go-formatjson) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatText`](#internal-report-health-go-formattext) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Health`](#internal-report-health-go-health) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`HealthReport`](#internal-report-health-go-healthreport) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`health_test.go`](#internal-report-health-test-go) | file | B+ | 88.3% | certified | 2026-04-23 |

<a id="internal-report-badge-go-badge"></a>
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

<a id="internal-report-badge-go-badgemarkdown"></a>
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

<a id="internal-report-badge-go-formatbadgejson"></a>
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

<a id="internal-report-badge-go-generatebadge"></a>
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

<a id="internal-report-badge-go-badgecolor"></a>
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

<a id="internal-report-badge-go-badgemessage"></a>
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

<a id="internal-report-badge-test-go"></a>
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

<a id="internal-report-card-go-card"></a>
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

<a id="internal-report-card-go-formatcardmarkdown"></a>
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

<a id="internal-report-card-go-formatcardtext"></a>
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

<a id="internal-report-card-go-generatecard"></a>
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

<a id="internal-report-card-go-issuecard"></a>
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

<a id="internal-report-card-go-languagecard"></a>
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

<a id="internal-report-card-go-buildlanguagecards"></a>
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

<a id="internal-report-card-go-buildtopissues"></a>
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

<a id="internal-report-card-go-gradeemoji"></a>
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

<a id="internal-report-card-test-go"></a>
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

<a id="internal-report-detailed-go-areasummary"></a>
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

<a id="internal-report-detailed-go-detailed"></a>
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

<a id="internal-report-detailed-go-detailedreport"></a>
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

<a id="internal-report-detailed-go-formatdetailedtext"></a>
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

<a id="internal-report-detailed-go-languagebreakdown"></a>
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

<a id="internal-report-detailed-go-unitsummary"></a>
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

<a id="internal-report-detailed-go-computedimensionaverages"></a>
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

<a id="internal-report-detailed-go-computelanguagebreakdowns"></a>
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

<a id="internal-report-detailed-go-explainstatus"></a>
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

<a id="internal-report-detailed-go-findexpiringsoon"></a>
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

<a id="internal-report-detailed-go-findfailing"></a>
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

<a id="internal-report-detailed-go-findhighestrisk"></a>
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

<a id="internal-report-detailed-go-findrecurrentlyfailing"></a>
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

<a id="internal-report-detailed-go-unitsummaryfrom"></a>
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

<a id="internal-report-detailed-test-go"></a>
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

<a id="internal-report-full-go-formatfullmarkdown"></a>
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

<a id="internal-report-full-go-fullreport"></a>
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

<a id="internal-report-full-go-generatefullreport"></a>
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

<a id="internal-report-full-go-languagedetail"></a>
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

<a id="internal-report-full-go-unitreport"></a>
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

<a id="internal-report-full-go-buildlanguagedetail"></a>
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

<a id="internal-report-full-go-dirof"></a>
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

<a id="internal-report-full-go-shortfile"></a>
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

<a id="internal-report-full-go-sortedkeys"></a>
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

<a id="internal-report-full-go-unitreportfrom"></a>
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

<a id="internal-report-full-go-writeaiinsights"></a>
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

<a id="internal-report-full-go-writeallunits"></a>
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

<a id="internal-report-full-go-writedimensionaverages"></a>
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

<a id="internal-report-full-go-writegradedistribution"></a>
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

<a id="internal-report-full-go-writeheader"></a>
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

<a id="internal-report-full-go-writelanguagedetail"></a>
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

<a id="internal-report-full-go-writesummary"></a>
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

<a id="internal-report-full-go-writeunitdetails"></a>
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

<a id="internal-report-full-test-go"></a>
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

<a id="internal-report-health-go-formatjson"></a>
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

<a id="internal-report-health-go-formattext"></a>
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

<a id="internal-report-health-go-health"></a>
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

<a id="internal-report-health-go-healthreport"></a>
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

<a id="internal-report-health-test-go"></a>
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
| [`DialogueParser`](#testdata-repos-ts-simple-src-parser-ts-dialogueparser) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`MAX_TOKENS`](#testdata-repos-ts-simple-src-parser-ts-max-tokens) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`parseNode`](#testdata-repos-ts-simple-src-parser-ts-parsenode) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`tokenizeDialogue`](#testdata-repos-ts-simple-src-parser-ts-tokenizedialogue) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`formatDate`](#testdata-repos-ts-simple-src-utils-ts-formatdate) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`log`](#testdata-repos-ts-simple-src-utils-ts-log) | function | B+ | 89.4% | certified | 2026-04-23 |

<a id="testdata-repos-ts-simple-src-parser-ts-dialogueparser"></a>
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

<a id="testdata-repos-ts-simple-src-parser-ts-max-tokens"></a>
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

<a id="testdata-repos-ts-simple-src-parser-ts-parsenode"></a>
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

<a id="testdata-repos-ts-simple-src-parser-ts-tokenizedialogue"></a>
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

<a id="testdata-repos-ts-simple-src-utils-ts-formatdate"></a>
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

<a id="testdata-repos-ts-simple-src-utils-ts-log"></a>
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
| [`RunResult`](#vscode-certify-src-certifybinary-ts-runresult) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`findCertifyBinary`](#vscode-certify-src-certifybinary-ts-findcertifybinary) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`listModels`](#vscode-certify-src-certifybinary-ts-listmodels) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`promptInstall`](#vscode-certify-src-certifybinary-ts-promptinstall) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`runCertify`](#vscode-certify-src-certifybinary-ts-runcertify) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`runCertifyJSON`](#vscode-certify-src-certifybinary-ts-runcertifyjson) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`runInTerminal`](#vscode-certify-src-certifybinary-ts-runinterminal) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`BRAND_COLORS`](#vscode-certify-src-constants-ts-brand-colors) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`DIMENSION_NAMES`](#vscode-certify-src-constants-ts-dimension-names) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`GRADE_COLORS`](#vscode-certify-src-constants-ts-grade-colors) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`GRADE_EMOJI`](#vscode-certify-src-constants-ts-grade-emoji) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`PROVIDER_PRESETS`](#vscode-certify-src-constants-ts-provider-presets) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`CertifyDataLoader`](#vscode-certify-src-dataloader-ts-certifydataloader) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`activate`](#vscode-certify-src-extension-ts-activate) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`deactivate`](#vscode-certify-src-extension-ts-deactivate) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`createStatusBarItem`](#vscode-certify-src-statusbar-ts-createstatusbaritem) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`AgentConfig`](#vscode-certify-src-types-ts-agentconfig) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`BadgeJSON`](#vscode-certify-src-types-ts-badgejson) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`CertifyCard`](#vscode-certify-src-types-ts-certifycard) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`CertifyConfig`](#vscode-certify-src-types-ts-certifyconfig) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`FullReport`](#vscode-certify-src-types-ts-fullreport) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`IndexEntry`](#vscode-certify-src-types-ts-indexentry) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`IssueCard`](#vscode-certify-src-types-ts-issuecard) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`LanguageCard`](#vscode-certify-src-types-ts-languagecard) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`LanguageDetail`](#vscode-certify-src-types-ts-languagedetail) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ModelAssignments`](#vscode-certify-src-types-ts-modelassignments) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ModelInfo`](#vscode-certify-src-types-ts-modelinfo) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ProviderConfig`](#vscode-certify-src-types-ts-providerconfig) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ProviderPreset`](#vscode-certify-src-types-ts-providerpreset) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`RecordJSON`](#vscode-certify-src-types-ts-recordjson) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`UnitReport`](#vscode-certify-src-types-ts-unitreport) | class | B+ | 88.3% | certified | 2026-04-23 |

<a id="vscode-certify-src-certifybinary-ts-runresult"></a>
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

<a id="vscode-certify-src-certifybinary-ts-findcertifybinary"></a>
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

<a id="vscode-certify-src-certifybinary-ts-listmodels"></a>
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

<a id="vscode-certify-src-certifybinary-ts-promptinstall"></a>
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

<a id="vscode-certify-src-certifybinary-ts-runcertify"></a>
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

<a id="vscode-certify-src-certifybinary-ts-runcertifyjson"></a>
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

<a id="vscode-certify-src-certifybinary-ts-runinterminal"></a>
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

<a id="vscode-certify-src-constants-ts-brand-colors"></a>
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

<a id="vscode-certify-src-constants-ts-dimension-names"></a>
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

<a id="vscode-certify-src-constants-ts-grade-colors"></a>
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

<a id="vscode-certify-src-constants-ts-grade-emoji"></a>
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

<a id="vscode-certify-src-constants-ts-provider-presets"></a>
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

<a id="vscode-certify-src-dataloader-ts-certifydataloader"></a>
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

<a id="vscode-certify-src-extension-ts-activate"></a>
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

<a id="vscode-certify-src-extension-ts-deactivate"></a>
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

<a id="vscode-certify-src-statusbar-ts-createstatusbaritem"></a>
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

<a id="vscode-certify-src-types-ts-agentconfig"></a>
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

<a id="vscode-certify-src-types-ts-badgejson"></a>
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

<a id="vscode-certify-src-types-ts-certifycard"></a>
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

<a id="vscode-certify-src-types-ts-certifyconfig"></a>
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

<a id="vscode-certify-src-types-ts-fullreport"></a>
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

<a id="vscode-certify-src-types-ts-indexentry"></a>
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

<a id="vscode-certify-src-types-ts-issuecard"></a>
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

<a id="vscode-certify-src-types-ts-languagecard"></a>
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

<a id="vscode-certify-src-types-ts-languagedetail"></a>
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

<a id="vscode-certify-src-types-ts-modelassignments"></a>
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

<a id="vscode-certify-src-types-ts-modelinfo"></a>
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

<a id="vscode-certify-src-types-ts-providerconfig"></a>
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

<a id="vscode-certify-src-types-ts-providerpreset"></a>
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

<a id="vscode-certify-src-types-ts-recordjson"></a>
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

<a id="vscode-certify-src-types-ts-unitreport"></a>
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
| [`CertifyCodeLensProvider`](#vscode-certify-src-codelens-certifycodelensprovider-ts-certifycodelensprovider) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`showDimensionScores`](#vscode-certify-src-codelens-certifycodelensprovider-ts-showdimensionscores) | function | B+ | 88.3% | certified | 2026-04-23 |

<a id="vscode-certify-src-codelens-certifycodelensprovider-ts-certifycodelensprovider"></a>
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

<a id="vscode-certify-src-codelens-certifycodelensprovider-ts-showdimensionscores"></a>
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
| [`ConfigPanel`](#vscode-certify-src-config-configpanel-ts-configpanel) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ConnectionTestResult`](#vscode-certify-src-config-configwriter-ts-connectiontestresult) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`readConfig`](#vscode-certify-src-config-configwriter-ts-readconfig) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`testConnection`](#vscode-certify-src-config-configwriter-ts-testconnection) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`writeConfig`](#vscode-certify-src-config-configwriter-ts-writeconfig) | function | B+ | 88.3% | certified | 2026-04-23 |

<a id="vscode-certify-src-config-configpanel-ts-configpanel"></a>
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

<a id="vscode-certify-src-config-configwriter-ts-connectiontestresult"></a>
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

<a id="vscode-certify-src-config-configwriter-ts-readconfig"></a>
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

<a id="vscode-certify-src-config-configwriter-ts-testconnection"></a>
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

<a id="vscode-certify-src-config-configwriter-ts-writeconfig"></a>
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
| [`DashboardPanel`](#vscode-certify-src-dashboard-dashboardpanel-ts-dashboardpanel) | class | B+ | 88.3% | certified | 2026-04-23 |

<a id="vscode-certify-src-dashboard-dashboardpanel-ts-dashboardpanel"></a>
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
| [`CertifyDiagnostics`](#vscode-certify-src-diagnostics-certifydiagnostics-ts-certifydiagnostics) | class | B+ | 88.3% | certified | 2026-04-23 |

<a id="vscode-certify-src-diagnostics-certifydiagnostics-ts-certifydiagnostics"></a>
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
| [`CertificationTreeProvider`](#vscode-certify-src-treeview-certificationtreeprovider-ts-certificationtreeprovider) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`CertifyTreeItem`](#vscode-certify-src-treeview-certificationtreeprovider-ts-certifytreeitem) | class | B+ | 88.3% | certified | 2026-04-23 |

<a id="vscode-certify-src-treeview-certificationtreeprovider-ts-certificationtreeprovider"></a>
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

<a id="vscode-certify-src-treeview-certificationtreeprovider-ts-certifytreeitem"></a>
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
| [`collections`](#website-src-content-config-ts-collections) | function | B+ | 89.4% | certified | 2026-04-23 |

<a id="website-src-content-config-ts-collections"></a>
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
