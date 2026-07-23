package cmd

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/exec"
	"github.com/spf13/cobra"
)

var appsetsCmd = &cobra.Command{
	Use:   "appsets",
	Short: "List Argo CD ApplicationSets",

	RunE: func(cmd *cobra.Command, args []string) error {

		out, err := exec.Run(
			"kubectl",
			"get",
			"appset",
			"-A",
		)

		if err != nil {
			return err
		}

		fmt.Print(out)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(appsetsCmd)
}