package provider

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/testpackages/a"
	"github.com/suger-131997/dein/internal/testpackages/b"
	"github.com/suger-131997/dein/internal/testpackages/c"
	"github.com/suger-131997/dein/internal/testutils"
	"reflect"
	"testing"
)

func TestNewConstructorProvider(t *testing.T) {
	tests := []struct {
		name string

		f        any
		hasError bool

		want    *Provider
		wantErr error
	}{
		{
			name: "no args function",

			f:        a.NewA1,
			hasError: false,

			want: &Provider{
				in:          []component.Component{},
				out:         testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{}))),
				pkgPaths:    []string{"github.com/suger-131997/dein/internal/testpackages/a"},
				markInvoked: false,
				err:         nil,
			},
			wantErr: nil,
		},
		{
			name: "one args function",

			f:        b.NewB,
			hasError: false,

			want: &Provider{
				in:  []component.Component{testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A1{})))},
				out: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&b.B{}))),
				pkgPaths: []string{
					"github.com/suger-131997/dein/internal/testpackages/a",
					"github.com/suger-131997/dein/internal/testpackages/b",
				},
				markInvoked: false,
				err:         nil,
			},
		},
		{
			name: "interface args function",

			f:        c.NewC,
			hasError: false,

			want: &Provider{
				in: []component.Component{testutils.Must[component.Component](t)(component.NewComponent(func() reflect.Type {
					var i a.IA1
					return reflect.TypeOf(&i).Elem()
				}()))},
				out: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(c.C{}))),
				pkgPaths: []string{
					"github.com/suger-131997/dein/internal/testpackages/a",
					"github.com/suger-131997/dein/internal/testpackages/c",
				},
				markInvoked: false,
				err:         nil,
			},
		},
		{
			name: "generics function",

			f:        a.NewA3[b.B],
			hasError: false,

			want: &Provider{
				in:  []component.Component{},
				out: testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(a.A3[b.B]{}))),
				pkgPaths: []string{
					"github.com/suger-131997/dein/internal/testpackages/b",
					"github.com/suger-131997/dein/internal/testpackages/a",
				},
				markInvoked: false,
				err:         nil,
			},
		},
		{
			name: "anonymous function",

			f:        func() a.A1 { return a.A1{} },
			hasError: false,

			want:    nil,
			wantErr: errors.New("anonymous function is not allowed"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got := NewConstructorProvider(tc.f, tc.hasError)
			if err := got.CheckError(); err != nil {
				if tc.wantErr == nil {
					tt.Errorf("unexpected error: %v", err)
					return
				}
				if err.Error() != tc.wantErr.Error() {
					tt.Errorf("error mismatch: got %v, want %v", err, tc.wantErr)
				}
				return
			}

			if diff := cmp.Diff(got, tc.want,
				cmp.AllowUnexported(Provider{}, component.Component{}),
				cmpopts.IgnoreFields(Provider{}, "buildGenerator"),
			); diff != "" {
				tt.Errorf("Provider is mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
