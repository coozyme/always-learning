package main

import "fmt"

/*
	Gunakan type deklaration untuk terlihat simple
*/

type Filter func(string) string

func sayHello(name string, filter Filter) {
	resultFilter := filter(name)
	fmt.Println(resultFilter)
}

func filter(v string) string {
	if v == "anjing" {
		return "..."
	} else {
		return "Hello, " + v
	}
}

func main() {

	sayHello("anjing", filter)
	sayHello("ary", filter)
}
