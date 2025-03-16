package car

import (
	"github.com/suger-131997/dein/example/car/engine"
	"github.com/suger-131997/dein/example/car/light"
	"github.com/suger-131997/dein/example/car/wheel"
)

type ICar interface {
	Run()
}

type Car struct {
	engine *engine.Engine
	wheel  wheel.Wheel
	light  light.Light
}

func NewCar(e *engine.Engine, w wheel.Wheel, l light.Light) *Car {
	return &Car{
		engine: e,
		wheel:  w,
		light:  l,
	}
}

func (c *Car) Run() {
	c.engine.Start()
	c.wheel.Spin()
	c.light.LightOn()
}
