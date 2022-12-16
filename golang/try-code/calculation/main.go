package main

import (
	"fmt"
	"log"
	"math"

	"github.com/shopspring/decimal"
)

func Percent(x, y float64) float64 {
	// y = y / 100
	var p float64
	log.Println("X", x, y)
	log.Println("X", int(x), int(y))
	p = x * y
	f, _ := decimal.NewFromFloat(p).RoundCeil(-2).Float64()
	// itn, _ := strconv.ParseInt(d, 3, 32)
	// log.Println("LOG-DES", itn)
	// f, _ := strconv.ParseFloat(p, 64)
	log.Println("LOG-FF", p)
	// it, _ := fmt.Printf("%.3f \n", p)
	// log.Println("to FLoat", float64(it))
	log.Println("LOG-P", f)
	// fmt.Printf("%.3f \n", p)
	round := math.Ceil(f)
	log.Println("LOG-round", round)
	return round
}

func RoundToEven(x float64, y float64) float64 {
	z := x * y
	// it, _ := fmt.Printf("%.2f \n", z)
	// log.Println("it", z)
	a, _ := fmt.Printf("%.2f \n", z)
	log.Println("A", a)
	roundUp := math.Ceil(z)
	return roundUp
}

func Ins(value, insuranceFee float64) float64 {
	// fmt.Println("Normal", value*insurance)
	return math.Round(value * insuranceFee)
}

// func NewFromFloat(value float64) Decimal
func main() {
	var price float64 = 789000
	ins := 0.20 / 100
	// fmt.Println(Ins(price, 0.02))
	// fmt.Println(Percent(price, ins))
	fmt.Println(RoundToEven(price, ins))
}
