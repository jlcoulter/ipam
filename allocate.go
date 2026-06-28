package main

import "fmt"
import "github.com/spf13/cobra"

var allocateCmd = &cobra.Command{
	Use:   "allocate <subnet> --host <hostname> [--ip <address>]",
	Short: "Allocate an IP address from a subnet",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		subnet := args[0]
		host, _ := cmd.Flags().GetString("host")
		// TODO: open DB, allocate IP
		fmt.Printf("allocating from %s for %s\n", subnet, host)
	},
}

func init() {
	allocateCmd.Flags().StringP("host", "", "", "hostname for the allocation")
	allocateCmd.Flags().StringP("ip", "", "", "specific IP to allocate")
}
