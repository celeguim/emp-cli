package renderers

import (
	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/compiler/contracts"
)

type ClusterRenderer struct {
	// ...
}

func NewClusterRenderer() *ClusterRenderer {
	return &ClusterRenderer{}
}

func (r *ClusterRenderer) Render(
	ctx *contracts.Context,
	cat *catalog.Catalog,
) error {

	// usa ctx.Root
	// percorre cat.Applications
	// escreve arquivos

	return nil
}
