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

func TestContainerGeneratorGenerate(t *testing.T) {
	tests := []struct {
		name string

		c component.Component

		want string
	}{
		{
			name: "normal component",

			c: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),

			want: "A1 a.A1",
		},
		{
			name: "pointer component",

			c: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&a.A1{}))),

			want: "A1 *a.A1",
		},
		{
			name: "generics component with one type parameter",

			c: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[b.B]{}))),

			want: "A3 a.A3[b.B]",
		},
		{
			name: "generics component with two type parameter",

			c: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A4[b.B, c.C]{}))),

			want: "A4 a.A4[b.B, c.C]",
		},
		{
			name: "generics component with build-in type parameter",

			c: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[int]{}))),

			want: "A3 a.A3[int]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			gen := generator.NewContainerGenerator(
				symbols.NewSymbols([]component.Component{tc.c}, []string{}),
				tc.c,
			)

			got := gen.Generate()
			if diff := cmp.Diff(got, tc.want); diff != "" {
				tt.Errorf("Generate() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
