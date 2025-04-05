package symbols

import (
	"fmt"
	"path"
	"sort"

	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/utils"
)

type Symbols struct {
	distPkgPath     string
	varNames        map[component.Component]string
	pkgNames        map[string]string
	orderedPkgPaths []string
}

func NewSymbols(distPkgPath string, _components []component.Component, _pkgPaths []string) *Symbols {
	components := utils.Uniq(_components)
	sort.Slice(components, func(i, j int) bool {
		return components[i].Less(components[j])
	})

	pkgPaths := make([]string, 0)
	for _, c := range components {
		pkgPaths = append(pkgPaths, c.PkgPaths()...)
	}

	pkgPaths = utils.Uniq(append(pkgPaths, _pkgPaths...))
	sort.Strings(pkgPaths)

	nameCounts := make(map[string]int)
	nameCounts[path.Base(distPkgPath)] = 1

	varNames := make(map[component.Component]string, len(components))

	for _, c := range components {
		name := utils.HeadToLower(c.Name())
		varNames[c] = name

		if count, ok := nameCounts[name]; ok {
			varNames[c] = fmt.Sprintf("%s_%d", name, count+1)
			nameCounts[name] = count + 1

			continue
		}

		varNames[c] = name
		nameCounts[name] = 1
	}

	pkgNames := make(map[string]string, len(pkgPaths))

	for _, p := range pkgPaths {
		if p == distPkgPath {
			continue
		}

		pkgName := path.Base(p)
		if count, ok := nameCounts[pkgName]; ok {
			pkgNames[p] = fmt.Sprintf("%s_%d", pkgName, count+1)
			nameCounts[pkgName] = count + 1

			continue
		}

		pkgNames[p] = pkgName
		nameCounts[pkgName] = 1
	}

	return &Symbols{
		distPkgPath:     distPkgPath,
		varNames:        varNames,
		pkgNames:        pkgNames,
		orderedPkgPaths: pkgPaths,
	}
}

func (s *Symbols) DistPkgName() string {
	return path.Base(s.distPkgPath)
}

func (s *Symbols) VarName(c component.Component) string {
	return s.varNames[c]
}

func (s *Symbols) PkgName(pkgPath string) string {
	if pkgPath == "" || pkgPath == s.distPkgPath {
		return ""
	}

	return s.pkgNames[pkgPath]
}

func (s *Symbols) Imports() [][]string {
	imports := make([][]string, 0, len(s.orderedPkgPaths))

	for _, p := range s.orderedPkgPaths {
		if p == s.distPkgPath {
			continue
		}

		imports = append(imports, []string{s.pkgNames[p], p})
	}

	return imports
}
