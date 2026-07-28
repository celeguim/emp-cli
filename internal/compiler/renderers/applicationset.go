package renderers

import (
	"os"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/compiler/contracts"
	"github.com/celeguim/emp-cli/internal/compiler/manifests"
	"github.com/celeguim/emp-cli/internal/resolved"
	"go.yaml.in/yaml/v2"
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

		data, err := yaml.Marshal(manifest)
		if err != nil {
			return err
		}

		filename := filepath.Join(
			ctx.Root,
			"applicationsets",
			appset.Name+".yaml",
		)

		if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
			return err
		}

		if err := os.WriteFile(filename, data, 0644); err != nil {
			return err
		}
	}

	return nil
}
