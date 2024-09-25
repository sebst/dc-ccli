package dcinstaller

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

// DCInstall downloads a package installer script and creates a "run.sh" script
func DCInstall(packagename string) error {
	// Construct the URL
	url := "https://github.com/sebst/dc-features/raw/main/features/" + packagename + "/install.sh"

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "dcinstaller")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v", err)
	}

	// Define the paths for install.sh and run.sh
	installPath := filepath.Join(tempDir, "install.sh")
	runPath := filepath.Join(tempDir, "run.sh")

	// Download the content of install.sh
	err = downloadFile(url, installPath)
	if err != nil {
		return fmt.Errorf("failed to download install.sh: %v", err)
	}

	// Create the run.sh file and write "echo hello"
	content := "#!/usr/bin/env sh\n\ncd " + tempDir + "\nsudo VERSION=latest " + tempDir + "/install.sh \n"
	contentAsBytes := []byte(content)
	err = os.WriteFile(runPath, contentAsBytes, 0777)
	if err != nil {
		return fmt.Errorf("failed to create run.sh: %v", err)
	}

	// set the permissions of the install.sh file
	err = os.Chmod(installPath, 0777)
	if err != nil {
		return fmt.Errorf("failed to set permissions: %v", err)
	}

	fmt.Printf("Files created in %s\n", tempDir)

	cmd := exec.Command("sh", runPath)
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return errors.New("command failed")
	}
	fmt.Println(string(stdout))

	// Remove the temporary directory
	err = os.RemoveAll(tempDir)
	if err != nil {
		return fmt.Errorf("failed to remove temp dir: %v", err)
	}

	return nil
}

// Helper function to download the file from the URL
func downloadFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to get URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: status code %d", resp.StatusCode)
	}

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	// Copy the content from the response to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}
