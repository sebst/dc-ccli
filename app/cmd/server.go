/*
Copyright © 2024 devcontainer.com
*/
package cmd

import (
	"fmt"

	"devcontainer.com/ccli/internal/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "devcontainer.com Container API & Dashboard",
	Long:  `The server component provides a REST API and a dashboard for your devcontainer.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		server.RunApp(6867)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
