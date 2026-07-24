package argocd

import "github.com/celeguim/emp-cli/internal/runner"

func (c *Client) RepoList() (*runner.Result, error) {
	return runner.Run(
		"argocd",
		"repo",
		"list",
	)
}

func (c *Client) RepoListJSON() (*runner.Result, error) {
	return runner.Run(
		"argocd",
		"repo",
		"list",
		"-o",
		"json",
	)
}
