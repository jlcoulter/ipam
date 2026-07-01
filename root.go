package main

import (
	"fmt"
	"os"

	"github.com/jlcoulter/ipam/internal/db"
	"github.com/jlcoulter/ipam/internal/ipam"
)

func openIPAM() (*ipam.IPAM, *db.DB, error) {
	database, err := db.Open(dbPath)
	if err != nil {
		return nil, nil, fmt.Errorf("open db: %w", err)
	}

	return ipam.NewIPAM(database), database, nil
}

func mustOpenIPAM() (*ipam.IPAM, *db.DB) {
	ipamSvc, dbObj, err := openIPAM()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return ipamSvc, dbObj
}
