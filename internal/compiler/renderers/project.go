package renderers

import (
	"os"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/compiler/contracts"
	"github.com/celeguim/emp-cli/internal/compiler/manifests"
	"github.com/celeguim/emp-cli/internal/resolved"
	"go.yaml.in/yaml/v2"
)

type Project struct{}

func NewProject() contracts.Renderer {
	return &Project{}
}

func (r *Project) Render(
	ctx *contracts.Context,
	cat *resolved.Catalog,
) error {

	for _, project := range cat.Projects {

		manifest := manifests.NewProject(project)

		data, err := yaml.Marshal(manifest)

		if err != nil {
			return err
		}

		filename := filepath.Join(
			ctx.Root,
			"projects",
			project.Name+".yaml",
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
