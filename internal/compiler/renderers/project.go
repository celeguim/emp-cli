package renderers

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/compiler/contracts"
	"github.com/celeguim/emp-cli/internal/compiler/manifests"
	"github.com/celeguim/emp-cli/internal/resolved"
)

type Project struct{}

func NewProject() contracts.Renderer {
	return &Project{}
}

func (r *Project) Render(ctx *contracts.Context, cat *resolved.Catalog) error {
	fmt.Println("render projects")

	for _, project := range cat.Projects {
		manifest := manifests.NewProject(project)
		if err := ctx.Writer.Write(projectKind, project.Name, manifest); err != nil {
			return err
		}
	}

	return nil
}
