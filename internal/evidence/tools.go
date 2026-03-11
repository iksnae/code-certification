package evidence

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

// --- ESLint (TypeScript/JavaScript) ---

// eslintJSONOutput represents ESLint JSON output format.
type eslintJSONOutput []eslintFileResult

type eslintFileResult struct {
	FilePath string      `json:"filePath"`
	Messages []eslintMsg `json:"messages"`
}

type eslintMsg struct {
	RuleID   string `json:"ruleId"`
	Severity int    `json:"severity"` // 1=warning, 2=error
	Message  string `json:"message"`
	Line     int    `json:"line"`
}

// ParseESLintJSON parses ESLint JSON output into a LintResult.
func ParseESLintJSON(output string) LintResult {
	result := LintResult{Tool: "eslint"}
	if strings.TrimSpace(output) == "" {
		return result
	}

	var parsed eslintJSONOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		return result
	}

	for _, file := range parsed {
		for _, msg := range file.Messages {
			sev := "warning"
			if msg.Severity >= 2 {
				sev = "error"
				result.ErrorCount++
			} else {
				result.WarnCount++
			}
			result.Findings = append(result.Findings, LintFinding{
				File:     file.FilePath,
				Line:     msg.Line,
				Message:  msg.Message,
				Severity: sev,
				Rule:     msg.RuleID,
			})
		}
	}
	return result
}

// --- Ruff (Python linter) ---

// ruffJSONOutput is a single ruff finding in JSON mode.
type ruffJSONFinding struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Location struct {
		Row int `json:"row"`
	} `json:"location"`
}

// ParseRuffJSON parses ruff JSON output (one object per line) into a LintResult.
func ParseRuffJSON(output string) LintResult {
	result := LintResult{Tool: "ruff"}
	if strings.TrimSpace(output) == "" {
		return result
	}

	// Ruff outputs a JSON array
	var findings []ruffJSONFinding
	if err := json.Unmarshal([]byte(output), &findings); err != nil {
		// Try line-by-line (older ruff versions)
		for _, line := range strings.Split(output, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			var f ruffJSONFinding
			if err := json.Unmarshal([]byte(line), &f); err == nil {
				findings = append(findings, f)
			}
		}
	}

	for _, f := range findings {
		sev := "warning"
		// E-prefixed codes are errors, most others are warnings
		if strings.HasPrefix(f.Code, "E") || strings.HasPrefix(f.Code, "F") {
			sev = "error"
			result.ErrorCount++
		} else {
			result.WarnCount++
		}
		result.Findings = append(result.Findings, LintFinding{
			File:     f.Filename,
			Line:     f.Location.Row,
			Message:  f.Message,
			Severity: sev,
			Rule:     f.Code,
		})
	}
	return result
}

// --- Cargo Clippy (Rust linter) ---

// cargoMsg represents a Cargo JSON diagnostic message.
type cargoMsg struct {
	Reason  string     `json:"reason"`
	Message *cargoDiag `json:"message,omitempty"`
}

type cargoDiag struct {
	Code    *cargoCode  `json:"code"`
	Level   string      `json:"level"` // "warning", "error"
	Message string      `json:"message"`
	Spans   []cargoSpan `json:"spans"`
}

type cargoCode struct {
	Code string `json:"code"`
}

type cargoSpan struct {
	FileName  string `json:"file_name"`
	LineStart int    `json:"line_start"`
}

// ParseCargoClippyJSON parses cargo clippy --message-format=json output.
func ParseCargoClippyJSON(output string) LintResult {
	result := LintResult{Tool: "cargo clippy"}
	if strings.TrimSpace(output) == "" {
		return result
	}

	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var msg cargoMsg
		if err := json.Unmarshal([]byte(line), &msg); err != nil {
			continue
		}
		if msg.Reason != "compiler-message" || msg.Message == nil {
			continue
		}
		diag := msg.Message
		if diag.Level != "warning" && diag.Level != "error" {
			continue
		}

		finding := LintFinding{
			Message:  diag.Message,
			Severity: diag.Level,
		}
		if diag.Code != nil {
			finding.Rule = diag.Code.Code
		}
		if len(diag.Spans) > 0 {
			finding.File = diag.Spans[0].FileName
			finding.Line = diag.Spans[0].LineStart
		}

		result.Findings = append(result.Findings, finding)
		if diag.Level == "error" {
			result.ErrorCount++
		} else {
			result.WarnCount++
		}
	}
	return result
}

// --- Jest / Vitest (TypeScript/JavaScript test runners) ---

// jestJSONOutput represents Jest --json output.
type jestJSONOutput struct {
	NumTotalTests   int  `json:"numTotalTests"`
	NumPassedTests  int  `json:"numPassedTests"`
	NumFailedTests  int  `json:"numFailedTests"`
	NumPendingTests int  `json:"numPendingTests"`
	Success         bool `json:"success"`
}

// ParseJestJSON parses Jest JSON output into a TestResult.
func ParseJestJSON(output string) TestResult {
	result := TestResult{Tool: "jest"}
	if strings.TrimSpace(output) == "" {
		return result
	}

	var parsed jestJSONOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		return result
	}

	result.TotalCount = parsed.NumTotalTests
	result.PassedCount = parsed.NumPassedTests
	result.FailedCount = parsed.NumFailedTests
	result.SkipCount = parsed.NumPendingTests
	return result
}

// --- pytest (Python test runner) ---

// pytestJUnitSuites represents pytest JUnit XML output.
type pytestJUnitSuites struct {
	XMLName xml.Name           `xml:"testsuites"`
	Suites  []pytestJUnitSuite `xml:"testsuite"`
}

type pytestJUnitSuite struct {
	Tests    int `xml:"tests,attr"`
	Errors   int `xml:"errors,attr"`
	Failures int `xml:"failures,attr"`
	Skipped  int `xml:"skipped,attr"`
}

// ParsePytestJUnitXML parses pytest JUnit XML output.
func ParsePytestJUnitXML(output string) TestResult {
	result := TestResult{Tool: "pytest"}
	if strings.TrimSpace(output) == "" {
		return result
	}

	var suites pytestJUnitSuites
	if err := xml.Unmarshal([]byte(output), &suites); err != nil {
		// Try single suite
		var suite pytestJUnitSuite
		if err := xml.Unmarshal([]byte(output), &suite); err == nil {
			suites.Suites = append(suites.Suites, suite)
		} else {
			return result
		}
	}

	for _, s := range suites.Suites {
		result.TotalCount += s.Tests
		result.FailedCount += s.Failures + s.Errors
		result.SkipCount += s.Skipped
	}
	result.PassedCount = result.TotalCount - result.FailedCount - result.SkipCount
	if result.PassedCount < 0 {
		result.PassedCount = 0
	}
	return result
}

// ParsePytestShort parses pytest short summary output for fallback.
// Matches patterns like "5 passed", "2 failed", "1 skipped"
func ParsePytestShort(output string) TestResult {
	result := TestResult{Tool: "pytest"}

	passedRe := regexp.MustCompile(`(\d+) passed`)
	failedRe := regexp.MustCompile(`(\d+) failed`)
	skippedRe := regexp.MustCompile(`(\d+) skipped`)
	errorRe := regexp.MustCompile(`(\d+) error`)

	if m := passedRe.FindStringSubmatch(output); len(m) > 1 {
		result.PassedCount, _ = strconv.Atoi(m[1])
	}
	if m := failedRe.FindStringSubmatch(output); len(m) > 1 {
		result.FailedCount, _ = strconv.Atoi(m[1])
	}
	if m := skippedRe.FindStringSubmatch(output); len(m) > 1 {
		result.SkipCount, _ = strconv.Atoi(m[1])
	}
	if m := errorRe.FindStringSubmatch(output); len(m) > 1 {
		n, _ := strconv.Atoi(m[1])
		result.FailedCount += n
	}
	result.TotalCount = result.PassedCount + result.FailedCount + result.SkipCount
	return result
}

// --- Cargo Test (Rust test runner) ---

// ParseCargoTestOutput parses cargo test stdout for test counts.
// Format: "test result: ok. 42 passed; 0 failed; 0 ignored; ..."
func ParseCargoTestOutput(output string) TestResult {
	result := TestResult{Tool: "cargo test"}

	re := regexp.MustCompile(`test result: \w+\. (\d+) passed; (\d+) failed; (\d+) ignored`)
	matches := re.FindAllStringSubmatch(output, -1)

	for _, m := range matches {
		if len(m) >= 4 {
			p, _ := strconv.Atoi(m[1])
			f, _ := strconv.Atoi(m[2])
			s, _ := strconv.Atoi(m[3])
			result.PassedCount += p
			result.FailedCount += f
			result.SkipCount += s
		}
	}
	result.TotalCount = result.PassedCount + result.FailedCount + result.SkipCount
	return result
}

// --- LCOV coverage parsing (used by Jest, c8, nyc, pytest-cov) ---

// ParseLCOV parses LCOV coverage format and returns coverage as 0.0-1.0.
func ParseLCOV(content string) float64 {
	var totalLines, hitLines int

	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "LF:") {
			n, _ := strconv.Atoi(strings.TrimPrefix(line, "LF:"))
			totalLines += n
		} else if strings.HasPrefix(line, "LH:") {
			n, _ := strconv.Atoi(strings.TrimPrefix(line, "LH:"))
			hitLines += n
		}
	}

	if totalLines == 0 {
		return 0
	}
	return float64(hitLines) / float64(totalLines)
}

// --- Cobertura coverage parsing (used by pytest-cov --cov-report=xml) ---

type coberturaReport struct {
	XMLName  xml.Name `xml:"coverage"`
	LineRate string   `xml:"line-rate,attr"`
}

// ParseCoberturaXML parses Cobertura XML and returns coverage as 0.0-1.0.
func ParseCoberturaXML(content string) float64 {
	var report coberturaReport
	if err := xml.Unmarshal([]byte(content), &report); err != nil {
		return 0
	}
	rate, err := strconv.ParseFloat(report.LineRate, 64)
	if err != nil {
		return 0
	}
	return rate
}

// --- Tool runners for non-Go languages ---

// runESLintAt runs ESLint from the given module directory.
func (te *ToolExecutor) runESLintAt(dir string) *domain.Evidence {
	// Try npx eslint first (most common in projects)
	eslintCmd := findESLint(dir)
	if eslintCmd == "" {
		return nil
	}

	var cmd *exec.Cmd
	if eslintCmd == "npx" {
		cmd = exec.Command("npx", "eslint", "--format", "json", ".")
	} else {
		cmd = exec.Command(eslintCmd, "--format", "json", ".")
	}
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil && len(output) == 0 {
		return nil
	}

	result := ParseESLintJSON(string(output))
	te.adjustFindingPaths(result.Findings, dir)
	te.rawLintFindings = append(te.rawLintFindings, result.Findings...)
	ev := result.ToEvidence()
	return &ev
}

// findESLint detects ESLint availability in order of preference.
func findESLint(dir string) string {
	// Check for local eslint in node_modules
	localBin := filepath.Join(dir, "node_modules", ".bin", "eslint")
	if _, err := os.Stat(localBin); err == nil {
		return localBin
	}

	// Check for npx
	if _, err := exec.LookPath("npx"); err == nil {
		// Verify eslint is in package.json devDeps
		pkgJSON := filepath.Join(dir, "package.json")
		if data, err := os.ReadFile(pkgJSON); err == nil {
			if strings.Contains(string(data), "eslint") {
				return "npx"
			}
		}
	}

	// Check global eslint
	if path, err := exec.LookPath("eslint"); err == nil {
		return path
	}

	return ""
}

// runJSTestAt runs Jest or Vitest in the given directory.
func (te *ToolExecutor) runJSTestAt(dir string) *domain.Evidence {
	runner, args := detectJSTestRunner(dir)
	if runner == "" {
		return nil
	}

	cmd := exec.Command(runner, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "CI=true") // Jest uses CI mode for non-interactive
	output, _ := cmd.CombinedOutput()

	result := ParseJestJSON(string(output))
	if result.TotalCount == 0 {
		// Jest JSON might be embedded in other output — try to extract
		if idx := strings.Index(string(output), "{"); idx >= 0 {
			result = ParseJestJSON(string(output)[idx:])
		}
	}
	if result.TotalCount == 0 {
		return nil // couldn't parse
	}

	// Try to collect coverage
	result.Coverage = te.collectJSCoverage(dir, runner)

	ev := result.ToEvidence()
	return &ev
}

// detectJSTestRunner finds which test runner is configured.
func detectJSTestRunner(dir string) (string, []string) {
	pkgJSON := filepath.Join(dir, "package.json")
	data, err := os.ReadFile(pkgJSON)
	if err != nil {
		return "", nil
	}
	content := string(data)

	// Check for vitest
	if strings.Contains(content, "vitest") {
		if npx, err := exec.LookPath("npx"); err == nil {
			return npx, []string{"vitest", "run", "--reporter=json"}
		}
	}

	// Check for jest
	if strings.Contains(content, "jest") {
		if npx, err := exec.LookPath("npx"); err == nil {
			return npx, []string{"jest", "--json", "--no-coverage"}
		}
	}

	return "", nil
}

// collectJSCoverage attempts to collect JS coverage data.
func (te *ToolExecutor) collectJSCoverage(dir string, runner string) float64 {
	// Look for existing LCOV coverage file
	lcovPaths := []string{
		filepath.Join(dir, "coverage", "lcov.info"),
		filepath.Join(dir, "coverage", "lcov", "lcov.info"),
	}
	for _, p := range lcovPaths {
		if data, err := os.ReadFile(p); err == nil {
			return ParseLCOV(string(data))
		}
	}
	return 0
}

// runRuffAt runs Ruff linter in the given directory.
func (te *ToolExecutor) runRuffAt(dir string) *domain.Evidence {
	ruffPath, err := exec.LookPath("ruff")
	if err != nil {
		return nil
	}

	cmd := exec.Command(ruffPath, "check", "--output-format", "json", ".")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil && len(output) == 0 {
		return nil
	}

	result := ParseRuffJSON(string(output))
	te.adjustFindingPaths(result.Findings, dir)
	te.rawLintFindings = append(te.rawLintFindings, result.Findings...)
	ev := result.ToEvidence()
	return &ev
}

// runPytestAt runs pytest in the given directory.
func (te *ToolExecutor) runPytestAt(dir string) *domain.Evidence {
	pytestPath, err := exec.LookPath("pytest")
	if err != nil {
		// Try python -m pytest
		if _, err := exec.LookPath("python3"); err != nil {
			if _, err := exec.LookPath("python"); err != nil {
				return nil
			}
		}
	}

	// Generate JUnit XML for structured parsing
	junitFile := filepath.Join(os.TempDir(), fmt.Sprintf("certify-pytest-%d.xml", time.Now().UnixNano()))
	defer os.Remove(junitFile)

	var cmd *exec.Cmd
	if pytestPath != "" {
		cmd = exec.Command(pytestPath, "--tb=no", "-q", fmt.Sprintf("--junitxml=%s", junitFile))
	} else if python3, err := exec.LookPath("python3"); err == nil {
		cmd = exec.Command(python3, "-m", "pytest", "--tb=no", "-q", fmt.Sprintf("--junitxml=%s", junitFile))
	} else {
		python, _ := exec.LookPath("python")
		cmd = exec.Command(python, "-m", "pytest", "--tb=no", "-q", fmt.Sprintf("--junitxml=%s", junitFile))
	}
	cmd.Dir = dir
	output, _ := cmd.CombinedOutput()

	// Try JUnit XML first
	if data, err := os.ReadFile(junitFile); err == nil {
		result := ParsePytestJUnitXML(string(data))
		if result.TotalCount > 0 {
			result.Coverage = te.collectPythonCoverage(dir)
			ev := result.ToEvidence()
			return &ev
		}
	}

	// Fallback to parsing stdout
	result := ParsePytestShort(string(output))
	if result.TotalCount > 0 {
		result.Coverage = te.collectPythonCoverage(dir)
		ev := result.ToEvidence()
		return &ev
	}

	return nil
}

// collectPythonCoverage looks for existing coverage data.
func (te *ToolExecutor) collectPythonCoverage(dir string) float64 {
	// Check for coverage.xml (Cobertura format from pytest-cov)
	cobPath := filepath.Join(dir, "coverage.xml")
	if data, err := os.ReadFile(cobPath); err == nil {
		return ParseCoberturaXML(string(data))
	}
	// Check for htmlcov/status.json or .coverage
	return 0
}

// runCargoClippyAt runs cargo clippy in the given directory.
func (te *ToolExecutor) runCargoClippyAt(dir string) *domain.Evidence {
	if _, err := exec.LookPath("cargo"); err != nil {
		return nil
	}

	cmd := exec.Command("cargo", "clippy", "--message-format=json", "--", "-W", "clippy::all")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil && len(output) == 0 {
		return nil
	}

	result := ParseCargoClippyJSON(string(output))
	te.adjustFindingPaths(result.Findings, dir)
	te.rawLintFindings = append(te.rawLintFindings, result.Findings...)
	ev := result.ToEvidence()
	return &ev
}

// runCargoTestAt runs cargo test in the given directory.
func (te *ToolExecutor) runCargoTestAt(dir string) *domain.Evidence {
	if _, err := exec.LookPath("cargo"); err != nil {
		return nil
	}

	cmd := exec.Command("cargo", "test", "--", "--test-threads=1")
	cmd.Dir = dir
	output, _ := cmd.CombinedOutput()

	result := ParseCargoTestOutput(string(output))
	if result.TotalCount == 0 {
		return nil
	}

	ev := result.ToEvidence()
	return &ev
}
