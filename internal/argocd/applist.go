package argocd

import "github.com/celeguim/emp-cli/internal/runner"

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) AppList() (*runner.Result, error) {

	return runner.Run(
		"argocd",
		"app",
		"list",
	)
}

func (c *Client) AppListJSON() (*runner.Result, error) {

	return runner.Run(
		"argocd",
		"app",
		"list",
		"-o",
		"json",
	)
}
