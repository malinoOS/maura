package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var Version string = "undefined"

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-v" || os.Args[1] == "--version" {
			fmt.Println("maura ps v" + Version)
		}
	}
	matches, err := filepath.Glob("/proc/*/exe")
	if err != nil {
		fmt.Printf("ps: could not list processes: %v\n", err)
		os.Exit(1)
	}
	for _, file := range matches {
		pid := filepath.Base(filepath.Dir(file))
		target, err := os.Readlink(file)
		if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
			fmt.Printf("ps: could not list processes: %v\n", err)
			//os.Exit(1)
		}
		if len(target) > 0 && !strings.Contains(pid, "self") {
			fmt.Printf("PID: %s, Process: %+v\n", pid, target)
		}
	}
}
