package argocd

import "github.com/celeguim/emp-cli/internal/runner"

func (c *Client) ProjectList() (*runner.Result, error) {
	return runner.Run(
		"argocd",
		"proj",
		"list",
	)
}

func (c *Client) ProjectListJSON() (*runner.Result, error) {
	return runner.Run(
		"argocd",
		"proj",
		"list",
		"-o",
		"json",
	)
}
