/*
Copyright © 2024 devcontainer.com
*/
package cmd

import (
	"fmt"

	"devcontainer.com/ccli/internal/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of the devcontainer.com cCLI",
	Long:  `Version of the devcontainer.com cCLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
		version := version.GetVersion()
		fmt.Println("devcontainer.com cCLI version:", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
