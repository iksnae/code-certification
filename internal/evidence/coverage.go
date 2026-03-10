package evidence

import (
	"strings"
)

// FileCoverage holds aggregated coverage data for a single file.
type FileCoverage struct {
	Statements int // Total statements in coverage blocks
	Covered    int // Statements with count > 0
}

// CoverageMap maps file paths to their aggregated coverage data.
type CoverageMap map[string]FileCoverage

// ParseCoverProfilePerFunc parses a Go coverage profile and returns
// per-file coverage data. The profile format is:
//
//	mode: set
//	file:startline.col,endline.col numstatements count
func ParseCoverProfilePerFunc(profile string) CoverageMap {
	if strings.TrimSpace(profile) == "" {
		return CoverageMap{}
	}

	cm := make(CoverageMap)
	lines := strings.Split(strings.TrimSpace(profile), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "mode:") {
			continue
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Format: file:startline.col,endline.col numStatements count
		// Find the file part (everything before the first colon+digit pattern)
		colonIdx := strings.Index(line, ":")
		if colonIdx < 0 {
			continue
		}
		file := line[:colonIdx]

		// Parse numStatements and count from the remaining fields
		rest := line[colonIdx+1:]
		// Find the space-separated fields after the line range
		spaceIdx := strings.Index(rest, " ")
		if spaceIdx < 0 {
			continue
		}
		fields := strings.Fields(rest[spaceIdx:])
		if len(fields) < 2 {
			continue
		}

		stmts := simpleAtoi(fields[0])
		count := simpleAtoi(fields[1])

		fc := cm[file]
		fc.Statements += stmts
		if count > 0 {
			fc.Covered += stmts
		}
		cm[file] = fc
	}

	// Remove zero-entry files
	for k, v := range cm {
		if v.Statements == 0 && v.Covered == 0 {
			delete(cm, k)
		}
	}

	return cm
}

// CoverageForFile returns the coverage ratio (0.0–1.0) for a specific file.
// Returns -1 if the file is not found in the coverage map.
func CoverageForFile(cm CoverageMap, filePath string) float64 {
	fc, ok := cm[filePath]
	if !ok {
		return -1
	}
	if fc.Statements == 0 {
		return 0
	}
	return float64(fc.Covered) / float64(fc.Statements)
}
