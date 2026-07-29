package catalog

type Loaded[T any] struct {
	Path   string
	Object T
}

type Loader interface {
	Load() (*Catalog, error)
	// LoadTest(path string) (*Catalog, error)
}

func (l *FilesystemLoader) Load() (*Catalog, error) {
	var c Catalog
	var err error

	c.Environments, err = l.loadEnvironments()
	if err != nil {
		return nil, err
	}

	c.Clusters, err = l.loadClusters()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// func (l *FilesystemLoader) LoadTest(path string) (*Catalog, error) {
// 	var c Catalog
// 	var err error

// 	c.Applications, err = l.loadApplicationsTest(path)
// 	if err != nil {
// 		return nil, err
// 	}
// }
