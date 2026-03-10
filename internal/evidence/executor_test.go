package evidence_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/evidence"
)

func TestToolExecutor_LintFindings_Default(t *testing.T) {
	te := evidence.NewToolExecutor(t.TempDir())
	findings := te.LintFindings()
	if findings != nil {
		t.Errorf("default LintFindings should be nil, got %d", len(findings))
	}
}

func TestToolExecutor_CoverageProfile_Default(t *testing.T) {
	te := evidence.NewToolExecutor(t.TempDir())
	profile := te.CoverageProfile()
	if profile != "" {
		t.Errorf("default CoverageProfile should be empty, got %q", profile)
	}
}
