package report_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/report"
)

// legacyStoreRecords writes the mixed fixture the way a binary predating the
// `unsupported` field wrote it — no key at all — and reads it back through the
// store. This is the shape of every record already committed under
// .certification/records/, and it is the input the client surfaces actually
// receive: reports are rendered from stored records, never from the in-memory
// pipeline result.
func legacyStoreRecords(t *testing.T) []domain.CertificationRecord {
	t.Helper()
	dir := t.TempDir()
	files := map[string]string{
		"a.json": `{"unit_id":"go://app/a.go#A","unit_type":"function","unit_path":"app/a.go",` +
			`"status":"certified","grade":"A","score":0.9,"confidence":1,` +
			`"certified_at":"2026-01-01T00:00:00Z","expires_at":"2027-01-01T00:00:00Z",` +
			`"source":"deterministic","version":1}`,
		"b.json": `{"unit_id":"swift://app/b.swift#b","unit_type":"function","unit_path":"app/b.swift",` +
			`"status":"exempt","grade":"N/A","score":0,"confidence":0,` +
			`"certified_at":"2026-01-01T00:00:00Z","expires_at":"2027-01-01T00:00:00Z",` +
			`"source":"deterministic","version":1}`,
		"c.json": `{"unit_id":"swift://app/c.swift#c","unit_type":"function","unit_path":"app/c.swift",` +
			`"status":"exempt","grade":"N/A","score":0,"confidence":0,` +
			`"certified_at":"2026-01-01T00:00:00Z","expires_at":"2027-01-01T00:00:00Z",` +
			`"source":"deterministic","version":1}`,
	}
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatalf("writing %s: %v", name, err)
		}
	}
	recs, err := record.NewStore(dir).ListAll()
	if err != nil {
		t.Fatalf("ListAll() error: %v", err)
	}
	if len(recs) != 3 {
		t.Fatalf("ListAll() returned %d records, want 3", len(recs))
	}
	return recs
}

// TestClientSurfaces_LegacyStoredRecords is #47 reproduced on stored data.
//
// Every assertion here passes trivially against in-memory fixtures, because
// those set Unsupported directly. Routed through the store, the same three
// units used to arrive with Unsupported=false and every surface reported one
// unassessed Swift package as three certified units. This is the test that
// distinguishes "the aggregation is fixed" from "the product is fixed".
func TestClientSurfaces_LegacyStoredRecords(t *testing.T) {
	recs := legacyStoreRecords(t)
	now := time.Now()

	c := report.GenerateCard(recs, "test/repo", "abc", now)
	if c.UnsupportedCount != 2 {
		t.Errorf("card UnsupportedCount = %d, want 2", c.UnsupportedCount)
	}
	if c.Passing != 1 {
		t.Errorf("card Passing = %d, want 1 — two Swift units were never assessed", c.Passing)
	}
	if c.AnalyzableUnits != 1 {
		t.Errorf("card AnalyzableUnits = %d, want 1", c.AnalyzableUnits)
	}
	if c.PassRate != 1.0 {
		t.Errorf("card PassRate = %.3f, want 1.000 over the single analyzable unit", c.PassRate)
	}

	fr := report.GenerateFullReport(recs, "test/repo", "abc", now)
	pkgs := report.BuildPackageSummaries(fr)
	if len(pkgs) != 1 {
		t.Fatalf("package count = %d, want 1", len(pkgs))
	}
	if pkgs[0].Unsupported != 2 {
		t.Errorf("package Unsupported = %d, want 2", pkgs[0].Unsupported)
	}
	if pkgs[0].Analyzable() != 1 {
		t.Errorf("package Analyzable() = %d, want 1", pkgs[0].Analyzable())
	}

	d := report.Detailed(recs, now)
	if got := d.ByLanguage["swift"].Unsupported; got != 2 {
		t.Errorf("detailed swift Unsupported = %d, want 2", got)
	}
	if got := d.ByLanguage["swift"].Passing; got != 0 {
		t.Errorf("detailed swift Passing = %d, want 0", got)
	}

	// Health derives from the language rather than the field, so it was already
	// right on stored data. Pinning it here proves the backfill agrees with it
	// instead of introducing a second answer.
	h := report.Health(recs)
	if h.AnalyzableUnits != 1 || h.Unsupported != 2 {
		t.Errorf("health AnalyzableUnits=%d Unsupported=%d, want 1 and 2 — the backfill must agree with the language-derived count",
			h.AnalyzableUnits, h.Unsupported)
	}
}
