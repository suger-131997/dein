package provider

import (
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/generator"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/utils"
	"reflect"
)

func NewFunctionProvider(r reflect.Type, hasError bool, tl ...reflect.Type) *Provider {
	pkgPaths := make([]string, 0)

	in := make([]component.Component, 0, len(tl))
	for _, t := range tl {
		c, err := component.NewComponent(t)
		if err != nil {
			return &Provider{
				err: err,
			}
		}
		in = append(in, c)
		pkgPaths = append(pkgPaths, c.PkgPaths()...)
	}

	out, err := component.NewComponent(r)
	if err != nil {
		return &Provider{
			err: err,
		}
	}
	pkgPaths = append(pkgPaths, out.PkgPaths()...)

	return &Provider{
		in:       in,
		out:      out,
		pkgPaths: utils.Uniq(pkgPaths),
		buildGenerator: func(syms *symbols.Symbols, isInvoked bool) generator.BodyGenerator {
			return generator.NewFunctionGenerator(syms, in, out, hasError, isInvoked)
		},
	}
}
