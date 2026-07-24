package cmd

import (
	"fmt"
	"strings"

	"github.com/celeguim/emp-cli/internal/argocd"
	"github.com/celeguim/emp-cli/internal/runner"
	"github.com/spf13/cobra"
)

var projectsOutput string

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List Argo CD projects",

	RunE: func(cmd *cobra.Command, args []string) error {

		client := argocd.New()

		result, err := projectList(client)
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

func projectList(client *argocd.Client) (*runner.Result, error) {

	switch projectsOutput {

	case "json":
		return client.ProjectListJSON()

	default:
		return client.ProjectList()
	}
}

func init() {

	rootCmd.AddCommand(projectsCmd)

	projectsCmd.Flags().StringVarP(
		&projectsOutput,
		"output",
		"o",
		"table",
		"Output format (table|json)",
	)
}
