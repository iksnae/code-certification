package agent

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// ProjectContext aggregates everything the architect reviewer needs.
type ProjectContext struct {
	RepoName      string
	CommitSHA     string
	Languages     []string
	Snapshot      *ArchSnapshot
	LowestUnits   []UnitSummary     // bottom 20 by score
	Documentation map[string]string // filename → content (README, PRD, etc.)
	FileTree      string
	PolicyRules   []string
	GitSummary    string
}

// UnitSummary is a lightweight unit representation for LLM context.
type UnitSummary struct {
	ID           string
	Score        float64
	Grade        string
	Observations []string
}

// GatherContext assembles a ProjectContext from certification records and the filesystem.
func GatherContext(root, certDir string, records []domain.CertificationRecord) (*ProjectContext, error) {
	pc := &ProjectContext{
		Documentation: make(map[string]string),
	}

	// Build snapshot
	pc.Snapshot = BuildSnapshot(records, root)

	// Detect languages from records
	langSet := make(map[string]bool)
	for _, r := range records {
		lang := r.UnitID.Language()
		if lang != "" && lang != "file" {
			langSet[lang] = true
		}
	}
	for lang := range langSet {
		pc.Languages = append(pc.Languages, lang)
	}
	sort.Strings(pc.Languages)

	// Load documentation files
	docFiles := []string{"README.md", "CLAUDE.md", "PRD.md", "FEATURES.md", "docs/internal/PRD.md"}
	for _, f := range docFiles {
		path := filepath.Join(root, f)
		data, err := os.ReadFile(path)
		if err == nil {
			content := string(data)
			if len(content) > 4000 {
				content = content[:4000] + "\n... (truncated)"
			}
			pc.Documentation[f] = content
		}
	}

	// Build lowest-scoring units
	type scoredUnit struct {
		id    string
		score float64
		grade string
		obs   []string
	}
	var units []scoredUnit
	for _, r := range records {
		units = append(units, scoredUnit{
			id:    r.UnitID.String(),
			score: r.Score,
			grade: r.Grade.String(),
			obs:   r.Observations,
		})
	}
	sort.Slice(units, func(i, j int) bool {
		return units[i].score < units[j].score
	})
	limit := 20
	if len(units) < limit {
		limit = len(units)
	}
	for _, u := range units[:limit] {
		pc.LowestUnits = append(pc.LowestUnits, UnitSummary{
			ID: u.id, Score: u.score, Grade: u.grade, Observations: u.obs,
		})
	}

	// Build file tree (limited depth)
	pc.FileTree = buildFileTree(root, 3)

	// Load policy rules
	policiesDir := filepath.Join(certDir, "policies")
	if entries, err := os.ReadDir(policiesDir); err == nil {
		for _, e := range entries {
			if strings.HasSuffix(e.Name(), ".yml") || strings.HasSuffix(e.Name(), ".yaml") {
				pc.PolicyRules = append(pc.PolicyRules, e.Name())
			}
		}
	}

	return pc, nil
}

// FormatForLLM serializes the project context into a structured text block for LLM consumption.
// maxTokensHint suggests a target size (~4 chars per token).
func (pc *ProjectContext) FormatForLLM(maxTokensHint int) string {
	var b strings.Builder
	maxChars := maxTokensHint * 4

	b.WriteString("# Architecture Snapshot\n\n")
	formatHeader(&b, pc)

	snap := pc.Snapshot
	if snap == nil {
		b.WriteString("No certification data available.\n")
		return b.String()
	}

	formatSnapshotMetrics(&b, snap)
	formatPackageMap(&b, snap)
	formatDependencyGraph(&b, snap)
	formatLayerMap(&b, snap)
	formatHotspots(&b, snap)
	formatCouplingPairs(&b, snap)
	formatTopObservations(&b, snap)
	formatStructuralMetrics(&b, snap)
	formatCoverageMetrics(&b, snap)
	formatCodeMetrics(&b, snap)
	formatQualitativeContext(&b, pc, maxChars)

	return b.String()
}

func formatHeader(b *strings.Builder, pc *ProjectContext) {
	if pc.RepoName != "" {
		fmt.Fprintf(b, "**Repository:** %s\n", pc.RepoName)
	}
	if pc.CommitSHA != "" {
		fmt.Fprintf(b, "**Commit:** %s\n", pc.CommitSHA)
	}
	if len(pc.Languages) > 0 {
		fmt.Fprintf(b, "**Languages:** %s\n", strings.Join(pc.Languages, ", "))
	}
	if pc.Snapshot != nil && pc.Snapshot.SchemaVersion > 0 {
		fmt.Fprintf(b, "**Snapshot Schema:** v%d\n", pc.Snapshot.SchemaVersion)
	}
	b.WriteString("\n")
}

func formatSnapshotMetrics(b *strings.Builder, snap *ArchSnapshot) {
	fmt.Fprintf(b, "## Aggregate Metrics\n")
	fmt.Fprintf(b, "- Total Units: %d\n", snap.Metrics.TotalUnits)
	fmt.Fprintf(b, "- Total Packages: %d\n", snap.Metrics.TotalPackages)
	fmt.Fprintf(b, "- Average Score: %.1f%%\n\n", snap.Metrics.AvgScore*100)

	b.WriteString("## Grade Distribution\n")
	for _, g := range []string{"A", "A-", "B+", "B", "C", "D", "F"} {
		if count, ok := snap.Metrics.GradeDistribution[g]; ok && count > 0 {
			fmt.Fprintf(b, "- %s: %d\n", g, count)
		}
	}
	b.WriteString("\n")
}

func formatPackageMap(b *strings.Builder, snap *ArchSnapshot) {
	b.WriteString("## Package Map\n")
	b.WriteString("| Package | Units | Avg Score | Grade | Observations | Top Issues |\n")
	b.WriteString("|---------|------:|----------:|:-----:|-------------:|------------|\n")
	for _, pkg := range snap.Packages {
		issues := "-"
		if len(pkg.TopIssues) > 0 {
			issues = strings.Join(pkg.TopIssues, ", ")
		}
		fmt.Fprintf(b, "| %s | %d | %.1f%% | %s | %d | %s |\n",
			pkg.Path, pkg.Units, pkg.AvgScore*100, pkg.Grade, pkg.Observations, issues)
	}
	b.WriteString("\n")
}

func formatDependencyGraph(b *strings.Builder, snap *ArchSnapshot) {
	if len(snap.DependencyEdges) == 0 {
		return
	}
	b.WriteString("## Dependency Graph\n```\n")
	depMap := make(map[string][]string)
	for _, e := range snap.DependencyEdges {
		depMap[e.From] = append(depMap[e.From], e.To)
	}
	var fromPkgs []string
	for k := range depMap {
		fromPkgs = append(fromPkgs, k)
	}
	sort.Strings(fromPkgs)
	for _, from := range fromPkgs {
		fmt.Fprintf(b, "%s → [%s]\n", from, strings.Join(depMap[from], ", "))
	}
	b.WriteString("```\n\n")
}

func formatLayerMap(b *strings.Builder, snap *ArchSnapshot) {
	if len(snap.Layers) == 0 {
		return
	}
	b.WriteString("## Layer Map\n")
	layerPkgs := make(map[string][]string)
	for pkg, layer := range snap.Layers {
		layerPkgs[layer] = append(layerPkgs[layer], pkg)
	}
	for _, layer := range []string{"cmd", "internal", "domain", "pkg", "other"} {
		pkgs := layerPkgs[layer]
		if len(pkgs) == 0 {
			continue
		}
		sort.Strings(pkgs)
		fmt.Fprintf(b, "- **%s**: %s\n", layer, strings.Join(pkgs, ", "))
	}
	b.WriteString("\n")
}

func formatHotspots(b *strings.Builder, snap *ArchSnapshot) {
	if len(snap.Hotspots) == 0 {
		return
	}
	b.WriteString("## Hotspots (highest risk)\n")
	b.WriteString("| Rank | Package | Units | Score | Risk Factor |\n")
	b.WriteString("|-----:|---------|------:|------:|------------:|\n")
	for i, h := range snap.Hotspots {
		risk := float64(h.Units) * (1.0 - h.AvgScore)
		fmt.Fprintf(b, "| %d | %s | %d | %.1f%% | %.2f |\n",
			i+1, h.Path, h.Units, h.AvgScore*100, risk)
	}
	b.WriteString("\n")
}

func formatCouplingPairs(b *strings.Builder, snap *ArchSnapshot) {
	if len(snap.CouplingPairs) == 0 {
		return
	}
	b.WriteString("## Coupling Pairs\n")
	b.WriteString("| Package A | Package B | Edges |\n")
	b.WriteString("|-----------|-----------|------:|\n")
	limit := 5
	if len(snap.CouplingPairs) < limit {
		limit = len(snap.CouplingPairs)
	}
	for _, cp := range snap.CouplingPairs[:limit] {
		fmt.Fprintf(b, "| %s | %s | %d |\n", cp.PkgA, cp.PkgB, cp.EdgeCount)
	}
	b.WriteString("\n")
}

func formatTopObservations(b *strings.Builder, snap *ArchSnapshot) {
	if len(snap.Metrics.TopObservations) == 0 {
		return
	}
	b.WriteString("## Top Observation Types\n")
	type obsCount struct {
		name  string
		count int
	}
	var obs []obsCount
	for k, v := range snap.Metrics.TopObservations {
		obs = append(obs, obsCount{k, v})
	}
	sort.Slice(obs, func(i, j int) bool {
		if obs[i].count != obs[j].count {
			return obs[i].count > obs[j].count
		}
		return obs[i].name < obs[j].name
	})
	limit := 10
	if len(obs) < limit {
		limit = len(obs)
	}
	for _, o := range obs[:limit] {
		fmt.Fprintf(b, "- %s: %d\n", o.name, o.count)
	}
	b.WriteString("\n")
}

func formatStructuralMetrics(b *strings.Builder, snap *ArchSnapshot) {
	s := snap.Metrics.Structural
	b.WriteString("## Structural Metrics (aggregated from all units)\n")
	b.WriteString("| Metric | Total | Description |\n")
	b.WriteString("|--------|------:|-------------|\n")
	fmt.Fprintf(b, "| panic_calls | %d | panic() calls in production code |\n", s.PanicCalls)
	fmt.Fprintf(b, "| os_exit_calls | %d | os.Exit() calls (1 in main is normal) |\n", s.OsExitCalls)
	fmt.Fprintf(b, "| global_mutable_count | %d | Package-level mutable var declarations |\n", s.GlobalMutableCount)
	fmt.Fprintf(b, "| defer_in_loop | %d | defer statements inside for/range loops |\n", s.DeferInLoop)
	fmt.Fprintf(b, "| init_func_count | %d | Files containing init() functions |\n", s.InitFuncCount)
	fmt.Fprintf(b, "| context_not_first | %d | Functions with context.Context not as first param |\n", s.ContextNotFirst)
	fmt.Fprintf(b, "| errors_ignored | %d | Error returns assigned to blank identifier |\n", s.ErrorsIgnored)
	fmt.Fprintf(b, "| naked_returns | %d | Bare return statements in named-return functions |\n", s.NakedReturns)
	fmt.Fprintf(b, "| recursive_calls | %d | Direct recursive function calls |\n", s.RecursiveCalls)
	fmt.Fprintf(b, "| max_nesting_depth | %d | Deepest loop nesting across all units |\n", s.MaxNestingDepth)
	fmt.Fprintf(b, "| nested_loop_pairs | %d | Nested loop pairs (O(n²) risk) |\n", s.NestedLoopPairs)
	fmt.Fprintf(b, "| quadratic_patterns | %d | Detected quadratic algorithm patterns |\n", s.QuadraticPatterns)
	fmt.Fprintf(b, "| total_func_lines | %d | Sum of function body lines |\n", s.TotalFuncLines)
	fmt.Fprintf(b, "| total_params | %d | Sum of function parameter counts |\n", s.TotalParams)
	fmt.Fprintf(b, "| total_returns | %d | Sum of function return value counts |\n", s.TotalReturns)
	fmt.Fprintf(b, "| total_methods | %d | Sum of type method counts |\n", s.TotalMethods)
	b.WriteString("\n")
}

func formatCoverageMetrics(b *strings.Builder, snap *ArchSnapshot) {
	c := snap.Metrics.Coverage
	if c.UnitsWithCoverage == 0 && c.UnitsWithoutCoverage == 0 {
		return
	}
	b.WriteString("## Coverage Metrics (aggregated from all units)\n")
	fmt.Fprintf(b, "- Units with coverage data: %d\n", c.UnitsWithCoverage)
	fmt.Fprintf(b, "- Units without coverage data: %d\n", c.UnitsWithoutCoverage)
	if c.UnitsWithCoverage > 0 {
		fmt.Fprintf(b, "- Average coverage: %.1f%%\n", c.AvgCoverage*100)
		fmt.Fprintf(b, "- Min coverage: %.1f%%\n", c.MinCoverage*100)
		fmt.Fprintf(b, "- Max coverage: %.1f%%\n", c.MaxCoverage*100)
	}
	b.WriteString("\n")
}

func formatCodeMetrics(b *strings.Builder, snap *ArchSnapshot) {
	cm := snap.Metrics.CodeMetrics
	if cm.TotalCodeLines == 0 {
		return
	}
	b.WriteString("## Code Metrics (aggregated from all units)\n")
	b.WriteString("| Metric | Value | Description |\n")
	b.WriteString("|--------|------:|-------------|\n")
	fmt.Fprintf(b, "| total_code_lines | %d | Lines of code (excluding blanks/comments) |\n", cm.TotalCodeLines)
	fmt.Fprintf(b, "| total_comment_lines | %d | Lines of comments |\n", cm.TotalCommentLines)
	fmt.Fprintf(b, "| total_complexity | %d | Sum of cyclomatic complexity |\n", cm.TotalComplexity)
	fmt.Fprintf(b, "| max_complexity | %d | Highest single-unit complexity |\n", cm.MaxComplexity)
	fmt.Fprintf(b, "| avg_complexity | %.1f | Average complexity per unit |\n", cm.AvgComplexity)
	fmt.Fprintf(b, "| total_todos | %d | TODO/FIXME markers in code |\n", cm.TotalTodos)
	b.WriteString("\n")
}

func formatQualitativeContext(b *strings.Builder, pc *ProjectContext, maxChars int) {
	b.WriteString("---\n\n# Qualitative Context\n\n")

	if len(pc.LowestUnits) > 0 {
		b.WriteString("## Lowest Scoring Units\n")
		b.WriteString("| Unit | Score | Grade | Observations |\n")
		b.WriteString("|------|------:|:-----:|-------------|\n")
		for _, u := range pc.LowestUnits {
			obs := "-"
			if len(u.Observations) > 0 {
				obs = strings.Join(u.Observations, "; ")
				if len(obs) > 100 {
					obs = obs[:100] + "..."
				}
			}
			fmt.Fprintf(b, "| %s | %.1f%% | %s | %s |\n",
				u.ID, u.Score*100, u.Grade, obs)
		}
		b.WriteString("\n")
	}

	if len(pc.Documentation) > 0 {
		b.WriteString("## Documentation\n")
		var docNames []string
		for k := range pc.Documentation {
			docNames = append(docNames, k)
		}
		sort.Strings(docNames)
		for _, name := range docNames {
			content := pc.Documentation[name]
			if b.Len()+len(content) > maxChars-500 {
				remaining := maxChars - b.Len() - 500
				if remaining > 200 {
					content = content[:remaining] + "\n... (truncated)"
				} else {
					content = "(truncated)"
				}
			}
			fmt.Fprintf(b, "### %s\n%s\n\n", name, content)
		}
	}

	if len(pc.PolicyRules) > 0 {
		b.WriteString("## Policy Packs\n")
		for _, r := range pc.PolicyRules {
			fmt.Fprintf(b, "- %s\n", r)
		}
		b.WriteString("\n")
	}

	if pc.FileTree != "" {
		b.WriteString("## File Tree\n```\n")
		tree := pc.FileTree
		if b.Len()+len(tree) > maxChars-100 {
			remaining := maxChars - b.Len() - 100
			if remaining > 100 {
				tree = tree[:remaining] + "\n... (truncated)"
			}
		}
		b.WriteString(tree)
		b.WriteString("\n```\n")
	}
}

// buildFileTree generates a simple tree of the directory.
func buildFileTree(root string, maxDepth int) string {
	var b strings.Builder
	buildTreeRecursive(&b, root, "", 0, maxDepth)
	return b.String()
}

func buildTreeRecursive(b *strings.Builder, dir, prefix string, depth, maxDepth int) {
	if depth >= maxDepth {
		return
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	// Filter out hidden dirs and common noise
	var filtered []os.DirEntry
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, ".") || name == "node_modules" || name == "vendor" || name == "build" {
			continue
		}
		filtered = append(filtered, e)
	}

	for i, e := range filtered {
		isLast := i == len(filtered)-1
		connector := "├── "
		if isLast {
			connector = "└── "
		}
		fmt.Fprintf(b, "%s%s%s\n", prefix, connector, e.Name())

		if e.IsDir() {
			nextPrefix := prefix + "│   "
			if isLast {
				nextPrefix = prefix + "    "
			}
			buildTreeRecursive(b, filepath.Join(dir, e.Name()), nextPrefix, depth+1, maxDepth)
		}
	}
}
