package main

import "github.com/spf13/cobra"

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start REST API server",
	Run: func(cmd *cobra.Command, args []string) {
		addr, _ := cmd.Flags().GetString("addr")
		// TODO: open DB, register API routes, start HTTP server
		_ = addr
	},
}

func init() {
	serveCmd.Flags().StringP("addr", "a", ":8080", "listen address")
}