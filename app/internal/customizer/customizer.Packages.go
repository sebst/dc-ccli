/*
Copyright © 2024 devcontainer.com
*/
package customizer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func aptInstall(packages []string) error {
	// Install the packages
	updateCommand := "sudo apt-get update"
	fmt.Println("Running command:", updateCommand)
	updateCmd := exec.Command("bash", "-c", updateCommand)
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	err := updateCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	command := "sudo apt-get install -y " + strings.Join(packages, " ")
	fmt.Println("Running command:", command)
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}
	return nil
}

func ApplyPackages(config *Config) error {
	return nil
	// aptPackages := []string{}
	// for _, pkg := range config.Packages {
	// 	if pkg.Manager == "apt" {
	// 		aptPackages = append(aptPackages, pkg.Name)
	// 	}
	// }
	// return aptInstall(aptPackages)
}
