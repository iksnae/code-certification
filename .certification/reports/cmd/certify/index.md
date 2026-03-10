# 🟢 `cmd/certify`

[← All Packages](../../index.md) · [← Report Card](../../../REPORT_CARD.md)

**Grade:** 🟢 B (85.3%)  
**Units:** 54 · **Passing:** 54 / 54

## Units

| Unit | Type | Grade | Score | Status | Expires |
|------|------|:-----:|------:|--------|--------:|
| [main](main.go/main.md) | function | 🟡 C | 77.9% | certified_with_observations | 2026-04-24 |
| [processQueue](certify_cmd.go/processQueue.md) | method | 🟡 C | 79.3% | certified_with_observations | 2026-04-24 |
| [setupArchitectProvider](architect_cmd.go/setupArchitectProvider.md) | function | 🟡 C | 79.3% | certified_with_observations | 2026-04-24 |
| [runWorkspaceReport](report_cmd.go/runWorkspaceReport.md) | function | 🟡 C | 79.3% | certified_with_observations | 2026-04-24 |
| [runCertify](certify_cmd.go/runCertify.md) | function | 🟡 C | 79.3% | certified_with_observations | 2026-04-24 |
| [detectAPIKeyOnly](certify_cmd.go/detectAPIKeyOnly.md) | function | 🟡 C | 80.0% | certified_with_observations | 2026-04-24 |
| [bindVersionInfo](version.go/bindVersionInfo.md) | function | 🟢 B | 80.7% | certified | 2026-04-24 |
| [setupExplicitAgent](certify_cmd.go/setupExplicitAgent.md) | function | 🟢 B | 81.4% | certified | 2026-04-24 |
| [languagePolicy](init_cmd.go/languagePolicy.md) | function | 🟢 B | 81.4% | certified | 2026-04-24 |
| [loadCertifyContext](certify_cmd.go/loadCertifyContext.md) | function | 🟢 B | 85.0% | certified | 2026-04-24 |
| [runWorkspaceCertify](certify_cmd.go/runWorkspaceCertify.md) | function | 🟢 B | 85.0% | certified | 2026-04-24 |
| [runWorkspaceInit](init_cmd.go/runWorkspaceInit.md) | function | 🟢 B | 85.0% | certified | 2026-04-24 |
| [gradeEmoji](certify_cmd.go/gradeEmoji.md) | function | 🟢 B | 85.0% | certified | 2026-04-24 |
| [runArchitect](architect_cmd.go/runArchitect.md) | function | 🟢 B | 85.0% | certified | 2026-04-24 |
| [bindExpireFlags](expire.go/bindExpireFlags.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [certifyContext](certify_cmd.go/certifyContext.md) | class | 🟢 B | 86.4% | certified | 2026-04-24 |
| [policyVersions](certify_cmd.go/policyVersions.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [printQueueStatus](certify_cmd.go/printQueueStatus.md) | method | 🟢 B | 86.4% | certified | 2026-04-24 |
| [printSummary](certify_cmd.go/printSummary.md) | method | 🟢 B | 86.4% | certified | 2026-04-24 |
| [isLocalURL](certify_cmd.go/isLocalURL.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [getCertifyFlags](certify_cmd.go/getCertifyFlags.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [runParams](certify_cmd.go/runParams.md) | class | 🟢 B | 86.4% | certified | 2026-04-24 |
| [formatETA](certify_cmd.go/formatETA.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [setupAgent](certify_cmd.go/setupAgent.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [setupConservativeAgent](certify_cmd.go/setupConservativeAgent.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [filterUnits](certify_cmd.go/filterUnits.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [bindCertifyFlags](certify_cmd.go/bindCertifyFlags.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [bindArchitectFlags](architect_cmd.go/bindArchitectFlags.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [runWorkspaceExpire](expire.go/runWorkspaceExpire.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [bindInitFlags](init_cmd.go/bindInitFlags.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [generateConfig](init_cmd.go/generateConfig.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [defaultConfigObj](certify_cmd.go/defaultConfigObj.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [certifyFlags](certify_cmd.go/certifyFlags.md) | class | 🟢 B | 86.4% | certified | 2026-04-24 |
| [loadQueue](certify_cmd.go/loadQueue.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [bindModelsFlags](models_cmd.go/bindModelsFlags.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [listFromProvider](models_cmd.go/listFromProvider.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [runModels](models_cmd.go/runModels.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [bindReportFlags](report_cmd.go/bindReportFlags.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [detectCommit](report_cmd.go/detectCommit.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [detectRepoName](report_cmd.go/detectRepoName.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [gradeEmojiShort](report_cmd.go/gradeEmojiShort.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [reportDirOf](report_cmd.go/reportDirOf.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [buildCertificationRun](certify_cmd.go/buildCertificationRun.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [bindReviewFlags](review.go/bindReviewFlags.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [flagBool](root.go/flagBool.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [flagInt](root.go/flagInt.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [flagString](root.go/flagString.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [flagStringSlice](root.go/flagStringSlice.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [registerCommands](root.go/registerCommands.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [bindScanFlags](scan.go/bindScanFlags.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [runWorkspaceScan](scan.go/runWorkspaceScan.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [tryScanSuggestions](scan.go/tryScanSuggestions.md) | function | 🟢 B | 86.4% | certified | 2026-04-24 |
| [cli_test.go](cli_test.go.md) | file | 🟢 B+ | 89.3% | certified | 2026-04-24 |
| [runSubcommand](workspace_dispatch.go/runSubcommand.md) | function | 🟢 A- | 90.7% | certified | 2026-04-24 |

---

*Generated by [Certify](https://github.com/iksnae/code-certification)*
