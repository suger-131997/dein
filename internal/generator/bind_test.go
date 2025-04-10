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

func TestBindGeneratorGenerateBody(t *testing.T) {
	tests := []struct {
		name string

		distPkgPath string
		bindTo      component.Component
		implement   component.Component
		markExposed bool

		want string
	}{
		{
			name: "b.B bind to a.IA1",

			distPkgPath: "main",
			bindTo: func() component.Component {
				var ia *a.IA1
				return testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(ia).Elem()))
			}(),
			implement:   testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
			markExposed: false,

			want: "var iA1 a.IA1 = b",
		},
		{
			name: "*b.B bind to a.IA1",

			distPkgPath: "main",
			bindTo: func() component.Component {
				var ia *a.IA1
				return testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(ia).Elem()))
			}(),
			implement:   testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&b.B{}))),
			markExposed: false,

			want: "var iA1 a.IA1 = b",
		},
		{
			name: "mark exposed",

			distPkgPath: "main",
			bindTo: func() component.Component {
				var ia *a.IA1
				return testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(ia).Elem()))
			}(),
			implement:   testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
			markExposed: true,

			want: `var iA1 a.IA1 = b
__c.iA1 = iA1`,
		},
		{
			name: "generics component with one type parameter",

			distPkgPath: "main",
			bindTo: func() component.Component {
				var ia *a.IA2[c.C]
				return testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(ia).Elem()))
			}(),
			implement:   testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
			markExposed: false,

			want: "var iA2 a.IA2[c.C] = b",
		},
		{
			name: "generics component with build-in type parameter",

			distPkgPath: "main",
			bindTo: func() component.Component {
				var ia *a.IA2[int]
				return testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(ia).Elem()))
			}(),
			implement:   testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
			markExposed: false,

			want: "var iA2 a.IA2[int] = b",
		},
		{
			name: "dist package the same as bind to component package",

			distPkgPath: "github.com/suger-131997/dein/internal/testpackages/a",
			bindTo: func() component.Component {
				var ia *a.IA1
				return testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(ia).Elem()))
			}(),
			implement:   testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
			markExposed: false,

			want: "var iA1 IA1 = b",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			gen := generator.NewBindGenerator(
				symbols.NewSymbols(tc.distPkgPath, []component.Component{tc.bindTo, tc.implement}, []string{}),
				tc.bindTo,
				tc.implement,
				tc.markExposed,
			)

			got := gen.GenerateBody()
			if diff := cmp.Diff(got, tc.want); diff != "" {
				tt.Errorf("GenerateBody() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
