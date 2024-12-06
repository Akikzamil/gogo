package utils

import (
	"fmt"
	"os/exec"
)

func RunCommand(name string, arg ...string) bool {
	cmd := exec.Command(name,arg...)
	output, err := cmd.CombinedOutput();
	if err != nil {
		fmt.Printf("Error executing command: %s\n", err.Error())
        fmt.Printf("Command output: %s\n", string(output))
		return false
	}

	fmt.Println(string(output))

	return true

}