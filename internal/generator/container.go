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

func (g *ContainerGenerator) GenerateField() string {
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

func (g *ContainerGenerator) GenerateMethod() string {
	var b strings.Builder

	b.WriteString("func (c *Container) ")
	b.WriteString(utils.HeadToUpper(g.symbols.VarName(g.c)))
	b.WriteString("() ")
	b.WriteString(g.c.Prefix())

	if pkgName := g.symbols.PkgName(g.c.PkgPath()); pkgName != "" {
		b.WriteString(pkgName)
		b.WriteString(".")
	}

	b.WriteString(g.c.Name())
	writeTypeParams(&b, g.symbols, g.c.TypeParams())
	b.WriteString(" {\n")

	b.WriteString("return c.")
	b.WriteString(g.symbols.VarName(g.c))

	b.WriteString("\n}")

	return b.String()
}
