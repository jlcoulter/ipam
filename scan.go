package main

import "github.com/spf13/cobra"

var scanCmd = &cobra.Command{
	Use:   "scan <cidr>",
	Short: "Ping sweep a subnet to discover live hosts",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cidr := args[0]
		// TODO: parse CIDR, ping sweep, upsert discovered hosts
		_ = cidr
	},
}