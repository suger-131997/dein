package provider_test

import (
	"github.com/suger-131997/dein/internal/provider"
	"github.com/suger-131997/dein/internal/testpackages/a"
	"testing"
)

func TestMark(t *testing.T) {
	p := provider.NewConstructorProvider(a.NewA1, false)

	if p.MarkExposed() {
		t.Error("MarkExposed() should return false")
	}

	p = provider.Mark(p)

	if !p.MarkExposed() {
		t.Error("MarkExposed() should return true")
	}
}
