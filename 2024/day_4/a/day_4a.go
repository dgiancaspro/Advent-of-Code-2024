package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var dataSlice []string
	var searchStringRev []string
	searchString := "XMAS"

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
		dataSlice = append(dataSlice, data.Text())
	}
	fmt.Println(dataSlice)
	fmt.Println(searchString)
	for i := len(searchString); i >= 0; i-- {
		append(searchStringRev, searchString[i])
	}
	fmt.Println(searchStringRev)
}
