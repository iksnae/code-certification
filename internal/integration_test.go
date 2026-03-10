package internal_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/config"
	"github.com/iksnae/code-certification/internal/discovery"
	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/engine"
	"github.com/iksnae/code-certification/internal/evidence"
	"github.com/iksnae/code-certification/internal/override"
	"github.com/iksnae/code-certification/internal/policy"
	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/report"
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

// TestE2E_Certifier proves the engine.Certifier is usable as a library
// without any CLI dependency.
func TestE2E_Certifier(t *testing.T) {
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

	// 2. Load policies
	policyDir := filepath.Join("..", "testdata", "policies")
	packs, err := config.LoadPolicyPacks(policyDir)
	if err != nil {
		t.Fatal(err)
	}

	// 3. Construct Certifier (library usage — no CLI)
	store := record.NewStore(filepath.Join(certDir, "records"))
	certifier := &engine.Certifier{
		Root:    root,
		Store:   store,
		Matcher: policy.NewMatcher(packs),
		ExpiryCfg: domain.ExpiryConfig{
			DefaultWindowDays: 90,
			MinWindowDays:     7,
			MaxWindowDays:     365,
		},
	}

	// 4. Certify each unit via the Certifier service
	now := time.Now()
	ctx := context.Background()
	var results []*engine.CertifyResult
	for _, unit := range units {
		result, err := certifier.Certify(ctx, unit, nil, now)
		if err != nil {
			t.Fatalf("Certify(%s) error: %v", unit.ID, err)
		}
		results = append(results, result)
	}

	// 5. Verify all records were saved
	allRecords, err := store.ListAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(allRecords) != len(units) {
		t.Errorf("saved %d records, expected %d", len(allRecords), len(units))
	}

	// 6. Verify records are loadable individually
	for _, result := range results {
		loaded, err := store.Load(result.Record.UnitID)
		if err != nil {
			t.Errorf("Load(%s) error: %v", result.Record.UnitID, err)
			continue
		}
		if loaded.Score != result.Record.Score {
			t.Errorf("Score mismatch for %s: loaded=%f, result=%f",
				result.Record.UnitID, loaded.Score, result.Record.Score)
		}
	}

	// 7. Verify SaveReportArtifactsFromStore works from the library
	if err := engine.SaveReportArtifactsFromStore(certDir, store, "test/repo", "abc123", now); err != nil {
		t.Fatalf("SaveReportArtifactsFromStore() error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(certDir, "REPORT_CARD.md")); err != nil {
		t.Error("REPORT_CARD.md not generated")
	}
	if _, err := os.Stat(filepath.Join(certDir, "badge.json")); err != nil {
		t.Error("badge.json not generated")
	}
}

func TestE2E_StructuralEvidence(t *testing.T) {
	// Create a temp Go repo with functions that have known structural properties
	tmpDir := t.TempDir()

	// Well-structured function: documented, few params, low nesting
	goodSrc := `package example

// Add adds two integers.
func Add(a, b int) int {
	return a + b
}
`
	if err := os.WriteFile(filepath.Join(tmpDir, "good.go"), []byte(goodSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	// Poorly-structured function: no doc, many params, deep nesting, ignored errors
	badSrc := `package example

import "os"

func Process(a int, b string, c bool, d float64, e []byte, f int64) {
	if a > 0 {
		for i := 0; i < a; i++ {
			if c {
				if d > 0 {
					_, _ = os.Open(b)
				}
			}
		}
	}
}
`
	if err := os.WriteFile(filepath.Join(tmpDir, "bad.go"), []byte(badSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	certifier := &engine.Certifier{
		Root: tmpDir,
		ExpiryCfg: domain.ExpiryConfig{
			DefaultWindowDays: 90,
			MinWindowDays:     7,
			MaxWindowDays:     365,
		},
	}

	ctx := context.Background()
	now := time.Now()

	// Certify good function
	goodUnit := domain.NewUnit(
		domain.NewUnitID("go", "good.go", "Add"),
		domain.UnitTypeFunction,
	)
	goodResult, err := certifier.Certify(ctx, goodUnit, nil, now)
	if err != nil {
		t.Fatalf("Certify(Add) error: %v", err)
	}

	// Certify bad function
	badUnit := domain.NewUnit(
		domain.NewUnitID("go", "bad.go", "Process"),
		domain.UnitTypeFunction,
	)
	badResult, err := certifier.Certify(ctx, badUnit, nil, now)
	if err != nil {
		t.Fatalf("Certify(Process) error: %v", err)
	}

	// Good function should have structural evidence with doc comment
	hasStructural := false
	for _, e := range goodResult.Record.Evidence {
		if e.Kind == domain.EvidenceKindStructural {
			hasStructural = true
			if e.Metrics["has_doc_comment"] != 1.0 {
				t.Errorf("Add: has_doc_comment = %f, want 1.0", e.Metrics["has_doc_comment"])
			}
			if e.Metrics["param_count"] != 2 {
				t.Errorf("Add: param_count = %f, want 2", e.Metrics["param_count"])
			}
		}
	}
	if !hasStructural {
		t.Error("Add should have structural evidence")
	}

	// Bad function should have structural evidence showing problems
	for _, e := range badResult.Record.Evidence {
		if e.Kind == domain.EvidenceKindStructural {
			if e.Metrics["has_doc_comment"] != 0.0 {
				t.Errorf("Process: has_doc_comment = %f, want 0.0", e.Metrics["has_doc_comment"])
			}
			if e.Metrics["param_count"] != 6 {
				t.Errorf("Process: param_count = %f, want 6", e.Metrics["param_count"])
			}
			if e.Metrics["max_nesting_depth"] < 3 {
				t.Errorf("Process: max_nesting_depth = %f, want >= 3", e.Metrics["max_nesting_depth"])
			}
			if e.Metrics["errors_ignored"] < 1 {
				t.Errorf("Process: errors_ignored = %f, want >= 1", e.Metrics["errors_ignored"])
			}
		}
	}

	// Good function should score higher than bad function
	if goodResult.Record.Score <= badResult.Record.Score {
		t.Errorf("Good score (%f) should be > bad score (%f)",
			goodResult.Record.Score, badResult.Record.Score)
	}

	// Good function readability should be better
	goodRead := goodResult.Record.Dimensions[domain.DimReadability]
	badRead := badResult.Record.Dimensions[domain.DimReadability]
	if goodRead <= badRead {
		t.Errorf("Good readability (%f) should be > bad readability (%f)", goodRead, badRead)
	}
}
