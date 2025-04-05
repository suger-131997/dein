package generator

import (
	"strings"

	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
)

type ArgumentGenerator interface {
	GenerateArgument() string
}

type BodyGenerator interface {
	GenerateBody() string
}

const errorHandlingSegment = `
if err != nil{
	return nil, err
}`

func writeTypeParams(b *strings.Builder, syms *symbols.Symbols, params []component.TypeParam) {
	if len(params) == 0 {
		return
	}

	b.WriteString("[")

	for i, p := range params {
		b.WriteString(p.Prefix())

		if path := p.PkgPath(); path != "" {
			b.WriteString(syms.PkgName(path))
			b.WriteString(".")
		}

		b.WriteString(p.Name())
		writeTypeParams(b, syms, p.TypeParams())

		if i != len(params)-1 {
			b.WriteString(", ")
		}
	}

	b.WriteString("]")
}
