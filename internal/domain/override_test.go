package domain_test

import (
	"testing"
	"time"

	"github.com/code-certification/certify/internal/domain"
)

func TestOverrideAction_String(t *testing.T) {
	tests := []struct {
		a    domain.OverrideAction
		want string
	}{
		{domain.OverrideExempt, "exempt"},
		{domain.OverrideExtendWindow, "extend_window"},
		{domain.OverrideShortenWindow, "shorten_window"},
		{domain.OverrideForceReview, "force_review"},
	}
	for _, tt := range tests {
		if got := tt.a.String(); got != tt.want {
			t.Errorf("OverrideAction(%d).String() = %q, want %q", tt.a, got, tt.want)
		}
	}
}

func TestOverride_RequiresRationale(t *testing.T) {
	o := domain.Override{
		UnitID:    domain.NewUnitID("go", "main.go", "main"),
		Action:    domain.OverrideExempt,
		Rationale: "",
		Actor:     "admin",
		Timestamp: time.Now(),
	}

	if err := o.Validate(); err == nil {
		t.Error("Override with empty rationale should fail validation")
	}

	o.Rationale = "Legacy code, will be removed in Q2"
	if err := o.Validate(); err != nil {
		t.Errorf("Override with rationale should pass validation: %v", err)
	}
}

func TestOverride_RequiresActor(t *testing.T) {
	o := domain.Override{
		UnitID:    domain.NewUnitID("go", "main.go", "main"),
		Action:    domain.OverrideExempt,
		Rationale: "Legacy code",
		Actor:     "",
		Timestamp: time.Now(),
	}

	if err := o.Validate(); err == nil {
		t.Error("Override with empty actor should fail validation")
	}
}
