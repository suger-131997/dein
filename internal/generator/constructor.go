package generator

import (
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/utils"
	"strings"
)

type ConstructorGenerator struct {
	symbols *symbols.Symbols

	in  []component.Component
	out component.Component

	constructorName    string
	constructorPkgPath string

	hasError bool

	markExposed bool
}

var _ BodyGenerator = &ConstructorGenerator{}

func NewConstructorGenerator(
	syms *symbols.Symbols,
	in []component.Component,
	out component.Component,
	constructorName string,
	constructorPkgPath string,
	hasError bool,
	markExposed bool,
) *ConstructorGenerator {
	return &ConstructorGenerator{
		symbols:            syms,
		in:                 in,
		out:                out,
		constructorName:    constructorName,
		constructorPkgPath: constructorPkgPath,
		hasError:           hasError,
		markExposed:        markExposed,
	}
}

func (g *ConstructorGenerator) GenerateBody() string {
	var b strings.Builder

	b.WriteString(g.symbols.VarName(g.out))

	if g.hasError {
		b.WriteString(", err")
	}

	b.WriteString(" := ")

	b.WriteString(g.symbols.PkgName(g.constructorPkgPath))
	b.WriteString(".")
	b.WriteString(g.constructorName)
	b.WriteString("(")
	for i := range len(g.in) {
		b.WriteString(g.symbols.VarName(g.in[i]))
		if i < len(g.in)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString(")")

	if g.hasError {
		b.WriteString(errorHandlingSegment)
	}

	if g.markExposed {
		b.WriteString("\nc.")
		b.WriteString(utils.HeadToUpper(g.symbols.VarName(g.out)))
		b.WriteString(" = ")
		b.WriteString(g.symbols.VarName(g.out))
	}

	return b.String()
}
