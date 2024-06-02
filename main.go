package main

import (
	"fmt"

	"github.com/malinoOS/malino/libmalino"
)

func main() {
	defer libmalino.ResetTerminalMode()
	fmt.Println("Welcome to maura! The official tech-demo OS for malino!")
	for { // Word of advice: Never let this app exit. Always end in an infinite loop or shutdown.
		fmt.Print("\033[91m#\033[39m ")
		input := libmalino.UserLine()
		switch input {
		case "ping":
			fmt.Println("Pong!")
		case "pong":
			fmt.Println("Ping!")
		case "stop":
			libmalino.ShutdownComputer()
		default:
			fmt.Println("What?")
		}
	}
}
