package runtime

type Application struct {
	AppName string `yaml:"appName"`
	Chart   string `yaml:"chart"`
}

type Environment struct {
	Name string `yaml:"name"`

	Project string `yaml:"project"`

	TargetRevision string `yaml:"targetRevision"`

	Namespace string `yaml:"namespace"`
}

type Cluster struct {
	ClusterName string `yaml:"clusterName"`

	Server string `yaml:"server"`

	Environment string `yaml:"environment"`
}

type RenderedCluster struct {
	ClusterName string

	Server string

	Project string

	TargetRevision string

	Namespace string
}
