package resolver

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/resolved"
)

func signatureOf(app resolved.Application) resolved.ApplicationSetSignature {

	return resolved.ApplicationSetSignature{
		Project:        app.Environment.Project,
		RepoURL:        app.Application.RepoURL,
		TargetRevision: app.Environment.TargetRevision,
		Server:         app.Cluster.Server,
		Namespace:      app.Environment.Namespace,
	}
}

func resolveApplicationSets(rc *resolved.Catalog) error {

	groups := map[string]*resolved.ApplicationSet{}

	for _, app := range rc.Applications {

		if app.Application.ApplicationSet == "" {
			continue
		}

		name := app.Application.ApplicationSet

		group, ok := groups[name]
		if !ok {

			group = &resolved.ApplicationSet{
				Name: name,
			}

			group.Signature = signatureOf(app)

			groups[name] = group
		}

		if group.Signature != signatureOf(app) {

			return fmt.Errorf(
				"applicationset %q: application %q has incompatible template configuration",
				name,
				app.Application.Name,
			)
		}

		group.Applications = append(group.Applications, app)
	}

	for _, g := range groups {
		rc.ApplicationSets = append(rc.ApplicationSets, *g)
	}

	return nil
}

func validateSignature(
	expected resolved.ApplicationSetSignature,
	app resolved.Application,
) error {

	actual := signatureOf(app)

	if expected.Project != actual.Project {
		return fmt.Errorf(
			"application %q has project %q, expected %q",
			app.Application.Name,
			actual.Project,
			expected.Project,
		)
	}

	if expected.RepoURL != actual.RepoURL {
		return fmt.Errorf(
			"application %q has repoURL %q, expected %q",
			app.Application.Name,
			actual.RepoURL,
			expected.RepoURL,
		)
	}

	if expected.TargetRevision != actual.TargetRevision {
		return fmt.Errorf(
			"application %q has targetRevision %q, expected %q",
			app.Application.Name,
			actual.TargetRevision,
			expected.TargetRevision,
		)
	}

	if expected.Server != actual.Server {
		return fmt.Errorf(
			"application %q has server %q, expected %q",
			app.Application.Name,
			actual.Server,
			expected.Server,
		)
	}

	if expected.Namespace != actual.Namespace {
		return fmt.Errorf(
			"application %q has namespace %q, expected %q",
			app.Application.Name,
			actual.Namespace,
			expected.Namespace,
		)
	}

	return nil
}
