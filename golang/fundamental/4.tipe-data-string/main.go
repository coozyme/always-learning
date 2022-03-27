package main

import "fmt"

func main() {

	fmt.Println("") //string kosong

	fmt.Println("Ary")
	fmt.Println("Ary Setya")
	fmt.Println("Ary Setya Pambudi")

	fmt.Println(len("Ary"))             // output 3 , dihitung ddari angka 1
	fmt.Println("Ary Setya Pambudi"[0]) // output 65 , karena 65 itu bite-nya dari A
	fmt.Println("Ary Setya Pambudi"[1]) // output 114 , karena 114 itu bite-nya dari r
}
