package main

import (
	"fmt"
)

func fizzBuzz(n int32) {
	// Write your code here
	for i := int32(1); i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}

}

func main() {
	fizzBuzz(16)
}
