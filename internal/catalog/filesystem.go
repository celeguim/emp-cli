package catalog

import (
	"fmt"
	"path/filepath"
)

type FilesystemLoader struct {
	Root string
}

func NewFilesystemLoader(root string) *FilesystemLoader {
	return &FilesystemLoader{
		Root: root,
	}
}

func (l *FilesystemLoader) applicationsDir() string {
	thepath := filepath.Join(l.Root, "catalog", "applications")

	fmt.Printf(
		"l.root: %s , thepath: %s\n",
		l.Root, thepath)

	return thepath
}

func (l *FilesystemLoader) environmentsDir() string {
	return filepath.Join(l.Root, "catalog", "environments")
}

func (l *FilesystemLoader) clustersDir() string {
	return filepath.Join(l.Root, "catalog", "clusters")
}

func (l *FilesystemLoader) LoadApplications() ([]Document[Application], error) {
	return loadYAMLFiles[Application](l.applicationsDir())
}

func (l *FilesystemLoader) LoadEnvironments() ([]Document[Environment], error) {
	return loadYAMLFiles[Environment](l.environmentsDir())
}

func (l *FilesystemLoader) LoadClusters() ([]Document[Cluster], error) {
	return loadYAMLFiles[Cluster](l.clustersDir())
}
