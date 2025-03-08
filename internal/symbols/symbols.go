package symbols

import (
	"fmt"
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/utils"
	"iter"
	"path"
	"sort"
)

type Symbols struct {
	varNames       map[component.Component]string
	pkgNames       map[string]string
	sortedPkgPaths []string
}

func NewSymbols(c []component.Component, p []string) *Symbols {
	components := utils.Uniq(c)
	sort.Slice(components, func(i, j int) bool {
		return components[i].Less(components[j])
	})
	pkgPaths := utils.Uniq(p)
	sort.Strings(pkgPaths)

	nameCounts := make(map[string]int)
	varNames := make(map[component.Component]string, len(components))
	for _, c := range components {
		name := utils.HeadToLower(c.Name())
		varNames[c] = name
		if count, ok := nameCounts[name]; ok {
			varNames[c] = fmt.Sprintf("%s%d", name, count)
			nameCounts[name] = count + 1
			continue
		}
		varNames[c] = name
		nameCounts[name] = 1
	}

	pkgNames := make(map[string]string, len(pkgPaths))
	for _, p := range pkgPaths {
		pkgName := path.Base(p)
		if count, ok := nameCounts[pkgName]; ok {
			pkgNames[p] = fmt.Sprintf("%s%d", pkgName, count)
			nameCounts[pkgName] = count + 1
			continue
		}
		pkgNames[p] = pkgName
		nameCounts[pkgName] = 1
	}

	return &Symbols{
		varNames:       varNames,
		pkgNames:       pkgNames,
		sortedPkgPaths: pkgPaths,
	}
}

func (s *Symbols) VarName(c component.Component) string {
	return s.varNames[c]
}

func (s *Symbols) PkgName(pkgPath string) string {
	return s.pkgNames[pkgPath]
}

func (s *Symbols) Imports() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for _, p := range s.sortedPkgPaths {
			if !yield(s.pkgNames[p], p) {
				return
			}
		}
	}
}
