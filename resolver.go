package dein

import (
	"errors"
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/generator"
	"github.com/suger-131997/dein/internal/provider"
	"github.com/suger-131997/dein/internal/symbols"
	"github.com/suger-131997/dein/internal/utils"
	"sort"
)

// Resolver is a struct that resolves the dependency graph of the providers.
type Resolver struct {
	providers []*provider.Provider
}

// NewResolver creates a new Resolver.
func NewResolver() *Resolver {
	return &Resolver{}
}

// Resolve resolves the dependency graph of the providers and returns a source code generator.
// It uses topological sorting to ensure the correct order of provider initialization.
func (r *Resolver) Resolve() (*Generator, error) {
	for _, p := range r.providers {
		if err := p.CheckError(); err != nil {
			return nil, err
		}
	}

	providers := make(map[component.Component]*provider.Provider)
	graph := make(map[component.Component][]component.Component)
	indegrees := make(map[component.Component]int)
	for _, p := range r.providers {
		if _, ok := graph[p.Out()]; ok {
			return nil, errors.New("duplicate component provided")
		}
		providers[p.Out()] = p
		graph[p.Out()] = make([]component.Component, 0)
		indegrees[p.Out()] = 0
	}

	argumentComponents := make([]component.Component, 0)
	for _, p := range r.providers {
		for _, in := range p.In() {
			if _, ok := graph[in]; !ok {
				argumentComponents = append(argumentComponents, in)
				continue
			}
			graph[in] = append(graph[in], p.Out())
			indegrees[p.Out()]++
		}
	}
	sort.Slice(argumentComponents, func(i, j int) bool {
		return argumentComponents[i].Less(argumentComponents[j])
	})

	resolvedProviders := make([]*provider.Provider, 0, len(graph))

	pq := utils.NewPriorityQueue(func(i, j *provider.Provider) bool {
		return i.Out().Less(j.Out())
	})

	for c, i := range indegrees {
		if i == 0 {
			utils.Push(pq, providers[c])
		}
	}

	for pq.Len() > 0 {
		from := utils.Pop(pq)

		resolvedProviders = append(resolvedProviders, from)

		for _, to := range graph[from.Out()] {
			indegrees[to]--
			if indegrees[to] == 0 {
				utils.Push(pq, providers[to])
			}
		}
	}

	components := make([]component.Component, 0)
	pkgPaths := make([]string, 0)
	for _, p := range providers {
		components = append(components, p.In()...)
		components = append(components, p.Out())
		pkgPaths = append(pkgPaths, p.PkgPaths()...)
	}
	syms := symbols.NewSymbols(components, pkgPaths)

	containerComponents := make([]component.Component, 0)
	generators := make([]generator.Generator, 0, len(resolvedProviders))
	for _, p := range resolvedProviders {
		if p.MarkInvoked() {
			containerComponents = append(containerComponents, p.Out())
		}
		generators = append(generators, p.Generator(syms))
	}

	sort.Slice(containerComponents, func(i, j int) bool {
		return containerComponents[i].Less(containerComponents[j])
	})

	return &Generator{
		symbols:             syms,
		containerComponents: containerComponents,
		argumentComponents:  argumentComponents,
		generators:          generators,
	}, nil
}
