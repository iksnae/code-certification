package discovery_test

import (
	"testing"

	"github.com/code-certification/certify/internal/discovery"
	"github.com/code-certification/certify/internal/domain"
)

func TestTSAdapter_Functions(t *testing.T) {
	adapter := discovery.NewTSAdapter()
	units, err := adapter.Scan(repoPath("ts-simple"))
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}

	ids := unitIDs(units)

	mustHave := []string{
		"ts://src/parser.ts#tokenizeDialogue",
		"ts://src/parser.ts#parseNode",
		"ts://src/utils.ts#formatDate",
		"ts://src/utils.ts#log",
	}
	for _, want := range mustHave {
		if !ids[want] {
			t.Errorf("missing unit: %s", want)
		}
	}
}

func TestTSAdapter_Classes(t *testing.T) {
	adapter := discovery.NewTSAdapter()
	units, err := adapter.Scan(repoPath("ts-simple"))
	if err != nil {
		t.Fatal(err)
	}

	ids := unitIDs(units)
	if !ids["ts://src/parser.ts#DialogueParser"] {
		t.Error("missing class: DialogueParser")
	}

	for _, u := range units {
		if u.ID.String() == "ts://src/parser.ts#DialogueParser" {
			if u.Type != domain.UnitTypeClass {
				t.Errorf("DialogueParser type = %v, want class", u.Type)
			}
		}
	}
}

func TestTSAdapter_Constants(t *testing.T) {
	adapter := discovery.NewTSAdapter()
	units, err := adapter.Scan(repoPath("ts-simple"))
	if err != nil {
		t.Fatal(err)
	}

	ids := unitIDs(units)
	if !ids["ts://src/parser.ts#MAX_TOKENS"] {
		t.Error("missing const: MAX_TOKENS")
	}
}

func TestTSAdapter_LanguageIsTS(t *testing.T) {
	adapter := discovery.NewTSAdapter()
	units, err := adapter.Scan(repoPath("ts-simple"))
	if err != nil {
		t.Fatal(err)
	}

	for _, u := range units {
		if u.ID.Language() != "ts" {
			t.Errorf("unit %s language = %q, want ts", u.ID, u.ID.Language())
		}
	}
}

func TestTSAdapter_StableIDs(t *testing.T) {
	adapter := discovery.NewTSAdapter()
	u1, _ := adapter.Scan(repoPath("ts-simple"))
	u2, _ := adapter.Scan(repoPath("ts-simple"))

	if len(u1) != len(u2) {
		t.Fatal("scan should be deterministic")
	}
	for i := range u1 {
		if u1[i].ID.String() != u2[i].ID.String() {
			t.Errorf("ID mismatch: %s != %s", u1[i].ID, u2[i].ID)
		}
	}
}
