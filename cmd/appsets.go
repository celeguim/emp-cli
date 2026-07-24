package cmd

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/argocd"
	"github.com/celeguim/emp-cli/internal/runner"
	"github.com/spf13/cobra"
)

func init() {
	appsetsCmd.Flags().StringVarP(
		&output,
		"output",
		"o",
		"table",
		"Output format (table|json)",
	)
}

func init() {
	rootCmd.AddCommand(appsetsCmd)
}

func appsetsList(client *argocd.Client) (*runner.Result, error) {

	switch output {
	case "json":
		return client.AppListJSON()

	default:
		return client.AppList()
	}
}

var appsetsCmd = &cobra.Command{
	Use:   "appsets",
	Short: "List Argo CD ApplicationSets",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := argocd.New()
		result, err := appsetsList(client)

		if err != nil {
			return argocd.HandleError(result.Stderr, err)
		}

		fmt.Print(result.Stdout)
		return nil
	},
}
