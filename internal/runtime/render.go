package runtime

import "github.com/celeguim/emp-cli/internal/catalog"

func (r *Runtime) Render(cat *catalog.Catalog) error {

	if err := r.CreateWorkspace(); err != nil {
		return err
	}

	if err := r.Write(cat); err != nil {
		return err
	}

	return r.WriteManifest()
}
