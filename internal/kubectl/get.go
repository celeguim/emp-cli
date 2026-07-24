package kubectl

import "github.com/celeguim/emp-cli/internal/runner"

func (c *Client) Get(args ...string) (*runner.Result, error) {

	cmd := []string{}

	if c.Context != "" {
		cmd = append(cmd, "--context", c.Context)
	}

	if c.Namespace != "" {
		cmd = append(cmd, "-n", c.Namespace)
	}

	cmd = append(cmd, "get")
	cmd = append(cmd, args...)

	return runner.Run("kubectl", cmd...)
}
