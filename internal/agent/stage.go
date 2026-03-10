package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// StageInput holds all data flowing through the pipeline.
type StageInput struct {
	Unit            domain.Unit
	SourceCode      string
	Evidence        []domain.Evidence
	EvidenceSummary string
	ReviewOutput    string // set by ReviewStage, read by ScoringStage
}

// StageResult accumulates output from pipeline stages.
type StageResult struct {
	Reviewed        bool
	ReviewOutput    string
	Scores          map[string]float64
	Status          string
	Actions         []string
	Remediation     []RemediationStep
	Confidence      float64
	TokensUsed      int
	ModelsUsed      []string // attribution: which models contributed
	PrescreenReason string   // brief quality assessment from prescreen
	Suggestions     []string // actionable improvement suggestions
}

// Stage is a single step in the review pipeline.
// Returns (result, shouldContinue, error).
type Stage interface {
	Execute(ctx context.Context, input StageInput) (StageResult, bool, error)
	Name() string
}

// --- Prescreen Stage ---

type prescreenStage struct {
	provider Provider
	model    string
}

// NewPrescreenStage creates the prescreen gate stage.
func NewPrescreenStage(provider Provider, model string) Stage {
	return &prescreenStage{provider: provider, model: model}
}

func (s *prescreenStage) Name() string { return "prescreen" }

func (s *prescreenStage) Execute(ctx context.Context, input StageInput) (StageResult, bool, error) {
	resp, err := s.provider.Chat(ctx, ChatRequest{
		Model: s.model,
		Messages: AdaptiveMessages(
			`You are a code quality prescreen filter. Respond with ONLY JSON:
{"needs_review": true/false, "reason": "one-line quality assessment", "confidence": 0.0-1.0, "suggestions": ["actionable suggestion 1", "suggestion 2"]}
The "reason" should always describe the code quality briefly. "suggestions" lists 0-3 brief improvement ideas (even for passing code). Units that pass all checks do NOT need detailed review.`,
			fmt.Sprintf("Unit: %s (%s)\nEvidence: %s\n\nAssess this unit's quality and determine if it needs detailed review.",
				input.Unit.ID, input.Unit.Type, input.EvidenceSummary),
			false, // user-only for maximum compatibility
		),
		Temperature: 0.1,
		MaxTokens:   512,
	})
	if err != nil {
		return StageResult{}, false, err
	}

	model := resp.Model
	if model == "" {
		model = s.model
	}
	result := StageResult{TokensUsed: resp.Usage.TotalTokens, ModelsUsed: []string{model}}
	content := resp.Content()

	// Try strict JSON parse first
	var prescreen PrescreenResponse
	if err := json.Unmarshal([]byte(extractJSON(content)), &prescreen); err == nil {
		result.Confidence = prescreen.Confidence
		result.PrescreenReason = prescreen.Reason
		result.Suggestions = prescreen.Suggestions
		return result, prescreen.NeedsReview, nil
	}

	// Loose parse: look for signals in natural language response
	lower := strings.ToLower(content)
	needsReview := looseParseNeedsReview(lower)
	result.Confidence = 0.5 // lower confidence for loose parse
	result.PrescreenReason = "could not parse structured response"
	return result, needsReview, nil
}

// looseParseNeedsReview extracts review intent from free-text responses.
func looseParseNeedsReview(text string) bool {
	// Negative signals — no review needed
	noSignals := []string{
		"needs_review\": false", "needs_review\":false",
		"no review needed", "does not need", "doesn't need",
		"no detailed review", "looks clean", "looks fine",
		"all checks pass", "no issues",
	}
	for _, sig := range noSignals {
		if strings.Contains(text, sig) {
			return false
		}
	}

	// Positive signals — review needed
	yesSignals := []string{
		"needs_review\": true", "needs_review\":true",
		"needs review", "should be reviewed", "recommend review",
		"issues", "concerns", "borderline", "problems",
	}
	for _, sig := range yesSignals {
		if strings.Contains(text, sig) {
			return true
		}
	}

	// Ambiguous → default to review (safer)
	return true
}

// --- Review Stage ---

type reviewStage struct {
	provider Provider
	model    string
}

// NewReviewStage creates the detailed code review stage.
func NewReviewStage(provider Provider, model string) Stage {
	return &reviewStage{provider: provider, model: model}
}

func (s *reviewStage) Name() string { return "review" }

func (s *reviewStage) Execute(ctx context.Context, input StageInput) (StageResult, bool, error) {
	codeSnippet := input.SourceCode
	if len(codeSnippet) > 4000 {
		codeSnippet = codeSnippet[:4000] + "\n... (truncated)"
	}

	resp, err := s.provider.Chat(ctx, ChatRequest{
		Model: s.model,
		Messages: AdaptiveMessages(
			"You are a senior code reviewer. Provide specific, actionable feedback on correctness, maintainability, readability, testability, and security.",
			fmt.Sprintf("Review this code unit:\n\nUnit: %s\nLanguage: %s\n\n```\n%s\n```\n\nEvidence: %s",
				input.Unit.ID, input.Unit.ID.Language(), codeSnippet, input.EvidenceSummary),
			false,
		),
		Temperature: 0.3,
		MaxTokens:   1024,
	})
	if err != nil {
		return StageResult{}, false, err
	}

	model := resp.Model
	if model == "" {
		model = s.model
	}
	return StageResult{
		ReviewOutput: resp.Content(),
		TokensUsed:   resp.Usage.TotalTokens,
		ModelsUsed:   []string{model},
	}, true, nil
}

// --- Scoring Stage ---

type scoringStage struct {
	provider Provider
	model    string
}

// NewScoringStage creates the dimension scoring stage.
func NewScoringStage(provider Provider, model string) Stage {
	return &scoringStage{provider: provider, model: model}
}

func (s *scoringStage) Name() string { return "scoring" }

func (s *scoringStage) Execute(ctx context.Context, input StageInput) (StageResult, bool, error) {
	resp, err := s.provider.Chat(ctx, ChatRequest{
		Model: s.model,
		Messages: AdaptiveMessages(
			"You are a code quality scoring system. Score each dimension 0.0-1.0. Respond with ONLY JSON.",
			fmt.Sprintf("Score unit %s.\nReview: %s\nEvidence: %s\n\nJSON format:\n{\"scores\":{\"correctness\":0.0,\"maintainability\":0.0,\"readability\":0.0,\"testability\":0.0,\"security\":0.0,\"architectural_fitness\":0.0,\"operational_quality\":0.0,\"performance_appropriateness\":0.0,\"change_risk\":0.0},\"confidence\":0.0,\"reasoning\":\"...\"}",
				input.Unit.ID, input.ReviewOutput, input.EvidenceSummary),
			false,
		),
		Temperature: 0.1,
		MaxTokens:   512,
	})
	if err != nil {
		return StageResult{}, false, err
	}

	model := resp.Model
	if model == "" {
		model = s.model
	}
	result := StageResult{TokensUsed: resp.Usage.TotalTokens, ModelsUsed: []string{model}}
	content := resp.Content()

	// Try JSON parse
	var scoring ScoringResponse
	if err := json.Unmarshal([]byte(extractJSON(content)), &scoring); err == nil && len(scoring.Scores) > 0 {
		result.Scores = scoring.Scores
		result.Confidence = scoring.Confidence
		return result, true, nil
	}

	// Fallback: produce default scores
	result.Scores = defaultScores()
	result.Confidence = 0.3
	return result, true, nil
}

func defaultScores() map[string]float64 {
	return map[string]float64{
		"correctness":                 0.80,
		"maintainability":             0.80,
		"readability":                 0.80,
		"testability":                 0.80,
		"security":                    0.80,
		"architectural_fitness":       0.80,
		"operational_quality":         0.80,
		"performance_appropriateness": 0.80,
		"change_risk":                 0.80,
	}
}

// extractJSON tries to find a JSON object in a mixed-content response.
func extractJSON(s string) string {
	start := strings.Index(s, "{")
	if start < 0 {
		return s
	}
	// Find matching closing brace
	depth := 0
	for i := start; i < len(s); i++ {
		switch s[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return s[start : i+1]
			}
		}
	}
	return s[start:]
}
