/*
Copyright © 2024 devcontainer.com
*/
package customizer

import (
	"fmt"
	"os"
	"os/exec"
)

func cloneRepo(repo, branch, path string) error {
	// Clone the repository
	command := "git clone --branch " + branch + " " + repo + " " + path
	fmt.Println("Running command:", command)
	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = os.Getenv("HOME")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}
	return nil
}

func runInstallScript(path string, script string) error {
	// Run the install script
	command := "./" + script
	fmt.Println("Running command:", command)
	fmt.Println("Path:", path)
	cmd := exec.Command(command)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}
	return nil
}

func ApplyDotfiles(config *Config) error {
	// Check if the Dotfiles section is empty
	if config.Dotfiles.Github.Repo == "" {
		return nil
	}

	// Clone the repository
	home := os.Getenv("HOME")
	path := home + "/" + config.Dotfiles.Github.Path
	repo := "https://github.com/" + config.Dotfiles.Github.Repo + ".git"
	if err := cloneRepo(repo, config.Dotfiles.Github.Branch, path); err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	// Run the install script
	if err := runInstallScript(path, config.Dotfiles.Github.Install); err != nil {
		return fmt.Errorf("failed to run install script: %w", err)
	}

	return nil
}
