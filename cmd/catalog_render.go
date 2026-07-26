package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/compiler"
	"github.com/celeguim/emp-cli/internal/resolver"
	"github.com/celeguim/emp-cli/internal/validator"
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

		report := validator.New().Validate(cat)
		if report.HasErrors() {
			return report
		}

		cat, err = resolver.New().Resolve(cat)
		if err != nil {
			return err
		}

		return compiler.NewCompiler(".").Compile(cat)

	},
}

func init() {
	catalogCmd.AddCommand(catalogRenderCmd)
}
