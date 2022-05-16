package main

import "fmt"

func main() {
	f, m, l := getFullName("ary", "setya", "pambudi")
	fmt.Println(f, m, l)
}

func getFullName(first, midle, last string) (firstName, midleName, lastName string) {
	firstName = first
	midleName = midle
	lastName = last
	/*
		kita tidak perlu membuat
		return firstName, midleName, lastName
		karena kita sudah deklarasikan variable return valuenya
	*/
	return
}
