package provider

import (
	"errors"
	"fmt"
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/generator"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/utils"
	"reflect"
)

func NewBindProvider(i, t reflect.Type) *Provider {
	if i.Kind() != reflect.Interface {
		return &Provider{
			err: errors.New("bind target must be an interface"),
		}
	}
	if !t.Implements(i) {
		return &Provider{
			err: fmt.Errorf("%s must implement the interface %s", t.String(), i.String()),
		}
	}

	in, err := component.NewComponent(t)
	if err != nil {
		return &Provider{
			err: err,
		}
	}
	out, err := component.NewComponent(i)
	if err != nil {
		return &Provider{
			err: err,
		}
	}

	return &Provider{
		in:       []component.Component{in},
		out:      out,
		pkgPaths: utils.Uniq(append(in.PkgPaths(), out.PkgPaths()...)),
		buildGenerator: func(syms *symbols.Symbols, isInvoked bool) generator.BodyGenerator {
			return generator.NewBindGenerator(syms, out, in, isInvoked)
		},
	}
}
