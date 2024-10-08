package utils

import "net"

func IsValidIp(ip string) bool {
	parsed := net.ParseIP(ip)
	return parsed != nil
}

func IsValidPort(port int) bool {
	return port > 0 && port < 65536
}

func IsValidProtocol(protocol string) bool {
	switch protocol {
	case PROTOCOL_TCP, PROTOCOL_UDP:
		return true
	}
	return false
}

const (
	PROTOCOL_TCP = "tcp"
	PROTOCOL_UDP = "udp"
)
