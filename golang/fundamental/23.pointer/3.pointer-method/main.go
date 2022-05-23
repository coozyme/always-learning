package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) isMarried() {
	man.Name = "Mr. " + man.Name
}

func main() {
	lukman := Man{"Lukman"}
	lukman.isMarried()
	fmt.Println(lukman.Name)
}
