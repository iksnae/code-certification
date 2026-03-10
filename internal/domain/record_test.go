package domain_test

import (
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

func TestStatus_String(t *testing.T) {
	tests := []struct {
		s    domain.Status
		want string
	}{
		{domain.StatusCertified, "certified"},
		{domain.StatusCertifiedWithObservations, "certified_with_observations"},
		{domain.StatusProbationary, "probationary"},
		{domain.StatusExpired, "expired"},
		{domain.StatusDecertified, "decertified"},
		{domain.StatusExempt, "exempt"},
	}
	for _, tt := range tests {
		if got := tt.s.String(); got != tt.want {
			t.Errorf("Status(%d).String() = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestParseStatus(t *testing.T) {
	tests := []struct {
		input string
		want  domain.Status
		ok    bool
	}{
		{"certified", domain.StatusCertified, true},
		{"certified_with_observations", domain.StatusCertifiedWithObservations, true},
		{"probationary", domain.StatusProbationary, true},
		{"expired", domain.StatusExpired, true},
		{"decertified", domain.StatusDecertified, true},
		{"exempt", domain.StatusExempt, true},
		{"unknown", 0, false},
		{"", 0, false},
	}
	for _, tt := range tests {
		got, err := domain.ParseStatus(tt.input)
		if tt.ok {
			if err != nil {
				t.Errorf("ParseStatus(%q) unexpected error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("ParseStatus(%q) = %v, want %v", tt.input, got, tt.want)
			}
		} else {
			if err == nil {
				t.Errorf("ParseStatus(%q) expected error, got nil", tt.input)
			}
		}
	}
}

func TestStatus_IsPassing(t *testing.T) {
	passing := []domain.Status{
		domain.StatusCertified,
		domain.StatusCertifiedWithObservations,
		domain.StatusExempt,
	}
	for _, s := range passing {
		if !s.IsPassing() {
			t.Errorf("%v.IsPassing() = false, want true", s)
		}
	}

	failing := []domain.Status{
		domain.StatusProbationary,
		domain.StatusExpired,
		domain.StatusDecertified,
	}
	for _, s := range failing {
		if s.IsPassing() {
			t.Errorf("%v.IsPassing() = true, want false", s)
		}
	}
}

func TestGenerateRunID(t *testing.T) {
	ts := time.Date(2026, 3, 10, 15, 52, 27, 0, time.UTC)
	got := domain.GenerateRunID(ts)
	want := "run-20260310T155227Z"
	if got != want {
		t.Errorf("GenerateRunID() = %q, want %q", got, want)
	}

	// Non-UTC time should be converted to UTC
	est := time.FixedZone("EST", -5*3600)
	ts2 := time.Date(2026, 3, 10, 10, 52, 27, 0, est) // same instant as UTC 15:52:27
	got2 := domain.GenerateRunID(ts2)
	if got2 != want {
		t.Errorf("GenerateRunID(EST) = %q, want %q", got2, want)
	}

	// Must start with "run-"
	if !strings.HasPrefix(got, "run-") {
		t.Errorf("RunID should start with 'run-', got %q", got)
	}
}

func TestCertificationRun_Fields(t *testing.T) {
	now := time.Now()
	run := domain.CertificationRun{
		ID:             "run-20260310T155227Z",
		StartedAt:      now,
		CompletedAt:    now.Add(5 * time.Minute),
		Commit:         "abc123",
		PolicyVersions: []string{"global@1.0", "go-strict@2.1"},
		UnitsProcessed: 100,
		UnitsCertified: 90,
		UnitsFailed:    10,
		OverallGrade:   "B+",
		OverallScore:   0.82,
	}

	if run.ID != "run-20260310T155227Z" {
		t.Errorf("ID = %q", run.ID)
	}
	if run.Commit != "abc123" {
		t.Errorf("Commit = %q", run.Commit)
	}
	if len(run.PolicyVersions) != 2 {
		t.Errorf("PolicyVersions len = %d, want 2", len(run.PolicyVersions))
	}
	if run.UnitsProcessed != 100 {
		t.Errorf("UnitsProcessed = %d, want 100", run.UnitsProcessed)
	}
	if run.UnitsCertified != 90 {
		t.Errorf("UnitsCertified = %d, want 90", run.UnitsCertified)
	}
	if run.OverallGrade != "B+" {
		t.Errorf("OverallGrade = %q, want B+", run.OverallGrade)
	}
}
