/*
Copyright © 2024 devcontainer.com
*/

// see: https://pkg.go.dev/oras.land/oras-go/v2#example-package-PullFilesFromRemoteRepository

package ocifeatureinstall

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	oras "oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/file"
	"oras.land/oras-go/v2/registry"
	"oras.land/oras-go/v2/registry/remote"
)

func untar(filename string) error {
	// Even though the file extension is `tgz`, the devcontainer feature is actually a tar file
	// using the standard tar format in go, it fails to extract the file
	// so we use the `tar` command to extract the file which seem to work

	// Open the tar or tgz file
	fileDirectory := filepath.Dir(filename)
	relativePath := filepath.Base(filename)

	// in working directory `fileDirectory`, run command `tar -xvf relativePath`
	cmd := exec.Command("tar", "-xf", relativePath)
	cmd.Dir = fileDirectory
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to untar file: %w", err)
	}

	return nil
}

func DownloadAndExtractOCIArtifact(artifact string) (string, error) {
	tmpDirName := filepath.Join("/tmp", fmt.Sprintf("oras_extract_%d", rand.Int()))

	// 0. Create a file store
	fs, err := file.New(tmpDirName)
	if err != nil {
		panic(err)
	}
	defer fs.Close()

	ref, err := registry.ParseReference(artifact)
	if err != nil {
		return "", fmt.Errorf("failed to parse artifact reference: %w", err)
	}

	// 1. Connect to a remote repository
	ctx := context.Background()
	repo, err := remote.NewRepository(ref.Host() + "/" + ref.Repository)
	if err != nil {
		panic(err)
	}

	// // Note: The below code can be omitted if authentication is not required
	// repo.Client = &auth.Client{
	// 	Client: retry.DefaultClient,
	// 	Cache:  auth.NewCache(),
	// 	Credential: auth.StaticCredential(reg, auth.Credential{
	// 		Username: "username",
	// 		Password: "password",
	// 	}),
	// }

	// 2. Copy from the remote repository to the file store
	var tag string
	if ref.Reference != "" {
		tag = ref.Reference
	} else {
		tag = "latest"
	}

	manifestDescriptor, err := oras.Copy(ctx, repo, tag, fs, tag, oras.DefaultCopyOptions)
	if err != nil {
		fmt.Println(manifestDescriptor)
		return "", fmt.Errorf("failed to copy from remote repository: %w", err)
	}

	// iterrate over the files in the directory
	files, err := os.ReadDir(tmpDirName)
	if err != nil {
		return "", fmt.Errorf("failed to read directory: %w", err)
	}
	for _, file := range files {
		// fmt.Println(file.Name())
		fileName := file.Name()
		// check if fileName starts with "devcontainer-feature-"
		if fileName[:21] == "devcontainer-feature-" {
			// check if fileName ends with ".tgz"
			if fileName[len(fileName)-4:] == ".tgz" {
				// fmt.Println("found devcontainer feature")
				err := untar(tmpDirName + "/" + fileName)
				if err != nil {
					return "", fmt.Errorf("failed to untar file: %w", err)
				}

				return tmpDirName, nil
				// tt, err := parseDevcontainerFeatureJSONFile(tmpDirName + "/devcontainer-feature.json")
				// if err != nil {
				// 	return "", fmt.Errorf("failed to parse devcontainer feature json file: %w", err)
				// }
				// // fmt.Println(tt.Options)
				// getDefaultOptions(tt.Options)

			}
		}
	}

	// return tmpDirName, nil
	return "", fmt.Errorf("artifact is not a valid devcontainer feature")

}

func install(dirName string, options map[string]string) error {
	installSh := filepath.Join(dirName, "install.sh")
	if _, err := os.Stat(installSh); os.IsNotExist(err) {
		return fmt.Errorf("install.sh not found in directory: %s", dirName)
	}
	// fmt.Println("Running install.sh")
	cmd := exec.Command(installSh)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dirName

	var envs []string
	for key, value := range options {
		envs = append(envs, fmt.Sprintf("%s=\"%s\"", key, value))
	}
	cmd.Env = append(os.Environ(), envs...)

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("failed to run install.sh: %w", err)
	}

	return nil
}

func DownloadAndInstallOCIArtifact(artifact string) error {
	tmpDirName, err := DownloadAndExtractOCIArtifact(artifact)
	if err != nil {
		return fmt.Errorf("failed to download and extract OCI artifact: %w", err)
	}
	devcontainerManifest, err := parseDevcontainerFeatureJSONFile(tmpDirName + "/devcontainer-feature.json")
	if err != nil {
		return fmt.Errorf("failed to parse devcontainer feature json file: %w", err)
	}
	defaultOptions := getDefaultOptions(devcontainerManifest.Options)
	// fmt.Println("setting default options", defaultOptions)
	install(tmpDirName, defaultOptions)
	// fmt.Println("installing")
	return nil
}

func getValueFromEnvOrArgument(key string, arguments []string) string {
	// TODO: implement this function
	// TODO: Actually use this function
	osEnv := os.Getenv(key)
	if osEnv != "" {
		return osEnv
	}
	return ""
}

func getDefaultOptions(options map[string]FeatureOption) map[string]string {

	var variables = make(map[string]string)

	for key, option := range options {
		keyUppercase := strings.ToUpper(key)
		variables[keyUppercase] = option.Default.(string)
	}

	return variables

	// for _, option := range options {
	// 	fmt.Println(option)
	// }
}
