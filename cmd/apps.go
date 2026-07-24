package cmd

import (
	"fmt"
	"strings"

	"github.com/celeguim/emp-cli/internal/argocd"
	"github.com/spf13/cobra"
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "List Argo CD applications",

	RunE: func(cmd *cobra.Command, args []string) error {

		client := argocd.New()

		result, err := client.AppList()
		if err != nil {
			if strings.Contains(result.Stderr, "Unauthenticated") {
				return fmt.Errorf("you are not logged into Argo CD. Run 'argocd login'")
			}
			return fmt.Errorf("%s", result.Stderr)
		}

		fmt.Print(result.Stdout)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(appsCmd)
}
