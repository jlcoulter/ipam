package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
		desc, _ := cmd.Flags().GetString("desc")
		vlan, _ := cmd.Flags().GetInt("vlan")

		ipamSvc, dbObj := mustOpenIPAM()
		defer dbObj.Close()

		// TODO: insert subnet into DB
		if err := ipamSvc.AddSubnet(cidr, name, desc, vlan); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("subnet %s added\n", cidr)
	},
}

var subnetListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered subnets",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: list all subnets
		ipamSvc, dbObj := mustOpenIPAM()
		defer dbObj.Close()

		subnets, err := ipamSvc.Subnets()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if len(subnets) == 0 {
			fmt.Println("no subnets registered ")
			return
		}

		fmt.Printf("%-4s %-20s %-15s %-5s %-20s\n", "ID", "CIDR", "Name", "VLAN", "Created")
		for _, s := range subnets {
			fmt.Printf("%-4d %-20s %-15s %-5d %-20s\n", s.ID, s.CIDR, s.Name, s.VLAN, s.CreatedAt)
		}
	},
}

func init() {
	subnetAddCmd.Flags().StringP("name", "n", "", "subnet name")
	subnetAddCmd.Flags().StringP("desc", "d", "", "description")
	subnetAddCmd.Flags().IntP("vlan", "v", 0, "VLAN ID")
	subnetCmd.AddCommand(subnetShowCmd, subnetAddCmd, subnetListCmd)
}

