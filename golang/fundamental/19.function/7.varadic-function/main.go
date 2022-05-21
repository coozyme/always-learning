package main

import "fmt"

func main() {
	fmt.Println(sumAll(10, 10, 10, 10))

	slice := []int{10, 10}

	fmt.Println(sumAll(slice...))
}

/*
	variable argument
*/
func sumAll(number ...int) int {
	total := 0
	for _, v := range number {
		total += v
	}

	return total
}
