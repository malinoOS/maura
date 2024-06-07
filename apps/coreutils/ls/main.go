package main

import (
	"fmt"
	"os"
	"strings"
)

var Version string = "undefined"

func main() {
	dir := "/"
	if len(os.Args) > 1 && os.Args[1] != "" {
		if strings.HasPrefix(os.Args[1], "/") {
			dir = os.Args[1]
		} else if os.Args[1] == "-v" || os.Args[1] == "--version" {
			fmt.Println("maura ls v" + Version)
		} else {
			var err error
			dir, err = os.Getwd()
			if err != nil {
				fmt.Printf("ls: could not find current path: %v\n", err)
				os.Exit(1)
			}
		}
	} else {
		var err error
		dir, err = os.Getwd()
		if err != nil {
			fmt.Printf("ls: could not find current path: %v\n", err)
			os.Exit(1)
		}
	}
	entries, err := os.ReadDir(dir)
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
