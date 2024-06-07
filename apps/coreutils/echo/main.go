package main

import (
	"fmt"
	"os"
	"strings"
)

var Version string = "undefined"

func main() {
	// check if there is any args
	if len(os.Args) > 1 {
		if os.Args[1] == "-n" {
			fmt.Print(strings.Join(os.Args[2:], " "))
		} else if os.Args[1] == "-v" || os.Args[1] == "--version" {
			fmt.Println("maura echo v" + Version)
		} else {
			fmt.Println(strings.Join(os.Args[1:], " "))
		}

	} else {
		fmt.Println()
	}
}
