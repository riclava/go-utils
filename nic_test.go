package goutils

import (
	"testing"
)

// TestGetHTTP test GetHTTP()
func TestGetNicInfo(t *testing.T) {
	info, err := GetNicInfo()
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("<", info, ">")
	}
}
