package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	enabled := true
	//ans := 0
	if len(os.Args) != 2 {
		fmt.Println("[!] Input file needed")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("[!] %s", err)
		os.Exit(1)
	}
	//m := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)
	dont := regexp.MustCompile(`don\'t\(\)`)
	do := regexp.MustCompile(`do\(\)`)
	src := bufio.NewScanner(f)

	for src.Scan() {
		line := src.Text()
		matchdont := dont.FindAllString(line, -1)
		if len(matchdont) != 0 && enabled {
			enabled = false
			fmt.Println(line)
		}
		matchdo := do.FindAllString(line, -1)
		if len(matchdo) != 0 && !enabled {
			enabled = true
		}

		// reInt := regexp.MustCompile(`\d{1,3}`)
		// for i := 0; i < len(matchSlice); i++ {
		// 	nums := reInt.FindAllString(matchSlice[i], -1)
		// 	num1, _ := strconv.Atoi(nums[0])
		// 	num2, _ := strconv.Atoi(nums[1])
		// 	ans += num1 * num2

	}

}

//fmt.Println(ans)
//}
