package main

import (
	"fmt"
	"os/exec"
)

// LogInfo prints a structured info message
func LogInfo(msg string) {
	fmt.Printf("[INFO] 🚀 %s\n", msg)
}

// LogSuccess prints a structured success message
func LogSuccess(msg string) {
	fmt.Printf("[SUCCESS] ✅ %s\n", msg)
}

// LogError prints a structured error message alongside the raw error and command output
func LogError(msg string, err error, output string) {
	fmt.Printf("[ERROR] ❌ %s\n", msg)
	if err != nil {
		fmt.Printf("Details: %s\n", err.Error())
	}
	if output != "" {
		fmt.Printf("Command Output:\n%s\n", output)
	}
}

// runCommandWrapper executes command using the Windows "cmd /C"
func runCommandWrapper(command string, args ...string) (string, error) {
	fullArgs := append([]string{"/C", command}, args...)
	cmd := exec.Command("cmd", fullArgs...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
