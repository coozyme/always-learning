package main

import (
	"always-learning/golang/programme/design-patern/creational/abstract-factory/sports"
	"fmt"
)

func main() {
	adidasFactory, _ := sports.GetSportsFactory("adidas")
	nikeFactory, _ := sports.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	fmt.Println("Nike")
	nikeShoe.SetLogo("Nike Shoes")
	fmt.Println("nikeShoe Logo: ", nikeShoe.GetLogo())
	fmt.Println("nikeShirtogo: ", nikeShirt.GetLogo())

	fmt.Println("Adidas")
	fmt.Println("adidasShoe Logo: ", adidasShoe.GetLogo())
	fmt.Println("adidasShirt Logo: ", adidasShirt.GetLogo())
	// printShoeDetails(nikeShoe)
	// printShirtDetails(nikeShirt)

	// printShoeDetails(adidasShoe)
	// printShirtDetails(adidasShirt)
}

// func printShoeDetails(s sports.IShoe) {
// 	fmt.Printf("Logo: %s", s.SetLogo("Nike"))
// 	fmt.Println()
// 	fmt.Printf("Size: %d", s.SetSize("IOKE"))
// 	fmt.Println()
// }

// func printShirtDetails(s sports.IShirt) {
// 	fmt.Printf("Logo: %s", s.SetLogo())
// 	fmt.Println()
// 	fmt.Printf("Size: %d", s.SetSize())
// 	fmt.Println()
// }
