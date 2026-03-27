package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds the dynamic configuration values from idp.yaml
type Config struct {
	AppName   string `yaml:"app_name"`
	AppPath   string `yaml:"app_path"`
	ImageName string `yaml:"image_name"`
	Tag       string `yaml:"tag"`
	Port      int    `yaml:"port"`
	Replicas  int    `yaml:"replicas"`
}

// loadConfig reads and parses the idp.yaml file and validates required fields
func loadConfig() (*Config, bool) {
	LogInfo("Loading configuration from idp.yaml...")

	// Read from project root
	data, err := os.ReadFile("../idp.yaml")
	if err != nil {
		LogError("Config file missing or cannot be read (make sure idp.yaml exists in the project root)", err, "")
		return nil, false
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		LogError("Failed to parse YAML configuration", err, "")
		return nil, false
	}

	// Validate required fields
	if config.AppName == "" || config.AppPath == "" || config.ImageName == "" || config.Tag == "" {
		LogError("Invalid config: app_name, app_path, image_name, and tag are required", nil, "")
		return nil, false
	}

	// Logging: Print loaded config clearly before execution
	LogSuccess("Configuration loaded successfully:")
	fmt.Printf("  - AppName:   %s\n", config.AppName)
	fmt.Printf("  - AppPath:   %s\n", config.AppPath)
	fmt.Printf("  - ImageName: %s\n", config.ImageName)
	fmt.Printf("  - Tag:       %s\n", config.Tag)
	fmt.Printf("  - Port:      %d\n", config.Port)
	fmt.Printf("  - Replicas:  %d\n", config.Replicas)

	return &config, true
}
