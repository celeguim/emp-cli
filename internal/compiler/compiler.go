package compiler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/compiler/contracts"
	"github.com/celeguim/emp-cli/internal/compiler/renderers"
	"github.com/celeguim/emp-cli/internal/resolved"
)

// Compiler transforms a validated catalog into GitOps artifacts.
type Compiler struct {
	context   *contracts.Context
	renderers []contracts.Renderer
}

func defaultRenderers() []contracts.Renderer {
	return []contracts.Renderer{
		renderers.NewApplication(),
		renderers.NewProject(),
		renderers.NewApplicationSet(),
		// renderers.NewEnvironment(),
		// renderers.NewCluster(),
	}
}

func NewCompiler(root string) *Compiler {
	return &Compiler{
		context: &contracts.Context{
			Root:   filepath.Join(root, ".emp"),
			Writer: renderers.NewWriter(filepath.Join(root, ".emp")),
		},
		renderers: defaultRenderers(),
	}
}

func (c *Compiler) Compile(cat *resolved.Catalog) error {
	if err := os.MkdirAll(c.context.Root, 0755); err != nil {
		return err
	}

	for _, r := range c.renderers {
		fmt.Printf("running %T\n", r)

		if err := r.Render(c.context, cat); err != nil {
			return err
		}
	}

	fmt.Println("renderers:", len(c.renderers))
	return nil
}
