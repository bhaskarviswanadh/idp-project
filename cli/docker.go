package main

import (
	"fmt"
)

// checkDocker checks if docker is installed and running
func checkDocker() bool {
	LogInfo("Checking Docker installation...")

	output, err := runCommandWrapper("docker --version")
	if err != nil {
		LogError("Docker is not installed or not running", err, output)
		return false
	}

	LogSuccess("Docker is installed")
	fmt.Print(output)
	return true
}

// buildDockerImage builds the Docker image from a specific directory context
func buildDockerImage(cfg *Config) bool {
	LogInfo("Starting Docker build...")
	LogInfo("Building Docker image " + cfg.ImageName + ":" + cfg.Tag + "...")

	// Requirement: docker build -t <image_name>:<tag> ../<app_path>
	dockerCmd := fmt.Sprintf("docker build -t %s:%s ../%s", cfg.ImageName, cfg.Tag, cfg.AppPath)
	output, err := runCommandWrapper(dockerCmd)
	if err != nil {
		LogError("Docker build failed", err, output)
		return false
	}

	LogSuccess("Docker image built successfully")
	fmt.Print(output)
	return true
}
