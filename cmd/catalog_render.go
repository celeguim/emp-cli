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

		fmt.Printf("Applications: %d\n", len(resolvedCatalog.Applications))
		fmt.Printf("Projects: %d\n", len(resolvedCatalog.Projects))
		fmt.Printf("ApplicationSets: %d\n", len(resolvedCatalog.ApplicationSets))

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
