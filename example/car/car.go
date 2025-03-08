package car

import (
	"github.com/suger-131997/dein/example/car/engine"
	"github.com/suger-131997/dein/example/car/wheel"
)

type Car struct {
	engine *engine.Engine
	wheel  wheel.Wheel
}

func NewCar(e *engine.Engine, w wheel.Wheel) *Car {
	return &Car{
		engine: e,
		wheel:  w,
	}
}

func (c *Car) Run() {
	c.engine.Start()
	c.wheel.Spin()
}
