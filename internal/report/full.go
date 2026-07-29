package report

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

// FullReport is a complete, per-unit certification report.
type FullReport struct {
	// Header
	Repository  string `json:"repository"`
	CommitSHA   string `json:"commit_sha,omitempty"`
	GeneratedAt string `json:"generated_at"`

	// Summary card
	Card Card `json:"card"`

	// Every unit with full details
	Units []UnitReport `json:"units"`

	// Dimension averages across all units
	DimensionAverages map[string]float64 `json:"dimension_averages"`

	// By-language detail
	LanguageDetail []LanguageDetail `json:"language_detail"`
}

// EvidenceSummary is a flattened view of a single evidence item for reports.
type EvidenceSummary struct {
	Kind    string             `json:"kind"`
	Source  string             `json:"source"`
	Passed  bool               `json:"passed"`
	Summary string             `json:"summary,omitempty"`
	Metrics map[string]float64 `json:"metrics,omitempty"`
}

// UnitReport is the complete certification detail for a single unit.
type UnitReport struct {
	UnitID     string  `json:"unit_id"`
	UnitType   string  `json:"unit_type"`
	Path       string  `json:"path"`
	Language   string  `json:"language"`
	Symbol     string  `json:"symbol,omitempty"`
	Status     string  `json:"status"`
	Grade      string  `json:"grade"`
	Score      float64 `json:"score"`
	Confidence float64 `json:"confidence"`
	// Unsupported marks a unit the engine could not analyse. Status alone
	// cannot express it — it is "exempt", which IsPassing() reads as true — so
	// every consumer counting over UnitReport needs this flag to avoid
	// certifying code that was never opened.
	Unsupported  bool               `json:"unsupported,omitempty"`
	Dimensions   map[string]float64 `json:"dimensions"`
	Evidence     []EvidenceSummary  `json:"evidence,omitempty"`
	Observations []string           `json:"observations,omitempty"`
	Actions      []string           `json:"actions,omitempty"`
	CertifiedAt  string             `json:"certified_at"`
	ExpiresAt    string             `json:"expires_at"`
	Source       string             `json:"source"`
}

// LanguageDetail is the unified language summary type used across all report formats.
// It replaces the former LanguageCard, LanguageBreakdown, and langRow types.
type LanguageDetail struct {
	Name    string `json:"name"`
	Units   int    `json:"units"`
	Passing int    `json:"passing"`
	// Unsupported counts this language's unassessed units, which are excluded
	// from Passing. See Card.
	Unsupported       int            `json:"unsupported"`
	AverageScore      float64        `json:"average_score"`
	Grade             string         `json:"grade"`
	GradeDistribution map[string]int `json:"grade_distribution"`
	TopScore          float64        `json:"top_score"`
	BottomScore       float64        `json:"bottom_score"`
}

// ScoreKnown reports whether this unit's Score is a measurement. An unassessed
// unit carries the placeholder zero the pipeline assigns when it declines to
// score, so rendering it as "0.0%" states a measured total failure next to a
// grade of N/A that says nothing was measured. See Card.ScoreKnown.
func (u UnitReport) ScoreKnown() bool { return !u.Unsupported }

// Analyzable is the number of this language's units about which a verdict was
// asserted — the denominator of any rate over the language.
func (l LanguageDetail) Analyzable() int { return l.Units - l.Unsupported }

// ScoreKnown reports whether AverageScore and Grade are measurements. See
// Card.ScoreKnown.
func (l LanguageDetail) ScoreKnown() bool { return l.Analyzable() > 0 }

// GenerateFullReport creates a comprehensive per-unit report.
func GenerateFullReport(records []domain.CertificationRecord, repo, commit string, now time.Time) FullReport {
	r := FullReport{
		Repository:        repo,
		CommitSHA:         commit,
		GeneratedAt:       now.Format(time.RFC3339),
		Card:              GenerateCard(records, repo, commit, now),
		DimensionAverages: computeDimensionAverages(records),
	}

	// Build per-unit reports
	r.Units = make([]UnitReport, 0, len(records))
	for _, rec := range records {
		r.Units = append(r.Units, unitReportFrom(rec))
	}
	sort.Slice(r.Units, func(i, j int) bool {
		return r.Units[i].UnitID < r.Units[j].UnitID
	})

	// Build language detail
	r.LanguageDetail = buildLanguageDetail(records)

	return r
}

func unitReportFrom(rec domain.CertificationRecord) UnitReport {
	dims := make(map[string]float64, len(rec.Dimensions))
	for d, v := range rec.Dimensions {
		dims[d.String()] = v
	}

	var evSummaries []EvidenceSummary
	for _, e := range rec.Evidence {
		es := EvidenceSummary{
			Kind:    e.Kind.String(),
			Source:  e.Source,
			Passed:  e.Passed,
			Summary: e.Summary,
		}
		if len(e.Metrics) > 0 {
			es.Metrics = make(map[string]float64, len(e.Metrics))
			for k, v := range e.Metrics {
				es.Metrics[k] = v
			}
		}
		evSummaries = append(evSummaries, es)
	}

	return UnitReport{
		UnitID:       rec.UnitID.String(),
		UnitType:     rec.UnitType.String(),
		Path:         rec.UnitPath,
		Language:     rec.UnitID.Language(),
		Symbol:       rec.UnitID.Symbol(),
		Status:       rec.Status.String(),
		Grade:        rec.Grade.String(),
		Score:        rec.Score,
		Confidence:   rec.Confidence,
		Unsupported:  rec.Unsupported,
		Dimensions:   dims,
		Evidence:     evSummaries,
		Observations: rec.Observations,
		Actions:      rec.Actions,
		CertifiedAt:  rec.CertifiedAt.Format(time.RFC3339),
		ExpiresAt:    rec.ExpiresAt.Format(time.RFC3339),
		Source:       rec.Source,
	}
}

func buildLanguageDetail(records []domain.CertificationRecord) []LanguageDetail {
	type langAccum struct {
		scores      []float64
		grades      map[string]int
		passing     int
		unsupported int
	}
	accum := make(map[string]*langAccum)

	for _, r := range records {
		lang := r.UnitID.Language()
		a, ok := accum[lang]
		if !ok {
			a = &langAccum{grades: make(map[string]int)}
			accum[lang] = a
		}
		a.scores = append(a.scores, r.Score)
		a.grades[r.Grade.String()]++
		// StatusExempt is IsPassing(), so an unassessed unit would otherwise be
		// counted as a passing unit of its language.
		if r.Unsupported {
			a.unsupported++
			continue
		}
		if r.Status.IsPassing() {
			a.passing++
		}
	}

	var details []LanguageDetail
	for lang, a := range accum {
		var sum, top, bottom float64
		bottom = 1.0
		for _, s := range a.scores {
			sum += s
			if s > top {
				top = s
			}
			if s < bottom {
				bottom = s
			}
		}
		// A language none of whose units could be analysed has no mean score and
		// no grade: every score summed above is the placeholder an unassessed
		// unit carries, and averaging placeholders produced a confident F for a
		// language the engine has no analyzer for. The denominator when some
		// unit WAS analysed is left at every unit, matching Card.OverallScore —
		// see #32 and statsForUnits.
		var avg float64
		grade := domain.GradeNA.String()
		if len(a.scores) > a.unsupported {
			avg = sum / float64(len(a.scores))
			grade = domain.GradeFromScore(avg).String()
		}
		details = append(details, LanguageDetail{
			Name:              lang,
			Units:             len(a.scores),
			Passing:           a.passing,
			Unsupported:       a.unsupported,
			AverageScore:      avg,
			Grade:             grade,
			GradeDistribution: a.grades,
			TopScore:          top,
			BottomScore:       bottom,
		})
	}
	sort.Slice(details, func(i, j int) bool {
		return details[i].Units > details[j].Units
	})
	return details
}

// FormatFullMarkdown renders the complete report as a markdown document.
func FormatFullMarkdown(r FullReport) string {
	var b strings.Builder
	writeHeader(&b, r)
	writeSummary(&b, r)
	writeGradeDistribution(&b, r)
	writeDimensionAverages(&b, r)
	writeAIInsights(&b, r)
	writeLanguageDetail(&b, r)
	writeAllUnits(&b, r)
	fmt.Fprintf(&b, "---\n\n*%d units. Generated by [Certify](https://github.com/iksnae/code-certification).*\n", len(r.Units))
	return b.String()
}

func writeHeader(b *strings.Builder, r FullReport) {
	emoji := gradeEmoji(r.Card.OverallGrade)
	fmt.Fprintf(b, "# %s Certify — Report Card\n\n", emoji)
	if r.Repository != "" {
		fmt.Fprintf(b, "**Repository:** `%s`  \n", r.Repository)
	}
	if r.CommitSHA != "" {
		fmt.Fprintf(b, "**Commit:** `%s`  \n", r.CommitSHA)
	}
	fmt.Fprintf(b, "**Generated:** %s  \n", r.GeneratedAt[:19])
	b.WriteString("\n")
}

func writeSummary(b *strings.Builder, r FullReport) {
	emoji := gradeEmoji(r.Card.OverallGrade)
	b.WriteString("---\n\n## Summary\n\n")
	fmt.Fprintf(b, "| Metric | Value |\n|--------|-------|\n")
	fmt.Fprintf(b, "| **Overall Grade** | %s **%s** |\n", emoji, r.Card.OverallGrade)
	fmt.Fprintf(b, "| **Overall Score** | %s |\n", FormatRate(r.Card.ScoreKnown(), r.Card.OverallScore, 1))
	fmt.Fprintf(b, "| **Total Units** | %d |\n", r.Card.TotalUnits)
	fmt.Fprintf(b, "| **Passing** | %d |\n", r.Card.Passing)
	fmt.Fprintf(b, "| **Failing** | %d |\n", r.Card.Failing)
	if r.Card.UnsupportedCount > 0 {
		fmt.Fprintf(b, "| **Not Assessed** | %d |\n", r.Card.UnsupportedCount)
		fmt.Fprintf(b, "| **Analyzable Units** | %d |\n", r.Card.AnalyzableUnits)
	}
	fmt.Fprintf(b, "| **Pass Rate** | %s |\n", formatPassRate(r.Card))
	fmt.Fprintf(b, "| **Observations** | %d |\n", r.Card.Observations)
	fmt.Fprintf(b, "| **Expired** | %d |\n\n", r.Card.Expired)
}

func writeGradeDistribution(b *strings.Builder, r FullReport) {
	b.WriteString("## Grade Distribution\n\n")
	b.WriteString("| Grade | Count | % | Bar |\n|:-----:|------:|----:|-----|\n")
	for _, g := range distributionGrades {
		count := r.Card.GradeDistribution[g]
		if count == 0 {
			continue
		}
		pct := float64(count) / float64(r.Card.TotalUnits) * 100
		barLen := int(pct / 2)
		if barLen < 1 {
			barLen = 1
		}
		fmt.Fprintf(b, "| %s | %d | %.1f%% | %s |\n", g, count, pct, strings.Repeat("█", barLen))
	}
	b.WriteString("\n")
}

func writeDimensionAverages(b *strings.Builder, r FullReport) {
	if len(r.DimensionAverages) == 0 {
		return
	}
	b.WriteString("## Dimension Averages\n\n")
	b.WriteString("| Dimension | Score | Bar |\n|-----------|------:|-----|\n")
	for _, d := range sortedKeys(r.DimensionAverages) {
		v := r.DimensionAverages[d]
		barLen := int(v * 20)
		bar := strings.Repeat("█", barLen) + strings.Repeat("░", 20-barLen)
		fmt.Fprintf(b, "| %s | %.1f%% | %s |\n", d, v*100, bar)
	}
	b.WriteString("\n")
}

func writeLanguageDetail(b *strings.Builder, r FullReport) {
	if len(r.LanguageDetail) == 0 {
		return
	}
	b.WriteString("## By Language\n\n")
	for _, lang := range r.LanguageDetail {
		fmt.Fprintf(b, "### %s — %s %s (%s)\n\n",
			lang.Name, gradeEmoji(lang.Grade), lang.Grade,
			FormatRate(lang.ScoreKnown(), lang.AverageScore, 1))
		fmt.Fprintf(b, "- **Units:** %d\n", lang.Units)
		fmt.Fprintf(b, "- **Score range:** %s – %s\n",
			FormatRate(lang.ScoreKnown(), lang.BottomScore, 1),
			FormatRate(lang.ScoreKnown(), lang.TopScore, 1))
		b.WriteString("- **Grades:** ")
		first := true
		for _, g := range []string{"A", "A-", "B+", "B", "C", "D", "F"} {
			if c := lang.GradeDistribution[g]; c > 0 {
				if !first {
					b.WriteString(", ")
				}
				fmt.Fprintf(b, "%d×%s", c, g)
				first = false
			}
		}
		b.WriteString("\n\n")
	}
}

func writeAllUnits(b *strings.Builder, r FullReport) {
	b.WriteString("## All Units\n\n")

	dirUnits := make(map[string][]UnitReport)
	var dirs []string
	for _, u := range r.Units {
		dir := dirOf(u.Path)
		if _, ok := dirUnits[dir]; !ok {
			dirs = append(dirs, dir)
		}
		dirUnits[dir] = append(dirUnits[dir], u)
	}
	sort.Strings(dirs)

	for _, dir := range dirs {
		units := dirUnits[dir]
		fmt.Fprintf(b, "### `%s/` (%d units)\n\n", dir, len(units))
		b.WriteString("| Unit | Type | Grade | Score | Status | Expires |\n")
		b.WriteString("|------|------|:-----:|------:|--------|--------|\n")
		for _, u := range units {
			name := u.Symbol
			if name == "" {
				name = shortFile(u.Path)
			}
			anchor := unitAnchor(u)
			fmt.Fprintf(b, "| [`%s`](#%s) | %s | %s | %s | %s | %s |\n",
				name, anchor, u.UnitType, u.Grade, FormatRate(u.ScoreKnown(), u.Score, 1),
				u.Status, u.ExpiresAt[:10])
		}
		b.WriteString("\n")
		writeUnitDetails(b, units)
	}
}

func writeAIInsights(b *strings.Builder, r FullReport) {
	// Collect all unique AI observations and suggestions
	var reasons []string
	suggestions := make(map[string]int) // suggestion → count
	aiModel := ""
	for _, u := range r.Units {
		for _, obs := range u.Observations {
			if len(obs) > 2 && obs[:len("🤖")] == "🤖" {
				reasons = append(reasons, obs)
			} else if len(obs) > 2 && obs[:len("💡")] == "💡" {
				suggestions[obs]++
			}
		}
		if aiModel == "" && strings.Contains(u.Source, "agent") {
			// Extract model name from source like "deterministic+agent-prescreen:model-name"
			if idx := strings.LastIndex(u.Source, ":"); idx >= 0 {
				aiModel = u.Source[idx+1:]
			}
		}
	}

	if len(suggestions) == 0 && len(reasons) == 0 {
		return
	}

	b.WriteString("## 🤖 AI Insights\n\n")
	if aiModel != "" {
		fmt.Fprintf(b, "*Powered by `%s` — %d units analyzed*\n\n", aiModel, len(r.Units))
	}

	// Top suggestions by frequency
	if len(suggestions) > 0 {
		b.WriteString("### Top Suggestions\n\n")
		type kv struct {
			key   string
			count int
		}
		var sorted []kv
		for k, v := range suggestions {
			sorted = append(sorted, kv{k, v})
		}
		sort.Slice(sorted, func(i, j int) bool { return sorted[i].count > sorted[j].count })

		shown := 0
		for _, s := range sorted {
			if shown >= 20 {
				fmt.Fprintf(b, "\n*...and %d more suggestions across individual units*\n", len(sorted)-20)
				break
			}
			if s.count > 1 {
				fmt.Fprintf(b, "- %s ×%d\n", s.key, s.count)
			} else {
				fmt.Fprintf(b, "- %s\n", s.key)
			}
			shown++
		}
		b.WriteString("\n")
	}
}

func writeUnitDetails(b *strings.Builder, units []UnitReport) {
	for _, u := range units {
		anchor := unitAnchor(u)
		fmt.Fprintf(b, "<a id=\"%s\"></a>\n<details>\n<summary>%s — %s details</summary>\n\n", anchor, u.Symbol, u.Status)
		if len(u.Dimensions) > 0 {
			b.WriteString("| Dimension | Score |\n|-----------|------:|\n")
			for _, d := range sortedKeys(u.Dimensions) {
				fmt.Fprintf(b, "| %s | %.1f%% |\n", d, u.Dimensions[d]*100)
			}
			b.WriteString("\n")
		}
		if len(u.Observations) > 0 {
			b.WriteString("**Observations:**\n")
			for _, obs := range u.Observations {
				fmt.Fprintf(b, "- %s\n", obs)
			}
			b.WriteString("\n")
		}
		b.WriteString("</details>\n\n")
	}
}

func dirOf(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx < 0 {
		return "."
	}
	return path[:idx]
}

func shortFile(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx < 0 {
		return path
	}
	return path[idx+1:]
}

// unitAnchor generates a stable, GitHub-compatible anchor ID for a unit.
func unitAnchor(u UnitReport) string {
	// Use path + symbol for uniqueness, sanitized for HTML IDs
	raw := u.Path
	if u.Symbol != "" {
		raw += "-" + u.Symbol
	}
	var sb strings.Builder
	for _, r := range strings.ToLower(raw) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			sb.WriteRune(r)
		case r == '/' || r == '.' || r == '_' || r == '-' || r == ' ':
			sb.WriteRune('-')
		}
	}
	return sb.String()
}

func sortedKeys(m map[string]float64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
