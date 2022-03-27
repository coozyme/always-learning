package main

import "fmt"

func main() {
	var nilaiUjian = 60
	if nilaiUjian > 60 {
		fmt.Println("Anda Lulus")
	} else if nilaiUjian == 60 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Println("TIdak Lulus")
	}

	/* If dengan Short Statement */
	if absen := 6; absen <= 5 {
		fmt.Println("Tidak bisa ikut ujian")
	} else {
		fmt.Println("Bisa ikut ujian")
	}
}
