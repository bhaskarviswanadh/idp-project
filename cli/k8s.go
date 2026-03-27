package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
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

// generateK8sYAML creates the deployment and service YAML dynamically
func generateK8sYAML(cfg *Config) string {
	yamlTemplate := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: %s
spec:
  replicas: %d
  selector:
    matchLabels:
      app: %s
  template:
    metadata:
      labels:
        app: %s
    spec:
      containers:
      - name: %s
        image: %s:%s
        imagePullPolicy: Never
        ports:
        - containerPort: %d
---
apiVersion: v1
kind: Service
metadata:
  name: %s-service
spec:
  type: NodePort
  selector:
    app: %s
  ports:
    - port: 80
      targetPort: %d
`
	return fmt.Sprintf(yamlTemplate,
		cfg.AppName, cfg.Replicas, cfg.AppName, cfg.AppName, cfg.AppName,
		cfg.ImageName, cfg.Tag, cfg.Port,
		cfg.AppName, cfg.AppName, cfg.Port,
	)
}

// deployToKubernetes runs the kubectl apply command using the deployment yaml
func deployToKubernetes(cfg *Config) bool {
	LogInfo("Generating dynamic Kubernetes manifests...")

	yamlStr := generateK8sYAML(cfg)
	fmt.Println("\n--- Generated YAML ---")
	fmt.Println(yamlStr)
	fmt.Println("----------------------")

	tempFile := "temp-deployment.yaml"
	err := os.WriteFile(tempFile, []byte(yamlStr), 0644)
	if err != nil {
		LogError("Failed to write temporary YAML file", err, "")
		return false
	}

	LogInfo("Deploying to Kubernetes...")

	// Use shell execution (cmd /C) to handle kubectl commands properly
	output, err := runCommandWrapper("kubectl apply -f " + tempFile)
	if err != nil {
		LogError("Failed to deploy to Kubernetes", err, output)
		LogInfo("Keeping " + tempFile + " for debugging.")
		return false
	}

	fmt.Print(output)
	LogSuccess("Deployment to Kubernetes successful")
	LogInfo("Access app using: minikube service " + cfg.AppName + "-service")

	// Delete temp file on success
	err = os.Remove(tempFile)
	if err != nil {
		LogInfo("Could not clean up temporary file: " + tempFile)
	}

	return true
}

// openServiceInBrowser opens the deployed application in the default web browser using Minikube
func openServiceInBrowser(cfg *Config) bool {
	LogInfo("Opening application in browser...")

	// minikube service --url blocks forever on Docker driver (it keeps the tunnel alive).
	// Solution: start it, read the first URL line from its output, kill the process, open browser.
	serviceName := fmt.Sprintf("%s-service", cfg.AppName)
	cmd := exec.Command("cmd", "/C", "minikube service "+serviceName+" --url")

	// Get a pipe to read stdout+stderr in real time
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		LogError("Failed to create output pipe", err, "")
		return false
	}
	cmd.Stderr = cmd.Stdout // redirect stderr to same pipe

	if err := cmd.Start(); err != nil {
		LogError("Failed to start minikube service command", err, "")
		return false
	}

	// Read lines until we find the URL, with a timeout fallback
	urlCh := make(chan string, 1)
	go func() {
		buf := make([]byte, 1024)
		var accumulated strings.Builder
		for {
			n, err := pipe.Read(buf)
			if n > 0 {
				accumulated.Write(buf[:n])
				// Check each line for a URL
				text := accumulated.String()
				for _, line := range strings.Split(text, "\n") {
					line = strings.TrimSpace(line)
					if strings.HasPrefix(line, "http") {
						urlCh <- line
						return
					}
				}
			}
			if err != nil {
				break
			}
		}
		urlCh <- ""
	}()

	// Wait up to 15 seconds for the URL
	var serviceURL string
	select {
	case serviceURL = <-urlCh:
	case <-time.After(15 * time.Second):
		LogError("Timed out waiting for minikube service URL", nil, "")
		cmd.Process.Kill()
		return false
	}

	// Kill the blocking minikube process now that we have the URL
	cmd.Process.Kill()

	if serviceURL == "" {
		LogError("Could not extract URL from minikube service output", nil, "")
		return false
	}

	LogInfo(fmt.Sprintf("Application URL: %s", serviceURL))

	// Open the browser — URL must be quoted separately for cmd start
	browserErr := exec.Command("cmd", "/C", "start", "", serviceURL).Run()
	if browserErr != nil {
		LogError("Failed to open browser", browserErr, "")
		return false
	}

	LogSuccess("Application opened successfully!")
	return true
}
