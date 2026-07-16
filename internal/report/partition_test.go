package report

import (
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
)

// makeCardRecordT is a test helper for internal tests.
func makeCardRecordT(lang, path, symbol string, status domain.Status, score float64) domain.CertificationRecord {
	return domain.CertificationRecord{
		UnitID:   domain.NewUnitID(lang, path, symbol),
		UnitType: domain.UnitTypeFunction,
		UnitPath: path,
		Status:   status,
		Score:    score,
		Grade:    domain.GradeFromScore(score),
	}
}

func TestPartitionByLanguage_AllSupported(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecordT("go", "a.go", "A", domain.StatusCertified, 0.95),
		makeCardRecordT("go", "b.go", "B", domain.StatusCertified, 0.85),
		makeCardRecordT("ts", "c.ts", "C", domain.StatusCertified, 0.80),
	}

	supported, unsupported := partitionByLanguage(records)

	if len(supported) != 3 {
		t.Errorf("supported = %d, want 3", len(supported))
	}
	if len(unsupported) != 0 {
		t.Errorf("unsupported = %d, want 0", len(unsupported))
	}
}

func TestPartitionByLanguage_AllUnsupported(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecordT("zz", "a.zz", "A", domain.StatusCertified, 0.95),
		makeCardRecordT("xx", "b.xx", "B", domain.StatusCertified, 0.85),
	}

	supported, unsupported := partitionByLanguage(records)

	if len(supported) != 0 {
		t.Errorf("supported = %d, want 0", len(supported))
	}
	if len(unsupported) != 2 {
		t.Errorf("unsupported = %d, want 2", len(unsupported))
	}
}

func TestPartitionByLanguage_Mixed(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecordT("go", "a.go", "A", domain.StatusCertified, 0.95),
		makeCardRecordT("zz", "b.zz", "B", domain.StatusCertified, 0.85),
		makeCardRecordT("ts", "c.ts", "C", domain.StatusCertified, 0.80),
		makeCardRecordT("xx", "d.xx", "D", domain.StatusCertified, 0.70),
	}

	supported, unsupported := partitionByLanguage(records)

	if len(supported) != 2 {
		t.Errorf("supported = %d, want 2", len(supported))
	}
	if len(unsupported) != 2 {
		t.Errorf("unsupported = %d, want 2", len(unsupported))
	}

	// Verify the supported ones are go and ts
	langs := make(map[string]bool)
	for _, r := range supported {
		langs[r.UnitID.Language()] = true
	}
	if !langs["go"] || !langs["ts"] {
		t.Errorf("expected supported languages go and ts, got %v", langs)
	}

	// Verify the unsupported ones are zz and xx
	unsupLangs := make(map[string]bool)
	for _, r := range unsupported {
		unsupLangs[r.UnitID.Language()] = true
	}
	if !unsupLangs["zz"] || !unsupLangs["xx"] {
		t.Errorf("expected unsupported languages zz and xx, got %v", unsupLangs)
	}
}

func TestPartitionByLanguage_Empty(t *testing.T) {
	supported, unsupported := partitionByLanguage(nil)

	if len(supported) != 0 {
		t.Errorf("supported = %d, want 0", len(supported))
	}
	if len(unsupported) != 0 {
		t.Errorf("unsupported = %d, want 0", len(unsupported))
	}
}

func TestCardNewFieldsDefaultZero(t *testing.T) {
	c := Card{}

	if c.UnsupportedCount != 0 {
		t.Errorf("UnsupportedCount = %d, want 0", c.UnsupportedCount)
	}
	if c.AnalyzableUnits != 0 {
		t.Errorf("AnalyzableUnits = %d, want 0", c.AnalyzableUnits)
	}
	if c.UnsupportedLanguages != nil {
		t.Errorf("UnsupportedLanguages = %v, want nil", c.UnsupportedLanguages)
	}
}
