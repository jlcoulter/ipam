package ipam

import "time"

// DiscovereredHost represents a host found during a network scan
type DiscoveredHost struct {
	IP       string
	Hostname string
	MAC      string
	Alive    bool
}

// ScanSubnet performs a ping swwp of a CIDR range
func ScanSubnet(cidr string) ([]DiscoveredHost, error) {
	// TODO: Parse CIDR, iterate all IPs
	// TODO: Ping each IP concurrently (goroutine pool, ~50 workers)
	// TODO: Collect responding hosts, attempt MAC resolution
	// TODO: Return list of discvoered hosts
	return nil, nil
}

// pingIP sends an ICMP echo request an waits for a reply
func pingIP(ip string, timeout time.Duration) bool {
	// TODO: Use net.Dial("ip4:icmp", ...) or exec("ping", "-c", "1", "-W", "1", ip)
	// TODO: Return true if host responds within timeout
	return false
}

// resolveMAC attempts to look up a MAC address from the ARP table
func resolveMAC(ip string) string {
	// TODO: Read /proc/net/arp (linux) or exec("arp", "-n", ip)
	// TODO: Return MAC string or empty string if not found
	return ""
}
