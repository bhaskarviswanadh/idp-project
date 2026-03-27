package main

import (
	"fmt"
)

// checkDocker checks if docker is installed and running
func checkDocker() bool {
	LogInfo("Checking Docker installation...")
	
	output, err := runCommandWrapper("docker", "--version")
	if err != nil {
		LogError("Docker is not installed or not running", err, output)
		return false
	}
	
	LogSuccess("Docker is installed")
	fmt.Print(output)
	return true
}

// buildDockerImage builds the Docker image from a specific directory context
func buildDockerImage() bool {
	LogInfo("Starting Docker build...")
	LogInfo("Building Docker image idp-app...")
	
	// Requirement: docker build -t idp-app ./sample-app
	output, err := runCommandWrapper("docker", "build", "-t", "idp-app", "./sample-app")
	if err != nil {
		LogError("Docker build failed", err, output)
		return false
	}
	
	LogSuccess("Docker image built successfully")
	fmt.Print(output)
	return true
}
