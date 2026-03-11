package evidence

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

// ToolExecutor runs external tools and collects evidence.
// After CollectAll(), raw lint findings and coverage profile are retained
// for per-unit attribution by the certification pipeline.
//
// ToolExecutor discovers nested module roots (go.mod in subdirectories)
// and runs tools from each module root. This supports monorepos and repos
// where the Go module is not at the repository root.
type ToolExecutor struct {
	root            string
	moduleRoots     []ModuleRoot // discovered module roots
	rawLintFindings []LintFinding
	rawCoverProfile string
}

// NewToolExecutor creates a tool executor rooted at the given directory.
// It discovers all module roots (go.mod, package.json, etc.) in the tree.
func NewToolExecutor(root string) *ToolExecutor {
	return &ToolExecutor{
		root:        root,
		moduleRoots: DiscoverModuleRoots(root),
	}
}

// CollectAll runs all available tool runners and returns collected evidence.
// Discovers module roots and runs language-appropriate tools from each.
func (te *ToolExecutor) CollectAll() []domain.Evidence {
	var ev []domain.Evidence

	// Go modules
	for _, mod := range filterModulesByLang(te.moduleRoots, "go") {
		if e := te.runGoVetAt(mod.Path); e != nil {
			ev = append(ev, *e)
		}
		if e := te.runGoTestAt(mod.Path); e != nil {
			ev = append(ev, *e)
		}
		if e := te.runGolangciLintAt(mod.Path); e != nil {
			ev = append(ev, *e)
		}
	}

	// TypeScript/JavaScript modules
	for _, mod := range filterModulesByLang(te.moduleRoots, "ts") {
		if e := te.runESLintAt(mod.Path); e != nil {
			ev = append(ev, *e)
		}
		if e := te.runJSTestAt(mod.Path); e != nil {
			ev = append(ev, *e)
		}
	}

	// Python modules
	for _, mod := range filterModulesByLang(te.moduleRoots, "py") {
		if e := te.runRuffAt(mod.Path); e != nil {
			ev = append(ev, *e)
		}
		if e := te.runPytestAt(mod.Path); e != nil {
			ev = append(ev, *e)
		}
	}

	// Rust modules
	for _, mod := range filterModulesByLang(te.moduleRoots, "rs") {
		if e := te.runCargoClippyAt(mod.Path); e != nil {
			ev = append(ev, *e)
		}
		if e := te.runCargoTestAt(mod.Path); e != nil {
			ev = append(ev, *e)
		}
	}

	if e := te.runGitStats(); e != nil {
		ev = append(ev, *e)
	}

	return ev
}

// filterModulesByLang returns module roots matching the given language.
func filterModulesByLang(roots []ModuleRoot, lang string) []ModuleRoot {
	var filtered []ModuleRoot
	for _, r := range roots {
		if r.Language == lang {
			filtered = append(filtered, r)
		}
	}
	return filtered
}

// HasGoMod returns true if any discovered module root is a Go module.
func (te *ToolExecutor) HasGoMod() bool {
	return len(GoModuleRoots(te.moduleRoots)) > 0
}

// ModuleRoots returns the discovered module roots.
func (te *ToolExecutor) ModuleRoots() []ModuleRoot {
	return te.moduleRoots
}

// HasPackageJSON returns true if the root has a package.json.
func (te *ToolExecutor) HasPackageJSON() bool {
	_, err := os.Stat(filepath.Join(te.root, "package.json"))
	return err == nil
}

func (te *ToolExecutor) runGoVetAt(dir string) *domain.Evidence {
	if _, err := exec.LookPath("go"); err != nil {
		return nil
	}

	cmd := exec.Command("go", "vet", "./...")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		}
	}

	result := ParseGoVet(string(output), exitCode)
	// Adjust finding paths to be relative to repo root, not module root
	te.adjustFindingPaths(result.Findings, dir)
	te.rawLintFindings = append(te.rawLintFindings, result.Findings...)
	ev := result.ToEvidence()
	return &ev
}

func (te *ToolExecutor) runGoTestAt(dir string) *domain.Evidence {
	if _, err := exec.LookPath("go"); err != nil {
		return nil
	}

	// Run go test with JSON output and coverage
	cmd := exec.Command("go", "test", "-json", "-count=1", "./...")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()

	result := ParseGoTestJSON(string(output))
	if err != nil && result.TotalCount == 0 {
		// Complete failure to run tests
		result.FailedCount = 1
		result.TotalCount = 1
	}

	// Try to get coverage
	coverFile := filepath.Join(os.TempDir(), fmt.Sprintf("certify-cover-%d.out", time.Now().UnixNano()))
	defer os.Remove(coverFile)

	coverCmd := exec.Command("go", "test", "-coverprofile", coverFile, "-count=1", "./...")
	coverCmd.Dir = dir
	coverCmd.Run() // Best effort

	if data, err := os.ReadFile(coverFile); err == nil {
		// Append to existing cover profile (multi-module)
		if te.rawCoverProfile != "" {
			te.rawCoverProfile += "\n" + string(data)
		} else {
			te.rawCoverProfile = string(data)
		}
		result.Coverage = ParseCoverProfile(string(data))
	}

	ev := result.ToEvidence()
	return &ev
}

func (te *ToolExecutor) runGolangciLintAt(dir string) *domain.Evidence {
	if _, err := exec.LookPath("golangci-lint"); err != nil {
		return nil // Not installed, skip
	}

	cmd := exec.Command("golangci-lint", "run", "--out-format", "json", "./...")
	cmd.Dir = dir
	// golangci-lint returns non-zero when findings exist, but output is still valid JSON
	output, err := cmd.CombinedOutput()
	if err != nil && len(output) == 0 {
		return nil // actual failure (not just lint findings)
	}

	result := ParseGolangciLintJSON(string(output))
	// Adjust finding paths to be relative to repo root
	te.adjustFindingPaths(result.Findings, dir)
	te.rawLintFindings = append(te.rawLintFindings, result.Findings...)
	ev := result.ToEvidence()
	return &ev
}

// adjustFindingPaths converts lint finding paths from module-relative to repo-relative.
// When running go vet from code/, a finding at "main.go" should become "code/main.go"
// for per-unit attribution to work.
func (te *ToolExecutor) adjustFindingPaths(findings []LintFinding, moduleDir string) {
	if moduleDir == te.root {
		return // No adjustment needed for root module
	}
	prefix, err := filepath.Rel(te.root, moduleDir)
	if err != nil || prefix == "." {
		return
	}
	prefix = filepath.ToSlash(prefix)
	for i := range findings {
		if findings[i].File != "" && !strings.HasPrefix(findings[i].File, prefix) {
			findings[i].File = prefix + "/" + findings[i].File
		}
	}
}

func (te *ToolExecutor) runGitStats() *domain.Evidence {
	if _, err := exec.LookPath("git"); err != nil {
		return nil
	}

	// Get overall repo stats
	cmd := exec.Command("git", "log", "--format=%H\t%an\t%ad", "--date=short")
	cmd.Dir = te.root
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil
	}

	// Find earliest date
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var earliest string
	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if len(parts) >= 3 {
			earliest = parts[2] // Last line = earliest commit
		}
	}

	stats := ParseGitLogWithAge(string(output), earliest)
	ev := stats.ToEvidence()
	return &ev
}

// LintFindings returns all lint findings collected during CollectAll().
// Returns nil if no lint tools were run or found no issues.
func (te *ToolExecutor) LintFindings() []LintFinding {
	return te.rawLintFindings
}

// CoverageProfile returns the raw Go coverage profile collected during CollectAll().
// Returns empty string if coverage was not collected.
func (te *ToolExecutor) CoverageProfile() string {
	return te.rawCoverProfile
}
