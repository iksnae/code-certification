// Package language_tiers provides tier classification for programming languages
// based on the depth of analysis support available.
package language_tiers

// Tier represents the level of analysis support for a language.
type Tier int

const (
	// TierNone means the language is explicitly not supported.
	TierNone Tier = iota
	// TierGeneric is the baseline tier for languages without specialized analysis.
	TierGeneric
	// TierLint means basic linting support is available.
	TierLint
	// TierParse means syntactic parsing support is available.
	TierParse
	// TierFull is the tier for languages with full structural analysis support.
	TierFull
)

// String returns the human-readable name of the tier.
func (t Tier) String() string {
	switch t {
	case TierNone:
		return "None"
	case TierGeneric:
		return "Generic"
	case TierLint:
		return "Lint"
	case TierParse:
		return "Parse"
	case TierFull:
		return "Full"
	default:
		return "Unknown"
	}
}

// languageTiers maps known language IDs to their analysis capability tier.
// This registry should be updated when new language support is added to discovery.
var languageTiers = map[string]Tier{
	// TierFull languages: full structural analysis available
	"go":   TierFull,
	"ts":   TierFull,
	"js":   TierFull,
	"py":   TierFull,
	"rs":   TierFull,

	// TierGeneric languages: baseline support only
	"rb":       TierGeneric,
	"java":     TierGeneric,
	"kt":       TierGeneric,
	"sh":       TierGeneric,
	"c":        TierGeneric,
	"cpp":      TierGeneric,
	"cs":       TierGeneric,
	"swift":    TierGeneric,
	"php":      TierGeneric,
	"elixir":   TierGeneric,
	"erlang":   TierGeneric,
	"scala":    TierGeneric,
	"lua":      TierGeneric,
	"r":        TierGeneric,
	"zig":      TierGeneric,
	"sql":      TierGeneric,
	"proto":    TierGeneric,
}

// TierForLanguage returns the analysis tier for a given language identifier.
// Returns the mapped tier from the registry, or TierGeneric for unknown languages.
func TierForLanguage(lang string) Tier {
	if tier, ok := languageTiers[lang]; ok {
		return tier
	}
	return TierGeneric
}

// IsSupported reports whether a language has enough analysis support for its
// code to be scored at all. At TierGeneric and below no dimension evidence can
// be produced, so any score or grade derived from such a unit would be
// fabricated rather than measured.
//
// This is the single source of truth for the supported/unsupported split.
// Callers must ask it directly rather than inferring the answer from the shape
// of a scoring result — an absent or empty score map means "nothing scored",
// which is not the same question.
func IsSupported(lang string) bool {
	return TierForLanguage(lang) > TierGeneric
}
