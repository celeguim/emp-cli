package cmd

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/argocd"
	"github.com/celeguim/emp-cli/internal/runner"
	"github.com/spf13/cobra"
)

var output string

func init() {
	appsCmd.Flags().StringVarP(
		&output,
		"output",
		"o",
		"table",
		"Output format (table|json)",
	)
}

func init() {
	rootCmd.AddCommand(appsCmd)
}

func appList(client *argocd.Client) (*runner.Result, error) {

	switch output {
	case "json":
		return client.AppListJSON()

	default:
		return client.AppList()
	}
}

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "List Argo CD Applications",

	RunE: func(cmd *cobra.Command, args []string) error {

		client := argocd.New()
		result, err := appList(client)

		if err != nil {
			return argocd.HandleError(result.Stderr, err)
		}

		fmt.Print(result.Stdout)

		return nil
	},
}
