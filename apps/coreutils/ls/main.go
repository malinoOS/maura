package main

import (
	"fmt"
	"os"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("ls: could not list directory: %v\n", err)
		os.Exit(1)
	}
	entries, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Printf("ls: could not list directory: %v\n", err)
		os.Exit(1)
	}
	for _, e := range entries {
		if e.IsDir() {
			fmt.Print("\033[94m")
		} else {
			fmt.Print("\033[39m")
		}
		fmt.Printf("%v ", e.Name())
	}
	fmt.Print("\033[39m\n")
}
