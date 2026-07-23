package language_tiers

import "testing"

func TestTierForLanguage(t *testing.T) {
	tests := []struct {
		name string
		lang string
		want Tier
	}{
		// TierFull languages: full structural analysis support
		{name: "Go", lang: "go", want: TierFull},
		{name: "TypeScript", lang: "ts", want: TierFull},
		{name: "Python", lang: "py", want: TierFull},
		{name: "Rust", lang: "rs", want: TierFull},
		{name: "JavaScript", lang: "js", want: TierFull},

		// TierGeneric languages: baseline support
		{name: "Ruby", lang: "rb", want: TierGeneric},
		{name: "Java", lang: "java", want: TierGeneric},
		{name: "Kotlin", lang: "kt", want: TierGeneric},
		{name: "Shell", lang: "sh", want: TierGeneric},
		{name: "C", lang: "c", want: TierGeneric},
		{name: "C++", lang: "cpp", want: TierGeneric},
		{name: "C#", lang: "cs", want: TierGeneric},
		{name: "Swift", lang: "swift", want: TierGeneric},
		{name: "PHP", lang: "php", want: TierGeneric},
		{name: "Elixir", lang: "elixir", want: TierGeneric},
		{name: "Erlang", lang: "erlang", want: TierGeneric},
		{name: "Scala", lang: "scala", want: TierGeneric},
		{name: "Lua", lang: "lua", want: TierGeneric},
		{name: "R", lang: "r", want: TierGeneric},
		{name: "Zig", lang: "zig", want: TierGeneric},
		{name: "SQL", lang: "sql", want: TierGeneric},
		{name: "Protobuf", lang: "proto", want: TierGeneric},

		// Unknown languages fall back to Generic
		{name: "Unknown", lang: "unknown", want: TierGeneric},
		{name: "Empty string", lang: "", want: TierGeneric},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TierForLanguage(tt.lang)
			if got != tt.want {
				t.Errorf("TierForLanguage(%q) = %v (%d), want %v (%d)", tt.lang, got, got, tt.want, tt.want)
			}
		})
	}
}

func TestTierString(t *testing.T) {
	tests := []struct {
		tier Tier
		want string
	}{
		{TierNone, "None"},
		{TierGeneric, "Generic"},
		{TierLint, "Lint"},
		{TierParse, "Parse"},
		{TierFull, "Full"},
		{Tier(99), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := tt.tier.String()
			if got != tt.want {
				t.Errorf("Tier(%d).String() = %q, want %q", tt.tier, got, tt.want)
			}
		})
	}
}

func TestTierIotaValues(t *testing.T) {
	// Ensure iota values are correct and TierGeneric remains at value 1 for backward compatibility
	if TierNone != 0 {
		t.Errorf("TierNone should be 0, got %d", TierNone)
	}
	if TierGeneric != 1 {
		t.Errorf("TierGeneric should be 1, got %d", TierGeneric)
	}
	if TierLint != 2 {
		t.Errorf("TierLint should be 2, got %d", TierLint)
	}
	if TierParse != 3 {
		t.Errorf("TierParse should be 3, got %d", TierParse)
	}
	if TierFull != 4 {
		t.Errorf("TierFull should be 4, got %d", TierFull)
	}
}

// TestLanguageTiersRegistry verifies all 22 language IDs have entries in the registry.
// This test fails CI when a new language extension is added to discovery without a tier entry.
func TestLanguageTiersRegistry(t *testing.T) {
	// Expected languages that should be in the registry
	expectedLanguages := []string{
		"go", "ts", "js", "py", "rs", // TierFull
		"rb", "java", "kt", "sh", "c", "cpp", "cs",
		"swift", "php", "elixir", "erlang", "scala", "lua",
		"r", "zig", "sql", "proto",
	}

	for _, lang := range expectedLanguages {
		tier, exists := languageTiers[lang]
		if !exists {
			t.Errorf("Language %q is not registered in languageTiers map - add it when adding new language support", lang)
			continue
		}
		if tier == TierNone {
			t.Errorf("Language %q is registered with TierNone - explicitly set a meaningful tier", lang)
		}
	}

	// Count languages to ensure we have exactly 22
	if len(languageTiers) != 22 {
		t.Errorf("Expected exactly 22 languages in registry, got %d - update registry when adding/removing language support", len(languageTiers))
	}
}
