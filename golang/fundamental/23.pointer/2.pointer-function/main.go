package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func CountryToIndonesia(a *Address) {
	a.Country = "Indonesia"
}

func main() {
	a := Address{"Jakarta", "DKI-Jakarta", ""}

	CountryToIndonesia(&a)
	fmt.Println(a)
}
