/*
Copyright © 2024 devcontainer.com
*/
package s6manager

import (
	"fmt"
	"os"
	"path/filepath"
)

func AddService(serviceName string,
	dependencies []string,
	producerFor string,
	serviceType string,
	runScript string) error {
	return addServiceFiles(serviceName, dependencies, producerFor, serviceType, runScript)
}

func addServiceFiles(serviceName string,
	dependencies []string,
	producerFor string,
	serviceType string,
	runScript string) error {

	fmt.Println("add service", serviceName)

	baseDirS6Rc := "/tmp/etc/s6-overlay/s6-rc.d/"
	serviceDir := filepath.Join(baseDirS6Rc, serviceName)

	// Check if the directory already exists
	if _, err := os.Stat(serviceDir); err == nil {
		// If directory exists, return an error
		return fmt.Errorf("directory %s already exists", serviceDir)
	} else if !os.IsNotExist(err) {
		// If stat returns an error other than "not exists", return it
		return fmt.Errorf("error checking if directory exists: %w", err)
	}

	// Create the directory
	if err := os.MkdirAll(serviceDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", serviceDir, err)
	}

	// Manage dependencies
	if len(dependencies) > 0 {
		// Create the dependencies file
		dependenciesDirectory := filepath.Join(serviceDir, "dependencies.d")
		if err := os.MkdirAll(dependenciesDirectory, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dependenciesDirectory, err)
		}

		// Write the dependencies
		for _, dep := range dependencies {
			dependencyFile := filepath.Join(dependenciesDirectory, dep)
			if _, err := os.Create(dependencyFile); err != nil {
				return fmt.Errorf("failed to create dependency file %s: %w", dependencyFile, err)
			}
		}
	}

	// Manage producerFor
	if producerFor != "" {
		// Create the producerFor file
		producerForFile := filepath.Join(serviceDir, "producer-for")
		if _, err := os.Create(producerForFile); err != nil {
			return fmt.Errorf("failed to create producerFor file %s: %w", producerForFile, err)
		}

		// Write the producerFor
		if err := os.WriteFile(producerForFile, []byte(producerFor), 0644); err != nil {
			return fmt.Errorf("failed to write producerFor file %s: %w", producerForFile, err)
		}
	}

	// Manage serviceType
	serviceTypeFile := filepath.Join(serviceDir, "type")
	if _, err := os.Create(serviceTypeFile); err != nil {
		return fmt.Errorf("failed to create serviceType file %s: %w", serviceTypeFile, err)
	}
	if err := os.WriteFile(serviceTypeFile, []byte(serviceType), 0644); err != nil {
		return fmt.Errorf("failed to write serviceType file %s: %w", serviceTypeFile, err)
	}

	// Manage runScript
	runScriptFile := filepath.Join(serviceDir, "run")
	if _, err := os.Create(runScriptFile); err != nil {
		return fmt.Errorf("failed to create runScript file %s: %w", runScriptFile, err)
	}
	if err := os.WriteFile(runScriptFile, []byte(runScript), 0755); err != nil {
		return fmt.Errorf("failed to write runScript file %s: %w", runScriptFile, err)
	}

	// fmt.Println("add service", serviceName)
	// fmt.Println("dependencies", dependencies)
	// fmt.Println("producerFor", producerFor)
	// fmt.Println("serviceType", serviceType)
	// fmt.Println("runScript", runScript)

	return nil
}

func CreateLoggerFor(serviceName string) error {
	fmt.Println("create logger for", serviceName)
	return nil
}
