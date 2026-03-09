// Package domain defines the core types for the Code Certification System.
// These types have zero external dependencies and form the shared vocabulary
// used by all other packages.
package domain

import (
	"fmt"
	"strings"
)

// UnitType identifies the kind of certifiable code unit.
type UnitType int

const (
	UnitTypeFile     UnitType = iota // A whole file
	UnitTypeFunction                 // A standalone function
	UnitTypeMethod                   // A method on a type
	UnitTypeClass                    // A class or type definition
	UnitTypeModule                   // A module (e.g. ES module)
	UnitTypePackage                  // A package (e.g. Go package)
)

var unitTypeStrings = map[UnitType]string{
	UnitTypeFile:     "file",
	UnitTypeFunction: "function",
	UnitTypeMethod:   "method",
	UnitTypeClass:    "class",
	UnitTypeModule:   "module",
	UnitTypePackage:  "package",
}

var stringToUnitType map[string]UnitType

func init() {
	stringToUnitType = make(map[string]UnitType, len(unitTypeStrings))
	for k, v := range unitTypeStrings {
		stringToUnitType[v] = k
	}
}

// String returns the string representation of a UnitType.
func (ut UnitType) String() string {
	if s, ok := unitTypeStrings[ut]; ok {
		return s
	}
	return fmt.Sprintf("UnitType(%d)", ut)
}

// ParseUnitType converts a string to a UnitType.
func ParseUnitType(s string) (UnitType, error) {
	if ut, ok := stringToUnitType[s]; ok {
		return ut, nil
	}
	return 0, fmt.Errorf("unknown unit type: %q", s)
}

// UnitID is a stable identifier for a certifiable code unit.
// Format: <language>://<path>[#<symbol>]
// Examples:
//
//	go://internal/service/sync.go#Apply
//	ts://src/parser/tokenize.ts#tokenizeDialogue
//	file://scripts/release.sh
type UnitID struct {
	language string
	path     string
	symbol   string
}

// NewUnitID creates a new UnitID. If language is empty, "file" is used.
func NewUnitID(language, path, symbol string) UnitID {
	if language == "" {
		language = "file"
	}
	return UnitID{language: language, path: path, symbol: symbol}
}

// Language returns the language component of the ID.
func (id UnitID) Language() string { return id.language }

// Path returns the file path component of the ID.
func (id UnitID) Path() string { return id.path }

// Symbol returns the symbol component of the ID (may be empty).
func (id UnitID) Symbol() string { return id.symbol }

// String returns the canonical string representation of the UnitID.
func (id UnitID) String() string {
	if id.symbol != "" {
		return fmt.Sprintf("%s://%s#%s", id.language, id.path, id.symbol)
	}
	return fmt.Sprintf("%s://%s", id.language, id.path)
}

// ParseUnitID parses a string into a UnitID.
func ParseUnitID(s string) (UnitID, error) {
	if s == "" {
		return UnitID{}, fmt.Errorf("empty unit ID")
	}

	// Split on "://"
	schemeIdx := strings.Index(s, "://")
	if schemeIdx < 0 {
		return UnitID{}, fmt.Errorf("invalid unit ID format (missing ://): %q", s)
	}

	language := s[:schemeIdx]
	rest := s[schemeIdx+3:]

	// Split on "#" for symbol
	var path, symbol string
	if hashIdx := strings.LastIndex(rest, "#"); hashIdx >= 0 {
		path = rest[:hashIdx]
		symbol = rest[hashIdx+1:]
	} else {
		path = rest
	}

	return UnitID{language: language, path: path, symbol: symbol}, nil
}

// Unit represents a certifiable code unit with its identity and type.
type Unit struct {
	ID   UnitID
	Type UnitType
}

// NewUnit creates a new Unit.
func NewUnit(id UnitID, unitType UnitType) Unit {
	return Unit{ID: id, Type: unitType}
}
