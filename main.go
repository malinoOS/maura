package main

import (
	"fmt"
	"time"

	"github.com/malinoOS/malino/libmalino"
)

var Version string = "undefined"

func main() {
	defer libmalino.ResetTerminalMode()
	fmt.Println("\033[97mWelcome to maura v" + Version + "!")

	fmt.Print("Mounting /proc... ")
	if err := libmalino.MountProcFS(); err != nil {
		mauraPanic(err, "mounting /proc")
	} else {
		fmt.Println("\033[92m[OK]\033[39m")
	}

	fmt.Print("Mounting /dev... ")
	if err := libmalino.MountDevFS(); err != nil {
		mauraPanic(err, "mounting /dev")
	} else {
		fmt.Println("\033[92m[OK]\033[39m")
	}

	fmt.Println("Starting /bin/msh (maura shell)...")

	if _, err := libmalino.SpawnProcess("/bin/msh", "/", []string{"SHELL=/bin/msh", "USER=root"}, true, ""); err != nil {
		mauraPanic(err, "running /bin/msh")
	}
	libmalino.ShutdownComputer()
}

func mauraPanic(err error, where string) {
	fmt.Println("\n--- maura \033[91mPANIC!\033[39m ---")
	fmt.Println(err.Error())
	fmt.Println("This happened while " + where)
	fmt.Println("\nThe system is halted.")
	for {
		time.Sleep(time.Hour)
	}
}
