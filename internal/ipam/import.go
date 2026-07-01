package ipam

import "encoding/xml"

// NmapXML represents the top-level nmap XML structure
type NmapXML struct {
	XMLName xml.Name   `xml:"nmaprun"`
	Hosts   []NmapHost `xml:"host"`
}

// NmapHOst represents a single host in nmap output
type NmapHost struct {
	Addresses []NmapAddress  `xml:"address"`
	Hostnames []NmapHostname `xml:"hostnames>hostname"`
	OS        NmapOS         `xml:"os"`
}

// NmapAddress represents an address element in nmap output
type NmapAddress struct {
	Addr string `xml:"addr,attr"`
	Type string `xml:"addrtype,attr"`
}

// NmapHostname represents a hostname in nmap output
type NmapHostname struct {
	Name string `xml:"name,attr"`
}

// NmapOS represents OS detection results in nmap output
type NmapOS struct {
	OSMatch []NmapOSMatch `xml:"osmatch"`
}

// NmapOSMatch represents an OS match in nmap output
type NmapOSMatch struct {
	Name string `xml:"name,attr"`
}

// ImportNmap parses an nmap XML file and returns discovered hosts
func ImportNmap(path string) ([]DiscoveredHost, error) {
	// TODO: Open and parse XML file (encoding/xml)
	// TODO: Extract IP (addrtype="ipv4"), hostname, MAC (addrtype="mac")
	// TODO: Map to DiscoveredHost structs
	// TODO: Return list
	return nil, nil
}
