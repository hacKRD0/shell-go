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
		_, err = handleCommand(strings.Split(strings.TrimSpace(input), " "))

		// Handle errors
		if err != nil {
			continue
		}
	}
}

func handleCommand(input []string) ([]byte, error) {
	// Parse user input
	cmd, args := strings.TrimSpace(input[0]), input[1:]

	// Built-in commands map
	builtins := map[string]int{
		"exit": 0,
		"echo": 1,
		"type": 2,
	}
	
	// Handle exit
	switch cmd {
		case "exit":
			os.Exit(0)
		case "echo": 
			fmt.Println(strings.Join(args, " "))
		case "type":
			// Check if the command is built-in
			k := strings.TrimSpace(args[0])
			_, ok := builtins[k]
			if ok {
				fmt.Println(k + " is a shell builtin")
			} else {
				fmt.Print(k, ": not found\n")
			} 
		default:
			fmt.Print(cmd, ": command not found\n")
	}

	return nil, nil
}
