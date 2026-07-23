package language_tiers

import "testing"

func TestTierForLanguage(t *testing.T) {
	tests := []struct {
		name string
		lang string
		want Tier
	}{
		// Registered languages with full analysis support
		{name: "Go", lang: "go", want: TierFull},
		{name: "TypeScript", lang: "ts", want: TierFull},
		{name: "Python", lang: "py", want: TierFull},
		{name: "Rust", lang: "rs", want: TierFull},
		{name: "JavaScript", lang: "js", want: TierFull},

		// Unregistered languages fall back to Generic
		{name: "Ruby", lang: "ruby", want: TierGeneric},
		{name: "Swift", lang: "swift", want: TierGeneric},
		{name: "C", lang: "c", want: TierGeneric},
		{name: "Empty string", lang: "", want: TierGeneric},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TierForLanguage(tt.lang)
			if got != tt.want {
				t.Errorf("TierForLanguage(%q) = %d, want %d", tt.lang, got, tt.want)
			}
		})
	}
}
