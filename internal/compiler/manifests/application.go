package manifests

import (
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

	// fmt.Printf("SyncPolicy: %+v\n", manifest.Spec.SyncPolicy)
	// manifest.Spec.SyncPolicy = buildSyncPolicy(app)

	return manifest
}

// func buildSyncPolicy(env Environment) *SyncPolicy {
// 	if env.SyncPolicy != "enabled" {
// 		return nil
// 	}

// 	switch env.Name {
// 	case "dev":
// 		return &SyncPolicy{
// 			Automated: &Automated{
// 				Enabled:  true,
// 				Prune:    true,
// 				SelfHeal: true,
// 			},
// 		}

// 	case "uat":
// 		return &SyncPolicy{
// 			Automated: &Automated{
// 				Enabled: true,
// 				Prune:   true,
// 			},
// 		}

// 	case "prd":
// 		return &SyncPolicy{
// 			Automated: &Automated{
// 				Enabled: false,
// 			},
// 		}
// 	}

// 	return nil
// }
