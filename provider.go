package dein

import (
	"reflect"

	"github.com/suger-131997/dein/internal/provider"
)

// Bind creates a new provider that binds a type T to an interface I.
func Bind[I, T any]() *provider.Provider {
	var i I

	var t T

	return provider.NewBindProvider(reflect.TypeOf(&i).Elem(), reflect.TypeOf(t))
}

func rt[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}
