package component

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

type testComponent struct {
}

type testGenericsComponent[T any] struct {
}

func TestNewComponent(t *testing.T) {
	tests := []struct {
		name string

		in reflect.Type

		want    Component
		wantErr error
	}{
		{
			name: "testComponent is value",
			in:   reflect.TypeOf(testComponent{}),
			want: Component{
				name:      "testComponent",
				pkgPath:   "github.com/suger-131997/dein/internal/component",
				isPointer: false,
			},
		},
		{
			name: "testComponent is pointer",
			in:   reflect.TypeOf(&testComponent{}),
			want: Component{
				name:      "testComponent",
				pkgPath:   "github.com/suger-131997/dein/internal/component",
				isPointer: true,
			},
		},
		{
			name: "testGenericsComponent[testComponent] is value",
			in:   reflect.TypeOf(testGenericsComponent[testComponent]{}),
			want: Component{
				name:      "testGenericsComponent[github.com/suger-131997/dein/internal/component.testComponent]",
				pkgPath:   "github.com/suger-131997/dein/internal/component",
				isPointer: false,
			},
		},
		{
			name:    "builtin type value is not supported",
			in:      reflect.TypeOf(0),
			wantErr: errors.New("builtin type is not supported"),
		},
		{
			name: "builtin type pinter is not supported",
			in: reflect.TypeOf(func() *int {
				i := 0
				return &i
			}()),
			wantErr: errors.New("builtin type is not supported"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got, err := NewComponent(tc.in)
			if err != nil {
				if tc.wantErr == nil {
					tt.Errorf("unexpected error: %v", err)
					return
				}
				if err.Error() != tc.wantErr.Error() {
					tt.Fatalf("error mismatch: got %v, want %v", err, tc.wantErr)
				}
				return
			}

			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(Component{})); diff != "" {
				tt.Errorf("Component is mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
