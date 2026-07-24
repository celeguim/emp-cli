package doctor

import "github.com/celeguim/emp-cli/internal/runner"

type Check struct {
	Name    string
	OK      bool
	Message string
}

func Run() []Check {

	return []Check{
		checkTool("kubectl", "version", "--client", "--short"),
		checkTool("argocd", "version", "--client"),
		checkTool("helm", "version", "--short"),
		checkTool("git", "--version"),
	}
}

func checkTool(name string, versionArgs ...string) Check {

	if !runner.Exists(name) {
		return Check{
			Name:    name,
			OK:      false,
			Message: "not installed",
		}
	}

	version, err := runner.Version(name, versionArgs...)
	if err != nil {
		version = "installed"
	}

	return Check{
		Name:    name,
		OK:      true,
		Message: version,
	}
}
