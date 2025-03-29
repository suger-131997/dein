package provider

import (
	"errors"
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/generator"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/utils"
)

func NewConstructorProvider(f any, hasError bool) *Provider {
	fv := reflect.ValueOf(f)

	if fv.Kind() != reflect.Func {
		return &Provider{
			err: errors.New("allow only function"),
		}
	}

	funcPath := strings.TrimSuffix(runtime.FuncForPC(fv.Pointer()).Name(), "[...]")

	l := strings.Split(path.Base(funcPath), ".")
	if len(l) > 2 {
		return &Provider{
			err: errors.New("anonymous function is not allowed"),
		}
	}

	constructorPkgPath := funcPath[:strings.LastIndex(funcPath, ".")]
	constructorName := l[1]

	pkgPaths := make([]string, 0)

	ft := fv.Type()

	in := make([]component.Component, 0, ft.NumIn())

	for i := 0; i < ft.NumIn(); i++ {
		c, err := component.NewComponent(ft.In(i))
		if err != nil {
			return &Provider{
				err: err,
			}
		}

		in = append(in, c)
		pkgPaths = append(pkgPaths, c.PkgPaths()...)
	}

	out, err := component.NewComponent(ft.Out(0))
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
		buildGenerator: func(syms *symbols.Symbols, markExposed bool) generator.BodyGenerator {
			return generator.NewConstructorGenerator(
				syms,
				in,
				out,
				constructorName,
				constructorPkgPath,
				hasError,
				markExposed,
			)
		},
	}
}
