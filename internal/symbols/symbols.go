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

	n := newNamer(path.Base(distPkgPath))

	varNames := make(map[component.Component]string, len(components))

	for _, c := range components {
		name := utils.HeadToLower(c.Name())
		varNames[c] = n.name(name)
	}

	pkgNames := make(map[string]string, len(pkgPaths))

	for _, p := range pkgPaths {
		if p == distPkgPath {
			continue
		}

		pkgNames[p] = n.name(path.Base(p))
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

type namer struct {
	names map[string]struct{}
	count map[string]int
}

func newNamer(inits ...string) *namer {
	names := make(map[string]struct{}, len(inits))
	count := make(map[string]int, len(inits))

	for _, init := range inits {
		names[init] = struct{}{}
		count[init] = 1
	}

	return &namer{
		names: names,
		count: count,
	}
}

func (n *namer) name(name string) string {
	if _, exists := n.names[name]; exists {
		count, ok := n.count[name]
		if !ok {
			count = 1
		}

		offset := 1

		for {
			if _, dupl := n.names[fmt.Sprintf("%s_%d", name, count+offset)]; !dupl {
				break
			}

			offset++
		}

		n.count[name] = count + offset

		newName := fmt.Sprintf("%s_%d", name, n.count[name])
		n.names[newName] = struct{}{}

		return newName
	}

	n.names[name] = struct{}{}
	n.count[name] = 1

	return name
}
