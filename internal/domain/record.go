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
	"certified":                   StatusCertified,
	"certified_with_observations": StatusCertifiedWithObservations,
	"probationary":                StatusProbationary,
	"expired":                     StatusExpired,
	"decertified":                 StatusDecertified,
	"exempt":                      StatusExempt,
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
	Status      Status          `json:"status"`
	Grade       Grade           `json:"grade"`
	Score       float64         `json:"score"`
	Confidence  float64         `json:"confidence"`
	Unsupported bool            `json:"unsupported"`
	Dimensions  DimensionScores `json:"dimensions,omitempty"`

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

// Verdict is the quality judgement a record asserts: the four fields that
// together say what the engine concluded about a unit.
type Verdict struct {
	Status     Status
	Grade      Grade
	Score      float64
	Confidence float64
}

// UnassessedVerdict is the verdict of a unit the engine cannot analyse: no
// judgement, expressed consistently in every field that could carry one.
//
// One function decides this, and both producers of an unassessed record call
// it — the pipeline, which builds one from a fresh run, and the store, which
// backfills one from a record written before the `unsupported` field existed.
// That is the whole point. A record that says "unassessed" in one field and
// "decertified / F / confidence 1.0" in the next is not a record with a
// display bug; it is a record that disagrees with itself, and every surface
// downstream then has to remember to distrust the fields and re-derive from
// the flag. Six review rounds each found one more site that had forgotten.
// Making the fields agree at the boundary is what removes the obligation.
func UnassessedVerdict() Verdict {
	return Verdict{
		Status:     StatusExempt,
		Grade:      GradeNA,
		Score:      0,
		Confidence: 0,
	}
}

// VerdictOf returns the verdict r currently asserts.
func (r CertificationRecord) VerdictOf() Verdict {
	return Verdict{Status: r.Status, Grade: r.Grade, Score: r.Score, Confidence: r.Confidence}
}

// WithUnassessedVerdict returns r carrying UnassessedVerdict and no dimension
// scores. It is idempotent: a record a fresh run already wrote as unassessed
// is returned unchanged.
func (r CertificationRecord) WithUnassessedVerdict() CertificationRecord {
	v := UnassessedVerdict()
	r.Status, r.Grade, r.Score, r.Confidence = v.Status, v.Grade, v.Score, v.Confidence
	r.Dimensions = nil
	return r
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
	// UnitsUnsupported counts units in languages the engine cannot analyse. They
	// are neither certified nor failed: no verdict was asserted about them.
	// Folding them into UnitsCertified writes a durable falsehood into
	// .certification/runs.jsonl — a claim of certification for code never opened.
	UnitsUnsupported int     `json:"units_unsupported,omitempty"`
	OverallGrade     string  `json:"overall_grade"`
	OverallScore     float64 `json:"overall_score"`
}

// GenerateRunID creates a timestamp-based run identifier.
func GenerateRunID(t time.Time) string {
	return "run-" + t.UTC().Format("20060102T150405Z")
}
