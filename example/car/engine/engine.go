package engine

import (
	"errors"
	"fmt"
)

type EngineType int

const (
	EngineTypeUnknown EngineType = iota
	EngineTypeGasoline
	EngineTypeHydrogen
)

func (t EngineType) String() string {
	switch t {
	case EngineTypeGasoline:
		return "Gasoline"
	case EngineTypeHydrogen:
		return "Hydrogen"
	default:
		return "Unknown"
	}
}

type Engine struct {
	engineType EngineType
}

func NewEngine(t EngineType) (*Engine, error) {
	if t == EngineTypeUnknown {
		return nil, errors.New("engine type is unknown")
	}

	return &Engine{
		engineType: t,
	}, nil
}

func (e *Engine) Start() {
	fmt.Printf("starting the %s engine\n", e.engineType)
}
