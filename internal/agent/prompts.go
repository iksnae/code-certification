package agent

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
