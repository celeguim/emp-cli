package cmd

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/argocd"
	"github.com/celeguim/emp-cli/internal/runner"
	"github.com/spf13/cobra"
)

var reposCmd = &cobra.Command{
	Use:   "repos",
	Short: "List Argo CD repositories",

	RunE: func(cmd *cobra.Command, args []string) error {

		client := argocd.New()
		result, err := repoList(client)

		if err != nil {
			return argocd.HandleError(result.Stderr, err)
		}

		fmt.Print(result.Stdout)

		return nil
	},
}

func repoList(client *argocd.Client) (*runner.Result, error) {

	switch output {
	case "json":
		return client.RepoListJSON()

	default:
		return client.RepoList()
	}
}

func init() {
	rootCmd.AddCommand(reposCmd)
}

func init() {
	reposCmd.Flags().StringVarP(
		&output,
		"output",
		"o",
		"table",
		"Output format (table|json)",
	)
}
