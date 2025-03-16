package dein

import (
	"github.com/suger-131997/dein/internal/provider"
	"reflect"
)

// Bind creates a new provider that binds a type T to an interface I.
func Bind[I any, T any]() *provider.Provider {
	var i *I
	var t T
	return provider.NewBindProvider(reflect.TypeOf(i).Elem(), reflect.TypeOf(t))
}
