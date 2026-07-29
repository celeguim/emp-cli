package compiler

import (
	"os"
	"path/filepath"
)

type Workspace struct {
	Root string
}

func (w *Workspace) Create() error {
	dirs := []string{
		w.Root,
		w.ApplicationsDir(),
		// w.ProjectsDir(),
		// w.ApplicationSetsDir(),
	}

	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	return nil
}

func (w *Workspace) CompilerDir() string {
	return filepath.Join(w.Root, ".emp", "compiler")
}

func (w *Workspace) ApplicationsDir() string {
	thepath := filepath.Join(w.Root, "generated", "applications")

	return thepath
}

func (l *Workspace) environmentsDir() string {
	return filepath.Join(l.Root, "generated", "environments")
}

func (l *Workspace) clustersDir() string {
	return filepath.Join(l.Root, "generated", "clusters")
}
