package domain_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
)

func TestUnitType_String(t *testing.T) {
	tests := []struct {
		ut   domain.UnitType
		want string
	}{
		{domain.UnitTypeFile, "file"},
		{domain.UnitTypeFunction, "function"},
		{domain.UnitTypeMethod, "method"},
		{domain.UnitTypeClass, "class"},
		{domain.UnitTypeModule, "module"},
		{domain.UnitTypePackage, "package"},
	}
	for _, tt := range tests {
		if got := tt.ut.String(); got != tt.want {
			t.Errorf("UnitType(%d).String() = %q, want %q", tt.ut, got, tt.want)
		}
	}
}

func TestParseUnitType(t *testing.T) {
	tests := []struct {
		input string
		want  domain.UnitType
		ok    bool
	}{
		{"file", domain.UnitTypeFile, true},
		{"function", domain.UnitTypeFunction, true},
		{"method", domain.UnitTypeMethod, true},
		{"class", domain.UnitTypeClass, true},
		{"module", domain.UnitTypeModule, true},
		{"package", domain.UnitTypePackage, true},
		{"unknown", 0, false},
		{"", 0, false},
	}
	for _, tt := range tests {
		got, err := domain.ParseUnitType(tt.input)
		if tt.ok {
			if err != nil {
				t.Errorf("ParseUnitType(%q) unexpected error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("ParseUnitType(%q) = %v, want %v", tt.input, got, tt.want)
			}
		} else {
			if err == nil {
				t.Errorf("ParseUnitType(%q) expected error, got nil", tt.input)
			}
		}
	}
}

func TestNewUnitID(t *testing.T) {
	tests := []struct {
		name    string
		lang    string
		path    string
		symbol  string
		wantStr string
	}{
		{
			name:    "go function",
			lang:    "go",
			path:    "internal/service/sync.go",
			symbol:  "Apply",
			wantStr: "go://internal/service/sync.go#Apply",
		},
		{
			name:    "ts function",
			lang:    "ts",
			path:    "src/parser/tokenize.ts",
			symbol:  "tokenizeDialogue",
			wantStr: "ts://src/parser/tokenize.ts#tokenizeDialogue",
		},
		{
			name:    "file-level fallback",
			lang:    "",
			path:    "scripts/release.sh",
			symbol:  "",
			wantStr: "file://scripts/release.sh",
		},
		{
			name:    "go file-level no symbol",
			lang:    "go",
			path:    "main.go",
			symbol:  "",
			wantStr: "go://main.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := domain.NewUnitID(tt.lang, tt.path, tt.symbol)
			if got := id.String(); got != tt.wantStr {
				t.Errorf("NewUnitID(%q, %q, %q).String() = %q, want %q",
					tt.lang, tt.path, tt.symbol, got, tt.wantStr)
			}
		})
	}
}

func TestUnitID_Components(t *testing.T) {
	id := domain.NewUnitID("go", "internal/service/sync.go", "Apply")
	if id.Language() != "go" {
		t.Errorf("Language() = %q, want %q", id.Language(), "go")
	}
	if id.Path() != "internal/service/sync.go" {
		t.Errorf("Path() = %q, want %q", id.Path(), "internal/service/sync.go")
	}
	if id.Symbol() != "Apply" {
		t.Errorf("Symbol() = %q, want %q", id.Symbol(), "Apply")
	}
}

func TestParseUnitID(t *testing.T) {
	tests := []struct {
		input    string
		wantLang string
		wantPath string
		wantSym  string
		ok       bool
	}{
		{"go://internal/service/sync.go#Apply", "go", "internal/service/sync.go", "Apply", true},
		{"ts://src/parser/tokenize.ts#tokenizeDialogue", "ts", "src/parser/tokenize.ts", "tokenizeDialogue", true},
		{"file://scripts/release.sh", "file", "scripts/release.sh", "", true},
		{"go://main.go", "go", "main.go", "", true},
		{"", "", "", "", false},
		{"invalid-no-scheme", "", "", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			id, err := domain.ParseUnitID(tt.input)
			if tt.ok {
				if err != nil {
					t.Fatalf("ParseUnitID(%q) unexpected error: %v", tt.input, err)
				}
				if id.Language() != tt.wantLang {
					t.Errorf("Language() = %q, want %q", id.Language(), tt.wantLang)
				}
				if id.Path() != tt.wantPath {
					t.Errorf("Path() = %q, want %q", id.Path(), tt.wantPath)
				}
				if id.Symbol() != tt.wantSym {
					t.Errorf("Symbol() = %q, want %q", id.Symbol(), tt.wantSym)
				}
			} else {
				if err == nil {
					t.Errorf("ParseUnitID(%q) expected error, got nil", tt.input)
				}
			}
		})
	}
}

func TestUnitID_Equality(t *testing.T) {
	a := domain.NewUnitID("go", "main.go", "main")
	b := domain.NewUnitID("go", "main.go", "main")
	c := domain.NewUnitID("go", "main.go", "other")

	if a.String() != b.String() {
		t.Error("identical UnitIDs should have equal String()")
	}
	if a.String() == c.String() {
		t.Error("different UnitIDs should have different String()")
	}
}

func TestNewUnit(t *testing.T) {
	id := domain.NewUnitID("go", "internal/service/sync.go", "Apply")
	unit := domain.NewUnit(id, domain.UnitTypeFunction)

	if unit.ID.String() != id.String() {
		t.Errorf("unit.ID = %v, want %v", unit.ID, id)
	}
	if unit.Type != domain.UnitTypeFunction {
		t.Errorf("unit.Type = %v, want %v", unit.Type, domain.UnitTypeFunction)
	}
}

func TestNewUnit_FileLevel(t *testing.T) {
	id := domain.NewUnitID("", "scripts/release.sh", "")
	unit := domain.NewUnit(id, domain.UnitTypeFile)

	if unit.Type != domain.UnitTypeFile {
		t.Errorf("unit.Type = %v, want %v", unit.Type, domain.UnitTypeFile)
	}
	if unit.ID.Language() != "file" {
		t.Errorf("Language() = %q, want %q", unit.ID.Language(), "file")
	}
}
