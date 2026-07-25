package resolver

import "github.com/celeguim/emp-cli/internal/catalog"

type Resolver struct{}

func New() *Resolver {
	return &Resolver{}
}

func (r *Resolver) Resolve(cat *catalog.Catalog) (*catalog.Catalog, error) {
	if err := r.resolveClusters(cat); err != nil {
		return nil, err
	}

	return cat, nil
}
