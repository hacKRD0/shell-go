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
		handleCommand(strings.Split(strings.TrimSpace(input), " "))
	}
}

func handleCommand(input []string) {
	// Parse user input
	cmd, argv := strings.TrimSpace(input[0]), input

	// Built-in commands map
	builtIns := map[string]int{
		"exit": 0,
		"echo": 1,
		"type": 2,
		"pwd": 3,
	}

	c := NewCommandsHandler(builtIns, cmd, argv)
	
	// Handle exit
	switch cmd {
		case "exit":
			c.Exit()
		case "echo": 
			c.Echo()
		case "type":
			c.Type()
		case "pwd":
			c.Pwd()
		default:
			c.Default()
	}
}