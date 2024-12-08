package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var list1 []string
	var list2 []string
	var ans int

	if len(os.Args) < 1 {
		fmt.Println("Need an Input file")
		os.Exit(1)
	}
	src, _ := os.Open(os.Args[1])
	data := bufio.NewScanner(src)
	for data.Scan() {
		line := data.Text()
		dp := strings.Split(line, "   ")
		list1 = append(list1, dp[0])
		list2 = append(list2, dp[1])
	}
	count := make(map[string]int)
	for i := 0; i < len(list1); i++ {
		//c := 0
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				count[list1[i]]++
			}
		}

	}
	//fmt.Println(count)
	for key, val := range count {
		x, _ := strconv.Atoi(key)
		x *= val
		ans += x
	}
	fmt.Printf("Answer is %d", ans)
}
