package evidence_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/evidence"
)

func TestParseCoverProfilePerFunc_Empty(t *testing.T) {
	cm := evidence.ParseCoverProfilePerFunc("")
	if len(cm) != 0 {
		t.Errorf("empty profile should return empty map, got %d entries", len(cm))
	}
}

func TestParseCoverProfilePerFunc_ModeLineOnly(t *testing.T) {
	cm := evidence.ParseCoverProfilePerFunc("mode: set\n")
	if len(cm) != 0 {
		t.Errorf("mode-only profile should return empty map, got %d entries", len(cm))
	}
}

func TestParseCoverProfilePerFunc_SingleFile(t *testing.T) {
	profile := `mode: set
github.com/example/pkg/foo.go:10.2,15.3 3 1
github.com/example/pkg/foo.go:17.2,20.3 2 0
`
	cm := evidence.ParseCoverProfilePerFunc(profile)
	if len(cm) != 1 {
		t.Fatalf("expected 1 file, got %d", len(cm))
	}
	fc, ok := cm["github.com/example/pkg/foo.go"]
	if !ok {
		t.Fatal("expected entry for foo.go")
	}
	// 3 stmts covered + 2 stmts not covered = 5 total, 3 covered
	if fc.Statements != 5 {
		t.Errorf("Statements = %d, want 5", fc.Statements)
	}
	if fc.Covered != 3 {
		t.Errorf("Covered = %d, want 3", fc.Covered)
	}
}

func TestParseCoverProfilePerFunc_MultipleFiles(t *testing.T) {
	profile := `mode: set
github.com/example/pkg/foo.go:10.2,15.3 3 1
github.com/example/pkg/bar.go:5.2,8.3 4 1
github.com/example/pkg/bar.go:10.2,12.3 2 0
`
	cm := evidence.ParseCoverProfilePerFunc(profile)
	if len(cm) != 2 {
		t.Fatalf("expected 2 files, got %d", len(cm))
	}
	if _, ok := cm["github.com/example/pkg/foo.go"]; !ok {
		t.Error("missing foo.go entry")
	}
	bar := cm["github.com/example/pkg/bar.go"]
	if bar.Statements != 6 {
		t.Errorf("bar Statements = %d, want 6", bar.Statements)
	}
	if bar.Covered != 4 {
		t.Errorf("bar Covered = %d, want 4", bar.Covered)
	}
}

func TestCoverageForFile_Found(t *testing.T) {
	profile := `mode: set
github.com/example/pkg/foo.go:10.2,15.3 4 1
github.com/example/pkg/foo.go:17.2,20.3 1 0
`
	cm := evidence.ParseCoverProfilePerFunc(profile)
	cov := evidence.CoverageForFile(cm, "github.com/example/pkg/foo.go")
	// 4 covered out of 5 = 0.8
	if cov < 0.79 || cov > 0.81 {
		t.Errorf("CoverageForFile = %f, want ~0.8", cov)
	}
}

func TestCoverageForFile_NotFound(t *testing.T) {
	cm := evidence.ParseCoverProfilePerFunc("mode: set\n")
	cov := evidence.CoverageForFile(cm, "nonexistent.go")
	if cov != -1 {
		t.Errorf("CoverageForFile for missing file = %f, want -1", cov)
	}
}

func TestCoverageForFile_ZeroStatements(t *testing.T) {
	// Edge case: file present but zero statements
	cm := evidence.CoverageMap{
		"empty.go": {Statements: 0, Covered: 0},
	}
	cov := evidence.CoverageForFile(cm, "empty.go")
	if cov != 0 {
		t.Errorf("CoverageForFile for zero-stmt file = %f, want 0", cov)
	}
}
