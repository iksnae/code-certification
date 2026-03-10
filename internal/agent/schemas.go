package agent

// PrescreenResponse is the structured output from the prescreen step.
type PrescreenResponse struct {
	NeedsReview bool     `json:"needs_review"`
	Reason      string   `json:"reason"`
	Confidence  float64  `json:"confidence"`
	Suggestions []string `json:"suggestions,omitempty"`
}

// ScoringResponse is the structured output from the scoring step.
type ScoringResponse struct {
	Scores     map[string]float64 `json:"scores"`
	Confidence float64            `json:"confidence"`
	Reasoning  string             `json:"reasoning"`
}

// DecisionResponse is the structured output from the decision step.
type DecisionResponse struct {
	Status    string   `json:"status"`
	Reasoning string   `json:"reasoning"`
	Actions   []string `json:"actions"`
}

// RemediationResponse is the structured output from the remediation step.
type RemediationResponse struct {
	Steps []RemediationStep `json:"steps"`
}

// RemediationStep is a single remediation action.
type RemediationStep struct {
	Priority    int    `json:"priority"`
	Dimension   string `json:"dimension"`
	Description string `json:"description"`
	Effort      string `json:"effort"` // low, medium, high
}
