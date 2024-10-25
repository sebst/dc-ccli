/*
Copyright © 2024 devcontainer.com
*/
package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

// ServiceConfig represents a service to run, based on the JSON configuration
type ServiceConfig struct {
	Name      string   `json:"name"`
	Run       string   `json:"run"`
	Prefix    string   `json:"prefix"`
	DependsOn []string `json:"dependsOn"`
}

// RunningService holds the command and configuration for each running service
type RunningService struct {
	Config ServiceConfig
	Cmd    *exec.Cmd
}

// PIDInfo represents the response structure for the /api/pids endpoint
type PIDInfo struct {
	PID    int           `json:"pid"`
	Config ServiceConfig `json:"config"`
}

var (
	// Map to keep track of running services by PID
	runningServices = make(map[int]*RunningService)
	// Mutex to protect concurrent access to runningServices
	mu sync.Mutex
	// Flag to indicate if the program is shutting down
	shuttingDown = false
)

func ServicesRunner(jsonFile string) {
	// Load the configuration from services.json
	services, err := loadConfig(jsonFile)
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		return
	}

	// Start each service specified in the configuration
	for _, service := range services {
		startService(service)
	}

	// Start the web server in a separate goroutine
	go startWebServer()

	// Set up a channel to capture interrupt signals (Ctrl+C)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT)

	// Wait for a signal (Ctrl+C)
	<-sigchan

	// Set the shutting down flag
	shuttingDown = true

	// Forward the SIGINT signal to all running commands
	mu.Lock()
	for pid, service := range runningServices {
		err = service.Cmd.Process.Signal(syscall.SIGINT)
		if err != nil {
			fmt.Printf("Failed to send signal to process PID %d: %v\n", pid, err)
		} else {
			fmt.Printf("Signal forwarded to PID %d\n", pid)
		}
	}
	mu.Unlock()

	// Wait for all processes to exit gracefully
	mu.Lock()
	for _, service := range runningServices {
		service.Cmd.Wait()
	}
	mu.Unlock()
}

// loadConfig loads the service configuration from a JSON file
func loadConfig(filePath string) ([]ServiceConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var services []ServiceConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&services); err != nil {
		return nil, err
	}

	return services, nil
}

// startService starts the external process specified in the service configuration
func startService(service ServiceConfig) {
	// Start the external process
	cmd := exec.Command(service.Run)

	// Create pipes to capture stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Failed to get stdout pipe for %s: %v\n", service.Name, err)
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Failed to get stderr pipe for %s: %v\n", service.Name, err)
		return
	}

	// Start the process in the background
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Failed to start %s: %v\n", service.Name, err)
		return
	}

	// Attach the process to the runningServices map
	mu.Lock()
	runningServices[cmd.Process.Pid] = &RunningService{
		Config: service,
		Cmd:    cmd,
	}
	mu.Unlock()

	// Print the process ID
	fmt.Printf("%s started as PID %d\n", service.Name, cmd.Process.Pid)

	// Start goroutines to capture and prefix the output from stdout and stderr
	go prefixOutput(service.Prefix, stdout)
	go prefixOutput(service.Prefix, stderr)

	// Monitor the process and restart it if it terminates (unless shutting down)
	go func() {
		// Wait for the process to exit
		cmd.Wait()

		if !shuttingDown { // Check if we're not shutting down
			fmt.Printf("%s (PID %d) terminated, restarting...\n", service.Name, cmd.Process.Pid)

			// Remove the terminated process from the runningServices map
			mu.Lock()
			delete(runningServices, cmd.Process.Pid)
			mu.Unlock()

			// Restart the service
			startService(service)
		} else {
			// Remove the terminated process from the runningServices map
			mu.Lock()
			delete(runningServices, cmd.Process.Pid)
			mu.Unlock()
			fmt.Printf("%s (PID %d) terminated and will not restart due to shutdown\n", service.Name, cmd.Process.Pid)
		}
	}()
}

// prefixOutput reads output line by line from the reader, and prefixes it with the specified prefix
func prefixOutput(prefix string, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("%s %s\n", prefix, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading output: %v\n", err)
	}
}

// startWebServer starts an HTTP server that handles various API endpoints
func startWebServer() {
	// List all running PIDs with their service configurations
	http.HandleFunc("/api/pids", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		var pids []PIDInfo
		for pid, service := range runningServices {
			pids = append(pids, PIDInfo{
				PID:    pid,
				Config: service.Config,
			})
		}

		// Send the list of PIDs and their configurations as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pids)
	})

	// Restart a process by its PID
	http.HandleFunc("/api/pids/", func(w http.ResponseWriter, r *http.Request) {
		pidStr := strings.TrimPrefix(r.URL.Path, "/api/pids/")
		pid, err := strconv.Atoi(pidStr)
		if err != nil || pid == 0 {
			http.Error(w, "Invalid PID", http.StatusBadRequest)
			return
		}

		// Send SIGINT to the process and restart it
		mu.Lock()
		service, exists := runningServices[pid]
		mu.Unlock()
		if !exists {
			http.Error(w, "Process not found", http.StatusNotFound)
			return
		}

		err = service.Cmd.Process.Signal(syscall.SIGINT)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send SIGINT: %v", err), http.StatusInternalServerError)
			return
		}

		// Wait for the process to shut down
		go func() {
			service.Cmd.Wait()

			// Restart the process
			startService(service.Config)
		}()

		fmt.Fprintf(w, "Process %d is restarting\n", pid)
	})

	// Start listening on port 8080
	fmt.Println("Starting web server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start web server: %v\n", err)
	}
}
