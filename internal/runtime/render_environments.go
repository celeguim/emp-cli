package runtime

import (
	"fmt"
	"path/filepath"
)

func (r *Runtime) RenderEnvironments() error {

	for _, doc := range r.Environments {

		name := doc.Object.EnvName
		if name == "" {
			return fmt.Errorf("environment %s has no metadata.name", doc.Path)
		}

		file := filepath.Join(
			r.EnvironmentsDir(),
			name+".yaml",
		)

		if err := writeYAML(file, doc.Object); err != nil {
			return err
		}
	}

	return nil
}
