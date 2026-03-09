package discovery_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/discovery"
	"github.com/iksnae/code-certification/internal/domain"
)

func TestGoAdapter_Functions(t *testing.T) {
	adapter := discovery.NewGoAdapter()
	units, err := adapter.Scan(repoPath("go-simple"))
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}

	// Should find functions: main, helper, Apply, Reset, Format
	// And types: Syncer
	ids := unitIDs(units)

	mustHave := []string{
		"go://main.go#main",
		"go://main.go#helper",
		"go://internal/service/sync.go#Apply",
		"go://internal/service/sync.go#Reset",
		"go://internal/service/sync.go#Format",
	}
	for _, want := range mustHave {
		if !ids[want] {
			t.Errorf("missing unit: %s", want)
		}
	}
}

func TestGoAdapter_Methods(t *testing.T) {
	adapter := discovery.NewGoAdapter()
	units, err := adapter.Scan(repoPath("go-simple"))
	if err != nil {
		t.Fatal(err)
	}

	// Apply and Reset should be methods
	for _, u := range units {
		id := u.ID.String()
		if id == "go://internal/service/sync.go#Apply" || id == "go://internal/service/sync.go#Reset" {
			if u.Type != domain.UnitTypeMethod {
				t.Errorf("%s type = %v, want method", id, u.Type)
			}
		}
	}
}

func TestGoAdapter_FreeStandingFunction(t *testing.T) {
	adapter := discovery.NewGoAdapter()
	units, err := adapter.Scan(repoPath("go-simple"))
	if err != nil {
		t.Fatal(err)
	}

	for _, u := range units {
		if u.ID.String() == "go://internal/service/sync.go#Format" {
			if u.Type != domain.UnitTypeFunction {
				t.Errorf("Format type = %v, want function", u.Type)
			}
			return
		}
	}
	t.Error("Format function not found")
}

func TestGoAdapter_Types(t *testing.T) {
	adapter := discovery.NewGoAdapter()
	units, err := adapter.Scan(repoPath("go-simple"))
	if err != nil {
		t.Fatal(err)
	}

	ids := unitIDs(units)
	if !ids["go://internal/service/sync.go#Syncer"] {
		t.Error("missing type: Syncer")
	}

	for _, u := range units {
		if u.ID.String() == "go://internal/service/sync.go#Syncer" {
			if u.Type != domain.UnitTypeClass {
				t.Errorf("Syncer type = %v, want class", u.Type)
			}
		}
	}
}

func TestGoAdapter_LanguageIsGo(t *testing.T) {
	adapter := discovery.NewGoAdapter()
	units, err := adapter.Scan(repoPath("go-simple"))
	if err != nil {
		t.Fatal(err)
	}

	for _, u := range units {
		if u.ID.Language() != "go" {
			t.Errorf("unit %s language = %q, want go", u.ID, u.ID.Language())
		}
	}
}

func TestGoAdapter_StableIDs(t *testing.T) {
	adapter := discovery.NewGoAdapter()
	u1, _ := adapter.Scan(repoPath("go-simple"))
	u2, _ := adapter.Scan(repoPath("go-simple"))

	if len(u1) != len(u2) {
		t.Fatal("scan should be deterministic")
	}
	for i := range u1 {
		if u1[i].ID.String() != u2[i].ID.String() {
			t.Errorf("ID mismatch: %s != %s", u1[i].ID, u2[i].ID)
		}
	}
}

func unitIDs(units []domain.Unit) map[string]bool {
	m := make(map[string]bool)
	for _, u := range units {
		m[u.ID.String()] = true
	}
	return m
}
