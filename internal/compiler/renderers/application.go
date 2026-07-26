package renderers

import (
	"os"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/compiler/contracts"
	"go.yaml.in/yaml/v2"
)

// compiler/renderers/application.go

type Application struct {
	// ...
}

func NewApplication() *Application {
	return &Application{}
}

func (r *Application) Render(ctx *contracts.Context, cat *catalog.Catalog) error {

	appsDir := filepath.Join(ctx.Root, "applications")

	if err := os.MkdirAll(appsDir, 0755); err != nil {
		return err
	}

	for _, app := range cat.Applications {

		data, err := yaml.Marshal(app.Object)
		if err != nil {
			return err
		}

		filename := filepath.Join(appsDir, app.Object.AppName+".yaml")

		if err := os.WriteFile(filename, data, 0644); err != nil {
			return err
		}
	}

	return nil
}
