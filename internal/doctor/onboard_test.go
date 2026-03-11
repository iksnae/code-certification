package doctor

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildOnboardPlan_EmptyDir(t *testing.T) {
	root := t.TempDir()

	plan := BuildOnboardPlan(root)

	if plan == nil {
		t.Fatal("expected non-nil plan")
	}
	if len(plan.Steps) != 6 {
		t.Fatalf("expected 6 steps, got %d", len(plan.Steps))
	}

	// Step 1 should be ready (no deps)
	if plan.Steps[0].Status != OnboardReady {
		t.Errorf("step 1 status = %v, want ready", plan.Steps[0].Status)
	}

	// Steps 2-6 should be pending
	for i := 1; i < 6; i++ {
		if plan.Steps[i].Status != OnboardPending {
			t.Errorf("step %d status = %v, want pending", i+1, plan.Steps[i].Status)
		}
	}

	// Next step should be step 1
	next := plan.NextStep()
	if next == nil {
		t.Fatal("expected non-nil next step")
	}
	if next.Number != 1 {
		t.Errorf("next step = %d, want 1", next.Number)
	}

	if plan.AllDone() {
		t.Error("expected not all done")
	}
}

func TestBuildOnboardPlan_AfterInit(t *testing.T) {
	root := t.TempDir()
	certDir := filepath.Join(root, ".certification")
	os.MkdirAll(certDir, 0o755)
	os.WriteFile(filepath.Join(certDir, "config.yml"), []byte("mode: advisory"), 0o644)

	plan := BuildOnboardPlan(root)

	// Step 1 done
	if plan.Steps[0].Status != OnboardDone {
		t.Errorf("step 1 status = %v, want done", plan.Steps[0].Status)
	}

	// Step 2 should be ready
	if plan.Steps[1].Status != OnboardReady {
		t.Errorf("step 2 status = %v, want ready", plan.Steps[1].Status)
	}

	// Next step should be step 2
	next := plan.NextStep()
	if next.Number != 2 {
		t.Errorf("next step = %d, want 2", next.Number)
	}
}

func TestBuildOnboardPlan_AfterScan(t *testing.T) {
	root := t.TempDir()
	certDir := filepath.Join(root, ".certification")
	os.MkdirAll(certDir, 0o755)
	os.WriteFile(filepath.Join(certDir, "config.yml"), []byte("mode: advisory"), 0o644)
	os.WriteFile(filepath.Join(certDir, "index.json"), []byte("[]"), 0o644)

	plan := BuildOnboardPlan(root)

	if plan.Steps[0].Status != OnboardDone {
		t.Errorf("step 1 status = %v, want done", plan.Steps[0].Status)
	}
	if plan.Steps[1].Status != OnboardDone {
		t.Errorf("step 2 status = %v, want done", plan.Steps[1].Status)
	}
	if plan.Steps[2].Status != OnboardReady {
		t.Errorf("step 3 status = %v, want ready", plan.Steps[2].Status)
	}

	next := plan.NextStep()
	if next.Number != 3 {
		t.Errorf("next step = %d, want 3", next.Number)
	}
}

func TestBuildOnboardPlan_AfterCertify(t *testing.T) {
	root := t.TempDir()
	certDir := filepath.Join(root, ".certification")
	os.MkdirAll(filepath.Join(certDir, "records"), 0o755)
	os.WriteFile(filepath.Join(certDir, "config.yml"), []byte("mode: advisory"), 0o644)
	os.WriteFile(filepath.Join(certDir, "index.json"), []byte("[]"), 0o644)
	os.WriteFile(filepath.Join(certDir, "records", "abc.json"), []byte("{}"), 0o644)

	plan := BuildOnboardPlan(root)

	if plan.Steps[2].Status != OnboardDone {
		t.Errorf("step 3 status = %v, want done", plan.Steps[2].Status)
	}
	if plan.Steps[3].Status != OnboardReady {
		t.Errorf("step 4 status = %v, want ready", plan.Steps[3].Status)
	}
	// Architect review (step 5) should also be ready since step 3 is done
	if plan.Steps[4].Status != OnboardReady {
		t.Errorf("step 5 (architect) status = %v, want ready", plan.Steps[4].Status)
	}
}

func TestBuildOnboardPlan_FullyDone(t *testing.T) {
	root := t.TempDir()
	certDir := filepath.Join(root, ".certification")
	os.MkdirAll(filepath.Join(certDir, "records"), 0o755)
	os.WriteFile(filepath.Join(certDir, "config.yml"), []byte("mode: advisory"), 0o644)
	os.WriteFile(filepath.Join(certDir, "index.json"), []byte("[]"), 0o644)
	os.WriteFile(filepath.Join(certDir, "records", "abc.json"), []byte("{}"), 0o644)
	os.WriteFile(filepath.Join(certDir, "REPORT_CARD.md"), []byte("# Report"), 0o644)
	os.WriteFile(filepath.Join(certDir, "ARCHITECT_REVIEW.md"), []byte("# Architect"), 0o644)
	os.WriteFile(filepath.Join(certDir, "badge.json"), []byte("{}"), 0o644)

	plan := BuildOnboardPlan(root)

	if !plan.AllDone() {
		for _, s := range plan.Steps {
			if s.Status != OnboardDone {
				t.Errorf("step %d (%s) not done: %v", s.Number, s.Title, s.Status)
			}
		}
	}

	next := plan.NextStep()
	if next != nil {
		t.Errorf("expected nil next step when all done, got step %d", next.Number)
	}
}

func TestFormatOnboardPlan_Structure(t *testing.T) {
	root := t.TempDir()
	plan := BuildOnboardPlan(root)

	output := FormatOnboardPlan(plan)

	required := []string{
		"Onboarding",
		"Initialize Certify",
		"Discover Code Units",
		"Run Certification",
		"Generate Report Card",
		"Architect Review",
		"Add Badge",
		"certify init",
		"Next step",
	}
	for _, s := range required {
		if !strings.Contains(output, s) {
			t.Errorf("output missing %q", s)
		}
	}
}

func TestFormatOnboardPlan_AllDone(t *testing.T) {
	root := t.TempDir()
	certDir := filepath.Join(root, ".certification")
	os.MkdirAll(filepath.Join(certDir, "records"), 0o755)
	os.WriteFile(filepath.Join(certDir, "config.yml"), []byte("mode: advisory"), 0o644)
	os.WriteFile(filepath.Join(certDir, "index.json"), []byte("[]"), 0o644)
	os.WriteFile(filepath.Join(certDir, "records", "abc.json"), []byte("{}"), 0o644)
	os.WriteFile(filepath.Join(certDir, "REPORT_CARD.md"), []byte("# Report"), 0o644)
	os.WriteFile(filepath.Join(certDir, "ARCHITECT_REVIEW.md"), []byte("# Architect"), 0o644)
	os.WriteFile(filepath.Join(certDir, "badge.json"), []byte("{}"), 0o644)

	plan := BuildOnboardPlan(root)
	output := FormatOnboardPlan(plan)

	if !strings.Contains(output, "All done") {
		t.Error("expected 'All done' message when fully onboarded")
	}
}

func TestOnboardStatus_Emoji(t *testing.T) {
	if OnboardDone.Emoji() != "✅" {
		t.Errorf("done emoji = %q, want ✅", OnboardDone.Emoji())
	}
	if OnboardReady.Emoji() != "👉" {
		t.Errorf("ready emoji = %q, want 👉", OnboardReady.Emoji())
	}
	if OnboardPending.Emoji() != "⬜" {
		t.Errorf("pending emoji = %q, want ⬜", OnboardPending.Emoji())
	}
}
