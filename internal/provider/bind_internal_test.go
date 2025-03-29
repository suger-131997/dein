package provider

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/testpackages/a"
	"github.com/suger-131997/dein/internal/testpackages/b"
	"github.com/suger-131997/dein/internal/testutils"
)

func TestNewBindProvider(t *testing.T) {
	tests := []struct {
		name string

		bindTo    reflect.Type
		implement reflect.Type

		want    *Provider
		wantErr error
	}{
		{
			name: "b.B bind to a.IA1",

			bindTo: func() reflect.Type {
				var ia *a.IA1
				return reflect.TypeOf(ia).Elem()
			}(),
			implement: reflect.TypeOf(b.B{}),

			want: &Provider{
				in: []component.Component{testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(b.B{})))},
				out: func() component.Component {
					var ia *a.IA1
					return testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(ia).Elem()))
				}(),
				pkgPaths: []string{
					"github.com/suger-131997/dein/internal/testpackages/b",
					"github.com/suger-131997/dein/internal/testpackages/a",
				},
				markExposed: false,
				err:         nil,
			},
		},
		{
			name: "*b.B bind to a.IA1",

			bindTo: func() reflect.Type {
				var ia *a.IA1
				return reflect.TypeOf(ia).Elem()
			}(),
			implement: reflect.TypeOf(&b.B{}),

			want: &Provider{
				in: []component.Component{testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(&b.B{})))},
				out: func() component.Component {
					var ia *a.IA1
					return testutils.Must[component.Component](t)(component.NewComponent(reflect.TypeOf(ia).Elem()))
				}(),
				pkgPaths: []string{
					"github.com/suger-131997/dein/internal/testpackages/b",
					"github.com/suger-131997/dein/internal/testpackages/a",
				},
				markExposed: false,
				err:         nil,
			},
		},
		{
			name: "bind to not interface",

			bindTo:    reflect.TypeOf(b.B{}),
			implement: reflect.TypeOf(b.B{}),

			want:    nil,
			wantErr: errors.New("bind target must be an interface"),
		},
		{
			name: "bind to not implement interface",

			bindTo: func() reflect.Type {
				var ia *a.IA1
				return reflect.TypeOf(ia).Elem()
			}(),
			implement: reflect.TypeOf(a.A1{}),

			want:    nil,
			wantErr: errors.New("a.A1 must implement the interface a.IA1"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got := NewBindProvider(tc.bindTo, tc.implement)
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
