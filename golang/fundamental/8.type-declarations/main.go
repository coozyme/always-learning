package main

import "fmt"

func main() {
	/*
		Mendeklarasikan dengan tipe data alias
	*/

	type NoKTP string
	type Married bool

	var noKtpAry NoKTP = "88764646545646"
	var isMarried Married = false

	fmt.Printf("%v bertipe -> %T \n", noKtpAry, noKtpAry)
	fmt.Printf("%v bertipe -> %T \n", isMarried, isMarried)
}
