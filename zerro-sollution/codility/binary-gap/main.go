package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Solution(9))
	fmt.Println(Solution(529))
	fmt.Println(Solution(32))
	fmt.Println(Solution(1041))
}

func Solution(n int64) int {

	splitBinary := strings.Split(strconv.FormatInt(n, 2), "")

	gaps := []string{}
	count := 0

	for i := 0; i < len(splitBinary); i++ {
		if splitBinary[i] == "1" {
			count++
			if count == 2 {
				break
			}
		} else if splitBinary[i] == "0" {
			gaps = append(gaps, splitBinary[i])
		}
	}

	if count < 2 {
		return 0
	}

	return len(gaps)
}
