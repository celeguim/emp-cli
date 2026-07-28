package manifests

import "github.com/celeguim/emp-cli/internal/resolved"

type Project struct {
	APIVersion string      `yaml:"apiVersion"`
	Kind       string      `yaml:"kind"`
	Metadata   Metadata    `yaml:"metadata"`
	Spec       ProjectSpec `yaml:"spec"`
}

type ProjectSpec struct {
	SourceRepos                []string      `yaml:"sourceRepos"`
	Destinations               []Destination `yaml:"destinations"`
	ClusterResourceWhitelist   []Resource    `yaml:"clusterResourceWhitelist"`
	NamespaceResourceWhitelist []Resource    `yaml:"namespaceResourceWhitelist"`
}

func NewProject(project resolved.Project) Project {

	manifest := Project{
		APIVersion: "argoproj.io/v1alpha1",
		Kind:       "AppProject",
		Metadata: Metadata{
			Name:      project.Name,
			Namespace: "argocd",
		},
		Spec: ProjectSpec{
			SourceRepos: project.SourceRepos,

			ClusterResourceWhitelist: []Resource{
				{
					Group: "*",
					Kind:  "*",
				},
			},

			NamespaceResourceWhitelist: []Resource{
				{
					Group: "*",
					Kind:  "*",
				},
			},
		},
	}

	for _, d := range project.Destinations {
		manifest.Spec.Destinations = append(
			manifest.Spec.Destinations,
			Destination{
				Server:    d.Server,
				Namespace: d.Namespace,
			},
		)
	}

	return manifest
}
