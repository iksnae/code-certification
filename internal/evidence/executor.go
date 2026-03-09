package evidence

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/code-certification/certify/internal/domain"
)

// ToolExecutor runs external tools and collects evidence.
type ToolExecutor struct {
	root string
}

// NewToolExecutor creates a tool executor rooted at the given directory.
func NewToolExecutor(root string) *ToolExecutor {
	return &ToolExecutor{root: root}
}

// CollectAll runs all available tool runners and returns collected evidence.
func (te *ToolExecutor) CollectAll() []domain.Evidence {
	var ev []domain.Evidence

	if e := te.runGoVet(); e != nil {
		ev = append(ev, *e)
	}
	if e := te.runGoTest(); e != nil {
		ev = append(ev, *e)
	}
	if e := te.runGolangciLint(); e != nil {
		ev = append(ev, *e)
	}
	if e := te.runGitStats(); e != nil {
		ev = append(ev, *e)
	}

	return ev
}

// HasGoMod returns true if the root directory contains a go.mod file.
func (te *ToolExecutor) HasGoMod() bool {
	_, err := os.Stat(filepath.Join(te.root, "go.mod"))
	return err == nil
}

// HasPackageJSON returns true if the root has a package.json.
func (te *ToolExecutor) HasPackageJSON() bool {
	_, err := os.Stat(filepath.Join(te.root, "package.json"))
	return err == nil
}

func (te *ToolExecutor) runGoVet() *domain.Evidence {
	if !te.HasGoMod() {
		return nil
	}
	if _, err := exec.LookPath("go"); err != nil {
		return nil
	}

	cmd := exec.Command("go", "vet", "./...")
	cmd.Dir = te.root
	output, err := cmd.CombinedOutput()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		}
	}

	result := ParseGoVet(string(output), exitCode)
	ev := result.ToEvidence()
	return &ev
}

func (te *ToolExecutor) runGoTest() *domain.Evidence {
	if !te.HasGoMod() {
		return nil
	}
	if _, err := exec.LookPath("go"); err != nil {
		return nil
	}

	// Run go test with JSON output and coverage
	cmd := exec.Command("go", "test", "-json", "-count=1", "./...")
	cmd.Dir = te.root
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
	coverCmd.Dir = te.root
	coverCmd.Run() // Best effort

	if data, err := os.ReadFile(coverFile); err == nil {
		result.Coverage = ParseCoverProfile(string(data))
	}

	ev := result.ToEvidence()
	return &ev
}

func (te *ToolExecutor) runGolangciLint() *domain.Evidence {
	if !te.HasGoMod() {
		return nil
	}
	if _, err := exec.LookPath("golangci-lint"); err != nil {
		return nil // Not installed, skip
	}

	cmd := exec.Command("golangci-lint", "run", "--out-format", "json", "./...")
	cmd.Dir = te.root
	output, _ := cmd.CombinedOutput()

	result := ParseGolangciLintJSON(string(output))
	ev := result.ToEvidence()
	return &ev
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
