package main

import (
	"fmt"
	"os"

	"github.com/malinoOS/malino/libmalino"
)

var Version string = "undefined"

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-v" || os.Args[1] == "--version" {
			fmt.Println("maura reboot v" + Version)
		}
	}
	libmalino.RebootComputer()
}
