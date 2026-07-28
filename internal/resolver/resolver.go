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

	// env map
	envIndex := make(map[string]catalog.Environment, len(cat.Environments))
	for _, env := range cat.Environments {
		envIndex[env.Object.Name] = env.Object
	}

	// cluster map
	clusterIndex := make(map[string]catalog.Cluster, len(cat.Clusters))
	for _, cluster := range cat.Clusters {
		clusterIndex[cluster.Object.Name] = cluster.Object
	}

	rc := &resolved.Catalog{
		Applications: make([]resolved.Application, 0, len(cat.Applications)),
	}

	for _, app := range cat.Applications {

		env, ok := envIndex[app.Object.Environment]
		if !ok {
			return nil, fmt.Errorf(
				"application %q references unknown environment %q",
				app.Object.Name,
				app.Object.Environment,
			)
		}

		cluster, ok := clusterIndex[env.Cluster]
		if !ok {
			return nil, fmt.Errorf(
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

	// project map
	projects := map[string]*resolved.Project{}
	// for _, p := range projects {
	for _, app := range rc.Applications {
		name := app.Environment.Project
		project, ok := projects[name]
		if !ok {
			project = &resolved.Project{
				Name: name,
			}
			projects[name] = project
		}
		rc.Projects = append(rc.Projects, *project)
	}

	return rc, nil
}
