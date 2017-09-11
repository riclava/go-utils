package goutils

import (
	"testing"
)

// TestGetDiskInfo test GetDiskInfo()
func TestGetDiskInfo(t *testing.T) {
	info, err := GetDiskInfo()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(info)
	}
}
