package main

import (
	"fmt"
	"strconv"
)

func main() {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke" + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}
	fmt.Println("SELESAI")
}
