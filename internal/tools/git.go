package tools

import (
	"strings"

	"github.com/celeguim/emp-cli/internal/runner"
)

type Git struct{}

func NewGit() *Git {
	return &Git{}
}

func (k *Git) Name() string {
	return "git"
}

func (k *Git) Exists() bool {
	return runner.Exists("git")
}

func (k *Git) Version() string {

	if !k.Exists() {
		return "not installed"
	}

	result, err := runner.Run(
		"git",
		"--version",
	)

	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(result.Stdout)
}
