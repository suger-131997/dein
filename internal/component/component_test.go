package component_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/testpackages/a"
	a2 "github.com/suger-131997/dein/internal/testpackages/a/a"
	"github.com/suger-131997/dein/internal/testpackages/b"
	"github.com/suger-131997/dein/internal/testutils"
	"reflect"
	"testing"
)

func TestComponentLess(t *testing.T) {
	tests := []struct {
		name string

		c1 component.Component
		c2 component.Component

		want bool
	}{
		{
			name: "c1 < c2: c1 = a.A1, c2 = b.B",

			c1: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			c2: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),

			want: true,
		},
		{
			name: "c1 > c2: c1 = b.B, c2 = a.A1",

			c1: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{}))),
			c2: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),

			want: false,
		},
		{
			name: "c1 < c2: c1 = a.A1, c2 = a.A2",

			c1: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			c2: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A2{}))),

			want: true,
		},
		{
			name: "c1 > c2: c1 = a.A2, c2 = a.A1",

			c1: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A2{}))),
			c2: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),

			want: false,
		},
		{
			name: "c1 < c2: c1 = a.A1, c2 = *a.A1",

			c1: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
			c2: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&a.A1{}))),

			want: true,
		},
		{
			name: "c1 > c2: c1 = a.A1, c2 = *a.A1",

			c1: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&a.A1{}))),
			c2: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),

			want: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got := tc.c1.Less(tc.c2)

			if got != tc.want {
				tt.Errorf("Less() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestComponentName(t *testing.T) {
	tests := []struct {
		name string

		component component.Component

		want string
	}{
		{
			name: "a.A1",

			component: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),

			want: "A1",
		},
		{
			name: "a.A3[a.A1]",

			component: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[a.A1]{}))),

			want: "A3",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got := tc.component.Name()

			if got != tc.want {
				tt.Errorf("Name() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestComponentPkgPaths(t *testing.T) {
	tests := []struct {
		name string

		component component.Component

		want []string
	}{
		{
			name: "no type parameter",

			component: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),

			want: []string{"github.com/suger-131997/dein/internal/testpackages/a"},
		},
		{
			name: "one type parameter",

			component: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[b.B]{}))),

			want: []string{
				"github.com/suger-131997/dein/internal/testpackages/a",
				"github.com/suger-131997/dein/internal/testpackages/b",
			},
		},
		{
			name: "twe type parameter",

			component: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A4[b.B, a2.A1]{}))),

			want: []string{
				"github.com/suger-131997/dein/internal/testpackages/a",
				"github.com/suger-131997/dein/internal/testpackages/b",
				"github.com/suger-131997/dein/internal/testpackages/a/a",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got := tc.component.PkgPaths()

			if diff := cmp.Diff(got, tc.want, cmpopts.SortSlices(func(i, j string) bool {
				return i < j
			})); diff != "" {
				tt.Errorf("PkgPaths() is mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
