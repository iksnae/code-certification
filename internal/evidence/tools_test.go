package evidence

import (
	"testing"
)

func TestParseESLintJSON_Clean(t *testing.T) {
	output := `[]`
	result := ParseESLintJSON(output)
	if result.ErrorCount != 0 || result.WarnCount != 0 {
		t.Errorf("clean ESLint: errors=%d, warnings=%d", result.ErrorCount, result.WarnCount)
	}
	if result.Tool != "eslint" {
		t.Errorf("tool = %q, want eslint", result.Tool)
	}
}

func TestParseESLintJSON_Findings(t *testing.T) {
	output := `[
		{
			"filePath": "/src/index.ts",
			"messages": [
				{"ruleId": "no-unused-vars", "severity": 2, "message": "'x' is defined but never used", "line": 5},
				{"ruleId": "prefer-const", "severity": 1, "message": "'y' is never reassigned", "line": 10}
			]
		},
		{
			"filePath": "/src/utils.ts",
			"messages": [
				{"ruleId": "no-console", "severity": 1, "message": "Unexpected console statement", "line": 3}
			]
		}
	]`

	result := ParseESLintJSON(output)
	if result.ErrorCount != 1 {
		t.Errorf("errors = %d, want 1", result.ErrorCount)
	}
	if result.WarnCount != 2 {
		t.Errorf("warnings = %d, want 2", result.WarnCount)
	}
	if len(result.Findings) != 3 {
		t.Errorf("findings = %d, want 3", len(result.Findings))
	}
	if result.Findings[0].Rule != "no-unused-vars" {
		t.Errorf("first rule = %q, want no-unused-vars", result.Findings[0].Rule)
	}
}

func TestParseESLintJSON_Empty(t *testing.T) {
	result := ParseESLintJSON("")
	if result.ErrorCount != 0 {
		t.Errorf("empty ESLint should have 0 errors, got %d", result.ErrorCount)
	}
}

func TestParseRuffJSON_Clean(t *testing.T) {
	result := ParseRuffJSON("[]")
	if result.ErrorCount != 0 || result.WarnCount != 0 {
		t.Errorf("clean ruff: errors=%d, warnings=%d", result.ErrorCount, result.WarnCount)
	}
	if result.Tool != "ruff" {
		t.Errorf("tool = %q, want ruff", result.Tool)
	}
}

func TestParseRuffJSON_Findings(t *testing.T) {
	output := `[
		{"code": "F401", "message": "os imported but unused", "filename": "main.py", "location": {"row": 1}},
		{"code": "E501", "message": "Line too long", "filename": "main.py", "location": {"row": 15}},
		{"code": "W291", "message": "Trailing whitespace", "filename": "utils.py", "location": {"row": 3}}
	]`

	result := ParseRuffJSON(output)
	if result.ErrorCount != 2 {
		t.Errorf("errors = %d, want 2 (F401+E501)", result.ErrorCount)
	}
	if result.WarnCount != 1 {
		t.Errorf("warnings = %d, want 1 (W291)", result.WarnCount)
	}
	if len(result.Findings) != 3 {
		t.Errorf("findings = %d, want 3", len(result.Findings))
	}
}

func TestParseCargoClippyJSON_Clean(t *testing.T) {
	// Cargo always emits a final "compiler-artifact" message
	output := `{"reason":"compiler-artifact","target":{"name":"mylib"}}
{"reason":"build-finished","success":true}`
	result := ParseCargoClippyJSON(output)
	if result.ErrorCount != 0 || result.WarnCount != 0 {
		t.Errorf("clean clippy: errors=%d, warnings=%d", result.ErrorCount, result.WarnCount)
	}
}

func TestParseCargoClippyJSON_Warnings(t *testing.T) {
	output := `{"reason":"compiler-message","message":{"code":{"code":"clippy::needless_return"},"level":"warning","message":"unneeded return statement","spans":[{"file_name":"src/main.rs","line_start":10}]}}`

	result := ParseCargoClippyJSON(output)
	if result.WarnCount != 1 {
		t.Errorf("warnings = %d, want 1", result.WarnCount)
	}
	if len(result.Findings) != 1 {
		t.Fatalf("findings = %d, want 1", len(result.Findings))
	}
	if result.Findings[0].Rule != "clippy::needless_return" {
		t.Errorf("rule = %q", result.Findings[0].Rule)
	}
	if result.Findings[0].File != "src/main.rs" {
		t.Errorf("file = %q", result.Findings[0].File)
	}
}

func TestParseJestJSON_Passing(t *testing.T) {
	output := `{"numTotalTests": 42, "numPassedTests": 42, "numFailedTests": 0, "numPendingTests": 0, "success": true}`
	result := ParseJestJSON(output)
	if result.TotalCount != 42 {
		t.Errorf("total = %d, want 42", result.TotalCount)
	}
	if result.PassedCount != 42 {
		t.Errorf("passed = %d, want 42", result.PassedCount)
	}
	if result.FailedCount != 0 {
		t.Errorf("failed = %d, want 0", result.FailedCount)
	}
}

func TestParseJestJSON_Failures(t *testing.T) {
	output := `{"numTotalTests": 10, "numPassedTests": 8, "numFailedTests": 2, "numPendingTests": 0, "success": false}`
	result := ParseJestJSON(output)
	if result.FailedCount != 2 {
		t.Errorf("failed = %d, want 2", result.FailedCount)
	}
}

func TestParseJestJSON_Empty(t *testing.T) {
	result := ParseJestJSON("")
	if result.TotalCount != 0 {
		t.Errorf("empty jest should have 0 total, got %d", result.TotalCount)
	}
}

func TestParsePytestJUnitXML(t *testing.T) {
	xml := `<?xml version="1.0" encoding="utf-8"?>
<testsuites>
  <testsuite name="pytest" tests="15" errors="1" failures="2" skipped="3">
  </testsuite>
</testsuites>`

	result := ParsePytestJUnitXML(xml)
	if result.TotalCount != 15 {
		t.Errorf("total = %d, want 15", result.TotalCount)
	}
	if result.FailedCount != 3 {
		t.Errorf("failed = %d, want 3 (1 error + 2 failures)", result.FailedCount)
	}
	if result.SkipCount != 3 {
		t.Errorf("skipped = %d, want 3", result.SkipCount)
	}
	if result.PassedCount != 9 {
		t.Errorf("passed = %d, want 9", result.PassedCount)
	}
}

func TestParsePytestJUnitXML_MultipleSuites(t *testing.T) {
	xml := `<?xml version="1.0" encoding="utf-8"?>
<testsuites>
  <testsuite name="tests.unit" tests="10" errors="0" failures="0" skipped="0"/>
  <testsuite name="tests.integration" tests="5" errors="0" failures="1" skipped="0"/>
</testsuites>`

	result := ParsePytestJUnitXML(xml)
	if result.TotalCount != 15 {
		t.Errorf("total = %d, want 15", result.TotalCount)
	}
	if result.FailedCount != 1 {
		t.Errorf("failed = %d, want 1", result.FailedCount)
	}
}

func TestParsePytestShort(t *testing.T) {
	output := `============================= test session starts ==============================
collected 20 items

tests/test_core.py ..........                                            [50%]
tests/test_utils.py ........FF                                           [100%]

=========================== short test summary info ============================
FAILED tests/test_utils.py::test_edge_case
FAILED tests/test_utils.py::test_boundary
==================== 18 passed, 2 failed in 1.23s ============================`

	result := ParsePytestShort(output)
	if result.PassedCount != 18 {
		t.Errorf("passed = %d, want 18", result.PassedCount)
	}
	if result.FailedCount != 2 {
		t.Errorf("failed = %d, want 2", result.FailedCount)
	}
	if result.TotalCount != 20 {
		t.Errorf("total = %d, want 20", result.TotalCount)
	}
}

func TestParsePytestShort_AllPassed(t *testing.T) {
	output := `========================= 5 passed in 0.12s =========================`
	result := ParsePytestShort(output)
	if result.PassedCount != 5 {
		t.Errorf("passed = %d, want 5", result.PassedCount)
	}
	if result.TotalCount != 5 {
		t.Errorf("total = %d, want 5", result.TotalCount)
	}
}

func TestParseCargoTestOutput(t *testing.T) {
	output := `running 3 tests
test tests::test_add ... ok
test tests::test_sub ... ok
test tests::test_mul ... FAILED

failures:
    tests::test_mul

test result: FAILED. 2 passed; 1 failed; 0 ignored; 0 measured; 0 filtered out`

	result := ParseCargoTestOutput(output)
	if result.PassedCount != 2 {
		t.Errorf("passed = %d, want 2", result.PassedCount)
	}
	if result.FailedCount != 1 {
		t.Errorf("failed = %d, want 1", result.FailedCount)
	}
	if result.TotalCount != 3 {
		t.Errorf("total = %d, want 3", result.TotalCount)
	}
}

func TestParseCargoTestOutput_AllPassing(t *testing.T) {
	output := `running 10 tests
test tests::test_1 ... ok
test tests::test_2 ... ok
...
test result: ok. 10 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.5s`

	result := ParseCargoTestOutput(output)
	if result.PassedCount != 10 {
		t.Errorf("passed = %d, want 10", result.PassedCount)
	}
	if result.FailedCount != 0 {
		t.Errorf("failed = %d, want 0", result.FailedCount)
	}
}

func TestParseCargoTestOutput_MultiCrate(t *testing.T) {
	output := `running 5 tests
test result: ok. 5 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out

running 3 tests
test result: ok. 3 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out`

	result := ParseCargoTestOutput(output)
	if result.PassedCount != 8 {
		t.Errorf("passed = %d, want 8 (5+3)", result.PassedCount)
	}
	if result.TotalCount != 8 {
		t.Errorf("total = %d, want 8", result.TotalCount)
	}
}

func TestParseLCOV(t *testing.T) {
	lcov := `SF:src/index.ts
DA:1,1
DA:2,1
DA:3,0
LF:3
LH:2
end_of_record
SF:src/utils.ts
DA:1,1
DA:2,0
LF:2
LH:1
end_of_record`

	coverage := ParseLCOV(lcov)
	if coverage != 0.6 { // 3 hit out of 5
		t.Errorf("coverage = %f, want 0.6", coverage)
	}
}

func TestParseLCOV_Empty(t *testing.T) {
	if coverage := ParseLCOV(""); coverage != 0 {
		t.Errorf("empty lcov should be 0, got %f", coverage)
	}
}

func TestParseCoberturaXML(t *testing.T) {
	xml := `<?xml version="1.0" ?>
<coverage version="6.5" timestamp="1234" lines-valid="100" lines-covered="85" line-rate="0.85" branch-rate="0">
</coverage>`

	coverage := ParseCoberturaXML(xml)
	if coverage != 0.85 {
		t.Errorf("coverage = %f, want 0.85", coverage)
	}
}

func TestParseCoberturaXML_Empty(t *testing.T) {
	if coverage := ParseCoberturaXML(""); coverage != 0 {
		t.Errorf("empty cobertura should be 0, got %f", coverage)
	}
}

// Test evidence conversion
func TestESLintResult_ToEvidence(t *testing.T) {
	result := ParseESLintJSON(`[{"filePath":"a.ts","messages":[{"ruleId":"r1","severity":2,"message":"bad","line":1}]}]`)
	ev := result.ToEvidence()
	if ev.Passed {
		t.Error("should not pass with errors")
	}
	if ev.Metrics["lint_errors"] != 1 {
		t.Errorf("lint_errors = %f, want 1", ev.Metrics["lint_errors"])
	}
}

func TestCargoTestResult_ToEvidence(t *testing.T) {
	result := ParseCargoTestOutput("test result: ok. 5 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out")
	ev := result.ToEvidence()
	if !ev.Passed {
		t.Error("should pass with 0 failures")
	}
	if ev.Metrics["test_passed"] != 5 {
		t.Errorf("test_passed = %f, want 5", ev.Metrics["test_passed"])
	}
}
