package compiler

import (
	"os"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/compiler/contracts"
	"github.com/celeguim/emp-cli/internal/compiler/renderers"
)

type Compiler struct {
	ctx       *contracts.Context
	renderers []contracts.Renderer
}

func New(root string) *Compiler {

	ctx := &contracts.Context{
		Root: root,
	}

	return &Compiler{
		ctx: ctx,
		renderers: []contracts.Renderer{
			renderers.NewApplicationRenderer(),
			renderers.NewEnvironmentRenderer(),
			renderers.NewClusterRenderer(),
		},
	}
}

func (c *Compiler) Compile(cat *catalog.Catalog) error {

	for _, r := range c.renderers {

		if err := r.Render(c.ctx, cat); err != nil {
			return err
		}
	}

	return nil
}

func (r *Compiler) CreateWorkspace() error {
	dirs := []string{
		filepath.Join(r.Root, ".emp", "Compiler"),
		filepath.Join(r.Root, ".emp", "Compiler", "applications"),
		filepath.Join(r.Root, ".emp", "Compiler", "environments"),
		filepath.Join(r.Root, ".emp", "Compiler", "clusters"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return nil
}

func (r *Compiler) CompilerDir() string {
	return filepath.Join(r.Root, ".emp", "Compiler")
}

func (r *Compiler) ApplicationsDir() string {
	return filepath.Join(r.CompilerDir(), "applications")
}

func (r *Compiler) EnvironmentsDir() string {
	return filepath.Join(r.CompilerDir(), "environments")
}

func (r *Compiler) ClustersDir() string {
	return filepath.Join(r.CompilerDir(), "clusters")
}
