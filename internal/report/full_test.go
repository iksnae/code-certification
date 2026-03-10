package report_test

import (
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

func makeFullRecord(lang, path, symbol string, utype domain.UnitType, status domain.Status, score float64) domain.CertificationRecord {
	now := time.Now()
	dims := domain.DimensionScores{
		domain.DimCorrectness:     score + 0.05,
		domain.DimMaintainability: score,
		domain.DimReadability:     score - 0.02,
		domain.DimTestability:     score + 0.03,
		domain.DimSecurity:        0.80,
	}
	return domain.CertificationRecord{
		UnitID:      domain.NewUnitID(lang, path, symbol),
		UnitType:    utype,
		UnitPath:    path,
		Status:      status,
		Score:       score,
		Grade:       domain.GradeFromScore(score),
		Confidence:  1.0,
		Dimensions:  dims,
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
		Source:      "deterministic",
	}
}

func TestGenerateFullReport(t *testing.T) {
	records := []domain.CertificationRecord{
		makeFullRecord("go", "internal/engine/scorer.go", "Score", domain.UnitTypeFunction, domain.StatusCertified, 0.90),
		makeFullRecord("go", "internal/engine/pipeline.go", "CertifyUnit", domain.UnitTypeFunction, domain.StatusCertified, 0.85),
		makeFullRecord("go", "cmd/certify/main.go", "main", domain.UnitTypeFunction, domain.StatusCertified, 0.88),
		makeFullRecord("ts", "src/parser.ts", "parse", domain.UnitTypeFunction, domain.StatusCertifiedWithObservations, 0.75),
	}
	records[3].Observations = []string{"complexity: 25 exceeds threshold 20"}

	r := report.GenerateFullReport(records, "test/repo", "abc123", time.Now())

	if len(r.Units) != 4 {
		t.Errorf("units = %d, want 4", len(r.Units))
	}
	if r.Card.TotalUnits != 4 {
		t.Errorf("card total = %d, want 4", r.Card.TotalUnits)
	}
	if len(r.DimensionAverages) == 0 {
		t.Error("dimension averages should not be empty")
	}
	if len(r.LanguageDetail) != 2 {
		t.Errorf("languages = %d, want 2", len(r.LanguageDetail))
	}

	// Check units are sorted by ID
	for i := 1; i < len(r.Units); i++ {
		if r.Units[i].UnitID < r.Units[i-1].UnitID {
			t.Error("units should be sorted by ID")
		}
	}
}

func TestFormatFullMarkdown(t *testing.T) {
	records := []domain.CertificationRecord{
		makeFullRecord("go", "internal/engine/scorer.go", "Score", domain.UnitTypeFunction, domain.StatusCertified, 0.90),
		makeFullRecord("go", "internal/engine/pipeline.go", "CertifyUnit", domain.UnitTypeFunction, domain.StatusCertified, 0.85),
		makeFullRecord("ts", "src/parser.ts", "parse", domain.UnitTypeFunction, domain.StatusCertifiedWithObservations, 0.75),
	}
	records[2].Observations = []string{"complexity too high"}

	r := report.GenerateFullReport(records, "test/repo", "abc123", time.Now())
	md := report.FormatFullMarkdown(r)

	checks := []string{
		"# ",                    // title
		"## Summary",            // summary section
		"## Grade Distribution", // grade dist
		"## Dimension Averages", // dimensions
		"## By Language",        // languages
		"## All Units",          // all units
		"| Unit |",              // table header
		"Score",                 // unit name in table
		"CertifyUnit",           // unit name in table
		"parse",                 // ts unit
		"complexity too high",   // observation detail
		"Certify",               // footer link
	}
	for _, check := range checks {
		if !strings.Contains(md, check) {
			t.Errorf("markdown should contain %q", check)
		}
	}

	// Verify links are self-contained anchors (not external report files)
	if strings.Contains(md, "](reports/") {
		t.Error("markdown should NOT contain external report file links — use #anchor links")
	}
	if !strings.Contains(md, "](#") {
		t.Error("markdown should contain self-contained #anchor links")
	}

	// Verify every unit has a corresponding <a id="..."> anchor for back-navigation
	for _, u := range r.Units {
		name := u.Symbol
		if name == "" {
			name = u.Path[strings.LastIndex(u.Path, "/")+1:]
		}
		// Find the anchor link in the table
		idx := strings.Index(md, "[`"+name+"`](#")
		if idx < 0 {
			t.Errorf("missing table link for unit %s", name)
			continue
		}
		// Extract anchor name from ](#anchor-name)
		rest := md[idx:]
		start := strings.Index(rest, "](#")
		end := strings.Index(rest[start+3:], ")")
		anchor := rest[start+3 : start+3+end]

		anchorTag := `<a id="` + anchor + `">`
		if !strings.Contains(md, anchorTag) {
			t.Errorf("missing anchor tag %s for unit %s", anchorTag, name)
		}
	}
}

func TestFormatFullMarkdown_AllUnitsHaveAnchors(t *testing.T) {
	// All certified, no observations — previously these were skipped
	records := []domain.CertificationRecord{
		makeFullRecord("go", "pkg/a.go", "FuncA", domain.UnitTypeFunction, domain.StatusCertified, 0.92),
		makeFullRecord("go", "pkg/b.go", "FuncB", domain.UnitTypeFunction, domain.StatusCertified, 0.88),
	}

	r := report.GenerateFullReport(records, "test/repo", "abc123", time.Now())
	md := report.FormatFullMarkdown(r)

	// Both units must have anchor tags even though they have no observations
	if !strings.Contains(md, `<a id="pkg-a-go-funca">`) {
		t.Error("missing anchor for FuncA (certified, no observations)")
	}
	if !strings.Contains(md, `<a id="pkg-b-go-funcb">`) {
		t.Error("missing anchor for FuncB (certified, no observations)")
	}

	// Should have self-contained anchor links (not external report files)
	if !strings.Contains(md, "](#") {
		t.Error("should contain self-contained #anchor links")
	}
}

func TestFullReport_LanguageDetail(t *testing.T) {
	records := []domain.CertificationRecord{
		makeFullRecord("go", "a.go", "A", domain.UnitTypeFunction, domain.StatusCertified, 0.95),
		makeFullRecord("go", "b.go", "B", domain.UnitTypeFunction, domain.StatusCertified, 0.85),
		makeFullRecord("go", "c.go", "C", domain.UnitTypeFunction, domain.StatusCertified, 0.75),
	}

	r := report.GenerateFullReport(records, "", "", time.Now())

	if len(r.LanguageDetail) != 1 {
		t.Fatalf("expected 1 language, got %d", len(r.LanguageDetail))
	}
	ld := r.LanguageDetail[0]
	if ld.Name != "go" {
		t.Errorf("lang = %s, want go", ld.Name)
	}
	if ld.TopScore < 0.94 || ld.TopScore > 0.96 {
		t.Errorf("top = %f, want ~0.95", ld.TopScore)
	}
	if ld.BottomScore < 0.74 || ld.BottomScore > 0.76 {
		t.Errorf("bottom = %f, want ~0.75", ld.BottomScore)
	}
}

func TestLanguageDetail_PassingCount(t *testing.T) {
	records := []domain.CertificationRecord{
		makeFullRecord("go", "a.go", "A", domain.UnitTypeFunction, domain.StatusCertified, 0.95),
		makeFullRecord("go", "b.go", "B", domain.UnitTypeFunction, domain.StatusDecertified, 0.30),
		makeFullRecord("go", "c.go", "C", domain.UnitTypeFunction, domain.StatusCertified, 0.85),
		makeFullRecord("ts", "d.ts", "D", domain.UnitTypeFunction, domain.StatusExpired, 0.70),
		makeFullRecord("ts", "e.ts", "E", domain.UnitTypeFunction, domain.StatusCertified, 0.88),
	}

	r := report.GenerateFullReport(records, "test/repo", "", time.Now())

	// Verify LanguageDetail has correct Passing counts
	langMap := make(map[string]report.LanguageDetail)
	for _, ld := range r.LanguageDetail {
		langMap[ld.Name] = ld
	}

	goLang := langMap["go"]
	if goLang.Units != 3 {
		t.Errorf("go units = %d, want 3", goLang.Units)
	}
	if goLang.Passing != 2 {
		t.Errorf("go passing = %d, want 2", goLang.Passing)
	}

	tsLang := langMap["ts"]
	if tsLang.Units != 2 {
		t.Errorf("ts units = %d, want 2", tsLang.Units)
	}
	if tsLang.Passing != 1 {
		t.Errorf("ts passing = %d, want 1", tsLang.Passing)
	}

	// Verify Card.Languages uses the same unified type with Passing
	cardLangMap := make(map[string]report.LanguageDetail)
	for _, ld := range r.Card.Languages {
		cardLangMap[ld.Name] = ld
	}
	if cardLangMap["go"].Passing != 2 {
		t.Errorf("card go passing = %d, want 2", cardLangMap["go"].Passing)
	}
	if cardLangMap["go"].GradeDistribution == nil {
		t.Error("card language should have GradeDistribution (unified type)")
	}
}

func TestLanguageDetail_AllPassing(t *testing.T) {
	records := []domain.CertificationRecord{
		makeFullRecord("go", "a.go", "A", domain.UnitTypeFunction, domain.StatusCertified, 0.90),
		makeFullRecord("go", "b.go", "B", domain.UnitTypeFunction, domain.StatusCertified, 0.85),
	}

	r := report.GenerateFullReport(records, "", "", time.Now())
	if len(r.LanguageDetail) != 1 {
		t.Fatalf("expected 1 language, got %d", len(r.LanguageDetail))
	}
	ld := r.LanguageDetail[0]
	if ld.Passing != ld.Units {
		t.Errorf("all units passing: passing=%d should equal units=%d", ld.Passing, ld.Units)
	}
}

func TestLanguageDetail_NonePassing(t *testing.T) {
	records := []domain.CertificationRecord{
		makeFullRecord("go", "a.go", "A", domain.UnitTypeFunction, domain.StatusDecertified, 0.20),
		makeFullRecord("go", "b.go", "B", domain.UnitTypeFunction, domain.StatusExpired, 0.30),
	}

	r := report.GenerateFullReport(records, "", "", time.Now())
	if len(r.LanguageDetail) != 1 {
		t.Fatalf("expected 1 language, got %d", len(r.LanguageDetail))
	}
	ld := r.LanguageDetail[0]
	if ld.Passing != 0 {
		t.Errorf("no units passing: passing=%d, want 0", ld.Passing)
	}
}
