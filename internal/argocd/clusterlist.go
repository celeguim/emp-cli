package argocd

import "github.com/celeguim/emp-cli/internal/runner"

func (c *Client) ClusterList() (*runner.Result, error) {
	return runner.Run(
		"argocd",
		"cluster",
		"list",
	)
}

func (c *Client) ClusterListJSON() (*runner.Result, error) {
	return runner.Run(
		"argocd",
		"cluster",
		"list",
		"-o",
		"json",
	)
}
