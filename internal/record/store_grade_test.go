package record_test

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/record"
)

// saveThenRewriteGrade persists a record, then rewrites the persisted "grade"
// string to the given value, simulating a truncated, hand-edited, or
// forward-version record on disk. It returns the record as the store reloads it.
func saveThenRewriteGrade(t *testing.T, gradeJSON string) domain.CertificationRecord {
	t.Helper()

	dir := t.TempDir()
	store := record.NewStore(dir)

	rec := domain.CertificationRecord{
		UnitID:      domain.NewUnitID("go", "internal/app/main.go", "main"),
		UnitType:    domain.UnitTypeFunction,
		UnitPath:    "internal/app/main.go",
		Status:      domain.StatusCertified,
		Grade:       domain.GradeA,
		Score:       0.95,
		Confidence:  1.0,
		CertifiedAt: time.Now(),
		ExpiresAt:   time.Now().Add(24 * time.Hour),
		Source:      "deterministic",
		Version:     1,
	}
	if err := store.Save(rec); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	matches, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil || len(matches) != 1 {
		t.Fatalf("expected exactly 1 record file, got %v (err %v)", matches, err)
	}
	raw, err := os.ReadFile(matches[0])
	if err != nil {
		t.Fatalf("ReadFile() error: %v", err)
	}
	gradeField := regexp.MustCompile(`"grade":\s*"[^"]*"`)
	if !gradeField.Match(raw) {
		t.Fatalf("could not find the grade field in:\n%s", raw)
	}
	mutated := gradeField.ReplaceAllLiteralString(string(raw), `"grade": `+gradeJSON)
	if err := os.WriteFile(matches[0], []byte(mutated), 0o644); err != nil {
		t.Fatalf("WriteFile() error: %v", err)
	}

	loaded, err := store.Load(rec.UnitID)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	return loaded
}

// TestStore_AbsentGradeIsNotF is the same defect class as the unsupported-language
// bug, one layer down. parseGrade's fallback returns GradeF for every string it
// does not recognise, so a truncated or partially-written record — one whose
// grade field is empty — is silently reloaded as a confident F. A failing verdict
// must never be manufactured from absent data.
func TestStore_AbsentGradeIsNotF(t *testing.T) {
	loaded := saveThenRewriteGrade(t, `""`)

	if loaded.Grade == domain.GradeF {
		t.Errorf("loaded Grade = F for an absent grade string — a failing verdict was fabricated from missing data")
	}
	if loaded.Grade != domain.GradeNA {
		t.Errorf("loaded Grade = %v, want N/A for an absent grade string", loaded.Grade)
	}
}

// TestStore_UnrecognisedGradeIsNotF covers the general case: any grade string the
// current binary does not know — a typo, a corrupted byte, a grade written by a
// newer version — must decay to "unknown", not to "failed".
func TestStore_UnrecognisedGradeIsNotF(t *testing.T) {
	for _, s := range []string{`"A+"`, `"n/a"`, `"???"`, `"Grade(9)"`} {
		t.Run(s, func(t *testing.T) {
			loaded := saveThenRewriteGrade(t, s)
			if loaded.Grade != domain.GradeNA {
				t.Errorf("grade %s loaded as %v, want N/A — an unrecognised grade is not a failing grade", s, loaded.Grade)
			}
		})
	}
}

// TestStore_KnownGradesStillParse guards against the fix over-reaching: every
// grade the domain can render must still round-trip to itself.
func TestStore_KnownGradesStillParse(t *testing.T) {
	for _, g := range []domain.Grade{
		domain.GradeA, domain.GradeAMinus, domain.GradeBPlus,
		domain.GradeB, domain.GradeC, domain.GradeD, domain.GradeF, domain.GradeNA,
	} {
		t.Run(g.String(), func(t *testing.T) {
			loaded := saveThenRewriteGrade(t, `"`+g.String()+`"`)
			if loaded.Grade != g {
				t.Errorf("grade %q loaded as %v, want %v", g.String(), loaded.Grade, g)
			}
		})
	}
}
