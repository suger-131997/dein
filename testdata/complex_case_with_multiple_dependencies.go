// Code generated by dein. DO NOT EDIT.
package main

import (
	a "github.com/suger-131997/dein/internal/testpackages/a"
	b_2 "github.com/suger-131997/dein/internal/testpackages/b"
	c_2 "github.com/suger-131997/dein/internal/testpackages/c"
)

type Container struct {
	c *c_2.C
}

func (c *Container) C() *c_2.C {
	return c.c
}

func NewContainer(
	a3 a.A3[int],
	__funcA4 func(a.A1) a.A4[int, string],
	__funcB func(a.A1, a.A3[int]) *b_2.B,
) (*Container, error) {
	__c := &Container{}

	a1 := a.NewA1()
	a4 := __funcA4(a1)
	b := __funcB(a1, a3)
	var iA1 a.IA1 = b
	c, err := c_2.NewC2(iA1, a4)
	if err != nil {
		return nil, err
	}
	__c.c = c

	return __c, nil
}
