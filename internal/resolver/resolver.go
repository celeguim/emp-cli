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

	for _, app := range cat.Applications {

		ra := resolved.ResolvedApplication{
			Application: app.Object,
		}

		// TODO:
		// localizar Environment

		// TODO:
		// localizar Cluster

		rc.Applications = append(rc.Applications, ra)
	}

	return rc, nil
}
