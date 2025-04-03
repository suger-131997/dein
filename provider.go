package dein

import (
	"reflect"

	"github.com/suger-131997/dein/internal/provider"
)

// Bind creates a new provider that binds a type T to an interface I.
func Bind[T, I any]() *provider.Provider {
	var t T

	var i I

	return provider.NewBindProvider(reflect.TypeOf(&i).Elem(), reflect.TypeOf(t))
}

func rt[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}
