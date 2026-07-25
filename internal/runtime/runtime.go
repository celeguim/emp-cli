package runtime

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/catalog"
)

type Runtime struct {
	Root string

	Applications []catalog.Document[catalog.Application]
	Environments []catalog.Document[catalog.Environment]
	Clusters     []catalog.Document[catalog.Cluster]
}

func New(root string) *Runtime {
	return &Runtime{
		Root: root,
	}
}

func (r *Runtime) CreateWorkspace() error {
	dirs := []string{
		filepath.Join(r.Root, ".emp", "runtime"),
		filepath.Join(r.Root, ".emp", "runtime", "applications"),
		filepath.Join(r.Root, ".emp", "runtime", "environments"),
		filepath.Join(r.Root, ".emp", "runtime", "clusters"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return nil
}

func (r *Runtime) LoadCatalogs() error {
	loader := catalog.NewFilesystemLoader(r.Root)

	var err error

	r.Applications, err = loader.LoadApplications()
	if err != nil {
		return err
	}

	r.Environments, err = loader.LoadEnvironments()
	if err != nil {
		return err
	}

	r.Clusters, err = loader.LoadClusters()
	if err != nil {
		return err
	}

	fmt.Printf(
		"\nLoaded %d applications, %d environments, %d clusters\n",
		len(r.Applications),
		len(r.Environments),
		len(r.Clusters),
	)

	return nil
}

func (r *Runtime) WriteManifest() error {
	return nil
}

func (r *Runtime) RuntimeDir() string {
	return filepath.Join(r.Root, ".emp", "runtime")
}

func (r *Runtime) ApplicationsDir() string {
	return filepath.Join(r.RuntimeDir(), "applications")
}

func (r *Runtime) EnvironmentsDir() string {
	return filepath.Join(r.RuntimeDir(), "environments")
}

func (r *Runtime) ClustersDir() string {
	return filepath.Join(r.RuntimeDir(), "clusters")
}
