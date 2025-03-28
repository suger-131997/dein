package provider

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/testpackages/a"
	"github.com/suger-131997/dein/internal/testpackages/b"
	"github.com/suger-131997/dein/internal/testpackages/c"
	"github.com/suger-131997/dein/internal/testutils"
	"reflect"
	"testing"
)

func TestNewFunctionProvider(t *testing.T) {
	tests := []struct {
		name string

		r        reflect.Type
		tl       []reflect.Type
		hasError bool

		want *Provider
	}{
		{
			name: "no arguments",

			r:        reflect.TypeOf(a.A1{}),
			tl:       []reflect.Type{},
			hasError: false,

			want: &Provider{
				in:          []component.Component{},
				out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
				pkgPaths:    []string{"github.com/suger-131997/dein/internal/testpackages/a"},
				markExposed: false,
				err:         nil,
			},
		},
		{
			name: "one arguments",

			r:        reflect.TypeOf(a.A1{}),
			tl:       []reflect.Type{reflect.TypeOf(b.B{})},
			hasError: false,

			want: &Provider{
				in:  []component.Component{testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{})))},
				out: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
				pkgPaths: []string{
					"github.com/suger-131997/dein/internal/testpackages/b",
					"github.com/suger-131997/dein/internal/testpackages/a",
				},
				markExposed: false,
				err:         nil,
			},
		},
		{
			name: "two arguments",

			r:        reflect.TypeOf(a.A1{}),
			tl:       []reflect.Type{reflect.TypeOf(b.B{}), reflect.TypeOf(c.C{})},
			hasError: false,

			want: &Provider{
				in: []component.Component{
					testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
					testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(c.C{}))),
				},
				out: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
				pkgPaths: []string{
					"github.com/suger-131997/dein/internal/testpackages/b",
					"github.com/suger-131997/dein/internal/testpackages/c",
					"github.com/suger-131997/dein/internal/testpackages/a",
				},
				markExposed: false,
				err:         nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got := NewFunctionProvider(tc.r, tc.hasError, tc.tl...)
			if err := got.CheckError(); err != nil {
				tt.Errorf("unexpected error: %v", err)
				return
			}

			if diff := cmp.Diff(got, tc.want,
				cmp.AllowUnexported(Provider{}, component.Component{}),
				cmpopts.IgnoreFields(Provider{}, "buildGenerator"),
			); diff != "" {
				tt.Errorf("Provider is mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
