package main

import (
	"log"
	"os"
)

type data struct {
	Name string
	Age  int
}

func main() {
	// datas := []data{
	// 	data{"ARY", 2},
	// 	data{"ary", 3},
	// 	data{"Setya", 8},
	// }
	nama := os.Getenv("NAME")
	log.Println("LOG-SD", nama)
	// arr := strings.Split(nama, ",")
	// for _, v := range datas {

	// }
}
