package cmd

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/kubectl"
	"github.com/spf13/cobra"
)

var appsetsCmd = &cobra.Command{
	Use:   "appsets",
	Short: "List Argo CD ApplicationSets",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := kubectl.New()

		result, err := client.Get(
			"appset",
			"-A",
		)
		if err != nil {
			return fmt.Errorf("%s", result.Stderr)
		}

		fmt.Print(result.Stdout)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(appsetsCmd)
}
