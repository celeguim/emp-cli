package renderers

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/compiler/contracts"
	"github.com/celeguim/emp-cli/internal/compiler/manifests"
	"github.com/celeguim/emp-cli/internal/resolved"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (r *Application) Render(ctx *contracts.Context, cat *resolved.Catalog) error {
	fmt.Println("render applications")

	for _, app := range cat.Applications {
		manifest := manifests.NewApplication(app)
		if err := ctx.Writer.Write(applicationKind, app.Application.Name, manifest); err != nil {
			return err
		}
	}

	return nil
}
