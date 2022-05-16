package main

import "fmt"

type dataDummy struct {
	name  string
	kelas int
}

func main() {
	dataku := []*dataDummy{
		{name: "ary", kelas: 10},
		{name: "setya", kelas: 13},
		{name: "pambudi", kelas: 14},
	}

	for index, value := range dataku {
		fmt.Printf("index ke %d name %s kelas %d \n", index, value.name, value.kelas)
	}
}
