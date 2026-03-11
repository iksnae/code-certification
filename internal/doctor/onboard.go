package doctor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// OnboardStep represents a single onboarding step.
type OnboardStep struct {
	Number      int
	Title       string
	Description string
	Command     string        // command to run (empty = informational)
	Status      OnboardStatus // computed from filesystem
	Detail      string        // additional context
}

// OnboardStatus represents whether a step is done.
type OnboardStatus int

const (
	OnboardPending OnboardStatus = iota
	OnboardDone
	OnboardReady // dependencies met, ready to run
)

func (s OnboardStatus) Emoji() string {
	switch s {
	case OnboardDone:
		return "✅"
	case OnboardReady:
		return "👉"
	case OnboardPending:
		return "⬜"
	}
	return "❓"
}

// OnboardPlan holds the full onboarding plan for a project.
type OnboardPlan struct {
	Steps []OnboardStep
	Root  string
}

// NextStep returns the first step that is Ready or Pending, or nil if all done.
func (p *OnboardPlan) NextStep() *OnboardStep {
	for i, s := range p.Steps {
		if s.Status == OnboardReady {
			return &p.Steps[i]
		}
	}
	for i, s := range p.Steps {
		if s.Status == OnboardPending {
			return &p.Steps[i]
		}
	}
	return nil
}

// AllDone returns true if every step is done.
func (p *OnboardPlan) AllDone() bool {
	for _, s := range p.Steps {
		if s.Status != OnboardDone {
			return false
		}
	}
	return true
}

// BuildOnboardPlan creates an onboarding plan by checking what's already done.
func BuildOnboardPlan(root string) *OnboardPlan {
	certDir := filepath.Join(root, ".certification")

	plan := &OnboardPlan{Root: root}

	// Step 1: Initialize
	step1 := OnboardStep{
		Number:      1,
		Title:       "Initialize Certify",
		Description: "Bootstrap .certification/ directory with config, policies, and CI workflows.",
		Command:     "certify init",
	}
	if _, err := os.Stat(filepath.Join(certDir, "config.yml")); err == nil {
		step1.Status = OnboardDone
		step1.Detail = ".certification/config.yml exists"
	} else {
		step1.Status = OnboardReady
		step1.Detail = "Run this in your repository root"
	}
	plan.Steps = append(plan.Steps, step1)

	// Step 2: Scan
	step2 := OnboardStep{
		Number:      2,
		Title:       "Discover Code Units",
		Description: "Scan the repository to find all functions, methods, types, and files.",
		Command:     "certify scan",
	}
	if _, err := os.Stat(filepath.Join(certDir, "index.json")); err == nil {
		step2.Status = OnboardDone
		step2.Detail = "index.json exists"
	} else if step1.Status == OnboardDone {
		step2.Status = OnboardReady
	} else {
		step2.Status = OnboardPending
		step2.Detail = "Complete step 1 first"
	}
	plan.Steps = append(plan.Steps, step2)

	// Step 3: Certify
	step3 := OnboardStep{
		Number:      3,
		Title:       "Run Certification",
		Description: "Evaluate every unit against policies, score across 9 dimensions, assign certification.",
		Command:     "certify certify",
	}
	recordsDir := filepath.Join(certDir, "records")
	if hasJSONFiles(recordsDir) {
		step3.Status = OnboardDone
		step3.Detail = "Certification records exist"
	} else if step2.Status == OnboardDone {
		step3.Status = OnboardReady
	} else {
		step3.Status = OnboardPending
		step3.Detail = "Complete step 2 first"
	}
	plan.Steps = append(plan.Steps, step3)

	// Step 4: Report
	step4 := OnboardStep{
		Number:      4,
		Title:       "Generate Report Card",
		Description: "Create your report card, badge, and per-unit report tree.",
		Command:     "certify report",
	}
	if _, err := os.Stat(filepath.Join(certDir, "REPORT_CARD.md")); err == nil {
		step4.Status = OnboardDone
		step4.Detail = "REPORT_CARD.md exists"
	} else if step3.Status == OnboardDone {
		step4.Status = OnboardReady
	} else {
		step4.Status = OnboardPending
		step4.Detail = "Complete step 3 first"
	}
	plan.Steps = append(plan.Steps, step4)

	// Step 5: Architect Review (optional)
	step5 := OnboardStep{
		Number:      5,
		Title:       "Architect Review (optional)",
		Description: "Run a 6-phase AI architectural analysis. Requires an AI provider.",
		Command:     "certify architect",
	}
	if _, err := os.Stat(filepath.Join(certDir, "ARCHITECT_REVIEW.md")); err == nil {
		step5.Status = OnboardDone
		step5.Detail = "ARCHITECT_REVIEW.md exists"
	} else if step3.Status == OnboardDone {
		step5.Status = OnboardReady
		step5.Detail = "Requires OPENROUTER_API_KEY or local model"
	} else {
		step5.Status = OnboardPending
		step5.Detail = "Complete step 3 first"
	}
	plan.Steps = append(plan.Steps, step5)

	// Step 6: Badge
	step6 := OnboardStep{
		Number:      6,
		Title:       "Add Badge to README",
		Description: "Get the markdown snippet for a live certification badge in your README.",
		Command:     "certify report --badge",
	}
	if _, err := os.Stat(filepath.Join(certDir, "badge.json")); err == nil {
		step6.Status = OnboardDone
		step6.Detail = "badge.json exists"
	} else if step4.Status == OnboardDone {
		step6.Status = OnboardReady
	} else {
		step6.Status = OnboardPending
	}
	plan.Steps = append(plan.Steps, step6)

	return plan
}

// FormatOnboardPlan renders the plan as a human-readable string.
func FormatOnboardPlan(plan *OnboardPlan) string {
	var b strings.Builder

	fmt.Fprintf(&b, "🚀 Certify Onboarding — %s\n\n", plan.Root)

	next := plan.NextStep()

	for _, s := range plan.Steps {
		fmt.Fprintf(&b, "  %s %d. %s\n", s.Status.Emoji(), s.Number, s.Title)
		if s.Status == OnboardReady && s == *next {
			fmt.Fprintf(&b, "     → %s\n", s.Command)
			if s.Detail != "" {
				fmt.Fprintf(&b, "       %s\n", s.Detail)
			}
		} else if s.Status == OnboardDone && s.Detail != "" {
			fmt.Fprintf(&b, "     %s\n", s.Detail)
		} else if s.Status == OnboardPending && s.Detail != "" {
			fmt.Fprintf(&b, "     %s\n", s.Detail)
		}
	}

	fmt.Fprintln(&b)
	if plan.AllDone() {
		fmt.Fprintf(&b, "  🎉 All done! Your project is fully onboarded.\n")
		fmt.Fprintf(&b, "     Run 'certify doctor' to check health at any time.\n")
	} else if next != nil {
		fmt.Fprintf(&b, "  Next step: %s\n", next.Command)
	}

	return b.String()
}

// hasJSONFiles checks if a directory contains any .json files.
func hasJSONFiles(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".json") && !strings.Contains(e.Name(), ".history.") {
			return true
		}
	}
	return false
}
