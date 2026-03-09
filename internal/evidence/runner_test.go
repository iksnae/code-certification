package evidence_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/evidence"
)

func TestParseGoVet_Clean(t *testing.T) {
	result := evidence.ParseGoVet("", 0)
	if result.ErrorCount != 0 {
		t.Errorf("ErrorCount = %d, want 0", result.ErrorCount)
	}
	ev := result.ToEvidence()
	if !ev.Passed {
		t.Error("clean go vet evidence should pass")
	}
}

func TestParseGoVet_WithErrors(t *testing.T) {
	stderr := `# github.com/example/pkg
./main.go:10:2: unreachable code
./main.go:15:6: result of fmt.Sprintf call not used
`
	result := evidence.ParseGoVet(stderr, 1)
	if result.ErrorCount != 2 {
		t.Errorf("ErrorCount = %d, want 2", result.ErrorCount)
	}
	if len(result.Findings) != 2 {
		t.Errorf("Findings = %d, want 2", len(result.Findings))
	}
	ev := result.ToEvidence()
	if ev.Passed {
		t.Error("go vet with errors should not pass")
	}
}

func TestParseGoTestJSON_AllPass(t *testing.T) {
	output := `{"Test":"TestFoo","Action":"pass","Elapsed":0.01}
{"Test":"TestBar","Action":"pass","Elapsed":0.02}
{"Action":"pass","Elapsed":0.5}
`
	result := evidence.ParseGoTestJSON(output)
	if result.PassedCount != 2 {
		t.Errorf("PassedCount = %d, want 2", result.PassedCount)
	}
	if result.FailedCount != 0 {
		t.Errorf("FailedCount = %d, want 0", result.FailedCount)
	}
	ev := result.ToEvidence()
	if !ev.Passed {
		t.Error("all-pass test evidence should pass")
	}
}

func TestParseGoTestJSON_WithFailures(t *testing.T) {
	output := `{"Test":"TestFoo","Action":"pass","Elapsed":0.01}
{"Test":"TestBroken","Action":"fail","Elapsed":0.1}
{"Action":"fail","Elapsed":0.5}
`
	result := evidence.ParseGoTestJSON(output)
	if result.PassedCount != 1 {
		t.Errorf("PassedCount = %d, want 1", result.PassedCount)
	}
	if result.FailedCount != 1 {
		t.Errorf("FailedCount = %d, want 1", result.FailedCount)
	}
	ev := result.ToEvidence()
	if ev.Passed {
		t.Error("test with failures should not pass")
	}
}

func TestParseGoTestJSON_Empty(t *testing.T) {
	result := evidence.ParseGoTestJSON("")
	if result.TotalCount != 0 {
		t.Errorf("empty output TotalCount = %d, want 0", result.TotalCount)
	}
}

func TestParseCoverProfile(t *testing.T) {
	profile := `mode: set
github.com/example/pkg/main.go:10.1,15.1 1 1
github.com/example/pkg/main.go:17.1,20.1 1 0
github.com/example/pkg/main.go:22.1,25.1 1 1
`
	coverage := evidence.ParseCoverProfile(profile)
	// 2 out of 3 statements covered = 66.7%
	if coverage < 0.66 || coverage > 0.68 {
		t.Errorf("coverage = %f, want ~0.667", coverage)
	}
}

func TestParseCoverProfile_Empty(t *testing.T) {
	coverage := evidence.ParseCoverProfile("")
	if coverage != 0 {
		t.Errorf("empty profile coverage = %f, want 0", coverage)
	}
}

func TestParseGitLogRunner(t *testing.T) {
	output := `abc1234	alice	2025-01-15
def5678	bob	2025-06-10
ghi9012	alice	2025-09-05
`
	stats := evidence.ParseGitLogWithAge(output, "2025-01-15")
	if stats.CommitCount != 3 {
		t.Errorf("CommitCount = %d, want 3", stats.CommitCount)
	}
	if stats.AuthorCount != 2 {
		t.Errorf("AuthorCount = %d, want 2", stats.AuthorCount)
	}
	if stats.AgeDays < 1 {
		t.Errorf("AgeDays = %d, want > 0", stats.AgeDays)
	}
}

func TestParseGolangciLintJSON(t *testing.T) {
	output := `{"Issues":[
		{"FromLinter":"govet","Text":"unreachable code","Pos":{"Filename":"main.go","Line":10},"Severity":"error"},
		{"FromLinter":"unused","Text":"unused variable","Pos":{"Filename":"main.go","Line":15},"Severity":"warning"}
	]}`
	result := evidence.ParseGolangciLintJSON(output)
	if result.ErrorCount != 1 {
		t.Errorf("ErrorCount = %d, want 1", result.ErrorCount)
	}
	if result.WarnCount != 1 {
		t.Errorf("WarnCount = %d, want 1", result.WarnCount)
	}
	if len(result.Findings) != 2 {
		t.Errorf("Findings = %d, want 2", len(result.Findings))
	}
}

func TestParseGolangciLintJSON_Clean(t *testing.T) {
	result := evidence.ParseGolangciLintJSON(`{"Issues":[]}`)
	ev := result.ToEvidence()
	if !ev.Passed {
		t.Error("clean golangci-lint should pass")
	}
}

func TestParseGolangciLintJSON_Empty(t *testing.T) {
	result := evidence.ParseGolangciLintJSON("")
	if result.ErrorCount != 0 {
		t.Error("empty output should be treated as clean")
	}
}
