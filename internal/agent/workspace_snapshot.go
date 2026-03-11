package agent

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// SubmoduleInfo holds the input data for building a workspace snapshot.
// This is populated by the CLI from workspace.Submodule + loaded records.
type SubmoduleInfo struct {
	Name    string
	Path    string
	Commit  string
	Records []domain.CertificationRecord
}

// WorkspaceArchSnapshot extends the architecture snapshot with cross-submodule data.
type WorkspaceArchSnapshot struct {
	SchemaVersion      int                      `json:"schema_version"`
	SubmoduleSnapshots []SubmoduleSnapshotEntry `json:"submodule_snapshots"`
	CrossDependencies  []CrossDepEdge           `json:"cross_dependencies"`
	InfraFiles         []string                 `json:"infra_files"`
	AggregateMetrics   WorkspaceMetrics         `json:"aggregate_metrics"`
}

// SubmoduleSnapshotEntry holds a per-submodule architecture snapshot.
type SubmoduleSnapshotEntry struct {
	Name     string        `json:"name"`
	Path     string        `json:"path"`
	Commit   string        `json:"commit"`
	Snapshot *ArchSnapshot `json:"snapshot"`
	Role     string        `json:"role"` // "service", "library", "tool", "infrastructure"
}

// CrossDepEdge represents a dependency between two submodules.
type CrossDepEdge struct {
	FromSubmodule string `json:"from_submodule"`
	ToSubmodule   string `json:"to_submodule"`
	Evidence      string `json:"evidence"`
	Weight        int    `json:"weight"`
}

// WorkspaceMetrics holds aggregate metrics across all submodules.
type WorkspaceMetrics struct {
	TotalSubmodules      int     `json:"total_submodules"`
	ConfiguredSubmodules int     `json:"configured_submodules"`
	TotalUnitsAcrossAll  int     `json:"total_units_across_all"`
	WeightedAvgScore     float64 `json:"weighted_avg_score"`
	WorstSubmodule       string  `json:"worst_submodule"`
	BestSubmodule        string  `json:"best_submodule"`
}

// BuildWorkspaceSnapshot builds a workspace-level architecture snapshot
// from submodule info. Each submodule gets its own ArchSnapshot built
// from its certification records.
func BuildWorkspaceSnapshot(root string, subs []SubmoduleInfo) *WorkspaceArchSnapshot {
	snap := &WorkspaceArchSnapshot{
		SchemaVersion: SnapshotSchemaVersion,
	}

	snap.InfraFiles = detectInfraFiles(root)

	if len(subs) == 0 {
		return snap
	}

	snap.AggregateMetrics.TotalSubmodules = len(subs)
	snap.AggregateMetrics.ConfiguredSubmodules = len(subs)

	var totalUnits int
	var totalWeightedScore float64
	bestScore := -1.0
	worstScore := 2.0

	for _, sub := range subs {
		subRoot := filepath.Join(root, sub.Path)
		subSnap := BuildSnapshot(sub.Records, subRoot)
		role := classifySubmoduleRole(root, sub.Path)

		entry := SubmoduleSnapshotEntry{
			Name:     sub.Name,
			Path:     sub.Path,
			Commit:   sub.Commit,
			Snapshot: subSnap,
			Role:     role,
		}
		snap.SubmoduleSnapshots = append(snap.SubmoduleSnapshots, entry)

		units := subSnap.Metrics.TotalUnits
		totalUnits += units

		if units > 0 {
			avg := subSnap.Metrics.AvgScore
			totalWeightedScore += avg * float64(units)

			if avg > bestScore {
				bestScore = avg
				snap.AggregateMetrics.BestSubmodule = sub.Name
			}
			if avg < worstScore {
				worstScore = avg
				snap.AggregateMetrics.WorstSubmodule = sub.Name
			}
		}
	}

	snap.AggregateMetrics.TotalUnitsAcrossAll = totalUnits
	if totalUnits > 0 {
		snap.AggregateMetrics.WeightedAvgScore = totalWeightedScore / float64(totalUnits)
	}

	// Detect cross-submodule dependencies
	snap.CrossDependencies = detectCrossDeps(root, subs)

	return snap
}

// classifySubmoduleRole determines a submodule's role based on filesystem heuristics.
func classifySubmoduleRole(root, subPath string) string {
	absPath := filepath.Join(root, subPath)

	// Check for cmd/ directory with main.go → service or tool
	cmdDir := filepath.Join(absPath, "cmd")
	if entries, err := os.ReadDir(cmdDir); err == nil {
		for _, e := range entries {
			if e.IsDir() {
				mainPath := filepath.Join(cmdDir, e.Name(), "main.go")
				if _, err := os.Stat(mainPath); err == nil {
					// Heuristic: "services/" in path → service, "tools/" → tool
					if strings.Contains(subPath, "service") || strings.Contains(subPath, "api") {
						return "service"
					}
					if strings.Contains(subPath, "tool") {
						return "tool"
					}
					return "service" // default for things with cmd/
				}
			}
		}
	}

	// Check for pkg/ directory → library
	if _, err := os.Stat(filepath.Join(absPath, "pkg")); err == nil {
		return "library"
	}

	// Check for exported Go packages (no cmd/) → library
	if _, err := os.Stat(filepath.Join(absPath, "go.mod")); err == nil {
		if _, err := os.Stat(cmdDir); os.IsNotExist(err) {
			return "library"
		}
	}

	return "component" // generic fallback
}

// detectCrossDeps finds cross-submodule dependencies by parsing go.mod replace directives.
func detectCrossDeps(root string, subs []SubmoduleInfo) []CrossDepEdge {
	// Build a map of submodule paths for quick lookup
	subPaths := make(map[string]string) // normalized path → sub.Path
	for _, sub := range subs {
		subPaths[sub.Path] = sub.Path
	}

	var edges []CrossDepEdge

	for _, sub := range subs {
		goModPath := filepath.Join(root, sub.Path, "go.mod")
		data, err := os.ReadFile(goModPath)
		if err != nil {
			continue
		}

		replacePaths := parseGoModReplaces(string(data))
		for _, rp := range replacePaths {
			// Resolve relative replace path to workspace-relative path
			target := resolveReplacePath(sub.Path, rp)
			if targetSub, ok := findMatchingSubmodule(target, subPaths); ok {
				edges = append(edges, CrossDepEdge{
					FromSubmodule: sub.Path,
					ToSubmodule:   targetSub,
					Evidence:      "go.mod replace",
					Weight:        1,
				})
			}
		}
	}

	return edges
}

// parseGoModReplaces extracts local replace directive paths from go.mod content.
func parseGoModReplaces(content string) []string {
	var paths []string
	inBlock := false

	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)

		if trimmed == "replace (" {
			inBlock = true
			continue
		}
		if inBlock && trimmed == ")" {
			inBlock = false
			continue
		}

		var replaceLine string
		if inBlock {
			replaceLine = trimmed
		} else if strings.HasPrefix(trimmed, "replace ") {
			replaceLine = strings.TrimPrefix(trimmed, "replace ")
		} else {
			continue
		}

		// Parse: module => path
		parts := strings.Split(replaceLine, "=>")
		if len(parts) != 2 {
			continue
		}
		target := strings.TrimSpace(parts[1])
		// Only interested in local (relative) replacements
		if strings.HasPrefix(target, ".") || strings.HasPrefix(target, "/") {
			paths = append(paths, target)
		}
	}
	return paths
}

// resolveReplacePath resolves a relative replace path against a submodule's location.
func resolveReplacePath(fromSubPath, replacePath string) string {
	joined := filepath.Join(fromSubPath, replacePath)
	cleaned := filepath.Clean(joined)
	return filepath.ToSlash(cleaned)
}

// findMatchingSubmodule finds a submodule whose path matches the target.
func findMatchingSubmodule(target string, subPaths map[string]string) (string, bool) {
	// Direct match
	if sp, ok := subPaths[target]; ok {
		return sp, true
	}
	// Try cleaning
	cleaned := filepath.ToSlash(filepath.Clean(target))
	if sp, ok := subPaths[cleaned]; ok {
		return sp, true
	}
	return "", false
}

// detectInfraFiles finds infrastructure files at the workspace root.
func detectInfraFiles(root string) []string {
	infraPatterns := []string{
		"Justfile",
		"Makefile",
		"docker-compose.yml",
		"docker-compose.yaml",
		"Dockerfile",
		"Tiltfile",
		"Vagrantfile",
		"Procfile",
		"netlify.toml",
		"vercel.json",
		"fly.toml",
	}

	var found []string

	// Check root-level infra files
	for _, pattern := range infraPatterns {
		path := filepath.Join(root, pattern)
		if _, err := os.Stat(path); err == nil {
			found = append(found, pattern)
		}
	}

	// Check .github/workflows/
	workflowDir := filepath.Join(root, ".github", "workflows")
	if entries, err := os.ReadDir(workflowDir); err == nil {
		for _, e := range entries {
			if !e.IsDir() && (strings.HasSuffix(e.Name(), ".yml") || strings.HasSuffix(e.Name(), ".yaml")) {
				found = append(found, filepath.ToSlash(filepath.Join(".github/workflows", e.Name())))
			}
		}
	}

	sort.Strings(found)
	return found
}

// WorkspaceProjectContext holds everything needed for a workspace-level architect review.
type WorkspaceProjectContext struct {
	RepoName      string
	CommitSHA     string
	Snapshot      *WorkspaceArchSnapshot
	Documentation map[string]string // workspace-level docs
}

// GatherWorkspaceContext assembles a WorkspaceProjectContext for architect review.
func GatherWorkspaceContext(root string, subs []SubmoduleInfo) *WorkspaceProjectContext {
	wpc := &WorkspaceProjectContext{
		Documentation: make(map[string]string),
	}

	wpc.Snapshot = BuildWorkspaceSnapshot(root, subs)
	wpc.RepoName = detectRepoNameFromDir(root)

	// Load workspace-level docs
	docFiles := []string{"README.md", "CLAUDE.md", "PRD.md", "FEATURES.md"}
	for _, f := range docFiles {
		path := filepath.Join(root, f)
		data, err := os.ReadFile(path)
		if err == nil {
			content := string(data)
			if len(content) > 4000 {
				content = content[:4000] + "\n... (truncated)"
			}
			wpc.Documentation[f] = content
		}
	}

	return wpc
}

// detectRepoNameFromDir attempts to detect the repo name from git remote.
func detectRepoNameFromDir(root string) string {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return filepath.Base(root)
	}
	url := strings.TrimSpace(string(out))
	url = strings.TrimSuffix(url, ".git")
	if idx := strings.Index(url, "github.com/"); idx >= 0 {
		return url[idx+len("github.com/"):]
	}
	return filepath.Base(root)
}

// FormatWorkspaceForLLM serializes the workspace snapshot for LLM consumption.
func FormatWorkspaceForLLM(snap *WorkspaceArchSnapshot, maxTokensHint int) string {
	var b strings.Builder
	maxChars := maxTokensHint * 4

	b.WriteString("# Workspace Architecture Snapshot\n\n")
	if snap.SchemaVersion > 0 {
		fmt.Fprintf(&b, "**Snapshot Schema:** v%d (workspace)\n", snap.SchemaVersion)
	}
	fmt.Fprintf(&b, "**Submodules:** %d (%d configured)\n",
		snap.AggregateMetrics.TotalSubmodules, snap.AggregateMetrics.ConfiguredSubmodules)
	fmt.Fprintf(&b, "**Total Units:** %d\n", snap.AggregateMetrics.TotalUnitsAcrossAll)
	if snap.AggregateMetrics.WeightedAvgScore > 0 {
		fmt.Fprintf(&b, "**Weighted Avg Score:** %.1f%%\n", snap.AggregateMetrics.WeightedAvgScore*100)
	}
	if snap.AggregateMetrics.BestSubmodule != "" {
		fmt.Fprintf(&b, "**Best:** %s · **Worst:** %s\n",
			snap.AggregateMetrics.BestSubmodule, snap.AggregateMetrics.WorstSubmodule)
	}
	b.WriteString("\n")

	// Submodule overview table
	if len(snap.SubmoduleSnapshots) > 0 {
		b.WriteString("## Submodule Overview\n")
		b.WriteString("| Submodule | Role | Units | Avg Score | Grade | Packages | Commit |\n")
		b.WriteString("|-----------|------|------:|----------:|:-----:|--------:|--------|\n")
		for _, entry := range snap.SubmoduleSnapshots {
			units, avg, grade, pkgs := 0, 0.0, "N/A", 0
			if entry.Snapshot != nil {
				units = entry.Snapshot.Metrics.TotalUnits
				avg = entry.Snapshot.Metrics.AvgScore
				grade = domain.GradeFromScore(avg).String()
				pkgs = entry.Snapshot.Metrics.TotalPackages
			}
			commit := entry.Commit
			if len(commit) > 7 {
				commit = commit[:7]
			}
			fmt.Fprintf(&b, "| %s | %s | %d | %.1f%% | %s | %d | `%s` |\n",
				entry.Name, entry.Role, units, avg*100, grade, pkgs, commit)
		}
		b.WriteString("\n")
	}

	// Cross-submodule dependencies
	if len(snap.CrossDependencies) > 0 {
		b.WriteString("## Cross-Submodule Dependencies\n")
		b.WriteString("| From | To | Evidence |\n")
		b.WriteString("|------|----|---------|\n")
		for _, dep := range snap.CrossDependencies {
			fmt.Fprintf(&b, "| %s | %s | %s |\n", dep.FromSubmodule, dep.ToSubmodule, dep.Evidence)
		}
		b.WriteString("\n")
	}

	// Infrastructure files
	if len(snap.InfraFiles) > 0 {
		b.WriteString("## Infrastructure Files\n")
		for _, f := range snap.InfraFiles {
			fmt.Fprintf(&b, "- %s\n", f)
		}
		b.WriteString("\n")
	}

	// Per-submodule condensed summaries (if within token budget)
	if b.Len() < maxChars-2000 {
		for _, entry := range snap.SubmoduleSnapshots {
			if entry.Snapshot == nil || entry.Snapshot.Metrics.TotalUnits == 0 {
				continue
			}
			if b.Len() > maxChars-1000 {
				b.WriteString("\n... (remaining submodule details truncated for token budget)\n")
				break
			}
			formatCondensedSubmodule(&b, entry)
		}
	}

	return b.String()
}

// formatCondensedSubmodule writes a condensed per-submodule summary.
func formatCondensedSubmodule(b *strings.Builder, entry SubmoduleSnapshotEntry) {
	snap := entry.Snapshot
	fmt.Fprintf(b, "### %s (%s) — %s\n", entry.Name, entry.Path, entry.Role)
	fmt.Fprintf(b, "- Units: %d · Packages: %d · Avg Score: %.1f%%\n",
		snap.Metrics.TotalUnits, snap.Metrics.TotalPackages, snap.Metrics.AvgScore*100)

	// Grade distribution
	if len(snap.Metrics.GradeDistribution) > 0 {
		var grades []string
		for _, g := range []string{"A", "A-", "B+", "B", "C", "D", "F"} {
			if c, ok := snap.Metrics.GradeDistribution[g]; ok && c > 0 {
				grades = append(grades, fmt.Sprintf("%s:%d", g, c))
			}
		}
		if len(grades) > 0 {
			fmt.Fprintf(b, "- Grades: %s\n", strings.Join(grades, " "))
		}
	}

	// Hotspots (top 3)
	if len(snap.Hotspots) > 0 {
		limit := 3
		if len(snap.Hotspots) < limit {
			limit = len(snap.Hotspots)
		}
		b.WriteString("- Hotspots: ")
		var hs []string
		for _, h := range snap.Hotspots[:limit] {
			hs = append(hs, fmt.Sprintf("%s(%.0f%%)", h.Path, h.AvgScore*100))
		}
		b.WriteString(strings.Join(hs, ", "))
		b.WriteString("\n")
	}

	// Key structural metrics (non-zero only)
	s := snap.Metrics.Structural
	var structIssues []string
	if s.PanicCalls > 0 {
		structIssues = append(structIssues, fmt.Sprintf("panic:%d", s.PanicCalls))
	}
	if s.ErrorsIgnored > 0 {
		structIssues = append(structIssues, fmt.Sprintf("errors_ignored:%d", s.ErrorsIgnored))
	}
	if s.GlobalMutableCount > 0 {
		structIssues = append(structIssues, fmt.Sprintf("global_mutable:%d", s.GlobalMutableCount))
	}
	if len(structIssues) > 0 {
		fmt.Fprintf(b, "- Structural: %s\n", strings.Join(structIssues, ", "))
	}

	b.WriteString("\n")
}
