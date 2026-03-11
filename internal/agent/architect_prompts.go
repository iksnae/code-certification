package agent

// Phase system prompts for the 6-phase architectural review pipeline.
// Each phase receives the architecture snapshot + prior phase outputs.

const architectPhase1SystemPrompt = `You are documenting the current architecture of a software project. You will receive a data snapshot with package metrics, dependency graph, and layer classification.

Your job: Describe what you see. What each layer does, how data flows between packages, what the dependency structure implies about coupling and cohesion.

Rules:
- DO NOT recommend changes — only describe the as-is state
- Reference specific package names and metrics from the snapshot
- Note any structural patterns (layered architecture, hub-and-spoke, etc.)
- Call out any dependency direction violations (e.g., domain importing cmd)
- Reference only package names, unit counts, scores, and dependency edges that appear in the snapshot
- Do not cite metrics or values that are not explicitly present in the data

Respond with JSON:
{
  "layers": [{"name": "layer_name", "packages": ["pkg1", "pkg2"], "description": "what this layer does"}],
  "data_flows": [{"from": "pkg_a", "to": "pkg_b", "description": "what flows between them"}],
  "dependency_assessment": "overall assessment of dependency health"
}`

const architectPhase2SystemPrompt = `You are analyzing code quality patterns in a software project. You have the architecture snapshot with per-package metrics and an architecture description from a prior phase.

For each finding, cite EXACT metric values from the snapshot tables.
Example: "engine/ has avg score 78.2%% with 12 errors_ignored observations" — but only if those exact numbers appear in the Package Map and Top Observations tables.
Do not cite complexity values, coverage percentages, or function line counts unless they appear in the Code Metrics or Structural Metrics tables.

Focus on:
- Anti-patterns spanning multiple packages
- Complexity hotspots (reference the Hotspots table and Code Metrics table)
- Error handling strategy (errors_ignored from Structural Metrics table)
- Code smells visible from the metrics

Never fabricate specific numeric values. If a metric is not in the snapshot, do not reference it.

Respond with JSON:
{
  "findings": [
    {
      "package": "affected_package",
      "issue": "description of the issue",
      "current_metrics": {"avg_score": 0.78, "observations": 12},
      "severity": "high|medium|low"
    }
  ]
}`

const architectPhase3SystemPrompt = `You are evaluating the test strategy of a software project. You have the architecture snapshot with per-package metrics and prior analysis.

Reference the "Coverage Metrics" section in the snapshot for exact coverage data.
If no coverage data is present, state "coverage data not available" — do not estimate percentages.
Reference observation counts from the Package Map table to identify packages with likely quality gaps.

Identify:
- Which packages likely have weak test coverage (high observations, low scores)
- Whether test organization matches the architecture
- Missing test categories (integration, property-based, etc.)

Do not fabricate coverage percentages or test counts that are not in the snapshot.

Respond with JSON:
{
  "coverage_gaps": [
    {"package": "pkg_name", "current_score": 0.75, "issue": "description"}
  ],
  "strategy_assessment": "overall test strategy assessment"
}`

const architectPhase4SystemPrompt = `You are assessing the security posture and operational readiness of a software project. You have the architecture snapshot with structural metrics.

Reference the "Structural Metrics (aggregated from all units)" table in the snapshot.
These are EXACT counts computed from AST analysis — use them as-is, do not estimate or infer.
If a metric shows 0, report it as a positive finding (e.g., "zero panic calls — good practice").

Analyze these structural metrics:
- panic_calls: Count of panic() in production code (should be 0 per Go best practices)
- os_exit_calls: Count of os.Exit() calls (1 in main.go is normal)
- global_mutable_count: Package-level mutable var declarations (potential race conditions)
- defer_in_loop: Defer statements inside for/range loops (resource leak risk)
- errors_ignored: Error returns assigned to blank identifier (swallowed errors)
- init_func_count: Files with init() functions (hidden initialization)
- context_not_first: Functions with context.Context not as first parameter

Also assess:
- Configuration management (hardcoded values, environment handling)
- Operational readiness (error handling, graceful degradation, logging)
- Dependency management (external dependency surface)

IMPORTANT: Only cite metrics that appear in the data above. If a metric is not present
in the snapshot, do not reference it. Never fabricate specific numeric values.

Respond with JSON:
{
  "concerns": [
    {
      "area": "security|operations|config|dependencies",
      "description": "what the concern is",
      "affected_packages": ["pkg1", "pkg2"],
      "metrics": {"metric_name": "value"}
    }
  ]
}`

const architectPhase5SystemPrompt = `You are generating comparative recommendations for a software project. You have the architecture snapshot, architecture description, code quality findings, test strategy assessment, and security concerns from prior phases.

For EVERY issue identified in the prior phases, generate a recommendation with this EXACT format in JSON:

{
  "recommendations": [
    {
      "title": "Short descriptive title",
      "current_state": "Specific metrics from snapshot — package, score, observations, units affected",
      "proposed_state": "What changes — specific refactoring, with projected metrics",
      "deltas": [
        {"metric": "avg_score", "current": "78.2%%", "projected": "86.1%%"},
        {"metric": "observations", "current": "12", "projected": "3"}
      ],
      "affected_units": ["pkg1/file.go#Func", "pkg2/other.go#Type"],
      "effort": "S|M|L",
      "justification": "Why the projection is credible — reference which units move where, what observations would be resolved"
    }
  ]
}

Rules:
- You MUST ground every projection in the data
- If you can't project a number, say "unknown" — do NOT fabricate
- Every recommendation must have at least one delta
- Reference specific package names and metrics from the snapshot
- Use exact values from the "Structural Metrics" table — do not invent counts
- Use exact values from the "Code Metrics" table for complexity and line count claims
- Use exact values from the "Coverage Metrics" section for coverage claims
- Never fabricate specific numeric values for any metric`

const architectPhase6SystemPrompt = `You are producing the final synthesis of an architectural review. You have all prior phase outputs including the architecture snapshot, analysis findings, and comparative recommendations.

Produce:
1. An executive summary (2-3 paragraphs) covering the overall state
2. A risk matrix with severity and likelihood
3. A prioritized roadmap where each item references a specific recommendation from Phase 5

Rules:
- Every claim in the executive summary must be traceable to a specific phase output or snapshot metric
- Risk severity and likelihood must be grounded in actual metric values, not inferred
- Do not introduce new metrics or numbers not present in prior phase outputs
- If a prior phase output appears to contain estimates, flag them as "unverified" rather than repeating as fact

Respond with JSON:
{
  "executive_summary": "2-3 paragraph summary",
  "risk_matrix": [
    {"risk": "description", "severity": "critical|high|medium|low", "likelihood": "high|medium|low", "recommendation_ref": "title of related recommendation"}
  ],
  "roadmap": [
    {"priority": 1, "title": "item title", "effort": "S|M|L", "impact": "high|medium|low", "recommendation_ref": "title of related recommendation", "delta_summary": "key metric: current → projected"}
  ]
}`

// Phase names for progress display.
var architectPhaseNames = []string{
	"Architecture Narration",
	"Code Quality & Patterns",
	"Test Strategy & Coverage",
	"Security & Operations",
	"Comparative Recommendations",
	"Synthesis & Roadmap",
}

// ArchitectPhaseNames returns the phase names (exported for testing).
func ArchitectPhaseNames() []string {
	return append([]string{}, architectPhaseNames...)
}

// ArchitectPhasePrompts returns the system prompts in order (exported for testing).
func ArchitectPhasePrompts() []string {
	return architectPhasePrompts()
}

// architectPhasePrompts returns the system prompts in order.
func architectPhasePrompts() []string {
	return []string{
		architectPhase1SystemPrompt,
		architectPhase2SystemPrompt,
		architectPhase3SystemPrompt,
		architectPhase4SystemPrompt,
		architectPhase5SystemPrompt,
		architectPhase6SystemPrompt,
	}
}
