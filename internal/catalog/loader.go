package catalog

import "fmt"

type Loaded[T any] struct {
	Path   string
	Object T
}

type Loader interface {
	Load() (*Catalog, error)
}

func (l *FilesystemLoader) Load() (*Catalog, error) {
	var c Catalog
	var err error

	c.Applications, err = l.loadApplications()
	if err != nil {
		return nil, err
	}

	fmt.Println("Apps:", len(c.Applications))
	fmt.Println("Envs:", len(c.Environments))
	fmt.Println("Clusters:", len(c.Clusters))

	return &c, nil
}
