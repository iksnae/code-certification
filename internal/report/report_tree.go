package report

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// PackageSummary holds roll-up stats for a single package directory.
type PackageSummary struct {
	Path     string  `json:"path"`
	Units    int     `json:"units"`
	Grade    string  `json:"grade"`
	AvgScore float64 `json:"avg_score"`
	PassRate float64 `json:"pass_rate"`
}

// BuildPackageSummaries computes per-package stats from a FullReport.
func BuildPackageSummaries(r FullReport) []PackageSummary {
	type accum struct {
		scores  []float64
		passing int
	}
	pkgs := make(map[string]*accum)
	var dirs []string

	for _, u := range r.Units {
		dir := dirOf(u.Path)
		a, ok := pkgs[dir]
		if !ok {
			a = &accum{}
			pkgs[dir] = a
			dirs = append(dirs, dir)
		}
		a.scores = append(a.scores, u.Score)
		s := statusFromString(u.Status)
		if s.IsPassing() {
			a.passing++
		}
	}

	sort.Strings(dirs)

	summaries := make([]PackageSummary, 0, len(dirs))
	for _, dir := range dirs {
		a := pkgs[dir]
		var total float64
		for _, s := range a.scores {
			total += s
		}
		avg := total / float64(len(a.scores))
		summaries = append(summaries, PackageSummary{
			Path:     dir,
			Units:    len(a.scores),
			Grade:    domain.GradeFromScore(avg).String(),
			AvgScore: avg,
			PassRate: float64(a.passing) / float64(len(a.scores)),
		})
	}
	return summaries
}

// GenerateReportTree writes a hierarchical markdown report tree.
//
// Structure:
//
//	outDir/
//	  index.md                          top-level index
//	  <pkg>/
//	    index.md                        package roll-up
//	    <filename>/
//	      <symbol>.md                   unit certificate (if symbol exists)
//	    <filename>.md                   unit certificate (file-level, no symbol)
//
// Returns the total number of files written.
func GenerateReportTree(r FullReport, outDir string) (int, error) {
	// Clean output directory to remove stale certs
	if err := cleanDir(outDir); err != nil {
		return 0, fmt.Errorf("cleaning output dir: %w", err)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return 0, fmt.Errorf("creating output dir: %w", err)
	}

	// Group units by package
	pkgUnits := make(map[string][]UnitReport)
	var pkgOrder []string
	for _, u := range r.Units {
		dir := dirOf(u.Path)
		if _, ok := pkgUnits[dir]; !ok {
			pkgOrder = append(pkgOrder, dir)
		}
		pkgUnits[dir] = append(pkgUnits[dir], u)
	}
	sort.Strings(pkgOrder)

	summaries := BuildPackageSummaries(r)
	count := 0

	// Write top-level index
	topIndex := formatReportTreeIndex(summaries, r)
	if err := os.WriteFile(filepath.Join(outDir, "index.md"), []byte(topIndex), 0o644); err != nil {
		return 0, fmt.Errorf("writing top-level index: %w", err)
	}
	count++

	// Write package indexes and unit certs
	for _, pkg := range pkgOrder {
		units := pkgUnits[pkg]
		pkgDir := filepath.Join(outDir, pkg)
		if err := os.MkdirAll(pkgDir, 0o755); err != nil {
			return 0, fmt.Errorf("creating package dir %s: %w", pkg, err)
		}

		// Relative path from package index back to top-level index
		depth := strings.Count(pkg, "/") + 1
		relRoot := strings.Repeat("../", depth)

		pkgIndex := formatPackageIndexMarkdown(pkg, units, r, relRoot)
		if err := os.WriteFile(filepath.Join(pkgDir, "index.md"), []byte(pkgIndex), 0o644); err != nil {
			return 0, fmt.Errorf("writing package index %s: %w", pkg, err)
		}
		count++

		// Write unit certs
		for _, u := range units {
			unitPath, unitContent := writeUnitCert(u, r, pkg)
			fullPath := filepath.Join(pkgDir, unitPath)
			if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
				return 0, fmt.Errorf("creating unit dir: %w", err)
			}
			if err := os.WriteFile(fullPath, []byte(unitContent), 0o644); err != nil {
				return 0, fmt.Errorf("writing unit cert %s: %w", u.UnitID, err)
			}
			count++
		}
	}

	return count, nil
}

// writeUnitCert returns the relative file path within the package dir and the content.
func writeUnitCert(u UnitReport, r FullReport, pkg string) (relPath string, content string) {
	filename := shortFile(u.Path)
	symbol := u.Symbol

	if symbol != "" {
		// <filename>/<symbol>.md
		relPath = filepath.Join(filename, symbol+".md")
		backLink := "../index.md"
		content = formatUnitMarkdownWithNav(u, r, pkg, backLink)
	} else {
		// <filename>.md (file-level unit)
		relPath = filename + ".md"
		backLink := "index.md"
		content = formatUnitMarkdownWithNav(u, r, pkg, backLink)
	}
	return
}

// formatUnitMarkdownWithNav renders a unit certificate with navigation back to the package index.
func formatUnitMarkdownWithNav(u UnitReport, r FullReport, pkg, backLink string) string {
	var b strings.Builder

	name := u.Symbol
	if name == "" {
		name = shortFile(u.Path)
	}

	emoji := gradeEmoji(u.Grade)
	fmt.Fprintf(&b, "# %s `%s`\n\n", emoji, name)
	fmt.Fprintf(&b, "[← %s](%s)\n\n", pkg, backLink)
	b.WriteString("---\n\n")

	// Identity
	b.WriteString("## Identity\n\n")
	b.WriteString("| Field | Value |\n")
	b.WriteString("|-------|-------|\n")
	fmt.Fprintf(&b, "| **Unit ID** | `%s` |\n", u.UnitID)
	fmt.Fprintf(&b, "| **Type** | %s |\n", u.UnitType)
	fmt.Fprintf(&b, "| **Path** | `%s` |\n", u.Path)
	fmt.Fprintf(&b, "| **Language** | %s |\n", u.Language)
	if u.Symbol != "" {
		fmt.Fprintf(&b, "| **Symbol** | `%s` |\n", u.Symbol)
	}
	b.WriteString("\n")

	// Certification
	b.WriteString("## Certification\n\n")
	b.WriteString("| Field | Value |\n")
	b.WriteString("|-------|-------|\n")
	fmt.Fprintf(&b, "| **Grade** | %s **%s** |\n", emoji, u.Grade)
	fmt.Fprintf(&b, "| **Score** | %.1f%% |\n", u.Score*100)
	fmt.Fprintf(&b, "| **Status** | %s |\n", u.Status)
	fmt.Fprintf(&b, "| **Confidence** | %.0f%% |\n", u.Confidence*100)
	fmt.Fprintf(&b, "| **Certified** | %s |\n", formatDate(u.CertifiedAt))
	fmt.Fprintf(&b, "| **Expires** | %s |\n", formatDate(u.ExpiresAt))
	fmt.Fprintf(&b, "| **Source** | `%s` |\n", u.Source)
	b.WriteString("\n")

	// Dimensions
	if len(u.Dimensions) > 0 {
		b.WriteString("## Dimension Scores\n\n")
		b.WriteString("| Dimension | Score | Bar |\n")
		b.WriteString("|-----------|------:|-----|\n")
		for _, d := range sortedKeys(u.Dimensions) {
			score := u.Dimensions[d]
			bar := scoreBar(score, 20)
			fmt.Fprintf(&b, "| %s | %.1f%% | %s |\n", d, score*100, bar)
		}
		b.WriteString("\n")
	}

	// AI Observations
	aiObs, suggestions, otherObs := splitObservations(u.Observations)
	if len(aiObs) > 0 || len(suggestions) > 0 {
		b.WriteString("## 🤖 AI Assessment\n\n")
		for _, obs := range aiObs {
			text := strings.TrimPrefix(obs, "🤖 ")
			fmt.Fprintf(&b, "%s\n\n", text)
		}
		if len(suggestions) > 0 {
			b.WriteString("### Suggestions\n\n")
			for _, s := range suggestions {
				text := strings.TrimPrefix(s, "💡 ")
				fmt.Fprintf(&b, "- %s\n", text)
			}
			b.WriteString("\n")
		}
	}

	// Other observations
	if len(otherObs) > 0 {
		b.WriteString("## Observations\n\n")
		for _, obs := range otherObs {
			fmt.Fprintf(&b, "- %s\n", obs)
		}
		b.WriteString("\n")
	}

	// Actions
	if len(u.Actions) > 0 {
		b.WriteString("## Required Actions\n\n")
		for _, a := range u.Actions {
			fmt.Fprintf(&b, "- [ ] %s\n", a)
		}
		b.WriteString("\n")
	}

	// Footer
	b.WriteString("---\n\n")
	fmt.Fprintf(&b, "*Repository: `%s` · Commit: `%s` · Generated: %s*\n",
		r.Repository, r.CommitSHA, r.GeneratedAt)
	fmt.Fprintf(&b, "*Generated by [Certify](https://github.com/iksnae/code-certification)*\n")

	return b.String()
}

// formatReportTreeIndex generates the top-level reports/index.md.
func formatReportTreeIndex(packages []PackageSummary, r FullReport) string {
	var b strings.Builder

	emoji := gradeEmoji(r.Card.OverallGrade)
	fmt.Fprintf(&b, "# %s Certification Report\n\n", emoji)
	fmt.Fprintf(&b, "[← Report Card](../REPORT_CARD.md)\n\n")

	if r.Repository != "" {
		fmt.Fprintf(&b, "**Repository:** `%s`  \n", r.Repository)
	}
	if r.CommitSHA != "" {
		fmt.Fprintf(&b, "**Commit:** `%s`  \n", r.CommitSHA)
	}
	fmt.Fprintf(&b, "**Overall:** %s %s (%.1f%%)  \n", emoji, r.Card.OverallGrade, r.Card.OverallScore*100)
	fmt.Fprintf(&b, "**Units:** %d · **Passing:** %d · **Failing:** %d\n\n", r.Card.TotalUnits, r.Card.Passing, r.Card.Failing)

	b.WriteString("## Packages\n\n")
	b.WriteString("| Package | Units | Grade | Score | Pass Rate |\n")
	b.WriteString("|---------|------:|:-----:|------:|----------:|\n")
	for _, p := range packages {
		fmt.Fprintf(&b, "| [%s](%s/index.md) | %d | %s %s | %.1f%% | %.0f%% |\n",
			p.Path, p.Path, p.Units, gradeEmoji(p.Grade), p.Grade, p.AvgScore*100, p.PassRate*100)
	}
	b.WriteString("\n---\n\n")
	b.WriteString("*Generated by [Certify](https://github.com/iksnae/code-certification)*\n")

	return b.String()
}

// formatPackageIndexMarkdown generates a package-level index.md.
func formatPackageIndexMarkdown(pkg string, units []UnitReport, r FullReport, relRoot string) string {
	var b strings.Builder

	// Compute package stats
	var totalScore float64
	passing := 0
	for _, u := range units {
		totalScore += u.Score
		s := statusFromString(u.Status)
		if s.IsPassing() {
			passing++
		}
	}
	avgScore := totalScore / float64(len(units))
	grade := domain.GradeFromScore(avgScore).String()
	emoji := gradeEmoji(grade)

	fmt.Fprintf(&b, "# %s `%s`\n\n", emoji, pkg)
	fmt.Fprintf(&b, "[← All Packages](%sindex.md) · [← Report Card](%s../REPORT_CARD.md)\n\n", relRoot, relRoot)

	fmt.Fprintf(&b, "**Grade:** %s %s (%.1f%%)  \n", emoji, grade, avgScore*100)
	fmt.Fprintf(&b, "**Units:** %d · **Passing:** %d / %d\n\n", len(units), passing, len(units))

	// Sort units by score ascending (worst first)
	sorted := make([]UnitReport, len(units))
	copy(sorted, units)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Score < sorted[j].Score
	})

	b.WriteString("## Units\n\n")
	b.WriteString("| Unit | Type | Grade | Score | Status | Expires |\n")
	b.WriteString("|------|------|:-----:|------:|--------|--------:|\n")

	for _, u := range sorted {
		name := u.Symbol
		link := unitTreePath(u)
		if name == "" {
			name = shortFile(u.Path)
		}
		fmt.Fprintf(&b, "| [%s](%s) | %s | %s %s | %.1f%% | %s | %s |\n",
			name, link, u.UnitType, gradeEmoji(u.Grade), u.Grade, u.Score*100, u.Status, formatDate(u.ExpiresAt))
	}

	b.WriteString("\n---\n\n")
	b.WriteString("*Generated by [Certify](https://github.com/iksnae/code-certification)*\n")

	return b.String()
}

// unitTreePath returns the relative path from the package index to the unit cert file.
func unitTreePath(u UnitReport) string {
	filename := shortFile(u.Path)
	if u.Symbol != "" {
		return filename + "/" + u.Symbol + ".md"
	}
	return filename + ".md"
}

// cleanDir removes all contents of a directory (but not the directory itself).
func cleanDir(dir string) error {
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	for _, e := range entries {
		if err := os.RemoveAll(filepath.Join(dir, e.Name())); err != nil {
			return err
		}
	}
	return nil
}
