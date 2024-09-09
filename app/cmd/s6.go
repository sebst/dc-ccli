/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// s6Cmd represents the s6 command
var s6Cmd = &cobra.Command{
	Use:   "s6",
	Short: "Manage s6-overlay",
	Long:  `This command manages the s6-overlay commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("s6 called")
	},
}

func init() {
	rootCmd.AddCommand(s6Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// s6Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// s6Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
