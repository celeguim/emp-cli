package resolved

import "github.com/celeguim/emp-cli/internal/catalog"

type Catalog struct {
	Applications []Application
}

type Application struct {
	Application catalog.Application
	Environment catalog.Environment
	Cluster     catalog.Cluster
}
