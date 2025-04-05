package generator

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/testpackages/a"
	"github.com/suger-131997/dein/internal/testpackages/b"
	"github.com/suger-131997/dein/internal/testpackages/c"
	"github.com/suger-131997/dein/internal/testutils"
)

func TestWriteTypeParams(t *testing.T) {
	tests := []struct {
		name string

		in component.Component

		want string
	}{
		{
			name: "no type params",

			in: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),

			want: "",
		},
		{
			name: "one type params",

			in: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[b.B]{}))),

			want: "[b.B]",
		},
		{
			name: "one pointer type params",

			in: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[*b.B]{}))),

			want: "[*b.B]",
		},
		{
			name: "one pointer of pointer type params",

			in: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[**b.B]{}))),

			want: "[**b.B]",
		},
		{
			name: "one array type params",

			in: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[[]b.B]{}))),

			want: "[[]b.B]",
		},
		{
			name: "one array pointer type params",

			in: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[*[]b.B]{}))),

			want: "[*[]b.B]",
		},
		{
			name: "two type params",

			in: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A4[b.B, c.C]{}))),

			want: "[b.B, c.C]",
		},
		{
			name: "build-in type params",

			in: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[int]{}))),

			want: "[int]",
		},
		{
			name: "build-in pointer type params",

			in: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[*int]{}))),

			want: "[*int]",
		},
		{
			name: "nest type params",

			in: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[a.A3[b.B]]{}))),

			want: "[a.A3[b.B]]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			var out strings.Builder

			writeTypeParams(&out, symbols.NewSymbols([]component.Component{tc.in}, []string{}), tc.in.TypeParams())
			got := out.String()

			if diff := cmp.Diff(got, tc.want); diff != "" {
				tt.Errorf("writeTypeParams() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
