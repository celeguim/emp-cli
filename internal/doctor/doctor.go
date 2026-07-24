package doctor

import (
	"strings"

	"github.com/celeguim/emp-cli/internal/runner"
)

type Check struct {
	Category string
	Name     string
	OK       bool
	Message  string
}

func Run() []Check {

	var checks []Check

	// ---------------------------------------------------------------------
	// Tools
	// ---------------------------------------------------------------------

	checks = append(checks,
		checkTool("kubectl"),
		checkTool("helm"),
		checkTool("git"),
		checkTool("argocd"),
	)

	// ---------------------------------------------------------------------
	// Connectivity
	// ---------------------------------------------------------------------

	checks = append(checks,
		checkKubernetes(),
		checkArgoCD(),
	)

	return checks
}

func checkTool(name string) Check {

	ok := runner.Exists(name)

	msg := "installed"
	if !ok {
		msg = "not installed"
	}

	return Check{
		Category: "Tools",
		Name:     name,
		OK:       ok,
		Message:  msg,
	}
}

func checkKubernetes() Check {

	result, err := runner.Run(
		"kubectl",
		"cluster-info",
	)

	if err != nil {

		msg := strings.TrimSpace(result.Stderr)
		if msg == "" {
			msg = err.Error()
		}

		return Check{
			Category: "Connectivity",
			Name:     "Kubernetes",
			OK:       false,
			Message:  msg,
		}
	}

	return Check{
		Category: "Connectivity",
		Name:     "Kubernetes",
		OK:       true,
		Message:  "connected",
	}
}

func checkArgoCD() Check {

	result, err := runner.Run(
		"argocd",
		"account",
		"get-user-info",
	)

	if err != nil {

		msg := strings.TrimSpace(result.Stderr)

		switch {

		case strings.Contains(msg, "Unauthenticated"):
			msg = "login required"

		case strings.Contains(msg, "connection refused"):
			msg = "server unreachable"

		case msg == "":
			msg = err.Error()
		}

		return Check{
			Category: "Connectivity",
			Name:     "Argo CD",
			OK:       false,
			Message:  msg,
		}
	}

	return Check{
		Category: "Connectivity",
		Name:     "Argo CD",
		OK:       true,
		Message:  "connected",
	}
}
