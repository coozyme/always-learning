package main

import "log"

type personal struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

type Category struct {
	Class  string
	Person *personal `json:"person"`
}

func main() {
	// newP := personal{Age: 1, Name: "aoj"}
	var ds personal

	ds.Age = 1
	ds.Name = "asd"
	newCategory := Category{Class: "IPA", Person: ds}

	log.Println("LOG-", newCategory)
}
