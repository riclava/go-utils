package goutils

import (
	"net"
	"strings"
)

// GetIpv4Addr get ipv4 addr of a machine only if the addr is not lo
func GetIpv4Addr() string {
	IPAddr := "unknown"
	ifaces, err := net.Interfaces()
	if err != nil {
		return IPAddr
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				break
			case *net.IPAddr:
				ip = v.IP
				break
			}
			ipStr := string(ip.String())
			if ipStr != "127.0.0.1" && ipStr != "::1" && !strings.HasPrefix(string(ipStr), "fe80") {
				IPAddr = ipStr
				break
			}
		}
	}
	return IPAddr
}
