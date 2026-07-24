package cmd

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/argocd"
	"github.com/celeguim/emp-cli/internal/runner"
	"github.com/spf13/cobra"
)

var clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "List Argo CD clusters",

	RunE: func(cmd *cobra.Command, args []string) error {

		client := argocd.New()
		result, err := clusterList(client)

		if err != nil {

			if err != nil {
				return argocd.HandleError(result.Stderr, err)
			}

			return err
		}

		fmt.Print(result.Stdout)

		return nil
	},
}

func clusterList(client *argocd.Client) (*runner.Result, error) {

	switch output {
	case "json":
		return client.ClusterListJSON()

	default:
		return client.ClusterList()
	}
}

func init() {
	rootCmd.AddCommand(clustersCmd)
}

func init() {
	clustersCmd.Flags().StringVarP(
		&output,
		"output",
		"o",
		"table",
		"Output format (table|json)",
	)
}
