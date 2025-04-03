package dein_test

import (
	"testing"

	"github.com/suger-131997/dein"
)

type (
	A struct{}
	B struct{}
)

func NewA(b *B) *A {
	return &A{}
}

func NewB(a *A) *B {
	return &B{}
}

func TestResolver(t *testing.T) {
	t.Run("circular dependency", func(tt *testing.T) {
		r := dein.NewResolver()

		dein.Register(r, dein.P1(NewA))
		dein.Register(r, dein.P1(NewB))

		_, err := r.Resolve()
		if err == nil {
			t.Error("expected error, got nil")
		} else if err.Error() != "circular dependency detected" {
			t.Errorf("expected 'circular dependency detected', got '%s'", err.Error())
		}
	})
	t.Run("duplicate component provided", func(tt *testing.T) {
		r := dein.NewResolver()

		dein.Register(r, dein.P1(NewA))
		dein.Register(r, dein.P1(NewA))

		_, err := r.Resolve()
		if err == nil {
			t.Error("expected error, got nil")
		} else if err.Error() != "duplicate component provided" {
			t.Errorf("expected 'duplicate component provided', got '%s'", err.Error())
		}
	})
}
