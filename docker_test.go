package goutils

import (
	"testing"
)

// TestIsDockerEnvironment test IsDockerEnvironment()
func TestIsDockerEnvironment(t *testing.T) {
	isDockerEnv := IsDockerEnvironment()
	if isDockerEnv {
		t.Log("docker env")
	} else {
		t.Log("non docker env")
	}
}
