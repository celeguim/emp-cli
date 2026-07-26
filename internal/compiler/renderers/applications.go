package renderers

import (
	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/compiler/contracts"
)

// compiler/renderers/application.go

type Application struct {
	// ...
}

func NewApplicationRenderer() *Application {
	return &Application{}
}

func (r *Application) Render(
	ctx *contracts.Context,
	cat *catalog.Catalog,
) error {

	// usa ctx.Root
	// percorre cat.Applications
	// escreve arquivos

	return nil
}
