package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var answer int
	var list1 []int
	var list2 []int

	if len(os.Args) == 1 {
		fmt.Println("Need an input file")
		os.Exit(0)
	}

	f, _ := os.Open(os.Args[1])
	in := bufio.NewScanner(f)
	for in.Scan() {
		line := in.Text()
		nums := strings.Split(line, "   ")
		val1, _ := strconv.Atoi(nums[0])
		val2, _ := strconv.Atoi(nums[1])
		list1 = append(list1, val1)
		list2 = append(list2, val2)

	}
	//fmt.Println(list1)
	//fmt.Println(list2)
	sort.Ints(list1)
	sort.Ints(list2)
	//fmt.Println(list1)
	//fmt.Println(list2)
	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			distance := list1[i] - list2[i]
			//fmt.Printf("%d - %d = %d", list1[i], list2[i], distance)
			answer += distance
		} else if list2[i] > list1[i] {
			distance := list2[i] - list1[i]
			//fmt.Printf("%d - %d = %d\n", list2[i], list1[i], distance)
			answer += distance
		} else {
			distance := list1[i] - list2[i]
			//fmt.Printf("%d - %d = %d", list1[i], list2[i], distance)
			answer += distance
		}
	}
	fmt.Printf("Answer = %d", answer)
}
