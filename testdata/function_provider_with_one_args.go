// Code generated by dein. DO NOT EDIT.
package main

import (
	a "github.com/suger-131997/dein/internal/testpackages/a"
	b_2 "github.com/suger-131997/dein/internal/testpackages/b"
)

type Container struct {
	b b_2.B
}

func (c *Container) B() b_2.B {
	return c.b
}

func NewContainer(
	a3 a.A3[int],
	__funcB func(a.A1, a.A3[int]) b_2.B,
) (*Container, error) {
	__c := &Container{}

	a1 := a.NewA1()
	b := __funcB(a1, a3)
	__c.b = b

	return __c, nil
}
