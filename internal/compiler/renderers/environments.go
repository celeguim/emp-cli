package renderers

import (
	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/compiler/contracts"
)

type EnvironmentRenderer struct {
	// ...
}

func NewEnvironmentRenderer() *EnvironmentRenderer {
	return &EnvironmentRenderer{}
}

func (r *EnvironmentRenderer) Render(
	ctx *contracts.Context,
	cat *catalog.Catalog,
) error {

	// usa ctx.Root
	// percorre cat.Applications
	// escreve arquivos

	return nil
}
