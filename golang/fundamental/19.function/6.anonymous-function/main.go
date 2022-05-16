package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blaclist BlackList) {
	if blaclist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}
func main() {
	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blackList)
	registerUser("agung", func(s string) bool {
		return s == "agung"
	})
}
