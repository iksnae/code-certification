package agent

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// SnapshotSchemaVersion tracks the data contract between BuildSnapshot and FormatForLLM.
// Increment when adding/removing/renaming fields in SnapshotMetrics or its sub-structs.
const SnapshotSchemaVersion = 2 // v1: PR #20 (structural only), v2: full alignment

// ArchSnapshot captures the current architecture as deterministic data.
// Computed from certification records — no LLM involved.
type ArchSnapshot struct {
	SchemaVersion   int               `json:"schema_version"`
	Packages        []PackageNode     // every package with metrics, sorted by path
	DependencyEdges []DepEdge         // package → package imports
	Layers          map[string]string // package → layer (cmd, internal, domain, external)
	Hotspots        []PackageNode     // top 10 by risk score (units × (1 - avg_score))
	CouplingPairs   []CouplingPair    // most cross-referenced package pairs
	Metrics         SnapshotMetrics   // aggregate metrics
}

// PackageNode holds metrics for a single package directory.
type PackageNode struct {
	Path         string   `json:"path"`
	Units        int      `json:"units"`
	AvgScore     float64  `json:"avg_score"`
	Grade        string   `json:"grade"`
	Observations int      `json:"observations"`
	TopIssues    []string `json:"top_issues,omitempty"`
	Imports      []string `json:"imports,omitempty"`
	ImportedBy   []string `json:"imported_by,omitempty"`
	ExportedAPI  []string `json:"exported_api,omitempty"`
}

// DepEdge represents an import dependency between two packages.
type DepEdge struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Weight int    `json:"weight"` // number of import sites
}

// CouplingPair represents two packages with high bidirectional coupling.
type CouplingPair struct {
	PkgA      string `json:"pkg_a"`
	PkgB      string `json:"pkg_b"`
	EdgeCount int    `json:"edge_count"` // total edges both directions
}

// SnapshotMetrics holds aggregate metrics across all packages.
type SnapshotMetrics struct {
	TotalUnits        int                   `json:"total_units"`
	TotalPackages     int                   `json:"total_packages"`
	AvgScore          float64               `json:"avg_score"`
	GradeDistribution map[string]int        `json:"grade_distribution"`
	TopObservations   map[string]int        `json:"top_observations"`
	PolicyViolations  map[string]int        `json:"policy_violations"`
	Structural        StructuralAggregates  `json:"structural"`
	Coverage          CoverageAggregates    `json:"coverage"`
	CodeMetrics       CodeMetricsAggregates `json:"code_metrics"`
}

// StructuralAggregates holds summed structural metrics across all units.
// These are computed from Evidence.Metrics in certification records, providing
// the LLM with real data instead of requiring it to guess.
type StructuralAggregates struct {
	PanicCalls         int `json:"panic_calls"`
	OsExitCalls        int `json:"os_exit_calls"`
	GlobalMutableCount int `json:"global_mutable_count"`
	DeferInLoop        int `json:"defer_in_loop"`
	InitFuncCount      int `json:"init_func_count"`
	ContextNotFirst    int `json:"context_not_first"`
	ErrorsIgnored      int `json:"errors_ignored"`
	NakedReturns       int `json:"naked_returns"`
	RecursiveCalls     int `json:"recursive_calls"`
	MaxNestingDepth    int `json:"max_nesting_depth"`
	NestedLoopPairs    int `json:"nested_loop_pairs"`
	QuadraticPatterns  int `json:"quadratic_patterns"`
	TotalFuncLines     int `json:"total_func_lines"`
	TotalParams        int `json:"total_params"`
	TotalReturns       int `json:"total_returns"`
	TotalMethods       int `json:"total_methods"`
}

// CoverageAggregates holds aggregated test coverage data across all units.
type CoverageAggregates struct {
	UnitsWithCoverage    int     `json:"units_with_coverage"`
	UnitsWithoutCoverage int     `json:"units_without_coverage"`
	AvgCoverage          float64 `json:"avg_coverage"`
	MinCoverage          float64 `json:"min_coverage"`
	MaxCoverage          float64 `json:"max_coverage"`
}

// CodeMetricsAggregates holds aggregated code metrics across all units.
type CodeMetricsAggregates struct {
	TotalCodeLines    int     `json:"total_code_lines"`
	TotalCommentLines int     `json:"total_comment_lines"`
	TotalComplexity   int     `json:"total_complexity"`
	MaxComplexity     int     `json:"max_complexity"`
	AvgComplexity     float64 `json:"avg_complexity"`
	TotalTodos        int     `json:"total_todos"`
}

// BuildSnapshot computes a deterministic architecture snapshot from certification records.
// The root parameter is used for Go import analysis (can be empty to skip).
func BuildSnapshot(records []domain.CertificationRecord, root string) *ArchSnapshot {
	snap := &ArchSnapshot{
		SchemaVersion: SnapshotSchemaVersion,
		Layers:        make(map[string]string),
		Metrics: SnapshotMetrics{
			GradeDistribution: make(map[string]int),
			TopObservations:   make(map[string]int),
			PolicyViolations:  make(map[string]int),
		},
	}

	if len(records) == 0 {
		return snap
	}

	// Group records by package path
	type pkgAccum struct {
		scores       []float64
		observations []string
		issueCounts  map[string]int
	}
	pkgMap := make(map[string]*pkgAccum)
	var pkgPaths []string

	var totalScore float64
	var coverageValues []float64
	var metricsUnitCount int

	for _, r := range records {
		pkgPath := packagePath(r.UnitPath)

		a, ok := pkgMap[pkgPath]
		if !ok {
			a = &pkgAccum{issueCounts: make(map[string]int)}
			pkgMap[pkgPath] = a
			pkgPaths = append(pkgPaths, pkgPath)
		}
		a.scores = append(a.scores, r.Score)
		a.observations = append(a.observations, r.Observations...)

		for _, obs := range r.Observations {
			issueType := extractIssueType(obs)
			a.issueCounts[issueType]++
			snap.Metrics.TopObservations[issueType]++
		}

		snap.Metrics.GradeDistribution[r.Grade.String()]++
		totalScore += r.Score

		// Aggregate evidence metrics by kind
		for _, ev := range r.Evidence {
			switch ev.Kind {
			case domain.EvidenceKindStructural:
				snap.Metrics.Structural.PanicCalls += int(ev.Metrics["panic_calls"])
				snap.Metrics.Structural.OsExitCalls += int(ev.Metrics["os_exit_calls"])
				snap.Metrics.Structural.GlobalMutableCount += int(ev.Metrics["global_mutable_count"])
				snap.Metrics.Structural.DeferInLoop += int(ev.Metrics["defer_in_loop"])
				snap.Metrics.Structural.ErrorsIgnored += int(ev.Metrics["errors_ignored"])
				snap.Metrics.Structural.NakedReturns += int(ev.Metrics["naked_returns"])
				snap.Metrics.Structural.RecursiveCalls += int(ev.Metrics["recursive_calls"])
				snap.Metrics.Structural.NestedLoopPairs += int(ev.Metrics["nested_loop_pairs"])
				snap.Metrics.Structural.QuadraticPatterns += int(ev.Metrics["quadratic_patterns"])
				snap.Metrics.Structural.TotalFuncLines += int(ev.Metrics["func_lines"])
				snap.Metrics.Structural.TotalParams += int(ev.Metrics["param_count"])
				snap.Metrics.Structural.TotalReturns += int(ev.Metrics["return_count"])
				snap.Metrics.Structural.TotalMethods += int(ev.Metrics["method_count"])
				if ev.Metrics["has_init_func"] > 0 {
					snap.Metrics.Structural.InitFuncCount++
				}
				snap.Metrics.Structural.ContextNotFirst += int(ev.Metrics["context_not_first"])
				if nd := int(ev.Metrics["loop_nesting_depth"]); nd > snap.Metrics.Structural.MaxNestingDepth {
					snap.Metrics.Structural.MaxNestingDepth = nd
				}
			case domain.EvidenceKindMetrics:
				snap.Metrics.CodeMetrics.TotalCodeLines += int(ev.Metrics["code_lines"])
				snap.Metrics.CodeMetrics.TotalCommentLines += int(ev.Metrics["comment_lines"])
				c := int(ev.Metrics["complexity"])
				snap.Metrics.CodeMetrics.TotalComplexity += c
				if c > snap.Metrics.CodeMetrics.MaxComplexity {
					snap.Metrics.CodeMetrics.MaxComplexity = c
				}
				snap.Metrics.CodeMetrics.TotalTodos += int(ev.Metrics["todo_count"])
				metricsUnitCount++
			}
			if ev.Source == "coverage:unit" {
				if cov, ok := ev.Metrics["unit_test_coverage"]; ok {
					coverageValues = append(coverageValues, cov)
				}
			}
		}
	}

	snap.Metrics.TotalUnits = len(records)
	snap.Metrics.AvgScore = totalScore / float64(len(records))

	// Finalize coverage aggregates
	snap.Metrics.Coverage.UnitsWithCoverage = len(coverageValues)
	snap.Metrics.Coverage.UnitsWithoutCoverage = len(records) - len(coverageValues)
	if len(coverageValues) > 0 {
		min, max, sum := coverageValues[0], coverageValues[0], 0.0
		for _, v := range coverageValues {
			sum += v
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		snap.Metrics.Coverage.AvgCoverage = sum / float64(len(coverageValues))
		snap.Metrics.Coverage.MinCoverage = min
		snap.Metrics.Coverage.MaxCoverage = max
	}

	// Finalize code metrics aggregates
	if metricsUnitCount > 0 {
		snap.Metrics.CodeMetrics.AvgComplexity = float64(snap.Metrics.CodeMetrics.TotalComplexity) / float64(metricsUnitCount)
	}

	// Sort package paths for deterministic output
	sort.Strings(pkgPaths)

	// Build package nodes
	for _, path := range pkgPaths {
		a := pkgMap[path]
		var sum float64
		for _, s := range a.scores {
			sum += s
		}
		avg := sum / float64(len(a.scores))

		node := PackageNode{
			Path:         path,
			Units:        len(a.scores),
			AvgScore:     avg,
			Grade:        domain.GradeFromScore(avg).String(),
			Observations: len(a.observations),
			TopIssues:    topNIssues(a.issueCounts, 5),
		}
		snap.Packages = append(snap.Packages, node)
		snap.Layers[path] = classifyLayer(path)
	}

	snap.Metrics.TotalPackages = len(snap.Packages)

	// Compute hotspots: sort by risk = units × (1.0 - avgScore)
	hotspots := make([]PackageNode, len(snap.Packages))
	copy(hotspots, snap.Packages)
	sort.Slice(hotspots, func(i, j int) bool {
		riskI := float64(hotspots[i].Units) * (1.0 - hotspots[i].AvgScore)
		riskJ := float64(hotspots[j].Units) * (1.0 - hotspots[j].AvgScore)
		if riskI != riskJ {
			return riskI > riskJ
		}
		return hotspots[i].Path < hotspots[j].Path
	})
	if len(hotspots) > 10 {
		hotspots = hotspots[:10]
	}
	snap.Hotspots = hotspots

	// Dependency analysis from Go source (if root provided)
	if root != "" {
		snap.analyzeDependencies(root, pkgPaths)
	}

	return snap
}

// analyzeDependencies parses Go imports to build dependency edges.
func (snap *ArchSnapshot) analyzeDependencies(root string, pkgPaths []string) {
	knownPkgs := make(map[string]bool, len(pkgPaths))
	for _, p := range pkgPaths {
		knownPkgs[p] = true
	}

	// Track edges: from → to → weight
	edges := make(map[string]map[string]int)

	// Detect module path from go.mod
	modulePath := detectModulePath(root)

	for _, pkgPath := range pkgPaths {
		absDir := filepath.Join(root, pkgPath)
		imports := parseGoImports(absDir)
		for _, imp := range imports {
			relImp := normalizeImport(imp, modulePath)
			if relImp == "" || !knownPkgs[relImp] || relImp == pkgPath {
				continue
			}
			if edges[pkgPath] == nil {
				edges[pkgPath] = make(map[string]int)
			}
			edges[pkgPath][relImp]++
		}
	}

	// Build DepEdge list and update PackageNode imports/importedBy
	importedBy := make(map[string][]string)

	for from, targets := range edges {
		for to, weight := range targets {
			snap.DependencyEdges = append(snap.DependencyEdges, DepEdge{
				From: from, To: to, Weight: weight,
			})
			importedBy[to] = append(importedBy[to], from)
		}
	}

	// Sort edges for deterministic output
	sort.Slice(snap.DependencyEdges, func(i, j int) bool {
		if snap.DependencyEdges[i].From != snap.DependencyEdges[j].From {
			return snap.DependencyEdges[i].From < snap.DependencyEdges[j].From
		}
		return snap.DependencyEdges[i].To < snap.DependencyEdges[j].To
	})

	// Update package nodes
	for i := range snap.Packages {
		pkg := &snap.Packages[i]
		if targets, ok := edges[pkg.Path]; ok {
			for to := range targets {
				pkg.Imports = append(pkg.Imports, to)
			}
			sort.Strings(pkg.Imports)
		}
		if importers, ok := importedBy[pkg.Path]; ok {
			sort.Strings(importers)
			seen := make(map[string]bool)
			for _, imp := range importers {
				if !seen[imp] {
					pkg.ImportedBy = append(pkg.ImportedBy, imp)
					seen[imp] = true
				}
			}
		}
	}

	snap.buildCouplingPairs(edges)
}

// buildCouplingPairs finds package pairs with the most cross-references.
func (snap *ArchSnapshot) buildCouplingPairs(edges map[string]map[string]int) {
	type pairKey struct{ a, b string }
	pairCounts := make(map[pairKey]int)

	for from, targets := range edges {
		for to, weight := range targets {
			a, b := from, to
			if a > b {
				a, b = b, a
			}
			pairCounts[pairKey{a, b}] += weight
		}
	}

	for pk, count := range pairCounts {
		snap.CouplingPairs = append(snap.CouplingPairs, CouplingPair{
			PkgA: pk.a, PkgB: pk.b, EdgeCount: count,
		})
	}

	sort.Slice(snap.CouplingPairs, func(i, j int) bool {
		if snap.CouplingPairs[i].EdgeCount != snap.CouplingPairs[j].EdgeCount {
			return snap.CouplingPairs[i].EdgeCount > snap.CouplingPairs[j].EdgeCount
		}
		return snap.CouplingPairs[i].PkgA < snap.CouplingPairs[j].PkgA
	})

	if len(snap.CouplingPairs) > 20 {
		snap.CouplingPairs = snap.CouplingPairs[:20]
	}
}

// packagePath extracts the directory from a file path.
func packagePath(filePath string) string {
	dir := filepath.Dir(filePath)
	if dir == "." {
		return "."
	}
	return filepath.ToSlash(dir)
}

// extractIssueType extracts the issue type from an observation string.
func extractIssueType(obs string) string {
	if idx := strings.Index(obs, ":"); idx > 0 {
		return strings.TrimSpace(obs[:idx])
	}
	if idx := strings.Index(obs, " "); idx > 0 {
		return strings.TrimSpace(obs[:idx])
	}
	return obs
}

// topNIssues returns the top N issue types by count.
func topNIssues(counts map[string]int, n int) []string {
	type kv struct {
		key   string
		count int
	}
	var items []kv
	for k, v := range counts {
		items = append(items, kv{k, v})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].count != items[j].count {
			return items[i].count > items[j].count
		}
		return items[i].key < items[j].key
	})
	var result []string
	for i, item := range items {
		if i >= n {
			break
		}
		result = append(result, item.key)
	}
	return result
}

// classifyLayer classifies a package path into an architectural layer.
func classifyLayer(pkgPath string) string {
	parts := strings.Split(pkgPath, "/")
	if len(parts) == 0 {
		return "other"
	}
	switch parts[0] {
	case "cmd":
		return "cmd"
	case "internal":
		if len(parts) >= 2 && parts[1] == "domain" {
			return "domain"
		}
		return "internal"
	case "pkg":
		return "pkg"
	case "vendor":
		return "vendor"
	default:
		return "other"
	}
}

// parseGoImports parses all Go files in a directory and returns import paths.
func parseGoImports(dir string) []string {
	entries, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil || len(entries) == 0 {
		return nil
	}

	var imports []string
	for _, f := range entries {
		data, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		// Only read first 8KB — imports are always near top
		src := string(data)
		if len(src) > 8192 {
			src = src[:8192]
		}
		imports = append(imports, extractImportPaths(src)...)
	}
	return imports
}

// extractImportPaths extracts import paths from Go source.
func extractImportPaths(src string) []string {
	var imports []string
	inImportBlock := false

	for _, line := range strings.Split(src, "\n") {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "import (") {
			inImportBlock = true
			continue
		}
		if inImportBlock && trimmed == ")" {
			inImportBlock = false
			continue
		}
		if inImportBlock {
			if path := extractQuotedImport(trimmed); path != "" {
				imports = append(imports, path)
			}
		}
		if strings.HasPrefix(trimmed, "import \"") || strings.HasPrefix(trimmed, "import\t\"") {
			if path := extractQuotedImport(trimmed); path != "" {
				imports = append(imports, path)
			}
		}
	}
	return imports
}

// extractQuotedImport extracts the quoted import path from a line.
func extractQuotedImport(line string) string {
	start := strings.Index(line, "\"")
	if start < 0 {
		return ""
	}
	end := strings.Index(line[start+1:], "\"")
	if end < 0 {
		return ""
	}
	return line[start+1 : start+1+end]
}

// normalizeImport converts a full Go import path to a relative package path.
func normalizeImport(importPath, modulePath string) string {
	if modulePath != "" && strings.HasPrefix(importPath, modulePath+"/") {
		return importPath[len(modulePath)+1:]
	}
	return ""
}

// detectModulePath reads the module path from go.mod.
func detectModulePath(root string) string {
	data, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(line[7:])
		}
	}
	return ""
}
