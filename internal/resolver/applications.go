package resolver

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/resolved"
)

func resolveApplications(cat *catalog.Catalog, rc *resolved.Catalog) error {

	envIndex := make(map[string]catalog.Environment, len(cat.Environments))
	for _, env := range cat.Environments {
		envIndex[env.Object.Name] = env.Object
	}

	clusterIndex := make(map[string]catalog.Cluster, len(cat.Clusters))
	for _, cluster := range cat.Clusters {
		clusterIndex[cluster.Object.Name] = cluster.Object
	}

	rc.Applications = make([]resolved.Application, 0, len(cat.Applications))

	for _, app := range cat.Applications {

		env, ok := envIndex[app.Object.Environment]
		if !ok {
			return fmt.Errorf(
				"application %q references unknown environment %q",
				app.Object.Name,
				app.Object.Environment,
			)
		}

		cluster, ok := clusterIndex[env.Cluster]
		if !ok {
			return fmt.Errorf(
				"environment %q references unknown cluster %q",
				env.Name,
				env.Cluster,
			)
		}

		rc.Applications = append(rc.Applications, resolved.Application{
			Application: app.Object,
			Environment: env,
			Cluster:     cluster,
		})
	}

	return nil
}
