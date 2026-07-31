package catalog

type SyncPolicy struct {
	Automated *Automated `yaml:"automated,omitempty"`
}

type Automated struct {
	Enabled  bool `yaml:"enabled,omitempty"`
	Prune    bool `yaml:"prune,omitempty"`
	SelfHeal bool `yaml:"selfHeal,omitempty"`
}
