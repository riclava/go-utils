package goutils

import (
	"testing"
)

// TestGetIpv4Addr test GetIpv4Addr()
func TestGetIpv4Addr(t *testing.T) {
	addr := GetIpv4Addr()
	if addr != "unknown" {
		t.Log(addr)
	} else {
		t.Error("Can not get Ipv4 addr")
	}
}
