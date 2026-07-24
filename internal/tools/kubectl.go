package tools

import (
	"strings"

	"github.com/celeguim/emp-cli/internal/runner"
)

type Kubectl struct{}

func NewKubectl() *Kubectl {
	return &Kubectl{}
}

func (k *Kubectl) Name() string {
	return "kubectl"
}

func (k *Kubectl) Exists() bool {
	return runner.Exists("kubectl")
}

func (k *Kubectl) Version() string {

	if !k.Exists() {
		return "not installed"
	}

	result, err := runner.Run(
		"kubectl",
		"version",
		"--client",
	)

	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(result.Stdout)
}
