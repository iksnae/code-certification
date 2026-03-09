package evidence_test

import (
	"testing"

	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/evidence"
)

func TestCodeMetrics_LineCount(t *testing.T) {
	src := `package main

import "fmt"

func main() {
	fmt.Println("hello")
}

func helper() string {
	return "world"
}
`
	m := evidence.ComputeMetrics(src)
	if m.TotalLines != 11 {
		t.Errorf("TotalLines = %d, want 11", m.TotalLines)
	}
}

func TestCodeMetrics_BlankLines(t *testing.T) {
	src := "line1\n\nline3\n\nline5\n"
	m := evidence.ComputeMetrics(src)
	if m.BlankLines != 2 {
		t.Errorf("BlankLines = %d, want 2", m.BlankLines)
	}
}

func TestCodeMetrics_CommentLines(t *testing.T) {
	src := `// Package main
package main

// main is the entry point
func main() {
	// do stuff
}
`
	m := evidence.ComputeMetrics(src)
	if m.CommentLines < 3 {
		t.Errorf("CommentLines = %d, want at least 3", m.CommentLines)
	}
}

func TestCodeMetrics_TodoCount(t *testing.T) {
	src := `func main() {
	// TODO: fix this
	// FIXME: also this
	// regular comment
	// TODO another one
}
`
	m := evidence.ComputeMetrics(src)
	if m.TodoCount != 3 {
		t.Errorf("TodoCount = %d, want 3", m.TodoCount)
	}
}

func TestCodeMetrics_Empty(t *testing.T) {
	m := evidence.ComputeMetrics("")
	if m.TotalLines != 0 {
		t.Errorf("empty source TotalLines = %d, want 0", m.TotalLines)
	}
}

func TestCodeMetrics_ToEvidence(t *testing.T) {
	m := evidence.CodeMetrics{
		TotalLines:   100,
		BlankLines:   10,
		CommentLines: 15,
		CodeLines:    75,
		TodoCount:    2,
	}
	ev := m.ToEvidence()
	if ev.Kind != domain.EvidenceKindMetrics {
		t.Errorf("Kind = %v, want metrics", ev.Kind)
	}
	if ev.Source != "metrics" {
		t.Errorf("Source = %q, want metrics", ev.Source)
	}
}
