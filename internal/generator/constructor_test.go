package generator_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/generator"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/testpackages/a"
	"github.com/suger-131997/dein/internal/testutils"
)

func TestConstructorGeneratorGenerateBody(t *testing.T) {
	tests := []struct {
		name string

		in                 []component.Component
		out                component.Component
		constructorName    string
		constructorPkgPath string
		hasError           bool
		markExposed        bool

		want string
	}{
		{
			name: "no arguments",

			in:                 []component.Component{},
			out:                testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			constructorName:    "NewA1",
			constructorPkgPath: "github.com/suger-131997/dein/internal/testpackages/a",
			hasError:           false,
			markExposed:        false,
			want:               "a1 := a.NewA1()",
		},
		{
			name: "one arguments",

			in: []component.Component{
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A2{}))),
			},
			out:                testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			constructorName:    "NewA1",
			constructorPkgPath: "github.com/suger-131997/dein/internal/testpackages/a",
			hasError:           false,
			markExposed:        false,
			want:               "a1 := a.NewA1(a2)",
		},
		{
			name: "two arguments",

			in: []component.Component{
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A2{}))),
				testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[int]{}))),
			},
			out:                testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			constructorName:    "NewA1",
			constructorPkgPath: "github.com/suger-131997/dein/internal/testpackages/a",
			hasError:           false,
			markExposed:        false,
			want:               "a1 := a.NewA1(a2, a3)",
		},
		{
			name: "has error",

			in:                 []component.Component{},
			out:                testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			constructorName:    "NewA1",
			constructorPkgPath: "github.com/suger-131997/dein/internal/testpackages/a",
			hasError:           true,
			markExposed:        false,
			want: `a1, err := a.NewA1()
if err != nil{
	return nil, err
}`,
		},
		{
			name: "mark exposed",

			in:                 []component.Component{},
			out:                testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			constructorName:    "NewA1",
			constructorPkgPath: "github.com/suger-131997/dein/internal/testpackages/a",
			hasError:           false,
			markExposed:        true,
			want: `a1 := a.NewA1()
__c.A1 = a1`,
		},
		{
			name: "has error and mark exposed",

			in:                 []component.Component{},
			out:                testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			constructorName:    "NewA1",
			constructorPkgPath: "github.com/suger-131997/dein/internal/testpackages/a",
			hasError:           true,
			markExposed:        true,
			want: `a1, err := a.NewA1()
if err != nil{
	return nil, err
}
__c.A1 = a1`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			gen := generator.NewConstructorGenerator(
				symbols.NewSymbols(append(tc.in, tc.out), []string{tc.constructorPkgPath}),
				tc.in,
				tc.out,
				tc.constructorName,
				tc.constructorPkgPath,
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
