package wheel

import "fmt"

type Wheel struct {
}

func NewWheel() Wheel {
	return Wheel{}
}

func (w Wheel) Spin() {
	fmt.Println("Wheel is spinning")
}
