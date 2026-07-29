package main

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/report"
)

// legacyStore opens the captured 6c9110de3 output: 5 Go units at B+ (0.8889)
// and 4 Swift units the engine cannot analyse. See scripts/gen-legacy-fixture.sh.
func legacyStore(t *testing.T) *record.Store {
	t.Helper()
	return record.NewStore(filepath.Join("..", "..", "testdata", "legacy-store", "records"))
}

// TestBuildCertificationRun_GradeMatchesReportCard is the assertion the grade
// arm never had. Every existing test of buildCertificationRun passes nil for
// the store (tally_test.go:83), which skips the scoring block entirely — the
// arm that computes the grade persisted to runs.jsonl and state.json was 0%
// covered, which is why it could contradict REPORT_CARD.md for six rounds.
//
// One command writes both artifacts. They are the same claim, so they get the
// same number, and this test is what holds them together.
func TestBuildCertificationRun_GradeMatchesReportCard(t *testing.T) {
	store := legacyStore(t)
	recs, err := store.ListAll()
	if err != nil {
		t.Fatalf("ListAll() error: %v", err)
	}

	run := buildCertificationRun(runParams{
		runID:  "run-test",
		commit: "abc123",
		tally:  runTally{certified: 5, unsupported: 4, processed: 9},
	}, store)

	card := report.GenerateCard(recs, "", "abc123", time.Now())

	if run.OverallGrade != card.OverallGrade {
		t.Errorf("run.OverallGrade = %q, card.OverallGrade = %q — runs.jsonl and REPORT_CARD.md must not contradict each other",
			run.OverallGrade, card.OverallGrade)
	}
	if run.OverallScore != card.OverallScore {
		t.Errorf("run.OverallScore = %v, card.OverallScore = %v", run.OverallScore, card.OverallScore)
	}
}

// TestBuildCertificationRun_ScoreExcludesUnassessedUnits pins the actual
// arithmetic, so the test above cannot be satisfied by making both surfaces
// wrong in the same way.
//
// 5 Go units at 0.8889 and 4 unassessed Swift units: the grade is B+ over the
// five that were measured, not 5 × 0.8889 / 9 = 0.4938 over all nine. 0.4938
// is an F, and an F is what landed in the git-tracked state.json.
func TestBuildCertificationRun_ScoreExcludesUnassessedUnits(t *testing.T) {
	run := buildCertificationRun(runParams{
		runID:  "run-test",
		commit: "abc123",
		tally:  runTally{certified: 5, unsupported: 4, processed: 9},
	}, legacyStore(t))

	const want = 0.8888888888888888
	if diff := run.OverallScore - want; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("run.OverallScore = %v, want %v — the four unassessed units must leave the denominator",
			run.OverallScore, want)
	}
	if run.OverallGrade != "B+" {
		t.Errorf("run.OverallGrade = %q, want \"B+\" — 0.4938 (the all-nine mean) grades F", run.OverallGrade)
	}
	if run.UnitsUnsupported != 4 {
		t.Errorf("run.UnitsUnsupported = %d, want 4", run.UnitsUnsupported)
	}
}
