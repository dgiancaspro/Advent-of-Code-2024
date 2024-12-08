package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var dataList []string
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
		dataList = strings.Split(data.Text(), "\n")
	}
	fmt.Println(dataList)
}