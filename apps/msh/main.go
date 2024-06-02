package main

import (
	"fmt"

	"github.com/malinoOS/malino/libmalino"
)

var Version string = "undefined"

func main() {
	fmt.Println("msh v" + Version)
	for {
		fmt.Print("\033[91m#\033[39m ")
		input := libmalino.UserLine()
		switch input {
		case "ping":
			fmt.Println("Pong!")
		case "pong":
			fmt.Println("Ping!")
		case "stop":
			libmalino.ShutdownComputer()
		case "":
			fmt.Println()
		default:
			fmt.Println("What?")
		}
	}
}
