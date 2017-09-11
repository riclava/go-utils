package goutils

// IsDockerEnvironment check whether current runtime is in docker
func IsDockerEnvironment() bool {
	// check if node is docker env
	cmd := "cat /proc/1/cgroup | grep kubepods | wc -l"
	procBody, _ := ExecCmd(cmd)
	if procBody != "0\n" {
		return true
	}
	return false
}
