package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ExecuteCommand runs the provided shell command.
func ExecuteCommand(command string) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command) // Use "bash" for Unix-like systems
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
	}
}

// HandleUserOptions presents the options to the user.
func HandleUserOptions(command string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Execute the command")
		fmt.Println("2. Revise the command")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			ExecuteCommand(command)
		case "2":
			fmt.Printf("Enter the revised command [%s]: ", command)
			revisedCommand, _ := reader.ReadString('\n')
			revisedCommand = strings.TrimSpace(revisedCommand)
			if revisedCommand == "" {
				revisedCommand = command
			}
			ExecuteCommand(revisedCommand)
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}