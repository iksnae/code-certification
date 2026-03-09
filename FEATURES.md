# Code Certification System — Feature Acceptance Checklist

## 1. Product Foundation
- [ ] The system exists as a standalone, production-capable solution rather than an ad hoc collection of scripts.
- [x] The system has a clearly defined core domain model independent of any single programming language being certified.
- [x] The system is implementation-language specific internally and certification-language agnostic externally.
- [ ] The system is designed so core governance logic is separated from repository host integrations, analyzer adapters, and reporting adapters.
- [ ] The system supports repository-native operation without requiring a permanently online centralized control plane for basic functionality.
- [ ] The system can operate in local CLI mode and GitHub workflow mode.
- [ ] The system supports advisory mode and enforcing mode.
- [ ] The system has a documented architecture, configuration model, data model, and operational model.

## 2. Repository Onboarding / Bootstrap
- [ ] The system can be initialized in an existing GitHub repository.
- [ ] Initialization can be triggered manually through a GitHub workflow dispatch and/or local CLI command.
- [ ] Initialization inspects the target repository structure before generating configuration.
- [ ] Initialization detects relevant repository languages and project characteristics.
- [ ] Initialization creates a repository-local `.certification/` structure.
- [ ] Initialization generates starter configuration files.
- [ ] Initialization generates starter policy packs.
- [ ] Initialization generates the recurring certification workflow file(s).
- [ ] Initialization generates an initial unit index for certifiable code units.
- [ ] Initialization proposes changes through a pull request rather than mutating the default branch directly.
- [ ] The initialization PR includes a human-readable summary of detected languages, assumptions, generated files, and next steps.
- [ ] Maintainers can review and edit the generated setup before activation.
- [ ] No recurring certification enforcement begins until the initialization PR is merged.

## 3. Repository-Local Configuration
- [ ] Configuration is stored in the repository.
- [ ] Configuration is versioned through normal source control.
- [ ] Configuration is human-readable and machine-readable.
- [ ] Configuration supports include/exclude patterns for certification scope.
- [ ] Configuration supports enabling or disabling policy packs.
- [ ] Configuration supports advisory vs enforcing mode.
- [ ] Configuration supports GitHub issue synchronization settings.
- [ ] Configuration supports recertification schedule settings.
- [ ] Configuration supports expiry thresholds and risk weighting parameters.
- [ ] Configuration supports analyzer/tool adapter settings.
- [ ] Configuration supports agent-review enablement and model/provider configuration where applicable.

## 4. Policy-as-Code
- [ ] Policies are stored in repository-local files.
- [ ] Policies are versioned.
- [ ] Policies can target all code globally.
- [ ] Policies can target specific languages.
- [ ] Policies can target specific paths, file patterns, or unit types.
- [ ] Policies can define thresholds for metrics such as size, complexity, test expectations, or disallowed constructs.
- [ ] Policies can define severity levels for violations.
- [ ] Policies can be updated via normal pull request workflows.
- [ ] Certification records retain the policy version under which the evaluation occurred.
- [ ] Policy changes can trigger recertification.
- [ ] The system can distinguish policy drift from code changes.
- [ ] The system supports repository-defined overrides or exemptions with explicit rationale.

## 5. Certifiable Unit Discovery
- [ ] The system can discover certifiable units automatically.
- [ ] The system supports file-level unit discovery at minimum.
- [ ] The system supports richer unit discovery such as function, method, class, module, package, route, workflow, or migration where adapters exist.
- [ ] The system assigns stable identifiers to discovered units.
- [ ] Stable identifiers include sufficient information to distinguish language, path, and symbol when available.
- [ ] Unsupported or partially supported languages fall back to a generic file-level model rather than failing the scan.
- [ ] Discovery respects include/exclude configuration.
- [ ] Discovery can exclude generated code, vendor code, build output, and other configured non-target areas.
- [ ] The system maintains a persistent unit index.
- [ ] The system can update the unit index over time.
- [ ] The system can identify newly added units.
- [ ] The system can identify removed units.
- [ ] The system can detect changed units between revisions.
- [ ] The system handles moved or renamed files honestly, preserving continuity where possible and clearly treating units as new where continuity cannot be established.

## 6. Language-Agnostic Support Model
- [ ] The system is not limited to certifying only the implementation language of the engine.
- [ ] The system supports polyglot repositories.
- [ ] The system supports a generic certification path for unsupported languages.
- [ ] The system supports language-aware adapters for supported languages.
- [ ] The core certification model does not embed hard-coded assumptions specific to Go, TypeScript, or any one language.
- [ ] New language adapters can be added without redesigning core domain concepts.
- [ ] Reports can segment results by language where relevant.

## 7. Evidence Collection
- [ ] The system can collect deterministic evidence from external analyzers and quality tools.
- [ ] The system can ingest lint results where configured.
- [ ] The system can ingest type-checking results where configured.
- [ ] The system can ingest test results where configured.
- [ ] The system can ingest static analysis results where configured.
- [ ] The system can ingest code metrics such as complexity, size, or churn where configured.
- [ ] The system can incorporate git history metadata relevant to stability and risk.
- [ ] Evidence from different tools is normalized into a common internal model.
- [ ] Evidence is attached to the relevant certification record.
- [ ] Missing evidence is explicitly marked as missing rather than silently ignored.
- [ ] Evidence collection failures are reported clearly.
- [ ] Partial evidence does not get represented as complete evidence.
- [ ] The system preserves enough evidence context to explain certification outcomes later.

## 8. Certification Evaluation Engine
- [ ] The system evaluates discovered units against one or more active policy packs.
- [ ] The system can evaluate multiple policy packs against a single unit.
- [ ] Deterministic checks are authoritative for deterministic policy violations.
- [ ] The system records rule violations explicitly.
- [ ] The system computes results across multiple quality dimensions.
- [ ] Weighted scoring is supported.
- [ ] The system assigns a status for each evaluated unit.
- [x] The system can assign statuses including `certified`, `certified_with_observations`, `probationary`, `expired`, `decertified`, and `exempt`.
- [ ] The system can assign an overall grade and/or score where configured.
- [ ] The system can attach confidence values where applicable.
- [ ] The system can attach required remediation actions where appropriate.
- [ ] The system can provide a concise human-readable explanation of why a unit passed, failed, expired, or was downgraded.
- [ ] The system does not falsely imply mathematical certainty where only heuristic evidence exists.

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
- [ ] Reports expose dimension-level breakdowns for evaluated units when configured.

## 10. Certification Records / Trust Ledger
- [ ] Certification records are stored in a structured format.
- [ ] Certification records are versionable and auditable.
- [ ] Each certification record includes the unit identifier.
- [ ] Each certification record includes the unit type and path.
- [ ] Each certification record includes the policy version used for evaluation.
- [ ] Each certification record includes status.
- [ ] Each certification record includes grade and/or score breakdown where configured.
- [ ] Each certification record includes evidence references or summaries.
- [ ] Each certification record includes observations.
- [ ] Each certification record includes required actions where applicable.
- [ ] Each certification record includes `certified_at` and `expires_at` for non-exempt states.
- [ ] Each certification record includes enough historical metadata to support trust trend analysis.
- [ ] Certification history can show repeated passes, failures, expiries, and overrides over time.

## 11. Expiry and Trust Decay
- [ ] Every non-exempt certified unit receives an expiration date.
- [ ] Expiration windows are computed from configurable rules.
- [ ] New code can receive shorter initial trust windows.
- [ ] Critical or high-risk code can receive shorter trust windows.
- [ ] Stable repeatedly passing code can receive longer trust windows.
- [ ] Maximum and minimum trust windows are configurable.
- [ ] Certification can expire automatically when the trust window elapses.
- [ ] Expired units are visibly marked as expired in reports and records.
- [ ] The system can shorten trust windows based on churn, risk, or poor historical performance.
- [ ] The system can lengthen trust windows based on sustained stability and repeated successful certification.
- [ ] The system can distinguish expiry caused by elapsed time from invalidation caused by code or policy changes.

## 12. Invalidation and Recertification
- [ ] Changes to code invalidate inherited trust for impacted units.
- [ ] Policy changes can invalidate prior certifications.
- [ ] Dependency or ecosystem drift can trigger recertification if configured.
- [ ] The system can target recertification to changed units rather than the whole repository where appropriate.
- [ ] The system can run a full repository recertification sweep when requested.
- [ ] Recertification events are recorded in history.
- [ ] Manual recertification can be triggered for a unit, file, directory, or repository.
- [ ] The system can identify expiring-soon units for proactive review.
- [ ] The system supports nightly, weekly, monthly, and/or annual recertification schedules.

## 13. GitHub Pull Request Workflow
- [ ] The system runs in pull requests.
- [ ] The PR workflow identifies changed units.
- [ ] The PR workflow evaluates changed units using active policies.
- [ ] The PR workflow reports newly certified units.
- [ ] The PR workflow reports newly decertified or downgraded units.
- [ ] The PR workflow reports newly introduced uncertified units.
- [ ] The PR workflow can show trust delta caused by the change set.
- [ ] The PR workflow annotates pull requests with actionable findings.
- [ ] PR output is concise enough to be usable and detailed enough to be actionable.
- [ ] The PR workflow supports advisory mode.
- [ ] The PR workflow supports blocking mode for configured failure conditions.
- [ ] Blocking behavior is configurable by severity, status, path, or criticality class.
- [ ] PR-time certification does not unnecessarily re-certify unrelated unchanged units.

## 14. Scheduled GitHub Workflows
- [ ] The system supports a scheduled nightly workflow.
- [ ] The system supports a scheduled weekly workflow.
- [ ] The system supports longer cadence full recertification workflows.
- [ ] Scheduled workflows can detect expiring units.
- [ ] Scheduled workflows can expire overdue certifications automatically.
- [ ] Scheduled workflows can open or update remediation issues when configured.
- [ ] Scheduled workflows generate reports without requiring manual intervention.
- [ ] Scheduled workflows can run incrementally to limit cost and noise.
- [ ] Scheduled workflows can be enabled or disabled through configuration.

## 15. GitHub Issue / Remediation Sync
- [ ] The system can create GitHub issues for failing or expired certifications when configured.
- [ ] Issue creation can be limited to selected statuses or risk classes.
- [ ] Issues include unit identifiers and human-readable context.
- [ ] Issues include policy violation summaries and/or evidence summaries.
- [ ] Issues include recommended remediation guidance where available.
- [ ] Issues are labeled appropriately by severity/type.
- [ ] The system avoids duplicate issue spam.
- [ ] The system can update existing issues as certification status evolves.
- [ ] The system can close or resolve issues automatically under configured conditions.
- [ ] The system supports single-unit remediation issues and grouped remediation issues.

## 16. Agent-Assisted Review
- [ ] Agent-assisted review is optional.
- [ ] Agent-assisted review can be enabled or disabled by configuration.
- [ ] Agent output is clearly distinguished from deterministic evidence.
- [ ] Agent review can provide explanatory observations for borderline or failing units.
- [ ] Agent review can provide suggested remediation steps.
- [ ] Agent review cannot silently override deterministic failures.
- [ ] Human governance remains explicit where agent-assisted review is used.
- [ ] The system can operate entirely without agent assistance.
- [ ] The system can support local and/or remote model backends through adapters where implemented.

## 17. Human Overrides and Governance
- [ ] The system supports explicit exemptions.
- [ ] The system supports explicit manual overrides.
- [ ] Overrides are stored in a dedicated structured format.
- [ ] Overrides require rationale.
- [ ] Overrides are visible in reports and certification records.
- [ ] Overrides can extend trust windows when justified.
- [ ] Overrides can shorten trust windows when justified.
- [ ] Overrides can force recertification when justified.
- [ ] The system can require human signoff for configured critical areas.
- [ ] Human override actions are auditable and not silently destructive to evidence history.

## 18. Reporting
- [ ] The system generates machine-readable reports.
- [ ] The system generates human-readable reports.
- [ ] Reports include total unit counts.
- [ ] Reports include counts by certification status.
- [ ] Reports include average grade and/or score summaries where configured.
- [ ] Reports include expiring-soon units.
- [ ] Reports include highest-risk units.
- [ ] Reports include recurrently failing areas.
- [ ] Reports include trend information over time.
- [ ] Reports can summarize by directory, package, module, or language where useful.
- [ ] Reports can explain why a unit is failing, expired, or high risk.
- [ ] Reports are stored in predictable locations.
- [ ] Reports do not require proprietary viewers to be useful.

## 19. Local CLI Experience
- [ ] The system provides a local CLI.
- [ ] The CLI can initialize certification configuration.
- [ ] The CLI can scan for units.
- [ ] The CLI can certify target scope.
- [ ] The CLI can expire outdated certifications.
- [ ] The CLI can generate reports.
- [ ] The CLI can target a specific file, directory, or changed set.
- [ ] Local CLI behavior respects repository configuration.
- [ ] Local CLI output is usable for developers before pushing changes.
- [ ] Local CLI output is reasonably consistent with CI results for the same scope.

## 20. Multi-Language Adapter Model
- [ ] The architecture includes a defined adapter boundary for language-specific discovery and evidence collection.
- [ ] The architecture includes a defined adapter boundary for analyzer execution and result normalization.
- [ ] Adding a new language adapter does not require changing core certification concepts.
- [ ] The system can support Go-specific analyzers through adapters.
- [ ] The system can support TypeScript/JavaScript-specific analyzers through adapters.
- [ ] The system can support future adapters for Python, Swift, shell, YAML/workflows, and other languages.
- [ ] The system can function with mixed levels of support across languages in the same repository.

## 21. Security and Permissions
- [ ] GitHub workflows declare explicit minimum required permissions.
- [ ] The bootstrap process documents any required GitHub repository settings.
- [ ] The system behaves safely under GitHub token permission constraints.
- [ ] The design accounts for fork PR limitations and untrusted contribution scenarios.
- [ ] The system avoids unsafe write operations in untrusted contexts.
- [ ] The system does not require unnecessary elevated permissions for basic read/evaluate workflows.
- [ ] Security-sensitive paths can be given stricter certification handling if configured.
- [ ] Any use of privileged automation is clearly documented and isolated.

## 22. Storage, Schema, and Data Integrity
- [x] Structured schemas exist for configuration, policies, unit index, certification records, and reports.
- [ ] Persisted artifacts are stable enough for downstream automation consumption.
- [ ] Data files are reviewable in pull requests.
- [ ] The system tolerates partially missing state by rebuilding what is derivable.
- [ ] The system fails honestly when required state is invalid or irrecoverable.
- [ ] The system can validate configuration and policy syntax before execution.
- [ ] The system preserves historical continuity where possible without inventing false lineage.

## 23. Rollout and Adoption
- [ ] The system supports incremental rollout in mature repositories.
- [ ] The system can focus on changed code first.
- [ ] The system can operate in advisory mode before enforcing mode.
- [ ] The system can scope certification to selected directories or languages during rollout.
- [ ] The system can distinguish baseline debt from newly introduced debt.
- [ ] The system can expand toward full-repository governance over time through configuration changes rather than redesign.

## 24. Operational Quality
- [ ] The system is documented well enough for maintainers to install, configure, run, and troubleshoot it.
- [ ] The system includes installation documentation.
- [ ] The system includes policy authoring documentation.
- [ ] The system includes operator/maintainer documentation.
- [ ] The system includes developer-facing usage documentation.
- [ ] The system includes troubleshooting guidance.
- [ ] The system includes example repository setup artifacts.
- [ ] The system includes test coverage for core logic.
- [ ] The system includes integration validation for GitHub workflow behavior.
- [ ] The system produces deterministic behavior wherever deterministic inputs are available.

## 25. Architecture and Code Quality Expectations
- [ ] Core domain logic is free of host-specific concerns where practical.
- [ ] Core domain logic is free of analyzer-specific parsing hacks where practical.
- [ ] No policy logic is hard-coded into orchestration paths that should be configuration-driven.
- [ ] State mutation is intentional and traceable.
- [ ] The solution demonstrates low tolerance for brittle backwards-compatibility hacks, dead code retention, and hidden behavior.
- [ ] The codebase is lint-compliant.
- [ ] The codebase is testable and modular.
- [ ] The codebase favors explicit contracts over ambiguous coupling.
- [ ] The implementation is maintainable enough to support future adapters, policy evolution, and platform growth.

## 26. Minimum v1 Readiness Criteria
- [ ] A repository can be onboarded through a reviewable initialization PR.
- [ ] Repository-local certification configuration is created successfully.
- [ ] Repository-local starter policies are created successfully.
- [ ] An initial unit index is created successfully.
- [ ] Pull request workflows certify changed units successfully.
- [ ] Scheduled workflows expire and recertify units successfully.
- [ ] Certification records are written and updated successfully.
- [ ] Reports are generated successfully in machine-readable and human-readable form.
- [ ] Remediation issues can be created and updated successfully when configured.
- [ ] Go and TypeScript/JavaScript receive at least initial language-aware support.
- [ ] Unsupported languages still receive generic certification treatment.
- [ ] The architecture and interfaces are ready for future language expansion without core redesign.

## 27. Program-Level Success Indicators
- [ ] Teams can identify which parts of a codebase are currently trusted.
- [ ] Teams can identify which parts of a codebase are expired, probationary, or decertified.
- [ ] Teams can tell under which standards a given unit was last certified.
- [ ] Teams can prioritize technical debt through evidence and trust decay rather than intuition alone.
- [ ] Teams can adopt the system incrementally without overwhelming initial noise.
- [ ] The system improves visibility into quality drift over time.
- [ ] The system creates a durable governance model rather than a one-off quality report.