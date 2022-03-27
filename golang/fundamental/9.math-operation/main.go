package main

import (
	"fmt"
)

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a, b int = 10, 10

	c := a * b
	fmt.Println(c)

	/*
		Augmented Argument
	*/

	var i = 10

	i += 5
	fmt.Println(i)
	i -= 5
	fmt.Println(i)
	i *= 3
	fmt.Println(i)
	i /= 2
	fmt.Println(i)
	i %= 10
	fmt.Println(i)

	/*
		Unary Operator
	*/

	var bool = true

	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
	fmt.Println(-i)
	fmt.Println(!bool)

}
