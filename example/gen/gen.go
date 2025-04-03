//go:generate go run . -output=../dein_gen.go

package main

import (
	"flag"
	"log"
	"os"

	"golang.org/x/tools/imports"

	"github.com/suger-131997/dein"
	"github.com/suger-131997/dein/example/car"
	"github.com/suger-131997/dein/example/car/engine"
	"github.com/suger-131997/dein/example/car/light"
	"github.com/suger-131997/dein/example/car/wheel"
)

var filename = flag.String("output", "dein_gen.go", "output file name")

func main() {
	flag.Parse()

	// Create a resolver
	r := dein.NewResolver()

	// Register the providers
	dein.Register(r, dein.Mark(dein.Bind[car.ICar, *car.Car]()))
	dein.Register(r, dein.PE1(engine.NewEngine))
	dein.Register(r, dein.P0(wheel.NewWheel))
	dein.Register(r, dein.P3(car.NewCar))
	dein.Register(r, dein.PF0[light.Light]())

	// Dependency resolution and create a generator
	gen, err := r.Resolve()
	if err != nil {
		log.Fatal(err)
	}

	// Generate code
	data, err := gen.Generate("main")
	if err != nil {
		log.Fatal(err)
	}

	// Format code
	data, err = imports.Process(*filename, data, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Write code to file
	err = os.WriteFile(*filename, data, 0o644)
	if err != nil {
		log.Fatal(err)
	}
}
