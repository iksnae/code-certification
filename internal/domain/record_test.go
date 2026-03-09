package domain_test

import (
	"testing"

	"github.com/code-certification/certify/internal/domain"
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
