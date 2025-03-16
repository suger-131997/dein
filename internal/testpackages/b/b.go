package b

import "github.com/suger-131997/dein/internal/testpackages/a"

type B struct{}

func NewB(a a.A1) *B {
	return &B{}
}

func (b B) A1() {}

var _ a.IA1 = B{}
