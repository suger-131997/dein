//go:generate go run . -output=../dein_gen.go

package main

import (
	"flag"
	"github.com/suger-131997/dein"
	"github.com/suger-131997/dein/example/car"
	"github.com/suger-131997/dein/example/car/engine"
	"github.com/suger-131997/dein/example/car/wheel"
	"golang.org/x/tools/imports"
	"log"
	"os"
)

var filename = flag.String("output", "dein_gen.go", "output file name")

func main() {
	flag.Parse()
	r := dein.NewResolver()

	dein.Register(r, dein.Mark(dein.Bind[car.ICar, *car.Car]()))
	dein.Register(r, dein.PE1(engine.NewEngine))
	dein.Register(r, dein.P0(wheel.NewWheel))
	dein.Register(r, dein.P2(car.NewCar))

	gen, err := r.Resolve()
	if err != nil {
		log.Fatal(err)
	}

	data, err := gen.Generate("main")
	if err != nil {
		log.Fatal(err)
	}

	data, err = imports.Process(*filename, data, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(*filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
