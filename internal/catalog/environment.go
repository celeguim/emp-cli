package catalog

type Environment struct {
	Name string `yaml:"name"`

	Project string `yaml:"project"`

	TargetRevision string `yaml:"targetRevision"`

	Namespace string `yaml:"namespace"`
}
