package github_test

import (
	"strings"
	"testing"

	gh "github.com/iksnae/code-certification/internal/github"
)

func TestGeneratePRWorkflow(t *testing.T) {
	yaml := gh.GeneratePRWorkflow()

	if !strings.Contains(yaml, "pull_request") {
		t.Error("should contain pull_request trigger")
	}
	if !strings.Contains(yaml, "OPENROUTER_API_KEY") {
		t.Error("should reference OPENROUTER_API_KEY secret")
	}
	if !strings.Contains(yaml, "certify certify") {
		t.Error("should run certify certify")
	}
	if !strings.Contains(yaml, "certify report") {
		t.Error("should run certify report")
	}
}

func TestGenerateNightlyWorkflow(t *testing.T) {
	yaml := gh.GenerateNightlyWorkflow()

	if !strings.Contains(yaml, "schedule") {
		t.Error("should contain schedule trigger")
	}
	if !strings.Contains(yaml, "cron") {
		t.Error("should contain cron schedule")
	}
	if !strings.Contains(yaml, "certify scan") {
		t.Error("should run certify scan")
	}
}

func TestGenerateWeeklyWorkflow(t *testing.T) {
	yaml := gh.GenerateWeeklyWorkflow()

	if !strings.Contains(yaml, "schedule") {
		t.Error("should contain schedule trigger")
	}
	if !strings.Contains(yaml, "certify report") {
		t.Error("should run certify report")
	}
}

func TestWorkflows_ValidYAML(t *testing.T) {
	// Basic validation — all workflows should be non-empty and start with name:
	for name, gen := range map[string]func() string{
		"pr":      gh.GeneratePRWorkflow,
		"nightly": gh.GenerateNightlyWorkflow,
		"weekly":  gh.GenerateWeeklyWorkflow,
	} {
		yaml := gen()
		if !strings.HasPrefix(yaml, "name:") {
			t.Errorf("%s workflow should start with 'name:'", name)
		}
		if len(yaml) < 100 {
			t.Errorf("%s workflow seems too short: %d bytes", name, len(yaml))
		}
	}
}
