package generator

import (
	"strings"

	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
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

	b.WriteString(g.c.Prefix())

	if pkgName := g.symbols.PkgName(g.c.PkgPath()); pkgName != "" {
		b.WriteString(pkgName)
		b.WriteString(".")
	}

	b.WriteString(g.c.Name())
	writeTypeParams(&b, g.symbols, g.c.TypeParams())

	return b.String()
}
