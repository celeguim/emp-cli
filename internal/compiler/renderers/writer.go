package renderers

import (
	"fmt"
	"os"
	"path/filepath"

	"go.yaml.in/yaml/v2"
)

const (
	applicationKind    = "application"
	projectKind        = "project"
	applicationSetKind = "applicationset"
)

type Writer struct {
	root string
}

func NewWriter(root string) *Writer {
	return &Writer{root: root}
}

func (w *Writer) Write(kind string, name string, manifest any) error {
	data, err := yaml.Marshal(manifest)
	if err != nil {
		return err
	}

	switch kind {
	case applicationKind:
		return w.writeApp(name, data)
	case projectKind:
		return w.writeProject(name, data)
	case applicationSetKind:
		return w.writeApplicationSet(name, data)
	default:
		return fmt.Errorf("unsupported manifest kind %q", kind)
	}
}

func (w *Writer) writeApp(name string, data []byte) error {
	return w.write("applications", name, data)
}

func (w *Writer) writeProject(name string, data []byte) error {
	return w.write("projects", name, data)
}

func (w *Writer) writeApplicationSet(name string, data []byte) error {
	return w.write("applicationsets", name, data)
}

func (w *Writer) write(directory string, name string, data []byte) error {
	path := filepath.Join(w.root, directory, name+".yaml")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}
