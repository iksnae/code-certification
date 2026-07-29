package report

import (
	"encoding/json"
	"fmt"
)

// Badge is a shields.io-compatible endpoint badge.
// See https://shields.io/badges/endpoint-badge
type Badge struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color"`
	Style         string `json:"style,omitempty"`
	NamedLogo     string `json:"namedLogo,omitempty"`
	LogoColor     string `json:"logoColor,omitempty"`
}

// GenerateBadge creates a shields.io endpoint badge from a report card.
func GenerateBadge(c Card) Badge {
	return Badge{
		SchemaVersion: 1,
		Label:         "certification",
		Message:       badgeMessage(c),
		Color:         badgeColor(badgeGrade(c)),
		NamedLogo:     "checkmarx",
		LogoColor:     "white",
	}
}

// badgeMessage is what the public README badge says. It is the most widely
// read output the product has, so it must never state a rate that was not
// measured: with nothing analyzable, "0%" and "100%" are equally false claims
// about code the engine never opened.
func badgeMessage(c Card) string {
	if c.TotalUnits == 0 {
		return "no data"
	}
	if !c.PassRateKnown() {
		return fmt.Sprintf("not assessed · 0 of %d units analyzable", c.TotalUnits)
	}
	return fmt.Sprintf("%s · %.0f%% · %d units",
		c.OverallGrade, c.PassRate*100, c.TotalUnits)
}

// badgeGrade is the grade the badge colours itself by. A card with no
// analyzable units has no verdict to colour, so it falls through to the
// neutral gray rather than painting the README red or green off a grade
// derived from units that were never scored.
func badgeGrade(c Card) string {
	if c.TotalUnits > 0 && !c.PassRateKnown() {
		return ""
	}
	return c.OverallGrade
}

// badgeColor maps grade to brand-consistent colors.
// Brand palette: Certified #2E8B57, Observations #E0A100,
// Probationary #F59E0B, Expired #9CA3AF, Decertified #DC2626.
func badgeColor(grade string) string {
	switch grade {
	case "A":
		return "2E8B57" // certified green
	case "A-":
		return "3DA06A" // certified green (lighter)
	case "B+":
		return "4A6B82" // steel blue (brand primary)
	case "B":
		return "4A6B82" // steel blue
	case "C":
		return "E0A100" // observations amber
	case "D":
		return "F59E0B" // probationary warning
	case "F":
		return "DC2626" // decertified red
	default:
		return "9CA3AF" // expired gray
	}
}

// FormatBadgeJSON marshals the badge to pretty JSON.
func FormatBadgeJSON(b Badge) ([]byte, error) {
	return json.MarshalIndent(b, "", "  ")
}

// BadgeMarkdown returns a markdown snippet to embed in a README.
// repo is "owner/repo", branch is typically "main".
func BadgeMarkdown(repo, branch string) string {
	badgeURL := fmt.Sprintf(
		"https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/%s/%s/.certification/badge.json",
		repo, branch,
	)
	reportURL := fmt.Sprintf(
		"https://github.com/%s/blob/%s/.certification/REPORT_CARD.md",
		repo, branch,
	)
	return fmt.Sprintf("[![Certification](%s)](%s)", badgeURL, reportURL)
}
