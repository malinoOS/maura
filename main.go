package main

import (
	"fmt"
	"os"
	"time"

	"github.com/malinoOS/malino/libmalino"
)

var Version string = "undefined"

func main() {
	defer libmalino.ResetTerminalMode()
	fmt.Println("Welcome to maura v" + Version + "!")
	fmt.Println("Starting /bin/msh (maura shell)...")

	if err := libmalino.SpawnProcess("/bin/msh", "/", []string{"SHELL=/bin/msh", "USER=root"}, []uintptr{os.Stdout.Fd(), os.Stdin.Fd(), os.Stderr.Fd()}, true, true); err != nil {
		fmt.Println("Error running /bin/msh: " + err.Error())
		fmt.Println("The system will shut down in 15 seconds.")
		time.Sleep(15 * time.Second)
	}
	libmalino.ShutdownComputer()
}
