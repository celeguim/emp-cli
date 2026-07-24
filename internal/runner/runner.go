package runner

import (
	"bytes"
	"os/exec"
	"strings"
)

type Result struct {
	Stdout string
	Stderr string
	Code   int
}

func Exists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func Version(command string, args ...string) (string, error) {

	result, err := Run(command, args...)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(result.Stdout), nil
}

func Run(name string, args ...string) (*Result, error) {

	cmd := exec.Command(name, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	result := &Result{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
		Code:   0,
	}

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			result.Code = exitErr.ExitCode()
		} else {
			result.Code = -1
		}
		return result, err
	}

	return result, nil
}
