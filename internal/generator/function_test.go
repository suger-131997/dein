package generator_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/generator"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/testpackages/a"
	"github.com/suger-131997/dein/internal/testpackages/b"
	"github.com/suger-131997/dein/internal/testutils"
)

func TestFunctionGeneratorGenerateArgument(t *testing.T) {
	tests := []struct {
		name string

		distPkgPath string
		in          []component.Component
		out         component.Component
		hasError    bool
		markExposed bool

		want string
	}{
		{
			name: "no arguments",

			distPkgPath: "main",
			in:          []component.Component{},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    false,
			markExposed: false,

			want: "__funcA1 func() (a.A1)",
		},
		{
			name: "one argument",

			distPkgPath: "main",
			in: []component.Component{
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A2{}))),
			},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    false,
			markExposed: false,

			want: "__funcA1 func(a.A2) (a.A1)",
		},
		{
			name: "two arguments",

			distPkgPath: "main",
			in: []component.Component{
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A2{}))),
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
			},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    false,
			markExposed: false,

			want: "__funcA1 func(a.A2, b_2.B) (a.A1)",
		},
		{
			name: "generics argument",

			distPkgPath: "main",
			in: []component.Component{
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[a.A2]{}))),
			},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    false,
			markExposed: false,

			want: "__funcA1 func(a.A3[a.A2]) (a.A1)",
		},
		{
			name: "pointer argument",

			distPkgPath: "main",
			in: []component.Component{
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&a.A2{}))),
			},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    false,
			markExposed: false,

			want: "__funcA1 func(*a.A2) (a.A1)",
		},
		{
			name: "pointer return",

			distPkgPath: "main",
			in: []component.Component{
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A2{}))),
			},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&a.A1{}))),
			hasError:    false,
			markExposed: false,

			want: "__funcA1 func(a.A2) (*a.A1)",
		},
		{
			name: "has error",

			distPkgPath: "main",
			in:          []component.Component{},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    true,
			markExposed: false,

			want: "__funcA1 func() (a.A1, error)",
		},
		{
			name: "in dist package type input and output",

			distPkgPath: "github.com/suger-131997/dein/internal/testpackages/a",
			in: []component.Component{
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A2{}))),
			},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&a.A1{}))),
			hasError:    false,
			markExposed: false,

			want: "__funcA1 func(A2) (*A1)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			gen := generator.NewFunctionGenerator(
				symbols.NewSymbols(tc.distPkgPath, append(tc.in, tc.out), []string{}),
				tc.in,
				tc.out,
				tc.hasError,
				tc.markExposed,
			)

			got := gen.GenerateArgument()
			if diff := cmp.Diff(got, tc.want); diff != "" {
				tt.Errorf("GenerateArgument() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestFunctionGeneratorGenerateBody(t *testing.T) {
	tests := []struct {
		name string

		distPkgPath string
		in          []component.Component
		out         component.Component
		hasError    bool
		markExposed bool

		want string
	}{
		{
			name: "no arguments",

			distPkgPath: "main",
			in:          []component.Component{},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    false,
			markExposed: false,

			want: "a1 := __funcA1()",
		},
		{
			name: "one argument",

			distPkgPath: "main",
			in: []component.Component{
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A2{}))),
			},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    false,
			markExposed: false,

			want: "a1 := __funcA1(a2)",
		},
		{
			name: "two arguments",

			distPkgPath: "main",
			in: []component.Component{
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A2{}))),
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
			},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    false,
			markExposed: false,

			want: "a1 := __funcA1(a2, b)",
		},
		{
			name: "has error",

			distPkgPath: "main",
			in:          []component.Component{},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    true,
			markExposed: false,

			want: `a1, err := __funcA1()
if err != nil{
	return nil, err
}`,
		},
		{
			name: "mark exposed",

			distPkgPath: "main",
			in:          []component.Component{},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    false,
			markExposed: true,

			want: `a1 := __funcA1()
__c.a1 = a1`,
		},

		{
			name: "has error and mark exposed",

			distPkgPath: "main",
			in:          []component.Component{},
			out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			hasError:    true,
			markExposed: true,

			want: `a1, err := __funcA1()
if err != nil{
	return nil, err
}
__c.a1 = a1`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			gen := generator.NewFunctionGenerator(
				symbols.NewSymbols(tc.distPkgPath, append(tc.in, tc.out), []string{}),
				tc.in,
				tc.out,
				tc.hasError,
				tc.markExposed,
			)

			got := gen.GenerateBody()
			if diff := cmp.Diff(got, tc.want); diff != "" {
				tt.Errorf("GenerateBody() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
