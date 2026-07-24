package cmd

import (
	"github.com/spf13/cobra"

	"github.com/seuusuario/emp/internal/runtime"
)

var catalogRenderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render runtime artifacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runtime.Render(".")
	},
}

func init() {
	catalogCmd.AddCommand(catalogRenderCmd)
}
