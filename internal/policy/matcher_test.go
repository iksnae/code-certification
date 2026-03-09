package policy_test

import (
	"testing"

	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/policy"
)

func TestMatcher_LanguageMatch(t *testing.T) {
	packs := []domain.PolicyPack{
		{Name: "go-standard", Language: "go"},
		{Name: "ts-standard", Language: "ts"},
		{Name: "global", Language: ""},
	}

	m := policy.NewMatcher(packs)

	goUnit := domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction)
	matched := m.Match(goUnit)

	names := packNames(matched)
	if !names["go-standard"] {
		t.Error("go unit should match go-standard")
	}
	if names["ts-standard"] {
		t.Error("go unit should not match ts-standard")
	}
	if !names["global"] {
		t.Error("go unit should match global")
	}
}

func TestMatcher_PathPatterns(t *testing.T) {
	packs := []domain.PolicyPack{
		{Name: "security", PathPatterns: []string{"auth/*", "security/*"}},
	}

	m := policy.NewMatcher(packs)

	authUnit := domain.NewUnit(domain.NewUnitID("go", "auth/handler.go", ""), domain.UnitTypeFile)
	matched := m.Match(authUnit)
	if len(matched) != 1 {
		t.Errorf("auth unit should match security pack, got %d matches", len(matched))
	}

	otherUnit := domain.NewUnit(domain.NewUnitID("go", "api/handler.go", ""), domain.UnitTypeFile)
	matched = m.Match(otherUnit)
	if len(matched) != 0 {
		t.Errorf("api unit should not match security pack, got %d matches", len(matched))
	}
}

func TestMatcher_NoPacks(t *testing.T) {
	m := policy.NewMatcher(nil)
	unit := domain.NewUnit(domain.NewUnitID("go", "main.go", ""), domain.UnitTypeFile)
	if len(m.Match(unit)) != 0 {
		t.Error("no packs should match nothing")
	}
}

func TestMatcher_GlobalMatchesAll(t *testing.T) {
	packs := []domain.PolicyPack{
		{Name: "global", Language: ""},
	}

	m := policy.NewMatcher(packs)

	for _, lang := range []string{"go", "ts", "py", "file"} {
		unit := domain.NewUnit(domain.NewUnitID(lang, "test."+lang, ""), domain.UnitTypeFile)
		matched := m.Match(unit)
		if len(matched) != 1 {
			t.Errorf("global should match %s unit, got %d", lang, len(matched))
		}
	}
}

func packNames(packs []domain.PolicyPack) map[string]bool {
	m := make(map[string]bool)
	for _, p := range packs {
		m[p.Name] = true
	}
	return m
}
