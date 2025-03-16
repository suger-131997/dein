package main

import (
	"github.com/suger-131997/dein/example/car/engine"
	"github.com/suger-131997/dein/example/car/light"
	"log"
)

func main() {
	c, err := NewContainer(
		engine.EngineTypeGasoline,
		func() light.Light {
			return light.Light{LightType: light.LightTypeTail}
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	myCar := c.ICar
	myCar.Run()
}
