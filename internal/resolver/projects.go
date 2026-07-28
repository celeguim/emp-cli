package resolver

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/resolved"
)

func resolveProjects(rc *resolved.Catalog) error {

	projects := map[string]*resolved.Project{}

	for _, app := range rc.Applications {

		projectName := app.Environment.Project

		fmt.Printf("Resolving project: %s\n", projectName)

		project, ok := projects[projectName]
		if !ok {

			project = &resolved.Project{
				Name: projectName,
			}

			projects[projectName] = project
		}

		// enriquecer o Project
		if !contains(project.SourceRepos, app.Application.RepoURL) {
			project.SourceRepos = append(project.SourceRepos, app.Application.RepoURL)
		}

		if !hasDestination(project,
			app.Cluster.Server,
			app.Environment.Namespace) {

			project.Destinations = append(
				project.Destinations,
				resolved.Destination{
					Server:    app.Cluster.Server,
					Namespace: app.Environment.Namespace,
				},
			)
		}

	}

	for _, project := range projects {
		rc.Projects = append(rc.Projects, *project)
	}

	fmt.Println("Projects resolved:", len(rc.Projects))
	for _, p := range rc.Projects {
		fmt.Printf("%+v\n", p)
	}

	return nil
}

func contains(ss []string, value string) bool {
	for _, s := range ss {
		if s == value {
			return true
		}
	}
	return false
}

func hasDestination(
	project *resolved.Project,
	server string,
	namespace string,
) bool {

	for _, d := range project.Destinations {

		if d.Server == server &&
			d.Namespace == namespace {

			return true
		}
	}

	return false
}
