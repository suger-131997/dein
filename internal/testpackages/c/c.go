package c

import "github.com/suger-131997/dein/internal/testpackages/a"

type C struct{}

func NewC(ia1 a.IA1) C {
	return C{}
}

func NewC2(ia1 a.IA1, a4 a.A4[int, string]) (*C, error) {
	return &C{}, nil
}
