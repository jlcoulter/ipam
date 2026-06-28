package main

import "github.com/spf13/cobra"

var showCmd = &cobra.Command{
	Use:   "show <ip-address>",
	Short: "Show details for a specific host",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := args[0]
		// TODO: open DB, show host details
		_ = ip
	},
}