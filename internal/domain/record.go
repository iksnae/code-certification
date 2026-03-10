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

var stringToStatus = map[string]Status{
	"certified":                  StatusCertified,
	"certified_with_observations": StatusCertifiedWithObservations,
	"probationary":               StatusProbationary,
	"expired":                    StatusExpired,
	"decertified":                StatusDecertified,
	"exempt":                     StatusExempt,
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

// CertificationRun captures metadata about a single certification invocation.
type CertificationRun struct {
	ID             string    `json:"id"`
	StartedAt      time.Time `json:"started_at"`
	CompletedAt    time.Time `json:"completed_at"`
	Commit         string    `json:"commit"`
	PolicyVersions []string  `json:"policy_versions,omitempty"`
	UnitsProcessed int       `json:"units_processed"`
	UnitsCertified int       `json:"units_certified"`
	UnitsFailed    int       `json:"units_failed"`
	OverallGrade   string    `json:"overall_grade"`
	OverallScore   float64   `json:"overall_score"`
}

// GenerateRunID creates a timestamp-based run identifier.
func GenerateRunID(t time.Time) string {
	return "run-" + t.UTC().Format("20060102T150405Z")
}
