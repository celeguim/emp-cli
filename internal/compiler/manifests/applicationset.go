package manifests

import "github.com/celeguim/emp-cli/internal/resolved"

func NewApplicationSet(appset resolved.ApplicationSet) ApplicationSet {

	manifest := ApplicationSet{
		APIVersion: "argoproj.io/v1alpha1",
		Kind:       "ApplicationSet",
		Metadata: Metadata{
			Name:      appset.Name,
			Namespace: "argocd",
		},
		Spec: ApplicationSetSpec{
			Generators: []Generator{
				{
					List: ListGenerator{},
				},
			},
			Template: ApplicationTemplate{
				Metadata: TemplateMetadata{
					Name: "{{name}}",
				},
				Spec: TemplateSpec{
					Project: appset.Signature.Project,
					Source: Source{
						RepoURL:        appset.Signature.RepoURL,
						TargetRevision: appset.Signature.TargetRevision,
						Path:           "{{path}}",
					},
					Destination: Destination{
						Server:    appset.Signature.Server,
						Namespace: appset.Signature.Namespace,
					},
				},
			},
		},
	}

	for _, app := range appset.Applications {

		manifest.Spec.Generators[0].List.Elements = append(
			manifest.Spec.Generators[0].List.Elements,
			Element{
				Name: app.Application.Name,
				Path: app.Application.Path,
			},
		)
	}

	return manifest
}

type ApplicationSet struct {
	APIVersion string             `yaml:"apiVersion"`
	Kind       string             `yaml:"kind"`
	Metadata   Metadata           `yaml:"metadata"`
	Spec       ApplicationSetSpec `yaml:"spec"`
}

type ApplicationSetSpec struct {
	Generators []Generator         `yaml:"generators"`
	Template   ApplicationTemplate `yaml:"template"`
}

type Generator struct {
	List ListGenerator `yaml:"list"`
}

type ListGenerator struct {
	Elements []Element `yaml:"elements"`
}

type Element struct {
	Name      string `yaml:"name"`
	Path      string `yaml:"path"`
	Namespace string `yaml:"namespace,omitempty"`
	Cluster   string `yaml:"cluster,omitempty"`
}

type ApplicationTemplate struct {
	Metadata TemplateMetadata `yaml:"metadata"`
	Spec     TemplateSpec     `yaml:"spec"`
}

type TemplateSpec struct {
	Project     string      `yaml:"project"`
	Source      Source      `yaml:"source"`
	Destination Destination `yaml:"destination"`
}

type TemplateMetadata struct {
	Name string `yaml:"name"`
}
