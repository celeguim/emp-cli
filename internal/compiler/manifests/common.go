package manifests

type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type Destination struct {
	Server    string `yaml:"server,omitempty"`
	Name      string `yaml:"name,omitempty"`
	Namespace string `yaml:"namespace"`
}

type Source struct {
	RepoURL        string `yaml:"repoURL"`
	Path           string `yaml:"path"`
	TargetRevision string `yaml:"targetRevision"`
}

type Resource struct {
	Group string `yaml:"group"`
	Kind  string `yaml:"kind"`
}
