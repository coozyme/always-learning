package main

import (
	"errors"
	"fmt"
)

func RemoveIndex(s []string, index int) []string {
	fmt.Println("e", s[:index])
	fmt.Println("ed", s[index+1:])
	fmt.Println("edo", s[index])
	return append(s[:index], s[index+1:]...)
}

func LoopingArr(data []string) {
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}

// Find Position Element
func Find(data []string, x string) (int, error) {
	for i, res := range data {
		if res == x {
			// fmt.Println("s", i)
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

// Mengecek apakah di dalam array terdapat nilai x.
func Contain(data []string, x string) bool {
	for _, v := range data {
		if v == x {
			return true
		}
	}
	return false
}

// Mencari nilai maksimal dari array
func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {
	dataArrStr := [5]string{"cat", "dog", "lion", "snake", "mouse"}
	dataArrInt := [5]int{5, 9, 10, 22, 11}

	// removeI := RemoveIndex(dataArrStr[:], 2)
	// fmt.Println("Remove Index", removeI)

	// LoopingArr(dataArrStr[:])
	// Find(dataArrStr[:], "lion")

	con := Contain(dataArrStr[:], "ook")
	fmt.Println(con)

	fMin, fMax := findMinAndMax(dataArrInt[:])
	fmt.Println("fMax", fMax)
	fmt.Println("fMin", fMin)
}
