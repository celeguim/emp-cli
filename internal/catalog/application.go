package catalog

type Application struct {
	Name   string `yaml:"name"`
	Chart  string `yaml:"chart"`
	Source Source `yaml:"source"`
}

type Source struct {
	RepoURL string `yaml:"repoURL"`
	Path    string `yaml:"path"`
}
