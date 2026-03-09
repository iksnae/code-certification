package domain

import (
	"fmt"
	"time"
)

// Status represents the certification status of a code unit.
type Status int

const (
	StatusCertified                 Status = iota // Fully compliant
	StatusCertifiedWithObservations               // Acceptable with minor issues
	StatusProbationary                            // Needs improvement within short window
	StatusExpired                                 // Certification window elapsed
	StatusDecertified                             // Fails policy requirements
	StatusExempt                                  // Excluded by explicit override
)

var statusStrings = map[Status]string{
	StatusCertified:                 "certified",
	StatusCertifiedWithObservations: "certified_with_observations",
	StatusProbationary:              "probationary",
	StatusExpired:                   "expired",
	StatusDecertified:               "decertified",
	StatusExempt:                    "exempt",
}

var stringToStatus map[string]Status

func init() {
	stringToStatus = make(map[string]Status, len(statusStrings))
	for k, v := range statusStrings {
		stringToStatus[v] = k
	}
}

// String returns the string representation of a Status.
func (s Status) String() string {
	if str, ok := statusStrings[s]; ok {
		return str
	}
	return fmt.Sprintf("Status(%d)", s)
}

// ParseStatus converts a string to a Status.
func ParseStatus(s string) (Status, error) {
	if st, ok := stringToStatus[s]; ok {
		return st, nil
	}
	return 0, fmt.Errorf("unknown status: %q", s)
}

// IsPassing returns true if the status represents a passing certification.
func (s Status) IsPassing() bool {
	switch s {
	case StatusCertified, StatusCertifiedWithObservations, StatusExempt:
		return true
	default:
		return false
	}
}

// CertificationRecord is the complete trust record for a code unit.
type CertificationRecord struct {
	// Identity
	UnitID   UnitID   `json:"unit_id"`
	UnitType UnitType `json:"unit_type"`
	UnitPath string   `json:"unit_path"`

	// Policy
	PolicyVersion string `json:"policy_version"`

	// Result
	Status     Status          `json:"status"`
	Grade      Grade           `json:"grade"`
	Score      float64         `json:"score"`
	Confidence float64         `json:"confidence"`
	Dimensions DimensionScores `json:"dimensions,omitempty"`

	// Evidence
	Evidence     []Evidence `json:"evidence,omitempty"`
	Observations []string   `json:"observations,omitempty"`
	Actions      []string   `json:"actions,omitempty"`

	// Timestamps
	CertifiedAt time.Time `json:"certified_at"`
	ExpiresAt   time.Time `json:"expires_at"`

	// Metadata
	Source  string `json:"source"` // "deterministic", "agent", "combined"
	RunID   string `json:"run_id,omitempty"`
	Version int    `json:"version"` // record schema version
}
