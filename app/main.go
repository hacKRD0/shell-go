package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')

		// Handle errors
		if err != nil {
			continue
		}

		// Handle user input
		_, err = handleCommand(cmd)

		// Handle errors
		if err != nil {
			continue
		}
	}
}

func handleCommand(cmd string) ([]byte, error) {
	// Print user input
	fmt.Fprint(os.Stdout, cmd[:len(cmd)-1] + ": command not found\n")

	return nil, nil
}
