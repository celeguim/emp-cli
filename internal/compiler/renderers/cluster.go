package renderers

import (
	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/compiler/contracts"
)

type ClusterRenderer struct {
	// ...
}

func NewCluster() *ClusterRenderer {
	return &ClusterRenderer{}
}

func (r *ClusterRenderer) Render(
	ctx *contracts.Context,
	cat *catalog.Catalog,
) error {

	return nil
}
