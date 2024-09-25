/*
Copyright © 2024 devcontainer.com
*/
package dcinstaller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type PackageList struct {
	Packages []Package `json:"packages"`
}

func InstallFromConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()
	data, err := io.ReadAll(io.Reader(file))
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return err
	}
	var pkgList PackageList
	err = json.Unmarshal(data, &pkgList)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
		return err
	}

	outFile := "/tmp/dc-ccli.packages.txt"
	out, err := os.Create(outFile)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
		return err
	}

	for _, pkg := range pkgList.Packages {
		fmt.Printf("Package Name: %s, Version: %s\n", pkg.Name, pkg.Version)
		_, err = out.WriteString(fmt.Sprintf("%s\t%s\n", pkg.Name, pkg.Version))
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}
	defer out.Close()
	return nil
}
