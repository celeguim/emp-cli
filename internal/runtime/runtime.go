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

func (r *Runtime) LoadCatalogs() error {
	return nil
}

func (r *Runtime) Write() error {
	return nil
}

func (r *Runtime) WriteManifest() error {
	return nil
}
