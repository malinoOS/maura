package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("If you don't see anything, that's because /proc hasn't been fully implemented yet in malino.")
	fmt.Println("Trust me, there is code for this command, but it doesn't work right now.")
	matches, err := filepath.Glob("/proc/*/exe")
	if err != nil {
		fmt.Printf("ps: could not list processes: %v\n", err)
		os.Exit(1)
	}
	for _, file := range matches {
		pid := filepath.Base(filepath.Dir(file))
		target, err := os.Readlink(file)
		if err != nil {
			fmt.Printf("ps: could not list processes: %v\n", err)
			os.Exit(1)
		}
		if len(target) > 0 {
			fmt.Printf("PID: %s, Process: %+v\n", pid, target)
		}
	}
}
