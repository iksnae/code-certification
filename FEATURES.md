# Code Certification System — Feature Acceptance Checklist

## 1. Product Foundation
- [x] The system exists as a standalone, production-capable solution rather than an ad hoc collection of scripts.
- [x] The system has a clearly defined core domain model independent of any single programming language being certified.
- [x] The system is implementation-language specific internally and certification-language agnostic externally.
- [x] The system is designed so core governance logic is separated from repository host integrations, analyzer adapters, and reporting adapters.
- [x] The system supports repository-native operation without requiring a permanently online centralized control plane for basic functionality.
- [x] The system can operate in local CLI mode and GitHub workflow mode.
- [x] The system supports advisory mode and enforcing mode.
- [x] The system has a documented architecture, configuration model, data model, and operational model.

## 2. Repository Onboarding / Bootstrap
- [x] The system can be initialized in an existing GitHub repository.
- [x] Initialization can be triggered manually through a GitHub workflow dispatch and/or local CLI command.
- [x] Initialization inspects the target repository structure before generating configuration.
- [x] Initialization detects relevant repository languages and project characteristics.
- [x] Initialization creates a repository-local `.certification/` structure.
- [x] Initialization generates starter configuration files.
- [x] Initialization generates starter policy packs.
- [x] Initialization generates the recurring certification workflow file(s).
- [x] Initialization generates an initial unit index for certifiable code units.
- [x] Initialization proposes changes through a pull request rather than mutating the default branch directly.
- [x] The initialization PR includes a human-readable summary of detected languages, assumptions, generated files, and next steps.
- [x] Maintainers can review and edit the generated setup before activation.
- [x] No recurring certification enforcement begins until the initialization PR is merged.

## 3. Repository-Local Configuration
- [x] Configuration is stored in the repository.
- [x] Configuration is versioned through normal source control.
- [x] Configuration is human-readable and machine-readable.
- [x] Configuration supports include/exclude patterns for certification scope.
- [x] Configuration supports enabling or disabling policy packs.
- [x] Configuration supports advisory vs enforcing mode.
- [x] Configuration supports GitHub issue synchronization settings.
- [x] Configuration supports recertification schedule settings.
- [x] Configuration supports expiry thresholds and risk weighting parameters.
- [x] Configuration supports analyzer/tool adapter settings.
- [x] Configuration supports agent-review enablement and model/provider configuration where applicable.

## 4. Policy-as-Code
- [x] Policies are stored in repository-local files.
- [x] Policies are versioned.
- [x] Policies can target all code globally.
- [x] Policies can target specific languages.
- [x] Policies can target specific paths, file patterns, or unit types.
- [x] Policies can define thresholds for metrics such as size, complexity, test expectations, or disallowed constructs.
- [x] Policies can define severity levels for violations.
- [x] Policies can be updated via normal pull request workflows.
- [x] Certification records retain the policy version under which the evaluation occurred.
- [x] Policy changes can trigger recertification.
- [x] The system can distinguish policy drift from code changes.
- [x] The system supports repository-defined overrides or exemptions with explicit rationale.

## 5. Certifiable Unit Discovery
- [x] The system can discover certifiable units automatically.
- [x] The system supports file-level unit discovery.at minimum.
- [x] The system supports richer unit discovery such as function, method, class, module, package, route, workflow, or migration where adapters exist.
- [x] The system assigns stable identifiers to discovered units.
- [x] Stable identifiers include sufficient information to distinguish language, path, and symbol when available.
- [x] Unsupported or partially supported languages fall back to a generic file-level model rather than failing the scan.
- [x] Discovery respects include/exclude configuration.
- [x] Discovery can exclude generated code, vendor code, build output, and other configured non-target areas.
- [x] The system maintains a persistent unit index.
- [x] The system can update the unit index over time.
- [x] The system can identify newly added units.
- [x] The system can identify removed units.
- [x] The system can detect changed units between revisions.
- [x] The system handles moved or renamed files honestly, preserving continuity where possible and clearly treating units as new where continuity cannot be established.

## 6. Language-Agnostic Support Model
- [x] The system is not limited to certifying only the implementation language of the engine.
- [x] The system supports polyglot repositories.
- [x] The system supports a generic certification path for unsupported languages.
- [x] The system supports language-aware adapters for supported languages.
- [x] The core certification model does not embed hard-coded assumptions specific to Go, TypeScript, or any one language.
- [x] New language adapters can be added without redesigning core domain concepts.
- [x] Reports can segment results by language where relevant.

## 7. Evidence Collection
- [x] The system can collect deterministic evidence from external analyzers and quality tools.
- [x] The system can ingest lint results where configured.
- [x] The system can ingest type-checking results where configured.
- [x] The system can ingest test results where configured.
- [x] The system can ingest static analysis results where configured.
- [x] The system can ingest code metrics such as complexity, size, or churn where configured.
- [x] The system can incorporate git history metadata relevant to stability and risk.
- [x] Evidence from different tools is normalized into a common internal model.
- [x] Evidence is attached to the relevant certification record.
- [x] Missing evidence is explicitly marked as missing rather than silently ignored.
- [x] Evidence collection failures are reported clearly.
- [x] Partial evidence does not get represented as complete evidence.
- [x] The system preserves enough evidence context to explain certification outcomes later.

## 8. Certification Evaluation Engine
- [x] The system evaluates discovered units against one or more active policy packs.
- [x] The system can evaluate multiple policy packs against a single unit.
- [x] Deterministic checks are authoritative for deterministic policy violations.
- [x] The system records rule violations explicitly.
- [x] The system computes results across multiple quality dimensions.
- [x] Weighted scoring is supported.
- [x] The system assigns a status for each evaluated unit.
- [x] The system can assign statuses including `certified`, `certified_with_observations`, `probationary`, `expired`, `decertified`, and `exempt`.
- [x] The system can assign an overall grade and/or score where configured.
- [x] The system can attach confidence values where applicable.
- [x] The system can attach required remediation actions where appropriate.
- [x] The system can provide a concise human-readable explanation of why a unit passed, failed, expired, or was downgraded.
- [x] The system does not falsely imply mathematical certainty where only heuristic evidence exists.

## 9. Certification Dimensions
- [x] The system supports evaluating correctness.
- [x] The system supports evaluating maintainability.
- [x] The system supports evaluating readability.
- [x] The system supports evaluating testability.
- [x] The system supports evaluating security.
- [x] The system supports evaluating architectural fitness.
- [x] The system supports evaluating operational quality.
- [x] The system supports evaluating performance appropriateness.
- [x] The system supports evaluating change risk.
- [x] The weighting of dimensions is configurable.
- [x] Reports expose dimension-level breakdowns for evaluated units when configured.

## 10. Certification Records / Trust Ledger
- [x] Certification records are stored in a structured format.
- [x] Certification records are versionable and auditable.
- [x] Each certification record includes the unit identifier.
- [x] Each certification record includes the unit type and path.
- [x] Each certification record includes the policy version used for evaluation.
- [x] Each certification record includes status.
- [x] Each certification record includes grade and/or score breakdown where configured.
- [x] Each certification record includes evidence references or summaries.
- [x] Each certification record includes observations.
- [x] Each certification record includes required actions where applicable.
- [x] Each certification record includes `certified_at` and `expires_at` for non-exempt states.
- [x] Each certification record includes enough historical metadata to support trust trend analysis.
- [x] Certification history can show repeated passes, failures, expiries, and overrides over time.

## 11. Expiry and Trust Decay
- [x] Every non-exempt certified unit receives an expiration date.
- [x] Expiration windows are computed from configurable rules.
- [x] New code can receive shorter initial trust windows.
- [x] Critical or high-risk code can receive shorter trust windows.
- [x] Stable repeatedly passing code can receive longer trust windows.
- [x] Maximum and minimum trust windows are configurable.
- [x] Certification can expire automatically when the trust window elapses.
- [x] Expired units are visibly marked as expired in reports and records.
- [x] The system can shorten trust windows based on churn, risk, or poor historical performance.
- [x] The system can lengthen trust windows based on sustained stability and repeated successful certification.
- [x] The system can distinguish expiry caused by elapsed time from invalidation caused by code or policy changes.

## 12. Invalidation and Recertification
- [x] Changes to code invalidate inherited trust for impacted units.
- [x] Policy changes can invalidate prior certifications.
- [x] Dependency or ecosystem drift can trigger recertification if configured.
- [x] The system can target recertification to changed units rather than the whole repository where appropriate.
- [x] The system can run a full repository recertification sweep when requested.
- [x] Recertification events are recorded in history.
- [x] Manual recertification can be triggered for a unit, file, directory, or repository.
- [x] The system can identify expiring-soon units for proactive review.
- [x] The system supports nightly, weekly, monthly, and/or annual recertification schedules.

## 13. GitHub Pull Request Workflow
- [x] The system runs in pull requests.
- [x] The PR workflow identifies changed units.
- [x] The PR workflow evaluates changed units using active policies.
- [x] The PR workflow reports newly certified units.
- [x] The PR workflow reports newly decertified or downgraded units.
- [x] The PR workflow reports newly introduced uncertified units.
- [x] The PR workflow can show trust delta caused by the change set.
- [x] The PR workflow annotates pull requests with actionable findings.
- [x] PR output is concise enough to be usable and detailed enough to be actionable.
- [x] The PR workflow supports advisory mode.
- [x] The PR workflow supports blocking mode for configured failure conditions.
- [x] Blocking behavior is configurable by severity, status, path, or criticality class.
- [x] PR-time certification does not unnecessarily re-certify unrelated unchanged units.

## 14. Scheduled GitHub Workflows
- [x] The system supports a scheduled nightly workflow.
- [x] The system supports a scheduled weekly workflow.
- [x] The system supports longer cadence full recertification workflows.
- [x] Scheduled workflows can detect expiring units.
- [x] Scheduled workflows can expire overdue certifications automatically.
- [x] Scheduled workflows can open or update remediation issues when configured.
- [x] Scheduled workflows generate reports without requiring manual intervention.
- [x] Scheduled workflows can run incrementally to limit cost and noise.
- [x] Scheduled workflows can be enabled or disabled through configuration.

## 15. GitHub Issue / Remediation Sync
- [x] The system can create GitHub issues for failing or expired certifications when configured.
- [x] Issue creation can be limited to selected statuses or risk classes.
- [x] Issues include unit identifiers and human-readable context.
- [x] Issues include policy violation summaries and/or evidence summaries.
- [x] Issues include recommended remediation guidance where available.
- [x] Issues are labeled appropriately by severity/type.
- [x] The system avoids duplicate issue spam.
- [x] The system can update existing issues as certification status evolves.
- [x] The system can close or resolve issues automatically under configured conditions.
- [x] The system supports single-unit remediation issues and grouped remediation issues.

## 16. Agent-Assisted Review
- [x] Agent-assisted review is optional.
- [x] Agent-assisted review can be enabled or disabled by configuration.
- [x] Agent output is clearly distinguished from deterministic evidence.
- [x] Agent review can provide explanatory observations for borderline or failing units.
- [x] Agent review can provide suggested remediation steps.
- [x] Agent review cannot silently override deterministic failures.
- [x] Human governance remains explicit where agent-assisted review is used.
- [x] The system can operate entirely without agent assistance.
- [x] The system can support local and/or remote model backends through adapters where implemented.

## 17. Human Overrides and Governance
- [x] The system supports explicit exemptions.
- [x] The system supports explicit manual overrides.
- [x] Overrides are stored in a dedicated structured format.
- [x] Overrides require rationale.
- [x] Overrides are visible in reports and certification records.
- [x] Overrides can extend trust windows when justified.
- [x] Overrides can shorten trust windows when justified.
- [x] Overrides can force recertification when justified.
- [x] The system can require human signoff for configured critical areas.
- [x] Human override actions are auditable and not silently destructive to evidence history.

## 18. Reporting
- [x] The system generates machine-readable reports.
- [x] The system generates human-readable reports.
- [x] Reports include total unit counts.
- [x] Reports include counts by certification status.
- [x] Reports include average grade and/or score summaries where configured.
- [x] Reports include expiring-soon units.
- [x] Reports include highest-risk units.
- [x] Reports include recurrently failing areas.
- [x] Reports include trend information over time.
- [x] Reports can summarize by directory, package, module, or language where useful.
- [x] Reports can explain why a unit is failing, expired, or high risk.
- [x] Reports are stored in predictable locations.
- [x] Reports do not require proprietary viewers to be useful.

## 19. Local CLI Experience
- [x] The system provides a local CLI.
- [x] The CLI can initialize certification configuration.
- [x] The CLI can scan for units.
- [x] The CLI can certify target scope.
- [x] The CLI can expire outdated certifications.
- [x] The CLI can generate reports.
- [x] The CLI can target a specific file, directory, or changed set.
- [x] Local CLI behavior respects repository configuration.
- [x] Local CLI output is usable for developers before pushing changes.
- [x] Local CLI output is reasonably consistent with CI results for the same scope.

## 20. Multi-Language Adapter Model
- [x] The architecture includes a defined adapter boundary for language-specific discovery and evidence collection.
- [x] The architecture includes a defined adapter boundary for analyzer execution and result normalization.
- [x] Adding a new language adapter does not require changing core certification concepts.
- [x] The system can support Go-specific analyzers through adapters.
- [x] The system can support TypeScript/JavaScript-specific analyzers through adapters.
- [x] The system can support future adapters for Python, Swift, shell, YAML/workflows, and other languages.
- [x] The system can function with mixed levels of support across languages in the same repository.

## 21. Security and Permissions
- [x] GitHub workflows declare explicit minimum required permissions.
- [x] The bootstrap process documents any required GitHub repository settings.
- [x] The system behaves safely under GitHub token permission constraints.
- [x] The design accounts for fork PR limitations and untrusted contribution scenarios.
- [x] The system avoids unsafe write operations in untrusted contexts.
- [x] The system does not require unnecessary elevated permissions for basic read/evaluate workflows.
- [x] Security-sensitive paths can be given stricter certification handling if configured.
- [x] Any use of privileged automation is clearly documented and isolated.

## 22. Storage, Schema, and Data Integrity
- [x] Structured schemas exist for configuration, policies, unit index, certification records, and reports.
- [x] Persisted artifacts are stable enough for downstream automation consumption.
- [x] Data files are reviewable in pull requests.
- [x] The system tolerates partially missing state by rebuilding what is derivable.
- [x] The system fails honestly when required state is invalid or irrecoverable.
- [x] The system can validate configuration and policy syntax before execution.
- [x] The system preserves historical continuity where possible without inventing false lineage.

## 23. Rollout and Adoption
- [x] The system supports incremental rollout in mature repositories.
- [x] The system can focus on changed code first.
- [x] The system can operate in advisory mode before enforcing mode.
- [x] The system can scope certification to selected directories or languages during rollout.
- [x] The system can distinguish baseline debt from newly introduced debt.
- [x] The system can expand toward full-repository governance over time through configuration changes rather than redesign.

## 24. Operational Quality
- [x] The system is documented well enough for maintainers to install, configure, run, and troubleshoot it.
- [x] The system includes installation documentation.
- [x] The system includes policy authoring documentation.
- [x] The system includes operator/maintainer documentation.
- [x] The system includes developer-facing usage documentation.
- [x] The system includes troubleshooting guidance.
- [x] The system includes example repository setup artifacts.
- [x] The system includes test coverage for core logic.
- [x] The system includes integration validation for GitHub workflow behavior.
- [x] The system produces deterministic behavior wherever deterministic inputs are available.

## 25. Architecture and Code Quality Expectations
- [x] Core domain logic is free of host-specific concerns where practical.
- [x] Core domain logic is free of analyzer-specific parsing hacks where practical.
- [x] No policy logic is hard-coded into orchestration paths that should be configuration-driven.
- [x] State mutation is intentional and traceable.
- [x] The solution demonstrates low tolerance for brittle backwards-compatibility hacks, dead code retention, and hidden behavior.
- [x] The codebase is lint-compliant.
- [x] The codebase is testable and modular.
- [x] The codebase favors explicit contracts over ambiguous coupling.
- [x] The implementation is maintainable enough to support future adapters, policy evolution, and platform growth.

## 26. Minimum v1 Readiness Criteria
- [x] A repository can be onboarded through a reviewable initialization PR.
- [x] Repository-local certification configuration is created successfully.
- [x] Repository-local starter policies are created successfully.
- [x] An initial unit index is created successfully.
- [x] Pull request workflows certify changed units successfully.
- [x] Scheduled workflows expire and recertify units successfully.
- [x] Certification records are written and updated successfully.
- [x] Reports are generated successfully in machine-readable and human-readable form.
- [x] Remediation issues can be created and updated successfully when configured.
- [x] Go and TypeScript/JavaScript receive at least initial language-aware support.
- [x] Unsupported languages still receive generic certification treatment.
- [x] The architecture and interfaces are ready for future language expansion without core redesign.

## 27. Program-Level Success Indicators
- [x] Teams can identify which parts of a codebase are currently trusted.
- [x] Teams can identify which parts of a codebase are expired, probationary, or decertified.
- [x] Teams can tell under which standards a given unit was last certified.
- [x] Teams can prioritize technical debt through evidence and trust decay rather than intuition alone.
- [x] Teams can adopt the system incrementally without overwhelming initial noise.
- [x] The system improves visibility into quality drift over time.
- [x] The system creates a durable governance model rather than a one-off quality report.