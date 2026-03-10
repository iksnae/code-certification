# 🟢 Certify — Report Card

**Repository:** `iksnae/code-certification`  
**Commit:** `dcf319d`  
**Generated:** 2026-03-10T02:55:35  

---

## Summary

| Metric | Value |
|--------|-------|
| **Overall Grade** | 🟢 **B+** |
| **Overall Score** | 87.2% |
| **Total Units** | 50 |
| **Passing** | 50 |
| **Failing** | 0 |
| **Pass Rate** | 100.0% |
| **Observations** | 0 |
| **Expired** | 0 |

## Grade Distribution

| Grade | Count | % | Bar |
|:-----:|------:|----:|-----|
| B+ | 32 | 64.0% | ████████████████████████████████ |
| B | 18 | 36.0% | ██████████████████ |

## Dimension Averages

| Dimension | Score | Bar |
|-----------|------:|-----|
| architectural_fitness | 80.0% | ███████████████░░░░░ |
| change_risk | 90.0% | █████████████████░░░ |
| correctness | 95.0% | ███████████████████░ |
| maintainability | 92.8% | ██████████████████░░ |
| operational_quality | 85.0% | █████████████████░░░ |
| performance_appropriateness | 80.0% | ███████████████░░░░░ |
| readability | 92.4% | ██████████████████░░ |
| security | 80.0% | ███████████████░░░░░ |
| testability | 90.0% | █████████████████░░░ |

## By Language

### go — 🟢 B+ (87.3%)

- **Units:** 46
- **Score range:** 85.0% – 87.8%
- **Grades:** 31×B+, 15×B

### ts — 🟢 B (86.8%)

- **Units:** 4
- **Score range:** 86.1% – 87.8%
- **Grades:** 1×B+, 3×B

## All Units

### `cmd/certify/` (3 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`loadCertifyContext`](reports/cmd-certify-certify-cmd-go-loadcertifycontext.md) | function | B | 86.7% | certified | 2026-04-24 |
| [`processQueue`](reports/cmd-certify-certify-cmd-go-processqueue.md) | method | B | 86.1% | certified | 2026-04-24 |
| [`saveReportArtifacts`](reports/cmd-certify-certify-cmd-go-savereportartifacts.md) | method | B+ | 87.8% | certified | 2026-04-24 |

<a id="cmd-certify-certify-cmd-go-loadcertifycontext"></a>
<details>
<summary>loadCertifyContext — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 95.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

<a id="cmd-certify-certify-cmd-go-processqueue"></a>
<details>
<summary>processQueue — certified details</summary>

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

<a id="cmd-certify-certify-cmd-go-savereportartifacts"></a>
<details>
<summary>saveReportArtifacts — certified details</summary>

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

### `internal/agent/` (13 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`DetectAPIKey`](reports/internal-agent-autodetect-go-detectapikey.md) | function | B+ | 87.8% | certified | 2026-04-24 |
| [`init`](reports/internal-agent-autodetect-go-init.md) | function | B+ | 87.8% | certified | 2026-04-24 |
| [`NewFallbackProvider`](reports/internal-agent-fallback-go-newfallbackprovider.md) | function | B+ | 87.8% | certified | 2026-04-24 |
| [`ollamaModel`](reports/internal-agent-models-go-ollamamodel.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`CoordinatorConfig`](reports/internal-agent-pipeline-go-coordinatorconfig.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`Strategy`](reports/internal-agent-pipeline-go-strategy.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`toResult`](reports/internal-agent-pipeline-go-toresult.md) | method | B+ | 87.8% | certified | 2026-04-24 |
| [`DetectProviders`](reports/internal-agent-providers-go-detectproviders.md) | function | B | 85.6% | certified | 2026-04-24 |
| [`Stage`](reports/internal-agent-stage-go-stage.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`defaultScores`](reports/internal-agent-stage-go-defaultscores.md) | function | B+ | 87.8% | certified | 2026-04-24 |
| [`looseParseNeedsReview`](reports/internal-agent-stage-go-looseparseneedsreview.md) | function | B+ | 87.8% | certified | 2026-04-24 |
| [`Message`](reports/internal-agent-types-go-message.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`TaskType`](reports/internal-agent-types-go-tasktype.md) | class | B+ | 87.8% | certified | 2026-04-24 |

<a id="internal-agent-autodetect-go-detectapikey"></a>
<details>
<summary>DetectAPIKey — certified details</summary>

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

<a id="internal-agent-autodetect-go-init"></a>
<details>
<summary>init — certified details</summary>

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

<a id="internal-agent-fallback-go-newfallbackprovider"></a>
<details>
<summary>NewFallbackProvider — certified details</summary>

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

<a id="internal-agent-models-go-ollamamodel"></a>
<details>
<summary>ollamaModel — certified details</summary>

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

<a id="internal-agent-pipeline-go-coordinatorconfig"></a>
<details>
<summary>CoordinatorConfig — certified details</summary>

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

<a id="internal-agent-pipeline-go-strategy"></a>
<details>
<summary>Strategy — certified details</summary>

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

<a id="internal-agent-pipeline-go-toresult"></a>
<details>
<summary>toResult — certified details</summary>

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

<a id="internal-agent-providers-go-detectproviders"></a>
<details>
<summary>DetectProviders — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

<a id="internal-agent-stage-go-stage"></a>
<details>
<summary>Stage — certified details</summary>

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

<a id="internal-agent-stage-go-defaultscores"></a>
<details>
<summary>defaultScores — certified details</summary>

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

<a id="internal-agent-stage-go-looseparseneedsreview"></a>
<details>
<summary>looseParseNeedsReview — certified details</summary>

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

<a id="internal-agent-types-go-message"></a>
<details>
<summary>Message — certified details</summary>

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

<a id="internal-agent-types-go-tasktype"></a>
<details>
<summary>TaskType — certified details</summary>

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

### `internal/config/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`LoadFromDir`](reports/internal-config-loader-go-loadfromdir.md) | function | B+ | 87.8% | certified | 2026-04-24 |

<a id="internal-config-loader-go-loadfromdir"></a>
<details>
<summary>LoadFromDir — certified details</summary>

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

### `internal/discovery/` (8 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`buildLanguageList`](reports/internal-discovery-detect-go-buildlanguagelist.md) | function | B+ | 87.8% | certified | 2026-04-24 |
| [`GoAdapter`](reports/internal-discovery-go-adapter-go-goadapter.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`NewGoAdapter`](reports/internal-discovery-go-adapter-go-newgoadapter.md) | function | B+ | 87.8% | certified | 2026-04-24 |
| [`NewIndex`](reports/internal-discovery-index-go-newindex.md) | function | B+ | 87.8% | certified | 2026-04-24 |
| [`indexEntry`](reports/internal-discovery-index-go-indexentry.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`Scanner`](reports/internal-discovery-scanner-go-scanner.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`scanner_test.go`](reports/internal-discovery-scanner-test-go.md) | file | B | 86.1% | certified | 2026-04-24 |
| [`NewTSAdapter`](reports/internal-discovery-ts-adapter-go-newtsadapter.md) | function | B+ | 87.8% | certified | 2026-04-24 |

<a id="internal-discovery-detect-go-buildlanguagelist"></a>
<details>
<summary>buildLanguageList — certified details</summary>

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

<a id="internal-discovery-go-adapter-go-goadapter"></a>
<details>
<summary>GoAdapter — certified details</summary>

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

<a id="internal-discovery-go-adapter-go-newgoadapter"></a>
<details>
<summary>NewGoAdapter — certified details</summary>

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

<a id="internal-discovery-index-go-newindex"></a>
<details>
<summary>NewIndex — certified details</summary>

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

<a id="internal-discovery-index-go-indexentry"></a>
<details>
<summary>indexEntry — certified details</summary>

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

<a id="internal-discovery-scanner-go-scanner"></a>
<details>
<summary>Scanner — certified details</summary>

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

<a id="internal-discovery-scanner-test-go"></a>
<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 80.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

<a id="internal-discovery-ts-adapter-go-newtsadapter"></a>
<details>
<summary>NewTSAdapter — certified details</summary>

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

### `internal/domain/` (6 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`config_test.go`](reports/internal-domain-config-test-go.md) | file | B | 86.7% | certified | 2026-04-24 |
| [`evidence_test.go`](reports/internal-domain-evidence-test-go.md) | file | B | 86.7% | certified | 2026-04-24 |
| [`policy_test.go`](reports/internal-domain-policy-test-go.md) | file | B | 86.7% | certified | 2026-04-24 |
| [`Status`](reports/internal-domain-record-go-status.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`record_test.go`](reports/internal-domain-record-test-go.md) | file | B | 86.7% | certified | 2026-04-24 |
| [`Unit`](reports/internal-domain-unit-go-unit.md) | class | B+ | 87.8% | certified | 2026-04-24 |

<a id="internal-domain-config-test-go"></a>
<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

<a id="internal-domain-evidence-test-go"></a>
<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

<a id="internal-domain-policy-test-go"></a>
<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

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

<a id="internal-domain-record-test-go"></a>
<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

<a id="internal-domain-unit-go-unit"></a>
<details>
<summary>Unit — certified details</summary>

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

### `internal/engine/` (3 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`CertifyUnit`](reports/internal-engine-pipeline-go-certifyunit.md) | function | B | 86.7% | certified | 2026-04-24 |
| [`Score`](reports/internal-engine-scorer-go-score.md) | function | B | 86.1% | certified | 2026-04-24 |
| [`extractSummaryFloat`](reports/internal-engine-scorer-go-extractsummaryfloat.md) | function | B | 86.7% | certified | 2026-04-24 |

<a id="internal-engine-pipeline-go-certifyunit"></a>
<details>
<summary>CertifyUnit — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 95.0% |
| security | 80.0% |
| testability | 90.0% |

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

<a id="internal-engine-scorer-go-extractsummaryfloat"></a>
<details>
<summary>extractSummaryFloat — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 95.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

### `internal/evidence/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`Collector`](reports/internal-evidence-collector-go-collector.md) | class | B+ | 87.8% | certified | 2026-04-24 |

<a id="internal-evidence-collector-go-collector"></a>
<details>
<summary>Collector — certified details</summary>

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

### `internal/override/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`rawOverrideFile`](reports/internal-override-loader-go-rawoverridefile.md) | class | B+ | 87.8% | certified | 2026-04-24 |

<a id="internal-override-loader-go-rawoverridefile"></a>
<details>
<summary>rawOverrideFile — certified details</summary>

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

### `internal/policy/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`extractMetric`](reports/internal-policy-evaluator-go-extractmetric.md) | function | B | 85.0% | certified | 2026-04-24 |

<a id="internal-policy-evaluator-go-extractmetric"></a>
<details>
<summary>extractMetric — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 80.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

**Observations:**
- todo_count: 2 exceeds threshold 0

</details>

### `internal/queue/` (3 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`Item`](reports/internal-queue-queue-go-item.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`Reset`](reports/internal-queue-queue-go-reset.md) | method | B+ | 87.8% | certified | 2026-04-24 |
| [`Save`](reports/internal-queue-queue-go-save.md) | method | B+ | 87.8% | certified | 2026-04-24 |

<a id="internal-queue-queue-go-item"></a>
<details>
<summary>Item — certified details</summary>

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

<a id="internal-queue-queue-go-reset"></a>
<details>
<summary>Reset — certified details</summary>

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

<a id="internal-queue-queue-go-save"></a>
<details>
<summary>Save — certified details</summary>

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

### `internal/report/` (6 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`badge_test.go`](reports/internal-report-badge-test-go.md) | file | B | 86.7% | certified | 2026-04-24 |
| [`buildTopIssues`](reports/internal-report-card-go-buildtopissues.md) | function | B | 86.7% | certified | 2026-04-24 |
| [`FormatDetailedText`](reports/internal-report-detailed-go-formatdetailedtext.md) | function | B | 85.0% | certified | 2026-04-24 |
| [`unitSummaryFrom`](reports/internal-report-detailed-go-unitsummaryfrom.md) | function | B+ | 87.8% | certified | 2026-04-24 |
| [`FullReport`](reports/internal-report-full-go-fullreport.md) | class | B+ | 87.8% | certified | 2026-04-24 |
| [`FormatText`](reports/internal-report-health-go-formattext.md) | function | B+ | 87.8% | certified | 2026-04-24 |

<a id="internal-report-badge-test-go"></a>
<details>
<summary> — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

<a id="internal-report-card-go-buildtopissues"></a>
<details>
<summary>buildTopIssues — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 85.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 95.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

<a id="internal-report-detailed-go-formatdetailedtext"></a>
<details>
<summary>FormatDetailedText — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 80.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

<a id="internal-report-detailed-go-unitsummaryfrom"></a>
<details>
<summary>unitSummaryFrom — certified details</summary>

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

<a id="internal-report-full-go-fullreport"></a>
<details>
<summary>FullReport — certified details</summary>

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

<a id="internal-report-health-go-formattext"></a>
<details>
<summary>FormatText — certified details</summary>

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

### `vscode-certify/src/` (2 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`CertifyDataLoader`](reports/vscode-certify-src-dataloader-ts-certifydataloader.md) | class | B | 86.1% | certified | 2026-04-24 |
| [`createStatusBarItem`](reports/vscode-certify-src-statusbar-ts-createstatusbaritem.md) | function | B+ | 87.8% | certified | 2026-04-24 |

<a id="vscode-certify-src-dataloader-ts-certifydataloader"></a>
<details>
<summary>CertifyDataLoader — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 80.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

<a id="vscode-certify-src-statusbar-ts-createstatusbaritem"></a>
<details>
<summary>createStatusBarItem — certified details</summary>

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

### `vscode-certify/src/config/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`ConnectionTestResult`](reports/vscode-certify-src-config-configwriter-ts-connectiontestresult.md) | class | B | 86.7% | certified | 2026-04-24 |

<a id="vscode-certify-src-config-configwriter-ts-connectiontestresult"></a>
<details>
<summary>ConnectionTestResult — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

### `vscode-certify/src/treeView/` (1 units)

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------|
| [`CertificationTreeProvider`](reports/vscode-certify-src-treeview-certificationtreeprovider-ts-certificationtreeprovider.md) | class | B | 86.7% | certified | 2026-04-24 |

<a id="vscode-certify-src-treeview-certificationtreeprovider-ts-certificationtreeprovider"></a>
<details>
<summary>CertificationTreeProvider — certified details</summary>

| Dimension | Score |
|-----------|------:|
| architectural_fitness | 80.0% |
| change_risk | 90.0% |
| correctness | 95.0% |
| maintainability | 95.0% |
| operational_quality | 85.0% |
| performance_appropriateness | 80.0% |
| readability | 85.0% |
| security | 80.0% |
| testability | 90.0% |

</details>

---

*50 units certified. Generated by [Certify](https://github.com/iksnae/code-certification).*
