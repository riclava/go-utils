package goutils

import (
	"os/exec"
)

// ExecCmd exec a shell command and return its output, if error error is not nil
func ExecCmd(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
