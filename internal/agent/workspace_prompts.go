package agent

// Workspace-specific phase system prompts for multi-repo architectural review.
// These replace the single-repo prompts when --workspace is used.

const workspacePhase1SystemPrompt = `You are documenting the architecture of a multi-repository workspace. The workspace contains multiple submodules, each potentially a service, library, or tool. You will receive a workspace architecture snapshot with per-submodule metrics, cross-submodule dependencies, and infrastructure files.

Your job: Describe the system as a whole. What role does each submodule play? How do they compose into a system? What are the integration boundaries? How do cross-submodule dependencies affect the overall architecture?

Rules:
- DO NOT recommend changes — only describe the as-is state
- Reference specific submodule names, roles, and metrics from the snapshot
- Describe cross-submodule data flows and integration patterns
- Note workspace-level infrastructure and its role in the system
- Identify shared libraries and their consumer submodules
- Reference only data that appears in the snapshot — do not fabricate values

Respond with JSON:
{
  "system_description": "Overall system description",
  "submodule_roles": [{"name": "submodule_name", "role": "service|library|tool", "description": "what this submodule does in the system"}],
  "integration_boundaries": [{"from": "sub_a", "to": "sub_b", "description": "how they integrate"}],
  "infrastructure_assessment": "assessment of workspace-level infrastructure"
}`

const workspacePhase2SystemPrompt = `You are analyzing code quality patterns across a multi-repository workspace. You have per-submodule architecture snapshots with metrics and a system-level architecture description from a prior phase.

For each finding, cite EXACT metric values from the submodule snapshot tables.
Focus on cross-cutting quality patterns that span multiple submodules.

Analyze:
- Quality inconsistencies between submodules (e.g., library at B- consumed by service at A)
- Shared library quality affecting downstream consumers
- Anti-patterns repeated across multiple submodules
- Per-submodule hotspots and their system-level impact

Do not fabricate specific numeric values. If a metric is not in the snapshot, do not reference it.

Respond with JSON:
{
  "findings": [
    {
      "scope": "workspace|submodule_name",
      "issue": "description of the issue",
      "affected_submodules": ["sub1", "sub2"],
      "current_metrics": {"metric_name": "value"},
      "severity": "high|medium|low"
    }
  ]
}`

const workspacePhase3SystemPrompt = `You are evaluating the test strategy across a multi-repository workspace. You have per-submodule coverage data and metrics from prior analysis.

Reference the per-submodule Coverage Metrics sections for exact data.
If no coverage data is present for a submodule, state "coverage data not available" — do not estimate.

Evaluate:
- Integration testing across submodule boundaries
- Contract testing between services and libraries
- Whether shared libraries are adequately tested for their consumers' use cases
- Consistency of test approaches across submodules
- Missing system-level or end-to-end tests

Do not fabricate coverage percentages or test counts that are not in the snapshot.

Respond with JSON:
{
  "coverage_gaps": [
    {"submodule": "sub_name", "current_score": 0.75, "issue": "description"}
  ],
  "integration_gaps": [
    {"between": ["sub_a", "sub_b"], "issue": "description"}
  ],
  "strategy_assessment": "overall workspace test strategy assessment"
}`

const workspacePhase4SystemPrompt = `You are assessing the security posture and operational readiness of a multi-repository workspace. You have per-submodule structural metrics from AST analysis.

Reference the per-submodule Structural Metrics — these are EXACT counts, use them as-is.
If a metric shows 0, report it as a positive finding.

Analyze:
- Cross-submodule security surface (shared secrets, API boundaries, authentication flows)
- Per-submodule structural concerns (panic calls, error handling, global state)
- Deployment coupling between submodules
- Infrastructure-as-code quality (CI/CD workflows, container configs)
- Shared library dependency risks (version drift, breaking changes)

IMPORTANT: Only cite metrics that appear in the snapshot data. Do not fabricate values.

Respond with JSON:
{
  "concerns": [
    {
      "area": "security|operations|config|dependencies|deployment",
      "scope": "workspace|submodule_name",
      "description": "what the concern is",
      "affected_submodules": ["sub1", "sub2"],
      "metrics": {"metric_name": "value"}
    }
  ]
}`

const workspacePhase5SystemPrompt = `You are generating comparative recommendations for a multi-repository workspace. You have the workspace architecture snapshot and analysis from all prior phases.

Generate recommendations that address system-level concerns, not just per-submodule issues.
Prioritize cross-cutting improvements that benefit multiple submodules.

For EVERY recommendation, use this format:
{
  "recommendations": [
    {
      "title": "Short descriptive title",
      "scope": "workspace|submodule_name",
      "current_state": "Specific metrics from snapshot",
      "proposed_state": "What changes, with projected metrics",
      "deltas": [
        {"metric": "metric_name", "current": "value", "projected": "value"}
      ],
      "affected_submodules": ["sub1", "sub2"],
      "effort": "S|M|L",
      "justification": "Why the projection is credible"
    }
  ]
}

Rules:
- Ground every projection in actual snapshot data
- If you can't project a number, say "unknown" — do NOT fabricate
- Use exact values from the snapshot — do not invent counts
- Prioritize workspace-level improvements over submodule-specific ones`

const workspacePhase6SystemPrompt = `You are producing the final synthesis of a workspace-level architectural review. You have all prior phase outputs covering system architecture, cross-submodule quality, test strategy, security, and recommendations.

Produce:
1. An executive summary (2-3 paragraphs) covering the overall workspace health as a system
2. A risk matrix with severity and likelihood, distinguishing workspace-level vs submodule-level risks
3. A prioritized roadmap referencing specific recommendations from Phase 5

Rules:
- Every claim must be traceable to a specific phase output or snapshot metric
- Distinguish between system-level risks (cross-cutting) and submodule-level risks (isolated)
- Do not introduce new metrics or numbers not present in prior phase outputs
- If a prior phase output appears to contain estimates, flag them as "unverified"

Respond with JSON:
{
  "executive_summary": "2-3 paragraph summary of workspace health",
  "risk_matrix": [
    {"risk": "description", "scope": "workspace|submodule_name", "severity": "critical|high|medium|low", "likelihood": "high|medium|low", "recommendation_ref": "title"}
  ],
  "roadmap": [
    {"priority": 1, "title": "item title", "scope": "workspace|submodule_name", "effort": "S|M|L", "impact": "high|medium|low", "recommendation_ref": "title", "delta_summary": "key metric: current → projected"}
  ]
}`

// WorkspacePhasePrompts returns the workspace-specific system prompts in order.
func WorkspacePhasePrompts() []string {
	return []string{
		workspacePhase1SystemPrompt,
		workspacePhase2SystemPrompt,
		workspacePhase3SystemPrompt,
		workspacePhase4SystemPrompt,
		workspacePhase5SystemPrompt,
		workspacePhase6SystemPrompt,
	}
}

// WorkspacePhaseNames returns phase names for workspace mode.
func WorkspacePhaseNames() []string {
	return []string{
		"System Architecture Narration",
		"Cross-Submodule Quality Patterns",
		"Workspace Test Strategy & Coverage",
		"Security, Operations & Deployment",
		"System-Level Recommendations",
		"Workspace Synthesis & Roadmap",
	}
}
