package main

import "fmt"

func main() {
	var name1 = "Ary"
	var name2 = "Setya"
	var name3 = "ary"

	fmt.Println(name1 > name2)
	fmt.Println(name1 < name2)
	fmt.Println(name1 != name2)
	fmt.Println(name1 == name3)
	fmt.Println(name1 <= name2)
	fmt.Println(name1 >= name2)
}
