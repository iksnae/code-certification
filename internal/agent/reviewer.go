package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

// ReviewInput holds everything needed for an agent review.
type ReviewInput struct {
	Unit       domain.Unit
	SourceCode string
	Evidence   []domain.Evidence
}

// ReviewResult holds the outcome of an agent review.
type ReviewResult struct {
	Reviewed     bool               `json:"reviewed"`
	Prescreened  bool               `json:"prescreened"`
	ReviewOutput string             `json:"review_output,omitempty"`
	Scores       map[string]float64 `json:"scores,omitempty"`
	Status       string             `json:"status,omitempty"`
	Actions      []string           `json:"actions,omitempty"`
	Remediation  []RemediationStep  `json:"remediation,omitempty"`
	Confidence   float64            `json:"confidence"`
	TokensUsed   int                `json:"tokens_used"`
	ModelsUsed   []string           `json:"models_used,omitempty"`
}

// ToEvidence converts the review result to a domain.Evidence.
// Model attribution is embedded in the Source field.
func (r ReviewResult) ToEvidence() domain.Evidence {
	source := "agent"
	if len(r.ModelsUsed) > 0 {
		source = "agent:" + joinModels(r.ModelsUsed)
	}

	summary := fmt.Sprintf("Agent review: %s (confidence: %.0f%%)", r.Status, r.Confidence*100)
	if len(r.ModelsUsed) > 0 {
		summary += fmt.Sprintf(" [models: %s]", joinModels(r.ModelsUsed))
	}

	return domain.Evidence{
		Kind:       domain.EvidenceKindAgentReview,
		Source:     source,
		Passed:     r.Status != "decertified",
		Summary:    summary,
		Details:    r,
		Timestamp:  time.Now(),
		Confidence: r.Confidence,
	}
}

// ToPrescreenEvidence converts a prescreen-only result to evidence.
// Used when AI evaluated but determined no detailed review was needed.
func (r ReviewResult) ToPrescreenEvidence() domain.Evidence {
	source := "agent-prescreen"
	if len(r.ModelsUsed) > 0 {
		source = "agent-prescreen:" + joinModels(r.ModelsUsed)
	}

	summary := fmt.Sprintf("AI prescreen: no issues found (confidence: %.0f%%)", r.Confidence*100)
	if len(r.ModelsUsed) > 0 {
		summary += fmt.Sprintf(" [model: %s]", joinModels(r.ModelsUsed))
	}

	return domain.Evidence{
		Kind:       domain.EvidenceKindAgentReview,
		Source:     source,
		Passed:     true,
		Summary:    summary,
		Details:    r,
		Timestamp:  time.Now(),
		Confidence: r.Confidence,
	}
}

func joinModels(models []string) string {
	// Deduplicate while preserving order
	seen := make(map[string]bool)
	var unique []string
	for _, m := range models {
		if !seen[m] {
			seen[m] = true
			unique = append(unique, m)
		}
	}
	result := ""
	for i, m := range unique {
		if i > 0 {
			result += ","
		}
		result += m
	}
	return result
}

// Reviewer orchestrates the 5-step agent review pipeline.
type Reviewer struct {
	provider Provider
	router   *Router
}

// NewReviewer creates a new agent reviewer.
func NewReviewer(provider Provider, router *Router) *Reviewer {
	return &Reviewer{provider: provider, router: router}
}

// Review runs the full agent review pipeline.
func (rv *Reviewer) Review(ctx context.Context, input ReviewInput) (ReviewResult, error) {
	if rv.provider == nil || rv.router == nil {
		return ReviewResult{}, nil
	}

	var result ReviewResult
	var totalTokens int

	// Step 1: Prescreen
	prescreenModel := rv.router.ModelFor(TaskPrescreen)
	if prescreenModel == "" {
		return ReviewResult{}, nil
	}

	resp, err := rv.provider.Chat(ctx, ChatRequest{
		Model: prescreenModel,
		Messages: []Message{
			{Role: "user", Content: fmt.Sprintf("Prescreen unit %s. Evidence count: %d", input.Unit.ID, len(input.Evidence))},
		},
		Temperature: 0.1,
		MaxTokens:   256,
	})
	if err != nil {
		// Graceful degradation: skip agent review
		return ReviewResult{}, nil
	}
	totalTokens += resp.Usage.TotalTokens

	var prescreen PrescreenResponse
	if err := json.Unmarshal([]byte(resp.Content()), &prescreen); err != nil {
		// Can't parse prescreen, skip review
		return ReviewResult{}, nil
	}

	if !prescreen.NeedsReview {
		return ReviewResult{Reviewed: false, Confidence: prescreen.Confidence, TokensUsed: totalTokens}, nil
	}

	// Step 2: Review
	reviewModel := rv.router.ModelFor(TaskReview)
	if reviewModel != "" {
		resp, err = rv.provider.Chat(ctx, ChatRequest{
			Model: reviewModel,
			Messages: []Message{
				{Role: "user", Content: fmt.Sprintf("Review code unit %s:\n%s", input.Unit.ID, input.SourceCode)},
			},
			Temperature: 0.3,
			MaxTokens:   2048,
		})
		if err == nil {
			result.ReviewOutput = resp.Content()
			totalTokens += resp.Usage.TotalTokens
		}
	}

	// Step 3: Scoring
	scoringModel := rv.router.ModelFor(TaskScoring)
	if scoringModel != "" {
		resp, err = rv.provider.Chat(ctx, ChatRequest{
			Model: scoringModel,
			Messages: []Message{
				{Role: "user", Content: fmt.Sprintf("Score unit %s. Review: %s", input.Unit.ID, result.ReviewOutput)},
			},
			Temperature: 0.1,
			MaxTokens:   512,
		})
		if err == nil {
			var scoring ScoringResponse
			if json.Unmarshal([]byte(resp.Content()), &scoring) == nil {
				result.Scores = scoring.Scores
				result.Confidence = scoring.Confidence
			}
			totalTokens += resp.Usage.TotalTokens
		}
	}

	// Step 4: Decision
	decisionModel := rv.router.ModelFor(TaskDecision)
	if decisionModel != "" {
		resp, err = rv.provider.Chat(ctx, ChatRequest{
			Model: decisionModel,
			Messages: []Message{
				{Role: "user", Content: fmt.Sprintf("Decide certification for %s. Scores: %v", input.Unit.ID, result.Scores)},
			},
			Temperature: 0.1,
			MaxTokens:   512,
		})
		if err == nil {
			var decision DecisionResponse
			if json.Unmarshal([]byte(resp.Content()), &decision) == nil {
				result.Status = decision.Status
				result.Actions = decision.Actions
			}
			totalTokens += resp.Usage.TotalTokens
		}
	}

	// Step 5: Remediation (only for non-certified)
	if result.Status != "certified" && result.Status != "" {
		remModel := rv.router.ModelFor(TaskRemediation)
		if remModel != "" {
			resp, err = rv.provider.Chat(ctx, ChatRequest{
				Model: remModel,
				Messages: []Message{
					{Role: "user", Content: fmt.Sprintf("Generate remediation for %s. Status: %s", input.Unit.ID, result.Status)},
				},
				Temperature: 0.3,
				MaxTokens:   1024,
			})
			if err == nil {
				var rem RemediationResponse
				if json.Unmarshal([]byte(resp.Content()), &rem) == nil {
					result.Remediation = rem.Steps
				}
				totalTokens += resp.Usage.TotalTokens
			}
		}
	}

	result.Reviewed = true
	result.TokensUsed = totalTokens
	return result, nil
}
