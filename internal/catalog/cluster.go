package catalog

type Cluster struct {
	ClusterName string `yaml:"clusterName"`

	Server string `yaml:"server"`

	Environment string `yaml:"environment"`
}
