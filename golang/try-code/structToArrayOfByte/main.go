package main

import (
	"encoding/json"
	"log"
)

type identity struct {
	Name string
	Age  int
}

func main() {
	newI := identity{
		Name: "Ary",
		Age:  12,
	}

	log.Println(newI)
	newByte, err := json.Marshal(newI)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("byte", newByte)

	// var data []byte
	i := identity{}
	err = json.Unmarshal(newByte, &i)

	if err != nil {
		log.Println("err", err)
	}

	log.Println(i)

}
