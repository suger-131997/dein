package a

type A1 struct{}
type A2 struct{}
type A3[T any] struct{}
type A4[T any, U any] struct{}

func NewA1() A1 {
	return A1{}
}

func NewA3[T any]() A3[T] {
	return A3[T]{}
}
