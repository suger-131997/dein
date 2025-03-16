package generator

import (
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/utils"
	"strings"
)

type FunctionGenerator struct {
	symbols *symbols.Symbols

	in  []component.Component
	out component.Component

	hasError bool

	isInvoked bool
}

var _ ArgumentGenerator = &FunctionGenerator{}
var _ BodyGenerator = &FunctionGenerator{}

func NewFunctionGenerator(
	syms *symbols.Symbols,
	in []component.Component,
	out component.Component,
	hasError bool,
	isInvoked bool,
) *FunctionGenerator {
	return &FunctionGenerator{
		symbols:   syms,
		in:        in,
		out:       out,
		hasError:  hasError,
		isInvoked: isInvoked,
	}
}

func (g *FunctionGenerator) GenerateArgument() string {
	var b strings.Builder

	b.WriteString(g.symbols.VarName(g.out))
	b.WriteString("Func func(")

	for i := range len(g.in) {
		if g.in[i].IsPointer() {
			b.WriteString("*")
		}
		b.WriteString(g.symbols.PkgName(g.in[i].PkgPath()))
		b.WriteString(".")
		b.WriteString(g.in[i].Name())
		writeTypeParams(&b, g.symbols, g.in[i].TypeParams())
		if i < len(g.in)-1 {
			b.WriteString(", ")
		}
	}

	b.WriteString(") (")

	if g.out.IsPointer() {
		b.WriteString("*")
	}
	b.WriteString(g.symbols.PkgName(g.out.PkgPath()))
	b.WriteString(".")
	b.WriteString(g.out.Name())
	writeTypeParams(&b, g.symbols, g.out.TypeParams())
	if g.hasError {
		b.WriteString(", error")
	}

	b.WriteString(")")

	return b.String()
}

func (g *FunctionGenerator) GenerateBody() string {
	var b strings.Builder

	b.WriteString(g.symbols.VarName(g.out))

	if g.hasError {
		b.WriteString(", err")
	}

	b.WriteString(" := ")

	b.WriteString(g.symbols.VarName(g.out))
	b.WriteString("Func(")
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

	if g.isInvoked {
		b.WriteString("\nc.")
		b.WriteString(utils.HeadToUpper(g.symbols.VarName(g.out)))
		b.WriteString(" = ")
		b.WriteString(g.symbols.VarName(g.out))
	}

	return b.String()
}
