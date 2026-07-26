package compiler

import (
	"os"
	"path/filepath"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/compiler/contracts"
	"github.com/celeguim/emp-cli/internal/compiler/renderers"
)

// Compiler transforms a validated catalog into GitOps artifacts.
type Compiler struct {
	context   *contracts.Context
	renderers []contracts.Renderer
}

func defaultRenderers() []contracts.Renderer {
	return []contracts.Renderer{
		renderers.NewApplication(),
		renderers.NewEnvironment(),
		renderers.NewCluster(),
	}
}

//	func NewCompiler(root string) *Compiler {
//		return &Compiler{
//			context:   &contracts.Context{},
//			renderers: defaultRenderers(),
//		}
//	}
func NewCompiler(root string) *Compiler {
	return &Compiler{
		context: &contracts.Context{
			Root: filepath.Join(root, ".emp"),
		},
		renderers: defaultRenderers(),
	}
}

// func (c *Compiler) Compile(cat *catalog.Catalog) error {
// 	fmt.Println("Compile()")

// 	for _, r := range c.renderers {
// 		fmt.Printf("%T\n", r)

// 		if err := r.Render(c.context, cat); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (c *Compiler) Compile(cat *catalog.Catalog) error {
// 	if err := c.workspace.Create(); err != nil {
// 		return err
// 	}

// 	for _, r := range c.renderers {
// 		if err := r.Render(c.context, cat); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

func (c *Compiler) Compile(cat *catalog.Catalog) error {
	if err := os.MkdirAll(c.context.Root, 0755); err != nil {
		return err
	}

	for _, r := range c.renderers {
		if err := r.Render(c.context, cat); err != nil {
			return err
		}
	}

	return nil
}
