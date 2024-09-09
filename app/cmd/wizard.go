/*
Copyright © 2024 devcontainer.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// wizardCmd represents the wizard command
var wizardCmd = &cobra.Command{
	Use:   "wizard",
	Short: "Run a devcontainer.com wizard",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wizard called")
	},
}

func init() {
	rootCmd.AddCommand(wizardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wizardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wizardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
