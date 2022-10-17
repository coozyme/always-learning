package main

import (
	"always-learning/golang/programme/design-patern/creational/factory-method/arms"
	"fmt"
)

func main() {
	ak47, _ := arms.GunFactory("ak47")
	musket, _ := arms.GunFactory("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g arms.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
