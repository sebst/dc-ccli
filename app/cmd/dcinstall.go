/*
Copyright © 2024 devcontainer.com
*/
package cmd

import (
	"fmt"

	"devcontainer.com/ccli/internal/dcinstaller"
	"github.com/spf13/cobra"
)

// dcinstallCmd represents the dcinstall command
var dcinstallCmd = &cobra.Command{
	Use:   "dcinstall",
	Short: "Installs devcontainer.com packages",
	Long: `Installs devcontainer.com packages. For example:
- dc-s6-overlay
- dc-sshd
- dc-sshd-config-global
- dc-s6-service-sshd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dcinstall called")
		for _, arg := range args {
			dcinstaller.DCInstall(arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(dcinstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dcinstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dcinstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
