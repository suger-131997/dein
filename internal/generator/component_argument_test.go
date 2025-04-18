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
	"github.com/suger-131997/dein/internal/testpackages/c"
	"github.com/suger-131997/dein/internal/testutils"
)

func TestComponentArgumentGeneratorGenerate(t *testing.T) {
	tests := []struct {
		name string

		distPkgPath string
		c           component.Component

		want string
	}{
		{
			name: "normal component",

			distPkgPath: "main",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),

			want: "a1 a.A1",
		},
		{
			name: "pointer of pointers component",

			distPkgPath: "main",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(new(*a.A1)))),

			want: "a1 **a.A1",
		},
		{
			name: "array component",

			distPkgPath: "main",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf([]a.A1{}))),

			want: "a1 []a.A1",
		},
		{
			name: "pointer component",

			distPkgPath: "main",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&a.A1{}))),

			want: "a1 *a.A1",
		},
		{
			name: "generics component with one type parameter",

			distPkgPath: "main",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[b.B]{}))),

			want: "a3 a.A3[b.B]",
		},
		{
			name: "generics component with two type parameter",

			distPkgPath: "main",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A4[b.B, c.C]{}))),

			want: "a4 a.A4[b.B, c.C]",
		},
		{
			name: "generics component with build-in type parameter",

			distPkgPath: "main",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[int]{}))),

			want: "a3 a.A3[int]",
		},
		{
			name: "dist package the same as component package",

			distPkgPath: "github.com/suger-131997/dein/internal/testpackages/a",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),

			want: "a1 A1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			gen := generator.NewComponentArgumentGenerator(
				symbols.NewSymbols(tc.distPkgPath, []component.Component{tc.c}, []string{}),
				tc.c,
			)

			got := gen.GenerateArgument()
			if diff := cmp.Diff(got, tc.want); diff != "" {
				tt.Errorf("GenerateArgument() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
