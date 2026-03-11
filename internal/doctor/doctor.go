// Package doctor provides health checks for certify setup and environment.
package doctor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/iksnae/code-certification/internal/agent"
	"github.com/iksnae/code-certification/internal/config"
	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/evidence"
)

// CheckStatus represents the result of a single health check.
type CheckStatus int

const (
	StatusPass CheckStatus = iota
	StatusWarn
	StatusFail
	StatusSkip
)

func (s CheckStatus) String() string {
	switch s {
	case StatusPass:
		return "pass"
	case StatusWarn:
		return "warn"
	case StatusFail:
		return "fail"
	case StatusSkip:
		return "skip"
	}
	return "unknown"
}

// Emoji returns the status emoji.
func (s CheckStatus) Emoji() string {
	switch s {
	case StatusPass:
		return "✅"
	case StatusWarn:
		return "⚠️"
	case StatusFail:
		return "❌"
	case StatusSkip:
		return "⏭️"
	}
	return "❓"
}

// Check represents a single health check result.
type Check struct {
	Name    string
	Group   string
	Status  CheckStatus
	Message string
	Fix     string // suggested fix command or action
}

// Report holds the full doctor report.
type Report struct {
	Checks []Check
	Root   string
}

// Summary returns pass/warn/fail/skip counts.
func (r *Report) Summary() (pass, warn, fail, skip int) {
	for _, c := range r.Checks {
		switch c.Status {
		case StatusPass:
			pass++
		case StatusWarn:
			warn++
		case StatusFail:
			fail++
		case StatusSkip:
			skip++
		}
	}
	return
}

// HasFailures returns true if any check failed.
func (r *Report) HasFailures() bool {
	for _, c := range r.Checks {
		if c.Status == StatusFail {
			return true
		}
	}
	return false
}

// RunAll runs all health checks for the given root directory.
func RunAll(root string) *Report {
	r := &Report{Root: root}

	r.checkEnvironment()
	r.checkModuleRoots(root)
	r.checkProjectSetup(root)
	r.checkConfiguration(root)
	r.checkTools()
	r.checkProviders()

	return r
}

// checkModuleRoots discovers language module roots (go.mod, package.json, etc.).
func (r *Report) checkModuleRoots(root string) {
	roots := evidence.DiscoverModuleRoots(root)
	if len(roots) == 0 {
		r.Checks = append(r.Checks, Check{
			Name:    "Module roots",
			Group:   "modules",
			Status:  StatusWarn,
			Message: "No language module markers found (go.mod, package.json, Cargo.toml, etc.)",
			Fix:     "Ensure your project has a module/package manifest",
		})
		return
	}

	// Group by language
	byLang := map[string][]evidence.ModuleRoot{}
	for _, m := range roots {
		byLang[m.Language] = append(byLang[m.Language], m)
	}

	for lang, mods := range byLang {
		var locations []string
		for _, m := range mods {
			loc := m.RelPath
			if loc == "" {
				loc = "."
			}
			locations = append(locations, fmt.Sprintf("%s (%s)", loc, m.Marker))
		}
		msg := fmt.Sprintf("%d module root(s): %s", len(mods), strings.Join(locations, ", "))

		status := StatusPass
		// Warn if module is not at repo root (common source of missing evidence)
		hasRoot := false
		for _, m := range mods {
			if m.RelPath == "" {
				hasRoot = true
			}
		}
		fix := ""
		if !hasRoot && lang == "go" {
			status = StatusPass // Not a problem anymore — we handle nested modules
			msg += " (nested module — tools will run from module directory)"
		}

		r.Checks = append(r.Checks, Check{
			Name:    fmt.Sprintf("Module roots (%s)", lang),
			Group:   "modules",
			Status:  status,
			Message: msg,
			Fix:     fix,
		})
	}
}

// checkEnvironment checks Go and Git availability.
func (r *Report) checkEnvironment() {
	// Go
	if version, err := commandOutput("go", "version"); err == nil {
		r.Checks = append(r.Checks, Check{
			Name:    "Go compiler",
			Group:   "environment",
			Status:  StatusPass,
			Message: strings.TrimSpace(version),
		})
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "Go compiler",
			Group:   "environment",
			Status:  StatusFail,
			Message: "go not found in PATH",
			Fix:     "Install Go 1.22+ from https://go.dev",
		})
	}

	// Git
	if version, err := commandOutput("git", "--version"); err == nil {
		r.Checks = append(r.Checks, Check{
			Name:    "Git",
			Group:   "environment",
			Status:  StatusPass,
			Message: strings.TrimSpace(version),
		})
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "Git",
			Group:   "environment",
			Status:  StatusFail,
			Message: "git not found in PATH",
			Fix:     "Install Git from https://git-scm.com",
		})
	}
}

// checkProjectSetup checks .certification/ directory structure.
func (r *Report) checkProjectSetup(root string) {
	certDir := filepath.Join(root, ".certification")

	// Config file
	configPath := filepath.Join(certDir, "config.yml")
	if _, err := os.Stat(configPath); err == nil {
		r.Checks = append(r.Checks, Check{
			Name:    "Configuration file",
			Group:   "project",
			Status:  StatusPass,
			Message: ".certification/config.yml exists",
		})
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "Configuration file",
			Group:   "project",
			Status:  StatusFail,
			Message: ".certification/config.yml not found",
			Fix:     "Run: certify init",
		})
		return // Skip downstream checks
	}

	// Policies
	policiesDir := filepath.Join(certDir, "policies")
	if entries, err := os.ReadDir(policiesDir); err == nil && len(entries) > 0 {
		var names []string
		for _, e := range entries {
			if strings.HasSuffix(e.Name(), ".yml") || strings.HasSuffix(e.Name(), ".yaml") {
				names = append(names, e.Name())
			}
		}
		r.Checks = append(r.Checks, Check{
			Name:    "Policy packs",
			Group:   "project",
			Status:  StatusPass,
			Message: fmt.Sprintf("%d policy pack(s): %s", len(names), strings.Join(names, ", ")),
		})
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "Policy packs",
			Group:   "project",
			Status:  StatusWarn,
			Message: "No policy packs found in .certification/policies/",
			Fix:     "Run: certify init",
		})
	}

	// Index
	indexPath := filepath.Join(certDir, "index.json")
	if info, err := os.Stat(indexPath); err == nil {
		r.Checks = append(r.Checks, Check{
			Name:    "Unit index",
			Group:   "project",
			Status:  StatusPass,
			Message: fmt.Sprintf(".certification/index.json (%d bytes)", info.Size()),
		})
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "Unit index",
			Group:   "project",
			Status:  StatusWarn,
			Message: "No index found — units not yet discovered",
			Fix:     "Run: certify scan",
		})
	}

	// Records
	recordsDir := filepath.Join(certDir, "records")
	if entries, err := os.ReadDir(recordsDir); err == nil {
		jsonCount := 0
		for _, e := range entries {
			if strings.HasSuffix(e.Name(), ".json") && !strings.Contains(e.Name(), ".history.") {
				jsonCount++
			}
		}
		if jsonCount > 0 {
			r.Checks = append(r.Checks, Check{
				Name:    "Certification records",
				Group:   "project",
				Status:  StatusPass,
				Message: fmt.Sprintf("%d certification record(s)", jsonCount),
			})
		} else {
			r.Checks = append(r.Checks, Check{
				Name:    "Certification records",
				Group:   "project",
				Status:  StatusWarn,
				Message: "No certification records found",
				Fix:     "Run: certify certify",
			})
		}
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "Certification records",
			Group:   "project",
			Status:  StatusWarn,
			Message: "No records directory",
			Fix:     "Run: certify certify",
		})
	}

	// Report card
	reportPath := filepath.Join(certDir, "REPORT_CARD.md")
	if _, err := os.Stat(reportPath); err == nil {
		r.Checks = append(r.Checks, Check{
			Name:    "Report card",
			Group:   "project",
			Status:  StatusPass,
			Message: "REPORT_CARD.md exists",
		})
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "Report card",
			Group:   "project",
			Status:  StatusWarn,
			Message: "No report card generated yet",
			Fix:     "Run: certify report",
		})
	}

	// Badge
	badgePath := filepath.Join(certDir, "badge.json")
	if _, err := os.Stat(badgePath); err == nil {
		r.Checks = append(r.Checks, Check{
			Name:    "Badge endpoint",
			Group:   "project",
			Status:  StatusPass,
			Message: "badge.json exists",
		})
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "Badge endpoint",
			Group:   "project",
			Status:  StatusSkip,
			Message: "badge.json not found (generated by certify report)",
		})
	}
}

// checkConfiguration validates config.yml content.
func (r *Report) checkConfiguration(root string) {
	certDir := filepath.Join(root, ".certification")

	// Skip if config doesn't exist (already reported in project setup)
	if _, err := os.Stat(filepath.Join(certDir, "config.yml")); os.IsNotExist(err) {
		return
	}

	cfg, err := config.LoadFromDir(certDir)
	if err != nil {
		r.Checks = append(r.Checks, Check{
			Name:    "Configuration valid",
			Group:   "config",
			Status:  StatusFail,
			Message: fmt.Sprintf("Failed to parse config: %v", err),
			Fix:     "Check .certification/config.yml for YAML syntax errors",
		})
		return
	}

	errs := config.ValidateConfig(cfg)
	if len(errs) == 0 {
		r.Checks = append(r.Checks, Check{
			Name:    "Configuration valid",
			Group:   "config",
			Status:  StatusPass,
			Message: fmt.Sprintf("mode=%s, expiry=%dd", cfg.Mode, cfg.Expiry.DefaultWindowDays),
		})
	} else {
		for _, ve := range errs {
			r.Checks = append(r.Checks, Check{
				Name:    "Configuration: " + ve.Field,
				Group:   "config",
				Status:  StatusWarn,
				Message: ve.Message,
				Fix:     "Edit .certification/config.yml",
			})
		}
	}

	// Check scope: include: [] means nothing included
	if len(cfg.Scope.Include) == 0 {
		r.Checks = append(r.Checks, Check{
			Name:    "Scope includes",
			Group:   "config",
			Status:  StatusPass,
			Message: "scope.include is empty (includes everything)",
		})
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "Scope includes",
			Group:   "config",
			Status:  StatusPass,
			Message: fmt.Sprintf("scope.include has %d pattern(s)", len(cfg.Scope.Include)),
		})
	}

	// Check agent config
	if cfg.Agent.Enabled {
		r.checkAgentConfig(cfg)
	} else if cfg.Agent.ExplicitlyDisabled {
		r.Checks = append(r.Checks, Check{
			Name:    "Agent review",
			Group:   "config",
			Status:  StatusSkip,
			Message: "Explicitly disabled in config",
		})
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "Agent review",
			Group:   "config",
			Status:  StatusPass,
			Message: "Not configured (auto-detection will be used)",
		})
	}

	// Load and validate policy packs
	policiesDir := filepath.Join(certDir, "policies")
	packs, err := config.LoadPolicyPacks(policiesDir)
	if err == nil && len(packs) > 0 {
		for _, p := range packs {
			packErrs := config.ValidatePolicyPack(p)
			if len(packErrs) == 0 {
				r.Checks = append(r.Checks, Check{
					Name:    fmt.Sprintf("Policy: %s@%s", p.Name, p.Version),
					Group:   "policies",
					Status:  StatusPass,
					Message: fmt.Sprintf("%d rule(s)", len(p.Rules)),
				})
			} else {
				for _, ve := range packErrs {
					r.Checks = append(r.Checks, Check{
						Name:    fmt.Sprintf("Policy: %s — %s", p.Name, ve.Field),
						Group:   "policies",
						Status:  StatusWarn,
						Message: ve.Message,
					})
				}
			}
		}
	}
}

// checkAgentConfig validates the agent provider configuration.
func (r *Report) checkAgentConfig(cfg domain.Config) {
	if cfg.Agent.Provider.BaseURL == "" {
		r.Checks = append(r.Checks, Check{
			Name:    "Agent: base URL",
			Group:   "agent",
			Status:  StatusFail,
			Message: "agent.provider.base_url is empty",
			Fix:     "Add base_url to .certification/config.yml agent section",
		})
		return
	}

	r.Checks = append(r.Checks, Check{
		Name:    "Agent: base URL",
		Group:   "agent",
		Status:  StatusPass,
		Message: cfg.Agent.Provider.BaseURL,
	})

	// Check API key
	if cfg.Agent.Provider.APIKeyEnv != "" {
		if key := os.Getenv(cfg.Agent.Provider.APIKeyEnv); key != "" {
			r.Checks = append(r.Checks, Check{
				Name:    "Agent: API key",
				Group:   "agent",
				Status:  StatusPass,
				Message: fmt.Sprintf("$%s is set (%d chars)", cfg.Agent.Provider.APIKeyEnv, len(key)),
			})
		} else {
			isLocal := strings.Contains(cfg.Agent.Provider.BaseURL, "localhost") ||
				strings.Contains(cfg.Agent.Provider.BaseURL, "127.0.0.1")
			if isLocal {
				r.Checks = append(r.Checks, Check{
					Name:    "Agent: API key",
					Group:   "agent",
					Status:  StatusPass,
					Message: fmt.Sprintf("$%s not set (not required for local provider)", cfg.Agent.Provider.APIKeyEnv),
				})
			} else {
				r.Checks = append(r.Checks, Check{
					Name:    "Agent: API key",
					Group:   "agent",
					Status:  StatusFail,
					Message: fmt.Sprintf("$%s is not set", cfg.Agent.Provider.APIKeyEnv),
					Fix:     fmt.Sprintf("export %s=your-key-here", cfg.Agent.Provider.APIKeyEnv),
				})
			}
		}
	}

	// Check models
	if cfg.Agent.Models.Review != "" {
		r.Checks = append(r.Checks, Check{
			Name:    "Agent: model",
			Group:   "agent",
			Status:  StatusPass,
			Message: cfg.Agent.Models.Review,
		})
	} else if cfg.Agent.Models.Prescreen != "" {
		r.Checks = append(r.Checks, Check{
			Name:    "Agent: model",
			Group:   "agent",
			Status:  StatusPass,
			Message: cfg.Agent.Models.Prescreen + " (prescreen only)",
		})
	}
}

// checkTools checks for optional tools that improve evidence quality.
func (r *Report) checkTools() {
	// golangci-lint
	if version, err := commandOutput("golangci-lint", "--version"); err == nil {
		short := strings.TrimSpace(version)
		if idx := strings.Index(short, "\n"); idx > 0 {
			short = short[:idx]
		}
		r.Checks = append(r.Checks, Check{
			Name:    "golangci-lint",
			Group:   "tools",
			Status:  StatusPass,
			Message: short,
		})
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "golangci-lint",
			Group:   "tools",
			Status:  StatusWarn,
			Message: "Not found — lint evidence will be limited to go vet",
			Fix:     "Install: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh",
		})
	}

	// ESLint (TypeScript/JavaScript)
	if _, err := exec.LookPath("eslint"); err == nil {
		r.Checks = append(r.Checks, Check{Name: "eslint", Group: "tools", Status: StatusPass, Message: "eslint found"})
	} else if _, err := exec.LookPath("npx"); err == nil {
		r.Checks = append(r.Checks, Check{Name: "eslint", Group: "tools", Status: StatusPass, Message: "available via npx"})
	} else {
		r.addToolWarn("eslint", "TS/JS lint evidence will be unavailable", "npm install -D eslint")
	}

	// Ruff (Python linter)
	if version, err := commandOutput("ruff", "--version"); err == nil {
		r.Checks = append(r.Checks, Check{Name: "ruff", Group: "tools", Status: StatusPass, Message: strings.TrimSpace(version)})
	} else {
		r.addToolWarn("ruff", "Python lint evidence will be unavailable", "pip install ruff")
	}

	// pytest (Python test runner)
	if version, err := commandOutput("pytest", "--version"); err == nil {
		short := strings.TrimSpace(version)
		if idx := strings.Index(short, "\n"); idx > 0 {
			short = short[:idx]
		}
		r.Checks = append(r.Checks, Check{Name: "pytest", Group: "tools", Status: StatusPass, Message: short})
	} else {
		r.addToolWarn("pytest", "Python test evidence will be unavailable", "pip install pytest")
	}

	// Cargo (Rust — clippy + test)
	if _, err := exec.LookPath("cargo"); err == nil {
		r.Checks = append(r.Checks, Check{Name: "cargo", Group: "tools", Status: StatusPass, Message: "cargo found (clippy + test)"})
	} else {
		r.addToolWarn("cargo", "Rust lint/test evidence will be unavailable", "Install: https://rustup.rs")
	}

	// gh CLI
	if version, err := commandOutput("gh", "--version"); err == nil {
		short := strings.TrimSpace(version)
		if idx := strings.Index(short, "\n"); idx > 0 {
			short = short[:idx]
		}
		r.Checks = append(r.Checks, Check{
			Name:    "GitHub CLI (gh)",
			Group:   "tools",
			Status:  StatusPass,
			Message: short,
		})
	} else {
		r.Checks = append(r.Checks, Check{
			Name:    "GitHub CLI (gh)",
			Group:   "tools",
			Status:  StatusSkip,
			Message: "Not found — GitHub issue/PR integration unavailable",
			Fix:     "Install: https://cli.github.com",
		})
	}
}

// checkProviders checks for available AI providers.
func (r *Report) checkProviders() {
	providers := agent.DetectProviders()

	if len(providers) == 0 {
		r.Checks = append(r.Checks, Check{
			Name:    "AI providers",
			Group:   "providers",
			Status:  StatusSkip,
			Message: "No AI providers detected (certification works without them)",
			Fix:     "Set OPENROUTER_API_KEY or start Ollama for AI-assisted review",
		})
		return
	}

	for _, p := range providers {
		status := StatusPass
		msg := p.BaseURL
		if p.Local {
			msg += " (local)"
		} else {
			msg += fmt.Sprintf(" (key: %d chars)", len(p.APIKey))
		}
		if len(p.Models) > 0 {
			msg += fmt.Sprintf(", models: %s", strings.Join(p.Models[:min(3, len(p.Models))], ", "))
			if len(p.Models) > 3 {
				msg += fmt.Sprintf(" (+%d more)", len(p.Models)-3)
			}
		}
		r.Checks = append(r.Checks, Check{
			Name:    fmt.Sprintf("Provider: %s", p.Name),
			Group:   "providers",
			Status:  status,
			Message: msg,
		})
	}
}

// FormatReport renders the report as a human-readable string.
func FormatReport(r *Report) string {
	var b strings.Builder

	fmt.Fprintf(&b, "🩺 Certify Doctor — %s\n\n", r.Root)

	currentGroup := ""
	for _, c := range r.Checks {
		if c.Group != currentGroup {
			currentGroup = c.Group
			fmt.Fprintf(&b, "  %s\n", groupTitle(currentGroup))
		}
		fmt.Fprintf(&b, "    %s %s: %s\n", c.Status.Emoji(), c.Name, c.Message)
		if c.Fix != "" && c.Status != StatusPass {
			fmt.Fprintf(&b, "       → %s\n", c.Fix)
		}
	}

	pass, warn, fail, skip := r.Summary()
	fmt.Fprintf(&b, "\n  Summary: %d passed, %d warnings, %d failed, %d skipped\n", pass, warn, fail, skip)

	if fail > 0 {
		fmt.Fprintf(&b, "\n  ❌ Issues found. Fix the failures above to get started.\n")
	} else if warn > 0 {
		fmt.Fprintf(&b, "\n  ⚠️  Setup looks good with minor recommendations.\n")
	} else {
		fmt.Fprintf(&b, "\n  ✅ Everything looks good!\n")
	}

	return b.String()
}

func groupTitle(group string) string {
	switch group {
	case "environment":
		return "── Environment ──"
	case "project":
		return "── Project Setup ──"
	case "config":
		return "── Configuration ──"
	case "policies":
		return "── Policy Packs ──"
	case "tools":
		return "── Optional Tools ──"
	case "agent":
		return "── Agent Configuration ──"
	case "providers":
		return "── AI Providers ──"
	default:
		return "── " + group + " ──"
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// addToolWarn appends a warning check for a missing optional tool.
func (r *Report) addToolWarn(name, msg, fix string) {
	r.Checks = append(r.Checks, Check{
		Name:    name,
		Group:   "tools",
		Status:  StatusWarn,
		Message: msg,
		Fix:     fix,
	})
}

// commandOutput runs a command and returns its stdout.
func commandOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
