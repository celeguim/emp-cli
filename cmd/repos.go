package cmd

import (
	"fmt"
	"strings"

	"github.com/celeguim/emp-cli/internal/argocd"
	"github.com/celeguim/emp-cli/internal/runner"
	"github.com/spf13/cobra"
)

var reposOutput string

var reposCmd = &cobra.Command{
	Use:   "repos",
	Short: "List Argo CD repositories",

	RunE: func(cmd *cobra.Command, args []string) error {

		client := argocd.New()

		result, err := repoList(client)
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

func repoList(client *argocd.Client) (*runner.Result, error) {

	switch reposOutput {

	case "json":
		return client.RepoListJSON()

	default:
		return client.RepoList()
	}
}

func init() {

	rootCmd.AddCommand(reposCmd)

	reposCmd.Flags().StringVarP(
		&reposOutput,
		"output",
		"o",
		"table",
		"Output format (table|json)",
	)
}
