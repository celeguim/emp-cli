package renderers

import (
	"os"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/argocd"
	"github.com/celeguim/emp-cli/internal/compiler/contracts"
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

		argoApp := argocd.Application{
			APIVersion: "argoproj.io/v1alpha1",
			Kind:       "Application",
			Metadata: argocd.Metadata{
				Name:      app.Application.Name,
				Namespace: "argocd",
			},
			Spec: argocd.ApplicationSpec{
				Project: app.Environment.Project,
				Source: argocd.Source{
					RepoURL:        app.Application.RepoURL,
					Path:           app.Application.Path,
					TargetRevision: app.Environment.TargetRevision,
				},
				Destination: argocd.Destination{
					Server:    app.Cluster.Server,
					Namespace: app.Environment.Namespace,
				},
			},
		}

		data, err := yaml.Marshal(argoApp)
		// data, err := yaml.Marshal(app.Object)

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
