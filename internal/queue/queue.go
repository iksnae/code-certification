// Package queue provides a persistent work queue for incremental certification.
// Items are persisted to disk after each state change so progress survives
// crashes, rate limits, and multi-run processing.
package queue

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// ItemStatus represents the processing state of a queue item.
type ItemStatus string

const (
	StatusPending    ItemStatus = "pending"
	StatusInProgress ItemStatus = "in_progress"
	StatusCompleted  ItemStatus = "completed"
	StatusSkipped    ItemStatus = "skipped"
	StatusFailed     ItemStatus = "failed"
)

const DefaultMaxRetries = 3

// Item is a single unit of work in the queue.
type Item struct {
	UnitID    string     `json:"unit_id"`
	FilePath  string     `json:"file_path"`
	Status    ItemStatus `json:"status"`
	Model     string     `json:"model,omitempty"`  // model that processed this item
	Error     string     `json:"error,omitempty"`  // last error message
	Reason    string     `json:"reason,omitempty"` // skip/complete reason
	Retries   int        `json:"retries"`          // number of failed attempts
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// Stats holds queue statistics.
type Stats struct {
	Total      int `json:"total"`
	Pending    int `json:"pending"`
	InProgress int `json:"in_progress"`
	Completed  int `json:"completed"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

// Queue is a persistent work queue for certification processing.
type Queue struct {
	items      []*Item
	index      map[string]int // unitID → items index
	maxRetries int
}

// New creates an empty queue.
func New() *Queue {
	return &Queue{
		index:      make(map[string]int),
		maxRetries: DefaultMaxRetries,
	}
}

// Len returns the total number of items.
func (q *Queue) Len() int {
	return len(q.items)
}

// Enqueue adds a unit to the queue if not already present.
func (q *Queue) Enqueue(unitID, filePath string) {
	if _, exists := q.index[unitID]; exists {
		return
	}
	now := time.Now()
	item := &Item{
		UnitID:    unitID,
		FilePath:  filePath,
		Status:    StatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}
	q.index[unitID] = len(q.items)
	q.items = append(q.items, item)
}

// Next returns the next pending or retryable item and marks it in_progress.
// Returns (item, true) if found, (zero, false) if queue is exhausted.
func (q *Queue) Next() (Item, bool) {
	for _, item := range q.items {
		if item.Status == StatusPending {
			item.Status = StatusInProgress
			item.UpdatedAt = time.Now()
			return *item, true
		}
		if item.Status == StatusFailed && item.Retries < q.maxRetries {
			item.Status = StatusInProgress
			item.UpdatedAt = time.Now()
			return *item, true
		}
	}
	return Item{}, false
}

// BatchNext returns up to n pending items, all marked in_progress.
func (q *Queue) BatchNext(n int) []Item {
	var batch []Item
	for _, item := range q.items {
		if len(batch) >= n {
			break
		}
		if item.Status == StatusPending || (item.Status == StatusFailed && item.Retries < q.maxRetries) {
			item.Status = StatusInProgress
			item.UpdatedAt = time.Now()
			batch = append(batch, *item)
		}
	}
	return batch
}

// Complete marks an item as successfully processed.
func (q *Queue) Complete(unitID, model string) {
	if idx, ok := q.index[unitID]; ok {
		q.items[idx].Status = StatusCompleted
		q.items[idx].Model = model
		q.items[idx].UpdatedAt = time.Now()
	}
}

// Skip marks an item as skipped (e.g., prescreen said no review needed).
func (q *Queue) Skip(unitID, reason string) {
	if idx, ok := q.index[unitID]; ok {
		q.items[idx].Status = StatusSkipped
		q.items[idx].Reason = reason
		q.items[idx].UpdatedAt = time.Now()
	}
}

// Fail marks an item as failed with an error message.
func (q *Queue) Fail(unitID, errMsg string) {
	if idx, ok := q.index[unitID]; ok {
		q.items[idx].Status = StatusFailed
		q.items[idx].Error = errMsg
		q.items[idx].Retries++
		q.items[idx].UpdatedAt = time.Now()
	}
}

// Stats returns current queue statistics.
func (q *Queue) Stats() Stats {
	var s Stats
	s.Total = len(q.items)
	for _, item := range q.items {
		switch item.Status {
		case StatusPending:
			s.Pending++
		case StatusInProgress:
			s.InProgress++
		case StatusCompleted:
			s.Completed++
		case StatusSkipped:
			s.Skipped++
		case StatusFailed:
			if item.Retries >= q.maxRetries {
				s.Failed++ // exhausted
			} else {
				s.Pending++ // retryable = still pending
			}
		}
	}
	return s
}

// Reset sets all items back to pending (for re-processing).
func (q *Queue) Reset() {
	for _, item := range q.items {
		item.Status = StatusPending
		item.Retries = 0
		item.Error = ""
		item.Model = ""
		item.Reason = ""
		item.UpdatedAt = time.Now()
	}
}

// --- Persistence ---

type persistedQueue struct {
	Items []*Item `json:"items"`
}

// Save writes the queue to a JSON file.
func (q *Queue) Save(path string) error {
	data, err := json.MarshalIndent(persistedQueue{Items: q.items}, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling queue: %w", err)
	}
	return os.WriteFile(path, data, 0o644)
}

// Load reads a queue from a JSON file.
func Load(path string) (*Queue, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading queue: %w", err)
	}

	var pq persistedQueue
	if err := json.Unmarshal(data, &pq); err != nil {
		return nil, fmt.Errorf("parsing queue: %w", err)
	}

	q := &Queue{
		items:      pq.Items,
		index:      make(map[string]int),
		maxRetries: DefaultMaxRetries,
	}
	for i, item := range q.items {
		q.index[item.UnitID] = i
		// Reset in_progress items back to pending (interrupted run)
		if item.Status == StatusInProgress {
			item.Status = StatusPending
		}
	}
	return q, nil
}
