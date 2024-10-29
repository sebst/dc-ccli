/*
Copyright © 2024 devcontainer.com
*/
package customizer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ReadConfigFromUrl(url string, apiKey string) (*Config, error) {
	// Perform the HTTP GET request
	// with HttpHeader "Authorization: Bearer apiKey"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := http.Client{}
	resp, err := client.Do(req)

	// resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JSON from URL: %w", err)
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the JSON into the Config struct
	var config Config
	if err := json.Unmarshal(body, &config); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &config, nil
}
