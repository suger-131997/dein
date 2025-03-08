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

	buildGenerator func(syms *symbols.Symbols, isInvoked bool) generator.Generator

	markInvoked bool

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

func (p *Provider) MarkInvoked() bool {
	return p.markInvoked
}

func (p *Provider) Generator(syms *symbols.Symbols) generator.Generator {
	return p.buildGenerator(syms, p.markInvoked)
}

func (p *Provider) CheckError() error {
	return p.err
}

func Mark(p *Provider) *Provider {
	p.markInvoked = true
	return p
}
