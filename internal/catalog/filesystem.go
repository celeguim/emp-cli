package catalog

type FilesystemLoader struct {
	Root string
}

func NewFilesystemLoader(root string) *FilesystemLoader {
	return &FilesystemLoader{
		Root: root,
	}
}

// func (l *FilesystemLoader) LoadApplications() ([]Document[Application], error) {
// 	return loadYAMLFiles[Application](l.applicationsDir())
// }

// func (l *FilesystemLoader) LoadEnvironments() ([]Document[Environment], error) {
// 	return loadYAMLFiles[Environment](l.environmentsDir())
// }

// func (l *FilesystemLoader) LoadClusters() ([]Document[Cluster], error) {
// 	return loadYAMLFiles[Cluster](l.clustersDir())
// }

func (l *FilesystemLoader) loadApplications() ([]Document[Application], error) {
	return loadYAMLFiles[Application](l.applicationsDir())
}

// func (l *FilesystemLoader) loadEnvironments() ([]Document[Environment], error) {
// 	return nil, nil
// }

// func (l *FilesystemLoader) loadClusters() ([]Document[Cluster], error) {
// 	return nil, nil
// }
