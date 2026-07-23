// Package language_tiers provides tier classification for programming languages
// based on the depth of analysis support available.
package language_tiers

import "github.com/iksnae/code-certification/internal/analysis"

// Tier represents the level of analysis support for a language.
type Tier int

const (
	// TierGeneric is the baseline tier for languages without specialized analysis.
	TierGeneric Tier = iota
	// TierFull is the tier for languages with full structural analysis support.
	TierFull
)

// TierForLanguage returns the analysis tier for a given language identifier.
// Languages registered with the analysis package return TierFull;
// unrecognized languages return TierGeneric.
func TierForLanguage(lang string) Tier {
	if analysis.ForLanguage(lang) != nil {
		return TierFull
	}
	return TierGeneric
}
