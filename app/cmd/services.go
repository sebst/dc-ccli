/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"devcontainer.com/ccli/internal/services"
	"github.com/spf13/cobra"
)

// servicesCmd represents the services command
var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("services called")
		if len(args) != 1 {
			fmt.Println("Please provide one argument: The JSON file")
			return
		}
		jsonFile := args[0]
		fmt.Println(jsonFile)
		// services.ServicesRunner(jsonFile)
		services.ServicesRunner(jsonFile)
	},
}

func init() {
	rootCmd.AddCommand(servicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// servicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// servicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
