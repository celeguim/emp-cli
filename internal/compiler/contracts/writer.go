package contracts

// type Writer interface {
// 	WriteApplication(app ApplicationManifest) error

// 	WriteCluster(cluster ClusterManifest) error

// 	WriteEnvironment(env EnvironmentManifest) error
// }

type Writer interface {
	Write(kind string, name string, data any) error
}
