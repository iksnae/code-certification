package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

// --- Deep Review Stage (for local models) ---

type deepReviewStage struct {
	provider Provider
	model    string
}

// NewDeepReviewStage creates a comprehensive review stage for local models.
// Unlike the standard prescreen→review flow, this always performs a full analysis
// since local model tokens are free.
func NewDeepReviewStage(provider Provider, model string) Stage {
	return &deepReviewStage{provider: provider, model: model}
}

func (s *deepReviewStage) Name() string { return "deep-review" }

func (s *deepReviewStage) Execute(ctx context.Context, input StageInput) (StageResult, bool, error) {
	codeSnippet := input.SourceCode
	// Local models: allow much larger snippets (no token cost)
	if len(codeSnippet) > 16000 {
		codeSnippet = codeSnippet[:16000] + "\n... (truncated)"
	}

	resp, err := s.provider.Chat(ctx, ChatRequest{
		Model: s.model,
		Messages: AdaptiveMessages(
			deepReviewSystemPrompt,
			fmt.Sprintf(deepReviewUserPrompt,
				input.Unit.ID, input.Unit.Type, input.Unit.ID.Language(),
				codeSnippet, input.EvidenceSummary),
			false,
		),
		Temperature: 0.3,
		MaxTokens:   4096,
	})
	if err != nil {
		return StageResult{}, false, err
	}

	model := resp.Model
	if model == "" {
		model = s.model
	}

	content := resp.Content()
	result := StageResult{
		ReviewOutput: content,
		TokensUsed:   resp.Usage.TotalTokens,
		ModelsUsed:   []string{model},
	}

	// Try to parse structured response
	var dr DeepReviewResponse
	if err := json.Unmarshal([]byte(extractJSON(content)), &dr); err == nil {
		result.PrescreenReason = dr.Summary
		result.Suggestions = dr.Suggestions
		if len(dr.Risks) > 0 {
			for _, r := range dr.Risks {
				result.Suggestions = append(result.Suggestions, "⚠️ "+r)
			}
		}
		if len(dr.SystemImpact) > 0 {
			for _, si := range dr.SystemImpact {
				result.Suggestions = append(result.Suggestions, "🔗 "+si)
			}
		}
		// Use structured review as output if present
		if dr.Review != "" {
			result.ReviewOutput = dr.Review
		}
	} else {
		// Free-text response — extract what we can
		result.PrescreenReason = extractFirstSentence(content)
	}

	return result, true, nil // always continue to scoring
}

// DeepReviewResponse is the structured output from deep review.
type DeepReviewResponse struct {
	Summary      string   `json:"summary"`
	Review       string   `json:"review"`
	Risks        []string `json:"risks"`
	SystemImpact []string `json:"system_impact"`
	Suggestions  []string `json:"suggestions"`
	Strengths    []string `json:"strengths"`
}

const deepReviewSystemPrompt = `You are a senior software engineer performing a thorough code review. Analyze the code for:

1. **Quality Assessment**: correctness, readability, maintainability, testability
2. **Risk Evaluation**: security vulnerabilities, error handling gaps, race conditions, resource leaks, edge cases
3. **System Impact**: how this unit affects the broader system — coupling, API surface, failure propagation, performance implications
4. **Actionable Suggestions**: specific, concrete improvements (not vague advice)

Respond with JSON:
{
  "summary": "one-line quality assessment",
  "review": "detailed review paragraph(s) with specific observations",
  "risks": ["specific risk 1", "specific risk 2"],
  "system_impact": ["how this unit affects X", "coupling concern with Y"],
  "suggestions": ["concrete suggestion 1", "concrete suggestion 2"],
  "strengths": ["what this code does well"]
}

Be specific. Reference line-level details. Identify real issues, not generic advice.`

const deepReviewUserPrompt = `Perform a deep review of this code unit:

**Unit:** %s
**Type:** %s
**Language:** %s

` + "```" + `
%s
` + "```" + `

**Automated Evidence:** %s

Provide your structured analysis.`

// extractFirstSentence gets the first sentence from a text block.
func extractFirstSentence(text string) string {
	text = strings.TrimSpace(text)
	// Strip markdown/JSON artifacts
	text = strings.TrimLeft(text, "#{`\"")
	for _, sep := range []string{". ", ".\n", "!\n"} {
		if idx := strings.Index(text, sep); idx > 0 && idx < 200 {
			return text[:idx+1]
		}
	}
	if len(text) > 200 {
		return text[:200] + "..."
	}
	return text
}

// --- Helpers for unit report enrichment ---

// FormatDeepObservations creates formatted observation strings from a ReviewResult.
func FormatDeepObservations(result ReviewResult) []string {
	var obs []string

	if result.PrescreenReason != "" {
		obs = append(obs, "🤖 "+result.PrescreenReason)
	}

	for _, s := range result.Suggestions {
		if s == "" {
			continue
		}
		// Already prefixed with emoji from deep review parsing
		if strings.HasPrefix(s, "⚠️") || strings.HasPrefix(s, "🔗") {
			obs = append(obs, s)
		} else {
			obs = append(obs, "💡 "+s)
		}
	}

	return obs
}

// IsDeepReview returns true if the result came from a deep review (local model).
func IsDeepReview(result ReviewResult) bool {
	return result.ReviewOutput != "" && (result.Reviewed || result.Prescreened)
}

// FormatReviewForRecord creates a compact review summary suitable for record storage.
func FormatReviewForRecord(result ReviewResult) string {
	if result.ReviewOutput == "" {
		return ""
	}
	review := result.ReviewOutput
	// Cap at 2000 chars for storage
	if len(review) > 2000 {
		review = review[:2000] + "\n... (truncated)"
	}
	return review
}

// ToDeepEvidence creates evidence from a deep review result.
func (r ReviewResult) ToDeepEvidence() domain.Evidence {
	source := "agent-deep-review"
	if len(r.ModelsUsed) > 0 {
		source = "agent-deep-review:" + joinModels(r.ModelsUsed)
	}

	summary := fmt.Sprintf("AI deep review (confidence: %.0f%%)", r.Confidence*100)
	if r.PrescreenReason != "" {
		summary = fmt.Sprintf("AI: %s", r.PrescreenReason)
	}
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
