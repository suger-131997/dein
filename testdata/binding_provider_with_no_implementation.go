// Code generated by dein. DO NOT EDIT.
package main

import (
	a "github.com/suger-131997/dein/internal/testpackages/a"
	b_2 "github.com/suger-131997/dein/internal/testpackages/b"
)

type Container struct {
}

func NewContainer(
	b b_2.B,
) (*Container, error) {
	__c := &Container{}

	var iA1 a.IA1 = b

	return __c, nil
}
