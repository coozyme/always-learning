package main

import "fmt"

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

var count int = 0

func CountDigits(num int) int {
	if num > 0 {
		CountDigits(num / 10)
		count++
		return count
	}
	return count
}

func main() {
	fmt.Println(factorialRecursive(5))

	hitungDigit := CountDigits(5123)

	fmt.Printf("Count of digits is: %d\n", hitungDigit)

}
