package main

import "github.com/spf13/cobra"

var releaseCmd = &cobra.Command{
	Use:   "release <ip-address>",
	Short: "Release an allocated IP address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := args[0]
		// TODO: open DB, release IP
		_ = ip
	},
}