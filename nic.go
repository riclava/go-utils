package goutils

import (
	"net"
	"strings"
)

// GetNicInfo retrieve eth network interfaces info and it's speed, if not ok return error
func GetNicInfo() (string, error) {
	nics, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	nicInfo := ""
	for _, nic := range nics {
		if strings.Contains(nic.Name, "eth") || strings.Contains(nic.Name, "en") {
			nicInfo += nic.Name + "->"
			nicInfo += nic.HardwareAddr.String() + ","
			speed, err := ExecCmd("ethtool " + nic.Name + " | grep Speed")
			if err != nil {
				speed = "N/A"
			}
			nicInfo += speed + ";"
		}
	}
	return nicInfo, nil
}
