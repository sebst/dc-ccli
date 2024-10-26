/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"devcontainer.com/ccli/internal/ocifeatureinstall"
	"github.com/spf13/cobra"
)

// ocifeatureinstallCmd represents the ocifeatureinstall command
var ocifeatureinstallCmd = &cobra.Command{
	Use:   "ocifeatureinstall",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: ccli ocifeatureinstall <ociFeature>")
			return
		}
		ociFeature := args[0]

		err := ocifeatureinstall.DownloadAndInstallOCIArtifact(ociFeature)
		if err != nil {
			fmt.Printf("Failed to download and extract OCI artifact: %v\n", err)
			return
		}
		// fmt.Println(dir)
	},
}

func init() {
	rootCmd.AddCommand(ocifeatureinstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ocifeatureinstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ocifeatureinstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
