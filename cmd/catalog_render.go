package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/compiler"
	"github.com/celeguim/emp-cli/internal/resolver"
)

var catalogRenderCmd = &cobra.Command{

	Use:   "render",
	Short: "Render runtime artifacts",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("catalog render")

		root, err := filepath.Abs(".")
		if err != nil {
			return err
		}

		loader := catalog.NewFilesystemLoader(root)
		cat, err := loader.Load()

		if err != nil {
			return err
		}

		resolvedCatalog, err := resolver.Resolve(cat)
		if err != nil {
			return err
		}

		c := compiler.NewCompiler(root)

		return c.Compile(resolvedCatalog)

	},
}

func init() {
	catalogCmd.AddCommand(catalogRenderCmd)
}
