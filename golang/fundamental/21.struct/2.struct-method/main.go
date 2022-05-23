package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (c Customer) sayHello(greeter string) {
	fmt.Printf("Hello %s saya dari %s \n", c.Name, greeter)
}
func main() {
	var asep Customer
	asep.Name = "Asep"
	asep.Address = "Tangerang"
	asep.Age = 22

	asep.sayHello("Customer Service")
}
