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
	Title          string
	CommitSHA      string
	GeneratedAt    string
	GradeEmoji     string
	OverallGrade   string
	OverallScore   float64
	TotalUnits     int
	PassRate       float64
	Passing        int
	Failing        int
	HasGrades      bool
	Grades         []gradeRow
	HasDimensions  bool
	Dimensions     []dimRow
	HasLanguages   bool
	Languages      []langRow
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

type langRow struct {
	Name         string
	Units        int
	Grade        string
	CSSClass     string
	AverageScore float64
}

type pkgRow struct {
	Path     string
	Units    int
	Grade    string
	CSSClass string
	AvgScore float64
}

type issueRow struct {
	Name     string
	Anchor   string
	Grade    string
	CSSClass string
	Score    float64
	Reason   string
}

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
		OverallScore:   r.Card.OverallScore,
		TotalUnits:     r.Card.TotalUnits,
		PassRate:       r.Card.PassRate,
		Passing:        r.Card.Passing,
		Failing:        r.Card.Failing,
		IncludeSearch:  cfg.IncludeSearch,
		ReportCardLink: true,
	}

	// Grade distribution
	gradeOrder := []string{"A", "A-", "B+", "B", "C", "D", "F"}
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
	for _, l := range r.LanguageDetail {
		data.Languages = append(data.Languages, langRow{
			Name:         l.Name,
			Units:        l.Units,
			Grade:        l.Grade,
			CSSClass:     gradeCSSClass(l.Grade),
			AverageScore: l.AverageScore,
		})
	}
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
			Name:     name,
			Anchor:   anchor,
			Grade:    issue.Grade,
			CSSClass: gradeCSSClass(issue.Grade),
			Score:    issue.Score,
			Reason:   issue.Reason,
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
	AvgScore    float64
	UnitCount   int
	PassRate    float64
	Units       []packageUnitRow
	IndexURL    string
}

type packageUnitRow struct {
	Name      string
	UnitType  string
	Grade     string
	CSSClass  string
	Score     float64
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

		// Compute stats
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
		passRate := float64(passing) / float64(len(units))
		grade := domain.GradeFromScore(avgScore).String()

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
				Score:     u.Score,
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
			AvgScore:    avgScore,
			UnitCount:   len(units),
			PassRate:    passRate,
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
	Title                string
	Name                 string
	GradeEmoji           string
	UnitID               string
	UnitType             string
	Path                 string
	Language             string
	Symbol               string
	Grade                string
	CSSClass             string
	Score                float64
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

	// Group units by directory for prev/next navigation
	dirUnits := make(map[string][]UnitReport)
	for _, u := range r.Units {
		dir := dirOf(u.Path)
		dirUnits[dir] = append(dirUnits[dir], u)
	}

	// Build prev/next maps
	type navLink struct {
		URL  string
		Name string
	}
	prevMap := make(map[string]navLink) // unitID -> prev
	nextMap := make(map[string]navLink) // unitID -> next

	for _, units := range dirUnits {
		for i, u := range units {
			if i > 0 {
				prev := units[i-1]
				prevName := prev.Symbol
				if prevName == "" {
					prevName = shortFile(prev.Path)
				}
				prevMap[u.UnitID] = navLink{
					URL:  unitAnchor(prev) + ".html",
					Name: prevName,
				}
			}
			if i < len(units)-1 {
				next := units[i+1]
				nextName := next.Symbol
				if nextName == "" {
					nextName = shortFile(next.Path)
				}
				nextMap[u.UnitID] = navLink{
					URL:  unitAnchor(next) + ".html",
					Name: nextName,
				}
			}
		}
	}

	for _, u := range r.Units {
		name := u.Symbol
		if name == "" {
			name = shortFile(u.Path)
		}
		dir := dirOf(u.Path)
		anchor := unitAnchor(u)

		// Relative paths from units/ to other directories
		packageURL := "../packages/" + dir + "/index.html"
		indexURL := "../index.html"

		ai, suggestions, other := splitObservations(u.Observations)
		// Strip emoji prefixes for display
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
			Score:                u.Score,
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
			PackageURL:           packageURL,
			IndexURL:             indexURL,
		}

		if prev, ok := prevMap[u.UnitID]; ok {
			data.PrevURL = prev.URL
			data.PrevName = prev.Name
		}
		if next, ok := nextMap[u.UnitID]; ok {
			data.NextURL = next.URL
			data.NextName = next.Name
		}

		outPath := filepath.Join(cfg.OutputDir, "units", anchor+".html")
		if err := writeTemplate(tmpl, data, outPath); err != nil {
			return fmt.Errorf("writing unit page %s: %w", u.UnitID, err)
		}
	}

	return nil
}

// --- Helpers ---

func buildPackageStats(r FullReport) []pkgRow {
	type pkgAccum struct {
		units int
		score float64
	}
	accum := make(map[string]*pkgAccum)
	var dirs []string

	for _, u := range r.Units {
		dir := dirOf(u.Path)
		a, ok := accum[dir]
		if !ok {
			a = &pkgAccum{}
			accum[dir] = a
			dirs = append(dirs, dir)
		}
		a.units++
		a.score += u.Score
	}

	sort.Strings(dirs)

	var rows []pkgRow
	for _, dir := range dirs {
		a := accum[dir]
		avg := a.score / float64(a.units)
		grade := domain.GradeFromScore(avg).String()
		rows = append(rows, pkgRow{
			Path:     dir,
			Units:    a.units,
			Grade:    grade,
			CSSClass: gradeCSSClass(grade),
			AvgScore: avg,
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
