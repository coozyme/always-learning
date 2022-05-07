package main

import (
	"fmt"
)

/*
	Disini kita akan belajar tentang subslice
	Subslice adalah sebuah slice yang dapat dibuat dari sebuah slice/array lainnya.

	baseContainer[low : high]       // two-index form
	baseContainer[low : high : max] // three-index form
*/

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := a[:]
	fmt.Println(s1)

	s2 := a[2:3]
	fmt.Println(s2)

	/*
		Jika uncomment s3 := a[:6:4] akan error atau merah karena nilai maxnya 6 , dan daya tampungnya 4
	*/
	// s3 := a[:6:4]
	s3 := a[:6:8]
	fmt.Println(s3) // output [0 1 2 3 4 5]

}
