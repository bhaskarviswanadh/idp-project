package main

import (
	"fmt"
)

// loadImageToMinikube pushes the local Docker image into the Minikube environment
func loadImageToMinikube() bool {
	LogInfo("Loading image into Minikube...")
	
	// Requirement: minikube cache or load (using absolute path for Windows safety)
	// We'll use the precise syntax: exec.Command("cmd", "/C", "minikube image load idp-app:latest")
	// Since we know minikube path issues, we'll keep the absolute path from before
	output, err := runCommandWrapper("\"C:\\Program Files\\Kubernetes\\Minikube\\minikube.exe\"", "image", "load", "idp-app:latest")
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
func deployToKubernetes() bool {
	LogInfo("Deploying to Kubernetes...")
	
	// Requirement: kubectl apply -f ../k8s/deployment.yaml
	output, err := runCommandWrapper("kubectl", "apply", "-f", "../k8s/deployment.yaml")
	if err != nil {
		LogError("Failed to deploy to Kubernetes", err, output)
		return false
	}
	
	fmt.Print(output)
	LogSuccess("Deployment to Kubernetes successful")
	LogInfo("Access app using: minikube service idp-service")
	return true
}
