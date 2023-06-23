package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	// Hint: Type return (a+b) below
	if (a < 1) || (b > 1000) {
		return 0
	}

	result := a + b
	return result
}

func main() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}