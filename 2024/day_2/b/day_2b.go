package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dampner(l []string) bool {
	/*	This function removes elements f the record one at a time and sees if there is a valid
		Record */

	var safe bool
	for e := 0; e < len(l); e++ {
		var t []string
		for i := 0; i < len(l); i++ {
			if i == e {
				continue
			} else {
				t = append(t, l[i])
			}
		}
		var asc, desc bool
		int1, _ := strconv.Atoi(t[0])
		int2, _ := strconv.Atoi(t[1])
		if int1 < int2 {
			asc = true
		} else if int1 > int2 {
			desc = true
		} else {
			continue
		}

		for i := 0; i < len(t)-1; i++ {
			int1, _ := strconv.Atoi(t[i])
			int2, _ := strconv.Atoi(t[i+1])
			if asc && int1 < int2 && int2-int1 < 4 {
				safe = true
				continue
			} else if desc && int2 < int1 && int1-int2 < 4 {
				safe = true
				continue
			} else {
				safe = false
				break
			}
		}
		// If the readings in the current record are safve return true
		if safe {
			return true
		}

	}
	// If the record is unsafe regardless of wht reading we remove return false
	return false
}

func main() {
	var ans int
	if len(os.Args) < 2 {
		fmt.Println("[!] Input File needed")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("[!] %s", err)
	}
	src := bufio.NewScanner(f)
	for src.Scan() {
		l := strings.Split(src.Text(), " ")
		if dampner(l) {
			ans++
		}
	}
	fmt.Printf("Total Safe Readings is %d\n", ans)
}
