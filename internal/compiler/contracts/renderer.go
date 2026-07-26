package contracts

import (
	"github.com/celeguim/emp-cli/internal/catalog"
)

// func (r *Compiler) Render(cat *catalog.Catalog) error {

// 	if err := r.CreateWorkspace(); err != nil {
// 		return err
// 	}

// 	if err := r.Write(cat); err != nil {
// 		return err
// 	}

// 	return r.WriteManifest()
// }

// type Renderer interface {
// 	Render(*catalog.Catalog) error
// }

type Renderer interface {
	Render(ctx *Context, cat *catalog.Catalog) error
}

func (r *Compiler) WriteManifest() error {
	return nil
}
