package generator

import (
	"strings"

	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/utils"
)

type BindGenerator struct {
	symbols *symbols.Symbols

	bindTo    component.Component
	implement component.Component

	markExposed bool
}

var _ BodyGenerator = &BindGenerator{}

func NewBindGenerator(symbols *symbols.Symbols, bindTo, implement component.Component, markExposed bool) *BindGenerator {
	return &BindGenerator{
		symbols:     symbols,
		bindTo:      bindTo,
		implement:   implement,
		markExposed: markExposed,
	}
}

func (g BindGenerator) GenerateBody() string {
	var b strings.Builder

	b.WriteString("var ")

	b.WriteString(g.symbols.VarName(g.bindTo))

	b.WriteString(" ")

	if pkgName := g.symbols.PkgName(g.bindTo.PkgPath()); pkgName != "" {
		b.WriteString(pkgName)
		b.WriteString(".")
	}

	b.WriteString(g.bindTo.Name())
	writeTypeParams(&b, g.symbols, g.bindTo.TypeParams())

	b.WriteString(" = ")

	b.WriteString(g.symbols.VarName(g.implement))

	if g.markExposed {
		b.WriteString("\n__c.")
		b.WriteString(utils.HeadToUpper(g.symbols.VarName(g.bindTo)))
		b.WriteString(" = ")
		b.WriteString(g.symbols.VarName(g.bindTo))
	}

	return b.String()
}
