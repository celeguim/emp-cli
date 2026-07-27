package argocd

type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type ApplicationSpec struct {
	Project     string      `yaml:"project"`
	Source      Source      `yaml:"source"`
	Destination Destination `yaml:"destination"`
	// SyncPolicy  *SyncPolicy `yaml:"syncPolicy,omitempty"`
}

type Source struct {
	RepoURL        string `yaml:"repoURL"`
	Path           string `yaml:"path"`
	TargetRevision string `yaml:"targetRevision"`
}

type Destination struct {
	Server    string `yaml:"server"`
	Namespace string `yaml:"namespace"`
}
