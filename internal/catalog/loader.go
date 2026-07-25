package catalog

type Loaded[T any] struct {
	Path   string
	Object T
}

type Loader interface {
	LoadApplications() ([]Document[Application], error)
	LoadEnvironments() ([]Document[Environment], error)
	LoadClusters() ([]Document[Cluster], error)
}
