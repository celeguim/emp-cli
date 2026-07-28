package manifests

import "github.com/celeguim/emp-cli/internal/resolved"

type Application struct {
	APIVersion string          `yaml:"apiVersion"`
	Kind       string          `yaml:"kind"`
	Metadata   Metadata        `yaml:"metadata"`
	Spec       ApplicationSpec `yaml:"spec"`
}

type ApplicationSpec struct {
	Project     string      `yaml:"project"`
	Source      Source      `yaml:"source"`
	Destination Destination `yaml:"destination"`
}

func NewApplication(app resolved.Application) Application {

	return Application{
		APIVersion: "argoproj.io/v1alpha1",
		Kind:       "Application",
		Metadata: Metadata{
			Name:      app.Application.Name,
			Namespace: "argocd",
		},
		Spec: ApplicationSpec{
			Project: app.Environment.Project,
			Source: Source{
				RepoURL:        app.Application.RepoURL,
				Path:           app.Application.Path,
				TargetRevision: app.Environment.TargetRevision,
			},
			Destination: Destination{
				Server:    app.Cluster.Server,
				Namespace: app.Environment.Namespace,
			},
		},
	}
}
