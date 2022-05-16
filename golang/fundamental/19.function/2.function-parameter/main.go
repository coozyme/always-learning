package main

import "fmt"

func main() {
	sayHelloTo("ary", "setya", 19)
}

/*
	Penulisan parameternya bisa di singkat menjadi
	(firstName, lastName string, age int0
*/
func sayHelloTo(firstName string, lastName string, age int) {
	fmt.Printf("hello %s %s %d ", firstName, lastName, age)
}
