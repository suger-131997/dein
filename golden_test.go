package dein_test

import (
	"go/format"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/suger-131997/dein"
	"github.com/suger-131997/dein/internal/testpackages/a"
	"github.com/suger-131997/dein/internal/testpackages/b"
	"github.com/suger-131997/dein/internal/testpackages/c"
)

func filename(n string) string {
	return "testdata/" + strings.ReplaceAll(n, " ", "_") + ".go"
}

func TestGolden(t *testing.T) {
	mode, err := strconv.ParseBool(os.Getenv("WRITE_GOLDEN_FILE_MODE"))
	if err != nil {
		mode = false
	}

	tests := []struct {
		name string

		pkgName  string
		register func(r *dein.Resolver)
	}{
		{
			name: "constructor provider with no args",

			pkgName: "main",
			register: func(r *dein.Resolver) {
				dein.Register(r, dein.Mark(dein.P0(a.NewA1)))
			},
		},
		{
			name:    "constructor provider with one args",
			pkgName: "main",
			register: func(r *dein.Resolver) {
				dein.Register(r, dein.Mark(dein.P1(b.NewB)))
				dein.Register(r, dein.P0(a.NewA1))
			},
		},
		{
			name:    "constructor provider with two args",
			pkgName: "main",
			register: func(r *dein.Resolver) {
				dein.Register(r, dein.Mark(dein.P2(a.NewA2)))
				dein.Register(r, dein.P0(a.NewA1))
			},
		},
		{
			name: "binding provider with no implementation",

			pkgName: "main",
			register: func(r *dein.Resolver) {
				dein.Register(r, dein.Bind[a.IA1, b.B]())
			},
		},
		{
			name: "binding provider with implementation",

			pkgName: "main",
			register: func(r *dein.Resolver) {
				dein.Register(r, dein.Bind[a.IA1, b.B]())
				dein.Register(r, dein.P1(b.NewB))
			},
		},
		{
			name: "function provider with no args",

			pkgName: "main",
			register: func(r *dein.Resolver) {
				dein.Register(r, dein.Mark(dein.PF2[a.A1, a.A3[int], b.B]()))
			},
		},
		{
			name:    "function provider with one args",
			pkgName: "main",
			register: func(r *dein.Resolver) {
				dein.Register(r, dein.Mark(dein.PF2[a.A1, a.A3[int], b.B]()))
				dein.Register(r, dein.P0(a.NewA1))
			},
		},
		{
			name:    "complex case with multiple dependencies",
			pkgName: "main",
			register: func(r *dein.Resolver) {
				dein.Register(r, dein.PF2[a.A1, a.A3[int], *b.B]())
				dein.Register(r, dein.PF1[a.A1, a.A4[int, string]]())
				dein.Register(r, dein.P0(a.NewA1))
				dein.Register(r, dein.Bind[a.IA1, *b.B]())
				dein.Register(r, dein.Mark(dein.PE2(c.NewC2)))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			r := dein.NewResolver()
			tc.register(r)

			gen, err := r.Resolve()
			if err != nil {
				tt.Fatalf("Resolve() unexpeted error: %v", err)
			}

			raw, err := gen.Generate(tc.pkgName)
			if err != nil {
				tt.Fatalf("Generate() unexpeted error: %v", err)
			}

			got, err := format.Source(raw)
			if err != nil {
				tt.Fatalf("format.Source() unexpeted error: %v", err)
			}

			golden := filename(tc.name)
			if mode {
				if err := os.WriteFile(golden, got, 0o777); err != nil {
					tt.Fatalf("os.WriteFile() unexpeted error: %v", err)
				}

				tt.Skip("write golden file")
			}

			want, err := os.ReadFile(golden)
			if err != nil {
				tt.Fatalf("os.ReadFile() unexpeted error: %v", err)
			}

			if diff := cmp.Diff(string(got), string(want)); diff != "" {
				tt.Errorf("diff (-got +want):\n%s", diff)
			}
		})
	}
}
