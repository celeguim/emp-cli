package catalog

type Environment struct {
	EnvName string `yaml:"name"`

	Project string `yaml:"project"`

	TargetRevision string `yaml:"targetRevision"`

	Namespace string `yaml:"namespace"`
}
