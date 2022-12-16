package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person1 := Person{
		Name: "Ary",
		Age:  20,
	}
	// person2 := Person{
	// 	Name: "Setya",
	// 	Age:  12,
	// }
	// persons := []Person{person1, person2}

	fmt.Printf("%s", person1)
}
