package domain

import (
	"fmt"
	"time"
)

// OverrideAction specifies what a manual override does.
type OverrideAction int

const (
	OverrideExempt        OverrideAction = iota // Exclude unit from certification
	OverrideExtendWindow                        // Grant a longer trust window
	OverrideShortenWindow                       // Require more frequent re-evaluation
	OverrideForceReview                         // Force immediate recertification
)

var overrideActionStrings = map[OverrideAction]string{
	OverrideExempt:        "exempt",
	OverrideExtendWindow:  "extend_window",
	OverrideShortenWindow: "shorten_window",
	OverrideForceReview:   "force_review",
}

// String returns the string representation of an OverrideAction.
func (a OverrideAction) String() string {
	if s, ok := overrideActionStrings[a]; ok {
		return s
	}
	return fmt.Sprintf("OverrideAction(%d)", a)
}

// Override represents a manual governance action on a unit.
type Override struct {
	UnitID    UnitID         `json:"unit_id" yaml:"unit_id"`
	Action    OverrideAction `json:"action" yaml:"action"`
	Rationale string         `json:"rationale" yaml:"rationale"`
	Actor     string         `json:"actor" yaml:"actor"`
	Timestamp time.Time      `json:"timestamp" yaml:"timestamp"`
}

// Validate checks that required fields are present.
func (o Override) Validate() error {
	if o.Rationale == "" {
		return fmt.Errorf("override requires rationale")
	}
	if o.Actor == "" {
		return fmt.Errorf("override requires actor")
	}
	return nil
}
