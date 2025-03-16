package main

import (
	"github.com/suger-131997/dein/example/car/engine"
	"log"
)

func main() {
	c, err := NewContainer(engine.EngineTypeGasoline)
	if err != nil {
		log.Fatal(err)
	}

	myCar := c.ICar
	myCar.Run()
}
