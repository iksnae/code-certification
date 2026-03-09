package evidence

import (
	"encoding/json"
	"math"
	"strings"
	"time"
)

// ParseGoVet parses go vet stderr output into a LintResult.
func ParseGoVet(stderr string, exitCode int) LintResult {
	result := LintResult{Tool: "go vet"}

	if exitCode == 0 && strings.TrimSpace(stderr) == "" {
		return result
	}

	lines := strings.Split(strings.TrimSpace(stderr), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, ": ", 2)
		finding := LintFinding{Severity: "error", Rule: "govet"}
		if len(parts) >= 2 {
			if colonIdx := strings.Index(parts[0], ":"); colonIdx >= 0 {
				finding.File = parts[0][:colonIdx]
			} else {
				finding.File = parts[0]
			}
			finding.Message = parts[1]
		} else {
			finding.Message = line
		}
		result.Findings = append(result.Findings, finding)
		result.ErrorCount++
	}

	return result
}

// goTestEvent represents a single line from `go test -json`.
type goTestEvent struct {
	Test    string  `json:"Test"`
	Action  string  `json:"Action"`
	Elapsed float64 `json:"Elapsed"`
}

// ParseGoTestJSON parses `go test -json` output into a TestResult.
func ParseGoTestJSON(output string) TestResult {
	result := TestResult{Tool: "go test"}

	if strings.TrimSpace(output) == "" {
		return result
	}

	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var ev goTestEvent
		if err := json.Unmarshal([]byte(line), &ev); err != nil {
			continue
		}
		if ev.Test == "" {
			continue
		}
		switch ev.Action {
		case "pass":
			result.PassedCount++
			result.TotalCount++
		case "fail":
			result.FailedCount++
			result.TotalCount++
		case "skip":
			result.SkipCount++
			result.TotalCount++
		}
	}

	return result
}

// ParseCoverProfile parses a Go coverage profile and returns coverage as 0.0-1.0.
func ParseCoverProfile(profile string) float64 {
	if strings.TrimSpace(profile) == "" {
		return 0
	}

	lines := strings.Split(strings.TrimSpace(profile), "\n")
	var totalStmts, coveredStmts int

	for _, line := range lines {
		if strings.HasPrefix(line, "mode:") {
			continue
		}
		// Format: file:startline.col,endline.col numStatements count
		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue
		}
		stmts := simpleAtoi(parts[1])
		count := simpleAtoi(parts[2])
		totalStmts += stmts
		if count > 0 {
			coveredStmts += stmts
		}
	}

	if totalStmts == 0 {
		return 0
	}
	return float64(coveredStmts) / float64(totalStmts)
}

func simpleAtoi(s string) int {
	n := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	return n
}

// ParseGitLogWithAge parses git log output and computes age from the earliest date.
func ParseGitLogWithAge(output string, earliestDate string) GitStats {
	stats := ParseGitLog(output)

	if earliestDate != "" {
		if earliest, err := time.Parse("2006-01-02", earliestDate); err == nil {
			stats.AgeDays = int(math.Round(time.Since(earliest).Hours() / 24))
		}
	}

	return stats
}

// golangciLintOutput represents golangci-lint JSON output.
type golangciLintOutput struct {
	Issues []golangciLintIssue `json:"Issues"`
}

type golangciLintIssue struct {
	FromLinter string `json:"FromLinter"`
	Text       string `json:"Text"`
	Severity   string `json:"Severity"`
	Pos        struct {
		Filename string `json:"Filename"`
		Line     int    `json:"Line"`
	} `json:"Pos"`
}

// ParseGolangciLintJSON parses golangci-lint JSON output into a LintResult.
func ParseGolangciLintJSON(output string) LintResult {
	result := LintResult{Tool: "golangci-lint"}

	if strings.TrimSpace(output) == "" {
		return result // No output = clean
	}

	var parsed golangciLintOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		return result // Can't parse = treat as clean
	}

	for _, issue := range parsed.Issues {
		sev := issue.Severity
		if sev == "" {
			sev = "warning"
		}
		result.Findings = append(result.Findings, LintFinding{
			File:     issue.Pos.Filename,
			Line:     issue.Pos.Line,
			Message:  issue.Text,
			Severity: sev,
			Rule:     issue.FromLinter,
		})
		if sev == "error" {
			result.ErrorCount++
		} else {
			result.WarnCount++
		}
	}

	return result
}
