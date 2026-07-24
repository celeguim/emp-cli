package doctor

import "github.com/celeguim/emp-cli/internal/runner"

type Check struct {
	Name    string
	OK      bool
	Message string
}

func Run() []Check {

	checks := []Check{
		checkTool("kubectl"),
		checkTool("argocd"),
		checkTool("helm"),
		checkTool("git"),
	}

	return checks
}

func checkTool(name string) Check {

	if runner.Exists(name) {
		return Check{
			Name:    name,
			OK:      true,
			Message: "installed",
		}
	}

	return Check{
		Name:    name,
		OK:      false,
		Message: "not installed",
	}
}
