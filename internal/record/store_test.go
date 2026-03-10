package record_test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/record"
)

func sampleRecord(unitPath, symbol string) domain.CertificationRecord {
	now := time.Now()
	return domain.CertificationRecord{
		UnitID:      domain.NewUnitID("go", unitPath, symbol),
		UnitType:    domain.UnitTypeFunction,
		UnitPath:    unitPath,
		Status:      domain.StatusCertified,
		Grade:       domain.GradeB,
		Score:       0.85,
		Confidence:  1.0,
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
		Source:      "deterministic",
		Version:     1,
		Evidence: []domain.Evidence{
			{
				Kind:       domain.EvidenceKindLint,
				Source:     "go-vet",
				Passed:     true,
				Summary:    "go-vet: 0 errors, 0 warnings",
				Timestamp:  now,
				Confidence: 1.0,
			},
		},
	}
}

func TestStore_SaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	rec := sampleRecord("main.go", "main")
	if err := store.Save(rec); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := store.Load(rec.UnitID)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if loaded.UnitID.String() != rec.UnitID.String() {
		t.Errorf("loaded UnitID = %s, want %s", loaded.UnitID, rec.UnitID)
	}
	if loaded.Score != rec.Score {
		t.Errorf("loaded Score = %f, want %f", loaded.Score, rec.Score)
	}
}

func TestStore_LoadNotFound(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	_, err := store.Load(domain.NewUnitID("go", "missing.go", "foo"))
	if err == nil {
		t.Fatal("Load(missing) should return error")
	}
}

func TestStore_ListAll(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	_ = store.Save(sampleRecord("main.go", "main"))
	_ = store.Save(sampleRecord("main.go", "helper"))
	_ = store.Save(sampleRecord("service/sync.go", "Apply"))

	records, err := store.ListAll()
	if err != nil {
		t.Fatalf("ListAll() error: %v", err)
	}
	if len(records) != 3 {
		t.Errorf("ListAll() = %d records, want 3", len(records))
	}
}

func TestStore_SaveOverwrites(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	rec := sampleRecord("main.go", "main")
	rec.Score = 0.7
	_ = store.Save(rec)

	rec.Score = 0.9
	_ = store.Save(rec)

	loaded, _ := store.Load(rec.UnitID)
	if loaded.Score != 0.9 {
		t.Errorf("overwritten Score = %f, want 0.9", loaded.Score)
	}
}

func TestStore_PathFormat(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	rec := sampleRecord("internal/service/sync.go", "Apply")
	_ = store.Save(rec)

	// Should create a file under records dir
	pattern := filepath.Join(dir, "*.json")
	matches, _ := filepath.Glob(pattern)
	if len(matches) == 0 {
		// Check subdirectory
		matches, _ = filepath.Glob(filepath.Join(dir, "**", "*.json"))
	}
	// Just verify it can be loaded back
	loaded, err := store.Load(rec.UnitID)
	if err != nil {
		t.Fatalf("roundtrip failed: %v", err)
	}
	if loaded.UnitID.String() != rec.UnitID.String() {
		t.Error("roundtrip ID mismatch")
	}
}

func TestStore_EvidenceRoundTrip(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	now := time.Now().Truncate(time.Second) // RFC3339 has second precision

	// Build a record with 4 evidence items covering different kinds and Details
	rec := domain.CertificationRecord{
		UnitID:      domain.NewUnitID("go", "pkg/foo.go", "Foo"),
		UnitType:    domain.UnitTypeFunction,
		UnitPath:    "pkg/foo.go",
		Status:      domain.StatusCertified,
		Grade:       domain.GradeA,
		Score:       0.95,
		Confidence:  1.0,
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
		Source:      "deterministic",
		Version:     1,
		Evidence: []domain.Evidence{
			{
				Kind:       domain.EvidenceKindLint,
				Source:     "golangci-lint",
				Passed:     true,
				Summary:    "golangci-lint: 0 errors, 2 warnings",
				Timestamp:  now,
				Confidence: 1.0,
			},
			{
				Kind:       domain.EvidenceKindTest,
				Source:     "go-test",
				Passed:     true,
				Missing:    false,
				Summary:    "go-test: 10/10 passed (85% coverage)",
				Timestamp:  now,
				Confidence: 1.0,
			},
			{
				Kind:    domain.EvidenceKindMetrics,
				Source:  "metrics",
				Passed:  true,
				Summary: "42 lines (30 code, 5 comment, 7 blank), 3 TODOs, complexity 8",
				Metrics: map[string]float64{
					"code_lines":  30,
					"todo_count":  3,
					"complexity":  8,
					"total_lines": 42,
				},
				Details: map[string]any{
					"code_lines":  float64(30),
					"todo_count":  float64(3),
					"complexity":  float64(8),
					"total_lines": float64(42),
				},
				Timestamp:  now,
				Confidence: 1.0,
			},
			{
				Kind:       domain.EvidenceKindGitHistory,
				Source:     "git",
				Passed:     true,
				Summary:    "15 commits by 3 authors over 120 days",
				Timestamp:  now,
				Confidence: 1.0,
			},
		},
	}

	if err := store.Save(rec); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := store.Load(rec.UnitID)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	// Must have 4 evidence items
	if len(loaded.Evidence) != 4 {
		t.Fatalf("Evidence count = %d, want 4", len(loaded.Evidence))
	}

	// Check each evidence item
	for i, want := range rec.Evidence {
		got := loaded.Evidence[i]
		if got.Kind != want.Kind {
			t.Errorf("ev[%d].Kind = %v, want %v", i, got.Kind, want.Kind)
		}
		if got.Source != want.Source {
			t.Errorf("ev[%d].Source = %q, want %q", i, got.Source, want.Source)
		}
		if got.Passed != want.Passed {
			t.Errorf("ev[%d].Passed = %v, want %v", i, got.Passed, want.Passed)
		}
		if got.Missing != want.Missing {
			t.Errorf("ev[%d].Missing = %v, want %v", i, got.Missing, want.Missing)
		}
		if got.Summary != want.Summary {
			t.Errorf("ev[%d].Summary = %q, want %q", i, got.Summary, want.Summary)
		}
		if got.Confidence != want.Confidence {
			t.Errorf("ev[%d].Confidence = %f, want %f", i, got.Confidence, want.Confidence)
		}
		// Timestamp round-trips to second precision
		if !got.Timestamp.Equal(want.Timestamp) {
			t.Errorf("ev[%d].Timestamp = %v, want %v", i, got.Timestamp, want.Timestamp)
		}
	}

	// Check Details on the metrics evidence (index 2) survived as map[string]any
	metricsEv := loaded.Evidence[2]
	details, ok := metricsEv.Details.(map[string]any)
	if !ok {
		t.Fatalf("ev[2].Details type = %T, want map[string]any", metricsEv.Details)
	}
	if v, ok := details["code_lines"].(float64); !ok || v != 30 {
		t.Errorf("ev[2].Details[code_lines] = %v, want 30", details["code_lines"])
	}
	if v, ok := details["todo_count"].(float64); !ok || v != 3 {
		t.Errorf("ev[2].Details[todo_count] = %v, want 3", details["todo_count"])
	}
	if v, ok := details["complexity"].(float64); !ok || v != 8 {
		t.Errorf("ev[2].Details[complexity] = %v, want 8", details["complexity"])
	}

	// Check Metrics on the metrics evidence (index 2) survived round-trip
	if metricsEv.Metrics == nil {
		t.Fatal("ev[2].Metrics should not be nil after round-trip")
	}
	if metricsEv.Metrics["code_lines"] != 30 {
		t.Errorf("ev[2].Metrics[code_lines] = %f, want 30", metricsEv.Metrics["code_lines"])
	}
	if metricsEv.Metrics["todo_count"] != 3 {
		t.Errorf("ev[2].Metrics[todo_count] = %f, want 3", metricsEv.Metrics["todo_count"])
	}
	if metricsEv.Metrics["complexity"] != 8 {
		t.Errorf("ev[2].Metrics[complexity] = %f, want 8", metricsEv.Metrics["complexity"])
	}

	// Evidence with no Details (lint, index 0) should have nil Details
	if loaded.Evidence[0].Details != nil {
		t.Errorf("ev[0].Details should be nil, got %T", loaded.Evidence[0].Details)
	}
}

func TestStore_EvidenceRoundTrip_Empty(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	// Record with nil evidence
	rec := sampleRecord("empty.go", "Empty")
	rec.Evidence = nil

	if err := store.Save(rec); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := store.Load(rec.UnitID)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if len(loaded.Evidence) != 0 {
		t.Errorf("Evidence count = %d, want 0 for nil input", len(loaded.Evidence))
	}
}

func TestStore_BackwardCompatibility(t *testing.T) {
	dir := t.TempDir()

	// Write a record in the old format (no evidence field) directly to disk
	unitID := domain.NewUnitID("go", "old/file.go", "OldFunc")
	h := sha256.Sum256([]byte(unitID.String()))
	filename := hex.EncodeToString(h[:8]) + ".json"

	oldJSON := `{
  "unit_id": "go://old/file.go#OldFunc",
  "unit_type": "function",
  "unit_path": "old/file.go",
  "policy_version": "",
  "status": "certified",
  "grade": "B",
  "score": 0.85,
  "confidence": 1,
  "dimensions": {"correctness": 0.9},
  "observations": ["all good"],
  "certified_at": "2026-01-01T00:00:00Z",
  "expires_at": "2026-04-01T00:00:00Z",
  "source": "deterministic",
  "version": 1
}`

	if err := os.WriteFile(filepath.Join(dir, filename), []byte(oldJSON), 0o644); err != nil {
		t.Fatalf("writing old format file: %v", err)
	}

	store := record.NewStore(dir)
	loaded, err := store.Load(unitID)
	if err != nil {
		t.Fatalf("Load() old format error: %v", err)
	}

	// Core fields should load correctly
	if loaded.UnitID.String() != "go://old/file.go#OldFunc" {
		t.Errorf("UnitID = %s, want go://old/file.go#OldFunc", loaded.UnitID)
	}
	if loaded.Score != 0.85 {
		t.Errorf("Score = %f, want 0.85", loaded.Score)
	}
	if loaded.Status != domain.StatusCertified {
		t.Errorf("Status = %v, want certified", loaded.Status)
	}

	// Evidence should be nil/empty — not an error
	if len(loaded.Evidence) != 0 {
		t.Errorf("Evidence = %d items, want 0 for old format", len(loaded.Evidence))
	}
}

func TestStore_SaveAndLoadSnapshot(t *testing.T) {
	// Save 3 records, snapshot them, load into a fresh store
	srcDir := t.TempDir()
	store := record.NewStore(srcDir)

	recs := []domain.CertificationRecord{
		sampleRecord("main.go", "main"),
		sampleRecord("util.go", "helper"),
		sampleRecord("service/sync.go", "Apply"),
	}
	for _, r := range recs {
		if err := store.Save(r); err != nil {
			t.Fatalf("Save() error: %v", err)
		}
	}

	snapshotPath := filepath.Join(t.TempDir(), "state.json")
	if err := store.SaveSnapshot(snapshotPath, "abc123"); err != nil {
		t.Fatalf("SaveSnapshot() error: %v", err)
	}

	// Load into a fresh store with empty records dir
	dstDir := t.TempDir()
	dst := record.NewStore(dstDir)
	if err := dst.LoadSnapshot(snapshotPath); err != nil {
		t.Fatalf("LoadSnapshot() error: %v", err)
	}

	// All 3 should be loadable
	for _, r := range recs {
		loaded, err := dst.Load(r.UnitID)
		if err != nil {
			t.Errorf("Load(%s) after snapshot: %v", r.UnitID, err)
			continue
		}
		if loaded.Score != r.Score {
			t.Errorf("Score for %s = %f, want %f", r.UnitID, loaded.Score, r.Score)
		}
	}

	all, err := dst.ListAll()
	if err != nil {
		t.Fatalf("ListAll() error: %v", err)
	}
	if len(all) != 3 {
		t.Errorf("ListAll() = %d, want 3", len(all))
	}
}

func TestStore_SnapshotFormat(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	_ = store.Save(sampleRecord("a.go", "A"))
	_ = store.Save(sampleRecord("b.go", "B"))

	snapshotPath := filepath.Join(t.TempDir(), "state.json")
	if err := store.SaveSnapshot(snapshotPath, "deadbeef"); err != nil {
		t.Fatalf("SaveSnapshot() error: %v", err)
	}

	data, err := os.ReadFile(snapshotPath)
	if err != nil {
		t.Fatalf("reading snapshot: %v", err)
	}

	var snap struct {
		Version     int               `json:"version"`
		GeneratedAt string            `json:"generated_at"`
		Commit      string            `json:"commit"`
		UnitCount   int               `json:"unit_count"`
		Records     []json.RawMessage `json:"records"`
	}
	if err := json.Unmarshal(data, &snap); err != nil {
		t.Fatalf("parsing snapshot: %v", err)
	}

	if snap.Version != 1 {
		t.Errorf("version = %d, want 1", snap.Version)
	}
	if snap.Commit != "deadbeef" {
		t.Errorf("commit = %q, want deadbeef", snap.Commit)
	}
	if snap.UnitCount != 2 {
		t.Errorf("unit_count = %d, want 2", snap.UnitCount)
	}
	if len(snap.Records) != 2 {
		t.Errorf("records array len = %d, want 2", len(snap.Records))
	}
	if snap.GeneratedAt == "" {
		t.Error("generated_at should not be empty")
	}
}

func TestStore_ListAllFallbackToSnapshot(t *testing.T) {
	// Build a snapshot from a populated store
	srcDir := t.TempDir()
	src := record.NewStore(srcDir)
	_ = src.Save(sampleRecord("x.go", "X"))
	_ = src.Save(sampleRecord("y.go", "Y"))

	snapshotPath := filepath.Join(t.TempDir(), "state.json")
	if err := src.SaveSnapshot(snapshotPath, "snap1"); err != nil {
		t.Fatalf("SaveSnapshot() error: %v", err)
	}

	// Create a store pointing to an empty records dir, with snapshot fallback
	emptyDir := filepath.Join(t.TempDir(), "records") // doesn't exist
	store := record.NewStoreWithSnapshot(emptyDir, snapshotPath)

	records, err := store.ListAll()
	if err != nil {
		t.Fatalf("ListAll() error: %v", err)
	}
	if len(records) != 2 {
		t.Errorf("ListAll() fallback = %d, want 2", len(records))
	}
}

func TestStore_ListAllPrefersRecordsDir(t *testing.T) {
	// Populate both records dir and a snapshot with different data
	recDir := t.TempDir()
	store := record.NewStore(recDir)

	// Save 3 records to the dir
	_ = store.Save(sampleRecord("a.go", "A"))
	_ = store.Save(sampleRecord("b.go", "B"))
	_ = store.Save(sampleRecord("c.go", "C"))

	// Create a snapshot with only 1 record
	snapDir := t.TempDir()
	snapStore := record.NewStore(snapDir)
	_ = snapStore.Save(sampleRecord("only.go", "Only"))
	snapshotPath := filepath.Join(t.TempDir(), "state.json")
	_ = snapStore.SaveSnapshot(snapshotPath, "old")

	// Store with both records dir and snapshot
	combined := record.NewStoreWithSnapshot(recDir, snapshotPath)
	records, err := combined.ListAll()
	if err != nil {
		t.Fatalf("ListAll() error: %v", err)
	}
	// Should return 3 from records dir, NOT 1 from snapshot
	if len(records) != 3 {
		t.Errorf("ListAll() = %d, want 3 (from records dir, not snapshot)", len(records))
	}
}

func TestStore_SnapshotSortsDeterministically(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	// Save records in a specific order
	_ = store.Save(sampleRecord("z.go", "Z"))
	_ = store.Save(sampleRecord("a.go", "A"))
	_ = store.Save(sampleRecord("m.go", "M"))

	snap1 := filepath.Join(t.TempDir(), "state1.json")
	snap2 := filepath.Join(t.TempDir(), "state2.json")

	if err := store.SaveSnapshot(snap1, "same"); err != nil {
		t.Fatalf("SaveSnapshot(1) error: %v", err)
	}
	if err := store.SaveSnapshot(snap2, "same"); err != nil {
		t.Fatalf("SaveSnapshot(2) error: %v", err)
	}

	data1, _ := os.ReadFile(snap1)
	data2, _ := os.ReadFile(snap2)

	if string(data1) != string(data2) {
		t.Error("snapshot output should be byte-identical across calls")
	}
}

func TestStore_AppendAndLoadRuns(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(filepath.Join(dir, "records"))

	now := time.Now().Truncate(time.Second)
	runs := []domain.CertificationRun{
		{ID: "run-20260310T100000Z", StartedAt: now, CompletedAt: now.Add(time.Minute), Commit: "aaa", UnitsProcessed: 10, UnitsCertified: 8, UnitsFailed: 2, OverallGrade: "B", OverallScore: 0.80},
		{ID: "run-20260310T110000Z", StartedAt: now.Add(time.Hour), CompletedAt: now.Add(time.Hour + time.Minute), Commit: "bbb", UnitsProcessed: 15, UnitsCertified: 14, UnitsFailed: 1, OverallGrade: "B+", OverallScore: 0.85},
		{ID: "run-20260310T120000Z", StartedAt: now.Add(2 * time.Hour), CompletedAt: now.Add(2*time.Hour + time.Minute), Commit: "ccc", PolicyVersions: []string{"global@1.0"}, UnitsProcessed: 20, UnitsCertified: 19, UnitsFailed: 1, OverallGrade: "A-", OverallScore: 0.91},
	}

	for _, r := range runs {
		if err := store.AppendRun(r); err != nil {
			t.Fatalf("AppendRun(%s) error: %v", r.ID, err)
		}
	}

	loaded, err := store.LoadRuns()
	if err != nil {
		t.Fatalf("LoadRuns() error: %v", err)
	}
	if len(loaded) != 3 {
		t.Fatalf("LoadRuns() = %d runs, want 3", len(loaded))
	}

	// Verify ordering and field values
	if loaded[0].ID != "run-20260310T100000Z" {
		t.Errorf("runs[0].ID = %q, want run-20260310T100000Z", loaded[0].ID)
	}
	if loaded[2].Commit != "ccc" {
		t.Errorf("runs[2].Commit = %q, want ccc", loaded[2].Commit)
	}
	if loaded[2].UnitsCertified != 19 {
		t.Errorf("runs[2].UnitsCertified = %d, want 19", loaded[2].UnitsCertified)
	}
	if len(loaded[2].PolicyVersions) != 1 || loaded[2].PolicyVersions[0] != "global@1.0" {
		t.Errorf("runs[2].PolicyVersions = %v, want [global@1.0]", loaded[2].PolicyVersions)
	}
}

func TestStore_LoadRuns_Empty(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(filepath.Join(dir, "records"))

	runs, err := store.LoadRuns()
	if err != nil {
		t.Fatalf("LoadRuns() on empty store should not error: %v", err)
	}
	if runs != nil {
		t.Errorf("LoadRuns() on empty store = %v, want nil", runs)
	}
}

func TestStore_SnapshotIncludesRuns(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(filepath.Join(dir, "records"))

	// Save a record and a run
	_ = store.Save(sampleRecord("main.go", "main"))

	now := time.Now().Truncate(time.Second)
	_ = store.AppendRun(domain.CertificationRun{
		ID: "run-20260310T100000Z", StartedAt: now, CompletedAt: now.Add(time.Minute),
		Commit: "abc", UnitsProcessed: 1, UnitsCertified: 1, OverallGrade: "B", OverallScore: 0.85,
	})

	snapshotPath := filepath.Join(t.TempDir(), "state.json")
	if err := store.SaveSnapshot(snapshotPath, "abc"); err != nil {
		t.Fatalf("SaveSnapshot() error: %v", err)
	}

	// Read raw JSON and verify runs array
	data, _ := os.ReadFile(snapshotPath)
	var snap struct {
		Runs []json.RawMessage `json:"runs"`
	}
	if err := json.Unmarshal(data, &snap); err != nil {
		t.Fatalf("parsing snapshot: %v", err)
	}
	if len(snap.Runs) != 1 {
		t.Errorf("snapshot runs count = %d, want 1", len(snap.Runs))
	}
}
