/*
Copyright © 2024 devcontainer.com
*/
package ocifeatureinstall

import (
	"encoding/json"
	"os"
)

// DevcontainerFeature represents the structure of a devcontainer-feature.json file.
type DevcontainerFeature struct {
	ID               string                   `json:"id"`
	Version          string                   `json:"version"`
	Name             string                   `json:"name"`
	Description      string                   `json:"description"`
	Options          map[string]FeatureOption `json:"options,omitempty"`
	InstallsAfter    []string                 `json:"installsAfter,omitempty"`
	Privileged       bool                     `json:"privileged,omitempty"`
	Init             bool                     `json:"init,omitempty"`
	DocumentationURL string                   `json:"documentationURL,omitempty"`
	InstallCommand   string                   `json:"installCommand,omitempty"`
	EntryPoints      []string                 `json:"entryPoints,omitempty"`
}

// FeatureOption defines an option that can be configured in a devcontainer feature.
type FeatureOption struct {
	Type        string        `json:"type"`
	Description string        `json:"description"`
	Default     interface{}   `json:"default,omitempty"`
	Enum        []interface{} `json:"enum,omitempty"`
}

func parseDevcontainerFeatureJSONFile(fileName string) (*DevcontainerFeature, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var feature DevcontainerFeature
	err = json.NewDecoder(f).Decode(&feature)
	if err != nil {
		return nil, err
	}

	return &feature, nil
}
