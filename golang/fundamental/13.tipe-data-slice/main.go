package main

import "fmt"

/*
	Perhatikan saat membuat slice karena berbeda dengan array
	Example:
		var fruitsA = []string{"apple", "grape"}      // slice
		var fruitsB = [2]string{"banana", "melon"}    // array
		var fruitsC = [...]string{"papaya", "grape"}  // array
*/
func main() {
	// data Array
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	fmt.Println(months)

	/*
		[4:7]
		pointer = 4 karena kita ambil dari index ke 4
		length = 3 karena kita ambil 3 value dari pointer yaitu index ke 4, nilai yang diambil itu ada dari index ke 4, 5, 6
	*/
	var slice1 = months[4:7]
	fmt.Println(slice1)      //output: [Mei Juni Juli]
	fmt.Println(len(slice1)) //output: 3
	fmt.Println(cap(slice1)) //output: 8

	slice1[0] = "Ubah-Mei" //Jika kita ubah nilai slice dari index tersebut, maka index itu mereference/mengarah pada index slice
	fmt.Println(slice1)    //output: [Ubah-Mei Juni Juli]
	fmt.Println(months)    //output: [Januari Februari Maret April Ubah-Mei Juni Juli Agustus September Oktober November Desember]

	fmt.Println("")

	slice2 := months[2:4]
	fmt.Println(slice2) //output; [Maret April]

	slice3 := append(slice2, "Ary")

	fmt.Println(slice3) //output: [Maret April Ary]
	slice3[1] = "Bukan-Desember"
	fmt.Println(slice3) //output: [Maret Bukan-Desember Ary]

	fmt.Println(slice2) //output: [Maret Bukan-Desember]
	fmt.Println(months) //output: [Januari Februari Maret Bukan-Desember Ary Juni Juli Agustus September Oktober November Desember]
	fmt.Println("")

	/*
		Jika kita mengappend slice diluar capacity slice maka itu akan membuat slice baru
	*/
	slice4 := months[10:]
	fmt.Println(slice4)

	append2 := append(slice4, "Addon-month")
	fmt.Println(append2)
	fmt.Println(months)
	fmt.Println("")

	/*
		Create new slice
	*/

	newSlice := make([]string, 2, 5) //len = 2, cap = 5

	newSlice[0] = "Ary"
	newSlice[1] = "Setya"

	fmt.Println(newSlice)      //output: [Ary Setya]
	fmt.Println(cap(newSlice)) //output: 5
	fmt.Println(len(newSlice)) //output: 2

	appendNewSlice := append(newSlice, "Pambudi")
	fmt.Println(appendNewSlice)          //output: [Ary Setya Pambudi]
	fmt.Println(newSlice, len(newSlice)) //output: Ary Setya] 2
	fmt.Println("")

	/*
		Copy value slice
		*Perhatikan len dan cap jika ingin mencopy data, jika len dan cap
	*/

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copySlice2 := make([]string, 1, cap(newSlice))

	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	copy(copySlice2, newSlice) //output: [Ary Setya]
	fmt.Println(copySlice2)    //output: [Ary]
}
