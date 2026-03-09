package evidence

import (
	"fmt"
	"strings"
	"time"

	"github.com/code-certification/certify/internal/domain"
)

// CodeMetrics holds basic code metrics for a file or symbol.
type CodeMetrics struct {
	TotalLines   int `json:"total_lines"`
	BlankLines   int `json:"blank_lines"`
	CommentLines int `json:"comment_lines"`
	CodeLines    int `json:"code_lines"`
	TodoCount    int `json:"todo_count"`
	Complexity   int `json:"complexity"`
}

// ComputeMetrics computes basic code metrics from source text.
func ComputeMetrics(src string) CodeMetrics {
	if src == "" {
		return CodeMetrics{}
	}

	lines := strings.Split(src, "\n")
	// Remove trailing empty string from final newline
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var m CodeMetrics
	m.TotalLines = len(lines)

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		switch {
		case trimmed == "":
			m.BlankLines++
		case strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "#") || strings.HasPrefix(trimmed, "/*"):
			m.CommentLines++
			if containsTodo(trimmed) {
				m.TodoCount++
			}
		default:
			m.CodeLines++
			// Check inline comments
			if containsTodo(trimmed) {
				m.TodoCount++
			}
		}
	}

	return m
}

func containsTodo(line string) bool {
	upper := strings.ToUpper(line)
	return strings.Contains(upper, "TODO") || strings.Contains(upper, "FIXME")
}

// ToEvidence converts CodeMetrics to a domain.Evidence.
func (m CodeMetrics) ToEvidence() domain.Evidence {
	return domain.Evidence{
		Kind:       domain.EvidenceKindMetrics,
		Source:     "metrics",
		Passed:     true,
		Summary:    fmt.Sprintf("%d lines (%d code, %d comment, %d blank), %d TODOs, complexity %d", m.TotalLines, m.CodeLines, m.CommentLines, m.BlankLines, m.TodoCount, m.Complexity),
		Details:    m,
		Timestamp:  time.Now(),
		Confidence: 1.0,
	}
}
