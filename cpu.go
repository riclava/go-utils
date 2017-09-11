package goutils

import "strings"

// GetCPUInfo get cpu info
func GetCPUInfo() (string, error) {
	model, err := ExecCmd("cat /proc/cpuinfo | grep \"model name\" | tail -1")
	if err != nil {
		return "N/A", err
	}
	cores, err := ExecCmd("cat /proc/cpuinfo | grep \"model name\" | wc -l")
	if err != nil {
		return "N/A", err
	}
	model = strings.Replace(model, "\n", "", -1)
	cores = strings.Replace(cores, "\n", "", -1)
	return "Cores: " + cores + "; " + model, nil
}
