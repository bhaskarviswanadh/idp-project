package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if a command is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: idp <command>")
		os.Exit(1)
	}

	// Read the first argument passed to the CLI
	command := os.Args[1]

	// Handle standard commands
	switch command {
	case "deploy":
		deploy()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

// deploy orchestrates standard deployment workflow
func deploy() {
	LogInfo("Starting deployment pipeline...")

	cfg, ok := loadConfig()
	if !ok {
		return
	}

	// 1. Check if Docker is installed
	if !checkDocker() {
		return
	}

	// 2. Build Docker Image
	if !buildDockerImage(cfg) {
		return
	}

	// 3. Load image into Minikube
	if !loadImageToMinikube(cfg) {
		return
	}

	// 4. Deploy to Kubernetes
	if !deployToKubernetes(cfg) {
		return
	}

	LogSuccess("Deployment pipeline completed successfully!")
}
