package catalog

type Environment struct {
	Name           string `yaml:"name"`
	Project        string `yaml:"project"`
	Cluster        string `yaml:"cluster"`
	RepoURL        string `yaml:"repo"`
	Namespace      string `yaml:"namespace"`
	TargetRevision string `yaml:"targetRevision"`
	SyncPolicy     string `yaml:"syncPolicy"`
}
