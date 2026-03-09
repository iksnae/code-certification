package agent

import (
	"context"
	"fmt"
	"strings"
)

// RepoSummary holds repository metadata for generating AI suggestions.
type RepoSummary struct {
	Languages    []string // Detected languages (e.g., "go", "typescript")
	UnitCount    int      // Total code units discovered
	FilePatterns []string // Top-level directories or patterns found
	Policies     []string // Currently configured policy pack names
}

// ScanSuggestion holds the result of an AI-powered scan suggestion.
type ScanSuggestion struct {
	Suggestions string // Human-readable suggestion text
	TokensUsed  int    // Tokens consumed by this call
	Model       string // Model that generated the suggestion
}

// SuggestForRepo sends a single LLM call to generate policy/scope suggestions
// based on the repository summary. Returns empty ScanSuggestion on any failure
// (graceful degradation — never blocks, never errors).
func SuggestForRepo(ctx context.Context, provider Provider, summary RepoSummary) ScanSuggestion {
	if provider == nil {
		return ScanSuggestion{}
	}

	prompt := buildSuggestPrompt(summary)

	resp, err := provider.Chat(ctx, ChatRequest{
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
		Temperature: 0.3,
		MaxTokens:   512,
	})
	if err != nil {
		return ScanSuggestion{}
	}

	content := resp.Content()
	if content == "" {
		return ScanSuggestion{}
	}

	model := resp.Model
	if model == "" {
		model = provider.Name()
	}

	return ScanSuggestion{
		Suggestions: strings.TrimSpace(content),
		TokensUsed:  resp.Usage.TotalTokens,
		Model:       model,
	}
}

func buildSuggestPrompt(summary RepoSummary) string {
	langs := "None detected"
	if len(summary.Languages) > 0 {
		langs = strings.Join(summary.Languages, ", ")
	}

	patterns := "None"
	if len(summary.FilePatterns) > 0 {
		patterns = strings.Join(summary.FilePatterns, ", ")
	}

	policies := "Default policies"
	if len(summary.Policies) > 0 {
		policies = strings.Join(summary.Policies, ", ")
	}

	return fmt.Sprintf(`You are a code certification advisor. Based on the repository summary below, suggest:
1. Policy adjustments (are the default policies appropriate?)
2. Scope refinements (should any paths be excluded or specifically included?)
3. Risk areas (which directories or patterns should get shorter certification windows?)

Keep suggestions brief and actionable. If the defaults look appropriate, say so.

## Repository Summary
Languages: %s
%d code units discovered
File patterns: %s
Current policies: %s`, langs, summary.UnitCount, patterns, policies)
}
