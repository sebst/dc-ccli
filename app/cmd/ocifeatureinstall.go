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
	Use:   "ocifeatureinstall <ociFeature>",
	Short: "Installs a devcontainer feature from an OCI registry",
	Long: `Installs a devcontainer feature from an OCI registry. 
		e.g. dc-ccli ocifeatureinstall "ghcr.io/sebst/devcontainer-features/debug-dump-env:1"
	`,
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
