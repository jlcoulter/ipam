package main

import (
	"os"

	"github.com/spf13/cobra"
)

var dbPath string

func main() {
	rootCmd := &cobra.Command{
		Use:   "ipam",
		Short: "IP address management and host tracker",
	}
	rootCmd.PersistentFlags().StringVar(&dbPath, "db", "ipam.db", "./")

	rootCmd.AddCommand(
		allocateCmd,
		releaseCmd,
		listCmd,
		showCmd,
		subnetCmd,
		serveCmd,
		scanCmd,
		importCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

