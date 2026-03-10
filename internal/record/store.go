// Package record handles persistence of certification records.
package record

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

// recordJSON is the JSON-serializable form of a CertificationRecord.
type recordJSON struct {
	UnitID       string             `json:"unit_id"`
	UnitType     string             `json:"unit_type"`
	UnitPath     string             `json:"unit_path"`
	PolicyVer    string             `json:"policy_version"`
	Status       string             `json:"status"`
	Grade        string             `json:"grade"`
	Score        float64            `json:"score"`
	Confidence   float64            `json:"confidence"`
	Dimensions   map[string]float64 `json:"dimensions,omitempty"`
	Evidence     []evidenceJSON     `json:"evidence,omitempty"`
	Observations []string           `json:"observations,omitempty"`
	Actions      []string           `json:"actions,omitempty"`
	CertifiedAt  string             `json:"certified_at"`
	ExpiresAt    string             `json:"expires_at"`
	Source       string             `json:"source"`
	RunID        string             `json:"run_id,omitempty"`
	Version      int                `json:"version"`
}

// evidenceJSON is the JSON-serializable form of a domain.Evidence.
// Details is stored as json.RawMessage to preserve whatever concrete type
// was serialized (LintResult, CodeMetrics, GitStats, ReviewResult) without
// the record store needing to know about those types.
type evidenceJSON struct {
	Kind       string             `json:"kind"`
	Source     string             `json:"source"`
	Passed     bool               `json:"passed"`
	Missing    bool               `json:"missing,omitempty"`
	Summary    string             `json:"summary"`
	Metrics    map[string]float64 `json:"metrics,omitempty"`
	Details    json.RawMessage    `json:"details,omitempty"`
	Timestamp  string             `json:"timestamp"`
	Confidence float64            `json:"confidence"`
}

// snapshotJSON is the JSON-serializable form of a state snapshot.
type snapshotJSON struct {
	Version     int          `json:"version"`
	GeneratedAt string       `json:"generated_at"`
	Commit      string       `json:"commit"`
	UnitCount   int          `json:"unit_count"`
	Records     []recordJSON `json:"records"`
}

// Store manages certification record files.
type Store struct {
	dir          string
	snapshotPath string // optional: path to state.json for ListAll fallback
}

// NewStore creates a new record store rooted at the given directory.
func NewStore(dir string) *Store {
	return &Store{dir: dir}
}

// NewStoreWithSnapshot creates a record store with a snapshot fallback path.
// When ListAll finds no individual record files, it reads from the snapshot.
func NewStoreWithSnapshot(dir, snapshotPath string) *Store {
	return &Store{dir: dir, snapshotPath: snapshotPath}
}

// SaveSnapshot writes all current records to a single JSON snapshot file.
// Records are sorted by UnitID for deterministic output (clean git diffs).
func (s *Store) SaveSnapshot(path string, commit string) error {
	records, err := s.listAllFromDir()
	if err != nil {
		return fmt.Errorf("listing records for snapshot: %w", err)
	}

	// Sort by UnitID for deterministic output
	sort.Slice(records, func(i, j int) bool {
		return records[i].UnitID.String() < records[j].UnitID.String()
	})

	var recs []recordJSON
	for _, r := range records {
		recs = append(recs, toJSON(r))
	}

	snap := snapshotJSON{
		Version:     1,
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Commit:      commit,
		UnitCount:   len(recs),
		Records:     recs,
	}

	data, err := json.MarshalIndent(snap, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling snapshot: %w", err)
	}

	if dir := filepath.Dir(path); dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("creating snapshot directory: %w", err)
		}
	}

	return os.WriteFile(path, data, 0o644)
}

// LoadSnapshot reads records from a snapshot file and populates the store
// by writing individual record files. Use this to restore state after clone.
func (s *Store) LoadSnapshot(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading snapshot: %w", err)
	}

	var snap snapshotJSON
	if err := json.Unmarshal(data, &snap); err != nil {
		return fmt.Errorf("parsing snapshot: %w", err)
	}

	for _, rj := range snap.Records {
		rec := fromJSON(rj)
		if err := s.Save(rec); err != nil {
			return fmt.Errorf("restoring record %s: %w", rj.UnitID, err)
		}
	}
	return nil
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
// If the records directory is empty or missing and a snapshot path is configured,
// it falls back to reading from the snapshot file.
func (s *Store) ListAll() ([]domain.CertificationRecord, error) {
	records, err := s.listAllFromDir()
	if err != nil {
		return nil, err
	}
	if len(records) > 0 {
		return records, nil
	}

	// Fallback to snapshot if configured and records dir is empty
	if s.snapshotPath != "" {
		return s.listAllFromSnapshot()
	}
	return records, nil
}

// listAllFromDir reads records from the individual files in the records directory.
func (s *Store) listAllFromDir() ([]domain.CertificationRecord, error) {
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

// listAllFromSnapshot reads records from a snapshot file without writing
// individual record files. Used as a read-only fallback for ListAll.
func (s *Store) listAllFromSnapshot() ([]domain.CertificationRecord, error) {
	data, err := os.ReadFile(s.snapshotPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("reading snapshot: %w", err)
	}

	var snap snapshotJSON
	if err := json.Unmarshal(data, &snap); err != nil {
		return nil, fmt.Errorf("parsing snapshot: %w", err)
	}

	records := make([]domain.CertificationRecord, 0, len(snap.Records))
	for _, rj := range snap.Records {
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
	var evJSON []evidenceJSON
	for _, ev := range rec.Evidence {
		evJSON = append(evJSON, evidenceToJSON(ev))
	}
	return recordJSON{
		UnitID:       rec.UnitID.String(),
		UnitType:     rec.UnitType.String(),
		UnitPath:     rec.UnitPath,
		PolicyVer:    rec.PolicyVersion,
		Status:       rec.Status.String(),
		Grade:        rec.Grade.String(),
		Score:        rec.Score,
		Confidence:   rec.Confidence,
		Dimensions:   dimensionsToMap(rec.Dimensions),
		Evidence:     evJSON,
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

	var evidence []domain.Evidence
	for _, ej := range rj.Evidence {
		evidence = append(evidence, evidenceFromJSON(ej))
	}

	return domain.CertificationRecord{
		UnitID:        id,
		UnitType:      ut,
		UnitPath:      rj.UnitPath,
		PolicyVersion: rj.PolicyVer,
		Status:        st,
		Grade:         parseGrade(rj.Grade),
		Score:         rj.Score,
		Confidence:    rj.Confidence,
		Dimensions:    mapToDimensions(rj.Dimensions),
		Evidence:      evidence,
		Observations:  rj.Observations,
		Actions:       rj.Actions,
		CertifiedAt:   certAt,
		ExpiresAt:     expAt,
		Source:        rj.Source,
		RunID:         rj.RunID,
		Version:       rj.Version,
	}
}

// AppendHistory appends a history entry for the given record.
// History is stored as a JSON-lines file alongside the record.
func (s *Store) AppendHistory(rec domain.CertificationRecord) error {
	histPath := s.historyPathFor(rec.UnitID)
	f, err := os.OpenFile(histPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("opening history file: %w", err)
	}
	defer f.Close()

	entry := historyEntry{
		Status:      rec.Status.String(),
		Score:       rec.Score,
		Grade:       rec.Grade.String(),
		CertifiedAt: rec.CertifiedAt.Format(time.RFC3339),
		Source:      rec.Source,
	}
	data, _ := json.Marshal(entry)
	data = append(data, '\n')
	_, err = f.Write(data)
	return err
}

// LoadHistory returns all history entries for a unit.
func (s *Store) LoadHistory(id domain.UnitID) ([]historyEntry, error) {
	histPath := s.historyPathFor(id)
	data, err := os.ReadFile(histPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var entries []historyEntry
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		if line == "" {
			continue
		}
		var e historyEntry
		if json.Unmarshal([]byte(line), &e) == nil {
			entries = append(entries, e)
		}
	}
	return entries, nil
}

type historyEntry struct {
	Status      string  `json:"status"`
	Score       float64 `json:"score"`
	Grade       string  `json:"grade"`
	CertifiedAt string  `json:"certified_at"`
	Source      string  `json:"source"`
}

func (s *Store) historyPathFor(id domain.UnitID) string {
	h := sha256.Sum256([]byte(id.String()))
	name := hex.EncodeToString(h[:8]) + ".history.jsonl"
	return filepath.Join(s.dir, name)
}

func evidenceToJSON(ev domain.Evidence) evidenceJSON {
	var details json.RawMessage
	if ev.Details != nil {
		data, err := json.Marshal(ev.Details)
		if err == nil {
			details = data
		}
	}
	return evidenceJSON{
		Kind:       ev.Kind.String(),
		Source:     ev.Source,
		Passed:     ev.Passed,
		Missing:    ev.Missing,
		Summary:    ev.Summary,
		Metrics:    ev.Metrics,
		Details:    details,
		Timestamp:  ev.Timestamp.Format(time.RFC3339),
		Confidence: ev.Confidence,
	}
}

func evidenceFromJSON(ej evidenceJSON) domain.Evidence {
	kind, _ := domain.ParseEvidenceKind(ej.Kind)
	ts, _ := time.Parse(time.RFC3339, ej.Timestamp)

	var details any
	if len(ej.Details) > 0 {
		_ = json.Unmarshal(ej.Details, &details)
	}

	return domain.Evidence{
		Kind:       kind,
		Source:     ej.Source,
		Passed:     ej.Passed,
		Missing:    ej.Missing,
		Summary:    ej.Summary,
		Metrics:    ej.Metrics,
		Details:    details,
		Timestamp:  ts,
		Confidence: ej.Confidence,
	}
}

func dimensionsToMap(dims domain.DimensionScores) map[string]float64 {
	if len(dims) == 0 {
		return nil
	}
	m := make(map[string]float64, len(dims))
	for d, v := range dims {
		m[d.String()] = v
	}
	return m
}

func mapToDimensions(m map[string]float64) domain.DimensionScores {
	if len(m) == 0 {
		return nil
	}
	// Build reverse lookup
	lookup := make(map[string]domain.Dimension)
	for _, d := range domain.AllDimensions() {
		lookup[d.String()] = d
	}
	dims := make(domain.DimensionScores, len(m))
	for k, v := range m {
		if d, ok := lookup[k]; ok {
			dims[d] = v
		}
	}
	return dims
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
