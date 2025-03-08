package dein

import (
	"github.com/suger-131997/dein/internal/provider"
)

func P0[R any](f func() R) *provider.Provider {
	return provider.NewConstructorProvider(f, false)
}

func P1[R, T1 any](f func(T1) R) *provider.Provider {
	return provider.NewConstructorProvider(f, false)
}

func PE1[R, T1 any](f func(T1) (R, error)) *provider.Provider {
	return provider.NewConstructorProvider(f, true)
}

func P2[R, T1, T2 any](f func(T1, T2) R) *provider.Provider {
	return provider.NewConstructorProvider(f, false)
}
