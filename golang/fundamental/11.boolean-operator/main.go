package main

import "fmt"

func main() {
	ujian := 80
	absen := 70

	lulusUjian := ujian >= 75
	lulusAbsen := absen >= 75

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsen)

	lulus := lulusUjian && lulusAbsen

	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absen >= 75)

}
