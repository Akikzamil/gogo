package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetModuleName() (string, error) {
	execDir, _ := os.Getwd()

	location:=filepath.Join(execDir,"go.mod")
	file, err := os.Open(location)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			// Extract the module name
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return "", fmt.Errorf("module name not found in %s", location)
}