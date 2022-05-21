package main

import "fmt"

func main() {
	name := "Ary"
	count := 0

	increment := func() {
		/*
			closure ini sama dengan scope si variable / function

			Jika name = "Setya" dicomment, maka output pada line 18 adalah pambudi
			karena valuenya diambil dari variable line 20

			dan jika name = "pambudi" yang dicoment, maka outputnya pada line 24 adalah setya
			karena valuenya diambil dari varible name pada line 6, yang valuenya adalah setya
		*/

		// name = "Setya"
		name := "pambudi"
		count++
		fmt.Println("increment")
		fmt.Println(name)
	}

	increment()

	fmt.Println(name)
}
