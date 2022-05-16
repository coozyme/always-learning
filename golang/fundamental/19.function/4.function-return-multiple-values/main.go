package main

import "fmt"

func main() {
	firstName, lastName := getFullName("ary", "setya")
	fmt.Println(firstName, lastName)
}

func getFullName(firstName, lastName string) (string, string) {
	return firstName, lastName
}
