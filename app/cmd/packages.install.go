/*
Copyright © 2024 devcontainer.com
*/
package cmd

import (
	"errors"
	"strings"

	"devcontainer.com/ccli/internal/packageinstaller"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "devcontainer.com package installer",
	Long:  `Installs devcontainer.com packages to your container.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		packages := []*packageinstaller.Package{}
		for _, arg := range args {
			packageParts := strings.Split(arg, "=")
			packageName := packageParts[0]
			if len(packageParts) == 1 {
				packageParts = append(packageParts, "latest")
			}
			packages = append(packages, &packageinstaller.Package{
				Name:    packageName,
				Version: packageParts[1],
			})

		}
		packageinstaller.InstallPackages(packages)

	},
}

func init() {
	packagesCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
