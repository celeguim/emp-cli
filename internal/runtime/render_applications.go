package runtime

import (
	"fmt"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func (r *Runtime) RenderApplications(
	apps []catalog.Document[catalog.Application],
) error {

	for _, doc := range apps {

		name := doc.Object.AppName

		fmt.Printf("rendering application %s\n", name)

		if name == "" {
			return fmt.Errorf("application %s has no metadata.name", doc.Path)
		}

		file := filepath.Join(
			r.ApplicationsDir(),
			name+".yaml",
		)

		if err := writeYAML(file, doc.Object); err != nil {
			return err
		}
	}

	return nil
}
