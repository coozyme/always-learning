package main

import "fmt"

/*
	Perhatikan saat membuat slice karena berbeda dengan array
	diArray tidak bisa mengubah ukuran datanya
	Example:
		var fruitsA = []string{"apple", "grape"}      // slice
		var fruitsB = [2]string{"banana", "melon"}    // array
		var fruitsC = [...]string{"papaya", "grape"}  // array
*/
func main() {
	var names [4]string
	var value = [...]int{
		10,
		30,
		25,
	}
	var unlimitedValueArr []int

	names[0] = "Ary"
	names[1] = "Setya"
	names[2] = "Pambudi"

	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names))

	fmt.Println(value)

	append1 := append(unlimitedValueArr, 100)
	append2 := append(unlimitedValueArr, 200)
	append3 := append(unlimitedValueArr, 300)

	fmt.Println(append1)
	fmt.Println(append2)
	fmt.Println(append3)
	fmt.Println(unlimitedValueArr) // variable unlimitedValueArr akan tetap [] karena diarray datanya tidak bisa diubah
}
