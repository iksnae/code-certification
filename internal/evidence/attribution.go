package evidence

// AttributeLintToFile filters lint findings to those belonging to a specific file
// and returns a LintResult scoped to that file.
func AttributeLintToFile(findings []LintFinding, filePath string) LintResult {
	result := LintResult{Tool: "golangci-lint:unit"}

	for _, f := range findings {
		if f.File != filePath {
			continue
		}
		result.Findings = append(result.Findings, f)
		if f.Severity == "error" {
			result.ErrorCount++
		} else {
			result.WarnCount++
		}
	}

	return result
}

// AttributeLintToUnit filters lint findings to those within a specific file
// and line range [startLine, endLine] inclusive.
func AttributeLintToUnit(findings []LintFinding, filePath string, startLine, endLine int) LintResult {
	result := LintResult{Tool: "golangci-lint:unit"}

	for _, f := range findings {
		if f.File != filePath {
			continue
		}
		if f.Line < startLine || f.Line > endLine {
			continue
		}
		result.Findings = append(result.Findings, f)
		if f.Severity == "error" {
			result.ErrorCount++
		} else {
			result.WarnCount++
		}
	}

	return result
}
