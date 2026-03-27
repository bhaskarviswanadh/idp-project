package main

import (
	"fmt"
)

// loadImageToMinikube pushes the local Docker image into the Minikube environment
func loadImageToMinikube(cfg *Config) bool {
	LogInfo("Loading image into Minikube...")
	
	// Use shell execution (cmd /C) to handle minikube commands properly
	loadCmd := fmt.Sprintf("minikube image load %s:%s", cfg.ImageName, cfg.Tag)
	output, err := runCommandWrapper(loadCmd)
	if err != nil {
		LogError("Failed to load image into Minikube", err, output)
		return false
	}
	
	if len(output) > 0 {
		fmt.Print(output)
	}
	return true
}

// deployToKubernetes runs the kubectl apply command using the deployment yaml
func deployToKubernetes(cfg *Config) bool {
	LogInfo("Deploying to Kubernetes...")
	
	// Use shell execution (cmd /C) to handle kubectl commands properly
	output, err := runCommandWrapper("kubectl apply -f ../k8s/deployment.yaml")
	if err != nil {
		LogError("Failed to deploy to Kubernetes", err, output)
		return false
	}
	
	fmt.Print(output)
	LogSuccess("Deployment to Kubernetes successful")
	LogInfo("Access app using: minikube service idp-service")
	return true
}
