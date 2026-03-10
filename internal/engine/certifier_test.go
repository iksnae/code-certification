package engine_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/engine"
	"github.com/iksnae/code-certification/internal/evidence"
	"github.com/iksnae/code-certification/internal/policy"
	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/report"
)

func testPacks() []domain.PolicyPack {
	return []domain.PolicyPack{
		{
			Name:    "test-global",
			Version: "1.0",
			Rules: []domain.PolicyRule{
				{
					ID:        "lint-clean",
					Dimension: domain.DimCorrectness,
					Severity:  domain.SeverityError,
					Metric:    "lint_errors",
					Threshold: 0,
				},
			},
		},
	}
}

func testExpiryCfg() domain.ExpiryConfig {
	return domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     7,
		MaxWindowDays:     365,
	}
}

func TestCertifier_Certify_Basic(t *testing.T) {
	root := filepath.Join("..", "..", "testdata", "repos", "go-simple")

	certifier := &engine.Certifier{
		Root:      root,
		Matcher:   policy.NewMatcher(testPacks()),
		ExpiryCfg: testExpiryCfg(),
	}

	unit := domain.NewUnit(
		domain.NewUnitID("go", "main.go", "main"),
		domain.UnitTypeFunction,
	)

	repoEv := []domain.Evidence{
		evidence.LintResult{Tool: "go vet", ErrorCount: 0}.ToEvidence(),
	}

	result, err := certifier.Certify(context.Background(), unit, repoEv, time.Now())
	if err != nil {
		t.Fatalf("Certify() error: %v", err)
	}
	if result == nil {
		t.Fatal("Certify() returned nil result")
	}

	rec := result.Record
	if rec.UnitID.String() != "go://main.go#main" {
		t.Errorf("UnitID = %s, want go://main.go#main", rec.UnitID)
	}
	if rec.Status != domain.StatusCertified {
		t.Errorf("Status = %v, want certified", rec.Status)
	}
	if rec.Source != "deterministic" {
		t.Errorf("Source = %q, want deterministic", rec.Source)
	}
	if result.AgentReview != nil {
		t.Error("AgentReview should be nil when no agent configured")
	}

	// Should have repo evidence (lint) + per-unit metrics
	if len(rec.Evidence) < 2 {
		t.Errorf("Evidence count = %d, want >= 2 (repo lint + unit metrics)", len(rec.Evidence))
	}

	// Verify metrics evidence was added (from reading the source file)
	hasMetrics := false
	for _, e := range rec.Evidence {
		if e.Kind == domain.EvidenceKindMetrics {
			hasMetrics = true
		}
	}
	if !hasMetrics {
		t.Error("should include metrics evidence from source file")
	}
}

func TestCertifier_Certify_SavesRecord(t *testing.T) {
	root := filepath.Join("..", "..", "testdata", "repos", "go-simple")
	storeDir := t.TempDir()
	store := record.NewStore(storeDir)

	certifier := &engine.Certifier{
		Root:      root,
		Store:     store,
		Matcher:   policy.NewMatcher(testPacks()),
		ExpiryCfg: testExpiryCfg(),
	}

	unit := domain.NewUnit(
		domain.NewUnitID("go", "main.go", "helper"),
		domain.UnitTypeFunction,
	)

	repoEv := []domain.Evidence{
		evidence.LintResult{Tool: "go vet", ErrorCount: 0}.ToEvidence(),
	}

	result, err := certifier.Certify(context.Background(), unit, repoEv, time.Now())
	if err != nil {
		t.Fatalf("Certify() error: %v", err)
	}
	if result == nil {
		t.Fatal("Certify() returned nil result")
	}

	// Record should be loadable from store
	loaded, err := store.Load(result.Record.UnitID)
	if err != nil {
		t.Fatalf("Store.Load() error: %v", err)
	}
	if loaded.UnitID.String() != "go://main.go#helper" {
		t.Errorf("loaded UnitID = %s, want go://main.go#helper", loaded.UnitID)
	}
	if loaded.Score != result.Record.Score {
		t.Errorf("loaded Score = %f, want %f", loaded.Score, result.Record.Score)
	}

	// History should have been appended
	history, err := store.LoadHistory(result.Record.UnitID)
	if err != nil {
		t.Fatalf("LoadHistory() error: %v", err)
	}
	if len(history) == 0 {
		t.Error("history should have at least one entry")
	}
}

func TestCertifier_Certify_AppliesOverrides(t *testing.T) {
	root := filepath.Join("..", "..", "testdata", "repos", "go-simple")

	unitID := domain.NewUnitID("go", "main.go", "main")
	unit := domain.NewUnit(unitID, domain.UnitTypeFunction)

	certifier := &engine.Certifier{
		Root:      root,
		Matcher:   policy.NewMatcher(testPacks()),
		ExpiryCfg: testExpiryCfg(),
		Overrides: []domain.Override{
			{
				UnitID:    unitID,
				Action:    domain.OverrideExempt,
				Rationale: "Legacy code",
				Actor:     "admin",
			},
		},
	}

	result, err := certifier.Certify(context.Background(), unit, nil, time.Now())
	if err != nil {
		t.Fatalf("Certify() error: %v", err)
	}
	if result == nil {
		t.Fatal("Certify() returned nil result")
	}
	if result.Record.Status != domain.StatusExempt {
		t.Errorf("Status = %v, want exempt (override should apply)", result.Record.Status)
	}
}

func TestCertifier_Certify_NoStore(t *testing.T) {
	root := filepath.Join("..", "..", "testdata", "repos", "go-simple")

	certifier := &engine.Certifier{
		Root:      root,
		Store:     nil, // explicitly nil
		ExpiryCfg: testExpiryCfg(),
	}

	unit := domain.NewUnit(
		domain.NewUnitID("go", "main.go", "main"),
		domain.UnitTypeFunction,
	)

	result, err := certifier.Certify(context.Background(), unit, nil, time.Now())
	if err != nil {
		t.Fatalf("Certify() with nil store should not error: %v", err)
	}
	if result == nil {
		t.Fatal("Certify() returned nil result")
	}
	if result.Record.UnitID.String() != "go://main.go#main" {
		t.Errorf("UnitID = %s, want go://main.go#main", result.Record.UnitID)
	}
}

func TestCertifier_Certify_ComputesSymbolMetrics(t *testing.T) {
	// Create a temp Go source file with a known function
	tmpDir := t.TempDir()
	src := `package example

func Add(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a + b
}
`
	if err := os.WriteFile(filepath.Join(tmpDir, "math.go"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	certifier := &engine.Certifier{
		Root:      tmpDir,
		ExpiryCfg: testExpiryCfg(),
	}

	unit := domain.NewUnit(
		domain.NewUnitID("go", "math.go", "Add"),
		domain.UnitTypeFunction,
	)

	result, err := certifier.Certify(context.Background(), unit, nil, time.Now())
	if err != nil {
		t.Fatalf("Certify() error: %v", err)
	}
	if result == nil {
		t.Fatal("Certify() returned nil result")
	}

	// Find the metrics evidence
	var metricsEv *domain.Evidence
	for i, e := range result.Record.Evidence {
		if e.Kind == domain.EvidenceKindMetrics {
			metricsEv = &result.Record.Evidence[i]
			break
		}
	}
	if metricsEv == nil {
		t.Fatal("no metrics evidence found")
	}

	// Symbol-level metrics should show the function's complexity
	if metricsEv.Source != "metrics" {
		t.Errorf("metrics Source = %q, want 'metrics'", metricsEv.Source)
	}
	// Summary should mention complexity (ComputeSymbolMetrics extracts it)
	if metricsEv.Summary == "" {
		t.Error("metrics Summary should not be empty")
	}
}

func TestCertifier_Certify_ComputesFileMetrics(t *testing.T) {
	// Create a temp non-Go source file
	tmpDir := t.TempDir()
	src := `export function greet(name: string): string {
  return "Hello, " + name;
}

export function add(a: number, b: number): number {
  return a + b;
}
`
	if err := os.WriteFile(filepath.Join(tmpDir, "utils.ts"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	certifier := &engine.Certifier{
		Root:      tmpDir,
		ExpiryCfg: testExpiryCfg(),
	}

	unit := domain.NewUnit(
		domain.NewUnitID("ts", "utils.ts", "greet"),
		domain.UnitTypeFunction,
	)

	result, err := certifier.Certify(context.Background(), unit, nil, time.Now())
	if err != nil {
		t.Fatalf("Certify() error: %v", err)
	}
	if result == nil {
		t.Fatal("Certify() returned nil result")
	}

	// Find the metrics evidence — should use file-level ComputeMetrics (not symbol)
	var metricsEv *domain.Evidence
	for i, e := range result.Record.Evidence {
		if e.Kind == domain.EvidenceKindMetrics {
			metricsEv = &result.Record.Evidence[i]
			break
		}
	}
	if metricsEv == nil {
		t.Fatal("no metrics evidence found for .ts file")
	}
	if metricsEv.Source != "metrics" {
		t.Errorf("metrics Source = %q, want 'metrics'", metricsEv.Source)
	}
}

func TestCertifier_Certify_StructuralEvidence(t *testing.T) {
	// Create a temp Go source file with a known exported function
	tmpDir := t.TempDir()
	src := `package example

// Add adds two integers and returns the result.
func Add(a, b int) int {
	return a + b
}
`
	if err := os.WriteFile(filepath.Join(tmpDir, "math.go"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	certifier := &engine.Certifier{
		Root:      tmpDir,
		ExpiryCfg: testExpiryCfg(),
	}

	unit := domain.NewUnit(
		domain.NewUnitID("go", "math.go", "Add"),
		domain.UnitTypeFunction,
	)

	result, err := certifier.Certify(context.Background(), unit, nil, time.Now())
	if err != nil {
		t.Fatalf("Certify() error: %v", err)
	}

	// Find structural evidence
	var structEv *domain.Evidence
	for i, e := range result.Record.Evidence {
		if e.Kind == domain.EvidenceKindStructural {
			structEv = &result.Record.Evidence[i]
			break
		}
	}
	if structEv == nil {
		t.Fatal("no structural evidence found for Go function")
	}
	if structEv.Source != "structural" {
		t.Errorf("Source = %q, want structural", structEv.Source)
	}
	if structEv.Metrics["has_doc_comment"] != 1.0 {
		t.Errorf("has_doc_comment = %f, want 1.0", structEv.Metrics["has_doc_comment"])
	}
	if structEv.Metrics["param_count"] != 2 {
		t.Errorf("param_count = %f, want 2", structEv.Metrics["param_count"])
	}
	if structEv.Metrics["exported_name"] != 1.0 {
		t.Errorf("exported_name = %f, want 1.0", structEv.Metrics["exported_name"])
	}
}

func TestCertifier_Certify_NoStructuralForNonGo(t *testing.T) {
	tmpDir := t.TempDir()
	src := `export function greet() { return "hi"; }`
	if err := os.WriteFile(filepath.Join(tmpDir, "greet.ts"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	certifier := &engine.Certifier{
		Root:      tmpDir,
		ExpiryCfg: testExpiryCfg(),
	}

	unit := domain.NewUnit(
		domain.NewUnitID("ts", "greet.ts", "greet"),
		domain.UnitTypeFunction,
	)

	result, err := certifier.Certify(context.Background(), unit, nil, time.Now())
	if err != nil {
		t.Fatalf("Certify() error: %v", err)
	}

	for _, e := range result.Record.Evidence {
		if e.Kind == domain.EvidenceKindStructural {
			t.Error("non-Go unit should not have structural evidence")
		}
	}
}

func TestCertifier_Certify_PerUnitLintAttribution(t *testing.T) {
	tmpDir := t.TempDir()
	src := `package example

func Foo() {}
`
	if err := os.WriteFile(filepath.Join(tmpDir, "foo.go"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	certifier := &engine.Certifier{
		Root:      tmpDir,
		ExpiryCfg: testExpiryCfg(),
	}

	unit := domain.NewUnit(
		domain.NewUnitID("go", "foo.go", "Foo"),
		domain.UnitTypeFunction,
	)

	// Simulate repo evidence with lint findings for a different file
	repoEv := []domain.Evidence{
		evidence.LintResult{
			Tool: "golangci-lint", ErrorCount: 1,
			Findings: []evidence.LintFinding{
				{File: "bar.go", Line: 5, Message: "unused", Severity: "error"},
			},
		}.ToEvidence(),
	}

	// Set lint findings on the certifier for per-unit attribution
	certifier.RepoLintFindings = []evidence.LintFinding{
		{File: "bar.go", Line: 5, Message: "unused", Severity: "error"},
	}

	result, err := certifier.Certify(context.Background(), unit, repoEv, time.Now())
	if err != nil {
		t.Fatalf("Certify() error: %v", err)
	}

	// Check that per-unit lint evidence is clean (findings were for bar.go, not foo.go)
	for _, e := range result.Record.Evidence {
		if e.Kind == domain.EvidenceKindLint && strings.Contains(e.Source, ":unit") {
			if !e.Passed {
				t.Error("per-unit lint for foo.go should pass (findings are in bar.go)")
			}
			if e.Metrics["unit_lint_errors"] != 0 {
				t.Errorf("unit_lint_errors = %f, want 0", e.Metrics["unit_lint_errors"])
			}
		}
	}
}

func TestCertifier_PopulatesRunID(t *testing.T) {
	root := filepath.Join("..", "..", "testdata", "repos", "go-simple")

	certifier := &engine.Certifier{
		Root:      root,
		Matcher:   policy.NewMatcher(testPacks()),
		ExpiryCfg: testExpiryCfg(),
		RunID:     "run-20260310T155227Z",
	}

	unit := domain.NewUnit(
		domain.NewUnitID("go", "main.go", "main"),
		domain.UnitTypeFunction,
	)

	repoEv := []domain.Evidence{
		evidence.LintResult{Tool: "go vet", ErrorCount: 0}.ToEvidence(),
	}

	result, err := certifier.Certify(context.Background(), unit, repoEv, time.Now())
	if err != nil {
		t.Fatalf("Certify() error: %v", err)
	}
	if result.Record.RunID != "run-20260310T155227Z" {
		t.Errorf("RunID = %q, want run-20260310T155227Z", result.Record.RunID)
	}
}

func TestCertifier_PopulatesPolicyVersion(t *testing.T) {
	root := filepath.Join("..", "..", "testdata", "repos", "go-simple")

	certifier := &engine.Certifier{
		Root:           root,
		Matcher:        policy.NewMatcher(testPacks()),
		ExpiryCfg:      testExpiryCfg(),
		RunID:          "run-test",
		PolicyVersions: []string{"test-global@1.0", "go-strict@2.1"},
	}

	unit := domain.NewUnit(
		domain.NewUnitID("go", "main.go", "main"),
		domain.UnitTypeFunction,
	)

	result, err := certifier.Certify(context.Background(), unit, nil, time.Now())
	if err != nil {
		t.Fatalf("Certify() error: %v", err)
	}
	want := "test-global@1.0,go-strict@2.1"
	if result.Record.PolicyVersion != want {
		t.Errorf("PolicyVersion = %q, want %q", result.Record.PolicyVersion, want)
	}
}

func sampleRecords(now time.Time) []domain.CertificationRecord {
	var recs []domain.CertificationRecord
	for _, sym := range []string{"main", "helper", "init"} {
		recs = append(recs, domain.CertificationRecord{
			UnitID:      domain.NewUnitID("go", "main.go", sym),
			UnitType:    domain.UnitTypeFunction,
			UnitPath:    "main.go",
			Status:      domain.StatusCertified,
			Grade:       domain.GradeB,
			Score:       0.85,
			Confidence:  1.0,
			CertifiedAt: now,
			ExpiresAt:   now.Add(90 * 24 * time.Hour),
			Source:      "deterministic",
			Version:     1,
		})
	}
	return recs
}

func verifyArtifacts(t *testing.T, certDir string) {
	t.Helper()

	// Verify REPORT_CARD.md
	cardData, err := os.ReadFile(filepath.Join(certDir, "REPORT_CARD.md"))
	if err != nil {
		t.Fatalf("REPORT_CARD.md not found: %v", err)
	}
	if len(cardData) == 0 {
		t.Error("REPORT_CARD.md is empty")
	}

	// Verify badge.json
	badgeData, err := os.ReadFile(filepath.Join(certDir, "badge.json"))
	if err != nil {
		t.Fatalf("badge.json not found: %v", err)
	}
	var badgeObj map[string]interface{}
	if err := json.Unmarshal(badgeData, &badgeObj); err != nil {
		t.Fatalf("badge.json is not valid JSON: %v", err)
	}

	// Verify per-unit reports directory
	entries, err := os.ReadDir(filepath.Join(certDir, "reports"))
	if err != nil {
		t.Fatalf("reports directory not found: %v", err)
	}
	if len(entries) == 0 {
		t.Error("reports directory should contain per-unit markdown files")
	}
}

func TestSaveReportArtifacts(t *testing.T) {
	certDir := t.TempDir()
	now := time.Now()
	recs := sampleRecords(now)

	fr := report.GenerateFullReport(recs, "test/repo", "abc123", now)
	if err := engine.SaveReportArtifacts(certDir, fr); err != nil {
		t.Fatalf("SaveReportArtifacts() error: %v", err)
	}

	verifyArtifacts(t, certDir)
}

func TestSaveReportArtifactsFromStore(t *testing.T) {
	certDir := t.TempDir()
	storeDir := filepath.Join(certDir, "records")
	store := record.NewStore(storeDir)

	now := time.Now()
	for _, rec := range sampleRecords(now) {
		if err := store.Save(rec); err != nil {
			t.Fatalf("Save() error: %v", err)
		}
	}

	if err := engine.SaveReportArtifactsFromStore(certDir, store, "test/repo", "abc123", now); err != nil {
		t.Fatalf("SaveReportArtifactsFromStore() error: %v", err)
	}

	verifyArtifacts(t, certDir)
}
