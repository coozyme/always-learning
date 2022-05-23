package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Tangerang", "Banten", "Indonesia"}
	address2 := &address1

	fmt.Println(address1)

	address2.City = "Serang"

	fmt.Println(address1)
	fmt.Println(address2)
}
