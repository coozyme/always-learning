package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	name := "Ary"
	e := name[0]         // var e Jadinya bertipe byte atau uint8
	eString := string(e) // Byte diubah jadi string

	fmt.Printf("%v bertipe -> %T \n", name, name)
	fmt.Printf("%v bertipe -> %T \n", e, e)
	fmt.Printf("%v bertipe -> %T \n", eString, eString)

}
