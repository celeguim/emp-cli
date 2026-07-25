package resolver

import "github.com/celeguim/emp-cli/internal/catalog"

type Resolver struct{}

func New() *Resolver {
	return &Resolver{}
}

func (r *Resolver) Resolve(cat *catalog.Catalog) (*catalog.Catalog, error) {
	return cat, nil
}
