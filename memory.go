package goutils

// GetMemInfo get mem info
func GetMemInfo() (string, error) {
	mem, err := ExecCmd("free -m | grep Mem | awk '{print $2}'")
	if err != nil {
		return "", err
	}
	return "Mem: " + mem + "MB", nil
}
