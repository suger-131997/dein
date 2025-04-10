package symbols_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/testpackages/a"
	a2 "github.com/suger-131997/dein/internal/testpackages/a/a"
	"github.com/suger-131997/dein/internal/testpackages/b"
	"github.com/suger-131997/dein/internal/testpackages/c"
	"github.com/suger-131997/dein/internal/testutils"
)

func Test_NewSymbols(t *testing.T) {
	tests := []struct {
		name string

		symbols *symbols.Symbols

		wantDistPkgName string
	}{
		{
			name: "normal case",

			symbols: symbols.NewSymbols(
				"main",
				[]component.Component{
					testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
					testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[b.B]{}))),
					testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a2.A1{}))),
					testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
					testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A4[int, c.C]{}))),
				},
				[]string{"github.com/suger-131997/dein/internal/testpackages/x"},
			),

			wantDistPkgName: "main",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.symbols.DistPkgName(); got != tc.wantDistPkgName {
				t.Errorf("DistPkgName() = %s, want %s", got, tc.wantDistPkgName)
			}
		})
	}
}

func TestNewSymbols(t *testing.T) {
	syms := symbols.NewSymbols(
		"main",
		[]component.Component{
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[b.B]{}))),
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a2.A1{}))),
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A4[int, c.C]{}))),
		},
		[]string{"github.com/suger-131997/dein/internal/testpackages/x"},
	)

	t.Run("get dist pkg name", func(tt *testing.T) {
		want := "main"
		if got := syms.DistPkgName(); got != want {
			tt.Errorf("DistPkgName() = %s, want %s", got, want)
		}
	})

	t.Run("get var name", func(tt *testing.T) {
		want := "b"
		if got := syms.VarName(testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{})))); got != want {
			tt.Errorf("VarName() = %s, want %s", got, want)
		}
	})

	t.Run("get pkg name", func(tt *testing.T) {
		want := "x"
		if got := syms.PkgName("github.com/suger-131997/dein/internal/testpackages/x"); got != want {
			tt.Errorf("PkgName() = %s, want %s", got, want)
		}
	})

	t.Run("get type param pkg name", func(tt *testing.T) {
		want := "c"
		if got := syms.PkgName("github.com/suger-131997/dein/internal/testpackages/c"); got != want {
			tt.Errorf("PkgName() = %s, want %s", got, want)
		}
	})

	t.Run("get duplicated var names", func(tt *testing.T) {
		want := []string{"a1", "a1_2"}

		got := make([]string, 0, len(want))
		got = append(got, syms.VarName(testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{})))))
		got = append(got, syms.VarName(testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a2.A1{})))))

		if diff := cmp.Diff(got, want, cmpopts.SortSlices(func(i, j string) bool {
			return i < j
		})); diff != "" {
			tt.Errorf("Provider is mismatch (-got +want):\n%s", diff)
		}
	})

	t.Run("get duplicated pkg names", func(tt *testing.T) {
		want := []string{"a", "a_2"}

		got := make([]string, 0, len(want))
		got = append(got, syms.PkgName("github.com/suger-131997/dein/internal/testpackages/a"))
		got = append(got, syms.PkgName("github.com/suger-131997/dein/internal/testpackages/a/a"))

		if diff := cmp.Diff(got, want, cmpopts.SortSlices(func(i, j string) bool {
			return i < j
		})); diff != "" {
			tt.Errorf("Provider is mismatch (-got +want):\n%s", diff)
		}
	})

	t.Run("get ordered pkg name and pkg path", func(tt *testing.T) {
		want := [][]string{
			{"a", "github.com/suger-131997/dein/internal/testpackages/a"},
			{"a_2", "github.com/suger-131997/dein/internal/testpackages/a/a"},
			{"b_2", "github.com/suger-131997/dein/internal/testpackages/b"},
			{"c", "github.com/suger-131997/dein/internal/testpackages/c"},
			{"x", "github.com/suger-131997/dein/internal/testpackages/x"},
		}

		got := syms.Imports()

		if diff := cmp.Diff(got, want); diff != "" {
			tt.Errorf("Imports() is mismatch (-got +want):\n%s", diff)
		}
	})
}

func TestNewSymbols_SpecificDistPkg(t *testing.T) {
	syms := symbols.NewSymbols(
		"github.com/suger-131997/dein/internal/testpackages/b",
		[]component.Component{
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
		},
		[]string{},
	)

	t.Run("get dist pkg name", func(tt *testing.T) {
		want := "b"
		if got := syms.DistPkgName(); got != want {
			tt.Errorf("DistPkgName() = %s, want %s", got, want)
		}
	})

	t.Run("get var name", func(tt *testing.T) {
		want := "b_2"
		if got := syms.VarName(testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{})))); got != want {
			tt.Errorf("VarName() = %s, want %s", got, want)
		}
	})

	t.Run("get pkg name", func(tt *testing.T) {
		want := ""
		if got := syms.PkgName("github.com/suger-131997/dein/internal/testpackages/b"); got != want {
			tt.Errorf("VarName() = %s, want %s", got, want)
		}
	})

	t.Run("get imports", func(tt *testing.T) {
		want := [][]string{{"a", "github.com/suger-131997/dein/internal/testpackages/a"}}
		if diff := cmp.Diff(syms.Imports(), want); diff != "" {
			tt.Errorf("Imports() is mismatch (-got +want):\n%s", diff)
		}
	})
}

func TestNewSymbols_DuplicatedVarName(t *testing.T) {
	syms := symbols.NewSymbols(
		"github.com/suger-131997/dein/internal/testpackages/b",
		[]component.Component{
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1_2{}))),
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&a.A1{}))),
		},
		[]string{},
	)

	t.Run("get dist pkg name", func(tt *testing.T) {
		want := []string{"a1", "a1_2", "a1_2_2"}

		got := make([]string, 0, len(want))
		got = append(got, syms.VarName(testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{})))))
		got = append(got, syms.VarName(testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&a.A1{})))))
		got = append(got, syms.VarName(testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1_2{})))))

		if diff := cmp.Diff(got, want); diff != "" {
			tt.Errorf("VarName() is mismatch (-got +want):\n%s", diff)
		}
	})
}
