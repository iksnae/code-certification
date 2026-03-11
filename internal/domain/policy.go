package domain

// PolicyPack is a versioned set of certification rules.
type PolicyPack struct {
	Name         string       `json:"name" yaml:"name"`
	Version      string       `json:"version" yaml:"version"`
	Language     string       `json:"language,omitempty" yaml:"language,omitempty"`           // Empty = global (all languages)
	PathPatterns []string     `json:"path_patterns,omitempty" yaml:"path_patterns,omitempty"` // Glob patterns
	Rules        []PolicyRule `json:"rules" yaml:"rules"`
}

// IsGlobal returns true if this policy applies to all languages.
func (p PolicyPack) IsGlobal() bool {
	return p.Language == ""
}

// PolicyRule defines a single certification requirement.
type PolicyRule struct {
	ID              string    `json:"id" yaml:"id"`
	Dimension       Dimension `json:"dimension" yaml:"dimension"`
	Description     string    `json:"description" yaml:"description"`
	Severity        Severity  `json:"severity" yaml:"severity"`
	Threshold       float64   `json:"threshold,omitempty" yaml:"threshold,omitempty"`               // Metric must be below this
	Metric          string    `json:"metric,omitempty" yaml:"metric,omitempty"`                     // Which metric to evaluate
	PathPatterns    []string  `json:"path_patterns,omitempty" yaml:"path_patterns,omitempty"`       // If set, rule only applies to matching paths
	ExcludePatterns []string  `json:"exclude_patterns,omitempty" yaml:"exclude_patterns,omitempty"` // If set, rule skips matching paths
}

// Violation records a specific policy rule failure for a unit.
type Violation struct {
	RuleID      string    `json:"rule_id"`
	PolicyName  string    `json:"policy_name"`
	Severity    Severity  `json:"severity"`
	Description string    `json:"description"`
	Dimension   Dimension `json:"dimension"`
}
