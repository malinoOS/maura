package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file := ""
	if len(os.Args) <= 1 && os.Args[1] == "" {
		fmt.Println("cat: no file")
		os.Exit(2)
	}
	if strings.HasPrefix(os.Args[1], "/") {
		file = os.Args[1]
	} else {
		currentDir, _ := os.Getwd()
		if currentDir == "/" {
			file = fmt.Sprintf("/%v", os.Args[1])
		} else {
			file = fmt.Sprintf("%v/%v", currentDir, os.Args[1])
		}
	}
	dat, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("cat: could not read file: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(string(dat))
}
