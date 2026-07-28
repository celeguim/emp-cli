package resolver

import (
	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/resolved"
)

type Resolver struct{}

func New() *Resolver {
	return &Resolver{}
}

func Resolve(cat *catalog.Catalog) (*resolved.Catalog, error) {

	rc := &resolved.Catalog{}

	if err := resolveApplications(cat, rc); err != nil {
		return nil, err
	}

	if err := resolveProjects(rc); err != nil {
		return nil, err
	}

	if err := resolveApplicationSets(rc); err != nil {
		return nil, err
	}

	return rc, nil
}
