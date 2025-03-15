package generator

import (
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
	"strings"
)

type ComponentArgumentGenerator struct {
	symbols *symbols.Symbols
	c       component.Component
}

var _ ArgumentGenerator = &ComponentArgumentGenerator{}

func NewComponentArgumentGenerator(syms *symbols.Symbols, c component.Component) *ComponentArgumentGenerator {
	return &ComponentArgumentGenerator{
		symbols: syms,
		c:       c,
	}
}

func (g *ComponentArgumentGenerator) GenerateArgument() string {
	var b strings.Builder
	b.WriteString(g.symbols.VarName(g.c))
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
