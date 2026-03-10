package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// ArchitectReviewer orchestrates the 6-phase architectural review pipeline.
type ArchitectReviewer struct {
	Provider     Provider
	Model        string
	Verbose      bool
	OnPhaseStart func(phase int, name string)
	OnPhaseDone  func(phase int, name string, tokens int)
}

// ArchitectResult holds the output of a complete architectural review.
type ArchitectResult struct {
	Snapshot       *ArchSnapshot
	Phase1         *ArchPhase1Result
	Phase2         *ArchPhase2Result
	Phase3         *ArchPhase3Result
	Phase4         *ArchPhase4Result
	Phase5         *ArchPhase5Result
	Phase6         *ArchPhase6Result
	RawOutputs     []string // raw LLM responses per phase
	Thinking       []string // chain-of-thought reasoning per phase (from <think> tags)
	TotalTokens    int
	Duration       time.Duration
	Model          string
	PhasesComplete int
	Errors         []string // errors per failed phase
}

// Phase response types

// ArchPhase1Result describes the as-is architecture.
type ArchPhase1Result struct {
	Layers               []ArchLayer    `json:"layers"`
	DataFlows            []ArchDataFlow `json:"data_flows"`
	DependencyAssessment string         `json:"dependency_assessment"`
}

type ArchLayer struct {
	Name        string   `json:"name"`
	Packages    []string `json:"packages"`
	Description string   `json:"description"`
}

type ArchDataFlow struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Description string `json:"description"`
}

// ArchPhase2Result holds code quality findings.
type ArchPhase2Result struct {
	Findings []ArchFinding `json:"findings"`
}

type ArchFinding struct {
	Package        string             `json:"package"`
	Issue          string             `json:"issue"`
	CurrentMetrics map[string]float64 `json:"current_metrics"`
	Severity       string             `json:"severity"`
}

// ArchPhase3Result holds test strategy assessment.
type ArchPhase3Result struct {
	CoverageGaps       []ArchCoverageGap `json:"coverage_gaps"`
	StrategyAssessment string            `json:"strategy_assessment"`
}

type ArchCoverageGap struct {
	Package      string  `json:"package"`
	CurrentScore float64 `json:"current_score"`
	Issue        string  `json:"issue"`
}

// ArchPhase4Result holds security/ops concerns.
type ArchPhase4Result struct {
	Concerns []ArchConcern `json:"concerns"`
}

type ArchConcern struct {
	Area             string         `json:"area"`
	Description      string         `json:"description"`
	AffectedPackages []string       `json:"affected_packages"`
	Metrics          map[string]any `json:"metrics"`
}

// ArchPhase5Result holds comparative recommendations.
type ArchPhase5Result struct {
	Recommendations []ArchRecommendation `json:"recommendations"`
}

type ArchRecommendation struct {
	Title         string      `json:"title"`
	CurrentState  string      `json:"current_state"`
	ProposedState string      `json:"proposed_state"`
	Deltas        []ArchDelta `json:"deltas"`
	AffectedUnits []string    `json:"affected_units"`
	Effort        string      `json:"effort"`
	Justification string      `json:"justification"`
}

type ArchDelta struct {
	Metric    string `json:"metric"`
	Current   string `json:"current"`
	Projected string `json:"projected"`
}

// ArchPhase6Result holds the final synthesis.
type ArchPhase6Result struct {
	ExecutiveSummary string         `json:"executive_summary"`
	RiskMatrix       []ArchRisk     `json:"risk_matrix"`
	Roadmap          []ArchRoadItem `json:"roadmap"`
}

type ArchRisk struct {
	Risk              string `json:"risk"`
	Severity          string `json:"severity"`
	Likelihood        string `json:"likelihood"`
	RecommendationRef string `json:"recommendation_ref"`
}

type ArchRoadItem struct {
	Priority          int    `json:"priority"`
	Title             string `json:"title"`
	Effort            string `json:"effort"`
	Impact            string `json:"impact"`
	RecommendationRef string `json:"recommendation_ref"`
	DeltaSummary      string `json:"delta_summary"`
}

// Review runs the 6-phase architectural review pipeline.
// Phases can be limited with the phases parameter (nil = all, else list of 1-indexed phase numbers).
func (ar *ArchitectReviewer) Review(ctx context.Context, pc *ProjectContext, phases []int) (*ArchitectResult, error) {
	if ar.Provider == nil {
		return nil, fmt.Errorf("no AI provider available")
	}

	start := time.Now()
	result := &ArchitectResult{
		Snapshot:   pc.Snapshot,
		Model:      ar.Model,
		RawOutputs: make([]string, 6),
		Thinking:   make([]string, 6),
	}

	// Build the shared context that all phases receive
	contextBlock := pc.FormatForLLM(4000)
	prompts := architectPhasePrompts()

	// Determine which phases to run
	runPhase := make(map[int]bool)
	if len(phases) == 0 {
		for i := 1; i <= 6; i++ {
			runPhase[i] = true
		}
	} else {
		for _, p := range phases {
			if p >= 1 && p <= 6 {
				runPhase[p] = true
			}
		}
	}

	// Track prior outputs for feed-forward
	var priorOutputs []string

	for i := 0; i < 6; i++ {
		phaseNum := i + 1
		phaseName := architectPhaseNames[i]

		if !runPhase[phaseNum] {
			priorOutputs = append(priorOutputs, "")
			continue
		}

		if ar.OnPhaseStart != nil {
			ar.OnPhaseStart(phaseNum, phaseName)
		}

		// Build user prompt with context + prior outputs
		userPrompt := buildArchitectUserPrompt(contextBlock, priorOutputs, phaseNum)

		// Generous token limits — architect phases produce detailed analysis
		maxTokens := 8192
		if phaseNum == 5 {
			maxTokens = 12288 // Phase 5 (comparative recommendations) needs the most room
		}

		resp, err := ar.Provider.Chat(ctx, ChatRequest{
			Model: ar.Model,
			Messages: []Message{
				{Role: "system", Content: prompts[i]},
				{Role: "user", Content: userPrompt},
			},
			Temperature: 0.3,
			MaxTokens:   maxTokens,
		})
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("Phase %d (%s): %v", phaseNum, phaseName, err))
			priorOutputs = append(priorOutputs, "")
			continue
		}

		content := resp.Content()
		result.RawOutputs[i] = content
		result.TotalTokens += resp.Usage.TotalTokens
		result.PhasesComplete++

		// Capture chain-of-thought reasoning before stripping
		result.Thinking[i] = extractThinking(content)

		// Strip <think>...</think> blocks (qwen3 chain-of-thought) before JSON extraction
		cleaned := stripThinkTags(content)

		// Parse structured response
		jsonStr := extractJSON(cleaned)
		ar.parsePhaseResult(result, phaseNum, jsonStr, content)

		priorOutputs = append(priorOutputs, content)

		if ar.OnPhaseDone != nil {
			ar.OnPhaseDone(phaseNum, phaseName, resp.Usage.TotalTokens)
		}
	}

	result.Duration = time.Since(start)

	// Validate Phase 5 recommendations
	if result.Phase5 != nil {
		for i := range result.Phase5.Recommendations {
			rec := &result.Phase5.Recommendations[i]
			if len(rec.Deltas) == 0 {
				rec.Deltas = []ArchDelta{{Metric: "unknown", Current: "unknown", Projected: "unknown"}}
			}
		}
	}

	return result, nil
}

// parsePhaseResult parses the LLM response into the appropriate phase result type.
func (ar *ArchitectReviewer) parsePhaseResult(result *ArchitectResult, phase int, jsonStr, raw string) {
	switch phase {
	case 1:
		var p ArchPhase1Result
		if json.Unmarshal([]byte(jsonStr), &p) == nil {
			result.Phase1 = &p
		} else {
			// Fallback: store raw text as dependency assessment
			result.Phase1 = &ArchPhase1Result{DependencyAssessment: raw}
		}
	case 2:
		var p ArchPhase2Result
		if json.Unmarshal([]byte(jsonStr), &p) == nil {
			result.Phase2 = &p
		} else {
			result.Phase2 = &ArchPhase2Result{
				Findings: []ArchFinding{{Issue: raw, Severity: "medium"}},
			}
		}
	case 3:
		var p ArchPhase3Result
		if json.Unmarshal([]byte(jsonStr), &p) == nil {
			result.Phase3 = &p
		} else {
			result.Phase3 = &ArchPhase3Result{StrategyAssessment: raw}
		}
	case 4:
		var p ArchPhase4Result
		if json.Unmarshal([]byte(jsonStr), &p) == nil {
			result.Phase4 = &p
		} else {
			result.Phase4 = &ArchPhase4Result{
				Concerns: []ArchConcern{{Description: raw, Area: "general"}},
			}
		}
	case 5:
		var p ArchPhase5Result
		if json.Unmarshal([]byte(jsonStr), &p) == nil {
			result.Phase5 = &p
		}
	case 6:
		var p ArchPhase6Result
		if json.Unmarshal([]byte(jsonStr), &p) == nil {
			result.Phase6 = &p
		} else {
			result.Phase6 = &ArchPhase6Result{ExecutiveSummary: raw}
		}
	}
}

// thinkTagRegex matches <think>...</think> blocks from chain-of-thought models.
// Qwen3 and similar models wrap reasoning in these tags.
var thinkTagRegex = regexp.MustCompile(`(?s)<think>(.*?)</think>`)

// extractThinking pulls all <think> block content from model output.
// Returns concatenated thinking text, or empty string if none found.
func extractThinking(s string) string {
	matches := thinkTagRegex.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return ""
	}
	var parts []string
	for _, m := range matches {
		if len(m) > 1 {
			t := strings.TrimSpace(m[1])
			if t != "" {
				parts = append(parts, t)
			}
		}
	}
	return strings.Join(parts, "\n\n")
}

// stripThinkTags removes <think>...</think> blocks from model output
// so JSON extraction doesn't pick up braces from reasoning text.
func stripThinkTags(s string) string {
	return strings.TrimSpace(thinkTagRegex.ReplaceAllString(s, ""))
}

// buildArchitectUserPrompt builds the user message for a phase.
func buildArchitectUserPrompt(contextBlock string, priorOutputs []string, phaseNum int) string {
	var b strings.Builder

	b.WriteString("## Project Data\n\n")
	b.WriteString(contextBlock)
	b.WriteString("\n\n")

	// Include prior phase outputs
	if len(priorOutputs) > 0 {
		hasPrior := false
		for i, output := range priorOutputs {
			if output == "" {
				continue
			}
			if !hasPrior {
				b.WriteString("## Prior Phase Results\n\n")
				hasPrior = true
			}
			fmt.Fprintf(&b, "### Phase %d: %s\n%s\n\n", i+1, architectPhaseNames[i], output)
		}
	}

	fmt.Fprintf(&b, "## Your Task\nPerform Phase %d: %s\n", phaseNum, architectPhaseNames[phaseNum-1])
	return b.String()
}
