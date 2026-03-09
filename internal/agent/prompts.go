package agent

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// PromptTemplate is a loaded prompt template with variable substitution.
type PromptTemplate struct {
	raw      string
	filename string
}

// LoadPrompt loads a prompt template from a file.
func LoadPrompt(path string) (*PromptTemplate, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("loading prompt: %w", err)
	}

	return &PromptTemplate{
		raw:      string(data),
		filename: filepath.Base(path),
	}, nil
}

// Render substitutes variables into the template.
// Variables are specified as {{.Key}} in the template.
func (p *PromptTemplate) Render(vars map[string]string) (string, error) {
	result := p.raw
	for key, value := range vars {
		placeholder := "{{." + key + "}}"
		result = strings.ReplaceAll(result, placeholder, value)
	}
	return result, nil
}

// Version extracts the version from the filename (e.g., "prescreen.v1.md" → "v1").
func (p *PromptTemplate) Version() string {
	name := strings.TrimSuffix(p.filename, filepath.Ext(p.filename))
	parts := strings.Split(name, ".")
	if len(parts) >= 2 {
		return parts[len(parts)-1]
	}
	return "unknown"
}

// PromptRegistry loads and caches prompt templates by task type.
type PromptRegistry struct {
	dir   string
	cache map[TaskType]*PromptTemplate
	mu    sync.Mutex
}

// NewPromptRegistry creates a registry that loads templates from a directory.
func NewPromptRegistry(dir string) *PromptRegistry {
	return &PromptRegistry{
		dir:   dir,
		cache: make(map[TaskType]*PromptTemplate),
	}
}

// taskFileNames maps task types to their template file patterns.
var taskFileNames = map[TaskType]string{
	TaskPrescreen:   "prescreen",
	TaskReview:      "review",
	TaskScoring:     "scoring",
	TaskDecision:    "decision",
	TaskRemediation: "remediation",
}

// Get returns the prompt template for a task type, loading and caching on first access.
func (r *PromptRegistry) Get(task TaskType) (*PromptTemplate, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if tmpl, ok := r.cache[task]; ok {
		return tmpl, nil
	}

	base, ok := taskFileNames[task]
	if !ok {
		return nil, fmt.Errorf("unknown task type: %v", task)
	}

	// Find the latest version (e.g., prescreen.v1.md)
	pattern := filepath.Join(r.dir, base+".*.md")
	matches, err := filepath.Glob(pattern)
	if err != nil || len(matches) == 0 {
		return nil, fmt.Errorf("no template found for %s in %s", base, r.dir)
	}

	// Use the last match (highest version by lexical sort)
	path := matches[len(matches)-1]
	tmpl, err := LoadPrompt(path)
	if err != nil {
		return nil, err
	}

	r.cache[task] = tmpl
	return tmpl, nil
}
