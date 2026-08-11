package manifests

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/resolved"
)

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
	SyncPolicy  *SyncPolicy `yaml:"syncPolicy,omitempty"`
}

type SyncPolicy struct {
	Automated *Automated `yaml:"automated,omitempty"`
}

type Automated struct {
	Prune    bool `yaml:"prune,omitempty"`
	SelfHeal bool `yaml:"selfHeal,omitempty"`
}

func NewApplication(app resolved.Application) Application {

	manifest := Application{
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

	manifest.Spec.SyncPolicy = buildSyncPolicy(app.Environment)
	fmt.Printf("SyncPolicy: %+v\n", manifest.Spec.SyncPolicy)

	return manifest
}

func buildSyncPolicy(env catalog.Environment) *SyncPolicy {
	if env.SyncPolicy != "enabled" {
		return nil
	}

	return &SyncPolicy{
		Automated: &Automated{
			Prune:    true,
			SelfHeal: true,
		},
	}
}
