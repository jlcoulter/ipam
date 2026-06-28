package main

import "github.com/spf13/cobra"

var subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "Manage subnets",
}

var subnetShowCmd = &cobra.Command{
	Use:   "show <cidr>",
	Short: "Show subnet summary (used/free/total)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cidr := args[0]
		// TODO: calculate subnet size, query DB for used count
		_ = cidr
	},
}

var subnetAddCmd = &cobra.Command{
	Use:   "add <cidr> --name <name>",
	Short: "Register a new subnet",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cidr := args[0]
		name, _ := cmd.Flags().GetString("name")
		// TODO: insert subnet into DB
		_, _ = cidr, name
	},
}

var subnetListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered subnets",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: list all subnets
	},
}

func init() {
	subnetAddCmd.Flags().StringP("name", "n", "", "subnet name")
	subnetAddCmd.Flags().StringP("desc", "d", "", "description")
	subnetAddCmd.Flags().IntP("vlan", "v", 0, "VLAN ID")
	subnetCmd.AddCommand(subnetShowCmd, subnetAddCmd, subnetListCmd)
}