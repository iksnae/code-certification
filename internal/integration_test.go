package internal_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/code-certification/certify/internal/config"
	"github.com/code-certification/certify/internal/discovery"
	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/engine"
	"github.com/code-certification/certify/internal/evidence"
	"github.com/code-certification/certify/internal/override"
	"github.com/code-certification/certify/internal/record"
	"github.com/code-certification/certify/internal/report"
)

func TestE2E_GoSimpleRepo(t *testing.T) {
	root := filepath.Join("..", "testdata", "repos", "go-simple")
	certDir := t.TempDir()

	// 1. Discover units
	goAdapter := discovery.NewGoAdapter()
	units, err := goAdapter.Scan(root)
	if err != nil {
		t.Fatalf("scan error: %v", err)
	}
	if len(units) == 0 {
		t.Fatal("should discover at least one unit")
	}

	// Save index
	idx := discovery.NewIndex(units)
	indexPath := filepath.Join(certDir, "index.json")
	if err := idx.Save(indexPath); err != nil {
		t.Fatal(err)
	}

	// 2. Load policies
	policyDir := filepath.Join("..", "testdata", "policies")
	packs, err := config.LoadPolicyPacks(policyDir)
	if err != nil {
		t.Fatal(err)
	}

	// 3. Certify each unit
	expiryCfg := domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     7,
		MaxWindowDays:     365,
	}
	store := record.NewStore(filepath.Join(certDir, "records"))
	now := time.Now()

	var records []domain.CertificationRecord
	for _, unit := range units {
		matcher := config.NewPolicyMatcher(packs)
		matched := matcher.Match(unit)

		var rules []domain.PolicyRule
		for _, p := range matched {
			rules = append(rules, p.Rules...)
		}

		// Collect basic evidence
		src, _ := os.ReadFile(filepath.Join(root, unit.ID.Path()))
		metrics := evidence.ComputeMetrics(string(src))
		ev := []domain.Evidence{metrics.ToEvidence()}

		rec := engine.CertifyUnit(unit, rules, ev, expiryCfg, now)

		if err := store.Save(rec); err != nil {
			t.Fatalf("save record error: %v", err)
		}
		records = append(records, rec)
	}

	// 4. Verify records were saved
	allRecords, err := store.ListAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(allRecords) != len(units) {
		t.Errorf("saved %d records, expected %d", len(allRecords), len(units))
	}

	// 5. Generate health report
	h := report.Health(records)
	if h.TotalUnits != len(units) {
		t.Errorf("report total = %d, want %d", h.TotalUnits, len(units))
	}

	// 6. Verify text and JSON formatters work
	text := report.FormatText(h)
	if len(text) == 0 {
		t.Error("text report should not be empty")
	}
	jsonData, err := report.FormatJSON(h)
	if err != nil || len(jsonData) == 0 {
		t.Error("JSON report should not be empty")
	}
}

func TestE2E_TSSimpleRepo(t *testing.T) {
	root := filepath.Join("..", "testdata", "repos", "ts-simple")

	tsAdapter := discovery.NewTSAdapter()
	units, err := tsAdapter.Scan(root)
	if err != nil {
		t.Fatalf("scan error: %v", err)
	}

	if len(units) < 4 {
		t.Errorf("should discover at least 4 TS units, got %d", len(units))
	}

	// Verify stable IDs
	for _, u := range units {
		if u.ID.Language() != "ts" {
			t.Errorf("language = %q, want ts", u.ID.Language())
		}
	}
}

func TestE2E_AdvisoryMode(t *testing.T) {
	// Advisory mode should never return error even with failing units
	root := filepath.Join("..", "testdata", "repos", "go-simple")
	goAdapter := discovery.NewGoAdapter()
	units, _ := goAdapter.Scan(root)

	expiryCfg := domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     7,
		MaxWindowDays:     365,
	}

	// Use impossible rules to guarantee failure
	rules := []domain.PolicyRule{
		{ID: "impossible", Dimension: domain.DimCorrectness, Severity: domain.SeverityError,
			Metric: "lint_errors", Threshold: 0},
	}

	now := time.Now()
	var hasFailing bool
	for _, unit := range units {
		rec := engine.CertifyUnit(unit, rules, nil, expiryCfg, now)
		if !rec.Status.IsPassing() {
			hasFailing = true
		}
	}

	if !hasFailing {
		t.Error("should have failing units with impossible rules")
	}
	// In advisory mode, this is fine — no error returned
}

func TestE2E_WithOverrides(t *testing.T) {
	unit := domain.NewUnit(domain.NewUnitID("go", "legacy/old.go", "Deprecated"), domain.UnitTypeFunction)

	rules := []domain.PolicyRule{
		{ID: "lint-clean", Dimension: domain.DimCorrectness, Severity: domain.SeverityError,
			Metric: "lint_errors", Threshold: 0},
	}

	expiryCfg := domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     7,
		MaxWindowDays:     365,
	}

	rec := engine.CertifyUnit(unit, rules, nil, expiryCfg, time.Now())
	if rec.Status.IsPassing() {
		t.Error("unit should fail without evidence")
	}

	// Apply exempt override
	overrides := []domain.Override{
		{
			UnitID:    unit.ID,
			Action:    domain.OverrideExempt,
			Rationale: "Legacy code",
			Actor:     "admin",
		},
	}

	rec = override.ApplyAll(rec, overrides)
	if rec.Status != domain.StatusExempt {
		t.Errorf("after override, status = %v, want exempt", rec.Status)
	}
	if !rec.Status.IsPassing() {
		t.Error("exempt should be passing")
	}
}

func TestE2E_IndexDiff(t *testing.T) {
	root := filepath.Join("..", "testdata", "repos", "go-simple")

	goAdapter := discovery.NewGoAdapter()
	units1, _ := goAdapter.Scan(root)
	idx1 := discovery.NewIndex(units1)

	// Simulate adding a unit
	units2 := append(units1, domain.NewUnit(
		domain.NewUnitID("go", "new_file.go", "NewFunc"),
		domain.UnitTypeFunction,
	))
	idx2 := discovery.NewIndex(units2)

	diff := discovery.Diff(idx1, idx2)
	if len(diff.Added) != 1 {
		t.Errorf("Added = %d, want 1", len(diff.Added))
	}
	if len(diff.Removed) != 0 {
		t.Errorf("Removed = %d, want 0", len(diff.Removed))
	}
}
