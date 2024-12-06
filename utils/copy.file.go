package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CopyFile(cliFilePath string, execFiltPath string,moduleName string, replaceFrom string,replaceTo string) bool {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return false
	}

	// Get the directory of the executable
	cliDir := filepath.Dir(execPath)
	sourceFile := filepath.Join(cliDir, cliFilePath)

	// Get the current execution directory
	execDir, err := os.Getwd()

	if err != nil {
		return false
	}

	// Construct the destination file path in the execution directory
	destFile := filepath.Join(execDir, execFiltPath)

	content, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return false
	}
	originalText := string(content)

	modifiedText := strings.ReplaceAll(originalText, "gogo", moduleName);
	
	if(replaceFrom != ""){
		modifiedText = strings.ReplaceAll(modifiedText, replaceFrom, replaceTo);
	}

	dir := filepath.Dir(destFile)
	err2:= os.MkdirAll(dir, 0755)
	if err2 != nil {
		return false
	}

	err = os.WriteFile(destFile, []byte(modifiedText), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return false
	}

	fmt.Println("File created")
	return true
}