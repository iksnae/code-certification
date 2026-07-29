package report

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// PackageSummary holds roll-up stats for a single package directory.
type PackageSummary struct {
	Path  string `json:"path"`
	Units int    `json:"units"`
	// Unsupported counts this package's unassessed units, which are excluded
	// from both sides of PassRate. See Card.
	Unsupported int     `json:"unsupported"`
	Grade       string  `json:"grade"`
	AvgScore    float64 `json:"avg_score"`
	PassRate    float64 `json:"pass_rate"`
}

// Analyzable is the number of units in this package about which a verdict was
// asserted — the denominator of PassRate.
func (p PackageSummary) Analyzable() int { return p.Units - p.Unsupported }

// PassRateKnown reports whether PassRate is a measurement. See Card.PassRateKnown.
func (p PackageSummary) PassRateKnown() bool { return p.Analyzable() > 0 }

// ScoreKnown reports whether AvgScore and Grade are measurements. See
// Card.ScoreKnown.
func (p PackageSummary) ScoreKnown() bool { return p.Analyzable() > 0 }

// packageStats is the one aggregation over a package's units. Three surfaces
// render these numbers — the markdown packages table, the markdown package page
// and the HTML package page — and each used to recompute them independently,
// which is how the unassessed-unit false positive survived at two of them after
// being fixed at the third. Adding a fourth reader must not mean adding a fourth
// chance to get the numerator wrong.
type packageStats struct {
	units       int
	passing     int
	unsupported int
	avgScore    float64
}

// statsForUnits aggregates one package's units. passing spans only the units
// about which a verdict was asserted: an unsupported unit carries status
// "exempt", whose IsPassing() is true, so counting it here is the report card's
// false positive one level down.
//
// avgScore spans the same set, for the same reason. An unassessed unit's score is
// the placeholder the pipeline assigns when it declines to score, so summing it
// pulls the mean toward zero as though a failure had been measured there. This is
// the denominator Card.OverallScore uses, and the two must stay identical: the
// package grade and the overall grade are printed in one artifact one table
// apart, and disagreeing is the contradiction this branch exists to remove.
func statsForUnits(units []UnitReport) packageStats {
	s := packageStats{units: len(units)}
	var totalScore float64
	for _, u := range units {
		if u.Unsupported {
			s.unsupported++
			continue
		}
		totalScore += u.Score
		if statusFromString(u.Status).IsPassing() {
			s.passing++
		}
	}
	if s.measured() {
		s.avgScore = totalScore / float64(s.analyzable())
	}
	return s
}

// analyzable is the number of units about which a verdict was asserted — the
// denominator of the pass rate.
func (s packageStats) analyzable() int { return s.units - s.unsupported }

// measured reports whether any unit in this package was assessed.
//
// It gates the pass rate AND the grade, because both are summaries of assessed
// code and both are undefined over none of it: 0/0 is neither 0% nor 100%, and
// a mean over nothing but placeholder zeroes is not an F. One predicate rather
// than two, so a surface cannot gate the rate and forget the grade — which is
// exactly how "Pass Rate: n/a" came to be printed beneath "🔴 F (0.0%)".
func (s packageStats) measured() bool { return s.analyzable() > 0 }

func (s packageStats) passRate() float64 {
	if !s.measured() {
		return 0
	}
	return float64(s.passing) / float64(s.analyzable())
}

func (s packageStats) grade() string {
	if !s.measured() {
		return domain.GradeNA.String()
	}
	return domain.GradeFromScore(s.avgScore).String()
}

// BuildPackageSummaries computes per-package stats from a FullReport.
func BuildPackageSummaries(r FullReport) []PackageSummary {
	pkgUnits := make(map[string][]UnitReport)
	var dirs []string

	for _, u := range r.Units {
		dir := dirOf(u.Path)
		if _, ok := pkgUnits[dir]; !ok {
			dirs = append(dirs, dir)
		}
		pkgUnits[dir] = append(pkgUnits[dir], u)
	}

	sort.Strings(dirs)

	summaries := make([]PackageSummary, 0, len(dirs))
	for _, dir := range dirs {
		s := statsForUnits(pkgUnits[dir])
		summaries = append(summaries, PackageSummary{
			Path:        dir,
			Units:       s.units,
			Unsupported: s.unsupported,
			Grade:       s.grade(),
			AvgScore:    s.avgScore,
			PassRate:    s.passRate(),
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

	writeUnitIdentity(&b, u)
	writeUnitCertification(&b, u, emoji)
	writeUnitDimensions(&b, u)
	writeUnitEvidence(&b, u)
	writeUnitObservations(&b, u)

	if len(u.Actions) > 0 {
		b.WriteString("## Required Actions\n\n")
		for _, a := range u.Actions {
			fmt.Fprintf(&b, "- [ ] %s\n", a)
		}
		b.WriteString("\n")
	}

	b.WriteString("---\n\n")
	fmt.Fprintf(&b, "*Repository: `%s` · Commit: `%s` · Generated: %s*\n",
		r.Repository, r.CommitSHA, r.GeneratedAt)
	fmt.Fprintf(&b, "*Generated by [Certify](https://github.com/iksnae/code-certification)*\n")

	return b.String()
}

func writeUnitIdentity(b *strings.Builder, u UnitReport) {
	b.WriteString("## Identity\n\n")
	b.WriteString("| Field | Value |\n")
	b.WriteString("|-------|-------|\n")
	fmt.Fprintf(b, "| **Unit ID** | `%s` |\n", u.UnitID)
	fmt.Fprintf(b, "| **Type** | %s |\n", u.UnitType)
	fmt.Fprintf(b, "| **Path** | `%s` |\n", u.Path)
	fmt.Fprintf(b, "| **Language** | %s |\n", u.Language)
	if u.Symbol != "" {
		fmt.Fprintf(b, "| **Symbol** | `%s` |\n", u.Symbol)
	}
	b.WriteString("\n")
}

func writeUnitCertification(b *strings.Builder, u UnitReport, emoji string) {
	b.WriteString("## Certification\n\n")
	b.WriteString("| Field | Value |\n")
	b.WriteString("|-------|-------|\n")
	fmt.Fprintf(b, "| **Grade** | %s **%s** |\n", emoji, u.Grade)
	fmt.Fprintf(b, "| **Score** | %s |\n", FormatRate(u.ScoreKnown(), u.Score, 1))
	fmt.Fprintf(b, "| **Status** | %s |\n", u.Status)
	fmt.Fprintf(b, "| **Confidence** | %.0f%% |\n", u.Confidence*100)
	fmt.Fprintf(b, "| **Certified** | %s |\n", formatDate(u.CertifiedAt))
	fmt.Fprintf(b, "| **Expires** | %s |\n", formatDate(u.ExpiresAt))
	fmt.Fprintf(b, "| **Source** | `%s` |\n", u.Source)
	b.WriteString("\n")
}

func writeUnitDimensions(b *strings.Builder, u UnitReport) {
	if len(u.Dimensions) == 0 {
		return
	}
	b.WriteString("## Dimension Scores\n\n")
	b.WriteString("| Dimension | Score | Bar |\n")
	b.WriteString("|-----------|------:|-----|\n")
	for _, d := range sortedKeys(u.Dimensions) {
		score := u.Dimensions[d]
		bar := scoreBar(score, 20)
		fmt.Fprintf(b, "| %s | %.1f%% | %s |\n", d, score*100, bar)
	}
	b.WriteString("\n")
}

func writeUnitEvidence(b *strings.Builder, u UnitReport) {
	if len(u.Evidence) == 0 {
		return
	}
	b.WriteString("## Evidence\n\n")
	for _, ev := range u.Evidence {
		pass := "✅"
		if !ev.Passed {
			pass = "❌"
		}
		fmt.Fprintf(b, "### %s %s (`%s`)\n\n", pass, ev.Kind, ev.Source)
		if ev.Summary != "" {
			fmt.Fprintf(b, "%s\n\n", ev.Summary)
		}
		if len(ev.Metrics) > 0 {
			b.WriteString("| Metric | Value |\n")
			b.WriteString("|--------|------:|\n")
			for _, k := range sortedKeys(ev.Metrics) {
				v := ev.Metrics[k]
				if v == float64(int(v)) {
					fmt.Fprintf(b, "| `%s` | %d |\n", k, int(v))
				} else {
					fmt.Fprintf(b, "| `%s` | %.2f |\n", k, v)
				}
			}
			b.WriteString("\n")
		}
	}
}

func writeUnitObservations(b *strings.Builder, u UnitReport) {
	aiObs, suggestions, otherObs := splitObservations(u.Observations)
	if len(aiObs) > 0 || len(suggestions) > 0 {
		b.WriteString("## 🤖 AI Assessment\n\n")
		for _, obs := range aiObs {
			text := strings.TrimPrefix(obs, "🤖 ")
			fmt.Fprintf(b, "%s\n\n", text)
		}
		if len(suggestions) > 0 {
			b.WriteString("### Suggestions\n\n")
			for _, s := range suggestions {
				text := strings.TrimPrefix(s, "💡 ")
				fmt.Fprintf(b, "- %s\n", text)
			}
			b.WriteString("\n")
		}
	}
	if len(otherObs) > 0 {
		b.WriteString("## Observations\n\n")
		for _, obs := range otherObs {
			fmt.Fprintf(b, "- %s\n", obs)
		}
		b.WriteString("\n")
	}
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
	fmt.Fprintf(&b, "**Overall:** %s %s (%s)  \n", emoji, r.Card.OverallGrade,
		FormatRate(r.Card.ScoreKnown(), r.Card.OverallScore, 1))
	fmt.Fprintf(&b, "**Units:** %d · **Passing:** %d · **Failing:** %d", r.Card.TotalUnits, r.Card.Passing, r.Card.Failing)
	if r.Card.UnsupportedCount > 0 {
		fmt.Fprintf(&b, " · **Not Assessed:** %d", r.Card.UnsupportedCount)
	}
	b.WriteString("\n\n")

	// A row reading "3 units · 100%" is read as three passing units, when the
	// rate covers only the analyzable ones. The Not Assessed column carries the
	// gap, and appears only when there is a gap — an all-assessed repo keeps the
	// table it had, the same rule the summary line above already follows.
	anyUnsupported := false
	for _, p := range packages {
		if p.Unsupported > 0 {
			anyUnsupported = true
			break
		}
	}

	b.WriteString("## Packages\n\n")
	if anyUnsupported {
		b.WriteString("| Package | Units | Not Assessed | Grade | Score | Pass Rate |\n")
		b.WriteString("|---------|------:|-------------:|:-----:|------:|----------:|\n")
	} else {
		b.WriteString("| Package | Units | Grade | Score | Pass Rate |\n")
		b.WriteString("|---------|------:|:-----:|------:|----------:|\n")
	}
	for _, p := range packages {
		cells := []string{
			fmt.Sprintf("[%s](%s/index.md)", p.Path, p.Path),
			strconv.Itoa(p.Units),
		}
		if anyUnsupported {
			cells = append(cells, strconv.Itoa(p.Unsupported))
		}
		cells = append(cells,
			fmt.Sprintf("%s %s", gradeEmoji(p.Grade), p.Grade),
			FormatRate(p.ScoreKnown(), p.AvgScore, 1),
			FormatRate(p.PassRateKnown(), p.PassRate, 0),
		)
		fmt.Fprintf(&b, "| %s |\n", strings.Join(cells, " | "))
	}
	b.WriteString("\n---\n\n")
	b.WriteString("*Generated by [Certify](https://github.com/iksnae/code-certification)*\n")

	return b.String()
}

// formatPackageIndexMarkdown generates a package-level index.md.
func formatPackageIndexMarkdown(pkg string, units []UnitReport, r FullReport, relRoot string) string {
	var b strings.Builder

	s := statsForUnits(units)
	grade := s.grade()
	emoji := gradeEmoji(grade)

	fmt.Fprintf(&b, "# %s `%s`\n\n", emoji, pkg)
	fmt.Fprintf(&b, "[← All Packages](%sindex.md) · [← Report Card](%s../REPORT_CARD.md)\n\n", relRoot, relRoot)

	fmt.Fprintf(&b, "**Grade:** %s %s (%s)  \n", emoji, grade, FormatRate(s.measured(), s.avgScore, 1))
	// Passing is stated over analyzable units, and only when some unit was
	// analyzed — "0 / 0" is not a modest claim, it is a claim about units the
	// engine never opened. This page sits one click from the packages table; the
	// two disagreeing put contradicting certification claims in one artifact.
	fmt.Fprintf(&b, "**Units:** %d", s.units)
	if s.measured() {
		fmt.Fprintf(&b, " · **Passing:** %d / %d", s.passing, s.analyzable())
	}
	if s.unsupported > 0 {
		fmt.Fprintf(&b, " · **Not Assessed:** %d", s.unsupported)
	}
	b.WriteString("\n\n")

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
		fmt.Fprintf(&b, "| [%s](%s) | %s | %s %s | %s | %s | %s |\n",
			name, link, u.UnitType, gradeEmoji(u.Grade), u.Grade,
			FormatRate(u.ScoreKnown(), u.Score, 1), u.Status, formatDate(u.ExpiresAt))
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
