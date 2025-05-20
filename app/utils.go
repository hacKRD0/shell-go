package main

import (
	"fmt"
	"os"
	"strings"
)

func FindInPath(cmd string) (string, bool) {
	path := os.Getenv("PATH")
	fmt.Println(path)
	for _, dir := range strings.Split(path, ";") {
		filepath := dir + "/" + cmd
		_, err := os.Stat(filepath);
		if err == nil {
			return filepath, err == nil
		}
	}
	return "", false
}