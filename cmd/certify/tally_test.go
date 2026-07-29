package main

import (
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
)

func unsupportedRec() domain.CertificationRecord {
	return domain.CertificationRecord{
		UnitID:      domain.NewUnitID("swift", "App.swift", "App"),
		Status:      domain.StatusExempt,
		Grade:       domain.GradeNA,
		Unsupported: true,
	}
}

func statusRec(s domain.Status) domain.CertificationRecord {
	return domain.CertificationRecord{
		UnitID: domain.NewUnitID("go", "main.go", "main"),
		Status: s,
		Grade:  domain.GradeB,
		Score:  0.85,
	}
}

// TestRunTally_UnsupportedIsNotCertified guards the durable half of the
// false-positive. An unsupported unit carries StatusExempt, which the outcome
// switch folds into `observations` and the run record then sums into
// `certified + observations`. That writes "we certified this unit" into
// .certification/runs.jsonl for code the engine never opened — a falsehood that
// outlives the console output that announced it.
func TestRunTally_UnsupportedIsNotCertified(t *testing.T) {
	var tally runTally
	tally.add(unsupportedRec())
	tally.add(unsupportedRec())
	tally.add(unsupportedRec())

	if tally.unsupported != 3 {
		t.Errorf("unsupported = %d, want 3", tally.unsupported)
	}
	if tally.certified != 0 {
		t.Errorf("certified = %d, want 0 — an unassessed unit was not certified", tally.certified)
	}
	if tally.observations != 0 {
		t.Errorf("observations = %d, want 0 — an unassessed unit has no observations", tally.observations)
	}
	if tally.failed != 0 {
		t.Errorf("failed = %d, want 0 — an unassessed unit did not fail", tally.failed)
	}
}

// TestRunTally_AssessedUnitsUnchanged guards against the fix over-reaching: the
// existing classification of genuinely assessed units must not move.
func TestRunTally_AssessedUnitsUnchanged(t *testing.T) {
	var tally runTally
	tally.add(statusRec(domain.StatusCertified))
	tally.add(statusRec(domain.StatusCertifiedWithObservations))
	tally.add(statusRec(domain.StatusExempt)) // operator-exempted, NOT unsupported
	tally.add(statusRec(domain.StatusDecertified))
	tally.add(statusRec(domain.StatusProbationary))

	if tally.certified != 1 {
		t.Errorf("certified = %d, want 1", tally.certified)
	}
	if tally.observations != 2 {
		t.Errorf("observations = %d, want 2 (with-observations + operator exempt)", tally.observations)
	}
	if tally.failed != 2 {
		t.Errorf("failed = %d, want 2", tally.failed)
	}
	if tally.unsupported != 0 {
		t.Errorf("unsupported = %d, want 0 — no record carried the Unsupported flag", tally.unsupported)
	}
}

// TestBuildCertificationRun_UnsupportedPersistedSeparately checks what actually
// lands on disk in .certification/runs.jsonl.
func TestBuildCertificationRun_UnsupportedPersistedSeparately(t *testing.T) {
	run := buildCertificationRun(runParams{
		runID: "run-test",
		tally: runTally{unsupported: 3, processed: 3},
	}, nil)

	if run.UnitsProcessed != 3 {
		t.Errorf("UnitsProcessed = %d, want 3 — runTally.processed is the only carrier", run.UnitsProcessed)
	}
	if run.UnitsCertified != 0 {
		t.Errorf("UnitsCertified = %d, want 0 — unassessed units must not be persisted as certified", run.UnitsCertified)
	}
	if run.UnitsFailed != 0 {
		t.Errorf("UnitsFailed = %d, want 0 — unassessed units must not be persisted as failed", run.UnitsFailed)
	}
	if run.UnitsUnsupported != 3 {
		t.Errorf("UnitsUnsupported = %d, want 3 — the run record must account for them", run.UnitsUnsupported)
	}
}

// TestEnforcingGateWarning covers the exit-code surface. Excluding unsupported
// units from `failed` means an all-unsupported enforcing repo now exits 0. The
// exit code stays two-state, so the run must say out loud that it certified
// nothing rather than letting a green CI imply that it did.
func TestEnforcingGateWarning(t *testing.T) {
	cases := []struct {
		name      string
		mode      domain.CertificationMode
		tally     runTally
		wantWarns bool
	}{
		{"all unsupported, enforcing", domain.ModeEnforcing, runTally{unsupported: 3, processed: 3}, true},
		{"all unsupported, advisory", domain.ModeAdvisory, runTally{unsupported: 3, processed: 3}, false},
		{"some assessed, enforcing", domain.ModeEnforcing, runTally{certified: 1, unsupported: 3, processed: 4}, false},
		{"nothing processed, enforcing", domain.ModeEnforcing, runTally{}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := enforcingGateWarning(tc.mode, tc.tally)
			if (got != "") != tc.wantWarns {
				t.Errorf("enforcingGateWarning() = %q, wantWarns = %v", got, tc.wantWarns)
			}
		})
	}
}
