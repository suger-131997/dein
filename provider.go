package dein

import (
	"github.com/suger-131997/dein/internal/provider"
	"reflect"
)

// Bind creates a new provider that binds a type T to an interface I.
func Bind[I any, T any]() *provider.Provider {
	var i I
	var t T
	return provider.NewBindProvider(reflect.TypeOf(&i).Elem(), reflect.TypeOf(t))
}

// P0 creates a new provider that initializes with a constructor that takes no arguments and returns a value of type R.

// PF2 creates a new provider that initializes with a function provided at build time, which takes no arguments and returns a value of type R.

func rt[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}
