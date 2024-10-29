/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"devcontainer.com/ccli/internal/customizer"
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var customizationApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apply called")
		if len(args) != 2 {
			fmt.Println("please provide id and api key")
			return
		}
		id := args[0]
		apiKey := args[1]
		url := "https://dc-dotfiles-web.contact-c10.workers.dev/api/profiles/" + id
		fmt.Println("id:", id)
		fmt.Println("url:", url)
		config, err := customizer.ReadConfigFromUrl(url, apiKey)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("config:", config)

		err = customizer.ApplyFiles(config)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		// err = customizer.ApplyPackages(config)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// 	return
		// }

		err = customizer.ApplyDotfiles(config)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	},
}

func init() {
	customizationCmd.AddCommand(customizationApplyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
