package record_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/record"
)

// writeRawRecord drops a record file into a store directory verbatim, which is
// the only way to produce the shape under test: a record written by a binary
// that predates the `unsupported` field. Saving a domain record through the
// store cannot produce it, because the current writer always has an opinion.
func writeRawRecord(t *testing.T, dir, name, body string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
		t.Fatalf("writing %s: %v", name, err)
	}
}

const legacySwiftRecord = `{"unit_id":"swift://app/b.swift#b","unit_type":"function",` +
	`"unit_path":"app/b.swift","status":"exempt","grade":"N/A","score":0,"confidence":0,` +
	`"certified_at":"2026-01-01T00:00:00Z","expires_at":"2027-01-01T00:00:00Z",` +
	`"source":"deterministic","version":1}`

const legacyGoRecord = `{"unit_id":"go://app/a.go#A","unit_type":"function",` +
	`"unit_path":"app/a.go","status":"certified","grade":"A","score":0.9,"confidence":1,` +
	`"certified_at":"2026-01-01T00:00:00Z","expires_at":"2027-01-01T00:00:00Z",` +
	`"source":"deterministic","version":1}`

// TestListAll_LegacyRecordWithoutUnsupportedKey is the reddening test for #50.
//
// The `unsupported` key did not exist before this branch, so every record
// already on disk arrives with no key at all. Decoding that into Go's zero
// value states "the engine assessed this unit and found it acceptable" — a
// definite verdict manufactured from a missing measurement, which is the shape
// of the defect this branch exists to close. The fix is inert on stored data
// until absence stops meaning `false`.
//
// The unit is Swift: the engine has no analyzer for it, so no assessment can
// ever have happened, whatever the byte on disk says.
func TestListAll_LegacyRecordWithoutUnsupportedKey(t *testing.T) {
	dir := t.TempDir()
	writeRawRecord(t, dir, "legacy_swift.json", legacySwiftRecord)

	records, err := record.NewStore(dir).ListAll()
	if err != nil {
		t.Fatalf("ListAll() error: %v", err)
	}
	if len(records) != 1 {
		t.Fatalf("ListAll() returned %d records, want 1", len(records))
	}

	if !records[0].Unsupported {
		t.Error("Unsupported = false for a Swift unit written before the field existed; " +
			"an absent key is no determination, not a determination of 'assessed'")
	}
	if records[0].Status != domain.StatusExempt {
		t.Errorf("Status = %v, want exempt — the backfill must not rewrite anything else", records[0].Status)
	}
}

// TestListAll_LegacyRecordInAnalyzableLanguage is the other half: a legacy Go
// record must stay assessed. The backfill derives from the language, so it
// reproduces exactly what the current writer would have recorded — it is not a
// blanket "absent means unsupported", which would decertify the entire corpus.
func TestListAll_LegacyRecordInAnalyzableLanguage(t *testing.T) {
	dir := t.TempDir()
	writeRawRecord(t, dir, "legacy_go.json", legacyGoRecord)

	records, err := record.NewStore(dir).ListAll()
	if err != nil {
		t.Fatalf("ListAll() error: %v", err)
	}
	if len(records) != 1 {
		t.Fatalf("ListAll() returned %d records, want 1", len(records))
	}
	if records[0].Unsupported {
		t.Error("Unsupported = true for a legacy Go unit; Go has an analyzer, so the record was assessed")
	}
}

// TestListAll_RecordedFalseIsHonoured pins the third case apart from the other
// two. A record that explicitly carries `"unsupported": false` recorded a
// determination, and a recorded measurement outranks a derived one — otherwise
// the store would be re-deciding rather than reading, and a future release that
// gains an analyzer could not express "this unit was assessed under the old
// registry" at all.
func TestListAll_RecordedFalseIsHonoured(t *testing.T) {
	dir := t.TempDir()
	explicit := `{"unit_id":"swift://app/b.swift#b","unit_type":"function",` +
		`"unit_path":"app/b.swift","status":"certified","grade":"A","score":0.9,"confidence":1,` +
		`"unsupported":false,` +
		`"certified_at":"2026-01-01T00:00:00Z","expires_at":"2027-01-01T00:00:00Z",` +
		`"source":"deterministic","version":1}`
	writeRawRecord(t, dir, "explicit.json", explicit)

	records, err := record.NewStore(dir).ListAll()
	if err != nil {
		t.Fatalf("ListAll() error: %v", err)
	}
	if len(records) != 1 {
		t.Fatalf("ListAll() returned %d records, want 1", len(records))
	}
	if records[0].Unsupported {
		t.Error("Unsupported = true, want false — the record recorded a determination and it must win over the derivation")
	}
}

// TestSave_DoesNotWriteAnUnsupportedKeyForAssessedUnits pins the on-disk bytes.
//
// Making absence meaningful must not turn into writing `"unsupported": false`
// into every record: that is a schema migration over the whole committed corpus
// (#48's territory), and it is redundant — the derivation reproduces the same
// answer from the language. The key appears only where it carries information.
func TestSave_DoesNotWriteAnUnsupportedKeyForAssessedUnits(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	assessed := domain.CertificationRecord{
		UnitID:   domain.NewUnitID("go", "app/a.go", "A"),
		UnitType: domain.UnitTypeFunction,
		UnitPath: "app/a.go",
		Status:   domain.StatusCertified,
		Grade:    domain.GradeA,
		Score:    0.9,
		Version:  1,
	}
	if err := store.Save(assessed); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("ReadDir() error: %v", err)
	}
	for _, e := range entries {
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			t.Fatalf("ReadFile(%s) error: %v", e.Name(), err)
		}
		if strings.Contains(string(data), `"unsupported"`) {
			t.Errorf("%s carries an unsupported key for an assessed unit; every committed record would be rewritten:\n%s",
				e.Name(), data)
		}
	}
}

// TestSave_WritesUnsupportedKeyWhenSet keeps the other direction pinned: the
// determination that does carry information is still persisted.
func TestSave_WritesUnsupportedKeyWhenSet(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	unassessed := domain.CertificationRecord{
		UnitID:      domain.NewUnitID("swift", "app/b.swift", "b"),
		UnitType:    domain.UnitTypeFunction,
		UnitPath:    "app/b.swift",
		Status:      domain.StatusExempt,
		Grade:       domain.GradeNA,
		Unsupported: true,
		Version:     1,
	}
	if err := store.Save(unassessed); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("ReadDir() error: %v", err)
	}
	found := false
	for _, e := range entries {
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			t.Fatalf("ReadFile(%s) error: %v", e.Name(), err)
		}
		if strings.Contains(string(data), `"unsupported": true`) {
			found = true
		}
	}
	if !found {
		t.Error(`no record file carries "unsupported": true`)
	}
}
