package main

import (
	"log"
	"os"
)

type data struct {
	Name string
	Age  int
}

type datas []data

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

	dataMap := make(map[string]datas)
	io := datas{}
	datasw := datas{data{Name: "AOK", Age: 123}, data{Name: "ARY", Age: 23}}
	datasw2 := datas{data{Name: "AOK2", Age: 123}, data{Name: "ARY2", Age: 23}}
	datasw3 := datas{data{Name: "AOK2", Age: 123}, data{Name: "ARY2", Age: 23}, data{Name: "ARY2", Age: 23}}
	dataMap["OK"] = datasw
	io = append(io, dataMap["OK"]...)
	log.Println("LOG-DT-MAP", dataMap)
	dataMap["OK"] = datasw2
	io = append(io, dataMap["OK"]...)
	log.Println("LOG-DT-MAP", dataMap)
	dataMap["OK"] = datasw3
	io = append(io, dataMap["OK"]...)
	log.Println("LOG-DT-MAP", dataMap)
	log.Println("io", io)
}
