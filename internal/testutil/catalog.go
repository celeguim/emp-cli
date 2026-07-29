package testutil

import "github.com/celeguim/emp-cli/internal/catalog"

func NewCatalog() *catalog.Catalog {
	return &catalog.Catalog{}
}

func DefaultCatalog() *catalog.Catalog {
	cat := NewCatalog()

	AddCluster(
		cat,
		"dev",
		"https://kubernetes.default.svc",
	)

	AddEnvironment(
		cat,
		"dev",
		"default",
		"dev",
		"default",
		"HEAD",
	)

	return cat
}

func AddCluster(
	cat *catalog.Catalog,
	name string,
	server string,
) {
	cat.Clusters = append(cat.Clusters, catalog.Document[catalog.Cluster]{
		Path: "test",
		Object: catalog.Cluster{
			Name:   name,
			Server: server,
		},
	})
}

func AddEnvironment(
	cat *catalog.Catalog,
	name string,
	project string,
	cluster string,
	namespace string,
	revision string,
) {
	cat.Environments = append(cat.Environments, catalog.Document[catalog.Environment]{
		Path: "test",
		Object: catalog.Environment{
			Name:           name,
			Project:        project,
			Cluster:        cluster,
			Namespace:      namespace,
			TargetRevision: revision,
		},
	})
}

func AddApplication(
	cat *catalog.Catalog,
	name string,
	appset string,
	environment string,
	repo string,
	path string,
) {
	cat.Applications = append(cat.Applications, catalog.Document[catalog.Application]{
		Path: "test",
		Object: catalog.Application{
			Name:           name,
			ApplicationSet: appset,
			Environment:    environment,
			RepoURL:        repo,
			Path:           path,
		},
	})
}
