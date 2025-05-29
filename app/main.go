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
		fullCommand, err := bufio.NewReader(os.Stdin).ReadString('\n')
		
		// Handle errors
		if err != nil {
			continue
		}
		
		// Handle user input
		handleCommand(strings.TrimSpace(fullCommand))
	}
}

func handleCommand(fullCommand string) {
	// Parse the command line input
	argv, err := NewParser(fullCommand).ParseTokens()
	if err != nil {
		fmt.Println("Error parsing command:", err)
		return
	}
	cmd := argv[0]

	// Built-in commands map
	builtIns := map[string]int{
		"exit": 0,
		"echo": 1,
		"type": 2,
		"pwd": 3,
		"cd": 4,
		"cat": 5,
	}

	c := NewCommandsHandler(builtIns, cmd, argv)
	
	// Handle exit
	switch c.cmd {
		case "exit":
			c.Exit()
		case "echo": 
			c.Echo()
		case "type":
			c.Type()
		case "pwd":
			c.Pwd()
		case "cd":
			c.Cd()
		case "cat":
			c.Cat()
		default:
			c.Default()
	}
}