package resolver

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/resolved"
)

type Resolver struct{}

func New() *Resolver {
	return &Resolver{}
}

func Resolve(cat *catalog.Catalog) (*resolved.Catalog, error) {
	environments := make(map[string]catalog.Environment, len(cat.Environments))
	for _, env := range cat.Environments {
		fmt.Println(env)
		environments[env.Object.Project] = env.Object
	}

	clusters := make(map[string]catalog.Cluster, len(cat.Clusters))
	for _, cluster := range cat.Clusters {
		fmt.Println(cluster)
		clusters[cluster.Object.Environment] = cluster.Object
	}

	rc := &resolved.Catalog{Applications: make([]resolved.ResolvedApplication, 0, len(cat.Applications))}

	for _, app := range cat.Applications {
		env, ok := environments[app.Object.Project]
		if !ok {
			return nil, fmt.Errorf("application %q references unknown project %q", app.Object.Name, app.Object.Project)
		}

		cluster, ok := clusters[env.Name]
		if !ok {
			return nil, fmt.Errorf("environment %q has no cluster", env.Name)
		}

		rc.Applications = append(rc.Applications, resolved.ResolvedApplication{
			Application: app.Object,
			Environment: env,
			Cluster:     cluster,
		})
	}

	return rc, nil
}
