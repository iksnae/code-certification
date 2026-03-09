package queue_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/iksnae/code-certification/internal/queue"
)

func TestNewQueue_Empty(t *testing.T) {
	q := queue.New()
	if q.Len() != 0 {
		t.Errorf("new queue len = %d, want 0", q.Len())
	}
	stats := q.Stats()
	if stats.Total != 0 {
		t.Errorf("total = %d, want 0", stats.Total)
	}
}

func TestQueue_Enqueue(t *testing.T) {
	q := queue.New()
	q.Enqueue("go://main.go#main", "main.go")
	q.Enqueue("go://main.go#helper", "main.go")
	q.Enqueue("go://utils.go#parse", "utils.go")

	if q.Len() != 3 {
		t.Errorf("len = %d, want 3", q.Len())
	}
	stats := q.Stats()
	if stats.Pending != 3 {
		t.Errorf("pending = %d, want 3", stats.Pending)
	}
}

func TestQueue_Enqueue_NoDuplicates(t *testing.T) {
	q := queue.New()
	q.Enqueue("go://main.go#main", "main.go")
	q.Enqueue("go://main.go#main", "main.go") // duplicate

	if q.Len() != 1 {
		t.Errorf("len = %d, want 1 (no dupes)", q.Len())
	}
}

func TestQueue_Next_ReturnsPending(t *testing.T) {
	q := queue.New()
	q.Enqueue("go://a.go#a", "a.go")
	q.Enqueue("go://b.go#b", "b.go")

	item, ok := q.Next()
	if !ok {
		t.Fatal("expected item")
	}
	if item.UnitID != "go://a.go#a" {
		t.Errorf("unit = %q, want go://a.go#a", item.UnitID)
	}
	if item.Status != queue.StatusInProgress {
		t.Errorf("status = %q, want in_progress", item.Status)
	}
}

func TestQueue_Next_SkipsCompleted(t *testing.T) {
	q := queue.New()
	q.Enqueue("go://a.go#a", "a.go")
	q.Enqueue("go://b.go#b", "b.go")

	item, _ := q.Next()
	q.Complete(item.UnitID, "")

	item2, ok := q.Next()
	if !ok {
		t.Fatal("expected second item")
	}
	if item2.UnitID != "go://b.go#b" {
		t.Errorf("unit = %q, want go://b.go#b", item2.UnitID)
	}
}

func TestQueue_Next_RetriesFailed(t *testing.T) {
	q := queue.New()
	q.Enqueue("go://a.go#a", "a.go")

	item, _ := q.Next()
	q.Fail(item.UnitID, "rate limited")

	// Failed items should be retryable
	item2, ok := q.Next()
	if !ok {
		t.Fatal("failed item should be retryable")
	}
	if item2.UnitID != "go://a.go#a" {
		t.Errorf("should retry failed item")
	}
	if item2.Retries != 1 {
		t.Errorf("retries = %d, want 1", item2.Retries)
	}
}

func TestQueue_Next_ExhaustsMaxRetries(t *testing.T) {
	q := queue.New()
	q.Enqueue("go://a.go#a", "a.go")

	// Fail 3 times (default max retries)
	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Fatalf("should get item on retry %d", i)
		}
		q.Fail(item.UnitID, "error")
	}

	// Should now be exhausted
	_, ok := q.Next()
	if ok {
		t.Error("should be exhausted after max retries")
	}
}

func TestQueue_Next_Empty(t *testing.T) {
	q := queue.New()
	_, ok := q.Next()
	if ok {
		t.Error("empty queue should return !ok")
	}
}

func TestQueue_Complete(t *testing.T) {
	q := queue.New()
	q.Enqueue("go://a.go#a", "a.go")

	item, _ := q.Next()
	q.Complete(item.UnitID, "mistralai/mistral-nemo")

	stats := q.Stats()
	if stats.Completed != 1 {
		t.Errorf("completed = %d, want 1", stats.Completed)
	}
	if stats.Pending != 0 {
		t.Errorf("pending = %d, want 0", stats.Pending)
	}
}

func TestQueue_Skip(t *testing.T) {
	q := queue.New()
	q.Enqueue("go://a.go#a", "a.go")

	item, _ := q.Next()
	q.Skip(item.UnitID, "prescreen: no review needed")

	stats := q.Stats()
	if stats.Skipped != 1 {
		t.Errorf("skipped = %d, want 1", stats.Skipped)
	}
}

func TestQueue_SaveLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "queue.json")

	q := queue.New()
	q.Enqueue("go://a.go#a", "a.go")
	q.Enqueue("go://b.go#b", "b.go")

	item, _ := q.Next()
	q.Complete(item.UnitID, "model-a")

	if err := q.Save(path); err != nil {
		t.Fatalf("save: %v", err)
	}

	q2, err := queue.Load(path)
	if err != nil {
		t.Fatalf("load: %v", err)
	}

	if q2.Len() != 2 {
		t.Errorf("loaded len = %d, want 2", q2.Len())
	}

	stats := q2.Stats()
	if stats.Completed != 1 {
		t.Errorf("loaded completed = %d, want 1", stats.Completed)
	}
	if stats.Pending != 1 {
		t.Errorf("loaded pending = %d, want 1", stats.Pending)
	}
}

func TestQueue_Load_NotFound(t *testing.T) {
	_, err := queue.Load("/nonexistent/queue.json")
	if err == nil {
		t.Error("should error on missing file")
	}
}

func TestQueue_SaveAfterEachItem(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "queue.json")

	q := queue.New()
	q.Enqueue("go://a.go#a", "a.go")
	q.Enqueue("go://b.go#b", "b.go")
	q.Enqueue("go://c.go#c", "c.go")

	// Process one, save
	item, _ := q.Next()
	q.Complete(item.UnitID, "")
	q.Save(path)

	// Simulate crash — reload
	q2, _ := queue.Load(path)
	stats := q2.Stats()
	if stats.Completed != 1 || stats.Pending != 2 {
		t.Errorf("after reload: completed=%d pending=%d, want 1/2", stats.Completed, stats.Pending)
	}

	// Process another
	item2, ok := q2.Next()
	if !ok {
		t.Fatal("should have pending items")
	}
	if item2.UnitID == item.UnitID {
		t.Error("should not re-process completed item")
	}
}

func TestQueue_Stats(t *testing.T) {
	q := queue.New()
	q.Enqueue("a", "a.go")
	q.Enqueue("b", "b.go")
	q.Enqueue("c", "c.go")
	q.Enqueue("d", "d.go")

	item, _ := q.Next()
	q.Complete(item.UnitID, "model-x")
	item, _ = q.Next()
	q.Skip(item.UnitID, "clean")
	item, _ = q.Next()
	q.Fail(item.UnitID, "429")

	stats := q.Stats()
	if stats.Total != 4 {
		t.Errorf("total = %d, want 4", stats.Total)
	}
	if stats.Completed != 1 {
		t.Errorf("completed = %d, want 1", stats.Completed)
	}
	if stats.Skipped != 1 {
		t.Errorf("skipped = %d, want 1", stats.Skipped)
	}
	// Failed once = retryable, so counts as pending not failed
	if stats.Pending != 2 {
		t.Errorf("pending = %d, want 2 (1 untouched + 1 retryable)", stats.Pending)
	}
	if stats.Failed != 0 {
		t.Errorf("failed = %d, want 0 (not yet exhausted)", stats.Failed)
	}
}

func TestQueue_Reset(t *testing.T) {
	q := queue.New()
	q.Enqueue("a", "a.go")
	item, _ := q.Next()
	q.Complete(item.UnitID, "")

	q.Reset()
	stats := q.Stats()
	if stats.Completed != 0 || stats.Pending != 1 {
		t.Errorf("after reset: completed=%d pending=%d, want 0/1", stats.Completed, stats.Pending)
	}
}

func TestQueue_BatchNext(t *testing.T) {
	q := queue.New()
	for i := 0; i < 10; i++ {
		q.Enqueue(string(rune('a'+i)), "f.go")
	}

	batch := q.BatchNext(3)
	if len(batch) != 3 {
		t.Errorf("batch size = %d, want 3", len(batch))
	}

	// All should be in_progress
	stats := q.Stats()
	if stats.InProgress != 3 {
		t.Errorf("in_progress = %d, want 3", stats.InProgress)
	}
}

func TestQueue_BatchNext_LessThanRequested(t *testing.T) {
	q := queue.New()
	q.Enqueue("a", "a.go")
	q.Enqueue("b", "b.go")

	batch := q.BatchNext(5)
	if len(batch) != 2 {
		t.Errorf("batch size = %d, want 2", len(batch))
	}
}

func TestQueue_PersistPath(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "queue.json")

	q := queue.New()
	q.Enqueue("a", "a.go")
	q.Save(path)

	_, err := os.Stat(path)
	if err != nil {
		t.Errorf("queue file should exist: %v", err)
	}
}
