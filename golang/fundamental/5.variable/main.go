package main

import "fmt"

func main() {
	var name string

	name = "Ary Setya"
	fmt.Printf("%v -> tipe: %T \n", name, name)

	name = "Ary Setya Pambudi"
	fmt.Printf("%v -> tipe: %T \n", name, name)

	/*
		name = 12	//Akan error karena variable name bertipe string, tidak bisa diisi dengan tipe data lain.
		name = true	//Akan error karena variable name bertipe string, tidak bisa diisi dengan tipe data lain.
	*/

	var animal = "cat" // Jika buat variable dengan langsung menginisialisasinya, kita tidak perlu lagi unuk memberikan tipe datanya, karena golang sudah pintar
	fmt.Printf("%v -> tipe: %T \n", animal, animal)

	var age = 10 // by default nilai init bertipe int, integer berapanya tergantung SO laptop, ada yang 32 / 64
	fmt.Printf("%v -> tipe: %T \n", age, age)

	var year int8 = 10 // by default nilai init bertipe int, integer berapanya tergantung SO laptop, ada yang 32 / 64, jika kita ingin mengganti sesuai dengan tipe data yang kita inginkan maka harus kita define tipe datanya
	fmt.Printf("%v -> tipe: %T \n", year, year)

	/* best practice -> Tidak perlu menggunakan var dan tipe data saat membuat variable */
	country := "Indonesia"
	fmt.Println(country)

	/* best practice -> Deklarasi multiple variable */
	var address, city string

	address = "Tigaraksa"
	city = "Tangerang"

	fmt.Println(address)
	fmt.Println(city)

	/* best practice -> Deklarasi dan inisialisasi multiple variable */
	var (
		firstName = "Ary Setya"
		lastName  = "Pambudi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
