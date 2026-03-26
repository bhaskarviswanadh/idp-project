package main

import (
	"fmt"
	"os"
	"os/exec"
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

func deploy() {
	fmt.Println("🚀 Starting deployment pipeline...")
	fmt.Println("🔍 Checking Docker installation...")
	
	cmd := exec.Command("docker", "--version")
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		fmt.Println("❌ Error running Docker command:")
		fmt.Println(err)
		fmt.Print(string(output))
		return
	}
	
	fmt.Println("✅ Docker is installed")
	fmt.Print(string(output))

	fmt.Println("🚀 Starting Docker build...")
	fmt.Println("📦 Building Docker image...")

	buildCmd := exec.Command("docker", "build", "-t", "idp-app", ".")
	buildOutput, buildErr := buildCmd.CombinedOutput()

	if buildErr != nil {
		fmt.Println("❌ Docker build failed")
		fmt.Println(buildErr)
		fmt.Print(string(buildOutput))
		return
	}

	fmt.Println("✅ Docker image built successfully")
	fmt.Print(string(buildOutput))

	fmt.Println("🧹 Cleaning up old containers...")
	_ = exec.Command("docker", "rm", "-f", "idp-container").Run()
	fmt.Println("✨ Previous container removed (if any)")

	fmt.Println("🚀 Running Docker container...")
	runCmd := exec.Command("docker", "run", "-d", "-p", "5000:5000", "--name", "idp-container", "idp-app")
	runOutput, runErr := runCmd.CombinedOutput()

	if runErr != nil {
		fmt.Println("❌ Failed to run container")
		fmt.Println(runErr)
		fmt.Print(string(runOutput))
		return
	}

	fmt.Println("✅ Container started successfully")
	fmt.Println("🌐 App running at: http://localhost:5000")
	fmt.Print(string(runOutput))
}
