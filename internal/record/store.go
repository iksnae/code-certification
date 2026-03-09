// Package record handles persistence of certification records.
package record

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/code-certification/certify/internal/domain"
)

// recordJSON is the JSON-serializable form of a CertificationRecord.
type recordJSON struct {
	UnitID       string   `json:"unit_id"`
	UnitType     string   `json:"unit_type"`
	UnitPath     string   `json:"unit_path"`
	PolicyVer    string   `json:"policy_version"`
	Status       string   `json:"status"`
	Grade        string   `json:"grade"`
	Score        float64  `json:"score"`
	Confidence   float64  `json:"confidence"`
	Observations []string `json:"observations,omitempty"`
	Actions      []string `json:"actions,omitempty"`
	CertifiedAt  string   `json:"certified_at"`
	ExpiresAt    string   `json:"expires_at"`
	Source       string   `json:"source"`
	RunID        string   `json:"run_id,omitempty"`
	Version      int      `json:"version"`
}

// Store manages certification record files.
type Store struct {
	dir string
}

// NewStore creates a new record store rooted at the given directory.
func NewStore(dir string) *Store {
	return &Store{dir: dir}
}

// Save writes a certification record to the store.
func (s *Store) Save(rec domain.CertificationRecord) error {
	if err := os.MkdirAll(s.dir, 0o755); err != nil {
		return fmt.Errorf("creating records directory: %w", err)
	}

	rj := toJSON(rec)
	data, err := json.MarshalIndent(rj, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling record: %w", err)
	}

	path := s.pathFor(rec.UnitID)
	return os.WriteFile(path, data, 0o644)
}

// Load reads a certification record for the given unit.
func (s *Store) Load(id domain.UnitID) (domain.CertificationRecord, error) {
	path := s.pathFor(id)
	data, err := os.ReadFile(path)
	if err != nil {
		return domain.CertificationRecord{}, fmt.Errorf("reading record: %w", err)
	}

	var rj recordJSON
	if err := json.Unmarshal(data, &rj); err != nil {
		return domain.CertificationRecord{}, fmt.Errorf("parsing record: %w", err)
	}

	return fromJSON(rj), nil
}

// ListAll returns all records in the store.
func (s *Store) ListAll() ([]domain.CertificationRecord, error) {
	entries, err := os.ReadDir(s.dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("reading records directory: %w", err)
	}

	var records []domain.CertificationRecord
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dir, e.Name()))
		if err != nil {
			continue
		}
		var rj recordJSON
		if err := json.Unmarshal(data, &rj); err != nil {
			continue
		}
		records = append(records, fromJSON(rj))
	}
	return records, nil
}

func (s *Store) pathFor(id domain.UnitID) string {
	// Use hash of UnitID to create a flat file structure
	h := sha256.Sum256([]byte(id.String()))
	name := hex.EncodeToString(h[:8]) + ".json"
	return filepath.Join(s.dir, name)
}

func toJSON(rec domain.CertificationRecord) recordJSON {
	return recordJSON{
		UnitID:       rec.UnitID.String(),
		UnitType:     rec.UnitType.String(),
		UnitPath:     rec.UnitPath,
		PolicyVer:    rec.PolicyVersion,
		Status:       rec.Status.String(),
		Grade:        rec.Grade.String(),
		Score:        rec.Score,
		Confidence:   rec.Confidence,
		Observations: rec.Observations,
		Actions:      rec.Actions,
		CertifiedAt:  rec.CertifiedAt.Format(time.RFC3339),
		ExpiresAt:    rec.ExpiresAt.Format(time.RFC3339),
		Source:       rec.Source,
		RunID:        rec.RunID,
		Version:      rec.Version,
	}
}

func fromJSON(rj recordJSON) domain.CertificationRecord {
	id, _ := domain.ParseUnitID(rj.UnitID)
	ut, _ := domain.ParseUnitType(rj.UnitType)
	st, _ := domain.ParseStatus(rj.Status)
	certAt, _ := time.Parse(time.RFC3339, rj.CertifiedAt)
	expAt, _ := time.Parse(time.RFC3339, rj.ExpiresAt)

	return domain.CertificationRecord{
		UnitID:        id,
		UnitType:      ut,
		UnitPath:      rj.UnitPath,
		PolicyVersion: rj.PolicyVer,
		Status:        st,
		Grade:         parseGrade(rj.Grade),
		Score:         rj.Score,
		Confidence:    rj.Confidence,
		Observations:  rj.Observations,
		Actions:       rj.Actions,
		CertifiedAt:   certAt,
		ExpiresAt:     expAt,
		Source:        rj.Source,
		RunID:         rj.RunID,
		Version:       rj.Version,
	}
}

func parseGrade(s string) domain.Grade {
	m := map[string]domain.Grade{
		"A": domain.GradeA, "A-": domain.GradeAMinus, "B+": domain.GradeBPlus,
		"B": domain.GradeB, "C": domain.GradeC, "D": domain.GradeD, "F": domain.GradeF,
	}
	if g, ok := m[s]; ok {
		return g
	}
	return domain.GradeF
}
