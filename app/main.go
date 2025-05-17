package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	for {
		cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		// Handle user input
		go handleCommand(cmd)
	}
}

func handleCommand(cmd string) {
	// Print user input
	fmt.Fprint(os.Stdout, cmd[:len(cmd)-1] + ": command not found\n")
}
