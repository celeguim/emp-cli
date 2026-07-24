package cmd

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/argocd"
	"github.com/celeguim/emp-cli/internal/runner"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List Argo CD projects",

	RunE: func(cmd *cobra.Command, args []string) error {

		client := argocd.New()
		result, err := projectList(client)

		if err != nil {
			return argocd.HandleError(result.Stderr, err)
		}

		fmt.Print(result.Stdout)

		return nil
	},
}

func projectList(client *argocd.Client) (*runner.Result, error) {

	switch output {
	case "json":
		return client.ProjectListJSON()

	default:
		return client.ProjectList()
	}
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}

func init() {
	projectsCmd.Flags().StringVarP(
		&output,
		"output",
		"o",
		"table",
		"Output format (table|json)",
	)
}
