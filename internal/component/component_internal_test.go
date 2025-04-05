package component

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type (
	testComponent                        struct{}
	testGenericsComponent[T any]         struct{}
	testMultiGenericsComponent[T, U any] struct{}
)

func TestNewComponent(t *testing.T) {
	tests := []struct {
		name string

		in reflect.Type

		want    Component
		wantErr error
	}{
		{
			name: "testComponent",
			in:   reflect.TypeOf(testComponent{}),
			want: Component{
				name:    "testComponent",
				pkgPath: "github.com/suger-131997/dein/internal/component",
				prefix:  "",
			},
		},
		{
			name: "&testComponent",
			in:   reflect.TypeOf(&testComponent{}),
			want: Component{
				name:    "testComponent",
				pkgPath: "github.com/suger-131997/dein/internal/component",
				prefix:  "*",
			},
		},
		{
			name: "testGenericsComponent[testComponent]",
			in:   reflect.TypeOf(testGenericsComponent[testComponent]{}),
			want: Component{
				name:    "testGenericsComponent[github.com/suger-131997/dein/internal/component.testComponent]",
				pkgPath: "github.com/suger-131997/dein/internal/component",
				prefix:  "",
			},
		},
		{
			name: "testGenericsComponent[*testComponent]",
			in:   reflect.TypeOf(testGenericsComponent[*testComponent]{}),
			want: Component{
				name:    "testGenericsComponent[*github.com/suger-131997/dein/internal/component.testComponent]",
				pkgPath: "github.com/suger-131997/dein/internal/component",
				prefix:  "",
			},
		},
		{
			name: "[]testComponent",
			in:   reflect.TypeOf([]testComponent{}),
			want: Component{
				name:    "testComponent",
				pkgPath: "github.com/suger-131997/dein/internal/component",
				prefix:  "[]",
			},
		},
		{
			name: "**testComponent",
			in:   reflect.TypeOf(new(*testComponent)),
			want: Component{
				name:    "testComponent",
				pkgPath: "github.com/suger-131997/dein/internal/component",
				prefix:  "**",
			},
		},
		{
			name: "*[]*testComponent",
			in:   reflect.TypeOf(&[]*testComponent{}),
			want: Component{
				name:    "testComponent",
				pkgPath: "github.com/suger-131997/dein/internal/component",
				prefix:  "*[]*",
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
		{
			name:    "anonymous struct is not supported",
			in:      reflect.TypeOf(struct{ Name string }{}),
			wantErr: errors.New("anonymous struct is not supported"),
		},
		{
			name:    "anonymous struct for type param is not supported",
			in:      reflect.TypeOf(testGenericsComponent[struct{ Name string }]{}),
			wantErr: errors.New("anonymous struct for type param is not supported"),
		},
		{
			name:    "map type is not supported",
			in:      reflect.TypeOf(map[testComponent]testComponent{}),
			wantErr: errors.New("map type is not supported"),
		},
		{
			name:    "map type is not supported in type param",
			in:      reflect.TypeOf(testGenericsComponent[map[testComponent]testComponent]{}),
			wantErr: errors.New("map type is not supported in type param"),
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
					tt.Errorf("error mismatch: got %v, want %v", err, tc.wantErr)
				}

				return
			}

			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(Component{})); diff != "" {
				tt.Errorf("Component is mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestComponentTypeParams(t *testing.T) {
	tests := []struct {
		name string

		in reflect.Type

		want []TypeParam
	}{
		{
			name: "no type params",

			in: reflect.TypeOf(testComponent{}),

			want: []TypeParam{},
		},
		{
			name: "one type params",

			in: reflect.TypeOf(testGenericsComponent[testComponent]{}),

			want: []TypeParam{{name: "testComponent", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: ""}},
		},
		{
			name: "two type params",

			in: reflect.TypeOf(testMultiGenericsComponent[testComponent, int]{}),

			want: []TypeParam{
				{name: "testComponent", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: ""},
				{name: "int", pkgPath: "", prefix: ""},
			},
		},
		{
			name: "one pointer of pointer type params",

			in: reflect.TypeOf(testGenericsComponent[**testComponent]{}),

			want: []TypeParam{{name: "testComponent", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: "**"}},
		},
		{
			name: "two pointer type params",

			in: reflect.TypeOf(testMultiGenericsComponent[*testComponent, *int]{}),

			want: []TypeParam{
				{name: "testComponent", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: "*"},
				{name: "int", pkgPath: "", prefix: "*"},
			},
		},
		{
			name: "one array type params",

			in: reflect.TypeOf(testGenericsComponent[[]testComponent]{}),

			want: []TypeParam{{name: "testComponent", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: "[]"}},
		},
		{
			name: "one array pointer type params",

			in: reflect.TypeOf(testGenericsComponent[*[]*testComponent]{}),

			want: []TypeParam{{name: "testComponent", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: "*[]*"}},
		},
		{
			name: "nest type params",

			in: reflect.TypeOf(testGenericsComponent[testGenericsComponent[testGenericsComponent[int]]]{}),

			want: []TypeParam{{name: "testGenericsComponent[github.com/suger-131997/dein/internal/component.testGenericsComponent[int]]", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: ""}},
		},
		{
			name: "nest pointer type params",

			in: reflect.TypeOf(testGenericsComponent[testGenericsComponent[*testGenericsComponent[*int]]]{}),

			want: []TypeParam{{name: "testGenericsComponent[*github.com/suger-131997/dein/internal/component.testGenericsComponent[*int]]", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: ""}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			c, err := NewComponent(tc.in)
			if err != nil {
				tt.Fatalf("unexpected error: %v", err)
				return
			}

			got := c.TypeParams()

			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(TypeParam{})); diff != "" {
				tt.Errorf("TypeParams() is mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestTypeParamTypeParams(t *testing.T) {
	tests := []struct {
		name string

		in TypeParam

		want []TypeParam
	}{
		{
			name: "no type params",

			in: TypeParam{name: "testComponent", prefix: ""},

			want: []TypeParam{},
		},
		{
			name: "one type params",

			in: TypeParam{name: "testGenericsComponent[github.com/suger-131997/dein/internal/component.testComponent]", prefix: ""},

			want: []TypeParam{{name: "testComponent", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: ""}},
		},
		{
			name: "two type params",

			in: TypeParam{name: "testMultiGenericsComponent[github.com/suger-131997/dein/internal/component.testComponent,int]", prefix: ""},

			want: []TypeParam{
				{name: "testComponent", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: ""},
				{name: "int", pkgPath: "", prefix: ""},
			},
		},
		{
			name: "one pointer of pointer type params",

			in: TypeParam{name: "testGenericsComponent[**github.com/suger-131997/dein/internal/component.testComponent]", prefix: ""},

			want: []TypeParam{{name: "testComponent", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: "**"}},
		},
		{
			name: "two pointer type params",

			in: TypeParam{name: "testMultiGenericsComponent[*github.com/suger-131997/dein/internal/component.testComponent,*int]", prefix: ""},

			want: []TypeParam{
				{name: "testComponent", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: "*"},
				{name: "int", pkgPath: "", prefix: "*"},
			},
		},
		{
			name: "one array type params",

			in: TypeParam{name: "testGenericsComponent[[]github.com/suger-131997/dein/internal/component.testComponent]", prefix: ""},

			want: []TypeParam{{name: "testComponent", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: "[]"}},
		},
		{
			name: "nest type params",

			in: TypeParam{name: "testGenericsComponent[github.com/suger-131997/dein/internal/component.testGenericsComponent[github.com/suger-131997/dein/internal/component.testGenericsComponent[int]]]"},

			want: []TypeParam{{name: "testGenericsComponent[github.com/suger-131997/dein/internal/component.testGenericsComponent[int]]", pkgPath: "github.com/suger-131997/dein/internal/component", prefix: ""}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got := tc.in.TypeParams()

			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(TypeParam{})); diff != "" {
				tt.Errorf("TypeParams() is mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
