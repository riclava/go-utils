package goutils

import (
	"testing"
)

// TestGetMemInfo test GetMemInfo()
func TestGetMemInfo(t *testing.T) {
	info, err := GetMemInfo()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(info)
	}
}
