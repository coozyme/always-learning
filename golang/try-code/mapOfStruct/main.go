package main

import "fmt"

type Movie struct {
	Title string
}

var movies = map[int]Movie{
	1: {
		"Spiderman",
	},
	2: {
		"Joker",
	},
	3: {
		"Escape Plan",
	},
	4: {
		"Lord of the Rings",
	},
}

type Data struct {
	movies Movie `json:"Data"`
}

func main() {
	idc, exit := movies[1]
	fmt.Println(idc, exit)

	responseData := make(map[string]string)
	responseData["Title"] = idc.Title

	Data.movies.Title[]
	fmt.Println(responseData["Title"])
}
