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

func TestContainerGeneratorGenerateField(t *testing.T) {
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
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[int]{}))),

			want: "a3 A3[int]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			gen := generator.NewContainerGenerator(
				symbols.NewSymbols(tc.distPkgPath, []component.Component{tc.c}, []string{}),
				tc.c,
			)

			got := gen.GenerateField()
			if diff := cmp.Diff(got, tc.want); diff != "" {
				tt.Errorf("Generate() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestContainerGeneratorGenerateMethod(t *testing.T) {
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

			want: `func (c *Container) A1() a.A1 {
return c.a1
}`,
		},
		{
			name: "pointer component",

			distPkgPath: "main",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&a.A1{}))),

			want: `func (c *Container) A1() *a.A1 {
return c.a1
}`,
		},
		{
			name: "generics component with one type parameter",

			distPkgPath: "main",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[b.B]{}))),

			want: `func (c *Container) A3() a.A3[b.B] {
return c.a3
}`,
		},
		{
			name: "generics component with two type parameter",

			distPkgPath: "main",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A4[b.B, c.C]{}))),

			want: `func (c *Container) A4() a.A4[b.B, c.C] {
return c.a4
}`,
		},
		{
			name: "generics component with build-in type parameter",

			distPkgPath: "main",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[int]{}))),

			want: `func (c *Container) A3() a.A3[int] {
return c.a3
}`,
		},
		{
			name: "dist package the same as component package",

			distPkgPath: "github.com/suger-131997/dein/internal/testpackages/a",
			c:           testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[int]{}))),

			want: `func (c *Container) A3() A3[int] {
return c.a3
}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			gen := generator.NewContainerGenerator(
				symbols.NewSymbols(tc.distPkgPath, []component.Component{tc.c}, []string{}),
				tc.c,
			)

			got := gen.GenerateMethod()
			if diff := cmp.Diff(got, tc.want); diff != "" {
				tt.Errorf("Generate() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
