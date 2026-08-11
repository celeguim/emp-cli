package catalog

type Application struct {
	Name           string `yaml:"name"`
	ApplicationSet string `yaml:"applicationset,omitempty"`
	Environment    string `yaml:"environment"`
	RepoURL        string `yaml:"repoURL"`
	Path           string `yaml:"path"`
}
