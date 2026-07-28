package catalog

type Environment struct {
	Name           string `yaml:"name"`
	Project        string `yaml:"project"`
	Cluster        string `yaml:"cluster"`
	Namespace      string `yaml:"namespace"`
	TargetRevision string `yaml:"targetRevision"`
}
