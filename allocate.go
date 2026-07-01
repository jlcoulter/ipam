package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var allocateCmd = &cobra.Command{
	Use:   "allocate <subnet> --host <hostname> [--ip <address>] [--mac <mac>] [--role <role>] [--desc <desc>] [--tags <tags>]",
	Short: "Allocate an IP address from a subnet",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		subnet := args[0]
		host, _ := cmd.Flags().GetString("host")
		specificIP, _ := cmd.Flags().GetString("ip")
		mac, _ := cmd.Flags().GetString("mac")
		role, _ := cmd.Flags().GetString("role")
		desc, _ := cmd.Flags().GetString("desc")
		tags, _ := cmd.Flags().GetString("tags")
		// TODO: open DB, allocate IP
		fmt.Printf("allocating from %s for %s (ip=%s mac=%s role=%s desc=%s tags=%s)\n",
			subnet, host, specificIP, mac, role, desc, tags)
	},
}

func init() {
	allocateCmd.Flags().StringP("host", "", "", "hostname for the allocation")
	allocateCmd.Flags().StringP("ip", "", "", "specific IP to allocate")
	allocateCmd.Flags().StringP("mac", "", "", "MAC address")
	allocateCmd.Flags().StringP("role", "", "", "host role")
	allocateCmd.Flags().StringP("desc", "", "", "description")
	allocateCmd.Flags().StringP("tags", "", "", "comma-separated tags")
}
