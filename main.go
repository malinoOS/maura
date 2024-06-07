package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/malinoOS/malino/libmalino"
)

var Version string = "undefined"

func main() {
	defer libmalino.ResetTerminalMode()
	fmt.Println("\033[97mWelcome to maura v" + Version + "!")
	fmt.Print("Mounting /proc... ")
	if err := os.Mkdir("/proc", 0777); err != nil {
		mauraPanic(err, "creating /proc")
	}
	if err := syscall.Mount("proc", "/proc", "proc", uintptr(0), ""); err != nil {
		mauraPanic(err, "mounting /proc")
	} else {
		fmt.Println("\033[92m[OK]\033[39m")
	}
	fmt.Println("")
	fmt.Println("Starting /bin/msh (maura shell)...")

	if err := libmalino.SpawnProcess("/bin/msh", "/", []string{"SHELL=/bin/msh", "USER=root"}, []uintptr{os.Stdout.Fd(), os.Stdin.Fd(), os.Stderr.Fd()}, true, true); err != nil {
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
	} // todo: don't spin the cpu
}
