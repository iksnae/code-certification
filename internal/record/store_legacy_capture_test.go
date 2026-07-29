package record_test

import (
	"path/filepath"
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/record"
)

// legacyFixtureDir is the on-disk output of the certify binary at 6c9110de3 —
// the commit before the `unsupported` field existed. Captured, not authored;
// see scripts/gen-legacy-fixture.sh.
//
// The provenance is the point. The previous legacy fixture in this repo was
// hand-written as {"status":"exempt","grade":"N/A"}, which is today's output
// with one key deleted. The binary that actually wrote the records sitting in
// every client's .certification/records/ wrote
// {"status":"decertified","grade":"F","score":0,"confidence":1}. Only the real
// shape reaches the defect, because the defect is precisely a stored grade
// that contradicts the backfilled flag — and a fixture whose stored grade is
// already "N/A" has no contradiction to expose.
func legacyFixtureDir(t *testing.T) string {
	t.Helper()
	return filepath.Join("..", "..", "testdata", "legacy-store", "records")
}

func loadLegacy(t *testing.T) []domain.CertificationRecord {
	t.Helper()
	recs, err := record.NewStore(legacyFixtureDir(t)).ListAll()
	if err != nil {
		t.Fatalf("ListAll() error: %v", err)
	}
	if len(recs) != 9 {
		t.Fatalf("ListAll() returned %d records, want 9 (5 Go + 4 Swift)", len(recs))
	}
	return recs
}

// TestStore_LegacyUnsupportedCarriesUnassessedVerdict is the root assertion.
//
// A record read from a legacy store used to arrive half-migrated: Unsupported
// backfilled to true from the language, while Status, Grade, Score and
// Confidence kept the confident-F verdict the old binary wrote. Every
// aggregation downstream then had to remember to branch on the flag rather
// than trust the fields — and six review rounds each found one more site that
// had forgotten. A record that is unassessed in one field and decertified/F in
// the next is the defect; the fields must agree at the boundary.
func TestStore_LegacyUnsupportedCarriesUnassessedVerdict(t *testing.T) {
	for _, r := range loadLegacy(t) {
		if !r.Unsupported {
			continue
		}
		if r.Status != domain.StatusExempt {
			t.Errorf("%s: Status = %v, want %v — an unassessed unit asserts no verdict",
				r.UnitID, r.Status, domain.StatusExempt)
		}
		if r.Grade != domain.GradeNA {
			t.Errorf("%s: Grade = %v, want %v — the stored F was a claim about code the engine never opened",
				r.UnitID, r.Grade, domain.GradeNA)
		}
		if r.Score != 0 {
			t.Errorf("%s: Score = %v, want 0", r.UnitID, r.Score)
		}
		if r.Confidence != 0 {
			t.Errorf("%s: Confidence = %v, want 0 — the old binary wrote 1.0 for a unit it could not analyse",
				r.UnitID, r.Confidence)
		}
		if len(r.Dimensions) != 0 {
			t.Errorf("%s: Dimensions = %v, want none", r.UnitID, r.Dimensions)
		}
	}
}

// TestStore_LegacyAssessedRecordsUntouched guards against the normalization
// over-reaching onto units that were genuinely assessed.
func TestStore_LegacyAssessedRecordsUntouched(t *testing.T) {
	var assessed int
	for _, r := range loadLegacy(t) {
		if r.Unsupported {
			continue
		}
		assessed++
		if r.Status != domain.StatusCertified {
			t.Errorf("%s: Status = %v, want %v", r.UnitID, r.Status, domain.StatusCertified)
		}
		if r.Grade != domain.GradeBPlus {
			t.Errorf("%s: Grade = %v, want B+", r.UnitID, r.Grade)
		}
		if r.Score == 0 {
			t.Errorf("%s: Score = 0, want the stored measurement", r.UnitID)
		}
	}
	if assessed != 5 {
		t.Errorf("assessed records = %d, want 5", assessed)
	}
}

// TestStore_BackfillAgreesWithFreshPipeline pins the invariant that makes the
// normalization correct rather than merely convenient: a legacy record read
// through the store must be indistinguishable, on every verdict field, from
// what a fresh run writes for the same unit. Two producers of an unassessed
// record, one answer. If the pipeline's unassessed verdict ever changes, this
// fails rather than letting the two drift.
func TestStore_BackfillAgreesWithFreshPipeline(t *testing.T) {
	want := domain.UnassessedVerdict()
	for _, r := range loadLegacy(t) {
		if !r.Unsupported {
			continue
		}
		got := domain.Verdict{
			Status:     r.Status,
			Grade:      r.Grade,
			Score:      r.Score,
			Confidence: r.Confidence,
		}
		if got != want {
			t.Errorf("%s: verdict = %+v, want %+v — the store backfill and the pipeline must produce one answer",
				r.UnitID, got, want)
		}
	}
}
