package generator

import (
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/utils"
	"strings"
)

type BindGenerator struct {
	symbols *symbols.Symbols

	bindTo    component.Component
	implement component.Component

	isInvoked bool
}

var _ BodyGenerator = &BindGenerator{}

func NewBindGenerator(symbols *symbols.Symbols, bindTo, implement component.Component, isInvoked bool) *BindGenerator {
	return &BindGenerator{
		symbols:   symbols,
		bindTo:    bindTo,
		implement: implement,
		isInvoked: isInvoked,
	}
}

func (g BindGenerator) GenerateBody() string {
	var b strings.Builder
	b.WriteString("var ")

	b.WriteString(g.symbols.VarName(g.bindTo))

	b.WriteString(" ")

	b.WriteString(g.symbols.PkgName(g.bindTo.PkgPath()))
	b.WriteString(".")
	b.WriteString(g.bindTo.Name())
	if params := g.bindTo.TypeParams(); len(params) != 0 {
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

	b.WriteString(" = ")

	b.WriteString(g.symbols.VarName(g.implement))

	if g.isInvoked {
		b.WriteString("\nc.")
		b.WriteString(utils.HeadToUpper(g.symbols.VarName(g.bindTo)))
		b.WriteString(" = ")
		b.WriteString(g.symbols.VarName(g.bindTo))
	}

	return b.String()
}
