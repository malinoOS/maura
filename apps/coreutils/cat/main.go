package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("cat: no file")
		os.Exit(2)
	}

	var file string

	if os.Args[1] == "-s" {
		if len(os.Args) <= 2 {
			fmt.Println("cat: no file")
			os.Exit(2)
		}
		file = os.Args[2]
		readFileByteByByte(file)
	} else {
		file = os.Args[1]
		readFileInOneGo(file)
	}
}

func getFilePath(fileName string) string {
	if strings.HasPrefix(fileName, "/") {
		return fileName
	}
	currentDir, _ := os.Getwd()
	if currentDir == "/" {
		return fmt.Sprintf("/%v", fileName)
	}
	return fmt.Sprintf("%v/%v", currentDir, fileName)
}

func readFileInOneGo(file string) {
	filePath := getFilePath(file)
	dat, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("cat: could not read file: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(string(dat))
}

func readFileByteByByte(file string) {
	filePath := getFilePath(file)
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("cat: could not open file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	buf := make([]byte, 1)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == os.ErrClosed || n == 0 {
				break
			}
			fmt.Printf("cat: could not read file: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(string(buf[:n]))
	}
}
