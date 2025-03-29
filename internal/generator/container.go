package generator

import (
	"strings"

	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/utils"
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
	writeTypeParams(&b, g.symbols, g.c.TypeParams())

	return b.String()
}
