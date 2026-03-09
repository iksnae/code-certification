package domain

import "fmt"

// CertificationMode determines whether certification results block merges.
type CertificationMode int

const (
	ModeAdvisory  CertificationMode = iota // Report findings, don't block
	ModeEnforcing                          // Block merges on configured failures
)

var modeStrings = map[CertificationMode]string{
	ModeAdvisory:  "advisory",
	ModeEnforcing: "enforcing",
}

// String returns the string representation of a CertificationMode.
func (m CertificationMode) String() string {
	if s, ok := modeStrings[m]; ok {
		return s
	}
	return fmt.Sprintf("CertificationMode(%d)", m)
}

// Config is the top-level certification configuration.
type Config struct {
	Mode      CertificationMode `yaml:"mode"`
	Scope     ScopeConfig       `yaml:"scope"`
	Policies  PolicyConfig      `yaml:"policies"`
	Analyzers AnalyzerConfig    `yaml:"analyzers"`
	Agent     AgentConfig       `yaml:"agent"`
	Schedule  ScheduleConfig    `yaml:"schedule"`
	Expiry    ExpiryConfig      `yaml:"expiry"`
	Issues    IssueConfig       `yaml:"issues"`
}

// AnalyzerConfig defines settings for tool adapters.
type AnalyzerConfig struct {
	GoVet         bool   `yaml:"go_vet"`         // Enable go vet (default: auto-detect)
	GoTest        bool   `yaml:"go_test"`        // Enable go test (default: auto-detect)
	GolangciLint  bool   `yaml:"golangci_lint"`  // Enable golangci-lint (default: auto-detect)
	ESLint        bool   `yaml:"eslint"`         // Enable eslint (default: auto-detect)
	CustomCommand string `yaml:"custom_command"` // Custom analyzer command
}

// DefaultConfig returns a Config with sensible defaults.
func DefaultConfig() Config {
	return Config{
		Mode: ModeAdvisory,
		Agent: AgentConfig{
			Enabled: false,
		},
		Expiry: ExpiryConfig{
			DefaultWindowDays: 90,
			MinWindowDays:     7,
			MaxWindowDays:     365,
		},
	}
}

// ScopeConfig defines which code paths are in/out of certification scope.
type ScopeConfig struct {
	Include []string `yaml:"include,omitempty"` // Glob patterns to include
	Exclude []string `yaml:"exclude,omitempty"` // Glob patterns to exclude
}

// PolicyConfig controls which policy packs are active.
type PolicyConfig struct {
	Enabled  []string `yaml:"enabled,omitempty"`  // Pack names to enable (empty = all)
	Disabled []string `yaml:"disabled,omitempty"` // Pack names to disable
}

// AgentConfig configures the optional agent-assisted review.
type AgentConfig struct {
	Enabled   bool             `yaml:"enabled"`
	Provider  ProviderConfig   `yaml:"provider"`
	Models    ModelAssignments `yaml:"models"`
	RateLimit RateLimitConfig  `yaml:"rate_limit"`
}

// ProviderConfig defines the LLM provider settings.
type ProviderConfig struct {
	Type        string `yaml:"type"` // openrouter, openai, local
	BaseURL     string `yaml:"base_url"`
	APIKeyEnv   string `yaml:"api_key_env"` // Env var name (not the key itself)
	HTTPReferer string `yaml:"http_referer"`
	XTitle      string `yaml:"x_title"`
}

// ModelAssignments maps certification tasks to specific models.
type ModelAssignments struct {
	Prescreen   string `yaml:"prescreen"`
	Review      string `yaml:"review"`
	Scoring     string `yaml:"scoring"`
	Decision    string `yaml:"decision"`
	Remediation string `yaml:"remediation"`
	Fallback    string `yaml:"fallback"`
}

// RateLimitConfig defines rate limiting for API calls.
type RateLimitConfig struct {
	RequestsPerMinute  int `yaml:"requests_per_minute"`
	RetryMax           int `yaml:"retry_max"`
	RetryBackoffBaseMs int `yaml:"retry_backoff_base_ms"`
}

// ScheduleConfig defines which scheduled workflows are enabled.
type ScheduleConfig struct {
	Nightly bool `yaml:"nightly"`
	Weekly  bool `yaml:"weekly"`
	Sweep   bool `yaml:"sweep"`
}

// ExpiryConfig defines bounds for certification expiry windows.
type ExpiryConfig struct {
	DefaultWindowDays int `yaml:"default_window_days"`
	MinWindowDays     int `yaml:"min_window_days"`
	MaxWindowDays     int `yaml:"max_window_days"`
}

// IssueConfig defines GitHub issue sync settings.
type IssueConfig struct {
	Enabled  bool     `yaml:"enabled"`
	Labels   []string `yaml:"labels,omitempty"`
	Grouping string   `yaml:"grouping,omitempty"` // "single", "directory", "policy"
}

// EnforcingConfig controls what blocks merges in enforcing mode.
type EnforcingConfig struct {
	BlockOnSeverity []string `yaml:"block_on_severity,omitempty"` // error, critical
	BlockOnStatus   []string `yaml:"block_on_status,omitempty"`   // decertified, probationary
	BlockPaths      []string `yaml:"block_paths,omitempty"`       // paths that must pass
	AllowPaths      []string `yaml:"allow_paths,omitempty"`       // paths exempt from blocking
}

// SignoffConfig controls human signoff requirements.
type SignoffConfig struct {
	RequiredPaths []string `yaml:"required_paths,omitempty"` // paths requiring signoff
	RequiredLabel string   `yaml:"required_label,omitempty"` // GitHub label for signoff
}
