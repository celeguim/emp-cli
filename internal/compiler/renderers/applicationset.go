package renderers

import (
	"github.com/celeguim/emp-cli/internal/compiler/contracts"
	"github.com/celeguim/emp-cli/internal/compiler/manifests"
	"github.com/celeguim/emp-cli/internal/resolved"
)

type ApplicationSet struct{}

func NewApplicationSet() contracts.Renderer {
	return &ApplicationSet{}
}

func (r *ApplicationSet) Render(
	ctx *contracts.Context,
	cat *resolved.Catalog,
) error {
	for _, appset := range cat.ApplicationSets {
		manifest := manifests.NewApplicationSet(appset)
		if err := ctx.Writer.Write(applicationSetKind, appset.Name, manifest); err != nil {
			return err
		}
	}

	return nil
}
