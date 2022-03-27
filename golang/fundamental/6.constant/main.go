package main

import "fmt"

func main() {
	const firstName string = "Ary Setya"
	const lastName = "Pambudi"

	/*
		constant tidak akan dikomplain jika tidak dipakai
		constant tidak bisa dibuat dengan cara firstName := "Ary", karena dianggapnya variable
	*/

	/* deklarasi dan inisialiasi multiple konstant */
	const (
		first  string = "A"
		second        = "B"
		value         = 20
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(value)

}
