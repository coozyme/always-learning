package main

import "fmt"

func main() {
	/*
		Map adalah tipe data kumpulan key-value
		deklarasi panjangnya seperti ini -> var person map[string]string = map[string]string

		*Function Map
		Operasi
		len(map)	-> Untuk mendapatkan jumlah data di map
		map[key] -> Mengambil data di map dengan key
		map[key] = value -> Mengganti data di map dengan value
		make(map[TypeKey]TypeaValue)	-> Membuat map baru
		delete(map, key) -> Menghapus data di map dengan key
	*/
	person := map[string]string{
		"name":    "Ary",
		"address": "Tangerang",
	}

	fmt.Println(person)            //output: map[address:Tangerang name:Ary]
	fmt.Println(person["name"])    //output: Ary
	fmt.Println(person["address"]) //output: Tangerang

	/* Menambahkan data pada mat */
	person["title"] = "Programmer"

	fmt.Println(person) //output: map[address:Tangerang name:Ary title:Programmer]
	fmt.Println("")

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Ary"
	book["upss"] = "Salah"

	fmt.Println(book, len(book)) //output: map[author:Ary title:Buku Go-Lang upss:Salah]
	delete(book, "upss")
	fmt.Println(book, len(book)) //output: map[author:Ary title:Buku Go-Lang]

	delete(book, "upssyy")       //walaupun key-nya tidak ada tetap tidak error
	fmt.Println(book, len(book)) //outputnya: map[author:Ary title:Buku Go-Lang
}
