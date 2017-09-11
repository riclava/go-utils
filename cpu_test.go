package goutils

import (
	"testing"
)

// TestGetCPUInfo test GetCPUInfo()
func TestGetCPUInfo(t *testing.T) {
	info, err := GetCPUInfo()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(info)
	}
}
