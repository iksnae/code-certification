package agent_test

import (
	"path/filepath"
	"testing"

	"github.com/code-certification/certify/internal/agent"
)

func promptsDir() string {
	return filepath.Join("..", "..", "templates", "prompts")
}

func TestLoadPrompt(t *testing.T) {
	tmpl, err := agent.LoadPrompt(filepath.Join(promptsDir(), "prescreen.v1.md"))
	if err != nil {
		t.Fatalf("LoadPrompt() error: %v", err)
	}

	vars := map[string]string{
		"UnitID":          "go://main.go#main",
		"Language":        "go",
		"Path":            "main.go",
		"UnitType":        "function",
		"EvidenceSummary": "lint: clean, tests: 10/10 pass",
	}

	result, err := tmpl.Render(vars)
	if err != nil {
		t.Fatalf("Render() error: %v", err)
	}

	if len(result) == 0 {
		t.Error("rendered template should not be empty")
	}
}

func TestLoadPrompt_AllTemplates(t *testing.T) {
	templates := []string{
		"prescreen.v1.md",
		"review.v1.md",
		"scoring.v1.md",
		"decision.v1.md",
		"remediation.v1.md",
	}

	for _, name := range templates {
		_, err := agent.LoadPrompt(filepath.Join(promptsDir(), name))
		if err != nil {
			t.Errorf("LoadPrompt(%s) error: %v", name, err)
		}
	}
}

func TestPromptVersion(t *testing.T) {
	tmpl, _ := agent.LoadPrompt(filepath.Join(promptsDir(), "prescreen.v1.md"))
	if tmpl.Version() != "v1" {
		t.Errorf("Version() = %q, want v1", tmpl.Version())
	}
}

func TestLoadPrompt_NotFound(t *testing.T) {
	_, err := agent.LoadPrompt("/nonexistent/prompt.md")
	if err == nil {
		t.Fatal("should error on missing file")
	}
}
