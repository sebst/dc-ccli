/*
Copyright © 2024 devcontainer.com
*/

package ocifeatureinstall

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"os"

	oras "oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/file"
	"oras.land/oras-go/v2/registry"
	"oras.land/oras-go/v2/registry/remote"
)

func untar(fileName string) {
	// Open the tar file for reading.
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// untar
	// Create a new tar reader.
	reader := tar.NewReader(file)
	// Iterate through the files in the archive.
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		// Print the name of the file.
		fmt.Println(header.Name)
		// Create a new file.
		newFile, err := os.Create(header.Name)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer newFile.Close()
		// Copy the file contents.
		_, err = io.Copy(newFile, reader)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
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
				untar(tmpDirName + fileName)
			}
		}
	}

	return manifestDescriptor.Digest.String(), nil

}
