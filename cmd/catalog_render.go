package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/celeguim/emp-cli/internal/runtime"
)

var catalogRenderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render runtime artifacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		rt := runtime.New(".")
		return rt.Render()
	},
}

func init() {
	abs_path, _ := filepath.Abs(".")
	fmt.Printf("catalog_render: abs_path: %s\n", abs_path)
	catalogCmd.AddCommand(catalogRenderCmd)
}
