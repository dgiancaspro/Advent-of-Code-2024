package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dampner(l []string, e int) []string {
	// This function removes the Element at e and returns the
	// Slice in the same order
	var t []string
	for i := 0; i < len(l); i++ {
		if i == e {
			continue
		} else {
			t = append(t, l[i])
		}
	}
	return t
}

func safeTest(l []string) bool {
	var asc, desc bool
	// Determine which direction the data is going Ascending or Descending
	// By checking the first 2 values

	int1, _ := strconv.Atoi(l[0])
	int2, _ := strconv.Atoi(l[1])
	if int1 < int2 {
		asc = true
	} else if int1 > int2 {
		desc = true
		//if the first 2 vaues are equal remove the first value and check again

	} else {
		l = dampner(l, 0)
		int1, _ := strconv.Atoi(l[0])
		int2, _ := strconv.Atoi(l[1])
		if int1 < int2 {
			asc = true
		} else if int1 > int2 {
			desc = true
			// If the first 2 values are equal again return false
		} else {
			return false
		}
	}

	// Now we step through the readings and check the two rule
	// 1. No Value Changes by more than 3
	// 2. The values are going in the same direction i.e. ascending or descending
	// If the values are safe move to the next value

	for i := 0; i < len(l)-1; i++ {
		int1, _ := strconv.Atoi(l[i])
		int2, _ := strconv.Atoi(l[i+1])
		if asc && int1 < int2 && int2-int1 < 4 {
			continue
		} else if desc && int2 < int1 && int1-int2 < 4 {
			continue

			// If values are deemed unsafe use the dampner function to remoove the current value
			// Check the updated Slice against the same rules

		} else {
			newl := dampner(l, i)
			for j := 0; j < len(newl)-1; j++ {
				int1, _ := strconv.Atoi(newl[j])
				int2, _ := strconv.Atoi(newl[j+1])
				if asc && int1 < int2 && int2-int1 < 4 {
					continue
				} else if desc && int2 < int1 && int1-int2 < 4 {
					continue
					// If the new Slice fails again the list unsafe

				} else {
					return false
				}
			}
		}
	}
	// Finally if everything is correct return True

	return true
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
		if safeTest(l) {
			ans++
		}
	}
	fmt.Printf("Total Safe Readings is %d\n", ans)
}
