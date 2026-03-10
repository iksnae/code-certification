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
	Reviewed        bool               `json:"reviewed"`
	Prescreened     bool               `json:"prescreened"`
	ReviewOutput    string             `json:"review_output,omitempty"`
	Scores          map[string]float64 `json:"scores,omitempty"`
	Status          string             `json:"status,omitempty"`
	Actions         []string           `json:"actions,omitempty"`
	Remediation     []RemediationStep  `json:"remediation,omitempty"`
	Confidence      float64            `json:"confidence"`
	TokensUsed      int                `json:"tokens_used"`
	ModelsUsed      []string           `json:"models_used,omitempty"`
	PrescreenReason string             `json:"prescreen_reason,omitempty"`
	Suggestions     []string           `json:"suggestions,omitempty"`
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

	metrics := make(map[string]float64, len(r.Scores)+2)
	for k, v := range r.Scores {
		metrics[k] = v
	}
	metrics["confidence"] = r.Confidence
	metrics["tokens_used"] = float64(r.TokensUsed)

	return domain.Evidence{
		Kind:       domain.EvidenceKindAgentReview,
		Source:     source,
		Passed:     r.Status != "decertified",
		Summary:    summary,
		Metrics:    metrics,
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
	if r.PrescreenReason != "" {
		summary = fmt.Sprintf("AI: %s (confidence: %.0f%%)", r.PrescreenReason, r.Confidence*100)
	}
	if len(r.ModelsUsed) > 0 {
		summary += fmt.Sprintf(" [model: %s]", joinModels(r.ModelsUsed))
	}

	return domain.Evidence{
		Kind:    domain.EvidenceKindAgentReview,
		Source:  source,
		Passed:  true,
		Summary: summary,
		Metrics: map[string]float64{
			"confidence": r.Confidence,
		},
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

	// Step 1: Prescreen — decide if full review is needed
	prescreen, tokens, err := rv.runPrescreen(ctx, input)
	if err != nil || !prescreen.NeedsReview {
		return ReviewResult{Reviewed: false, Confidence: prescreen.Confidence, TokensUsed: tokens}, nil
	}

	// Steps 2-5: Full review pipeline
	result := ReviewResult{Reviewed: true, TokensUsed: tokens}
	rv.runCodeReview(ctx, input, &result)
	rv.runScoring(ctx, input, &result)
	rv.runDecision(ctx, input, &result)
	rv.runRemediation(ctx, input, &result)

	return result, nil
}

func (rv *Reviewer) runPrescreen(ctx context.Context, input ReviewInput) (PrescreenResponse, int, error) {
	model := rv.router.ModelFor(TaskPrescreen)
	if model == "" {
		return PrescreenResponse{}, 0, fmt.Errorf("no prescreen model")
	}
	resp, err := rv.provider.Chat(ctx, ChatRequest{
		Model:       model,
		Messages:    []Message{{Role: "user", Content: fmt.Sprintf("Prescreen unit %s. Evidence count: %d", input.Unit.ID, len(input.Evidence))}},
		Temperature: 0.1,
		MaxTokens:   256,
	})
	if err != nil {
		return PrescreenResponse{}, 0, err
	}
	var prescreen PrescreenResponse
	if err := json.Unmarshal([]byte(resp.Content()), &prescreen); err != nil {
		return PrescreenResponse{}, resp.Usage.TotalTokens, err
	}
	return prescreen, resp.Usage.TotalTokens, nil
}

func (rv *Reviewer) runCodeReview(ctx context.Context, input ReviewInput, result *ReviewResult) {
	model := rv.router.ModelFor(TaskReview)
	if model == "" {
		return
	}
	resp, err := rv.provider.Chat(ctx, ChatRequest{
		Model:       model,
		Messages:    []Message{{Role: "user", Content: fmt.Sprintf("Review code unit %s:\n%s", input.Unit.ID, input.SourceCode)}},
		Temperature: 0.3,
		MaxTokens:   2048,
	})
	if err == nil {
		result.ReviewOutput = resp.Content()
		result.TokensUsed += resp.Usage.TotalTokens
	}
}

func (rv *Reviewer) runScoring(ctx context.Context, input ReviewInput, result *ReviewResult) {
	model := rv.router.ModelFor(TaskScoring)
	if model == "" {
		return
	}
	resp, err := rv.provider.Chat(ctx, ChatRequest{
		Model:       model,
		Messages:    []Message{{Role: "user", Content: fmt.Sprintf("Score unit %s. Review: %s", input.Unit.ID, result.ReviewOutput)}},
		Temperature: 0.1,
		MaxTokens:   512,
	})
	if err == nil {
		var scoring ScoringResponse
		if json.Unmarshal([]byte(resp.Content()), &scoring) == nil {
			result.Scores = scoring.Scores
			result.Confidence = scoring.Confidence
		}
		result.TokensUsed += resp.Usage.TotalTokens
	}
}

func (rv *Reviewer) runDecision(ctx context.Context, input ReviewInput, result *ReviewResult) {
	model := rv.router.ModelFor(TaskDecision)
	if model == "" {
		return
	}
	resp, err := rv.provider.Chat(ctx, ChatRequest{
		Model:       model,
		Messages:    []Message{{Role: "user", Content: fmt.Sprintf("Decide certification for %s. Scores: %v", input.Unit.ID, result.Scores)}},
		Temperature: 0.1,
		MaxTokens:   512,
	})
	if err == nil {
		var decision DecisionResponse
		if json.Unmarshal([]byte(resp.Content()), &decision) == nil {
			result.Status = decision.Status
			result.Actions = decision.Actions
		}
		result.TokensUsed += resp.Usage.TotalTokens
	}
}

func (rv *Reviewer) runRemediation(ctx context.Context, input ReviewInput, result *ReviewResult) {
	if result.Status == "certified" || result.Status == "" {
		return
	}
	model := rv.router.ModelFor(TaskRemediation)
	if model == "" {
		return
	}
	resp, err := rv.provider.Chat(ctx, ChatRequest{
		Model:       model,
		Messages:    []Message{{Role: "user", Content: fmt.Sprintf("Generate remediation for %s. Status: %s", input.Unit.ID, result.Status)}},
		Temperature: 0.3,
		MaxTokens:   1024,
	})
	if err == nil {
		var rem RemediationResponse
		if json.Unmarshal([]byte(resp.Content()), &rem) == nil {
			result.Remediation = rem.Steps
		}
		result.TokensUsed += resp.Usage.TotalTokens
	}
}
