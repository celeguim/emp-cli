package exec

import (
	"bytes"
	"os/exec"
)

func Run(name string, args ...string) (string, error) {

	cmd := exec.Command(name, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return stderr.String(), err
	}

	return out.String(), nil
}