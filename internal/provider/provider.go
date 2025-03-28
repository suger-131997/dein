package provider

import (
	"github.com/suger-131997/dein/internal/component"
	"github.com/suger-131997/dein/internal/generator"
	"github.com/suger-131997/dein/internal/symbols"
)

type Provider struct {
	in  []component.Component
	out component.Component

	pkgPaths []string

	buildGenerator func(syms *symbols.Symbols, markExposed bool) generator.BodyGenerator

	markExposed bool

	err error
}

func (p *Provider) In() []component.Component {
	return p.in
}

func (p *Provider) Out() component.Component {
	return p.out
}

func (p *Provider) PkgPaths() []string {
	return p.pkgPaths
}

func (p *Provider) MarkExposed() bool {
	return p.markExposed
}

func (p *Provider) Generator(syms *symbols.Symbols) generator.BodyGenerator {
	return p.buildGenerator(syms, p.markExposed)
}

func (p *Provider) CheckError() error {
	return p.err
}

func Mark(p *Provider) *Provider {
	p.markExposed = true
	return p
}
