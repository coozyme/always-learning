package main

import "fmt"

/*
	defer direkomendasikan ditempatkan diawal karena, jika ditempatkan diakhir ketika ditengah proses ada error maka defer tidak akan berjalan
*/
func logging() {
	fmt.Println("Call Logging")
}

func runApp(v int) {
	defer logging()

	fmt.Println("Run App")
	result := 10 / v
	fmt.Println("result", result)

}
func main() {
	runApp(10)
}
