package evidence

import (
	"fmt"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
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
			// Only count TODOs in inline comments (after //), not in code
			if idx := strings.Index(trimmed, "//"); idx >= 0 {
				if containsTodo(trimmed[idx:]) {
					m.TodoCount++
				}
			}
		}
	}

	return m
}

func containsTodo(line string) bool {
	upper := strings.ToUpper(line)
	for _, keyword := range []string{"TODO", "FIXME"} {
		idx := strings.Index(upper, keyword)
		if idx < 0 {
			continue
		}
		// Word boundary check: the character before the keyword must not be
		// a letter (to exclude identifiers like "extractTodoCount").
		if idx > 0 && isLetter(upper[idx-1]) {
			continue
		}
		// Also check after: "TODOCOUNT" is an identifier, not a marker.
		end := idx + len(keyword)
		if end < len(upper) && isLetter(upper[end]) {
			continue
		}
		// Check if the keyword is inside a quoted string in the comment.
		before := line[:idx]
		if strings.Count(before, "\"")%2 == 1 {
			continue
		}
		return true
	}
	return false
}

// isLetter returns true if b is an ASCII letter.
func isLetter(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

// ToEvidence converts CodeMetrics to a domain.Evidence.
func (m CodeMetrics) ToEvidence() domain.Evidence {
	return domain.Evidence{
		Kind:    domain.EvidenceKindMetrics,
		Source:  "metrics",
		Passed:  true,
		Summary: fmt.Sprintf("%d lines (%d code, %d comment, %d blank), %d TODOs, complexity %d", m.TotalLines, m.CodeLines, m.CommentLines, m.BlankLines, m.TodoCount, m.Complexity),
		Metrics: map[string]float64{
			"total_lines":   float64(m.TotalLines),
			"code_lines":    float64(m.CodeLines),
			"comment_lines": float64(m.CommentLines),
			"blank_lines":   float64(m.BlankLines),
			"todo_count":    float64(m.TodoCount),
			"complexity":    float64(m.Complexity),
		},
		Details:    m,
		Timestamp:  time.Now(),
		Confidence: 1.0,
	}
}
