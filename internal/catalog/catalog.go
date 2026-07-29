package catalog

import (
	"path/filepath"
)

type Catalog struct {
	Applications []Document[Application]
	Environments []Document[Environment]
	Clusters     []Document[Cluster]
}

func (l *FilesystemLoader) applicationsDir() string {
	thepath := filepath.Join(l.Root, "catalog", "applications")
	return thepath

}

func (l *FilesystemLoader) environmentsDir() string {
	return filepath.Join(l.Root, "catalog", "environments")
}

func (l *FilesystemLoader) clustersDir() string {
	return filepath.Join(l.Root, "catalog", "clusters")
}
