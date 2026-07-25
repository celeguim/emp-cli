package runtime

import (
	"fmt"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func (r *Runtime) RenderEnvironments(
	envs []catalog.Document[catalog.Environment],
) error {

	for _, doc := range envs {

		name := doc.Object.EnvName
		if name == "" {
			return fmt.Errorf("environment %s has no metadata.name", doc.Path)
		}

		file := filepath.Join(r.EnvironmentsDir(), name+".yaml")

		if err := writeYAML(file, doc.Object); err != nil {
			return err
		}
	}

	return nil
}
