package report

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// SiteConfig configures static site generation.
type SiteConfig struct {
	OutputDir     string // Directory to write the site into
	Title         string // Site title (typically repo name)
	BasePath      string // Base path for GitHub Pages subdirectory hosting
	IncludeSearch bool   // Whether to include client-side search (default: true)
}

// GenerateSite generates a complete static HTML site from a FullReport.
// The site is self-contained with embedded CSS and works via file:// protocol.
func GenerateSite(r FullReport, cfg SiteConfig) error {
	if cfg.OutputDir == "" {
		return fmt.Errorf("output directory is required")
	}
	if cfg.Title == "" {
		cfg.Title = "Certification Report"
	}

	// Create output directories
	for _, dir := range []string{
		cfg.OutputDir,
		filepath.Join(cfg.OutputDir, "packages"),
		filepath.Join(cfg.OutputDir, "units"),
	} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("creating directory %s: %w", dir, err)
		}
	}

	if err := generateIndex(r, cfg); err != nil {
		return fmt.Errorf("generating index: %w", err)
	}
	if err := generatePackagePages(r, cfg); err != nil {
		return fmt.Errorf("generating package pages: %w", err)
	}
	if err := generateUnitPages(r, cfg); err != nil {
		return fmt.Errorf("generating unit pages: %w", err)
	}

	return nil
}

// --- Index page ---

type indexData struct {
	Title        string
	CommitSHA    string
	GeneratedAt  string
	GradeEmoji   string
	OverallGrade string
	// OverallScore is preformatted, like PassRate below: the undefined case
	// (nothing analyzable) has no number to render, and one formatter for both
	// keeps the markers from drifting apart on the same page.
	OverallScore   string
	TotalUnits     int
	PassRate       string
	Passing        int
	Failing        int
	Unsupported    int
	HasGrades      bool
	Grades         []gradeRow
	HasDimensions  bool
	Dimensions     []dimRow
	HasLanguages   bool
	Languages      []LanguageDetail
	HasPackages    bool
	Packages       []pkgRow
	HasTopIssues   bool
	TopIssues      []issueRow
	IncludeSearch  bool
	ReportCardLink bool
}

type gradeRow struct {
	Name     string
	CSSClass string
	Count    int
	Pct      float64
}

type dimRow struct {
	Name  string
	Score float64
}

type pkgRow struct {
	Path     string
	Units    int
	Grade    string
	CSSClass string
	AvgScore string
}

type issueRow struct {
	Name     string
	Anchor   string
	Grade    string
	CSSClass string
	Score    float64
	// Unsupported carries IssueCard.Unsupported through to the template, which
	// gates on ScoreKnown. See IssueCard.ScoreKnown.
	Unsupported bool
	Reason      string
}

// ScoreKnown reports whether this row's Score is a measurement.
func (r issueRow) ScoreKnown() bool { return !r.Unsupported }

func generateIndex(r FullReport, cfg SiteConfig) error {
	tmpl, err := template.New("index").Funcs(siteFuncMap).Parse(indexTemplateStr)
	if err != nil {
		return fmt.Errorf("parsing index template: %w", err)
	}

	data := indexData{
		Title:          cfg.Title,
		CommitSHA:      r.CommitSHA,
		GeneratedAt:    r.GeneratedAt,
		GradeEmoji:     gradeEmoji(r.Card.OverallGrade),
		OverallGrade:   r.Card.OverallGrade,
		OverallScore:   FormatRate(r.Card.ScoreKnown(), r.Card.OverallScore, 1),
		TotalUnits:     r.Card.TotalUnits,
		PassRate:       formatPassRate(r.Card),
		Passing:        r.Card.Passing,
		Failing:        r.Card.Failing,
		Unsupported:    r.Card.UnsupportedCount,
		IncludeSearch:  cfg.IncludeSearch,
		ReportCardLink: true,
	}

	// Grade distribution
	gradeOrder := distributionGrades
	for _, g := range gradeOrder {
		if count, ok := r.Card.GradeDistribution[g]; ok && count > 0 {
			data.Grades = append(data.Grades, gradeRow{
				Name:     g,
				CSSClass: gradeCSSClass(g),
				Count:    count,
				Pct:      float64(count) / float64(max(r.Card.TotalUnits, 1)),
			})
		}
	}
	data.HasGrades = len(data.Grades) > 0

	// Dimension averages
	for _, key := range sortedKeys(r.DimensionAverages) {
		data.Dimensions = append(data.Dimensions, dimRow{
			Name:  key,
			Score: r.DimensionAverages[key],
		})
	}
	data.HasDimensions = len(data.Dimensions) > 0

	// Languages
	data.Languages = r.LanguageDetail
	data.HasLanguages = len(data.Languages) > 0

	// Packages
	pkgStats := buildPackageStats(r)
	for _, p := range pkgStats {
		data.Packages = append(data.Packages, p)
	}
	data.HasPackages = len(data.Packages) > 0

	// Top issues
	for _, issue := range r.Card.TopIssues {
		name := issue.UnitID
		anchor := ""
		// Find matching unit for anchor
		for _, u := range r.Units {
			if u.UnitID == issue.UnitID {
				name = u.Symbol
				if name == "" {
					name = shortFile(u.Path)
				}
				anchor = unitAnchor(u)
				break
			}
		}
		data.TopIssues = append(data.TopIssues, issueRow{
			Name:        name,
			Anchor:      anchor,
			Grade:       issue.Grade,
			CSSClass:    gradeCSSClass(issue.Grade),
			Score:       issue.Score,
			Unsupported: issue.Unsupported,
			Reason:      issue.Reason,
		})
	}
	data.HasTopIssues = len(data.TopIssues) > 0

	// Search index — write as separate JS file to keep index.html small
	if cfg.IncludeSearch && len(r.Units) > 0 {
		entries := BuildSearchIndex(r)
		jsContent := FormatSearchIndexJS(entries)
		jsPath := filepath.Join(cfg.OutputDir, "search-index.js")
		if err := os.WriteFile(jsPath, []byte(jsContent), 0o644); err != nil {
			return fmt.Errorf("writing search index: %w", err)
		}
		// Search index loaded via <script src="search-index.js">
	}

	return writeTemplate(tmpl, data, filepath.Join(cfg.OutputDir, "index.html"))
}

// --- Package pages ---

type packagePageData struct {
	Title       string
	PackagePath string
	GradeEmoji  string
	Grade       string
	AvgScore    string
	UnitCount   int
	PassRate    string
	Unsupported int
	Units       []packageUnitRow
	IndexURL    string
}

type packageUnitRow struct {
	Name     string
	UnitType string
	Grade    string
	CSSClass string
	// Score is preformatted: an unassessed unit's score is a placeholder, and
	// there is no number to render for it. See UnitReport.ScoreKnown.
	Score     string
	Status    string
	ExpiresAt string
	UnitURL   string
}

func generatePackagePages(r FullReport, cfg SiteConfig) error {
	tmpl, err := template.New("package").Funcs(siteFuncMap).Parse(packageTemplateStr)
	if err != nil {
		return fmt.Errorf("parsing package template: %w", err)
	}

	// Group units by directory
	dirUnits := make(map[string][]UnitReport)
	for _, u := range r.Units {
		dir := dirOf(u.Path)
		dirUnits[dir] = append(dirUnits[dir], u)
	}

	for dir, units := range dirUnits {
		// Sort by score ascending (worst first)
		sort.Slice(units, func(i, j int) bool {
			return units[i].Score < units[j].Score
		})

		// One aggregation, shared with the markdown surfaces. See packageStats.
		stats := statsForUnits(units)
		passRate := FormatRate(stats.measured(), stats.passRate(), 1)
		grade := stats.grade()

		// Compute relative path from package page to site root
		depth := strings.Count(dir, "/") + 1 // +1 for packages/ prefix
		indexURL := strings.Repeat("../", depth+1) + "index.html"

		var unitRows []packageUnitRow
		for _, u := range units {
			name := u.Symbol
			if name == "" {
				name = shortFile(u.Path)
			}
			anchor := unitAnchor(u)
			unitURL := strings.Repeat("../", depth+1) + "units/" + anchor + ".html"

			unitRows = append(unitRows, packageUnitRow{
				Name:      name,
				UnitType:  u.UnitType,
				Grade:     u.Grade,
				CSSClass:  gradeCSSClass(u.Grade),
				Score:     FormatRate(u.ScoreKnown(), u.Score, 1),
				Status:    u.Status,
				ExpiresAt: formatDate(u.ExpiresAt),
				UnitURL:   unitURL,
			})
		}

		data := packagePageData{
			Title:       cfg.Title,
			PackagePath: dir,
			GradeEmoji:  gradeEmoji(grade),
			Grade:       grade,
			AvgScore:    FormatRate(stats.measured(), stats.avgScore, 1),
			UnitCount:   stats.units,
			PassRate:    passRate,
			Unsupported: stats.unsupported,
			Units:       unitRows,
			IndexURL:    indexURL,
		}

		outPath := filepath.Join(cfg.OutputDir, "packages", dir, "index.html")
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return fmt.Errorf("creating package dir: %w", err)
		}
		if err := writeTemplate(tmpl, data, outPath); err != nil {
			return fmt.Errorf("writing package page %s: %w", dir, err)
		}
	}

	return nil
}

// --- Unit pages ---

type unitPageData struct {
	Title      string
	Name       string
	GradeEmoji string
	UnitID     string
	UnitType   string
	Path       string
	Language   string
	Symbol     string
	Grade      string
	CSSClass   string
	// Score is preformatted. See packageUnitRow.Score.
	Score                string
	Status               string
	Confidence           float64
	CertifiedAt          string
	ExpiresAt            string
	Source               string
	HasDimensions        bool
	Dimensions           []dimRow
	HasAIObservations    bool
	AIObservations       []string
	HasSuggestions       bool
	Suggestions          []string
	HasOtherObservations bool
	OtherObservations    []string
	HasActions           bool
	Actions              []string
	PackagePath          string
	PackageURL           string
	IndexURL             string
	PrevURL              string
	PrevName             string
	NextURL              string
	NextName             string
}

func generateUnitPages(r FullReport, cfg SiteConfig) error {
	tmpl, err := template.New("unit").Funcs(siteFuncMap).Parse(unitTemplateStr)
	if err != nil {
		return fmt.Errorf("parsing unit template: %w", err)
	}

	prevMap, nextMap := buildUnitNavMaps(r.Units)

	for _, u := range r.Units {
		data := buildUnitPageData(u, cfg, prevMap, nextMap)
		anchor := unitAnchor(u)
		outPath := filepath.Join(cfg.OutputDir, "units", anchor+".html")
		if err := writeTemplate(tmpl, data, outPath); err != nil {
			return fmt.Errorf("writing unit page %s: %w", u.UnitID, err)
		}
	}
	return nil
}

type navLink struct {
	URL  string
	Name string
}

func buildUnitNavMaps(units []UnitReport) (prev, next map[string]navLink) {
	dirUnits := make(map[string][]UnitReport)
	for _, u := range units {
		dir := dirOf(u.Path)
		dirUnits[dir] = append(dirUnits[dir], u)
	}

	prev = make(map[string]navLink)
	next = make(map[string]navLink)
	for _, units := range dirUnits {
		for i, u := range units {
			if i > 0 {
				p := units[i-1]
				name := p.Symbol
				if name == "" {
					name = shortFile(p.Path)
				}
				prev[u.UnitID] = navLink{URL: unitAnchor(p) + ".html", Name: name}
			}
			if i < len(units)-1 {
				n := units[i+1]
				name := n.Symbol
				if name == "" {
					name = shortFile(n.Path)
				}
				next[u.UnitID] = navLink{URL: unitAnchor(n) + ".html", Name: name}
			}
		}
	}
	return prev, next
}

func buildUnitPageData(u UnitReport, cfg SiteConfig, prevMap, nextMap map[string]navLink) unitPageData {
	name := u.Symbol
	if name == "" {
		name = shortFile(u.Path)
	}
	dir := dirOf(u.Path)

	ai, suggestions, other := splitObservations(u.Observations)
	cleanAI := make([]string, len(ai))
	for i, o := range ai {
		cleanAI[i] = strings.TrimPrefix(o, "🤖 ")
	}
	cleanSuggestions := make([]string, len(suggestions))
	for i, o := range suggestions {
		cleanSuggestions[i] = strings.TrimPrefix(o, "💡 ")
	}

	var dims []dimRow
	for _, key := range sortedKeys(u.Dimensions) {
		dims = append(dims, dimRow{Name: key, Score: u.Dimensions[key]})
	}

	data := unitPageData{
		Title:                cfg.Title,
		Name:                 name,
		GradeEmoji:           gradeEmoji(u.Grade),
		UnitID:               u.UnitID,
		UnitType:             u.UnitType,
		Path:                 u.Path,
		Language:             u.Language,
		Symbol:               u.Symbol,
		Grade:                u.Grade,
		CSSClass:             gradeCSSClass(u.Grade),
		Score:                FormatRate(u.ScoreKnown(), u.Score, 1),
		Status:               u.Status,
		Confidence:           u.Confidence,
		CertifiedAt:          formatDate(u.CertifiedAt),
		ExpiresAt:            formatDate(u.ExpiresAt),
		Source:               u.Source,
		HasDimensions:        len(dims) > 0,
		Dimensions:           dims,
		HasAIObservations:    len(cleanAI) > 0,
		AIObservations:       cleanAI,
		HasSuggestions:       len(cleanSuggestions) > 0,
		Suggestions:          cleanSuggestions,
		HasOtherObservations: len(other) > 0,
		OtherObservations:    other,
		HasActions:           len(u.Actions) > 0,
		Actions:              u.Actions,
		PackagePath:          dir,
		PackageURL:           "../packages/" + dir + "/index.html",
		IndexURL:             "../index.html",
	}

	if p, ok := prevMap[u.UnitID]; ok {
		data.PrevURL = p.URL
		data.PrevName = p.Name
	}
	if n, ok := nextMap[u.UnitID]; ok {
		data.NextURL = n.URL
		data.NextName = n.Name
	}
	return data
}

// --- Helpers ---

// buildPackageStats builds the site dashboard's packages table.
//
// It was the fourth independent aggregation over a package's units, with its
// own accumulator and its own call to GradeFromScore, so the unassessed-unit
// correction reached the other three and left this one grading an unopened
// package F. It now reads the shared aggregation, which is what packageStats
// exists to guarantee: a new reader must not be a new chance to get it wrong.
func buildPackageStats(r FullReport) []pkgRow {
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

	var rows []pkgRow
	for _, dir := range dirs {
		s := statsForUnits(pkgUnits[dir])
		grade := s.grade()
		rows = append(rows, pkgRow{
			Path:     dir,
			Units:    s.units,
			Grade:    grade,
			CSSClass: gradeCSSClass(grade),
			AvgScore: FormatRate(s.measured(), s.avgScore, 1),
		})
	}
	return rows
}

func gradeCSSClass(grade string) string {
	g := strings.ToLower(grade)
	g = strings.ReplaceAll(g, "+", "plus")
	if g == "" || g == "n/a" {
		return "na"
	}
	return g
}

func statusFromString(s string) domain.Status {
	switch s {
	case "certified":
		return domain.StatusCertified
	case "certified_with_observations":
		return domain.StatusCertifiedWithObservations
	case "probationary":
		return domain.StatusProbationary
	case "expired":
		return domain.StatusExpired
	case "decertified":
		return domain.StatusDecertified
	case "exempt":
		return domain.StatusExempt
	default:
		return domain.StatusDecertified
	}
}

func writeTemplate(tmpl *template.Template, data any, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, data)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
