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
	Pwd()
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

func (c *commands) Pwd() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wd)
}

func (c *commands) Exit() {
	os.Exit(0)
}

func (c *commands) Echo() {
	fmt.Println(strings.Join(c.argv[1:], " "))
}

func (c *commands) Type() {
	// Check if the argument is built-in
	k := strings.TrimSpace(c.argv[1])
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
	_, found := FindInPath(c.cmd)
	if found {
		executable := exec.Command(c.cmd, c.argv[1:]...)
		executable.Stdout = os.Stdout
		executable.Stderr = os.Stderr
		executable.Run()
	} else {
		fmt.Println(c.cmd + ": command not found")
	}
}
