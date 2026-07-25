package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/runtime"
	"github.com/celeguim/emp-cli/internal/validator"
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

		report := validator.New().Validate(cat)
		if report.HasErrors() {
			return report
		}

		return runtime.New(root).Render(cat)

	},
}

func init() {
	catalogCmd.AddCommand(catalogRenderCmd)
}
