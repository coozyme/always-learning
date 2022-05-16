package main

import "fmt"

func main() {
	fmt.Println(getName("gilang"))
	fmt.Println(getNumberPhone(8322308873))
}

func getName(fullname string) string {
	return "your name is" + fullname
}

func getNumberPhone(number int) int {
	return number
}
