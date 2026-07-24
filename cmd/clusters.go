package cmd

import (
	"fmt"
	"strings"

	"github.com/celeguim/emp-cli/internal/argocd"
	"github.com/celeguim/emp-cli/internal/runner"
	"github.com/spf13/cobra"
)

var clustersOutput string

var clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "List Argo CD clusters",

	RunE: func(cmd *cobra.Command, args []string) error {

		client := argocd.New()

		result, err := clusterList(client)
		if err != nil {

			if strings.Contains(result.Stderr, "Unauthenticated") {
				return fmt.Errorf("you are not logged into Argo CD. Run 'argocd login'")
			}

			return err
		}

		fmt.Print(result.Stdout)

		return nil
	},
}

func clusterList(client *argocd.Client) (*runner.Result, error) {

	switch clustersOutput {

	case "json":
		return client.ClusterListJSON()

	default:
		return client.ClusterList()
	}
}

func init() {

	rootCmd.AddCommand(clustersCmd)

	clustersCmd.Flags().StringVarP(
		&clustersOutput,
		"output",
		"o",
		"table",
		"Output format (table|json)",
	)
}
