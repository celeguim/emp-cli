package runtime

import (
	"os"
	"path/filepath"
)

type Runtime struct {
	Root string
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
