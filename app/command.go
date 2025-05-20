package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Commands interface {
	Exit()
	Echo()
	Type()
	Default()
}

type commands struct {
	builtIns map[string]int
	argv []string
	cmd string
}

func NewCommandsHandler(builtIns map[string]int ,cmd string, argv []string) *commands {
	return &commands{
		builtIns: builtIns,
		argv: argv,
		cmd: cmd,
	}
}

func (c *commands) Exit() {
	os.Exit(0)
}

func (c *commands) Echo() {
	fmt.Println(strings.Join(c.argv, " "))
}

func (c *commands) Type() {
	// Check if the argument is built-in
	k := strings.TrimSpace(c.argv[0])
	_, ok := c.builtIns[k]
	if ok {
		fmt.Println(k + " is a shell builtin")
		return 
	}

	// Check if the argument is in a directory defined in the path variable 
	path, found := FindInPath(k)
	if found {
		fmt.Println(k, "is " + path)	
		return
	} 
	fmt.Println(k + ": not found")
}


func (c * commands) Default() {
	path, found := FindInPath(c.cmd)
	if found {
		executable := exec.Command(path, c.argv...)
		fmt.Printf("Program was passed %d args (including program name).\n", len(c.argv) + 1)
		_, _ = executable.Output()
	} else {
		fmt.Println(c.cmd + ": command not found")
	}
}
