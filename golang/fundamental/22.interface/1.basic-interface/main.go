package main

import "fmt"

type HashName interface {
	GetName() string
}

type Animal struct {
	Name string
}

type Person struct {
	Name string
}

func sayHello(hasName HashName) {
	fmt.Println("Hello, ", hasName.GetName())
}

func (a Animal) GetName() string {
	return a.Name
}

func (p Person) GetName() string {
	return p.Name
}

func main() {
	ary := Person{
		Name: "Ary",
	}
	sayHello(ary)

	cat := Animal{
		Name: "cat",
	}

	sayHello(cat)
}
