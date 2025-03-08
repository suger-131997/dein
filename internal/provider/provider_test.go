package provider_test

import (
	"github.com/suger-131997/dein/internal/provider"
	"github.com/suger-131997/dein/internal/testpackages/a"
	"testing"
)

func TestMark(t *testing.T) {
	p := provider.NewConstructorProvider(a.NewA1, false)

	if p.MarkInvoked() {
		t.Error("MarkInvoked() should return false")
	}

	p = provider.Mark(p)

	if !p.MarkInvoked() {
		t.Error("MarkInvoked() should return true")
	}
}
