package generator

import (
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/utils"
	"strings"
)

type ContainerGenerator struct {
	symbols *symbols.Symbols
	c       component.Component
}

func NewContainerGenerator(syms *symbols.Symbols, c component.Component) *ContainerGenerator {
	return &ContainerGenerator{
		symbols: syms,
		c:       c,
	}
}

func (g *ContainerGenerator) Generate() string {
	var b strings.Builder
	b.WriteString(utils.HeadToUpper(g.symbols.VarName(g.c)))

	b.WriteString(" ")

	if g.c.IsPointer() {
		b.WriteString("*")
	}
	b.WriteString(g.symbols.PkgName(g.c.PkgPath()))
	b.WriteString(".")
	b.WriteString(g.c.Name())

	if params := g.c.TypeParams(); len(params) != 0 {
		b.WriteString("[")
		for i, p := range params {
			if path := p.PkgPath(); path != "" {
				b.WriteString(g.symbols.PkgName(path))
				b.WriteString(".")
			}
			b.WriteString(p.Name())
			if i != len(params)-1 {
				b.WriteString(", ")
			}
		}
		b.WriteString("]")
	}

	return b.String()
}
