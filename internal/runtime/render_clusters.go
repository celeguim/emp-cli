package runtime

import (
	"fmt"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func (r *Runtime) RenderClusters(
	clusters []catalog.Document[catalog.Cluster],
) error {

	for _, doc := range clusters {

		name := doc.Object.ClusterName

		if name == "" {
			return fmt.Errorf("cluster %s has no metadata.name", doc.Path)
		}

		file := filepath.Join(r.ClustersDir(), name+".yaml")

		if err := writeYAML(file, doc.Object); err != nil {
			return err
		}
	}

	return nil
}
