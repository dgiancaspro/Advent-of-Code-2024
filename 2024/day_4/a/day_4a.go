package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
)

func hSearch(line string) int {
	// fmt.Println(line)
	searchString := "XMAS"
	reSearch := regexp.MustCompile(searchString)
	lineMatch := reSearch.FindAllString(line, -1)
	return len(lineMatch)
}

func reverse(line string) string {
	revLine := ""
	for i := len(line) - 1; i >= 0; i-- {
		revLine += string(line[i])
	}
	// fmt.Println(line, revLine)
	return revLine
}

func main() {
	var dataSlice []string
	var answer int

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
	matrix := len(dataSlice[0])
	// var dataMatrix [matrix][matrix]string
	fmt.Println(dataSlice)
	fmt.Println(reflect.TypeOf(matrix), matrix)
	fmt.Printf("[+] Answer:%d\n", answer)
}

// for i := 0; i < len(dataSlice); i++ {
// 	answer += hSearch(dataSlice[i])
// 	revLine := reverse(dataSlice[i])
// 	answer += hSearch(revLine)
// } The Quick Brown fox jumped over the lazy dogs 76 +
