package a

type (
	A1               struct{}
	A1_2             struct{}
	A2               struct{}
	A3[T any]        struct{}
	A4[T any, U any] struct{}
	IA1              interface {
		A1()
	}
)
type IA2[T any] interface{}

func NewA1() A1 {
	return A1{}
}

func NewA2(A1, *A1) A2 {
	return A2{}
}

func NewA3[T any]() A3[T] {
	return A3[T]{}
}
