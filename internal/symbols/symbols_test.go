package symbols_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/testpackages/a"
	a2 "github.com/suger-131997/dein/internal/testpackages/a/a"
	"github.com/suger-131997/dein/internal/testpackages/b"
	"github.com/suger-131997/dein/internal/testpackages/c"
	"github.com/suger-131997/dein/internal/testutils"
	"reflect"
	"testing"
)

func TestNewSymbols(t *testing.T) {
	syms := symbols.NewSymbols(
		[]component.Component{
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[b.B]{}))),
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a2.A1{}))),
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
			testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A4[int, c.C]{}))),
		},
		[]string{"github.com/suger-131997/dein/internal/testpackages/x"},
	)

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
		wantName := []string{"a", "a_2", "b_2", "c", "x"}
		wantPath := []string{
			"github.com/suger-131997/dein/internal/testpackages/a",
			"github.com/suger-131997/dein/internal/testpackages/a/a",
			"github.com/suger-131997/dein/internal/testpackages/b",
			"github.com/suger-131997/dein/internal/testpackages/c",
			"github.com/suger-131997/dein/internal/testpackages/x",
		}

		gotName := make([]string, 0, len(wantName))
		gotPath := make([]string, 0, len(wantPath))
		for name, path := range syms.Imports() {
			gotName = append(gotName, name)
			gotPath = append(gotPath, path)
		}

		if diff := cmp.Diff(gotName, wantName); diff != "" {
			tt.Errorf("Imports().name is mismatch (-got +want):\n%s", diff)
		}
		if diff := cmp.Diff(gotPath, wantPath); diff != "" {
			tt.Errorf("Imports().path is mismatch (-got +want):\n%s", diff)
		}
	})
}
