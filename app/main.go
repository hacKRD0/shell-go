package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		// Handle errors
		if err != nil {
			continue
		}

		// Handle user input
		_, err = handleCommand(strings.Split(input, " "))

		// Handle errors
		if err != nil {
			continue
		}
	}
}

func handleCommand(input []string) ([]byte, error) {
	// Parse user input
	cmd, _ := strings.TrimSpace(input[0]), input[1:]
	
	// Handle exit
	if cmd == "exit\n" {
		os.Exit(0)
	}

	// Print user input
	fmt.Fprint(os.Stdout, cmd + ": command not found\n")

	return nil, nil
}
