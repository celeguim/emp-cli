package runtime

import (
	"fmt"
	"path/filepath"
)

func (r *Runtime) RenderClusters() error {

	for _, doc := range r.Clusters {

		name := doc.Object.ClusterName
		if name == "" {
			return fmt.Errorf("cluster %s has no metadata.name", doc.Path)
		}

		file := filepath.Join(
			r.ClustersDir(),
			name+".yaml",
		)

		if err := writeYAML(file, doc.Object); err != nil {
			return err
		}
	}

	return nil
}
