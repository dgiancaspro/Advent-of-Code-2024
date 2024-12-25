package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	enabled := true
	ans := 0
	if len(os.Args) != 2 {
		fmt.Println("[!] Input file needed")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("[!] %s", err)
		os.Exit(1)
	}
	src := bufio.NewScanner(f)
	for src.Scan() {
		matchRE := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)|don\'t\(\)|do\(\)`)
		reInt := regexp.MustCompile(`\d{1,3}`)
		line := src.Text()
		matched := matchRE.FindAllString(line, -1)
		fmt.Println(matched)
		fmt.Println(matched[102])
		nums := reInt.FindAllString(matched[102], -1)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		fmt.Println(num1, num2)
		for i := 0; i < len(matched); i++ {
			cmd := matched[i]
			fmt.Println(i, cmd, enabled)
			if cmd == "don't()" && enabled {
				enabled = false
			} else if cmd == "don't()" && !enabled {
				continue
			} else if cmd == "do()" && !enabled {
				enabled = true
			} else if cmd == "do()" && enabled {
				continue
			} else if enabled {
				nums := reInt.FindAllString(cmd, -1)
				fmt.Println(nums)
				num1, _ := strconv.Atoi(nums[0])
				num2, _ := strconv.Atoi(nums[1])
				ans += num1 * num2
			}
		}
	}
	fmt.Println(ans)
}
