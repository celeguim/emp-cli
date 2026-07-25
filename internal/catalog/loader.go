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

func (l *FilesystemLoader) Load() (*Catalog, error) {
	var c Catalog
	var err error

	c.Applications, err = l.LoadApplications()
	if err != nil {
		return nil, err
	}

	c.Environments, err = l.LoadEnvironments()
	if err != nil {
		return nil, err
	}

	c.Clusters, err = l.LoadClusters()
	if err != nil {
		return nil, err
	}

	return &c, nil
}
