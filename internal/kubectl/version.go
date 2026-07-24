package kubectl

import "github.com/celeguim/emp-cli/internal/runner"

func (c *Client) Version() (*runner.Result, error) {

	return runner.Run(
		"kubectl",
		"version",
		"--client",
	)
}
