package runtime

import (
	"github.com/celeguim/emp-cli/internal/catalog"
)

func (r *Runtime) Write(cat *catalog.Catalog) error {

	if err := r.RenderApplications(cat.Applications); err != nil {
		return err
	}

	if err := r.RenderEnvironments(cat.Environments); err != nil {
		return err
	}

	if err := r.RenderClusters(cat.Clusters); err != nil {
		return err
	}

	return nil
}
