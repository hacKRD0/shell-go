package main

import (
	"os"
	"strings"
)

func FindExecutable(cmd string) (string, bool) {
	path := os.Getenv("PATH")
	for _, dir := range strings.Split(path, ":") {
		filepath := dir + "/" + cmd
		_, err := os.Stat(filepath);
		if err == nil {
			return filepath, err == nil
		}
	}
	return "", false
}