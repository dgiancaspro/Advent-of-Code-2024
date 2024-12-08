package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ans := 0
	if len(os.Args) < 2 {
		fmt.Println("[!] Gonna need an input file there buddy")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("[!] Error: %s\n", err)
		os.Exit(1)
	}
	d := bufio.NewScanner(f)
	rec := 0
	for d.Scan() {
		var asc, desc bool
		unsafe := 0
		rec++
		l := strings.Split(d.Text(), " ")
		i1, _ := strconv.Atoi(l[0])
		i2, _ := strconv.Atoi(l[1])
		if i1 < i2 {
			asc = true
		} else if i1 > i2 {
			desc = true
		} else {
			unsafe++
		}
		for i := 0; i < len(l)-1; i++ {
			int1, _ := strconv.Atoi(l[i])
			int2, _ := strconv.Atoi(l[i+1])
			if int1 < int2 && int2-int1 < 4 && asc {
				continue
			} else if int1 > int2 && int1-int2 < 4 && desc {
				continue
			} else {
				unsafe++
			}
		}
		if unsafe == 0 {
			ans++
		}
	}
	fmt.Printf("[+] There are %d Safe Records", ans)
}
