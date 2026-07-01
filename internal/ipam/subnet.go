package ipam

import (
	//"encoding/binary"
	"net"
)

// ParseCIDR parses a CIDR notion string into a net.IPNET
func ParseCIDR(cidr string) (*net.IPNet, error) {
	// TODO: net.ParseCIDR
	return nil, nil
}

func NextFreeIP(subnet *net.IPNet, used []string) (net.IP, error) {
	// TODO: Iterate from subnet start, skip used IPs
	// TODO: Skip network and broadcast addresses
	// TODO: Return first free IP
	return nil, nil
}

// IPToInt converts an IPv4 address to a uint32 using encoding/binary
func IPToInt(ip net.IP) uint32 {
	// TODO: Use encoding/binary.BigEndian.Uint32 on IP bytes
	return 0
}

// IntoToIp converts a uint32 to an IPv4 address
func IntoToIp(n uint32) net.IP {
	// TODO Use encoding/binary.BigEndian.PutUint32
	return nil
}

// SubnetSize calculates the total number of addresses in a subnet
func SubnetSize(subnet *net.IPNet) int {
	// TODO: 2^(32 - ones) for IPv4
	return 0
}
