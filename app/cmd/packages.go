/*
Copyright © 2024 devcontainer.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// packagesCmd represents the packages command
var packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "devcontainer.com package installer",
	Long:  `Manage packages in your container.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("packages called")
	},
}

func init() {
	rootCmd.AddCommand(packagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// packagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// packagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
