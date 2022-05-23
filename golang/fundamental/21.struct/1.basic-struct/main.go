package main

import "fmt"

type Person struct {
	FirstName, MidleName, LastName string
	Age                            int
}

func main() {
	person1 := &Person{
		FirstName: "Ary",
		MidleName: "Setya",
		LastName:  "Pambudi",
		Age:       19,
	}

	var person2 Person

	person2.FirstName = "Malik"
	person2.LastName = "Ramadan"
	person2.Age = 13

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.FirstName)
	fmt.Println(person1.Age)

	fmt.Println(person2.FirstName)
	fmt.Println(person2.Age)
}
