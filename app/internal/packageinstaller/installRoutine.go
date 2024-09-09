package packageinstaller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"runtime"
)

func GetSystemArchitecture() string {
	arch := runtime.GOARCH
	switch arch {
	case "amd64":
		return "x86-64"
	case "arm64":
		return "arm64"
	default:
		return "unknown architecture"
	}
}

// Define the struct types
type VersionInfo struct {
	Sha256sum string `json:"sha256sum"`
	Size      int    `json:"size"`
}

type Architecture struct {
	Versions map[string]VersionInfo `json:"versions"`
}

type DistInfo struct {
	X86_64  Architecture `json:"x86-64"`
	Aarch64 Architecture `json:"aarch64"`
}

type PackageInfo struct {
	Dist DistInfo `json:"dist"`
}

type Packages map[string]PackageInfo

func fetchAndParseJSON() (Packages, error) {
	url := "https://raw.githubusercontent.com/sebst/brewkit-test/main/tools/_optim_db-minified.json"
	// Fetch the data from the URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a Packages object
	var packages Packages
	err = json.Unmarshal(body, &packages)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func InstallPackages(packages []*Package) error {
	arch := GetSystemArchitecture()

	tools, err := fetchAndParseJSON()
	if err != nil {
		fmt.Print(err)
	}

	for _, pkg := range packages {
		// Check if the package exists in the JSON data
		_, packageNameOk := tools[pkg.Name]
		packageVersionOk := false
		if pkg.Version != "latest" {
			packageVersionOk = false
			if arch == "x86-64" {
				_, packageVersionOk = tools[pkg.Name].Dist.X86_64.Versions[pkg.Version]
			} else if arch == "arm64" {
				_, packageVersionOk = tools[pkg.Name].Dist.Aarch64.Versions[pkg.Version]
			}
		} else {
			packageVersionOk = true
		}
		if packageNameOk && packageVersionOk {
			fmt.Println("Installing package:", pkg.Name, "version:", pkg.Version)
		} else {
			fmt.Println("Package not found:", pkg.Name, "version:", pkg.Version)
			return errors.New("package not found")
		}
	}
	arguments := []string{"install"}
	for _, pkg := range packages {
		if pkg.Version == "latest" {
			arguments = append(arguments, pkg.Name)
		} else {
			arguments = append(arguments, pkg.Name+"="+pkg.Version)
		}
	}
	fmt.Println("Running command:", "pkgx", arguments)

	cmd := exec.Command("pkgx", arguments...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return errors.New("command failed")
	}

	fmt.Println(string(stdout))
	return nil
}
