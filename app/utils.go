package main

import (
	"os"
	"path/filepath"
	"strings"
)

func FindInPath(cmd string) (string, bool) {
	path := os.Getenv("PATH")
	for _, dir := range strings.Split(path, ";") {
		filepath := filepath.Join(dir, cmd)
		_, err := os.Stat(filepath);
		if err == nil {
			return filepath, err == nil
		}
	}
	return cmd, false
}