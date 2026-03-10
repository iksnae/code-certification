# 🟢 Certify — Report Card

**Repository:** `iksnae/code-certification`  
**Commit:** `b7a7b83`  
**Generated:** 2026-03-10T12:46:56  
**[Browse full report →](site/index.html)**  

---

## Summary

| Metric | Value |
|--------|-------|
| **Overall Grade** | 🟢 **B+** |
| **Overall Score** | 89.0% |
| **Total Units** | 559 |
| **Passing** | 559 |
| **Failing** | 0 |
| **Pass Rate** | 100.0% |
| **Observations** | 0 |
| **Expired** | 0 |

## Grade Distribution

| Grade | Count | % | Bar |
|:-----:|------:|----:|-----|
| B+ | 558 | 99.8% | █████████████████████████████████████████████████ |
| B | 1 | 0.2% | █ |

## Dimension Averages

| Dimension | Score | Bar |
|-----------|------:|-----|
| architectural_fitness | 85.0% | ████████████████░░░░ |
| change_risk | 90.0% | █████████████████░░░ |
| correctness | 95.0% | ██████████████████░░ |
| maintainability | 93.5% | ██████████████████░░ |
| operational_quality | 85.0% | █████████████████░░░ |
| performance_appropriateness | 85.0% | ████████████████░░░░ |
| readability | 92.6% | ██████████████████░░ |
| security | 85.0% | ████████████████░░░░ |
| testability | 90.0% | █████████████████░░░ |

## 🤖 AI Insights

*Powered by `qwen/qwen3-coder-30b` — 559 units analyzed*

### Top Suggestions

- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps ×19
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly ×19
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection ×19
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#') ×19
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation ×19
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time ×19
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers ×17
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints ×17
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use ×16
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig ×16
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers ×16
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions ×16
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent ×16
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme ×15
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error ×15
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail ×15
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted ×15
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store ×15
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing ×15
- 💡 Make file permissions configurable via environment variables or config options ×15

*...and 625 more suggestions across individual units*

## By Language

### go — 🟢 B+ (89.1%)

- **Units:** 492
- **Score range:** 86.1% – 89.4%
- **Grades:** 491×B+, 1×B

### ts — 🟢 B+ (88.5%)

- **Units:** 67
- **Score range:** 88.3% – 89.4%
- **Grades:** 67×B+

## All Units

### `cmd/certify/` (36 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`certifyContext`](reports/cmd-certify-certify-cmd-go-certifycontext.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`certifyUnit`](reports/cmd-certify-certify-cmd-go-certifyunit.md) | method | B+ | 87.2% | certified | 2026-04-23 |
| [`collectRepoEvidence`](reports/cmd-certify-certify-cmd-go-collectrepoevidence.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`defaultConfigObj`](reports/cmd-certify-certify-cmd-go-defaultconfigobj.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`filterUnits`](reports/cmd-certify-certify-cmd-go-filterunits.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](reports/cmd-certify-certify-cmd-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`isLocalURL`](reports/cmd-certify-certify-cmd-go-islocalurl.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`loadCertifyContext`](reports/cmd-certify-certify-cmd-go-loadcertifycontext.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`loadQueue`](reports/cmd-certify-certify-cmd-go-loadqueue.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`printQueueStatus`](reports/cmd-certify-certify-cmd-go-printqueuestatus.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`printSummary`](reports/cmd-certify-certify-cmd-go-printsummary.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`processQueue`](reports/cmd-certify-certify-cmd-go-processqueue.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`runCertify`](reports/cmd-certify-certify-cmd-go-runcertify.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`saveReportArtifacts`](reports/cmd-certify-certify-cmd-go-savereportartifacts.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`setupAgent`](reports/cmd-certify-certify-cmd-go-setupagent.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`setupConservativeAgent`](reports/cmd-certify-certify-cmd-go-setupconservativeagent.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`setupExplicitAgent`](reports/cmd-certify-certify-cmd-go-setupexplicitagent.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`cli_test.go`](reports/cmd-certify-cli-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`init`](reports/cmd-certify-expire-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`generateConfig`](reports/cmd-certify-init-cmd-go-generateconfig.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](reports/cmd-certify-init-cmd-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`languagePolicy`](reports/cmd-certify-init-cmd-go-languagepolicy.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`main`](reports/cmd-certify-main-go-main.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](reports/cmd-certify-models-cmd-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`listFromProvider`](reports/cmd-certify-models-cmd-go-listfromprovider.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`runModels`](reports/cmd-certify-models-cmd-go-runmodels.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`detectCommit`](reports/cmd-certify-report-cmd-go-detectcommit.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`detectRepoName`](reports/cmd-certify-report-cmd-go-detectreponame.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](reports/cmd-certify-report-cmd-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`saveBadge`](reports/cmd-certify-report-cmd-go-savebadge.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`saveReportCard`](reports/cmd-certify-report-cmd-go-savereportcard.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](reports/cmd-certify-review-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](reports/cmd-certify-root-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](reports/cmd-certify-scan-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`tryScanSuggestions`](reports/cmd-certify-scan-go-tryscansuggestions.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`version.go`](reports/cmd-certify-version-go.md) | file | B+ | 89.4% | certified | 2026-04-23 |

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
| readability | 85.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The isLocalURL function has a security vulnerability due to incomplete URL matching and potential bypass via malformed URLs.
- 💡 Replace isLocalURL with a proper URL parsing and validation function using net/url package to correctly identify local endpoints
- 💡 Add comprehensive test cases for isLocalURL covering edge cases like encoded URLs, different protocols, and port numbers
- ⚠️ Insecure URL validation in isLocalURL function allows potential bypass of local provider detection
- ⚠️ Improper handling of encoded or malformed URLs may lead to incorrect local/remote provider logic
- 🔗 This function affects the agent configuration and provider selection logic, which can change token limits, review strategies, and overall certification behavior
- 🔗 Improper local URL detection could result in unintended use of remote providers for local models, leading to incorrect token usage or strategy application

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
- 🤖 The test file has functional tests but lacks proper error handling, uses unsafe os operations, and contains fragile assumptions about directory structure.
- 💡 Replace `os.Getwd()` with a check for its error and use a more robust way to capture the current directory before changing it.
- 💡 Explicitly declare or inject `initCmd` and `versionCmd` in tests to avoid relying on global state.
- 💡 Add assertions for command output or exit codes in `TestVersionCmd` to make the test meaningful.
- 💡 Add validation of file content (e.g., config.yml) in `TestInitCmd` to ensure it's correctly initialized.
- 💡 Use `t.Cleanup()` for setting up and cleaning up temporary directories to ensure proper cleanup even in case of test failure.
- ⚠️ Unvalidated `os.Getwd()` in `TestInitCmd` can lead to runtime panics or incorrect behavior if the working directory is not accessible.
- ⚠️ Assumption that `initCmd` is a global variable without explicit declaration or initialization can cause runtime panics or undefined behavior.
- 🔗 This test file tightly couples to global command variables (`initCmd`, `versionCmd`) and assumes their existence, increasing system fragility.
- 🔗 The tests do not isolate the CLI logic from the filesystem or git operations, making them brittle and dependent on system state.

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
- 🤖 The code correctly implements a certification expiration checker but has several risk factors including error handling gaps, lack of proper logging, and potential race conditions in concurrent access.
- 💡 Replace `_` with proper error handling for `os.Getwd()` to prevent silent failures when the working directory is inaccessible
- 💡 Add synchronization mechanism (e.g., file locking or database transactions) to prevent race conditions during concurrent record updates
- 💡 Validate that `rec.ExpiresAt` is a valid time before comparing it to `now` to avoid unexpected panics or incorrect behavior
- 💡 Use structured logging instead of printing warnings to stderr for better observability and debugging support
- ⚠️ Silent failure on `os.Getwd()` due to suppression of error with `_`
- ⚠️ Race condition when multiple processes attempt to save expired records concurrently
- 🔗 This unit tightly couples to the filesystem structure and assumes a specific directory layout (.certification), increasing system fragility
- 🔗 The command modifies records in-place without transactional safety, potentially causing partial updates or data corruption under concurrent access

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
- 🤖 Well-structured CLI initialization command with solid directory and file creation logic, but lacks proper error handling for Git operations and has a hardcoded dependency on `gh` CLI.
- 💡 Add explicit error checking and exit on failure for all Git commands in the PR creation path to prevent partial initialization
- 💡 Implement a fallback or warning mechanism when `gh` is not available during PR creation, rather than silently ignoring the failure
- 💡 Validate and sanitize `initPath` input to prevent path traversal or invalid directory creation
- 💡 Consider using a more robust language detection or policy generation approach that can handle unknown languages gracefully
- ⚠️ Git command execution failure in PR creation path is not handled properly, leading to potential incomplete or broken initialization
- ⚠️ Hardcoded reliance on the `gh` CLI for PR creation without fallback or error handling if it's not installed or accessible
- 🔗 This unit tightly couples with external tools like `gh` and Git, increasing failure propagation risk when these are unavailable or misconfigured
- 🔗 The PR creation logic introduces a side effect that modifies the local repository state and assumes user environment is properly configured

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
- 🤖 Well-structured CLI initialization command with solid directory and file creation logic, but lacks proper error handling for Git operations and has a hardcoded dependency on `gh` CLI.
- 💡 Add explicit error checking and exit on failure for all Git commands in the PR creation path to prevent partial initialization
- 💡 Implement a fallback or warning mechanism when `gh` is not available during PR creation, rather than silently ignoring the failure
- 💡 Validate and sanitize `initPath` input to prevent path traversal or invalid directory creation
- 💡 Consider using a more robust language detection or policy generation approach that can handle unknown languages gracefully
- ⚠️ Git command execution failure in PR creation path is not handled properly, leading to potential incomplete or broken initialization
- ⚠️ Hardcoded reliance on the `gh` CLI for PR creation without fallback or error handling if it's not installed or accessible
- 🔗 This unit tightly couples with external tools like `gh` and Git, increasing failure propagation risk when these are unavailable or misconfigured
- 🔗 The PR creation logic introduces a side effect that modifies the local repository state and assumes user environment is properly configured

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
- 🤖 Well-structured CLI initialization command with solid directory and file creation logic, but lacks proper error handling for Git operations and has a hardcoded dependency on `gh` CLI.
- 💡 Add explicit error checking and exit on failure for all Git commands in the PR creation path to prevent partial initialization
- 💡 Implement a fallback or warning mechanism when `gh` is not available during PR creation, rather than silently ignoring the failure
- 💡 Validate and sanitize `initPath` input to prevent path traversal or invalid directory creation
- 💡 Consider using a more robust language detection or policy generation approach that can handle unknown languages gracefully
- ⚠️ Git command execution failure in PR creation path is not handled properly, leading to potential incomplete or broken initialization
- ⚠️ Hardcoded reliance on the `gh` CLI for PR creation without fallback or error handling if it's not installed or accessible
- 🔗 This unit tightly couples with external tools like `gh` and Git, increasing failure propagation risk when these are unavailable or misconfigured
- 🔗 The PR creation logic introduces a side effect that modifies the local repository state and assumes user environment is properly configured

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
- 🤖 Minimal main function with proper error handling but lacks comprehensive CLI setup and testing coverage
- 💡 Add structured logging capability to capture error context and application state for debugging purposes
- 💡 Implement signal handling to allow graceful shutdown of the CLI process when receiving termination signals
- ⚠️ Missing error context - errors printed to stderr lack structured logging or additional context that would help debugging
- ⚠️ No signal handling for graceful shutdown - application doesn't handle SIGTERM/SIGINT for clean termination
- 🔗 This unit is a thin wrapper that couples tightly to the external rootCmd implementation, making it difficult to test in isolation
- 🔗 The main function's simplicity creates a single point of failure for CLI execution without any error recovery or fallback mechanisms

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
- 🤖 The code is functionally correct but has several security and error handling gaps, including potential exposure of API keys in logs, lack of input validation, and missing error recovery for concurrent provider checks.
- 💡 Sanitize or mask API key values before printing to stderr in case of missing keys.
- 💡 Replace sequential provider loop with concurrent checks or at least attempt all providers before failing, to improve resilience.
- 💡 Add input validation for `baseURL` to prevent malformed URLs from causing runtime errors or injection.
- 💡 Use structured logging instead of `fmt.Fprintf` to stderr for better observability and consistency.
- 💡 Consider adding retry logic or exponential backoff when calling `agent.ListModels` to handle transient network issues.
- ⚠️ Potential exposure of API keys in logs or error messages due to direct use of `os.Getenv` without sanitization.
- ⚠️ Sequential provider checking with early return prevents fallback to other providers if one fails.
- 🔗 This unit tightly couples with the `agent` package's `DetectProviders` and `ListModels` functions, increasing system fragility if those change.
- 🔗 The command uses global variables (`modelsProviderURL`, `modelsAPIKeyEnv`) which makes it harder to unit test and increases coupling.

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
- 🤖 The code is functionally correct but has several security and error handling gaps, including potential exposure of API keys in logs, lack of input validation, and missing error recovery for concurrent provider checks.
- 💡 Sanitize or mask API key values before printing to stderr in case of missing keys.
- 💡 Replace sequential provider loop with concurrent checks or at least attempt all providers before failing, to improve resilience.
- 💡 Add input validation for `baseURL` to prevent malformed URLs from causing runtime errors or injection.
- 💡 Use structured logging instead of `fmt.Fprintf` to stderr for better observability and consistency.
- 💡 Consider adding retry logic or exponential backoff when calling `agent.ListModels` to handle transient network issues.
- ⚠️ Potential exposure of API keys in logs or error messages due to direct use of `os.Getenv` without sanitization.
- ⚠️ Sequential provider checking with early return prevents fallback to other providers if one fails.
- 🔗 This unit tightly couples with the `agent` package's `DetectProviders` and `ListModels` functions, increasing system fragility if those change.
- 🔗 The command uses global variables (`modelsProviderURL`, `modelsAPIKeyEnv`) which makes it harder to unit test and increases coupling.

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
- 🤖 The code is functionally correct but has several security and error handling gaps, including potential exposure of API keys in logs, lack of input validation, and missing error recovery for concurrent provider checks.
- 💡 Sanitize or mask API key values before printing to stderr in case of missing keys.
- 💡 Replace sequential provider loop with concurrent checks or at least attempt all providers before failing, to improve resilience.
- 💡 Add input validation for `baseURL` to prevent malformed URLs from causing runtime errors or injection.
- 💡 Use structured logging instead of `fmt.Fprintf` to stderr for better observability and consistency.
- 💡 Consider adding retry logic or exponential backoff when calling `agent.ListModels` to handle transient network issues.
- ⚠️ Potential exposure of API keys in logs or error messages due to direct use of `os.Getenv` without sanitization.
- ⚠️ Sequential provider checking with early return prevents fallback to other providers if one fails.
- 🔗 This unit tightly couples with the `agent` package's `DetectProviders` and `ListModels` functions, increasing system fragility if those change.
- 🔗 The command uses global variables (`modelsProviderURL`, `modelsAPIKeyEnv`) which makes it harder to unit test and increases coupling.

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
- 🤖 The code is functionally correct but has several critical issues including unhandled errors, inconsistent error handling, and potential security vulnerabilities from shell command execution.
- 💡 Replace os.Getwd() with explicit error handling instead of suppressing errors
- 💡 Add proper logging or error reporting for failed JSON formatting in saveBadge
- 💡 Sanitize and validate inputs before executing git commands to prevent command injection
- 💡 Implement proper file system checks (e.g., disk space, permissions) before writing files in saveReportCard and saveBadge
- 💡 Use context with timeout for git command execution to prevent hanging processes
- ⚠️ Command injection vulnerability in detectRepoName and detectCommit due to direct use of exec.Command
- ⚠️ Silent failure in saveBadge when JSON formatting fails
- 🔗 The report command tightly couples to the git repository structure and assumes local environment consistency
- 🔗 Unreliable error handling in saveBadge can cause silent data loss or invalid badge generation

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
- 🤖 The code is functionally correct but has several critical issues including unhandled errors, inconsistent error handling, and potential security vulnerabilities from shell command execution.
- 💡 Replace os.Getwd() with explicit error handling instead of suppressing errors
- 💡 Add proper logging or error reporting for failed JSON formatting in saveBadge
- 💡 Sanitize and validate inputs before executing git commands to prevent command injection
- 💡 Implement proper file system checks (e.g., disk space, permissions) before writing files in saveReportCard and saveBadge
- 💡 Use context with timeout for git command execution to prevent hanging processes
- ⚠️ Command injection vulnerability in detectRepoName and detectCommit due to direct use of exec.Command
- ⚠️ Silent failure in saveBadge when JSON formatting fails
- 🔗 The report command tightly couples to the git repository structure and assumes local environment consistency
- 🔗 Unreliable error handling in saveBadge can cause silent data loss or invalid badge generation

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
- 🤖 The code is functionally correct but has several critical issues including unhandled errors, inconsistent error handling, and potential security vulnerabilities from shell command execution.
- 💡 Replace os.Getwd() with explicit error handling instead of suppressing errors
- 💡 Add proper logging or error reporting for failed JSON formatting in saveBadge
- 💡 Sanitize and validate inputs before executing git commands to prevent command injection
- 💡 Implement proper file system checks (e.g., disk space, permissions) before writing files in saveReportCard and saveBadge
- 💡 Use context with timeout for git command execution to prevent hanging processes
- ⚠️ Command injection vulnerability in detectRepoName and detectCommit due to direct use of exec.Command
- ⚠️ Silent failure in saveBadge when JSON formatting fails
- 🔗 The report command tightly couples to the git repository structure and assumes local environment consistency
- 🔗 Unreliable error handling in saveBadge can cause silent data loss or invalid badge generation

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
- 🤖 The code is functionally correct but has several critical issues including unhandled errors, inconsistent error handling, and potential security vulnerabilities from shell command execution.
- 💡 Replace os.Getwd() with explicit error handling instead of suppressing errors
- 💡 Add proper logging or error reporting for failed JSON formatting in saveBadge
- 💡 Sanitize and validate inputs before executing git commands to prevent command injection
- 💡 Implement proper file system checks (e.g., disk space, permissions) before writing files in saveReportCard and saveBadge
- 💡 Use context with timeout for git command execution to prevent hanging processes
- ⚠️ Command injection vulnerability in detectRepoName and detectCommit due to direct use of exec.Command
- ⚠️ Silent failure in saveBadge when JSON formatting fails
- 🔗 The report command tightly couples to the git repository structure and assumes local environment consistency
- 🔗 Unreliable error handling in saveBadge can cause silent data loss or invalid badge generation

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
- 🤖 The code is functionally correct but has several critical issues including unhandled errors, inconsistent error handling, and potential security vulnerabilities from shell command execution.
- 💡 Replace os.Getwd() with explicit error handling instead of suppressing errors
- 💡 Add proper logging or error reporting for failed JSON formatting in saveBadge
- 💡 Sanitize and validate inputs before executing git commands to prevent command injection
- 💡 Implement proper file system checks (e.g., disk space, permissions) before writing files in saveReportCard and saveBadge
- 💡 Use context with timeout for git command execution to prevent hanging processes
- ⚠️ Command injection vulnerability in detectRepoName and detectCommit due to direct use of exec.Command
- ⚠️ Silent failure in saveBadge when JSON formatting fails
- 🔗 The report command tightly couples to the git repository structure and assumes local environment consistency
- 🔗 Unreliable error handling in saveBadge can cause silent data loss or invalid badge generation

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including improper error recovery and missing input validation.
- 💡 Add proper error logging and handling for config loading failure on line 27, e.g., log the error and continue or exit gracefully
- 💡 Validate and sanitize `reviewPath` input before constructing paths to prevent path traversal or injection vulnerabilities
- 💡 Ensure `defaultConfigObj()` is either defined in this file or imported to avoid runtime panics
- 💡 Add error handling for `os.Getwd()` on line 26 to prevent silent failures in case of permission issues or other errors
- ⚠️ Runtime panic due to undefined `defaultConfigObj()` function
- ⚠️ Improper error handling leading to silent config loading failures
- 🔗 This unit introduces a dependency on an undefined global config object, increasing system fragility
- 🔗 The command does not validate or sanitize input paths, creating potential path traversal or injection risks

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
- 🤖 Code is functionally correct but has missing command definitions and potential runtime errors due to undefined variables.
- 💡 Define or import the missing command variables (versionCmd, initCmd, scanCmd, certifyCmd, reportCmd, modelsCmd) with proper command structure including Run functions
- 💡 Add proper error handling for command registration and execution in the init() function
- 💡 Include a main() function that calls rootCmd.Execute() to make the CLI functional
- 💡 Add proper command descriptions and flag definitions for each subcommand
- ⚠️ Compilation failure due to undefined variables (versionCmd, initCmd, scanCmd, certifyCmd, reportCmd, modelsCmd)
- ⚠️ Runtime panic if any command is executed since command functions are not defined
- 🔗 This unit creates a dependency on external cobra library but doesn't properly integrate with the command execution flow
- 🔗 The unit introduces tight coupling between root command and undefined subcommands, making system behavior unpredictable

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
- 🤖 The function `tryScanSuggestions` has a clean structure but contains a critical logic flaw in how it handles AI provider selection and lacks proper error handling for the AI suggestion flow.
- 💡 Validate that `dp.BaseURL` and `dp.APIKey` are not empty before constructing the provider to prevent runtime panics or invalid HTTP requests
- 💡 Add explicit error handling and logging for `agent.SuggestForRepo` to ensure failures are surfaced rather than silently ignored
- 💡 Extract the AI suggestion logic into a separate function with configurable timeout and retry behavior to decouple from hardcoded values
- 💡 Refactor output operations (`fmt.Printf`, `fmt.Fprintf`) to accept an io.Writer interface for better testability and flexibility
- ⚠️ Potential panic or invalid request due to missing validation of `dp.BaseURL` and `dp.APIKey` before constructing a provider
- ⚠️ Silent failure in AI suggestion logic due to lack of error propagation when `agent.SuggestForRepo` fails
- 🔗 The function introduces a soft dependency on external AI services that can cause silent failures or degraded UX without proper error reporting
- 🔗 Tight coupling to `os.Stderr` and `os.Stdout` makes this function non-testable and harder to integrate into other systems or UIs

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
- 🤖 The function `tryScanSuggestions` has a clean structure but contains a critical logic flaw in how it handles AI provider selection and lacks proper error handling for the AI suggestion flow.
- 💡 Validate that `dp.BaseURL` and `dp.APIKey` are not empty before constructing the provider to prevent runtime panics or invalid HTTP requests
- 💡 Add explicit error handling and logging for `agent.SuggestForRepo` to ensure failures are surfaced rather than silently ignored
- 💡 Extract the AI suggestion logic into a separate function with configurable timeout and retry behavior to decouple from hardcoded values
- 💡 Refactor output operations (`fmt.Printf`, `fmt.Fprintf`) to accept an io.Writer interface for better testability and flexibility
- ⚠️ Potential panic or invalid request due to missing validation of `dp.BaseURL` and `dp.APIKey` before constructing a provider
- ⚠️ Silent failure in AI suggestion logic due to lack of error propagation when `agent.SuggestForRepo` fails
- 🔗 The function introduces a soft dependency on external AI services that can cause silent failures or degraded UX without proper error reporting
- 🔗 Tight coupling to `os.Stderr` and `os.Stdout` makes this function non-testable and harder to integrate into other systems or UIs

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
- 🤖 Simple version command with build-time variable injection, but lacks proper error handling and testing coverage for the version output format
- 💡 Add input validation and sanitization for Version, Commit, and Date variables before printing to prevent formatting or injection issues
- 💡 Implement unit tests for the version command that verify output format and handle edge cases like empty or malformed version strings
- 💡 Consider using a structured logging approach instead of direct printf to better handle output formatting and potential security issues
- 💡 Remove or update the static Release: v0.1.3 comment to avoid confusion with dynamic build-time variables
- ⚠️ Potential command injection or formatting issues if Version, Commit, or Date contain special characters that break printf formatting
- ⚠️ Security risk from exposing raw build metadata without sanitization or validation of fields that could contain malicious content
- 🔗 This unit introduces a public API endpoint (version command) that exposes internal build metadata to users, creating potential information disclosure
- 🔗 The hardcoded printf format creates tight coupling between the version string format and the command implementation, making future format changes more difficult

</details>

### `extensions/` (18 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`agent-chain.ts`](reports/extensions-agent-chain-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`agent-team.ts`](reports/extensions-agent-team-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`cross-agent.ts`](reports/extensions-cross-agent-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`damage-control.ts`](reports/extensions-damage-control-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`minimal.ts`](reports/extensions-minimal-ts.md) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`pi-pi.ts`](reports/extensions-pi-pi-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`pure-focus.ts`](reports/extensions-pure-focus-ts.md) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`purpose-gate.ts`](reports/extensions-purpose-gate-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`session-replay.ts`](reports/extensions-session-replay-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`subagent-widget.ts`](reports/extensions-subagent-widget-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`system-select.ts`](reports/extensions-system-select-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`theme-cycler.ts`](reports/extensions-theme-cycler-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`THEME_MAP`](reports/extensions-thememap-ts-theme-map.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`applyExtensionDefaults`](reports/extensions-thememap-ts-applyextensiondefaults.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`applyExtensionTheme`](reports/extensions-thememap-ts-applyextensiontheme.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`tilldone.ts`](reports/extensions-tilldone-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`tool-counter-widget.ts`](reports/extensions-tool-counter-widget-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`tool-counter.ts`](reports/extensions-tool-counter-ts.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 Code has functional correctness but contains critical security and resource management flaws including unsafe subprocess spawning, potential path traversal, and missing error handling in key areas.
- 💡 Validate and sanitize all subprocess arguments in `runAgent` using a whitelist approach or proper escaping
- 💡 Implement strict input validation for file paths in `scanAgentDirs` and `loadChains` to prevent directory traversal
- 💡 Add proper error handling in `runAgent` to capture and report stderr content for debugging
- 💡 Use atomic file operations or locking when managing session files to prevent race conditions
- 💡 Improve YAML parsing logic to handle multi-line prompts and nested structures properly
- 💡 Add proper type checking for all user inputs and tool parameters to prevent runtime errors
- ⚠️ Command injection vulnerability in subprocess execution via unvalidated prompt parameters
- ⚠️ Path traversal and directory traversal issues in agent file scanning and session management
- 🔗 The extension creates a security boundary violation by spawning subprocesses with user-controlled parameters
- 🔗 High coupling to the Pi agent system through direct subprocess calls and session file manipulation
- 🔗 Failure propagation risk in chain execution where one failed agent stops the entire pipeline without graceful handling

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
- 🤖 The code implements a Pi agent team dispatcher with UI rendering but has critical resource leak risks, unhandled edge cases, and lacks proper error recovery.
- 💡 Ensure all spawned process streams (stdout/stderr) are properly closed in error handlers to prevent resource leaks
- 💡 Validate that theme color names like 'accent' are supported before calling `theme.fg()` to avoid runtime errors
- ⚠️ Resource leak in spawned process stdout/stderr handling
- ⚠️ Potential runtime errors from invalid theme color names
- 🔗 The extension introduces a new tool called 'dispatch_agent' that can spawn arbitrary subprocesses, creating potential security vulnerabilities if the task input is not properly sanitized
- 🔗 The extension creates a new UI widget that depends on TUI state management, which could cause cascading failures if the TUI context is not properly initialized

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
- 🤖 Code has functional correctness but suffers from poor error handling, security risks via unsafe file operations, and UI rendering issues.
- 💡 Replace `readdirSync` and `readFileSync` with async versions and add proper error handling around each operation to prevent crashes
- 💡 Implement input validation and sanitization for all discovered content (command names, agent names, descriptions) before registering or rendering
- 💡 Add timeout and retry logic for UI notifications to prevent race conditions
- 💡 Use `fs.promises` instead of sync methods and handle errors explicitly to avoid blocking the event loop
- 💡 Implement proper frontmatter parsing with regex escaping and safer field extraction to prevent DoS or injection
- 💡 Add fallbacks for ANSI color codes when TTY is not supported
- ⚠️ Uncaught filesystem errors in directory scanning can crash the process
- ⚠️ Unsafe argument expansion in `expandArgs` may lead to command injection or unexpected behavior
- 🔗 This extension introduces tight coupling with filesystem structure and assumes specific directory layouts
- 🔗 The UI notification system is tightly coupled to terminal rendering, potentially breaking in non-TTY environments

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
- 🤖 The damage-control.ts extension has functional rule-based access control but contains critical security flaws in path matching, regex handling, and unsafe bash command analysis.
- 💡 Replace the path matching logic with a proper glob library like `micromatch` or implement a more robust regex that correctly handles directory vs file matching
- 💡 Implement comprehensive input validation and sanitization for all paths and bash commands before any pattern matching occurs
- 💡 Add proper escaping of user-provided regex patterns in bash rule matching to prevent injection attacks
- 💡 Implement strict path validation that ensures resolved paths remain within the project directory scope
- 💡 Add unit tests for edge cases including path traversal, glob patterns, and bash command obfuscation
- ⚠️ Path traversal vulnerability due to insecure path resolution and lack of validation
- ⚠️ Bash command bypass through simple string matching and regex injection
- 🔗 This extension introduces a critical security boundary that can be easily circumvented, making it ineffective as a damage control mechanism
- 🔗 The extension's rule loading and validation logic creates a potential failure point that can crash or misbehave on malformed YAML files

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
- 🤖 The extension implements a minimal UI footer with model name and context usage bar, but has critical correctness issues in rendering logic and lacks proper error handling.
- 💡 Replace `visibleWidth` with a proper width calculation that accounts for ANSI color codes in styled strings before computing padding
- 💡 Add validation and sanitization of `pct` value to prevent NaN or invalid bar rendering when `getContextUsage()` returns malformed data
- 💡 Implement proper lifecycle methods (`dispose`, `invalidate`) to allow for clean resource management and reactivity in the TUI system
- 💡 Consider returning an array of strings from `render` that can support multi-line footers if needed in future extensions
- ⚠️ Incorrect string width calculation leading to misaligned or truncated UI elements
- ⚠️ Potential runtime errors from invalid percentage values in context usage bar
- 🔗 This extension tightly couples to the `@mariozechner/pi-tui` package's `visibleWidth` and `truncateToWidth` functions, creating brittle dependency on TUI rendering behavior
- 🔗 The footer's render method does not follow standard TUI component conventions, which could break expectations in larger UI systems or cause unexpected layout behavior

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
- 🤖 The Pi Pi extension is a complex multi-agent system with good structure but critical resource leak and error handling issues.
- 💡 Add explicit process cleanup in queryExpert error handlers using proc.kill()
- 💡 Replace the regex in parseAgentFile with safer parsing logic or add input validation to prevent catastrophic backtracking
- 💡 Fix ANSI color code handling in renderCard by ensuring proper reset sequences are applied
- 💡 Implement proper input validation for expert names to prevent command injection
- 💡 Add timeout handling to queryExpert to prevent indefinite hanging subprocesses
- 💡 Fix the JSON parsing logic in queryExpert to properly handle incomplete buffers
- ⚠️ Resource leak in child process management - processes are not killed on error
- ⚠️ Potential catastrophic backtracking in regex parsing of agent files
- 🔗 This extension creates subprocesses that can accumulate and exhaust system resources if not properly cleaned up
- 🔗 The TUI rendering uses hardcoded ANSI codes that can corrupt terminal output in complex environments

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
- 🤖 The code implements a UI extension to hide footer and status line, but has critical correctness and maintainability issues.
- 💡 Replace `import.meta.url` with a proper theme path or remove the call to `applyExtensionDefaults` since it's not used for this extension
- 💡 Add proper error handling around `ctx.ui.setFooter` and validate that the context UI object exists before calling methods on it
- 💡 Investigate how the actual TUI system handles footer rendering and implement a proper override mechanism instead of returning an empty render function
- 💡 Add type assertion or validation to ensure `ctx.ui` has the expected methods before calling them
- ⚠️ Runtime error when `applyExtensionDefaults` is called with `import.meta.url` as a path
- ⚠️ Extension does not actually hide footer UI due to incorrect implementation pattern
- 🔗 This extension introduces a brittle hook that assumes specific UI structure and behavior
- 🔗 The extension's reliance on `import.meta.url` makes it non-portable and breaks in environments without proper module resolution

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
- 🤖 The code implements a session-level purpose gate with UI integration but has critical runtime and type safety issues.
- 💡 Replace `void askForPurpose(ctx)` with proper async handling to ensure the input is completed before proceeding
- 💡 Add null checks or default values for `purpose` before rendering to prevent runtime crashes
- 💡 Validate that `ctx.ui.input` exists and is callable before using it to prevent undefined method errors
- 💡 Ensure that `invalidate()` in the widget actually triggers re-rendering or remove it if not needed
- 💡 Add a check to ensure that `purpose` is not only set but also valid before modifying system prompts
- ⚠️ Uncaught runtime error when accessing `purpose!` in render function if purpose is never set
- ⚠️ Potential crash or undefined behavior due to improper async handling of input prompt
- 🔗 This extension introduces a hard dependency on UI input/output methods that may not exist in all environments
- 🔗 The widget rendering logic is tightly coupled to specific TUI formatting and lacks fallbacks for other output modes

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
- 🤖 
    "summary": "The SessionReplayUI component has functional rendering and navigation but suffers from poor input handling, hardcoded UI assumptions, and missing error boundaries that could lead to s...

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
- 🤖 The subagent widget implementation is functionally correct but has critical resource leak risks, missing error handling, and unsafe process management.
- 💡 Replace all `any` types with proper TypeScript interfaces and use strict type checking
- 💡 Add proper error handling around fs.mkdirSync() with try/catch and handle permission errors
- 💡 Implement proper cleanup of child processes in all exit paths (e.g., on /subrm, /subclear, session_start)
- 💡 Fix the process.kill() call to check if process exists before killing it
- 💡 Add input sanitization for session file paths to prevent path traversal attacks
- 💡 Fix the processLine() function to properly handle stderr data in addition to stdout
- 💡 Use a more robust session file naming strategy to prevent race conditions
- 💡 Implement proper widget invalidation logic with container state reset
- ⚠️ Uncaught exceptions in JSON parsing can silently drop data or crash the process
- ⚠️ Process handle leak when subagents are removed without proper cleanup of child processes
- 🔗 This unit introduces a critical resource leak in child process management that can cause system instability
- 🔗 The widget rendering logic has a deterministic UI corruption bug when widgets are updated

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
- 🤖 The code implements a system prompt selector with agent scanning and UI integration, but has critical security and correctness issues including path traversal, unsafe file reading, and missing error handling.
- 💡 Add proper error handling around filesystem operations with meaningful logging
- 💡 Implement strict path validation to prevent directory traversal attacks by ensuring all scanned paths are within allowed directories
- 💡 Fix the displayName function to handle empty strings properly by checking length before accessing character methods
- 💡 Add input validation and sanitization for all user-provided paths in the directory scanning logic
- 💡 Fix parseFrontmatter to properly handle multi-line YAML fields and edge cases in frontmatter parsing
- 💡 Implement proper resource cleanup for file operations to prevent memory leaks
- ⚠️ Path traversal vulnerability in directory scanning allowing arbitrary file reads
- ⚠️ Uncaught filesystem errors causing silent failures in agent loading
- 🔗 This extension directly modifies system prompts and tool sets, creating tight coupling with the core Pi agent behavior
- 🔗 The extension introduces a security risk by allowing arbitrary file reading from project and home directories

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
- 🤖 The theme cycler has functional keyboard shortcuts and commands but suffers from critical race conditions, improper state management, and missing error handling for UI operations.
- 💡 Use a local context reference in each handler instead of relying on a shared `currentCtx` variable to prevent race conditions
- 💡 Add proper error handling around `ctx.ui.setWidget()` and `ctx.ui.select()` calls to ensure that UI state remains consistent even if operations fail
- 💡 Ensure `swatchTimer` is always cleared in all code paths after setting or clearing widgets, including error cases
- 💡 Validate that `ctx.ui.theme.name` exists in the theme list before attempting to find its index, and handle missing themes gracefully
- 💡 Remove unused import `truncateToWidth` to improve clarity
- ⚠️ Race condition with `currentCtx` being overwritten by concurrent shortcut/command handlers
- ⚠️ Potential memory leak from unhandled clearTimeout in widget update failure cases
- 🔗 This extension introduces a shared mutable state (`currentCtx`) that couples it tightly to session lifecycle and can cause stale context usage across concurrent operations
- 🔗 The extension's widget rendering relies on an undocumented TUI API pattern (`render` function returning array), which increases fragility and makes the extension harder to maintain or extend

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
- 🤖 The code is functionally correct but has a critical race condition in theme application and lacks proper error handling for UI operations.
- 💡 Replace the heuristic primary extension detection with a more robust mechanism, such as an environment variable or centralized state tracking
- 💡 Add explicit error logging when theme application fails and fallback occurs, to aid debugging in production
- 💡 Validate that `fileUrl` and `ctx` are valid before proceeding with theme application
- 💡 Make the fallback theme configurable or validate that it exists at runtime
- 💡 Return `false` from `applyExtensionTheme` when skipping theme application for a secondary extension, to clearly indicate that no action was taken
- ⚠️ Race condition in theme application due to reliance on process.argv ordering
- ⚠️ Silent UI failure with no error reporting when theme setting fails
- 🔗 This unit introduces a brittle dependency on process.argv order, which can cause incorrect theme application in complex Pi setups
- 🔗 The function's return behavior is inconsistent and may mislead callers into thinking theme application succeeded when it didn't

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
- 🤖 The code is functionally correct but has a critical race condition in theme application and lacks proper error handling for UI operations.
- 💡 Replace the heuristic primary extension detection with a more robust mechanism, such as an environment variable or centralized state tracking
- 💡 Add explicit error logging when theme application fails and fallback occurs, to aid debugging in production
- 💡 Validate that `fileUrl` and `ctx` are valid before proceeding with theme application
- 💡 Make the fallback theme configurable or validate that it exists at runtime
- 💡 Return `false` from `applyExtensionTheme` when skipping theme application for a secondary extension, to clearly indicate that no action was taken
- ⚠️ Race condition in theme application due to reliance on process.argv ordering
- ⚠️ Silent UI failure with no error reporting when theme setting fails
- 🔗 This unit introduces a brittle dependency on process.argv order, which can cause incorrect theme application in complex Pi setups
- 🔗 The function's return behavior is inconsistent and may mislead callers into thinking theme application succeeded when it didn't

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
- 🤖 The code is functionally correct but has a critical race condition in theme application and lacks proper error handling for UI operations.
- 💡 Replace the heuristic primary extension detection with a more robust mechanism, such as an environment variable or centralized state tracking
- 💡 Add explicit error logging when theme application fails and fallback occurs, to aid debugging in production
- 💡 Validate that `fileUrl` and `ctx` are valid before proceeding with theme application
- 💡 Make the fallback theme configurable or validate that it exists at runtime
- 💡 Return `false` from `applyExtensionTheme` when skipping theme application for a secondary extension, to clearly indicate that no action was taken
- ⚠️ Race condition in theme application due to reliance on process.argv ordering
- ⚠️ Silent UI failure with no error reporting when theme setting fails
- 🔗 This unit introduces a brittle dependency on process.argv order, which can cause incorrect theme application in complex Pi setups
- 🔗 The function's return behavior is inconsistent and may mislead callers into thinking theme application succeeded when it didn't

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
- 🤖 Well-structured task management extension with UI rendering, but has critical logic gaps in tool execution and session state handling.
- 💡 Complete the tilldone execute function with proper handling of all action cases (list, add, remove, update, clear)
- 💡 Add type guards and error handling for session data in reconstructState to prevent runtime crashes
- 💡 Refactor refreshWidget to reuse DynamicBorder instances instead of recreating them on every render call
- 💡 Simplify repeated filtering logic in UI rendering functions by caching filtered results
- ⚠️ Unimplemented tool actions (list, add, remove, update, clear) in tilldone execute function
- ⚠️ Potential runtime errors from malformed session data in reconstructState
- 🔗 The tilldone tool is a required gate for other tools, so incomplete implementation blocks agent usage
- 🔗 Session state reconstruction can silently fail or crash if session data structure changes

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
- 🤖 The code implements a tool counter widget with basic functionality but has critical correctness and maintainability issues.
- 💡 Add proper error handling around `applyExtensionDefaults` and UI rendering to prevent silent failures
- 💡 Implement a mechanism to reset `counts` and `total` on session start to ensure accurate counting
- 💡 Use a more robust color assignment strategy that persists across sessions or resets appropriately
- 💡 Validate and sanitize ANSI escape code usage to prevent terminal injection vulnerabilities
- 💡 Add unit tests for the widget rendering logic and session handling
- ⚠️ Race condition in concurrent access to `counts` and `toolColors` maps
- ⚠️ Potential security vulnerability from direct use of ANSI escape codes without validation
- 🔗 The widget creates a persistent UI element that may not be properly cleaned up on session termination
- 🔗 The hardcoded color palette limits extensibility and theming capabilities

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
- 🤖 The tool counter extension is functionally correct but has several maintainability and correctness issues including potential race conditions, improper disposal of event listeners, and lack of error handling.
- 💡 Use a Map or proper locking mechanism to synchronize access to the counts object
- 💡 Call `unsub()` in a cleanup function or use a proper disposal pattern to prevent memory leaks
- 💡 Add error handling around session branch iteration and context usage access
- 💡 Implement proper type guards when accessing `entry.message` to ensure it's an AssistantMessage
- 💡 Extract token counting logic into a separate function for better testability and maintainability
- ⚠️ Race condition in counts object access
- ⚠️ Memory leak due to unmanaged event listener subscription
- 🔗 This extension tightly couples to session state and UI rendering, making it sensitive to changes in the underlying session manager or TUI components
- 🔗 The extension's footer rendering logic could cause UI flickering or rendering issues if not properly disposed of

</details>

### `internal/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`integration_test.go`](reports/internal-integration-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The integration test suite has good end-to-end coverage but suffers from poor error handling, hardcoded paths, and missing assertions that could mask failures.
- 💡 Replace `src, _ := os.ReadFile(...)` with explicit error handling to ensure test integrity when files are missing or unreadable
- 💡 Use `t.Cleanup` or explicit temp directory management to ensure test data isolation and prevent resource leaks
- 💡 Make paths configurable or use `t.TempDir()` consistently for all temporary directories to improve portability
- 💡 Add assertions on specific fields in `TestE2E_AdvisoryMode` to validate that the certification record has expected failure details
- 💡 Validate that `TestE2E_TSSimpleRepo` has a stable and documented minimum number of discovered units to avoid flaky tests
- ⚠️ Silent failure due to ignored file read errors in evidence collection
- ⚠️ Brittle test dependencies on hardcoded relative paths and directory structure
- 🔗 The tests rely on hardcoded paths and external test data, increasing coupling to local environment
- 🔗 Failure in one test (e.g., due to missing file) may mask issues in other tests or lead to false positives

</details>

### `internal/agent/` (138 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`attribution_test.go`](reports/internal-agent-attribution-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`DetectAPIKey`](reports/internal-agent-autodetect-go-detectapikey.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatProviderSummary`](reports/internal-agent-autodetect-go-formatprovidersummary.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`HasAnyProvider`](reports/internal-agent-autodetect-go-hasanyprovider.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewConservativeCoordinator`](reports/internal-agent-autodetect-go-newconservativecoordinator.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`init`](reports/internal-agent-autodetect-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`autodetect_test.go`](reports/internal-agent-autodetect-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Chat`](reports/internal-agent-circuit-go-chat.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`CircuitBreaker`](reports/internal-agent-circuit-go-circuitbreaker.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`IsOpen`](reports/internal-agent-circuit-go-isopen.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Name`](reports/internal-agent-circuit-go-name.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewCircuitBreaker`](reports/internal-agent-circuit-go-newcircuitbreaker.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`AdaptiveMessages`](reports/internal-agent-fallback-go-adaptivemessages.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Chat`](reports/internal-agent-fallback-go-chat.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`FallbackProvider`](reports/internal-agent-fallback-go-fallbackprovider.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ModelChain`](reports/internal-agent-fallback-go-modelchain.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Name`](reports/internal-agent-fallback-go-name.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewFallbackProvider`](reports/internal-agent-fallback-go-newfallbackprovider.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewModelChain`](reports/internal-agent-fallback-go-newmodelchain.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`modelPinnedProvider`](reports/internal-agent-fallback-go-modelpinnedprovider.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`fallback_test.go`](reports/internal-agent-fallback-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`ListModels`](reports/internal-agent-models-go-listmodels.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ModelInfo`](reports/internal-agent-models-go-modelinfo.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`listOllamaModels`](reports/internal-agent-models-go-listollamamodels.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`listOpenAIModels`](reports/internal-agent-models-go-listopenaimodels.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ollamaModel`](reports/internal-agent-models-go-ollamamodel.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ollamaTagsResponse`](reports/internal-agent-models-go-ollamatagsresponse.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`openAIModel`](reports/internal-agent-models-go-openaimodel.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`openAIModelsResponse`](reports/internal-agent-models-go-openaimodelsresponse.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`models_test.go`](reports/internal-agent-models-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`APIError`](reports/internal-agent-openrouter-go-apierror.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Chat`](reports/internal-agent-openrouter-go-chat.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`Error`](reports/internal-agent-openrouter-go-error.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Name`](reports/internal-agent-openrouter-go-name.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewLocalProvider`](reports/internal-agent-openrouter-go-newlocalprovider.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewOpenRouterProvider`](reports/internal-agent-openrouter-go-newopenrouterprovider.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`OpenRouterProvider`](reports/internal-agent-openrouter-go-openrouterprovider.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`doRequest`](reports/internal-agent-openrouter-go-dorequest.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`isAPIError`](reports/internal-agent-openrouter-go-isapierror.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`isAuthError`](reports/internal-agent-openrouter-go-isautherror.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`isBudgetError`](reports/internal-agent-openrouter-go-isbudgeterror.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`isRetryable`](reports/internal-agent-openrouter-go-isretryable.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`openrouter_test.go`](reports/internal-agent-openrouter-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Coordinator`](reports/internal-agent-pipeline-go-coordinator.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`CoordinatorConfig`](reports/internal-agent-pipeline-go-coordinatorconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`IsLocal`](reports/internal-agent-pipeline-go-islocal.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewCoordinator`](reports/internal-agent-pipeline-go-newcoordinator.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewPipeline`](reports/internal-agent-pipeline-go-newpipeline.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`Pipeline`](reports/internal-agent-pipeline-go-pipeline.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`PipelineConfig`](reports/internal-agent-pipeline-go-pipelineconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ReviewUnit`](reports/internal-agent-pipeline-go-reviewunit.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`Run`](reports/internal-agent-pipeline-go-run.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`Stats`](reports/internal-agent-pipeline-go-stats.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Strategy`](reports/internal-agent-pipeline-go-strategy.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`toResult`](reports/internal-agent-pipeline-go-toresult.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Get`](reports/internal-agent-prompts-go-get.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`LoadPrompt`](reports/internal-agent-prompts-go-loadprompt.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewPromptRegistry`](reports/internal-agent-prompts-go-newpromptregistry.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`PromptRegistry`](reports/internal-agent-prompts-go-promptregistry.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`PromptTemplate`](reports/internal-agent-prompts-go-prompttemplate.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Render`](reports/internal-agent-prompts-go-render.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Version`](reports/internal-agent-prompts-go-version.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`prompts_test.go`](reports/internal-agent-prompts-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Provider`](reports/internal-agent-provider-go-provider.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`provider_multi_test.go`](reports/internal-agent-provider-multi-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`DetectProviders`](reports/internal-agent-providers-go-detectproviders.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`DetectedProvider`](reports/internal-agent-providers-go-detectedprovider.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ProviderNames`](reports/internal-agent-providers-go-providernames.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](reports/internal-agent-providers-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`normalizeLocalURL`](reports/internal-agent-providers-go-normalizelocalurl.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`probeLocal`](reports/internal-agent-providers-go-probelocal.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Allow`](reports/internal-agent-ratelimit-go-allow.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewRateLimiter`](reports/internal-agent-ratelimit-go-newratelimiter.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`RateLimiter`](reports/internal-agent-ratelimit-go-ratelimiter.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Wait`](reports/internal-agent-ratelimit-go-wait.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`refill`](reports/internal-agent-ratelimit-go-refill.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ratelimit_test.go`](reports/internal-agent-ratelimit-test-go.md) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`NewReviewer`](reports/internal-agent-reviewer-go-newreviewer.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Review`](reports/internal-agent-reviewer-go-review.md) | method | B+ | 87.2% | certified | 2026-04-23 |
| [`ReviewInput`](reports/internal-agent-reviewer-go-reviewinput.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ReviewResult`](reports/internal-agent-reviewer-go-reviewresult.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Reviewer`](reports/internal-agent-reviewer-go-reviewer.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ToEvidence`](reports/internal-agent-reviewer-go-toevidence.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ToPrescreenEvidence`](reports/internal-agent-reviewer-go-toprescreenevidence.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`joinModels`](reports/internal-agent-reviewer-go-joinmodels.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`reviewer_test.go`](reports/internal-agent-reviewer-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`ModelFor`](reports/internal-agent-router-go-modelfor.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewRouter`](reports/internal-agent-router-go-newrouter.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Router`](reports/internal-agent-router-go-router.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`directAssignment`](reports/internal-agent-router-go-directassignment.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`router_test.go`](reports/internal-agent-router-test-go.md) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`DecisionResponse`](reports/internal-agent-schemas-go-decisionresponse.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`PrescreenResponse`](reports/internal-agent-schemas-go-prescreenresponse.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`RemediationResponse`](reports/internal-agent-schemas-go-remediationresponse.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`RemediationStep`](reports/internal-agent-schemas-go-remediationstep.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ScoringResponse`](reports/internal-agent-schemas-go-scoringresponse.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`schemas_test.go`](reports/internal-agent-schemas-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Execute`](reports/internal-agent-stage-go-execute.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Name`](reports/internal-agent-stage-go-name.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewPrescreenStage`](reports/internal-agent-stage-go-newprescreenstage.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewReviewStage`](reports/internal-agent-stage-go-newreviewstage.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewScoringStage`](reports/internal-agent-stage-go-newscoringstage.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Stage`](reports/internal-agent-stage-go-stage.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`StageInput`](reports/internal-agent-stage-go-stageinput.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`StageResult`](reports/internal-agent-stage-go-stageresult.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`defaultScores`](reports/internal-agent-stage-go-defaultscores.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`extractJSON`](reports/internal-agent-stage-go-extractjson.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`looseParseNeedsReview`](reports/internal-agent-stage-go-looseparseneedsreview.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`prescreenStage`](reports/internal-agent-stage-go-prescreenstage.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`reviewStage`](reports/internal-agent-stage-go-reviewstage.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`scoringStage`](reports/internal-agent-stage-go-scoringstage.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`DeepReviewResponse`](reports/internal-agent-stage-deep-go-deepreviewresponse.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Execute`](reports/internal-agent-stage-deep-go-execute.md) | method | B+ | 87.2% | certified | 2026-04-23 |
| [`FormatDeepObservations`](reports/internal-agent-stage-deep-go-formatdeepobservations.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`FormatReviewForRecord`](reports/internal-agent-stage-deep-go-formatreviewforrecord.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`IsDeepReview`](reports/internal-agent-stage-deep-go-isdeepreview.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Name`](reports/internal-agent-stage-deep-go-name.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewDeepReviewStage`](reports/internal-agent-stage-deep-go-newdeepreviewstage.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ToDeepEvidence`](reports/internal-agent-stage-deep-go-todeepevidence.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`deepReviewStage`](reports/internal-agent-stage-deep-go-deepreviewstage.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`extractFirstSentence`](reports/internal-agent-stage-deep-go-extractfirstsentence.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`stage_test.go`](reports/internal-agent-stage-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`RepoSummary`](reports/internal-agent-suggest-go-reposummary.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ScanSuggestion`](reports/internal-agent-suggest-go-scansuggestion.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`SuggestForRepo`](reports/internal-agent-suggest-go-suggestforrepo.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`buildSuggestPrompt`](reports/internal-agent-suggest-go-buildsuggestprompt.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`suggest_test.go`](reports/internal-agent-suggest-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`ChatRequest`](reports/internal-agent-types-go-chatrequest.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ChatResponse`](reports/internal-agent-types-go-chatresponse.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Choice`](reports/internal-agent-types-go-choice.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Content`](reports/internal-agent-types-go-content.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Message`](reports/internal-agent-types-go-message.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ModelConfig`](reports/internal-agent-types-go-modelconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ResponseFormat`](reports/internal-agent-types-go-responseformat.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`String`](reports/internal-agent-types-go-string.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`TaskType`](reports/internal-agent-types-go-tasktype.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Usage`](reports/internal-agent-types-go-usage.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`types_test.go`](reports/internal-agent-types-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The test file has good coverage of model attribution logic but lacks proper mocking and validation of actual model usage in execution paths.
- 💡 Add assertions to verify that each `mockProvider` or `sequenceProvider` is called with the correct model name in `TestStageResult_TracksModel` and `TestPipeline_CollectsAllModels`.
- 💡 Replace hardcoded string expectations like `"agent:qwen/qwen3-coder:free,mistralai/mistral-nemo"` with dynamic generation or validation of the source field to ensure robustness.
- 💡 Ensure that `sequenceProvider` validates that each response corresponds to the correct model or stage, and assert that all models in the pipeline are actually invoked.
- ⚠️ Mock validation gaps: Tests do not verify that actual model names are used in execution paths, leading to false positives if mocks don't reflect real behavior.
- ⚠️ Hardcoded string assertions: String formatting for evidence sources is not validated dynamically, which can cause silent failures if the format changes.
- 🔗 This unit tests model attribution logic but doesn't validate end-to-end execution, potentially masking issues in how models are selected or invoked across the pipeline.
- 🔗 The tests do not validate that model tracking is consistent across all stages of the agent lifecycle, which could lead to incomplete attribution in production.

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
- 🤖 The code is functionally correct but has several maintainability and risk issues, including hardcoded values, missing error handling in provider creation, and a potential logic flaw in prescreen assignment.
- 💡 Validate that `providers[0].Models` is not empty before accessing `providers[0].Models[0]` to prevent runtime panics
- 💡 Add error handling and validation when instantiating local and cloud providers to ensure robustness
- 💡 Parameterize `ConservativeCircuitThreshold` and `ConservativeTokenBudget` to allow configuration
- 💡 Use consistent provider initialization (e.g., avoid hardcoded strings like the GitHub URL) to prevent misconfiguration
- ⚠️ Index out of bounds panic if any provider has an empty Models list
- ⚠️ Unvalidated provider creation may lead to runtime errors or incorrect behavior
- 🔗 This function tightly couples to hardcoded provider configurations and assumes provider structure, increasing system fragility
- 🔗 The use of a fixed prescreen model from the first provider can cause incorrect behavior if that model is not suitable for pre-screening

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
- 🤖 The code is functionally correct but has several maintainability and risk issues, including hardcoded values, missing error handling in provider creation, and a potential logic flaw in prescreen assignment.
- 💡 Validate that `providers[0].Models` is not empty before accessing `providers[0].Models[0]` to prevent runtime panics
- 💡 Add error handling and validation when instantiating local and cloud providers to ensure robustness
- 💡 Parameterize `ConservativeCircuitThreshold` and `ConservativeTokenBudget` to allow configuration
- 💡 Use consistent provider initialization (e.g., avoid hardcoded strings like the GitHub URL) to prevent misconfiguration
- ⚠️ Index out of bounds panic if any provider has an empty Models list
- ⚠️ Unvalidated provider creation may lead to runtime errors or incorrect behavior
- 🔗 This function tightly couples to hardcoded provider configurations and assumes provider structure, increasing system fragility
- 🔗 The use of a fixed prescreen model from the first provider can cause incorrect behavior if that model is not suitable for pre-screening

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
- 🤖 The code is functionally correct but has several maintainability and risk issues, including hardcoded values, missing error handling in provider creation, and a potential logic flaw in prescreen assignment.
- 💡 Validate that `providers[0].Models` is not empty before accessing `providers[0].Models[0]` to prevent runtime panics
- 💡 Add error handling and validation when instantiating local and cloud providers to ensure robustness
- 💡 Parameterize `ConservativeCircuitThreshold` and `ConservativeTokenBudget` to allow configuration
- 💡 Use consistent provider initialization (e.g., avoid hardcoded strings like the GitHub URL) to prevent misconfiguration
- ⚠️ Index out of bounds panic if any provider has an empty Models list
- ⚠️ Unvalidated provider creation may lead to runtime errors or incorrect behavior
- 🔗 This function tightly couples to hardcoded provider configurations and assumes provider structure, increasing system fragility
- 🔗 The use of a fixed prescreen model from the first provider can cause incorrect behavior if that model is not suitable for pre-screening

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
- 🤖 The code is functionally correct but has several maintainability and risk issues, including hardcoded values, missing error handling in provider creation, and a potential logic flaw in prescreen assignment.
- 💡 Validate that `providers[0].Models` is not empty before accessing `providers[0].Models[0]` to prevent runtime panics
- 💡 Add error handling and validation when instantiating local and cloud providers to ensure robustness
- 💡 Parameterize `ConservativeCircuitThreshold` and `ConservativeTokenBudget` to allow configuration
- 💡 Use consistent provider initialization (e.g., avoid hardcoded strings like the GitHub URL) to prevent misconfiguration
- ⚠️ Index out of bounds panic if any provider has an empty Models list
- ⚠️ Unvalidated provider creation may lead to runtime errors or incorrect behavior
- 🔗 This function tightly couples to hardcoded provider configurations and assumes provider structure, increasing system fragility
- 🔗 The use of a fixed prescreen model from the first provider can cause incorrect behavior if that model is not suitable for pre-screening

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
- 🤖 The code is functionally correct but has several maintainability and risk issues, including hardcoded values, missing error handling in provider creation, and a potential logic flaw in prescreen assignment.
- 💡 Validate that `providers[0].Models` is not empty before accessing `providers[0].Models[0]` to prevent runtime panics
- 💡 Add error handling and validation when instantiating local and cloud providers to ensure robustness
- 💡 Parameterize `ConservativeCircuitThreshold` and `ConservativeTokenBudget` to allow configuration
- 💡 Use consistent provider initialization (e.g., avoid hardcoded strings like the GitHub URL) to prevent misconfiguration
- ⚠️ Index out of bounds panic if any provider has an empty Models list
- ⚠️ Unvalidated provider creation may lead to runtime errors or incorrect behavior
- 🔗 This function tightly couples to hardcoded provider configurations and assumes provider structure, increasing system fragility
- 🔗 The use of a fixed prescreen model from the first provider can cause incorrect behavior if that model is not suitable for pre-screening

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
- 🤖 The test file has good coverage but suffers from poor test isolation, redundant setup logic, and lacks proper mocking or cleanup for environment state.
- 💡 Use `t.Setenv()` and `t.Unsetenv()` instead of `os.Setenv()` and `os.Unsetenv()` to ensure proper test isolation in Go 1.18+
- 💡 Refactor `TestDetectAPIKey_*` tests to use a helper function that resets and sets only the specific environment variable needed for that test, rather than clearing all env vars
- 💡 Add assertions in `TestHasAnyProvider_NoProviders` to validate that no providers are detected when env vars are cleared
- 💡 Validate that `ConservativeModels` contains at least one model and that all models are valid (e.g., not empty strings) in `TestConservativeModels`
- 💡 Add a test case to verify that `DetectAPIKey` returns the correct key and envVar when multiple keys are set in priority order (e.g., OPENROUTER_API_KEY and CERTIFY_API_KEY)
- ⚠️ Race condition risk due to shared global environment state between tests
- ⚠️ Test isolation failure leading to flaky or unreliable test results
- 🔗 Tests that mutate global environment state increase risk of cross-contamination between test runs
- 🔗 Lack of proper cleanup in some tests can cause cascading failures in other parts of the system that rely on environment variables

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
- 🤖 The circuit breaker has a race condition due to inconsistent locking and lacks proper state management for recovery.
- 💡 Use a single lock for all operations in the `Chat` method to ensure atomicity of state checks and updates
- 💡 Validate that threshold is positive in `NewCircuitBreaker` to prevent invalid behavior
- 💡 Include the actual error from the provider in the returned error message to aid debugging
- ⚠️ Race condition in circuit state transitions leading to inconsistent behavior
- ⚠️ Potential panic or incorrect operation if threshold is zero or negative
- 🔗 The circuit breaker introduces a race condition that can cause concurrent calls to bypass circuit logic, leading to cascading failures in downstream systems
- 🔗 The unit tightly couples to the Provider interface and assumes it can be safely called under concurrent access without proper synchronization

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
- 🤖 The circuit breaker has a race condition due to inconsistent locking and lacks proper state management for recovery.
- 💡 Use a single lock for all operations in the `Chat` method to ensure atomicity of state checks and updates
- 💡 Validate that threshold is positive in `NewCircuitBreaker` to prevent invalid behavior
- 💡 Include the actual error from the provider in the returned error message to aid debugging
- ⚠️ Race condition in circuit state transitions leading to inconsistent behavior
- ⚠️ Potential panic or incorrect operation if threshold is zero or negative
- 🔗 The circuit breaker introduces a race condition that can cause concurrent calls to bypass circuit logic, leading to cascading failures in downstream systems
- 🔗 The unit tightly couples to the Provider interface and assumes it can be safely called under concurrent access without proper synchronization

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
- 🤖 The circuit breaker has a race condition due to inconsistent locking and lacks proper state management for recovery.
- 💡 Use a single lock for all operations in the `Chat` method to ensure atomicity of state checks and updates
- 💡 Validate that threshold is positive in `NewCircuitBreaker` to prevent invalid behavior
- 💡 Include the actual error from the provider in the returned error message to aid debugging
- ⚠️ Race condition in circuit state transitions leading to inconsistent behavior
- ⚠️ Potential panic or incorrect operation if threshold is zero or negative
- 🔗 The circuit breaker introduces a race condition that can cause concurrent calls to bypass circuit logic, leading to cascading failures in downstream systems
- 🔗 The unit tightly couples to the Provider interface and assumes it can be safely called under concurrent access without proper synchronization

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
- 🤖 The circuit breaker has a race condition due to inconsistent locking and lacks proper state management for recovery.
- 💡 Use a single lock for all operations in the `Chat` method to ensure atomicity of state checks and updates
- 💡 Validate that threshold is positive in `NewCircuitBreaker` to prevent invalid behavior
- 💡 Include the actual error from the provider in the returned error message to aid debugging
- ⚠️ Race condition in circuit state transitions leading to inconsistent behavior
- ⚠️ Potential panic or incorrect operation if threshold is zero or negative
- 🔗 The circuit breaker introduces a race condition that can cause concurrent calls to bypass circuit logic, leading to cascading failures in downstream systems
- 🔗 The unit tightly couples to the Provider interface and assumes it can be safely called under concurrent access without proper synchronization

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
- 🤖 The circuit breaker has a race condition due to inconsistent locking and lacks proper state management for recovery.
- 💡 Use a single lock for all operations in the `Chat` method to ensure atomicity of state checks and updates
- 💡 Validate that threshold is positive in `NewCircuitBreaker` to prevent invalid behavior
- 💡 Include the actual error from the provider in the returned error message to aid debugging
- ⚠️ Race condition in circuit state transitions leading to inconsistent behavior
- ⚠️ Potential panic or incorrect operation if threshold is zero or negative
- 🔗 The circuit breaker introduces a race condition that can cause concurrent calls to bypass circuit logic, leading to cascading failures in downstream systems
- 🔗 The unit tightly couples to the Provider interface and assumes it can be safely called under concurrent access without proper synchronization

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
- 🤖 The AdaptiveMessages function has a logic flaw in message construction and the fallback provider does not properly handle context cancellation or concurrent access.
- 💡 Modify AdaptiveMessages to sanitize or escape newlines in systemPrompt and userContent before concatenation, or use a structured format like JSON to avoid content corruption.
- 💡 Add context cancellation checks in the fallback loop of Chat() to ensure that long-running or blocked requests do not prevent other providers from being tried.
- 💡 Implement a timeout or retry mechanism in modelPinnedProvider to prevent indefinite blocking of the fallback chain.
- 💡 Add logging or metrics for failed provider attempts in FallbackProvider to aid debugging and monitoring.
- ⚠️ Improper message concatenation in AdaptiveMessages can cause malformed or duplicated content when systemPrompt or userContent contains newlines.
- ⚠️ Context cancellation is not respected in fallback providers, leading to potential resource leaks or indefinite waits.
- 🔗 The AdaptiveMessages function directly affects message formatting in chat flows, which could break downstream consumers expecting specific message structures.
- 🔗 The FallbackProvider and ModelChain introduce tight coupling between providers and the fallback logic, making it harder to swap or extend individual providers without affecting the entire chain.

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
- 🤖 The AdaptiveMessages function has a logic flaw in message construction and the fallback provider does not properly handle context cancellation or concurrent access.
- 💡 Modify AdaptiveMessages to sanitize or escape newlines in systemPrompt and userContent before concatenation, or use a structured format like JSON to avoid content corruption.
- 💡 Add context cancellation checks in the fallback loop of Chat() to ensure that long-running or blocked requests do not prevent other providers from being tried.
- 💡 Implement a timeout or retry mechanism in modelPinnedProvider to prevent indefinite blocking of the fallback chain.
- 💡 Add logging or metrics for failed provider attempts in FallbackProvider to aid debugging and monitoring.
- ⚠️ Improper message concatenation in AdaptiveMessages can cause malformed or duplicated content when systemPrompt or userContent contains newlines.
- ⚠️ Context cancellation is not respected in fallback providers, leading to potential resource leaks or indefinite waits.
- 🔗 The AdaptiveMessages function directly affects message formatting in chat flows, which could break downstream consumers expecting specific message structures.
- 🔗 The FallbackProvider and ModelChain introduce tight coupling between providers and the fallback logic, making it harder to swap or extend individual providers without affecting the entire chain.

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
- 🤖 The AdaptiveMessages function has a logic flaw in message construction and the fallback provider does not properly handle context cancellation or concurrent access.
- 💡 Modify AdaptiveMessages to sanitize or escape newlines in systemPrompt and userContent before concatenation, or use a structured format like JSON to avoid content corruption.
- 💡 Add context cancellation checks in the fallback loop of Chat() to ensure that long-running or blocked requests do not prevent other providers from being tried.
- 💡 Implement a timeout or retry mechanism in modelPinnedProvider to prevent indefinite blocking of the fallback chain.
- 💡 Add logging or metrics for failed provider attempts in FallbackProvider to aid debugging and monitoring.
- ⚠️ Improper message concatenation in AdaptiveMessages can cause malformed or duplicated content when systemPrompt or userContent contains newlines.
- ⚠️ Context cancellation is not respected in fallback providers, leading to potential resource leaks or indefinite waits.
- 🔗 The AdaptiveMessages function directly affects message formatting in chat flows, which could break downstream consumers expecting specific message structures.
- 🔗 The FallbackProvider and ModelChain introduce tight coupling between providers and the fallback logic, making it harder to swap or extend individual providers without affecting the entire chain.

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
- 🤖 The AdaptiveMessages function has a logic flaw in message construction and the fallback provider does not properly handle context cancellation or concurrent access.
- 💡 Modify AdaptiveMessages to sanitize or escape newlines in systemPrompt and userContent before concatenation, or use a structured format like JSON to avoid content corruption.
- 💡 Add context cancellation checks in the fallback loop of Chat() to ensure that long-running or blocked requests do not prevent other providers from being tried.
- 💡 Implement a timeout or retry mechanism in modelPinnedProvider to prevent indefinite blocking of the fallback chain.
- 💡 Add logging or metrics for failed provider attempts in FallbackProvider to aid debugging and monitoring.
- ⚠️ Improper message concatenation in AdaptiveMessages can cause malformed or duplicated content when systemPrompt or userContent contains newlines.
- ⚠️ Context cancellation is not respected in fallback providers, leading to potential resource leaks or indefinite waits.
- 🔗 The AdaptiveMessages function directly affects message formatting in chat flows, which could break downstream consumers expecting specific message structures.
- 🔗 The FallbackProvider and ModelChain introduce tight coupling between providers and the fallback logic, making it harder to swap or extend individual providers without affecting the entire chain.

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
- 🤖 The AdaptiveMessages function has a logic flaw in message construction and the fallback provider does not properly handle context cancellation or concurrent access.
- 💡 Modify AdaptiveMessages to sanitize or escape newlines in systemPrompt and userContent before concatenation, or use a structured format like JSON to avoid content corruption.
- 💡 Add context cancellation checks in the fallback loop of Chat() to ensure that long-running or blocked requests do not prevent other providers from being tried.
- 💡 Implement a timeout or retry mechanism in modelPinnedProvider to prevent indefinite blocking of the fallback chain.
- 💡 Add logging or metrics for failed provider attempts in FallbackProvider to aid debugging and monitoring.
- ⚠️ Improper message concatenation in AdaptiveMessages can cause malformed or duplicated content when systemPrompt or userContent contains newlines.
- ⚠️ Context cancellation is not respected in fallback providers, leading to potential resource leaks or indefinite waits.
- 🔗 The AdaptiveMessages function directly affects message formatting in chat flows, which could break downstream consumers expecting specific message structures.
- 🔗 The FallbackProvider and ModelChain introduce tight coupling between providers and the fallback logic, making it harder to swap or extend individual providers without affecting the entire chain.

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
- 🤖 The AdaptiveMessages function has a logic flaw in message construction and the fallback provider does not properly handle context cancellation or concurrent access.
- 💡 Modify AdaptiveMessages to sanitize or escape newlines in systemPrompt and userContent before concatenation, or use a structured format like JSON to avoid content corruption.
- 💡 Add context cancellation checks in the fallback loop of Chat() to ensure that long-running or blocked requests do not prevent other providers from being tried.
- 💡 Implement a timeout or retry mechanism in modelPinnedProvider to prevent indefinite blocking of the fallback chain.
- 💡 Add logging or metrics for failed provider attempts in FallbackProvider to aid debugging and monitoring.
- ⚠️ Improper message concatenation in AdaptiveMessages can cause malformed or duplicated content when systemPrompt or userContent contains newlines.
- ⚠️ Context cancellation is not respected in fallback providers, leading to potential resource leaks or indefinite waits.
- 🔗 The AdaptiveMessages function directly affects message formatting in chat flows, which could break downstream consumers expecting specific message structures.
- 🔗 The FallbackProvider and ModelChain introduce tight coupling between providers and the fallback logic, making it harder to swap or extend individual providers without affecting the entire chain.

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
- 🤖 The AdaptiveMessages function has a logic flaw in message construction and the fallback provider does not properly handle context cancellation or concurrent access.
- 💡 Modify AdaptiveMessages to sanitize or escape newlines in systemPrompt and userContent before concatenation, or use a structured format like JSON to avoid content corruption.
- 💡 Add context cancellation checks in the fallback loop of Chat() to ensure that long-running or blocked requests do not prevent other providers from being tried.
- 💡 Implement a timeout or retry mechanism in modelPinnedProvider to prevent indefinite blocking of the fallback chain.
- 💡 Add logging or metrics for failed provider attempts in FallbackProvider to aid debugging and monitoring.
- ⚠️ Improper message concatenation in AdaptiveMessages can cause malformed or duplicated content when systemPrompt or userContent contains newlines.
- ⚠️ Context cancellation is not respected in fallback providers, leading to potential resource leaks or indefinite waits.
- 🔗 The AdaptiveMessages function directly affects message formatting in chat flows, which could break downstream consumers expecting specific message structures.
- 🔗 The FallbackProvider and ModelChain introduce tight coupling between providers and the fallback logic, making it harder to swap or extend individual providers without affecting the entire chain.

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
- 🤖 The AdaptiveMessages function has a logic flaw in message construction and the fallback provider does not properly handle context cancellation or concurrent access.
- 💡 Modify AdaptiveMessages to sanitize or escape newlines in systemPrompt and userContent before concatenation, or use a structured format like JSON to avoid content corruption.
- 💡 Add context cancellation checks in the fallback loop of Chat() to ensure that long-running or blocked requests do not prevent other providers from being tried.
- 💡 Implement a timeout or retry mechanism in modelPinnedProvider to prevent indefinite blocking of the fallback chain.
- 💡 Add logging or metrics for failed provider attempts in FallbackProvider to aid debugging and monitoring.
- ⚠️ Improper message concatenation in AdaptiveMessages can cause malformed or duplicated content when systemPrompt or userContent contains newlines.
- ⚠️ Context cancellation is not respected in fallback providers, leading to potential resource leaks or indefinite waits.
- 🔗 The AdaptiveMessages function directly affects message formatting in chat flows, which could break downstream consumers expecting specific message structures.
- 🔗 The FallbackProvider and ModelChain introduce tight coupling between providers and the fallback logic, making it harder to swap or extend individual providers without affecting the entire chain.

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
- 🤖 The code is functionally correct but has critical test design flaws, lacks proper error handling in helpers, and introduces potential race conditions in concurrent testing.
- 💡 Use sync.Mutex or atomic operations to protect the calls slice in trackingProvider
- 💡 Replace custom contains function with strings.Contains for better performance and clarity
- 💡 Add assertions to verify that providers are called in the correct order in TestFallbackProvider_TriesModelsInOrder
- 💡 Add validation to ensure that each model in the chain is correctly instantiated and used by the ModelChain
- 💡 Add test cases for context cancellation and timeout handling in fallback logic
- ⚠️ Race condition in trackingProvider due to unsynchronized access to shared calls slice
- ⚠️ Inefficient and redundant custom contains function that can be replaced with strings.Contains
- 🔗 The trackingProvider is not thread-safe and may cause flaky tests or data corruption in concurrent environments
- 🔗 The fallback logic may silently skip important error types (e.g., 400, 422) instead of trying fallbacks, leading to incorrect behavior in production

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
- 🤖 The code correctly implements fallback logic for model listing from OpenAI-compatible and Ollama APIs, but has critical security and error handling flaws.
- 💡 Fix the fallback logic to return Ollama models when OpenAI fails, not the original error
- 💡 Validate and sanitize the baseURL input before constructing API endpoints to prevent URL injection
- 💡 Add proper error handling for io.ReadAll and ensure all errors are logged or returned
- 💡 Implement retry logic with exponential backoff for HTTP requests to handle transient failures
- ⚠️ URL injection vulnerability due to direct string concatenation with unvalidated base URLs
- ⚠️ Resource leak from discarding error when reading HTTP response body
- 🔗 The fallback logic is fundamentally broken - it returns the first error instead of the successful fallback result
- 🔗 The system may silently fail to discover models from Ollama if OpenAI endpoint fails, leading to incorrect behavior in downstream components

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
- 🤖 The code correctly implements fallback logic for model listing from OpenAI-compatible and Ollama APIs, but has critical security and error handling flaws.
- 💡 Fix the fallback logic to return Ollama models when OpenAI fails, not the original error
- 💡 Validate and sanitize the baseURL input before constructing API endpoints to prevent URL injection
- 💡 Add proper error handling for io.ReadAll and ensure all errors are logged or returned
- 💡 Implement retry logic with exponential backoff for HTTP requests to handle transient failures
- ⚠️ URL injection vulnerability due to direct string concatenation with unvalidated base URLs
- ⚠️ Resource leak from discarding error when reading HTTP response body
- 🔗 The fallback logic is fundamentally broken - it returns the first error instead of the successful fallback result
- 🔗 The system may silently fail to discover models from Ollama if OpenAI endpoint fails, leading to incorrect behavior in downstream components

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
- 🤖 The code correctly implements fallback logic for model listing from OpenAI-compatible and Ollama APIs, but has critical security and error handling flaws.
- 💡 Fix the fallback logic to return Ollama models when OpenAI fails, not the original error
- 💡 Validate and sanitize the baseURL input before constructing API endpoints to prevent URL injection
- 💡 Add proper error handling for io.ReadAll and ensure all errors are logged or returned
- 💡 Implement retry logic with exponential backoff for HTTP requests to handle transient failures
- ⚠️ URL injection vulnerability due to direct string concatenation with unvalidated base URLs
- ⚠️ Resource leak from discarding error when reading HTTP response body
- 🔗 The fallback logic is fundamentally broken - it returns the first error instead of the successful fallback result
- 🔗 The system may silently fail to discover models from Ollama if OpenAI endpoint fails, leading to incorrect behavior in downstream components

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
- 🤖 The code correctly implements fallback logic for model listing from OpenAI-compatible and Ollama APIs, but has critical security and error handling flaws.
- 💡 Fix the fallback logic to return Ollama models when OpenAI fails, not the original error
- 💡 Validate and sanitize the baseURL input before constructing API endpoints to prevent URL injection
- 💡 Add proper error handling for io.ReadAll and ensure all errors are logged or returned
- 💡 Implement retry logic with exponential backoff for HTTP requests to handle transient failures
- ⚠️ URL injection vulnerability due to direct string concatenation with unvalidated base URLs
- ⚠️ Resource leak from discarding error when reading HTTP response body
- 🔗 The fallback logic is fundamentally broken - it returns the first error instead of the successful fallback result
- 🔗 The system may silently fail to discover models from Ollama if OpenAI endpoint fails, leading to incorrect behavior in downstream components

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
- 🤖 The code correctly implements fallback logic for model listing from OpenAI-compatible and Ollama APIs, but has critical security and error handling flaws.
- 💡 Fix the fallback logic to return Ollama models when OpenAI fails, not the original error
- 💡 Validate and sanitize the baseURL input before constructing API endpoints to prevent URL injection
- 💡 Add proper error handling for io.ReadAll and ensure all errors are logged or returned
- 💡 Implement retry logic with exponential backoff for HTTP requests to handle transient failures
- ⚠️ URL injection vulnerability due to direct string concatenation with unvalidated base URLs
- ⚠️ Resource leak from discarding error when reading HTTP response body
- 🔗 The fallback logic is fundamentally broken - it returns the first error instead of the successful fallback result
- 🔗 The system may silently fail to discover models from Ollama if OpenAI endpoint fails, leading to incorrect behavior in downstream components

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
- 🤖 The code correctly implements fallback logic for model listing from OpenAI-compatible and Ollama APIs, but has critical security and error handling flaws.
- 💡 Fix the fallback logic to return Ollama models when OpenAI fails, not the original error
- 💡 Validate and sanitize the baseURL input before constructing API endpoints to prevent URL injection
- 💡 Add proper error handling for io.ReadAll and ensure all errors are logged or returned
- 💡 Implement retry logic with exponential backoff for HTTP requests to handle transient failures
- ⚠️ URL injection vulnerability due to direct string concatenation with unvalidated base URLs
- ⚠️ Resource leak from discarding error when reading HTTP response body
- 🔗 The fallback logic is fundamentally broken - it returns the first error instead of the successful fallback result
- 🔗 The system may silently fail to discover models from Ollama if OpenAI endpoint fails, leading to incorrect behavior in downstream components

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
- 🤖 The code correctly implements fallback logic for model listing from OpenAI-compatible and Ollama APIs, but has critical security and error handling flaws.
- 💡 Fix the fallback logic to return Ollama models when OpenAI fails, not the original error
- 💡 Validate and sanitize the baseURL input before constructing API endpoints to prevent URL injection
- 💡 Add proper error handling for io.ReadAll and ensure all errors are logged or returned
- 💡 Implement retry logic with exponential backoff for HTTP requests to handle transient failures
- ⚠️ URL injection vulnerability due to direct string concatenation with unvalidated base URLs
- ⚠️ Resource leak from discarding error when reading HTTP response body
- 🔗 The fallback logic is fundamentally broken - it returns the first error instead of the successful fallback result
- 🔗 The system may silently fail to discover models from Ollama if OpenAI endpoint fails, leading to incorrect behavior in downstream components

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
- 🤖 The code correctly implements fallback logic for model listing from OpenAI-compatible and Ollama APIs, but has critical security and error handling flaws.
- 💡 Fix the fallback logic to return Ollama models when OpenAI fails, not the original error
- 💡 Validate and sanitize the baseURL input before constructing API endpoints to prevent URL injection
- 💡 Add proper error handling for io.ReadAll and ensure all errors are logged or returned
- 💡 Implement retry logic with exponential backoff for HTTP requests to handle transient failures
- ⚠️ URL injection vulnerability due to direct string concatenation with unvalidated base URLs
- ⚠️ Resource leak from discarding error when reading HTTP response body
- 🔗 The fallback logic is fundamentally broken - it returns the first error instead of the successful fallback result
- 🔗 The system may silently fail to discover models from Ollama if OpenAI endpoint fails, leading to incorrect behavior in downstream components

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
- 🤖 The test suite for ListModels has good coverage but lacks proper error handling in mock servers and has inconsistent behavior with auth headers.
- 💡 Add a test case for malformed JSON responses to ensure proper error handling
- 💡 Ensure all possible URL paths are handled consistently in the Ollama format test server to prevent unexpected 404s
- 💡 Validate that all model fields (Created, OwnedBy) are correctly populated in test assertions
- 💡 Add a test for context cancellation to ensure proper cleanup and prevent goroutine leaks
- ⚠️ Inconsistent handling of Ollama format where /api/tags is only handled in one case but not others
- ⚠️ Potential panic or incorrect parsing if server returns malformed JSON
- 🔗 The tests do not verify proper context cancellation handling, which could lead to resource leaks in the actual implementation
- 🔗 Mock server behavior does not fully simulate real-world API responses, potentially masking bugs in the ListModels function

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The code implements a robust OpenRouter-compatible provider with retry logic and error handling, but has missing validation for URL construction, potential security risks in header injection, and lacks proper test coverage for edge cases.
- 💡 Validate and sanitize the `baseURL` before appending `/chat/completions` to prevent URL manipulation
- 💡 Sanitize or validate `httpReferer` and `xTitle` values before setting them as HTTP headers to prevent header injection
- 💡 Add validation for `maxRetries` and make it configurable via constructor to allow tuning in production
- 💡 Implement proper error handling for 4xx status codes by checking response body before parsing JSON to avoid panics
- 💡 Use a more robust retry policy that includes client errors (400s) when appropriate, and document which status codes are retryable
- ⚠️ URL concatenation without validation can lead to malformed URLs or path traversal vulnerabilities
- ⚠️ Unsanitized HTTP headers (HTTP-Referer, X-Title) may allow header injection attacks
- 🔗 This provider tightly couples to the HTTP client and assumes all OpenAI-compatible endpoints follow the same schema, limiting flexibility for future API changes
- 🔗 The lack of proper error categorization and retry logic for client errors (4xx) can cause silent failures in downstream systems

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
- 🤖 The test suite for OpenRouter provider is well-structured and covers several key scenarios including success, headers, error handling, rate limiting, and name validation, but has some gaps in error handling and test isolation.
- 💡 Replace the global `calls` variable with a local counter or use a mockable retry mechanism to improve test isolation and reliability.
- 💡 Add explicit error checking after `io.ReadAll(r.Body)` (line 20) to ensure I/O errors are handled gracefully.
- 💡 Use `t.Run()` to encapsulate each test case and ensure proper cleanup, especially for the header test that uses a global variable.
- 💡 Assert specific error types and messages in `TestOpenRouter_MissingAPIKey` to ensure correct error handling.
- 💡 Add a test for malformed JSON response from the server to verify proper error handling in that scenario.
- ⚠️ Use of global variable `calls` in rate limit test (line 72) introduces potential race conditions and makes tests brittle.
- ⚠️ Unvalidated request body parsing (line 20) can lead to runtime panics or incorrect test behavior if the request is malformed.
- 🔗 These tests validate the OpenRouter provider's integration points, but lack coverage for concurrent access patterns that might reveal race conditions.
- 🔗 The test suite does not validate error propagation or retry logic under realistic network conditions, which could mask failures in production.

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
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

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
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

</details>

<a id="internal-agent-pipeline-go-islocal"></a>
<details>
<summary>IsLocal — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

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
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

</details>

<a id="internal-agent-pipeline-go-newpipeline"></a>
<details>
<summary>NewPipeline — certified details</summary>

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
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

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
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

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
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

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
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

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
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

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
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

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
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

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
- 🤖 The code is functionally correct but has a critical race condition in Stats() and lacks proper error handling for token budgeting.
- 💡 Add mutex locking around c.reviewed map access in Stats() and ReviewUnit to prevent race conditions
- 💡 Move token budget update (c.tokensSpent += result.TokensUsed) before returning the result in ReviewUnit to ensure atomicity
- 💡 Replace fmt.Sprintf("%v", summaries) with a proper string concatenation or JSON marshaling for better performance and correctness
- 💡 Clarify or remove the fallback logic in StrategyLocal where Review model defaults to Prescreen model
- 💡 Add unit tests for concurrent access to Stats() and ReviewUnit to validate thread safety
- ⚠️ Race condition in Stats() due to unsynchronized access to c.reviewed map
- ⚠️ Token budgeting logic is not thread-safe and can lead to overconsumption
- 🔗 Stats() method is called by external systems for monitoring and can cause data races
- 🔗 Coordinator is tightly coupled to Pipeline and Provider interfaces, making it hard to test in isolation

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
- 🤖 The code is functionally correct but has critical race conditions, poor error handling for file operations, and lacks proper versioning logic.
- 💡 Replace `strings.ReplaceAll` with a proper templating engine like `text/template` to prevent incorrect variable substitution
- 💡 Sort glob matches using semantic version comparison instead of lexicographical sorting to ensure correct version selection
- 💡 Add input validation for filename patterns and enforce consistent naming conventions
- 💡 Implement proper concurrency control around file loading to prevent race conditions during template initialization
- ⚠️ Race condition in concurrent access to the same file path during loading
- ⚠️ Incorrect variable substitution due to use of strings.ReplaceAll instead of proper templating engine
- 🔗 The registry introduces tight coupling between task types and file naming conventions
- 🔗 Failure propagation through the registry can cause cascading errors in downstream systems that depend on prompt templates

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
- 🤖 The code is functionally correct but has critical race conditions, poor error handling for file operations, and lacks proper versioning logic.
- 💡 Replace `strings.ReplaceAll` with a proper templating engine like `text/template` to prevent incorrect variable substitution
- 💡 Sort glob matches using semantic version comparison instead of lexicographical sorting to ensure correct version selection
- 💡 Add input validation for filename patterns and enforce consistent naming conventions
- 💡 Implement proper concurrency control around file loading to prevent race conditions during template initialization
- ⚠️ Race condition in concurrent access to the same file path during loading
- ⚠️ Incorrect variable substitution due to use of strings.ReplaceAll instead of proper templating engine
- 🔗 The registry introduces tight coupling between task types and file naming conventions
- 🔗 Failure propagation through the registry can cause cascading errors in downstream systems that depend on prompt templates

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
- 🤖 The code is functionally correct but has critical race conditions, poor error handling for file operations, and lacks proper versioning logic.
- 💡 Replace `strings.ReplaceAll` with a proper templating engine like `text/template` to prevent incorrect variable substitution
- 💡 Sort glob matches using semantic version comparison instead of lexicographical sorting to ensure correct version selection
- 💡 Add input validation for filename patterns and enforce consistent naming conventions
- 💡 Implement proper concurrency control around file loading to prevent race conditions during template initialization
- ⚠️ Race condition in concurrent access to the same file path during loading
- ⚠️ Incorrect variable substitution due to use of strings.ReplaceAll instead of proper templating engine
- 🔗 The registry introduces tight coupling between task types and file naming conventions
- 🔗 Failure propagation through the registry can cause cascading errors in downstream systems that depend on prompt templates

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
- 🤖 The code is functionally correct but has critical race conditions, poor error handling for file operations, and lacks proper versioning logic.
- 💡 Replace `strings.ReplaceAll` with a proper templating engine like `text/template` to prevent incorrect variable substitution
- 💡 Sort glob matches using semantic version comparison instead of lexicographical sorting to ensure correct version selection
- 💡 Add input validation for filename patterns and enforce consistent naming conventions
- 💡 Implement proper concurrency control around file loading to prevent race conditions during template initialization
- ⚠️ Race condition in concurrent access to the same file path during loading
- ⚠️ Incorrect variable substitution due to use of strings.ReplaceAll instead of proper templating engine
- 🔗 The registry introduces tight coupling between task types and file naming conventions
- 🔗 Failure propagation through the registry can cause cascading errors in downstream systems that depend on prompt templates

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
- 🤖 The code is functionally correct but has critical race conditions, poor error handling for file operations, and lacks proper versioning logic.
- 💡 Replace `strings.ReplaceAll` with a proper templating engine like `text/template` to prevent incorrect variable substitution
- 💡 Sort glob matches using semantic version comparison instead of lexicographical sorting to ensure correct version selection
- 💡 Add input validation for filename patterns and enforce consistent naming conventions
- 💡 Implement proper concurrency control around file loading to prevent race conditions during template initialization
- ⚠️ Race condition in concurrent access to the same file path during loading
- ⚠️ Incorrect variable substitution due to use of strings.ReplaceAll instead of proper templating engine
- 🔗 The registry introduces tight coupling between task types and file naming conventions
- 🔗 Failure propagation through the registry can cause cascading errors in downstream systems that depend on prompt templates

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
- 🤖 The code is functionally correct but has critical race conditions, poor error handling for file operations, and lacks proper versioning logic.
- 💡 Replace `strings.ReplaceAll` with a proper templating engine like `text/template` to prevent incorrect variable substitution
- 💡 Sort glob matches using semantic version comparison instead of lexicographical sorting to ensure correct version selection
- 💡 Add input validation for filename patterns and enforce consistent naming conventions
- 💡 Implement proper concurrency control around file loading to prevent race conditions during template initialization
- ⚠️ Race condition in concurrent access to the same file path during loading
- ⚠️ Incorrect variable substitution due to use of strings.ReplaceAll instead of proper templating engine
- 🔗 The registry introduces tight coupling between task types and file naming conventions
- 🔗 Failure propagation through the registry can cause cascading errors in downstream systems that depend on prompt templates

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
- 🤖 The code is functionally correct but has critical race conditions, poor error handling for file operations, and lacks proper versioning logic.
- 💡 Replace `strings.ReplaceAll` with a proper templating engine like `text/template` to prevent incorrect variable substitution
- 💡 Sort glob matches using semantic version comparison instead of lexicographical sorting to ensure correct version selection
- 💡 Add input validation for filename patterns and enforce consistent naming conventions
- 💡 Implement proper concurrency control around file loading to prevent race conditions during template initialization
- ⚠️ Race condition in concurrent access to the same file path during loading
- ⚠️ Incorrect variable substitution due to use of strings.ReplaceAll instead of proper templating engine
- 🔗 The registry introduces tight coupling between task types and file naming conventions
- 🔗 Failure propagation through the registry can cause cascading errors in downstream systems that depend on prompt templates

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
- 🤖 The test file has good coverage for basic prompt loading and rendering but lacks proper error handling in tests, has hardcoded paths that are brittle, and does not validate rendered content.
- 💡 Replace hardcoded relative paths in `promptsDir()` with a configurable or dynamically resolved path to make tests more robust
- 💡 Add assertions to validate that rendered template content contains expected variable substitutions and matches known good outputs
- 💡 Use `t.Errorf` instead of `t.Fatalf` in `TestLoadPrompt` to allow other tests to continue running even if one fails
- 💡 Implement a golden file approach or compare rendered output against known reference values to ensure template integrity
- ⚠️ Brittle file path assumptions in `promptsDir()` lead to test instability when project structure changes
- ⚠️ Lack of content validation in rendered templates means incorrect or malformed output may go undetected
- 🔗 This test file is tightly coupled to the internal directory structure of templates, increasing maintenance burden when templates are reorganized
- 🔗 The tests do not validate correctness of rendered prompts, which could lead to silent failures in downstream systems using these templates

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
- 🤖 Minimal interface with no implementation details, lacks proper error handling and validation patterns
- 💡 Add specific documentation for expected error types and categorization patterns (e.g., retryable vs non-retryable errors)
- 💡 Define concrete ChatRequest and ChatResponse types with validation rules to ensure consistent behavior across implementations
- 💡 Add context timeout handling documentation or constraints to prevent unbounded request processing
- 💡 Specify naming conventions for provider names to ensure uniqueness and avoid conflicts
- ⚠️ No input validation or sanitization patterns - implementations may not handle malformed requests properly
- ⚠️ No documented error types or categories - client code cannot properly distinguish between transient and permanent failures
- 🔗 High coupling to concrete implementation details - any changes to ChatRequest/ChatResponse will break all implementations
- 🔗 Limited testability due to lack of constraints - cannot create meaningful unit tests without concrete types

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
- 🤖 The test file has good coverage for provider detection and fallback logic but contains several critical race conditions, missing error handling in mocks, and improper environment variable cleanup that could lead to flaky tests or incorrect behavior.
- 💡 Add mutex to mockProvider to prevent race conditions when used in concurrent tests
- 💡 Fix clearProviderEnvVars to properly unset all relevant environment variables including GROQ_API_KEY, OPENAI_API_KEY, OLLAMA_HOST, and LM_STUDIO_URL
- 💡 Wrap httptest.NewServer in a proper test context to ensure the server is fully initialized before setting env vars
- 💡 Add validation in TestOpenRouterProvider_LocalNoAuth to ensure the test server is actually serving content before proceeding
- ⚠️ Race condition in mockProvider due to lack of synchronization
- ⚠️ Improper environment variable cleanup leading to test pollution and flaky tests
- 🔗 The mockProvider race condition could cause intermittent test failures that mask real concurrency issues in the actual provider implementations
- 🔗 Improper environment variable handling could cause cross-contamination between tests, making the test suite unreliable and causing cascading failures

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
- 🤖 The code is functionally correct but has security and maintainability issues including hardcoded API key exposure, lack of input validation, and fragile local provider detection logic.
- 💡 Sanitize or redact API keys before logging or storing them in memory
- 💡 Add authentication checks or model-specific probes to `probeLocal` to prevent false positives
- 💡 Replace `init()`-based population of `DefaultModels` with a function that accepts the map as input to improve testability
- 💡 Validate and normalize URLs more carefully in `normalizeLocalURL` to support various local LLM server formats
- 💡 Use a more efficient slice-to-string conversion in `ProviderNames` (e.g., `make([]string, 0, len(providers))`)
- ⚠️ Hardcoded API key exposure in memory and logs
- ⚠️ False positive local provider detection due to basic HTTP probe
- 🔗 Relies on external environment variables and local services, increasing system fragility
- 🔗 Uses global state in `DefaultModels` map that's populated at init time, causing testability and determinism issues

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
- 🤖 The code is functionally correct but has security and maintainability issues including hardcoded API key exposure, lack of input validation, and fragile local provider detection logic.
- 💡 Sanitize or redact API keys before logging or storing them in memory
- 💡 Add authentication checks or model-specific probes to `probeLocal` to prevent false positives
- 💡 Replace `init()`-based population of `DefaultModels` with a function that accepts the map as input to improve testability
- 💡 Validate and normalize URLs more carefully in `normalizeLocalURL` to support various local LLM server formats
- 💡 Use a more efficient slice-to-string conversion in `ProviderNames` (e.g., `make([]string, 0, len(providers))`)
- ⚠️ Hardcoded API key exposure in memory and logs
- ⚠️ False positive local provider detection due to basic HTTP probe
- 🔗 Relies on external environment variables and local services, increasing system fragility
- 🔗 Uses global state in `DefaultModels` map that's populated at init time, causing testability and determinism issues

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
- 🤖 The code is functionally correct but has security and maintainability issues including hardcoded API key exposure, lack of input validation, and fragile local provider detection logic.
- 💡 Sanitize or redact API keys before logging or storing them in memory
- 💡 Add authentication checks or model-specific probes to `probeLocal` to prevent false positives
- 💡 Replace `init()`-based population of `DefaultModels` with a function that accepts the map as input to improve testability
- 💡 Validate and normalize URLs more carefully in `normalizeLocalURL` to support various local LLM server formats
- 💡 Use a more efficient slice-to-string conversion in `ProviderNames` (e.g., `make([]string, 0, len(providers))`)
- ⚠️ Hardcoded API key exposure in memory and logs
- ⚠️ False positive local provider detection due to basic HTTP probe
- 🔗 Relies on external environment variables and local services, increasing system fragility
- 🔗 Uses global state in `DefaultModels` map that's populated at init time, causing testability and determinism issues

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
- 🤖 The code is functionally correct but has security and maintainability issues including hardcoded API key exposure, lack of input validation, and fragile local provider detection logic.
- 💡 Sanitize or redact API keys before logging or storing them in memory
- 💡 Add authentication checks or model-specific probes to `probeLocal` to prevent false positives
- 💡 Replace `init()`-based population of `DefaultModels` with a function that accepts the map as input to improve testability
- 💡 Validate and normalize URLs more carefully in `normalizeLocalURL` to support various local LLM server formats
- 💡 Use a more efficient slice-to-string conversion in `ProviderNames` (e.g., `make([]string, 0, len(providers))`)
- ⚠️ Hardcoded API key exposure in memory and logs
- ⚠️ False positive local provider detection due to basic HTTP probe
- 🔗 Relies on external environment variables and local services, increasing system fragility
- 🔗 Uses global state in `DefaultModels` map that's populated at init time, causing testability and determinism issues

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
- 🤖 The code is functionally correct but has security and maintainability issues including hardcoded API key exposure, lack of input validation, and fragile local provider detection logic.
- 💡 Sanitize or redact API keys before logging or storing them in memory
- 💡 Add authentication checks or model-specific probes to `probeLocal` to prevent false positives
- 💡 Replace `init()`-based population of `DefaultModels` with a function that accepts the map as input to improve testability
- 💡 Validate and normalize URLs more carefully in `normalizeLocalURL` to support various local LLM server formats
- 💡 Use a more efficient slice-to-string conversion in `ProviderNames` (e.g., `make([]string, 0, len(providers))`)
- ⚠️ Hardcoded API key exposure in memory and logs
- ⚠️ False positive local provider detection due to basic HTTP probe
- 🔗 Relies on external environment variables and local services, increasing system fragility
- 🔗 Uses global state in `DefaultModels` map that's populated at init time, causing testability and determinism issues

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
- 🤖 The code is functionally correct but has security and maintainability issues including hardcoded API key exposure, lack of input validation, and fragile local provider detection logic.
- 💡 Sanitize or redact API keys before logging or storing them in memory
- 💡 Add authentication checks or model-specific probes to `probeLocal` to prevent false positives
- 💡 Replace `init()`-based population of `DefaultModels` with a function that accepts the map as input to improve testability
- 💡 Validate and normalize URLs more carefully in `normalizeLocalURL` to support various local LLM server formats
- 💡 Use a more efficient slice-to-string conversion in `ProviderNames` (e.g., `make([]string, 0, len(providers))`)
- ⚠️ Hardcoded API key exposure in memory and logs
- ⚠️ False positive local provider detection due to basic HTTP probe
- 🔗 Relies on external environment variables and local services, increasing system fragility
- 🔗 Uses global state in `DefaultModels` map that's populated at init time, causing testability and determinism issues

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
- 🤖 The rate limiter has a race condition in `Wait()` and inefficient token refilling logic; correctness and performance are compromised.
- 💡 Replace the busy loop in `Wait()` with a channel-based or condition variable approach to avoid spinning and reduce CPU overhead
- 💡 Fix `refill()` logic to properly compute token additions based on elapsed time rather than assuming full intervals
- 💡 Ensure `min()` is properly imported or defined in the package to avoid runtime errors
- ⚠️ Race condition in `Wait()` due to repeated locking and unlocking of the same mutex without proper synchronization
- ⚠️ Token refill logic incorrectly assumes full intervals passed, leading to overfilling and incorrect rate limiting behavior
- 🔗 The `Wait()` method creates a busy loop that can consume excessive CPU resources under high load
- 🔗 The flawed refill logic introduces inconsistency in rate limiting behavior, affecting downstream systems that depend on accurate throttling

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
- 🤖 The rate limiter has a race condition in `Wait()` and inefficient token refilling logic; correctness and performance are compromised.
- 💡 Replace the busy loop in `Wait()` with a channel-based or condition variable approach to avoid spinning and reduce CPU overhead
- 💡 Fix `refill()` logic to properly compute token additions based on elapsed time rather than assuming full intervals
- 💡 Ensure `min()` is properly imported or defined in the package to avoid runtime errors
- ⚠️ Race condition in `Wait()` due to repeated locking and unlocking of the same mutex without proper synchronization
- ⚠️ Token refill logic incorrectly assumes full intervals passed, leading to overfilling and incorrect rate limiting behavior
- 🔗 The `Wait()` method creates a busy loop that can consume excessive CPU resources under high load
- 🔗 The flawed refill logic introduces inconsistency in rate limiting behavior, affecting downstream systems that depend on accurate throttling

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
- 🤖 The rate limiter has a race condition in `Wait()` and inefficient token refilling logic; correctness and performance are compromised.
- 💡 Replace the busy loop in `Wait()` with a channel-based or condition variable approach to avoid spinning and reduce CPU overhead
- 💡 Fix `refill()` logic to properly compute token additions based on elapsed time rather than assuming full intervals
- 💡 Ensure `min()` is properly imported or defined in the package to avoid runtime errors
- ⚠️ Race condition in `Wait()` due to repeated locking and unlocking of the same mutex without proper synchronization
- ⚠️ Token refill logic incorrectly assumes full intervals passed, leading to overfilling and incorrect rate limiting behavior
- 🔗 The `Wait()` method creates a busy loop that can consume excessive CPU resources under high load
- 🔗 The flawed refill logic introduces inconsistency in rate limiting behavior, affecting downstream systems that depend on accurate throttling

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
- 🤖 The rate limiter has a race condition in `Wait()` and inefficient token refilling logic; correctness and performance are compromised.
- 💡 Replace the busy loop in `Wait()` with a channel-based or condition variable approach to avoid spinning and reduce CPU overhead
- 💡 Fix `refill()` logic to properly compute token additions based on elapsed time rather than assuming full intervals
- 💡 Ensure `min()` is properly imported or defined in the package to avoid runtime errors
- ⚠️ Race condition in `Wait()` due to repeated locking and unlocking of the same mutex without proper synchronization
- ⚠️ Token refill logic incorrectly assumes full intervals passed, leading to overfilling and incorrect rate limiting behavior
- 🔗 The `Wait()` method creates a busy loop that can consume excessive CPU resources under high load
- 🔗 The flawed refill logic introduces inconsistency in rate limiting behavior, affecting downstream systems that depend on accurate throttling

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
- 🤖 The rate limiter has a race condition in `Wait()` and inefficient token refilling logic; correctness and performance are compromised.
- 💡 Replace the busy loop in `Wait()` with a channel-based or condition variable approach to avoid spinning and reduce CPU overhead
- 💡 Fix `refill()` logic to properly compute token additions based on elapsed time rather than assuming full intervals
- 💡 Ensure `min()` is properly imported or defined in the package to avoid runtime errors
- ⚠️ Race condition in `Wait()` due to repeated locking and unlocking of the same mutex without proper synchronization
- ⚠️ Token refill logic incorrectly assumes full intervals passed, leading to overfilling and incorrect rate limiting behavior
- 🔗 The `Wait()` method creates a busy loop that can consume excessive CPU resources under high load
- 🔗 The flawed refill logic introduces inconsistency in rate limiting behavior, affecting downstream systems that depend on accurate throttling

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
- 🤖 The rate limiter tests are basic and lack coverage for edge cases, concurrency, and precise timing validation.
- 💡 Add concurrent tests using `t.Run` with multiple goroutines to verify that `Allow()` and `Wait()` are thread-safe
- 💡 Replace the 40ms threshold in `TestRateLimiter_Wait` with a more precise assertion using `time.Since(start) >= 50*time.Millisecond` to ensure the wait duration matches the configured refill time
- 💡 Add a test case for `Allow()` returning false immediately after exhaustion and before any refill period has elapsed, to ensure no race condition exists between checking allowance and refilling
- 💡 Add a test that verifies `Wait()` returns immediately if a token is available, and only blocks when necessary
- ⚠️ Arbitrary timing thresholds in `TestRateLimiter_Wait` make test flaky and unreliable across environments
- ⚠️ No concurrency tests to detect race conditions in token management or allowance logic
- 🔗 These tests provide minimal confidence in the correctness of the rate limiter under real-world usage patterns
- 🔗 The lack of concurrency testing increases risk of subtle bugs in multi-threaded environments where the rate limiter might be used

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
- 🤖 The code has functional correctness but suffers from poor error handling, missing input validation, and a lack of proper state management in the Review method.
- 💡 Change early return in Review method from returning ReviewResult{} to returning an error indicating that provider or router is nil
- 💡 Fix joinModels function by changing loop condition to avoid prepending comma to first element: change `if i > 0` to `if i > 0 { result += "," }`
- 💡 Add proper error handling to each AI call in Review so that partial failures don't silently drop data
- 💡 Validate r.Status and r.Confidence in ToEvidence and ToPrescreenEvidence methods before constructing domain.Evidence
- 💡 Log or track token usage and AI call failures for observability and debugging purposes
- ⚠️ Early return with empty ReviewResult when provider/router is nil leads to silent failure instead of proper error propagation
- ⚠️ Improper comma joining in joinModels function causes malformed model strings
- 🔗 The Review method can silently fail and return incomplete results, affecting downstream systems that depend on consistent review outcomes
- 🔗 The ToEvidence methods embed raw user-provided data into domain.Evidence without validation, creating potential injection or malformed data risks

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
- 🤖 The code has functional correctness but suffers from poor error handling, missing input validation, and a lack of proper state management in the Review method.
- 💡 Change early return in Review method from returning ReviewResult{} to returning an error indicating that provider or router is nil
- 💡 Fix joinModels function by changing loop condition to avoid prepending comma to first element: change `if i > 0` to `if i > 0 { result += "," }`
- 💡 Add proper error handling to each AI call in Review so that partial failures don't silently drop data
- 💡 Validate r.Status and r.Confidence in ToEvidence and ToPrescreenEvidence methods before constructing domain.Evidence
- 💡 Log or track token usage and AI call failures for observability and debugging purposes
- ⚠️ Early return with empty ReviewResult when provider/router is nil leads to silent failure instead of proper error propagation
- ⚠️ Improper comma joining in joinModels function causes malformed model strings
- 🔗 The Review method can silently fail and return incomplete results, affecting downstream systems that depend on consistent review outcomes
- 🔗 The ToEvidence methods embed raw user-provided data into domain.Evidence without validation, creating potential injection or malformed data risks

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
- 🤖 The code has functional correctness but suffers from poor error handling, missing input validation, and a lack of proper state management in the Review method.
- 💡 Change early return in Review method from returning ReviewResult{} to returning an error indicating that provider or router is nil
- 💡 Fix joinModels function by changing loop condition to avoid prepending comma to first element: change `if i > 0` to `if i > 0 { result += "," }`
- 💡 Add proper error handling to each AI call in Review so that partial failures don't silently drop data
- 💡 Validate r.Status and r.Confidence in ToEvidence and ToPrescreenEvidence methods before constructing domain.Evidence
- 💡 Log or track token usage and AI call failures for observability and debugging purposes
- ⚠️ Early return with empty ReviewResult when provider/router is nil leads to silent failure instead of proper error propagation
- ⚠️ Improper comma joining in joinModels function causes malformed model strings
- 🔗 The Review method can silently fail and return incomplete results, affecting downstream systems that depend on consistent review outcomes
- 🔗 The ToEvidence methods embed raw user-provided data into domain.Evidence without validation, creating potential injection or malformed data risks

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
- 🤖 The code has functional correctness but suffers from poor error handling, missing input validation, and a lack of proper state management in the Review method.
- 💡 Change early return in Review method from returning ReviewResult{} to returning an error indicating that provider or router is nil
- 💡 Fix joinModels function by changing loop condition to avoid prepending comma to first element: change `if i > 0` to `if i > 0 { result += "," }`
- 💡 Add proper error handling to each AI call in Review so that partial failures don't silently drop data
- 💡 Validate r.Status and r.Confidence in ToEvidence and ToPrescreenEvidence methods before constructing domain.Evidence
- 💡 Log or track token usage and AI call failures for observability and debugging purposes
- ⚠️ Early return with empty ReviewResult when provider/router is nil leads to silent failure instead of proper error propagation
- ⚠️ Improper comma joining in joinModels function causes malformed model strings
- 🔗 The Review method can silently fail and return incomplete results, affecting downstream systems that depend on consistent review outcomes
- 🔗 The ToEvidence methods embed raw user-provided data into domain.Evidence without validation, creating potential injection or malformed data risks

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
- 🤖 The code has functional correctness but suffers from poor error handling, missing input validation, and a lack of proper state management in the Review method.
- 💡 Change early return in Review method from returning ReviewResult{} to returning an error indicating that provider or router is nil
- 💡 Fix joinModels function by changing loop condition to avoid prepending comma to first element: change `if i > 0` to `if i > 0 { result += "," }`
- 💡 Add proper error handling to each AI call in Review so that partial failures don't silently drop data
- 💡 Validate r.Status and r.Confidence in ToEvidence and ToPrescreenEvidence methods before constructing domain.Evidence
- 💡 Log or track token usage and AI call failures for observability and debugging purposes
- ⚠️ Early return with empty ReviewResult when provider/router is nil leads to silent failure instead of proper error propagation
- ⚠️ Improper comma joining in joinModels function causes malformed model strings
- 🔗 The Review method can silently fail and return incomplete results, affecting downstream systems that depend on consistent review outcomes
- 🔗 The ToEvidence methods embed raw user-provided data into domain.Evidence without validation, creating potential injection or malformed data risks

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
- 🤖 The code has functional correctness but suffers from poor error handling, missing input validation, and a lack of proper state management in the Review method.
- 💡 Change early return in Review method from returning ReviewResult{} to returning an error indicating that provider or router is nil
- 💡 Fix joinModels function by changing loop condition to avoid prepending comma to first element: change `if i > 0` to `if i > 0 { result += "," }`
- 💡 Add proper error handling to each AI call in Review so that partial failures don't silently drop data
- 💡 Validate r.Status and r.Confidence in ToEvidence and ToPrescreenEvidence methods before constructing domain.Evidence
- 💡 Log or track token usage and AI call failures for observability and debugging purposes
- ⚠️ Early return with empty ReviewResult when provider/router is nil leads to silent failure instead of proper error propagation
- ⚠️ Improper comma joining in joinModels function causes malformed model strings
- 🔗 The Review method can silently fail and return incomplete results, affecting downstream systems that depend on consistent review outcomes
- 🔗 The ToEvidence methods embed raw user-provided data into domain.Evidence without validation, creating potential injection or malformed data risks

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
- 🤖 The code has functional correctness but suffers from poor error handling, missing input validation, and a lack of proper state management in the Review method.
- 💡 Change early return in Review method from returning ReviewResult{} to returning an error indicating that provider or router is nil
- 💡 Fix joinModels function by changing loop condition to avoid prepending comma to first element: change `if i > 0` to `if i > 0 { result += "," }`
- 💡 Add proper error handling to each AI call in Review so that partial failures don't silently drop data
- 💡 Validate r.Status and r.Confidence in ToEvidence and ToPrescreenEvidence methods before constructing domain.Evidence
- 💡 Log or track token usage and AI call failures for observability and debugging purposes
- ⚠️ Early return with empty ReviewResult when provider/router is nil leads to silent failure instead of proper error propagation
- ⚠️ Improper comma joining in joinModels function causes malformed model strings
- 🔗 The Review method can silently fail and return incomplete results, affecting downstream systems that depend on consistent review outcomes
- 🔗 The ToEvidence methods embed raw user-provided data into domain.Evidence without validation, creating potential injection or malformed data risks

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
- 🤖 The code has functional correctness but suffers from poor error handling, missing input validation, and a lack of proper state management in the Review method.
- 💡 Change early return in Review method from returning ReviewResult{} to returning an error indicating that provider or router is nil
- 💡 Fix joinModels function by changing loop condition to avoid prepending comma to first element: change `if i > 0` to `if i > 0 { result += "," }`
- 💡 Add proper error handling to each AI call in Review so that partial failures don't silently drop data
- 💡 Validate r.Status and r.Confidence in ToEvidence and ToPrescreenEvidence methods before constructing domain.Evidence
- 💡 Log or track token usage and AI call failures for observability and debugging purposes
- ⚠️ Early return with empty ReviewResult when provider/router is nil leads to silent failure instead of proper error propagation
- ⚠️ Improper comma joining in joinModels function causes malformed model strings
- 🔗 The Review method can silently fail and return incomplete results, affecting downstream systems that depend on consistent review outcomes
- 🔗 The ToEvidence methods embed raw user-provided data into domain.Evidence without validation, creating potential injection or malformed data risks

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
- 🤖 The test suite has good coverage for core logic but lacks proper error handling in mock server and has a critical nil pointer dereference in one test case.
- 💡 Add error checking after `json.NewEncoder(w).Encode(resp)` in `mockServer` to ensure responses are valid JSON
- 💡 Fix `TestReviewer_NoProvider` to properly assert that nil provider results in an empty ReviewResult with no output fields
- 💡 Reset `call` counter in `mockServer` or make it a parameter to avoid test state leakage
- 💡 Validate that all mocked model names in tests match actual registered models or use a validation layer
- ⚠️ Uncaught JSON encoding errors in mock server may mask real encoding issues in production
- ⚠️ Nil provider test case does not align with actual behavior of Reviewer.Review() leading to false positives
- 🔗 Mock server implementation tightly couples test behavior to internal encoder behavior, reducing robustness
- 🔗 Test suite does not isolate state between tests due to shared variable `call`, increasing risk of flaky tests

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
- 🤖 The code is functionally correct but has a critical logic flaw in fallback handling and lacks proper error handling for invalid task types.
- 💡 Add logging or error reporting for unknown task types in `directAssignment` to aid debugging and detect configuration issues
- 💡 Validate that `r.assignments.Fallback` is not empty before returning it in `ModelFor` to prevent silent failures
- ⚠️ Fallback behavior is ambiguous when `r.assignments.Fallback` is empty or undefined, potentially leading to silent task routing failures
- ⚠️ No handling of unknown `TaskType` values in the switch statement, which can cause silent misrouting or runtime panics if new task types are introduced without updating the switch
- 🔗 This unit tightly couples to `domain.ModelAssignments` and assumes all task types are known at compile time, limiting extensibility
- 🔗 The fallback logic introduces a potential single point of failure if `Fallback` is misconfigured or missing

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
- 🤖 The code is functionally correct but has a critical logic flaw in fallback handling and lacks proper error handling for invalid task types.
- 💡 Add logging or error reporting for unknown task types in `directAssignment` to aid debugging and detect configuration issues
- 💡 Validate that `r.assignments.Fallback` is not empty before returning it in `ModelFor` to prevent silent failures
- ⚠️ Fallback behavior is ambiguous when `r.assignments.Fallback` is empty or undefined, potentially leading to silent task routing failures
- ⚠️ No handling of unknown `TaskType` values in the switch statement, which can cause silent misrouting or runtime panics if new task types are introduced without updating the switch
- 🔗 This unit tightly couples to `domain.ModelAssignments` and assumes all task types are known at compile time, limiting extensibility
- 🔗 The fallback logic introduces a potential single point of failure if `Fallback` is misconfigured or missing

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
- 🤖 The code is functionally correct but has a critical logic flaw in fallback handling and lacks proper error handling for invalid task types.
- 💡 Add logging or error reporting for unknown task types in `directAssignment` to aid debugging and detect configuration issues
- 💡 Validate that `r.assignments.Fallback` is not empty before returning it in `ModelFor` to prevent silent failures
- ⚠️ Fallback behavior is ambiguous when `r.assignments.Fallback` is empty or undefined, potentially leading to silent task routing failures
- ⚠️ No handling of unknown `TaskType` values in the switch statement, which can cause silent misrouting or runtime panics if new task types are introduced without updating the switch
- 🔗 This unit tightly couples to `domain.ModelAssignments` and assumes all task types are known at compile time, limiting extensibility
- 🔗 The fallback logic introduces a potential single point of failure if `Fallback` is misconfigured or missing

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
- 🤖 The code is functionally correct but has a critical logic flaw in fallback handling and lacks proper error handling for invalid task types.
- 💡 Add logging or error reporting for unknown task types in `directAssignment` to aid debugging and detect configuration issues
- 💡 Validate that `r.assignments.Fallback` is not empty before returning it in `ModelFor` to prevent silent failures
- ⚠️ Fallback behavior is ambiguous when `r.assignments.Fallback` is empty or undefined, potentially leading to silent task routing failures
- ⚠️ No handling of unknown `TaskType` values in the switch statement, which can cause silent misrouting or runtime panics if new task types are introduced without updating the switch
- 🔗 This unit tightly couples to `domain.ModelAssignments` and assumes all task types are known at compile time, limiting extensibility
- 🔗 The fallback logic introduces a potential single point of failure if `Fallback` is misconfigured or missing

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
- 🤖 The test suite for the Router component is well-structured but lacks coverage for edge cases and fallback behavior in the presence of partial assignments.
- 💡 Add a test case for `TestRouter_PartialAssignments` to verify that when some tasks have models and others do not, the fallback is correctly used for unassigned tasks.
- 💡 Add a test case for `TestRouter_InvalidTaskType` to ensure that invalid or unknown task types do not cause panics or unexpected behavior in the `ModelFor` method.
- ⚠️ Missing test coverage for partial model assignments where some tasks have models and others do not, potentially leading to incorrect fallback logic.
- ⚠️ No validation of edge cases such as empty or malformed task types, which may result in runtime panics or incorrect behavior.
- 🔗 This unit tests the Router's model selection logic, which is a core component of the agent's decision-making process. If the Router fails to correctly select or fallback on models, it can lead to incorrect task execution and downstream failures.
- 🔗 The Router's behavior directly impacts the agent's ability to route tasks to appropriate models. If not tested thoroughly, it can introduce subtle bugs that propagate through the system during task execution.

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
- 🤖 Well-structured Go data models with clear JSON tags, but missing validation and potential inconsistency in field naming.
- 💡 Add validation logic or custom JSON unmarshaling for `Confidence` to ensure it's between 0 and 1
- 💡 Add a check in `ScoringResponse` to ensure `Scores` is not nil before use, or make it a required field with proper initialization
- 💡 Standardize pluralization of fields across response types (e.g., use `Suggestions` consistently instead of mixing `Suggestion` and `Scores`)
- 💡 Introduce a type or enum for `Effort` to enforce valid values ('low', 'medium', 'high') and avoid string literals in code
- ⚠️ Missing input validation for fields like Confidence (should be between 0 and 1) and Scores (should not be nil)
- ⚠️ Inconsistent naming and structure across response types (e.g., pluralization of fields)
- 🔗 These structs form part of the API contract between steps in the agent pipeline, so any inconsistency or missing validation can propagate errors across system boundaries
- 🔗 The lack of type constraints on Effort (low/medium/high) introduces potential runtime errors or invalid data if not handled properly by downstream consumers

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
- 🤖 Well-structured Go data models with clear JSON tags, but missing validation and potential inconsistency in field naming.
- 💡 Add validation logic or custom JSON unmarshaling for `Confidence` to ensure it's between 0 and 1
- 💡 Add a check in `ScoringResponse` to ensure `Scores` is not nil before use, or make it a required field with proper initialization
- 💡 Standardize pluralization of fields across response types (e.g., use `Suggestions` consistently instead of mixing `Suggestion` and `Scores`)
- 💡 Introduce a type or enum for `Effort` to enforce valid values ('low', 'medium', 'high') and avoid string literals in code
- ⚠️ Missing input validation for fields like Confidence (should be between 0 and 1) and Scores (should not be nil)
- ⚠️ Inconsistent naming and structure across response types (e.g., pluralization of fields)
- 🔗 These structs form part of the API contract between steps in the agent pipeline, so any inconsistency or missing validation can propagate errors across system boundaries
- 🔗 The lack of type constraints on Effort (low/medium/high) introduces potential runtime errors or invalid data if not handled properly by downstream consumers

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
- 🤖 Well-structured Go data models with clear JSON tags, but missing validation and potential inconsistency in field naming.
- 💡 Add validation logic or custom JSON unmarshaling for `Confidence` to ensure it's between 0 and 1
- 💡 Add a check in `ScoringResponse` to ensure `Scores` is not nil before use, or make it a required field with proper initialization
- 💡 Standardize pluralization of fields across response types (e.g., use `Suggestions` consistently instead of mixing `Suggestion` and `Scores`)
- 💡 Introduce a type or enum for `Effort` to enforce valid values ('low', 'medium', 'high') and avoid string literals in code
- ⚠️ Missing input validation for fields like Confidence (should be between 0 and 1) and Scores (should not be nil)
- ⚠️ Inconsistent naming and structure across response types (e.g., pluralization of fields)
- 🔗 These structs form part of the API contract between steps in the agent pipeline, so any inconsistency or missing validation can propagate errors across system boundaries
- 🔗 The lack of type constraints on Effort (low/medium/high) introduces potential runtime errors or invalid data if not handled properly by downstream consumers

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
- 🤖 Well-structured Go data models with clear JSON tags, but missing validation and potential inconsistency in field naming.
- 💡 Add validation logic or custom JSON unmarshaling for `Confidence` to ensure it's between 0 and 1
- 💡 Add a check in `ScoringResponse` to ensure `Scores` is not nil before use, or make it a required field with proper initialization
- 💡 Standardize pluralization of fields across response types (e.g., use `Suggestions` consistently instead of mixing `Suggestion` and `Scores`)
- 💡 Introduce a type or enum for `Effort` to enforce valid values ('low', 'medium', 'high') and avoid string literals in code
- ⚠️ Missing input validation for fields like Confidence (should be between 0 and 1) and Scores (should not be nil)
- ⚠️ Inconsistent naming and structure across response types (e.g., pluralization of fields)
- 🔗 These structs form part of the API contract between steps in the agent pipeline, so any inconsistency or missing validation can propagate errors across system boundaries
- 🔗 The lack of type constraints on Effort (low/medium/high) introduces potential runtime errors or invalid data if not handled properly by downstream consumers

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
- 🤖 Well-structured Go data models with clear JSON tags, but missing validation and potential inconsistency in field naming.
- 💡 Add validation logic or custom JSON unmarshaling for `Confidence` to ensure it's between 0 and 1
- 💡 Add a check in `ScoringResponse` to ensure `Scores` is not nil before use, or make it a required field with proper initialization
- 💡 Standardize pluralization of fields across response types (e.g., use `Suggestions` consistently instead of mixing `Suggestion` and `Scores`)
- 💡 Introduce a type or enum for `Effort` to enforce valid values ('low', 'medium', 'high') and avoid string literals in code
- ⚠️ Missing input validation for fields like Confidence (should be between 0 and 1) and Scores (should not be nil)
- ⚠️ Inconsistent naming and structure across response types (e.g., pluralization of fields)
- 🔗 These structs form part of the API contract between steps in the agent pipeline, so any inconsistency or missing validation can propagate errors across system boundaries
- 🔗 The lack of type constraints on Effort (low/medium/high) introduces potential runtime errors or invalid data if not handled properly by downstream consumers

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
- 🤖 The test file validates JSON unmarshaling of schema types but lacks comprehensive validation and proper error handling for edge cases.
- 💡 Add table-driven tests with various valid and invalid JSON inputs to ensure robust parsing and error handling for each schema type.
- 💡 Validate that all fields in the structs (e.g. confidence, reason) are within expected ranges or constraints (e.g., 0 <= confidence <= 1).
- 💡 Add tests for edge cases like missing keys, incorrect types (e.g. string instead of float), and malformed JSON to ensure graceful failure.
- 💡 Use `t.Run()` to group related test cases, and consider parameterizing tests with different JSON payloads to improve maintainability.
- ⚠️ No validation of invalid JSON inputs (e.g. missing fields, wrong types) which could lead to runtime panics or incorrect behavior if the schema changes.
- ⚠️ Lack of input sanitization and validation for fields like `confidence` (e.g., negative or >1 values) which could cause downstream logic errors.
- 🔗 These tests only validate parsing, not correctness of the underlying data structures or business logic, so they provide limited assurance that the system behaves correctly in production.
- 🔗 The tests are tightly coupled to specific JSON structures and field names, increasing coupling between test code and implementation details.

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

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
- 🤖 The code implements a pipeline of stages for code review, but contains critical parsing and error handling issues that could lead to incorrect behavior or security vulnerabilities.
- 💡 Replace `extractJSON` with a robust JSON parser that validates structure and handles malformed input gracefully to prevent injection or incorrect parsing
- 💡 Implement strict validation of the parsed JSON in `prescreenStage.Execute` to ensure all required fields are present and correctly typed before using them
- 💡 Refactor `looseParseNeedsReview` to use more precise NLP techniques or regex patterns instead of simple substring matching to reduce false positives
- 💡 Add input validation for `input.Unit.ID` and `input.Unit.Type` to prevent runtime panics or incorrect processing
- 💡 Improve error handling in `scoringStage.Execute` to return errors instead of silently falling back to default scores when the JSON structure is invalid
- 💡 Add logging or metrics for cases where fallback logic is triggered to enable monitoring and debugging
- ⚠️ Insecure JSON parsing in prescreenStage.Execute that can lead to incorrect behavior or injection vulnerabilities due to lack of strict validation
- ⚠️ Loose natural language parsing in prescreenStage.Execute that is highly susceptible to false positives and can misclassify code quality
- 🔗 The prescreen stage's incorrect parsing logic can cause cascading failures in downstream stages by incorrectly flagging code for review or skipping necessary reviews
- 🔗 The scoring stage's fallback to default scores can mask real errors in model responses, leading to inaccurate quality assessments and potentially incorrect system behavior
- 🔗 The code's reliance on string truncation without clear user feedback can silently lose important context needed for accurate review

</details>

<a id="internal-agent-stage-deep-go-deepreviewresponse"></a>
<details>
<summary>DeepReviewResponse — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The deep review stage has functional correctness but lacks robust error handling, security checks on untrusted input, and is tightly coupled to specific model outputs.
- 💡 Implement strict validation and sanitization of JSON responses from the LLM before unmarshaling into `DeepReviewResponse`
- 💡 Add proper UTF-8 boundary checking when truncating code snippets to prevent invalid Unicode in prompts
- 💡 Add context deadline checks or timeout handling before calling `s.provider.Chat(...)`
- 💡 Validate that input fields like `input.Unit.ID`, `input.Unit.Type` are non-empty before constructing the prompt
- 💡 Make `MaxTokens` configurable or dynamically determined based on model capabilities and available context window
- ⚠️ Potential injection of malicious JSON from untrusted LLM responses
- ⚠️ Truncation of code snippet at invalid UTF-8 boundary leading to malformed prompts
- 🔗 This unit tightly couples to the specific structure of LLM responses and assumes consistent output format, making it brittle to model changes
- 🔗 The stage does not properly isolate its input/output from other stages, increasing failure propagation risk in multi-stage pipelines

</details>

<a id="internal-agent-stage-deep-go-execute"></a>
<details>
<summary>Execute — certified details</summary>

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
- 🤖 The deep review stage has functional correctness but lacks robust error handling, security checks on untrusted input, and is tightly coupled to specific model outputs.
- 💡 Implement strict validation and sanitization of JSON responses from the LLM before unmarshaling into `DeepReviewResponse`
- 💡 Add proper UTF-8 boundary checking when truncating code snippets to prevent invalid Unicode in prompts
- 💡 Add context deadline checks or timeout handling before calling `s.provider.Chat(...)`
- 💡 Validate that input fields like `input.Unit.ID`, `input.Unit.Type` are non-empty before constructing the prompt
- 💡 Make `MaxTokens` configurable or dynamically determined based on model capabilities and available context window
- ⚠️ Potential injection of malicious JSON from untrusted LLM responses
- ⚠️ Truncation of code snippet at invalid UTF-8 boundary leading to malformed prompts
- 🔗 This unit tightly couples to the specific structure of LLM responses and assumes consistent output format, making it brittle to model changes
- 🔗 The stage does not properly isolate its input/output from other stages, increasing failure propagation risk in multi-stage pipelines

</details>

<a id="internal-agent-stage-deep-go-formatdeepobservations"></a>
<details>
<summary>FormatDeepObservations — certified details</summary>

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
- 🤖 The deep review stage has functional correctness but lacks robust error handling, security checks on untrusted input, and is tightly coupled to specific model outputs.
- 💡 Implement strict validation and sanitization of JSON responses from the LLM before unmarshaling into `DeepReviewResponse`
- 💡 Add proper UTF-8 boundary checking when truncating code snippets to prevent invalid Unicode in prompts
- 💡 Add context deadline checks or timeout handling before calling `s.provider.Chat(...)`
- 💡 Validate that input fields like `input.Unit.ID`, `input.Unit.Type` are non-empty before constructing the prompt
- 💡 Make `MaxTokens` configurable or dynamically determined based on model capabilities and available context window
- ⚠️ Potential injection of malicious JSON from untrusted LLM responses
- ⚠️ Truncation of code snippet at invalid UTF-8 boundary leading to malformed prompts
- 🔗 This unit tightly couples to the specific structure of LLM responses and assumes consistent output format, making it brittle to model changes
- 🔗 The stage does not properly isolate its input/output from other stages, increasing failure propagation risk in multi-stage pipelines

</details>

<a id="internal-agent-stage-deep-go-formatreviewforrecord"></a>
<details>
<summary>FormatReviewForRecord — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The deep review stage has functional correctness but lacks robust error handling, security checks on untrusted input, and is tightly coupled to specific model outputs.
- 💡 Implement strict validation and sanitization of JSON responses from the LLM before unmarshaling into `DeepReviewResponse`
- 💡 Add proper UTF-8 boundary checking when truncating code snippets to prevent invalid Unicode in prompts
- 💡 Add context deadline checks or timeout handling before calling `s.provider.Chat(...)`
- 💡 Validate that input fields like `input.Unit.ID`, `input.Unit.Type` are non-empty before constructing the prompt
- 💡 Make `MaxTokens` configurable or dynamically determined based on model capabilities and available context window
- ⚠️ Potential injection of malicious JSON from untrusted LLM responses
- ⚠️ Truncation of code snippet at invalid UTF-8 boundary leading to malformed prompts
- 🔗 This unit tightly couples to the specific structure of LLM responses and assumes consistent output format, making it brittle to model changes
- 🔗 The stage does not properly isolate its input/output from other stages, increasing failure propagation risk in multi-stage pipelines

</details>

<a id="internal-agent-stage-deep-go-isdeepreview"></a>
<details>
<summary>IsDeepReview — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The deep review stage has functional correctness but lacks robust error handling, security checks on untrusted input, and is tightly coupled to specific model outputs.
- 💡 Implement strict validation and sanitization of JSON responses from the LLM before unmarshaling into `DeepReviewResponse`
- 💡 Add proper UTF-8 boundary checking when truncating code snippets to prevent invalid Unicode in prompts
- 💡 Add context deadline checks or timeout handling before calling `s.provider.Chat(...)`
- 💡 Validate that input fields like `input.Unit.ID`, `input.Unit.Type` are non-empty before constructing the prompt
- 💡 Make `MaxTokens` configurable or dynamically determined based on model capabilities and available context window
- ⚠️ Potential injection of malicious JSON from untrusted LLM responses
- ⚠️ Truncation of code snippet at invalid UTF-8 boundary leading to malformed prompts
- 🔗 This unit tightly couples to the specific structure of LLM responses and assumes consistent output format, making it brittle to model changes
- 🔗 The stage does not properly isolate its input/output from other stages, increasing failure propagation risk in multi-stage pipelines

</details>

<a id="internal-agent-stage-deep-go-name"></a>
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
- 🤖 The deep review stage has functional correctness but lacks robust error handling, security checks on untrusted input, and is tightly coupled to specific model outputs.
- 💡 Implement strict validation and sanitization of JSON responses from the LLM before unmarshaling into `DeepReviewResponse`
- 💡 Add proper UTF-8 boundary checking when truncating code snippets to prevent invalid Unicode in prompts
- 💡 Add context deadline checks or timeout handling before calling `s.provider.Chat(...)`
- 💡 Validate that input fields like `input.Unit.ID`, `input.Unit.Type` are non-empty before constructing the prompt
- 💡 Make `MaxTokens` configurable or dynamically determined based on model capabilities and available context window
- ⚠️ Potential injection of malicious JSON from untrusted LLM responses
- ⚠️ Truncation of code snippet at invalid UTF-8 boundary leading to malformed prompts
- 🔗 This unit tightly couples to the specific structure of LLM responses and assumes consistent output format, making it brittle to model changes
- 🔗 The stage does not properly isolate its input/output from other stages, increasing failure propagation risk in multi-stage pipelines

</details>

<a id="internal-agent-stage-deep-go-newdeepreviewstage"></a>
<details>
<summary>NewDeepReviewStage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The deep review stage has functional correctness but lacks robust error handling, security checks on untrusted input, and is tightly coupled to specific model outputs.
- 💡 Implement strict validation and sanitization of JSON responses from the LLM before unmarshaling into `DeepReviewResponse`
- 💡 Add proper UTF-8 boundary checking when truncating code snippets to prevent invalid Unicode in prompts
- 💡 Add context deadline checks or timeout handling before calling `s.provider.Chat(...)`
- 💡 Validate that input fields like `input.Unit.ID`, `input.Unit.Type` are non-empty before constructing the prompt
- 💡 Make `MaxTokens` configurable or dynamically determined based on model capabilities and available context window
- ⚠️ Potential injection of malicious JSON from untrusted LLM responses
- ⚠️ Truncation of code snippet at invalid UTF-8 boundary leading to malformed prompts
- 🔗 This unit tightly couples to the specific structure of LLM responses and assumes consistent output format, making it brittle to model changes
- 🔗 The stage does not properly isolate its input/output from other stages, increasing failure propagation risk in multi-stage pipelines

</details>

<a id="internal-agent-stage-deep-go-todeepevidence"></a>
<details>
<summary>ToDeepEvidence — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The deep review stage has functional correctness but lacks robust error handling, security checks on untrusted input, and is tightly coupled to specific model outputs.
- 💡 Implement strict validation and sanitization of JSON responses from the LLM before unmarshaling into `DeepReviewResponse`
- 💡 Add proper UTF-8 boundary checking when truncating code snippets to prevent invalid Unicode in prompts
- 💡 Add context deadline checks or timeout handling before calling `s.provider.Chat(...)`
- 💡 Validate that input fields like `input.Unit.ID`, `input.Unit.Type` are non-empty before constructing the prompt
- 💡 Make `MaxTokens` configurable or dynamically determined based on model capabilities and available context window
- ⚠️ Potential injection of malicious JSON from untrusted LLM responses
- ⚠️ Truncation of code snippet at invalid UTF-8 boundary leading to malformed prompts
- 🔗 This unit tightly couples to the specific structure of LLM responses and assumes consistent output format, making it brittle to model changes
- 🔗 The stage does not properly isolate its input/output from other stages, increasing failure propagation risk in multi-stage pipelines

</details>

<a id="internal-agent-stage-deep-go-deepreviewstage"></a>
<details>
<summary>deepReviewStage — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The deep review stage has functional correctness but lacks robust error handling, security checks on untrusted input, and is tightly coupled to specific model outputs.
- 💡 Implement strict validation and sanitization of JSON responses from the LLM before unmarshaling into `DeepReviewResponse`
- 💡 Add proper UTF-8 boundary checking when truncating code snippets to prevent invalid Unicode in prompts
- 💡 Add context deadline checks or timeout handling before calling `s.provider.Chat(...)`
- 💡 Validate that input fields like `input.Unit.ID`, `input.Unit.Type` are non-empty before constructing the prompt
- 💡 Make `MaxTokens` configurable or dynamically determined based on model capabilities and available context window
- ⚠️ Potential injection of malicious JSON from untrusted LLM responses
- ⚠️ Truncation of code snippet at invalid UTF-8 boundary leading to malformed prompts
- 🔗 This unit tightly couples to the specific structure of LLM responses and assumes consistent output format, making it brittle to model changes
- 🔗 The stage does not properly isolate its input/output from other stages, increasing failure propagation risk in multi-stage pipelines

</details>

<a id="internal-agent-stage-deep-go-extractfirstsentence"></a>
<details>
<summary>extractFirstSentence — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The deep review stage has functional correctness but lacks robust error handling, security checks on untrusted input, and is tightly coupled to specific model outputs.
- 💡 Implement strict validation and sanitization of JSON responses from the LLM before unmarshaling into `DeepReviewResponse`
- 💡 Add proper UTF-8 boundary checking when truncating code snippets to prevent invalid Unicode in prompts
- 💡 Add context deadline checks or timeout handling before calling `s.provider.Chat(...)`
- 💡 Validate that input fields like `input.Unit.ID`, `input.Unit.Type` are non-empty before constructing the prompt
- 💡 Make `MaxTokens` configurable or dynamically determined based on model capabilities and available context window
- ⚠️ Potential injection of malicious JSON from untrusted LLM responses
- ⚠️ Truncation of code snippet at invalid UTF-8 boundary leading to malformed prompts
- 🔗 This unit tightly couples to the specific structure of LLM responses and assumes consistent output format, making it brittle to model changes
- 🔗 The stage does not properly isolate its input/output from other stages, increasing failure propagation risk in multi-stage pipelines

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
- 🤖 The test suite provides good coverage for core agent stages and pipeline behavior, but has several gaps in error handling, test isolation, and correctness assumptions.
- 💡 Fix TestPrescreenStage_MalformedJSON_FallsThrough to use a response that actually contains 'no review' or similar trigger phrase
- 💡 Add proper synchronization to conditionalProvider to prevent race conditions in concurrent tests
- 💡 Implement a more comprehensive test for fallback scoring logic in TestScoringStage_MalformedJSON_FallbackScores
- 💡 Add assertions to verify that pipeline stages receive correct inputs in TestPipeline_StandardStrategy
- 💡 Validate token usage in TestCoordinator_RespectsTokenBudget by asserting specific token counts rather than just checking if reviewed is false
- ⚠️ Hardcoded string matching in TestPrescreenStage_MalformedJSON_FallsThrough does not actually trigger the intended fallback logic
- ⚠️ Race condition in conditionalProvider used in TestCircuitBreaker_ClosesAfterSuccess due to unsynchronized call count increment
- 🔗 The test suite does not adequately validate error propagation from the pipeline, which could lead to unhandled exceptions in production
- 🔗 The test suite assumes deterministic behavior from sequenceProvider and mockProvider, which may not reflect real-world concurrent usage

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
- 🤖 The function handles AI suggestion generation with graceful failure handling but lacks input validation, has a hardcoded prompt template, and does not expose token usage or model info for logging/tracing.
- 💡 Validate `summary.UnitCount` and other numeric fields before including them in the prompt to prevent malformed or misleading prompts.
- 💡 Add a field to `ScanSuggestion` to capture the original error or status code from the provider call, enabling better observability and debugging.
- 💡 Extract `buildSuggestPrompt` into a configurable template or function that can be overridden or tested independently.
- 💡 Log or return error context (e.g., provider name, error type) when returning empty `ScanSuggestion` to aid in troubleshooting.
- ⚠️ Hardcoded prompt template makes it difficult to customize or localize suggestions without code changes.
- ⚠️ No input validation on `summary.UnitCount` or other fields can lead to malformed prompts or unexpected behavior.
- 🔗 This unit tightly couples to a specific `Provider` interface and assumes a particular response structure from the LLM, increasing system fragility if that contract changes.
- 🔗 The function returns a zero-value `ScanSuggestion` on any error, which prevents proper error propagation or retry logic in higher layers.

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
- 🤖 The function handles AI suggestion generation with graceful failure handling but lacks input validation, has a hardcoded prompt template, and does not expose token usage or model info for logging/tracing.
- 💡 Validate `summary.UnitCount` and other numeric fields before including them in the prompt to prevent malformed or misleading prompts.
- 💡 Add a field to `ScanSuggestion` to capture the original error or status code from the provider call, enabling better observability and debugging.
- 💡 Extract `buildSuggestPrompt` into a configurable template or function that can be overridden or tested independently.
- 💡 Log or return error context (e.g., provider name, error type) when returning empty `ScanSuggestion` to aid in troubleshooting.
- ⚠️ Hardcoded prompt template makes it difficult to customize or localize suggestions without code changes.
- ⚠️ No input validation on `summary.UnitCount` or other fields can lead to malformed prompts or unexpected behavior.
- 🔗 This unit tightly couples to a specific `Provider` interface and assumes a particular response structure from the LLM, increasing system fragility if that contract changes.
- 🔗 The function returns a zero-value `ScanSuggestion` on any error, which prevents proper error propagation or retry logic in higher layers.

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
- 🤖 The function handles AI suggestion generation with graceful failure handling but lacks input validation, has a hardcoded prompt template, and does not expose token usage or model info for logging/tracing.
- 💡 Validate `summary.UnitCount` and other numeric fields before including them in the prompt to prevent malformed or misleading prompts.
- 💡 Add a field to `ScanSuggestion` to capture the original error or status code from the provider call, enabling better observability and debugging.
- 💡 Extract `buildSuggestPrompt` into a configurable template or function that can be overridden or tested independently.
- 💡 Log or return error context (e.g., provider name, error type) when returning empty `ScanSuggestion` to aid in troubleshooting.
- ⚠️ Hardcoded prompt template makes it difficult to customize or localize suggestions without code changes.
- ⚠️ No input validation on `summary.UnitCount` or other fields can lead to malformed prompts or unexpected behavior.
- 🔗 This unit tightly couples to a specific `Provider` interface and assumes a particular response structure from the LLM, increasing system fragility if that contract changes.
- 🔗 The function returns a zero-value `ScanSuggestion` on any error, which prevents proper error propagation or retry logic in higher layers.

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
- 🤖 The function handles AI suggestion generation with graceful failure handling but lacks input validation, has a hardcoded prompt template, and does not expose token usage or model info for logging/tracing.
- 💡 Validate `summary.UnitCount` and other numeric fields before including them in the prompt to prevent malformed or misleading prompts.
- 💡 Add a field to `ScanSuggestion` to capture the original error or status code from the provider call, enabling better observability and debugging.
- 💡 Extract `buildSuggestPrompt` into a configurable template or function that can be overridden or tested independently.
- 💡 Log or return error context (e.g., provider name, error type) when returning empty `ScanSuggestion` to aid in troubleshooting.
- ⚠️ Hardcoded prompt template makes it difficult to customize or localize suggestions without code changes.
- ⚠️ No input validation on `summary.UnitCount` or other fields can lead to malformed prompts or unexpected behavior.
- 🔗 This unit tightly couples to a specific `Provider` interface and assumes a particular response structure from the LLM, increasing system fragility if that contract changes.
- 🔗 The function returns a zero-value `ScanSuggestion` on any error, which prevents proper error propagation or retry logic in higher layers.

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
- 🤖 The test suite for SuggestForRepo is functionally correct but lacks comprehensive coverage and has weak assertions that don't validate the actual content or structure of suggestions.
- 💡 Add assertions to verify that suggestion content includes relevant language-specific recommendations based on the repo summary
- 💡 Implement more comprehensive test cases covering multiple choices, nil responses, and malformed chat responses to improve robustness
- ⚠️ Lack of content validation in success case may mask incorrect suggestion formatting or missing language-specific recommendations
- ⚠️ Mock provider does not simulate realistic error conditions or partial response scenarios that could occur in production
- 🔗 This test file only affects the unit under test (SuggestForRepo) but does not adequately cover its dependencies or integration points
- 🔗 The test suite provides minimal confidence in the correctness of suggestion generation logic due to incomplete coverage

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
- 🤖 Well-structured Go types for LLM chat interactions with minor correctness and maintainability gaps.
- 💡 Replace `any` with a concrete type or interface for `ResponseFormat.Schema` to improve type safety and enable compile-time validation.
- 💡 Add input validation for `ChatRequest.Temperature` (should be 0–2) and `MaxTokens` (should be positive integer).
- 💡 Add a check in `ChatResponse.Content()` to ensure `r.Choices[0].Message.Content` is not empty or malformed before returning it.
- 💡 Consider adding unit tests for `TaskType.String()` with unknown enum values to ensure fallback behavior works as expected.
- ⚠️ Unvalidated optional fields in ChatRequest (e.g., Temperature, MaxTokens) can cause unexpected behavior or invalid API calls.
- ⚠️ Use of `any` for ResponseFormat.Schema introduces runtime type safety risks and makes it harder to enforce schema validation.
- 🔗 This unit tightly couples with OpenAI-compatible APIs and assumes specific JSON field names, limiting flexibility if the API changes.
- 🔗 The use of `any` in ResponseFormat increases coupling to unstructured data, making downstream consumers more fragile.

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
- 🤖 Well-structured Go types for LLM chat interactions with minor correctness and maintainability gaps.
- 💡 Replace `any` with a concrete type or interface for `ResponseFormat.Schema` to improve type safety and enable compile-time validation.
- 💡 Add input validation for `ChatRequest.Temperature` (should be 0–2) and `MaxTokens` (should be positive integer).
- 💡 Add a check in `ChatResponse.Content()` to ensure `r.Choices[0].Message.Content` is not empty or malformed before returning it.
- 💡 Consider adding unit tests for `TaskType.String()` with unknown enum values to ensure fallback behavior works as expected.
- ⚠️ Unvalidated optional fields in ChatRequest (e.g., Temperature, MaxTokens) can cause unexpected behavior or invalid API calls.
- ⚠️ Use of `any` for ResponseFormat.Schema introduces runtime type safety risks and makes it harder to enforce schema validation.
- 🔗 This unit tightly couples with OpenAI-compatible APIs and assumes specific JSON field names, limiting flexibility if the API changes.
- 🔗 The use of `any` in ResponseFormat increases coupling to unstructured data, making downstream consumers more fragile.

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
- 🤖 Well-structured Go types for LLM chat interactions with minor correctness and maintainability gaps.
- 💡 Replace `any` with a concrete type or interface for `ResponseFormat.Schema` to improve type safety and enable compile-time validation.
- 💡 Add input validation for `ChatRequest.Temperature` (should be 0–2) and `MaxTokens` (should be positive integer).
- 💡 Add a check in `ChatResponse.Content()` to ensure `r.Choices[0].Message.Content` is not empty or malformed before returning it.
- 💡 Consider adding unit tests for `TaskType.String()` with unknown enum values to ensure fallback behavior works as expected.
- ⚠️ Unvalidated optional fields in ChatRequest (e.g., Temperature, MaxTokens) can cause unexpected behavior or invalid API calls.
- ⚠️ Use of `any` for ResponseFormat.Schema introduces runtime type safety risks and makes it harder to enforce schema validation.
- 🔗 This unit tightly couples with OpenAI-compatible APIs and assumes specific JSON field names, limiting flexibility if the API changes.
- 🔗 The use of `any` in ResponseFormat increases coupling to unstructured data, making downstream consumers more fragile.

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
- 🤖 Well-structured Go types for LLM chat interactions with minor correctness and maintainability gaps.
- 💡 Replace `any` with a concrete type or interface for `ResponseFormat.Schema` to improve type safety and enable compile-time validation.
- 💡 Add input validation for `ChatRequest.Temperature` (should be 0–2) and `MaxTokens` (should be positive integer).
- 💡 Add a check in `ChatResponse.Content()` to ensure `r.Choices[0].Message.Content` is not empty or malformed before returning it.
- 💡 Consider adding unit tests for `TaskType.String()` with unknown enum values to ensure fallback behavior works as expected.
- ⚠️ Unvalidated optional fields in ChatRequest (e.g., Temperature, MaxTokens) can cause unexpected behavior or invalid API calls.
- ⚠️ Use of `any` for ResponseFormat.Schema introduces runtime type safety risks and makes it harder to enforce schema validation.
- 🔗 This unit tightly couples with OpenAI-compatible APIs and assumes specific JSON field names, limiting flexibility if the API changes.
- 🔗 The use of `any` in ResponseFormat increases coupling to unstructured data, making downstream consumers more fragile.

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
- 🤖 Well-structured Go types for LLM chat interactions with minor correctness and maintainability gaps.
- 💡 Replace `any` with a concrete type or interface for `ResponseFormat.Schema` to improve type safety and enable compile-time validation.
- 💡 Add input validation for `ChatRequest.Temperature` (should be 0–2) and `MaxTokens` (should be positive integer).
- 💡 Add a check in `ChatResponse.Content()` to ensure `r.Choices[0].Message.Content` is not empty or malformed before returning it.
- 💡 Consider adding unit tests for `TaskType.String()` with unknown enum values to ensure fallback behavior works as expected.
- ⚠️ Unvalidated optional fields in ChatRequest (e.g., Temperature, MaxTokens) can cause unexpected behavior or invalid API calls.
- ⚠️ Use of `any` for ResponseFormat.Schema introduces runtime type safety risks and makes it harder to enforce schema validation.
- 🔗 This unit tightly couples with OpenAI-compatible APIs and assumes specific JSON field names, limiting flexibility if the API changes.
- 🔗 The use of `any` in ResponseFormat increases coupling to unstructured data, making downstream consumers more fragile.

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
- 🤖 Well-structured Go types for LLM chat interactions with minor correctness and maintainability gaps.
- 💡 Replace `any` with a concrete type or interface for `ResponseFormat.Schema` to improve type safety and enable compile-time validation.
- 💡 Add input validation for `ChatRequest.Temperature` (should be 0–2) and `MaxTokens` (should be positive integer).
- 💡 Add a check in `ChatResponse.Content()` to ensure `r.Choices[0].Message.Content` is not empty or malformed before returning it.
- 💡 Consider adding unit tests for `TaskType.String()` with unknown enum values to ensure fallback behavior works as expected.
- ⚠️ Unvalidated optional fields in ChatRequest (e.g., Temperature, MaxTokens) can cause unexpected behavior or invalid API calls.
- ⚠️ Use of `any` for ResponseFormat.Schema introduces runtime type safety risks and makes it harder to enforce schema validation.
- 🔗 This unit tightly couples with OpenAI-compatible APIs and assumes specific JSON field names, limiting flexibility if the API changes.
- 🔗 The use of `any` in ResponseFormat increases coupling to unstructured data, making downstream consumers more fragile.

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
- 🤖 Well-structured Go types for LLM chat interactions with minor correctness and maintainability gaps.
- 💡 Replace `any` with a concrete type or interface for `ResponseFormat.Schema` to improve type safety and enable compile-time validation.
- 💡 Add input validation for `ChatRequest.Temperature` (should be 0–2) and `MaxTokens` (should be positive integer).
- 💡 Add a check in `ChatResponse.Content()` to ensure `r.Choices[0].Message.Content` is not empty or malformed before returning it.
- 💡 Consider adding unit tests for `TaskType.String()` with unknown enum values to ensure fallback behavior works as expected.
- ⚠️ Unvalidated optional fields in ChatRequest (e.g., Temperature, MaxTokens) can cause unexpected behavior or invalid API calls.
- ⚠️ Use of `any` for ResponseFormat.Schema introduces runtime type safety risks and makes it harder to enforce schema validation.
- 🔗 This unit tightly couples with OpenAI-compatible APIs and assumes specific JSON field names, limiting flexibility if the API changes.
- 🔗 The use of `any` in ResponseFormat increases coupling to unstructured data, making downstream consumers more fragile.

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
- 🤖 Well-structured Go types for LLM chat interactions with minor correctness and maintainability gaps.
- 💡 Replace `any` with a concrete type or interface for `ResponseFormat.Schema` to improve type safety and enable compile-time validation.
- 💡 Add input validation for `ChatRequest.Temperature` (should be 0–2) and `MaxTokens` (should be positive integer).
- 💡 Add a check in `ChatResponse.Content()` to ensure `r.Choices[0].Message.Content` is not empty or malformed before returning it.
- 💡 Consider adding unit tests for `TaskType.String()` with unknown enum values to ensure fallback behavior works as expected.
- ⚠️ Unvalidated optional fields in ChatRequest (e.g., Temperature, MaxTokens) can cause unexpected behavior or invalid API calls.
- ⚠️ Use of `any` for ResponseFormat.Schema introduces runtime type safety risks and makes it harder to enforce schema validation.
- 🔗 This unit tightly couples with OpenAI-compatible APIs and assumes specific JSON field names, limiting flexibility if the API changes.
- 🔗 The use of `any` in ResponseFormat increases coupling to unstructured data, making downstream consumers more fragile.

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
- 🤖 Well-structured Go types for LLM chat interactions with minor correctness and maintainability gaps.
- 💡 Replace `any` with a concrete type or interface for `ResponseFormat.Schema` to improve type safety and enable compile-time validation.
- 💡 Add input validation for `ChatRequest.Temperature` (should be 0–2) and `MaxTokens` (should be positive integer).
- 💡 Add a check in `ChatResponse.Content()` to ensure `r.Choices[0].Message.Content` is not empty or malformed before returning it.
- 💡 Consider adding unit tests for `TaskType.String()` with unknown enum values to ensure fallback behavior works as expected.
- ⚠️ Unvalidated optional fields in ChatRequest (e.g., Temperature, MaxTokens) can cause unexpected behavior or invalid API calls.
- ⚠️ Use of `any` for ResponseFormat.Schema introduces runtime type safety risks and makes it harder to enforce schema validation.
- 🔗 This unit tightly couples with OpenAI-compatible APIs and assumes specific JSON field names, limiting flexibility if the API changes.
- 🔗 The use of `any` in ResponseFormat increases coupling to unstructured data, making downstream consumers more fragile.

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
- 🤖 Well-structured Go types for LLM chat interactions with minor correctness and maintainability gaps.
- 💡 Replace `any` with a concrete type or interface for `ResponseFormat.Schema` to improve type safety and enable compile-time validation.
- 💡 Add input validation for `ChatRequest.Temperature` (should be 0–2) and `MaxTokens` (should be positive integer).
- 💡 Add a check in `ChatResponse.Content()` to ensure `r.Choices[0].Message.Content` is not empty or malformed before returning it.
- 💡 Consider adding unit tests for `TaskType.String()` with unknown enum values to ensure fallback behavior works as expected.
- ⚠️ Unvalidated optional fields in ChatRequest (e.g., Temperature, MaxTokens) can cause unexpected behavior or invalid API calls.
- ⚠️ Use of `any` for ResponseFormat.Schema introduces runtime type safety risks and makes it harder to enforce schema validation.
- 🔗 This unit tightly couples with OpenAI-compatible APIs and assumes specific JSON field names, limiting flexibility if the API changes.
- 🔗 The use of `any` in ResponseFormat increases coupling to unstructured data, making downstream consumers more fragile.

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
- 🤖 The test suite is well-structured and covers basic functionality, but lacks comprehensive edge case handling and has minor structural issues in test assertions.
- 💡 Replace `parsed["messages"].([]any)` with a more robust type assertion or use `json.RawMessage` to inspect nested structures safely.
- 💡 Add a test case for invalid or unknown `TaskType` values to ensure that `String()` handles them gracefully (e.g., returning a default or panic).
- 💡 Add tests for `ChatResponse` edge cases like missing fields, malformed JSON, or empty choices array.
- 💡 Ensure that `TestChatResponse_Empty` also verifies that all fields of `ChatResponse` are initialized to zero values, not just `Content()`.
- ⚠️ Fragile JSON parsing in `TestChatRequest_JSON` due to type assertion on `parsed["messages"]` as `[]any` without checking type compatibility.
- ⚠️ Potential silent failure if `TaskType` values are modified or extended, since there's no validation that all valid enum values are covered in tests.
- 🔗 This file is a test file and does not directly affect runtime behavior, but poor test coverage increases risk of regressions in the agent package.
- 🔗 The tests rely on hardcoded string values and assumptions about JSON structure that may not hold in future versions or with different input formats.

</details>

### `internal/config/` (22 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`Load`](reports/internal-config-loader-go-load.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`LoadFile`](reports/internal-config-loader-go-loadfile.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`LoadFromDir`](reports/internal-config-loader-go-loadfromdir.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`rawAgent`](reports/internal-config-loader-go-rawagent.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`rawConfig`](reports/internal-config-loader-go-rawconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`validate`](reports/internal-config-loader-go-validate.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`loader_test.go`](reports/internal-config-loader-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`FilterPolicyPacks`](reports/internal-config-matcher-go-filterpolicypacks.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`NewPolicyMatcher`](reports/internal-config-matcher-go-newpolicymatcher.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`LoadPolicyPack`](reports/internal-config-policy-go-loadpolicypack.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`LoadPolicyPacks`](reports/internal-config-policy-go-loadpolicypacks.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`parseDimension`](reports/internal-config-policy-go-parsedimension.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`parsePolicyPack`](reports/internal-config-policy-go-parsepolicypack.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`parseSeverity`](reports/internal-config-policy-go-parseseverity.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`rawPolicyPack`](reports/internal-config-policy-go-rawpolicypack.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`rawPolicyRule`](reports/internal-config-policy-go-rawpolicyrule.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`policy_test.go`](reports/internal-config-policy-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Error`](reports/internal-config-validator-go-error.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ValidateConfig`](reports/internal-config-validator-go-validateconfig.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ValidatePolicyPack`](reports/internal-config-validator-go-validatepolicypack.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ValidationError`](reports/internal-config-validator-go-validationerror.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`validator_test.go`](reports/internal-config-validator-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The config loader has correctness issues with agent enabled detection and YAML parsing, lacks proper error handling for malformed input, and introduces potential logic gaps in config validation.
- 💡 Replace the double YAML unmarshaling with a single pass that uses yaml.Node to inspect explicit field presence for agent.enabled
- 💡 Fix the agent config application logic by checking raw.Agent.Enabled directly and applying the full agent config when it's explicitly set
- 💡 Ensure that all struct fields in domain.Config are properly initialized before applying overrides and avoid blind assignments like `cfg.Schedule = raw.Schedule`
- ⚠️ Incorrect agent enabled detection due to double YAML unmarshaling and flawed conditional logic
- ⚠️ Potential panic or incorrect config application if raw fields are not properly initialized
- 🔗 This unit tightly couples to YAML parsing and assumes specific struct layouts in domain.Config, increasing fragility when domain changes
- 🔗 The double YAML unmarshaling introduces performance overhead and increases surface area for parsing errors

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
- 🤖 The config loader has correctness issues with agent enabled detection and YAML parsing, lacks proper error handling for malformed input, and introduces potential logic gaps in config validation.
- 💡 Replace the double YAML unmarshaling with a single pass that uses yaml.Node to inspect explicit field presence for agent.enabled
- 💡 Fix the agent config application logic by checking raw.Agent.Enabled directly and applying the full agent config when it's explicitly set
- 💡 Ensure that all struct fields in domain.Config are properly initialized before applying overrides and avoid blind assignments like `cfg.Schedule = raw.Schedule`
- ⚠️ Incorrect agent enabled detection due to double YAML unmarshaling and flawed conditional logic
- ⚠️ Potential panic or incorrect config application if raw fields are not properly initialized
- 🔗 This unit tightly couples to YAML parsing and assumes specific struct layouts in domain.Config, increasing fragility when domain changes
- 🔗 The double YAML unmarshaling introduces performance overhead and increases surface area for parsing errors

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
- 🤖 The config loader has correctness issues with agent enabled detection and YAML parsing, lacks proper error handling for malformed input, and introduces potential logic gaps in config validation.
- 💡 Replace the double YAML unmarshaling with a single pass that uses yaml.Node to inspect explicit field presence for agent.enabled
- 💡 Fix the agent config application logic by checking raw.Agent.Enabled directly and applying the full agent config when it's explicitly set
- 💡 Ensure that all struct fields in domain.Config are properly initialized before applying overrides and avoid blind assignments like `cfg.Schedule = raw.Schedule`
- ⚠️ Incorrect agent enabled detection due to double YAML unmarshaling and flawed conditional logic
- ⚠️ Potential panic or incorrect config application if raw fields are not properly initialized
- 🔗 This unit tightly couples to YAML parsing and assumes specific struct layouts in domain.Config, increasing fragility when domain changes
- 🔗 The double YAML unmarshaling introduces performance overhead and increases surface area for parsing errors

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
- 🤖 The config loader has correctness issues with agent enabled detection and YAML parsing, lacks proper error handling for malformed input, and introduces potential logic gaps in config validation.
- 💡 Replace the double YAML unmarshaling with a single pass that uses yaml.Node to inspect explicit field presence for agent.enabled
- 💡 Fix the agent config application logic by checking raw.Agent.Enabled directly and applying the full agent config when it's explicitly set
- 💡 Ensure that all struct fields in domain.Config are properly initialized before applying overrides and avoid blind assignments like `cfg.Schedule = raw.Schedule`
- ⚠️ Incorrect agent enabled detection due to double YAML unmarshaling and flawed conditional logic
- ⚠️ Potential panic or incorrect config application if raw fields are not properly initialized
- 🔗 This unit tightly couples to YAML parsing and assumes specific struct layouts in domain.Config, increasing fragility when domain changes
- 🔗 The double YAML unmarshaling introduces performance overhead and increases surface area for parsing errors

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
- 🤖 The config loader has correctness issues with agent enabled detection and YAML parsing, lacks proper error handling for malformed input, and introduces potential logic gaps in config validation.
- 💡 Replace the double YAML unmarshaling with a single pass that uses yaml.Node to inspect explicit field presence for agent.enabled
- 💡 Fix the agent config application logic by checking raw.Agent.Enabled directly and applying the full agent config when it's explicitly set
- 💡 Ensure that all struct fields in domain.Config are properly initialized before applying overrides and avoid blind assignments like `cfg.Schedule = raw.Schedule`
- ⚠️ Incorrect agent enabled detection due to double YAML unmarshaling and flawed conditional logic
- ⚠️ Potential panic or incorrect config application if raw fields are not properly initialized
- 🔗 This unit tightly couples to YAML parsing and assumes specific struct layouts in domain.Config, increasing fragility when domain changes
- 🔗 The double YAML unmarshaling introduces performance overhead and increases surface area for parsing errors

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
- 🤖 The config loader has correctness issues with agent enabled detection and YAML parsing, lacks proper error handling for malformed input, and introduces potential logic gaps in config validation.
- 💡 Replace the double YAML unmarshaling with a single pass that uses yaml.Node to inspect explicit field presence for agent.enabled
- 💡 Fix the agent config application logic by checking raw.Agent.Enabled directly and applying the full agent config when it's explicitly set
- 💡 Ensure that all struct fields in domain.Config are properly initialized before applying overrides and avoid blind assignments like `cfg.Schedule = raw.Schedule`
- ⚠️ Incorrect agent enabled detection due to double YAML unmarshaling and flawed conditional logic
- ⚠️ Potential panic or incorrect config application if raw fields are not properly initialized
- 🔗 This unit tightly couples to YAML parsing and assumes specific struct layouts in domain.Config, increasing fragility when domain changes
- 🔗 The double YAML unmarshaling introduces performance overhead and increases surface area for parsing errors

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
- 🤖 The test suite for config loading is comprehensive but lacks proper error validation and uses hardcoded paths that reduce portability.
- 💡 Replace hardcoded `testdataPath` with a utility that resolves paths relative to the test file's location using `t.Cleanup` or `filepath.Abs` for robustness
- 💡 Add assertions on specific error messages in `TestLoadConfig_Invalid` to ensure validation logic is correctly implemented
- 💡 Add assertions for all default fields in `TestLoadConfig_FromBytes` to ensure proper defaults are applied
- 💡 Add a test for `TestLoadConfig_Dir` that verifies only one config file is loaded when multiple exist in the directory
- ⚠️ Hardcoded relative paths in `testdataPath` can cause test failures in CI or different working directories
- ⚠️ Lack of assertion on error message content in `TestLoadConfig_Invalid` may mask incorrect validation logic
- 🔗 Tests rely on hardcoded paths that make them tightly coupled to project structure
- 🔗 No validation of error messages could lead to silent failures in config validation logic

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
- 🤖 The FilterPolicyPacks function has correct logic for filtering policy packs but lacks input validation and could benefit from clearer error handling and documentation.
- 💡 Add validation to ensure `p.Name` is not empty before using it in the map lookup
- 💡 Add explicit checks for nil or invalid `cfg.Enabled` and `cfg.Disabled` slices to prevent runtime panics
- 💡 Consider returning an error or logging a warning when `cfg.Enabled` and `cfg.Disabled` overlap, as this can lead to ambiguous behavior
- 💡 Refactor conditionals to explicitly check `len(cfg.Enabled) > 0` instead of checking `len(enabled) > 0` to improve clarity
- ⚠️ Potential panic if `p.Name` is nil or empty string, leading to runtime errors
- ⚠️ Incorrect filtering logic when `cfg.Enabled` is non-empty but does not include a policy name
- 🔗 This function tightly couples the configuration layer with domain policy packs, increasing system fragility if either changes
- 🔗 The filtering logic can silently produce incorrect results without clear error signaling, affecting downstream policy evaluation

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
- 🤖 The FilterPolicyPacks function has correct logic for filtering policy packs but lacks input validation and could benefit from clearer error handling and documentation.
- 💡 Add validation to ensure `p.Name` is not empty before using it in the map lookup
- 💡 Add explicit checks for nil or invalid `cfg.Enabled` and `cfg.Disabled` slices to prevent runtime panics
- 💡 Consider returning an error or logging a warning when `cfg.Enabled` and `cfg.Disabled` overlap, as this can lead to ambiguous behavior
- 💡 Refactor conditionals to explicitly check `len(cfg.Enabled) > 0` instead of checking `len(enabled) > 0` to improve clarity
- ⚠️ Potential panic if `p.Name` is nil or empty string, leading to runtime errors
- ⚠️ Incorrect filtering logic when `cfg.Enabled` is non-empty but does not include a policy name
- 🔗 This function tightly couples the configuration layer with domain policy packs, increasing system fragility if either changes
- 🔗 The filtering logic can silently produce incorrect results without clear error signaling, affecting downstream policy evaluation

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
- 🤖 The code is functionally correct but has a critical logic flaw in `LoadPolicyPacks` that prevents loading multiple policy packs and lacks proper error handling for malformed YAML.
- 💡 Modify `LoadPolicyPacks` to collect errors and continue processing other files instead of failing fast
- 💡 Add unit tests for `parseDimension` to ensure all valid dimension strings are mapped correctly and invalid ones return errors
- ⚠️ Early termination in `LoadPolicyPacks` prevents loading of valid policy packs when one file fails to parse
- ⚠️ Potential panic or silent failure in `parseSeverity` if `domain.ParseSeverity` is not properly validated
- 🔗 The `LoadPolicyPacks` function is tightly coupled to file system I/O and assumes all YAML files in a directory are valid policy packs, increasing failure propagation risk
- 🔗 The `parseDimension` function is decoupled and safe in isolation, but its usage within a loop that fails fast introduces systemic fragility

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
- 🤖 The code is functionally correct but has a critical logic flaw in `LoadPolicyPacks` that prevents loading multiple policy packs and lacks proper error handling for malformed YAML.
- 💡 Modify `LoadPolicyPacks` to collect errors and continue processing other files instead of failing fast
- 💡 Add unit tests for `parseDimension` to ensure all valid dimension strings are mapped correctly and invalid ones return errors
- ⚠️ Early termination in `LoadPolicyPacks` prevents loading of valid policy packs when one file fails to parse
- ⚠️ Potential panic or silent failure in `parseSeverity` if `domain.ParseSeverity` is not properly validated
- 🔗 The `LoadPolicyPacks` function is tightly coupled to file system I/O and assumes all YAML files in a directory are valid policy packs, increasing failure propagation risk
- 🔗 The `parseDimension` function is decoupled and safe in isolation, but its usage within a loop that fails fast introduces systemic fragility

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
- 🤖 The code is functionally correct but has a critical logic flaw in `LoadPolicyPacks` that prevents loading multiple policy packs and lacks proper error handling for malformed YAML.
- 💡 Modify `LoadPolicyPacks` to collect errors and continue processing other files instead of failing fast
- 💡 Add unit tests for `parseDimension` to ensure all valid dimension strings are mapped correctly and invalid ones return errors
- ⚠️ Early termination in `LoadPolicyPacks` prevents loading of valid policy packs when one file fails to parse
- ⚠️ Potential panic or silent failure in `parseSeverity` if `domain.ParseSeverity` is not properly validated
- 🔗 The `LoadPolicyPacks` function is tightly coupled to file system I/O and assumes all YAML files in a directory are valid policy packs, increasing failure propagation risk
- 🔗 The `parseDimension` function is decoupled and safe in isolation, but its usage within a loop that fails fast introduces systemic fragility

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
- 🤖 The code is functionally correct but has a critical logic flaw in `LoadPolicyPacks` that prevents loading multiple policy packs and lacks proper error handling for malformed YAML.
- 💡 Modify `LoadPolicyPacks` to collect errors and continue processing other files instead of failing fast
- 💡 Add unit tests for `parseDimension` to ensure all valid dimension strings are mapped correctly and invalid ones return errors
- ⚠️ Early termination in `LoadPolicyPacks` prevents loading of valid policy packs when one file fails to parse
- ⚠️ Potential panic or silent failure in `parseSeverity` if `domain.ParseSeverity` is not properly validated
- 🔗 The `LoadPolicyPacks` function is tightly coupled to file system I/O and assumes all YAML files in a directory are valid policy packs, increasing failure propagation risk
- 🔗 The `parseDimension` function is decoupled and safe in isolation, but its usage within a loop that fails fast introduces systemic fragility

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
- 🤖 The code is functionally correct but has a critical logic flaw in `LoadPolicyPacks` that prevents loading multiple policy packs and lacks proper error handling for malformed YAML.
- 💡 Modify `LoadPolicyPacks` to collect errors and continue processing other files instead of failing fast
- 💡 Add unit tests for `parseDimension` to ensure all valid dimension strings are mapped correctly and invalid ones return errors
- ⚠️ Early termination in `LoadPolicyPacks` prevents loading of valid policy packs when one file fails to parse
- ⚠️ Potential panic or silent failure in `parseSeverity` if `domain.ParseSeverity` is not properly validated
- 🔗 The `LoadPolicyPacks` function is tightly coupled to file system I/O and assumes all YAML files in a directory are valid policy packs, increasing failure propagation risk
- 🔗 The `parseDimension` function is decoupled and safe in isolation, but its usage within a loop that fails fast introduces systemic fragility

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
- 🤖 The code is functionally correct but has a critical logic flaw in `LoadPolicyPacks` that prevents loading multiple policy packs and lacks proper error handling for malformed YAML.
- 💡 Modify `LoadPolicyPacks` to collect errors and continue processing other files instead of failing fast
- 💡 Add unit tests for `parseDimension` to ensure all valid dimension strings are mapped correctly and invalid ones return errors
- ⚠️ Early termination in `LoadPolicyPacks` prevents loading of valid policy packs when one file fails to parse
- ⚠️ Potential panic or silent failure in `parseSeverity` if `domain.ParseSeverity` is not properly validated
- 🔗 The `LoadPolicyPacks` function is tightly coupled to file system I/O and assumes all YAML files in a directory are valid policy packs, increasing failure propagation risk
- 🔗 The `parseDimension` function is decoupled and safe in isolation, but its usage within a loop that fails fast introduces systemic fragility

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
- 🤖 The code is functionally correct but has a critical logic flaw in `LoadPolicyPacks` that prevents loading multiple policy packs and lacks proper error handling for malformed YAML.
- 💡 Modify `LoadPolicyPacks` to collect errors and continue processing other files instead of failing fast
- 💡 Add unit tests for `parseDimension` to ensure all valid dimension strings are mapped correctly and invalid ones return errors
- ⚠️ Early termination in `LoadPolicyPacks` prevents loading of valid policy packs when one file fails to parse
- ⚠️ Potential panic or silent failure in `parseSeverity` if `domain.ParseSeverity` is not properly validated
- 🔗 The `LoadPolicyPacks` function is tightly coupled to file system I/O and assumes all YAML files in a directory are valid policy packs, increasing failure propagation risk
- 🔗 The `parseDimension` function is decoupled and safe in isolation, but its usage within a loop that fails fast introduces systemic fragility

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
- 🤖 The test suite for policy loading is functional but lacks comprehensive validation and has brittle path handling.
- 💡 Replace hardcoded relative paths in `policiesPath` with a configurable or dynamically resolved path using `t.Cleanup()` and `filepath.Abs()` to ensure tests run reliably regardless of working directory
- 💡 Improve error assertion in `TestLoadPolicyPack_Invalid` by checking that the returned error is of a specific type (e.g., `*yaml.TypeError`) or includes expected keywords to validate that the error is meaningful and actionable
- 💡 Add assertions for all fields in `domain.Rule` (e.g., Description, Enabled, Parameters) to ensure full fidelity of loaded policies and prevent silent data loss or misinterpretation
- 💡 Add a test case for `TestLoadPolicyPack_Invalid` that validates missing language field or other required fields to ensure robust error handling for malformed policies
- 💡 Use `t.TempDir()` in `TestLoadPolicyPack_Invalid` to ensure proper cleanup and avoid potential side effects from file system mutations
- ⚠️ Brittle path handling in `policiesPath` can cause test failures when run from unexpected directories
- ⚠️ Insufficient error validation in `TestLoadPolicyPack_Invalid` makes it unclear whether errors are properly handled or just returned
- 🔗 Tests depend on hardcoded relative paths, increasing coupling to project structure and reducing portability
- 🔗 Lack of validation for all rule fields increases risk of silent data corruption or incorrect policy enforcement

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
- 🤖 The validation logic is functionally correct but has several maintainability and error handling weaknesses, including hardcoded strings, missing validation for edge cases like zero values in policy packs, and lack of input sanitization.
- 💡 Replace hardcoded field path strings with a helper function or structured error builder to improve maintainability and reduce coupling
- 💡 Add nil checks before calling .String() on rule.Severity and rule.Dimension in ValidatePolicyPack to prevent panics
- 💡 Consider validating that DefaultWindowDays > 0 is a strict requirement or allow zero as a special case with documentation
- ⚠️ Potential panic from calling .String() on nil or malformed Severity/Dimension fields in ValidatePolicyPack
- ⚠️ Brittle field path strings that can break on config schema changes
- 🔗 The validation functions are tightly coupled to the internal domain model structure, making them fragile to schema changes
- 🔗 Error messages are hardcoded and not localized or customizable, which limits extensibility

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
- 🤖 The validation logic is functionally correct but has several maintainability and error handling weaknesses, including hardcoded strings, missing validation for edge cases like zero values in policy packs, and lack of input sanitization.
- 💡 Replace hardcoded field path strings with a helper function or structured error builder to improve maintainability and reduce coupling
- 💡 Add nil checks before calling .String() on rule.Severity and rule.Dimension in ValidatePolicyPack to prevent panics
- 💡 Consider validating that DefaultWindowDays > 0 is a strict requirement or allow zero as a special case with documentation
- ⚠️ Potential panic from calling .String() on nil or malformed Severity/Dimension fields in ValidatePolicyPack
- ⚠️ Brittle field path strings that can break on config schema changes
- 🔗 The validation functions are tightly coupled to the internal domain model structure, making them fragile to schema changes
- 🔗 Error messages are hardcoded and not localized or customizable, which limits extensibility

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
- 🤖 The validation logic is functionally correct but has several maintainability and error handling weaknesses, including hardcoded strings, missing validation for edge cases like zero values in policy packs, and lack of input sanitization.
- 💡 Replace hardcoded field path strings with a helper function or structured error builder to improve maintainability and reduce coupling
- 💡 Add nil checks before calling .String() on rule.Severity and rule.Dimension in ValidatePolicyPack to prevent panics
- 💡 Consider validating that DefaultWindowDays > 0 is a strict requirement or allow zero as a special case with documentation
- ⚠️ Potential panic from calling .String() on nil or malformed Severity/Dimension fields in ValidatePolicyPack
- ⚠️ Brittle field path strings that can break on config schema changes
- 🔗 The validation functions are tightly coupled to the internal domain model structure, making them fragile to schema changes
- 🔗 Error messages are hardcoded and not localized or customizable, which limits extensibility

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
- 🤖 The validation logic is functionally correct but has several maintainability and error handling weaknesses, including hardcoded strings, missing validation for edge cases like zero values in policy packs, and lack of input sanitization.
- 💡 Replace hardcoded field path strings with a helper function or structured error builder to improve maintainability and reduce coupling
- 💡 Add nil checks before calling .String() on rule.Severity and rule.Dimension in ValidatePolicyPack to prevent panics
- 💡 Consider validating that DefaultWindowDays > 0 is a strict requirement or allow zero as a special case with documentation
- ⚠️ Potential panic from calling .String() on nil or malformed Severity/Dimension fields in ValidatePolicyPack
- ⚠️ Brittle field path strings that can break on config schema changes
- 🔗 The validation functions are tightly coupled to the internal domain model structure, making them fragile to schema changes
- 🔗 Error messages are hardcoded and not localized or customizable, which limits extensibility

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
- 🤖 The test suite for config validation has good coverage of basic scenarios but lacks robustness in edge case handling and error message consistency.
- 💡 Add assertions to verify that specific validation errors are returned with correct field paths and descriptive messages in `TestValidateConfig_InvalidExpiry`
- 💡 Add tests for boundary conditions such as empty rule ID and invalid metric names in `TestValidatePolicyPack_InvalidRuleMetric` to ensure all fields are properly validated
- 💡 Ensure that `TestValidatePolicyPack_MissingFields` explicitly checks for expected error types (e.g., required field errors) rather than just counting errors
- ⚠️ Lack of specific error message assertions makes tests brittle to refactoring or changes in validation logic
- ⚠️ Missing edge case testing for invalid rule IDs and metrics (e.g., empty or malformed values)
- 🔗 These tests provide minimal assurance that validation logic correctly handles all valid and invalid configurations, potentially masking silent failures in real-world usage
- 🔗 The lack of consistent error message validation increases coupling between the validator and consumers who depend on structured error outputs

</details>

### `internal/discovery/` (39 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`DetectLanguages`](reports/internal-discovery-detect-go-detectlanguages.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`DetectedAdapters`](reports/internal-discovery-detect-go-detectedadapters.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`LanguageInfo`](reports/internal-discovery-detect-go-languageinfo.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`buildLanguageList`](reports/internal-discovery-detect-go-buildlanguagelist.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`detect_test.go`](reports/internal-discovery-detect-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`ChangedFiles`](reports/internal-discovery-diff-go-changedfiles.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`DetectMoves`](reports/internal-discovery-diff-go-detectmoves.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FilterByPaths`](reports/internal-discovery-diff-go-filterbypaths.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`FilterChanged`](reports/internal-discovery-diff-go-filterchanged.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`MovedFile`](reports/internal-discovery-diff-go-movedfile.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`diff_test.go`](reports/internal-discovery-diff-test-go.md) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`GenericScanner`](reports/internal-discovery-generic-go-genericscanner.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`NewGenericScanner`](reports/internal-discovery-generic-go-newgenericscanner.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Scan`](reports/internal-discovery-generic-go-scan.md) | method | B+ | 87.2% | certified | 2026-04-23 |
| [`matchAny`](reports/internal-discovery-generic-go-matchany.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`GoAdapter`](reports/internal-discovery-go-adapter-go-goadapter.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`NewGoAdapter`](reports/internal-discovery-go-adapter-go-newgoadapter.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Scan`](reports/internal-discovery-go-adapter-go-scan.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`parseFile`](reports/internal-discovery-go-adapter-go-parsefile.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`go_adapter_test.go`](reports/internal-discovery-go-adapter-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Diff`](reports/internal-discovery-index-go-diff.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`DiffResult`](reports/internal-discovery-index-go-diffresult.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Index`](reports/internal-discovery-index-go-index.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`LoadIndex`](reports/internal-discovery-index-go-loadindex.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewIndex`](reports/internal-discovery-index-go-newindex.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Save`](reports/internal-discovery-index-go-save.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Units`](reports/internal-discovery-index-go-units.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`indexEntry`](reports/internal-discovery-index-go-indexentry.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`index_test.go`](reports/internal-discovery-index-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`DeduplicateFileLevel`](reports/internal-discovery-scanner-go-deduplicatefilelevel.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`Merge`](reports/internal-discovery-scanner-go-merge.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`Scanner`](reports/internal-discovery-scanner-go-scanner.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`UnitList`](reports/internal-discovery-scanner-go-unitlist.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`scanner_test.go`](reports/internal-discovery-scanner-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`NewTSAdapter`](reports/internal-discovery-ts-adapter-go-newtsadapter.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Scan`](reports/internal-discovery-ts-adapter-go-scan.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`TSAdapter`](reports/internal-discovery-ts-adapter-go-tsadapter.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`parseFile`](reports/internal-discovery-ts-adapter-go-parsefile.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`ts_adapter_test.go`](reports/internal-discovery-ts-adapter-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The code correctly detects programming languages and adapters but has a critical bug in adapter detection logic and lacks proper error handling for filesystem operations.
- 💡 Fix `DetectedAdapters` to properly deduplicate adapters by checking unique language or adapter identifiers instead of just the raw adapter field.
- 💡 Modify `DetectLanguages` to return an error when `filepath.WalkDir` encounters a filesystem issue rather than silently ignoring it.
- 💡 Add validation to ensure all supported languages in `extToLanguage` have corresponding entries in `languageToAdapter`, or explicitly define fallbacks.
- 💡 Add unit tests for edge cases like duplicate adapters, missing config files, and filesystem errors.
- ⚠️ Inconsistent adapter mapping: 'typescript' and 'javascript' both map to 'ts', but the logic in `DetectedAdapters` does not group them correctly, potentially causing redundant or incorrect adapter usage.
- ⚠️ Silent filesystem errors: `filepath.WalkDir` errors are ignored by returning nil, which can mask real issues like permission failures or broken symbolic links.
- 🔗 The `DetectedAdapters` function affects downstream systems that rely on unique adapter identifiers, potentially causing duplicate or incorrect adapter loading.
- 🔗 The lack of proper error handling in `DetectLanguages` introduces a failure propagation risk where silent failures in directory traversal can mask real problems.

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
- 🤖 The code correctly detects programming languages and adapters but has a critical bug in adapter detection logic and lacks proper error handling for filesystem operations.
- 💡 Fix `DetectedAdapters` to properly deduplicate adapters by checking unique language or adapter identifiers instead of just the raw adapter field.
- 💡 Modify `DetectLanguages` to return an error when `filepath.WalkDir` encounters a filesystem issue rather than silently ignoring it.
- 💡 Add validation to ensure all supported languages in `extToLanguage` have corresponding entries in `languageToAdapter`, or explicitly define fallbacks.
- 💡 Add unit tests for edge cases like duplicate adapters, missing config files, and filesystem errors.
- ⚠️ Inconsistent adapter mapping: 'typescript' and 'javascript' both map to 'ts', but the logic in `DetectedAdapters` does not group them correctly, potentially causing redundant or incorrect adapter usage.
- ⚠️ Silent filesystem errors: `filepath.WalkDir` errors are ignored by returning nil, which can mask real issues like permission failures or broken symbolic links.
- 🔗 The `DetectedAdapters` function affects downstream systems that rely on unique adapter identifiers, potentially causing duplicate or incorrect adapter loading.
- 🔗 The lack of proper error handling in `DetectLanguages` introduces a failure propagation risk where silent failures in directory traversal can mask real problems.

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
- 🤖 The code correctly detects programming languages and adapters but has a critical bug in adapter detection logic and lacks proper error handling for filesystem operations.
- 💡 Fix `DetectedAdapters` to properly deduplicate adapters by checking unique language or adapter identifiers instead of just the raw adapter field.
- 💡 Modify `DetectLanguages` to return an error when `filepath.WalkDir` encounters a filesystem issue rather than silently ignoring it.
- 💡 Add validation to ensure all supported languages in `extToLanguage` have corresponding entries in `languageToAdapter`, or explicitly define fallbacks.
- 💡 Add unit tests for edge cases like duplicate adapters, missing config files, and filesystem errors.
- ⚠️ Inconsistent adapter mapping: 'typescript' and 'javascript' both map to 'ts', but the logic in `DetectedAdapters` does not group them correctly, potentially causing redundant or incorrect adapter usage.
- ⚠️ Silent filesystem errors: `filepath.WalkDir` errors are ignored by returning nil, which can mask real issues like permission failures or broken symbolic links.
- 🔗 The `DetectedAdapters` function affects downstream systems that rely on unique adapter identifiers, potentially causing duplicate or incorrect adapter loading.
- 🔗 The lack of proper error handling in `DetectLanguages` introduces a failure propagation risk where silent failures in directory traversal can mask real problems.

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
- 🤖 The code correctly detects programming languages and adapters but has a critical bug in adapter detection logic and lacks proper error handling for filesystem operations.
- 💡 Fix `DetectedAdapters` to properly deduplicate adapters by checking unique language or adapter identifiers instead of just the raw adapter field.
- 💡 Modify `DetectLanguages` to return an error when `filepath.WalkDir` encounters a filesystem issue rather than silently ignoring it.
- 💡 Add validation to ensure all supported languages in `extToLanguage` have corresponding entries in `languageToAdapter`, or explicitly define fallbacks.
- 💡 Add unit tests for edge cases like duplicate adapters, missing config files, and filesystem errors.
- ⚠️ Inconsistent adapter mapping: 'typescript' and 'javascript' both map to 'ts', but the logic in `DetectedAdapters` does not group them correctly, potentially causing redundant or incorrect adapter usage.
- ⚠️ Silent filesystem errors: `filepath.WalkDir` errors are ignored by returning nil, which can mask real issues like permission failures or broken symbolic links.
- 🔗 The `DetectedAdapters` function affects downstream systems that rely on unique adapter identifiers, potentially causing duplicate or incorrect adapter loading.
- 🔗 The lack of proper error handling in `DetectLanguages` introduces a failure propagation risk where silent failures in directory traversal can mask real problems.

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
- 🤖 The test file has minimal coverage and lacks proper error handling, making it brittle and potentially misleading.
- 💡 Refactor `repoPath` to be a configurable or mockable function instead of relying on hardcoded paths
- 💡 Add error handling and assertions for `discovery.DetectLanguages` return values to prevent silent failures
- 💡 Use table-driven tests for `TestDetectLanguages_*` to improve maintainability and reduce duplication
- ⚠️ Flaky tests due to reliance on external repo files and directory structure
- ⚠️ Potential panic or incorrect behavior if `repoPath` returns invalid paths or fails to resolve
- 🔗 These tests do not affect core system behavior but introduce fragility in CI/CD pipelines due to reliance on external state
- 🔗 Tight coupling between test logic and internal `repoPath` utility increases maintenance burden

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
- 🤖 The code is functionally correct but has security and performance concerns due to direct shell command execution and inefficient filtering logic.
- 💡 Sanitize git ref inputs in `ChangedFiles` to prevent command injection by validating or escaping special characters
- 💡 Replace the O(n*m) loop in `FilterByPaths` with a more efficient data structure like a trie or prefix tree for better performance
- 💡 Make git reference handling in `DetectMoves` configurable or parameterized to support different branches or commits
- 💡 Add proper error handling and logging for git command failures in `ChangedFiles` and `DetectMoves`
- 💡 Consider using a more robust path matching mechanism in `FilterChanged` that accounts for relative vs absolute paths or normalization
- ⚠️ Command injection vulnerability in `ChangedFiles` due to direct use of user-provided git ref arguments without sanitization
- ⚠️ Potential performance degradation in `FilterByPaths` with O(n*m) complexity for large lists of units and paths
- 🔗 The use of `exec.Command` ties this module tightly to the system's git environment and shell, increasing coupling and reducing portability
- 🔗 Hardcoded git references in `DetectMoves` increase fragility and reduce testability of the module

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
- 🤖 The code is functionally correct but has security and performance concerns due to direct shell command execution and inefficient filtering logic.
- 💡 Sanitize git ref inputs in `ChangedFiles` to prevent command injection by validating or escaping special characters
- 💡 Replace the O(n*m) loop in `FilterByPaths` with a more efficient data structure like a trie or prefix tree for better performance
- 💡 Make git reference handling in `DetectMoves` configurable or parameterized to support different branches or commits
- 💡 Add proper error handling and logging for git command failures in `ChangedFiles` and `DetectMoves`
- 💡 Consider using a more robust path matching mechanism in `FilterChanged` that accounts for relative vs absolute paths or normalization
- ⚠️ Command injection vulnerability in `ChangedFiles` due to direct use of user-provided git ref arguments without sanitization
- ⚠️ Potential performance degradation in `FilterByPaths` with O(n*m) complexity for large lists of units and paths
- 🔗 The use of `exec.Command` ties this module tightly to the system's git environment and shell, increasing coupling and reducing portability
- 🔗 Hardcoded git references in `DetectMoves` increase fragility and reduce testability of the module

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
- 🤖 The code is functionally correct but has security and performance concerns due to direct shell command execution and inefficient filtering logic.
- 💡 Sanitize git ref inputs in `ChangedFiles` to prevent command injection by validating or escaping special characters
- 💡 Replace the O(n*m) loop in `FilterByPaths` with a more efficient data structure like a trie or prefix tree for better performance
- 💡 Make git reference handling in `DetectMoves` configurable or parameterized to support different branches or commits
- 💡 Add proper error handling and logging for git command failures in `ChangedFiles` and `DetectMoves`
- 💡 Consider using a more robust path matching mechanism in `FilterChanged` that accounts for relative vs absolute paths or normalization
- ⚠️ Command injection vulnerability in `ChangedFiles` due to direct use of user-provided git ref arguments without sanitization
- ⚠️ Potential performance degradation in `FilterByPaths` with O(n*m) complexity for large lists of units and paths
- 🔗 The use of `exec.Command` ties this module tightly to the system's git environment and shell, increasing coupling and reducing portability
- 🔗 Hardcoded git references in `DetectMoves` increase fragility and reduce testability of the module

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
- 🤖 The code is functionally correct but has security and performance concerns due to direct shell command execution and inefficient filtering logic.
- 💡 Sanitize git ref inputs in `ChangedFiles` to prevent command injection by validating or escaping special characters
- 💡 Replace the O(n*m) loop in `FilterByPaths` with a more efficient data structure like a trie or prefix tree for better performance
- 💡 Make git reference handling in `DetectMoves` configurable or parameterized to support different branches or commits
- 💡 Add proper error handling and logging for git command failures in `ChangedFiles` and `DetectMoves`
- 💡 Consider using a more robust path matching mechanism in `FilterChanged` that accounts for relative vs absolute paths or normalization
- ⚠️ Command injection vulnerability in `ChangedFiles` due to direct use of user-provided git ref arguments without sanitization
- ⚠️ Potential performance degradation in `FilterByPaths` with O(n*m) complexity for large lists of units and paths
- 🔗 The use of `exec.Command` ties this module tightly to the system's git environment and shell, increasing coupling and reducing portability
- 🔗 Hardcoded git references in `DetectMoves` increase fragility and reduce testability of the module

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
- 🤖 The code is functionally correct but has security and performance concerns due to direct shell command execution and inefficient filtering logic.
- 💡 Sanitize git ref inputs in `ChangedFiles` to prevent command injection by validating or escaping special characters
- 💡 Replace the O(n*m) loop in `FilterByPaths` with a more efficient data structure like a trie or prefix tree for better performance
- 💡 Make git reference handling in `DetectMoves` configurable or parameterized to support different branches or commits
- 💡 Add proper error handling and logging for git command failures in `ChangedFiles` and `DetectMoves`
- 💡 Consider using a more robust path matching mechanism in `FilterChanged` that accounts for relative vs absolute paths or normalization
- ⚠️ Command injection vulnerability in `ChangedFiles` due to direct use of user-provided git ref arguments without sanitization
- ⚠️ Potential performance degradation in `FilterByPaths` with O(n*m) complexity for large lists of units and paths
- 🔗 The use of `exec.Command` ties this module tightly to the system's git environment and shell, increasing coupling and reducing portability
- 🔗 Hardcoded git references in `DetectMoves` increase fragility and reduce testability of the module

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
- 🤖 The test file has minimal coverage and lacks proper validation of edge cases or error conditions in the discovery package.
- 💡 Add tests for edge cases in FilterChanged: e.g., duplicate paths, missing files, and incorrect path formats
- 💡 Add tests for FilterByPaths to ensure that overlapping or nested paths are handled correctly and do not produce unexpected results
- ⚠️ Missing validation of malformed or invalid paths in FilterChanged and FilterByPaths functions
- ⚠️ Lack of test coverage for overlapping or nested path filters in FilterByPaths
- 🔗 This unit only tests the discovery package's filtering logic, but does not provide sufficient validation of its behavior in real-world usage
- 🔗 The tests do not simulate concurrent access or race conditions, which could be an issue if these functions are used in a multi-threaded environment

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
- 🤖 The code implements a file scanner with glob filtering but contains critical runtime errors due to undefined variable references and incorrect logic in path matching.
- 💡 Replace `skipDirs` with `scanSkipDirs` on line 93 to fix undefined variable reference
- 💡 Refactor `matchAny` function to correctly implement glob pattern matching using `filepath.Match` for full paths and handle special cases like "**" properly
- 💡 Validate that the `root` path is a valid directory before calling `filepath.WalkDir` to prevent runtime errors
- 💡 Add unit tests for the `matchAny` function with various glob patterns to ensure correct matching behavior
- ⚠️ Runtime panic from undefined variable `skipDirs`
- ⚠️ Incorrect glob pattern matching leading to incorrect file inclusion/exclusion
- 🔗 The scanner's filtering logic is broken, causing incorrect file discovery and potentially exposing non-certifiable files
- 🔗 The use of undefined variable `skipDirs` introduces a silent failure or panic that can break the entire discovery pipeline

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
- 🤖 The code implements a file scanner with glob filtering but contains critical runtime errors due to undefined variable references and incorrect logic in path matching.
- 💡 Replace `skipDirs` with `scanSkipDirs` on line 93 to fix undefined variable reference
- 💡 Refactor `matchAny` function to correctly implement glob pattern matching using `filepath.Match` for full paths and handle special cases like "**" properly
- 💡 Validate that the `root` path is a valid directory before calling `filepath.WalkDir` to prevent runtime errors
- 💡 Add unit tests for the `matchAny` function with various glob patterns to ensure correct matching behavior
- ⚠️ Runtime panic from undefined variable `skipDirs`
- ⚠️ Incorrect glob pattern matching leading to incorrect file inclusion/exclusion
- 🔗 The scanner's filtering logic is broken, causing incorrect file discovery and potentially exposing non-certifiable files
- 🔗 The use of undefined variable `skipDirs` introduces a silent failure or panic that can break the entire discovery pipeline

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
- 🤖 The code implements a file scanner with glob filtering but contains critical runtime errors due to undefined variable references and incorrect logic in path matching.
- 💡 Replace `skipDirs` with `scanSkipDirs` on line 93 to fix undefined variable reference
- 💡 Refactor `matchAny` function to correctly implement glob pattern matching using `filepath.Match` for full paths and handle special cases like "**" properly
- 💡 Validate that the `root` path is a valid directory before calling `filepath.WalkDir` to prevent runtime errors
- 💡 Add unit tests for the `matchAny` function with various glob patterns to ensure correct matching behavior
- ⚠️ Runtime panic from undefined variable `skipDirs`
- ⚠️ Incorrect glob pattern matching leading to incorrect file inclusion/exclusion
- 🔗 The scanner's filtering logic is broken, causing incorrect file discovery and potentially exposing non-certifiable files
- 🔗 The use of undefined variable `skipDirs` introduces a silent failure or panic that can break the entire discovery pipeline

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
- 🤖 The code implements a file scanner with glob filtering but contains critical runtime errors due to undefined variable references and incorrect logic in path matching.
- 💡 Replace `skipDirs` with `scanSkipDirs` on line 93 to fix undefined variable reference
- 💡 Refactor `matchAny` function to correctly implement glob pattern matching using `filepath.Match` for full paths and handle special cases like "**" properly
- 💡 Validate that the `root` path is a valid directory before calling `filepath.WalkDir` to prevent runtime errors
- 💡 Add unit tests for the `matchAny` function with various glob patterns to ensure correct matching behavior
- ⚠️ Runtime panic from undefined variable `skipDirs`
- ⚠️ Incorrect glob pattern matching leading to incorrect file inclusion/exclusion
- 🔗 The scanner's filtering logic is broken, causing incorrect file discovery and potentially exposing non-certifiable files
- 🔗 The use of undefined variable `skipDirs` introduces a silent failure or panic that can break the entire discovery pipeline

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
- 🤖 The GoAdapter correctly discovers Go code units but has a critical error handling flaw where parsing errors lead to silent fallbacks instead of proper error propagation.
- 💡 Replace the graceful fallback in Scan() with a proper error return or logging to ensure invalid files are flagged
- 💡 Add support for handling ast.ValueSpec and ast.VarDecl in parseFile() to capture variable units
- 💡 Improve directory filtering by checking if the root path is a subdirectory of any skipped directory to prevent traversing into vendor or testdata
- 💡 Ensure that unit ID string comparison in sort.Slice is robust and doesn't rely on lexicographic ordering of complex IDs
- ⚠️ Silent error recovery in Scan() leads to incorrect reporting of malformed or unreadable Go files as valid units
- ⚠️ Missing support for other AST node types (e.g., var declarations) results in incomplete unit discovery
- 🔗 This unit introduces a weak failure mode where invalid Go files are silently treated as valid, causing downstream systems to process incorrect or incomplete data
- 🔗 The unit tightly couples to the internal domain.Unit structure and assumes specific ID formats, increasing system fragility

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
- 🤖 The GoAdapter correctly discovers Go code units but has a critical error handling flaw where parsing errors lead to silent fallbacks instead of proper error propagation.
- 💡 Replace the graceful fallback in Scan() with a proper error return or logging to ensure invalid files are flagged
- 💡 Add support for handling ast.ValueSpec and ast.VarDecl in parseFile() to capture variable units
- 💡 Improve directory filtering by checking if the root path is a subdirectory of any skipped directory to prevent traversing into vendor or testdata
- 💡 Ensure that unit ID string comparison in sort.Slice is robust and doesn't rely on lexicographic ordering of complex IDs
- ⚠️ Silent error recovery in Scan() leads to incorrect reporting of malformed or unreadable Go files as valid units
- ⚠️ Missing support for other AST node types (e.g., var declarations) results in incomplete unit discovery
- 🔗 This unit introduces a weak failure mode where invalid Go files are silently treated as valid, causing downstream systems to process incorrect or incomplete data
- 🔗 The unit tightly couples to the internal domain.Unit structure and assumes specific ID formats, increasing system fragility

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
- 🤖 The GoAdapter correctly discovers Go code units but has a critical error handling flaw where parsing errors lead to silent fallbacks instead of proper error propagation.
- 💡 Replace the graceful fallback in Scan() with a proper error return or logging to ensure invalid files are flagged
- 💡 Add support for handling ast.ValueSpec and ast.VarDecl in parseFile() to capture variable units
- 💡 Improve directory filtering by checking if the root path is a subdirectory of any skipped directory to prevent traversing into vendor or testdata
- 💡 Ensure that unit ID string comparison in sort.Slice is robust and doesn't rely on lexicographic ordering of complex IDs
- ⚠️ Silent error recovery in Scan() leads to incorrect reporting of malformed or unreadable Go files as valid units
- ⚠️ Missing support for other AST node types (e.g., var declarations) results in incomplete unit discovery
- 🔗 This unit introduces a weak failure mode where invalid Go files are silently treated as valid, causing downstream systems to process incorrect or incomplete data
- 🔗 The unit tightly couples to the internal domain.Unit structure and assumes specific ID formats, increasing system fragility

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
- 🤖 The GoAdapter correctly discovers Go code units but has a critical error handling flaw where parsing errors lead to silent fallbacks instead of proper error propagation.
- 💡 Replace the graceful fallback in Scan() with a proper error return or logging to ensure invalid files are flagged
- 💡 Add support for handling ast.ValueSpec and ast.VarDecl in parseFile() to capture variable units
- 💡 Improve directory filtering by checking if the root path is a subdirectory of any skipped directory to prevent traversing into vendor or testdata
- 💡 Ensure that unit ID string comparison in sort.Slice is robust and doesn't rely on lexicographic ordering of complex IDs
- ⚠️ Silent error recovery in Scan() leads to incorrect reporting of malformed or unreadable Go files as valid units
- ⚠️ Missing support for other AST node types (e.g., var declarations) results in incomplete unit discovery
- 🔗 This unit introduces a weak failure mode where invalid Go files are silently treated as valid, causing downstream systems to process incorrect or incomplete data
- 🔗 The unit tightly couples to the internal domain.Unit structure and assumes specific ID formats, increasing system fragility

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
- 🤖 The test suite verifies Go adapter functionality but lacks robustness in error handling, deterministic behavior checks, and API consistency validation.
- 💡 Replace `t.Fatalf` in `TestGoAdapter_Functions` with a loop that collects all missing units and reports them at once to improve debuggability
- 💡 Refactor repeated logic in `TestGoAdapter_Methods`, `TestGoAdapter_FreeStandingFunction`, and `TestGoAdapter_Types` into a shared helper function to reduce redundancy
- 💡 Add a test case for scanning an empty or invalid repository path to ensure graceful error handling
- 💡 Validate that unit metadata (e.g., line numbers, file paths) is stable and correctly parsed across multiple scans
- ⚠️ Potential false positives in unit ID matching due to reliance on string-based equality without normalization or canonicalization
- ⚠️ Race condition risk if `adapter.Scan()` is not thread-safe and tests run concurrently (not currently observed but possible in future)
- 🔗 This file only contains unit tests and does not directly affect system behavior or API surface
- 🔗 If the adapter under test has bugs, these tests might pass incorrectly due to incomplete validation logic

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
- 🤖 The code is functionally correct but has critical error handling issues, lacks proper validation, and introduces potential security and performance risks.
- 💡 Replace `ut, _ := domain.ParseUnitType(e.Type)` with proper error handling to ensure invalid types are not silently accepted
- 💡 Validate that `domain.Unit.ID.String()` is a safe and stable key before using it in maps for diffing, or use a more robust identifier like a hash
- 💡 Add input validation and sanitization for file paths in `Save` to prevent directory traversal or overwriting sensitive files
- 💡 Make `Units()` return a copy of the slice instead of the direct reference to prevent external mutation
- 💡 Add unit tests for edge cases such as nil inputs to `Diff` and malformed JSON in `LoadIndex`
- ⚠️ Silent failure in unit type parsing during index load
- ⚠️ Potential data corruption or incorrect diffing due to reliance on `domain.Unit.ID.String()` as a unique key
- 🔗 The `LoadIndex` function introduces a failure propagation risk by silently ignoring unit type parsing errors, which could lead to invalid or corrupted index data being used downstream
- 🔗 The `Units()` method exposes internal state, increasing coupling and reducing encapsulation; any consumer can mutate the index directly

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
- 🤖 The code is functionally correct but has critical error handling issues, lacks proper validation, and introduces potential security and performance risks.
- 💡 Replace `ut, _ := domain.ParseUnitType(e.Type)` with proper error handling to ensure invalid types are not silently accepted
- 💡 Validate that `domain.Unit.ID.String()` is a safe and stable key before using it in maps for diffing, or use a more robust identifier like a hash
- 💡 Add input validation and sanitization for file paths in `Save` to prevent directory traversal or overwriting sensitive files
- 💡 Make `Units()` return a copy of the slice instead of the direct reference to prevent external mutation
- 💡 Add unit tests for edge cases such as nil inputs to `Diff` and malformed JSON in `LoadIndex`
- ⚠️ Silent failure in unit type parsing during index load
- ⚠️ Potential data corruption or incorrect diffing due to reliance on `domain.Unit.ID.String()` as a unique key
- 🔗 The `LoadIndex` function introduces a failure propagation risk by silently ignoring unit type parsing errors, which could lead to invalid or corrupted index data being used downstream
- 🔗 The `Units()` method exposes internal state, increasing coupling and reducing encapsulation; any consumer can mutate the index directly

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
- 🤖 The code is functionally correct but has critical error handling issues, lacks proper validation, and introduces potential security and performance risks.
- 💡 Replace `ut, _ := domain.ParseUnitType(e.Type)` with proper error handling to ensure invalid types are not silently accepted
- 💡 Validate that `domain.Unit.ID.String()` is a safe and stable key before using it in maps for diffing, or use a more robust identifier like a hash
- 💡 Add input validation and sanitization for file paths in `Save` to prevent directory traversal or overwriting sensitive files
- 💡 Make `Units()` return a copy of the slice instead of the direct reference to prevent external mutation
- 💡 Add unit tests for edge cases such as nil inputs to `Diff` and malformed JSON in `LoadIndex`
- ⚠️ Silent failure in unit type parsing during index load
- ⚠️ Potential data corruption or incorrect diffing due to reliance on `domain.Unit.ID.String()` as a unique key
- 🔗 The `LoadIndex` function introduces a failure propagation risk by silently ignoring unit type parsing errors, which could lead to invalid or corrupted index data being used downstream
- 🔗 The `Units()` method exposes internal state, increasing coupling and reducing encapsulation; any consumer can mutate the index directly

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
- 🤖 The code is functionally correct but has critical error handling issues, lacks proper validation, and introduces potential security and performance risks.
- 💡 Replace `ut, _ := domain.ParseUnitType(e.Type)` with proper error handling to ensure invalid types are not silently accepted
- 💡 Validate that `domain.Unit.ID.String()` is a safe and stable key before using it in maps for diffing, or use a more robust identifier like a hash
- 💡 Add input validation and sanitization for file paths in `Save` to prevent directory traversal or overwriting sensitive files
- 💡 Make `Units()` return a copy of the slice instead of the direct reference to prevent external mutation
- 💡 Add unit tests for edge cases such as nil inputs to `Diff` and malformed JSON in `LoadIndex`
- ⚠️ Silent failure in unit type parsing during index load
- ⚠️ Potential data corruption or incorrect diffing due to reliance on `domain.Unit.ID.String()` as a unique key
- 🔗 The `LoadIndex` function introduces a failure propagation risk by silently ignoring unit type parsing errors, which could lead to invalid or corrupted index data being used downstream
- 🔗 The `Units()` method exposes internal state, increasing coupling and reducing encapsulation; any consumer can mutate the index directly

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
- 🤖 The code is functionally correct but has critical error handling issues, lacks proper validation, and introduces potential security and performance risks.
- 💡 Replace `ut, _ := domain.ParseUnitType(e.Type)` with proper error handling to ensure invalid types are not silently accepted
- 💡 Validate that `domain.Unit.ID.String()` is a safe and stable key before using it in maps for diffing, or use a more robust identifier like a hash
- 💡 Add input validation and sanitization for file paths in `Save` to prevent directory traversal or overwriting sensitive files
- 💡 Make `Units()` return a copy of the slice instead of the direct reference to prevent external mutation
- 💡 Add unit tests for edge cases such as nil inputs to `Diff` and malformed JSON in `LoadIndex`
- ⚠️ Silent failure in unit type parsing during index load
- ⚠️ Potential data corruption or incorrect diffing due to reliance on `domain.Unit.ID.String()` as a unique key
- 🔗 The `LoadIndex` function introduces a failure propagation risk by silently ignoring unit type parsing errors, which could lead to invalid or corrupted index data being used downstream
- 🔗 The `Units()` method exposes internal state, increasing coupling and reducing encapsulation; any consumer can mutate the index directly

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
- 🤖 The code is functionally correct but has critical error handling issues, lacks proper validation, and introduces potential security and performance risks.
- 💡 Replace `ut, _ := domain.ParseUnitType(e.Type)` with proper error handling to ensure invalid types are not silently accepted
- 💡 Validate that `domain.Unit.ID.String()` is a safe and stable key before using it in maps for diffing, or use a more robust identifier like a hash
- 💡 Add input validation and sanitization for file paths in `Save` to prevent directory traversal or overwriting sensitive files
- 💡 Make `Units()` return a copy of the slice instead of the direct reference to prevent external mutation
- 💡 Add unit tests for edge cases such as nil inputs to `Diff` and malformed JSON in `LoadIndex`
- ⚠️ Silent failure in unit type parsing during index load
- ⚠️ Potential data corruption or incorrect diffing due to reliance on `domain.Unit.ID.String()` as a unique key
- 🔗 The `LoadIndex` function introduces a failure propagation risk by silently ignoring unit type parsing errors, which could lead to invalid or corrupted index data being used downstream
- 🔗 The `Units()` method exposes internal state, increasing coupling and reducing encapsulation; any consumer can mutate the index directly

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
- 🤖 The code is functionally correct but has critical error handling issues, lacks proper validation, and introduces potential security and performance risks.
- 💡 Replace `ut, _ := domain.ParseUnitType(e.Type)` with proper error handling to ensure invalid types are not silently accepted
- 💡 Validate that `domain.Unit.ID.String()` is a safe and stable key before using it in maps for diffing, or use a more robust identifier like a hash
- 💡 Add input validation and sanitization for file paths in `Save` to prevent directory traversal or overwriting sensitive files
- 💡 Make `Units()` return a copy of the slice instead of the direct reference to prevent external mutation
- 💡 Add unit tests for edge cases such as nil inputs to `Diff` and malformed JSON in `LoadIndex`
- ⚠️ Silent failure in unit type parsing during index load
- ⚠️ Potential data corruption or incorrect diffing due to reliance on `domain.Unit.ID.String()` as a unique key
- 🔗 The `LoadIndex` function introduces a failure propagation risk by silently ignoring unit type parsing errors, which could lead to invalid or corrupted index data being used downstream
- 🔗 The `Units()` method exposes internal state, increasing coupling and reducing encapsulation; any consumer can mutate the index directly

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
- 🤖 The code is functionally correct but has critical error handling issues, lacks proper validation, and introduces potential security and performance risks.
- 💡 Replace `ut, _ := domain.ParseUnitType(e.Type)` with proper error handling to ensure invalid types are not silently accepted
- 💡 Validate that `domain.Unit.ID.String()` is a safe and stable key before using it in maps for diffing, or use a more robust identifier like a hash
- 💡 Add input validation and sanitization for file paths in `Save` to prevent directory traversal or overwriting sensitive files
- 💡 Make `Units()` return a copy of the slice instead of the direct reference to prevent external mutation
- 💡 Add unit tests for edge cases such as nil inputs to `Diff` and malformed JSON in `LoadIndex`
- ⚠️ Silent failure in unit type parsing during index load
- ⚠️ Potential data corruption or incorrect diffing due to reliance on `domain.Unit.ID.String()` as a unique key
- 🔗 The `LoadIndex` function introduces a failure propagation risk by silently ignoring unit type parsing errors, which could lead to invalid or corrupted index data being used downstream
- 🔗 The `Units()` method exposes internal state, increasing coupling and reducing encapsulation; any consumer can mutate the index directly

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
- 🤖 Well-structured tests with good coverage of core functionality, but lacks assertions for order preservation and has potential issues with directory creation behavior.
- 💡 Replace string comparison in `TestIndex_Roundtrip_PreservesOrder` with direct unit list comparison to ensure actual order preservation
- 💡 Add explicit assertions in `TestIndex_FileCreatesDirectory` to verify that intermediate directories are created before attempting to write the file
- 💡 Add a test case for `TestDiff` that verifies complex diff scenarios involving multiple added/removed units and mixed changes
- 💡 Ensure all file system interactions in tests use `t.TempDir()` consistently to avoid test interference and improve reliability
- ⚠️ Order preservation in JSON roundtrip not fully validated due to string comparison instead of checking actual unit list order
- ⚠️ Directory creation behavior in `TestIndex_FileCreatesDirectory` is not thoroughly tested for concurrent access or permission issues
- 🔗 This test suite directly impacts the reliability of the discovery index component by ensuring correctness of save/load and diff operations
- 🔗 The tests do not isolate the unit under test from external filesystem behavior, which could lead to flaky tests in environments with restricted permissions or concurrent access

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
- 🤖 The Merge function has a logic flaw in type prioritization and lacks proper handling of concurrent access patterns; the DeduplicateFileLevel function is correct but could benefit from performance optimization.
- 💡 Ensure u.Type is an enum or has a defined precedence before comparing directly in Merge; consider using a priority map or switch statement to enforce correct ordering
- 💡 Use a stable sorting mechanism after collecting results from Merge to ensure deterministic output order, or sort the keys before iterating over seen map
- 💡 Add nil checks for u.ID.Path() and u.ID.Symbol() in DeduplicateFileLevel to prevent runtime panics
- 💡 Consider precomputing the path keys in DeduplicateFileLevel to avoid repeated calls to u.ID.Path()
- ⚠️ Incorrect type prioritization in Merge due to improper comparison of u.Type
- ⚠️ Unstable result ordering in Merge because map iteration order is not deterministic
- 🔗 The Merge function's incorrect type comparison can lead to wrong deduplication decisions and affect downstream processing logic that depends on unit type precedence
- 🔗 The lack of stable ordering in Merge results can introduce flaky behavior in systems that expect consistent output

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
- 🤖 The Merge function has a logic flaw in type prioritization and lacks proper handling of concurrent access patterns; the DeduplicateFileLevel function is correct but could benefit from performance optimization.
- 💡 Ensure u.Type is an enum or has a defined precedence before comparing directly in Merge; consider using a priority map or switch statement to enforce correct ordering
- 💡 Use a stable sorting mechanism after collecting results from Merge to ensure deterministic output order, or sort the keys before iterating over seen map
- 💡 Add nil checks for u.ID.Path() and u.ID.Symbol() in DeduplicateFileLevel to prevent runtime panics
- 💡 Consider precomputing the path keys in DeduplicateFileLevel to avoid repeated calls to u.ID.Path()
- ⚠️ Incorrect type prioritization in Merge due to improper comparison of u.Type
- ⚠️ Unstable result ordering in Merge because map iteration order is not deterministic
- 🔗 The Merge function's incorrect type comparison can lead to wrong deduplication decisions and affect downstream processing logic that depends on unit type precedence
- 🔗 The lack of stable ordering in Merge results can introduce flaky behavior in systems that expect consistent output

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
- 🤖 The Merge function has a logic flaw in type prioritization and lacks proper handling of concurrent access patterns; the DeduplicateFileLevel function is correct but could benefit from performance optimization.
- 💡 Ensure u.Type is an enum or has a defined precedence before comparing directly in Merge; consider using a priority map or switch statement to enforce correct ordering
- 💡 Use a stable sorting mechanism after collecting results from Merge to ensure deterministic output order, or sort the keys before iterating over seen map
- 💡 Add nil checks for u.ID.Path() and u.ID.Symbol() in DeduplicateFileLevel to prevent runtime panics
- 💡 Consider precomputing the path keys in DeduplicateFileLevel to avoid repeated calls to u.ID.Path()
- ⚠️ Incorrect type prioritization in Merge due to improper comparison of u.Type
- ⚠️ Unstable result ordering in Merge because map iteration order is not deterministic
- 🔗 The Merge function's incorrect type comparison can lead to wrong deduplication decisions and affect downstream processing logic that depends on unit type precedence
- 🔗 The lack of stable ordering in Merge results can introduce flaky behavior in systems that expect consistent output

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
- 🤖 The Merge function has a logic flaw in type prioritization and lacks proper handling of concurrent access patterns; the DeduplicateFileLevel function is correct but could benefit from performance optimization.
- 💡 Ensure u.Type is an enum or has a defined precedence before comparing directly in Merge; consider using a priority map or switch statement to enforce correct ordering
- 💡 Use a stable sorting mechanism after collecting results from Merge to ensure deterministic output order, or sort the keys before iterating over seen map
- 💡 Add nil checks for u.ID.Path() and u.ID.Symbol() in DeduplicateFileLevel to prevent runtime panics
- 💡 Consider precomputing the path keys in DeduplicateFileLevel to avoid repeated calls to u.ID.Path()
- ⚠️ Incorrect type prioritization in Merge due to improper comparison of u.Type
- ⚠️ Unstable result ordering in Merge because map iteration order is not deterministic
- 🔗 The Merge function's incorrect type comparison can lead to wrong deduplication decisions and affect downstream processing logic that depends on unit type precedence
- 🔗 The lack of stable ordering in Merge results can introduce flaky behavior in systems that expect consistent output

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
- 🤖 The test suite for the GenericScanner is comprehensive but has several correctness and maintainability issues, including brittle path matching logic, missing test coverage for edge cases, and insufficient validation of scanner behavior.
- 💡 Fix the path matching in TestGenericScanner_ExcludePatterns to use `strings.HasPrefix(path, "internal/")` or similar robust logic instead of checking only immediate parent directories
- 💡 Replace hardcoded build output directory checks in TestGenericScanner_SkipsBuildOutputDirs with assertions that validate the scanner's actual exclusion behavior based on its configuration
- 💡 Add assertions in TestGenericScanner_SkipsNonCertifiable to verify that each discovered unit's language and file type match expectations
- 💡 Add a test case for scanning a directory with no readable files to ensure proper error handling
- ⚠️ Flawed exclusion pattern matching logic in TestGenericScanner_ExcludePatterns that fails for deeply nested paths
- ⚠️ Hardcoded build output directory prefixes in TestGenericScanner_SkipsBuildOutputDirs that don't reflect actual scanner behavior
- 🔗 The test suite's flawed path matching logic could lead to incorrect assumptions about scanner filtering behavior, causing downstream issues in systems that depend on accurate file discovery
- 🔗 Lack of validation for certifiable vs non-certifiable file filtering in TestGenericScanner_SkipsNonCertifiable may mask bugs in the scanner's language or file type detection logic

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
- 🤖 The code correctly implements a TypeScript symbol scanner using regex but has critical correctness issues in symbol classification and lacks proper error handling for file parsing.
- 💡 Change `tsInterfaceRE`, `tsTypeRE`, and `tsEnumRE` matches to use appropriate unit types (e.g., `domain.UnitTypeInterface`, `domain.UnitTypeType`, `domain.UnitTypeEnum`) instead of `domain.UnitTypeClass`
- 💡 Change `tsConstRE` matches to use `domain.UnitTypeConstant` or similar instead of incorrectly classifying them as functions
- 💡 Implement multi-line scanning logic to handle declarations that span across multiple lines
- 💡 Add proper error logging or reporting when `parseFile` fails to help with debugging and monitoring
- 💡 Consider using a proper TypeScript parser (like AST-based tools) instead of regex for more accurate symbol detection
- ⚠️ Incorrect symbol classification leading to semantic misrepresentation in downstream systems
- ⚠️ False positives/negatives in symbol detection due to fragile regex parsing and lack of multi-line support
- 🔗 This unit's incorrect classification of interfaces/types/enums as classes affects downstream systems that rely on accurate symbol types for code analysis or documentation generation
- 🔗 The unit introduces coupling to domain.UnitTypeClass and domain.UnitTypeFunction with incorrect semantic mappings, making it hard to extend or maintain

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
- 🤖 The code correctly implements a TypeScript symbol scanner using regex but has critical correctness issues in symbol classification and lacks proper error handling for file parsing.
- 💡 Change `tsInterfaceRE`, `tsTypeRE`, and `tsEnumRE` matches to use appropriate unit types (e.g., `domain.UnitTypeInterface`, `domain.UnitTypeType`, `domain.UnitTypeEnum`) instead of `domain.UnitTypeClass`
- 💡 Change `tsConstRE` matches to use `domain.UnitTypeConstant` or similar instead of incorrectly classifying them as functions
- 💡 Implement multi-line scanning logic to handle declarations that span across multiple lines
- 💡 Add proper error logging or reporting when `parseFile` fails to help with debugging and monitoring
- 💡 Consider using a proper TypeScript parser (like AST-based tools) instead of regex for more accurate symbol detection
- ⚠️ Incorrect symbol classification leading to semantic misrepresentation in downstream systems
- ⚠️ False positives/negatives in symbol detection due to fragile regex parsing and lack of multi-line support
- 🔗 This unit's incorrect classification of interfaces/types/enums as classes affects downstream systems that rely on accurate symbol types for code analysis or documentation generation
- 🔗 The unit introduces coupling to domain.UnitTypeClass and domain.UnitTypeFunction with incorrect semantic mappings, making it hard to extend or maintain

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
- 🤖 The code correctly implements a TypeScript symbol scanner using regex but has critical correctness issues in symbol classification and lacks proper error handling for file parsing.
- 💡 Change `tsInterfaceRE`, `tsTypeRE`, and `tsEnumRE` matches to use appropriate unit types (e.g., `domain.UnitTypeInterface`, `domain.UnitTypeType`, `domain.UnitTypeEnum`) instead of `domain.UnitTypeClass`
- 💡 Change `tsConstRE` matches to use `domain.UnitTypeConstant` or similar instead of incorrectly classifying them as functions
- 💡 Implement multi-line scanning logic to handle declarations that span across multiple lines
- 💡 Add proper error logging or reporting when `parseFile` fails to help with debugging and monitoring
- 💡 Consider using a proper TypeScript parser (like AST-based tools) instead of regex for more accurate symbol detection
- ⚠️ Incorrect symbol classification leading to semantic misrepresentation in downstream systems
- ⚠️ False positives/negatives in symbol detection due to fragile regex parsing and lack of multi-line support
- 🔗 This unit's incorrect classification of interfaces/types/enums as classes affects downstream systems that rely on accurate symbol types for code analysis or documentation generation
- 🔗 The unit introduces coupling to domain.UnitTypeClass and domain.UnitTypeFunction with incorrect semantic mappings, making it hard to extend or maintain

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
- 🤖 The code correctly implements a TypeScript symbol scanner using regex but has critical correctness issues in symbol classification and lacks proper error handling for file parsing.
- 💡 Change `tsInterfaceRE`, `tsTypeRE`, and `tsEnumRE` matches to use appropriate unit types (e.g., `domain.UnitTypeInterface`, `domain.UnitTypeType`, `domain.UnitTypeEnum`) instead of `domain.UnitTypeClass`
- 💡 Change `tsConstRE` matches to use `domain.UnitTypeConstant` or similar instead of incorrectly classifying them as functions
- 💡 Implement multi-line scanning logic to handle declarations that span across multiple lines
- 💡 Add proper error logging or reporting when `parseFile` fails to help with debugging and monitoring
- 💡 Consider using a proper TypeScript parser (like AST-based tools) instead of regex for more accurate symbol detection
- ⚠️ Incorrect symbol classification leading to semantic misrepresentation in downstream systems
- ⚠️ False positives/negatives in symbol detection due to fragile regex parsing and lack of multi-line support
- 🔗 This unit's incorrect classification of interfaces/types/enums as classes affects downstream systems that rely on accurate symbol types for code analysis or documentation generation
- 🔗 The unit introduces coupling to domain.UnitTypeClass and domain.UnitTypeFunction with incorrect semantic mappings, making it hard to extend or maintain

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
- 🤖 The test suite for TSAdapter lacks proper error handling, has brittle assertions, and does not validate the content or correctness of scanned units beyond existence.
- 💡 Add assertions to verify unit metadata (e.g., type, parameters, return types) in addition to presence checks
- 💡 Replace `t.Fatalf` with `t.Fatal` for consistency and ensure all tests handle errors uniformly
- 💡 Sort unit slices before comparing IDs in `TestTSAdapter_StableIDs` to avoid flaky failures due to ordering
- 💡 Use a mock or isolated test repo instead of relying on `repoPath("ts-simple")` to make tests more reliable and self-contained
- ⚠️ Brittle test assertions that do not validate unit content or metadata correctness
- ⚠️ Inconsistent error handling (`t.Fatalf` vs `t.Fatal`) across tests
- 🔗 Tests do not validate the actual correctness of parsed units, potentially masking bugs in the discovery logic
- 🔗 Reliance on external repo paths without proper isolation or mocks increases coupling to filesystem state

</details>

### `internal/domain/` (69 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`AgentConfig`](reports/internal-domain-config-go-agentconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`AnalyzerConfig`](reports/internal-domain-config-go-analyzerconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`CertificationMode`](reports/internal-domain-config-go-certificationmode.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Config`](reports/internal-domain-config-go-config.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`DefaultConfig`](reports/internal-domain-config-go-defaultconfig.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`EnforcingConfig`](reports/internal-domain-config-go-enforcingconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ExpiryConfig`](reports/internal-domain-config-go-expiryconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`IssueConfig`](reports/internal-domain-config-go-issueconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ModelAssignments`](reports/internal-domain-config-go-modelassignments.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`PolicyConfig`](reports/internal-domain-config-go-policyconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ProviderConfig`](reports/internal-domain-config-go-providerconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`RateLimitConfig`](reports/internal-domain-config-go-ratelimitconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ScheduleConfig`](reports/internal-domain-config-go-scheduleconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ScopeConfig`](reports/internal-domain-config-go-scopeconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`SignoffConfig`](reports/internal-domain-config-go-signoffconfig.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`String`](reports/internal-domain-config-go-string.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`config_test.go`](reports/internal-domain-config-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`AllDimensions`](reports/internal-domain-dimension-go-alldimensions.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Dimension`](reports/internal-domain-dimension-go-dimension.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`DimensionScores`](reports/internal-domain-dimension-go-dimensionscores.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`DimensionWeights`](reports/internal-domain-dimension-go-dimensionweights.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Grade`](reports/internal-domain-dimension-go-grade.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`GradeFromScore`](reports/internal-domain-dimension-go-gradefromscore.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`String`](reports/internal-domain-dimension-go-string.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`WeightedAverage`](reports/internal-domain-dimension-go-weightedaverage.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`dimension_test.go`](reports/internal-domain-dimension-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Evidence`](reports/internal-domain-evidence-go-evidence.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`EvidenceKind`](reports/internal-domain-evidence-go-evidencekind.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseSeverity`](reports/internal-domain-evidence-go-parseseverity.md) | function | B+ | 87.8% | certified | 2026-04-24 |
| [`Severity`](reports/internal-domain-evidence-go-severity.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`String`](reports/internal-domain-evidence-go-string.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](reports/internal-domain-evidence-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`evidence_test.go`](reports/internal-domain-evidence-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Duration`](reports/internal-domain-expiry-go-duration.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ExpiryFactors`](reports/internal-domain-expiry-go-expiryfactors.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ExpiryWindow`](reports/internal-domain-expiry-go-expirywindow.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`IsExpired`](reports/internal-domain-expiry-go-isexpired.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`RemainingAt`](reports/internal-domain-expiry-go-remainingat.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`expiry_test.go`](reports/internal-domain-expiry-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Override`](reports/internal-domain-override-go-override.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`OverrideAction`](reports/internal-domain-override-go-overrideaction.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`String`](reports/internal-domain-override-go-string.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Validate`](reports/internal-domain-override-go-validate.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`override_test.go`](reports/internal-domain-override-test-go.md) | file | B+ | 89.4% | certified | 2026-04-23 |
| [`IsGlobal`](reports/internal-domain-policy-go-isglobal.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`PolicyPack`](reports/internal-domain-policy-go-policypack.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`PolicyRule`](reports/internal-domain-policy-go-policyrule.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Violation`](reports/internal-domain-policy-go-violation.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`policy_test.go`](reports/internal-domain-policy-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`CertificationRecord`](reports/internal-domain-record-go-certificationrecord.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`IsPassing`](reports/internal-domain-record-go-ispassing.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseStatus`](reports/internal-domain-record-go-parsestatus.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Status`](reports/internal-domain-record-go-status.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`String`](reports/internal-domain-record-go-string.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](reports/internal-domain-record-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`record_test.go`](reports/internal-domain-record-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Language`](reports/internal-domain-unit-go-language.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewUnit`](reports/internal-domain-unit-go-newunit.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`NewUnitID`](reports/internal-domain-unit-go-newunitid.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseUnitID`](reports/internal-domain-unit-go-parseunitid.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseUnitType`](reports/internal-domain-unit-go-parseunittype.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Path`](reports/internal-domain-unit-go-path.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`String`](reports/internal-domain-unit-go-string.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Symbol`](reports/internal-domain-unit-go-symbol.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Unit`](reports/internal-domain-unit-go-unit.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`UnitID`](reports/internal-domain-unit-go-unitid.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`UnitType`](reports/internal-domain-unit-go-unittype.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`init`](reports/internal-domain-unit-go-init.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`unit_test.go`](reports/internal-domain-unit-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 Well-structured configuration types with clear intent but missing validation and potential runtime safety issues in ExpiryConfig.
- 💡 Add validation logic to ExpiryConfig to ensure DefaultWindowDays is between MinWindowDays and MaxWindowDays, and that all three fields are positive integers
- 💡 Add unit tests for DefaultConfig() to verify that ExpiryConfig defaults are valid and consistent
- 💡 Consider adding a Validate() method or function to Config that checks the integrity of all nested config types, particularly ExpiryConfig
- ⚠️ Missing validation in ExpiryConfig fields (DefaultWindowDays, MinWindowDays, MaxWindowDays) can lead to invalid or inconsistent configuration
- ⚠️ No runtime validation or input sanitization for ExpiryConfig, which may cause unexpected behavior if values are out of range
- 🔗 ExpiryConfig is part of the top-level Config struct and affects how certification windows are managed, so incorrect values could impact merge blocking logic
- 🔗 The Config struct is likely used throughout the system for policy enforcement and scheduling, making it a high-coupling point with downstream components

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
- 🤖 The test file has minimal coverage and lacks comprehensive validation of configuration behavior, with several test cases that only check field values without asserting functional correctness or edge case handling.
- 💡 Add integration-style tests that simulate actual usage of `domain.DefaultConfig()` in a real-world scenario (e.g., loading from environment or config file) to ensure defaults are correctly applied
- 💡 Implement validation logic in `domain.ModelAssignments` to check that model names conform to expected format (e.g., regex pattern) and add corresponding tests
- 💡 Add a test case to `TestConfig_ScopePatterns` that validates actual filtering behavior using a mock file system or pattern matcher to ensure include/exclude logic works as expected
- 💡 Add test cases for edge cases such as empty lists, nil pointers, or invalid enum values in `CertificationMode.String()` to ensure robustness
- ⚠️ Missing validation of model name formats in `TestAgentConfig_ModelAssignments` - could allow invalid or malformed model identifiers to pass through
- ⚠️ Lack of edge case testing for scope patterns in `TestConfig_ScopePatterns` - no checks for invalid glob patterns or overlapping include/exclude rules
- 🔗 The tests do not validate that configurations can be safely used in concurrent contexts or that defaults are properly applied across different environments
- 🔗 The test suite does not verify that configuration changes propagate correctly to dependent systems or that invalid configurations are rejected early

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
- 🤖 Well-structured domain types with clear intent but lacks input validation and has a hardcoded dimension list that's not easily extensible.
- 💡 Replace hardcoded `AllDimensions()` with a dynamically built list using reflection or a registry pattern to support extensibility
- 💡 Extract grade thresholds into named constants to improve readability and reduce magic number risks in `GradeFromScore`
- 💡 Add validation or documentation to ensure that all dimensions used in `WeightedAverage` are properly scored and weighted, or handle missing values explicitly
- 💡 Consider adding unit tests for edge cases such as zero weights, missing scores, or empty maps in `WeightedAverage`
- ⚠️ Hardcoded dimension list in `AllDimensions()` makes it brittle to future additions or removals
- ⚠️ Magic numbers in `GradeFromScore` reduce maintainability and increase risk of incorrect grading
- 🔗 The `AllDimensions()` function tightly couples the dimension list to a fixed set, making it difficult to extend or dynamically manage dimensions in larger systems
- 🔗 The `WeightedAverage` function assumes all dimensions are scored and weighted, which can lead to silent failures or incorrect averages if some dimensions are missing from either input map

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
- 🤖 Well-structured domain types with clear intent but lacks input validation and has a hardcoded dimension list that's not easily extensible.
- 💡 Replace hardcoded `AllDimensions()` with a dynamically built list using reflection or a registry pattern to support extensibility
- 💡 Extract grade thresholds into named constants to improve readability and reduce magic number risks in `GradeFromScore`
- 💡 Add validation or documentation to ensure that all dimensions used in `WeightedAverage` are properly scored and weighted, or handle missing values explicitly
- 💡 Consider adding unit tests for edge cases such as zero weights, missing scores, or empty maps in `WeightedAverage`
- ⚠️ Hardcoded dimension list in `AllDimensions()` makes it brittle to future additions or removals
- ⚠️ Magic numbers in `GradeFromScore` reduce maintainability and increase risk of incorrect grading
- 🔗 The `AllDimensions()` function tightly couples the dimension list to a fixed set, making it difficult to extend or dynamically manage dimensions in larger systems
- 🔗 The `WeightedAverage` function assumes all dimensions are scored and weighted, which can lead to silent failures or incorrect averages if some dimensions are missing from either input map

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
- 🤖 Well-structured domain types with clear intent but lacks input validation and has a hardcoded dimension list that's not easily extensible.
- 💡 Replace hardcoded `AllDimensions()` with a dynamically built list using reflection or a registry pattern to support extensibility
- 💡 Extract grade thresholds into named constants to improve readability and reduce magic number risks in `GradeFromScore`
- 💡 Add validation or documentation to ensure that all dimensions used in `WeightedAverage` are properly scored and weighted, or handle missing values explicitly
- 💡 Consider adding unit tests for edge cases such as zero weights, missing scores, or empty maps in `WeightedAverage`
- ⚠️ Hardcoded dimension list in `AllDimensions()` makes it brittle to future additions or removals
- ⚠️ Magic numbers in `GradeFromScore` reduce maintainability and increase risk of incorrect grading
- 🔗 The `AllDimensions()` function tightly couples the dimension list to a fixed set, making it difficult to extend or dynamically manage dimensions in larger systems
- 🔗 The `WeightedAverage` function assumes all dimensions are scored and weighted, which can lead to silent failures or incorrect averages if some dimensions are missing from either input map

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
- 🤖 Well-structured domain types with clear intent but lacks input validation and has a hardcoded dimension list that's not easily extensible.
- 💡 Replace hardcoded `AllDimensions()` with a dynamically built list using reflection or a registry pattern to support extensibility
- 💡 Extract grade thresholds into named constants to improve readability and reduce magic number risks in `GradeFromScore`
- 💡 Add validation or documentation to ensure that all dimensions used in `WeightedAverage` are properly scored and weighted, or handle missing values explicitly
- 💡 Consider adding unit tests for edge cases such as zero weights, missing scores, or empty maps in `WeightedAverage`
- ⚠️ Hardcoded dimension list in `AllDimensions()` makes it brittle to future additions or removals
- ⚠️ Magic numbers in `GradeFromScore` reduce maintainability and increase risk of incorrect grading
- 🔗 The `AllDimensions()` function tightly couples the dimension list to a fixed set, making it difficult to extend or dynamically manage dimensions in larger systems
- 🔗 The `WeightedAverage` function assumes all dimensions are scored and weighted, which can lead to silent failures or incorrect averages if some dimensions are missing from either input map

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
- 🤖 Well-structured domain types with clear intent but lacks input validation and has a hardcoded dimension list that's not easily extensible.
- 💡 Replace hardcoded `AllDimensions()` with a dynamically built list using reflection or a registry pattern to support extensibility
- 💡 Extract grade thresholds into named constants to improve readability and reduce magic number risks in `GradeFromScore`
- 💡 Add validation or documentation to ensure that all dimensions used in `WeightedAverage` are properly scored and weighted, or handle missing values explicitly
- 💡 Consider adding unit tests for edge cases such as zero weights, missing scores, or empty maps in `WeightedAverage`
- ⚠️ Hardcoded dimension list in `AllDimensions()` makes it brittle to future additions or removals
- ⚠️ Magic numbers in `GradeFromScore` reduce maintainability and increase risk of incorrect grading
- 🔗 The `AllDimensions()` function tightly couples the dimension list to a fixed set, making it difficult to extend or dynamically manage dimensions in larger systems
- 🔗 The `WeightedAverage` function assumes all dimensions are scored and weighted, which can lead to silent failures or incorrect averages if some dimensions are missing from either input map

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
- 🤖 Well-structured domain types with clear intent but lacks input validation and has a hardcoded dimension list that's not easily extensible.
- 💡 Replace hardcoded `AllDimensions()` with a dynamically built list using reflection or a registry pattern to support extensibility
- 💡 Extract grade thresholds into named constants to improve readability and reduce magic number risks in `GradeFromScore`
- 💡 Add validation or documentation to ensure that all dimensions used in `WeightedAverage` are properly scored and weighted, or handle missing values explicitly
- 💡 Consider adding unit tests for edge cases such as zero weights, missing scores, or empty maps in `WeightedAverage`
- ⚠️ Hardcoded dimension list in `AllDimensions()` makes it brittle to future additions or removals
- ⚠️ Magic numbers in `GradeFromScore` reduce maintainability and increase risk of incorrect grading
- 🔗 The `AllDimensions()` function tightly couples the dimension list to a fixed set, making it difficult to extend or dynamically manage dimensions in larger systems
- 🔗 The `WeightedAverage` function assumes all dimensions are scored and weighted, which can lead to silent failures or incorrect averages if some dimensions are missing from either input map

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
- 🤖 Well-structured domain types with clear intent but lacks input validation and has a hardcoded dimension list that's not easily extensible.
- 💡 Replace hardcoded `AllDimensions()` with a dynamically built list using reflection or a registry pattern to support extensibility
- 💡 Extract grade thresholds into named constants to improve readability and reduce magic number risks in `GradeFromScore`
- 💡 Add validation or documentation to ensure that all dimensions used in `WeightedAverage` are properly scored and weighted, or handle missing values explicitly
- 💡 Consider adding unit tests for edge cases such as zero weights, missing scores, or empty maps in `WeightedAverage`
- ⚠️ Hardcoded dimension list in `AllDimensions()` makes it brittle to future additions or removals
- ⚠️ Magic numbers in `GradeFromScore` reduce maintainability and increase risk of incorrect grading
- 🔗 The `AllDimensions()` function tightly couples the dimension list to a fixed set, making it difficult to extend or dynamically manage dimensions in larger systems
- 🔗 The `WeightedAverage` function assumes all dimensions are scored and weighted, which can lead to silent failures or incorrect averages if some dimensions are missing from either input map

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
- 🤖 Well-structured domain types with clear intent but lacks input validation and has a hardcoded dimension list that's not easily extensible.
- 💡 Replace hardcoded `AllDimensions()` with a dynamically built list using reflection or a registry pattern to support extensibility
- 💡 Extract grade thresholds into named constants to improve readability and reduce magic number risks in `GradeFromScore`
- 💡 Add validation or documentation to ensure that all dimensions used in `WeightedAverage` are properly scored and weighted, or handle missing values explicitly
- 💡 Consider adding unit tests for edge cases such as zero weights, missing scores, or empty maps in `WeightedAverage`
- ⚠️ Hardcoded dimension list in `AllDimensions()` makes it brittle to future additions or removals
- ⚠️ Magic numbers in `GradeFromScore` reduce maintainability and increase risk of incorrect grading
- 🔗 The `AllDimensions()` function tightly couples the dimension list to a fixed set, making it difficult to extend or dynamically manage dimensions in larger systems
- 🔗 The `WeightedAverage` function assumes all dimensions are scored and weighted, which can lead to silent failures or incorrect averages if some dimensions are missing from either input map

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
- 🤖 The tests are well-structured and cover the expected behavior of domain types, but lack validation for edge cases and error conditions in weighted averages.
- 💡 Replace tolerance-based floating-point checks with a more robust comparison using math.Abs or a custom epsilon function
- 💡 Add tests for edge cases such as zero scores, zero weights, and negative values in WeightedAverage logic
- 💡 Use a constant or reflection-based check instead of hardcoding the expected list in TestDimension_AllNine
- 💡 Fix the format string in TestGrade_String to use %v instead of %d for clarity
- ⚠️ Floating-point comparison with tolerance may cause flaky tests in different environments
- ⚠️ Missing validation for invalid inputs (negative scores, weights) in weighted average logic
- 🔗 These tests directly validate the correctness of domain types and their transformations, which are core to the certification logic
- 🔗 If these tests pass but don't cover edge cases, incorrect behavior in downstream systems using the domain types could go undetected

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
- 🤖 The ParseSeverity function has a critical runtime bug due to incorrect map initialization, and the code lacks test coverage for key logic paths.
- 💡 Fix the `init()` function to correctly populate `stringToSeverity` by swapping key and value: `stringToSeverity[v] = k` should be `stringToSeverity[v] = k` (but the current logic is backwards). Correctly assign keys and values so that `stringToSeverity["info"]` resolves to `SeverityInfo`.
- 💡 Replace `any` with a concrete type or interface for `Details` to improve type safety and maintainability.
- ⚠️ Runtime panic or incorrect severity parsing due to flawed map initialization in `init()`
- ⚠️ Type safety degradation with use of `any` for `Details` field
- 🔗 The `ParseSeverity` function is part of the policy evaluation system and may misinterpret severity levels, leading to incorrect policy enforcement decisions
- 🔗 The use of `any` in `Evidence.Details` increases coupling between components and makes refactoring more difficult

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
- 🤖 The ParseSeverity function has a critical runtime bug due to incorrect map initialization, and the code lacks test coverage for key logic paths.
- 💡 Fix the `init()` function to correctly populate `stringToSeverity` by swapping key and value: `stringToSeverity[v] = k` should be `stringToSeverity[v] = k` (but the current logic is backwards). Correctly assign keys and values so that `stringToSeverity["info"]` resolves to `SeverityInfo`.
- 💡 Replace `any` with a concrete type or interface for `Details` to improve type safety and maintainability.
- ⚠️ Runtime panic or incorrect severity parsing due to flawed map initialization in `init()`
- ⚠️ Type safety degradation with use of `any` for `Details` field
- 🔗 The `ParseSeverity` function is part of the policy evaluation system and may misinterpret severity levels, leading to incorrect policy enforcement decisions
- 🔗 The use of `any` in `Evidence.Details` increases coupling between components and makes refactoring more difficult

</details>

<a id="internal-domain-evidence-go-parseseverity"></a>
<details>
<summary>ParseSeverity — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 95.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

<a id="internal-domain-evidence-go-severity"></a>
<details>
<summary>Severity — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 95.0% |
| security | 80.0% |
| testability | 90.0% |

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
- 🤖 The ParseSeverity function has a critical runtime bug due to incorrect map initialization, and the code lacks test coverage for key logic paths.
- 💡 Fix the `init()` function to correctly populate `stringToSeverity` by swapping key and value: `stringToSeverity[v] = k` should be `stringToSeverity[v] = k` (but the current logic is backwards). Correctly assign keys and values so that `stringToSeverity["info"]` resolves to `SeverityInfo`.
- 💡 Replace `any` with a concrete type or interface for `Details` to improve type safety and maintainability.
- ⚠️ Runtime panic or incorrect severity parsing due to flawed map initialization in `init()`
- ⚠️ Type safety degradation with use of `any` for `Details` field
- 🔗 The `ParseSeverity` function is part of the policy evaluation system and may misinterpret severity levels, leading to incorrect policy enforcement decisions
- 🔗 The use of `any` in `Evidence.Details` increases coupling between components and makes refactoring more difficult

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
- 🤖 The ParseSeverity function has a critical runtime bug due to incorrect map initialization, and the code lacks test coverage for key logic paths.
- 💡 Fix the `init()` function to correctly populate `stringToSeverity` by swapping key and value: `stringToSeverity[v] = k` should be `stringToSeverity[v] = k` (but the current logic is backwards). Correctly assign keys and values so that `stringToSeverity["info"]` resolves to `SeverityInfo`.
- 💡 Replace `any` with a concrete type or interface for `Details` to improve type safety and maintainability.
- ⚠️ Runtime panic or incorrect severity parsing due to flawed map initialization in `init()`
- ⚠️ Type safety degradation with use of `any` for `Details` field
- 🔗 The `ParseSeverity` function is part of the policy evaluation system and may misinterpret severity levels, leading to incorrect policy enforcement decisions
- 🔗 The use of `any` in `Evidence.Details` increases coupling between components and makes refactoring more difficult

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
- 🤖 The test file has good coverage for string representations and parsing of evidence kinds and severity, but lacks comprehensive validation for edge cases and error conditions.
- 💡 Change t.Errorf format strings from %d to %v or use proper enum name formatting for clarity
- 💡 Add test cases for ParseSeverity with uppercase/lowercase inputs and whitespace to ensure robust parsing
- 💡 Add test case for Evidence.Missing = false with Passed = true to verify mutually exclusive states
- 💡 Add test case for invalid EvidenceKind values to ensure proper handling of unknown enum values
- ⚠️ Improper error message formatting in t.Errorf calls that use %d for non-integer types
- ⚠️ Lack of validation for case-insensitive parsing in ParseSeverity
- 🔗 These tests don't directly affect system behavior but provide confidence in evidence handling logic
- 🔗 Weak test coverage may mask edge-case failures that could propagate to downstream systems

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
- 🤖 The code is functionally correct but has subtle issues with time handling, missing validation, and unclear semantics around expiration logic.
- 💡 Simplify `IsExpired` to use only `at.After(w.ExpiresAt)` since equality is already implied by the semantics of After.
- 💡 Add validation logic in `ExpiryFactors` to ensure that fields like `ChurnRate`, `TestCoverage`, and `Complexity` are within expected ranges (e.g., 0.0–1.0 for coverage and churn rate).
- 💡 Add unit tests covering edge cases such as when `CertifiedAt` is after `ExpiresAt`, or when both times are equal.
- 💡 Consider adding documentation to explain how `ExpiryFactors` fields contribute to computing an actual expiry window (e.g., whether they're used in a formula or just influence policy).
- 💡 Standardize JSON struct tag formatting (e.g., always use `json:"field_name"` without spaces).
- ⚠️ Unvalidated input in `ExpiryFactors` fields such as `ChurnRate`, `TestCoverage`, and `Complexity` may result in incorrect or undefined behavior during calculations.
- ⚠️ The use of `time.Time` fields without explicit timezone handling can introduce subtle bugs in systems that rely on UTC vs local time.
- 🔗 The `ExpiryWindow` type tightly couples to Go's time package and assumes UTC or local timezone semantics, which can cause issues in distributed systems or multi-timezone environments.
- 🔗 The `ExpiryFactors` struct introduces a new API surface that is not validated or constrained, increasing the risk of misuse in downstream logic.

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
- 🤖 The code is functionally correct but has subtle issues with time handling, missing validation, and unclear semantics around expiration logic.
- 💡 Simplify `IsExpired` to use only `at.After(w.ExpiresAt)` since equality is already implied by the semantics of After.
- 💡 Add validation logic in `ExpiryFactors` to ensure that fields like `ChurnRate`, `TestCoverage`, and `Complexity` are within expected ranges (e.g., 0.0–1.0 for coverage and churn rate).
- 💡 Add unit tests covering edge cases such as when `CertifiedAt` is after `ExpiresAt`, or when both times are equal.
- 💡 Consider adding documentation to explain how `ExpiryFactors` fields contribute to computing an actual expiry window (e.g., whether they're used in a formula or just influence policy).
- 💡 Standardize JSON struct tag formatting (e.g., always use `json:"field_name"` without spaces).
- ⚠️ Unvalidated input in `ExpiryFactors` fields such as `ChurnRate`, `TestCoverage`, and `Complexity` may result in incorrect or undefined behavior during calculations.
- ⚠️ The use of `time.Time` fields without explicit timezone handling can introduce subtle bugs in systems that rely on UTC vs local time.
- 🔗 The `ExpiryWindow` type tightly couples to Go's time package and assumes UTC or local timezone semantics, which can cause issues in distributed systems or multi-timezone environments.
- 🔗 The `ExpiryFactors` struct introduces a new API surface that is not validated or constrained, increasing the risk of misuse in downstream logic.

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
- 🤖 The code is functionally correct but has subtle issues with time handling, missing validation, and unclear semantics around expiration logic.
- 💡 Simplify `IsExpired` to use only `at.After(w.ExpiresAt)` since equality is already implied by the semantics of After.
- 💡 Add validation logic in `ExpiryFactors` to ensure that fields like `ChurnRate`, `TestCoverage`, and `Complexity` are within expected ranges (e.g., 0.0–1.0 for coverage and churn rate).
- 💡 Add unit tests covering edge cases such as when `CertifiedAt` is after `ExpiresAt`, or when both times are equal.
- 💡 Consider adding documentation to explain how `ExpiryFactors` fields contribute to computing an actual expiry window (e.g., whether they're used in a formula or just influence policy).
- 💡 Standardize JSON struct tag formatting (e.g., always use `json:"field_name"` without spaces).
- ⚠️ Unvalidated input in `ExpiryFactors` fields such as `ChurnRate`, `TestCoverage`, and `Complexity` may result in incorrect or undefined behavior during calculations.
- ⚠️ The use of `time.Time` fields without explicit timezone handling can introduce subtle bugs in systems that rely on UTC vs local time.
- 🔗 The `ExpiryWindow` type tightly couples to Go's time package and assumes UTC or local timezone semantics, which can cause issues in distributed systems or multi-timezone environments.
- 🔗 The `ExpiryFactors` struct introduces a new API surface that is not validated or constrained, increasing the risk of misuse in downstream logic.

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
- 🤖 The code is functionally correct but has subtle issues with time handling, missing validation, and unclear semantics around expiration logic.
- 💡 Simplify `IsExpired` to use only `at.After(w.ExpiresAt)` since equality is already implied by the semantics of After.
- 💡 Add validation logic in `ExpiryFactors` to ensure that fields like `ChurnRate`, `TestCoverage`, and `Complexity` are within expected ranges (e.g., 0.0–1.0 for coverage and churn rate).
- 💡 Add unit tests covering edge cases such as when `CertifiedAt` is after `ExpiresAt`, or when both times are equal.
- 💡 Consider adding documentation to explain how `ExpiryFactors` fields contribute to computing an actual expiry window (e.g., whether they're used in a formula or just influence policy).
- 💡 Standardize JSON struct tag formatting (e.g., always use `json:"field_name"` without spaces).
- ⚠️ Unvalidated input in `ExpiryFactors` fields such as `ChurnRate`, `TestCoverage`, and `Complexity` may result in incorrect or undefined behavior during calculations.
- ⚠️ The use of `time.Time` fields without explicit timezone handling can introduce subtle bugs in systems that rely on UTC vs local time.
- 🔗 The `ExpiryWindow` type tightly couples to Go's time package and assumes UTC or local timezone semantics, which can cause issues in distributed systems or multi-timezone environments.
- 🔗 The `ExpiryFactors` struct introduces a new API surface that is not validated or constrained, increasing the risk of misuse in downstream logic.

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
- 🤖 The code is functionally correct but has subtle issues with time handling, missing validation, and unclear semantics around expiration logic.
- 💡 Simplify `IsExpired` to use only `at.After(w.ExpiresAt)` since equality is already implied by the semantics of After.
- 💡 Add validation logic in `ExpiryFactors` to ensure that fields like `ChurnRate`, `TestCoverage`, and `Complexity` are within expected ranges (e.g., 0.0–1.0 for coverage and churn rate).
- 💡 Add unit tests covering edge cases such as when `CertifiedAt` is after `ExpiresAt`, or when both times are equal.
- 💡 Consider adding documentation to explain how `ExpiryFactors` fields contribute to computing an actual expiry window (e.g., whether they're used in a formula or just influence policy).
- 💡 Standardize JSON struct tag formatting (e.g., always use `json:"field_name"` without spaces).
- ⚠️ Unvalidated input in `ExpiryFactors` fields such as `ChurnRate`, `TestCoverage`, and `Complexity` may result in incorrect or undefined behavior during calculations.
- ⚠️ The use of `time.Time` fields without explicit timezone handling can introduce subtle bugs in systems that rely on UTC vs local time.
- 🔗 The `ExpiryWindow` type tightly couples to Go's time package and assumes UTC or local timezone semantics, which can cause issues in distributed systems or multi-timezone environments.
- 🔗 The `ExpiryFactors` struct introduces a new API surface that is not validated or constrained, increasing the risk of misuse in downstream logic.

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
- 🤖 The test suite for ExpiryWindow and ExpiryFactors is functionally correct but lacks comprehensive coverage for edge cases and error conditions.
- 💡 Add tests for edge cases: equal CertifiedAt and ExpiresAt, negative durations, and invalid time ordering
- 💡 Mock or stabilize time.Now() in tests to prevent flakiness when running under different system clocks
- 💡 Add a test for ExpiryFactors to validate computed properties or business logic if any exists
- ⚠️ Flaky tests due to reliance on time.Now() without mocking or stabilization
- ⚠️ Missing validation for invalid time ordering (ExpiresAt < CertifiedAt) in ExpiryWindow logic
- 🔗 These tests provide minimal confidence in the correctness of ExpiryWindow logic under edge cases
- 🔗 Lack of boundary condition testing increases risk of regressions in downstream systems that depend on accurate expiration handling

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
- 🤖 The code is functionally correct but has a subtle bug in string representation and lacks test coverage for edge cases.
- 💡 Add unit tests for `String()` method with out-of-range enum values to ensure consistent fallback behavior
- 💡 Add validation for `UnitID` and `Timestamp` in the `Validate()` method to ensure all required fields are present and valid
- 💡 Consider using a switch statement instead of a map lookup in `String()` for better performance and clarity
- ⚠️ Unvalidated enum values can lead to misleading string representations
- ⚠️ Missing validation of required fields like UnitID and Timestamp
- 🔗 The `String()` method affects logging and API responses where enum values are serialized to strings
- 🔗 The `Validate()` method's incomplete checks can cause downstream systems to process invalid overrides

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
- 🤖 The code is functionally correct but has a subtle bug in string representation and lacks test coverage for edge cases.
- 💡 Add unit tests for `String()` method with out-of-range enum values to ensure consistent fallback behavior
- 💡 Add validation for `UnitID` and `Timestamp` in the `Validate()` method to ensure all required fields are present and valid
- 💡 Consider using a switch statement instead of a map lookup in `String()` for better performance and clarity
- ⚠️ Unvalidated enum values can lead to misleading string representations
- ⚠️ Missing validation of required fields like UnitID and Timestamp
- 🔗 The `String()` method affects logging and API responses where enum values are serialized to strings
- 🔗 The `Validate()` method's incomplete checks can cause downstream systems to process invalid overrides

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
- 🤖 The code is functionally correct but has a subtle bug in string representation and lacks test coverage for edge cases.
- 💡 Add unit tests for `String()` method with out-of-range enum values to ensure consistent fallback behavior
- 💡 Add validation for `UnitID` and `Timestamp` in the `Validate()` method to ensure all required fields are present and valid
- 💡 Consider using a switch statement instead of a map lookup in `String()` for better performance and clarity
- ⚠️ Unvalidated enum values can lead to misleading string representations
- ⚠️ Missing validation of required fields like UnitID and Timestamp
- 🔗 The `String()` method affects logging and API responses where enum values are serialized to strings
- 🔗 The `Validate()` method's incomplete checks can cause downstream systems to process invalid overrides

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
- 🤖 The code is functionally correct but has a subtle bug in string representation and lacks test coverage for edge cases.
- 💡 Add unit tests for `String()` method with out-of-range enum values to ensure consistent fallback behavior
- 💡 Add validation for `UnitID` and `Timestamp` in the `Validate()` method to ensure all required fields are present and valid
- 💡 Consider using a switch statement instead of a map lookup in `String()` for better performance and clarity
- ⚠️ Unvalidated enum values can lead to misleading string representations
- ⚠️ Missing validation of required fields like UnitID and Timestamp
- 🔗 The `String()` method affects logging and API responses where enum values are serialized to strings
- 🔗 The `Validate()` method's incomplete checks can cause downstream systems to process invalid overrides

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
- 🤖 The test file has good coverage for basic validation logic but lacks comprehensive edge case testing and does not validate error message content.
- 💡 Add parameterized tests for all `OverrideAction` enum values in `TestOverrideAction_String` to ensure complete coverage
- 💡 Modify `TestOverride_RequiresRationale` and `TestOverride_RequiresActor` to assert specific error types (e.g., `errors.Is(err, ErrInvalidRationale)` or similar) rather than just checking for nil/non-nil errors
- 💡 Add a test case that validates the `Validate()` method returns an error when `UnitID` is invalid (e.g., malformed or empty unit ID)
- 💡 Add a test case that checks that `Validate()` properly handles invalid timestamps (e.g., future timestamps or zero time values)
- ⚠️ Missing validation of error message content - tests only check for presence of error but not correctness of error messages
- ⚠️ Incomplete validation coverage - does not test all enum values or edge cases for validation logic
- 🔗 This unit tests a domain model that directly affects the certification and override workflow; any missing validation can propagate to system-wide data integrity issues
- 🔗 The test suite does not adequately cover the full validation surface area, which increases risk of undetected bugs in production

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
- 🤖 The PolicyPack struct and related types are well-defined but lack validation, input sanitization, and error handling for critical fields like path patterns or rule thresholds.
- 💡 Add a `Validate()` method to `PolicyPack` that checks for valid `PathPatterns`, ensures `Rules` is not empty, and validates each rule's `ID` and `Severity`.
- 💡 Implement input validation for `PolicyRule.Threshold` to ensure it is a finite, non-negative number and that `Metric` is one of the allowed values or matches a known pattern.
- 💡 Add a constructor function like `NewPolicyPack(name, version string, rules []PolicyRule)` that ensures required fields are not empty and validates inputs.
- 💡 Add unit tests for `IsGlobal()` with edge cases like whitespace-only language strings to ensure robustness.
- ⚠️ Unvalidated glob patterns in `PathPatterns` can lead to unexpected behavior or injection vulnerabilities.
- ⚠️ Unvalidated `Threshold` in `PolicyRule` can cause runtime panics or incorrect policy evaluation.
- 🔗 The lack of validation in `PolicyPack` and `PolicyRule` makes this unit a potential failure point for downstream systems that rely on these policies.
- 🔗 The `IsGlobal()` method introduces an implicit contract that is not enforced or documented, increasing coupling between this type and consumers who assume empty language implies global scope.

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
- 🤖 The PolicyPack struct and related types are well-defined but lack validation, input sanitization, and error handling for critical fields like path patterns or rule thresholds.
- 💡 Add a `Validate()` method to `PolicyPack` that checks for valid `PathPatterns`, ensures `Rules` is not empty, and validates each rule's `ID` and `Severity`.
- 💡 Implement input validation for `PolicyRule.Threshold` to ensure it is a finite, non-negative number and that `Metric` is one of the allowed values or matches a known pattern.
- 💡 Add a constructor function like `NewPolicyPack(name, version string, rules []PolicyRule)` that ensures required fields are not empty and validates inputs.
- 💡 Add unit tests for `IsGlobal()` with edge cases like whitespace-only language strings to ensure robustness.
- ⚠️ Unvalidated glob patterns in `PathPatterns` can lead to unexpected behavior or injection vulnerabilities.
- ⚠️ Unvalidated `Threshold` in `PolicyRule` can cause runtime panics or incorrect policy evaluation.
- 🔗 The lack of validation in `PolicyPack` and `PolicyRule` makes this unit a potential failure point for downstream systems that rely on these policies.
- 🔗 The `IsGlobal()` method introduces an implicit contract that is not enforced or documented, increasing coupling between this type and consumers who assume empty language implies global scope.

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
- 🤖 The PolicyPack struct and related types are well-defined but lack validation, input sanitization, and error handling for critical fields like path patterns or rule thresholds.
- 💡 Add a `Validate()` method to `PolicyPack` that checks for valid `PathPatterns`, ensures `Rules` is not empty, and validates each rule's `ID` and `Severity`.
- 💡 Implement input validation for `PolicyRule.Threshold` to ensure it is a finite, non-negative number and that `Metric` is one of the allowed values or matches a known pattern.
- 💡 Add a constructor function like `NewPolicyPack(name, version string, rules []PolicyRule)` that ensures required fields are not empty and validates inputs.
- 💡 Add unit tests for `IsGlobal()` with edge cases like whitespace-only language strings to ensure robustness.
- ⚠️ Unvalidated glob patterns in `PathPatterns` can lead to unexpected behavior or injection vulnerabilities.
- ⚠️ Unvalidated `Threshold` in `PolicyRule` can cause runtime panics or incorrect policy evaluation.
- 🔗 The lack of validation in `PolicyPack` and `PolicyRule` makes this unit a potential failure point for downstream systems that rely on these policies.
- 🔗 The `IsGlobal()` method introduces an implicit contract that is not enforced or documented, increasing coupling between this type and consumers who assume empty language implies global scope.

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
- 🤖 The PolicyPack struct and related types are well-defined but lack validation, input sanitization, and error handling for critical fields like path patterns or rule thresholds.
- 💡 Add a `Validate()` method to `PolicyPack` that checks for valid `PathPatterns`, ensures `Rules` is not empty, and validates each rule's `ID` and `Severity`.
- 💡 Implement input validation for `PolicyRule.Threshold` to ensure it is a finite, non-negative number and that `Metric` is one of the allowed values or matches a known pattern.
- 💡 Add a constructor function like `NewPolicyPack(name, version string, rules []PolicyRule)` that ensures required fields are not empty and validates inputs.
- 💡 Add unit tests for `IsGlobal()` with edge cases like whitespace-only language strings to ensure robustness.
- ⚠️ Unvalidated glob patterns in `PathPatterns` can lead to unexpected behavior or injection vulnerabilities.
- ⚠️ Unvalidated `Threshold` in `PolicyRule` can cause runtime panics or incorrect policy evaluation.
- 🔗 The lack of validation in `PolicyPack` and `PolicyRule` makes this unit a potential failure point for downstream systems that rely on these policies.
- 🔗 The `IsGlobal()` method introduces an implicit contract that is not enforced or documented, increasing coupling between this type and consumers who assume empty language implies global scope.

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
- 🤖 The test file has basic coverage for policy domain structures but lacks comprehensive validation, proper error handling in tests, and does not cover edge cases or internal logic.
- 💡 Replace hardcoded assertions with table-driven tests to improve maintainability and ensure all fields are validated across different scenarios.
- 💡 Add assertions for `IsGlobal()` behavior to confirm that it correctly identifies global policies when language is empty.
- 💡 Add tests for edge cases such as empty `Rules` or `PathPatterns` slices to prevent runtime panics.
- 💡 Include validation checks for fields like `Threshold` or `Metric` to ensure they meet expected constraints (e.g., positive numbers).
- 💡 Ensure that `Violation` struct fields like `Description` are validated or formatted according to expected standards.
- ⚠️ Missing validation assertions for struct fields that could be invalid or malformed (e.g., empty rule IDs, invalid thresholds).
- ⚠️ Lack of test coverage for edge cases such as nil or empty slices in `Rules` or `PathPatterns`, which could lead to panics or incorrect behavior.
- 🔗 This unit is a test file and does not directly affect runtime behavior, but poor test coverage can mask bugs in dependent systems.
- 🔗 If `PolicyPack.IsGlobal()` or similar methods are not tested, incorrect logic may propagate to downstream systems that rely on this behavior.

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
- 🤖 Well-structured status enum with string conversion, but contains a critical race condition in initialization and lacks proper error handling for malformed inputs.
- 💡 Use atomic operations or sync.Once to ensure `stringToStatus` initialization is safe for concurrent access.
- 💡 Add input sanitization in `ParseStatus` to trim whitespace and normalize case before lookup.
- 💡 Validate `Score` and `Confidence` fields to ensure they fall within expected ranges (e.g., 0.0 to 1.0).
- 💡 Add validation logic in `CertificationRecord` struct methods or constructors to ensure data integrity (e.g., `ExpiresAt` after `CertifiedAt`).
- 💡 Consider adding unit tests for edge cases in `ParseStatus` (e.g., empty string, whitespace-only strings, case variations).
- ⚠️ Race condition in `init()` function where `stringToStatus` is populated non-atomically, leading to potential data races during concurrent access.
- ⚠️ Lack of input sanitization in `ParseStatus` — does not trim whitespace or normalize case, which can cause silent failures or misinterpretation of inputs.
- 🔗 The `Status` enum and its string conversion functions form part of the public API surface for certification systems, so any change here affects all consumers relying on string parsing.
- 🔗 The `CertificationRecord` struct tightly couples status, grade, and evidence fields without validation or consistency checks, increasing the risk of propagating invalid or inconsistent state across services.

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
- 🤖 Well-structured status enum with string conversion, but contains a critical race condition in initialization and lacks proper error handling for malformed inputs.
- 💡 Use atomic operations or sync.Once to ensure `stringToStatus` initialization is safe for concurrent access.
- 💡 Add input sanitization in `ParseStatus` to trim whitespace and normalize case before lookup.
- 💡 Validate `Score` and `Confidence` fields to ensure they fall within expected ranges (e.g., 0.0 to 1.0).
- 💡 Add validation logic in `CertificationRecord` struct methods or constructors to ensure data integrity (e.g., `ExpiresAt` after `CertifiedAt`).
- 💡 Consider adding unit tests for edge cases in `ParseStatus` (e.g., empty string, whitespace-only strings, case variations).
- ⚠️ Race condition in `init()` function where `stringToStatus` is populated non-atomically, leading to potential data races during concurrent access.
- ⚠️ Lack of input sanitization in `ParseStatus` — does not trim whitespace or normalize case, which can cause silent failures or misinterpretation of inputs.
- 🔗 The `Status` enum and its string conversion functions form part of the public API surface for certification systems, so any change here affects all consumers relying on string parsing.
- 🔗 The `CertificationRecord` struct tightly couples status, grade, and evidence fields without validation or consistency checks, increasing the risk of propagating invalid or inconsistent state across services.

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
- 🤖 Well-structured status enum with string conversion, but contains a critical race condition in initialization and lacks proper error handling for malformed inputs.
- 💡 Use atomic operations or sync.Once to ensure `stringToStatus` initialization is safe for concurrent access.
- 💡 Add input sanitization in `ParseStatus` to trim whitespace and normalize case before lookup.
- 💡 Validate `Score` and `Confidence` fields to ensure they fall within expected ranges (e.g., 0.0 to 1.0).
- 💡 Add validation logic in `CertificationRecord` struct methods or constructors to ensure data integrity (e.g., `ExpiresAt` after `CertifiedAt`).
- 💡 Consider adding unit tests for edge cases in `ParseStatus` (e.g., empty string, whitespace-only strings, case variations).
- ⚠️ Race condition in `init()` function where `stringToStatus` is populated non-atomically, leading to potential data races during concurrent access.
- ⚠️ Lack of input sanitization in `ParseStatus` — does not trim whitespace or normalize case, which can cause silent failures or misinterpretation of inputs.
- 🔗 The `Status` enum and its string conversion functions form part of the public API surface for certification systems, so any change here affects all consumers relying on string parsing.
- 🔗 The `CertificationRecord` struct tightly couples status, grade, and evidence fields without validation or consistency checks, increasing the risk of propagating invalid or inconsistent state across services.

</details>

<a id="internal-domain-record-go-status"></a>
<details>
<summary>Status — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 95.0% |
| security | 80.0% |
| testability | 90.0% |

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
- 🤖 Well-structured status enum with string conversion, but contains a critical race condition in initialization and lacks proper error handling for malformed inputs.
- 💡 Use atomic operations or sync.Once to ensure `stringToStatus` initialization is safe for concurrent access.
- 💡 Add input sanitization in `ParseStatus` to trim whitespace and normalize case before lookup.
- 💡 Validate `Score` and `Confidence` fields to ensure they fall within expected ranges (e.g., 0.0 to 1.0).
- 💡 Add validation logic in `CertificationRecord` struct methods or constructors to ensure data integrity (e.g., `ExpiresAt` after `CertifiedAt`).
- 💡 Consider adding unit tests for edge cases in `ParseStatus` (e.g., empty string, whitespace-only strings, case variations).
- ⚠️ Race condition in `init()` function where `stringToStatus` is populated non-atomically, leading to potential data races during concurrent access.
- ⚠️ Lack of input sanitization in `ParseStatus` — does not trim whitespace or normalize case, which can cause silent failures or misinterpretation of inputs.
- 🔗 The `Status` enum and its string conversion functions form part of the public API surface for certification systems, so any change here affects all consumers relying on string parsing.
- 🔗 The `CertificationRecord` struct tightly couples status, grade, and evidence fields without validation or consistency checks, increasing the risk of propagating invalid or inconsistent state across services.

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
- 🤖 Well-structured status enum with string conversion, but contains a critical race condition in initialization and lacks proper error handling for malformed inputs.
- 💡 Use atomic operations or sync.Once to ensure `stringToStatus` initialization is safe for concurrent access.
- 💡 Add input sanitization in `ParseStatus` to trim whitespace and normalize case before lookup.
- 💡 Validate `Score` and `Confidence` fields to ensure they fall within expected ranges (e.g., 0.0 to 1.0).
- 💡 Add validation logic in `CertificationRecord` struct methods or constructors to ensure data integrity (e.g., `ExpiresAt` after `CertifiedAt`).
- 💡 Consider adding unit tests for edge cases in `ParseStatus` (e.g., empty string, whitespace-only strings, case variations).
- ⚠️ Race condition in `init()` function where `stringToStatus` is populated non-atomically, leading to potential data races during concurrent access.
- ⚠️ Lack of input sanitization in `ParseStatus` — does not trim whitespace or normalize case, which can cause silent failures or misinterpretation of inputs.
- 🔗 The `Status` enum and its string conversion functions form part of the public API surface for certification systems, so any change here affects all consumers relying on string parsing.
- 🔗 The `CertificationRecord` struct tightly couples status, grade, and evidence fields without validation or consistency checks, increasing the risk of propagating invalid or inconsistent state across services.

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
- 🤖 Well-structured tests for Status enum functionality with good coverage but missing edge case handling and test robustness.
- 💡 Add a test case to verify that Status values are correctly mapped to their string representations and that the enum order matches expectations
- 💡 Add a test case for invalid Status values (e.g., domain.Status(999)) to ensure proper error handling in String() and IsPassing() methods
- 💡 Improve error messages in TestParseStatus to include the actual parsed value and input string for better debugging
- 💡 Add test coverage for case-sensitive parsing edge cases (e.g., "Certified" vs "certified") to ensure robustness
- ⚠️ Missing validation of enum value consistency and ordering, leading to brittle tests if Status constants change
- ⚠️ Potential silent failure or panic with out-of-range enum values in String() and IsPassing() methods
- 🔗 These tests directly validate the Status enum's behavior, which is used across the domain layer and affects downstream logic in certification workflows
- 🔗 The test suite has low coupling to other modules but high importance due to its role in validating critical domain logic

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured domain types with minor correctness and concurrency issues.
- 💡 Use `sync.Once` or similar synchronization mechanism in `init()` to ensure thread-safe population of `stringToUnitType`.
- 💡 Add input validation in `ParseUnitID` to ensure that the language field is non-empty and conforms to a valid format (e.g., alphanumeric characters only).
- 💡 Consider adding a public constructor or helper function for `UnitID` to facilitate unit testing with specific test cases.
- 💡 Add validation in `NewUnitID` or `ParseUnitID` to reject invalid paths (e.g., those containing null bytes or other control characters).
- ⚠️ Race condition in `init()` function where `stringToUnitType` is populated, potentially causing data races if accessed concurrently before initialization completes.
- ⚠️ Lack of validation in `ParseUnitID` for language component (e.g., empty or invalid characters) can lead to malformed identifiers being accepted.
- 🔗 The `UnitID` type is used as a core identifier across the system, so any parsing or validation issues will propagate to all consumers.
- 🔗 The global `stringToUnitType` map introduces tight coupling between this package and any code that depends on it, making future extensibility harder.

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
- 🤖 Well-structured unit tests with good coverage but missing edge case handling and potential string comparison issues.
- 💡 Add tests for case-insensitive parsing in ParseUnitType to ensure robustness against different input formats
- 💡 Add validation for malformed or invalid inputs in ParseUnitID to ensure proper error handling and prevent silent failures
- ⚠️ Edge case handling missing in ParseUnitType and ParseUnitID for malformed inputs
- ⚠️ Lack of validation for case sensitivity in unit type parsing
- 🔗 These tests validate core domain logic, so any failure in parsing or string conversion could propagate to downstream systems that depend on UnitID
- 🔗 If ParseUnitType or ParseUnitID has unhandled edge cases, it could lead to incorrect unit identification in code analysis tools

</details>

### `internal/engine/` (10 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`CertifyUnit`](reports/internal-engine-pipeline-go-certifyunit.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`pipeline_test.go`](reports/internal-engine-pipeline-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Score`](reports/internal-engine-scorer-go-score.md) | function | B | 86.1% | certified | 2026-04-24 |
| [`StatusFromScore`](reports/internal-engine-scorer-go-statusfromscore.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`extractSummaryFloat`](reports/internal-engine-scorer-go-extractsummaryfloat.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`extractSummaryInt`](reports/internal-engine-scorer-go-extractsummaryint.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`scoreFromGitHistory`](reports/internal-engine-scorer-go-scorefromgithistory.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`scoreFromMetrics`](reports/internal-engine-scorer-go-scorefrommetrics.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`severityPenalty`](reports/internal-engine-scorer-go-severitypenalty.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`scorer_test.go`](reports/internal-engine-scorer-test-go.md) | file | B+ | 87.2% | certified | 2026-04-23 |

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
- 🤖 The CertifyUnit function is well-structured but has several maintainability and correctness issues including hardcoded confidence, lack of error handling, and potential source attribution logic flaws.
- 💡 Replace hardcoded confidence value with a dynamic calculation based on evidence quality or policy evaluation strength
- 💡 Refactor source attribution logic to use a more robust method than string prefix matching (e.g., enum-based evidence kind or structured source format)
- 💡 Validate inputs like unit.ID, rules, and ev before processing to prevent nil pointer dereferences
- 💡 Add error handling for sub-functions like policy.Evaluate and Score to allow failure propagation and logging
- 💡 Use a constant or configuration for Version instead of hardcoding 1
- ⚠️ Hardcoded confidence value (1.0) that does not reflect actual evidence quality or policy evaluation strength
- ⚠️ Brittle source attribution logic relying on string prefix matching that may break with future evidence formats
- 🔗 The function tightly couples to internal domain types and assumes specific evidence formats, increasing system fragility
- 🔗 Lack of error propagation makes failure diagnosis difficult and can mask upstream issues in evidence collection or policy evaluation

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
- 🤖 The test suite for CertifyUnit has basic coverage but lacks robustness in edge case handling and validation of policy rule behavior.
- 💡 Use a mock or fixed time value in tests to ensure deterministic behavior and avoid flaky test runs
- 💡 Add assertions for specific score calculation logic, including edge cases like zero vs. non-zero lint errors or missing evidence
- 💡 Add a test case for invalid or malformed evidence to ensure proper error handling and status reporting
- 💡 Parameterize the test cases with different combinations of evidence, rules, and configurations to increase coverage
- ⚠️ Flaky tests due to reliance on time.Now() without mocking or deterministic setup
- ⚠️ Lack of validation for edge cases such as missing evidence, invalid rule formats, or unexpected metric values
- 🔗 These tests do not directly impact system components but provide minimal confidence in the correctness of CertifyUnit under various conditions
- 🔗 The tests are tightly coupled to implementation details like hardcoded strings and specific scoring thresholds, increasing maintenance burden

</details>

<a id="internal-engine-scorer-go-score"></a>
<details>
<summary>Score — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 80.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 95.0% |
| security | 80.0% |
| testability | 90.0% |

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
- 🤖 The scoring logic is functionally sound but has several correctness and robustness issues related to parsing, edge case handling, and inconsistent use of min/max functions.
- 💡 Replace `max(scores[domain.DimMaintainability], 0.95)` with direct assignment `scores[domain.DimMaintainability] = 0.95` to avoid overwriting higher scores
- 💡 Fix `extractSummaryInt` to properly parse numeric values by using regex or a more robust scanning method that handles all edge cases including numbers followed by non-numeric characters
- 💡 Add nil checks for `e.Summary` in all parsing functions and return early or default to safe values
- 💡 Use consistent logic for applying scores from evidence (e.g., always use max or always assign) to prevent overwriting
- 💡 Add unit tests for `extractSummaryInt` and `extractSummaryFloat` with various malformed inputs to ensure robustness
- ⚠️ Potential panic from `strings.Index(summary, keyword)` when summary is nil or malformed
- ⚠️ Incorrect score calculation due to improper use of max() in metrics and git history scoring
- 🔗 The `Score` function is a core part of certification logic and directly influences status determination, so any incorrect score can cause certification failures or false positives
- 🔗 The parsing functions (`extractSummaryInt`, `extractSummaryFloat`) are tightly coupled to evidence format and can break if input changes, affecting downstream systems that depend on consistent parsing

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
- 🤖 The scoring logic is functionally sound but has several correctness and robustness issues related to parsing, edge case handling, and inconsistent use of min/max functions.
- 💡 Replace `max(scores[domain.DimMaintainability], 0.95)` with direct assignment `scores[domain.DimMaintainability] = 0.95` to avoid overwriting higher scores
- 💡 Fix `extractSummaryInt` to properly parse numeric values by using regex or a more robust scanning method that handles all edge cases including numbers followed by non-numeric characters
- 💡 Add nil checks for `e.Summary` in all parsing functions and return early or default to safe values
- 💡 Use consistent logic for applying scores from evidence (e.g., always use max or always assign) to prevent overwriting
- 💡 Add unit tests for `extractSummaryInt` and `extractSummaryFloat` with various malformed inputs to ensure robustness
- ⚠️ Potential panic from `strings.Index(summary, keyword)` when summary is nil or malformed
- ⚠️ Incorrect score calculation due to improper use of max() in metrics and git history scoring
- 🔗 The `Score` function is a core part of certification logic and directly influences status determination, so any incorrect score can cause certification failures or false positives
- 🔗 The parsing functions (`extractSummaryInt`, `extractSummaryFloat`) are tightly coupled to evidence format and can break if input changes, affecting downstream systems that depend on consistent parsing

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
- 🤖 The scoring logic is functionally sound but has several correctness and robustness issues related to parsing, edge case handling, and inconsistent use of min/max functions.
- 💡 Replace `max(scores[domain.DimMaintainability], 0.95)` with direct assignment `scores[domain.DimMaintainability] = 0.95` to avoid overwriting higher scores
- 💡 Fix `extractSummaryInt` to properly parse numeric values by using regex or a more robust scanning method that handles all edge cases including numbers followed by non-numeric characters
- 💡 Add nil checks for `e.Summary` in all parsing functions and return early or default to safe values
- 💡 Use consistent logic for applying scores from evidence (e.g., always use max or always assign) to prevent overwriting
- 💡 Add unit tests for `extractSummaryInt` and `extractSummaryFloat` with various malformed inputs to ensure robustness
- ⚠️ Potential panic from `strings.Index(summary, keyword)` when summary is nil or malformed
- ⚠️ Incorrect score calculation due to improper use of max() in metrics and git history scoring
- 🔗 The `Score` function is a core part of certification logic and directly influences status determination, so any incorrect score can cause certification failures or false positives
- 🔗 The parsing functions (`extractSummaryInt`, `extractSummaryFloat`) are tightly coupled to evidence format and can break if input changes, affecting downstream systems that depend on consistent parsing

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
- 🤖 The scoring logic is functionally sound but has several correctness and robustness issues related to parsing, edge case handling, and inconsistent use of min/max functions.
- 💡 Replace `max(scores[domain.DimMaintainability], 0.95)` with direct assignment `scores[domain.DimMaintainability] = 0.95` to avoid overwriting higher scores
- 💡 Fix `extractSummaryInt` to properly parse numeric values by using regex or a more robust scanning method that handles all edge cases including numbers followed by non-numeric characters
- 💡 Add nil checks for `e.Summary` in all parsing functions and return early or default to safe values
- 💡 Use consistent logic for applying scores from evidence (e.g., always use max or always assign) to prevent overwriting
- 💡 Add unit tests for `extractSummaryInt` and `extractSummaryFloat` with various malformed inputs to ensure robustness
- ⚠️ Potential panic from `strings.Index(summary, keyword)` when summary is nil or malformed
- ⚠️ Incorrect score calculation due to improper use of max() in metrics and git history scoring
- 🔗 The `Score` function is a core part of certification logic and directly influences status determination, so any incorrect score can cause certification failures or false positives
- 🔗 The parsing functions (`extractSummaryInt`, `extractSummaryFloat`) are tightly coupled to evidence format and can break if input changes, affecting downstream systems that depend on consistent parsing

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
- 🤖 The scoring logic is functionally sound but has several correctness and robustness issues related to parsing, edge case handling, and inconsistent use of min/max functions.
- 💡 Replace `max(scores[domain.DimMaintainability], 0.95)` with direct assignment `scores[domain.DimMaintainability] = 0.95` to avoid overwriting higher scores
- 💡 Fix `extractSummaryInt` to properly parse numeric values by using regex or a more robust scanning method that handles all edge cases including numbers followed by non-numeric characters
- 💡 Add nil checks for `e.Summary` in all parsing functions and return early or default to safe values
- 💡 Use consistent logic for applying scores from evidence (e.g., always use max or always assign) to prevent overwriting
- 💡 Add unit tests for `extractSummaryInt` and `extractSummaryFloat` with various malformed inputs to ensure robustness
- ⚠️ Potential panic from `strings.Index(summary, keyword)` when summary is nil or malformed
- ⚠️ Incorrect score calculation due to improper use of max() in metrics and git history scoring
- 🔗 The `Score` function is a core part of certification logic and directly influences status determination, so any incorrect score can cause certification failures or false positives
- 🔗 The parsing functions (`extractSummaryInt`, `extractSummaryFloat`) are tightly coupled to evidence format and can break if input changes, affecting downstream systems that depend on consistent parsing

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
- 🤖 The scoring logic is functionally sound but has several correctness and robustness issues related to parsing, edge case handling, and inconsistent use of min/max functions.
- 💡 Replace `max(scores[domain.DimMaintainability], 0.95)` with direct assignment `scores[domain.DimMaintainability] = 0.95` to avoid overwriting higher scores
- 💡 Fix `extractSummaryInt` to properly parse numeric values by using regex or a more robust scanning method that handles all edge cases including numbers followed by non-numeric characters
- 💡 Add nil checks for `e.Summary` in all parsing functions and return early or default to safe values
- 💡 Use consistent logic for applying scores from evidence (e.g., always use max or always assign) to prevent overwriting
- 💡 Add unit tests for `extractSummaryInt` and `extractSummaryFloat` with various malformed inputs to ensure robustness
- ⚠️ Potential panic from `strings.Index(summary, keyword)` when summary is nil or malformed
- ⚠️ Incorrect score calculation due to improper use of max() in metrics and git history scoring
- 🔗 The `Score` function is a core part of certification logic and directly influences status determination, so any incorrect score can cause certification failures or false positives
- 🔗 The parsing functions (`extractSummaryInt`, `extractSummaryFloat`) are tightly coupled to evidence format and can break if input changes, affecting downstream systems that depend on consistent parsing

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
- 🤖 The test file has good coverage of scoring logic but lacks proper error handling, has brittle assertions, and does not validate edge cases or invalid inputs.
- 💡 Replace hardcoded score thresholds (e.g., 0.7, 0.8, 0.9) with configurable or documented tolerance ranges to make tests more robust
- 💡 Add test cases for nil evidence, empty evidence slices, and invalid evidence types to ensure robustness of engine.Score
- 💡 Validate that all relevant dimensions are properly scored in tests like TestScorer_GitHistoryBoostsScores
- 💡 Add assertions to verify that specific dimensions (e.g., correctness) are penalized in violation scenarios
- ⚠️ Hardcoded score thresholds in tests can cause false negatives when scoring logic changes
- ⚠️ Missing validation of invalid evidence types or malformed inputs could lead to runtime panics
- 🔗 These tests do not validate the full API surface of engine.Score, potentially missing critical edge cases
- 🔗 The tests assume specific scoring behaviors that may not hold in production due to lack of input validation

</details>

### `internal/evidence/` (39 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`Collector`](reports/internal-evidence-collector-go-collector.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ComputeGoComplexity`](reports/internal-evidence-complexity-go-computegocomplexity.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ComputeSymbolMetrics`](reports/internal-evidence-complexity-go-computesymbolmetrics.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`funcName`](reports/internal-evidence-complexity-go-funcname.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`complexity_test.go`](reports/internal-evidence-complexity-test-go.md) | file | B+ | 87.2% | certified | 2026-04-23 |
| [`CollectAll`](reports/internal-evidence-executor-go-collectall.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`HasGoMod`](reports/internal-evidence-executor-go-hasgomod.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`HasPackageJSON`](reports/internal-evidence-executor-go-haspackagejson.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`NewToolExecutor`](reports/internal-evidence-executor-go-newtoolexecutor.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ToolExecutor`](reports/internal-evidence-executor-go-toolexecutor.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`runGitStats`](reports/internal-evidence-executor-go-rungitstats.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`runGoTest`](reports/internal-evidence-executor-go-rungotest.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`runGoVet`](reports/internal-evidence-executor-go-rungovet.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`runGolangciLint`](reports/internal-evidence-executor-go-rungolangcilint.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ChurnRate`](reports/internal-evidence-git-go-churnrate.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`GitStats`](reports/internal-evidence-git-go-gitstats.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseGitLog`](reports/internal-evidence-git-go-parsegitlog.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ToEvidence`](reports/internal-evidence-git-go-toevidence.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`git_test.go`](reports/internal-evidence-git-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`LintFinding`](reports/internal-evidence-lint-go-lintfinding.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`LintResult`](reports/internal-evidence-lint-go-lintresult.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`TestResult`](reports/internal-evidence-lint-go-testresult.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ToEvidence`](reports/internal-evidence-lint-go-toevidence.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`lint_test.go`](reports/internal-evidence-lint-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`CodeMetrics`](reports/internal-evidence-metrics-go-codemetrics.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ComputeMetrics`](reports/internal-evidence-metrics-go-computemetrics.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`ToEvidence`](reports/internal-evidence-metrics-go-toevidence.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`containsTodo`](reports/internal-evidence-metrics-go-containstodo.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`metrics_test.go`](reports/internal-evidence-metrics-test-go.md) | file | B+ | 87.2% | certified | 2026-04-23 |
| [`ParseCoverProfile`](reports/internal-evidence-runner-go-parsecoverprofile.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ParseGitLogWithAge`](reports/internal-evidence-runner-go-parsegitlogwithage.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ParseGoTestJSON`](reports/internal-evidence-runner-go-parsegotestjson.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ParseGoVet`](reports/internal-evidence-runner-go-parsegovet.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ParseGolangciLintJSON`](reports/internal-evidence-runner-go-parsegolangcilintjson.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`goTestEvent`](reports/internal-evidence-runner-go-gotestevent.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`golangciLintIssue`](reports/internal-evidence-runner-go-golangcilintissue.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`golangciLintOutput`](reports/internal-evidence-runner-go-golangcilintoutput.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`simpleAtoi`](reports/internal-evidence-runner-go-simpleatoi.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`runner_test.go`](reports/internal-evidence-runner-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 Minimal interface with no implementation, appears to be a placeholder or incomplete design
- 💡 Implement concrete collector types with proper error handling and validation
- 💡 Add input sanitization for root path parameter to prevent directory traversal attacks
- 💡 Define clear contract documentation for the Collect method including expected error types and behavior
- 💡 Add unit tests for the interface methods with proper mock implementations
- ⚠️ Missing implementation leads to runtime panic if any method is called
- ⚠️ No input validation or sanitization for root path parameter
- 🔗 This interface creates tight coupling to external domain types without clear contract definition
- 🔗 The lack of concrete implementation makes this module non-functional and breaks the intended evidence collection pipeline

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
- 🤖 The code computes cyclomatic complexity and symbol metrics from Go source, but has critical logic flaws in symbol matching and function extraction.
- 💡 Fix `funcName` to properly handle all receiver types including nested identifiers and qualified names
- 💡 Replace line-based source extraction in `ComputeSymbolMetrics` with a more robust AST node range extraction
- 💡 Ensure that symbol matching logic in `ComputeSymbolMetrics` correctly handles both method and function name matches
- 💡 Add unit tests to verify that complex receiver types are handled correctly in `funcName`
- ⚠️ Incorrect function name generation in `funcName` for complex receiver types leads to incorrect complexity tracking
- ⚠️ Line-based source extraction in `ComputeSymbolMetrics` is fragile and can produce incorrect or incomplete function code
- 🔗 The complexity tracking is broken for methods with complex receivers, causing incorrect metrics in downstream systems
- 🔗 The symbol extraction logic introduces a coupling between AST parsing and string manipulation that makes the code brittle and hard to maintain

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
- 🤖 The code computes cyclomatic complexity and symbol metrics from Go source, but has critical logic flaws in symbol matching and function extraction.
- 💡 Fix `funcName` to properly handle all receiver types including nested identifiers and qualified names
- 💡 Replace line-based source extraction in `ComputeSymbolMetrics` with a more robust AST node range extraction
- 💡 Ensure that symbol matching logic in `ComputeSymbolMetrics` correctly handles both method and function name matches
- 💡 Add unit tests to verify that complex receiver types are handled correctly in `funcName`
- ⚠️ Incorrect function name generation in `funcName` for complex receiver types leads to incorrect complexity tracking
- ⚠️ Line-based source extraction in `ComputeSymbolMetrics` is fragile and can produce incorrect or incomplete function code
- 🔗 The complexity tracking is broken for methods with complex receivers, causing incorrect metrics in downstream systems
- 🔗 The symbol extraction logic introduces a coupling between AST parsing and string manipulation that makes the code brittle and hard to maintain

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
- 🤖 The code computes cyclomatic complexity and symbol metrics from Go source, but has critical logic flaws in symbol matching and function extraction.
- 💡 Fix `funcName` to properly handle all receiver types including nested identifiers and qualified names
- 💡 Replace line-based source extraction in `ComputeSymbolMetrics` with a more robust AST node range extraction
- 💡 Ensure that symbol matching logic in `ComputeSymbolMetrics` correctly handles both method and function name matches
- 💡 Add unit tests to verify that complex receiver types are handled correctly in `funcName`
- ⚠️ Incorrect function name generation in `funcName` for complex receiver types leads to incorrect complexity tracking
- ⚠️ Line-based source extraction in `ComputeSymbolMetrics` is fragile and can produce incorrect or incomplete function code
- 🔗 The complexity tracking is broken for methods with complex receivers, causing incorrect metrics in downstream systems
- 🔗 The symbol extraction logic introduces a coupling between AST parsing and string manipulation that makes the code brittle and hard to maintain

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
- 🤖 The test suite for Go complexity and symbol metrics is functionally correct but lacks robustness in edge case handling and has inconsistent test expectations.
- 💡 Clarify and document the complexity calculation rules (e.g., how logical operators, switch cases, and nested control structures contribute to complexity) so that test expectations are consistent and maintainable
- 💡 Add tests for malformed or invalid Go code to ensure that `ComputeGoComplexity` and `ComputeSymbolMetrics` handle errors gracefully without panicking or returning incorrect values
- 💡 Refactor test assertions to be more precise and aligned with documented behavior, e.g., instead of asserting `result["complex"] < 5`, assert that the complexity is exactly 8 based on a well-defined rule
- 💡 Improve test naming to better reflect what is being tested (e.g., `TestComputeGoComplexity_WithBranches_ExpectedComplexityIsEight`)
- ⚠️ Inconsistent complexity calculation logic in test expectations may mask bugs in the actual implementation
- ⚠️ Lack of input validation or error handling tests for malformed Go source code
- 🔗 These tests do not isolate or mock external dependencies, making them fragile to changes in the underlying implementation
- 🔗 The tests rely on string parsing and function name resolution logic that may not scale or be resilient to refactorings in the main codebase

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
- 🤖 The code has functional correctness but suffers from poor error handling, resource leaks, and missing input validation.
- 💡 Add error checking for `coverCmd.Run()` in `runGoTest` and ensure the cover file is properly cleaned up
- 💡 Remove the underscore in `output, _ := cmd.CombinedOutput()` on line 105 to handle errors from golangci-lint command
- 💡 Improve git log parsing in `runGitStats` to reliably extract earliest commit date instead of assuming last line is earliest
- 💡 Add context-based timeouts to all `exec.Command` invocations to prevent hanging processes
- ⚠️ Resource leak due to missing error handling in `runGoTest` when removing temp cover file
- ⚠️ Silent failure in `runGolangciLint` due to ignored error from command execution
- 🔗 The tool executor tightly couples to external tools (go, golangci-lint, git) and assumes their presence and correct behavior
- 🔗 Lack of timeouts or concurrency control in `CollectAll` can cause performance bottlenecks and system instability under high load

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
- 🤖 The code has functional correctness but suffers from poor error handling, resource leaks, and missing input validation.
- 💡 Add error checking for `coverCmd.Run()` in `runGoTest` and ensure the cover file is properly cleaned up
- 💡 Remove the underscore in `output, _ := cmd.CombinedOutput()` on line 105 to handle errors from golangci-lint command
- 💡 Improve git log parsing in `runGitStats` to reliably extract earliest commit date instead of assuming last line is earliest
- 💡 Add context-based timeouts to all `exec.Command` invocations to prevent hanging processes
- ⚠️ Resource leak due to missing error handling in `runGoTest` when removing temp cover file
- ⚠️ Silent failure in `runGolangciLint` due to ignored error from command execution
- 🔗 The tool executor tightly couples to external tools (go, golangci-lint, git) and assumes their presence and correct behavior
- 🔗 Lack of timeouts or concurrency control in `CollectAll` can cause performance bottlenecks and system instability under high load

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
- 🤖 The code has functional correctness but suffers from poor error handling, resource leaks, and missing input validation.
- 💡 Add error checking for `coverCmd.Run()` in `runGoTest` and ensure the cover file is properly cleaned up
- 💡 Remove the underscore in `output, _ := cmd.CombinedOutput()` on line 105 to handle errors from golangci-lint command
- 💡 Improve git log parsing in `runGitStats` to reliably extract earliest commit date instead of assuming last line is earliest
- 💡 Add context-based timeouts to all `exec.Command` invocations to prevent hanging processes
- ⚠️ Resource leak due to missing error handling in `runGoTest` when removing temp cover file
- ⚠️ Silent failure in `runGolangciLint` due to ignored error from command execution
- 🔗 The tool executor tightly couples to external tools (go, golangci-lint, git) and assumes their presence and correct behavior
- 🔗 Lack of timeouts or concurrency control in `CollectAll` can cause performance bottlenecks and system instability under high load

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
- 🤖 The code has functional correctness but suffers from poor error handling, resource leaks, and missing input validation.
- 💡 Add error checking for `coverCmd.Run()` in `runGoTest` and ensure the cover file is properly cleaned up
- 💡 Remove the underscore in `output, _ := cmd.CombinedOutput()` on line 105 to handle errors from golangci-lint command
- 💡 Improve git log parsing in `runGitStats` to reliably extract earliest commit date instead of assuming last line is earliest
- 💡 Add context-based timeouts to all `exec.Command` invocations to prevent hanging processes
- ⚠️ Resource leak due to missing error handling in `runGoTest` when removing temp cover file
- ⚠️ Silent failure in `runGolangciLint` due to ignored error from command execution
- 🔗 The tool executor tightly couples to external tools (go, golangci-lint, git) and assumes their presence and correct behavior
- 🔗 Lack of timeouts or concurrency control in `CollectAll` can cause performance bottlenecks and system instability under high load

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
- 🤖 The code has functional correctness but suffers from poor error handling, resource leaks, and missing input validation.
- 💡 Add error checking for `coverCmd.Run()` in `runGoTest` and ensure the cover file is properly cleaned up
- 💡 Remove the underscore in `output, _ := cmd.CombinedOutput()` on line 105 to handle errors from golangci-lint command
- 💡 Improve git log parsing in `runGitStats` to reliably extract earliest commit date instead of assuming last line is earliest
- 💡 Add context-based timeouts to all `exec.Command` invocations to prevent hanging processes
- ⚠️ Resource leak due to missing error handling in `runGoTest` when removing temp cover file
- ⚠️ Silent failure in `runGolangciLint` due to ignored error from command execution
- 🔗 The tool executor tightly couples to external tools (go, golangci-lint, git) and assumes their presence and correct behavior
- 🔗 Lack of timeouts or concurrency control in `CollectAll` can cause performance bottlenecks and system instability under high load

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
- 🤖 The code has functional correctness but suffers from poor error handling, resource leaks, and missing input validation.
- 💡 Add error checking for `coverCmd.Run()` in `runGoTest` and ensure the cover file is properly cleaned up
- 💡 Remove the underscore in `output, _ := cmd.CombinedOutput()` on line 105 to handle errors from golangci-lint command
- 💡 Improve git log parsing in `runGitStats` to reliably extract earliest commit date instead of assuming last line is earliest
- 💡 Add context-based timeouts to all `exec.Command` invocations to prevent hanging processes
- ⚠️ Resource leak due to missing error handling in `runGoTest` when removing temp cover file
- ⚠️ Silent failure in `runGolangciLint` due to ignored error from command execution
- 🔗 The tool executor tightly couples to external tools (go, golangci-lint, git) and assumes their presence and correct behavior
- 🔗 Lack of timeouts or concurrency control in `CollectAll` can cause performance bottlenecks and system instability under high load

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
- 🤖 The code has functional correctness but suffers from poor error handling, resource leaks, and missing input validation.
- 💡 Add error checking for `coverCmd.Run()` in `runGoTest` and ensure the cover file is properly cleaned up
- 💡 Remove the underscore in `output, _ := cmd.CombinedOutput()` on line 105 to handle errors from golangci-lint command
- 💡 Improve git log parsing in `runGitStats` to reliably extract earliest commit date instead of assuming last line is earliest
- 💡 Add context-based timeouts to all `exec.Command` invocations to prevent hanging processes
- ⚠️ Resource leak due to missing error handling in `runGoTest` when removing temp cover file
- ⚠️ Silent failure in `runGolangciLint` due to ignored error from command execution
- 🔗 The tool executor tightly couples to external tools (go, golangci-lint, git) and assumes their presence and correct behavior
- 🔗 Lack of timeouts or concurrency control in `CollectAll` can cause performance bottlenecks and system instability under high load

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
- 🤖 The code has functional correctness but suffers from poor error handling, resource leaks, and missing input validation.
- 💡 Add error checking for `coverCmd.Run()` in `runGoTest` and ensure the cover file is properly cleaned up
- 💡 Remove the underscore in `output, _ := cmd.CombinedOutput()` on line 105 to handle errors from golangci-lint command
- 💡 Improve git log parsing in `runGitStats` to reliably extract earliest commit date instead of assuming last line is earliest
- 💡 Add context-based timeouts to all `exec.Command` invocations to prevent hanging processes
- ⚠️ Resource leak due to missing error handling in `runGoTest` when removing temp cover file
- ⚠️ Silent failure in `runGolangciLint` due to ignored error from command execution
- 🔗 The tool executor tightly couples to external tools (go, golangci-lint, git) and assumes their presence and correct behavior
- 🔗 Lack of timeouts or concurrency control in `CollectAll` can cause performance bottlenecks and system instability under high load

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
- 🤖 The code is functionally correct but has several critical issues including unhandled errors, resource leaks, and missing error propagation that could lead to silent failures.
- 💡 Add error handling for `coverCmd.Run()` in runGoTest to ensure coverage collection errors are not silently ignored
- 💡 Handle the error return from `cmd.CombinedOutput()` in runGolangciLint to ensure linting errors are reported
- 💡 Use a more robust temporary file generation mechanism in runGoTest to avoid race conditions and ensure cleanup
- 💡 Add timeout to all external command executions using exec.CommandContext to prevent hanging processes
- 💡 Implement proper error aggregation in CollectAll to surface partial failures instead of silently dropping evidence
- ⚠️ Uncaught error in golangci-lint execution (line 72) leads to silent failures
- ⚠️ Resource leak in runGoTest due to deferred file removal without checking for errors
- 🔗 This unit tightly couples to external tool availability and assumes specific command outputs without robust error handling
- 🔗 The lack of proper error propagation in CollectAll can mask tool failures and cause downstream systems to receive incomplete evidence

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
- 🤖 The code is functionally correct but has a critical logic flaw in age calculation and lacks proper error handling for malformed git log data.
- 💡 Modify ParseGitLog to parse the third field (date) and compute actual age in days between earliest and latest commits
- 💡 Add validation to ensure the date field is parseable before computing age, or return an error if malformed
- 💡 Add unit tests for ParseGitLog with various date formats and edge cases to prevent silent failures
- ⚠️ Incorrect age calculation in GitStats due to missing date parsing logic
- ⚠️ Potential panic or incorrect data if git log output contains invalid date formats
- 🔗 The age field in GitStats affects downstream systems that rely on accurate temporal metrics for code health or risk assessment
- 🔗 The ToEvidence method's confidence value assumes deterministic git data, but incorrect age leads to misleading evidence

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
- 🤖 The code is functionally correct but has a critical logic flaw in age calculation and lacks proper error handling for malformed git log data.
- 💡 Modify ParseGitLog to parse the third field (date) and compute actual age in days between earliest and latest commits
- 💡 Add validation to ensure the date field is parseable before computing age, or return an error if malformed
- 💡 Add unit tests for ParseGitLog with various date formats and edge cases to prevent silent failures
- ⚠️ Incorrect age calculation in GitStats due to missing date parsing logic
- ⚠️ Potential panic or incorrect data if git log output contains invalid date formats
- 🔗 The age field in GitStats affects downstream systems that rely on accurate temporal metrics for code health or risk assessment
- 🔗 The ToEvidence method's confidence value assumes deterministic git data, but incorrect age leads to misleading evidence

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
- 🤖 The code is functionally correct but has a critical logic flaw in age calculation and lacks proper error handling for malformed git log data.
- 💡 Modify ParseGitLog to parse the third field (date) and compute actual age in days between earliest and latest commits
- 💡 Add validation to ensure the date field is parseable before computing age, or return an error if malformed
- 💡 Add unit tests for ParseGitLog with various date formats and edge cases to prevent silent failures
- ⚠️ Incorrect age calculation in GitStats due to missing date parsing logic
- ⚠️ Potential panic or incorrect data if git log output contains invalid date formats
- 🔗 The age field in GitStats affects downstream systems that rely on accurate temporal metrics for code health or risk assessment
- 🔗 The ToEvidence method's confidence value assumes deterministic git data, but incorrect age leads to misleading evidence

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
- 🤖 The code is functionally correct but has a critical logic flaw in age calculation and lacks proper error handling for malformed git log data.
- 💡 Modify ParseGitLog to parse the third field (date) and compute actual age in days between earliest and latest commits
- 💡 Add validation to ensure the date field is parseable before computing age, or return an error if malformed
- 💡 Add unit tests for ParseGitLog with various date formats and edge cases to prevent silent failures
- ⚠️ Incorrect age calculation in GitStats due to missing date parsing logic
- ⚠️ Potential panic or incorrect data if git log output contains invalid date formats
- 🔗 The age field in GitStats affects downstream systems that rely on accurate temporal metrics for code health or risk assessment
- 🔗 The ToEvidence method's confidence value assumes deterministic git data, but incorrect age leads to misleading evidence

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
- 🤖 The code is functionally correct but lacks robustness in parsing and has weak test coverage for edge cases.
- 💡 Add tests for malformed or inconsistent git log input to ensure robust parsing behavior.
- 💡 Use a more robust floating-point comparison (e.g., using `math.Abs` with epsilon) in churn rate test instead of direct inequality.
- 💡 Verify that all fields from GitStats (e.g., CommitCount, AuthorCount) are correctly included in the ToEvidence output.
- 💡 Add a test case for `ParseGitLog` with multiple whitespace or inconsistent spacing between fields to ensure robustness.
- ⚠️ Improper parsing of git log input can lead to incorrect statistics or panics if fields are missing or malformed.
- ⚠️ Floating-point comparison in churn rate test is brittle and may cause flaky test failures.
- 🔗 The parsing logic directly affects downstream evidence processing, so incorrect parsing could propagate errors into certification logic.
- 🔗 The test suite does not cover edge cases like malformed log lines or inconsistent field counts, increasing risk of undetected bugs in production.

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
- 🤖 The code is functionally correct but has weak type safety and lacks proper error handling in evidence conversion.
- 💡 Inject a time provider or use dependency injection for time.Now() to enable testability
- 💡 Add validation logic to sanitize fields like Severity, Message, and Tool before inclusion in evidence
- 💡 Consider adding a check to ensure that the Details field in domain.Evidence is properly typed when embedding LintResult or TestResult
- 💡 Add unit tests to cover edge cases like zero counts, invalid coverage values, and nil findings
- ⚠️ Hardcoded time.Now() in ToEvidence() prevents testability and makes behavior non-deterministic
- ⚠️ No input validation on fields like Severity, Message, or Tool can lead to malformed evidence
- 🔗 The ToEvidence methods tightly couple this module to the domain.Evidence type and its assumptions about struct embedding
- 🔗 The use of time.Now() introduces non-deterministic behavior that can affect system state consistency and testability

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
- 🤖 The code is functionally correct but has weak type safety and lacks proper error handling in evidence conversion.
- 💡 Inject a time provider or use dependency injection for time.Now() to enable testability
- 💡 Add validation logic to sanitize fields like Severity, Message, and Tool before inclusion in evidence
- 💡 Consider adding a check to ensure that the Details field in domain.Evidence is properly typed when embedding LintResult or TestResult
- 💡 Add unit tests to cover edge cases like zero counts, invalid coverage values, and nil findings
- ⚠️ Hardcoded time.Now() in ToEvidence() prevents testability and makes behavior non-deterministic
- ⚠️ No input validation on fields like Severity, Message, or Tool can lead to malformed evidence
- 🔗 The ToEvidence methods tightly couple this module to the domain.Evidence type and its assumptions about struct embedding
- 🔗 The use of time.Now() introduces non-deterministic behavior that can affect system state consistency and testability

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
- 🤖 The code is functionally correct but has weak type safety and lacks proper error handling in evidence conversion.
- 💡 Inject a time provider or use dependency injection for time.Now() to enable testability
- 💡 Add validation logic to sanitize fields like Severity, Message, and Tool before inclusion in evidence
- 💡 Consider adding a check to ensure that the Details field in domain.Evidence is properly typed when embedding LintResult or TestResult
- 💡 Add unit tests to cover edge cases like zero counts, invalid coverage values, and nil findings
- ⚠️ Hardcoded time.Now() in ToEvidence() prevents testability and makes behavior non-deterministic
- ⚠️ No input validation on fields like Severity, Message, or Tool can lead to malformed evidence
- 🔗 The ToEvidence methods tightly couple this module to the domain.Evidence type and its assumptions about struct embedding
- 🔗 The use of time.Now() introduces non-deterministic behavior that can affect system state consistency and testability

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
- 🤖 The code is functionally correct but has weak type safety and lacks proper error handling in evidence conversion.
- 💡 Inject a time provider or use dependency injection for time.Now() to enable testability
- 💡 Add validation logic to sanitize fields like Severity, Message, and Tool before inclusion in evidence
- 💡 Consider adding a check to ensure that the Details field in domain.Evidence is properly typed when embedding LintResult or TestResult
- 💡 Add unit tests to cover edge cases like zero counts, invalid coverage values, and nil findings
- ⚠️ Hardcoded time.Now() in ToEvidence() prevents testability and makes behavior non-deterministic
- ⚠️ No input validation on fields like Severity, Message, or Tool can lead to malformed evidence
- 🔗 The ToEvidence methods tightly couple this module to the domain.Evidence type and its assumptions about struct embedding
- 🔗 The use of time.Now() introduces non-deterministic behavior that can affect system state consistency and testability

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
- 🤖 The test file has good coverage of core logic but lacks validation for edge cases and error conditions in ToEvidence() conversion.
- 💡 Add a test case for TestLintResult_WithErrors to verify that the returned evidence contains exactly 4 findings with correct file, line, message and severity fields
- 💡 Add a test case for TestLintResult_WarningsOnly to assert that the returned evidence has Passed=true and Kind=domain.EvidenceKindLint
- 💡 Add a test case for TestLintResult with ErrorCount=0, WarnCount=0 and empty Findings to ensure ToEvidence() correctly returns a clean pass
- 💡 Add a test case for TestLintResult with ErrorCount=0, WarnCount=1 and severity='error' to ensure that strict lint tools are handled correctly
- 💡 Add a test case for TestTestResult with Coverage=0.0 to ensure that zero coverage is handled properly in the evidence conversion
- ⚠️ Missing validation for ToEvidence() behavior with empty or malformed inputs
- ⚠️ No verification that lint findings are correctly mapped to evidence structure
- 🔗 This unit tests conversion logic between domain-specific types and evidence types, which are used in certification workflows
- 🔗 If ToEvidence() doesn't correctly handle edge cases, it could lead to incorrect certification decisions in downstream systems

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
- 🤖 The code correctly computes basic metrics but has a critical logic flaw in TODO detection and lacks proper handling of multi-line comments.
- 💡 Replace `strings.Contains(upper, "TODO")` with regex pattern matching to ensure word boundaries: `regexp.MustCompile(`(?i)\bTODO\b|\bFIXME\b`).FindString(line) != ""`
- 💡 Add support for multi-line comments by checking for `/*` and `*/` in the line or tracking state across multiple lines
- 💡 Update the Complexity field to reflect actual complexity calculation logic (currently unused)
- 💡 Add handling for Windows line endings (`\r\n`) by normalizing them before splitting or using `bufio.Scanner`
- ⚠️ Incorrect TODO detection due to substring matching instead of word-boundary matching
- ⚠️ Missing support for multi-line comments (e.g., C-style /* ... */)
- 🔗 The function's incorrect TODO detection can lead to false positives in code quality checks
- 🔗 Inconsistent line ending handling may cause incorrect metrics on Windows systems

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
- 🤖 The code correctly computes basic metrics but has a critical logic flaw in TODO detection and lacks proper handling of multi-line comments.
- 💡 Replace `strings.Contains(upper, "TODO")` with regex pattern matching to ensure word boundaries: `regexp.MustCompile(`(?i)\bTODO\b|\bFIXME\b`).FindString(line) != ""`
- 💡 Add support for multi-line comments by checking for `/*` and `*/` in the line or tracking state across multiple lines
- 💡 Update the Complexity field to reflect actual complexity calculation logic (currently unused)
- 💡 Add handling for Windows line endings (`\r\n`) by normalizing them before splitting or using `bufio.Scanner`
- ⚠️ Incorrect TODO detection due to substring matching instead of word-boundary matching
- ⚠️ Missing support for multi-line comments (e.g., C-style /* ... */)
- 🔗 The function's incorrect TODO detection can lead to false positives in code quality checks
- 🔗 Inconsistent line ending handling may cause incorrect metrics on Windows systems

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
- 🤖 The code correctly computes basic metrics but has a critical logic flaw in TODO detection and lacks proper handling of multi-line comments.
- 💡 Replace `strings.Contains(upper, "TODO")` with regex pattern matching to ensure word boundaries: `regexp.MustCompile(`(?i)\bTODO\b|\bFIXME\b`).FindString(line) != ""`
- 💡 Add support for multi-line comments by checking for `/*` and `*/` in the line or tracking state across multiple lines
- 💡 Update the Complexity field to reflect actual complexity calculation logic (currently unused)
- 💡 Add handling for Windows line endings (`\r\n`) by normalizing them before splitting or using `bufio.Scanner`
- ⚠️ Incorrect TODO detection due to substring matching instead of word-boundary matching
- ⚠️ Missing support for multi-line comments (e.g., C-style /* ... */)
- 🔗 The function's incorrect TODO detection can lead to false positives in code quality checks
- 🔗 Inconsistent line ending handling may cause incorrect metrics on Windows systems

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
- 🤖 The code correctly computes basic metrics but has a critical logic flaw in TODO detection and lacks proper handling of multi-line comments.
- 💡 Replace `strings.Contains(upper, "TODO")` with regex pattern matching to ensure word boundaries: `regexp.MustCompile(`(?i)\bTODO\b|\bFIXME\b`).FindString(line) != ""`
- 💡 Add support for multi-line comments by checking for `/*` and `*/` in the line or tracking state across multiple lines
- 💡 Update the Complexity field to reflect actual complexity calculation logic (currently unused)
- 💡 Add handling for Windows line endings (`\r\n`) by normalizing them before splitting or using `bufio.Scanner`
- ⚠️ Incorrect TODO detection due to substring matching instead of word-boundary matching
- ⚠️ Missing support for multi-line comments (e.g., C-style /* ... */)
- 🔗 The function's incorrect TODO detection can lead to false positives in code quality checks
- 🔗 Inconsistent line ending handling may cause incorrect metrics on Windows systems

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
- 🤖 The test suite for code metrics is functionally correct but lacks robustness in edge case handling and has inconsistent assertions that could mask bugs.
- 💡 Change `if m.CommentLines < 3` to `if m.CommentLines != 3` in TestCodeMetrics_CommentLines to enforce exact matching
- 💡 Add a test case for files with trailing newlines or mixed line endings (e.g., \r\n) to ensure robustness of line counting
- 💡 Add a test for `TestCodeMetrics_ToEvidence` that verifies all fields in the returned evidence struct are correctly populated
- 💡 Add a test for edge case inputs like empty files with trailing newlines or whitespace-only lines to ensure robustness
- ⚠️ Inconsistent assertion in TestCodeMetrics_CommentLines allows incorrect comment counting logic to pass
- ⚠️ Fragile line count assumptions in TestCodeMetrics_LineCount may break with different input formatting
- 🔗 These tests do not adequately cover edge cases that could cause downstream systems to misinterpret code metrics
- 🔗 The ToEvidence conversion test is incomplete and does not ensure full fidelity of data mapping

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
- 🤖 The code is functionally correct but has several issues including missing error handling in JSON parsing, potential runtime panics from unsafe string operations, and lack of input validation that could lead to incorrect linting results.
- 💡 Add explicit nil checks before accessing nested JSON fields like issue.Pos.Filename in ParseGolangciLintJSON
- 💡 Replace simpleAtoi with a safer integer parsing function that handles overflow and invalid input gracefully
- 💡 Use strings.LastIndex instead of strings.Index in ParseGoVet to correctly extract file paths with colons
- 💡 Add comprehensive unit tests covering edge cases like malformed JSON, missing fields, and invalid date formats
- 💡 Implement proper error propagation instead of silently returning clean results when parsing fails
- ⚠️ Potential nil pointer dereference in ParseGolangciLintJSON when accessing issue.Pos.Filename without checking if issue.Pos exists
- ⚠️ Integer overflow in simpleAtoi function when parsing large numeric strings in ParseCoverProfile
- 🔗 The parsing functions directly affect linting and test result reporting, so any incorrect parsing can cause false positives/negatives in code quality metrics
- 🔗 Tight coupling between these parsers and specific tool output formats makes the system brittle to changes in Go tooling output format

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
- 🤖 The code is functionally correct but has several issues including missing error handling in JSON parsing, potential runtime panics from unsafe string operations, and lack of input validation that could lead to incorrect linting results.
- 💡 Add explicit nil checks before accessing nested JSON fields like issue.Pos.Filename in ParseGolangciLintJSON
- 💡 Replace simpleAtoi with a safer integer parsing function that handles overflow and invalid input gracefully
- 💡 Use strings.LastIndex instead of strings.Index in ParseGoVet to correctly extract file paths with colons
- 💡 Add comprehensive unit tests covering edge cases like malformed JSON, missing fields, and invalid date formats
- 💡 Implement proper error propagation instead of silently returning clean results when parsing fails
- ⚠️ Potential nil pointer dereference in ParseGolangciLintJSON when accessing issue.Pos.Filename without checking if issue.Pos exists
- ⚠️ Integer overflow in simpleAtoi function when parsing large numeric strings in ParseCoverProfile
- 🔗 The parsing functions directly affect linting and test result reporting, so any incorrect parsing can cause false positives/negatives in code quality metrics
- 🔗 Tight coupling between these parsers and specific tool output formats makes the system brittle to changes in Go tooling output format

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
- 🤖 The code is functionally correct but has several issues including missing error handling in JSON parsing, potential runtime panics from unsafe string operations, and lack of input validation that could lead to incorrect linting results.
- 💡 Add explicit nil checks before accessing nested JSON fields like issue.Pos.Filename in ParseGolangciLintJSON
- 💡 Replace simpleAtoi with a safer integer parsing function that handles overflow and invalid input gracefully
- 💡 Use strings.LastIndex instead of strings.Index in ParseGoVet to correctly extract file paths with colons
- 💡 Add comprehensive unit tests covering edge cases like malformed JSON, missing fields, and invalid date formats
- 💡 Implement proper error propagation instead of silently returning clean results when parsing fails
- ⚠️ Potential nil pointer dereference in ParseGolangciLintJSON when accessing issue.Pos.Filename without checking if issue.Pos exists
- ⚠️ Integer overflow in simpleAtoi function when parsing large numeric strings in ParseCoverProfile
- 🔗 The parsing functions directly affect linting and test result reporting, so any incorrect parsing can cause false positives/negatives in code quality metrics
- 🔗 Tight coupling between these parsers and specific tool output formats makes the system brittle to changes in Go tooling output format

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
- 🤖 The code is functionally correct but has several issues including missing error handling in JSON parsing, potential runtime panics from unsafe string operations, and lack of input validation that could lead to incorrect linting results.
- 💡 Add explicit nil checks before accessing nested JSON fields like issue.Pos.Filename in ParseGolangciLintJSON
- 💡 Replace simpleAtoi with a safer integer parsing function that handles overflow and invalid input gracefully
- 💡 Use strings.LastIndex instead of strings.Index in ParseGoVet to correctly extract file paths with colons
- 💡 Add comprehensive unit tests covering edge cases like malformed JSON, missing fields, and invalid date formats
- 💡 Implement proper error propagation instead of silently returning clean results when parsing fails
- ⚠️ Potential nil pointer dereference in ParseGolangciLintJSON when accessing issue.Pos.Filename without checking if issue.Pos exists
- ⚠️ Integer overflow in simpleAtoi function when parsing large numeric strings in ParseCoverProfile
- 🔗 The parsing functions directly affect linting and test result reporting, so any incorrect parsing can cause false positives/negatives in code quality metrics
- 🔗 Tight coupling between these parsers and specific tool output formats makes the system brittle to changes in Go tooling output format

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
- 🤖 The code is functionally correct but has several issues including missing error handling in JSON parsing, potential runtime panics from unsafe string operations, and lack of input validation that could lead to incorrect linting results.
- 💡 Add explicit nil checks before accessing nested JSON fields like issue.Pos.Filename in ParseGolangciLintJSON
- 💡 Replace simpleAtoi with a safer integer parsing function that handles overflow and invalid input gracefully
- 💡 Use strings.LastIndex instead of strings.Index in ParseGoVet to correctly extract file paths with colons
- 💡 Add comprehensive unit tests covering edge cases like malformed JSON, missing fields, and invalid date formats
- 💡 Implement proper error propagation instead of silently returning clean results when parsing fails
- ⚠️ Potential nil pointer dereference in ParseGolangciLintJSON when accessing issue.Pos.Filename without checking if issue.Pos exists
- ⚠️ Integer overflow in simpleAtoi function when parsing large numeric strings in ParseCoverProfile
- 🔗 The parsing functions directly affect linting and test result reporting, so any incorrect parsing can cause false positives/negatives in code quality metrics
- 🔗 Tight coupling between these parsers and specific tool output formats makes the system brittle to changes in Go tooling output format

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
- 🤖 The code is functionally correct but has several issues including missing error handling in JSON parsing, potential runtime panics from unsafe string operations, and lack of input validation that could lead to incorrect linting results.
- 💡 Add explicit nil checks before accessing nested JSON fields like issue.Pos.Filename in ParseGolangciLintJSON
- 💡 Replace simpleAtoi with a safer integer parsing function that handles overflow and invalid input gracefully
- 💡 Use strings.LastIndex instead of strings.Index in ParseGoVet to correctly extract file paths with colons
- 💡 Add comprehensive unit tests covering edge cases like malformed JSON, missing fields, and invalid date formats
- 💡 Implement proper error propagation instead of silently returning clean results when parsing fails
- ⚠️ Potential nil pointer dereference in ParseGolangciLintJSON when accessing issue.Pos.Filename without checking if issue.Pos exists
- ⚠️ Integer overflow in simpleAtoi function when parsing large numeric strings in ParseCoverProfile
- 🔗 The parsing functions directly affect linting and test result reporting, so any incorrect parsing can cause false positives/negatives in code quality metrics
- 🔗 Tight coupling between these parsers and specific tool output formats makes the system brittle to changes in Go tooling output format

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
- 🤖 The code is functionally correct but has several issues including missing error handling in JSON parsing, potential runtime panics from unsafe string operations, and lack of input validation that could lead to incorrect linting results.
- 💡 Add explicit nil checks before accessing nested JSON fields like issue.Pos.Filename in ParseGolangciLintJSON
- 💡 Replace simpleAtoi with a safer integer parsing function that handles overflow and invalid input gracefully
- 💡 Use strings.LastIndex instead of strings.Index in ParseGoVet to correctly extract file paths with colons
- 💡 Add comprehensive unit tests covering edge cases like malformed JSON, missing fields, and invalid date formats
- 💡 Implement proper error propagation instead of silently returning clean results when parsing fails
- ⚠️ Potential nil pointer dereference in ParseGolangciLintJSON when accessing issue.Pos.Filename without checking if issue.Pos exists
- ⚠️ Integer overflow in simpleAtoi function when parsing large numeric strings in ParseCoverProfile
- 🔗 The parsing functions directly affect linting and test result reporting, so any incorrect parsing can cause false positives/negatives in code quality metrics
- 🔗 Tight coupling between these parsers and specific tool output formats makes the system brittle to changes in Go tooling output format

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
- 🤖 The code is functionally correct but has several issues including missing error handling in JSON parsing, potential runtime panics from unsafe string operations, and lack of input validation that could lead to incorrect linting results.
- 💡 Add explicit nil checks before accessing nested JSON fields like issue.Pos.Filename in ParseGolangciLintJSON
- 💡 Replace simpleAtoi with a safer integer parsing function that handles overflow and invalid input gracefully
- 💡 Use strings.LastIndex instead of strings.Index in ParseGoVet to correctly extract file paths with colons
- 💡 Add comprehensive unit tests covering edge cases like malformed JSON, missing fields, and invalid date formats
- 💡 Implement proper error propagation instead of silently returning clean results when parsing fails
- ⚠️ Potential nil pointer dereference in ParseGolangciLintJSON when accessing issue.Pos.Filename without checking if issue.Pos exists
- ⚠️ Integer overflow in simpleAtoi function when parsing large numeric strings in ParseCoverProfile
- 🔗 The parsing functions directly affect linting and test result reporting, so any incorrect parsing can cause false positives/negatives in code quality metrics
- 🔗 Tight coupling between these parsers and specific tool output formats makes the system brittle to changes in Go tooling output format

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
- 🤖 The code is functionally correct but has several issues including missing error handling in JSON parsing, potential runtime panics from unsafe string operations, and lack of input validation that could lead to incorrect linting results.
- 💡 Add explicit nil checks before accessing nested JSON fields like issue.Pos.Filename in ParseGolangciLintJSON
- 💡 Replace simpleAtoi with a safer integer parsing function that handles overflow and invalid input gracefully
- 💡 Use strings.LastIndex instead of strings.Index in ParseGoVet to correctly extract file paths with colons
- 💡 Add comprehensive unit tests covering edge cases like malformed JSON, missing fields, and invalid date formats
- 💡 Implement proper error propagation instead of silently returning clean results when parsing fails
- ⚠️ Potential nil pointer dereference in ParseGolangciLintJSON when accessing issue.Pos.Filename without checking if issue.Pos exists
- ⚠️ Integer overflow in simpleAtoi function when parsing large numeric strings in ParseCoverProfile
- 🔗 The parsing functions directly affect linting and test result reporting, so any incorrect parsing can cause false positives/negatives in code quality metrics
- 🔗 Tight coupling between these parsers and specific tool output formats makes the system brittle to changes in Go tooling output format

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
- 🤖 The test file has good coverage of parsing logic but lacks proper error handling and validation for malformed inputs, with several test cases that rely on brittle assumptions about parsing behavior.
- 💡 Add tests for malformed or corrupted inputs to `ParseGoVet`, `ParseGoTestJSON`, and `ParseGolangciLintJSON` to ensure they don't panic or return incorrect results
- 💡 Replace the floating-point tolerance in `TestParseCoverProfile` with an exact assertion to make test results deterministic and easier to debug
- 💡 Add assertions in `TestParseGoVet_WithErrors` to validate that each finding contains expected fields like filename and line number
- 💡 Validate that `ParseGitLogWithAge` correctly handles date parsing edge cases (e.g., invalid date formats, timezone differences)
- ⚠️ Parsing errors from malformed input may cause panics or incorrect results in production
- ⚠️ Inaccurate coverage calculation due to floating point tolerance and missing edge case handling
- 🔗 This unit tests parsing logic that feeds evidence into certification decisions; any incorrect parsing can lead to false positives or negatives in compliance checks
- 🔗 The test suite does not cover error paths, so the reliability of the parsing logic in production is uncertain

</details>

### `internal/expiry/` (2 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`Calculate`](reports/internal-expiry-calculator-go-calculate.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`calculator_test.go`](reports/internal-expiry-calculator-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The code implements a certification expiry calculator with logical flow but contains several mathematical edge cases and potential logic flaws that could lead to incorrect window calculations.
- 💡 Validate that `cfg.DefaultWindowDays` is positive and non-zero before using it as a fallback
- 💡 Add explicit bounds checking for `factors.TestCoverage` to ensure it's within [0, 1] range before applying the multiplier formula
- 💡 Ensure that `multiplier` never drops below 0.01 to prevent zero or negative window calculations
- 💡 Add validation for `factors.PriorPassCount` and `factors.PriorFailCount` to ensure they are non-negative
- 💡 Clamp the final `days` value before rounding to prevent clamping from being bypassed by rounding
- ⚠️ Mathematical edge case in TestCoverage multiplier calculation that can produce negative multipliers
- ⚠️ Potential zero or negative expiry window due to incorrect multiplier logic
- 🔗 The function directly affects certification expiry decisions, which can impact system security and compliance
- 🔗 High coupling to domain.ExpiryConfig and domain.ExpiryFactors, making it difficult to change or test in isolation

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
- 🤖 Tests are well-structured but lack precision in time comparisons and edge case validation.
- 💡 Use `time.Now().Truncate(time.Second)` or mock time in tests to ensure deterministic behavior and avoid flaky test runs
- 💡 Add tolerance (e.g., ±1 second) when comparing `CertifiedAt` fields in time-based assertions to handle microsecond-level variance
- 💡 Add assertions that verify the calculated duration matches expected values based on known input combinations (e.g., base window + factor adjustments) to ensure deterministic output
- ⚠️ Time comparison flakiness due to reliance on `time.Now()` without mocking or tolerance
- ⚠️ Potential incorrect assertion logic in duration-based tests due to lack of tolerance for floating-point or rounding errors
- 🔗 These tests do not directly affect system components but validate the correctness of the expiry logic, which influences certification lifecycle decisions
- 🔗 Tests are tightly coupled to implementation details of `expiry.Calculate`, making them brittle if internal logic changes

</details>

### `internal/github/` (17 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`BuildIssueCloseCommand`](reports/internal-github-issues-go-buildissueclosecommand.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`BuildIssueCreateCommand`](reports/internal-github-issues-go-buildissuecreatecommand.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`BuildIssueSearchCommand`](reports/internal-github-issues-go-buildissuesearchcommand.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`BuildIssueUpdateCommand`](reports/internal-github-issues-go-buildissueupdatecommand.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatGroupedIssueBody`](reports/internal-github-issues-go-formatgroupedissuebody.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatIssueBody`](reports/internal-github-issues-go-formatissuebody.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatIssueTitle`](reports/internal-github-issues-go-formatissuetitle.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`issues_test.go`](reports/internal-github-issues-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`BuildPRCommentCommand`](reports/internal-github-pr-go-buildprcommentcommand.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`ComputeTrustDelta`](reports/internal-github-pr-go-computetrustdelta.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`FormatPRComment`](reports/internal-github-pr-go-formatprcomment.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`TrustDelta`](reports/internal-github-pr-go-trustdelta.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`pr_test.go`](reports/internal-github-pr-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`GenerateNightlyWorkflow`](reports/internal-github-workflows-go-generatenightlyworkflow.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`GeneratePRWorkflow`](reports/internal-github-workflows-go-generateprworkflow.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`GenerateWeeklyWorkflow`](reports/internal-github-workflows-go-generateweeklyworkflow.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`workflows_test.go`](reports/internal-github-workflows-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 Well-structured, readable Go code with minor correctness and security concerns in string formatting and command construction.
- 💡 Sanitize all user-provided inputs (e.g., `rec.UnitPath`, `rec.Status`) before inserting into markdown in FormatIssueBody to prevent injection
- 💡 Implement proper escaping or validation for command arguments in BuildIssueCreateCommand and related functions to mitigate shell injection risks
- 💡 Add unit tests for FormatIssueBody with edge cases such as empty observations/actions or special characters in fields
- 💡 Validate that `records` in FormatGroupedIssueBody is not nil or contains only valid CertificationRecord entries before processing
- ⚠️ Potential command injection via unsanitized inputs in BuildIssueCreateCommand and related functions
- ⚠️ Markdown injection due to lack of sanitization in FormatIssueBody and FormatGroupedIssueBody
- 🔗 The command-building functions directly construct shell commands, coupling this module tightly to the underlying system's command-line interface
- 🔗 The formatting functions are used in issue creation/update flows, meaning any malformed output could propagate to external systems like GitHub

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
- 🤖 Well-structured, readable Go code with minor correctness and security concerns in string formatting and command construction.
- 💡 Sanitize all user-provided inputs (e.g., `rec.UnitPath`, `rec.Status`) before inserting into markdown in FormatIssueBody to prevent injection
- 💡 Implement proper escaping or validation for command arguments in BuildIssueCreateCommand and related functions to mitigate shell injection risks
- 💡 Add unit tests for FormatIssueBody with edge cases such as empty observations/actions or special characters in fields
- 💡 Validate that `records` in FormatGroupedIssueBody is not nil or contains only valid CertificationRecord entries before processing
- ⚠️ Potential command injection via unsanitized inputs in BuildIssueCreateCommand and related functions
- ⚠️ Markdown injection due to lack of sanitization in FormatIssueBody and FormatGroupedIssueBody
- 🔗 The command-building functions directly construct shell commands, coupling this module tightly to the underlying system's command-line interface
- 🔗 The formatting functions are used in issue creation/update flows, meaning any malformed output could propagate to external systems like GitHub

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
- 🤖 Well-structured, readable Go code with minor correctness and security concerns in string formatting and command construction.
- 💡 Sanitize all user-provided inputs (e.g., `rec.UnitPath`, `rec.Status`) before inserting into markdown in FormatIssueBody to prevent injection
- 💡 Implement proper escaping or validation for command arguments in BuildIssueCreateCommand and related functions to mitigate shell injection risks
- 💡 Add unit tests for FormatIssueBody with edge cases such as empty observations/actions or special characters in fields
- 💡 Validate that `records` in FormatGroupedIssueBody is not nil or contains only valid CertificationRecord entries before processing
- ⚠️ Potential command injection via unsanitized inputs in BuildIssueCreateCommand and related functions
- ⚠️ Markdown injection due to lack of sanitization in FormatIssueBody and FormatGroupedIssueBody
- 🔗 The command-building functions directly construct shell commands, coupling this module tightly to the underlying system's command-line interface
- 🔗 The formatting functions are used in issue creation/update flows, meaning any malformed output could propagate to external systems like GitHub

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
- 🤖 Well-structured, readable Go code with minor correctness and security concerns in string formatting and command construction.
- 💡 Sanitize all user-provided inputs (e.g., `rec.UnitPath`, `rec.Status`) before inserting into markdown in FormatIssueBody to prevent injection
- 💡 Implement proper escaping or validation for command arguments in BuildIssueCreateCommand and related functions to mitigate shell injection risks
- 💡 Add unit tests for FormatIssueBody with edge cases such as empty observations/actions or special characters in fields
- 💡 Validate that `records` in FormatGroupedIssueBody is not nil or contains only valid CertificationRecord entries before processing
- ⚠️ Potential command injection via unsanitized inputs in BuildIssueCreateCommand and related functions
- ⚠️ Markdown injection due to lack of sanitization in FormatIssueBody and FormatGroupedIssueBody
- 🔗 The command-building functions directly construct shell commands, coupling this module tightly to the underlying system's command-line interface
- 🔗 The formatting functions are used in issue creation/update flows, meaning any malformed output could propagate to external systems like GitHub

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
- 🤖 Well-structured, readable Go code with minor correctness and security concerns in string formatting and command construction.
- 💡 Sanitize all user-provided inputs (e.g., `rec.UnitPath`, `rec.Status`) before inserting into markdown in FormatIssueBody to prevent injection
- 💡 Implement proper escaping or validation for command arguments in BuildIssueCreateCommand and related functions to mitigate shell injection risks
- 💡 Add unit tests for FormatIssueBody with edge cases such as empty observations/actions or special characters in fields
- 💡 Validate that `records` in FormatGroupedIssueBody is not nil or contains only valid CertificationRecord entries before processing
- ⚠️ Potential command injection via unsanitized inputs in BuildIssueCreateCommand and related functions
- ⚠️ Markdown injection due to lack of sanitization in FormatIssueBody and FormatGroupedIssueBody
- 🔗 The command-building functions directly construct shell commands, coupling this module tightly to the underlying system's command-line interface
- 🔗 The formatting functions are used in issue creation/update flows, meaning any malformed output could propagate to external systems like GitHub

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
- 🤖 Well-structured, readable Go code with minor correctness and security concerns in string formatting and command construction.
- 💡 Sanitize all user-provided inputs (e.g., `rec.UnitPath`, `rec.Status`) before inserting into markdown in FormatIssueBody to prevent injection
- 💡 Implement proper escaping or validation for command arguments in BuildIssueCreateCommand and related functions to mitigate shell injection risks
- 💡 Add unit tests for FormatIssueBody with edge cases such as empty observations/actions or special characters in fields
- 💡 Validate that `records` in FormatGroupedIssueBody is not nil or contains only valid CertificationRecord entries before processing
- ⚠️ Potential command injection via unsanitized inputs in BuildIssueCreateCommand and related functions
- ⚠️ Markdown injection due to lack of sanitization in FormatIssueBody and FormatGroupedIssueBody
- 🔗 The command-building functions directly construct shell commands, coupling this module tightly to the underlying system's command-line interface
- 🔗 The formatting functions are used in issue creation/update flows, meaning any malformed output could propagate to external systems like GitHub

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
- 🤖 Well-structured, readable Go code with minor correctness and security concerns in string formatting and command construction.
- 💡 Sanitize all user-provided inputs (e.g., `rec.UnitPath`, `rec.Status`) before inserting into markdown in FormatIssueBody to prevent injection
- 💡 Implement proper escaping or validation for command arguments in BuildIssueCreateCommand and related functions to mitigate shell injection risks
- 💡 Add unit tests for FormatIssueBody with edge cases such as empty observations/actions or special characters in fields
- 💡 Validate that `records` in FormatGroupedIssueBody is not nil or contains only valid CertificationRecord entries before processing
- ⚠️ Potential command injection via unsanitized inputs in BuildIssueCreateCommand and related functions
- ⚠️ Markdown injection due to lack of sanitization in FormatIssueBody and FormatGroupedIssueBody
- 🔗 The command-building functions directly construct shell commands, coupling this module tightly to the underlying system's command-line interface
- 🔗 The formatting functions are used in issue creation/update flows, meaning any malformed output could propagate to external systems like GitHub

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
- 🤖 The test file has minimal coverage and lacks proper test structure, with brittle string-based assertions that don't validate actual command construction or content formatting.
- 💡 Replace string containment checks with structured assertions to validate content format and ensure robustness
- 💡 Add tests for edge cases like empty title, invalid labels, or special characters in file paths to improve reliability
- ⚠️ Brittle string-based assertions in FormatIssueBody and FormatIssueTitle tests may fail on minor output formatting changes
- ⚠️ BuildIssueCreateCommand test does not validate that labels or title are correctly passed to the command
- 🔗 This test file only validates GitHub integration utilities, so its impact on the broader system is minimal
- 🔗 The lack of proper test structure increases risk of regressions in GitHub-related functionality

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
- 🤖 The code is functionally correct but has a critical logic flaw in trust delta calculation and lacks input validation.
- 💡 Replace fixed epsilon (0.01) in ComputeTrustDelta with a configurable tolerance or use math/big for exact comparisons
- 💡 Add input validation to BuildPRCommentCommand to ensure prNumber and body are safe for command execution
- 💡 Consider extracting the command construction logic into a separate package or interface to reduce tight coupling
- ⚠️ Floating-point comparison with fixed epsilon (0.01) in ComputeTrustDelta can cause incorrect delta tracking
- ⚠️ Potential shell injection risk in BuildPRCommentCommand if prNumber or body are not sanitized
- 🔗 ComputeTrustDelta affects downstream trust metrics and decision-making logic that relies on accurate score deltas
- 🔗 BuildPRCommentCommand tightly couples to external CLI tool usage, increasing system fragility and reducing testability

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
- 🤖 The code is functionally correct but has a critical logic flaw in trust delta calculation and lacks input validation.
- 💡 Replace fixed epsilon (0.01) in ComputeTrustDelta with a configurable tolerance or use math/big for exact comparisons
- 💡 Add input validation to BuildPRCommentCommand to ensure prNumber and body are safe for command execution
- 💡 Consider extracting the command construction logic into a separate package or interface to reduce tight coupling
- ⚠️ Floating-point comparison with fixed epsilon (0.01) in ComputeTrustDelta can cause incorrect delta tracking
- ⚠️ Potential shell injection risk in BuildPRCommentCommand if prNumber or body are not sanitized
- 🔗 ComputeTrustDelta affects downstream trust metrics and decision-making logic that relies on accurate score deltas
- 🔗 BuildPRCommentCommand tightly couples to external CLI tool usage, increasing system fragility and reducing testability

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
- 🤖 The code is functionally correct but has a critical logic flaw in trust delta calculation and lacks input validation.
- 💡 Replace fixed epsilon (0.01) in ComputeTrustDelta with a configurable tolerance or use math/big for exact comparisons
- 💡 Add input validation to BuildPRCommentCommand to ensure prNumber and body are safe for command execution
- 💡 Consider extracting the command construction logic into a separate package or interface to reduce tight coupling
- ⚠️ Floating-point comparison with fixed epsilon (0.01) in ComputeTrustDelta can cause incorrect delta tracking
- ⚠️ Potential shell injection risk in BuildPRCommentCommand if prNumber or body are not sanitized
- 🔗 ComputeTrustDelta affects downstream trust metrics and decision-making logic that relies on accurate score deltas
- 🔗 BuildPRCommentCommand tightly couples to external CLI tool usage, increasing system fragility and reducing testability

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
- 🤖 The code is functionally correct but has a critical logic flaw in trust delta calculation and lacks input validation.
- 💡 Replace fixed epsilon (0.01) in ComputeTrustDelta with a configurable tolerance or use math/big for exact comparisons
- 💡 Add input validation to BuildPRCommentCommand to ensure prNumber and body are safe for command execution
- 💡 Consider extracting the command construction logic into a separate package or interface to reduce tight coupling
- ⚠️ Floating-point comparison with fixed epsilon (0.01) in ComputeTrustDelta can cause incorrect delta tracking
- ⚠️ Potential shell injection risk in BuildPRCommentCommand if prNumber or body are not sanitized
- 🔗 ComputeTrustDelta affects downstream trust metrics and decision-making logic that relies on accurate score deltas
- 🔗 BuildPRCommentCommand tightly couples to external CLI tool usage, increasing system fragility and reducing testability

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
- 🤖 The test file has minimal logic and is well-structured, but lacks comprehensive coverage for edge cases and error conditions.
- 💡 Replace `strings.Contains` checks with more precise assertions that validate the structure and content of formatted comments using regex or structured parsing.
- 💡 Add test cases for edge conditions such as empty record slices, records with zero scores, and records with invalid or malformed paths to ensure robustness.
- ⚠️ Fragile string-based assertions in tests may pass even if the output format changes slightly, leading to silent regression.
- ⚠️ Lack of validation for edge cases such as empty record sets or invalid status values in `FormatPRComment`.
- 🔗 This test file ensures correctness of PR comment formatting logic, which is part of the GitHub integration pipeline. If these tests fail or are incomplete, it could lead to incorrect PR comments being posted.
- 🔗 The tests do not isolate or mock external dependencies, increasing coupling between this test file and the actual GitHub API or command execution.

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
- 🤖 The code is functionally correct but has critical security and maintainability issues due to hardcoded dependencies, lack of input validation, and insecure credential handling.
- 💡 Replace hardcoded `github.com/iksnae/code-certification/cmd/certify@latest` with a configurable version or tag to allow for stable, reproducible builds
- 💡 Remove `OPENROUTER_API_KEY` from the PR comment step as it's not required there and introduces unnecessary exposure
- 💡 Parameterize workflow generation to support configurable cron schedules, Go versions, and report formats
- 💡 Introduce a validation step or YAML parser to ensure generated workflows are syntactically correct before returning them
- ⚠️ Hardcoded external dependency may break or introduce breaking changes
- ⚠️ Exposure of OPENROUTER_API_KEY in PR comment step increases risk of credential leakage
- 🔗 These functions tightly couple the system to a specific external tool (`certify`) and version, increasing maintenance burden
- 🔗 The inconsistent permissions across workflows increase the risk of privilege escalation or misconfiguration in CI/CD pipelines

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
- 🤖 The code is functionally correct but has critical security and maintainability issues due to hardcoded dependencies, lack of input validation, and insecure credential handling.
- 💡 Replace hardcoded `github.com/iksnae/code-certification/cmd/certify@latest` with a configurable version or tag to allow for stable, reproducible builds
- 💡 Remove `OPENROUTER_API_KEY` from the PR comment step as it's not required there and introduces unnecessary exposure
- 💡 Parameterize workflow generation to support configurable cron schedules, Go versions, and report formats
- 💡 Introduce a validation step or YAML parser to ensure generated workflows are syntactically correct before returning them
- ⚠️ Hardcoded external dependency may break or introduce breaking changes
- ⚠️ Exposure of OPENROUTER_API_KEY in PR comment step increases risk of credential leakage
- 🔗 These functions tightly couple the system to a specific external tool (`certify`) and version, increasing maintenance burden
- 🔗 The inconsistent permissions across workflows increase the risk of privilege escalation or misconfiguration in CI/CD pipelines

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
- 🤖 The code is functionally correct but has critical security and maintainability issues due to hardcoded dependencies, lack of input validation, and insecure credential handling.
- 💡 Replace hardcoded `github.com/iksnae/code-certification/cmd/certify@latest` with a configurable version or tag to allow for stable, reproducible builds
- 💡 Remove `OPENROUTER_API_KEY` from the PR comment step as it's not required there and introduces unnecessary exposure
- 💡 Parameterize workflow generation to support configurable cron schedules, Go versions, and report formats
- 💡 Introduce a validation step or YAML parser to ensure generated workflows are syntactically correct before returning them
- ⚠️ Hardcoded external dependency may break or introduce breaking changes
- ⚠️ Exposure of OPENROUTER_API_KEY in PR comment step increases risk of credential leakage
- 🔗 These functions tightly couple the system to a specific external tool (`certify`) and version, increasing maintenance burden
- 🔗 The inconsistent permissions across workflows increase the risk of privilege escalation or misconfiguration in CI/CD pipelines

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
- 🤖 The tests are brittle and rely on string matching instead of structured YAML validation, leading to fragile test cases and poor maintainability.
- 💡 Replace string-based assertions with YAML parsing and validation to ensure generated workflows are syntactically correct and semantically meaningful
- 💡 Add tests for edge cases such as empty or malformed inputs to the workflow generation functions, and verify that errors are handled appropriately
- 💡 Use table-driven tests with structured data to validate specific fields like triggers, jobs, and secrets in the generated YAML
- ⚠️ Brittle string-based assertions may break on whitespace or formatting changes, causing false test failures
- ⚠️ Lack of YAML parsing validation means invalid workflows could pass tests without detection
- 🔗 These tests provide minimal confidence in the correctness of generated workflows, increasing risk of deployment issues
- 🔗 The reliance on string matching increases coupling between test logic and implementation details, making refactoring harder

</details>

### `internal/override/` (9 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`Apply`](reports/internal-override-applier-go-apply.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`ApplyAll`](reports/internal-override-applier-go-applyall.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`applier_test.go`](reports/internal-override-applier-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`LoadDir`](reports/internal-override-loader-go-loaddir.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`LoadFile`](reports/internal-override-loader-go-loadfile.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`parseAction`](reports/internal-override-loader-go-parseaction.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`rawOverride`](reports/internal-override-loader-go-rawoverride.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`rawOverrideFile`](reports/internal-override-loader-go-rawoverridefile.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`loader_test.go`](reports/internal-override-loader-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The code correctly applies overrides to certification records but has a critical logic flaw in window extension and shortening that can lead to incorrect expiration times.
- 💡 Fix the time math in OverrideExtendWindow: use `rec.ExpiresAt = rec.ExpiresAt.Add(remaining / 2)` to correctly extend the window by half of the remaining time
- 💡 Fix the time math in OverrideShortenWindow: ensure that setting `rec.ExpiresAt` to half of the original window is done correctly by computing it as `rec.CertifiedAt.Add(remaining / 2)`
- 💡 Add unit tests for edge cases such as zero or negative time differences to validate behavior under unusual conditions
- ⚠️ Incorrect time calculation in OverrideExtendWindow and OverrideShortenWindow logic
- ⚠️ Potential for incorrect certification window behavior due to flawed time math
- 🔗 This unit directly modifies certification record state, affecting downstream systems that depend on accurate expiration times
- 🔗 Improper window handling can cause cascading failures in recertification workflows

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
- 🤖 The code correctly applies overrides to certification records but has a critical logic flaw in window extension and shortening that can lead to incorrect expiration times.
- 💡 Fix the time math in OverrideExtendWindow: use `rec.ExpiresAt = rec.ExpiresAt.Add(remaining / 2)` to correctly extend the window by half of the remaining time
- 💡 Fix the time math in OverrideShortenWindow: ensure that setting `rec.ExpiresAt` to half of the original window is done correctly by computing it as `rec.CertifiedAt.Add(remaining / 2)`
- 💡 Add unit tests for edge cases such as zero or negative time differences to validate behavior under unusual conditions
- ⚠️ Incorrect time calculation in OverrideExtendWindow and OverrideShortenWindow logic
- ⚠️ Potential for incorrect certification window behavior due to flawed time math
- 🔗 This unit directly modifies certification record state, affecting downstream systems that depend on accurate expiration times
- 🔗 Improper window handling can cause cascading failures in recertification workflows

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
- 🤖 The test suite has good coverage for override logic but lacks proper error handling and deterministic time-based assertions.
- 💡 Use a fixed time reference (e.g., `time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)`) in tests to make time-based assertions deterministic
- 💡 Add assertions to verify that the returned CertificationRecord is a new instance (e.g., by checking pointer equality or deep copy behavior)
- 💡 Add tests for invalid override actions and edge cases like missing rationale or actor fields
- 💡 Ensure `ApplyAll` correctly applies only the matching override and not others in the list
- ⚠️ Non-deterministic time-based assertions may cause flaky tests due to reliance on `time.Now()`
- ⚠️ Potential mutation of input records rather than returning a new copy, leading to unintended side effects
- 🔗 These tests validate behavior that directly affects certification status and expiration logic in the broader system
- 🔗 If these tests pass but don't properly validate override application, they may mask bugs in the actual `override.Apply` logic

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
- 🤖 The code is functionally correct but has critical security and validation gaps in file handling and input parsing.
- 💡 Validate and sanitize file paths before reading to prevent directory traversal attacks, e.g., use filepath.Clean(path) and check against allowed base directories
- 💡 Add input validation for Rationale and Actor fields to ensure they are not empty or contain unexpected characters before constructing domain.Override objects
- 💡 Make parseAction case-insensitive by converting input to lowercase before lookup or use a more robust enum-like structure with validation
- 💡 Consider using a configurable base directory for LoadDir to avoid loading from arbitrary locations
- ⚠️ Directory traversal vulnerability due to unsanitized file paths in LoadDir
- ⚠️ Lack of input sanitization for Rationale and Actor fields leading to potential malformed domain objects
- 🔗 The function allows loading of arbitrary YAML files from disk, increasing attack surface if used with untrusted input
- 🔗 The hardcoded string map in parseAction introduces tight coupling to specific action values and makes future extensibility difficult

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
- 🤖 The code is functionally correct but has critical security and validation gaps in file handling and input parsing.
- 💡 Validate and sanitize file paths before reading to prevent directory traversal attacks, e.g., use filepath.Clean(path) and check against allowed base directories
- 💡 Add input validation for Rationale and Actor fields to ensure they are not empty or contain unexpected characters before constructing domain.Override objects
- 💡 Make parseAction case-insensitive by converting input to lowercase before lookup or use a more robust enum-like structure with validation
- 💡 Consider using a configurable base directory for LoadDir to avoid loading from arbitrary locations
- ⚠️ Directory traversal vulnerability due to unsanitized file paths in LoadDir
- ⚠️ Lack of input sanitization for Rationale and Actor fields leading to potential malformed domain objects
- 🔗 The function allows loading of arbitrary YAML files from disk, increasing attack surface if used with untrusted input
- 🔗 The hardcoded string map in parseAction introduces tight coupling to specific action values and makes future extensibility difficult

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
- 🤖 The code is functionally correct but has critical security and validation gaps in file handling and input parsing.
- 💡 Validate and sanitize file paths before reading to prevent directory traversal attacks, e.g., use filepath.Clean(path) and check against allowed base directories
- 💡 Add input validation for Rationale and Actor fields to ensure they are not empty or contain unexpected characters before constructing domain.Override objects
- 💡 Make parseAction case-insensitive by converting input to lowercase before lookup or use a more robust enum-like structure with validation
- 💡 Consider using a configurable base directory for LoadDir to avoid loading from arbitrary locations
- ⚠️ Directory traversal vulnerability due to unsanitized file paths in LoadDir
- ⚠️ Lack of input sanitization for Rationale and Actor fields leading to potential malformed domain objects
- 🔗 The function allows loading of arbitrary YAML files from disk, increasing attack surface if used with untrusted input
- 🔗 The hardcoded string map in parseAction introduces tight coupling to specific action values and makes future extensibility difficult

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
- 🤖 The code is functionally correct but has critical security and validation gaps in file handling and input parsing.
- 💡 Validate and sanitize file paths before reading to prevent directory traversal attacks, e.g., use filepath.Clean(path) and check against allowed base directories
- 💡 Add input validation for Rationale and Actor fields to ensure they are not empty or contain unexpected characters before constructing domain.Override objects
- 💡 Make parseAction case-insensitive by converting input to lowercase before lookup or use a more robust enum-like structure with validation
- 💡 Consider using a configurable base directory for LoadDir to avoid loading from arbitrary locations
- ⚠️ Directory traversal vulnerability due to unsanitized file paths in LoadDir
- ⚠️ Lack of input sanitization for Rationale and Actor fields leading to potential malformed domain objects
- 🔗 The function allows loading of arbitrary YAML files from disk, increasing attack surface if used with untrusted input
- 🔗 The hardcoded string map in parseAction introduces tight coupling to specific action values and makes future extensibility difficult

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
- 🤖 The code is functionally correct but has critical security and validation gaps in file handling and input parsing.
- 💡 Validate and sanitize file paths before reading to prevent directory traversal attacks, e.g., use filepath.Clean(path) and check against allowed base directories
- 💡 Add input validation for Rationale and Actor fields to ensure they are not empty or contain unexpected characters before constructing domain.Override objects
- 💡 Make parseAction case-insensitive by converting input to lowercase before lookup or use a more robust enum-like structure with validation
- 💡 Consider using a configurable base directory for LoadDir to avoid loading from arbitrary locations
- ⚠️ Directory traversal vulnerability due to unsanitized file paths in LoadDir
- ⚠️ Lack of input sanitization for Rationale and Actor fields leading to potential malformed domain objects
- 🔗 The function allows loading of arbitrary YAML files from disk, increasing attack surface if used with untrusted input
- 🔗 The hardcoded string map in parseAction introduces tight coupling to specific action values and makes future extensibility difficult

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
- 🤖 The test file has good coverage for basic functionality but lacks comprehensive validation of override data and does not test edge cases or error conditions beyond file existence.
- 💡 Use `t.Cleanup` or ensure temp directory cleanup in `TestLoadOverrides_Empty` to prevent file system pollution
- 💡 Add assertions for all loaded overrides' fields (e.g., `Reason`, `Window`, `Expiry`) in `TestLoadOverrides` instead of just the first entry
- 💡 Refactor `TestLoadOverrides_Actions` to avoid redundant file loading and instead iterate over the already-loaded overrides
- 💡 Replace hardcoded absolute path in `TestLoadOverrides_NotFound` with a relative or dynamically constructed path to improve portability
- 💡 Add a test case for `LoadFile` with invalid YAML content to verify proper error handling
- ⚠️ Hardcoded relative path in `testdataPath` function may break if working directory changes or test is run from different location
- ⚠️ No validation of actual override content beyond first item and action type, leading to potential silent failures in data parsing or structure
- 🔗 This unit tests the override loading logic in isolation, but lacks integration with domain validation or mock configuration scenarios
- 🔗 The test suite does not cover edge cases like malformed YAML, duplicate unit IDs, or invalid actions, which could propagate errors into the broader override system

</details>

### `internal/policy/` (14 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`Evaluate`](reports/internal-policy-evaluator-go-evaluate.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`EvaluationResult`](reports/internal-policy-evaluator-go-evaluationresult.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`evaluateRule`](reports/internal-policy-evaluator-go-evaluaterule.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`extractComplexity`](reports/internal-policy-evaluator-go-extractcomplexity.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`extractCoverage`](reports/internal-policy-evaluator-go-extractcoverage.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`extractMetric`](reports/internal-policy-evaluator-go-extractmetric.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`extractTodoCount`](reports/internal-policy-evaluator-go-extracttodocount.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`evaluator_test.go`](reports/internal-policy-evaluator-test-go.md) | file | B+ | 87.2% | certified | 2026-04-23 |
| [`Match`](reports/internal-policy-matcher-go-match.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Matcher`](reports/internal-policy-matcher-go-matcher.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`NewMatcher`](reports/internal-policy-matcher-go-newmatcher.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`matchPath`](reports/internal-policy-matcher-go-matchpath.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`matchesPack`](reports/internal-policy-matcher-go-matchespack.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`matcher_test.go`](reports/internal-policy-matcher-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The policy evaluation logic is functionally correct but has several maintainability and robustness issues including brittle string parsing, inconsistent evidence handling, and lack of input validation.
- 💡 Replace string-based parsing in extractTodoCount with structured data access or proper regex matching to make it resilient to format changes
- 💡 Add proper error handling and type assertion checks in extractMetric before returning values to prevent runtime panics
- 💡 Add unit tests for edge cases like malformed summaries, missing fields in evidence details, and inconsistent evidence types
- 💡 Implement a more robust validation or schema-based extraction for evidence fields instead of relying on string parsing
- ⚠️ String parsing in extractTodoCount and extractComplexity is brittle and will fail on any format change
- ⚠️ Missing input validation for evidence details can lead to runtime panics or incorrect evaluation results
- 🔗 This unit tightly couples to specific evidence detail formats and summary string structures, making it fragile to changes in evidence generation
- 🔗 The function extractMetric has inconsistent behavior for different evidence kinds, creating unpredictable evaluation outcomes

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
- 🤖 The policy evaluation logic is functionally correct but has several maintainability and robustness issues including brittle string parsing, inconsistent evidence handling, and lack of input validation.
- 💡 Replace string-based parsing in extractTodoCount with structured data access or proper regex matching to make it resilient to format changes
- 💡 Add proper error handling and type assertion checks in extractMetric before returning values to prevent runtime panics
- 💡 Add unit tests for edge cases like malformed summaries, missing fields in evidence details, and inconsistent evidence types
- 💡 Implement a more robust validation or schema-based extraction for evidence fields instead of relying on string parsing
- ⚠️ String parsing in extractTodoCount and extractComplexity is brittle and will fail on any format change
- ⚠️ Missing input validation for evidence details can lead to runtime panics or incorrect evaluation results
- 🔗 This unit tightly couples to specific evidence detail formats and summary string structures, making it fragile to changes in evidence generation
- 🔗 The function extractMetric has inconsistent behavior for different evidence kinds, creating unpredictable evaluation outcomes

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
- 🤖 The policy evaluation logic is functionally correct but has several maintainability and robustness issues including brittle string parsing, inconsistent evidence handling, and lack of input validation.
- 💡 Replace string-based parsing in extractTodoCount with structured data access or proper regex matching to make it resilient to format changes
- 💡 Add proper error handling and type assertion checks in extractMetric before returning values to prevent runtime panics
- 💡 Add unit tests for edge cases like malformed summaries, missing fields in evidence details, and inconsistent evidence types
- 💡 Implement a more robust validation or schema-based extraction for evidence fields instead of relying on string parsing
- ⚠️ String parsing in extractTodoCount and extractComplexity is brittle and will fail on any format change
- ⚠️ Missing input validation for evidence details can lead to runtime panics or incorrect evaluation results
- 🔗 This unit tightly couples to specific evidence detail formats and summary string structures, making it fragile to changes in evidence generation
- 🔗 The function extractMetric has inconsistent behavior for different evidence kinds, creating unpredictable evaluation outcomes

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
- 🤖 The policy evaluation logic is functionally correct but has several maintainability and robustness issues including brittle string parsing, inconsistent evidence handling, and lack of input validation.
- 💡 Replace string-based parsing in extractTodoCount with structured data access or proper regex matching to make it resilient to format changes
- 💡 Add proper error handling and type assertion checks in extractMetric before returning values to prevent runtime panics
- 💡 Add unit tests for edge cases like malformed summaries, missing fields in evidence details, and inconsistent evidence types
- 💡 Implement a more robust validation or schema-based extraction for evidence fields instead of relying on string parsing
- ⚠️ String parsing in extractTodoCount and extractComplexity is brittle and will fail on any format change
- ⚠️ Missing input validation for evidence details can lead to runtime panics or incorrect evaluation results
- 🔗 This unit tightly couples to specific evidence detail formats and summary string structures, making it fragile to changes in evidence generation
- 🔗 The function extractMetric has inconsistent behavior for different evidence kinds, creating unpredictable evaluation outcomes

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
- 🤖 The code has functional correctness but suffers from fragile string parsing, poor error handling in extraction logic, and potential runtime panics due to unchecked type assertions.
- 💡 Replace hardcoded string parsing in `extractCoverage` with a regex pattern or structured data extraction if possible to make it more resilient to format changes
- 💡 Add bounds checking and input validation in `extractTodoCount` before accessing string indices to prevent runtime panics
- 💡 Standardize return values from `extractMetric` family functions to always return 0 for missing metrics instead of mixing -1 and 0
- 💡 Add logging or error reporting when type assertions fail in `extractTodoCount` and `extractComplexity` to aid debugging
- 💡 Validate that `e.Summary` is not empty before attempting to parse it in `extractCoverage` and related functions
- ⚠️ Fragile string parsing in `extractCoverage` can silently misparse coverage values or fail to parse at all if format changes
- ⚠️ Index out of bounds panic in `extractTodoCount` due to unchecked loop bounds when walking backwards through summary string
- 🔗 The `extractCoverage` function is tightly coupled to test evidence format, making it fragile to changes in output formatting from external tools like test runners
- 🔗 The inconsistent return values from `extractMetric` family functions increase coupling between this module and callers who must handle special cases for -1 vs. 0

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
- 🤖 The policy evaluation logic is functionally correct but has several maintainability and robustness issues including brittle string parsing, inconsistent evidence handling, and lack of input validation.
- 💡 Replace string-based parsing in extractTodoCount with structured data access or proper regex matching to make it resilient to format changes
- 💡 Add proper error handling and type assertion checks in extractMetric before returning values to prevent runtime panics
- 💡 Add unit tests for edge cases like malformed summaries, missing fields in evidence details, and inconsistent evidence types
- 💡 Implement a more robust validation or schema-based extraction for evidence fields instead of relying on string parsing
- ⚠️ String parsing in extractTodoCount and extractComplexity is brittle and will fail on any format change
- ⚠️ Missing input validation for evidence details can lead to runtime panics or incorrect evaluation results
- 🔗 This unit tightly couples to specific evidence detail formats and summary string structures, making it fragile to changes in evidence generation
- 🔗 The function extractMetric has inconsistent behavior for different evidence kinds, creating unpredictable evaluation outcomes

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
- 🤖 The policy evaluation logic is functionally correct but has several maintainability and robustness issues including brittle string parsing, inconsistent evidence handling, and lack of input validation.
- 💡 Replace string-based parsing in extractTodoCount with structured data access or proper regex matching to make it resilient to format changes
- 💡 Add proper error handling and type assertion checks in extractMetric before returning values to prevent runtime panics
- 💡 Add unit tests for edge cases like malformed summaries, missing fields in evidence details, and inconsistent evidence types
- 💡 Implement a more robust validation or schema-based extraction for evidence fields instead of relying on string parsing
- ⚠️ String parsing in extractTodoCount and extractComplexity is brittle and will fail on any format change
- ⚠️ Missing input validation for evidence details can lead to runtime panics or incorrect evaluation results
- 🔗 This unit tightly couples to specific evidence detail formats and summary string structures, making it fragile to changes in evidence generation
- 🔗 The function extractMetric has inconsistent behavior for different evidence kinds, creating unpredictable evaluation outcomes

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
- 🤖 The test suite for policy evaluation is functionally correct but lacks comprehensive edge case coverage and has potential logic gaps in violation handling.
- 💡 Add a test case for evidence that does not match any policy rule to ensure proper handling of unmatched metrics
- 💡 Add a test case for a warning rule with threshold > 0 that is violated to verify it still passes but generates a violation
- 💡 Add assertions for the full structure of violations (e.g., Severity, Dimension, Metric) to ensure correctness beyond just existence
- ⚠️ Missing evidence handling behavior is not clearly defined, leading to potential misinterpretation of policy outcomes
- ⚠️ No validation that violation details (severity, dimension, metric) are correctly populated in the result
- 🔗 This test file only affects the policy evaluation unit and does not introduce coupling to external systems
- 🔗 If the policy evaluator's behavior around missing evidence changes, these tests may silently pass while actual logic fails

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
- 🤖 The code has functional policy matching logic but contains critical path matching bugs, missing error handling, and lacks test coverage for edge cases.
- 💡 Replace `matchPath` with a more robust glob matching implementation using a dedicated library like `github.com/gobwas/glob` or fix the current logic to properly handle nested wildcards and path normalization
- 💡 Add proper error handling around `filepath.Match` calls to ensure malformed patterns are logged or rejected rather than silently ignored
- ⚠️ Incorrect glob pattern matching due to improper handling of nested wildcards and path components
- ⚠️ Silent failure in glob pattern matching due to unchecked errors from filepath.Match
- 🔗 This unit introduces a critical flaw in policy matching that can lead to incorrect policy application, affecting system-wide compliance and security checks
- 🔗 The path matching logic is tightly coupled to filesystem assumptions that may not hold in all environments (e.g., case sensitivity, path separators)

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
- 🤖 The code has functional policy matching logic but contains critical path matching bugs, missing error handling, and lacks test coverage for edge cases.
- 💡 Replace `matchPath` with a more robust glob matching implementation using a dedicated library like `github.com/gobwas/glob` or fix the current logic to properly handle nested wildcards and path normalization
- 💡 Add proper error handling around `filepath.Match` calls to ensure malformed patterns are logged or rejected rather than silently ignored
- ⚠️ Incorrect glob pattern matching due to improper handling of nested wildcards and path components
- ⚠️ Silent failure in glob pattern matching due to unchecked errors from filepath.Match
- 🔗 This unit introduces a critical flaw in policy matching that can lead to incorrect policy application, affecting system-wide compliance and security checks
- 🔗 The path matching logic is tightly coupled to filesystem assumptions that may not hold in all environments (e.g., case sensitivity, path separators)

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
- 🤖 The code has functional policy matching logic but contains critical path matching bugs, missing error handling, and lacks test coverage for edge cases.
- 💡 Replace `matchPath` with a more robust glob matching implementation using a dedicated library like `github.com/gobwas/glob` or fix the current logic to properly handle nested wildcards and path normalization
- 💡 Add proper error handling around `filepath.Match` calls to ensure malformed patterns are logged or rejected rather than silently ignored
- ⚠️ Incorrect glob pattern matching due to improper handling of nested wildcards and path components
- ⚠️ Silent failure in glob pattern matching due to unchecked errors from filepath.Match
- 🔗 This unit introduces a critical flaw in policy matching that can lead to incorrect policy application, affecting system-wide compliance and security checks
- 🔗 The path matching logic is tightly coupled to filesystem assumptions that may not hold in all environments (e.g., case sensitivity, path separators)

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
- 🤖 The code has functional policy matching logic but contains critical path matching bugs, missing error handling, and lacks test coverage for edge cases.
- 💡 Replace `matchPath` with a more robust glob matching implementation using a dedicated library like `github.com/gobwas/glob` or fix the current logic to properly handle nested wildcards and path normalization
- 💡 Add proper error handling around `filepath.Match` calls to ensure malformed patterns are logged or rejected rather than silently ignored
- ⚠️ Incorrect glob pattern matching due to improper handling of nested wildcards and path components
- ⚠️ Silent failure in glob pattern matching due to unchecked errors from filepath.Match
- 🔗 This unit introduces a critical flaw in policy matching that can lead to incorrect policy application, affecting system-wide compliance and security checks
- 🔗 The path matching logic is tightly coupled to filesystem assumptions that may not hold in all environments (e.g., case sensitivity, path separators)

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
- 🤖 The code has functional policy matching logic but contains critical path matching bugs, missing error handling, and lacks test coverage for edge cases.
- 💡 Replace `matchPath` with a more robust glob matching implementation using a dedicated library like `github.com/gobwas/glob` or fix the current logic to properly handle nested wildcards and path normalization
- 💡 Add proper error handling around `filepath.Match` calls to ensure malformed patterns are logged or rejected rather than silently ignored
- ⚠️ Incorrect glob pattern matching due to improper handling of nested wildcards and path components
- ⚠️ Silent failure in glob pattern matching due to unchecked errors from filepath.Match
- 🔗 This unit introduces a critical flaw in policy matching that can lead to incorrect policy application, affecting system-wide compliance and security checks
- 🔗 The path matching logic is tightly coupled to filesystem assumptions that may not hold in all environments (e.g., case sensitivity, path separators)

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
- 🤖 The test suite for the policy matcher is functionally correct but lacks comprehensive coverage and has some structural weaknesses in test design.
- 💡 Add tests for multiple matching packs to ensure proper precedence or conflict resolution behavior
- 💡 Include edge case tests for malformed path patterns, empty language strings, and units without defined paths or languages
- 💡 Refactor `TestMatcher_GlobalMatchesAll` to explicitly verify that the matched pack has a blank language field rather than just asserting count
- 💡 Add concurrent access tests to ensure thread safety if the matcher is intended for use in a multi-threaded environment
- ⚠️ Missing validation of matched pack content or metadata, which could lead to silent misbehavior if the matcher returns packs with incorrect configurations
- ⚠️ No test coverage for concurrent access or race conditions in matcher operations, which could cause issues in multi-threaded environments
- 🔗 This test file only affects the local unit testing of policy matching logic, but its design can mask deeper integration issues if not expanded to cover more complex scenarios
- 🔗 The use of hardcoded strings like 'go-standard', 'ts-standard', and 'global' in tests increases coupling to specific policy pack naming conventions, which may not be resilient to refactoring

</details>

### `internal/queue/` (17 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`BatchNext`](reports/internal-queue-queue-go-batchnext.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`Complete`](reports/internal-queue-queue-go-complete.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Enqueue`](reports/internal-queue-queue-go-enqueue.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Fail`](reports/internal-queue-queue-go-fail.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Item`](reports/internal-queue-queue-go-item.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`ItemStatus`](reports/internal-queue-queue-go-itemstatus.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Len`](reports/internal-queue-queue-go-len.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Load`](reports/internal-queue-queue-go-load.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`New`](reports/internal-queue-queue-go-new.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Next`](reports/internal-queue-queue-go-next.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Queue`](reports/internal-queue-queue-go-queue.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Reset`](reports/internal-queue-queue-go-reset.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Save`](reports/internal-queue-queue-go-save.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Skip`](reports/internal-queue-queue-go-skip.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Stats`](reports/internal-queue-queue-go-stats.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`persistedQueue`](reports/internal-queue-queue-go-persistedqueue.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`queue_test.go`](reports/internal-queue-queue-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The Len() method is functionally correct but has a subtle concurrency risk due to lack of synchronization in the Queue struct.
- 💡 Add a mutex to the Queue struct and synchronize all public methods (Len, Enqueue, Next, etc.) to make it safe for concurrent use
- 💡 Consider adding a unit test that verifies the behavior of Len() under concurrent access to expose race conditions
- ⚠️ Race condition in concurrent access to Queue methods (e.g., Len(), Enqueue(), Next()) due to missing synchronization
- ⚠️ Inconsistent Stats() logic where retryable failed items are counted as pending, potentially causing confusion in metrics
- 🔗 The Queue type is not thread-safe, which can lead to data races and inconsistent state in multi-threaded environments
- 🔗 The Load() function resets StatusInProgress items to StatusPending, which can cause issues if the queue is used concurrently or if items are being processed while loading

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
- 🤖 The test suite is comprehensive and well-structured, covering most functional and edge cases for the queue implementation, but has some gaps in error handling coverage and lacks assertions on internal state consistency.
- 💡 Add assertions in `TestQueue_Enqueue_NoDuplicates` to verify that the correct item (first added) is retained after duplicate insertion
- 💡 Replace `t.Fatal` with `t.Errorf` in `TestQueue_Next_ExhaustsMaxRetries` to maintain consistent error reporting style
- 💡 Add a test case for `TestQueue_SaveLoad` that validates the content of the saved JSON file matches expected structure
- 💡 Add a test for `TestQueue_BatchNext` that verifies items are returned in FIFO order and that they're marked as in progress
- 💡 Add a test for `TestQueue_Save` to verify that it correctly handles invalid file paths (e.g., directory does not exist)
- ⚠️ Missing assertions on internal state consistency can mask bugs in duplicate handling logic
- ⚠️ Inconsistent error handling practices make it harder to distinguish between recoverable and fatal test failures
- 🔗 This test suite provides strong coverage for the queue's public API surface, ensuring that all major operations behave as documented
- 🔗 The tests do not directly impact other modules but rely heavily on the queue implementation's correctness to pass

</details>

### `internal/record/` (17 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`AppendHistory`](reports/internal-record-store-go-appendhistory.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`ListAll`](reports/internal-record-store-go-listall.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`Load`](reports/internal-record-store-go-load.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`LoadHistory`](reports/internal-record-store-go-loadhistory.md) | method | B+ | 88.3% | certified | 2026-04-23 |
| [`NewStore`](reports/internal-record-store-go-newstore.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Save`](reports/internal-record-store-go-save.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`Store`](reports/internal-record-store-go-store.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`dimensionsToMap`](reports/internal-record-store-go-dimensionstomap.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`fromJSON`](reports/internal-record-store-go-fromjson.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`historyEntry`](reports/internal-record-store-go-historyentry.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`historyPathFor`](reports/internal-record-store-go-historypathfor.md) | method | B+ | 89.4% | certified | 2026-04-23 |
| [`mapToDimensions`](reports/internal-record-store-go-maptodimensions.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`parseGrade`](reports/internal-record-store-go-parsegrade.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`pathFor`](reports/internal-record-store-go-pathfor.md) | method | B+ | 87.8% | certified | 2026-04-24 |
| [`recordJSON`](reports/internal-record-store-go-recordjson.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`toJSON`](reports/internal-record-store-go-tojson.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`store_test.go`](reports/internal-record-store-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

</details>

<a id="internal-record-store-go-pathfor"></a>
<details>
<summary>pathFor — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 95.0% |
| security | 80.0% |
| testability | 90.0% |

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The code is functionally correct but has several maintainability and error handling issues, including unhandled errors in parsing and potential security risks from file naming.
- 💡 Replace `domain.ParseUnitID(...)` with error checking and return an appropriate error instead of ignoring the error
- 💡 Change `pathFor` to use full 32-byte hash or implement a more robust collision-resistant naming scheme
- 💡 Add logging or error propagation in `ListAll` when individual file reads fail
- 💡 Use sync.RWMutex or similar to protect concurrent access to files in the store
- 💡 Make file permissions configurable via environment variables or config options
- ⚠️ Uncaught parsing errors from domain.Parse* and time.Parse in fromJSON
- ⚠️ SHA256 hash collision risk in pathFor due to truncation to 8 bytes
- 🔗 The Store type tightly couples to filesystem I/O, making it hard to swap out for other storage backends like databases or cloud storage
- 🔗 The pathFor method creates a flat directory structure that can become unwieldy with many records and increases risk of name collisions

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
- 🤖 The test suite for the record store is functionally correct but lacks robustness in error handling, file system assumptions, and test isolation.
- 💡 Replace `filepath.Glob(pattern)` with explicit checks for the expected file path using `filepath.Join(dir, "records", ...)` to avoid brittle globbing logic and ensure compatibility across platforms
- 💡 Add assertions for error types in TestStore_LoadNotFound to validate that the returned error is of the expected kind (e.g., domain.ErrRecordNotFound) rather than just checking for nil
- 💡 Use `t.Cleanup()` to ensure temporary directories are cleaned up properly and isolate each test's environment more strictly
- 💡 Validate that the saved JSON file content matches expectations by reading back and comparing fields to ensure persistence correctness
- ⚠️ Assumption of internal file structure in TestStore_PathFormat may break on different platforms or future refactors
- ⚠️ Lack of error validation in TestStore_SaveAndLoad can mask silent failures in the Save operation
- 🔗 This unit tests a component that interacts with the file system and persists domain objects; any incorrect assumptions about paths or file formats can lead to data loss or corruption
- 🔗 The test suite does not isolate tests properly, as each uses t.TempDir() which creates a new directory per test but doesn't ensure clean state between tests in case of shared resources or side effects

</details>

### `internal/report/` (62 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`Badge`](reports/internal-report-badge-go-badge.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`BadgeMarkdown`](reports/internal-report-badge-go-badgemarkdown.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatBadgeJSON`](reports/internal-report-badge-go-formatbadgejson.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`GenerateBadge`](reports/internal-report-badge-go-generatebadge.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`badgeColor`](reports/internal-report-badge-go-badgecolor.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`badgeMessage`](reports/internal-report-badge-go-badgemessage.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`badge_test.go`](reports/internal-report-badge-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`Card`](reports/internal-report-card-go-card.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatCardMarkdown`](reports/internal-report-card-go-formatcardmarkdown.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`FormatCardText`](reports/internal-report-card-go-formatcardtext.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`GenerateCard`](reports/internal-report-card-go-generatecard.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`IssueCard`](reports/internal-report-card-go-issuecard.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`LanguageCard`](reports/internal-report-card-go-languagecard.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`buildLanguageCards`](reports/internal-report-card-go-buildlanguagecards.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`buildTopIssues`](reports/internal-report-card-go-buildtopissues.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`gradeEmoji`](reports/internal-report-card-go-gradeemoji.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`card_test.go`](reports/internal-report-card-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`AreaSummary`](reports/internal-report-detailed-go-areasummary.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`Detailed`](reports/internal-report-detailed-go-detailed.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`DetailedReport`](reports/internal-report-detailed-go-detailedreport.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatDetailedText`](reports/internal-report-detailed-go-formatdetailedtext.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`LanguageBreakdown`](reports/internal-report-detailed-go-languagebreakdown.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`UnitSummary`](reports/internal-report-detailed-go-unitsummary.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`computeDimensionAverages`](reports/internal-report-detailed-go-computedimensionaverages.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`computeLanguageBreakdowns`](reports/internal-report-detailed-go-computelanguagebreakdowns.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`explainStatus`](reports/internal-report-detailed-go-explainstatus.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`findExpiringSoon`](reports/internal-report-detailed-go-findexpiringsoon.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`findFailing`](reports/internal-report-detailed-go-findfailing.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`findHighestRisk`](reports/internal-report-detailed-go-findhighestrisk.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`findRecurrentlyFailing`](reports/internal-report-detailed-go-findrecurrentlyfailing.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`unitSummaryFrom`](reports/internal-report-detailed-go-unitsummaryfrom.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`detailed_test.go`](reports/internal-report-detailed-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`FormatFullMarkdown`](reports/internal-report-full-go-formatfullmarkdown.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FullReport`](reports/internal-report-full-go-fullreport.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`GenerateFullReport`](reports/internal-report-full-go-generatefullreport.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`LanguageDetail`](reports/internal-report-full-go-languagedetail.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`UnitReport`](reports/internal-report-full-go-unitreport.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`buildLanguageDetail`](reports/internal-report-full-go-buildlanguagedetail.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`dirOf`](reports/internal-report-full-go-dirof.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`shortFile`](reports/internal-report-full-go-shortfile.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`sortedKeys`](reports/internal-report-full-go-sortedkeys.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`unitAnchor`](reports/internal-report-full-go-unitanchor.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`unitReportFrom`](reports/internal-report-full-go-unitreportfrom.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`writeAIInsights`](reports/internal-report-full-go-writeaiinsights.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`writeAllUnits`](reports/internal-report-full-go-writeallunits.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`writeDimensionAverages`](reports/internal-report-full-go-writedimensionaverages.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`writeGradeDistribution`](reports/internal-report-full-go-writegradedistribution.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`writeHeader`](reports/internal-report-full-go-writeheader.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`writeLanguageDetail`](reports/internal-report-full-go-writelanguagedetail.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`writeSummary`](reports/internal-report-full-go-writesummary.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`writeUnitDetails`](reports/internal-report-full-go-writeunitdetails.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`full_test.go`](reports/internal-report-full-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`FormatJSON`](reports/internal-report-health-go-formatjson.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`FormatText`](reports/internal-report-health-go-formattext.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`Health`](reports/internal-report-health-go-health.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`HealthReport`](reports/internal-report-health-go-healthreport.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`health_test.go`](reports/internal-report-health-test-go.md) | file | B+ | 88.3% | certified | 2026-04-23 |
| [`FormatUnitMarkdown`](reports/internal-report-unit-report-go-formatunitmarkdown.md) | function | B+ | 87.2% | certified | 2026-04-23 |
| [`GenerateUnitReports`](reports/internal-report-unit-report-go-generateunitreports.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`formatDate`](reports/internal-report-unit-report-go-formatdate.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`scoreBar`](reports/internal-report-unit-report-go-scorebar.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`splitObservations`](reports/internal-report-unit-report-go-splitobservations.md) | function | B+ | 89.4% | certified | 2026-04-23 |

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
- 🤖 The code is functionally correct but has several maintainability and correctness issues including hardcoded values, lack of input validation, and missing error handling in JSON marshaling.
- 💡 Add input validation in `badgeMessage` to handle edge cases like negative pass rates or invalid total units
- 💡 Make `badgeColor` configurable or use a lookup table from an external configuration file to allow branding updates without code changes
- 💡 Add error checking in `FormatBadgeJSON` to propagate JSON marshaling errors instead of silently returning nil
- 💡 Parameterize `Label` and `NamedLogo` in `GenerateBadge` to support reuse across different certification systems
- 💡 Validate and sanitize inputs in `BadgeMarkdown` to prevent injection vulnerabilities or malformed URLs
- ⚠️ Hardcoded color values in `badgeColor` may not align with updated brand guidelines or cause visual inconsistency
- ⚠️ Silent failure in `FormatBadgeJSON` if JSON marshaling fails due to invalid data
- 🔗 The `GenerateBadge` function tightly couples the badge label and logo to specific values, increasing coupling with internal branding
- 🔗 The `BadgeMarkdown` function assumes a fixed GitHub URL structure, which increases fragility in environments with different repo hosting or path structures

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
- 🤖 The code is functionally correct but has several maintainability and correctness issues including hardcoded values, lack of input validation, and missing error handling in JSON marshaling.
- 💡 Add input validation in `badgeMessage` to handle edge cases like negative pass rates or invalid total units
- 💡 Make `badgeColor` configurable or use a lookup table from an external configuration file to allow branding updates without code changes
- 💡 Add error checking in `FormatBadgeJSON` to propagate JSON marshaling errors instead of silently returning nil
- 💡 Parameterize `Label` and `NamedLogo` in `GenerateBadge` to support reuse across different certification systems
- 💡 Validate and sanitize inputs in `BadgeMarkdown` to prevent injection vulnerabilities or malformed URLs
- ⚠️ Hardcoded color values in `badgeColor` may not align with updated brand guidelines or cause visual inconsistency
- ⚠️ Silent failure in `FormatBadgeJSON` if JSON marshaling fails due to invalid data
- 🔗 The `GenerateBadge` function tightly couples the badge label and logo to specific values, increasing coupling with internal branding
- 🔗 The `BadgeMarkdown` function assumes a fixed GitHub URL structure, which increases fragility in environments with different repo hosting or path structures

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
- 🤖 The code is functionally correct but has several maintainability and correctness issues including hardcoded values, lack of input validation, and missing error handling in JSON marshaling.
- 💡 Add input validation in `badgeMessage` to handle edge cases like negative pass rates or invalid total units
- 💡 Make `badgeColor` configurable or use a lookup table from an external configuration file to allow branding updates without code changes
- 💡 Add error checking in `FormatBadgeJSON` to propagate JSON marshaling errors instead of silently returning nil
- 💡 Parameterize `Label` and `NamedLogo` in `GenerateBadge` to support reuse across different certification systems
- 💡 Validate and sanitize inputs in `BadgeMarkdown` to prevent injection vulnerabilities or malformed URLs
- ⚠️ Hardcoded color values in `badgeColor` may not align with updated brand guidelines or cause visual inconsistency
- ⚠️ Silent failure in `FormatBadgeJSON` if JSON marshaling fails due to invalid data
- 🔗 The `GenerateBadge` function tightly couples the badge label and logo to specific values, increasing coupling with internal branding
- 🔗 The `BadgeMarkdown` function assumes a fixed GitHub URL structure, which increases fragility in environments with different repo hosting or path structures

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
- 🤖 The code is functionally correct but has several maintainability and correctness issues including hardcoded values, lack of input validation, and missing error handling in JSON marshaling.
- 💡 Add input validation in `badgeMessage` to handle edge cases like negative pass rates or invalid total units
- 💡 Make `badgeColor` configurable or use a lookup table from an external configuration file to allow branding updates without code changes
- 💡 Add error checking in `FormatBadgeJSON` to propagate JSON marshaling errors instead of silently returning nil
- 💡 Parameterize `Label` and `NamedLogo` in `GenerateBadge` to support reuse across different certification systems
- 💡 Validate and sanitize inputs in `BadgeMarkdown` to prevent injection vulnerabilities or malformed URLs
- ⚠️ Hardcoded color values in `badgeColor` may not align with updated brand guidelines or cause visual inconsistency
- ⚠️ Silent failure in `FormatBadgeJSON` if JSON marshaling fails due to invalid data
- 🔗 The `GenerateBadge` function tightly couples the badge label and logo to specific values, increasing coupling with internal branding
- 🔗 The `BadgeMarkdown` function assumes a fixed GitHub URL structure, which increases fragility in environments with different repo hosting or path structures

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
- 🤖 The code is functionally correct but has several maintainability and correctness issues including hardcoded values, lack of input validation, and missing error handling in JSON marshaling.
- 💡 Add input validation in `badgeMessage` to handle edge cases like negative pass rates or invalid total units
- 💡 Make `badgeColor` configurable or use a lookup table from an external configuration file to allow branding updates without code changes
- 💡 Add error checking in `FormatBadgeJSON` to propagate JSON marshaling errors instead of silently returning nil
- 💡 Parameterize `Label` and `NamedLogo` in `GenerateBadge` to support reuse across different certification systems
- 💡 Validate and sanitize inputs in `BadgeMarkdown` to prevent injection vulnerabilities or malformed URLs
- ⚠️ Hardcoded color values in `badgeColor` may not align with updated brand guidelines or cause visual inconsistency
- ⚠️ Silent failure in `FormatBadgeJSON` if JSON marshaling fails due to invalid data
- 🔗 The `GenerateBadge` function tightly couples the badge label and logo to specific values, increasing coupling with internal branding
- 🔗 The `BadgeMarkdown` function assumes a fixed GitHub URL structure, which increases fragility in environments with different repo hosting or path structures

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
- 🤖 The code is functionally correct but has several maintainability and correctness issues including hardcoded values, lack of input validation, and missing error handling in JSON marshaling.
- 💡 Add input validation in `badgeMessage` to handle edge cases like negative pass rates or invalid total units
- 💡 Make `badgeColor` configurable or use a lookup table from an external configuration file to allow branding updates without code changes
- 💡 Add error checking in `FormatBadgeJSON` to propagate JSON marshaling errors instead of silently returning nil
- 💡 Parameterize `Label` and `NamedLogo` in `GenerateBadge` to support reuse across different certification systems
- 💡 Validate and sanitize inputs in `BadgeMarkdown` to prevent injection vulnerabilities or malformed URLs
- ⚠️ Hardcoded color values in `badgeColor` may not align with updated brand guidelines or cause visual inconsistency
- ⚠️ Silent failure in `FormatBadgeJSON` if JSON marshaling fails due to invalid data
- 🔗 The `GenerateBadge` function tightly couples the badge label and logo to specific values, increasing coupling with internal branding
- 🔗 The `BadgeMarkdown` function assumes a fixed GitHub URL structure, which increases fragility in environments with different repo hosting or path structures

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
- 🤖 The test suite for badge generation is functional but lacks comprehensive coverage for edge cases and error conditions.
- 💡 Add tests for GenerateCard with nil records, empty records slice, and mixed status records to ensure robustness in grade calculation
- 💡 Add comprehensive tests for FormatBadgeJSON that validate all fields are correctly serialized and handle edge cases like nil or zero-value badge fields
- 💡 Add a test case for BadgeMarkdown that validates the constructed URL format and ensures proper encoding of repository names and branches
- 💡 Include tests for edge cases such as very high or low scores (e.g., >1.0 or <0.0) to ensure color mapping logic is resilient
- ⚠️ Potential incorrect badge color assignment due to insufficient test coverage for mixed record types and score combinations
- ⚠️ Risk of malformed JSON output from FormatBadgeJSON if badge struct contains unexpected field types or nil values
- 🔗 This test suite only validates the badge generation logic in isolation, not how it integrates with the broader reporting pipeline or external dependencies like shields.io
- 🔗 The tests do not simulate failure propagation from badge generation to downstream consumers, potentially masking issues in error handling or integration points

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
- 🤖 Well-structured report card generation with minor correctness and maintainability concerns.
- 💡 Fix the observation counting logic in `GenerateCard` to avoid double-counting units with `StatusCertifiedWithObservations`
- 💡 Reorder logic in `buildTopIssues` to prioritize failing units before low-scoring passing units
- 💡 Replace hardcoded string slicing (`c.GeneratedAt[:19]`) in `FormatCardText` with a robust time formatting function
- 💡 Initialize all possible grades (A, A-, B+, B, C, D, F) in `GradeDistribution` map to ensure all grades appear in the output
- 💡 Escape special characters in unit IDs when rendering markdown to prevent malformed markdown
- ⚠️ Incorrect observation counting logic in `GenerateCard` where units with `StatusCertifiedWithObservations` are counted as both passing and observations
- ⚠️ Top issues prioritization in `buildTopIssues` does not correctly prioritize failing units over low-scoring passing units
- 🔗 The `GenerateCard` function affects downstream consumers by potentially misreporting pass rates and observation counts
- 🔗 The `buildTopIssues` function affects the UI or reporting logic by not correctly prioritizing failing units, which may mislead users

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
- 🤖 Well-structured report card generation with minor correctness and maintainability concerns.
- 💡 Fix the observation counting logic in `GenerateCard` to avoid double-counting units with `StatusCertifiedWithObservations`
- 💡 Reorder logic in `buildTopIssues` to prioritize failing units before low-scoring passing units
- 💡 Replace hardcoded string slicing (`c.GeneratedAt[:19]`) in `FormatCardText` with a robust time formatting function
- 💡 Initialize all possible grades (A, A-, B+, B, C, D, F) in `GradeDistribution` map to ensure all grades appear in the output
- 💡 Escape special characters in unit IDs when rendering markdown to prevent malformed markdown
- ⚠️ Incorrect observation counting logic in `GenerateCard` where units with `StatusCertifiedWithObservations` are counted as both passing and observations
- ⚠️ Top issues prioritization in `buildTopIssues` does not correctly prioritize failing units over low-scoring passing units
- 🔗 The `GenerateCard` function affects downstream consumers by potentially misreporting pass rates and observation counts
- 🔗 The `buildTopIssues` function affects the UI or reporting logic by not correctly prioritizing failing units, which may mislead users

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
- 🤖 Well-structured report card generation with minor correctness and maintainability concerns.
- 💡 Fix the observation counting logic in `GenerateCard` to avoid double-counting units with `StatusCertifiedWithObservations`
- 💡 Reorder logic in `buildTopIssues` to prioritize failing units before low-scoring passing units
- 💡 Replace hardcoded string slicing (`c.GeneratedAt[:19]`) in `FormatCardText` with a robust time formatting function
- 💡 Initialize all possible grades (A, A-, B+, B, C, D, F) in `GradeDistribution` map to ensure all grades appear in the output
- 💡 Escape special characters in unit IDs when rendering markdown to prevent malformed markdown
- ⚠️ Incorrect observation counting logic in `GenerateCard` where units with `StatusCertifiedWithObservations` are counted as both passing and observations
- ⚠️ Top issues prioritization in `buildTopIssues` does not correctly prioritize failing units over low-scoring passing units
- 🔗 The `GenerateCard` function affects downstream consumers by potentially misreporting pass rates and observation counts
- 🔗 The `buildTopIssues` function affects the UI or reporting logic by not correctly prioritizing failing units, which may mislead users

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
- 🤖 Well-structured report card generation with minor correctness and maintainability concerns.
- 💡 Fix the observation counting logic in `GenerateCard` to avoid double-counting units with `StatusCertifiedWithObservations`
- 💡 Reorder logic in `buildTopIssues` to prioritize failing units before low-scoring passing units
- 💡 Replace hardcoded string slicing (`c.GeneratedAt[:19]`) in `FormatCardText` with a robust time formatting function
- 💡 Initialize all possible grades (A, A-, B+, B, C, D, F) in `GradeDistribution` map to ensure all grades appear in the output
- 💡 Escape special characters in unit IDs when rendering markdown to prevent malformed markdown
- ⚠️ Incorrect observation counting logic in `GenerateCard` where units with `StatusCertifiedWithObservations` are counted as both passing and observations
- ⚠️ Top issues prioritization in `buildTopIssues` does not correctly prioritize failing units over low-scoring passing units
- 🔗 The `GenerateCard` function affects downstream consumers by potentially misreporting pass rates and observation counts
- 🔗 The `buildTopIssues` function affects the UI or reporting logic by not correctly prioritizing failing units, which may mislead users

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
- 🤖 Well-structured report card generation with minor correctness and maintainability concerns.
- 💡 Fix the observation counting logic in `GenerateCard` to avoid double-counting units with `StatusCertifiedWithObservations`
- 💡 Reorder logic in `buildTopIssues` to prioritize failing units before low-scoring passing units
- 💡 Replace hardcoded string slicing (`c.GeneratedAt[:19]`) in `FormatCardText` with a robust time formatting function
- 💡 Initialize all possible grades (A, A-, B+, B, C, D, F) in `GradeDistribution` map to ensure all grades appear in the output
- 💡 Escape special characters in unit IDs when rendering markdown to prevent malformed markdown
- ⚠️ Incorrect observation counting logic in `GenerateCard` where units with `StatusCertifiedWithObservations` are counted as both passing and observations
- ⚠️ Top issues prioritization in `buildTopIssues` does not correctly prioritize failing units over low-scoring passing units
- 🔗 The `GenerateCard` function affects downstream consumers by potentially misreporting pass rates and observation counts
- 🔗 The `buildTopIssues` function affects the UI or reporting logic by not correctly prioritizing failing units, which may mislead users

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
- 🤖 Well-structured report card generation with minor correctness and maintainability concerns.
- 💡 Fix the observation counting logic in `GenerateCard` to avoid double-counting units with `StatusCertifiedWithObservations`
- 💡 Reorder logic in `buildTopIssues` to prioritize failing units before low-scoring passing units
- 💡 Replace hardcoded string slicing (`c.GeneratedAt[:19]`) in `FormatCardText` with a robust time formatting function
- 💡 Initialize all possible grades (A, A-, B+, B, C, D, F) in `GradeDistribution` map to ensure all grades appear in the output
- 💡 Escape special characters in unit IDs when rendering markdown to prevent malformed markdown
- ⚠️ Incorrect observation counting logic in `GenerateCard` where units with `StatusCertifiedWithObservations` are counted as both passing and observations
- ⚠️ Top issues prioritization in `buildTopIssues` does not correctly prioritize failing units over low-scoring passing units
- 🔗 The `GenerateCard` function affects downstream consumers by potentially misreporting pass rates and observation counts
- 🔗 The `buildTopIssues` function affects the UI or reporting logic by not correctly prioritizing failing units, which may mislead users

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
- 🤖 Well-structured report card generation with minor correctness and maintainability concerns.
- 💡 Fix the observation counting logic in `GenerateCard` to avoid double-counting units with `StatusCertifiedWithObservations`
- 💡 Reorder logic in `buildTopIssues` to prioritize failing units before low-scoring passing units
- 💡 Replace hardcoded string slicing (`c.GeneratedAt[:19]`) in `FormatCardText` with a robust time formatting function
- 💡 Initialize all possible grades (A, A-, B+, B, C, D, F) in `GradeDistribution` map to ensure all grades appear in the output
- 💡 Escape special characters in unit IDs when rendering markdown to prevent malformed markdown
- ⚠️ Incorrect observation counting logic in `GenerateCard` where units with `StatusCertifiedWithObservations` are counted as both passing and observations
- ⚠️ Top issues prioritization in `buildTopIssues` does not correctly prioritize failing units over low-scoring passing units
- 🔗 The `GenerateCard` function affects downstream consumers by potentially misreporting pass rates and observation counts
- 🔗 The `buildTopIssues` function affects the UI or reporting logic by not correctly prioritizing failing units, which may mislead users

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
- 🤖 Well-structured report card generation with minor correctness and maintainability concerns.
- 💡 Fix the observation counting logic in `GenerateCard` to avoid double-counting units with `StatusCertifiedWithObservations`
- 💡 Reorder logic in `buildTopIssues` to prioritize failing units before low-scoring passing units
- 💡 Replace hardcoded string slicing (`c.GeneratedAt[:19]`) in `FormatCardText` with a robust time formatting function
- 💡 Initialize all possible grades (A, A-, B+, B, C, D, F) in `GradeDistribution` map to ensure all grades appear in the output
- 💡 Escape special characters in unit IDs when rendering markdown to prevent malformed markdown
- ⚠️ Incorrect observation counting logic in `GenerateCard` where units with `StatusCertifiedWithObservations` are counted as both passing and observations
- ⚠️ Top issues prioritization in `buildTopIssues` does not correctly prioritize failing units over low-scoring passing units
- 🔗 The `GenerateCard` function affects downstream consumers by potentially misreporting pass rates and observation counts
- 🔗 The `buildTopIssues` function affects the UI or reporting logic by not correctly prioritizing failing units, which may mislead users

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
- 🤖 Well-structured report card generation with minor correctness and maintainability concerns.
- 💡 Fix the observation counting logic in `GenerateCard` to avoid double-counting units with `StatusCertifiedWithObservations`
- 💡 Reorder logic in `buildTopIssues` to prioritize failing units before low-scoring passing units
- 💡 Replace hardcoded string slicing (`c.GeneratedAt[:19]`) in `FormatCardText` with a robust time formatting function
- 💡 Initialize all possible grades (A, A-, B+, B, C, D, F) in `GradeDistribution` map to ensure all grades appear in the output
- 💡 Escape special characters in unit IDs when rendering markdown to prevent malformed markdown
- ⚠️ Incorrect observation counting logic in `GenerateCard` where units with `StatusCertifiedWithObservations` are counted as both passing and observations
- ⚠️ Top issues prioritization in `buildTopIssues` does not correctly prioritize failing units over low-scoring passing units
- 🔗 The `GenerateCard` function affects downstream consumers by potentially misreporting pass rates and observation counts
- 🔗 The `buildTopIssues` function affects the UI or reporting logic by not correctly prioritizing failing units, which may mislead users

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
- 🤖 The test suite for GenerateCard and related formatting functions is well-structured but lacks comprehensive edge case coverage and has some brittle assertions.
- 💡 Replace string-based containment checks in `TestFormatCardText` and `TestFormatCardMarkdown` with assertions that validate structured output (e.g., use regex to match expected table formats or validate JSON-like structures)
- 💡 Add tests for `GenerateCard` edge cases such as records with nil observations, duplicate unit IDs, and invalid status values to ensure robustness
- 💡 Add assertions in `TestGenerateCard_WithFailures` to verify that `TopIssues` contains the correct number of entries and that each entry includes the expected fields from failing records
- 💡 Modify `makeCardRecord` to accept a unit type parameter so that different unit types can be tested without hardcoding `domain.UnitTypeFunction`
- ⚠️ Fragile string-based assertions in `TestFormatCardText` and `TestFormatCardMarkdown` can break with minor formatting changes or whitespace differences
- ⚠️ Lack of validation for `TopIssues` population logic in `TestGenerateCard_WithFailures` may hide bugs in issue tracking
- 🔗 These tests provide minimal coverage for edge cases in `GenerateCard` and do not adequately test failure propagation or error handling paths
- 🔗 The string-based format tests increase coupling between the report generation logic and specific text formatting patterns, making future refactoring more difficult

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 Well-structured report generation with solid logic but missing error handling and potential performance bottlenecks in sorting.
- 💡 Remove duplicate comment on line 63-64 in `Detailed` function
- 💡 Add validation or error handling for `dim.String()` and `ExpiresAt.Format(time.RFC3339)` to prevent runtime panics
- 💡 Refactor `findHighestRisk` and `findRecurrentlyFailing` to avoid copying slices before sorting; use indices instead
- 💡 Replace hardcoded date substring `u.ExpiresAt[:10]` with proper time formatting using a consistent format like `time.RFC3339` or `time.DateOnly`
- 💡 Validate that `r.UnitID.String()` and `r.Status.String()` are non-empty before including in report output to prevent malformed entries
- ⚠️ Potential panic on `dim.String()` call if `dim` is not a valid type or nil
- ⚠️ Runtime panic on `ExpiresAt.Format(time.RFC3339)` if time is zero or invalid
- 🔗 The function tightly couples to `domain.CertificationRecord` and assumes specific field behaviors (e.g., `IsPassing()`, `String()`), increasing system coupling
- 🔗 The sorting logic in `findHighestRisk` and `findRecurrentlyFailing` creates unnecessary copies of slices, impacting performance for large datasets

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
- 🤖 The test suite for DetailedReport is well-structured but has several correctness and maintainability issues including hardcoded expectations, fragile string matching, and missing edge case handling.
- 💡 Replace hardcoded tolerance checks like `v < 0.79 || v > 0.81` with `math.Abs(v - expected) < epsilon` comparisons
- 💡 Replace `strings.Contains(d.ExpiringSoon[0].UnitID, "expiring.go")` with an assertion that verifies the slice is sorted by `ExpiresAt` in ascending order
- 💡 Use `reflect.DeepEqual` or structured comparison instead of string matching for validating formatted output in `TestFormatDetailedText`
- 💡 Add test cases for edge conditions such as zero scores, nil dimensions, and records with identical timestamps
- 💡 Use a fixed time in `makeDetailedRecord` to ensure deterministic behavior in time-sensitive tests
- ⚠️ Hardcoded tolerance ranges in dimension average tests can cause false negatives when implementation details change
- ⚠️ String-based assertions in expiring soon test do not validate proper sorting logic and are fragile
- 🔗 These tests do not directly affect system behavior but can mask bugs in the report generation logic if they pass despite incorrect implementation
- 🔗 Fragile string matching in `TestFormatDetailedText` increases risk of missing regressions in text formatting

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

</details>

<a id="internal-report-full-go-unitanchor"></a>
<details>
<summary>unitAnchor — certified details</summary>

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 Well-structured report generation with solid logic but several correctness and maintainability issues in string handling, edge case management, and performance.
- 💡 Replace `r.GeneratedAt[:19]` in writeSummary with a proper time formatting function or use time.RFC3339 directly
- 💡 Initialize `bottom` in buildLanguageDetail to the first score value or use math.Inf(1) to ensure correct min calculation
- 💡 Validate and sanitize AI observation prefixes in writeAIInsights to prevent malformed markdown or injection
- 💡 Fix unitAnchor to handle all special characters that can break HTML IDs (e.g., '+', '#')
- 💡 Add error handling for time formatting in unitReportFrom to prevent panics on invalid timestamps
- 💡 Optimize sortedKeys to reuse or pre-sort keys rather than allocating new slices each time
- ⚠️ Hardcoded time truncation in writeSummary can break with future time formats
- ⚠️ Incorrect bottom score calculation in buildLanguageDetail when scores < 1.0
- 🔗 The writeSummary and buildLanguageDetail functions directly affect report accuracy and can propagate incorrect data to downstream consumers
- 🔗 The unitAnchor function's flawed sanitization introduces potential HTML ID conflicts that may break UI rendering in the generated markdown

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
- 🤖 The test suite for full report generation is well-structured but has a critical logic error in test data setup and lacks proper validation of generated report content.
- 💡 Fix the test data setup in `TestGenerateFullReport` by ensuring that `records[3]` is properly added to the slice or removed from the observation assignment
- 💡 Add comprehensive assertions in `TestFormatFullMarkdown` to verify that markdown output contains specific formatted values like scores, unit names, and properly structured tables instead of just string presence checks
- 💡 Ensure that all generated report fields (e.g., `CertifiedAt`, `ExpiresAt`) are validated in tests to confirm they match expected values or formats
- ⚠️ Runtime panic due to out-of-bounds slice access in TestGenerateFullReport
- ⚠️ Insufficient validation of report content and formatting in TestFormatFullMarkdown
- 🔗 This test file directly tests the report generation logic and can mask bugs in the actual `GenerateFullReport` function if tests are not fixed
- 🔗 The test data corruption affects the entire test suite's ability to validate correctness of report generation

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
- 🤖 The code is functionally correct but has a critical division-by-zero bug and lacks proper error handling for edge cases.
- 💡 Add explicit zero-check for `len(records)` before computing `PassRate` and `AverageScore` to prevent division-by-zero errors
- 💡 Validate that `r.Score` is non-negative and finite before including it in totalScore calculations to avoid incorrect averages
- 💡 Add unit tests for edge cases such as empty input, zero scores, and NaN scores
- 💡 Consider returning a more descriptive error or handling edge cases explicitly in the `Health` function instead of silently returning an empty struct
- ⚠️ Division-by-zero in `PassRate` and `AverageScore` when records is empty
- ⚠️ Potential invalid score handling if `r.Score` is NaN or negative
- 🔗 The `HealthReport` struct and its methods form a tightly coupled component in the certification reporting pipeline, affecting downstream consumers that depend on accurate metrics
- 🔗 If `Health` is called with an empty slice and the returned struct is used without checking for zero values, downstream logic may misinterpret results

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
- 🤖 The code is functionally correct but has a critical division-by-zero bug and lacks proper error handling for edge cases.
- 💡 Add explicit zero-check for `len(records)` before computing `PassRate` and `AverageScore` to prevent division-by-zero errors
- 💡 Validate that `r.Score` is non-negative and finite before including it in totalScore calculations to avoid incorrect averages
- 💡 Add unit tests for edge cases such as empty input, zero scores, and NaN scores
- 💡 Consider returning a more descriptive error or handling edge cases explicitly in the `Health` function instead of silently returning an empty struct
- ⚠️ Division-by-zero in `PassRate` and `AverageScore` when records is empty
- ⚠️ Potential invalid score handling if `r.Score` is NaN or negative
- 🔗 The `HealthReport` struct and its methods form a tightly coupled component in the certification reporting pipeline, affecting downstream consumers that depend on accurate metrics
- 🔗 If `Health` is called with an empty slice and the returned struct is used without checking for zero values, downstream logic may misinterpret results

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
- 🤖 The code is functionally correct but has a critical division-by-zero bug and lacks proper error handling for edge cases.
- 💡 Add explicit zero-check for `len(records)` before computing `PassRate` and `AverageScore` to prevent division-by-zero errors
- 💡 Validate that `r.Score` is non-negative and finite before including it in totalScore calculations to avoid incorrect averages
- 💡 Add unit tests for edge cases such as empty input, zero scores, and NaN scores
- 💡 Consider returning a more descriptive error or handling edge cases explicitly in the `Health` function instead of silently returning an empty struct
- ⚠️ Division-by-zero in `PassRate` and `AverageScore` when records is empty
- ⚠️ Potential invalid score handling if `r.Score` is NaN or negative
- 🔗 The `HealthReport` struct and its methods form a tightly coupled component in the certification reporting pipeline, affecting downstream consumers that depend on accurate metrics
- 🔗 If `Health` is called with an empty slice and the returned struct is used without checking for zero values, downstream logic may misinterpret results

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
- 🤖 The code is functionally correct but has a critical division-by-zero bug and lacks proper error handling for edge cases.
- 💡 Add explicit zero-check for `len(records)` before computing `PassRate` and `AverageScore` to prevent division-by-zero errors
- 💡 Validate that `r.Score` is non-negative and finite before including it in totalScore calculations to avoid incorrect averages
- 💡 Add unit tests for edge cases such as empty input, zero scores, and NaN scores
- 💡 Consider returning a more descriptive error or handling edge cases explicitly in the `Health` function instead of silently returning an empty struct
- ⚠️ Division-by-zero in `PassRate` and `AverageScore` when records is empty
- ⚠️ Potential invalid score handling if `r.Score` is NaN or negative
- 🔗 The `HealthReport` struct and its methods form a tightly coupled component in the certification reporting pipeline, affecting downstream consumers that depend on accurate metrics
- 🔗 If `Health` is called with an empty slice and the returned struct is used without checking for zero values, downstream logic may misinterpret results

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
- 🤖 The test suite for health report logic is well-structured but lacks comprehensive edge case coverage and includes a potential logic error in pass rate calculation.
- 💡 Fix TestHealthReport_Summary to correctly calculate pass rate as (Certified + CertifiedWithObs) / TotalUnits
- 💡 Replace hardcoded time durations in makeRecord with configurable or mockable times to improve test stability
- ⚠️ Incorrect pass rate calculation logic in TestHealthReport_Summary
- ⚠️ Flaky tests due to hardcoded time values and reliance on system clock
- 🔗 This unit tests health report logic which is used in reporting and monitoring systems; incorrect pass rate calculation can propagate false insights
- 🔗 The test suite does not fully exercise boundary conditions or error paths in the health report generation logic

</details>

<a id="internal-report-unit-report-go-formatunitmarkdown"></a>
<details>
<summary>FormatUnitMarkdown — certified details</summary>

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
- 🤖 The code correctly generates unit reports but has critical security and correctness issues related to path traversal, file overwrites, and improper error handling.
- 💡 Validate and sanitize `unitAnchor(u)` to prevent path traversal by ensuring it only contains alphanumeric characters, hyphens, and underscores
- 💡 Use atomic file writes (e.g., write to temp file then rename) to prevent partial or corrupted files
- 💡 Change error return logic to return 0 when any write fails, and log the specific file that caused failure
- 💡 Add input validation for `u.Symbol` and `u.Path` to ensure they are safe for use in filenames
- 💡 Implement proper file permission checks and enforce stricter permissions if needed
- ⚠️ Path traversal vulnerability due to unsafe filename generation from unitAnchor()
- ⚠️ Misleading error handling where partial writes are reported as full success
- 🔗 This unit can be exploited to write arbitrary files outside the intended output directory, leading to potential data corruption or system compromise
- 🔗 Improper error handling causes partial success reporting that masks real failures in report generation

</details>

<a id="internal-report-unit-report-go-generateunitreports"></a>
<details>
<summary>GenerateUnitReports — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The code correctly generates unit reports but has critical security and correctness issues related to path traversal, file overwrites, and improper error handling.
- 💡 Validate and sanitize `unitAnchor(u)` to prevent path traversal by ensuring it only contains alphanumeric characters, hyphens, and underscores
- 💡 Use atomic file writes (e.g., write to temp file then rename) to prevent partial or corrupted files
- 💡 Change error return logic to return 0 when any write fails, and log the specific file that caused failure
- 💡 Add input validation for `u.Symbol` and `u.Path` to ensure they are safe for use in filenames
- 💡 Implement proper file permission checks and enforce stricter permissions if needed
- ⚠️ Path traversal vulnerability due to unsafe filename generation from unitAnchor()
- ⚠️ Misleading error handling where partial writes are reported as full success
- 🔗 This unit can be exploited to write arbitrary files outside the intended output directory, leading to potential data corruption or system compromise
- 🔗 Improper error handling causes partial success reporting that masks real failures in report generation

</details>

<a id="internal-report-unit-report-go-formatdate"></a>
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
- 🤖 The code correctly generates unit reports but has critical security and correctness issues related to path traversal, file overwrites, and improper error handling.
- 💡 Validate and sanitize `unitAnchor(u)` to prevent path traversal by ensuring it only contains alphanumeric characters, hyphens, and underscores
- 💡 Use atomic file writes (e.g., write to temp file then rename) to prevent partial or corrupted files
- 💡 Change error return logic to return 0 when any write fails, and log the specific file that caused failure
- 💡 Add input validation for `u.Symbol` and `u.Path` to ensure they are safe for use in filenames
- 💡 Implement proper file permission checks and enforce stricter permissions if needed
- ⚠️ Path traversal vulnerability due to unsafe filename generation from unitAnchor()
- ⚠️ Misleading error handling where partial writes are reported as full success
- 🔗 This unit can be exploited to write arbitrary files outside the intended output directory, leading to potential data corruption or system compromise
- 🔗 Improper error handling causes partial success reporting that masks real failures in report generation

</details>

<a id="internal-report-unit-report-go-scorebar"></a>
<details>
<summary>scoreBar — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The code correctly generates unit reports but has critical security and correctness issues related to path traversal, file overwrites, and improper error handling.
- 💡 Validate and sanitize `unitAnchor(u)` to prevent path traversal by ensuring it only contains alphanumeric characters, hyphens, and underscores
- 💡 Use atomic file writes (e.g., write to temp file then rename) to prevent partial or corrupted files
- 💡 Change error return logic to return 0 when any write fails, and log the specific file that caused failure
- 💡 Add input validation for `u.Symbol` and `u.Path` to ensure they are safe for use in filenames
- 💡 Implement proper file permission checks and enforce stricter permissions if needed
- ⚠️ Path traversal vulnerability due to unsafe filename generation from unitAnchor()
- ⚠️ Misleading error handling where partial writes are reported as full success
- 🔗 This unit can be exploited to write arbitrary files outside the intended output directory, leading to potential data corruption or system compromise
- 🔗 Improper error handling causes partial success reporting that masks real failures in report generation

</details>

<a id="internal-report-unit-report-go-splitobservations"></a>
<details>
<summary>splitObservations — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 85.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 85.0% |
| readability | 95.0% |
| security | 85.0% |
| testability | 90.0% |

**Observations:**
- 🤖 The code correctly generates unit reports but has critical security and correctness issues related to path traversal, file overwrites, and improper error handling.
- 💡 Validate and sanitize `unitAnchor(u)` to prevent path traversal by ensuring it only contains alphanumeric characters, hyphens, and underscores
- 💡 Use atomic file writes (e.g., write to temp file then rename) to prevent partial or corrupted files
- 💡 Change error return logic to return 0 when any write fails, and log the specific file that caused failure
- 💡 Add input validation for `u.Symbol` and `u.Path` to ensure they are safe for use in filenames
- 💡 Implement proper file permission checks and enforce stricter permissions if needed
- ⚠️ Path traversal vulnerability due to unsafe filename generation from unitAnchor()
- ⚠️ Misleading error handling where partial writes are reported as full success
- 🔗 This unit can be exploited to write arbitrary files outside the intended output directory, leading to potential data corruption or system compromise
- 🔗 Improper error handling causes partial success reporting that masks real failures in report generation

</details>

### `testdata/repos/ts-simple/src/` (6 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`DialogueParser`](reports/testdata-repos-ts-simple-src-parser-ts-dialogueparser.md) | class | B+ | 89.4% | certified | 2026-04-23 |
| [`MAX_TOKENS`](reports/testdata-repos-ts-simple-src-parser-ts-max-tokens.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`parseNode`](reports/testdata-repos-ts-simple-src-parser-ts-parsenode.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`tokenizeDialogue`](reports/testdata-repos-ts-simple-src-parser-ts-tokenizedialogue.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`formatDate`](reports/testdata-repos-ts-simple-src-utils-ts-formatdate.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`log`](reports/testdata-repos-ts-simple-src-utils-ts-log.md) | function | B+ | 89.4% | certified | 2026-04-23 |

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
- 🤖 Code is incomplete and has critical functional gaps - parseNode is unused and DialogueParser doesn't actually parse nodes
- 💡 Implement actual parsing logic in DialogueParser.parse() method to process tokens into structured nodes
- 💡 Add input validation and token limit checking using the MAX_TOKENS constant to prevent resource exhaustion
- 💡 Remove or properly implement parseNode function to either process tokens or remove it entirely
- ⚠️ Missing parsing implementation - the DialogueParser class has no actual parsing logic despite claiming to parse
- ⚠️ Unvalidated token processing - tokens are never validated or processed beyond simple splitting
- 🔗 This unit creates false expectations of parsing functionality - clients will assume it can parse dialogue but it cannot
- 🔗 The class structure suggests a parsing pipeline that doesn't exist, creating coupling to non-existent interfaces

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
- 🤖 Code is incomplete and has critical functional gaps - parseNode is unused and DialogueParser doesn't actually parse nodes
- 💡 Implement actual parsing logic in DialogueParser.parse() method to process tokens into structured nodes
- 💡 Add input validation and token limit checking using the MAX_TOKENS constant to prevent resource exhaustion
- 💡 Remove or properly implement parseNode function to either process tokens or remove it entirely
- ⚠️ Missing parsing implementation - the DialogueParser class has no actual parsing logic despite claiming to parse
- ⚠️ Unvalidated token processing - tokens are never validated or processed beyond simple splitting
- 🔗 This unit creates false expectations of parsing functionality - clients will assume it can parse dialogue but it cannot
- 🔗 The class structure suggests a parsing pipeline that doesn't exist, creating coupling to non-existent interfaces

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
- 🤖 Code is incomplete and has critical functional gaps - parseNode is unused and DialogueParser doesn't actually parse nodes
- 💡 Implement actual parsing logic in DialogueParser.parse() method to process tokens into structured nodes
- 💡 Add input validation and token limit checking using the MAX_TOKENS constant to prevent resource exhaustion
- 💡 Remove or properly implement parseNode function to either process tokens or remove it entirely
- ⚠️ Missing parsing implementation - the DialogueParser class has no actual parsing logic despite claiming to parse
- ⚠️ Unvalidated token processing - tokens are never validated or processed beyond simple splitting
- 🔗 This unit creates false expectations of parsing functionality - clients will assume it can parse dialogue but it cannot
- 🔗 The class structure suggests a parsing pipeline that doesn't exist, creating coupling to non-existent interfaces

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
- 🤖 Code is incomplete and has critical functional gaps - parseNode is unused and DialogueParser doesn't actually parse nodes
- 💡 Implement actual parsing logic in DialogueParser.parse() method to process tokens into structured nodes
- 💡 Add input validation and token limit checking using the MAX_TOKENS constant to prevent resource exhaustion
- 💡 Remove or properly implement parseNode function to either process tokens or remove it entirely
- ⚠️ Missing parsing implementation - the DialogueParser class has no actual parsing logic despite claiming to parse
- ⚠️ Unvalidated token processing - tokens are never validated or processed beyond simple splitting
- 🔗 This unit creates false expectations of parsing functionality - clients will assume it can parse dialogue but it cannot
- 🔗 The class structure suggests a parsing pipeline that doesn't exist, creating coupling to non-existent interfaces

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
- 🤖 Simple utility functions with minimal functionality but significant architectural and operational concerns
- 💡 Rename formatDate to serializeDate or toISOString to accurately reflect its behavior, or implement proper date formatting logic
- 💡 Replace direct console.log with configurable logger interface that can be mocked or swapped for different logging backends
- 💡 Add input validation to formatDate to handle invalid Date objects gracefully
- 💡 Consider adding timezone support or locale-aware formatting options for the formatDate function
- ⚠️ Misleading function name that doesn't match actual behavior - formatDate returns ISO string instead of formatted date
- ⚠️ Direct console.log usage makes logging non-configurable and hard to test/migrate in different environments
- 🔗 The logging function creates tight coupling between the utility module and console output, making it impossible to switch to structured logging or different output destinations
- 🔗 The formatDate function creates false expectations about date formatting capabilities and may break if consumers expect different output formats

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
- 🤖 Simple utility functions with minimal functionality but significant architectural and operational concerns
- 💡 Rename formatDate to serializeDate or toISOString to accurately reflect its behavior, or implement proper date formatting logic
- 💡 Replace direct console.log with configurable logger interface that can be mocked or swapped for different logging backends
- 💡 Add input validation to formatDate to handle invalid Date objects gracefully
- 💡 Consider adding timezone support or locale-aware formatting options for the formatDate function
- ⚠️ Misleading function name that doesn't match actual behavior - formatDate returns ISO string instead of formatted date
- ⚠️ Direct console.log usage makes logging non-configurable and hard to test/migrate in different environments
- 🔗 The logging function creates tight coupling between the utility module and console output, making it impossible to switch to structured logging or different output destinations
- 🔗 The formatDate function creates false expectations about date formatting capabilities and may break if consumers expect different output formats

</details>

### `vscode-certify/src/` (31 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`RunResult`](reports/vscode-certify-src-certifybinary-ts-runresult.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`findCertifyBinary`](reports/vscode-certify-src-certifybinary-ts-findcertifybinary.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`listModels`](reports/vscode-certify-src-certifybinary-ts-listmodels.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`promptInstall`](reports/vscode-certify-src-certifybinary-ts-promptinstall.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`runCertify`](reports/vscode-certify-src-certifybinary-ts-runcertify.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`runCertifyJSON`](reports/vscode-certify-src-certifybinary-ts-runcertifyjson.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`runInTerminal`](reports/vscode-certify-src-certifybinary-ts-runinterminal.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`BRAND_COLORS`](reports/vscode-certify-src-constants-ts-brand-colors.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`DIMENSION_NAMES`](reports/vscode-certify-src-constants-ts-dimension-names.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`GRADE_COLORS`](reports/vscode-certify-src-constants-ts-grade-colors.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`GRADE_EMOJI`](reports/vscode-certify-src-constants-ts-grade-emoji.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`PROVIDER_PRESETS`](reports/vscode-certify-src-constants-ts-provider-presets.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`CertifyDataLoader`](reports/vscode-certify-src-dataloader-ts-certifydataloader.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`activate`](reports/vscode-certify-src-extension-ts-activate.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`deactivate`](reports/vscode-certify-src-extension-ts-deactivate.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`createStatusBarItem`](reports/vscode-certify-src-statusbar-ts-createstatusbaritem.md) | function | B+ | 89.4% | certified | 2026-04-23 |
| [`AgentConfig`](reports/vscode-certify-src-types-ts-agentconfig.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`BadgeJSON`](reports/vscode-certify-src-types-ts-badgejson.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`CertifyCard`](reports/vscode-certify-src-types-ts-certifycard.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`CertifyConfig`](reports/vscode-certify-src-types-ts-certifyconfig.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`FullReport`](reports/vscode-certify-src-types-ts-fullreport.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`IndexEntry`](reports/vscode-certify-src-types-ts-indexentry.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`IssueCard`](reports/vscode-certify-src-types-ts-issuecard.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`LanguageCard`](reports/vscode-certify-src-types-ts-languagecard.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`LanguageDetail`](reports/vscode-certify-src-types-ts-languagedetail.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ModelAssignments`](reports/vscode-certify-src-types-ts-modelassignments.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ModelInfo`](reports/vscode-certify-src-types-ts-modelinfo.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ProviderConfig`](reports/vscode-certify-src-types-ts-providerconfig.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ProviderPreset`](reports/vscode-certify-src-types-ts-providerpreset.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`RecordJSON`](reports/vscode-certify-src-types-ts-recordjson.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`UnitReport`](reports/vscode-certify-src-types-ts-unitreport.md) | class | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The code is functionally correct but has several security and reliability issues including command injection risks, improper error handling, and lack of input validation.
- 💡 Use a more secure method than 'which' command to locate binaries, such as checking specific directories or using a whitelist of known locations
- 💡 Add proper input validation and sanitization for all user-provided paths, particularly the 'binaryPath' config setting
- 💡 Implement proper locking or atomic operations for the cachedBinaryPath variable to prevent race conditions
- 💡 Validate that returned paths from external commands are actually executable files before returning them as valid binaries
- 💡 Add type guards or more specific error handling in runCertify to ensure proper error extraction from execFileAsync errors
- 💡 Consider using a more robust binary detection library or implementing a whitelist of allowed binary locations
- ⚠️ Command injection vulnerability in execFileAsync calls
- ⚠️ Race condition in cachedBinaryPath variable access
- 🔗 The function is tightly coupled to VS Code's workspace and configuration APIs, making it hard to test in isolation
- 🔗 The caching mechanism introduces a global state dependency that can cause flaky behavior in concurrent environments

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
- 🤖 The code is functionally correct but has several security and reliability issues including command injection risks, improper error handling, and lack of input validation.
- 💡 Use a more secure method than 'which' command to locate binaries, such as checking specific directories or using a whitelist of known locations
- 💡 Add proper input validation and sanitization for all user-provided paths, particularly the 'binaryPath' config setting
- 💡 Implement proper locking or atomic operations for the cachedBinaryPath variable to prevent race conditions
- 💡 Validate that returned paths from external commands are actually executable files before returning them as valid binaries
- 💡 Add type guards or more specific error handling in runCertify to ensure proper error extraction from execFileAsync errors
- 💡 Consider using a more robust binary detection library or implementing a whitelist of allowed binary locations
- ⚠️ Command injection vulnerability in execFileAsync calls
- ⚠️ Race condition in cachedBinaryPath variable access
- 🔗 The function is tightly coupled to VS Code's workspace and configuration APIs, making it hard to test in isolation
- 🔗 The caching mechanism introduces a global state dependency that can cause flaky behavior in concurrent environments

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
- 🤖 The code is functionally correct but has several security and reliability issues including command injection risks, improper error handling, and lack of input validation.
- 💡 Use a more secure method than 'which' command to locate binaries, such as checking specific directories or using a whitelist of known locations
- 💡 Add proper input validation and sanitization for all user-provided paths, particularly the 'binaryPath' config setting
- 💡 Implement proper locking or atomic operations for the cachedBinaryPath variable to prevent race conditions
- 💡 Validate that returned paths from external commands are actually executable files before returning them as valid binaries
- 💡 Add type guards or more specific error handling in runCertify to ensure proper error extraction from execFileAsync errors
- 💡 Consider using a more robust binary detection library or implementing a whitelist of allowed binary locations
- ⚠️ Command injection vulnerability in execFileAsync calls
- ⚠️ Race condition in cachedBinaryPath variable access
- 🔗 The function is tightly coupled to VS Code's workspace and configuration APIs, making it hard to test in isolation
- 🔗 The caching mechanism introduces a global state dependency that can cause flaky behavior in concurrent environments

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
- 🤖 The code is functionally correct but has several security and reliability issues including command injection risks, improper error handling, and lack of input validation.
- 💡 Use a more secure method than 'which' command to locate binaries, such as checking specific directories or using a whitelist of known locations
- 💡 Add proper input validation and sanitization for all user-provided paths, particularly the 'binaryPath' config setting
- 💡 Implement proper locking or atomic operations for the cachedBinaryPath variable to prevent race conditions
- 💡 Validate that returned paths from external commands are actually executable files before returning them as valid binaries
- 💡 Add type guards or more specific error handling in runCertify to ensure proper error extraction from execFileAsync errors
- 💡 Consider using a more robust binary detection library or implementing a whitelist of allowed binary locations
- ⚠️ Command injection vulnerability in execFileAsync calls
- ⚠️ Race condition in cachedBinaryPath variable access
- 🔗 The function is tightly coupled to VS Code's workspace and configuration APIs, making it hard to test in isolation
- 🔗 The caching mechanism introduces a global state dependency that can cause flaky behavior in concurrent environments

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
- 🤖 The code is functionally correct but has several security and reliability issues including command injection risks, improper error handling, and lack of input validation.
- 💡 Use a more secure method than 'which' command to locate binaries, such as checking specific directories or using a whitelist of known locations
- 💡 Add proper input validation and sanitization for all user-provided paths, particularly the 'binaryPath' config setting
- 💡 Implement proper locking or atomic operations for the cachedBinaryPath variable to prevent race conditions
- 💡 Validate that returned paths from external commands are actually executable files before returning them as valid binaries
- 💡 Add type guards or more specific error handling in runCertify to ensure proper error extraction from execFileAsync errors
- 💡 Consider using a more robust binary detection library or implementing a whitelist of allowed binary locations
- ⚠️ Command injection vulnerability in execFileAsync calls
- ⚠️ Race condition in cachedBinaryPath variable access
- 🔗 The function is tightly coupled to VS Code's workspace and configuration APIs, making it hard to test in isolation
- 🔗 The caching mechanism introduces a global state dependency that can cause flaky behavior in concurrent environments

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
- 🤖 The code is functionally correct but has several security and reliability issues including command injection risks, improper error handling, and lack of input validation.
- 💡 Use a more secure method than 'which' command to locate binaries, such as checking specific directories or using a whitelist of known locations
- 💡 Add proper input validation and sanitization for all user-provided paths, particularly the 'binaryPath' config setting
- 💡 Implement proper locking or atomic operations for the cachedBinaryPath variable to prevent race conditions
- 💡 Validate that returned paths from external commands are actually executable files before returning them as valid binaries
- 💡 Add type guards or more specific error handling in runCertify to ensure proper error extraction from execFileAsync errors
- 💡 Consider using a more robust binary detection library or implementing a whitelist of allowed binary locations
- ⚠️ Command injection vulnerability in execFileAsync calls
- ⚠️ Race condition in cachedBinaryPath variable access
- 🔗 The function is tightly coupled to VS Code's workspace and configuration APIs, making it hard to test in isolation
- 🔗 The caching mechanism introduces a global state dependency that can cause flaky behavior in concurrent environments

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
- 🤖 The code is functionally correct but has several security and reliability issues including command injection risks, improper error handling, and lack of input validation.
- 💡 Use a more secure method than 'which' command to locate binaries, such as checking specific directories or using a whitelist of known locations
- 💡 Add proper input validation and sanitization for all user-provided paths, particularly the 'binaryPath' config setting
- 💡 Implement proper locking or atomic operations for the cachedBinaryPath variable to prevent race conditions
- 💡 Validate that returned paths from external commands are actually executable files before returning them as valid binaries
- 💡 Add type guards or more specific error handling in runCertify to ensure proper error extraction from execFileAsync errors
- 💡 Consider using a more robust binary detection library or implementing a whitelist of allowed binary locations
- ⚠️ Command injection vulnerability in execFileAsync calls
- ⚠️ Race condition in cachedBinaryPath variable access
- 🔗 The function is tightly coupled to VS Code's workspace and configuration APIs, making it hard to test in isolation
- 🔗 The caching mechanism introduces a global state dependency that can cause flaky behavior in concurrent environments

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
- 🤖 Well-structured constants with minor redundancy and missing type safety in emoji mapping.
- 💡 Use a union type for GRADE_EMOJI and GRADE_COLORS keys to enforce valid grades at compile time instead of using Record<string, string>.
- 💡 Consider defining a strict enum or union type for valid grades to prevent runtime access with invalid keys.
- 💡 Add runtime validation or a check in tests that ensures all valid grades have corresponding entries in GRADE_EMOJI and GRADE_COLORS.
- ⚠️ Grade emoji mapping inconsistency: Multiple grades map to the same emoji, which can obscure distinctions in UI or reporting.
- ⚠️ Lack of strict typing for GRADE_EMOJI and GRADE_COLORS — runtime errors possible if invalid keys are accessed.
- 🔗 GRADE_EMOJI and GRADE_COLORS are used across UI and analytics modules, so any inconsistency or missing key can propagate errors throughout the system.
- 🔗 PROVIDER_PRESETS are used in configuration and API setup logic, making them tightly coupled to external services and potentially causing misconfigurations or runtime issues if endpoints are unreachable.

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
- 🤖 Well-structured constants with minor redundancy and missing type safety in emoji mapping.
- 💡 Use a union type for GRADE_EMOJI and GRADE_COLORS keys to enforce valid grades at compile time instead of using Record<string, string>.
- 💡 Consider defining a strict enum or union type for valid grades to prevent runtime access with invalid keys.
- 💡 Add runtime validation or a check in tests that ensures all valid grades have corresponding entries in GRADE_EMOJI and GRADE_COLORS.
- ⚠️ Grade emoji mapping inconsistency: Multiple grades map to the same emoji, which can obscure distinctions in UI or reporting.
- ⚠️ Lack of strict typing for GRADE_EMOJI and GRADE_COLORS — runtime errors possible if invalid keys are accessed.
- 🔗 GRADE_EMOJI and GRADE_COLORS are used across UI and analytics modules, so any inconsistency or missing key can propagate errors throughout the system.
- 🔗 PROVIDER_PRESETS are used in configuration and API setup logic, making them tightly coupled to external services and potentially causing misconfigurations or runtime issues if endpoints are unreachable.

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
- 🤖 Well-structured constants with minor redundancy and missing type safety in emoji mapping.
- 💡 Use a union type for GRADE_EMOJI and GRADE_COLORS keys to enforce valid grades at compile time instead of using Record<string, string>.
- 💡 Consider defining a strict enum or union type for valid grades to prevent runtime access with invalid keys.
- 💡 Add runtime validation or a check in tests that ensures all valid grades have corresponding entries in GRADE_EMOJI and GRADE_COLORS.
- ⚠️ Grade emoji mapping inconsistency: Multiple grades map to the same emoji, which can obscure distinctions in UI or reporting.
- ⚠️ Lack of strict typing for GRADE_EMOJI and GRADE_COLORS — runtime errors possible if invalid keys are accessed.
- 🔗 GRADE_EMOJI and GRADE_COLORS are used across UI and analytics modules, so any inconsistency or missing key can propagate errors throughout the system.
- 🔗 PROVIDER_PRESETS are used in configuration and API setup logic, making them tightly coupled to external services and potentially causing misconfigurations or runtime issues if endpoints are unreachable.

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
- 🤖 Well-structured constants with minor redundancy and missing type safety in emoji mapping.
- 💡 Use a union type for GRADE_EMOJI and GRADE_COLORS keys to enforce valid grades at compile time instead of using Record<string, string>.
- 💡 Consider defining a strict enum or union type for valid grades to prevent runtime access with invalid keys.
- 💡 Add runtime validation or a check in tests that ensures all valid grades have corresponding entries in GRADE_EMOJI and GRADE_COLORS.
- ⚠️ Grade emoji mapping inconsistency: Multiple grades map to the same emoji, which can obscure distinctions in UI or reporting.
- ⚠️ Lack of strict typing for GRADE_EMOJI and GRADE_COLORS — runtime errors possible if invalid keys are accessed.
- 🔗 GRADE_EMOJI and GRADE_COLORS are used across UI and analytics modules, so any inconsistency or missing key can propagate errors throughout the system.
- 🔗 PROVIDER_PRESETS are used in configuration and API setup logic, making them tightly coupled to external services and potentially causing misconfigurations or runtime issues if endpoints are unreachable.

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
- 🤖 Well-structured constants with minor redundancy and missing type safety in emoji mapping.
- 💡 Use a union type for GRADE_EMOJI and GRADE_COLORS keys to enforce valid grades at compile time instead of using Record<string, string>.
- 💡 Consider defining a strict enum or union type for valid grades to prevent runtime access with invalid keys.
- 💡 Add runtime validation or a check in tests that ensures all valid grades have corresponding entries in GRADE_EMOJI and GRADE_COLORS.
- ⚠️ Grade emoji mapping inconsistency: Multiple grades map to the same emoji, which can obscure distinctions in UI or reporting.
- ⚠️ Lack of strict typing for GRADE_EMOJI and GRADE_COLORS — runtime errors possible if invalid keys are accessed.
- 🔗 GRADE_EMOJI and GRADE_COLORS are used across UI and analytics modules, so any inconsistency or missing key can propagate errors throughout the system.
- 🔗 PROVIDER_PRESETS are used in configuration and API setup logic, making them tightly coupled to external services and potentially causing misconfigurations or runtime issues if endpoints are unreachable.

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
- 🤖 The CertifyDataLoader class has good structure but contains several critical error handling gaps, missing type safety, and potential race conditions in file system operations.
- 💡 Add specific error logging in catch blocks to help with debugging and monitoring
- 💡 Validate the structure of loaded records before using them in `buildReportFromRecords` to prevent runtime errors
- 💡 Use proper error types instead of generic catch-all blocks to enable better error propagation and handling
- 💡 Implement input validation for `getUnitsForFile` to handle path normalization consistently across platforms
- ⚠️ Silent failures in file reading operations that could mask corrupted data or permission issues
- ⚠️ Potential runtime errors from malformed JSON or missing fields in records during report building
- 🔗 This class is tightly coupled to VSCode's file system watcher and workspace APIs, making it difficult to test in isolation
- 🔗 The `loadFullReport` fallback mechanism introduces a secondary data source dependency that isn't properly isolated or validated

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
- 🤖 The activate function is broadly functional but has several correctness and resource management issues, including missing error handling in command execution, improper disposal of resources, and potential race conditions.
- 💡 Fix ensureBinary to return true after successful installation or handle the case where user installs and retries
- 💡 Change diagnostics disposal to use proper subscription pattern: context.subscriptions.push(diagnostics)
- 💡 Add validation checks in certify.openUnit to ensure filePath is a valid file path before attempting to open it
- 💡 Implement error handling in runInTerminal for terminal creation and command execution failures
- ⚠️ Command execution may silently fail due to missing error handling in terminal commands
- ⚠️ Improper resource disposal in diagnostics and data loader may lead to memory leaks or dangling references
- 🔗 The extension's commands like 'certify.scan' and 'certify.certify' will not function properly if the binary is installed after initial activation
- 🔗 The hardcoded tree view ID 'certifyUnits' introduces a potential conflict with other extensions or future versions of this extension

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
- 🤖 The activate function is broadly functional but has several correctness and resource management issues, including missing error handling in command execution, improper disposal of resources, and potential race conditions.
- 💡 Fix ensureBinary to return true after successful installation or handle the case where user installs and retries
- 💡 Change diagnostics disposal to use proper subscription pattern: context.subscriptions.push(diagnostics)
- 💡 Add validation checks in certify.openUnit to ensure filePath is a valid file path before attempting to open it
- 💡 Implement error handling in runInTerminal for terminal creation and command execution failures
- ⚠️ Command execution may silently fail due to missing error handling in terminal commands
- ⚠️ Improper resource disposal in diagnostics and data loader may lead to memory leaks or dangling references
- 🔗 The extension's commands like 'certify.scan' and 'certify.certify' will not function properly if the binary is installed after initial activation
- 🔗 The hardcoded tree view ID 'certifyUnits' introduces a potential conflict with other extensions or future versions of this extension

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
- 🤖 The function creates a status bar item but has critical async handling and resource management issues.
- 💡 Add proper error handling in the async update function with try/catch and error logging
- 💡 Implement a dispose method or cleanup mechanism to remove the event listener when the status bar item is no longer needed
- 💡 Add validation for badge.message format before using it as a key in GRADE_COLORS
- 💡 Ensure the status bar item is disposed when the extension deactivates or component is destroyed
- ⚠️ Silent async errors in badge loading that can cause UI to be out of sync with actual data
- ⚠️ Memory leak from accumulating event listeners when status bar items are repeatedly created
- 🔗 Creates a persistent listener that accumulates over time, leading to memory bloat and potential performance degradation
- 🔗 The status bar item does not properly clean up its resources, causing potential resource leaks in long-running VS Code sessions

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

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
- 🤖 Well-structured TypeScript interfaces with clear JSON schema alignment, but lacks runtime validation and type safety for dynamic usage.
- 💡 Add runtime validation or Zod schemas to ensure data conforms to defined interfaces before processing
- 💡 Replace `Record<string, number>` with strongly typed enums or known key-value mappings where possible (e.g., for `grade_distribution` or `dimensions`) to prevent invalid keys from being accepted
- ⚠️ Missing runtime type validation for all interfaces, leading to potential runtime errors when data doesn't match expected structure
- ⚠️ Use of generic Record types (e.g., `Record<string, number>`) allows invalid keys to pass through without validation or constraints
- 🔗 These interfaces are used as type definitions for JSON parsing and API responses, so incorrect assumptions about structure can propagate throughout the system
- 🔗 The lack of strict type checking in consumers of these interfaces increases coupling and makes refactoring harder

</details>

### `vscode-certify/src/codeLens/` (2 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`CertifyCodeLensProvider`](reports/vscode-certify-src-codelens-certifycodelensprovider-ts-certifycodelensprovider.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`showDimensionScores`](reports/vscode-certify-src-codelens-certifycodelensprovider-ts-showdimensionscores.md) | function | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The code is functionally correct but has several maintainability and correctness issues including unsafe string operations, missing error handling, and potential performance bottlenecks.
- 💡 Replace `record.unit_id.split('#').pop()!` with a safer null check and fallback handling
- 💡 Add proper error handling in `loadRecords` to prevent unhandled promise rejections
- 💡 Improve the fallback symbol matching in `findSymbolLine` to use more precise text matching
- 💡 Add input validation for `record.dimensions` in `showDimensionScores` before accessing properties
- ⚠️ Runtime error from unsafe null assertion on line 52
- ⚠️ Incorrect CodeLens placement due to incomplete symbol matching logic
- 🔗 This unit tightly couples to the dataLoader and workspace root, creating a strong dependency that makes testing difficult
- 🔗 The CodeLens provider can cause performance degradation in large files due to linear search patterns

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
- 🤖 The code is functionally correct but has several maintainability and correctness issues including unsafe string operations, missing error handling, and potential performance bottlenecks.
- 💡 Replace `record.unit_id.split('#').pop()!` with a safer null check and fallback handling
- 💡 Add proper error handling in `loadRecords` to prevent unhandled promise rejections
- 💡 Improve the fallback symbol matching in `findSymbolLine` to use more precise text matching
- 💡 Add input validation for `record.dimensions` in `showDimensionScores` before accessing properties
- ⚠️ Runtime error from unsafe null assertion on line 52
- ⚠️ Incorrect CodeLens placement due to incomplete symbol matching logic
- 🔗 This unit tightly couples to the dataLoader and workspace root, creating a strong dependency that makes testing difficult
- 🔗 The CodeLens provider can cause performance degradation in large files due to linear search patterns

</details>

### `vscode-certify/src/config/` (5 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`ConfigPanel`](reports/vscode-certify-src-config-configpanel-ts-configpanel.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`ConnectionTestResult`](reports/vscode-certify-src-config-configwriter-ts-connectiontestresult.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`readConfig`](reports/vscode-certify-src-config-configwriter-ts-readconfig.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`testConnection`](reports/vscode-certify-src-config-configwriter-ts-testconnection.md) | function | B+ | 88.3% | certified | 2026-04-23 |
| [`writeConfig`](reports/vscode-certify-src-config-configwriter-ts-writeconfig.md) | function | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The ConfigPanel class has good structure but contains critical security and error handling flaws including hardcoded API key exposure, missing input validation, and improper error recovery.
- 💡 Remove API key data from webview DOM elements and use only internal state tracking
- 💡 Validate and sanitize all incoming message parameters before processing, especially API key environment variables
- 💡 Implement proper error recovery in the model fetching flow to prevent UI state inconsistency
- 💡 Escape HTML content when rendering models to prevent XSS vulnerabilities
- 💡 Use VSCode's theme variables consistently instead of hardcoded color values
- ⚠️ Hardcoded API key exposure in webview HTML and model lists
- ⚠️ Insecure API key handling with direct process.env access and no validation
- 🔗 This component creates a high-security surface area by handling sensitive API keys and exposing them in webview DOM
- 🔗 The component tightly couples to VSCode's secret storage and webview APIs, making it difficult to test or reuse outside the extension context

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
- 🤖 The writeConfig function has a critical race condition and lacks proper error handling for file I/O operations.
- 💡 Implement file locking or atomic write patterns in writeConfig to prevent race conditions
- 💡 Add explicit validation of parsed YAML structure before merging in writeConfig
- 💡 Use a proper HTTP client library with built-in timeout handling instead of raw http/https modules in testConnection
- 💡 Add proper error logging and re-throwing of exceptions in readConfig to surface configuration issues
- 💡 Validate that baseURL parameter is a valid URL before constructing the request in testConnection
- ⚠️ Race condition in writeConfig when multiple processes attempt to write to the same config file
- ⚠️ Potential data loss or corruption due to lack of atomic writes in writeConfig
- 🔗 writeConfig creates a potential point of failure for configuration persistence that could cause service instability
- 🔗 The testConnection function introduces a dependency on external HTTP services and lacks proper input validation

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
- 🤖 The writeConfig function has a critical race condition and lacks proper error handling for file I/O operations.
- 💡 Implement file locking or atomic write patterns in writeConfig to prevent race conditions
- 💡 Add explicit validation of parsed YAML structure before merging in writeConfig
- 💡 Use a proper HTTP client library with built-in timeout handling instead of raw http/https modules in testConnection
- 💡 Add proper error logging and re-throwing of exceptions in readConfig to surface configuration issues
- 💡 Validate that baseURL parameter is a valid URL before constructing the request in testConnection
- ⚠️ Race condition in writeConfig when multiple processes attempt to write to the same config file
- ⚠️ Potential data loss or corruption due to lack of atomic writes in writeConfig
- 🔗 writeConfig creates a potential point of failure for configuration persistence that could cause service instability
- 🔗 The testConnection function introduces a dependency on external HTTP services and lacks proper input validation

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
- 🤖 The writeConfig function has a critical race condition and lacks proper error handling for file I/O operations.
- 💡 Implement file locking or atomic write patterns in writeConfig to prevent race conditions
- 💡 Add explicit validation of parsed YAML structure before merging in writeConfig
- 💡 Use a proper HTTP client library with built-in timeout handling instead of raw http/https modules in testConnection
- 💡 Add proper error logging and re-throwing of exceptions in readConfig to surface configuration issues
- 💡 Validate that baseURL parameter is a valid URL before constructing the request in testConnection
- ⚠️ Race condition in writeConfig when multiple processes attempt to write to the same config file
- ⚠️ Potential data loss or corruption due to lack of atomic writes in writeConfig
- 🔗 writeConfig creates a potential point of failure for configuration persistence that could cause service instability
- 🔗 The testConnection function introduces a dependency on external HTTP services and lacks proper input validation

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
- 🤖 The writeConfig function has a critical race condition and lacks proper error handling for file I/O operations.
- 💡 Implement file locking or atomic write patterns in writeConfig to prevent race conditions
- 💡 Add explicit validation of parsed YAML structure before merging in writeConfig
- 💡 Use a proper HTTP client library with built-in timeout handling instead of raw http/https modules in testConnection
- 💡 Add proper error logging and re-throwing of exceptions in readConfig to surface configuration issues
- 💡 Validate that baseURL parameter is a valid URL before constructing the request in testConnection
- ⚠️ Race condition in writeConfig when multiple processes attempt to write to the same config file
- ⚠️ Potential data loss or corruption due to lack of atomic writes in writeConfig
- 🔗 writeConfig creates a potential point of failure for configuration persistence that could cause service instability
- 🔗 The testConnection function introduces a dependency on external HTTP services and lacks proper input validation

</details>

### `vscode-certify/src/dashboard/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`DashboardPanel`](reports/vscode-certify-src-dashboard-dashboardpanel-ts-dashboardpanel.md) | class | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The DashboardPanel class implements a VS Code webview dashboard with basic rendering and data handling, but has critical security and correctness issues including XSS vulnerabilities, missing error handling, and improper resource management.
- 💡 Implement proper HTML sanitization for all user-provided data before embedding into webview HTML
- 💡 Add try/catch blocks around data loading in update() method to handle and report errors gracefully
- 💡 Fix the event listener disposal by ensuring `this.dataLoader.onDataChanged()` returns a disposable that gets added to disposables
- 💡 Escape all dynamic content in the HTML template using a proper escaping function or library
- 💡 Add validation that `c.overall_grade` exists in GRADE_COLORS before accessing it to prevent undefined behavior
- ⚠️ XSS vulnerability through direct embedding of user-provided paths into DOM attributes
- ⚠️ Uncaught exceptions in data loading causing silent failures or crashes
- 🔗 The dashboard is tightly coupled to the CertifyDataLoader interface and will fail if that contract changes
- 🔗 The webview panel creates a persistent UI element that can leak memory if not properly disposed of
- 🔗 The update() method blocks the main thread during data loading, potentially freezing the UI

</details>

### `vscode-certify/src/diagnostics/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`CertifyDiagnostics`](reports/vscode-certify-src-diagnostics-certifydiagnostics-ts-certifydiagnostics.md) | class | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The CertifyDiagnostics class has functional diagnostics but suffers from critical correctness issues, poor error handling, and potential race conditions.
- 💡 Replace hardcoded diagnostic ranges with actual line/column positions from parsed content or use a more appropriate range like the first line of the file
- 💡 Add proper error handling around data loading and path operations with try/catch blocks and logging
- 💡 Implement a check to ensure documents are still valid before updating diagnostics in updateDocument()
- 💡 Add validation for record.expires_at to ensure it's a valid date string before parsing
- 💡 Use a WeakMap or similar structure to track document validity and prevent accessing disposed documents
- ⚠️ Hardcoded diagnostic range (0,0)-(0,0) will not correctly highlight relevant content in files
- ⚠️ No error handling for data loading or path operations, leading to silent failures
- 🔗 This class tightly couples with CertifyDataLoader and VS Code workspace, increasing system fragility
- 🔗 Improper disposal of disposables may cause memory leaks or resource contention in VS Code extension

</details>

### `vscode-certify/src/treeView/` (2 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`CertificationTreeProvider`](reports/vscode-certify-src-treeview-certificationtreeprovider-ts-certificationtreeprovider.md) | class | B+ | 88.3% | certified | 2026-04-23 |
| [`CertifyTreeItem`](reports/vscode-certify-src-treeview-certificationtreeprovider-ts-certifytreeitem.md) | class | B+ | 88.3% | certified | 2026-04-23 |

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
- 🤖 The code implements a tree view provider for certification data, but has critical runtime errors due to unsafe assumptions about data structure and missing error handling.
- 💡 Add null checks and fallbacks for `r.unit_path` before calling `path.dirname()` on line 37
- 💡 Fix the logic in `worstGrade` to properly identify worst grade by comparing indices correctly
- 💡 Wrap `dataLoader.loadAllRecords()` in try/catch and handle errors gracefully to prevent extension crashes
- 💡 Add validation for `record.grade` before using it as a key in GRADE_EMOJI
- 💡 Add proper unit tests for edge cases like missing fields, invalid paths, and malformed data
- ⚠️ Runtime crash when `r.unit_path` is undefined or invalid
- ⚠️ Incorrect grade calculation in worstGrade() function
- 🔗 This component is a core UI element that directly affects the user experience in VS Code
- 🔗 The tree view will show incorrect directory grouping and unit display when data is malformed

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
- 🤖 The code implements a tree view provider for certification data, but has critical runtime errors due to unsafe assumptions about data structure and missing error handling.
- 💡 Add null checks and fallbacks for `r.unit_path` before calling `path.dirname()` on line 37
- 💡 Fix the logic in `worstGrade` to properly identify worst grade by comparing indices correctly
- 💡 Wrap `dataLoader.loadAllRecords()` in try/catch and handle errors gracefully to prevent extension crashes
- 💡 Add validation for `record.grade` before using it as a key in GRADE_EMOJI
- 💡 Add proper unit tests for edge cases like missing fields, invalid paths, and malformed data
- ⚠️ Runtime crash when `r.unit_path` is undefined or invalid
- ⚠️ Incorrect grade calculation in worstGrade() function
- 🔗 This component is a core UI element that directly affects the user experience in VS Code
- 🔗 The tree view will show incorrect directory grouping and unit display when data is malformed

</details>

### `website/src/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`collections`](reports/website-src-content-config-ts-collections.md) | function | B+ | 89.4% | certified | 2026-04-23 |

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
- 🤖 Well-structured content collection configuration with minimal risk but limited extensibility
- 💡 Add proper error handling and validation around the collection definition to prevent application crashes on content loading failures
- 💡 Implement configuration parameters or environment variables to make the loader and schema more flexible for different content types or environments
- 💡 Consider adding support for multiple collections (docs, blog, api) to make this configuration more extensible
- 💡 Add documentation or comments explaining the purpose and limitations of this collection configuration
- ⚠️ No validation or error handling for content loading - if docsLoader() or docsSchema() fail, the application will crash without graceful handling
- ⚠️ Hardcoded configuration prevents extensibility - any additional content types or custom schemas would require modifying this file directly
- 🔗 This unit creates a tight coupling between the content configuration and specific Starlight dependencies, making it difficult to switch content management systems or adapt to different documentation requirements
- 🔗 The single collection definition limits the system's ability to handle multiple content types (blog posts, guides, API references) that would typically be needed in a comprehensive documentation site

</details>

---

*559 units certified. Generated by [Certify](https://github.com/iksnae/code-certification).*
