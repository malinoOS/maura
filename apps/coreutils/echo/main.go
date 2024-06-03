package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// check if there is any args
	if len(os.Args) > 1 {
		if os.Args[1] == "-n" {
			fmt.Print(strings.Join(os.Args[2:], " "))
		} else {
			fmt.Println(strings.Join(os.Args[1:], " "))
		}

	} else {
		fmt.Println()
	}
}
