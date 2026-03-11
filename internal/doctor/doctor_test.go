package doctor

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunAll_EmptyDir(t *testing.T) {
	root := t.TempDir()

	r := RunAll(root)

	if r == nil {
		t.Fatal("expected non-nil report")
	}
	if r.Root != root {
		t.Errorf("Root = %q, want %q", r.Root, root)
	}

	// Should have environment checks (go, git) regardless
	var envChecks int
	for _, c := range r.Checks {
		if c.Group == "environment" {
			envChecks++
		}
	}
	if envChecks < 2 {
		t.Errorf("expected at least 2 environment checks, got %d", envChecks)
	}

	// Should have a failure for missing config
	found := false
	for _, c := range r.Checks {
		if c.Name == "Configuration file" && c.Status == StatusFail {
			found = true
			if c.Fix == "" {
				t.Error("missing config check should have a fix suggestion")
			}
		}
	}
	if !found {
		t.Error("expected a failing check for missing .certification/config.yml")
	}
}

func TestRunAll_InitializedProject(t *testing.T) {
	root := t.TempDir()
	certDir := filepath.Join(root, ".certification")
	os.MkdirAll(filepath.Join(certDir, "policies"), 0o755)
	os.MkdirAll(filepath.Join(certDir, "records"), 0o755)

	// Write minimal config
	configContent := `mode: advisory
scope:
  exclude:
    - "vendor/**"
expiry:
  default_window_days: 90
  min_window_days: 7
  max_window_days: 365
`
	os.WriteFile(filepath.Join(certDir, "config.yml"), []byte(configContent), 0o644)

	// Write a policy pack
	policyContent := `name: test-policy
version: "1.0.0"
rules:
  - id: test-rule
    dimension: correctness
    description: "test"
    severity: error
    metric: lint_errors
    threshold: 0
`
	os.WriteFile(filepath.Join(certDir, "policies", "test.yml"), []byte(policyContent), 0o644)

	r := RunAll(root)

	// Config should pass
	var configCheck *Check
	for i, c := range r.Checks {
		if c.Name == "Configuration file" {
			configCheck = &r.Checks[i]
			break
		}
	}
	if configCheck == nil {
		t.Fatal("expected Configuration file check")
	}
	if configCheck.Status != StatusPass {
		t.Errorf("config check status = %v, want pass", configCheck.Status)
	}

	// Policy should pass
	var policyFound bool
	for _, c := range r.Checks {
		if strings.HasPrefix(c.Name, "Policy:") && c.Status == StatusPass {
			policyFound = true
		}
	}
	if !policyFound {
		t.Error("expected a passing policy check")
	}

	// Index should warn (not yet scanned)
	var indexCheck *Check
	for i, c := range r.Checks {
		if c.Name == "Unit index" {
			indexCheck = &r.Checks[i]
			break
		}
	}
	if indexCheck == nil {
		t.Fatal("expected Unit index check")
	}
	if indexCheck.Status != StatusWarn {
		t.Errorf("index check status = %v, want warn", indexCheck.Status)
	}
}

func TestRunAll_WithRecords(t *testing.T) {
	root := t.TempDir()
	certDir := filepath.Join(root, ".certification")
	os.MkdirAll(filepath.Join(certDir, "records"), 0o755)
	os.MkdirAll(filepath.Join(certDir, "policies"), 0o755)

	os.WriteFile(filepath.Join(certDir, "config.yml"), []byte("mode: advisory\nexpiry:\n  default_window_days: 90\n  min_window_days: 7\n  max_window_days: 365\n"), 0o644)
	os.WriteFile(filepath.Join(certDir, "index.json"), []byte("{}"), 0o644)
	os.WriteFile(filepath.Join(certDir, "records", "abc123.json"), []byte("{}"), 0o644)
	os.WriteFile(filepath.Join(certDir, "records", "def456.json"), []byte("{}"), 0o644)
	os.WriteFile(filepath.Join(certDir, "REPORT_CARD.md"), []byte("# Report"), 0o644)
	os.WriteFile(filepath.Join(certDir, "badge.json"), []byte("{}"), 0o644)

	r := RunAll(root)

	// Records should pass with count
	var recordCheck *Check
	for i, c := range r.Checks {
		if c.Name == "Certification records" {
			recordCheck = &r.Checks[i]
			break
		}
	}
	if recordCheck == nil {
		t.Fatal("expected Certification records check")
	}
	if recordCheck.Status != StatusPass {
		t.Errorf("records check status = %v, want pass", recordCheck.Status)
	}
	if !strings.Contains(recordCheck.Message, "2") {
		t.Errorf("records message = %q, expected to mention 2 records", recordCheck.Message)
	}

	// Report card should pass
	var reportCheck *Check
	for i, c := range r.Checks {
		if c.Name == "Report card" {
			reportCheck = &r.Checks[i]
			break
		}
	}
	if reportCheck != nil && reportCheck.Status != StatusPass {
		t.Errorf("report card check status = %v, want pass", reportCheck.Status)
	}
}

func TestRunAll_InvalidConfig(t *testing.T) {
	root := t.TempDir()
	certDir := filepath.Join(root, ".certification")
	os.MkdirAll(certDir, 0o755)

	// Write invalid YAML
	os.WriteFile(filepath.Join(certDir, "config.yml"), []byte("{{invalid yaml"), 0o644)

	r := RunAll(root)

	var configValid *Check
	for i, c := range r.Checks {
		if c.Name == "Configuration valid" {
			configValid = &r.Checks[i]
			break
		}
	}
	if configValid == nil {
		t.Fatal("expected Configuration valid check")
	}
	if configValid.Status != StatusFail {
		t.Errorf("config valid status = %v, want fail", configValid.Status)
	}
}

func TestSummary(t *testing.T) {
	r := &Report{
		Checks: []Check{
			{Status: StatusPass},
			{Status: StatusPass},
			{Status: StatusWarn},
			{Status: StatusFail},
			{Status: StatusSkip},
		},
	}

	pass, warn, fail, skip := r.Summary()
	if pass != 2 {
		t.Errorf("pass = %d, want 2", pass)
	}
	if warn != 1 {
		t.Errorf("warn = %d, want 1", warn)
	}
	if fail != 1 {
		t.Errorf("fail = %d, want 1", fail)
	}
	if skip != 1 {
		t.Errorf("skip = %d, want 1", skip)
	}
}

func TestHasFailures(t *testing.T) {
	r := &Report{
		Checks: []Check{
			{Status: StatusPass},
			{Status: StatusWarn},
		},
	}
	if r.HasFailures() {
		t.Error("expected no failures")
	}

	r.Checks = append(r.Checks, Check{Status: StatusFail})
	if !r.HasFailures() {
		t.Error("expected failures")
	}
}

func TestFormatReport_ContainsSections(t *testing.T) {
	r := &Report{
		Root: "/tmp/test",
		Checks: []Check{
			{Name: "Go compiler", Group: "environment", Status: StatusPass, Message: "go version go1.22"},
			{Name: "Git", Group: "environment", Status: StatusPass, Message: "git version 2.43"},
			{Name: "Configuration file", Group: "project", Status: StatusFail, Message: "not found", Fix: "Run: certify init"},
			{Name: "golangci-lint", Group: "tools", Status: StatusWarn, Message: "not found", Fix: "Install it"},
		},
	}

	output := FormatReport(r)

	required := []string{
		"Certify Doctor",
		"Environment",
		"Project Setup",
		"Optional Tools",
		"✅",
		"❌",
		"⚠️",
		"Run: certify init",
		"Summary",
		"2 passed",
		"1 warnings",
		"1 failed",
	}
	for _, s := range required {
		if !strings.Contains(output, s) {
			t.Errorf("output missing %q", s)
		}
	}
}

func TestCheckStatus_String(t *testing.T) {
	tests := []struct {
		status CheckStatus
		want   string
	}{
		{StatusPass, "pass"},
		{StatusWarn, "warn"},
		{StatusFail, "fail"},
		{StatusSkip, "skip"},
	}
	for _, tt := range tests {
		if got := tt.status.String(); got != tt.want {
			t.Errorf("CheckStatus(%d).String() = %q, want %q", tt.status, got, tt.want)
		}
	}
}

func TestCheckAnalysisTiers(t *testing.T) {
	r := &Report{Root: "/tmp/test"}
	r.checkAnalysisTiers()

	// Should have at least 4 checks (Go + TS + Py + Rs)
	if len(r.Checks) < 4 {
		t.Errorf("expected >= 4 analysis tier checks, got %d", len(r.Checks))
	}

	// Go should always be Tier 2
	goCheck := r.Checks[0]
	if goCheck.Status != StatusPass {
		t.Errorf("Go analysis should be pass, got %s", goCheck.Status)
	}
	if !strings.Contains(goCheck.Message, "Tier 2") {
		t.Errorf("Go analysis message should mention Tier 2, got %q", goCheck.Message)
	}

	// All checks should be in the "analysis" group
	for _, c := range r.Checks {
		if c.Group != "analysis" {
			t.Errorf("check %q has group %q, want analysis", c.Name, c.Group)
		}
	}
}

func TestFormatReport_AnalysisTierSection(t *testing.T) {
	r := &Report{
		Root: "/tmp/test",
		Checks: []Check{
			{Name: "Go analysis", Group: "analysis", Status: StatusPass, Message: "Tier 2 (go/types)"},
			{Name: "TypeScript analysis", Group: "analysis", Status: StatusWarn, Message: "Tier 1", Fix: "npm i -g typescript-language-server"},
		},
	}
	output := FormatReport(r)
	if !strings.Contains(output, "Analysis Tiers") {
		t.Errorf("output missing Analysis Tiers section")
	}
}

func TestCheckStatus_Emoji(t *testing.T) {
	if StatusPass.Emoji() != "✅" {
		t.Errorf("pass emoji = %q, want ✅", StatusPass.Emoji())
	}
	if StatusFail.Emoji() != "❌" {
		t.Errorf("fail emoji = %q, want ❌", StatusFail.Emoji())
	}
}
