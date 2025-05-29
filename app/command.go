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
	Cd()
	Cat()
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

func (c *commands) Cat() {
	if len(c.argv) < 2 {
		fmt.Println("cat: missing file operand")
		return
	}
	for _, file := range c.argv[1:] {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("cat: %s: %v\n", file, err)
			continue
		}
		defer f.Close()
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("cat: %s: %v\n", file, err)
			continue
		}
		fmt.Println(string(content))
	}
}

func (c *commands) Cd() {
	if len(c.argv) > 2 {
		fmt.Println("bash: cd: too many arguments")
		return
	}

 	dir := os.Getenv("HOME")
	if len(c.argv) == 2 && c.argv[1] != "~" {
		dir = strings.TrimSpace(c.argv[1])
	}

	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("cd: " + dir + ": No such file or directory")
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
	for _, arg := range c.argv[1:] {
		// Check if the argument is built-in
		k := strings.TrimSpace(arg)
		_, ok := c.builtIns[k]
		if ok {
			fmt.Println(k + " is a shell builtin")
			continue 
		}
	
		// Check if the argument is in a directory defined in the path variable 
		path, found := FindExecutable(k)
		if found {
			fmt.Println(k, "is " + path)	
			continue
		} 
		fmt.Println(k + ": not found")
	}
}


func (c * commands) Default() {
	_, found := FindExecutable(c.cmd)
	if found {
		executable := exec.Command(c.cmd, c.argv[1:]...)
		executable.Stdout = os.Stdout
		executable.Stderr = os.Stderr
		executable.Run()
	} else {
		fmt.Println(c.cmd + ": command not found")
	}
}
