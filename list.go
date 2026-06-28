package main

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list <subnet>",
	Short: "List all hosts in a subnet",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		subnet := args[0]
		// TODO: open DB, list hosts
		_ = subnet
	},
}