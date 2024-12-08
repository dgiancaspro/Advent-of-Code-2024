package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[!] Need a file name Please")
		os.Exit(2)
	}
	srcFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	data := bufio.NewScanner(srcFile)
	for data.Scan() {
		fmt.Println(data.Text())
	}
}
