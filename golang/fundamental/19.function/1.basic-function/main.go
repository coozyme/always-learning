package main

import "fmt"

func main() {
	sayHello()
	sayHi()
	sayHello()
}

func sayHello() {
	fmt.Println("hello")
}

func sayHi() {
	fmt.Println("Hi")
}
