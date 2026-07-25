package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/runtime"
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

		rt := runtime.New(root)
		return rt.Render(cat)
	},
}

func init() {
	catalogCmd.AddCommand(catalogRenderCmd)
}
