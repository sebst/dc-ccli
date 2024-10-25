/*
Copyright © 2024 devcontainer.com
*/

// see: https://pkg.go.dev/oras.land/oras-go/v2#example-package-PullFilesFromRemoteRepository

package ocifeatureinstall

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
	// 0. Create a file store
	fs, err := file.New("/tmp/oras/")
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
		panic(err)
	}

	tmpDirName := "/tmp/oras/"
	// iterrate over the files in the directory
	files, err := os.ReadDir(tmpDirName)
	if err != nil {
		return "", fmt.Errorf("failed to read directory: %w", err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
		fileName := file.Name()
		// check if fileName starts with "devcontainer-feature-"
		if fileName[:21] == "devcontainer-feature-" {
			// check if fileName ends with ".tgz"
			if fileName[len(fileName)-4:] == ".tgz" {
				fmt.Println("found devcontainer feature")
				err := untar(tmpDirName + fileName)
				fmt.Println("Err", err)
			}
		}
	}

	return manifestDescriptor.Digest.String(), nil

}
