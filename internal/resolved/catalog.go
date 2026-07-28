package resolved

import "github.com/celeguim/emp-cli/internal/catalog"

type Catalog struct {
	Applications []Application
	Projects     []Project
}

type Application struct {
	Application catalog.Application
	Environment catalog.Environment
	Cluster     catalog.Cluster
}
