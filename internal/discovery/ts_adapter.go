package discovery

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/code-certification/certify/internal/domain"
)

// Regex patterns for TypeScript symbol discovery.
var (
	// export function name(
	// export default function name(
	tsFuncRE = regexp.MustCompile(`^export\s+(?:default\s+)?(?:async\s+)?function\s+(\w+)`)

	// export class Name
	tsClassRE = regexp.MustCompile(`^export\s+(?:default\s+)?(?:abstract\s+)?class\s+(\w+)`)

	// export const NAME =
	// export let name =
	tsConstRE = regexp.MustCompile(`^export\s+(?:default\s+)?(?:const|let|var)\s+(\w+)`)

	// export interface Name
	tsInterfaceRE = regexp.MustCompile(`^export\s+(?:default\s+)?interface\s+(\w+)`)

	// export type Name
	tsTypeRE = regexp.MustCompile(`^export\s+(?:default\s+)?type\s+(\w+)`)

	// export enum Name
	tsEnumRE = regexp.MustCompile(`^export\s+(?:default\s+)?enum\s+(\w+)`)
)

// TSAdapter discovers TypeScript symbols using regex patterns.
type TSAdapter struct{}

// NewTSAdapter creates a new TypeScript discovery adapter.
func NewTSAdapter() *TSAdapter {
	return &TSAdapter{}
}

// Scan discovers exported TypeScript symbols in .ts/.tsx files.
func (a *TSAdapter) Scan(root string) ([]domain.Unit, error) {
	var units []domain.Unit

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, ".") || name == "node_modules" || name == "dist" || name == "build" {
				return filepath.SkipDir
			}
			return nil
		}

		ext := filepath.Ext(path)
		if ext != ".ts" && ext != ".tsx" {
			return nil
		}
		// Skip test and declaration files
		base := d.Name()
		if strings.HasSuffix(base, ".test.ts") || strings.HasSuffix(base, ".spec.ts") || strings.HasSuffix(base, ".d.ts") {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)

		fileUnits, err := a.parseFile(path, rel)
		if err != nil {
			// Fallback to file-level
			id := domain.NewUnitID("ts", rel, "")
			units = append(units, domain.NewUnit(id, domain.UnitTypeFile))
			return nil
		}
		units = append(units, fileUnits...)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(units, func(i, j int) bool {
		return units[i].ID.String() < units[j].ID.String()
	})

	return units, nil
}

func (a *TSAdapter) parseFile(path, rel string) ([]domain.Unit, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var units []domain.Unit
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if m := tsFuncRE.FindStringSubmatch(line); m != nil {
			id := domain.NewUnitID("ts", rel, m[1])
			units = append(units, domain.NewUnit(id, domain.UnitTypeFunction))
			continue
		}
		if m := tsClassRE.FindStringSubmatch(line); m != nil {
			id := domain.NewUnitID("ts", rel, m[1])
			units = append(units, domain.NewUnit(id, domain.UnitTypeClass))
			continue
		}
		if m := tsInterfaceRE.FindStringSubmatch(line); m != nil {
			id := domain.NewUnitID("ts", rel, m[1])
			units = append(units, domain.NewUnit(id, domain.UnitTypeClass))
			continue
		}
		if m := tsTypeRE.FindStringSubmatch(line); m != nil {
			id := domain.NewUnitID("ts", rel, m[1])
			units = append(units, domain.NewUnit(id, domain.UnitTypeClass))
			continue
		}
		if m := tsEnumRE.FindStringSubmatch(line); m != nil {
			id := domain.NewUnitID("ts", rel, m[1])
			units = append(units, domain.NewUnit(id, domain.UnitTypeClass))
			continue
		}
		if m := tsConstRE.FindStringSubmatch(line); m != nil {
			id := domain.NewUnitID("ts", rel, m[1])
			units = append(units, domain.NewUnit(id, domain.UnitTypeFunction)) // Treat exported const as function-level
			continue
		}
	}

	return units, scanner.Err()
}
