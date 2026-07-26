package renderers

import (
	"fmt"

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
	fmt.Println("ClusterRenderer")

	// usa ctx.Root
	// percorre cat.Applications
	// escreve arquivos

	return nil
}
