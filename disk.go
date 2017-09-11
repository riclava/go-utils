package goutils

import "strings"

// GetDiskInfo get disk size
func GetDiskInfo() (string, error) {
	disks, err := ExecCmd("fdisk -l | grep \"Disk /dev/\" | awk '{print $2\":\"$3\":\"$4}'")
	if err != nil {
		return "N/A", err
	}
	disks = strings.Replace(disks, ",\n", ";", -1)
	disks = strings.Replace(disks, "::", ":", -1)
	return disks, nil
}
