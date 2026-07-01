package db

import (
	"database/sql"
	"fmt"
)

// Host represents a sing host on the network
type Host struct {
	IP          string
	SubnetCIDR  string
	Hostname    string
	MAC         string
	Role        string
	Description string
	Tags        string
	AssisngedAt string
	LastSeen    string
}

// Subnet represents a registered network subnet
type Subnet struct {
	ID          int
	CIDR        string
	Name        string
	Description string
	VLAN        int
	CreatedAt   string
}

// DB wraps the SQLite database connection
type DB struct {
	db *sql.DB
}

// Open creates the SQLite database and ensures the schema exists
func Open(path string) (*DB, error) {
	// TODO: sql.Open with SQLite driver
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	// TODO: PRAGMA journal_mode=WAL
	// TODO: PRGRMA foreign_keys=ON
	prags := []string{
		"PRAGMA journal_mode=WAL",
		"PRAGMA foreign_keys=ON",
		"PRAGMA busy_timeout=5000",
	}
	for _, p := range prags {
		if _, err := db.Exec(p); err != nil {
			return nil, fmt.Errorf("%s: %w", p, err)
		}
	}
	// TODO: Create tables if not exists for subnets and hosts
	schema := `
	CREATE TABLE IF NOT EXISTS subnets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		cidr TEXT NOT NULL UNIQUE,
		name TEXT,
		description TEXT,
		vlan INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS hosts (
		ip TEXT PRIMARY KEY,
		subnet_cidr TEXT NOT NULL,
		hostname TEXT,
		mac TEXT,
		role TEXT,
		description TEXT,
		tags TEXT,
	assigned_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	last_seen DATETIME,
	FOREIGN KEY (subnet_cidr) REFERENCES subnets(cidr)
	);`

	if _, err := db.Exec(schema); err != nil {
		return nil, fmt.Errorf("create scheme: %w", err)
	}
	return &DB{db: db}, nil
}

// CreateSubnet inserts a new subnet record
func (d *DB) CreateSubnet(cidr, name, desc string, vlan int) error {
	// TODO: INSERT INTO subnets (cidr, name, description, vlan) VALUES (...)
	_, err := d.db.Exec(
		"INSERT INTO subnets (cidr, name, description, vlan) VALUES (?, ?, ?, ?)",
		cidr, name, desc, vlan,
	)
	return err
}

// GetSubnet retrieves a subnet by CIDR
func (d *DB) GetSubnet(cidr string) (*Subnet, error) {
	return nil, nil
	// TODO: SELECT * FROM subnets WHERE cidr = ?
}

// ListSubnets returns all registered subnets
func (d *DB) ListSubnets() ([]Subnet, error) {
	// TODO: SELECT * FROM subnets ORDER BY cidr
	rows, err := d.db.Query(
		"SELECT id, cidr, name, description, vlan, created_at FROM subnets ORDER BY cidr",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subnets []Subnet
	for rows.Next() {
		var s Subnet
		if err := rows.Scan(
			&s.ID,
			&s.CIDR,
			&s.Name,
			&s.Description,
			&s.VLAN,
			&s.CreatedAt,
		); err != nil {
			return nil, err
		}
		subnets = append(subnets, s)
	}
	return subnets, rows.Err()
}

// UpsertHost inserts or updates a host record (ON CONFLICT ip)
func (d *DB) UpsertHost(host *Host) error {
	return nil
	// TODO: INSERT OR REPLACE INTO hosts VALUES (...)
}

// AllocateIP inserts a new host record with an allocated IP
func (d *DB) AllocateIP(ip, subnet, host, mac, role, desc, tags string) error {
	return nil
	// TODO: INSERT INTO hosts (ip, subnet_cidr, hostname, mac, role, description, tags) VALUES (...)
}

// ReleaseIP removes a host record
func (d *DB) ReleaseIP(ip string) error {
	return nil
	// TODO: DELETE FROM hosts WHERE ip = ?
}

// GetHost retrieves a host by IP address
func (d *DB) GetHost(ip string) (*Host, error) {
	return nil, nil
	// TODO: SELECT * FROM hosts WHERE ip = ?
}

// ListHosts returns all hosts in a subnet
func (d *DB) ListHosts(subnet string) ([]Host, error) {
	return nil, nil
	// TODO: SELECT * FROM hosts WHERE subnet_cidr = ?
}

// GetUsedCount returns the number of hosts in a subnet
func (d *DB) GetUsedCount(subnet string) (int, error) {
	return 0, nil
	// TODO: SELECT COUNT(*) FROM hosts WHERE subnet_cidr = ?
}

// Close the database connection
func (d *DB) Close() error {
	return d.db.Close()
}
