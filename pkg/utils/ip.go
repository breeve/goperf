package utils

import "net"

func IsValidIp(ip string) bool {
	parsed := net.ParseIP(ip)
	return parsed != nil
}

func IsValidPort(port int) bool {
	return port > 0 && port < 65536
}
