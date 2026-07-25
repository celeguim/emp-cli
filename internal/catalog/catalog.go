package catalog

type Catalog struct {
	Applications []Document[Application]
	Environments []Document[Environment]
	Clusters     []Document[Cluster]
}
