package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/malinoOS/malino/libmalino"
)

var Version string = "undefined"

func main() {
	fmt.Println("msh v" + Version)
	for {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("msh: could not get working directory: " + err.Error())
		}
		fmt.Printf("\033[91m%v #\033[39m ", currentDir)
		input := libmalino.UserLine()
		cmd := strings.Split(input, " ")
		switch cmd[0] {
		case "ping":
			fmt.Println("Pong!")
		case "pong":
			fmt.Println("Ping!")
		case "exit":
			os.Exit(0)
		case "":
			fmt.Println()
		case "cd":
			cd(cmd)
		default:
			entries, err := os.ReadDir("/bin")
			if err != nil {
				fmt.Printf("Failed to read contents of /bin: %v\n", err)
				os.Exit(1)
			}
			foundCmd := false
			for _, e := range entries {
				if !e.IsDir() {
					if cmd[0] == e.Name() {
						foundCmd = true
						if err := libmalino.SpawnProcess("/bin/"+e.Name(), currentDir, []string{"SHELL=/bin/msh", "USER=root"}, []uintptr{os.Stdout.Fd(), os.Stdin.Fd(), os.Stderr.Fd()}, true, true, cmd[1:]...); err != nil {
							fmt.Println("Error running /bin/" + e.Name() + ": " + err.Error())
						}
					}
				}
			}
			if !foundCmd {
				fmt.Println("")
			}
		}
	}
}

func cd(arrCommandStr []string) {
	dir := ""
	if arrCommandStr[1] == ".." {
		currentDir, _ := os.Getwd()
		dir = filepath.Dir(currentDir)
	} else if strings.HasPrefix(arrCommandStr[1], "/") {
		dir = arrCommandStr[1]
	} else {
		currentDir, _ := os.Getwd()
		dir = fmt.Sprintf("%v/%v", currentDir, arrCommandStr[1])
	}
	fileInfo, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("cd: could not change directory: %v doesn't exist\n", dir)
			return
		} else {
			fmt.Printf("cd: could not change directory: %v\n", err.Error())
			return
		}
	} else {
		if fileInfo.IsDir() {
			err := syscall.Chdir(dir)
			if err != nil {
				fmt.Printf("cd: could not change directory: %v\n", err.Error())
				return
			}
		} else {
			fmt.Printf("cd: could not change directory: %v isn't a directory\n", dir)
			return
		}
	}
}
