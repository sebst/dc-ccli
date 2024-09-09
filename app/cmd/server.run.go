/*
Copyright © 2024 devcontainer.com
*/
package cmd

import (
	"devcontainer.com/ccli/internal/server"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "devcontainer.com Container API & Dashboard",
	Long:  `This command will start the devcontainer.com Container API & Dashboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := server.GetApp()
		server.RunServer(app)
	},
}

func init() {
	serverCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
