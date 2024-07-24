package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [args...]")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	for {
		cmd := exec.Command(command, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Command '%s' failed with error: %s\n", command, err)
			fmt.Println("Restarting command...")
			time.Sleep(1 * time.Second) // Sleep for a second before restarting
		} else {
			fmt.Printf("Command '%s' completed successfully\n", command)
			break
		}
	}
}