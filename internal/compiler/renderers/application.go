package renderers

import (
	"os"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/compiler/contracts"
	"github.com/celeguim/emp-cli/internal/compiler/manifests"
	"github.com/celeguim/emp-cli/internal/resolved"
	"go.yaml.in/yaml/v2"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (r *Application) Render(ctx *contracts.Context, cat *resolved.Catalog) error {

	appsDir := filepath.Join(ctx.Root, "applications")

	if err := os.MkdirAll(appsDir, 0755); err != nil {
		return err
	}

	for _, app := range cat.Applications {

		manifest := manifests.NewApplication(app)

		data, err := yaml.Marshal(manifest)
		if err != nil {
			return err
		}

		filename := filepath.Join(appsDir, app.Application.Name+".yaml")

		if err := os.WriteFile(filename, data, 0644); err != nil {
			return err
		}
	}

	return nil
}
