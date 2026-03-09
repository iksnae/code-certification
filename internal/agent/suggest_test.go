package agent

import (
	"context"
	"fmt"
	"testing"
)

func TestSuggestForRepo_Success(t *testing.T) {
	mock := &mockSuggestProvider{
		response: ChatResponse{
			Choices: []Choice{
				{Message: Message{Content: "1. Consider adding security policy for crypto packages\n2. Exclude vendor/ from scope"}},
			},
			Usage: Usage{TotalTokens: 120},
		},
	}

	summary := RepoSummary{
		Languages:    []string{"go", "typescript"},
		UnitCount:    150,
		FilePatterns: []string{"internal/", "cmd/", "pkg/"},
		Policies:     []string{"go-standard"},
	}

	result := SuggestForRepo(context.Background(), mock, summary)
	if result.Suggestions == "" {
		t.Fatal("SuggestForRepo() returned empty suggestions")
	}
	if result.TokensUsed != 120 {
		t.Errorf("TokensUsed = %d, want 120", result.TokensUsed)
	}
}

func TestSuggestForRepo_Error(t *testing.T) {
	mock := &mockSuggestProvider{
		err: fmt.Errorf("rate limited"),
	}

	summary := RepoSummary{
		Languages: []string{"go"},
		UnitCount: 50,
	}

	result := SuggestForRepo(context.Background(), mock, summary)
	if result.Suggestions != "" {
		t.Errorf("SuggestForRepo() on error should return empty suggestions, got %q", result.Suggestions)
	}
	if result.TokensUsed != 0 {
		t.Errorf("TokensUsed = %d, want 0", result.TokensUsed)
	}
}

func TestSuggestForRepo_EmptyResponse(t *testing.T) {
	mock := &mockSuggestProvider{
		response: ChatResponse{
			Choices: []Choice{},
		},
	}

	summary := RepoSummary{
		Languages: []string{"python"},
		UnitCount: 10,
	}

	result := SuggestForRepo(context.Background(), mock, summary)
	if result.Suggestions != "" {
		t.Errorf("SuggestForRepo() on empty response should return empty, got %q", result.Suggestions)
	}
}

// mockSuggestProvider implements Provider for testing.
type mockSuggestProvider struct {
	response ChatResponse
	err      error
}

func (m *mockSuggestProvider) Chat(_ context.Context, _ ChatRequest) (ChatResponse, error) {
	return m.response, m.err
}

func (m *mockSuggestProvider) Name() string { return "mock" }
