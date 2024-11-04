/*
Copyright © 2024 devcontainer.com
*/
package customizer

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// func writeTextFile(filePath string, content string, mode int) error {
// 	if err := os.WriteFile(filePath, []byte(content), mode); err != nil {
// 		return fmt.Errorf("failed to write file: %w", err)
// 	}
// 	return nil
// }

func writeTextFile(filePath string, content string) error {
	// Write the content to the specified path
	if filePath[0] != '/' {
		home := os.Getenv("HOME")
		filePath = filepath.Join(home, filePath)
	}
	const fileMode = 0644
	if err := os.WriteFile(filePath, []byte(content), fileMode); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func stringToFileMode(s string) (os.FileMode, error) {
	// Parse the string as an octal number
	mode, err := strconv.ParseUint(s, 8, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid file mode: %w", err)
	}

	// Convert the parsed number to os.FileMode
	return os.FileMode(mode), nil
}

func createFile(filePath string, file *File) error {
	if filePath[0] != '/' {
		basePath := getBasePath()
		filePath = filepath.Join(basePath, filePath)
	}
	var fileContent []byte
	if file.Content.Text != "" {
		fileContent = []byte(file.Content.Text)
	}
	// fileMode, err := stringToFileMode(file.Permissions)
	// if err != nil {
	// 	return fmt.Errorf("failed to parse file mode: %w", err)
	// }

	// Write the file content to the specified path
	fileMode := os.FileMode(0644)
	if err := os.WriteFile(filePath, fileContent, fileMode); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	// set the file owner and group
	// if file.Owner.UID != 0 || file.Group.GID != 0 {
	// 	if err := os.Chown(filePath, file.Owner.UID, file.Group.GID); err != nil {
	// 		return fmt.Errorf("failed to change file owner and group: %w", err)
	// 	}
	// }

	return nil
}

// func applyGitHubDorFilesRepo(repo string, branch string, destination string, script string) error {
// 	fmt.Println("Cloning GitHub repo:", repo)
// 	fmt.Println("Branch:", branch)
// 	fmt.Println("Destination:", destination)
// 	fmt.Println("Script:", script)

// 	return nil
// }

func ApplyFiles(config *Config) error {
	// basePath := os.Getenv("HOME")

	for _, file := range config.Files {
		destination := file.Path
		err := createFile(destination, &file)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
	}

	// for file := range config.Files {
	// 	destination := filepath.Join(basePath, filePath)
	// 	err := createFile(destination, &file)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to create file: %w", err)
	// 	}
	// }

	// gitHubRepo := config.Dotfiles.Github.Repo
	// gitHubRepoBranch := config.Dotfiles.Github.Branch
	// gitHubRepoPath := config.Dotfiles.Github.Path
	// gitHubRepoDestination := filepath.Join(home, gitHubRepoPath)
	// gitHubRepoScript := config.Dotfiles.Github.Install

	// if gitHubRepo != "" && gitHubRepoBranch != "" && gitHubRepoPath != "" {
	// 	err := applyGitHubDorFilesRepo(gitHubRepo, gitHubRepoBranch, gitHubRepoDestination, gitHubRepoScript)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to clone GitHub repo: %w", err)
	// 	}
	// }

	return nil
}
