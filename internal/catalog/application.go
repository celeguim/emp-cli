package catalog

type Application struct {
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	RepoURL     string `yaml:"repoURL"`
	Path        string `yaml:"path"`
}
