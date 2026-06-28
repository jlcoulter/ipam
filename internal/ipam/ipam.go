package ipam

import "github.com/jlcoulter/ipam/internal/db"

// IPAM provides IP address management operations
type IPAM struct {
	db *db.DB
}

func NewIPAM(database *db.DB) *IPAM {
	return &IPAM{db: database}
}

// Allocate finds and allocates an IP address in the given subnet
func (i *IPAM) Allocate(subnet, host, specificIP, mac, role, desc, tags string) (string, error) {
	// TODO: If specificIP provided check that it's free and allocate
	// TODO: Otherwise, find next free ip in subnet and allocate
	// TODO: Return allocated IP
	return "", nil
}

// Release returns an IP address to the free pool
func (i *IPAM) Release(ip string) error {
	//TODO: Release IP back to pool
	return nil
}

// List returns all hosts in a subnet
func (i *IPAM) List(subnet string) ([]db.Host, error) {
	return nil, nil
	// TODO: List all hosts in subnet
}

// Show returns details for a specific host
func (i *IPAM) Show(ip string) (*db.Host, error) {
	return nil, nil
	// TODO: Show details for host
}

// AddSubnet registered a new subnet
func (i *IPAM) AddSubnet(cidr, name, desc string, vlan int) error {
	return nil
	// TODO: Validate cidr, create subnet record
}

// Subnets returns all registered subnets
func (i *IPAM) Subnets() ([]db.Subnet, error) {
	return nil, nil
	// TODO: List all subnets
}

// UpsertHost inserts or updates a host record - used by scan/import
func (i *IPAM) UpsertHost(host *db.Host) error {
	return nil
	// TODO: Upsert host record
}
