package main

import "fmt"

func sayGoodBye(name string) string {
	return "Goodby " + name
}

func main() {
	goodBye := sayGoodBye

	fmt.Println(goodBye("ary"))
}
