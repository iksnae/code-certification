package engine_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/engine"
	"github.com/iksnae/code-certification/internal/evidence"
	"github.com/iksnae/code-certification/internal/policy"
	"github.com/iksnae/code-certification/internal/record"
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

func TestSaveReportArtifacts(t *testing.T) {
	certDir := t.TempDir()
	storeDir := filepath.Join(certDir, "records")
	store := record.NewStore(storeDir)

	// Save some sample records
	now := time.Now()
	for _, sym := range []string{"main", "helper", "init"} {
		rec := domain.CertificationRecord{
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
		}
		if err := store.Save(rec); err != nil {
			t.Fatalf("Save(%s) error: %v", sym, err)
		}
	}

	err := engine.SaveReportArtifacts(certDir, store, "test/repo", "abc123", now)
	if err != nil {
		t.Fatalf("SaveReportArtifacts() error: %v", err)
	}

	// Verify REPORT_CARD.md
	cardPath := filepath.Join(certDir, "REPORT_CARD.md")
	cardData, err := os.ReadFile(cardPath)
	if err != nil {
		t.Fatalf("REPORT_CARD.md not found: %v", err)
	}
	if len(cardData) == 0 {
		t.Error("REPORT_CARD.md is empty")
	}

	// Verify badge.json
	badgePath := filepath.Join(certDir, "badge.json")
	badgeData, err := os.ReadFile(badgePath)
	if err != nil {
		t.Fatalf("badge.json not found: %v", err)
	}
	var badgeObj map[string]interface{}
	if err := json.Unmarshal(badgeData, &badgeObj); err != nil {
		t.Fatalf("badge.json is not valid JSON: %v", err)
	}

	// Verify per-unit reports directory exists with files
	reportsDir := filepath.Join(certDir, "reports")
	entries, err := os.ReadDir(reportsDir)
	if err != nil {
		t.Fatalf("reports directory not found: %v", err)
	}
	if len(entries) == 0 {
		t.Error("reports directory should contain per-unit markdown files")
	}
}
