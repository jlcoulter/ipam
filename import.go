package main

import "github.com/spf13/cobra"

var importCmd = &cobra.Command{
	Use:   "import <nmap-xml-file>",
	Short: "Import nmap XML scan results",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		// TODO: parse nmap XML, upsert hosts
		_ = path
	},
}