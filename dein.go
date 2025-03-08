package dein

import "github.com/suger-131997/dein/internal/provider"

// Register registers a provider to the resolver.
func Register(r *Resolver, p *provider.Provider) {
	r.providers = append(r.providers, p)
}

// Mark marks a provider to be included in the container.
func Mark(p *provider.Provider) *provider.Provider {
	provider.Mark(p)
	return p
}
