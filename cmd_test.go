package goutils

import (
	"testing"
)

// TestExecCmd test ExecCmd()
func TestExecCmd(t *testing.T) {
	body, err := ExecCmd("uname -a")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(body)
	}
}
