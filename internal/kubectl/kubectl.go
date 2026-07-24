package kubectl

import "github.com/celeguim/emp-cli/internal/runner"

func Get(args ...string) (*runner.Result, error) {

	cmd := append([]string{"get"}, args...)

	return runner.Run("kubectl", cmd...)
}

func Version() (*runner.Result, error) {

	return runner.Run(
		"kubectl",
		"version",
		"--client",
	)
}
