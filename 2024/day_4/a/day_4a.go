/*
Notes
A strings Characters can be accessed indivdually but they are returned as bytes.  You need to convert
the	back to string characters with the string function ... string(line[index])
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	//"reflect"
	//"regexp"
)

// func hSearch(line string) int {
// 	// fmt.Println(line)
// 	searchString := "XMAS"
// 	reSearch := regexp.MustCompile(searchString)
// 	lineMatch := reSearch.FindAllString(line, -1)
// 	return len(lineMatch)
// }

// func reverse(line string) string {
// 	revLine := ""
// 	for i := len(line) - 1; i >= 0; i-- {
// 		revLine += string(line[i])
// 	}
// 	// fmt.Println(line, revLine)
// 	return revLine
// }

func main() {
	var dataSlice []string
	searchString := "XMAS"
	//revString := "SAMX"
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
	fmt.Println(dataSlice)
	for i, val := range dataSlice {
		fmt.Printf("[-] Line number %d\n", i)
		for j := 0; j < len(val)-len(searchString); j++ {
			fmt.Printf("Char %d: %s\n", j, string(val[j]))

			// answer += hSearch(dataSlice[i])
			// revLine := reverse(dataSlice[i])
			// answer += hSearch(revLine)
		}

	}
	fmt.Println(dataSlice)
	fmt.Println(answer)
}
